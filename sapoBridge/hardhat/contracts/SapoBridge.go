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

// ContractsMetaData contains all meta data concerning the Contracts contract.
var ContractsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sapoJob\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"JobExecutionRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"}],\"name\":\"failAndRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"request\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"saveResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b506040516119883803806119888339818101604052810190610032919061011c565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610149565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100e9826100be565b9050919050565b6100f9816100de565b811461010457600080fd5b50565b600081519050610116816100f0565b92915050565b600060208284031215610132576101316100b9565b5b600061014084828501610107565b91505092915050565b611830806101586000396000f3fe608060405234801561001057600080fd5b50600436106100575760003560e01c80630ae523e51461005c5780631a3cbef4146100785780632c199889146100a85780636f82f8fd146100c457806385afc1ab146100f4575b600080fd5b610076600480360381019061007191906105cf565b610110565b005b610092600480360381019061008d91906105cf565b610201565b60405161009f91906106ba565b60405180910390f35b6100c260048036038101906100bd9190610822565b6102ce565b005b6100de60048036038101906100d991906108a1565b610405565b6040516100eb91906108f0565b60405180910390f35b61010e6004803603810190610109919061090b565b610453565b005b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161461019e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610195906109ea565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16638c45c8656040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156101e657600080fd5b505af11580156101fa573d6000803e3d6000fd5b5050505050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156102c257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610278575b50505050509050919050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516102ff90610550565b61030991906108f0565b604051809103906000f080158015610325573d6000803e3d6000fd5b5090507f7843f28b8cbd00b16af4f1efc48e12641ca994dc385bff44528582ad78f240908183604051610359929190610a78565b60405180910390a1600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050565b6002602052816000526040600020818154811061042157600080fd5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146104e1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104d8906109ea565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663877db67d826040518263ffffffff1660e01b815260040161051a9190610aa8565b600060405180830381600087803b15801561053457600080fd5b505af1158015610548573d6000803e3d6000fd5b505050505050565b610d3080610acb83390190565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061059c82610571565b9050919050565b6105ac81610591565b81146105b757600080fd5b50565b6000813590506105c9816105a3565b92915050565b6000602082840312156105e5576105e4610567565b5b60006105f3848285016105ba565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b61063181610591565b82525050565b60006106438383610628565b60208301905092915050565b6000602082019050919050565b6000610667826105fc565b6106718185610607565b935061067c83610618565b8060005b838110156106ad5781516106948882610637565b975061069f8361064f565b925050600181019050610680565b5085935050505092915050565b600060208201905081810360008301526106d4818461065c565b905092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61072f826106e6565b810181811067ffffffffffffffff8211171561074e5761074d6106f7565b5b80604052505050565b600061076161055d565b905061076d8282610726565b919050565b600067ffffffffffffffff82111561078d5761078c6106f7565b5b610796826106e6565b9050602081019050919050565b82818337600083830152505050565b60006107c56107c084610772565b610757565b9050828152602081018484840111156107e1576107e06106e1565b5b6107ec8482856107a3565b509392505050565b600082601f830112610809576108086106dc565b5b81356108198482602086016107b2565b91505092915050565b60006020828403121561083857610837610567565b5b600082013567ffffffffffffffff8111156108565761085561056c565b5b610862848285016107f4565b91505092915050565b6000819050919050565b61087e8161086b565b811461088957600080fd5b50565b60008135905061089b81610875565b92915050565b600080604083850312156108b8576108b7610567565b5b60006108c6858286016105ba565b92505060206108d78582860161088c565b9150509250929050565b6108ea81610591565b82525050565b600060208201905061090560008301846108e1565b92915050565b6000806040838503121561092257610921610567565b5b6000610930858286016105ba565b925050602083013567ffffffffffffffff8111156109515761095061056c565b5b61095d858286016107f4565b9150509250929050565b600082825260208201905092915050565b7f4f6e6c79206272696467652063616e2063616c6c20746869732066756e63746960008201527f6f6e2e0000000000000000000000000000000000000000000000000000000000602082015250565b60006109d4602383610967565b91506109df82610978565b604082019050919050565b60006020820190508181036000830152610a03816109c7565b9050919050565b600081519050919050565b60005b83811015610a33578082015181840152602081019050610a18565b60008484015250505050565b6000610a4a82610a0a565b610a548185610967565b9350610a64818560208601610a15565b610a6d816106e6565b840191505092915050565b6000604082019050610a8d60008301856108e1565b8181036020830152610a9f8184610a3f565b90509392505050565b60006020820190508181036000830152610ac28184610a3f565b90509291505056fe60806040526000600160146101000a81548160ff021916908315150217905550604051610d30380380610d308339818101604052810190610040919061012a565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610157565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100f7826100cc565b9050919050565b610107816100ec565b811461011257600080fd5b50565b600081519050610124816100fe565b92915050565b6000602082840312156101405761013f6100c7565b5b600061014e84828501610115565b91505092915050565b610bca806101666000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063877db67d146100515780638c45c8651461006d5780639d9a7fe914610077578063de29278914610095575b600080fd5b61006b60048036038101906100669190610602565b6100b3565b005b6100756101b6565b005b61007f61035c565b60405161008c9190610666565b60405180910390f35b61009d61036f565b6040516100aa9190610700565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610143576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013a9061076e565b60405180910390fd5b600160149054906101000a900460ff161561015d57600080fd5b60018060146101000a81548160ff021916908315150217905550806002908161018691906109a4565b507f85c9037835c422f4d8688ffaff3f926bde639d28fa6a92c2960c123aa96191cb60405160405180910390a150565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610246576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023d9061076e565b60405180910390fd5b600160149054906101000a900460ff161561026057600080fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16476040516102a790610aa7565b60006040518083038185875af1925050503d80600081146102e4576040519150601f19603f3d011682016040523d82523d6000602084013e6102e9565b606091505b505090508061032d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032490610b08565b60405180910390fd5b7ff6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e60405160405180910390a150565b600160149054906101000a900460ff1681565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146103ff576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f690610b74565b60405180910390fd5b600160149054906101000a900460ff1661041857600080fd5b60028054610425906107bd565b80601f0160208091040260200160405190810160405280929190818152602001828054610451906107bd565b801561049e5780601f106104735761010080835404028352916020019161049e565b820191906000526020600020905b81548152906001019060200180831161048157829003601f168201915b5050505050905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b61050f826104c6565b810181811067ffffffffffffffff8211171561052e5761052d6104d7565b5b80604052505050565b60006105416104a8565b905061054d8282610506565b919050565b600067ffffffffffffffff82111561056d5761056c6104d7565b5b610576826104c6565b9050602081019050919050565b82818337600083830152505050565b60006105a56105a084610552565b610537565b9050828152602081018484840111156105c1576105c06104c1565b5b6105cc848285610583565b509392505050565b600082601f8301126105e9576105e86104bc565b5b81356105f9848260208601610592565b91505092915050565b600060208284031215610618576106176104b2565b5b600082013567ffffffffffffffff811115610636576106356104b7565b5b610642848285016105d4565b91505092915050565b60008115159050919050565b6106608161064b565b82525050565b600060208201905061067b6000830184610657565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156106bb5780820151818401526020810190506106a0565b60008484015250505050565b60006106d282610681565b6106dc818561068c565b93506106ec81856020860161069d565b6106f5816104c6565b840191505092915050565b6000602082019050818103600083015261071a81846106c7565b905092915050565b7f43616c6c6572206973206e6f7420627269646765000000000000000000000000600082015250565b600061075860148361068c565b915061076382610722565b602082019050919050565b600060208201905081810360008301526107878161074b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806107d557607f821691505b6020821081036107e8576107e761078e565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026108507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610813565b61085a8683610813565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006108a161089c61089784610872565b61087c565b610872565b9050919050565b6000819050919050565b6108bb83610886565b6108cf6108c7826108a8565b848454610820565b825550505050565b600090565b6108e46108d7565b6108ef8184846108b2565b505050565b5b81811015610913576109086000826108dc565b6001810190506108f5565b5050565b601f82111561095857610929816107ee565b61093284610803565b81016020851015610941578190505b61095561094d85610803565b8301826108f4565b50505b505050565b600082821c905092915050565b600061097b6000198460080261095d565b1980831691505092915050565b6000610994838361096a565b9150826002028217905092915050565b6109ad82610681565b67ffffffffffffffff8111156109c6576109c56104d7565b5b6109d082546107bd565b6109db828285610917565b600060209050601f831160018114610a0e57600084156109fc578287015190505b610a068582610988565b865550610a6e565b601f198416610a1c866107ee565b60005b82811015610a4457848901518255600182019150602085019450602081019050610a1f565b86831015610a615784890151610a5d601f89168261096a565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b50565b6000610a91600083610a76565b9150610a9c82610a81565b600082019050919050565b6000610ab282610a84565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610af260148361068c565b9150610afd82610abc565b602082019050919050565b60006020820190508181036000830152610b2181610ae5565b9050919050565b7f43616c6c6572206973206e6f7420696e69746961746f72000000000000000000600082015250565b6000610b5e60178361068c565b9150610b6982610b28565b602082019050919050565b60006020820190508181036000830152610b8d81610b51565b905091905056fea264697066735822122009b3210d60d5b586d43673e3466d53cab2595952ffc385c9db75e43a8c509d7e64736f6c63430008110033a2646970667358221220909480709ffc033fbcb77f555328bc5f2d9bd5a3d8ef30bc58d637e7a42b419b64736f6c63430008110033",
}

// ContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractsMetaData.ABI instead.
var ContractsABI = ContractsMetaData.ABI

// ContractsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractsMetaData.Bin instead.
var ContractsBin = ContractsMetaData.Bin

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend, bridgeAddress common.Address) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := ContractsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractsBin), backend, bridgeAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// GetJobs is a free data retrieval call binding the contract method 0x1a3cbef4.
//
// Solidity: function getJobs(address user) view returns(address[])
func (_Contracts *ContractsCaller) GetJobs(opts *bind.CallOpts, user common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getJobs", user)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetJobs is a free data retrieval call binding the contract method 0x1a3cbef4.
//
// Solidity: function getJobs(address user) view returns(address[])
func (_Contracts *ContractsSession) GetJobs(user common.Address) ([]common.Address, error) {
	return _Contracts.Contract.GetJobs(&_Contracts.CallOpts, user)
}

// GetJobs is a free data retrieval call binding the contract method 0x1a3cbef4.
//
// Solidity: function getJobs(address user) view returns(address[])
func (_Contracts *ContractsCallerSession) GetJobs(user common.Address) ([]common.Address, error) {
	return _Contracts.Contract.GetJobs(&_Contracts.CallOpts, user)
}

// Jobs is a free data retrieval call binding the contract method 0x6f82f8fd.
//
// Solidity: function jobs(address , uint256 ) view returns(address)
func (_Contracts *ContractsCaller) Jobs(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "jobs", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Jobs is a free data retrieval call binding the contract method 0x6f82f8fd.
//
// Solidity: function jobs(address , uint256 ) view returns(address)
func (_Contracts *ContractsSession) Jobs(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Contracts.Contract.Jobs(&_Contracts.CallOpts, arg0, arg1)
}

// Jobs is a free data retrieval call binding the contract method 0x6f82f8fd.
//
// Solidity: function jobs(address , uint256 ) view returns(address)
func (_Contracts *ContractsCallerSession) Jobs(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Contracts.Contract.Jobs(&_Contracts.CallOpts, arg0, arg1)
}

// FailAndRefund is a paid mutator transaction binding the contract method 0x0ae523e5.
//
// Solidity: function failAndRefund(address job) returns()
func (_Contracts *ContractsTransactor) FailAndRefund(opts *bind.TransactOpts, job common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "failAndRefund", job)
}

// FailAndRefund is a paid mutator transaction binding the contract method 0x0ae523e5.
//
// Solidity: function failAndRefund(address job) returns()
func (_Contracts *ContractsSession) FailAndRefund(job common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.FailAndRefund(&_Contracts.TransactOpts, job)
}

// FailAndRefund is a paid mutator transaction binding the contract method 0x0ae523e5.
//
// Solidity: function failAndRefund(address job) returns()
func (_Contracts *ContractsTransactorSession) FailAndRefund(job common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.FailAndRefund(&_Contracts.TransactOpts, job)
}

// Request is a paid mutator transaction binding the contract method 0x2c199889.
//
// Solidity: function request(string cid) returns()
func (_Contracts *ContractsTransactor) Request(opts *bind.TransactOpts, cid string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "request", cid)
}

