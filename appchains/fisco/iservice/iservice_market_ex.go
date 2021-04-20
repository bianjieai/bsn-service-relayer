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

// IServiceMarketExABI is the input ABI used to generate the binding from.
const IServiceMarketExABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"getServiceFee\",\"outputs\":[{\"name\":\"fee\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"getServiceProvider\",\"outputs\":[{\"name\":\"provider\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"getServiceBinding\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"string\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_serviceName\",\"type\":\"string\"},{\"name\":\"_provider\",\"type\":\"string\"},{\"name\":\"_serviceFee\",\"type\":\"string\"},{\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"updateServiceBinding\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"getQoS\",\"outputs\":[{\"name\":\"qos\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getServiceBindingCount\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"serviceBindingExists\",\"outputs\":[{\"name\":\"exist\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_serviceName\",\"type\":\"string\"},{\"name\":\"_schemas\",\"type\":\"string\"},{\"name\":\"_provider\",\"type\":\"string\"},{\"name\":\"_serviceFee\",\"type\":\"string\"},{\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"addServiceBinding\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"removeServiceBinding\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_serviceName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_schemas\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_provider\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_serviceFee\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"ServiceBindingAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_serviceName\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_provider\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_serviceFee\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"ServiceBindingUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"ServiceBindingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// IServiceMarketExBin is the compiled bytecode used for deploying new contracts.
var IServiceMarketExBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3612500806100dc6000396000f3006080604052600436106100ba576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff16806318274da4146100bf5780631ebc0772146101a1578063434590df146102835780636e29ea60146104b0578063834cbd7414610525578063842675a3146105a25780638da5cb5b146105cd5780638f32d59b146106245780638fd2daf61461065357806392ce31c4146106d4578063bce9f63f14610761578063f2fde38b1461079c575b600080fd5b3480156100cb57600080fd5b50610126600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506107df565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561016657808201518184015260208101905061014b565b50505050905090810190601f1680156101935780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101ad57600080fd5b50610208600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061095e565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561024857808201518184015260208101905061022d565b50505050905090810190601f1680156102755780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561028f57600080fd5b506102ea600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610add565b604051808060200180602001806020018060200186815260200185810385528a818151815260200191508051906020019080838360005b8381101561033c578082015181840152602081019050610321565b50505050905090810190601f1680156103695780820380516001836020036101000a031916815260200191505b50858103845289818151815260200191508051906020019080838360005b838110156103a2578082015181840152602081019050610387565b50505050905090810190601f1680156103cf5780820380516001836020036101000a031916815260200191505b50858103835288818151815260200191508051906020019080838360005b838110156104085780820151818401526020810190506103ed565b50505050905090810190601f1680156104355780820380516001836020036101000a031916815260200191505b50858103825287818151815260200191508051906020019080838360005b8381101561046e578082015181840152602081019050610453565b50505050905090810190601f16801561049b5780820380516001836020036101000a031916815260200191505b50995050505050505050505060405180910390f35b3480156104bc57600080fd5b5061052360048036038101908080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939080359060200190820180359060200191909192939192939080359060200190929190505050610eb0565b005b34801561053157600080fd5b5061058c600480360381019080803590602001908201803590602001908080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505091929192905050506111b6565b6040518082815260200191505060405180910390f35b3480156105ae57600080fd5b506105b761129d565b6040518082815260200191505060405180910390f35b3480156105d957600080fd5b506105e26112aa565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561063057600080fd5b506106396112d3565b604051808215151515815260200191505060405180910390f35b34801561065f57600080fd5b506106ba600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061132a565b604051808215151515815260200191505060405180910390f35b3480156106e057600080fd5b5061075f600480360381019080803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803590602001908201803590602001919091929391929390803590602001909291905050506113af565b005b34801561076d57600080fd5b5061079a600480360381019080803590602001908201803590602001919091929391929390505050611a75565b005b3480156107a857600080fd5b506107dd600480360381019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611d5c565b005b60606107e961231b565b6002836040518082805190602001908083835b60208310151561082157805182526020820191506020810190506020830392506107fc565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020604080519081016040529081600082015481526020016001820160009054906101000a900460ff161515151581525050905080602001511561095457600181600001518154811015156108a357fe5b90600052602060002090600502016003018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109485780601f1061091d57610100808354040283529160200191610948565b820191906000526020600020905b81548152906001019060200180831161092b57829003601f168201915b50505050509150610958565b8191505b50919050565b606061096861231b565b6002836040518082805190602001908083835b6020831015156109a0578051825260208201915060208101905060208303925061097b565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020604080519081016040529081600082015481526020016001820160009054906101000a900460ff1615151515815250509050806020015115610ad35760018160000151815481101515610a2257fe5b90600052602060002090600502016002018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ac75780601f10610a9c57610100808354040283529160200191610ac7565b820191906000526020600020905b815481529060010190602001808311610aaa57829003601f168201915b50505050509150610ad7565b8191505b50919050565b6060806060806000610aed61231b565b60006002886040518082805190602001908083835b602083101515610b275780518252602082019150602081019050602083039250610b02565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020604080519081016040529081600082015481526020016001820160009054906101000a900460ff1615151515815250509150816020015115610e4d5760018260000151815481101515610ba957fe5b90600052602060002090600502019050806000018160010182600201836003018460040154848054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c635780601f10610c3857610100808354040283529160200191610c63565b820191906000526020600020905b815481529060010190602001808311610c4657829003601f168201915b50505050509450838054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cff5780601f10610cd457610100808354040283529160200191610cff565b820191906000526020600020905b815481529060010190602001808311610ce257829003601f168201915b50505050509350828054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d9b5780601f10610d7057610100808354040283529160200191610d9b565b820191906000526020600020905b815481529060010190602001808311610d7e57829003601f168201915b50505050509250818054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e375780601f10610e0c57610100808354040283529160200191610e37565b820191906000526020600020905b815481529060010190602001808311610e1a57829003601f168201915b5050505050915096509650965096509650610ea5565b6000602060405190810160405280600081525090602060405190810160405280600081525090602060405190810160405280600081525090602060405190810160405280600081525090809050965096509650965096505b505091939590929450565b6000610eba6112d3565b1515610f2e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b87878080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050506002816040518082805190602001908083835b602083101515610f995780518252602082019150602081019050602083039250610f74565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff161515611077576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001807f69536572766963654d61726b657445783a20736572766963652062696e64696e81526020017f6720646f6573206e6f742065786973740000000000000000000000000000000081525060400191505060405180910390fd5b600160028a8a604051808383808284378201915050925050509081526020016040518091039020600001548154811015156110ae57fe5b9060005260206000209060050201915060008787905011156110df5786868360020191906110dd929190612337565b505b60008585905011156111005784848360030191906110fe929190612337565b505b6000831115611113578282600401819055505b88886040518083838082843782019150509250505060405180910390207fac059a9f400ca20e9086f7e73b998697778b9a48fe40d2392bf649752c5538b28888888888604051808060200180602001848152602001838103835288888281815260200192508082843782019150508381038252868682818152602001925080828437820191505097505050505050505060405180910390a2505050505050505050565b60006111c061231b565b6002836040518082805190602001908083835b6020831015156111f857805182526020820191506020810190506020830392506111d3565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020604080519081016040529081600082015481526020016001820160009054906101000a900460ff1615151515815250509050806020015115611293576001816000015181548110151561127a57fe5b9060005260206000209060050201600401549150611297565b8191505b50919050565b6000600180549050905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b60006002826040518082805190602001908083835b602083101515611364578051825260208201915060208101905060208303925061133f565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff169050919050565b6113b76123b7565b6113bf6112d3565b1515611433576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b89898080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505088888080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505087878080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505086868080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050856000855111151561159f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602f8152602001807f69536572766963654d61726b657445783a2073657276696365206e616d65206381526020017f616e206e6f7420626520656d707479000000000000000000000000000000000081525060400191505060405180910390fd5b6000845111151561163e576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260328152602001807f69536572766963654d61726b657445783a207365727669636520736368656d6181526020017f732063616e206e6f7420626520656d707479000000000000000000000000000081525060400191505060405180910390fd5b600083511115156116dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260338152602001807f69536572766963654d61726b657445783a20736572766963652070726f76696481526020017f65722063616e206e6f7420626520656d7074790000000000000000000000000081525060400191505060405180910390fd5b6000825111151561177c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f69536572766963654d61726b657445783a20736572766963652066656520636181526020017f6e206e6f7420626520656d70747900000000000000000000000000000000000081525060400191505060405180910390fd5b60008111151561181a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252602e8152602001807f69536572766963654d61726b657445783a20716f732073686f756c642062652081526020017f67726561746572207468616e203000000000000000000000000000000000000081525060400191505060405180910390fd5b8e8e8080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050506002816040518082805190602001908083835b6020831015156118855780518252602082019150602081019050602083039250611860565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff16151515611964576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001807f69536572766963654d61726b657445783a20736572766963652062696e64696e81526020017f6720616c7265616479206578697374730000000000000000000000000000000081525060400191505060405180910390fd5b8f8f8080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505087600001819052508d8d8080601f01602080910402602001604051908101604052809392919081815260200183838082843782019150505050505087602001819052508b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050876040018190525089898080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050876060018190525087876080018181525050611a6387611de4565b50505050505050505050505050505050565b611a7d6112d3565b1515611af1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b81818080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050506002816040518082805190602001908083835b602083101515611b5c5780518252602082019150602081019050602083039250611b37565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160009054906101000a900460ff161515611c3a576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260308152602001807f69536572766963654d61726b657445783a20736572766963652062696e64696e81526020017f6720646f6573206e6f742065786973740000000000000000000000000000000081525060400191505060405180910390fd5b60016002848460405180838380828437820191505092505050908152602001604051809103902060000154815481101515611c7157fe5b906000526020600020906005020160008082016000611c9091906123e7565b600182016000611ca091906123e7565b600282016000611cb091906123e7565b600382016000611cc091906123e7565b60048201600090555050600283836040518083838082843782019150509250505090815260200160405180910390206000808201600090556001820160006101000a81549060ff0219169055505082826040518083838082843782019150509250505060405180910390207f045edfc0de1510eef86edd984a529791dd2924769af547ab858d860139c7a54d60405160405180910390a2505050565b611d646112d3565b1515611dd8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281525060200191505060405180910390fd5b611de181612192565b50565b600180549050600282600001516040518082805190602001908083835b602083101515611e265780518252602082019150602081019050602083039250611e01565b6001836020036101000a0380198251168184511680821785525050505050509050019150509081526020016040518091039020600001819055506001600282600001516040518082805190602001908083835b602083101515611e9e5780518252602082019150602081019050602083039250611e79565b6001836020036101000a038019825116818451168082178552505050505050905001915050908152602001604051809103902060010160006101000a81548160ff021916908315150217905550600181908060018154018082558091505090600182039060005260206000209060050201600090919290919091506000820151816000019080519060200190611f3592919061242f565b506020820151816001019080519060200190611f5292919061242f565b506040820151816002019080519060200190611f6f92919061242f565b506060820151816003019080519060200190611f8c92919061242f565b506080820151816004015550505080600001516040518082805190602001908083835b602083101515611fd45780518252602082019150602081019050602083039250611faf565b6001836020036101000a03801982511681845116808217855250505050505090500191505060405180910390207f365f7e387fe001fa02d1c008ec11e2c21f47c6a2102f497b300a17c83127dc69826020015183604001518460600151856080015160405180806020018060200180602001858152602001848103845288818151815260200191508051906020019080838360005b83811015612084578082015181840152602081019050612069565b50505050905090810190601f1680156120b15780820380516001836020036101000a031916815260200191505b50848103835287818151815260200191508051906020019080838360005b838110156120ea5780820151818401526020810190506120cf565b50505050905090810190601f1680156121175780820380516001836020036101000a031916815260200191505b50848103825286818151815260200191508051906020019080838360005b83811015612150578082015181840152602081019050612135565b50505050905090810190601f16801561217d5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390a250565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561225d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260268152602001807f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206181526020017f646472657373000000000000000000000000000000000000000000000000000081525060400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040805190810160405280600081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061237857803560ff19168380011785556123a6565b828001600101855582156123a6579182015b828111156123a557823582559160200191906001019061238a565b5b5090506123b391906124af565b5090565b60a06040519081016040528060608152602001606081526020016060815260200160608152602001600081525090565b50805460018160011615610100020316600290046000825580601f1061240d575061242c565b601f01602090049060005260206000209081019061242b91906124af565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061247057805160ff191683800117855561249e565b8280016001018555821561249e579182015b8281111561249d578251825591602001919060010190612482565b5b5090506124ab91906124af565b5090565b6124d191905b808211156124cd5760008160009055506001016124b5565b5090565b905600a165627a7a72305820e46d909fcecad98052403044d83c0073bca6778682ab7617471b14139cceb1890029"

