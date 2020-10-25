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

// iServiceCoreExResponse is an auto generated low-level Go binding around an user-defined struct.
type iServiceCoreExResponse struct {
	ErrMsg      string
	Output      string
	IcRequestID string
}

// IServiceCoreExABI is the input ABI used to generate the binding from.
const IServiceCoreExABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"ServiceInvoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackFunction\",\"type\":\"bytes4\"}],\"name\":\"callService\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"}],\"name\":\"getResponse\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"errMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"icRequestID\",\"type\":\"string\"}],\"internalType\":\"structiServiceCoreEx.Response\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_errMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_icRequestID\",\"type\":\"string\"}],\"name\":\"setResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IServiceCoreExBin is the compiled bytecode used for deploying new contracts.
var IServiceCoreExBin = "0x60806040523480156200001157600080fd5b5060405162001f8838038062001f88833981810160405281019062000037919062000188565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f5c7c30d4a0f08950cb23be4132957b357fa5dfdb0fcf218f81b86a1c036e47d060405160405180910390a3600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16146200016a5780600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5062000202565b6000815190506200018281620001e8565b92915050565b6000602082840312156200019b57600080fd5b6000620001ab8482850162000171565b91505092915050565b6000620001c182620001c8565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620001f381620001b4565b8114620001ff57600080fd5b50565b611d7680620002126000396000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c8063744e931a1161005b578063744e931a14610115578063781c89c6146101455780637a1f2dea14610175578063ede8e5291461019157610088565b806302b43c191461008d57806316cad12a146100bd5780633a938f2d146100d95780635089e2c8146100f7575b600080fd5b6100a760048036038101906100a29190611381565b6101af565b6040516100b49190611985565b60405180910390f35b6100d760048036038101906100d29190611276565b61044f565b005b6100e16104a2565b6040516100ee9190611b65565b60405180910390f35b6100ff6104a8565b60405161010c919061194f565b60405180910390f35b61012f600480360381019061012a919061129f565b6104d1565b60405161013c9190611b43565b60405180910390f35b61015f600480360381019061015a91906112c8565b6106e4565b60405161016c919061196a565b60405180910390f35b61018f600480360381019061018a9190611276565b610cd7565b005b610199610dd2565b6040516101a6919061196a565b60405180910390f35b600087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505086868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050856000835111610280576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161027790611ae3565b60405180910390fd5b60008251116102c4576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016102bb90611a43565b60405180910390fd5b60008111610307576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016102fe90611aa3565b60405180910390fd5b61030f6110ae565b8b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050816000018190525089898080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081602001819052508781604001818152505086816060019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508581608001907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152505061043e81610e29565b945050505050979650505050505050565b610457610dd2565b610496576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161048d90611ac3565b60405180910390fd5b61049f81610f80565b50565b60035481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6104d9611112565b60026000838152602001908152602001600020604051806060016040529081600082018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156105905780601f1061056557610100808354040283529160200191610590565b820191906000526020600020905b81548152906001019060200180831161057357829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106325780601f1061060757610100808354040283529160200191610632565b820191906000526020600020905b81548152906001019060200180831161061557829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106d45780601f106106a9576101008083540402835291602001916106d4565b820191906000526020600020905b8154815290600101906020018083116106b757829003601f168201915b5050505050815250509050919050565b6000600460009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610776576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161076d90611b03565b60405180910390fd5b87600060016000838152602001908152602001600020600001805460018160011615610100020316600290049050116107e4576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016107db90611a83565b60405180910390fd5b60006002600083815260200190815260200160002060020180546001816001161561010002031660029004905014610851576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161084890611a23565b60405180910390fd5b610859611112565b88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050816000018190525084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508160400181905250600089899050141561094d5786868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081602001819052505b80600260008c8152602001908152602001600020600082015181600001908051906020019061097d929190611133565b50602082015181600101908051906020019061099a929190611133565b5060408201518160020190805190602001906109b7929190611133565b509050506109c36110ae565b600160008c81526020019081526020016000206040518060a0016040529081600082018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610a7a5780601f10610a4f57610100808354040283529160200191610a7a565b820191906000526020600020905b815481529060010190602001808311610a5d57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b1c5780601f10610af157610100808354040283529160200191610b1c565b820191906000526020600020905b815481529060010190602001808311610aff57829003601f168201915b50505050508152602001600282015481526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160149054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152505090506000816060015173ffffffffffffffffffffffffffffffffffffffff1682608001518d8560200151604051602401610c159291906119a0565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610c7f919061190c565b6000604051808303816000865af19150503d8060008114610cbc576040519150601f19603f3d011682016040523d82523d6000602084013e610cc1565b606091505b5050905080945050505050979650505050505050565b610cdf610dd2565b610d1e576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610d1590611ac3565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610d8e576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610d8590611b23565b60405180910390fd5b80600460006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b600030600354604051602001610e40929190611923565b60405160208183030381529060405280519060200120905081600160008381526020019081526020016000206000820151816000019080519060200190610e88929190611133565b506020820151816001019080519060200190610ea5929190611133565b506040820151816002015560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160030160146101000a81548163ffffffff021916908360e01c021790555090505060016003600082825401925050819055507f7a04e9df14128bb22525f90cdf151ca578b5f9d224ca7803d01d70a0b6f70b1481836000015184602001518560400151604051610f7094939291906119d0565b60405180910390a1809050919050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610ff0576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610fe790611a63565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f5c7c30d4a0f08950cb23be4132957b357fa5dfdb0fcf218f81b86a1c036e47d060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060a00160405280606081526020016060815260200160008152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525090565b60405180606001604052806060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061117457805160ff19168380011785556111a2565b828001600101855582156111a2579182015b828111156111a1578251825591602001919060010190611186565b5b5090506111af91906111b3565b5090565b6111d591905b808211156111d15760008160009055506001016111b9565b5090565b90565b6000813590506111e781611ce4565b92915050565b6000813590506111fc81611cfb565b92915050565b60008135905061121181611d12565b92915050565b60008083601f84011261122957600080fd5b8235905067ffffffffffffffff81111561124257600080fd5b60208301915083600182028301111561125a57600080fd5b9250929050565b60008135905061127081611d29565b92915050565b60006020828403121561128857600080fd5b6000611296848285016111d8565b91505092915050565b6000602082840312156112b157600080fd5b60006112bf848285016111ed565b91505092915050565b60008060008060008060006080888a0312156112e357600080fd5b60006112f18a828b016111ed565b975050602088013567ffffffffffffffff81111561130e57600080fd5b61131a8a828b01611217565b9650965050604088013567ffffffffffffffff81111561133957600080fd5b6113458a828b01611217565b9450945050606088013567ffffffffffffffff81111561136457600080fd5b6113708a828b01611217565b925092505092959891949750929550565b600080600080600080600060a0888a03121561139c57600080fd5b600088013567ffffffffffffffff8111156113b657600080fd5b6113c28a828b01611217565b9750975050602088013567ffffffffffffffff8111156113e157600080fd5b6113ed8a828b01611217565b955095505060406114008a828b01611261565b93505060606114118a828b016111d8565b92505060806114228a828b01611202565b91505092959891949750929550565b61143a81611bc3565b82525050565b61144981611bd5565b82525050565b61145881611be1565b82525050565b600061146982611b80565b6114738185611b96565b9350611483818560208601611c65565b80840191505092915050565b6114a061149b82611c41565b611c98565b82525050565b60006114b182611b8b565b6114bb8185611ba1565b93506114cb818560208601611c65565b6114d481611cc6565b840191505092915050565b60006114ea82611b8b565b6114f48185611bb2565b9350611504818560208601611c65565b61150d81611cc6565b840191505092915050565b6000611525602a83611bb2565b91507f6953657276696365436f726545783a207265717565737420686173206265656e60008301527f20726573706f6e646564000000000000000000000000000000000000000000006020830152604082019050919050565b600061158b602e83611bb2565b91507f6953657276696365436f726545783a207265717565737420696e70757420636160008301527f6e206e6f7420626520656d7074790000000000000000000000000000000000006020830152604082019050919050565b60006115f1602683611bb2565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611657602683611bb2565b91507f6953657276696365436f726545783a207265717565737420646f6573206e6f7460008301527f20657869737400000000000000000000000000000000000000000000000000006020830152604082019050919050565b60006116bd603683611bb2565b91507f6953657276696365436f726545783a20726571756573742074696d656f75742060008301527f6d7573742062652067726561746572207468616e2030000000000000000000006020830152604082019050919050565b6000611723602083611bb2565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000611763602d83611bb2565b91507f6953657276696365436f726545783a2073657276696365206e616d652063616e60008301527f206e6f7420626520656d707479000000000000000000000000000000000000006020830152604082019050919050565b60006117c9602983611bb2565b91507f6953657276696365436f726545783a2073656e646572206973206e6f7420746860008301527f652072656c6179657200000000000000000000000000000000000000000000006020830152604082019050919050565b600061182f602f83611bb2565b91507f6953657276696365436f726545783a2072656c6179657220616464726573732060008301527f63616e206e6f74206265207a65726f00000000000000000000000000000000006020830152604082019050919050565b600060608301600083015184820360008601526118a582826114a6565b915050602083015184820360208601526118bf82826114a6565b915050604083015184820360408601526118d982826114a6565b9150508091505092915050565b6118ef81611c37565b82525050565b61190661190182611c37565b611cbc565b82525050565b6000611918828461145e565b915081905092915050565b600061192f828561148f565b60148201915061193f82846118f5565b6020820191508190509392505050565b60006020820190506119646000830184611431565b92915050565b600060208201905061197f6000830184611440565b92915050565b600060208201905061199a600083018461144f565b92915050565b60006040820190506119b5600083018561144f565b81810360208301526119c781846114df565b90509392505050565b60006080820190506119e5600083018761144f565b81810360208301526119f781866114df565b90508181036040830152611a0b81856114df565b9050611a1a60608301846118e6565b95945050505050565b60006020820190508181036000830152611a3c81611518565b9050919050565b60006020820190508181036000830152611a5c8161157e565b9050919050565b60006020820190508181036000830152611a7c816115e4565b9050919050565b60006020820190508181036000830152611a9c8161164a565b9050919050565b60006020820190508181036000830152611abc816116b0565b9050919050565b60006020820190508181036000830152611adc81611716565b9050919050565b60006020820190508181036000830152611afc81611756565b9050919050565b60006020820190508181036000830152611b1c816117bc565b9050919050565b60006020820190508181036000830152611b3c81611822565b9050919050565b60006020820190508181036000830152611b5d8184611888565b905092915050565b6000602082019050611b7a60008301846118e6565b92915050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000611bce82611c17565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b6000611c4c82611c53565b9050919050565b6000611c5e82611c17565b9050919050565b60005b83811015611c83578082015181840152602081019050611c68565b83811115611c92576000848401525b50505050565b6000611ca382611caa565b9050919050565b6000611cb582611cd7565b9050919050565b6000819050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b611ced81611bc3565b8114611cf857600080fd5b50565b611d0481611be1565b8114611d0f57600080fd5b50565b611d1b81611beb565b8114611d2657600080fd5b50565b611d3281611c37565b8114611d3d57600080fd5b5056fea264697066735822122054b85b4555c69f555587cae715cd547dc37d818a51e8ff668a2de987365e692d64736f6c634300060a0033"

