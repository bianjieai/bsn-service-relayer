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
var IServiceMarketExBin = "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3612245806100dc6000396000f3fe608060405234801561001057600080fd5b50600436106100935760003560e01c8063a7c6396911610066578063a7c639691461010c578063bce9f63f1461013c578063bd36df1614610158578063f2fde38b14610188578063f8bdb54a146101a457610093565b80636e29ea60146100985780638da5cb5b146100b45780638f32d59b146100d257806392ce31c4146100f0575b600080fd5b6100b260048036038101906100ad9190611800565b6101c2565b005b6100bc6103b5565b6040516100c99190611e7f565b60405180910390f35b6100da6103de565b6040516100e79190611e9a565b60405180910390f35b61010a60048036038101906101059190611718565b610435565b005b610126600480360381019061012191906118fa565b6108f9565b6040516101339190612078565b60405180910390f35b610156600480360381019061015191906116d3565b610bd6565b005b610172600480360381019061016d91906118b9565b610de1565b60405161017f9190612078565b60405180910390f35b6101a2600480360381019061019d91906116aa565b611119565b005b6101ac61116c565b6040516101b9919061209a565b60405180910390f35b6101ca6103de565b610209576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161020090611fd8565b60405180910390fd5b86868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505060028160405161025e9190611e68565b908152602001604051809103902060010160009054906101000a900460ff166102bc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102b390611fb8565b60405180910390fd5b6000600160028a8a6040516102d2929190611e4f565b908152602001604051809103902060000154815481106102ee57fe5b90600052602060002090600502019050600087879050111561031f57868682600201919061031d92919061142a565b505b600085859050111561034057848482600301919061033e92919061142a565b505b6000831115610353578281600401819055505b8888604051610363929190611e4f565b60405180910390207fac059a9f400ca20e9086f7e73b998697778b9a48fe40d2392bf649752c5538b288888888886040516103a2959493929190611eb5565b60405180910390a2505050505050505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b61043d6103de565b61047c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161047390611fd8565b60405180910390fd5b88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505085858080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508460008551116105d5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105cc90612038565b60405180910390fd5b6000845111610619576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161061090611f98565b60405180910390fd5b600083511161065d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161065490612058565b60405180910390fd5b60008251116106a1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161069890611f78565b60405180910390fd5b600081116106e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106db90612018565b60405180910390fd5b8d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050506002816040516107399190611e68565b908152602001604051809103902060010160009054906101000a900460ff1615610798576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161078f90611ff8565b60405180910390fd5b6107a06114aa565b8f8f8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081600001819052508d8d8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081602001819052508b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050816040018190525089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508160600181905250878160800181815250506108e781611179565b50505050505050505050505050505050565b6109016114aa565b600180549050821015610bcd576001828154811061091b57fe5b90600052602060002090600502016040518060a0016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156109cd5780601f106109a2576101008083540402835291602001916109cd565b820191906000526020600020905b8154815290600101906020018083116109b057829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a6f5780601f10610a4457610100808354040283529160200191610a6f565b820191906000526020600020905b815481529060010190602001808311610a5257829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b115780601f10610ae657610100808354040283529160200191610b11565b820191906000526020600020905b815481529060010190602001808311610af457829003601f168201915b50505050508152602001600382018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610bb35780601f10610b8857610100808354040283529160200191610bb3565b820191906000526020600020905b815481529060010190602001808311610b9657829003601f168201915b505050505081526020016004820154815250509050610bd1565b8090505b919050565b610bde6103de565b610c1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c1490611fd8565b60405180910390fd5b81818080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050600281604051610c729190611e68565b908152602001604051809103902060010160009054906101000a900460ff16610cd0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cc790611fb8565b60405180910390fd5b600160028484604051610ce4929190611e4f565b90815260200160405180910390206000015481548110610d0057fe5b906000526020600020906005020160008082016000610d1f91906114d9565b600182016000610d2f91906114d9565b600282016000610d3f91906114d9565b600382016000610d4f91906114d9565b6004820160009055505060028383604051610d6b929190611e4f565b90815260200160405180910390206000808201600090556001820160006101000a81549060ff021916905550508282604051610da8929190611e4f565b60405180910390207f045edfc0de1510eef86edd984a529791dd2924769af547ab858d860139c7a54d60405160405180910390a2505050565b610de96114aa565b610df1611521565b600283604051610e019190611e68565b9081526020016040518091039020604051806040016040529081600082015481526020016001820160009054906101000a900460ff161515151581525050905080602001511561110f576001816000015181548110610e5c57fe5b90600052602060002090600502016040518060a0016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610f0e5780601f10610ee357610100808354040283529160200191610f0e565b820191906000526020600020905b815481529060010190602001808311610ef157829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fb05780601f10610f8557610100808354040283529160200191610fb0565b820191906000526020600020905b815481529060010190602001808311610f9357829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110525780601f1061102757610100808354040283529160200191611052565b820191906000526020600020905b81548152906001019060200180831161103557829003601f168201915b50505050508152602001600382018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110f45780601f106110c9576101008083540402835291602001916110f4565b820191906000526020600020905b8154815290600101906020018083116110d757829003601f168201915b50505050508152602001600482015481525050915050611114565b819150505b919050565b6111216103de565b611160576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115790611fd8565b60405180910390fd5b611169816112fc565b50565b6000600180549050905090565b600180549050600282600001516040516111939190611e68565b9081526020016040518091039020600001819055506001600282600001516040516111be9190611e68565b908152602001604051809103902060010160006101000a81548160ff0219169083151502179055506001819080600181540180825580915050600190039060005260206000209060050201600090919091909150600082015181600001908051906020019061122e92919061153d565b50602082015181600101908051906020019061124b92919061153d565b50604082015181600201908051906020019061126892919061153d565b50606082015181600301908051906020019061128592919061153d565b5060808201518160040155505080600001516040516112a49190611e68565b60405180910390207f365f7e387fe001fa02d1c008ec11e2c21f47c6a2102f497b300a17c83127dc6982602001518360400151846060015185608001516040516112f19493929190611efe565b60405180910390a250565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141561136c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161136390611f58565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061146b57803560ff1916838001178555611499565b82800160010185558215611499579182015b8281111561149857823582559160200191906001019061147d565b5b5090506114a691906115bd565b5090565b6040518060a0016040528060608152602001606081526020016060815260200160608152602001600081525090565b50805460018160011615610100020316600290046000825580601f106114ff575061151e565b601f01602090049060005260206000209081019061151d91906115bd565b5b50565b6040518060400160405280600081526020016000151581525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061157e57805160ff19168380011785556115ac565b828001600101855582156115ac579182015b828111156115ab578251825591602001919060010190611590565b5b5090506115b991906115bd565b5090565b6115df91905b808211156115db5760008160009055506001016115c3565b5090565b90565b6000813590506115f1816121e1565b92915050565b60008083601f84011261160957600080fd5b8235905067ffffffffffffffff81111561162257600080fd5b60208301915083600182028301111561163a57600080fd5b9250929050565b600082601f83011261165257600080fd5b8135611665611660826120e2565b6120b5565b9150808252602083016020830185838301111561168157600080fd5b61168c83828461218e565b50505092915050565b6000813590506116a4816121f8565b92915050565b6000602082840312156116bc57600080fd5b60006116ca848285016115e2565b91505092915050565b600080602083850312156116e657600080fd5b600083013567ffffffffffffffff81111561170057600080fd5b61170c858286016115f7565b92509250509250929050565b600080600080600080600080600060a08a8c03121561173657600080fd5b60008a013567ffffffffffffffff81111561175057600080fd5b61175c8c828d016115f7565b995099505060208a013567ffffffffffffffff81111561177b57600080fd5b6117878c828d016115f7565b975097505060408a013567ffffffffffffffff8111156117a657600080fd5b6117b28c828d016115f7565b955095505060608a013567ffffffffffffffff8111156117d157600080fd5b6117dd8c828d016115f7565b935093505060806117f08c828d01611695565b9150509295985092959850929598565b60008060008060008060006080888a03121561181b57600080fd5b600088013567ffffffffffffffff81111561183557600080fd5b6118418a828b016115f7565b9750975050602088013567ffffffffffffffff81111561186057600080fd5b61186c8a828b016115f7565b9550955050604088013567ffffffffffffffff81111561188b57600080fd5b6118978a828b016115f7565b935093505060606118aa8a828b01611695565b91505092959891949750929550565b6000602082840312156118cb57600080fd5b600082013567ffffffffffffffff8111156118e557600080fd5b6118f184828501611641565b91505092915050565b60006020828403121561190c57600080fd5b600061191a84828501611695565b91505092915050565b61192c81612146565b82525050565b61193b81612158565b82525050565b600061194d838561212a565b935061195a83858461218e565b611963836121d0565b840190509392505050565b600061197a838561213b565b935061198783858461218e565b82840190509392505050565b600061199e8261210e565b6119a88185612119565b93506119b881856020860161219d565b6119c1816121d0565b840191505092915050565b60006119d78261210e565b6119e1818561212a565b93506119f181856020860161219d565b6119fa816121d0565b840191505092915050565b6000611a108261210e565b611a1a818561213b565b9350611a2a81856020860161219d565b80840191505092915050565b6000611a4360268361212a565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611aa9602e8361212a565b91507f69536572766963654d61726b657445783a20736572766963652066656520636160008301527f6e206e6f7420626520656d7074790000000000000000000000000000000000006020830152604082019050919050565b6000611b0f60328361212a565b91507f69536572766963654d61726b657445783a207365727669636520736368656d6160008301527f732063616e206e6f7420626520656d70747900000000000000000000000000006020830152604082019050919050565b6000611b7560308361212a565b91507f69536572766963654d61726b657445783a20736572766963652062696e64696e60008301527f6720646f6573206e6f74206578697374000000000000000000000000000000006020830152604082019050919050565b6000611bdb60208361212a565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000611c1b60308361212a565b91507f69536572766963654d61726b657445783a20736572766963652062696e64696e60008301527f6720616c726561647920657869737473000000000000000000000000000000006020830152604082019050919050565b6000611c81602e8361212a565b91507f69536572766963654d61726b657445783a20716f732073686f756c642062652060008301527f67726561746572207468616e20300000000000000000000000000000000000006020830152604082019050919050565b6000611ce7602f8361212a565b91507f69536572766963654d61726b657445783a2073657276696365206e616d65206360008301527f616e206e6f7420626520656d70747900000000000000000000000000000000006020830152604082019050919050565b6000611d4d60338361212a565b91507f69536572766963654d61726b657445783a20736572766963652070726f76696460008301527f65722063616e206e6f7420626520656d707479000000000000000000000000006020830152604082019050919050565b600060a0830160008301518482036000860152611dc38282611993565b91505060208301518482036020860152611ddd8282611993565b91505060408301518482036040860152611df78282611993565b91505060608301518482036060860152611e118282611993565b9150506080830151611e266080860182611e31565b508091505092915050565b611e3a81612184565b82525050565b611e4981612184565b82525050565b6000611e5c82848661196e565b91508190509392505050565b6000611e748284611a05565b915081905092915050565b6000602082019050611e946000830184611923565b92915050565b6000602082019050611eaf6000830184611932565b92915050565b60006060820190508181036000830152611ed0818789611941565b90508181036020830152611ee5818587611941565b9050611ef46040830184611e40565b9695505050505050565b60006080820190508181036000830152611f1881876119cc565b90508181036020830152611f2c81866119cc565b90508181036040830152611f4081856119cc565b9050611f4f6060830184611e40565b95945050505050565b60006020820190508181036000830152611f7181611a36565b9050919050565b60006020820190508181036000830152611f9181611a9c565b9050919050565b60006020820190508181036000830152611fb181611b02565b9050919050565b60006020820190508181036000830152611fd181611b68565b9050919050565b60006020820190508181036000830152611ff181611bce565b9050919050565b6000602082019050818103600083015261201181611c0e565b9050919050565b6000602082019050818103600083015261203181611c74565b9050919050565b6000602082019050818103600083015261205181611cda565b9050919050565b6000602082019050818103600083015261207181611d40565b9050919050565b600060208201905081810360008301526120928184611da6565b905092915050565b60006020820190506120af6000830184611e40565b92915050565b6000604051905081810181811067ffffffffffffffff821117156120d857600080fd5b8060405250919050565b600067ffffffffffffffff8211156120f957600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600082825260208201905092915050565b600081905092915050565b600061215182612164565b9050919050565b60008115159050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b82818337600083830152505050565b60005b838110156121bb5780820151818401526020810190506121a0565b838111156121ca576000848401525b50505050565b6000601f19601f8301169050919050565b6121ea81612146565b81146121f557600080fd5b50565b61220181612184565b811461220c57600080fd5b5056fea26469706673582212201e7a39b7f75a6788aa69a23a304943e77a8636c2158eac80869e9e9baa9e991164736f6c634300060a0033"

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

