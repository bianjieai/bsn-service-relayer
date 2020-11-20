package bcos

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
	bcosclient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"

	"relayer/appchains/bcos/iservice"
	"relayer/common"
	"relayer/core"
	"relayer/logging"
	"relayer/store"
)

// BCOSChain defines the BCOS chain
type BCOSChain struct {
	ChainID string
	Client  *bcosclient.Client

	IServiceCoreAddr      string                            // iService Core Extension contract address
	IServiceMarketAddr    string                            // iService Market Extension contract address
	IServiceCoreSession   *iservice.IServiceCoreExSession   // iService Core Extension contract session
	IServiceMarketSession *iservice.IServiceMarketExSession // iService Market Extension contract session
	IServiceCoreABI       abi.ABI                           // parsed iService Core Extension ABI

	monitorInterval time.Duration // monitoring interval
	lastHeight      int64         // last height
	store           *store.Store  // store instance

	eventHandler core.InterchainEventHandler // handler for the interchain event
}

// NewBCOSChain constructs a new BCOSChain instance
func NewBCOSChain(
	configFile string,
	iServiceCoreAddr string,
	iServiceMarketAddr string,
	monitorInterval uint64,
	store *store.Store,
) BCOSChain {
	configs, err := conf.ParseConfigFile(configFile)
	if err != nil {
		logging.Logger.Panicf("failed to parse bcos config file: %s", err)
	}

	client, err := bcosclient.Dial(&configs[0])
	if err != nil {
		logging.Logger.Panicf("failed to connect to bcos node: %s", err)
	}

	chainID, err := client.GetChainID(context.Background())
	if err != nil {
		logging.Logger.Panicf("failed to get the chain id: %s", err)
	}

	iServiceCoreABI, err := abi.JSON(strings.NewReader(iservice.IServiceCoreExABI))
	if err != nil {
		logging.Logger.Panic("failed to parse iService Core Extension ABI: %s", err)
	}

	iServiceCore, err := iservice.NewIServiceCoreEx(ethcmn.HexToAddress(iServiceCoreAddr), client)
	if err != nil {
		logging.Logger.Panicf("failed to instantiate the iservice core contract: %s", err)
	}

	iServiceMarket, err := iservice.NewIServiceMarketEx(ethcmn.HexToAddress(iServiceMarketAddr), client)
	if err != nil {
		logging.Logger.Panicf("failed to instantiate the iservice market contract: %s", err)
	}

	if monitorInterval == 0 {
		monitorInterval = DefaultMonitorInterval
	}

	bcos := BCOSChain{
		ChainID: ChainIDPrefix + chainID.String(),
		Client:  client,

		IServiceCoreAddr:      iServiceCoreAddr,
		IServiceMarketAddr:    iServiceMarketAddr,
		IServiceCoreSession:   &iservice.IServiceCoreExSession{Contract: iServiceCore, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()},
		IServiceMarketSession: &iservice.IServiceMarketExSession{Contract: iServiceMarket, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()},
		IServiceCoreABI:       iServiceCoreABI,

		monitorInterval: time.Duration(monitorInterval) * time.Second,
		store:           store,
	}

	if store != nil {
		bcos.lastHeight, _ = store.GetInt64(GetHeightKey())
	}

	return bcos
}

// MakeBCOSChain builds a BCOS chain from the given config
func MakeBCOSChain(config Config, store *store.Store) BCOSChain {
	return NewBCOSChain(
		config.ConfigFile,
		config.IServiceCoreAddr,
		config.IServiceMarketAddr,
		config.MonitorInterval,
		store,
	)
}

// GetChainID implements AppChainI
func (b BCOSChain) GetChainID() string {
	return b.ChainID
}

// InterchainEventListener implements AppChainI
func (b BCOSChain) InterchainEventListener(cb core.InterchainEventHandler) error {
	// hold the callback
	b.eventHandler = cb

	// start the monitor
	go b.monitor()

	select {}
}

// SendResponse implements AppChainI
func (b BCOSChain) SendResponse(requestID string, response core.ResponseI) error {
	requestIDBytes, err := hex.DecodeString(requestID)
	if err != nil {
		return err
	}

	var requestID32Bytes [32]byte
	copy(requestID32Bytes[:], requestIDBytes)

	tx, _, err := b.IServiceCoreSession.SetResponse(requestID32Bytes, response.GetErrMsg(), response.GetOutput(), response.GetInterchainRequestID())
	if err != nil {
		return err
	}

	return b.waitForReceipt(tx, "SetResponse")
}

// AddServiceBinding implements AppChainI
func (b BCOSChain) AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error {
	tx, _, err := b.IServiceMarketSession.AddServiceBinding(serviceName, schemas, provider, serviceFee, big.NewInt(int64(qos)))
	if err != nil {
		return fmt.Errorf("failed to send AddServiceBinding transaction: %s", err)
	}

	return b.waitForReceipt(tx, "AddServiceBinding")
}

