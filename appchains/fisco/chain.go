package fisco

import (
	"context"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"
	"time"

	ethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/FISCO-BCOS/go-sdk/abi"
	fiscoclient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/core/types"

	"relayer/appchains/fisco/iservice"
	"relayer/common"
	"relayer/core"
	"relayer/indexer"
	"relayer/logging"
	"relayer/store"
)

// FISCOChain defines the FISCO chain
type FISCOChain struct {
	Config  Config
	Client  *fiscoclient.Client
	ChainID string // unique chain ID

	IServiceCoreSession   *iservice.IServiceCoreExSession   // iService Core Extension contract session
	IServiceMarketSession *iservice.IServiceMarketExSession // iService Market Extension contract session
	IServiceCoreABI       abi.ABI                           // parsed iService Core Extension ABI

	store      *store.Store // store backend instance
	lastHeight int64        // last height

	done    bool                          // indicates if the chain monitor is done
	handler core.InterchainRequestHandler // handler for the interchain request
}

// NewFISCOChain constructs a new FISCOChain instance
func NewFISCOChain(
	config Config,
	store *store.Store,
) (*FISCOChain, error) {
	clientConfig := BuildClientConfig(config)

	client, err := fiscoclient.Dial(clientConfig)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to fisco node: %s", err)
	}

	iServiceCoreABI, err := abi.JSON(strings.NewReader(iservice.IServiceCoreExABI))
	if err != nil {
		return nil, fmt.Errorf("failed to parse iService Core Extension ABI: %s", err)
	}

	iServiceCore, err := iservice.NewIServiceCoreEx(ethcmn.HexToAddress(config.IServiceCoreAddr), client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the iService Core Extension contract: %s", err)
	}

	iServiceMarket, err := iservice.NewIServiceMarketEx(ethcmn.HexToAddress(config.IServiceMarketAddr), client)
	if err != nil {
		return nil, fmt.Errorf("failed to instantiate the iService Market Extension contract: %s", err)
	}

	if config.MonitorInterval == 0 {
		config.MonitorInterval = DefaultMonitorInterval
	}

	chainID := GetChainID(config.ChainParams)

	fisco := &FISCOChain{
		Config:                config,
		Client:                client,
		ChainID:               chainID,
		IServiceCoreSession:   &iservice.IServiceCoreExSession{Contract: iServiceCore, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()},
		IServiceMarketSession: &iservice.IServiceMarketExSession{Contract: iServiceMarket, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()},
		IServiceCoreABI:       iServiceCoreABI,
		store:                 store,
		done:                  true,
	}

	err = fisco.storeChainParams()
	if err != nil {
		return nil, err
	}

	store.GetInt64(HeightKey(chainID))

	return fisco, nil
}

// BuildFISCOChain builds a FISCOChain instance from the given chain params and store
func BuildFISCOChain(
	chainParams []byte,
	store *store.Store,
) (*FISCOChain, error) {
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

	config := Config{
		BaseConfig:  baseConfig,
		ChainParams: params,
	}

	return NewFISCOChain(config, store)
}

// GetChainID implements AppChainI
func (f *FISCOChain) GetChainID() string {
	return f.ChainID
}

