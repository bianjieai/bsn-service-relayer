package crossChainCode

import (
	"fmt"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

const (
	Core_Type   = "callserviceinfo"
)

var successMsg = []byte("success")
var err_NoFunc = shim.Error("function not found")

func coreKey(key string) string {
	return fmt.Sprintf("core_%s", key)
}

type CrossChainCode struct {
}

func (c *CrossChainCode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("chainCode Init")

	return shim.Success(successMsg)
}

func (c *CrossChainCode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	fmt.Println("chainCode Invoke")
	function, args := stub.GetFunctionAndParameters()

	if strings.ToLower(function) == "sendrequest" {
		return c.sendRequest(stub, args)
	}

	if strings.ToLower(function) == "setresponse" {
		return c.setResponse(stub, args)
	}

	if strings.ToLower(function) == "getresponse" {
		return c.getResponse(stub, args)
	}

	//if strings.ToLower(function) == "callback" {
	//	return c.callBack(stub, args)
	//}

	if strings.ToLower(function) == "query" {
		return c.query(stub, args)
	}


	return err_NoFunc
}