// DeployIServiceMarketEx deploys a new contract, binding an instance of IServiceMarketEx to it.
func DeployIServiceMarketEx(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *IServiceMarketEx, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceMarketExABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IServiceMarketExBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IServiceMarketEx{IServiceMarketExCaller: IServiceMarketExCaller{contract: contract}, IServiceMarketExTransactor: IServiceMarketExTransactor{contract: contract}, IServiceMarketExFilterer: IServiceMarketExFilterer{contract: contract}}, nil
}

func AsyncDeployIServiceMarketEx(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceMarketExABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(IServiceMarketExBin), backend)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// IServiceMarketEx is an auto generated Go binding around a Solidity contract.
type IServiceMarketEx struct {
	IServiceMarketExCaller     // Read-only binding to the contract
	IServiceMarketExTransactor // Write-only binding to the contract
	IServiceMarketExFilterer   // Log filterer for contract events
}

// IServiceMarketExCaller is an auto generated read-only Go binding around a Solidity contract.
type IServiceMarketExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceMarketExTransactor is an auto generated write-only Go binding around a Solidity contract.
type IServiceMarketExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceMarketExFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IServiceMarketExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceMarketExSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IServiceMarketExSession struct {
	Contract     *IServiceMarketEx // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IServiceMarketExCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IServiceMarketExCallerSession struct {
	Contract *IServiceMarketExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IServiceMarketExTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IServiceMarketExTransactorSession struct {
	Contract     *IServiceMarketExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IServiceMarketExRaw is an auto generated low-level Go binding around a Solidity contract.
type IServiceMarketExRaw struct {
	Contract *IServiceMarketEx // Generic contract binding to access the raw methods on
}

// IServiceMarketExCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IServiceMarketExCallerRaw struct {
	Contract *IServiceMarketExCaller // Generic read-only contract binding to access the raw methods on
}

// IServiceMarketExTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IServiceMarketExTransactorRaw struct {
	Contract *IServiceMarketExTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIServiceMarketEx creates a new instance of IServiceMarketEx, bound to a specific deployed contract.
func NewIServiceMarketEx(address common.Address, backend bind.ContractBackend) (*IServiceMarketEx, error) {
	contract, err := bindIServiceMarketEx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IServiceMarketEx{IServiceMarketExCaller: IServiceMarketExCaller{contract: contract}, IServiceMarketExTransactor: IServiceMarketExTransactor{contract: contract}, IServiceMarketExFilterer: IServiceMarketExFilterer{contract: contract}}, nil
}

// NewIServiceMarketExCaller creates a new read-only instance of IServiceMarketEx, bound to a specific deployed contract.
func NewIServiceMarketExCaller(address common.Address, caller bind.ContractCaller) (*IServiceMarketExCaller, error) {
	contract, err := bindIServiceMarketEx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceMarketExCaller{contract: contract}, nil
}

// NewIServiceMarketExTransactor creates a new write-only instance of IServiceMarketEx, bound to a specific deployed contract.
func NewIServiceMarketExTransactor(address common.Address, transactor bind.ContractTransactor) (*IServiceMarketExTransactor, error) {
	contract, err := bindIServiceMarketEx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceMarketExTransactor{contract: contract}, nil
}

// NewIServiceMarketExFilterer creates a new log filterer instance of IServiceMarketEx, bound to a specific deployed contract.
func NewIServiceMarketExFilterer(address common.Address, filterer bind.ContractFilterer) (*IServiceMarketExFilterer, error) {
	contract, err := bindIServiceMarketEx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IServiceMarketExFilterer{contract: contract}, nil
}

// bindIServiceMarketEx binds a generic wrapper to an already deployed contract.
func bindIServiceMarketEx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceMarketExABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IServiceMarketEx *IServiceMarketExRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IServiceMarketEx.Contract.IServiceMarketExCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IServiceMarketEx *IServiceMarketExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.IServiceMarketExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceMarketEx *IServiceMarketExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.IServiceMarketExTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IServiceMarketEx *IServiceMarketExCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IServiceMarketEx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IServiceMarketEx *IServiceMarketExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceMarketEx *IServiceMarketExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.contract.Transact(opts, method, params...)
}

