// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package iservice

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
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
const IServiceMarketExABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_schemas\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_provider\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_serviceFee\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"ServiceBindingAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"ServiceBindingRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_provider\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_serviceFee\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"ServiceBindingUpdated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"GetServiceBinding\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"schemas\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"provider\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceFee\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"qos\",\"type\":\"uint256\"}],\"internalType\":\"structiServiceMarketEx.ServiceBinding\",\"name\":\"binding\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"GetServiceBindingByIndex\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"schemas\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"provider\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"serviceFee\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"qos\",\"type\":\"uint256\"}],\"internalType\":\"structiServiceMarketEx.ServiceBinding\",\"name\":\"binding\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GetServiceBindingCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_schemas\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_provider\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_serviceFee\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"addServiceBinding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"}],\"name\":\"removeServiceBinding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_provider\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_serviceFee\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_qos\",\"type\":\"uint256\"}],\"name\":\"updateServiceBinding\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IServiceMarketEx is an auto generated Go binding around an Ethereum contract.
type IServiceMarketEx struct {
	IServiceMarketExCaller     // Read-only binding to the contract
	IServiceMarketExTransactor // Write-only binding to the contract
	IServiceMarketExFilterer   // Log filterer for contract events
}

// IServiceMarketExCaller is an auto generated read-only Go binding around an Ethereum contract.
type IServiceMarketExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceMarketExTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IServiceMarketExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceMarketExFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IServiceMarketExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceMarketExSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IServiceMarketExSession struct {
	Contract     *IServiceMarketEx // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IServiceMarketExCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IServiceMarketExCallerSession struct {
	Contract *IServiceMarketExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IServiceMarketExTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IServiceMarketExTransactorSession struct {
	Contract     *IServiceMarketExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IServiceMarketExRaw is an auto generated low-level Go binding around an Ethereum contract.
type IServiceMarketExRaw struct {
	Contract *IServiceMarketEx // Generic contract binding to access the raw methods on
}

// IServiceMarketExCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IServiceMarketExCallerRaw struct {
	Contract *IServiceMarketExCaller // Generic read-only contract binding to access the raw methods on
}

// IServiceMarketExTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
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
func (_IServiceMarketEx *IServiceMarketExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.IServiceMarketExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceMarketEx *IServiceMarketExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
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
func (_IServiceMarketEx *IServiceMarketExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceMarketEx *IServiceMarketExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.contract.Transact(opts, method, params...)
}

// GetServiceBinding is a free data retrieval call binding the contract method 0xbd36df16.
//
// Solidity: function GetServiceBinding(string _serviceName) view returns((string,string,string,string,uint256) binding)
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
// Solidity: function GetServiceBinding(string _serviceName) view returns((string,string,string,string,uint256) binding)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBinding(_serviceName string) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBinding(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceBinding is a free data retrieval call binding the contract method 0xbd36df16.
//
// Solidity: function GetServiceBinding(string _serviceName) view returns((string,string,string,string,uint256) binding)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBinding(_serviceName string) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBinding(&_IServiceMarketEx.CallOpts, _serviceName)
}

// GetServiceBindingByIndex is a free data retrieval call binding the contract method 0xa7c63969.
//
// Solidity: function GetServiceBindingByIndex(uint256 _index) view returns((string,string,string,string,uint256) binding)
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
// Solidity: function GetServiceBindingByIndex(uint256 _index) view returns((string,string,string,string,uint256) binding)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBindingByIndex(_index *big.Int) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingByIndex(&_IServiceMarketEx.CallOpts, _index)
}

// GetServiceBindingByIndex is a free data retrieval call binding the contract method 0xa7c63969.
//
// Solidity: function GetServiceBindingByIndex(uint256 _index) view returns((string,string,string,string,uint256) binding)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBindingByIndex(_index *big.Int) (iServiceMarketExServiceBinding, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingByIndex(&_IServiceMarketEx.CallOpts, _index)
}

// GetServiceBindingCount is a free data retrieval call binding the contract method 0xf8bdb54a.
//
// Solidity: function GetServiceBindingCount() view returns(uint256)
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
// Solidity: function GetServiceBindingCount() view returns(uint256)
func (_IServiceMarketEx *IServiceMarketExSession) GetServiceBindingCount() (*big.Int, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingCount(&_IServiceMarketEx.CallOpts)
}

