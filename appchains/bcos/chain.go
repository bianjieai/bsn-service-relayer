package bcos

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"

	ethcmn "github.com/ethereum/go-ethereum/common"

	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	bcosclient "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"
	"github.com/FISCO-BCOS/go-sdk/core/types"

	"relayer/appchains/bcos/iservice"
	"relayer/core"
	"relayer/logging"
)

const (
	ChainIDPrefix = "bcos-"
)

// BCOSChain defines the BCOS chain
type BCOSChain struct {
	ChainID string
	Client  *bcosclient.Client

	IServiceCoreAddr   string // iService Core Extension contract address
	IServiceMarketAddr string // iService Market Extension contract address

	IServiceCoreSession   *iservice.IServiceCoreExSession   // iService Core Extension contract session
	IServiceMarketSession *iservice.IServiceMarketExSession // iService Market Extension contract session
}

// NewBCOSChain constructs a new BCOSChain instance
func NewBCOSChain(
	configFile string,
	iServiceCoreAddr string,
	iServiceMarketAddr string,
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

	iServiceCore, err := iservice.NewIServiceCoreEx(ethcmn.HexToAddress(iServiceCoreAddr), client)
	if err != nil {
		logging.Logger.Panicf("failed to instantiate the iservice core contract: %s", err)
	}

	iServiceMarket, err := iservice.NewIServiceMarketEx(ethcmn.HexToAddress(iServiceMarketAddr), client)
	if err != nil {
		logging.Logger.Panicf("failed to instantiate the iservice market contract: %s", err)
	}

	bcos := BCOSChain{
		ChainID: ChainIDPrefix + chainID.String(),
		Client:  client,

		IServiceCoreAddr:   iServiceCoreAddr,
		IServiceMarketAddr: iServiceMarketAddr,

		IServiceCoreSession:   &iservice.IServiceCoreExSession{Contract: iServiceCore, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()},
		IServiceMarketSession: &iservice.IServiceMarketExSession{Contract: iServiceMarket, CallOpts: *client.GetCallOpts(), TransactOpts: *client.GetTransactOpts()},
	}

	return bcos
}

// MakeBCOSChain builds a BCOS chain from the given config
func MakeBCOSChain(config Config) BCOSChain {
	return NewBCOSChain(
		config.ConfigFile,
		config.IServiceCoreAddr,
		config.IServiceMarketAddr,
	)
}

// GetChainID implements AppChainI
func (b BCOSChain) GetChainID() string {
	return b.ChainID
}

// InterchainEventListener implements AppChainI
func (b BCOSChain) InterchainEventListener(cb core.InterchainEventHandler) error {
	options := bind.WatchOpts{}
	ch := make(chan *iservice.IServiceCoreExServiceInvoked)

	sub, err := b.IServiceCoreSession.Contract.WatchServiceInvoked(&options, ch, nil)
	if err != nil {
		return err
	}

	for {
		select {
		case e := <-ch:
			cb(b.buildInterchainEvent(e))
		case err := <-sub.Err():
			logging.Logger.Errorf("Error on event subscription: %s", err)
		}
	}
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
