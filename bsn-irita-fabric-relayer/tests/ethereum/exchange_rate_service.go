// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ethereum

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

// ExchangeRateServiceABI is the input ABI used to generate the binding from.
const ExchangeRateServiceABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_iServiceContract\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_serviceName\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_input\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"_timeout\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"_rate\",\"type\":\"string\"}],\"name\":\"RequestResponded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"}],\"name\":\"RequestSent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"_requestID\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"_output\",\"type\":\"string\"}],\"name\":\"onResponse\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rate\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sendRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ExchangeRateService is an auto generated Go binding around an Ethereum contract.
type ExchangeRateService struct {
	ExchangeRateServiceCaller     // Read-only binding to the contract
	ExchangeRateServiceTransactor // Write-only binding to the contract
	ExchangeRateServiceFilterer   // Log filterer for contract events
}

// ExchangeRateServiceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExchangeRateServiceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeRateServiceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExchangeRateServiceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeRateServiceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExchangeRateServiceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExchangeRateServiceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExchangeRateServiceSession struct {
	Contract     *ExchangeRateService // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ExchangeRateServiceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExchangeRateServiceCallerSession struct {
	Contract *ExchangeRateServiceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// ExchangeRateServiceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExchangeRateServiceTransactorSession struct {
	Contract     *ExchangeRateServiceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ExchangeRateServiceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExchangeRateServiceRaw struct {
	Contract *ExchangeRateService // Generic contract binding to access the raw methods on
}

// ExchangeRateServiceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExchangeRateServiceCallerRaw struct {
	Contract *ExchangeRateServiceCaller // Generic read-only contract binding to access the raw methods on
}

// ExchangeRateServiceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExchangeRateServiceTransactorRaw struct {
	Contract *ExchangeRateServiceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExchangeRateService creates a new instance of ExchangeRateService, bound to a specific deployed contract.
