package fabric

import (
	"encoding/json"
	"fmt"
	eventfab "github.com/hyperledger/fabric-sdk-go/pkg/common/providers/fab"
	"github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	"relayer/appchains/fabric/config"
	"relayer/appchains/fabric/entity"
	"relayer/appchains/fabric/store"
	"relayer/errors"
	"relayer/logging"
	"time"

	"relayer/core"

	"github.com/hyperledger/fabric-sdk-go/pkg/client/channel"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/event"
	"github.com/hyperledger/fabric-sdk-go/pkg/client/ledger"
	"github.com/hyperledger/fabric-sdk-go/pkg/fabsdk"
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
)

// FabricChain defines the Fabric chain
type FabricChain struct {
	ChainInfo *entity.FabricRelayer
	config    *config.FabricConfig

	sdk           *fabsdk.FabricSDK
	channelClient *channel.Client
	eventClient   *event.Client
	ledgerClient  *ledger.Client

	handler core.InterchainRequestHandler
	isStop  bool
	stop    chan bool
}

// NewFabricChain constructs a new FabricChain instance
func NewFabricChain(chainInfo *entity.FabricRelayer, sdkConf *config.FabricConfig) (*FabricChain, error) {
	fabric := &FabricChain{
		ChainInfo: chainInfo,
		config:    sdkConf,
		stop:      make(chan bool),
		isStop:    true,
	}

	sdk, err := fabric.fabSdk()
	if err != nil {
		logging.Logger.Errorf("fabric sdk init failed %s", err)
		return nil, errors.New("fabric sdk init failed %s", err)

	}
	fabric.sdk = sdk
	channelProvider := fabric.sdk.ChannelContext(fabric.ChainInfo.ChannelId,
		fabsdk.WithOrg(fabric.config.OrgName),
		fabsdk.WithUser(fabric.config.MspUserName),
	)

	client, err := channel.New(channelProvider)
	if err != nil {
		logging.Logger.Errorf("fabric channel client init failed %s", err)
		return nil, errors.New("fabric channel client init failed %s", err)
	}

	ec, err := event.New(channelProvider, event.WithBlockEvents()) // event.WithSeekType(seek.Newest)
	if err != nil {
		logging.Logger.Errorf("fabric event client init failed %s", err)
		return nil, errors.New("fabric event client init failed %s", err)
	}

	lc, err := ledger.New(channelProvider)
	if err != nil {
		logging.Logger.Errorf("fabric ledger client init failed %s", err)
		return nil, errors.New("fabric ledger client init failed %s", err)
	}

	fabric.eventClient = ec
	fabric.channelClient = client
	fabric.ledgerClient = lc

	return fabric, nil
}

func (f *FabricChain) fabSdk() (*fabsdk.FabricSDK, error) {

	conf := f.config.GetSdkConfig(f.ChainInfo.ChannelId, f.ChainInfo.GetNodes())
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
func (fc *FabricChain) GetChainID() string {
	return fc.ChainInfo.GetChainId()
}

func (f *FabricChain) Start(handler core.InterchainRequestHandler) error {
	f.handler = handler

	chanRes := make(chan *errors.ChanError)

	go f.InterchainEventListener(chanRes)

	select {
	case err, _ := <-chanRes:
		{
			if err.HasError {
				return err.Err
			}
		}
	case <-time.After(60 * time.Second):
		{
			logging.Logger.Errorf("Service %s start time out", f.ChainInfo.GetChainId())
			return errors.New("Service %s start time out", f.ChainInfo.GetChainId())
		}
	}

	return nil
}

func (f *FabricChain) Stop() error {
	f.isStop = true
	f.stop <- true
	return nil
}

// InterchainEventListener implements AppChainI
func (fc *FabricChain) InterchainEventListener(chanErr chan *errors.ChanError) {

	fi := func(block *common.Block) bool {
		return true
	}

	logging.Logger.Infof("Into InterchainEventListener chainID：%s", fc.ChainInfo.GetChainId())

	reg, eventch, err := fc.eventClient.RegisterBlockEvent(fi) //.channelClient.RegisterChaincodeEvent(fc.ChainCodeID, "[\\S\\s]*")
	if err != nil {
		logging.Logger.Errorf("fabric event failed :%s", err)
		chanErr <- errors.NewChanError(errors.New(fmt.Sprintf("fabric event failed :%s", err)))
		return
	}
	defer fc.eventClient.Unregister(reg)

	chanErr <- errors.NewChanSuccess()
	for {
		select {
		case eventch, ok := <-eventch:
			if ok {
				fc.blockevent(eventch)
			}
		case stop, ok := <-fc.stop:
			{
				if ok && stop {
					logging.Logger.Infof("the chainId %s fabric relayer event is stop", fc.ChainInfo.GetChainId())
					return
				}

			}
		}

	}

}

func (fc *FabricChain) blockevent(event *eventfab.BlockEvent) {

	block, err := ParseBlock(event.Block)
	if err != nil || len(block.Transactions) <= 0 {
		logging.Logger.Errorf("ParseBlock has error is %v", err)
		return
	}
	logging.Logger.Infof("channelID:%s,blockNumber:%d;blockHash:%s,chainID:%s", fc.ChainInfo.ChannelId, block.BlockNumber, block.BlockHash, fc.ChainInfo.GetChainId())
	for _, trans := range block.Transactions {
		logging.Logger.Infof("TxId:%s,channelID:%s ;", trans.TxId, trans.ChannelId)
		for _, rw := range trans.NameSpaceSets {
			logging.Logger.Infof("chaincodeId:%s,", rw.NameSpace)
			if rw.NameSpace == fc.ChainInfo.CrossChainCode {
				for _, w := range rw.Writes {
					if w.Key == "CallService" {
						logging.Logger.Infof("chaincodeId:%s is CallService,value is %s", rw.NameSpace, w.Value)
						reqId := w.Value
						request, err := fc.getServiceInfo(reqId)
						var endpointInfo EndpointInfo
						err = json.Unmarshal([]byte(request.Request.EndpointInfo), &endpointInfo)
						if err != nil {
							logging.Logger.Errorf("failed to decode endpointInfo: %s", err)
						}
						if err == nil && request != nil && request.Response == nil {
							event := core.InterchainRequest{
								ID:              request.Request.RequestId,
								SourceChainID:   fc.GetChainID(),
								DestChainID:     endpointInfo.DestChainID,
								DestSubChainID:  endpointInfo.DestSubChainID,
								DestChainType:   endpointInfo.DestChainType,
								EndpointAddress: endpointInfo.EndpointAddress,
								EndpointType:    endpointInfo.DestChainType,
								Method:          request.Request.Method,
								CallData:        []byte(request.Request.CallData),
								TxHash:          trans.TxId,
								Sender:          trans.CreateName,
							}
							err = fc.handler(fc.GetChainID(), event, trans.TxId)
							if err != nil {

							}
						}
					}
				}

			}
		}

		//if trans.Events != nil {
		//	eventName := trans.Events.EventName
		//	logging.Logger.Infof("eventName:%s;chaincodeId:%s,", eventName, trans.Events.ChaincodeId)
		//	if strings.Contains(eventName, "CallService_") {
		//		reqId := strings.ReplaceAll(eventName, "CallService_", "")
		//		request, err := fc.getServiceInfo(reqId)
		//		if err == nil && request != nil {
		//
		//			event := iservice.IServiceRequestEvent{
		//				RequestID:   reqId,
		//				Provider:    request.Service.Provider,
		//				ServiceName: request.Request.ServiceName,
		//				Input:       request.Request.Input,
		//				Timeout:     request.Request.Timeout,
		//			}
		//			cb(event)
		//		}
		//	}
		//}

	}

}

func (fc *FabricChain) getServiceInfo(requestId string) (*serviceCallInfo, error) {
	request := channel.Request{
		ChaincodeID: fc.ChainInfo.CrossChainCode,
		Fcn:         "query",
		Args:        [][]byte{[]byte(requestId)},
	}

	fabres, err := fc.channelClient.Query(request)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("call fabric ServiceInfo failed :%s", err))
	}
	if fabres.TxValidationCode != pb.TxValidationCode_VALID {
		return nil, errors.New(fmt.Sprintf("call fabric ServiceInfo TxValidationCode is %s", fabres.TxValidationCode.String()))
	}
	res := &serviceCallInfo{}
	err = json.Unmarshal(fabres.Payload, res)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("call fabric ServiceInfo Unmarshal failed :%s", err))
	}
	return res, nil
}

