package crossChainCode

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func (c *CrossChainCode) addServiceBinding(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) == 0 {
		return shim.Error("the args cannot be empty")
	}
	sbi := &serviceBindInfo{}
	data := args[0]

	if err := json.Unmarshal([]byte(data), sbi); err != nil {
		return shim.Error("the args serialize failed")
	}

	si := &serviceInfo{
		Service: sbi,
		Type:    market_Type,
	}

	sib, err := json.Marshal(si)
	if err != nil {
		return shim.Error("")
	}
	if err := stub.PutState(marketKey(sbi.Name), sib); err != nil {
		return shim.Error(fmt.Sprintf("put service bind info error；%s", err))
	}

	return shim.Success(successMsg)
}

func (c *CrossChainCode) updateServiceBinding(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) == 0 {
		return shim.Error("the args cannot be empty")
	}

	usi := &updateServiceInfo{}
	err := json.Unmarshal([]byte(args[0]), usi)
	if err != nil {
		return shim.Error("the args serialize failed")
	}

	key := marketKey(usi.ServiceName)
	ser, err := stub.GetState(key)
	if err != nil || len(ser) == 0 {
		return shim.Error("the service invalid")
	}

	si := &serviceInfo{}
	err = json.Unmarshal(ser, si)
	if err != nil {
		return shim.Error("the service unmarshal failed")
	}

	if si.Service != nil {
		si.Service.Provider = usi.Provider
		si.Service.ServiceFee = usi.ServiceFee
		si.Service.Qos = usi.Qos
	} else {
		return shim.Error("the service bind info invalid")
	}

	sib, err := json.Marshal(si)
	if err != nil {
		return shim.Error("the service info marshal failed")
	}
	if err := stub.PutState(key, sib); err != nil {
		return shim.Error(fmt.Sprintf("put service bind info error；%s", err))
	}

	return shim.Success(successMsg)

}

func (c *CrossChainCode) getServiceBindings(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	queryString := `{
	"selector":{
		"type":"servicebindinfo"
		}
	}`

	res, err := stub.GetQueryResult(queryString)
	if err != nil {
		return shim.Error("QueryResult Error " + err.Error())
	}

	var serverList []*serviceBindInfo
	for {
		if res.HasNext() {
			kv, err := res.Next()
			if err != nil {
				continue
			}
			sbi := &serviceInfo{}
			data := kv.Value
			err = json.Unmarshal(data, sbi)
			if err != nil {
				continue
			}
			serverList = append(serverList, sbi.Service)

		} else {
			break
		}

	}
	list, err := json.Marshal(&serverList)
	if err != nil {
		return shim.Error("json Marshal Error " + err.Error())
	}
	return shim.Success(list)
}

func (c *CrossChainCode) getServiceBinding(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) == 0 {
		return shim.Error("the args cannot be empty")
	}

	key := marketKey(args[0])
	ser, err := stub.GetState(key)
	if err != nil || len(ser) == 0 {
		return shim.Error("the service invalid")
	}

	si := &serviceInfo{}
	err = json.Unmarshal(ser, si)
	if err != nil {
		return shim.Error("the service marshal failed")
	}

	sbib, err := json.Marshal(si.Service)
	if err != nil {
		return shim.Error("the service bind info marshal failed")
	}
	return shim.Success(sbib)

}
