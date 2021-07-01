package fabric

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/BSNDA/fabric-sdk-go-gm/pkg/client/channel"
	"github.com/BSNDA/fabric-sdk-go-gm/pkg/client/event"
	"github.com/BSNDA/fabric-sdk-go-gm/pkg/client/ledger"
	fcCommon "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"github.com/BSNDA/fabric-sdk-go-gm/pkg/fabsdk"
	"github.com/FISCO-BCOS/go-sdk/core/types"

	"relayer/appchains/fabric/iservice"
	"relayer/common"
	"relayer/core"
	"relayer/logging"
	"relayer/mysql"
	"relayer/store"
)

// FabricChain defines the Fabric chain
type FabricChain struct {
	config  Config
	ChainID string // unique chain ID

	sdk           *fabsdk.FabricSDK
	channelClient *channel.Client
	eventClient   *event.Client
	ledgerClient  *ledger.Client

	store      *store.Store // store backend instance
	lastHeight int64        // last height

	done    bool                          // indicates if the chain monitor is done
	handler core.InterchainRequestHandler // handler for the interchain request
}

// NewFISCOChain constructs a new FISCOChain instance
func NewFabricChain(
	config Config,
	store *store.Store,
) (*FabricChain, error) {
	chainID := common.GetChainID("fabric", config.ChannelId, strconv.FormatUint(config.ChainId,10))
	fabric := &FabricChain{
		ChainID:   chainID,
		config:    config,
		done:      true,
		store:     store,
	}

	sdk, err := fabric.fabSdk()
	if err != nil {
		return nil, fmt.Errorf("fabric sdk init failed %s", err)

	}
	fabric.sdk = sdk
	channelProvider := fabric.sdk.ChannelContext(fabric.config.ChannelId,
		fabsdk.WithOrg(fabric.config.OrgName),
		fabsdk.WithUser(fabric.config.MspUserName),
	)

	client, err := channel.New(channelProvider)
	if err != nil {
		logging.Logger.Errorf("fabric channel client init failed %s", err)
		return nil, fmt.Errorf("fabric channel client init failed %s", err)
	}

	ec, err := event.New(channelProvider, event.WithBlockEvents()) //, event.WithSeekType(seek.Newest)
	if err != nil {
		logging.Logger.Errorf("fabric event client init failed %s", err)
		return nil, fmt.Errorf("fabric event client init failed %s", err)
	}

	lc, err := ledger.New(channelProvider)
	if err != nil {
		logging.Logger.Errorf("fabric ledger client init failed %s", err)
		return nil, fmt.Errorf("fabric ledger client init failed %s", err)
	}

	fabric.eventClient = ec
	fabric.channelClient = client
	fabric.ledgerClient = lc

	err = fabric.storeChainParams()
	if err != nil {
		return nil, err
	}

	err = fabric.storeChainID()
	if err != nil {
		return nil, err
	}

	store.GetInt64(HeightKey(chainID))

	return fabric, nil
}

// BuildFISCOChain builds a FISCOChain instance from the given chain params and store
func BuildFabricChain(
	chainParams []byte,
	store *store.Store,
) (*FabricChain, error) {
	var params ChainParams
	err := json.Unmarshal(chainParams, &params)
	if err != nil {
		return nil, err
	}

	baseCfgBz, err := store.Get(BaseConfigKey())
	if err != nil {
		return nil, err
	}

	var baseConfig BaseConfig
	err = json.Unmarshal(baseCfgBz, &baseConfig)
	if err != nil {
		return nil, err
	}

	fabricRelayer := params.ToStoreData(baseConfig.OrgCode)

	config := Config{
		BaseConfig:  baseConfig,
		FabricRelayer: fabricRelayer,
	}

	return NewFabricChain(config, store)
}

