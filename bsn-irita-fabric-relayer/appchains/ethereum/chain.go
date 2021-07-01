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
	"github.com/ethereum/go-ethereum/ethclient"

	"relayer/appchains/ethereum/iservice"
	"relayer/core"
	"relayer/logging"
)

// EthChain defines the Ethereum chain
type EthChain struct {
	ChainID     string
	NodeRPCAddr string
	Client      *ethclient.Client
	GasLimit    uint64
	GasPrice    *big.Int
	Key         string
	Passphrase  string

	IServiceCoreAddr   string // iService Core Extension contract address
	IServiceMarketAddr string // iService Market Extension contract address
	IServiceEventName  string // iService event name to be monitored
	IServiceEventSig   string // iService event signature to be monitored

	IServiceCoreABI   abi.ABI // iService Core Extension contract ABI
	IServiceMarketABI abi.ABI // iService Market Extension contract ABI

	IServiceCoreContract   *iservice.IServiceCoreEx   // iService Core Extension contract
	IServiceMarketContract *iservice.IServiceMarketEx // iService Market Extension contract
}

// NewEthChain constructs a new EthChain instance
func NewEthChain(
	chainID string,
	nodeRPCAddr string,
	gasLimit uint64,
	gasPrice uint64,
	key string,
	passphrase string,
	iServiceCoreAddr string,
	iServiceMarketAddr string,
	iServiceEventName string,
	iServiceEventSig string,
) EthChain {
	iServiceCoreABI, err := ParseABI(iservice.IServiceCoreExABI)
	if err != nil {
		logging.Logger.Panicf("failed to parse iservice core abi: %s", err)
	}

	iServiceMarketABI, err := ParseABI(iservice.IServiceMarketExABI)
	if err != nil {
		logging.Logger.Panicf("failed to parse iservice market abi: %s", err)
	}

	client, err := ethclient.Dial(nodeRPCAddr)
	if err != nil {
		logging.Logger.Panicf("failed to connect to %s: %s", nodeRPCAddr, err)
	}

	iServiceCoreContract, err := iservice.NewIServiceCoreEx(ethcmn.HexToAddress(iServiceCoreAddr), client)
	if err != nil {
		logging.Logger.Panicf("failed to instantiate the iservice core contract: %s", err)
	}

	iServiceMarketContract, err := iservice.NewIServiceMarketEx(ethcmn.HexToAddress(iServiceMarketAddr), client)
	if err != nil {
		logging.Logger.Panicf("failed to instantiate the iservice market contract: %s", err)
	}

	eth := EthChain{
		ChainID:                chainID,
		NodeRPCAddr:            nodeRPCAddr,
		Client:                 client,
		GasLimit:               gasLimit,
		GasPrice:               big.NewInt(int64(gasPrice)),
		Key:                    key,
		Passphrase:             passphrase,
		IServiceCoreAddr:       iServiceCoreAddr,
		IServiceMarketAddr:     iServiceMarketAddr,
		IServiceEventName:      iServiceEventName,
		IServiceEventSig:       iServiceEventSig,
		IServiceCoreABI:        iServiceCoreABI,
		IServiceMarketABI:      iServiceMarketABI,
		IServiceCoreContract:   iServiceCoreContract,
		IServiceMarketContract: iServiceMarketContract,
	}

	return eth
}