// GetQoS is a free data retrieval call binding the contract method 0xeac81af8.
//
// Solidity: function getQoS(string _serviceName) constant returns(uint256 qos)
func (_IServiceMarketEx *IServiceMarketExCaller) GetQoS(opts *bind.CallOpts, _serviceName string) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IServiceMarketEx.contract.Call(opts, out, "getQoS", _serviceName)
	return *ret0, err
}

// GetQoS is a free data retrieval call binding the contract method 0xeac81af8.
//
// Solidity: function getQoS(string _serviceName) constant returns(uint256 qos)
func (_IServiceMarketEx *IServiceMarketExSession) GetQoS(_serviceName string) (*big.Int, error) {
	return _IServiceMarketEx.Contract.GetQoS(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetQoS is a free data retrieval call binding the contract method 0xeac81af8.
//
// Solidity: function getQoS(string _serviceName) constant returns(uint256 qos)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetQoS(_serviceName string) (*big.Int, error) {
	return _IServiceMarketEx.Contract.GetQoS(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceBinding is a free data retrieval call binding the contract method 0x81b92697.
//
// Solidity: function getServiceBinding(string _serviceName) constant returns(string, string, string, string, uint256)
func (_IServiceMarketEx *IServiceMarketExCaller) GetServiceBinding(opts *bind.CallOpts, _serviceName string) (string, string, string, string, *big.Int, error) {
	var (
		ret0 = new(string)
		ret1 = new(string)
		ret2 = new(string)
		ret3 = new(string)
		ret4 = new(*big.Int)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
		ret3,
		ret4,
	}
	err := _IServiceMarketEx.contract.Call(opts, out, "getServiceBinding", _serviceName)
	return *ret0, *ret1, *ret2, *ret3, *ret4, err
}

// GetServiceBinding is a free data retrieval call binding the contract method 0x81b92697.
//
// Solidity: function getServiceBinding(string _serviceName) constant returns(string, string, string, string, uint256)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBinding(_serviceName string) (string, string, string, string, *big.Int, error) {
	return _IServiceMarketEx.Contract.GetServiceBinding(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceBinding is a free data retrieval call binding the contract method 0x81b92697.
//
// Solidity: function getServiceBinding(string _serviceName) constant returns(string, string, string, string, uint256)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBinding(_serviceName string) (string, string, string, string, *big.Int, error) {
	return _IServiceMarketEx.Contract.GetServiceBinding(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceBindingCount is a free data retrieval call binding the contract method 0x0a9bdf07.
//
// Solidity: function getServiceBindingCount() constant returns(uint256 count)
func (_IServiceMarketEx *IServiceMarketExCaller) GetServiceBindingCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IServiceMarketEx.contract.Call(opts, out, "getServiceBindingCount")
	return *ret0, err
}

// GetServiceBindingCount is a free data retrieval call binding the contract method 0x0a9bdf07.
//
// Solidity: function getServiceBindingCount() constant returns(uint256 count)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBindingCount() (*big.Int, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingCount(&_IServiceMarketEx.CallOpts)
}

// GetServiceBindingCount is a free data retrieval call binding the contract method 0x0a9bdf07.
//
// Solidity: function getServiceBindingCount() constant returns(uint256 count)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBindingCount() (*big.Int, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingCount(&_IServiceMarketEx.CallOpts)
}

// GetServiceFee is a free data retrieval call binding the contract method 0x9b0a965c.
//
// Solidity: function getServiceFee(string _serviceName) constant returns(string fee)
func (_IServiceMarketEx *IServiceMarketExCaller) GetServiceFee(opts *bind.CallOpts, _serviceName string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IServiceMarketEx.contract.Call(opts, out, "getServiceFee", _serviceName)
	return *ret0, err
}

// GetServiceFee is a free data retrieval call binding the contract method 0x9b0a965c.
//
// Solidity: function getServiceFee(string _serviceName) constant returns(string fee)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceFee(_serviceName string) (string, error) {
	return _IServiceMarketEx.Contract.GetServiceFee(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceFee is a free data retrieval call binding the contract method 0x9b0a965c.
//
// Solidity: function getServiceFee(string _serviceName) constant returns(string fee)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceFee(_serviceName string) (string, error) {
	return _IServiceMarketEx.Contract.GetServiceFee(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceProvider is a free data retrieval call binding the contract method 0x438166c6.
//
// Solidity: function getServiceProvider(string _serviceName) constant returns(string provider)
func (_IServiceMarketEx *IServiceMarketExCaller) GetServiceProvider(opts *bind.CallOpts, _serviceName string) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _IServiceMarketEx.contract.Call(opts, out, "getServiceProvider", _serviceName)
	return *ret0, err
}

// GetServiceProvider is a free data retrieval call binding the contract method 0x438166c6.
//
// Solidity: function getServiceProvider(string _serviceName) constant returns(string provider)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceProvider(_serviceName string) (string, error) {
	return _IServiceMarketEx.Contract.GetServiceProvider(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceProvider is a free data retrieval call binding the contract method 0x438166c6.
//
// Solidity: function getServiceProvider(string _serviceName) constant returns(string provider)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceProvider(_serviceName string) (string, error) {
	return _IServiceMarketEx.Contract.GetServiceProvider(&_IServiceMarketEx.CallOpts, _serviceName)
}

// IsOwner is a free data retrieval call binding the contract method 0xede8e529.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceMarketEx *IServiceMarketExCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IServiceMarketEx.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0xede8e529.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceMarketEx *IServiceMarketExSession) IsOwner() (bool, error) {
	return _IServiceMarketEx.Contract.IsOwner(&_IServiceMarketEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0xede8e529.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceMarketEx *IServiceMarketExCallerSession) IsOwner() (bool, error) {
	return _IServiceMarketEx.Contract.IsOwner(&_IServiceMarketEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x5089e2c8.
//
// Solidity: function owner() constant returns(address)
func (_IServiceMarketEx *IServiceMarketExCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IServiceMarketEx.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x5089e2c8.
//
// Solidity: function owner() constant returns(address)
func (_IServiceMarketEx *IServiceMarketExSession) Owner() (common.Address, error) {
	return _IServiceMarketEx.Contract.Owner(&_IServiceMarketEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x5089e2c8.
//
// Solidity: function owner() constant returns(address)
func (_IServiceMarketEx *IServiceMarketExCallerSession) Owner() (common.Address, error) {
	return _IServiceMarketEx.Contract.Owner(&_IServiceMarketEx.CallOpts)
}

// ServiceBindingExists is a free data retrieval call binding the contract method 0x27e41750.
//
// Solidity: function serviceBindingExists(string _serviceName) constant returns(bool exist)
func (_IServiceMarketEx *IServiceMarketExCaller) ServiceBindingExists(opts *bind.CallOpts, _serviceName string) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IServiceMarketEx.contract.Call(opts, out, "serviceBindingExists", _serviceName)
	return *ret0, err
}

// ServiceBindingExists is a free data retrieval call binding the contract method 0x27e41750.
//
// Solidity: function serviceBindingExists(string _serviceName) constant returns(bool exist)
func (_IServiceMarketEx *IServiceMarketExSession) ServiceBindingExists(_serviceName string) (bool, error) {
	return _IServiceMarketEx.Contract.ServiceBindingExists(&_IServiceMarketEx.CallOpts, _serviceName)
}

// ServiceBindingExists is a free data retrieval call binding the contract method 0x27e41750.
//
// Solidity: function serviceBindingExists(string _serviceName) constant returns(bool exist)
func (_IServiceMarketEx *IServiceMarketExCallerSession) ServiceBindingExists(_serviceName string) (bool, error) {
	return _IServiceMarketEx.Contract.ServiceBindingExists(&_IServiceMarketEx.CallOpts, _serviceName)
}

// AddServiceBinding is a paid mutator transaction binding the contract method 0x7b6397fa.
//
// Solidity: function addServiceBinding(string _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) AddServiceBinding(opts *bind.TransactOpts, _serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.contract.Transact(opts, "addServiceBinding", _serviceName, _schemas, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExTransactor) AsyncAddServiceBinding(handler func(*types.Receipt, error), opts *bind.TransactOpts, _serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.AsyncTransact(opts, handler, "addServiceBinding", _serviceName, _schemas, _provider, _serviceFee, _qos)
}

// AddServiceBinding is a paid mutator transaction binding the contract method 0x7b6397fa.
//
// Solidity: function addServiceBinding(string _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExSession) AddServiceBinding(_serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.AddServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _schemas, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExSession) AsyncAddServiceBinding(handler func(*types.Receipt, error), _serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncAddServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName, _schemas, _provider, _serviceFee, _qos)
}

// AddServiceBinding is a paid mutator transaction binding the contract method 0x7b6397fa.
//
// Solidity: function addServiceBinding(string _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) AddServiceBinding(_serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.AddServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _schemas, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExTransactorSession) AsyncAddServiceBinding(handler func(*types.Receipt, error), _serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncAddServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName, _schemas, _provider, _serviceFee, _qos)
}

// RemoveServiceBinding is a paid mutator transaction binding the contract method 0xf10b3735.
//
// Solidity: function removeServiceBinding(string _serviceName) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) RemoveServiceBinding(opts *bind.TransactOpts, _serviceName string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.contract.Transact(opts, "removeServiceBinding", _serviceName)
}

func (_IServiceMarketEx *IServiceMarketExTransactor) AsyncRemoveServiceBinding(handler func(*types.Receipt, error), opts *bind.TransactOpts, _serviceName string) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.AsyncTransact(opts, handler, "removeServiceBinding", _serviceName)
}

// RemoveServiceBinding is a paid mutator transaction binding the contract method 0xf10b3735.
//
// Solidity: function removeServiceBinding(string _serviceName) returns()
func (_IServiceMarketEx *IServiceMarketExSession) RemoveServiceBinding(_serviceName string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.RemoveServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName)
}

func (_IServiceMarketEx *IServiceMarketExSession) AsyncRemoveServiceBinding(handler func(*types.Receipt, error), _serviceName string) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncRemoveServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName)
}

