package ethereum

import (
	"context"
	"encoding/hex"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcmn "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	bcos "github.com/FISCO-BCOS/go-sdk/client"
	"github.com/FISCO-BCOS/go-sdk/conf"

	"relayer/appchains/ethereum/iservice"
	"relayer/core"
	"relayer/logging"
)

// BCOSChain defines the BCOS chain
type BCOSChain struct {
	ChainID    string
	ConfigFile string
	Client     bcos.Client

	IServiceCoreAddr   string // iService Core Extension contract address
	IServiceMarketAddr string // iService Market Extension contract address
	IServiceEventName  string // iService event name to be monitored
	IServiceEventSig   string // iService event signature to be monitored

	IServiceCoreABI   abi.ABI // iService Core Extension contract ABI
	IServiceMarketABI abi.ABI // iService Market Extension contract ABI

	IServiceCoreContract   *iservice.IServiceCoreEx   // iService Core Extension contract
	IServiceMarketContract *iservice.IServiceMarketEx // iService Market Extension contract
}

// NewBCOSChain constructs a new BCOSChain instance
func NewBCOSChain(
	configFile string,
	iServiceCoreAddr string,
	iServiceMarketAddr string,
	iServiceEventName string,
	iServiceEventSig string,
) BCOSChain {
	iServiceCoreABI, err := ParseABI(iservice.IServiceCoreExABI)
	if err != nil {
		logging.Logger.Panicf("failed to parse iservice core abi: %s", err)
	}

	iServiceMarketABI, err := ParseABI(iservice.IServiceMarketExABI)
	if err != nil {
		logging.Logger.Panicf("failed to parse iservice market abi: %s", err)
	}

	configs, err := &conf.ParseConfig([]byte(configFile))
	if err != nil {
		logging.Logger.Panicf("failed to parse bcos config file: %s", err)
	}

	client, err := bcos.Dial(configs[0])
	if err != nil {
		logging.Logger.Panicf("failed to connect to bcos node: %s", err)
	}

	iServiceCoreContract, err := iservice.NewIServiceCoreEx(ethcmn.HexToAddress(iServiceCoreAddr), client)
	if err != nil {
		logging.Logger.Panicf("failed to instantiate the iservice core contract: %s", err)
	}

	iServiceMarketContract, err := iservice.NewIServiceMarketEx(ethcmn.HexToAddress(iServiceMarketAddr), client)
	if err != nil {
		logging.Logger.Panicf("failed to instantiate the iservice market contract: %s", err)
	}

	bcos := BCOSChain{
		ChainID: "",
		Client:  client,

		IServiceCoreAddr:       iServiceCoreAddr,
		IServiceMarketAddr:     iServiceMarketAddr,
		IServiceEventName:      iServiceEventName,
		IServiceEventSig:       iServiceEventSig,
		IServiceCoreABI:        iServiceCoreABI,
		IServiceMarketABI:      iServiceMarketABI,
		IServiceCoreContract:   iServiceCoreContract,
		IServiceMarketContract: iServiceMarketContract,
	}

	return bcos
}

// MakeBCOSChain builds a BCOS chain from the given config
func MakeBCOSChain(config Config) BCOSChain {
	return NewBCOSChain(
		config.ChainID,
		config.NodeRPCAddr,
		config.GasLimit,
		config.GasPrice,
		config.Key,
		config.Passphrase,
		config.IServiceCoreAddr,
		config.IServiceMarketAddr,
		config.IServiceEventName,
		config.IServiceEventSig,
	)
}

// GetChainID implements AppChainI
func (b BCOSChain) GetChainID() string {
	return b.ChainID
}

// InterchainEventListener implements AppChainI
func (b BCOSChain) InterchainEventListener(cb core.InterchainEventHandler) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filterQuery := ethereum.FilterQuery{
		Addresses: []ethcmn.Address{ethcmn.HexToAddress(b.IServiceCoreAddr)},
		Topics:    [][]ethcmn.Hash{{crypto.Keccak256Hash([]byte(b.IServiceEventSig))}},
	}

	ch := make(chan ethtypes.Log)

	sub, err := b.Client.SubscribeFilterLogs(ctx, filterQuery, ch)
	if err != nil {
		return err
	}

	logHandler := func(log ethtypes.Log) {
		iServiceRequestEvent, err := b.parseLog(log)

		if err != nil {
			logging.Logger.Errorf("failed to parse log %+v: %s", log, err)
		} else {
			cb(iServiceRequestEvent)
		}
	}

	b.logListener(sub, ch, logHandler)

	return nil
}