// DeployIServiceCoreEx deploys a new contract, binding an instance of IServiceCoreEx to it.
func DeployIServiceCoreEx(auth *bind.TransactOpts, backend bind.ContractBackend, _relayer common.Address) (common.Address, *types.Transaction, *IServiceCoreEx, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceCoreExABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IServiceCoreExBin), backend, _relayer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IServiceCoreEx{IServiceCoreExCaller: IServiceCoreExCaller{contract: contract}, IServiceCoreExTransactor: IServiceCoreExTransactor{contract: contract}, IServiceCoreExFilterer: IServiceCoreExFilterer{contract: contract}}, nil
}

func AsyncDeployIServiceCoreEx(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _relayer common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceCoreExABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(IServiceCoreExBin), backend, _relayer)
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
// Solidity: function getResponse(bytes32 _requestID) constant returns(iServiceCoreExResponse)
func (_IServiceCoreEx *IServiceCoreExCaller) GetResponse(opts *bind.CallOpts, _requestID [32]byte) (iServiceCoreExResponse, error) {
	var (
		ret0 = new(iServiceCoreExResponse)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "getResponse", _requestID)
	return *ret0, err
}

// GetResponse is a free data retrieval call binding the contract method 0x744e931a.
//
// Solidity: function getResponse(bytes32 _requestID) constant returns(iServiceCoreExResponse)
func (_IServiceCoreEx *IServiceCoreExSession) GetResponse(_requestID [32]byte) (iServiceCoreExResponse, error) {
	return _IServiceCoreEx.Contract.GetResponse(&_IServiceCoreEx.CallOpts, _requestID)
}

