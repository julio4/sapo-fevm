// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// EventsMetaData contains all meta data concerning the Events contract.
var EventsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"number\",\"type\":\"int256\"}],\"name\":\"NumberEvent\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"nb\",\"type\":\"int256\"}],\"name\":\"emitNb\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507fb7b401727864049638c23b3bf52321ea398b0fd65c72e465daf5c8b4564d7583600060405161008191906100dd565b60405180910390a16100f8565b6000819050919050565b6000819050919050565b6000819050919050565b60006100c76100c26100bd8461008e565b6100a2565b610098565b9050919050565b6100d7816100ac565b82525050565b60006020820190506100f260008301846100ce565b92915050565b610259806101076000396000f3fe608060405234801561001057600080fd5b506004361061002b5760003560e01c806341367af214610030575b600080fd5b61004a6004803603810190610045919061014f565b61004c565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146100da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016100d1906101d9565b60405180910390fd5b7fb7b401727864049638c23b3bf52321ea398b0fd65c72e465daf5c8b4564d7583816040516101099190610208565b60405180910390a150565b600080fd5b6000819050919050565b61012c81610119565b811461013757600080fd5b50565b60008135905061014981610123565b92915050565b60006020828403121561016557610164610114565b5b60006101738482850161013a565b91505092915050565b600082825260208201905092915050565b7f43616c6c6572206973206e6f74206f776e657200000000000000000000000000600082015250565b60006101c360138361017c565b91506101ce8261018d565b602082019050919050565b600060208201905081810360008301526101f2816101b6565b9050919050565b61020281610119565b82525050565b600060208201905061021d60008301846101f9565b9291505056fea26469706673582212203b65ce178f4eb70038c7c1aadbb007532f4b4bb41de0c214f69c6cb6778345a664736f6c63430008110033",
}

// EventsABI is the input ABI used to generate the binding from.
// Deprecated: Use EventsMetaData.ABI instead.
var EventsABI = EventsMetaData.ABI

// EventsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EventsMetaData.Bin instead.
var EventsBin = EventsMetaData.Bin

