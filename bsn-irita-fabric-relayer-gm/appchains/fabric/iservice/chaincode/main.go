package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	"relayer/appchains/fabric/iservice/chaincode/crossChainCode"
)

func main() {

	err := shim.Start(new(crossChainCode.CrossChainCode))
	if err != nil {
		fmt.Printf("Error starting CrossChainCode: %s", err)
	}

}
