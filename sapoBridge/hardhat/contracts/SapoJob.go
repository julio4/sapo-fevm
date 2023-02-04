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

// JobMetaData contains all meta data concerning the Job contract.
var JobMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"requestInitiator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"sapoBridge\",\"type\":\"address\"}],\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"JobFailed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"JobSucceeded\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"failAndRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInitiator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResult\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStatus\",\"outputs\":[{\"internalType\":\"enumSapoJob.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"executionResult\",\"type\":\"string\"}],\"name\":\"saveResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526000600160146101000a81548160ff021916908360028111156200002d576200002c620000e0565b5b02179055506040516200114638038062001146833981810160405281019062000057919062000179565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620001c0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001418262000114565b9050919050565b620001538162000134565b81146200015f57600080fd5b50565b600081519050620001738162000148565b92915050565b600080604083850312156200019357620001926200010f565b5b6000620001a38582860162000162565b9250506020620001b68582860162000162565b9150509250929050565b610f7680620001d06000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80634e69d5601461006757806384de1ea514610085578063877db67d146100a3578063b422a51e146100bf578063de292789146100db578063fb32c508146100f9575b600080fd5b61006f610117565b60405161007c919061073b565b60405180910390f35b61008d61012e565b60405161009a9190610797565b60405180910390f35b6100bd60048036038101906100b8919061090c565b610157565b005b6100d960048036038101906100d4919061090c565b6102c6565b005b6100e3610504565b6040516100f091906109d4565b60405180910390f35b61010161069a565b60405161010e9190610797565b60405180910390f35b6000600160149054906101000a900460ff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101de90610a42565b60405180910390fd5b600060028111156101fb576101fa6106c4565b5b600160149054906101000a900460ff16600281111561021d5761021c6106c4565b5b1461025d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025490610aae565b60405180910390fd5b60018060146101000a81548160ff02191690836002811115610282576102816106c4565b5b021790555080600290816102969190610ce4565b507f7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b60405160405180910390a150565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610356576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034d90610a42565b60405180910390fd5b6000600281111561036a576103696106c4565b5b600160149054906101000a900460ff16600281111561038c5761038b6106c4565b5b146103cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c390610aae565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff164760405161041390610de7565b60006040518083038185875af1925050503d8060008114610450576040519150601f19603f3d011682016040523d82523d6000602084013e610455565b606091505b5050905080610499576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049090610e48565b60405180910390fd5b6002600160146101000a81548160ff021916908360028111156104bf576104be6106c4565b5b021790555081600290816104d39190610ce4565b507ff6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e60405160405180910390a15050565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610594576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058b90610eb4565b60405180910390fd5b600060028111156105a8576105a76106c4565b5b600160149054906101000a900460ff1660028111156105ca576105c96106c4565b5b0361060a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060190610f20565b60405180910390fd5b6002805461061790610afd565b80601f016020809104026020016040519081016040528092919081815260200182805461064390610afd565b80156106905780601f1061066557610100808354040283529160200191610690565b820191906000526020600020905b81548152906001019060200180831161067357829003601f168201915b5050505050905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110610704576107036106c4565b5b50565b6000819050610715826106f3565b919050565b600061072582610707565b9050919050565b6107358161071a565b82525050565b6000602082019050610750600083018461072c565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061078182610756565b9050919050565b61079181610776565b82525050565b60006020820190506107ac6000830184610788565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610819826107d0565b810181811067ffffffffffffffff82111715610838576108376107e1565b5b80604052505050565b600061084b6107b2565b90506108578282610810565b919050565b600067ffffffffffffffff821115610877576108766107e1565b5b610880826107d0565b9050602081019050919050565b82818337600083830152505050565b60006108af6108aa8461085c565b610841565b9050828152602081018484840111156108cb576108ca6107cb565b5b6108d684828561088d565b509392505050565b600082601f8301126108f3576108f26107c6565b5b813561090384826020860161089c565b91505092915050565b600060208284031215610922576109216107bc565b5b600082013567ffffffffffffffff8111156109405761093f6107c1565b5b61094c848285016108de565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561098f578082015181840152602081019050610974565b60008484015250505050565b60006109a682610955565b6109b08185610960565b93506109c0818560208601610971565b6109c9816107d0565b840191505092915050565b600060208201905081810360008301526109ee818461099b565b905092915050565b7f43616c6c6572206973206e6f7420627269646765000000000000000000000000600082015250565b6000610a2c601483610960565b9150610a37826109f6565b602082019050919050565b60006020820190508181036000830152610a5b81610a1f565b9050919050565b7f4a6f62206973206e6f742070656e64696e670000000000000000000000000000600082015250565b6000610a98601283610960565b9150610aa382610a62565b602082019050919050565b60006020820190508181036000830152610ac781610a8b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610b1557607f821691505b602082108103610b2857610b27610ace565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610b53565b610b9a8683610b53565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610be1610bdc610bd784610bb2565b610bbc565b610bb2565b9050919050565b6000819050919050565b610bfb83610bc6565b610c0f610c0782610be8565b848454610b60565b825550505050565b600090565b610c24610c17565b610c2f818484610bf2565b505050565b5b81811015610c5357610c48600082610c1c565b600181019050610c35565b5050565b601f821115610c9857610c6981610b2e565b610c7284610b43565b81016020851015610c81578190505b610c95610c8d85610b43565b830182610c34565b50505b505050565b600082821c905092915050565b6000610cbb60001984600802610c9d565b1980831691505092915050565b6000610cd48383610caa565b9150826002028217905092915050565b610ced82610955565b67ffffffffffffffff811115610d0657610d056107e1565b5b610d108254610afd565b610d1b828285610c57565b600060209050601f831160018114610d4e5760008415610d3c578287015190505b610d468582610cc8565b865550610dae565b601f198416610d5c86610b2e565b60005b82811015610d8457848901518255600182019150602085019450602081019050610d5f565b86831015610da15784890151610d9d601f891682610caa565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b50565b6000610dd1600083610db6565b9150610ddc82610dc1565b600082019050919050565b6000610df282610dc4565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610e32601483610960565b9150610e3d82610dfc565b602082019050919050565b60006020820190508181036000830152610e6181610e25565b9050919050565b7f43616c6c6572206973206e6f7420696e69746961746f72000000000000000000600082015250565b6000610e9e601783610960565b9150610ea982610e68565b602082019050919050565b60006020820190508181036000830152610ecd81610e91565b9050919050565b7f4a6f62206973206e6f7420636f6d706c65746564000000000000000000000000600082015250565b6000610f0a601483610960565b9150610f1582610ed4565b602082019050919050565b60006020820190508181036000830152610f3981610efd565b905091905056fea2646970667358221220c6e1d6395916b7526155d5a58e6dab059a8428f8b12edca0e99dddb705f8088964736f6c63430008110033",
}

