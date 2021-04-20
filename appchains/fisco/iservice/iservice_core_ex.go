// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iservice

import (
	"math/big"
	"strings"

	"github.com/FISCO-BCOS/go-sdk/abi"
	"github.com/FISCO-BCOS/go-sdk/abi/bind"
	"github.com/FISCO-BCOS/go-sdk/core/types"
	"github.com/FISCO-BCOS/go-sdk/event"
	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// IServiceCoreExABI is the input ABI used to generate the binding from.
const IServiceCoreExABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"bytes32\"}],\"name\":\"getResponse\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"name\":\"_errMsg\",\"type\":\"string\"},{\"name\":\"_output\",\"type\":\"string\"},{\"name\":\"_icRequestID\",\"type\":\"string\"}],\"name\":\"setResponse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setIServiceMarket\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"iServiceMarket\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_serviceName\",\"type\":\"string\"},{\"name\":\"_input\",\"type\":\"string\"},{\"name\":\"_timeout\",\"type\":\"uint256\"},{\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"name\":\"_callbackFunction\",\"type\":\"bytes4\"}],\"name\":\"callService\",\"outputs\":[{\"name\":\"requestID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_iServiceMarket\",\"type\":\"address\"},{\"name\":\"_relayer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_serviceName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_provider\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_input\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_serviceFeeCap\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"ServiceInvoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// IServiceCoreExBin is the compiled bytecode used for deploying new contracts.
var IServiceCoreExBin = "0x60806040526040805190810160405280601781526020017f7b22686561646572223a7b7d2c22626f6479223a7b7d7d00000000000000000081525060079080519060200190620000519291906200037c565b503480156200005f57600080fd5b50604051604080620029b88339810180604052810190808051906020019092919080519060200190929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3620001638262000243640100000000026401000000009004565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515620001e15780600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200023b565b620001fa62000353640100000000026401000000009004565b600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50506200042b565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156200030f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260378152602001807f6953657276696365436f726545783a206953657276696365206d61726b65742081526020017f616464726573732063616e206e6f74206265207a65726f00000000000000000081525060400191505060405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620003bf57805160ff1916838001178555620003f0565b82800160010185558215620003f0579182015b82811115620003ef578251825591602001919060010190620003d2565b5b509050620003ff919062000403565b5090565b6200042891905b80821115620004245760008160009055506001016200040a565b5090565b90565b61257d806200043b6000396000f3006080604052600436106100af576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff168063415a2bc1146100b45780634bcee3e7146102365780635badbe4c146102c75780636548e9bc146102f25780637344f7a6146103355780638406c079146103785780638da5cb5b146103cf5780638f32d59b14610426578063b8aa1cc114610455578063cfb21ef6146104ac578063f2fde38b1461056e575b600080fd5b3480156100c057600080fd5b506100e360048036038101908080356000191690602001909291905050506105b1565b60405180806020018060200180602001848103845287818151815260200191508051906020019080838360005b8381101561012b578082015181840152602081019050610110565b50505050905090810190601f1680156101585780820380516001836020036101000a031916815260200191505b50848103835286818151815260200191508051906020019080838360005b83811015610191578082015181840152602081019050610176565b50505050905090810190601f1680156101be5780820380516001836020036101000a031916815260200191505b50848103825285818151815260200191508051906020019080838360005b838110156101f75780820151818401526020810190506101dc565b50505050905090810190601f1680156102245780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34801561024257600080fd5b506102ad60048036038101908080356000191690602001909291908035906020019082018035906020019190919293919293908035906020019082018035906020019190919293919293908035906020019082018035906020019190919293919293905050506107f1565b604051808215151515815260200191505060405180910390f35b3480156102d357600080fd5b506102dc610e3e565b6040518082815260200191505060405180910390f35b3480156102fe57600080fd5b50610333600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610e44565b005b34801561034157600080fd5b50610376600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610fcf565b005b34801561038457600080fd5b5061038d611057565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156103db57600080fd5b506103e461107d565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561043257600080fd5b5061043b6110a6565b604051808215151515815260200191505060405180910390f35b34801561046157600080fd5b5061046a6110fd565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b3480156104b857600080fd5b5061055060048036038101908080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080357bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19169060200190929190505050611123565b60405180826000191660001916815260200191505060405180910390f35b34801561057a57600080fd5b506105af600480360381019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919050505061144c565b005b6060806060600360008560001916600019168152602001908152602001600020600001600360008660001916600019168152602001908152602001600020600101600360008760001916600019168152602001908152602001600020600201828054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106a55780601f1061067a576101008083540402835291602001916106a5565b820191906000526020600020905b81548152906001019060200180831161068857829003601f168201915b50505050509250818054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107415780601f1061071657610100808354040283529160200191610741565b820191906000526020600020905b81548152906001019060200180831161072457829003601f168201915b50505050509150808054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107dd5780601f106107b2576101008083540402835291602001916107dd565b820191906000526020600020905b8154815290600101906020018083116107c057829003601f168201915b505050505090509250925092509193909250565b60006107fb6123e4565b610803612406565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156108f0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260298152602001807f6953657276696365436f726545783a2073656e646572206973206e6f7420746881526020017f652072656c61796572000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8a60006001600083600019166000191681526020019081526020016000206002018054600181600116156101000203166002900490501115156109c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f6953657276696365436f726545783a207265717565737420646f6573206e6f7481526020017f206578697374000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b6000600360008360001916600019168152602001908152602001600020600201805460018160011615610100020316600290049050141515610a91576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602a8152602001807f6953657276696365436f726545783a207265717565737420686173206265656e81526020017f20726573706f6e6465640000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050846000018190525086868080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050846040018190525060008b8b90501415610b4f5788888080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505084602001819052505b83600360008e600019166000191681526020019081526020016000206000820151816000019080519060200190610b87929190612455565b506020820151816001019080519060200190610ba4929190612455565b506040820151816002019080519060200190610bc1929190612455565b50905050600260008d600019166000191681526020019081526020016000206040805190810160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a90047c0100000000000000000000000000000000000000000000000000000000027bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815250509250826000015173ffffffffffffffffffffffffffffffffffffffff1683602001518d866020015160405160240180836000191660001916815260200180602001828103825283818151815260200191508051906020019080838360005b83811015610d2b578082015181840152602081019050610d10565b50505050905090810190601f168015610d585780820380516001836020036101000a031916815260200191505b509350505050604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405180828051906020019080838360005b83811015610de8578082015181840152602081019050610dcd565b50505050905090810190601f168015610e155780820380516001836020036101000a031916815260200191505b509150506000604051808303816000865af1915050915081945050505050979650505050505050565b60045481565b610e4c6110a6565b1515610ec0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610f8b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f6953657276696365436f726545783a2072656c6179657220616464726573732081526020017f63616e206e6f74206265207a65726f000000000000000000000000000000000081525060400191505060405180910390fd5b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b610fd76110a6565b151561104b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b611054816114d4565b50565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600061112d6124d5565b88888080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050508560008251111515611200576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602d8152602001807f6953657276696365436f726545783a2073657276696365206e616d652063616e81526020017f206e6f7420626520656d7074790000000000000000000000000000000000000081525060400191505060405180910390fd5b60008111151561129e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260368152602001807f6953657276696365436f726545783a20726571756573742074696d656f75742081526020017f6d7573742062652067726561746572207468616e20300000000000000000000081525060400191505060405180910390fd5b33836020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050836040018190525088888080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050508360800181905250868360c0018181525050600083608001515114156114085760078054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156113fa5780601f106113cf576101008083540402835291602001916113fa565b820191906000526020600020905b8154815290600101906020018083116113dd57829003601f168201915b505050505083608001819052505b611411836115e3565b61141a83611ced565b9350838360000190600019169081600019168152505061143b848787612124565b839350505050979650505050505050565b6114546110a6565b15156114c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b6114d18161225b565b50565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561159f576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260378152602001807f6953657276696365436f726545783a206953657276696365206d61726b65742081526020017f616464726573732063616e206e6f74206265207a65726f00000000000000000081525060400191505060405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600080606080600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638fd2daf686604001516040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561169757808201518184015260208101905061167c565b50505050905090810190601f1680156116c45780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1580156116e357600080fd5b505af11580156116f7573d6000803e3d6000fd5b505050506040513d602081101561170d57600080fd5b810190808051906020019092919050505093508315156117bb576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252603c8152602001807f6953657276696365436f726545783a207365727669636520646f6573206e6f7481526020017f20657869737420696e207468652073657276696365206d61726b65740000000081525060400191505060405180910390fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663834cbd7486604001516040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561186957808201518184015260208101905061184e565b50505050905090810190601f1680156118965780820380516001836020036101000a031916815260200191505b5092505050602060405180830381600087803b1580156118b557600080fd5b505af11580156118c9573d6000803e3d6000fd5b505050506040513d60208110156118df57600080fd5b81019080805190602001909291905050509250828560c00151101515156119ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260508152602001807f6953657276696365436f726545783a20726571756573742074696d656f75742081526020017f6d7573742062652067726561746572207468616e206f7220657175616c20746f81526020017f20746865207365727669636520516f530000000000000000000000000000000081525060600191505060405180910390fd5b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16631ebc077286604001516040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611a68578082015181840152602081019050611a4d565b50505050905090810190601f168015611a955780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015611ab457600080fd5b505af1158015611ac8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015611af257600080fd5b810190808051640100000000811115611b0a57600080fd5b82810190506020810184811115611b2057600080fd5b8151856001820283011164010000000082111715611b3d57600080fd5b50509291905050509150600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166318274da486604001516040518263ffffffff167c01000000000000000000000000000000000000000000000000000000000281526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611bf5578082015181840152602081019050611bda565b50505050905090810190601f168015611c225780820380516001836020036101000a031916815260200191505b5092505050600060405180830381600087803b158015611c4157600080fd5b505af1158015611c55573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052506020811015611c7f57600080fd5b810190808051640100000000811115611c9757600080fd5b82810190506020810184811115611cad57600080fd5b8151856001820283011164010000000082111715611cca57600080fd5b50509291905050509050818560600181905250808560a001819052505050505050565b600030600454604051602001808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166c01000000000000000000000000028152601401828152602001925050506040516020818303038152906040526040518082805190602001908083835b602083101515611d895780518252602082019150602081019050602083039250611d64565b6001836020036101000a038019825116818451168082178552505050505050905001915050604051809103902090508082600001906000191690816000191681525050816001600083600019166000191681526020019081526020016000206000820151816000019060001916905560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040820151816002019080519060200190611e5b929190612455565b506060820151816003019080519060200190611e78929190612455565b506080820151816004019080519060200190611e95929190612455565b5060a0820151816005019080519060200190611eb2929190612455565b5060c0820151816006015590505060016004600082825401925050819055507f40213482a53a92631c29a845a3c30e046e8914e970c2bb843b81fe5a11c55baf826000015183602001518460400151856060015186608001518760a001518860c001516040518088600019166000191681526020018773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020018060200180602001806020018060200186815260200185810385528a818151815260200191508051906020019080838360005b83811015611fa7578082015181840152602081019050611f8c565b50505050905090810190601f168015611fd45780820380516001836020036101000a031916815260200191505b50858103845289818151815260200191508051906020019080838360005b8381101561200d578082015181840152602081019050611ff2565b50505050905090810190601f16801561203a5780820380516001836020036101000a031916815260200191505b50858103835288818151815260200191508051906020019080838360005b83811015612073578082015181840152602081019050612058565b50505050905090810190601f1680156120a05780820380516001836020036101000a031916815260200191505b50858103825287818151815260200191508051906020019080838360005b838110156120d95780820151818401526020810190506120be565b50505050905090810190601f1680156121065780820380516001836020036101000a031916815260200191505b509b50505050505050505050505060405180910390a1809050919050565b61212c612406565b82816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508181602001907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815250508060026000866000191660001916815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548163ffffffff02191690837c01000000000000000000000000000000000000000000000000000000009004021790555090505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515612326576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526020017f646472657373000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6060604051908101604052806060815260200160608152602001606081525090565b6040805190810160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061249657805160ff19168380011785556124c4565b828001600101855582156124c4579182015b828111156124c35782518255916020019190600101906124a8565b5b5090506124d1919061252c565b5090565b60e06040519081016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001606081526020016060815260200160608152602001600081525090565b61254e91905b8082111561254a576000816000905550600101612532565b5090565b905600a165627a7a7230582030ab71efbc8d1bb329c7fb16e62e277e2feb2f68d97090df8dae688f8d194ebf0029"

// DeployIServiceCoreEx deploys a new contract, binding an instance of IServiceCoreEx to it.
func DeployIServiceCoreEx(auth *bind.TransactOpts, backend bind.ContractBackend, _iServiceMarket common.Address, _relayer common.Address) (common.Address, *types.Transaction, *IServiceCoreEx, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceCoreExABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IServiceCoreExBin), backend, _iServiceMarket, _relayer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IServiceCoreEx{IServiceCoreExCaller: IServiceCoreExCaller{contract: contract}, IServiceCoreExTransactor: IServiceCoreExTransactor{contract: contract}, IServiceCoreExFilterer: IServiceCoreExFilterer{contract: contract}}, nil
}