// UpdateServiceBinding implements AppChainI
func (b BCOSChain) UpdateServiceBinding(serviceName, provider, serviceFee string, qos uint64) error {
	tx, _, err := b.IServiceMarketSession.UpdateServiceBinding(serviceName, provider, serviceFee, big.NewInt(int64(qos)))
	if err != nil {
		return fmt.Errorf("failed to send UpdateServiceBinding transaction: %s", err)
	}

	return b.waitForReceipt(tx, "UpdateServiceBinding")
}

// GetServiceBinding implements AppChainI
func (b BCOSChain) GetServiceBinding(serviceName string) (core.IServiceBinding, error) {
	serviceBinding, err := b.IServiceMarketSession.GetServiceBinding(serviceName)
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

// buildInterchainEvent builds an interchain request event from the appchain-specific event
func (b BCOSChain) buildInterchainEvent(e *iservice.IServiceCoreExServiceInvoked) iservice.IServiceRequestEvent {
	return iservice.NewIServiceRequestEvent(
		e.RequestID,
		e.ServiceName,
		e.Input,
		e.Timeout,
	)
}

// waitForReceipt waits for the receipt of the given tx
func (b BCOSChain) waitForReceipt(tx *types.Transaction, name string) error {
	logging.Logger.Infof("%s: transaction sent to %s, hash: %s", name, b.GetChainID(), tx.Hash().Hex())

	receipt, err := b.Client.WaitMined(tx)
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
func (b *BCOSChain) monitor() {
	for {
		b.scan()
		time.Sleep(b.monitorInterval)
	}
}

// scan performs chain scanning
func (b *BCOSChain) scan() {
	currentHeight, err := b.getBlockNumber()
	if err != nil {
		logging.Logger.Errorf("failed to get the current block height: %s", err)
		return
	}

	if b.lastHeight == 0 {
		b.lastHeight = currentHeight - 1
	}

	if currentHeight <= b.lastHeight {
		return
	}

	b.scanBlocks(b.lastHeight+1, currentHeight)
}

// scanBlocks scans the blocks of the specified range
func (b *BCOSChain) scanBlocks(startHeight int64, endHeight int64) {
	for h := startHeight; h <= endHeight; {
		block, err := b.getBlock(h)
		if err != nil {
			logging.Logger.Errorf(err.Error())
			continue
		}

		b.parseInterchainEventsFromBlock(block)

		err = b.updateHeight(h)
		if err != nil {
			logging.Logger.Errorf("failed to update height in store: %s", err)
		}

		h++
	}
}

// getBlockNumber retrieves the current block number
func (b *BCOSChain) getBlockNumber() (int64, error) {
	blockNumber, err := b.Client.GetBlockNumber(context.Background())
	if err != nil {
		return -1, err
	}

	blockNumberStr := string(blockNumber)

	return common.Hex2Decimal(blockNumberStr[3 : len(blockNumberStr)-1])
}

// getBlock gets the block in the given height
func (b *BCOSChain) getBlock(height int64) (block CompactBlock, err error) {
	blockBz, err := b.Client.GetBlockByNumber(context.Background(), fmt.Sprintf("0x%x", height), false)
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
func (b *BCOSChain) parseInterchainEventsFromBlock(block CompactBlock) {
	for _, txHash := range block.Txs {
		receipt, err := b.Client.GetTransactionReceipt(context.Background(), txHash)
		if err != nil {
			logging.Logger.Errorf("failed to get the receipt, tx: %s, err: %s", txHash, err)
			continue
		}

		if receipt.Status != types.Success {
			continue
		}

		b.parseServiceInvokedEvents(receipt)
	}
}

// parseServiceInvokedEvents parses the ServiceInvoked events from the receipt
func (b *BCOSChain) parseServiceInvokedEvents(receipt *types.Receipt) {
	for _, log := range receipt.Logs {
		if log.Address != b.IServiceCoreAddr {
			continue
		}

		data, err := hex.DecodeString(log.Data[2:])
		if err != nil {
			logging.Logger.Errorf("failed to decode the log data: %s", err)
			continue
		}

		var event iservice.IServiceCoreExServiceInvoked
		err = b.IServiceCoreABI.Unpack(&event, "ServiceInvoked", data)
		if err != nil {
			logging.Logger.Errorf("failed to unpack the log data: %s", err)
			continue
		}

		icEvent := b.buildInterchainEvent(&event)
		b.eventHandler(icEvent)
	}
}

// updateHeight updates the height
func (b *BCOSChain) updateHeight(height int64) error {
	b.lastHeight = height
	return b.store.SetInt64(GetHeightKey(), height)
}