// GetServiceBinding is a free data retrieval call binding the contract method 0xbd36df16.
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

// GetServiceBinding is a free data retrieval call binding the contract method 0xbd36df16.
//
// Solidity: function GetServiceBinding(string _serviceName) constant returns(iServiceMarketExServiceBinding binding)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBinding(_serviceName string) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBinding(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceBinding is a free data retrieval call binding the contract method 0xbd36df16.
//
// Solidity: function GetServiceBinding(string _serviceName) constant returns(iServiceMarketExServiceBinding binding)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBinding(_serviceName string) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBinding(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceBindingByIndex is a free data retrieval call binding the contract method 0xa7c63969.
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

// GetServiceBindingByIndex is a free data retrieval call binding the contract method 0xa7c63969.
//
// Solidity: function GetServiceBindingByIndex(uint256 _index) constant returns(iServiceMarketExServiceBinding binding)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBindingByIndex(_index *big.Int) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingByIndex(&_IServiceMarketEx.CallOpts, _index)
}

// GetServiceBindingByIndex is a free data retrieval call binding the contract method 0xa7c63969.
//
// Solidity: function GetServiceBindingByIndex(uint256 _index) constant returns(iServiceMarketExServiceBinding binding)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBindingByIndex(_index *big.Int) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingByIndex(&_IServiceMarketEx.CallOpts, _index)
}