func AsyncDeployIServiceCoreEx(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _iServiceMarket common.Address, _relayer common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceCoreExABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(IServiceCoreExBin), backend, _iServiceMarket, _relayer)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// IServiceCoreEx is an auto generated Go binding around a Solidity contract.
type IServiceCoreEx struct {
	IServiceCoreExCaller     // Read-only binding to the contract
	IServiceCoreExTransactor // Write-only binding to the contract
	IServiceCoreExFilterer   // Log filterer for contract events
}

// IServiceCoreExCaller is an auto generated read-only Go binding around a Solidity contract.
type IServiceCoreExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExTransactor is an auto generated write-only Go binding around a Solidity contract.
type IServiceCoreExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IServiceCoreExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IServiceCoreExSession struct {
	Contract     *IServiceCoreEx   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IServiceCoreExCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IServiceCoreExCallerSession struct {
	Contract *IServiceCoreExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IServiceCoreExTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IServiceCoreExTransactorSession struct {
	Contract     *IServiceCoreExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IServiceCoreExRaw is an auto generated low-level Go binding around a Solidity contract.
type IServiceCoreExRaw struct {
	Contract *IServiceCoreEx // Generic contract binding to access the raw methods on
}

// IServiceCoreExCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IServiceCoreExCallerRaw struct {
	Contract *IServiceCoreExCaller // Generic read-only contract binding to access the raw methods on
}

// IServiceCoreExTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IServiceCoreExTransactorRaw struct {
	Contract *IServiceCoreExTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIServiceCoreEx creates a new instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreEx(address common.Address, backend bind.ContractBackend) (*IServiceCoreEx, error) {
	contract, err := bindIServiceCoreEx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreEx{IServiceCoreExCaller: IServiceCoreExCaller{contract: contract}, IServiceCoreExTransactor: IServiceCoreExTransactor{contract: contract}, IServiceCoreExFilterer: IServiceCoreExFilterer{contract: contract}}, nil
}

// NewIServiceCoreExCaller creates a new read-only instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreExCaller(address common.Address, caller bind.ContractCaller) (*IServiceCoreExCaller, error) {
	contract, err := bindIServiceCoreEx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExCaller{contract: contract}, nil
}

// NewIServiceCoreExTransactor creates a new write-only instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreExTransactor(address common.Address, transactor bind.ContractTransactor) (*IServiceCoreExTransactor, error) {
	contract, err := bindIServiceCoreEx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExTransactor{contract: contract}, nil
}

