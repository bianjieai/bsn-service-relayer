package crossChainCode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

/*
服务消费者调用此接口发起服务调用，同时将触发一个跨链请求事件；
如果 Relayer 指定了非零服务费，则此接口负责托管服务消费者的服务费用

requestID 使用交易ID
*/
func (c *CrossChainCode) sendRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) == 0 {
		return shim.Error("the args cannot be empty")
	}
	if len(args[0]) == 0 {
		return shim.Error("the endpointInfo cannot be empty")
	}
	if len(args[1]) == 0 {
		return shim.Error("the method cannot be empty")
	}
	if len(args[2]) == 0 {
		return shim.Error("the callData cannot be empty")
	}

	serReq := &serviceRequest{}

	txId := stub.GetTxID()
	fmt.Println("txID:", txId)
	serReq.RequestId = txId
	serReq.EndpointInfo = args[0]
	serReq.Method = args[1]
	serReq.CallData = args[2]

	callBackInfo := &CallBackInfo{}
	callBackInfo.ChainCode = args[3]
	callBackInfo.FuncName = args[4]

	serReq.CallBack = callBackInfo

	callser := &serviceCallInfo{
		Request: serReq,
		Type:    Core_Type,
	}

	data, _ := json.Marshal(callser)

	if err := stub.PutState(coreKey(txId), data); err != nil {
		return shim.Error(fmt.Sprintf("put service info error；%s", err))
	}

	if err := stub.PutState("CallService", []byte(txId)); err != nil {
		return shim.Error(fmt.Sprintf("put CallService info error；%s", err))
	}

	stub.SetEvent(fmt.Sprintf("CallService_%s", txId), []byte(txId))
	return shim.Success([]byte(txId))
}

func (c *CrossChainCode) setResponse(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) == 0 {
		return shim.Error("the args cannot be empty")
	}

	resInfo := []byte(args[0])

	serReq := &serviceResponse{}
	err := json.Unmarshal(resInfo, serReq)
	if err != nil {
		return shim.Error("the args serialize failed")
	}

	key := coreKey(serReq.RequestId)
	ser, err := stub.GetState(key)
	if err != nil || len(ser) == 0 {
		return shim.Error("the requestID invalid")
	}

	serInfo := &serviceCallInfo{}
	err = json.Unmarshal(ser, serInfo)
	if err != nil {
		return shim.Error("the requestID invalid")
	}

	if serInfo.Response != nil{
		return shim.Error("duplicated response!")
	}

	serInfo.Response = serReq
	data, _ := json.Marshal(serInfo)
	if err := stub.PutState(key, data); err != nil {
		return shim.Error(fmt.Sprintf("put service info error；%s", err))
	}
	stub.SetEvent(fmt.Sprintf("SetResponse_%s", serReq.RequestId), []byte(serReq.RequestId))

	if serInfo.Request != nil && serInfo.Request.CallBack != nil {
		c.callBack(stub, serInfo.Request.CallBack, serReq)
	}
	return shim.Success(successMsg)
}

func (c *CrossChainCode) getResponse(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) == 0 {
		return shim.Error("the args cannot be empty")
	}

	key := coreKey(args[0])
	ser, err := stub.GetState(key)
	if err != nil || len(ser) == 0 {
		return shim.Error("the requestID invalid")
	}

	serInfo := &serviceCallInfo{}
	err = json.Unmarshal(ser, serInfo)
	if err != nil {
		return shim.Error("the requestID invalid")
	}

	if serInfo.Response == nil {
		return shim.Error("the response is null")
	}

	if serInfo.Response.ErrMsg != "" {
		return shim.Error(serInfo.Response.ErrMsg)
	}

	return shim.Success([]byte(serInfo.Response.Output))
}

func (c *CrossChainCode) callBack(stub shim.ChaincodeStubInterface, callbackInfo *CallBackInfo, res *serviceResponse) peer.Response {

	resB, _ := json.Marshal(res)

	var args [][]byte
	args = append(args, []byte(callbackInfo.FuncName))
	args = append(args, resB)

	stub.InvokeChaincode(callbackInfo.ChainCode, args, "")

	return shim.Success(successMsg)
}

func (c *CrossChainCode) query(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) == 0 {
		return shim.Error("the args cannot be empty")
	}
	key := coreKey(args[0])
	ser, err := stub.GetState(key)
	if err != nil || len(ser) == 0 {
		return shim.Error("the requestID invalid")
	}

	return shim.Success(ser)
}
