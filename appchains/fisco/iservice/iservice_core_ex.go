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
const IServiceCoreExABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iServiceMarket\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"_contractAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_provider\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_serviceFeeCap\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"ServiceInvoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackFunction\",\"type\":\"bytes4\"}],\"name\":\"callService\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"}],\"name\":\"getResponse\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"errMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"icRequestID\",\"type\":\"string\"}],\"internalType\":\"structiServiceCoreEx.Response\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"iServiceMarket\",\"outputs\":[{\"internalType\":\"contractiServiceMarketInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"relayer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setIServiceMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_errMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_icRequestID\",\"type\":\"string\"}],\"name\":\"setResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IServiceCoreExBin is the compiled bytecode used for deploying new contracts.
var IServiceCoreExBin = "0x60806040526040518060400160405280601781526020017f7b22686561646572223a7b7d2c22626f6479223a7b7d7d000000000000000000815250600790805190602001906200005192919062000307565b503480156200005f57600080fd5b5060405162002b5e38038062002b5e8339818101604052810190620000859190620003cd565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f5c7c30d4a0f08950cb23be4132957b357fa5dfdb0fcf218f81b86a1c036e47d060405160405180910390a362000152826200022760201b60201c565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614620001ce5780600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506200021f565b620001de620002de60201b60201c565b600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5050620004f7565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156200029a576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401620002919062000476565b60405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200034a57805160ff19168380011785556200037b565b828001600101855582156200037b579182015b828111156200037a5782518255916020019190600101906200035d565b5b5090506200038a91906200038e565b5090565b620003b391905b80821115620003af57600081600090555060010162000395565b5090565b90565b600081519050620003c781620004dd565b92915050565b60008060408385031215620003e157600080fd5b6000620003f185828601620003b6565b92505060206200040485828601620003b6565b9150509250929050565b60006200041d60378362000498565b91507f6953657276696365436f726545783a206953657276696365206d61726b65742060008301527f616464726573732063616e206e6f74206265207a65726f0000000000000000006020830152604082019050919050565b6000602082019050818103600083015262000491816200040e565b9050919050565b600082825260208201905092915050565b6000620004b682620004bd565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b620004e881620004a9565b8114620004f457600080fd5b50565b61265780620005076000396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063744e931a11610071578063744e931a14610152578063781c89c6146101825780637a1f2dea146101b2578063a70ac4dc146101ce578063e9667acd146101ec578063ede8e5291461020a576100a9565b806302b43c19146100ae57806316cad12a146100de5780633a938f2d146100fa5780635089e2c814610118578063650d6e8514610136575b600080fd5b6100c860048036038101906100c391906119ae565b610228565b6040516100d5919061211d565b60405180910390f35b6100f860048036038101906100f3919061187a565b6104c7565b005b61010261051a565b60405161010f91906123b2565b60405180910390f35b610120610520565b60405161012d91906120e7565b60405180910390f35b610150600480360381019061014b919061187a565b610549565b005b61016c600480360381019061016791906118cc565b61059c565b6040516101799190612390565b60405180910390f35b61019c600480360381019061019791906118f5565b6107af565b6040516101a99190612102565b60405180910390f35b6101cc60048036038101906101c7919061187a565b610c54565b005b6101d6610d4f565b6040516101e391906121f3565b60405180910390f35b6101f4610d75565b60405161020191906120e7565b60405180910390f35b610212610d9b565b60405161021f9190612102565b60405180910390f35b600087878080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508460008251116102b4576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016102ab90612310565b60405180910390fd5b600081116102f7576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016102ee906122d0565b60405180910390fd5b6102ff6115f3565b33816020019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508a8a8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050816040018190525088888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508160800181905250868160c00181815250506000816080015151141561048d5760078054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561047f5780601f106104545761010080835404028352916020019161047f565b820191906000526020600020905b81548152906001019060200180831161046257829003601f168201915b505050505081608001819052505b61049681610df2565b61049f81611168565b9350838160000181815250506104b68487876112ff565b839350505050979650505050505050565b6104cf610d9b565b61050e576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610505906122f0565b60405180910390fd5b61051781611411565b50565b60045481565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b610551610d9b565b610590576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610587906122f0565b60405180910390fd5b6105998161153f565b50565b6105a4611649565b60036000838152602001908152602001600020604051806060016040529081600082018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561065b5780601f106106305761010080835404028352916020019161065b565b820191906000526020600020905b81548152906001019060200180831161063e57829003601f168201915b50505050508152602001600182018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106fd5780601f106106d2576101008083540402835291602001916106fd565b820191906000526020600020905b8154815290600101906020018083116106e057829003601f168201915b50505050508152602001600282018054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561079f5780601f106107745761010080835404028352916020019161079f565b820191906000526020600020905b81548152906001019060200180831161078257829003601f168201915b5050505050815250509050919050565b6000600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610841576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161083890612330565b60405180910390fd5b87600060016000838152602001908152602001600020600201805460018160011615610100020316600290049050116108af576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016108a690612290565b60405180910390fd5b6000600360008381526020019081526020016000206002018054600181600116156101000203166002900490501461091c576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161091390612230565b60405180910390fd5b610924611649565b88888080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050816000018190525084848080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081604001819052506000898990501415610a185786868080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f8201169050808301925050505050505081602001819052505b80600360008c81526020019081526020016000206000820151816000019080519060200190610a4892919061166a565b506020820151816001019080519060200190610a6592919061166a565b506040820151816002019080519060200190610a8292919061166a565b50905050610a8e6116ea565b600260008c81526020019081526020016000206040518060400160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016000820160149054906101000a900460e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19168152505090506000816000015173ffffffffffffffffffffffffffffffffffffffff1682602001518d8560200151604051602401610b929291906121c3565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610bfc91906120a4565b6000604051808303816000865af19150503d8060008114610c39576040519150601f19603f3d011682016040523d82523d6000602084013e610c3e565b606091505b5050905080945050505050979650505050505050565b610c5c610d9b565b610c9b576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610c92906122f0565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415610d0b576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610d0290612350565b60405180910390fd5b80600560006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b600560009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614905090565b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166327e4175083604001516040518263ffffffff1660e01b8152600401610e53919061220e565b60206040518083038186803b158015610e6b57600080fd5b505afa158015610e7f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610ea391906118a3565b905080610ee5576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610edc90612270565b60405180910390fd5b6000600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663eac81af884604001516040518263ffffffff1660e01b8152600401610f46919061220e565b60206040518083038186803b158015610f5e57600080fd5b505afa158015610f72573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f969190611a9f565b9050808360c001511015610fdf576040517fc703cb12000000000000000000000000000000000000000000000000000000008152600401610fd6906122b0565b60405180910390fd5b6060600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663438166c685604001516040518263ffffffff1660e01b8152600401611040919061220e565b60006040518083038186803b15801561105857600080fd5b505afa15801561106c573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906110959190611a5e565b90506060600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16639b0a965c86604001516040518263ffffffff1660e01b81526004016110f8919061220e565b60006040518083038186803b15801561111057600080fd5b505afa158015611124573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061114d9190611a5e565b9050818560600181905250808560a001819052505050505050565b60003060045460405160200161117f9291906120bb565b6040516020818303038152906040528051906020012090508082600001818152505081600160008381526020019081526020016000206000820151816000015560208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550604082015181600201908051906020019061122292919061166a565b50606082015181600301908051906020019061123f92919061166a565b50608082015181600401908051906020019061125c92919061166a565b5060a082015181600501908051906020019061127992919061166a565b5060c0820151816006015590505060016004600082825401925050819055507f9806d3fe568da5236606718b2a2b923b61679261c9301d4719c3b84b8a4229e2826000015183602001518460400151856060015186608001518760a001518860c001516040516112ef9796959493929190612138565b60405180910390a1809050919050565b6113076116ea565b82816000019073ffffffffffffffffffffffffffffffffffffffff16908173ffffffffffffffffffffffffffffffffffffffff16815250508181602001907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191690817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525050806002600086815260200190815260200160002060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160000160146101000a81548163ffffffff021916908360e01c021790555090505050505050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415611481576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040161147890612250565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f5c7c30d4a0f08950cb23be4132957b357fa5dfdb0fcf218f81b86a1c036e47d060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614156115af576040517fc703cb120000000000000000000000000000000000000000000000000000000081526004016115a690612370565b60405180910390fd5b80600660006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6040518060e0016040528060008019168152602001600073ffffffffffffffffffffffffffffffffffffffff16815260200160608152602001606081526020016060815260200160608152602001600081525090565b60405180606001604052806060815260200160608152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106116ab57805160ff19168380011785556116d9565b828001600101855582156116d9579182015b828111156116d85782518255916020019190600101906116bd565b5b5090506116e69190611739565b5090565b6040518060400160405280600073ffffffffffffffffffffffffffffffffffffffff16815260200160007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191681525090565b61175b91905b8082111561175757600081600090555060010161173f565b5090565b90565b60008135905061176d816125ae565b92915050565b600081519050611782816125c5565b92915050565b600081359050611797816125dc565b92915050565b6000813590506117ac816125f3565b92915050565b60008083601f8401126117c457600080fd5b8235905067ffffffffffffffff8111156117dd57600080fd5b6020830191508360018202830111156117f557600080fd5b9250929050565b600082601f83011261180d57600080fd5b815161182061181b826123fa565b6123cd565b9150808252602083016020830185838301111561183c57600080fd5b61184783828461252f565b50505092915050565b60008135905061185f8161260a565b92915050565b6000815190506118748161260a565b92915050565b60006020828403121561188c57600080fd5b600061189a8482850161175e565b91505092915050565b6000602082840312156118b557600080fd5b60006118c384828501611773565b91505092915050565b6000602082840312156118de57600080fd5b60006118ec84828501611788565b91505092915050565b60008060008060008060006080888a03121561191057600080fd5b600061191e8a828b01611788565b975050602088013567ffffffffffffffff81111561193b57600080fd5b6119478a828b016117b2565b9650965050604088013567ffffffffffffffff81111561196657600080fd5b6119728a828b016117b2565b9450945050606088013567ffffffffffffffff81111561199157600080fd5b61199d8a828b016117b2565b925092505092959891949750929550565b600080600080600080600060a0888a0312156119c957600080fd5b600088013567ffffffffffffffff8111156119e357600080fd5b6119ef8a828b016117b2565b9750975050602088013567ffffffffffffffff811115611a0e57600080fd5b611a1a8a828b016117b2565b95509550506040611a2d8a828b01611850565b9350506060611a3e8a828b0161175e565b9250506080611a4f8a828b0161179d565b91505092959891949750929550565b600060208284031215611a7057600080fd5b600082015167ffffffffffffffff811115611a8a57600080fd5b611a96848285016117fc565b91505092915050565b600060208284031215611ab157600080fd5b6000611abf84828501611865565b91505092915050565b611ad181612469565b82525050565b611ae08161247b565b82525050565b611aef81612487565b82525050565b6000611b0082612426565b611b0a818561243c565b9350611b1a81856020860161252f565b80840191505092915050565b611b37611b32826124e7565b612562565b82525050565b611b468161250b565b82525050565b6000611b5782612431565b611b618185612447565b9350611b7181856020860161252f565b611b7a81612590565b840191505092915050565b6000611b9082612431565b611b9a8185612458565b9350611baa81856020860161252f565b611bb381612590565b840191505092915050565b6000611bcb602a83612458565b91507f6953657276696365436f726545783a207265717565737420686173206265656e60008301527f20726573706f6e646564000000000000000000000000000000000000000000006020830152604082019050919050565b6000611c31602683612458565b91507f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160008301527f64647265737300000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611c97603c83612458565b91507f6953657276696365436f726545783a207365727669636520646f6573206e6f7460008301527f20657869737420696e207468652073657276696365206d61726b6574000000006020830152604082019050919050565b6000611cfd602683612458565b91507f6953657276696365436f726545783a207265717565737420646f6573206e6f7460008301527f20657869737400000000000000000000000000000000000000000000000000006020830152604082019050919050565b6000611d63605083612458565b91507f6953657276696365436f726545783a20726571756573742074696d656f75742060008301527f6d7573742062652067726561746572207468616e206f7220657175616c20746f60208301527f20746865207365727669636520516f53000000000000000000000000000000006040830152606082019050919050565b6000611def603683612458565b91507f6953657276696365436f726545783a20726571756573742074696d656f75742060008301527f6d7573742062652067726561746572207468616e2030000000000000000000006020830152604082019050919050565b6000611e55602083612458565b91507f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726000830152602082019050919050565b6000611e95602d83612458565b91507f6953657276696365436f726545783a2073657276696365206e616d652063616e60008301527f206e6f7420626520656d707479000000000000000000000000000000000000006020830152604082019050919050565b6000611efb602983612458565b91507f6953657276696365436f726545783a2073656e646572206973206e6f7420746860008301527f652072656c6179657200000000000000000000000000000000000000000000006020830152604082019050919050565b6000611f61602f83612458565b91507f6953657276696365436f726545783a2072656c6179657220616464726573732060008301527f63616e206e6f74206265207a65726f00000000000000000000000000000000006020830152604082019050919050565b6000611fc7603783612458565b91507f6953657276696365436f726545783a206953657276696365206d61726b65742060008301527f616464726573732063616e206e6f74206265207a65726f0000000000000000006020830152604082019050919050565b6000606083016000830151848203600086015261203d8282611b4c565b915050602083015184820360208601526120578282611b4c565b915050604083015184820360408601526120718282611b4c565b9150508091505092915050565b612087816124dd565b82525050565b61209e612099826124dd565b612586565b82525050565b60006120b08284611af5565b915081905092915050565b60006120c78285611b26565b6014820191506120d7828461208d565b6020820191508190509392505050565b60006020820190506120fc6000830184611ac8565b92915050565b60006020820190506121176000830184611ad7565b92915050565b60006020820190506121326000830184611ae6565b92915050565b600060e08201905061214d600083018a611ae6565b61215a6020830189611ac8565b818103604083015261216c8188611b85565b905081810360608301526121808187611b85565b905081810360808301526121948186611b85565b905081810360a08301526121a88185611b85565b90506121b760c083018461207e565b98975050505050505050565b60006040820190506121d86000830185611ae6565b81810360208301526121ea8184611b85565b90509392505050565b60006020820190506122086000830184611b3d565b92915050565b600060208201905081810360008301526122288184611b85565b905092915050565b6000602082019050818103600083015261224981611bbe565b9050919050565b6000602082019050818103600083015261226981611c24565b9050919050565b6000602082019050818103600083015261228981611c8a565b9050919050565b600060208201905081810360008301526122a981611cf0565b9050919050565b600060208201905081810360008301526122c981611d56565b9050919050565b600060208201905081810360008301526122e981611de2565b9050919050565b6000602082019050818103600083015261230981611e48565b9050919050565b6000602082019050818103600083015261232981611e88565b9050919050565b6000602082019050818103600083015261234981611eee565b9050919050565b6000602082019050818103600083015261236981611f54565b9050919050565b6000602082019050818103600083015261238981611fba565b9050919050565b600060208201905081810360008301526123aa8184612020565b905092915050565b60006020820190506123c7600083018461207e565b92915050565b6000604051905081810181811067ffffffffffffffff821117156123f057600080fd5b8060405250919050565b600067ffffffffffffffff82111561241157600080fd5b601f19601f8301169050602081019050919050565b600081519050919050565b600081519050919050565b600081905092915050565b600082825260208201905092915050565b600082825260208201905092915050565b6000612474826124bd565b9050919050565b60008115159050919050565b6000819050919050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082169050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006124f2826124f9565b9050919050565b6000612504826124bd565b9050919050565b60006125168261251d565b9050919050565b6000612528826124bd565b9050919050565b60005b8381101561254d578082015181840152602081019050612532565b8381111561255c576000848401525b50505050565b600061256d82612574565b9050919050565b600061257f826125a1565b9050919050565b6000819050919050565b6000601f19601f8301169050919050565b60008160601b9050919050565b6125b781612469565b81146125c257600080fd5b50565b6125ce8161247b565b81146125d957600080fd5b50565b6125e581612487565b81146125f057600080fd5b50565b6125fc81612491565b811461260757600080fd5b50565b612613816124dd565b811461261e57600080fd5b5056fea2646970667358221220ca6e7205e3a11688ccb2d6588483210468d422ac0efad33f76c7d202912455ed64736f6c634300060a0033"

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
