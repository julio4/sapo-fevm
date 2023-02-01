// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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

// JobsMetaData contains all meta data concerning the Jobs contract.
var JobsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sapoBridge\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"JobFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"JobSuceeded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"completed\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"failAndRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"executionResult\",\"type\":\"string\"}],\"name\":\"saveResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600160146101000a81548160ff021916908315150217905550604051610d30380380610d308339818101604052810190610040919061012a565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610157565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100f7826100cc565b9050919050565b610107816100ec565b811461011257600080fd5b50565b600081519050610124816100fe565b92915050565b6000602082840312156101405761013f6100c7565b5b600061014e84828501610115565b91505092915050565b610bca806101666000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063877db67d146100515780638c45c8651461006d5780639d9a7fe914610077578063de29278914610095575b600080fd5b61006b60048036038101906100669190610602565b6100b3565b005b6100756101b6565b005b61007f61035c565b60405161008c9190610666565b60405180910390f35b61009d61036f565b6040516100aa9190610700565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610143576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013a9061076e565b60405180910390fd5b600160149054906101000a900460ff161561015d57600080fd5b60018060146101000a81548160ff021916908315150217905550806002908161018691906109a4565b507f85c9037835c422f4d8688ffaff3f926bde639d28fa6a92c2960c123aa96191cb60405160405180910390a150565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610246576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023d9061076e565b60405180910390fd5b600160149054906101000a900460ff161561026057600080fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16476040516102a790610aa7565b60006040518083038185875af1925050503d80600081146102e4576040519150601f19603f3d011682016040523d82523d6000602084013e6102e9565b606091505b505090508061032d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032490610b08565b60405180910390fd5b7ff6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e60405160405180910390a150565b600160149054906101000a900460ff1681565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f690610b74565b60405180910390fd5b600160149054906101000a900460ff1661041857600080fd5b60028054610425906107bd565b80601f0160208091040260200160405190810160405280929190818152602001828054610451906107bd565b801561049e5780601f106104735761010080835404028352916020019161049e565b820191906000526020600020905b81548152906001019060200180831161048157829003601f168201915b5050505050905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61050f826104c6565b810181811067ffffffffffffffff8211171561052e5761052d6104d7565b5b80604052505050565b60006105416104a8565b905061054d8282610506565b919050565b600067ffffffffffffffff82111561056d5761056c6104d7565b5b610576826104c6565b9050602081019050919050565b82818337600083830152505050565b60006105a56105a084610552565b610537565b9050828152602081018484840111156105c1576105c06104c1565b5b6105cc848285610583565b509392505050565b600082601f8301126105e9576105e86104bc565b5b81356105f9848260208601610592565b91505092915050565b600060208284031215610618576106176104b2565b5b600082013567ffffffffffffffff811115610636576106356104b7565b5b610642848285016105d4565b91505092915050565b60008115159050919050565b6106608161064b565b82525050565b600060208201905061067b6000830184610657565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156106bb5780820151818401526020810190506106a0565b60008484015250505050565b60006106d282610681565b6106dc818561068c565b93506106ec81856020860161069d565b6106f5816104c6565b840191505092915050565b6000602082019050818103600083015261071a81846106c7565b905092915050565b7f43616c6c6572206973206e6f7420627269646765000000000000000000000000600082015250565b600061075860148361068c565b915061076382610722565b602082019050919050565b600060208201905081810360008301526107878161074b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806107d557607f821691505b6020821081036107e8576107e761078e565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026108507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610813565b61085a8683610813565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006108a161089c61089784610872565b61087c565b610872565b9050919050565b6000819050919050565b6108bb83610886565b6108cf6108c7826108a8565b848454610820565b825550505050565b600090565b6108e46108d7565b6108ef8184846108b2565b505050565b5b81811015610913576109086000826108dc565b6001810190506108f5565b5050565b601f82111561095857610929816107ee565b61093284610803565b81016020851015610941578190505b61095561094d85610803565b8301826108f4565b50505b505050565b600082821c905092915050565b600061097b6000198460080261095d565b1980831691505092915050565b6000610994838361096a565b9150826002028217905092915050565b6109ad82610681565b67ffffffffffffffff8111156109c6576109c56104d7565b5b6109d082546107bd565b6109db828285610917565b600060209050601f831160018114610a0e57600084156109fc578287015190505b610a068582610988565b865550610a6e565b601f198416610a1c866107ee565b60005b82811015610a4457848901518255600182019150602085019450602081019050610a1f565b86831015610a615784890151610a5d601f89168261096a565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b50565b6000610a91600083610a76565b9150610a9c82610a81565b600082019050919050565b6000610ab282610a84565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610af260148361068c565b9150610afd82610abc565b602082019050919050565b60006020820190508181036000830152610b2181610ae5565b9050919050565b7f43616c6c6572206973206e6f7420696e69746961746f72000000000000000000600082015250565b6000610b5e60178361068c565b9150610b6982610b28565b602082019050919050565b60006020820190508181036000830152610b8d81610b51565b905091905056fea264697066735822122009b3210d60d5b586d43673e3466d53cab2595952ffc385c9db75e43a8c509d7e64736f6c63430008110033",
}