// GetServiceBindingCount is a free data retrieval call binding the contract method 0xf8bdb54a.
//
// Solidity: function GetServiceBindingCount() view returns(uint256)
func (_IServiceMarketEx *IServiceMarketExCallerSession) GetServiceBindingCount() (*big.Int, error) {
	return _IServiceMarketEx.Contract.GetServiceBindingCount(&_IServiceMarketEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
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
// Solidity: function isOwner() view returns(bool)
func (_IServiceMarketEx *IServiceMarketExSession) IsOwner() (bool, error) {
	return _IServiceMarketEx.Contract.IsOwner(&_IServiceMarketEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IServiceMarketEx *IServiceMarketExCallerSession) IsOwner() (bool, error) {
	return _IServiceMarketEx.Contract.IsOwner(&_IServiceMarketEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
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
// Solidity: function owner() view returns(address)
func (_IServiceMarketEx *IServiceMarketExSession) Owner() (common.Address, error) {
	return _IServiceMarketEx.Contract.Owner(&_IServiceMarketEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IServiceMarketEx *IServiceMarketExCallerSession) Owner() (common.Address, error) {
	return _IServiceMarketEx.Contract.Owner(&_IServiceMarketEx.CallOpts)
}

// AddServiceBinding is a paid mutator transaction binding the contract method 0x92ce31c4.
//
// Solidity: function addServiceBinding(string _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) AddServiceBinding(opts *bind.TransactOpts, _serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.Transact(opts, "addServiceBinding", _serviceName, _schemas, _provider, _serviceFee, _qos)
}

// AddServiceBinding is a paid mutator transaction binding the contract method 0x92ce31c4.
//
// Solidity: function addServiceBinding(string _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExSession) AddServiceBinding(_serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AddServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _schemas, _provider, _serviceFee, _qos)
}

// AddServiceBinding is a paid mutator transaction binding the contract method 0x92ce31c4.
//
// Solidity: function addServiceBinding(string _serviceName, string _schemas, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) AddServiceBinding(_serviceName string, _schemas string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.AddServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _schemas, _provider, _serviceFee, _qos)
}

// RemoveServiceBinding is a paid mutator transaction binding the contract method 0xbce9f63f.
//
// Solidity: function removeServiceBinding(string _serviceName) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) RemoveServiceBinding(opts *bind.TransactOpts, _serviceName string) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.Transact(opts, "removeServiceBinding", _serviceName)
}

// RemoveServiceBinding is a paid mutator transaction binding the contract method 0xbce9f63f.
//
// Solidity: function removeServiceBinding(string _serviceName) returns()
func (_IServiceMarketEx *IServiceMarketExSession) RemoveServiceBinding(_serviceName string) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.RemoveServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName)
}

// RemoveServiceBinding is a paid mutator transaction binding the contract method 0xbce9f63f.
//
// Solidity: function removeServiceBinding(string _serviceName) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) RemoveServiceBinding(_serviceName string) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.RemoveServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceMarketEx *IServiceMarketExSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.TransferOwnership(&_IServiceMarketEx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.TransferOwnership(&_IServiceMarketEx.TransactOpts, newOwner)
}

// UpdateServiceBinding is a paid mutator transaction binding the contract method 0x6e29ea60.
//
// Solidity: function updateServiceBinding(string _serviceName, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactor) UpdateServiceBinding(opts *bind.TransactOpts, _serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.contract.Transact(opts, "updateServiceBinding", _serviceName, _provider, _serviceFee, _qos)
}

// UpdateServiceBinding is a paid mutator transaction binding the contract method 0x6e29ea60.
//
// Solidity: function updateServiceBinding(string _serviceName, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExSession) UpdateServiceBinding(_serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.UpdateServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _provider, _serviceFee, _qos)
}

// UpdateServiceBinding is a paid mutator transaction binding the contract method 0x6e29ea60.
//
// Solidity: function updateServiceBinding(string _serviceName, string _provider, string _serviceFee, uint256 _qos) returns()
func (_IServiceMarketEx *IServiceMarketExTransactorSession) UpdateServiceBinding(_serviceName string, _provider string, _serviceFee string, _qos *big.Int) (*types.Transaction, error) {
	return _IServiceMarketEx.Contract.UpdateServiceBinding(&_IServiceMarketEx.TransactOpts, _serviceName, _provider, _serviceFee, _qos)
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