// MakeEthChain builds an Ethereum chain from the given config
func MakeEthChain(config Config) EthChain {
	return NewEthChain(
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
func (ec EthChain) GetChainID() string {
	return ec.ChainID
}

// InterchainEventListener implements AppChainI
func (ec EthChain) InterchainEventListener(cb core.InterchainRequestHandler) error {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	filterQuery := ethereum.FilterQuery{
		Addresses: []ethcmn.Address{ethcmn.HexToAddress(ec.IServiceCoreAddr)},
		Topics:    [][]ethcmn.Hash{{crypto.Keccak256Hash([]byte(ec.IServiceEventSig))}},
	}

	ch := make(chan ethtypes.Log)

	sub, err := ec.Client.SubscribeFilterLogs(ctx, filterQuery, ch)
	if err != nil {
		return err
	}

	logHandler := func(log ethtypes.Log) {
		interchainRequest, err := ec.parseLog(log)

		if err != nil {
			logging.Logger.Errorf("failed to parse log %+v: %s", log, err)
		} else {
			cb(ec.ChainID, interchainRequest)
		}
	}

	ec.logListener(sub, ch, logHandler)

	return nil
}

// SendResponse implements AppChainI
func (ec EthChain) SendResponse(requestID string, response core.ResponseI) error {
	auth, err := ec.buildAuthTransactor()
	if err != nil {
		return err
	}

	requestIDBytes, err := hex.DecodeString(requestID)
	if err != nil {
		return err
	}

	var requestID32Bytes [32]byte
	copy(requestID32Bytes[:], requestIDBytes)

	tx, err := ec.IServiceCoreContract.SetResponse(auth, requestID32Bytes, response.GetErrMsg(), response.GetOutput(), response.GetInterchainRequestID())
	if err != nil {
		return err
	}

	err = ec.waitForReceipt(tx, "SetResponse")
	if err != nil {
		return err
	}

	return nil
}

// AddServiceBinding implements AppChainI
func (ec EthChain) AddServiceBinding(serviceName, schemas, provider, serviceFee string, qos uint64) error {
	auth, err := ec.buildAuthTransactor()
	if err != nil {
		return err
	}

	tx, err := ec.IServiceMarketContract.AddServiceBinding(auth, serviceName, schemas, provider, serviceFee, big.NewInt(int64(qos)))
	if err != nil {
		return fmt.Errorf("failed to send AddServiceBinding transaction: %s", err)
	}

	err = ec.waitForReceipt(tx, "AddServiceBinding")
	if err != nil {
		return err
	}

	return nil
}

// UpdateServiceBinding implements AppChainI
func (ec EthChain) UpdateServiceBinding(serviceName, provider, serviceFee string, qos uint64) error {
	auth, err := ec.buildAuthTransactor()
	if err != nil {
		return err
	}

	tx, err := ec.IServiceMarketContract.UpdateServiceBinding(auth, serviceName, provider, serviceFee, big.NewInt(int64(qos)))
	if err != nil {
		return fmt.Errorf("failed to send UpdateServiceBinding transaction: %s", err)
	}

	err = ec.waitForReceipt(tx, "UpdateServiceBinding")
	if err != nil {
		return err
	}

	return nil
}

// GetServiceBinding implements AppChainI
func (ec EthChain) GetServiceBinding(serviceName string) (core.ServiceBindingI, error) {
	serviceBinding, err := ec.IServiceMarketContract.GetServiceBinding(nil, serviceName)
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
func (ec EthChain) buildAuthTransactor() (*bind.TransactOpts, error) {
	privKey, err := crypto.HexToECDSA(ec.Key)
	if err != nil {
		return nil, err
	}

	auth := bind.NewKeyedTransactor(privKey)

	nextNonce, err := ec.Client.PendingNonceAt(context.Background(), auth.From)
	if err != nil {
		return nil, err
	}

	auth.GasLimit = ec.GasLimit
	auth.GasPrice = ec.GasPrice
	auth.Nonce = big.NewInt(int64(nextNonce))

	return auth, nil
}

// logListener listens to the log sent by the given channel and handles it with the specified handler
func (ec EthChain) logListener(sub ethereum.Subscription, logChan chan ethtypes.Log, handler func(log ethtypes.Log)) {
	for {
		select {
		case log := <-logChan:
			handler(log)
		case err := <-sub.Err():
			logging.Logger.Errorf("Error on log subscription: %s", err)
		}
	}
}

// parseLog parses the given log to InterchainRequest
func (ec EthChain) parseLog(log ethtypes.Log) (core.InterchainRequest, error) {
	var interchainRequest core.InterchainRequest

	err := ec.IServiceCoreABI.Unpack(&interchainRequest, ec.IServiceEventName, log.Data)
	if err != nil {
		return interchainRequest, err
	}

	return interchainRequest, nil
}

// waitForReceipt waits for the receipt of the given tx
func (ec EthChain) waitForReceipt(tx *ethtypes.Transaction, name string) error {
	logging.Logger.Infof("%s: transaction sent to %s, hash: %s", name, ec.GetChainID(), tx.Hash().Hex())

	receipt, err := bind.WaitMined(context.Background(), ec.Client, tx)
	if err != nil {
		return fmt.Errorf("%s: failed to mint the transaction %s: %s", name, tx.Hash().Hex(), err)
	}

	if receipt.Status == ethtypes.ReceiptStatusFailed {
		return fmt.Errorf("%s: transaction %s execution failed", name, tx.Hash().Hex())
	}

	return nil
}