// SendResponse implements AppChainI
func (b BCOSChain) SendResponse(requestID string, response core.ResponseI) error {
	auth, err := b.buildAuthTransactor()
	if err != nil {
		return err
	}

	requestIDBytes, err := hex.DecodeString(requestID)
	if err != nil {
		return err
	}

	var requestID32Bytes [32]byte
	copy(requestID32Bytes[:], requestIDBytes)

	tx, err := b.IServiceCoreContract.SetResponse(auth, requestID32Bytes, response.GetErrMsg(), response.GetOutput(), response.GetInterchainRequestID())
	if err != nil {
		return err
	}

	err = b.waitForReceipt(tx, "SetResponse")
	if err != nil {
		return err
	}

	return nil
}

// AddServiceBinding implements AppChainI
func (b BCOSChain) AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error {
	auth, err := b.buildAuthTransactor()
	if err != nil {
		return err
	}

	tx, err := b.IServiceMarketContract.AddServiceBinding(auth, serviceName, schemas, provider, serviceFee, big.NewInt(int64(qos)))
	if err != nil {
		return fmt.Errorf("failed to send AddServiceBinding transaction: %s", err)
	}

	err = b.waitForReceipt(tx, "AddServiceBinding")
	if err != nil {
		return err
	}

	return nil
}

// UpdateServiceBinding implements AppChainI
func (b BCOSChain) UpdateServiceBinding(serviceName, provider, serviceFee string, qos uint64) error {
	auth, err := b.buildAuthTransactor()
	if err != nil {
		return err
	}

	tx, err := b.IServiceMarketContract.UpdateServiceBinding(auth, serviceName, provider, serviceFee, big.NewInt(int64(qos)))
	if err != nil {
		return fmt.Errorf("failed to send UpdateServiceBinding transaction: %s", err)
	}

	err = b.waitForReceipt(tx, "UpdateServiceBinding")
	if err != nil {
		return err
	}

	return nil
}

// GetServiceBinding implements AppChainI
func (b BCOSChain) GetServiceBinding(serviceName string) (core.IServiceBinding, error) {
	serviceBinding, err := b.IServiceMarketContract.GetServiceBinding(nil, serviceName)
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

// buildAuthTransactor builds an authenticated transactor
func (b BCOSChain) buildAuthTransactor() (*bind.TransactOpts, error) {
	privKey, err := crypto.HexToECDSA(b.Key)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privKey)

	nextNonce, err := b.Client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}

	auth.GasLimit = b.GasLimit
	auth.GasPrice = b.GasPrice
	auth.Nonce = big.NewInt(int64(nextNonce))

	return auth, nil
}

// logListener listens to the log sent by the given channel and handles it with the specified handler
func (b BCOSChain) logListener(sub ethereum.Subscription, logChan chan ethtypes.Log, handler func(log ethtypes.Log)) {
	for {
		select {
		case log := <-logChan:
			handler(log)
		case err := <-sub.Err():
			logging.Logger.Errorf("Error on log subscription: %s", err)
		}
	}
}

// parseLog parses the given log to IServiceRequestEvent
func (b BCOSChain) parseLog(log ethtypes.Log) (iservice.IServiceRequestEvent, error) {
	var iServiceRequestEvent iservice.IServiceRequestEvent

	err := b.IServiceCoreABI.Unpack(&iServiceRequestEvent, b.IServiceEventName, log.Data)
	if err != nil {
		return iServiceRequestEvent, err
	}

	return iServiceRequestEvent, nil
}

// waitForReceipt waits for the receipt of the given tx
func (b BCOSChain) waitForReceipt(tx *ethtypes.Transaction, name string) error {
	logging.Logger.Infof("%s: transaction sent to %s, hash: %s", name, b.GetChainID(), tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), b.Client, tx)
	if err != nil {
		return fmt.Errorf("%s: failed to mint the transaction %s: %s", name, tx.Hash().Hex(), err)
	}

	if receipt.Status == ethtypes.ReceiptStatusFailed {
		return fmt.Errorf("%s: transaction %s execution failed", name, tx.Hash().Hex())
	}

	return nil
}