// GetServiceBindingCount is a free data retrieval call binding the contract method 0xf8bdb54a.
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

// GetServiceBindingCount is a free data retrieval call binding the contract method 0xf8bdb54a.
//
// Solidity: function GetServiceBindingCount() constant returns(uint256 count)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBindingCount() (*big.Int, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingCount(&_IServiceMarketEx.CallOpts)
}

// GetServiceBindingCount is a free data retrieval call binding the contract method 0xf8bdb54a.
//
// Solidity: function GetServiceBindingCount() constant returns(uint256 count)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBindingCount() (*big.Int, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingCount(&_IServiceMarketEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
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

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceMarketEx *IServiceMarketExSession) IsOwner() (bool, error) {
	return _IServiceMarketEx.Contract.IsOwner(&_IServiceMarketEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceMarketEx *IServiceMarketExCallerSession) IsOwner() (bool, error) {
	return _IServiceMarketEx.Contract.IsOwner(&_IServiceMarketEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IServiceMarketEx *IServiceMarketExSession) Owner() (common.Address, error) {
	return _IServiceMarketEx.Contract.Owner(&_IServiceMarketEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IServiceMarketEx *IServiceMarketExCallerSession) Owner() (common.Address, error) {
	return _IServiceMarketEx.Contract.Owner(&_IServiceMarketEx.CallOpts)
}

// AddServiceBinding is a paid mutator transaction binding the contract method 0x92ce31c4.
//
// Solidity: function addServiceBinding(string _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) AddServiceBinding(opts *bind.TransactOpts, _serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.contract.Transact(opts, "addServiceBinding", _serviceName, _schemas, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExTransactor) AsyncAddServiceBinding(handler func(*types.Receipt, error), opts *bind.TransactOpts, _serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.AsyncTransact(opts, handler, "addServiceBinding", _serviceName, _schemas, _provider, _serviceFee, _qos)
}

// AddServiceBinding is a paid mutator transaction binding the contract method 0x92ce31c4.
//
// Solidity: function addServiceBinding(string _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExSession) AddServiceBinding(_serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.AddServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _schemas, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExSession) AsyncAddServiceBinding(handler func(*types.Receipt, error), _serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncAddServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName, _schemas, _provider, _serviceFee, _qos)
}

// AddServiceBinding is a paid mutator transaction binding the contract method 0x92ce31c4.
//
// Solidity: function addServiceBinding(string _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) AddServiceBinding(_serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.AddServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _schemas, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExTransactorSession) AsyncAddServiceBinding(handler func(*types.Receipt, error), _serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncAddServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName, _schemas, _provider, _serviceFee, _qos)
}