func (f *FabricChain) fabSdk() (*fabsdk.FabricSDK, error) {

	conf := f.config.GetSdkConfig(f.config.ChannelId, f.config.GetNodes())
	c, err := conf()
	if err != nil {
		ps, _ := c[0].Lookup("channels")
		logging.Logger.Infof("New Fabric SDK Channels is  %s", ps)
	}

	sdk, err := fabsdk.New(conf)

	if err != nil {
		logging.Logger.Errorf("New Fabric SDK has error %s", err.Error())
	}

	return sdk, err
}
// GetChainID implements AppChainI
func (f *FabricChain) GetChainID() string {
	return f.ChainID
}

// Start implements AppChainI
func (f *FabricChain) Start(handler core.InterchainRequestHandler) error {
	if !f.done {
		return fmt.Errorf("chain %s has been started", f.ChainID)
	}

	f.done = false
	f.handler = handler

	go f.monitor()

	logging.Logger.Infof("chain %s started", f.ChainID)

	return nil
}

// Stop implements AppChainI
func (f *FabricChain) Stop() error {
	logging.Logger.Infof("stopping chain %s", f.ChainID)
	f.done = true

	return nil
}

// GetHeight implements AppChainI
func (f *FabricChain) GetHeight() int64 {
	return f.lastHeight
}

// SendResponse implements AppChainI
func (f *FabricChain) SendResponse(requestID string, response core.ResponseI) error {
	requestIDBytes, err := hex.DecodeString(requestID)
	if err != nil {
		return err
	}

	var requestID32Bytes [32]byte
	copy(requestID32Bytes[:], requestIDBytes)

	tx, _, err := f.IServiceCoreSession.SetResponse(requestID32Bytes, response.GetErrMsg(), response.GetOutput())
	if err != nil {
		mysql.TxErrCollection(requestID, err.Error())
		return err
	}

	// TODO
	mysql.OnInterchainRequestResponseSent(requestID, tx.Hash().Hex())

	err = f.waitForReceipt(tx, "SetResponse")
	if err != nil {
		return err
	}

	// TODO
	mysql.OnInterchainRequestSucceeded(requestID)

	return nil
}

// buildInterchainRequest builds an interchain request from the interchain event
func (f *FabricChain) buildInterchainRequest(e *iservice.IServiceCoreExCrossChainRequestSent) core.InterchainRequest {
	var endpointInfo EndpointInfo
	err := json.Unmarshal([]byte(e.EndpointInfo), &endpointInfo)
	if err != nil {
		logging.Logger.Errorf("failed to decode endpointInfo: %s", err)
	}
	return core.InterchainRequest{
		ID:              hex.EncodeToString(e.RequestID[:]),
		SourceChainID:   f.ChainID,
		DestChainID:     endpointInfo.DestChainID,
		EndpointAddress: endpointInfo.EndpointAddress,
		EndpointType:    endpointInfo.EndpointType,
		Method:          e.Method,
		CallData:        e.CallData,
		Sender:          e.Sender.String(),
	}
}

// waitForReceipt waits for the receipt of the given tx
func (f *FabricChain) waitForReceipt(tx *types.Transaction, name string) error {
	logging.Logger.Infof("%s: transaction sent to %s, hash: %s", name, f.GetChainID(), tx.Hash().Hex())

	receipt, err := f.Client.WaitMined(tx)
	if err != nil {
		return fmt.Errorf("failed to mint the transaction %s: %s", tx.Hash().Hex(), err)
	}

	if receipt.Status != types.Success {
		return fmt.Errorf("transaction %s execution failed: %s", tx.Hash().Hex(), receipt.GetErrorMessage())
	}

	logging.Logger.Infof("%s: transaction %s execution succeeded", name, tx.Hash().Hex())

	return nil
}

// monitor is responsible for monitoring the chain
func (f *FabricChain) monitor() {
	for {
		f.scan()

		if f.done {
			return
		}

		time.Sleep(time.Duration(f.config.MonitorInterval) * time.Second)
	}
}