// Start implements AppChainI
func (f *FISCOChain) Start(handler core.InterchainRequestHandler) error {
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
func (f *FISCOChain) Stop() error {
	logging.Logger.Infof("stopping chain %s", f.ChainID)
	f.done = true

	return nil
}

// GetHeight implements AppChainI
func (f *FISCOChain) GetHeight() int64 {
	return f.lastHeight
}

// SendResponse implements AppChainI
func (f *FISCOChain) SendResponse(requestID string, response core.ResponseI) error {
	requestIDBytes, err := hex.DecodeString(requestID)
	if err != nil {
		return err
	}

	var requestID32Bytes [32]byte
	copy(requestID32Bytes[:], requestIDBytes)

	tx, _, err := f.IServiceCoreSession.SetResponse(requestID32Bytes, response.GetErrMsg(), response.GetOutput(), response.GetInterchainRequestID())
	if err != nil {
		return err
	}

	// TODO
	indexer.OnInterchainRequestResponseSent()

	err = f.waitForReceipt(tx, "SetResponse")
	if err != nil {
		return err
	}

	// TODO
	indexer.OnInterchainRequestSucceeded()

	return nil
}

// AddServiceBinding implements AppChainI
func (f *FISCOChain) AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error {
	tx, _, err := f.IServiceMarketSession.AddServiceBinding(serviceName, schemas, provider, serviceFee, big.NewInt(int64(qos)))
	if err != nil {
		return fmt.Errorf("failed to send AddServiceBinding transaction: %s", err)
	}

	return f.waitForReceipt(tx, "AddServiceBinding")
}

// UpdateServiceBinding implements AppChainI
func (f *FISCOChain) UpdateServiceBinding(serviceName, provider, serviceFee string, qos uint64) error {
	tx, _, err := f.IServiceMarketSession.UpdateServiceBinding(serviceName, provider, serviceFee, big.NewInt(int64(qos)))
	if err != nil {
		return fmt.Errorf("failed to send UpdateServiceBinding transaction: %s", err)
	}

	return f.waitForReceipt(tx, "UpdateServiceBinding")
}

// GetServiceBinding implements AppChainI
func (f *FISCOChain) GetServiceBinding(serviceName string) (core.ServiceBindingI, error) {
	serviceBinding, err := f.IServiceMarketSession.GetServiceBinding(serviceName)
	if err != nil {
		return nil, err
	}

	return iservice.ServiceBinding{
		ServiceName: serviceBinding.ServiceName,
		Schemas:     serviceBinding.Schemas,
		Provider:    serviceBinding.Provider,
		ServiceFee:  serviceBinding.ServiceFee,
		QoS:         serviceBinding.Qos,
	}, nil
}

// buildInterchainRequest builds an interchain request from the interchain event
func (f *FISCOChain) buildInterchainRequest(e *iservice.IServiceCoreExServiceInvoked) core.InterchainRequest {
	return core.InterchainRequest{
		ID:              hex.EncodeToString(e.RequestID[:]),
		ChainID:         f.ChainID,
		ContractAddress: e.ContractAddress.String(),
		ServiceName:     e.ServiceName,
		Provider:        e.Provider,
		Input:           e.Input,
		Timeout:         e.Timeout.Uint64(),
		ServiceFeeCap:   e.ServiceFeeCap,
	}
}

// waitForReceipt waits for the receipt of the given tx
func (f *FISCOChain) waitForReceipt(tx *types.Transaction, name string) error {
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
func (f *FISCOChain) monitor() {
	for {
		f.scan()

		if f.done {
			return
		}

		time.Sleep(time.Duration(f.Config.MonitorInterval) * time.Second)
	}
}

// scan performs chain scanning
func (f *FISCOChain) scan() {
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
func (f *FISCOChain) scanBlocks(startHeight int64, endHeight int64) {
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

// getBlockNumber retrieves the current block number
func (f *FISCOChain) getBlockNumber() (int64, error) {
	blockNumber, err := f.Client.GetBlockNumber(context.Background())
	if err != nil {
		return -1, err
	}

	blockNumberStr := string(blockNumber)

	return common.Hex2Decimal(blockNumberStr[3 : len(blockNumberStr)-1])
}

// getBlock gets the block in the given height
func (f *FISCOChain) getBlock(height int64) (block CompactBlock, err error) {
	blockBz, err := f.Client.GetBlockByNumber(context.Background(), fmt.Sprintf("0x%x", height), false)
	if err != nil {
		return block, fmt.Errorf("failed to retrieve the block, height: %d, err: %s", height, err)
	}

	err = json.Unmarshal(blockBz, &block)
	if err != nil {
		return block, fmt.Errorf("failed to unmarshal the block, height: %d, err: %s", height, err)
	}

	return
}

// parseInterchainEventsFromBlock parses the interchain events from the block
func (f *FISCOChain) parseInterchainEventsFromBlock(block CompactBlock) {
	for _, txHash := range block.Txs {
		receipt, err := f.Client.GetTransactionReceipt(context.Background(), txHash)
		if err != nil {
			logging.Logger.Errorf("failed to get the receipt, tx: %s, err: %s", txHash, err)
			continue
		}

		if receipt.Status != types.Success {
			continue
		}

		f.parseServiceInvokedEvents(receipt)
	}
}

// parseServiceInvokedEvents parses the ServiceInvoked events from the receipt
func (f *FISCOChain) parseServiceInvokedEvents(receipt *types.Receipt) {
	for _, log := range receipt.Logs {
		if log.Address != f.Config.IServiceCoreAddr {
			continue
		}

		data, err := hex.DecodeString(log.Data[2:])
		if err != nil {
			logging.Logger.Errorf("failed to decode the log data: %s", err)
			continue
		}

		var event iservice.IServiceCoreExServiceInvoked
		err = f.IServiceCoreABI.Unpack(&event, "ServiceInvoked", data)
		if err != nil {
			logging.Logger.Errorf("failed to unpack the log data: %s", err)
			continue
		}

		request := f.buildInterchainRequest(&event)
		f.handler(f.ChainID, request)
	}
}

// storeChainParams stores the chain params
func (f *FISCOChain) storeChainParams() error {
	bz, err := json.Marshal(f.Config.ChainParams)
	if err != nil {
		return err
	}

	return f.store.Set(ChainParamsKey(f.ChainID), bz)
}

// updateHeight updates the height
func (f *FISCOChain) updateHeight(height int64) error {
	f.lastHeight = height
	return f.store.SetInt64(HeightKey(f.ChainID), height)
}
