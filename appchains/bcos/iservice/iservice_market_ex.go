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

// iServiceMarketExServiceBinding is an auto generated low-level Go binding around an user-defined struct.
type iServiceMarketExServiceBinding struct {
	ServiceName string
	Schemas     string
	Provider    string
	ServiceFee  string
	Qos         *big.Int
}

// IServiceMarketExABI is the input ABI used to generate the binding from.
const IServiceMarketExABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_schemas\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_provider\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_serviceFee\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"ServiceBindingAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"ServiceBindingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_provider\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_serviceFee\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"ServiceBindingUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"GetServiceBinding\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"schemas\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"provider\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceFee\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"qos\",\"type\":\"uint256\"}],\"internalType\":\"structiServiceMarketEx.ServiceBinding\",\"name\":\"binding\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"GetServiceBindingByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"schemas\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"provider\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceFee\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"qos\",\"type\":\"uint256\"}],\"internalType\":\"structiServiceMarketEx.ServiceBinding\",\"name\":\"binding\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetServiceBindingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"count\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_schemas\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_provider\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_serviceFee\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"addServiceBinding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"removeServiceBinding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_provider\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_serviceFee\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"updateServiceBinding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IServiceMarketExBin is the compiled bytecode used for deploying new contracts.
var IServiceMarketExBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f5c7c30d4a0f08950cb23be4132957b357fa5dfdb0fcf218f81b86a1c036e47d060405160405180910390a3612245806100dc6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c80638fd4d350116100665780638fd4d3501461010a578063c6a9fbec14610128578063cbb950e414610158578063ede8e52914610188578063f10b3735146101a657610093565b806316cad12a1461009857806347b71720146100b45780635089e2c8146100d05780637b6397fa146100ee575b600080fd5b6100b260048036038101906100ad91906116aa565b6101c2565b005b6100ce60048036038101906100c99190611800565b610215565b005b6100d8610408565b6040516100e59190611e7f565b60405180910390f35b61010860048036038101906101039190611718565b610431565b005b6101126108f5565b60405161011f919061209a565b60405180910390f35b610142600480360381019061013d91906118b9565b610902565b60405161014f9190612078565b60405180910390f35b610172600480360381019061016d91906118fa565b610c3a565b60405161017f9190612078565b60405180910390f35b610190610f17565b60405161019d9190611e9a565b60405180910390f35b6101c060048036038101906101bb91906116d3565b610f6e565b005b6101ca610f17565b610209576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161020090612038565b60405180910390fd5b61021281611179565b50565b61021d610f17565b61025c576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161025390612038565b60405180910390fd5b86868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506002816040516102b19190611e68565b908152602001604051809103902060010160009054906101000a900460ff1661030f576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161030690611ff8565b60405180910390fd5b6000600160028a8a604051610325929190611e4f565b9081526020016040518091039020600001548154811061034157fe5b90600052602060002090600502019050600087879050111561037257868682600201919061037092919061142a565b505b600085859050111561039357848482600301919061039192919061142a565b505b60008311156103a6578281600401819055505b88886040516103b6929190611e4f565b60405180910390207f7a7669f9ae0472fda70561457e748c4df8b930fa1cf1457f5469e1ba3b8b3d4288888888886040516103f5959493929190611eb5565b60405180910390a2505050505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610439610f17565b610478576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161046f90612038565b60405180910390fd5b88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508460008551116105d1576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016105c890611fd8565b60405180910390fd5b6000845111610615576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161060c90611f78565b60405180910390fd5b6000835111610659576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161065090611f58565b60405180910390fd5b600082511161069d576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161069490611fb8565b60405180910390fd5b600081116106e0576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016106d790612018565b60405180910390fd5b8d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506002816040516107359190611e68565b908152602001604051809103902060010160009054906101000a900460ff1615610794576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161078b90612058565b60405180910390fd5b61079c6114aa565b8f8f8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081600001819052508d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081602001819052508b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050816040018190525089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508160600181905250878160800181815250506108e3816112a7565b50505050505050505050505050505050565b6000600180549050905090565b61090a6114aa565b6109126114d9565b6002836040516109229190611e68565b9081526020016040518091039020604051806040016040529081600082015481526020016001820160009054906101000a900460ff1615151515815250509050806020015115610c3057600181600001518154811061097d57fe5b90600052602060002090600502016040518060a0016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a2f5780601f10610a0457610100808354040283529160200191610a2f565b820191906000526020600020905b815481529060010190602001808311610a1257829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ad15780601f10610aa657610100808354040283529160200191610ad1565b820191906000526020600020905b815481529060010190602001808311610ab457829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b735780601f10610b4857610100808354040283529160200191610b73565b820191906000526020600020905b815481529060010190602001808311610b5657829003601f168201915b50505050508152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c155780601f10610bea57610100808354040283529160200191610c15565b820191906000526020600020905b815481529060010190602001808311610bf857829003601f168201915b50505050508152602001600482015481525050915050610c35565b819150505b919050565b610c426114aa565b600180549050821015610f0e5760018281548110610c5c57fe5b90600052602060002090600502016040518060a0016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d0e5780601f10610ce357610100808354040283529160200191610d0e565b820191906000526020600020905b815481529060010190602001808311610cf157829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610db05780601f10610d8557610100808354040283529160200191610db0565b820191906000526020600020905b815481529060010190602001808311610d9357829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e525780601f10610e2757610100808354040283529160200191610e52565b820191906000526020600020905b815481529060010190602001808311610e3557829003601f168201915b50505050508152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610ef45780601f10610ec957610100808354040283529160200191610ef4565b820191906000526020600020905b815481529060010190602001808311610ed757829003601f168201915b505050505081526020016004820154815250509050610f12565b8090505b919050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b610f76610f17565b610fb5576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610fac90612038565b60405180910390fd5b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505060028160405161100a9190611e68565b908152602001604051809103902060010160009054906101000a900460ff16611068576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161105f90611ff8565b60405180910390fd5b60016002848460405161107c929190611e4f565b9081526020016040518091039020600001548154811061109857fe5b9060005260206000209060050201600080820160006110b791906114f5565b6001820160006110c791906114f5565b6002820160006110d791906114f5565b6003820160006110e791906114f5565b6004820160009055505060028383604051611103929190611e4f565b90815260200160405180910390206000808201600090556001820160006101000a81549060ff021916905550508282604051611140929190611e4f565b60405180910390207f667378d5061e84f0e8940230e3022c567aa1d8c2e425ee81aace26e7d8d6eaa460405160405180910390a2505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156111e9576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016111e090611f98565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f5c7c30d4a0f08950cb23be4132957b357fa5dfdb0fcf218f81b86a1c036e47d060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600180549050600282600001516040516112c19190611e68565b9081526020016040518091039020600001819055506001600282600001516040516112ec9190611e68565b908152602001604051809103902060010160006101000a81548160ff0219169083151502179055506001819080600181540180825580915050600190039060005260206000209060050201600090919091909150600082015181600001908051906020019061135c92919061153d565b50602082015181600101908051906020019061137992919061153d565b50604082015181600201908051906020019061139692919061153d565b5060608201518160030190805190602001906113b392919061153d565b5060808201518160040155505080600001516040516113d29190611e68565b60405180910390207f3a092fad231e05e29f7061a78203ceb4110c59c67d0a2f0d1b4ccb2fe25ecc90826020015183604001518460600151856080015160405161141f9493929190611efe565b60405180910390a250565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061146b57803560ff1916838001178555611499565b82800160010185558215611499579182015b8281111561149857823582559160200191906001019061147d565b5b5090506114a691906115bd565b5090565b6040518060a0016040528060608152602001606081526020016060815260200160608152602001600081525090565b6040518060400160405280600081526020016000151581525090565b50805460018160011615610100020316600290046000825580601f1061151b575061153a565b601f01602090049060005260206000209081019061153991906115bd565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061157e57805160ff19168380011785556115ac565b828001600101855582156115ac579182015b828111156115ab578251825591602001919060010190611590565b5b5090506115b991906115bd565b5090565b6115df91905b808211156115db5760008160009055506001016115c3565b5090565b90565b6000813590506115f1816121e1565b92915050565b60008083601f84011261160957600080fd5b8235905067ffffffffffffffff81111561162257600080fd5b60208301915083600182028301111561163a57600080fd5b9250929050565b600082601f83011261165257600080fd5b8135611665611660826120e2565b6120b5565b9150808252602083016020830185838301111561168157600080fd5b61168c83828461218e565b50505092915050565b6000813590506116a4816121f8565b92915050565b6000602082840312156116bc57600080fd5b60006116ca848285016115e2565b91505092915050565b600080602083850312156116e657600080fd5b600083013567ffffffffffffffff81111561170057600080fd5b61170c858286016115f7565b92509250509250929050565b600080600080600080600080600060a08a8c03121561173657600080fd5b60008a013567ffffffffffffffff81111561175057600080fd5b61175c8c828d016115f7565b995099505060208a013567ffffffffffffffff81111561177b57600080fd5b6117878c828d016115f7565b975097505060408a013567ffffffffffffffff8111156117a657600080fd5b6117b28c828d016115f7565b955095505060608a013567ffffffffffffffff8111156117d157600080fd5b6117dd8c828d016115f7565b935093505060806117f08c828d01611695565b9150509295985092959850929598565b60008060008060008060006080888a03121561181b57600080fd5b600088013567ffffffffffffffff81111561183557600080fd5b6118418a828b016115f7565b9750975050602088013567ffffffffffffffff81111561186057600080fd5b61186c8a828b016115f7565b9550955050604088013567ffffffffffffffff81111561188b57600080fd5b6118978a828b016115f7565b935093505060606118aa8a828b01611695565b91505092959891949750929550565b6000602082840312156118cb57600080fd5b600082013567ffffffffffffffff8111156118e557600080fd5b6118f184828501611641565b91505092915050565b60006020828403121561190c57600080fd5b600061191a84828501611695565b91505092915050565b61192c81612146565b82525050565b61193b81612158565b82525050565b600061194d838561212a565b935061195a83858461218e565b611963836121d0565b840190509392505050565b600061197a838561213b565b935061198783858461218e565b82840190509392505050565b600061199e8261210e565b6119a88185612119565b93506119b881856020860161219d565b6119c1816121d0565b840191505092915050565b60006119d78261210e565b6119e1818561212a565b93506119f181856020860161219d565b6119fa816121d0565b840191505092915050565b6000611a108261210e565b611a1a818561213b565b9350611a2a81856020860161219d565b80840191505092915050565b6000611a4360338361212a565b91507f69536572766963654d61726b657445783a20736572766963652070726f76696460008301527f65722063616e206e6f7420626520656d707479000000000000000000000000006020830152604082019050919050565b6000611aa960328361212a565b91507f69536572766963654d61726b657445783a207365727669636520736368656d6160008301527f732063616e206e6f7420626520656d70747900000000000000000000000000006020830152604082019050919050565b6000611b0f60268361212a565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611b75602e8361212a565b91507f69536572766963654d61726b657445783a20736572766963652066656520636160008301527f6e206e6f7420626520656d7074790000000000000000000000000000000000006020830152604082019050919050565b6000611bdb602f8361212a565b91507f69536572766963654d61726b657445783a2073657276696365206e616d65206360008301527f616e206e6f7420626520656d70747900000000000000000000000000000000006020830152604082019050919050565b6000611c4160308361212a565b91507f69536572766963654d61726b657445783a20736572766963652062696e64696e60008301527f6720646f6573206e6f74206578697374000000000000000000000000000000006020830152604082019050919050565b6000611ca7602e8361212a565b91507f69536572766963654d61726b657445783a20716f732073686f756c642062652060008301527f67726561746572207468616e20300000000000000000000000000000000000006020830152604082019050919050565b6000611d0d60208361212a565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000611d4d60308361212a565b91507f69536572766963654d61726b657445783a20736572766963652062696e64696e60008301527f6720616c726561647920657869737473000000000000000000000000000000006020830152604082019050919050565b600060a0830160008301518482036000860152611dc38282611993565b91505060208301518482036020860152611ddd8282611993565b91505060408301518482036040860152611df78282611993565b91505060608301518482036060860152611e118282611993565b9150506080830151611e266080860182611e31565b508091505092915050565b611e3a81612184565b82525050565b611e4981612184565b82525050565b6000611e5c82848661196e565b91508190509392505050565b6000611e748284611a05565b915081905092915050565b6000602082019050611e946000830184611923565b92915050565b6000602082019050611eaf6000830184611932565b92915050565b60006060820190508181036000830152611ed0818789611941565b90508181036020830152611ee5818587611941565b9050611ef46040830184611e40565b9695505050505050565b60006080820190508181036000830152611f1881876119cc565b90508181036020830152611f2c81866119cc565b90508181036040830152611f4081856119cc565b9050611f4f6060830184611e40565b95945050505050565b60006020820190508181036000830152611f7181611a36565b9050919050565b60006020820190508181036000830152611f9181611a9c565b9050919050565b60006020820190508181036000830152611fb181611b02565b9050919050565b60006020820190508181036000830152611fd181611b68565b9050919050565b60006020820190508181036000830152611ff181611bce565b9050919050565b6000602082019050818103600083015261201181611c34565b9050919050565b6000602082019050818103600083015261203181611c9a565b9050919050565b6000602082019050818103600083015261205181611d00565b9050919050565b6000602082019050818103600083015261207181611d40565b9050919050565b600060208201905081810360008301526120928184611da6565b905092915050565b60006020820190506120af6000830184611e40565b92915050565b6000604051905081810181811067ffffffffffffffff821117156120d857600080fd5b8060405250919050565b600067ffffffffffffffff8211156120f957600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061215182612164565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156121bb5780820151818401526020810190506121a0565b838111156121ca576000848401525b50505050565b6000601f19601f8301169050919050565b6121ea81612146565b81146121f557600080fd5b50565b61220181612184565b811461220c57600080fd5b5056fea264697066735822122077893a19c39f3c64dfb6081c1b46fdefbd8ffe7d4bb760892e3c111a662b6c2a64736f6c634300060a0033"

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

