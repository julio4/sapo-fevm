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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sapoJob\",\"type\":\"address\"}],\"name\":\"JobExecutionRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"}],\"name\":\"failAndRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"jobs\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"request\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"saveResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50604051611fdf380380611fdf8339818101604052810190610032919061011c565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050610149565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b60006100e9826100be565b9050919050565b6100f9816100de565b811461010457600080fd5b50565b600081519050610116816100f0565b92915050565b600060208284031215610132576101316100b9565b5b600061014084828501610107565b91505092915050565b611e87806101586000396000f3fe60806040523480156200001157600080fd5b50600436106200005e5760003560e01c806301d511f114620000635780630ae523e514620000835780631a3cbef414620000a35780636f82f8fd14620000d957806385afc1ab146200010f575b600080fd5b6200008160048036038101906200007b91906200070c565b6200012f565b005b620000a160048036038101906200009b9190620007f6565b62000270565b005b620000c16004803603810190620000bb9190620007f6565b62000366565b604051620000d09190620008f6565b60405180910390f35b620000f76004803603810190620000f1919062000955565b62000435565b604051620001069190620009ad565b60405180910390f35b6200012d6004803603810190620001279190620009ca565b62000484565b005b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168383604051620001649062000588565b620001729392919062000ab9565b604051809103906000f0801580156200018f573d6000803e3d6000fd5b5090507fcba8ad1dbc943e28edc86bdb6f019b6378796e41c26f97cedd0fbcc94d5ccde081604051620001c39190620009ad565b60405180910390a1600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000301576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002f89062000b7a565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16638c45c8656040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156200034a57600080fd5b505af11580156200035f573d6000803e3d6000fd5b5050505050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156200042957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311620003de575b50505050509050919050565b600260205281600052604060002081815481106200045257600080fd5b906000526020600020016000915091509054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000515576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200050c9062000b7a565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff1663877db67d826040518263ffffffff1660e01b815260040162000550919062000b9c565b600060405180830381600087803b1580156200056b57600080fd5b505af115801562000580573d6000803e3d6000fd5b505050505050565b6112918062000bc183390190565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620005ff82620005b4565b810181811067ffffffffffffffff82111715620006215762000620620005c5565b5b80604052505050565b60006200063662000596565b9050620006448282620005f4565b919050565b600067ffffffffffffffff821115620006675762000666620005c5565b5b6200067282620005b4565b9050602081019050919050565b82818337600083830152505050565b6000620006a56200069f8462000649565b6200062a565b905082815260208101848484011115620006c457620006c3620005af565b5b620006d18482856200067f565b509392505050565b600082601f830112620006f157620006f0620005aa565b5b8135620007038482602086016200068e565b91505092915050565b60008060408385031215620007265762000725620005a0565b5b600083013567ffffffffffffffff811115620007475762000746620005a5565b5b6200075585828601620006d9565b925050602083013567ffffffffffffffff811115620007795762000778620005a5565b5b6200078785828601620006d9565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620007be8262000791565b9050919050565b620007d081620007b1565b8114620007dc57600080fd5b50565b600081359050620007f081620007c5565b92915050565b6000602082840312156200080f576200080e620005a0565b5b60006200081f84828501620007df565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6200085f81620007b1565b82525050565b600062000873838362000854565b60208301905092915050565b6000602082019050919050565b6000620008998262000828565b620008a5818562000833565b9350620008b28362000844565b8060005b83811015620008e9578151620008cd888262000865565b9750620008da836200087f565b925050600181019050620008b6565b5085935050505092915050565b600060208201905081810360008301526200091281846200088c565b905092915050565b6000819050919050565b6200092f816200091a565b81146200093b57600080fd5b50565b6000813590506200094f8162000924565b92915050565b600080604083850312156200096f576200096e620005a0565b5b60006200097f85828601620007df565b925050602062000992858286016200093e565b9150509250929050565b620009a781620007b1565b82525050565b6000602082019050620009c460008301846200099c565b92915050565b60008060408385031215620009e457620009e3620005a0565b5b6000620009f485828601620007df565b925050602083013567ffffffffffffffff81111562000a185762000a17620005a5565b5b62000a2685828601620006d9565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b8381101562000a6c57808201518184015260208101905062000a4f565b60008484015250505050565b600062000a858262000a30565b62000a91818562000a3b565b935062000aa381856020860162000a4c565b62000aae81620005b4565b840191505092915050565b600060608201905062000ad060008301866200099c565b818103602083015262000ae4818562000a78565b9050818103604083015262000afa818462000a78565b9050949350505050565b7f4f6e6c79206272696467652063616e2063616c6c20746869732066756e63746960008201527f6f6e2e0000000000000000000000000000000000000000000000000000000000602082015250565b600062000b6260238362000a3b565b915062000b6f8262000b04565b604082019050919050565b6000602082019050818103600083015262000b958162000b53565b9050919050565b6000602082019050818103600083015262000bb8818462000a78565b90509291505056fe60806040526000600460006101000a81548160ff02191690831515021790555060405162001291380380620012918339818101604052810190620000449190620002ea565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160029081620000d69190620005cf565b508060039081620000e89190620005cf565b50505050620006b6565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001338262000106565b9050919050565b620001458162000126565b81146200015157600080fd5b50565b60008151905062000165816200013a565b92915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b620001c08262000175565b810181811067ffffffffffffffff82111715620001e257620001e162000186565b5b80604052505050565b6000620001f7620000f2565b9050620002058282620001b5565b919050565b600067ffffffffffffffff82111562000228576200022762000186565b5b620002338262000175565b9050602081019050919050565b60005b838110156200026057808201518184015260208101905062000243565b60008484015250505050565b6000620002836200027d846200020a565b620001eb565b905082815260208101848484011115620002a257620002a162000170565b5b620002af84828562000240565b509392505050565b600082601f830112620002cf57620002ce6200016b565b5b8151620002e18482602086016200026c565b91505092915050565b600080600060608486031215620003065762000305620000fc565b5b6000620003168682870162000154565b935050602084015167ffffffffffffffff8111156200033a576200033962000101565b5b6200034886828701620002b7565b925050604084015167ffffffffffffffff8111156200036c576200036b62000101565b5b6200037a86828701620002b7565b9150509250925092565b600081519050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680620003d757607f821691505b602082108103620003ed57620003ec6200038f565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302620004577fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8262000418565b62000463868362000418565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000620004b0620004aa620004a4846200047b565b62000485565b6200047b565b9050919050565b6000819050919050565b620004cc836200048f565b620004e4620004db82620004b7565b84845462000425565b825550505050565b600090565b620004fb620004ec565b62000508818484620004c1565b505050565b5b81811015620005305762000524600082620004f1565b6001810190506200050e565b5050565b601f8211156200057f576200054981620003f3565b620005548462000408565b8101602085101562000564578190505b6200057c620005738562000408565b8301826200050d565b50505b505050565b600082821c905092915050565b6000620005a46000198460080262000584565b1980831691505092915050565b6000620005bf838362000591565b9150826002028217905092915050565b620005da8262000384565b67ffffffffffffffff811115620005f657620005f562000186565b5b620006028254620003be565b6200060f82828562000534565b600060209050601f83116001811462000647576000841562000632578287015190505b6200063e8582620005b1565b865550620006ae565b601f1984166200065786620003f3565b60005b8281101562000681578489015182556001820191506020850194506020810190506200065a565b86831015620006a157848901516200069d601f89168262000591565b8355505b6001600288020188555050505b505050505050565b610bcb80620006c66000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c8063877db67d146100515780638c45c8651461006d5780639d9a7fe914610077578063de29278914610095575b600080fd5b61006b60048036038101906100669190610603565b6100b3565b005b6100756101b7565b005b61007f61035d565b60405161008c9190610667565b60405180910390f35b61009d610370565b6040516100aa9190610701565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610143576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161013a9061076f565b60405180910390fd5b600460009054906101000a900460ff161561015d57600080fd5b6001600460006101000a81548160ff021916908315150217905550806005908161018791906109a5565b507f85c9037835c422f4d8688ffaff3f926bde639d28fa6a92c2960c123aa96191cb60405160405180910390a150565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610247576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023e9061076f565b60405180910390fd5b600460009054906101000a900460ff161561026157600080fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16476040516102a890610aa8565b60006040518083038185875af1925050503d80600081146102e5576040519150601f19603f3d011682016040523d82523d6000602084013e6102ea565b606091505b505090508061032e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161032590610b09565b60405180910390fd5b7ff6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e60405160405180910390a150565b600460009054906101000a900460ff1681565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610400576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103f790610b75565b60405180910390fd5b600460009054906101000a900460ff1661041957600080fd5b60058054610426906107be565b80601f0160208091040260200160405190810160405280929190818152602001828054610452906107be565b801561049f5780601f106104745761010080835404028352916020019161049f565b820191906000526020600020905b81548152906001019060200180831161048257829003601f168201915b5050505050905090565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610510826104c7565b810181811067ffffffffffffffff8211171561052f5761052e6104d8565b5b80604052505050565b60006105426104a9565b905061054e8282610507565b919050565b600067ffffffffffffffff82111561056e5761056d6104d8565b5b610577826104c7565b9050602081019050919050565b82818337600083830152505050565b60006105a66105a184610553565b610538565b9050828152602081018484840111156105c2576105c16104c2565b5b6105cd848285610584565b509392505050565b600082601f8301126105ea576105e96104bd565b5b81356105fa848260208601610593565b91505092915050565b600060208284031215610619576106186104b3565b5b600082013567ffffffffffffffff811115610637576106366104b8565b5b610643848285016105d5565b91505092915050565b60008115159050919050565b6106618161064c565b82525050565b600060208201905061067c6000830184610658565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156106bc5780820151818401526020810190506106a1565b60008484015250505050565b60006106d382610682565b6106dd818561068d565b93506106ed81856020860161069e565b6106f6816104c7565b840191505092915050565b6000602082019050818103600083015261071b81846106c8565b905092915050565b7f43616c6c6572206973206e6f7420627269646765000000000000000000000000600082015250565b600061075960148361068d565b915061076482610723565b602082019050919050565b600060208201905081810360008301526107888161074c565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b600060028204905060018216806107d657607f821691505b6020821081036107e9576107e861078f565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b6000600883026108517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610814565b61085b8683610814565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b60006108a261089d61089884610873565b61087d565b610873565b9050919050565b6000819050919050565b6108bc83610887565b6108d06108c8826108a9565b848454610821565b825550505050565b600090565b6108e56108d8565b6108f08184846108b3565b505050565b5b81811015610914576109096000826108dd565b6001810190506108f6565b5050565b601f8211156109595761092a816107ef565b61093384610804565b81016020851015610942578190505b61095661094e85610804565b8301826108f5565b50505b505050565b600082821c905092915050565b600061097c6000198460080261095e565b1980831691505092915050565b6000610995838361096b565b9150826002028217905092915050565b6109ae82610682565b67ffffffffffffffff8111156109c7576109c66104d8565b5b6109d182546107be565b6109dc828285610918565b600060209050601f831160018114610a0f57600084156109fd578287015190505b610a078582610989565b865550610a6f565b601f198416610a1d866107ef565b60005b82811015610a4557848901518255600182019150602085019450602081019050610a20565b86831015610a625784890151610a5e601f89168261096b565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b50565b6000610a92600083610a77565b9150610a9d82610a82565b600082019050919050565b6000610ab382610a85565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610af360148361068d565b9150610afe82610abd565b602082019050919050565b60006020820190508181036000830152610b2281610ae6565b9050919050565b7f43616c6c6572206973206e6f7420696e69746961746f72000000000000000000600082015250565b6000610b5f60178361068d565b9150610b6a82610b29565b602082019050919050565b60006020820190508181036000830152610b8e81610b52565b905091905056fea2646970667358221220d2106e450511dfe68b1b3ca96c8199857daac610472d6e8de608ce0e4aec43a564736f6c63430008110033a264697066735822122041102bf17df2242897501c78378a681e2c1b4438f683376a9f581a6f73fbe5a464736f6c63430008110033",
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