// GetResponse is a free data retrieval call binding the contract method 0x744e931a.
//
// Solidity: function getResponse(bytes32 _requestID) constant returns(iServiceCoreExResponse)
func (_IServiceCoreEx *IServiceCoreExCallerSession) GetResponse(_requestID [32]byte) (iServiceCoreExResponse, error) {
	return _IServiceCoreEx.Contract.GetResponse(&_IServiceCoreEx.CallOpts, _requestID)
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
	RequestID   [32]byte
	ServiceName string
	Input       string
	Timeout     *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterServiceInvoked is a free log retrieval operation binding the contract event 0x000000000000000000000000000000000000000000000000000000007a04e9df.
//
// Solidity: event ServiceInvoked(bytes32 _requestID, string _serviceName, string _input, uint256 _timeout)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterServiceInvoked(opts *bind.FilterOpts) (*IServiceCoreExServiceInvokedIterator, error) {

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "ServiceInvoked")
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExServiceInvokedIterator{contract: _IServiceCoreEx.contract, event: "ServiceInvoked", logs: logs, sub: sub}, nil
}

// WatchServiceInvoked is a free log subscription operation binding the contract event 0x000000000000000000000000000000000000000000000000000000007a04e9df.
//
// Solidity: event ServiceInvoked(bytes32 _requestID, string _serviceName, string _input, uint256 _timeout)
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

// ParseServiceInvoked is a log parse operation binding the contract event 0x000000000000000000000000000000000000000000000000000000007a04e9df.
//
// Solidity: event ServiceInvoked(bytes32 _requestID, string _serviceName, string _input, uint256 _timeout)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseServiceInvoked(log types.Log) (*IServiceCoreExServiceInvoked, error) {
	event := new(IServiceCoreExServiceInvoked)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "ServiceInvoked", log); err != nil {
		return nil, err
	}
	return event, nil
}
