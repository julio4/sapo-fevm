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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestInitiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sapoBridge\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"JobFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"JobSucceeded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"failAndRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitiator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumSapoJob.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"exResult1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"exResult2\",\"type\":\"bytes32\"}],\"name\":\"saveResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600260146101000a81548160ff021916908360028111156200002d576200002c62000121565b5b021790555060405162001292380380620012928339818101604052810190620000579190620001ba565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000201565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001828262000155565b9050919050565b620001948162000175565b8114620001a057600080fd5b50565b600081519050620001b48162000189565b92915050565b60008060408385031215620001d457620001d362000150565b5b6000620001e485828601620001a3565b9250506020620001f785828601620001a3565b9150509250929050565b61108180620002116000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80634e69d56014610067578063836fbaca1461008557806384de1ea5146100a15780638c45c865146100bf578063de292789146100c9578063fb32c508146100e7575b600080fd5b61006f610105565b60405161007c9190610991565b60405180910390f35b61009f600480360381019061009a91906109f6565b61011c565b005b6100a96103b2565b6040516100b69190610a77565b60405180910390f35b6100c76103db565b005b6100d1610732565b6040516100de9190610b22565b60405180910390f35b6100ef6108f0565b6040516100fc9190610a77565b60405180910390f35b6000600260149054906101000a900460ff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806101c55750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610204576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fb90610b90565b60405180910390fd5b600060028111156102185761021761091a565b5b600260149054906101000a900460ff16600281111561023a5761023961091a565b5b1461027a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027190610bfc565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16476040516102c290610c4d565b60006040518083038185875af1925050503d80600081146102ff576040519150601f19603f3d011682016040523d82523d6000602084013e610304565b606091505b5050905080610348576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033f90610cae565b60405180910390fd5b6001600260146101000a81548160ff0219169083600281111561036e5761036d61091a565b5b021790555082600381905550816004819055507f7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b60405160405180910390a1505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806104845750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b6104c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ba90610b90565b60405180910390fd5b600060028111156104d7576104d661091a565b5b600260149054906101000a900460ff1660028111156104f9576104f861091a565b5b14610539576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053090610bfc565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1666b1a2bc2ec5000060405161058790610c4d565b60006040518083038185875af1925050503d80600081146105c4576040519150601f19603f3d011682016040523d82523d6000602084013e6105c9565b606091505b505090508061060d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060490610d1a565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff164760405161065190610c4d565b60006040518083038185875af1925050503d806000811461068e576040519150601f19603f3d011682016040523d82523d6000602084013e610693565b606091505b505080915050806106d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d090610dac565b60405180910390fd5b60028060146101000a81548160ff021916908360028111156106fe576106fd61091a565b5b02179055507ff6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e60405160405180910390a150565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b990610e18565b60405180910390fd5b600060028111156107d6576107d561091a565b5b600260149054906101000a900460ff1660028111156107f8576107f761091a565b5b03610838576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082f90610e84565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff16631453d7566003546004546040518363ffffffff1660e01b81526004016108a0929190610eb3565b600060405180830381865afa1580156108bd573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906108e69190611002565b9050809250505090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6003811061095a5761095961091a565b5b50565b600081905061096b82610949565b919050565b600061097b8261095d565b9050919050565b61098b81610970565b82525050565b60006020820190506109a66000830184610982565b92915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6109d3816109c0565b81146109de57600080fd5b50565b6000813590506109f0816109ca565b92915050565b60008060408385031215610a0d57610a0c6109b6565b5b6000610a1b858286016109e1565b9250506020610a2c858286016109e1565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a6182610a36565b9050919050565b610a7181610a56565b82525050565b6000602082019050610a8c6000830184610a68565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610acc578082015181840152602081019050610ab1565b60008484015250505050565b6000601f19601f8301169050919050565b6000610af482610a92565b610afe8185610a9d565b9350610b0e818560208601610aae565b610b1781610ad8565b840191505092915050565b60006020820190508181036000830152610b3c8184610ae9565b905092915050565b7f43616c6c6572206973206e6f7420627269646765000000000000000000000000600082015250565b6000610b7a601483610a9d565b9150610b8582610b44565b602082019050919050565b60006020820190508181036000830152610ba981610b6d565b9050919050565b7f4a6f62206973206e6f742070656e64696e670000000000000000000000000000600082015250565b6000610be6601283610a9d565b9150610bf182610bb0565b602082019050919050565b60006020820190508181036000830152610c1581610bd9565b9050919050565b600081905092915050565b50565b6000610c37600083610c1c565b9150610c4282610c27565b600082019050919050565b6000610c5882610c2a565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610c98601483610a9d565b9150610ca382610c62565b602082019050919050565b60006020820190508181036000830152610cc781610c8b565b9050919050565b7f4661696c656420746f2073656e6420457468657220746f206272696467650000600082015250565b6000610d04601e83610a9d565b9150610d0f82610cce565b602082019050919050565b60006020820190508181036000830152610d3381610cf7565b9050919050565b7f4661696c656420746f2073656e6420457468657220746f20696e69746961746f60008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b6000610d96602183610a9d565b9150610da182610d3a565b604082019050919050565b60006020820190508181036000830152610dc581610d89565b9050919050565b7f43616c6c6572206973206e6f7420696e69746961746f72000000000000000000600082015250565b6000610e02601783610a9d565b9150610e0d82610dcc565b602082019050919050565b60006020820190508181036000830152610e3181610df5565b9050919050565b7f4a6f62206973206e6f7420636f6d706c65746564000000000000000000000000600082015250565b6000610e6e601483610a9d565b9150610e7982610e38565b602082019050919050565b60006020820190508181036000830152610e9d81610e61565b9050919050565b610ead816109c0565b82525050565b6000604082019050610ec86000830185610ea4565b610ed56020830184610ea4565b9392505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f1e82610ad8565b810181811067ffffffffffffffff82111715610f3d57610f3c610ee6565b5b80604052505050565b6000610f506109ac565b9050610f5c8282610f15565b919050565b600067ffffffffffffffff821115610f7c57610f7b610ee6565b5b610f8582610ad8565b9050602081019050919050565b6000610fa5610fa084610f61565b610f46565b905082815260208101848484011115610fc157610fc0610ee1565b5b610fcc848285610aae565b509392505050565b600082601f830112610fe957610fe8610edc565b5b8151610ff9848260208601610f92565b91505092915050565b600060208284031215611018576110176109b6565b5b600082015167ffffffffffffffff811115611036576110356109bb565b5b61104284828501610fd4565b9150509291505056fea26469706673582212208c68a6b9766ac716d1d319f5d0f6fd3e30628b0b9ea0079a24e6aea5521fb0e264736f6c63430008110033",
}