// DeployEvents deploys a new Ethereum contract, binding an instance of Events to it.
func DeployEvents(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Events, error) {
	parsed, err := EventsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EventsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// Events is an auto generated Go binding around an Ethereum contract.
type Events struct {
	EventsCaller     // Read-only binding to the contract
	EventsTransactor // Write-only binding to the contract
	EventsFilterer   // Log filterer for contract events
}

// EventsCaller is an auto generated read-only Go binding around an Ethereum contract.
type EventsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EventsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EventsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EventsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EventsSession struct {
	Contract     *Events           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EventsCallerSession struct {
	Contract *EventsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EventsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EventsTransactorSession struct {
	Contract     *EventsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EventsRaw is an auto generated low-level Go binding around an Ethereum contract.
type EventsRaw struct {
	Contract *Events // Generic contract binding to access the raw methods on
}

// EventsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EventsCallerRaw struct {
	Contract *EventsCaller // Generic read-only contract binding to access the raw methods on
}

// EventsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EventsTransactorRaw struct {
	Contract *EventsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEvents creates a new instance of Events, bound to a specific deployed contract.
func NewEvents(address common.Address, backend bind.ContractBackend) (*Events, error) {
	contract, err := bindEvents(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Events{EventsCaller: EventsCaller{contract: contract}, EventsTransactor: EventsTransactor{contract: contract}, EventsFilterer: EventsFilterer{contract: contract}}, nil
}

// NewEventsCaller creates a new read-only instance of Events, bound to a specific deployed contract.
func NewEventsCaller(address common.Address, caller bind.ContractCaller) (*EventsCaller, error) {
	contract, err := bindEvents(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EventsCaller{contract: contract}, nil
}

// NewEventsTransactor creates a new write-only instance of Events, bound to a specific deployed contract.
func NewEventsTransactor(address common.Address, transactor bind.ContractTransactor) (*EventsTransactor, error) {
	contract, err := bindEvents(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EventsTransactor{contract: contract}, nil
}

// NewEventsFilterer creates a new log filterer instance of Events, bound to a specific deployed contract.
func NewEventsFilterer(address common.Address, filterer bind.ContractFilterer) (*EventsFilterer, error) {
	contract, err := bindEvents(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EventsFilterer{contract: contract}, nil
}

// bindEvents binds a generic wrapper to an already deployed contract.
func bindEvents(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EventsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Events.Contract.EventsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.EventsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Events *EventsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Events.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Events *EventsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Events.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Events *EventsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Events.Contract.contract.Transact(opts, method, params...)
}

// EmitNb is a paid mutator transaction binding the contract method 0x41367af2.
//
// Solidity: function emitNb(int256 nb) returns()
func (_Events *EventsTransactor) EmitNb(opts *bind.TransactOpts, nb *big.Int) (*types.Transaction, error) {
	return _Events.contract.Transact(opts, "emitNb", nb)
}

// EmitNb is a paid mutator transaction binding the contract method 0x41367af2.
//
// Solidity: function emitNb(int256 nb) returns()
func (_Events *EventsSession) EmitNb(nb *big.Int) (*types.Transaction, error) {
	return _Events.Contract.EmitNb(&_Events.TransactOpts, nb)
}

// EmitNb is a paid mutator transaction binding the contract method 0x41367af2.
//
// Solidity: function emitNb(int256 nb) returns()
func (_Events *EventsTransactorSession) EmitNb(nb *big.Int) (*types.Transaction, error) {
	return _Events.Contract.EmitNb(&_Events.TransactOpts, nb)
}

// EventsNumberEventIterator is returned from FilterNumberEvent and is used to iterate over the raw logs and unpacked data for NumberEvent events raised by the Events contract.
type EventsNumberEventIterator struct {
	Event *EventsNumberEvent // Event containing the contract specifics and raw log

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
func (it *EventsNumberEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EventsNumberEvent)
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
		it.Event = new(EventsNumberEvent)
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
func (it *EventsNumberEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EventsNumberEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EventsNumberEvent represents a NumberEvent event raised by the Events contract.
type EventsNumberEvent struct {
	Number *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNumberEvent is a free log retrieval operation binding the contract event 0xb7b401727864049638c23b3bf52321ea398b0fd65c72e465daf5c8b4564d7583.
//
// Solidity: event NumberEvent(int256 number)
func (_Events *EventsFilterer) FilterNumberEvent(opts *bind.FilterOpts) (*EventsNumberEventIterator, error) {

	logs, sub, err := _Events.contract.FilterLogs(opts, "NumberEvent")
	if err != nil {
		return nil, err
	}
	return &EventsNumberEventIterator{contract: _Events.contract, event: "NumberEvent", logs: logs, sub: sub}, nil
}

// WatchNumberEvent is a free log subscription operation binding the contract event 0xb7b401727864049638c23b3bf52321ea398b0fd65c72e465daf5c8b4564d7583.
//
// Solidity: event NumberEvent(int256 number)
func (_Events *EventsFilterer) WatchNumberEvent(opts *bind.WatchOpts, sink chan<- *EventsNumberEvent) (event.Subscription, error) {

	logs, sub, err := _Events.contract.WatchLogs(opts, "NumberEvent")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EventsNumberEvent)
				if err := _Events.contract.UnpackLog(event, "NumberEvent", log); err != nil {
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

// ParseNumberEvent is a log parse operation binding the contract event 0xb7b401727864049638c23b3bf52321ea398b0fd65c72e465daf5c8b4564d7583.
//
// Solidity: event NumberEvent(int256 number)
func (_Events *EventsFilterer) ParseNumberEvent(log types.Log) (*EventsNumberEvent, error) {
	event := new(EventsNumberEvent)
	if err := _Events.contract.UnpackLog(event, "NumberEvent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