// scan performs chain scanning
func (f *FabricChain) scan() {
	fi := func(block *fcCommon.Block) bool {
		return true
	}

	logging.Logger.Infof("Into InterchainEventListener chainID：%s", f.config.ChainId)

	reg, eventch, err := f.eventClient.RegisterBlockEvent(fi) //.channelClient.RegisterChaincodeEvent(fc.ChainCodeID, "[\\S\\s]*")
	if err != nil {
		logging.Logger.Errorf("fabric event failed :%s", err)
		return
	}
	defer f.eventClient.Unregister(reg)
	for {
		select {
		case eventch, ok := <-eventch:
			if ok {
				f.blockevent(eventch)
			}
		case stop, ok := <-f.stop:
			{
				if ok && stop {
					logging.Logger.Infof("the chainId %s fabric relayer event is stop", fc.ChainInfo.GetChainId())
					return
				}

			}
		}

	}
	currentHeight, err := f.getBlockNumber()
	if err != nil {
		logging.Logger.Errorf("failed to get the current block height: %s", err)
		return
	}

	if f.lastHeight == 0 {
		f.lastHeight = currentHeight - 1
	}

	if currentHeight <= f.lastHeight {
		return
	}

	f.scanBlocks(f.lastHeight+1, currentHeight)
}

// scanBlocks scans the blocks of the specified range
func (f *FabricChain) scanBlocks(startHeight int64, endHeight int64) {
	for h := startHeight; h <= endHeight; {
		block, err := f.getBlock(h)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			continue
		}

		f.parseInterchainEventsFromBlock(block)

		err = f.updateHeight(h)
		if err != nil {
			logging.Logger.Errorf("failed to update height: %s", err)
		}

		h++
	}
}




// parseInterchainEventsFromBlock parses the interchain events from the block
func (f *FabricChain) parseInterchainEventsFromBlock(block CompactBlock) {
	for _, txHash := range block.Txs {
		receipt, err := f.Client.GetTransactionReceipt(context.Background(), txHash)
		if err != nil {
			logging.Logger.Errorf("failed to get the receipt, tx: %s, err: %s", txHash, err)
			continue
		}

		if receipt.Status != types.Success {
			continue
		}

		f.parseCrossChaiRequestSentEvents(receipt)
	}
}

// parseServiceInvokedEvents parses the ServiceInvoked events from the receipt
func (f *FabricChain) parseCrossChaiRequestSentEvents(receipt *types.Receipt) {
	for _, log := range receipt.Logs {
		if !strings.EqualFold(log.Address, f.Config.IServiceCoreAddr) {
			continue
		}

		data, err := hex.DecodeString(log.Data[2:])
		if err != nil {
			logging.Logger.Errorf("failed to decode the log data: %s", err)
			continue
		}

		var event iservice.IServiceCoreExCrossChainRequestSent
		err = f.IServiceCoreABI.Unpack(&event, "CrossChainRequestSent", data)
		if err != nil {
			logging.Logger.Errorf("failed to unpack the log data: %s", err)
			continue
		}

		request := f.buildInterchainRequest(&event)
		f.handler(f.ChainID, request, receipt.TransactionHash)
	}
}

// storeChainParams stores the chain params
func (f *FabricChain) storeChainParams() error {
	bz, err := json.Marshal(f.config.FabricRelayer)
	if err != nil {
		return err
	}

	return f.store.Set(ChainParamsKey(f.ChainID), bz)
}

func (f *FabricChain) storeChainID() error {
	chainIDsbz, err := f.store.Get([]byte("chainIDs"))
	if err != nil {
		return err
	}
	chainIDs := map[string]string{}
	err = json.Unmarshal(chainIDsbz, &chainIDs)
	if err != nil {
		return err
	}
	chainIDs[f.ChainID] = "fabric"
	bz, err := json.Marshal(chainIDs)
	return f.store.Set([]byte("chainIDs"), bz)
}

// updateHeight updates the height
func (f *FabricChain) updateHeight(height int64) error {
	f.lastHeight = height
	return f.store.SetInt64(HeightKey(f.ChainID), height)
}