// GetServiceBinding is a free data retrieval call binding the contract method 0xc6a9fbec.
//
// Solidity: function GetServiceBinding(string _serviceName) constant returns(iServiceMarketExServiceBinding binding)
func (_IServiceMarketEx *IServiceMarketExCaller) GetServiceBinding(opts *bind.CallOpts, _serviceName string) (iServiceMarketExServiceBinding, error) {
	var (
		ret0 = new(iServiceMarketExServiceBinding)
	)
	out := ret0
	err := _IServiceMarketEx.contract.Call(opts, out, "GetServiceBinding", _serviceName)
	return *ret0, err
}

// GetServiceBinding is a free data retrieval call binding the contract method 0xc6a9fbec.
//
// Solidity: function GetServiceBinding(string _serviceName) constant returns(iServiceMarketExServiceBinding binding)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBinding(_serviceName string) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBinding(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceBinding is a free data retrieval call binding the contract method 0xc6a9fbec.
//
// Solidity: function GetServiceBinding(string _serviceName) constant returns(iServiceMarketExServiceBinding binding)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBinding(_serviceName string) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBinding(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceBindingByIndex is a free data retrieval call binding the contract method 0xcbb950e4.
//
// Solidity: function GetServiceBindingByIndex(uint256 _index) constant returns(iServiceMarketExServiceBinding binding)
func (_IServiceMarketEx *IServiceMarketExCaller) GetServiceBindingByIndex(opts *bind.CallOpts, _index *big.Int) (iServiceMarketExServiceBinding, error) {
	var (
		ret0 = new(iServiceMarketExServiceBinding)
	)
	out := ret0
	err := _IServiceMarketEx.contract.Call(opts, out, "GetServiceBindingByIndex", _index)
	return *ret0, err
}