// RemoveServiceBinding is a paid mutator transaction binding the contract method 0xbce9f63f.
//
// Solidity: function removeServiceBinding(string _serviceName) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) RemoveServiceBinding(opts *bind.TransactOpts, _serviceName string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.contract.Transact(opts, "removeServiceBinding", _serviceName)
}

func (_IServiceMarketEx *IServiceMarketExTransactor) AsyncRemoveServiceBinding(handler func(*types.Receipt, error), opts *bind.TransactOpts, _serviceName string) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.AsyncTransact(opts, handler, "removeServiceBinding", _serviceName)
}

// RemoveServiceBinding is a paid mutator transaction binding the contract method 0xbce9f63f.
//
// Solidity: function removeServiceBinding(string _serviceName) returns()
func (_IServiceMarketEx *IServiceMarketExSession) RemoveServiceBinding(_serviceName string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.RemoveServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName)
}

func (_IServiceMarketEx *IServiceMarketExSession) AsyncRemoveServiceBinding(handler func(*types.Receipt, error), _serviceName string) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncRemoveServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName)
}

// RemoveServiceBinding is a paid mutator transaction binding the contract method 0xbce9f63f.
//
// Solidity: function removeServiceBinding(string _serviceName) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) RemoveServiceBinding(_serviceName string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.RemoveServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName)
}