func NewExchangeRateService(address common.Address, backend bind.ContractBackend) (*ExchangeRateService, error) {
	contract, err := bindExchangeRateService(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ExchangeRateService{ExchangeRateServiceCaller: ExchangeRateServiceCaller{contract: contract}, ExchangeRateServiceTransactor: ExchangeRateServiceTransactor{contract: contract}, ExchangeRateServiceFilterer: ExchangeRateServiceFilterer{contract: contract}}, nil
}

// NewExchangeRateServiceCaller creates a new read-only instance of ExchangeRateService, bound to a specific deployed contract.
func NewExchangeRateServiceCaller(address common.Address, caller bind.ContractCaller) (*ExchangeRateServiceCaller, error) {
	contract, err := bindExchangeRateService(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeRateServiceCaller{contract: contract}, nil
}

// NewExchangeRateServiceTransactor creates a new write-only instance of ExchangeRateService, bound to a specific deployed contract.
func NewExchangeRateServiceTransactor(address common.Address, transactor bind.ContractTransactor) (*ExchangeRateServiceTransactor, error) {
	contract, err := bindExchangeRateService(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExchangeRateServiceTransactor{contract: contract}, nil
}

// NewExchangeRateServiceFilterer creates a new log filterer instance of ExchangeRateService, bound to a specific deployed contract.
func NewExchangeRateServiceFilterer(address common.Address, filterer bind.ContractFilterer) (*ExchangeRateServiceFilterer, error) {
	contract, err := bindExchangeRateService(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExchangeRateServiceFilterer{contract: contract}, nil
}

// bindExchangeRateService binds a generic wrapper to an already deployed contract.
func bindExchangeRateService(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExchangeRateServiceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeRateService *ExchangeRateServiceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeRateService.Contract.ExchangeRateServiceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeRateService *ExchangeRateServiceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeRateService.Contract.ExchangeRateServiceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeRateService *ExchangeRateServiceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeRateService.Contract.ExchangeRateServiceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ExchangeRateService *ExchangeRateServiceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ExchangeRateService.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ExchangeRateService *ExchangeRateServiceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeRateService.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ExchangeRateService *ExchangeRateServiceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ExchangeRateService.Contract.contract.Transact(opts, method, params...)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(string)
func (_ExchangeRateService *ExchangeRateServiceCaller) Rate(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _ExchangeRateService.contract.Call(opts, out, "rate")
	return *ret0, err
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(string)
func (_ExchangeRateService *ExchangeRateServiceSession) Rate() (string, error) {
	return _ExchangeRateService.Contract.Rate(&_ExchangeRateService.CallOpts)
}

// Rate is a free data retrieval call binding the contract method 0x2c4e722e.
//
// Solidity: function rate() view returns(string)
func (_ExchangeRateService *ExchangeRateServiceCallerSession) Rate() (string, error) {
	return _ExchangeRateService.Contract.Rate(&_ExchangeRateService.CallOpts)
}

// OnResponse is a paid mutator transaction binding the contract method 0x78a78112.
//
// Solidity: function onResponse(bytes32 _requestID, string _output) returns()
func (_ExchangeRateService *ExchangeRateServiceTransactor) OnResponse(opts *bind.TransactOpts, _requestID [32]byte, _output string) (*types.Transaction, error) {
	return _ExchangeRateService.contract.Transact(opts, "onResponse", _requestID, _output)
}

// OnResponse is a paid mutator transaction binding the contract method 0x78a78112.
//
// Solidity: function onResponse(bytes32 _requestID, string _output) returns()
func (_ExchangeRateService *ExchangeRateServiceSession) OnResponse(_requestID [32]byte, _output string) (*types.Transaction, error) {
	return _ExchangeRateService.Contract.OnResponse(&_ExchangeRateService.TransactOpts, _requestID, _output)
}

// OnResponse is a paid mutator transaction binding the contract method 0x78a78112.
//
// Solidity: function onResponse(bytes32 _requestID, string _output) returns()
func (_ExchangeRateService *ExchangeRateServiceTransactorSession) OnResponse(_requestID [32]byte, _output string) (*types.Transaction, error) {
	return _ExchangeRateService.Contract.OnResponse(&_ExchangeRateService.TransactOpts, _requestID, _output)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bee2f2a.
//
// Solidity: function sendRequest() returns()
func (_ExchangeRateService *ExchangeRateServiceTransactor) SendRequest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ExchangeRateService.contract.Transact(opts, "sendRequest")
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bee2f2a.
//
// Solidity: function sendRequest() returns()
func (_ExchangeRateService *ExchangeRateServiceSession) SendRequest() (*types.Transaction, error) {
	return _ExchangeRateService.Contract.SendRequest(&_ExchangeRateService.TransactOpts)
}

// SendRequest is a paid mutator transaction binding the contract method 0x1bee2f2a.
//
// Solidity: function sendRequest() returns()
func (_ExchangeRateService *ExchangeRateServiceTransactorSession) SendRequest() (*types.Transaction, error) {
	return _ExchangeRateService.Contract.SendRequest(&_ExchangeRateService.TransactOpts)
}

// ExchangeRateServiceRequestRespondedIterator is returned from FilterRequestResponded and is used to iterate over the raw logs and unpacked data for RequestResponded events raised by the ExchangeRateService contract.
type ExchangeRateServiceRequestRespondedIterator struct {
	Event *ExchangeRateServiceRequestResponded // Event containing the contract specifics and raw log

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
func (it *ExchangeRateServiceRequestRespondedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeRateServiceRequestResponded)
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
		it.Event = new(ExchangeRateServiceRequestResponded)
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
func (it *ExchangeRateServiceRequestRespondedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeRateServiceRequestRespondedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeRateServiceRequestResponded represents a RequestResponded event raised by the ExchangeRateService contract.
type ExchangeRateServiceRequestResponded struct {
	RequestID [32]byte
	Rate      string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestResponded is a free log retrieval operation binding the contract event 0x5d337c667731d0f2a5d07da5189c409af4260a06a0d5662de9e5139750902468.
//
// Solidity: event RequestResponded(bytes32 _requestID, string _rate)
func (_ExchangeRateService *ExchangeRateServiceFilterer) FilterRequestResponded(opts *bind.FilterOpts) (*ExchangeRateServiceRequestRespondedIterator, error) {

	logs, sub, err := _ExchangeRateService.contract.FilterLogs(opts, "RequestResponded")
	if err != nil {
		return nil, err
	}
	return &ExchangeRateServiceRequestRespondedIterator{contract: _ExchangeRateService.contract, event: "RequestResponded", logs: logs, sub: sub}, nil
}

// WatchRequestResponded is a free log subscription operation binding the contract event 0x5d337c667731d0f2a5d07da5189c409af4260a06a0d5662de9e5139750902468.
//
// Solidity: event RequestResponded(bytes32 _requestID, string _rate)
func (_ExchangeRateService *ExchangeRateServiceFilterer) WatchRequestResponded(opts *bind.WatchOpts, sink chan<- *ExchangeRateServiceRequestResponded) (event.Subscription, error) {

	logs, sub, err := _ExchangeRateService.contract.WatchLogs(opts, "RequestResponded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeRateServiceRequestResponded)
				if err := _ExchangeRateService.contract.UnpackLog(event, "RequestResponded", log); err != nil {
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

// ParseRequestResponded is a log parse operation binding the contract event 0x5d337c667731d0f2a5d07da5189c409af4260a06a0d5662de9e5139750902468.
//
// Solidity: event RequestResponded(bytes32 _requestID, string _rate)
func (_ExchangeRateService *ExchangeRateServiceFilterer) ParseRequestResponded(log types.Log) (*ExchangeRateServiceRequestResponded, error) {
	event := new(ExchangeRateServiceRequestResponded)
	if err := _ExchangeRateService.contract.UnpackLog(event, "RequestResponded", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ExchangeRateServiceRequestSentIterator is returned from FilterRequestSent and is used to iterate over the raw logs and unpacked data for RequestSent events raised by the ExchangeRateService contract.
type ExchangeRateServiceRequestSentIterator struct {
	Event *ExchangeRateServiceRequestSent // Event containing the contract specifics and raw log

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
func (it *ExchangeRateServiceRequestSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ExchangeRateServiceRequestSent)
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
		it.Event = new(ExchangeRateServiceRequestSent)
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
func (it *ExchangeRateServiceRequestSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ExchangeRateServiceRequestSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ExchangeRateServiceRequestSent represents a RequestSent event raised by the ExchangeRateService contract.
type ExchangeRateServiceRequestSent struct {
	RequestID [32]byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRequestSent is a free log retrieval operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 _requestID)
func (_ExchangeRateService *ExchangeRateServiceFilterer) FilterRequestSent(opts *bind.FilterOpts) (*ExchangeRateServiceRequestSentIterator, error) {

	logs, sub, err := _ExchangeRateService.contract.FilterLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return &ExchangeRateServiceRequestSentIterator{contract: _ExchangeRateService.contract, event: "RequestSent", logs: logs, sub: sub}, nil
}

// WatchRequestSent is a free log subscription operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 _requestID)
func (_ExchangeRateService *ExchangeRateServiceFilterer) WatchRequestSent(opts *bind.WatchOpts, sink chan<- *ExchangeRateServiceRequestSent) (event.Subscription, error) {

	logs, sub, err := _ExchangeRateService.contract.WatchLogs(opts, "RequestSent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ExchangeRateServiceRequestSent)
				if err := _ExchangeRateService.contract.UnpackLog(event, "RequestSent", log); err != nil {
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

// ParseRequestSent is a log parse operation binding the contract event 0x1131472297a800fee664d1d89cfa8f7676ff07189ecc53f80bbb5f4969099db8.
//
// Solidity: event RequestSent(bytes32 _requestID)
func (_ExchangeRateService *ExchangeRateServiceFilterer) ParseRequestSent(log types.Log) (*ExchangeRateServiceRequestSent, error) {
	event := new(ExchangeRateServiceRequestSent)
	if err := _ExchangeRateService.contract.UnpackLog(event, "RequestSent", log); err != nil {
		return nil, err
	}
	return event, nil
}