// NewIServiceCoreExFilterer creates a new log filterer instance of IServiceCoreEx, bound to a specific deployed contract.
func NewIServiceCoreExFilterer(address common.Address, filterer bind.ContractFilterer) (*IServiceCoreExFilterer, error) {
	contract, err := bindIServiceCoreEx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExFilterer{contract: contract}, nil
}

// bindIServiceCoreEx binds a generic wrapper to an already deployed contract.
func bindIServiceCoreEx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceCoreExABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IServiceCoreEx *IServiceCoreExRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IServiceCoreEx.Contract.IServiceCoreExCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IServiceCoreEx *IServiceCoreExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.IServiceCoreExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceCoreEx *IServiceCoreExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.IServiceCoreExTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IServiceCoreEx *IServiceCoreExCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IServiceCoreEx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IServiceCoreEx *IServiceCoreExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceCoreEx *IServiceCoreExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.contract.Transact(opts, method, params...)
}

// GetResponse is a free data retrieval call binding the contract method 0x744e931a.
//
// Solidity: function getResponse(bytes32 _requestID) constant returns(string, string, string)
func (_IServiceCoreEx *IServiceCoreExCaller) GetResponse(opts *bind.CallOpts, _requestID [32]byte) (string, string, string, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _IServiceCoreEx.contract.Call(opts, out, "getResponse", _requestID)
	return *ret0, *ret1, *ret2, err
}

