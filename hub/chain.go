package hub

import (
	"fmt"

	iritaclient "github.com/bianjieai/irita-sdk-go"
	iservice "github.com/bianjieai/irita-sdk-go/modules/service"
	iritasdk "github.com/bianjieai/irita-sdk-go/types"
	"github.com/bianjieai/irita-sdk-go/types/store"

	"relayer/core"
	"relayer/logging"
)

// IritaHubChain defines the Irita-Hub chain
type IritaHubChain struct {
	ChainID     string
	NodeRPCAddr string

	KeyPath    string
	KeyName    string
	Passphrase string

	Client iritaclient.IRITAClient
}

// NewIritaHubChain constructs a new Irita-Hub chain
func NewIritaHubChain(
	chainID string,
	nodeRPCAddr string,
	keyPath string,
	keyName string,
	passphrase string,
) IritaHubChain {
	if len(chainID) == 0 {
		chainID = defaultChainID
	}

	if len(nodeRPCAddr) == 0 {
		nodeRPCAddr = defaultNodeRPCAddr
	}

	if len(keyPath) == 0 {
		keyPath = defaultKeyPath
	}

	config := iritasdk.ClientConfig{
		NodeURI: nodeRPCAddr,
		ChainID: chainID,
		Gas:     defaultGas,
		Mode:    defaultBroadcastMode,
		Algo:    defaultKeyAlgorithm,
		KeyDAO:  store.NewFileDAO(keyPath),
	}

	hub := IritaHubChain{
		ChainID:     chainID,
		NodeRPCAddr: nodeRPCAddr,
		KeyPath:     keyPath,
		KeyName:     keyName,
		Passphrase:  passphrase,
		Client:      iritaclient.NewIRITAClient(config),
	}

	return hub
}

// MakeIritaHubChain builds an Irita-Hub from the given config
func MakeIritaHubChain(config Config) IritaHubChain {
	return NewIritaHubChain(config.ChainID, config.NodeRPCAddr, config.KeyPath, config.KeyName, config.Passphrase)
}

// GetChainID implements IritaHubChainI
func (ic IritaHubChain) GetChainID() string {
	return ic.ChainID
}

// SendInterchainRequest implements IritaHubChainI
func (ic IritaHubChain) SendInterchainRequest(
	request core.InterchainRequestI,
	cb core.ResponseCallback,
) error {
	if request.GetServiceName() == "oracle" {
		return ic.HandleOracleRequest(request, cb)
	}

	invokeServiceReq, err := ic.BuildServiceInvocationRequest(request)
	if err != nil {
		return err
	}

	reqCtxID, err := ic.Client.Service.InvokeService(invokeServiceReq, ic.BuildBaseTx())
	if err != nil {
		return err
	}

	logging.Logger.Infof("request context created on %s: %s", ic.ChainID, reqCtxID)

	requests, err := ic.Client.Service.QueryRequestsByReqCtx(reqCtxID, 1)
	if err != nil {
		return err
	}

	if len(requests) == 0 {
		return fmt.Errorf("no service request initiated on %s", ic.ChainID)
	}

	logging.Logger.Infof("service request initiated on %s: %s", ic.ChainID, requests[0].ID)

	return ic.ResponseListener(reqCtxID, cb)
}

// BuildServiceInvocationRequest builds the service invocation request from the given interchain request
func (ic IritaHubChain) BuildServiceInvocationRequest(
	request core.InterchainRequestI,
) (iservice.InvokeServiceRequest, error) {
	serviceFeeCap, err := iritasdk.ParseDecCoins(request.GetServiceFeeCap())
	if err != nil {
		return iservice.InvokeServiceRequest{}, err
	}

	return iservice.InvokeServiceRequest{
		ServiceName:   request.GetServiceName(),
		Providers:     []string{request.GetProvider()},
		Input:         request.GetInput(),
		Timeout:       int64(request.GetTimeout()),
		ServiceFeeCap: serviceFeeCap,
	}, nil
}

// ResponseListener gets and handles the response of the given request context ID by event subscription
func (ic IritaHubChain) ResponseListener(reqCtxID string, cb core.ResponseCallback) error {
	callbackWrapper := func(reqCtxID, requestID, response string) {
		resp := core.ResponseAdaptor{
			StatusCode:  200,
			Output:      response,
			ICRequestID: requestID,
		}

		cb(requestID, resp)
	}

	logging.Logger.Infof("waiting for the service response on %s", ic.ChainID)

	_, err := ic.Client.Service.SubscribeServiceResponse(reqCtxID, callbackWrapper)
	if err != nil {
		return err
	}

	return nil
}

// HandleOracleRequest handles the oracle request
func (ic IritaHubChain) HandleOracleRequest(
	request core.InterchainRequestI,
	cb core.ResponseCallback,
) error {
	value, err := ic.QueryOracle([]byte(request.GetInput()))
	if err != nil {
		return err
	}

	resp := core.ResponseAdaptor{
		StatusCode:  200,
		Output:      string(value),
		ICRequestID: "",
	}

	cb("", resp)

	return nil
}

// QueryOracle queries the oracle data by the given key
func (ic IritaHubChain) QueryOracle(key []byte) (value []byte, err error) {
	// TODO
	return []byte{}, nil
}

// BuildBaseTx builds a base tx
func (ic IritaHubChain) BuildBaseTx() iritasdk.BaseTx {
	return iritasdk.BaseTx{
		From:     ic.KeyName,
		Password: ic.Passphrase,
	}
}