// RemoveServiceBinding is a paid mutator transaction binding the contract method 0xf10b3735.
//
// Solidity: function removeServiceBinding(string _serviceName) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) RemoveServiceBinding(_serviceName string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.RemoveServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName)
}

func (_IServiceMarketEx *IServiceMarketExTransactorSession) AsyncRemoveServiceBinding(handler func(*types.Receipt, error), _serviceName string) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncRemoveServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x16cad12a.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_IServiceMarketEx *IServiceMarketExTransactor) AsyncTransferOwnership(handler func(*types.Receipt, error), opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.AsyncTransact(opts, handler, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x16cad12a.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceMarketEx *IServiceMarketExSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.TransferOwnership(&_IServiceMarketEx.TransactOpts, newOwner)
}

func (_IServiceMarketEx *IServiceMarketExSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncTransferOwnership(handler, &_IServiceMarketEx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0x16cad12a.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.TransferOwnership(&_IServiceMarketEx.TransactOpts, newOwner)
}

func (_IServiceMarketEx *IServiceMarketExTransactorSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncTransferOwnership(handler, &_IServiceMarketEx.TransactOpts, newOwner)
}

// UpdateServiceBinding is a paid mutator transaction binding the contract method 0x47b71720.
//
// Solidity: function updateServiceBinding(string _serviceName, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) UpdateServiceBinding(opts *bind.TransactOpts, _serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.contract.Transact(opts, "updateServiceBinding", _serviceName, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExTransactor) AsyncUpdateServiceBinding(handler func(*types.Receipt, error), opts *bind.TransactOpts, _serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.AsyncTransact(opts, handler, "updateServiceBinding", _serviceName, _provider, _serviceFee, _qos)
}

// UpdateServiceBinding is a paid mutator transaction binding the contract method 0x47b71720.
//
// Solidity: function updateServiceBinding(string _serviceName, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExSession) UpdateServiceBinding(_serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.UpdateServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExSession) AsyncUpdateServiceBinding(handler func(*types.Receipt, error), _serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncUpdateServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName, _provider, _serviceFee, _qos)
}