// Request is a paid mutator transaction binding the contract method 0x2c199889.
//
// Solidity: function request(string cid) returns()
func (_Contracts *ContractsSession) Request(cid string) (*types.Transaction, error) {
	return _Contracts.Contract.Request(&_Contracts.TransactOpts, cid)
}

// Request is a paid mutator transaction binding the contract method 0x2c199889.
//
// Solidity: function request(string cid) returns()
func (_Contracts *ContractsTransactorSession) Request(cid string) (*types.Transaction, error) {
	return _Contracts.Contract.Request(&_Contracts.TransactOpts, cid)
}

// SaveResult is a paid mutator transaction binding the contract method 0x85afc1ab.
//
// Solidity: function saveResult(address job, string result) returns()
func (_Contracts *ContractsTransactor) SaveResult(opts *bind.TransactOpts, job common.Address, result string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "saveResult", job, result)
}

// SaveResult is a paid mutator transaction binding the contract method 0x85afc1ab.
//
// Solidity: function saveResult(address job, string result) returns()
func (_Contracts *ContractsSession) SaveResult(job common.Address, result string) (*types.Transaction, error) {
	return _Contracts.Contract.SaveResult(&_Contracts.TransactOpts, job, result)
}

// SaveResult is a paid mutator transaction binding the contract method 0x85afc1ab.
//
// Solidity: function saveResult(address job, string result) returns()
func (_Contracts *ContractsTransactorSession) SaveResult(job common.Address, result string) (*types.Transaction, error) {
	return _Contracts.Contract.SaveResult(&_Contracts.TransactOpts, job, result)
}