// JobsABI is the input ABI used to generate the binding from.
// Deprecated: Use JobsMetaData.ABI instead.
var JobsABI = JobsMetaData.ABI

// JobsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JobsMetaData.Bin instead.
var JobsBin = JobsMetaData.Bin

// DeployJobs deploys a new Ethereum contract, binding an instance of Jobs to it.
func DeployJobs(auth *bind.TransactOpts, backend bind.ContractBackend, requestInitiator common.Address, sapoBridge common.Address) (common.Address, *types.Transaction, *Jobs, error) {
	parsed, err := JobsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JobsBin), backend, requestInitiator, sapoBridge)
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

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address)
func (_Jobs *JobsCaller) GetBridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobs.contract.Call(opts, &out, "getBridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address)
func (_Jobs *JobsSession) GetBridgeAddress() (common.Address, error) {
	return _Jobs.Contract.GetBridgeAddress(&_Jobs.CallOpts)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address)
func (_Jobs *JobsCallerSession) GetBridgeAddress() (common.Address, error) {
	return _Jobs.Contract.GetBridgeAddress(&_Jobs.CallOpts)
}

// GetInitiator is a free data retrieval call binding the contract method 0x84de1ea5.
//
// Solidity: function getInitiator() view returns(address)
func (_Jobs *JobsCaller) GetInitiator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Jobs.contract.Call(opts, &out, "getInitiator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInitiator is a free data retrieval call binding the contract method 0x84de1ea5.
//
// Solidity: function getInitiator() view returns(address)
func (_Jobs *JobsSession) GetInitiator() (common.Address, error) {
	return _Jobs.Contract.GetInitiator(&_Jobs.CallOpts)
}

// GetInitiator is a free data retrieval call binding the contract method 0x84de1ea5.
//
// Solidity: function getInitiator() view returns(address)
func (_Jobs *JobsCallerSession) GetInitiator() (common.Address, error) {
	return _Jobs.Contract.GetInitiator(&_Jobs.CallOpts)
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

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_Jobs *JobsCaller) GetStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Jobs.contract.Call(opts, &out, "getStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_Jobs *JobsSession) GetStatus() (uint8, error) {
	return _Jobs.Contract.GetStatus(&_Jobs.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_Jobs *JobsCallerSession) GetStatus() (uint8, error) {
	return _Jobs.Contract.GetStatus(&_Jobs.CallOpts)
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

// SaveResult is a paid mutator transaction binding the contract method 0x836fbaca.
//
// Solidity: function saveResult(bytes32 exResult1, bytes32 exResult2) returns()
func (_Jobs *JobsTransactor) SaveResult(opts *bind.TransactOpts, exResult1 [32]byte, exResult2 [32]byte) (*types.Transaction, error) {
	return _Jobs.contract.Transact(opts, "saveResult", exResult1, exResult2)
}

// SaveResult is a paid mutator transaction binding the contract method 0x836fbaca.
//
// Solidity: function saveResult(bytes32 exResult1, bytes32 exResult2) returns()
func (_Jobs *JobsSession) SaveResult(exResult1 [32]byte, exResult2 [32]byte) (*types.Transaction, error) {
	return _Jobs.Contract.SaveResult(&_Jobs.TransactOpts, exResult1, exResult2)
}

// SaveResult is a paid mutator transaction binding the contract method 0x836fbaca.
//
// Solidity: function saveResult(bytes32 exResult1, bytes32 exResult2) returns()
func (_Jobs *JobsTransactorSession) SaveResult(exResult1 [32]byte, exResult2 [32]byte) (*types.Transaction, error) {
	return _Jobs.Contract.SaveResult(&_Jobs.TransactOpts, exResult1, exResult2)
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

// JobsJobSucceededIterator is returned from FilterJobSucceeded and is used to iterate over the raw logs and unpacked data for JobSucceeded events raised by the Jobs contract.
type JobsJobSucceededIterator struct {
	Event *JobsJobSucceeded // Event containing the contract specifics and raw log

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
func (it *JobsJobSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobsJobSucceeded)
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
		it.Event = new(JobsJobSucceeded)
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
func (it *JobsJobSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobsJobSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobsJobSucceeded represents a JobSucceeded event raised by the Jobs contract.
type JobsJobSucceeded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterJobSucceeded is a free log retrieval operation binding the contract event 0x7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b.
//
// Solidity: event JobSucceeded()
func (_Jobs *JobsFilterer) FilterJobSucceeded(opts *bind.FilterOpts) (*JobsJobSucceededIterator, error) {

	logs, sub, err := _Jobs.contract.FilterLogs(opts, "JobSucceeded")
	if err != nil {
		return nil, err
	}
	return &JobsJobSucceededIterator{contract: _Jobs.contract, event: "JobSucceeded", logs: logs, sub: sub}, nil
}

// WatchJobSucceeded is a free log subscription operation binding the contract event 0x7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b.
//
// Solidity: event JobSucceeded()
func (_Jobs *JobsFilterer) WatchJobSucceeded(opts *bind.WatchOpts, sink chan<- *JobsJobSucceeded) (event.Subscription, error) {

	logs, sub, err := _Jobs.contract.WatchLogs(opts, "JobSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobsJobSucceeded)
				if err := _Jobs.contract.UnpackLog(event, "JobSucceeded", log); err != nil {
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

// ParseJobSucceeded is a log parse operation binding the contract event 0x7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b.
//
// Solidity: event JobSucceeded()
func (_Jobs *JobsFilterer) ParseJobSucceeded(log types.Log) (*JobsJobSucceeded, error) {
	event := new(JobsJobSucceeded)
	if err := _Jobs.contract.UnpackLog(event, "JobSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
