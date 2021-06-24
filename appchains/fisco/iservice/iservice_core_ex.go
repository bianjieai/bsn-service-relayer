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
const IServiceCoreExABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"name\":\"_errMsg\",\"type\":\"string\"},{\"name\":\"_output\",\"type\":\"string\"}],\"name\":\"setResponse\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"requestCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_endpointInfo\",\"type\":\"string\"},{\"name\":\"_method\",\"type\":\"string\"},{\"name\":\"_callData\",\"type\":\"bytes\"},{\"name\":\"_callbackAddress\",\"type\":\"address\"},{\"name\":\"_callbackFunction\",\"type\":\"bytes4\"}],\"name\":\"sendRequest\",\"outputs\":[{\"name\":\"requestID\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setRelayer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"name\":\"_callData\",\"type\":\"bytes\"}],\"name\":\"callService\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"relayer\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_icRequestID\",\"type\":\"bytes32\"},{\"name\":\"_endpointAddress\",\"type\":\"address\"},{\"name\":\"_callData\",\"type\":\"bytes\"}],\"name\":\"callService\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_relayer\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_endpointInfo\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_method\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"_methodAndArgs\",\"type\":\"bytes\"},{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"}],\"name\":\"CrossChainRequestSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_icRequestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_result\",\"type\":\"bytes\"}],\"name\":\"CrossChainResponseSent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

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

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
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

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) IsOwner() (bool, error) {
	return _IServiceCoreEx.Contract.IsOwner(&_IServiceCoreEx.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() constant returns(bool)
func (_IServiceCoreEx *IServiceCoreExCallerSession) IsOwner() (bool, error) {
	return _IServiceCoreEx.Contract.IsOwner(&_IServiceCoreEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExSession) Owner() (common.Address, error) {
	return _IServiceCoreEx.Contract.Owner(&_IServiceCoreEx.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCallerSession) Owner() (common.Address, error) {
	return _IServiceCoreEx.Contract.Owner(&_IServiceCoreEx.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
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

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExSession) Relayer() (common.Address, error) {
	return _IServiceCoreEx.Contract.Relayer(&_IServiceCoreEx.CallOpts)
}

// Relayer is a free data retrieval call binding the contract method 0x8406c079.
//
// Solidity: function relayer() constant returns(address)
func (_IServiceCoreEx *IServiceCoreExCallerSession) Relayer() (common.Address, error) {
	return _IServiceCoreEx.Contract.Relayer(&_IServiceCoreEx.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
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

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() constant returns(uint256)
func (_IServiceCoreEx *IServiceCoreExSession) RequestCount() (*big.Int, error) {
	return _IServiceCoreEx.Contract.RequestCount(&_IServiceCoreEx.CallOpts)
}

// RequestCount is a free data retrieval call binding the contract method 0x5badbe4c.
//
// Solidity: function requestCount() constant returns(uint256)
func (_IServiceCoreEx *IServiceCoreExCallerSession) RequestCount() (*big.Int, error) {
	return _IServiceCoreEx.Contract.RequestCount(&_IServiceCoreEx.CallOpts)
}

// CallService is a paid mutator transaction binding the contract method 0x364c4b79.
//
// Solidity: function callService(bytes32 _requestID, bytes _callData) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactor) CallService(opts *bind.TransactOpts, _requestID [32]byte, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "callService", _requestID, _callData)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncCallService(handler func(*types.Receipt, error), opts *bind.TransactOpts, _requestID [32]byte, _callData []byte) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "callService", _requestID, _callData)
}

// CallService is a paid mutator transaction binding the contract method 0x364c4b79.
//
// Solidity: function callService(bytes32 _requestID, bytes _callData) returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) CallService(_requestID [32]byte, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.CallService(&_IServiceCoreEx.TransactOpts, _requestID, _callData)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncCallService(handler func(*types.Receipt, error), _requestID [32]byte, _callData []byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncCallService(handler, &_IServiceCoreEx.TransactOpts, _requestID, _callData)
}

// CallService is a paid mutator transaction binding the contract method 0x364c4b79.
//
// Solidity: function callService(bytes32 _requestID, bytes _callData) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) CallService(_requestID [32]byte, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.CallService(&_IServiceCoreEx.TransactOpts, _requestID, _callData)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncCallService(handler func(*types.Receipt, error), _requestID [32]byte, _callData []byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncCallService(handler, &_IServiceCoreEx.TransactOpts, _requestID, _callData)
}

// CallService0 is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _icRequestID, address _endpointAddress, bytes _callData) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) CallService0(opts *bind.TransactOpts, _icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "callService0", _icRequestID, _endpointAddress, _callData)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncCallService0(handler func(*types.Receipt, error), opts *bind.TransactOpts, _icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "callService0", _icRequestID, _endpointAddress, _callData)
}

// CallService0 is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _icRequestID, address _endpointAddress, bytes _callData) returns()
func (_IServiceCoreEx *IServiceCoreExSession) CallService0(_icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.CallService0(&_IServiceCoreEx.TransactOpts, _icRequestID, _endpointAddress, _callData)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncCallService0(handler func(*types.Receipt, error), _icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncCallService0(handler, &_IServiceCoreEx.TransactOpts, _icRequestID, _endpointAddress, _callData)
}

// CallService0 is a paid mutator transaction binding the contract method 0x75d219fa.
//
// Solidity: function callService(bytes32 _icRequestID, address _endpointAddress, bytes _callData) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) CallService0(_icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.CallService0(&_IServiceCoreEx.TransactOpts, _icRequestID, _endpointAddress, _callData)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncCallService0(handler func(*types.Receipt, error), _icRequestID [32]byte, _endpointAddress common.Address, _callData []byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncCallService0(handler, &_IServiceCoreEx.TransactOpts, _icRequestID, _endpointAddress, _callData)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bd75284.
//
// Solidity: function sendRequest(string _endpointInfo, string _method, bytes _callData, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExTransactor) SendRequest(opts *bind.TransactOpts, _endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "sendRequest", _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncSendRequest(handler func(*types.Receipt, error), opts *bind.TransactOpts, _endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "sendRequest", _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bd75284.
//
// Solidity: function sendRequest(string _endpointInfo, string _method, bytes _callData, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExSession) SendRequest(_endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SendRequest(&_IServiceCoreEx.TransactOpts, _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncSendRequest(handler func(*types.Receipt, error), _endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSendRequest(handler, &_IServiceCoreEx.TransactOpts, _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bd75284.
//
// Solidity: function sendRequest(string _endpointInfo, string _method, bytes _callData, address _callbackAddress, bytes4 _callbackFunction) returns(bytes32 requestID)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SendRequest(_endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SendRequest(&_IServiceCoreEx.TransactOpts, _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncSendRequest(handler func(*types.Receipt, error), _endpointInfo string, _method string, _callData []byte, _callbackAddress common.Address, _callbackFunction [4]byte) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSendRequest(handler, &_IServiceCoreEx.TransactOpts, _endpointInfo, _method, _callData, _callbackAddress, _callbackFunction)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) SetRelayer(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setRelayer", _address)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncSetRelayer(handler func(*types.Receipt, error), opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "setRelayer", _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExSession) SetRelayer(_address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetRelayer(&_IServiceCoreEx.TransactOpts, _address)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncSetRelayer(handler func(*types.Receipt, error), _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetRelayer(handler, &_IServiceCoreEx.TransactOpts, _address)
}

// SetRelayer is a paid mutator transaction binding the contract method 0x6548e9bc.
//
// Solidity: function setRelayer(address _address) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetRelayer(_address common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetRelayer(&_IServiceCoreEx.TransactOpts, _address)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncSetRelayer(handler func(*types.Receipt, error), _address common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetRelayer(handler, &_IServiceCoreEx.TransactOpts, _address)
}

// SetResponse is a paid mutator transaction binding the contract method 0x8892bb6a.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactor) SetResponse(opts *bind.TransactOpts, _requestID [32]byte, _errMsg string, _output string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "setResponse", _requestID, _errMsg, _output)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncSetResponse(handler func(*types.Receipt, error), opts *bind.TransactOpts, _requestID [32]byte, _errMsg string, _output string) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "setResponse", _requestID, _errMsg, _output)
}

// SetResponse is a paid mutator transaction binding the contract method 0x8892bb6a.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output) returns(bool)
func (_IServiceCoreEx *IServiceCoreExSession) SetResponse(_requestID [32]byte, _errMsg string, _output string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetResponse(&_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncSetResponse(handler func(*types.Receipt, error), _requestID [32]byte, _errMsg string, _output string) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetResponse(handler, &_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output)
}

// SetResponse is a paid mutator transaction binding the contract method 0x8892bb6a.
//
// Solidity: function setResponse(bytes32 _requestID, string _errMsg, string _output) returns(bool)
func (_IServiceCoreEx *IServiceCoreExTransactorSession) SetResponse(_requestID [32]byte, _errMsg string, _output string) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.SetResponse(&_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncSetResponse(handler func(*types.Receipt, error), _requestID [32]byte, _errMsg string, _output string) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncSetResponse(handler, &_IServiceCoreEx.TransactOpts, _requestID, _errMsg, _output)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.contract.Transact(opts, "transferOwnership", newOwner)
}

func (_IServiceCoreEx *IServiceCoreExTransactor) AsyncTransferOwnership(handler func(*types.Receipt, error), opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.contract.AsyncTransact(opts, handler, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.TransferOwnership(&_IServiceCoreEx.TransactOpts, newOwner)
}

func (_IServiceCoreEx *IServiceCoreExSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncTransferOwnership(handler, &_IServiceCoreEx.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IServiceCoreEx *IServiceCoreExTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, *types.Receipt, error) {
	return _IServiceCoreEx.Contract.TransferOwnership(&_IServiceCoreEx.TransactOpts, newOwner)
}

func (_IServiceCoreEx *IServiceCoreExTransactorSession) AsyncTransferOwnership(handler func(*types.Receipt, error), newOwner common.Address) (*types.Transaction, error) {
	return _IServiceCoreEx.Contract.AsyncTransferOwnership(handler, &_IServiceCoreEx.TransactOpts, newOwner)
}

// IServiceCoreExCrossChainRequestSentIterator is returned from FilterCrossChainRequestSent and is used to iterate over the raw logs and unpacked data for CrossChainRequestSent events raised by the IServiceCoreEx contract.
type IServiceCoreExCrossChainRequestSentIterator struct {
	Event *IServiceCoreExCrossChainRequestSent // Event containing the contract specifics and raw log

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
func (it *IServiceCoreExCrossChainRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceCoreExCrossChainRequestSent)
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
		it.Event = new(IServiceCoreExCrossChainRequestSent)
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
func (it *IServiceCoreExCrossChainRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceCoreExCrossChainRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceCoreExCrossChainRequestSent represents a CrossChainRequestSent event raised by the IServiceCoreEx contract.
type IServiceCoreExCrossChainRequestSent struct {
	RequestID     [32]byte
	EndpointInfo  string
	Method        string
	MethodAndArgs []byte
	Sender        common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCrossChainRequestSent is a free log retrieval operation binding the contract event 0x0faa824f1be7109f16e32fb016edbc264ea83711298f4b3219e1b8fa5aaa8cd7.
//
// Solidity: event CrossChainRequestSent(bytes32 _requestID, string _endpointInfo, string _method, bytes _methodAndArgs, address _sender)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterCrossChainRequestSent(opts *bind.FilterOpts) (*IServiceCoreExCrossChainRequestSentIterator, error) {

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "CrossChainRequestSent")
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExCrossChainRequestSentIterator{contract: _IServiceCoreEx.contract, event: "CrossChainRequestSent", logs: logs, sub: sub}, nil
}

// WatchCrossChainRequestSent is a free log subscription operation binding the contract event 0x0faa824f1be7109f16e32fb016edbc264ea83711298f4b3219e1b8fa5aaa8cd7.
//
// Solidity: event CrossChainRequestSent(bytes32 _requestID, string _endpointInfo, string _method, bytes _methodAndArgs, address _sender)
func (_IServiceCoreEx *IServiceCoreExFilterer) WatchCrossChainRequestSent(opts *bind.WatchOpts, sink chan<- *IServiceCoreExCrossChainRequestSent) (event.Subscription, error) {

	logs, sub, err := _IServiceCoreEx.contract.WatchLogs(opts, "CrossChainRequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceCoreExCrossChainRequestSent)
				if err := _IServiceCoreEx.contract.UnpackLog(event, "CrossChainRequestSent", log); err != nil {
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

// ParseCrossChainRequestSent is a log parse operation binding the contract event 0x0faa824f1be7109f16e32fb016edbc264ea83711298f4b3219e1b8fa5aaa8cd7.
//
// Solidity: event CrossChainRequestSent(bytes32 _requestID, string _endpointInfo, string _method, bytes _methodAndArgs, address _sender)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseCrossChainRequestSent(log types.Log) (*IServiceCoreExCrossChainRequestSent, error) {
	event := new(IServiceCoreExCrossChainRequestSent)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "CrossChainRequestSent", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IServiceCoreExCrossChainResponseSentIterator is returned from FilterCrossChainResponseSent and is used to iterate over the raw logs and unpacked data for CrossChainResponseSent events raised by the IServiceCoreEx contract.
type IServiceCoreExCrossChainResponseSentIterator struct {
	Event *IServiceCoreExCrossChainResponseSent // Event containing the contract specifics and raw log

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
func (it *IServiceCoreExCrossChainResponseSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IServiceCoreExCrossChainResponseSent)
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
		it.Event = new(IServiceCoreExCrossChainResponseSent)
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
func (it *IServiceCoreExCrossChainResponseSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IServiceCoreExCrossChainResponseSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IServiceCoreExCrossChainResponseSent represents a CrossChainResponseSent event raised by the IServiceCoreEx contract.
type IServiceCoreExCrossChainResponseSent struct {
	IcRequestID [32]byte
	Result      []byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterCrossChainResponseSent is a free log retrieval operation binding the contract event 0xdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c.
//
// Solidity: event CrossChainResponseSent(bytes32 _icRequestID, bytes _result)
func (_IServiceCoreEx *IServiceCoreExFilterer) FilterCrossChainResponseSent(opts *bind.FilterOpts) (*IServiceCoreExCrossChainResponseSentIterator, error) {

	logs, sub, err := _IServiceCoreEx.contract.FilterLogs(opts, "CrossChainResponseSent")
	if err != nil {
		return nil, err
	}
	return &IServiceCoreExCrossChainResponseSentIterator{contract: _IServiceCoreEx.contract, event: "CrossChainResponseSent", logs: logs, sub: sub}, nil
}

// WatchCrossChainResponseSent is a free log subscription operation binding the contract event 0xdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c.
//
// Solidity: event CrossChainResponseSent(bytes32 _icRequestID, bytes _result)
func (_IServiceCoreEx *IServiceCoreExFilterer) WatchCrossChainResponseSent(opts *bind.WatchOpts, sink chan<- *IServiceCoreExCrossChainResponseSent) (event.Subscription, error) {

	logs, sub, err := _IServiceCoreEx.contract.WatchLogs(opts, "CrossChainResponseSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IServiceCoreExCrossChainResponseSent)
				if err := _IServiceCoreEx.contract.UnpackLog(event, "CrossChainResponseSent", log); err != nil {
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

// ParseCrossChainResponseSent is a log parse operation binding the contract event 0xdd9b7291a69fc50c9dbd1ee123efce8f2f5086ebe46cea5855a5331b2eb6df4c.
//
// Solidity: event CrossChainResponseSent(bytes32 _icRequestID, bytes _result)
func (_IServiceCoreEx *IServiceCoreExFilterer) ParseCrossChainResponseSent(log types.Log) (*IServiceCoreExCrossChainResponseSent, error) {
	event := new(IServiceCoreExCrossChainResponseSent)
	if err := _IServiceCoreEx.contract.UnpackLog(event, "CrossChainResponseSent", log); err != nil {
		return nil, err
	}
	return event, nil
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
