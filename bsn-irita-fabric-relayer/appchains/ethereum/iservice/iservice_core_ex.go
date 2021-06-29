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

// iServiceCoreExResponse is an auto generated low-level Go binding around an user-defined struct.
type iServiceCoreExResponse struct {
	ErrMsg      string
	Output      string
	IcRequestID string
}

// IServiceCoreExABI is the input ABI used to generate the binding from.
const IServiceCoreExABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_relayer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"name\":\"ServiceInvoked\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"_callbackFunction\",\"type\":\"bytes4\"}],\"name\":\"callService\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"requestID\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"}],\"name\":\"getResponse\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"errMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"output\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"icRequestID\",\"type\":\"string\"}],\"internalType\":\"structiServiceCoreEx.Response\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_errMsg\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_icRequestID\",\"type\":\"string\"}],\"name\":\"setResponse\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IServiceCoreEx is an auto generated Go binding around an Ethereum contract.
type IServiceCoreEx struct {
	IServiceCoreExCaller     // Read-only binding to the contract
	IServiceCoreExTransactor // Write-only binding to the contract
	IServiceCoreExFilterer   // Log filterer for contract events
}

// IServiceCoreExCaller is an auto generated read-only Go binding around an Ethereum contract.
type IServiceCoreExCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IServiceCoreExTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IServiceCoreExFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IServiceCoreExSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IServiceCoreExSession struct {
	Contract     *IServiceCoreEx   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IServiceCoreExCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IServiceCoreExCallerSession struct {
	Contract *IServiceCoreExCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IServiceCoreExTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IServiceCoreExTransactorSession struct {
	Contract     *IServiceCoreExTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IServiceCoreExRaw is an auto generated low-level Go binding around an Ethereum contract.
type IServiceCoreExRaw struct {
	Contract *IServiceCoreEx // Generic contract binding to access the raw methods on
}

// IServiceCoreExCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IServiceCoreExCallerRaw struct {
	Contract *IServiceCoreExCaller // Generic read-only contract binding to access the raw methods on
}

// IServiceCoreExTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
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
func (_IServiceCoreEx *IServiceCoreExRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.IServiceCoreExTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceCoreEx *IServiceCoreExRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
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
func (_IServiceCoreEx *IServiceCoreExTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IServiceCoreEx *IServiceCoreExTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.contract.Transact(opts, method, params...)
}

// GetResponse is a free data retrieval call binding the contract method 0x415a2bc1.
//
// Solidity: function getResponse(bytes32 _requestID) view returns((string,string,string))
func (_IServiceCoreEx *IServiceCoreExCaller) GetResponse(opts *bind.CallOpts, _requestID [32]byte) (iServiceCoreExResponse, error) {
	var (
		ret0 = new(iServiceCoreExResponse)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "getResponse", _requestID)
	return *ret0, err
}

// GetResponse is a free data retrieval call binding the contract method 0x415a2bc1.
//
// Solidity: function getResponse(bytes32 _requestID) view returns((string,string,string))
func (_IServiceCoreEx *IServiceCoreExSession) GetResponse(_requestID [32]byte) (iServiceCoreExResponse, error) {
	return _IServiceCoreEx.Contract.GetResponse(&_IServiceCoreEx.CallOpts, _requestID)
}

// GetResponse is a free data retrieval call binding the contract method 0x415a2bc1.
//
// Solidity: function getResponse(bytes32 _requestID) view returns((string,string,string))
func (_IServiceCoreEx *IServiceCoreExCallerSession) GetResponse(_requestID [32]byte) (iServiceCoreExResponse, error) {
	return _IServiceCoreEx.Contract.GetResponse(&_IServiceCoreEx.CallOpts, _requestID)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IServiceCoreEx *IServiceCoreExCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "isOwner")
	return *ret0, err
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) IsOwner() (bool, error) {
	return _IServiceCoreEx.Contract.IsOwner(&_IServiceCoreEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_IServiceCoreEx *IServiceCoreExCallerSession) IsOwner() (bool, error) {
	return _IServiceCoreEx.Contract.IsOwner(&_IServiceCoreEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IServiceCoreEx *IServiceCoreExCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IServiceCoreEx *IServiceCoreExSession) Owner() (common.Address, error) {
	return _IServiceCoreEx.Contract.Owner(&_IServiceCoreEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_IServiceCoreEx *IServiceCoreExCallerSession) Owner() (common.Address, error) {
	return _IServiceCoreEx.Contract.Owner(&_IServiceCoreEx.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_IServiceCoreEx *IServiceCoreExCaller) RequestCount(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IServiceCoreEx.contract.Call(opts, out, "requestCount")
	return *ret0, err
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_IServiceCoreEx *IServiceCoreExSession) RequestCount() (*big.Int, error) {
	return _IServiceCoreEx.Contract.RequestCount(&_IServiceCoreEx.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() view returns(uint256)
func (_IServiceCoreEx *IServiceCoreExCallerSession) RequestCount() (*big.Int, error) {
	return _IServiceCoreEx.Contract.RequestCount(&_IServiceCoreEx.CallOpts)
}

// CallService is a paid mutator transaction binding the contract method 0xcfb21ef6.
//
// Solidity: function callService(string _serviceName, string _input, uint256 _timeout, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExTransactor) CallService(opts *bind.TransactOpts, _serviceName string, _input string, _timeout *big.Int, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.Transact(opts, "callService", _serviceName, _input, _timeout, _callbackAddress, _callbackFunction)
}

// CallService is a paid mutator transaction binding the contract method 0xcfb21ef6.
//
// Solidity: function callService(string _serviceName, string _input, uint256 _timeout, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExSession) CallService(_serviceName string, _input string, _timeout *big.Int, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.CallService(&_IServiceCoreEx.TransactOpts, _serviceName, _input, _timeout, _callbackAddress, _callbackFunction)
}

// CallService is a paid mutator transaction binding the contract method 0xcfb21ef6.
//
// Solidity: function callService(string _serviceName, string _input, uint256 _timeout, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) CallService(_serviceName string, _input string, _timeout *big.Int, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.CallService(&_IServiceCoreEx.TransactOpts, _serviceName, _input, _timeout, _callbackAddress, _callbackFunction)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) SetRelayer(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setRelayer", _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExSession) SetRelayer(_address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.SetRelayer(&_IServiceCoreEx.TransactOpts, _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetRelayer(_address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.SetRelayer(&_IServiceCoreEx.TransactOpts, _address)
}

// SetResponse is a paid mutator transaction binding the contract method 0x4bcee3e7.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output, string _icRequestID) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactor) SetResponse(opts *bind.TransactOpts, _requestID [32]byte, _errMsg string, _output string, _icRequestID string) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setResponse", _requestID, _errMsg, _output, _icRequestID)
}

// SetResponse is a paid mutator transaction binding the contract method 0x4bcee3e7.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output, string _icRequestID) returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) SetResponse(_requestID [32]byte, _errMsg string, _output string, _icRequestID string) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.SetResponse(&_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output, _icRequestID)
}

// SetResponse is a paid mutator transaction binding the contract method 0x4bcee3e7.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output, string _icRequestID) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetResponse(_requestID [32]byte, _errMsg string, _output string, _icRequestID string) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.SetResponse(&_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output, _icRequestID)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.TransferOwnership(&_IServiceCoreEx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.TransferOwnership(&_IServiceCoreEx.TransactOpts, newOwner)
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

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
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

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
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

// FilterServiceInvoked is a free log retrieval operation binding the contract event 0x22be23eea0c19f673d933857312686db7270c76305be3a112b4e5d21630ddc82.
//
// Solidity: event ServiceInvoked(bytes32 _requestID, string _serviceName, string _input, uint256 _timeout)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterServiceInvoked(opts *bind.FilterOpts) (*IServiceCoreExServiceInvokedIterator, error) {

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "ServiceInvoked")
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExServiceInvokedIterator{contract: _IServiceCoreEx.contract, event: "ServiceInvoked", logs: logs, sub: sub}, nil
}

// WatchServiceInvoked is a free log subscription operation binding the contract event 0x22be23eea0c19f673d933857312686db7270c76305be3a112b4e5d21630ddc82.
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

// ParseServiceInvoked is a log parse operation binding the contract event 0x22be23eea0c19f673d933857312686db7270c76305be3a112b4e5d21630ddc82.
//
// Solidity: event ServiceInvoked(bytes32 _requestID, string _serviceName, string _input, uint256 _timeout)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseServiceInvoked(log types.Log) (*IServiceCoreExServiceInvoked, error) {
	event := new(IServiceCoreExServiceInvoked)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "ServiceInvoked", log); err != nil {
		return nil, err
	}
	return event, nil
}