// GetResponse is a free data retrieval call binding the contract method 0x744e931a.
//
// Solidity: function getResponse(bytes32 _requestID) constant returns(string, string, string)
func (_IServiceCoreEx *IServiceCoreExSession) GetResponse(_requestID [32]byte) (string, string, string, error) {
	return _IServiceCoreEx.Contract.GetResponse(&_IServiceCoreEx.CallOpts, _requestID)
}

// GetResponse is a free data retrieval call binding the contract method 0x744e931a.
//
// Solidity: function getResponse(bytes32 _requestID) constant returns(string, string, string)
func (_IServiceCoreEx *IServiceCoreExCallerSession) GetResponse(_requestID [32]byte) (string, string, string, error) {
	return _IServiceCoreEx.Contract.GetResponse(&_IServiceCoreEx.CallOpts, _requestID)
}

// IServiceMarket is a free data retrieval call binding the contract method 0xa70ac4dc.
//
// Solidity: function iServiceMarket() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCaller) IServiceMarket(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "iServiceMarket")
	return *ret0, err
}

// IServiceMarket is a free data retrieval call binding the contract method 0xa70ac4dc.
//
// Solidity: function iServiceMarket() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExSession) IServiceMarket() (common.Address, error) {
	return _IServiceCoreEx.Contract.IServiceMarket(&_IServiceCoreEx.CallOpts)
}