// JobABI is the input ABI used to generate the binding from.
// Deprecated: Use JobMetaData.ABI instead.
var JobABI = JobMetaData.ABI

// JobBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use JobMetaData.Bin instead.
var JobBin = JobMetaData.Bin

// DeployJob deploys a new Ethereum contract, binding an instance of Job to it.
func DeployJob(auth *bind.TransactOpts, backend bind.ContractBackend, requestInitiator common.Address, sapoBridge common.Address) (common.Address, *types.Transaction, *Job, error) {
	parsed, err := JobMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(JobBin), backend, requestInitiator, sapoBridge)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Job{JobCaller: JobCaller{contract: contract}, JobTransactor: JobTransactor{contract: contract}, JobFilterer: JobFilterer{contract: contract}}, nil
}

// Job is an auto generated Go binding around an Ethereum contract.
type Job struct {
	JobCaller     // Read-only binding to the contract
	JobTransactor // Write-only binding to the contract
	JobFilterer   // Log filterer for contract events
}

// JobCaller is an auto generated read-only Go binding around an Ethereum contract.
type JobCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobTransactor is an auto generated write-only Go binding around an Ethereum contract.
type JobTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type JobFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// JobSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type JobSession struct {
	Contract     *Job              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type JobCallerSession struct {
	Contract *JobCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// JobTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type JobTransactorSession struct {
	Contract     *JobTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// JobRaw is an auto generated low-level Go binding around an Ethereum contract.
type JobRaw struct {
	Contract *Job // Generic contract binding to access the raw methods on
}

// JobCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type JobCallerRaw struct {
	Contract *JobCaller // Generic read-only contract binding to access the raw methods on
}

// JobTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type JobTransactorRaw struct {
	Contract *JobTransactor // Generic write-only contract binding to access the raw methods on
}