func (fc *FabricChain) GetHeight() int64 {
	cfg, err := fc.ledgerClient.QueryInfo()

	if err != nil {
		return 0
	}

	height := cfg.BCI.Height

	return int64(height)
}

// SendResponse implements AppChainI
func (fc *FabricChain) SendResponse(requestID string, response core.ResponseI) error {

	logging.Logger.Infof("SendResponse.requestID: %s", requestID)
	logging.Logger.Infof("SendResponse.InterchainRequestID: %s", response.GetInterchainRequestID())
	logging.Logger.Infof("SendResponse.GetOutput: %s", response.GetOutput())
	logging.Logger.Infof("SendResponse.GetErrMsg: %s", response.GetErrMsg())

	res := &serviceResponse{
		RequestId:   requestID,
		Output:      response.GetOutput(),
		IcRequestId: response.GetInterchainRequestID(),
	}

	resb, _ := json.Marshal(res)
	request := channel.Request{
		ChaincodeID: fc.ChainInfo.CrossChainCode,
		Fcn:         "setResponse",
		Args:        [][]byte{resb},
	}

	fabres, err := fc.channelClient.Execute(request)
	if err != nil {
		return errors.New(fmt.Sprintf("call fabric setResponse failed :%s", err))
	}

	if fabres.TxValidationCode != pb.TxValidationCode_VALID {
		return errors.New(fmt.Sprintf("call fabric setResponse TxValidationCode is %s", fabres.TxValidationCode.String()))
	}

	// 返回的交易记录 update where 来源 = 0
	// request_id  requestID
	// ic_request_id  IcRequestId
	// from_res_tx fabres.TransactionID
	// tx_time
	// tx_status 根据output解析交易状态，if response.GetErrMsg() == "" 则 1 elif 则为 2 并且加 response.GetErrMsg() 错误信息

	if response.GetErrMsg() == "" {
		logging.Logger.Infoln("callback no err info")
		data := entity.FabricRelayerTx{
			Request_id:     requestID,
			Ic_request_id:  res.IcRequestId,
			From_res_tx:    string(fabres.TransactionID),
			Tx_status:      1,
			Tx_time:        time.Now(),
			Source_service: 0,
		}
		store.CallBackSendResponse(&data)

	} else {
		logging.Logger.Infof("callback has err msg :%s", response.GetErrMsg())
		data := entity.FabricRelayerTx{
			Request_id:     requestID,
			Ic_request_id:  res.IcRequestId,
			From_res_tx:    string(fabres.TransactionID),
			Tx_status:      2,
			Error:          response.GetErrMsg(),
			Tx_time:        time.Now(),
			Source_service: 0,
		}
		store.CallBackSendResponse(&data)
	}

	return nil
}
