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

// IServiceDelegatorABI is the input ABI used to generate the binding from.
const IServiceDelegatorABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_iServiceCore\",\"type\":\"address\"}],\"name\":\"setIServiceCore\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"iServiceCore\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_iServiceCore\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"}]"

// IServiceDelegatorBin is the compiled bytecode used for deploying new contracts.
var IServiceDelegatorBin = "0x608060405234801561001057600080fd5b506040516020806104488339810180604052810190808051906020019092919050505033600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610384806100c46000396000f300608060405260043610610057576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680635089e2c8146100b0578063714902b014610107578063f221ba2e1461014a575b34801561006357600080fd5b506040513680600083378082016040526000808284346000546127105a03f1806001811461009857600081146100a5576100aa565b3d604051816000823e8181f35b600080fd5b50505050005b3480156100bc57600080fd5b506100c56101a1565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561011357600080fd5b50610148600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506101c7565b005b34801561015657600080fd5b5061015f610333565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b80600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515610293576040517fc703cb1200000000000000000000000000000000000000000000000000000000815260040180806020018281038252602b8152602001807f695365727669636544656c656761746f723a2061646472657373206d7573742081526020017f6e6f74206265207a65726f00000000000000000000000000000000000000000081525060400191505060405180910390fd5b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156102ef57600080fd5b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a723058202326e463480bffb7ba875440a0d71104ad0000f926be8554ba57ba9b9bfdfeab0029"

// DeployIServiceDelegator deploys a new contract, binding an instance of IServiceDelegator to it.
func DeployIServiceDelegator(auth *bind.TransactOpts, backend bind.ContractBackend, _iServiceCore common.Address) (common.Address, *types.Transaction, *IServiceDelegator, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceDelegatorABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IServiceDelegatorBin), backend, _iServiceCore)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IServiceDelegator{IServiceDelegatorCaller: IServiceDelegatorCaller{contract: contract}, IServiceDelegatorTransactor: IServiceDelegatorTransactor{contract: contract}, IServiceDelegatorFilterer: IServiceDelegatorFilterer{contract: contract}}, nil
}

func AsyncDeployIServiceDelegator(auth *bind.TransactOpts, handler func(*types.Receipt, error), backend bind.ContractBackend, _iServiceCore common.Address) (*types.Transaction, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceDelegatorABI))
	if err != nil {
		return nil, err
	}

	tx, err := bind.AsyncDeployContract(auth, handler, parsed, common.FromHex(IServiceDelegatorBin), backend, _iServiceCore)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

// IServiceDelegator is an auto generated Go binding around a Solidity contract.
type IServiceDelegator struct {
	IServiceDelegatorCaller     // Read-only binding to the contract
	IServiceDelegatorTransactor // Write-only binding to the contract
	IServiceDelegatorFilterer   // Log filterer for contract events
}