// GetServiceBindingByIndex is a free data retrieval call binding the contract method 0xcbb950e4.
//
// Solidity: function GetServiceBindingByIndex(uint256 _index) constant returns(iServiceMarketExServiceBinding binding)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBindingByIndex(_index *big.Int) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingByIndex(&_IServiceMarketEx.CallOpts, _index)
}

// GetServiceBindingByIndex is a free data retrieval call binding the contract method 0xcbb950e4.
//
// Solidity: function GetServiceBindingByIndex(uint256 _index) constant returns(iServiceMarketExServiceBinding binding)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBindingByIndex(_index *big.Int) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingByIndex(&_IServiceMarketEx.CallOpts, _index)
}

// GetServiceBindingCount is a free data retrieval call binding the contract method 0x8fd4d350.
//
// Solidity: function GetServiceBindingCount() constant returns(uint256 count)
func (_IServiceMarketEx *IServiceMarketExCaller) GetServiceBindingCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IServiceMarketEx.contract.Call(opts, out, "GetServiceBindingCount")
	return *ret0, err
}

// GetServiceBindingCount is a free data retrieval call binding the contract method 0x8fd4d350.
//
// Solidity: function GetServiceBindingCount() constant returns(uint256 count)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBindingCount() (*big.Int, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingCount(&_IServiceMarketEx.CallOpts)
}

// GetServiceBindingCount is a free data retrieval call binding the contract method 0x8fd4d350.
//
// Solidity: function GetServiceBindingCount() constant returns(uint256 count)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBindingCount() (*big.Int, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingCount(&_IServiceMarketEx.CallOpts)
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