// UpdateServiceBinding is a paid mutator transaction binding the contract method 0x47b71720.
//
// Solidity: function updateServiceBinding(string _serviceName, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) UpdateServiceBinding(_serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.UpdateServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExTransactorSession) AsyncUpdateServiceBinding(handler func(*types.Receipt, error), _serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncUpdateServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName, _provider, _serviceFee, _qos)
}

// IServiceMarketExOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the IServiceMarketEx contract.
type IServiceMarketExOwnershipTransferredIterator struct {
	Event *IServiceMarketExOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *IServiceMarketExOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceMarketExOwnershipTransferred)
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
		it.Event = new(IServiceMarketExOwnershipTransferred)
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
func (it *IServiceMarketExOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceMarketExOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceMarketExOwnershipTransferred represents a OwnershipTransferred event raised by the IServiceMarketEx contract.
type IServiceMarketExOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x000000000000000000000000000000000000000000000000000000005c7c30d4.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceMarketEx *IServiceMarketExFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*IServiceMarketExOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IServiceMarketEx.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &IServiceMarketExOwnershipTransferredIterator{contract: _IServiceMarketEx.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x000000000000000000000000000000000000000000000000000000005c7c30d4.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_IServiceMarketEx *IServiceMarketExFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *IServiceMarketExOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _IServiceMarketEx.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceMarketExOwnershipTransferred)
				if err := _IServiceMarketEx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_IServiceMarketEx *IServiceMarketExFilterer) ParseOwnershipTransferred(log types.Log) (*IServiceMarketExOwnershipTransferred, error) {
	event := new(IServiceMarketExOwnershipTransferred)
	if err := _IServiceMarketEx.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IServiceMarketExServiceBindingAddedIterator is returned from FilterServiceBindingAdded and is used to iterate over the raw logs and unpacked data for ServiceBindingAdded events raised by the IServiceMarketEx contract.
type IServiceMarketExServiceBindingAddedIterator struct {
	Event *IServiceMarketExServiceBindingAdded // Event containing the contract specifics and raw log

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
func (it *IServiceMarketExServiceBindingAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceMarketExServiceBindingAdded)
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
		it.Event = new(IServiceMarketExServiceBindingAdded)
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
func (it *IServiceMarketExServiceBindingAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceMarketExServiceBindingAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceMarketExServiceBindingAdded represents a ServiceBindingAdded event raised by the IServiceMarketEx contract.
type IServiceMarketExServiceBindingAdded struct {
	ServiceName common.Hash
	Schemas     string
	Provider    string
	ServiceFee  string
	Qos         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceBindingAdded is a free log retrieval operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000021bc4ee.
//
// Solidity: event ServiceBindingAdded(string indexed _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos)
func (_IServiceMarketEx *IServiceMarketExFilterer) FilterServiceBindingAdded(opts *bind.FilterOpts, _serviceName []string) (*IServiceMarketExServiceBindingAddedIterator, error) {

	var _serviceNameRule []interface{}
	for _, _serviceNameItem := range _serviceName {
		_serviceNameRule = append(_serviceNameRule, _serviceNameItem)
	}

	logs, sub, err := _IServiceMarketEx.contract.FilterLogs(opts, "ServiceBindingAdded", _serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &IServiceMarketExServiceBindingAddedIterator{contract: _IServiceMarketEx.contract, event: "ServiceBindingAdded", logs: logs, sub: sub}, nil
}

// WatchServiceBindingAdded is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000021bc4ee.
//
// Solidity: event ServiceBindingAdded(string indexed _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos)
func (_IServiceMarketEx *IServiceMarketExFilterer) WatchServiceBindingAdded(opts *bind.WatchOpts, sink chan<- *IServiceMarketExServiceBindingAdded, _serviceName []string) (event.Subscription, error) {

	var _serviceNameRule []interface{}
	for _, _serviceNameItem := range _serviceName {
		_serviceNameRule = append(_serviceNameRule, _serviceNameItem)
	}

	logs, sub, err := _IServiceMarketEx.contract.WatchLogs(opts, "ServiceBindingAdded", _serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceMarketExServiceBindingAdded)
				if err := _IServiceMarketEx.contract.UnpackLog(event, "ServiceBindingAdded", log); err != nil {
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

// ParseServiceBindingAdded is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000021bc4ee.
//
// Solidity: event ServiceBindingAdded(string indexed _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos)
func (_IServiceMarketEx *IServiceMarketExFilterer) ParseServiceBindingAdded(log types.Log) (*IServiceMarketExServiceBindingAdded, error) {
	event := new(IServiceMarketExServiceBindingAdded)
	if err := _IServiceMarketEx.contract.UnpackLog(event, "ServiceBindingAdded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IServiceMarketExServiceBindingRemovedIterator is returned from FilterServiceBindingRemoved and is used to iterate over the raw logs and unpacked data for ServiceBindingRemoved events raised by the IServiceMarketEx contract.
type IServiceMarketExServiceBindingRemovedIterator struct {
	Event *IServiceMarketExServiceBindingRemoved // Event containing the contract specifics and raw log

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
func (it *IServiceMarketExServiceBindingRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceMarketExServiceBindingRemoved)
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
		it.Event = new(IServiceMarketExServiceBindingRemoved)
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
func (it *IServiceMarketExServiceBindingRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceMarketExServiceBindingRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceMarketExServiceBindingRemoved represents a ServiceBindingRemoved event raised by the IServiceMarketEx contract.
type IServiceMarketExServiceBindingRemoved struct {
	ServiceName common.Hash
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceBindingRemoved is a free log retrieval operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000667378d5.
//
// Solidity: event ServiceBindingRemoved(string indexed _serviceName)
func (_IServiceMarketEx *IServiceMarketExFilterer) FilterServiceBindingRemoved(opts *bind.FilterOpts, _serviceName []string) (*IServiceMarketExServiceBindingRemovedIterator, error) {

	var _serviceNameRule []interface{}
	for _, _serviceNameItem := range _serviceName {
		_serviceNameRule = append(_serviceNameRule, _serviceNameItem)
	}

	logs, sub, err := _IServiceMarketEx.contract.FilterLogs(opts, "ServiceBindingRemoved", _serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &IServiceMarketExServiceBindingRemovedIterator{contract: _IServiceMarketEx.contract, event: "ServiceBindingRemoved", logs: logs, sub: sub}, nil
}

// WatchServiceBindingRemoved is a free log subscription operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000667378d5.
//
// Solidity: event ServiceBindingRemoved(string indexed _serviceName)
func (_IServiceMarketEx *IServiceMarketExFilterer) WatchServiceBindingRemoved(opts *bind.WatchOpts, sink chan<- *IServiceMarketExServiceBindingRemoved, _serviceName []string) (event.Subscription, error) {

	var _serviceNameRule []interface{}
	for _, _serviceNameItem := range _serviceName {
		_serviceNameRule = append(_serviceNameRule, _serviceNameItem)
	}

	logs, sub, err := _IServiceMarketEx.contract.WatchLogs(opts, "ServiceBindingRemoved", _serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceMarketExServiceBindingRemoved)
				if err := _IServiceMarketEx.contract.UnpackLog(event, "ServiceBindingRemoved", log); err != nil {
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

// ParseServiceBindingRemoved is a log parse operation binding the contract event 0x00000000000000000000000000000000000000000000000000000000667378d5.
//
// Solidity: event ServiceBindingRemoved(string indexed _serviceName)
func (_IServiceMarketEx *IServiceMarketExFilterer) ParseServiceBindingRemoved(log types.Log) (*IServiceMarketExServiceBindingRemoved, error) {
	event := new(IServiceMarketExServiceBindingRemoved)
	if err := _IServiceMarketEx.contract.UnpackLog(event, "ServiceBindingRemoved", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IServiceMarketExServiceBindingUpdatedIterator is returned from FilterServiceBindingUpdated and is used to iterate over the raw logs and unpacked data for ServiceBindingUpdated events raised by the IServiceMarketEx contract.
type IServiceMarketExServiceBindingUpdatedIterator struct {
	Event *IServiceMarketExServiceBindingUpdated // Event containing the contract specifics and raw log

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
func (it *IServiceMarketExServiceBindingUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceMarketExServiceBindingUpdated)
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
		it.Event = new(IServiceMarketExServiceBindingUpdated)
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
func (it *IServiceMarketExServiceBindingUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceMarketExServiceBindingUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceMarketExServiceBindingUpdated represents a ServiceBindingUpdated event raised by the IServiceMarketEx contract.
type IServiceMarketExServiceBindingUpdated struct {
	ServiceName common.Hash
	Provider    string
	ServiceFee  string
	Qos         *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceBindingUpdated is a free log retrieval operation binding the contract event 0x000000000000000000000000000000000000000000000000000000007a7669f9.
//
// Solidity: event ServiceBindingUpdated(string indexed _serviceName, string _provider, string _serviceFee, uint256 _qos)
func (_IServiceMarketEx *IServiceMarketExFilterer) FilterServiceBindingUpdated(opts *bind.FilterOpts, _serviceName []string) (*IServiceMarketExServiceBindingUpdatedIterator, error) {

	var _serviceNameRule []interface{}
	for _, _serviceNameItem := range _serviceName {
		_serviceNameRule = append(_serviceNameRule, _serviceNameItem)
	}

	logs, sub, err := _IServiceMarketEx.contract.FilterLogs(opts, "ServiceBindingUpdated", _serviceNameRule)
	if err != nil {
		return nil, err
	}
	return &IServiceMarketExServiceBindingUpdatedIterator{contract: _IServiceMarketEx.contract, event: "ServiceBindingUpdated", logs: logs, sub: sub}, nil
}

// WatchServiceBindingUpdated is a free log subscription operation binding the contract event 0x000000000000000000000000000000000000000000000000000000007a7669f9.
//
// Solidity: event ServiceBindingUpdated(string indexed _serviceName, string _provider, string _serviceFee, uint256 _qos)
func (_IServiceMarketEx *IServiceMarketExFilterer) WatchServiceBindingUpdated(opts *bind.WatchOpts, sink chan<- *IServiceMarketExServiceBindingUpdated, _serviceName []string) (event.Subscription, error) {

	var _serviceNameRule []interface{}
	for _, _serviceNameItem := range _serviceName {
		_serviceNameRule = append(_serviceNameRule, _serviceNameItem)
	}

	logs, sub, err := _IServiceMarketEx.contract.WatchLogs(opts, "ServiceBindingUpdated", _serviceNameRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceMarketExServiceBindingUpdated)
				if err := _IServiceMarketEx.contract.UnpackLog(event, "ServiceBindingUpdated", log); err != nil {
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

// ParseServiceBindingUpdated is a log parse operation binding the contract event 0x000000000000000000000000000000000000000000000000000000007a7669f9.
//
// Solidity: event ServiceBindingUpdated(string indexed _serviceName, string _provider, string _serviceFee, uint256 _qos)
func (_IServiceMarketEx *IServiceMarketExFilterer) ParseServiceBindingUpdated(log types.Log) (*IServiceMarketExServiceBindingUpdated, error) {
	event := new(IServiceMarketExServiceBindingUpdated)
	if err := _IServiceMarketEx.contract.UnpackLog(event, "ServiceBindingUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