// Request is a paid mutator transaction binding the contract method 0x01d511f1.
//
// Solidity: function request(string image, string cid) returns()
func (_Contracts *ContractsTransactor) Request(opts *bind.TransactOpts, image string, cid string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "request", image, cid)
}

// Request is a paid mutator transaction binding the contract method 0x01d511f1.
//
// Solidity: function request(string image, string cid) returns()
func (_Contracts *ContractsSession) Request(image string, cid string) (*types.Transaction, error) {
	return _Contracts.Contract.Request(&_Contracts.TransactOpts, image, cid)
}

// Request is a paid mutator transaction binding the contract method 0x01d511f1.
//
// Solidity: function request(string image, string cid) returns()
func (_Contracts *ContractsTransactorSession) Request(image string, cid string) (*types.Transaction, error) {
	return _Contracts.Contract.Request(&_Contracts.TransactOpts, image, cid)
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
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterJobExecutionRequest is a free log retrieval operation binding the contract event 0xcba8ad1dbc943e28edc86bdb6f019b6378796e41c26f97cedd0fbcc94d5ccde0.
//
// Solidity: event JobExecutionRequest(address sapoJob)
func (_Contracts *ContractsFilterer) FilterJobExecutionRequest(opts *bind.FilterOpts) (*ContractsJobExecutionRequestIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "JobExecutionRequest")
	if err != nil {
		return nil, err
	}
	return &ContractsJobExecutionRequestIterator{contract: _Contracts.contract, event: "JobExecutionRequest", logs: logs, sub: sub}, nil
}

// WatchJobExecutionRequest is a free log subscription operation binding the contract event 0xcba8ad1dbc943e28edc86bdb6f019b6378796e41c26f97cedd0fbcc94d5ccde0.
//
// Solidity: event JobExecutionRequest(address sapoJob)
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

// ParseJobExecutionRequest is a log parse operation binding the contract event 0xcba8ad1dbc943e28edc86bdb6f019b6378796e41c26f97cedd0fbcc94d5ccde0.
//
// Solidity: event JobExecutionRequest(address sapoJob)
func (_Contracts *ContractsFilterer) ParseJobExecutionRequest(log types.Log) (*ContractsJobExecutionRequest, error) {
	event := new(ContractsJobExecutionRequest)
	if err := _Contracts.contract.UnpackLog(event, "JobExecutionRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