// NewJob creates a new instance of Job, bound to a specific deployed contract.
func NewJob(address common.Address, backend bind.ContractBackend) (*Job, error) {
	contract, err := bindJob(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Job{JobCaller: JobCaller{contract: contract}, JobTransactor: JobTransactor{contract: contract}, JobFilterer: JobFilterer{contract: contract}}, nil
}

// NewJobCaller creates a new read-only instance of Job, bound to a specific deployed contract.
func NewJobCaller(address common.Address, caller bind.ContractCaller) (*JobCaller, error) {
	contract, err := bindJob(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &JobCaller{contract: contract}, nil
}

// NewJobTransactor creates a new write-only instance of Job, bound to a specific deployed contract.
func NewJobTransactor(address common.Address, transactor bind.ContractTransactor) (*JobTransactor, error) {
	contract, err := bindJob(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &JobTransactor{contract: contract}, nil
}

// NewJobFilterer creates a new log filterer instance of Job, bound to a specific deployed contract.
func NewJobFilterer(address common.Address, filterer bind.ContractFilterer) (*JobFilterer, error) {
	contract, err := bindJob(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &JobFilterer{contract: contract}, nil
}

// bindJob binds a generic wrapper to an already deployed contract.
func bindJob(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(JobABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Job *JobRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Job.Contract.JobCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Job *JobRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Job.Contract.JobTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Job *JobRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Job.Contract.JobTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Job *JobCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Job.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Job *JobTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Job.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Job *JobTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Job.Contract.contract.Transact(opts, method, params...)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address)
func (_Job *JobCaller) GetBridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Job.contract.Call(opts, &out, "getBridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address)
func (_Job *JobSession) GetBridgeAddress() (common.Address, error) {
	return _Job.Contract.GetBridgeAddress(&_Job.CallOpts)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address)
func (_Job *JobCallerSession) GetBridgeAddress() (common.Address, error) {
	return _Job.Contract.GetBridgeAddress(&_Job.CallOpts)
}

// GetInitiator is a free data retrieval call binding the contract method 0x84de1ea5.
//
// Solidity: function getInitiator() view returns(address)
func (_Job *JobCaller) GetInitiator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Job.contract.Call(opts, &out, "getInitiator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInitiator is a free data retrieval call binding the contract method 0x84de1ea5.
//
// Solidity: function getInitiator() view returns(address)
func (_Job *JobSession) GetInitiator() (common.Address, error) {
	return _Job.Contract.GetInitiator(&_Job.CallOpts)
}

// GetInitiator is a free data retrieval call binding the contract method 0x84de1ea5.
//
// Solidity: function getInitiator() view returns(address)
func (_Job *JobCallerSession) GetInitiator() (common.Address, error) {
	return _Job.Contract.GetInitiator(&_Job.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(string)
func (_Job *JobCaller) GetResult(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Job.contract.Call(opts, &out, "getResult")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(string)
func (_Job *JobSession) GetResult() (string, error) {
	return _Job.Contract.GetResult(&_Job.CallOpts)
}

// GetResult is a free data retrieval call binding the contract method 0xde292789.
//
// Solidity: function getResult() view returns(string)
func (_Job *JobCallerSession) GetResult() (string, error) {
	return _Job.Contract.GetResult(&_Job.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_Job *JobCaller) GetStatus(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Job.contract.Call(opts, &out, "getStatus")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_Job *JobSession) GetStatus() (uint8, error) {
	return _Job.Contract.GetStatus(&_Job.CallOpts)
}

// GetStatus is a free data retrieval call binding the contract method 0x4e69d560.
//
// Solidity: function getStatus() view returns(uint8)
func (_Job *JobCallerSession) GetStatus() (uint8, error) {
	return _Job.Contract.GetStatus(&_Job.CallOpts)
}

// FailAndRefund is a paid mutator transaction binding the contract method 0xb422a51e.
//
// Solidity: function failAndRefund(string reason) returns()
func (_Job *JobTransactor) FailAndRefund(opts *bind.TransactOpts, reason string) (*types.Transaction, error) {
	return _Job.contract.Transact(opts, "failAndRefund", reason)
}

// FailAndRefund is a paid mutator transaction binding the contract method 0xb422a51e.
//
// Solidity: function failAndRefund(string reason) returns()
func (_Job *JobSession) FailAndRefund(reason string) (*types.Transaction, error) {
	return _Job.Contract.FailAndRefund(&_Job.TransactOpts, reason)
}

// FailAndRefund is a paid mutator transaction binding the contract method 0xb422a51e.
//
// Solidity: function failAndRefund(string reason) returns()
func (_Job *JobTransactorSession) FailAndRefund(reason string) (*types.Transaction, error) {
	return _Job.Contract.FailAndRefund(&_Job.TransactOpts, reason)
}

// SaveResult is a paid mutator transaction binding the contract method 0x877db67d.
//
// Solidity: function saveResult(string executionResult) returns()
func (_Job *JobTransactor) SaveResult(opts *bind.TransactOpts, executionResult string) (*types.Transaction, error) {
	return _Job.contract.Transact(opts, "saveResult", executionResult)
}

// SaveResult is a paid mutator transaction binding the contract method 0x877db67d.
//
// Solidity: function saveResult(string executionResult) returns()
func (_Job *JobSession) SaveResult(executionResult string) (*types.Transaction, error) {
	return _Job.Contract.SaveResult(&_Job.TransactOpts, executionResult)
}

// SaveResult is a paid mutator transaction binding the contract method 0x877db67d.
//
// Solidity: function saveResult(string executionResult) returns()
func (_Job *JobTransactorSession) SaveResult(executionResult string) (*types.Transaction, error) {
	return _Job.Contract.SaveResult(&_Job.TransactOpts, executionResult)
}

// JobJobFailedIterator is returned from FilterJobFailed and is used to iterate over the raw logs and unpacked data for JobFailed events raised by the Job contract.
type JobJobFailedIterator struct {
	Event *JobJobFailed // Event containing the contract specifics and raw log

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
func (it *JobJobFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobJobFailed)
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
		it.Event = new(JobJobFailed)
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
func (it *JobJobFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobJobFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobJobFailed represents a JobFailed event raised by the Job contract.
type JobJobFailed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterJobFailed is a free log retrieval operation binding the contract event 0xf6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e.
//
// Solidity: event JobFailed()
func (_Job *JobFilterer) FilterJobFailed(opts *bind.FilterOpts) (*JobJobFailedIterator, error) {

	logs, sub, err := _Job.contract.FilterLogs(opts, "JobFailed")
	if err != nil {
		return nil, err
	}
	return &JobJobFailedIterator{contract: _Job.contract, event: "JobFailed", logs: logs, sub: sub}, nil
}

// WatchJobFailed is a free log subscription operation binding the contract event 0xf6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e.
//
// Solidity: event JobFailed()
func (_Job *JobFilterer) WatchJobFailed(opts *bind.WatchOpts, sink chan<- *JobJobFailed) (event.Subscription, error) {

	logs, sub, err := _Job.contract.WatchLogs(opts, "JobFailed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobJobFailed)
				if err := _Job.contract.UnpackLog(event, "JobFailed", log); err != nil {
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
func (_Job *JobFilterer) ParseJobFailed(log types.Log) (*JobJobFailed, error) {
	event := new(JobJobFailed)
	if err := _Job.contract.UnpackLog(event, "JobFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// JobJobSucceededIterator is returned from FilterJobSucceeded and is used to iterate over the raw logs and unpacked data for JobSucceeded events raised by the Job contract.
type JobJobSucceededIterator struct {
	Event *JobJobSucceeded // Event containing the contract specifics and raw log

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
func (it *JobJobSucceededIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(JobJobSucceeded)
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
		it.Event = new(JobJobSucceeded)
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
func (it *JobJobSucceededIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *JobJobSucceededIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// JobJobSucceeded represents a JobSucceeded event raised by the Job contract.
type JobJobSucceeded struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterJobSucceeded is a free log retrieval operation binding the contract event 0x7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b.
//
// Solidity: event JobSucceeded()
func (_Job *JobFilterer) FilterJobSucceeded(opts *bind.FilterOpts) (*JobJobSucceededIterator, error) {

	logs, sub, err := _Job.contract.FilterLogs(opts, "JobSucceeded")
	if err != nil {
		return nil, err
	}
	return &JobJobSucceededIterator{contract: _Job.contract, event: "JobSucceeded", logs: logs, sub: sub}, nil
}

// WatchJobSucceeded is a free log subscription operation binding the contract event 0x7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b.
//
// Solidity: event JobSucceeded()
func (_Job *JobFilterer) WatchJobSucceeded(opts *bind.WatchOpts, sink chan<- *JobJobSucceeded) (event.Subscription, error) {

	logs, sub, err := _Job.contract.WatchLogs(opts, "JobSucceeded")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(JobJobSucceeded)
				if err := _Job.contract.UnpackLog(event, "JobSucceeded", log); err != nil {
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
func (_Job *JobFilterer) ParseJobSucceeded(log types.Log) (*JobJobSucceeded, error) {
	event := new(JobJobSucceeded)
	if err := _Job.contract.UnpackLog(event, "JobSucceeded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