// ContractsJobExecutionRequestIterator is returned from FilterJobExecutionRequest and is used to iterate over the raw logs and unpacked data for JobExecutionRequest events raised by the Contracts contract.
type ContractsJobExecutionRequestIterator struct {
	Event *ContractsJobExecutionRequest // Event containing the contract specifics and raw log

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
func (it *ContractsJobExecutionRequestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractsJobExecutionRequest)
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
		it.Event = new(ContractsJobExecutionRequest)
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
func (it *ContractsJobExecutionRequestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractsJobExecutionRequestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractsJobExecutionRequest represents a JobExecutionRequest event raised by the Contracts contract.
type ContractsJobExecutionRequest struct {
	SapoJob common.Address
	Cid     string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterJobExecutionRequest is a free log retrieval operation binding the contract event 0x7843f28b8cbd00b16af4f1efc48e12641ca994dc385bff44528582ad78f24090.
//
// Solidity: event JobExecutionRequest(address sapoJob, string cid)
func (_Contracts *ContractsFilterer) FilterJobExecutionRequest(opts *bind.FilterOpts) (*ContractsJobExecutionRequestIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "JobExecutionRequest")
	if err != nil {
		return nil, err
	}
	return &ContractsJobExecutionRequestIterator{contract: _Contracts.contract, event: "JobExecutionRequest", logs: logs, sub: sub}, nil
}

// WatchJobExecutionRequest is a free log subscription operation binding the contract event 0x7843f28b8cbd00b16af4f1efc48e12641ca994dc385bff44528582ad78f24090.
//
// Solidity: event JobExecutionRequest(address sapoJob, string cid)
func (_Contracts *ContractsFilterer) WatchJobExecutionRequest(opts *bind.WatchOpts, sink chan<- *ContractsJobExecutionRequest) (event.Subscription, error) {

	logs, sub, err := _Contracts.contract.WatchLogs(opts, "JobExecutionRequest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractsJobExecutionRequest)
				if err := _Contracts.contract.UnpackLog(event, "JobExecutionRequest", log); err != nil {
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

// ParseJobExecutionRequest is a log parse operation binding the contract event 0x7843f28b8cbd00b16af4f1efc48e12641ca994dc385bff44528582ad78f24090.
//
// Solidity: event JobExecutionRequest(address sapoJob, string cid)
func (_Contracts *ContractsFilterer) ParseJobExecutionRequest(log types.Log) (*ContractsJobExecutionRequest, error) {
	event := new(ContractsJobExecutionRequest)
	if err := _Contracts.contract.UnpackLog(event, "JobExecutionRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