// IServiceMarket is a free data retrieval call binding the contract method 0xa70ac4dc.
//
// Solidity: function iServiceMarket() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCallerSession) IServiceMarket() (common.Address, error) {
	return _IServiceCoreEx.Contract.IServiceMarket(&_IServiceCoreEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0xede8e529.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceCoreEx *IServiceCoreExCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0xede8e529.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) IsOwner() (bool, error) {
	return _IServiceCoreEx.Contract.IsOwner(&_IServiceCoreEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0xede8e529.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceCoreEx *IServiceCoreExCallerSession) IsOwner() (bool, error) {
	return _IServiceCoreEx.Contract.IsOwner(&_IServiceCoreEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x5089e2c8.
//
// Solidity: function owner() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x5089e2c8.
//
// Solidity: function owner() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExSession) Owner() (common.Address, error) {
	return _IServiceCoreEx.Contract.Owner(&_IServiceCoreEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x5089e2c8.
//
// Solidity: function owner() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCallerSession) Owner() (common.Address, error) {
	return _IServiceCoreEx.Contract.Owner(&_IServiceCoreEx.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0xe9667acd.
//
// Solidity: function relayer() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCaller) Relayer(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "relayer")
	return *ret0, err
}

// Relayer is a free data retrieval call binding the contract method 0xe9667acd.
//
// Solidity: function relayer() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExSession) Relayer() (common.Address, error) {
	return _IServiceCoreEx.Contract.Relayer(&_IServiceCoreEx.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0xe9667acd.
//
// Solidity: function relayer() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCallerSession) Relayer() (common.Address, error) {
	return _IServiceCoreEx.Contract.Relayer(&_IServiceCoreEx.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x3a938f2d.
//
// Solidity: function requestCount() constant returns(uint256)
func (_IServiceCoreEx *IServiceCoreExCaller) RequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "requestCount")
	return *ret0, err
}

// RequestCount is a free data retrieval call binding the contract method 0x3a938f2d.
//
// Solidity: function requestCount() constant returns(uint256)
func (_IServiceCoreEx *IServiceCoreExSession) RequestCount() (*big.Int, error) {
	return _IServiceCoreEx.Contract.RequestCount(&_IServiceCoreEx.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x3a938f2d.
//
// Solidity: function requestCount() constant returns(uint256)
func (_IServiceCoreEx *IServiceCoreExCallerSession) RequestCount() (*big.Int, error) {
	return _IServiceCoreEx.Contract.RequestCount(&_IServiceCoreEx.CallOpts)
}

// CallService is a paid mutator transaction binding the contract method 0x02b43c19.
//
// Solidity: function callService(string _serviceName, string _input, uint256 _timeout, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExTransactor) CallService(opts *bind.TransactOpts, _serviceName string, _input string, _timeout *big.Int, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "callService", _serviceName, _input, _timeout, _callbackAddress, _callbackFunction)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncCallService(handler func(*types.Receipt, error), opts *bind.TransactOpts, _serviceName string, _input string, _timeout *big.Int, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "callService", _serviceName, _input, _timeout, _callbackAddress, _callbackFunction)
}

// CallService is a paid mutator transaction binding the contract method 0x02b43c19.
//
// Solidity: function callService(string _serviceName, string _input, uint256 _timeout, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExSession) CallService(_serviceName string, _input string, _timeout *big.Int, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.CallService(&_IServiceCoreEx.TransactOpts, _serviceName, _input, _timeout, _callbackAddress, _callbackFunction)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncCallService(handler func(*types.Receipt, error), _serviceName string, _input string, _timeout *big.Int, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncCallService(handler, &_IServiceCoreEx.TransactOpts, _serviceName, _input, _timeout, _callbackAddress, _callbackFunction)
}

// CallService is a paid mutator transaction binding the contract method 0x02b43c19.
//
// Solidity: function callService(string _serviceName, string _input, uint256 _timeout, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) CallService(_serviceName string, _input string, _timeout *big.Int, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.CallService(&_IServiceCoreEx.TransactOpts, _serviceName, _input, _timeout, _callbackAddress, _callbackFunction)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncCallService(handler func(*types.Receipt, error), _serviceName string, _input string, _timeout *big.Int, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncCallService(handler, &_IServiceCoreEx.TransactOpts, _serviceName, _input, _timeout, _callbackAddress, _callbackFunction)
}

// SetIServiceMarket is a paid mutator transaction binding the contract method 0x650d6e85.
//
// Solidity: function setIServiceMarket(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) SetIServiceMarket(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setIServiceMarket", _address)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncSetIServiceMarket(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "setIServiceMarket", _address)
}

// SetIServiceMarket is a paid mutator transaction binding the contract method 0x650d6e85.
//
// Solidity: function setIServiceMarket(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExSession) SetIServiceMarket(_address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetIServiceMarket(&_IServiceCoreEx.TransactOpts, _address)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncSetIServiceMarket(handler func(*types.Receipt, error), _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetIServiceMarket(handler, &_IServiceCoreEx.TransactOpts, _address)
}

// SetIServiceMarket is a paid mutator transaction binding the contract method 0x650d6e85.
//
// Solidity: function setIServiceMarket(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetIServiceMarket(_address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetIServiceMarket(&_IServiceCoreEx.TransactOpts, _address)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncSetIServiceMarket(handler func(*types.Receipt, error), _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetIServiceMarket(handler, &_IServiceCoreEx.TransactOpts, _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x7a1f2dea.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) SetRelayer(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setRelayer", _address)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncSetRelayer(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "setRelayer", _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x7a1f2dea.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExSession) SetRelayer(_address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetRelayer(&_IServiceCoreEx.TransactOpts, _address)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncSetRelayer(handler func(*types.Receipt, error), _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetRelayer(handler, &_IServiceCoreEx.TransactOpts, _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x7a1f2dea.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetRelayer(_address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetRelayer(&_IServiceCoreEx.TransactOpts, _address)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncSetRelayer(handler func(*types.Receipt, error), _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetRelayer(handler, &_IServiceCoreEx.TransactOpts, _address)
}

