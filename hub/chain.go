package hub

import (
	"fmt"

	service "github.com/irisnet/service-sdk-go"
	servicetypes "github.com/irisnet/service-sdk-go/service"
	"github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/types/store"

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

	ServiceClient service.ServiceClient
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

	config := types.ClientConfig{
		NodeURI: nodeRPCAddr,
		ChainID: chainID,
		Gas:     defaultGas,
		Mode:    defaultBroadcastMode,
		Algo:    defaultKeyAlgorithm,
		KeyDAO:  store.NewFileDAO(keyPath),
	}

	hub := IritaHubChain{
		ChainID:       chainID,
		NodeRPCAddr:   nodeRPCAddr,
		KeyPath:       keyPath,
		KeyName:       keyName,
		Passphrase:    passphrase,
		ServiceClient: service.NewServiceClient(config),
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
	invokeServiceReq, err := ic.BuildServiceInvocationRequest(request)
	if err != nil {
		return err
	}

	reqCtxID, err := ic.ServiceClient.InvokeService(invokeServiceReq, ic.BuildBaseTx())
	if err != nil {
		return err
	}

	logging.Logger.Infof("request context created on %s: %s", ic.ChainID, reqCtxID)

	requests, err := ic.ServiceClient.QueryRequestsByReqCtx(reqCtxID, 1)
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
) (servicetypes.InvokeServiceRequest, error) {
	serviceFeeCap, err := types.ParseDecCoins(request.GetServiceFeeCap())
	if err != nil {
		return servicetypes.InvokeServiceRequest{}, err
	}

	return servicetypes.InvokeServiceRequest{
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

	_, err := ic.ServiceClient.SubscribeServiceResponse(reqCtxID, callbackWrapper)
	if err != nil {
		return err
	}

	return nil
}

// BuildBaseTx builds a base tx
func (ic IritaHubChain) BuildBaseTx() types.BaseTx {
	return types.BaseTx{
		From:     ic.KeyName,
		Password: ic.Passphrase,
	}
}