func (_IServiceMarketEx *IServiceMarketExTransactorSession) AsyncRemoveServiceBinding(handler func(*types.Receipt, error), _serviceName string) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncRemoveServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_IServiceMarketEx *IServiceMarketExTransactor) AsyncTransferOwnership(handler func(*types.Receipt, error), opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.AsyncTransact(opts, handler, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceMarketEx *IServiceMarketExSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.TransferOwnership(&_IServiceMarketEx.TransactOpts, newOwner)
}

func (_IServiceMarketEx *IServiceMarketExSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncTransferOwnership(handler, &_IServiceMarketEx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.TransferOwnership(&_IServiceMarketEx.TransactOpts, newOwner)
}

func (_IServiceMarketEx *IServiceMarketExTransactorSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncTransferOwnership(handler, &_IServiceMarketEx.TransactOpts, newOwner)
}

// UpdateServiceBinding is a paid mutator transaction binding the contract method 0x6e29ea60.
//
// Solidity: function updateServiceBinding(string _serviceName, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) UpdateServiceBinding(opts *bind.TransactOpts, _serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.contract.Transact(opts, "updateServiceBinding", _serviceName, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExTransactor) AsyncUpdateServiceBinding(handler func(*types.Receipt, error), opts *bind.TransactOpts, _serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.AsyncTransact(opts, handler, "updateServiceBinding", _serviceName, _provider, _serviceFee, _qos)
}

// UpdateServiceBinding is a paid mutator transaction binding the contract method 0x6e29ea60.
//
// Solidity: function updateServiceBinding(string _serviceName, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExSession) UpdateServiceBinding(_serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, *types.Receipt, error) {
	return _IServiceMarketEx.Contract.UpdateServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _provider, _serviceFee, _qos)
}

func (_IServiceMarketEx *IServiceMarketExSession) AsyncUpdateServiceBinding(handler func(*types.Receipt, error), _serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AsyncUpdateServiceBinding(handler, &_IServiceMarketEx.TransactOpts, _serviceName, _provider, _serviceFee, _qos)
}

// UpdateServiceBinding is a paid mutator transaction binding the contract method 0x6e29ea60.
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

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
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

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
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

// FilterServiceBindingAdded is a free log retrieval operation binding the contract event 0x365f7e387fe001fa02d1c008ec11e2c21f47c6a2102f497b300a17c83127dc69.
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

// WatchServiceBindingAdded is a free log subscription operation binding the contract event 0x365f7e387fe001fa02d1c008ec11e2c21f47c6a2102f497b300a17c83127dc69.
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

// ParseServiceBindingAdded is a log parse operation binding the contract event 0x365f7e387fe001fa02d1c008ec11e2c21f47c6a2102f497b300a17c83127dc69.
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

// FilterServiceBindingRemoved is a free log retrieval operation binding the contract event 0x045edfc0de1510eef86edd984a529791dd2924769af547ab858d860139c7a54d.
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

// WatchServiceBindingRemoved is a free log subscription operation binding the contract event 0x045edfc0de1510eef86edd984a529791dd2924769af547ab858d860139c7a54d.
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

// ParseServiceBindingRemoved is a log parse operation binding the contract event 0x045edfc0de1510eef86edd984a529791dd2924769af547ab858d860139c7a54d.
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

// FilterServiceBindingUpdated is a free log retrieval operation binding the contract event 0xac059a9f400ca20e9086f7e73b998697778b9a48fe40d2392bf649752c5538b2.
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

// WatchServiceBindingUpdated is a free log subscription operation binding the contract event 0xac059a9f400ca20e9086f7e73b998697778b9a48fe40d2392bf649752c5538b2.
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

// ParseServiceBindingUpdated is a log parse operation binding the contract event 0xac059a9f400ca20e9086f7e73b998697778b9a48fe40d2392bf649752c5538b2.
//
// Solidity: event ServiceBindingUpdated(string indexed _serviceName, string _provider, string _serviceFee, uint256 _qos)
func (_IServiceMarketEx *IServiceMarketExFilterer) ParseServiceBindingUpdated(log types.Log) (*IServiceMarketExServiceBindingUpdated, error) {
	event := new(IServiceMarketExServiceBindingUpdated)
	if err := _IServiceMarketEx.contract.UnpackLog(event, "ServiceBindingUpdated", log); err != nil {
		return nil, err
	}
	return event, nil
}
