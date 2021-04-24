package hub

import (
	"encoding/json"
	"errors"
	"fmt"
	servicesdk "github.com/irisnet/service-sdk-go"
	"github.com/irisnet/service-sdk-go/service"
	"github.com/irisnet/service-sdk-go/types"
	"github.com/irisnet/service-sdk-go/types/store"
	"strconv"

	"relayer/core"
	"relayer/logging"
	"relayer/mysql"
)

// IritaHubChain defines the Irita-Hub chain
type IritaHubChain struct {
	ChainID     string
	NodeRPCAddr string

	KeyPath    string
	KeyName    string
	Passphrase string

	ServiceClient servicesdk.ServiceClient
}

// NewIritaHubChain constructs a new Irita-Hub chain
func NewIritaHubChain(
	chainID string,
	nodeRPCAddr string,
	nodeGRPCAddr string,
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

	if len(nodeGRPCAddr) == 0 {
		nodeGRPCAddr = defaultNodeGRPCAddr
	}

	if len(keyPath) == 0 {
		keyPath = defaultKeyPath
	}

	fee, err := types.ParseDecCoins(defaultFee)
	if err != nil {
		panic(err)
	}

	config := types.ClientConfig{
		NodeURI:  nodeRPCAddr,
		GRPCAddr: nodeGRPCAddr,
		ChainID:  chainID,
		Gas:      defaultGas,
		Fee:      fee,
		Mode:     defaultBroadcastMode,
		Algo:     defaultKeyAlgorithm,
		KeyDAO:   store.NewFileDAO(keyPath),
		Level:    "debug",
	}

	hub := IritaHubChain{
		ChainID:       chainID,
		NodeRPCAddr:   nodeRPCAddr,
		KeyPath:       keyPath,
		KeyName:       keyName,
		Passphrase:    passphrase,
		ServiceClient: servicesdk.NewServiceClient(config),
	}

	return hub
}

// BuildIritaHubChain builds an Irita-Hub instance from the given config
func BuildIritaHubChain(config Config) IritaHubChain {
	return NewIritaHubChain(
		config.ChainID,
		config.NodeRPCAddr,
		config.NodeGRPCAddr,
		config.KeyPath,
		config.KeyName,
		config.Passphrase,
	)
}

// GetChainID implements IritaHubChainI
func (ic IritaHubChain) GetChainID() string {
	return ic.ChainID
}

// SendInterchainRequest implements IritaHubChainI
func (ic IritaHubChain) SendInterchainRequest(
	request core.InterchainRequest,
	cb core.ResponseCallback,
) error {
	invokeServiceReq, err := ic.BuildServiceInvocationRequest(request)
	if err != nil {
		return err
	}

	reqCtxID, resTx, err := ic.ServiceClient.InvokeService(invokeServiceReq, ic.BuildBaseTx())
	if err != nil {
		mysql.TxErrCollection(request.ID, err.Error())
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

	// TODO
	mysql.OnInterchainRequestSent(request.ID, requests[0].ID, resTx.Hash)

	logging.Logger.Infof("service request initiated on %s: %s", ic.ChainID, requests[0].ID)

	return ic.ResponseListener(reqCtxID, requests[0].ID, cb)
}

// BuildServiceInvocationRequest builds the service invocation request from the given interchain request
func (ic IritaHubChain) BuildServiceInvocationRequest(
	request core.InterchainRequest,
) (service.InvokeServiceRequest, error) {
	serviceFeeCap, err := types.ParseDecCoins(request.ServiceFeeCap)
	if err != nil {
		return service.InvokeServiceRequest{}, err
	}
	// TODO: request id, chain id and contract address will be added into the header on-chain when IBC enabled
	var input ServiceInput
	inputStr, err := strconv.Unquote("\"" + request.Input + "\"")
	err = json.Unmarshal([]byte(inputStr), &input)
	if err != nil {
		logging.Logger.Errorf("Unmarshal error %v", err)
		return service.InvokeServiceRequest{}, err
	}
	if input.Header == nil {
		return service.InvokeServiceRequest{}, errors.New("input header cannot be empty")
	}
	input.AddHeader("RequestID", request.ID)
	input.AddHeader("SourceChainID", request.ChainID)
	input.AddHeader("ContractAddress", request.ContractAddress)

	serviceInput, err := json.Marshal(input)
	if err != nil {
		return service.InvokeServiceRequest{}, err
	}

	return service.InvokeServiceRequest{
		ServiceName:   request.ServiceName,
		Providers:     []string{request.Provider},
		Input:         string(serviceInput),
		Timeout:       int64(request.Timeout),
		ServiceFeeCap: serviceFeeCap,
	}, nil
}

// ResponseListener gets and handles the response of the given request context ID by event subscription
func (ic IritaHubChain) ResponseListener(reqCtxID string, requestID string, cb core.ResponseCallback) error {
	response, err := ic.ServiceClient.QueryServiceResponse(requestID)
	if response.RequestContextID == reqCtxID {
		resp := core.ResponseAdaptor{
			StatusCode:  200,
			Result:      response.Result,
			Output:      response.Output,
			ICRequestID: requestID,
		}

		cb(requestID, resp)

		return nil
	}

	callbackWrapper := func(reqCtxID, requestID, result string, response string) {
		resp := core.ResponseAdaptor{
			StatusCode:  200,
			Result:      result,
			Output:      response,
			ICRequestID: requestID,
		}

		cb(requestID, resp)
	}

	logging.Logger.Infof("waiting for the service response on %s", ic.ChainID)

	_, err = ic.ServiceClient.SubscribeServiceResponse(reqCtxID, callbackWrapper)
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