// IServiceDelegatorCaller is an auto generated read-only Go binding around a Solidity contract.
type IServiceDelegatorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceDelegatorTransactor is an auto generated write-only Go binding around a Solidity contract.
type IServiceDelegatorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceDelegatorFilterer is an auto generated log filtering Go binding around a Solidity contract events.
type IServiceDelegatorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceDelegatorSession is an auto generated Go binding around a Solidity contract,
// with pre-set call and transact options.
type IServiceDelegatorSession struct {
	Contract     *IServiceDelegator // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// IServiceDelegatorCallerSession is an auto generated read-only Go binding around a Solidity contract,
// with pre-set call options.
type IServiceDelegatorCallerSession struct {
	Contract *IServiceDelegatorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// IServiceDelegatorTransactorSession is an auto generated write-only Go binding around a Solidity contract,
// with pre-set transact options.
type IServiceDelegatorTransactorSession struct {
	Contract     *IServiceDelegatorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// IServiceDelegatorRaw is an auto generated low-level Go binding around a Solidity contract.
type IServiceDelegatorRaw struct {
	Contract *IServiceDelegator // Generic contract binding to access the raw methods on
}

// IServiceDelegatorCallerRaw is an auto generated low-level read-only Go binding around a Solidity contract.
type IServiceDelegatorCallerRaw struct {
	Contract *IServiceDelegatorCaller // Generic read-only contract binding to access the raw methods on
}

// IServiceDelegatorTransactorRaw is an auto generated low-level write-only Go binding around a Solidity contract.
type IServiceDelegatorTransactorRaw struct {
	Contract *IServiceDelegatorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIServiceDelegator creates a new instance of IServiceDelegator, bound to a specific deployed contract.
func NewIServiceDelegator(address common.Address, backend bind.ContractBackend) (*IServiceDelegator, error) {
	contract, err := bindIServiceDelegator(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IServiceDelegator{IServiceDelegatorCaller: IServiceDelegatorCaller{contract: contract}, IServiceDelegatorTransactor: IServiceDelegatorTransactor{contract: contract}, IServiceDelegatorFilterer: IServiceDelegatorFilterer{contract: contract}}, nil
}

// NewIServiceDelegatorCaller creates a new read-only instance of IServiceDelegator, bound to a specific deployed contract.
func NewIServiceDelegatorCaller(address common.Address, caller bind.ContractCaller) (*IServiceDelegatorCaller, error) {
	contract, err := bindIServiceDelegator(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceDelegatorCaller{contract: contract}, nil
}

// NewIServiceDelegatorTransactor creates a new write-only instance of IServiceDelegator, bound to a specific deployed contract.
func NewIServiceDelegatorTransactor(address common.Address, transactor bind.ContractTransactor) (*IServiceDelegatorTransactor, error) {
	contract, err := bindIServiceDelegator(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IServiceDelegatorTransactor{contract: contract}, nil
}

// NewIServiceDelegatorFilterer creates a new log filterer instance of IServiceDelegator, bound to a specific deployed contract.
func NewIServiceDelegatorFilterer(address common.Address, filterer bind.ContractFilterer) (*IServiceDelegatorFilterer, error) {
	contract, err := bindIServiceDelegator(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IServiceDelegatorFilterer{contract: contract}, nil
}

// bindIServiceDelegator binds a generic wrapper to an already deployed contract.
func bindIServiceDelegator(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IServiceDelegatorABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IServiceDelegator *IServiceDelegatorRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IServiceDelegator.Contract.IServiceDelegatorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IServiceDelegator *IServiceDelegatorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IServiceDelegator.Contract.IServiceDelegatorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceDelegator *IServiceDelegatorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IServiceDelegator.Contract.IServiceDelegatorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IServiceDelegator *IServiceDelegatorCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IServiceDelegator.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IServiceDelegator *IServiceDelegatorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, *types.Receipt, error) {
	return _IServiceDelegator.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceDelegator *IServiceDelegatorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, *types.Receipt, error) {
	return _IServiceDelegator.Contract.contract.Transact(opts, method, params...)
}

// IServiceCore is a free data retrieval call binding the contract method 0xf221ba2e.
//
// Solidity: function iServiceCore() constant returns(address)
func (_IServiceDelegator *IServiceDelegatorCaller) IServiceCore(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IServiceDelegator.contract.Call(opts, out, "iServiceCore")
	return *ret0, err
}

// IServiceCore is a free data retrieval call binding the contract method 0xf221ba2e.
//
// Solidity: function iServiceCore() constant returns(address)
func (_IServiceDelegator *IServiceDelegatorSession) IServiceCore() (common.Address, error) {
	return _IServiceDelegator.Contract.IServiceCore(&_IServiceDelegator.CallOpts)
}

// IServiceCore is a free data retrieval call binding the contract method 0xf221ba2e.
//
// Solidity: function iServiceCore() constant returns(address)
func (_IServiceDelegator *IServiceDelegatorCallerSession) IServiceCore() (common.Address, error) {
	return _IServiceDelegator.Contract.IServiceCore(&_IServiceDelegator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x5089e2c8.
//
// Solidity: function owner() constant returns(address)
func (_IServiceDelegator *IServiceDelegatorCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IServiceDelegator.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x5089e2c8.
//
// Solidity: function owner() constant returns(address)
func (_IServiceDelegator *IServiceDelegatorSession) Owner() (common.Address, error) {
	return _IServiceDelegator.Contract.Owner(&_IServiceDelegator.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x5089e2c8.
//
// Solidity: function owner() constant returns(address)
func (_IServiceDelegator *IServiceDelegatorCallerSession) Owner() (common.Address, error) {
	return _IServiceDelegator.Contract.Owner(&_IServiceDelegator.CallOpts)
}

// SetIServiceCore is a paid mutator transaction binding the contract method 0x714902b0.
//
// Solidity: function setIServiceCore(address _iServiceCore) returns()
func (_IServiceDelegator *IServiceDelegatorTransactor) SetIServiceCore(opts *bind.TransactOpts, _iServiceCore common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceDelegator.contract.Transact(opts, "setIServiceCore", _iServiceCore)
}

func (_IServiceDelegator *IServiceDelegatorTransactor) AsyncSetIServiceCore(handler func(*types.Receipt, error), opts *bind.TransactOpts, _iServiceCore common.Address) (*types.Transaction, error) {
	return _IServiceDelegator.contract.AsyncTransact(opts, handler, "setIServiceCore", _iServiceCore)
}

// SetIServiceCore is a paid mutator transaction binding the contract method 0x714902b0.
//
// Solidity: function setIServiceCore(address _iServiceCore) returns()
func (_IServiceDelegator *IServiceDelegatorSession) SetIServiceCore(_iServiceCore common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceDelegator.Contract.SetIServiceCore(&_IServiceDelegator.TransactOpts, _iServiceCore)
}

func (_IServiceDelegator *IServiceDelegatorSession) AsyncSetIServiceCore(handler func(*types.Receipt, error), _iServiceCore common.Address) (*types.Transaction, error) {
	return _IServiceDelegator.Contract.AsyncSetIServiceCore(handler, &_IServiceDelegator.TransactOpts, _iServiceCore)
}

// SetIServiceCore is a paid mutator transaction binding the contract method 0x714902b0.
//
// Solidity: function setIServiceCore(address _iServiceCore) returns()
func (_IServiceDelegator *IServiceDelegatorTransactorSession) SetIServiceCore(_iServiceCore common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceDelegator.Contract.SetIServiceCore(&_IServiceDelegator.TransactOpts, _iServiceCore)
}

func (_IServiceDelegator *IServiceDelegatorTransactorSession) AsyncSetIServiceCore(handler func(*types.Receipt, error), _iServiceCore common.Address) (*types.Transaction, error) {
	return _IServiceDelegator.Contract.AsyncSetIServiceCore(handler, &_IServiceDelegator.TransactOpts, _iServiceCore)
}