// JobsABI is the input ABI used to generate the binding from.
// Deprecated: Use JobsMetaData.ABI instead.
var JobsABI = JobsMetaData.ABI

// JobsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JobsMetaData.Bin instead.
var JobsBin = JobsMetaData.Bin

// DeployJobs deploys a new Ethereum contract, binding an instance of Jobs to it.
func DeployJobs(auth *bind.TransactOpts, backend bind.ContractBackend, sapoBridge common.Address) (common.Address, *types.Transaction, *Jobs, error) {
	parsed, err := JobsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JobsBin), backend, sapoBridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Jobs{JobsCaller: JobsCaller{contract: contract}, JobsTransactor: JobsTransactor{contract: contract}, JobsFilterer: JobsFilterer{contract: contract}}, nil
}

// Jobs is an auto generated Go binding around an Ethereum contract.
type Jobs struct {
	JobsCaller     // Read-only binding to the contract
	JobsTransactor // Write-only binding to the contract
	JobsFilterer   // Log filterer for contract events
}

// JobsCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobsSession struct {
	Contract     *Jobs             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobsCallerSession struct {
	Contract *JobsCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// JobsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobsTransactorSession struct {
	Contract     *JobsTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobsRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobsRaw struct {
	Contract *Jobs // Generic contract binding to access the raw methods on
}

// JobsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobsCallerRaw struct {
	Contract *JobsCaller // Generic read-only contract binding to access the raw methods on
}

// JobsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobsTransactorRaw struct {
	Contract *JobsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJobs creates a new instance of Jobs, bound to a specific deployed contract.