// SetResponse is a paid mutator transaction binding the contract method 0x781c89c6.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output, string _icRequestID) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactor) SetResponse(opts *bind.TransactOpts, _requestID [32]byte, _errMsg string, _output string, _icRequestID string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setResponse", _requestID, _errMsg, _output, _icRequestID)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncSetResponse(handler func(*types.Receipt, error), opts *bind.TransactOpts, _requestID [32]byte, _errMsg string, _output string, _icRequestID string) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "setResponse", _requestID, _errMsg, _output, _icRequestID)
}

// SetResponse is a paid mutator transaction binding the contract method 0x781c89c6.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output, string _icRequestID) returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) SetResponse(_requestID [32]byte, _errMsg string, _output string, _icRequestID string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetResponse(&_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output, _icRequestID)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncSetResponse(handler func(*types.Receipt, error), _requestID [32]byte, _errMsg string, _output string, _icRequestID string) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetResponse(handler, &_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output, _icRequestID)
}

// SetResponse is a paid mutator transaction binding the contract method 0x781c89c6.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output, string _icRequestID) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetResponse(_requestID [32]byte, _errMsg string, _output string, _icRequestID string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetResponse(&_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output, _icRequestID)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncSetResponse(handler func(*types.Receipt, error), _requestID [32]byte, _errMsg string, _output string, _icRequestID string) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetResponse(handler, &_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output, _icRequestID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x16cad12a.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncTransferOwnership(handler func(*types.Receipt, error), opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x16cad12a.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.TransferOwnership(&_IServiceCoreEx.TransactOpts, newOwner)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncTransferOwnership(handler, &_IServiceCoreEx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x16cad12a.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.TransferOwnership(&_IServiceCoreEx.TransactOpts, newOwner)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncTransferOwnership(handler, &_IServiceCoreEx.TransactOpts, newOwner)
}

// IServiceCoreExOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IServiceCoreEx contract.
type IServiceCoreExOwnershipTransferredIterator struct {
	Event *IServiceCoreExOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IServiceCoreExOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceCoreExOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IServiceCoreExOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IServiceCoreExOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceCoreExOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceCoreExOwnershipTransferred represents a OwnershipTransferred event raised by the IServiceCoreEx contract.
type IServiceCoreExOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x000000000000000000000000000000000000000000000000000000005c7c30d4.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IServiceCoreExOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExOwnershipTransferredIterator{contract: _IServiceCoreEx.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x000000000000000000000000000000000000000000000000000000005c7c30d4.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceCoreEx *IServiceCoreExFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IServiceCoreExOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IServiceCoreEx.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceCoreExOwnershipTransferred)
				if err := _IServiceCoreEx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x000000000000000000000000000000000000000000000000000000005c7c30d4.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseOwnershipTransferred(log types.Log) (*IServiceCoreExOwnershipTransferred, error) {
	event := new(IServiceCoreExOwnershipTransferred)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IServiceCoreExServiceInvokedIterator is returned from FilterServiceInvoked and is used to iterate over the raw logs and unpacked data for ServiceInvoked events raised by the IServiceCoreEx contract.
type IServiceCoreExServiceInvokedIterator struct {
	Event *IServiceCoreExServiceInvoked // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *IServiceCoreExServiceInvokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceCoreExServiceInvoked)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(IServiceCoreExServiceInvoked)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *IServiceCoreExServiceInvokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceCoreExServiceInvokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceCoreExServiceInvoked represents a ServiceInvoked event raised by the IServiceCoreEx contract.
type IServiceCoreExServiceInvoked struct {
	RequestID       [32]byte
	ContractAddress common.Address
	ServiceName     string
	Provider        string
	Input           string
	ServiceFeeCap   string
	Timeout         *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterServiceInvoked is a free log retrieval operation binding the contract event 0x000000000000000000000000000000000000000000000000000000009806d3fe.
//
// Solidity: event ServiceInvoked(bytes32 _requestID, address _contractAddress, string _serviceName, string _provider, string _input, string _serviceFeeCap, uint256 _timeout)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterServiceInvoked(opts *bind.FilterOpts) (*IServiceCoreExServiceInvokedIterator, error) {

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "ServiceInvoked")
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExServiceInvokedIterator{contract: _IServiceCoreEx.contract, event: "ServiceInvoked", logs: logs, sub: sub}, nil
}

// WatchServiceInvoked is a free log subscription operation binding the contract event 0x000000000000000000000000000000000000000000000000000000009806d3fe.
//
// Solidity: event ServiceInvoked(bytes32 _requestID, address _contractAddress, string _serviceName, string _provider, string _input, string _serviceFeeCap, uint256 _timeout)
func (_IServiceCoreEx *IServiceCoreExFilterer) WatchServiceInvoked(opts *bind.WatchOpts, sink chan<- *IServiceCoreExServiceInvoked) (event.Subscription, error) {

	logs, sub, err := _IServiceCoreEx.contract.WatchLogs(opts, "ServiceInvoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceCoreExServiceInvoked)
				if err := _IServiceCoreEx.contract.UnpackLog(event, "ServiceInvoked", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseServiceInvoked is a log parse operation binding the contract event 0x000000000000000000000000000000000000000000000000000000009806d3fe.
//
// Solidity: event ServiceInvoked(bytes32 _requestID, address _contractAddress, string _serviceName, string _provider, string _input, string _serviceFeeCap, uint256 _timeout)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseServiceInvoked(log types.Log) (*IServiceCoreExServiceInvoked, error) {
	event := new(IServiceCoreExServiceInvoked)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "ServiceInvoked", log); err != nil {
		return nil, err
	}
	return event, nil
}