func NewJobs(address common.Address, backend bind.ContractBackend) (*Jobs, error) {
	contract, err := bindJobs(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Jobs{JobsCaller: JobsCaller{contract: contract}, JobsTransactor: JobsTransactor{contract: contract}, JobsFilterer: JobsFilterer{contract: contract}}, nil
}

// NewJobsCaller creates a new read-only instance of Jobs, bound to a specific deployed contract.
func NewJobsCaller(address common.Address, caller bind.ContractCaller) (*JobsCaller, error) {
	contract, err := bindJobs(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobsCaller{contract: contract}, nil
}

// NewJobsTransactor creates a new write-only instance of Jobs, bound to a specific deployed contract.
func NewJobsTransactor(address common.Address, transactor bind.ContractTransactor) (*JobsTransactor, error) {
	contract, err := bindJobs(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobsTransactor{contract: contract}, nil
}

// NewJobsFilterer creates a new log filterer instance of Jobs, bound to a specific deployed contract.
func NewJobsFilterer(address common.Address, filterer bind.ContractFilterer) (*JobsFilterer, error) {
	contract, err := bindJobs(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobsFilterer{contract: contract}, nil
}

// bindJobs binds a generic wrapper to an already deployed contract.
func bindJobs(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JobsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobs *JobsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jobs.Contract.JobsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobs *JobsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobs.Contract.JobsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobs *JobsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobs.Contract.JobsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Jobs *JobsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Jobs.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Jobs *JobsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobs.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Jobs *JobsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Jobs.Contract.contract.Transact(opts, method, params...)
}

// Completed is a free data retrieval call binding the contract method 0x9d9a7fe9.
//
// Solidity: function completed() view returns(bool)
func (_Jobs *JobsCaller) Completed(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Jobs.contract.Call(opts, &out, "completed")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Completed is a free data retrieval call binding the contract method 0x9d9a7fe9.
//
// Solidity: function completed() view returns(bool)
func (_Jobs *JobsSession) Completed() (bool, error) {
	return _Jobs.Contract.Completed(&_Jobs.CallOpts)
}

// Completed is a free data retrieval call binding the contract method 0x9d9a7fe9.
//
// Solidity: function completed() view returns(bool)
func (_Jobs *JobsCallerSession) Completed() (bool, error) {
	return _Jobs.Contract.Completed(&_Jobs.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(string)
func (_Jobs *JobsCaller) GetResult(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Jobs.contract.Call(opts, &out, "getResult")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(string)
func (_Jobs *JobsSession) GetResult() (string, error) {
	return _Jobs.Contract.GetResult(&_Jobs.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(string)
func (_Jobs *JobsCallerSession) GetResult() (string, error) {
	return _Jobs.Contract.GetResult(&_Jobs.CallOpts)
}

// FailAndRefund is a paid mutator transaction binding the contract method 0x8c45c865.
//
// Solidity: function failAndRefund() returns()
func (_Jobs *JobsTransactor) FailAndRefund(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Jobs.contract.Transact(opts, "failAndRefund")
}

// FailAndRefund is a paid mutator transaction binding the contract method 0x8c45c865.
//
// Solidity: function failAndRefund() returns()
func (_Jobs *JobsSession) FailAndRefund() (*types.Transaction, error) {
	return _Jobs.Contract.FailAndRefund(&_Jobs.TransactOpts)
}

// FailAndRefund is a paid mutator transaction binding the contract method 0x8c45c865.
//
// Solidity: function failAndRefund() returns()
func (_Jobs *JobsTransactorSession) FailAndRefund() (*types.Transaction, error) {
	return _Jobs.Contract.FailAndRefund(&_Jobs.TransactOpts)
}

// SaveResult is a paid mutator transaction binding the contract method 0x877db67d.
//
// Solidity: function saveResult(string executionResult) returns()
func (_Jobs *JobsTransactor) SaveResult(opts *bind.TransactOpts, executionResult string) (*types.Transaction, error) {
	return _Jobs.contract.Transact(opts, "saveResult", executionResult)
}

// SaveResult is a paid mutator transaction binding the contract method 0x877db67d.
//
// Solidity: function saveResult(string executionResult) returns()
func (_Jobs *JobsSession) SaveResult(executionResult string) (*types.Transaction, error) {
	return _Jobs.Contract.SaveResult(&_Jobs.TransactOpts, executionResult)
}

// SaveResult is a paid mutator transaction binding the contract method 0x877db67d.
//
// Solidity: function saveResult(string executionResult) returns()
func (_Jobs *JobsTransactorSession) SaveResult(executionResult string) (*types.Transaction, error) {
	return _Jobs.Contract.SaveResult(&_Jobs.TransactOpts, executionResult)
}

// JobsJobFailedIterator is returned from FilterJobFailed and is used to iterate over the raw logs and unpacked data for JobFailed events raised by the Jobs contract.
type JobsJobFailedIterator struct {
	Event *JobsJobFailed // Event containing the contract specifics and raw log

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
func (it *JobsJobFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobsJobFailed)
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
		it.Event = new(JobsJobFailed)
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
func (it *JobsJobFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobsJobFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobsJobFailed represents a JobFailed event raised by the Jobs contract.
type JobsJobFailed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterJobFailed is a free log retrieval operation binding the contract event 0xf6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e.
//
// Solidity: event JobFailed()
func (_Jobs *JobsFilterer) FilterJobFailed(opts *bind.FilterOpts) (*JobsJobFailedIterator, error) {

	logs, sub, err := _Jobs.contract.FilterLogs(opts, "JobFailed")
	if err != nil {
		return nil, err
	}
	return &JobsJobFailedIterator{contract: _Jobs.contract, event: "JobFailed", logs: logs, sub: sub}, nil
}

// WatchJobFailed is a free log subscription operation binding the contract event 0xf6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e.
//
// Solidity: event JobFailed()
func (_Jobs *JobsFilterer) WatchJobFailed(opts *bind.WatchOpts, sink chan<- *JobsJobFailed) (event.Subscription, error) {

	logs, sub, err := _Jobs.contract.WatchLogs(opts, "JobFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobsJobFailed)
				if err := _Jobs.contract.UnpackLog(event, "JobFailed", log); err != nil {
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

// ParseJobFailed is a log parse operation binding the contract event 0xf6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e.
//
// Solidity: event JobFailed()
func (_Jobs *JobsFilterer) ParseJobFailed(log types.Log) (*JobsJobFailed, error) {
	event := new(JobsJobFailed)
	if err := _Jobs.contract.UnpackLog(event, "JobFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobsJobSuceededIterator is returned from FilterJobSuceeded and is used to iterate over the raw logs and unpacked data for JobSuceeded events raised by the Jobs contract.
type JobsJobSuceededIterator struct {
	Event *JobsJobSuceeded // Event containing the contract specifics and raw log

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
func (it *JobsJobSuceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobsJobSuceeded)
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
		it.Event = new(JobsJobSuceeded)
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
func (it *JobsJobSuceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobsJobSuceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobsJobSuceeded represents a JobSuceeded event raised by the Jobs contract.
type JobsJobSuceeded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterJobSuceeded is a free log retrieval operation binding the contract event 0x85c9037835c422f4d8688ffaff3f926bde639d28fa6a92c2960c123aa96191cb.
//
// Solidity: event JobSuceeded()
func (_Jobs *JobsFilterer) FilterJobSuceeded(opts *bind.FilterOpts) (*JobsJobSuceededIterator, error) {

	logs, sub, err := _Jobs.contract.FilterLogs(opts, "JobSuceeded")
	if err != nil {
		return nil, err
	}
	return &JobsJobSuceededIterator{contract: _Jobs.contract, event: "JobSuceeded", logs: logs, sub: sub}, nil
}

// WatchJobSuceeded is a free log subscription operation binding the contract event 0x85c9037835c422f4d8688ffaff3f926bde639d28fa6a92c2960c123aa96191cb.
//
// Solidity: event JobSuceeded()
func (_Jobs *JobsFilterer) WatchJobSuceeded(opts *bind.WatchOpts, sink chan<- *JobsJobSuceeded) (event.Subscription, error) {

	logs, sub, err := _Jobs.contract.WatchLogs(opts, "JobSuceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobsJobSuceeded)
				if err := _Jobs.contract.UnpackLog(event, "JobSuceeded", log); err != nil {
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

// ParseJobSuceeded is a log parse operation binding the contract event 0x85c9037835c422f4d8688ffaff3f926bde639d28fa6a92c2960c123aa96191cb.
//
// Solidity: event JobSuceeded()
func (_Jobs *JobsFilterer) ParseJobSuceeded(log types.Log) (*JobsJobSuceeded, error) {
	event := new(JobsJobSuceeded)
	if err := _Jobs.contract.UnpackLog(event, "JobSuceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
