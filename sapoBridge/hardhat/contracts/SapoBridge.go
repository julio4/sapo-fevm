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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sapoJob\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"JobExecutionRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"failAndRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"request\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"result\",\"type\":\"string\"}],\"name\":\"saveResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200225838038062002258833981810160405281019062000037919062000129565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200015b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000f182620000c4565b9050919050565b6200010381620000e4565b81146200010f57600080fd5b50565b6000815190506200012381620000f8565b92915050565b600060208284031215620001425762000141620000bf565b5b6000620001528482850162000112565b91505092915050565b6120ed806200016b6000396000f3fe6080604052600436106200005c5760003560e01c80631a3cbef414620000615780632c19988914620000a557806385afc1ab14620000c55780638dd1480214620000f3578063f947bead1462000121578063fb32c508146200014f575b600080fd5b3480156200006e57600080fd5b506200008d6004803603810190620000879190620008f0565b6200017f565b6040516200009c9190620009f0565b60405180910390f35b620000c36004803603810190620000bd919062000b76565b6200024e565b005b348015620000d257600080fd5b50620000f16004803603810190620000eb919062000bc7565b62000446565b005b3480156200010057600080fd5b506200011f6004803603810190620001199190620008f0565b620005d8565b005b3480156200012e57600080fd5b506200014d600480360381019062000147919062000bc7565b620006ad565b005b3480156200015c57600080fd5b50620001676200083f565b60405162000176919062000c3e565b60405180910390f35b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000208054806020026020016040519081016040528092919081815260200182805480156200024257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311620001f7575b50505050509050919050565b67016345785d8a00003410156200029c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002939062000cbc565b60405180910390fd5b600033600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051620002d09062000869565b620002dd92919062000cde565b604051809103906000f080158015620002fa573d6000803e3d6000fd5b5090507f7843f28b8cbd00b16af4f1efc48e12641ca994dc385bff44528582ad78f2409081836040516200033092919062000d83565b60405180910390a1600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f1935050505015801562000441573d6000803e3d6000fd5b505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620004d9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620004d09062000e2d565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff168160405160240162000503919062000e4f565b6040516020818303038152906040527f064d6b87000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516200058f919062000ec0565b600060405180830381855af49150503d8060008114620005cc576040519150601f19603f3d011682016040523d82523d6000602084013e620005d1565b606091505b5050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000669576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620006609062000f4f565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000740576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620007379062000e2d565b60405180910390fd5b8173ffffffffffffffffffffffffffffffffffffffff16816040516024016200076a919062000e4f565b6040516020818303038152906040527f7a373e6a000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051620007f6919062000ec0565b600060405180830381855af49150503d806000811462000833576040519150601f19603f3d011682016040523d82523d6000602084013e62000838565b606091505b5050505050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6111468062000f7283390190565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620008b8826200088b565b9050919050565b620008ca81620008ab565b8114620008d657600080fd5b50565b600081359050620008ea81620008bf565b92915050565b60006020828403121562000909576200090862000881565b5b60006200091984828501620008d9565b91505092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b6200095981620008ab565b82525050565b60006200096d83836200094e565b60208301905092915050565b6000602082019050919050565b6000620009938262000922565b6200099f81856200092d565b9350620009ac836200093e565b8060005b83811015620009e3578151620009c788826200095f565b9750620009d48362000979565b925050600181019050620009b0565b5085935050505092915050565b6000602082019050818103600083015262000a0c818462000986565b905092915050565b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000a698262000a1e565b810181811067ffffffffffffffff8211171562000a8b5762000a8a62000a2f565b5b80604052505050565b600062000aa062000877565b905062000aae828262000a5e565b919050565b600067ffffffffffffffff82111562000ad15762000ad062000a2f565b5b62000adc8262000a1e565b9050602081019050919050565b82818337600083830152505050565b600062000b0f62000b098462000ab3565b62000a94565b90508281526020810184848401111562000b2e5762000b2d62000a19565b5b62000b3b84828562000ae9565b509392505050565b600082601f83011262000b5b5762000b5a62000a14565b5b813562000b6d84826020860162000af8565b91505092915050565b60006020828403121562000b8f5762000b8e62000881565b5b600082013567ffffffffffffffff81111562000bb05762000baf62000886565b5b62000bbe8482850162000b43565b91505092915050565b6000806040838503121562000be15762000be062000881565b5b600062000bf185828601620008d9565b925050602083013567ffffffffffffffff81111562000c155762000c1462000886565b5b62000c238582860162000b43565b9150509250929050565b62000c3881620008ab565b82525050565b600060208201905062000c55600083018462000c2d565b92915050565b600082825260208201905092915050565b7f4d696e696d756d20636f6c61746572616c20697320302e315446494c00000000600082015250565b600062000ca4601c8362000c5b565b915062000cb18262000c6c565b602082019050919050565b6000602082019050818103600083015262000cd78162000c95565b9050919050565b600060408201905062000cf5600083018562000c2d565b62000d04602083018462000c2d565b9392505050565b600081519050919050565b60005b8381101562000d3657808201518184015260208101905062000d19565b60008484015250505050565b600062000d4f8262000d0b565b62000d5b818562000c5b565b935062000d6d81856020860162000d16565b62000d788162000a1e565b840191505092915050565b600060408201905062000d9a600083018562000c2d565b818103602083015262000dae818462000d42565b90509392505050565b7f4f6e6c79204272696467652063616e2063616c6c20746869732066756e63746960008201527f6f6e000000000000000000000000000000000000000000000000000000000000602082015250565b600062000e1560228362000c5b565b915062000e228262000db7565b604082019050919050565b6000602082019050818103600083015262000e488162000e06565b9050919050565b6000602082019050818103600083015262000e6b818462000d42565b905092915050565b600081519050919050565b600081905092915050565b600062000e968262000e73565b62000ea2818562000e7e565b935062000eb481856020860162000d16565b80840191505092915050565b600062000ece828462000e89565b915081905092915050565b7f4f6e6c79204f776e65722063616e2063616c6c20746869732066756e6374696f60008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b600062000f3760218362000c5b565b915062000f448262000ed9565b604082019050919050565b6000602082019050818103600083015262000f6a8162000f28565b905091905056fe60806040526000600160146101000a81548160ff021916908360028111156200002d576200002c620000e0565b5b02179055506040516200114638038062001146833981810160405281019062000057919062000179565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505050620001c0565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001418262000114565b9050919050565b620001538162000134565b81146200015f57600080fd5b50565b600081519050620001738162000148565b92915050565b600080604083850312156200019357620001926200010f565b5b6000620001a38582860162000162565b9250506020620001b68582860162000162565b9150509250929050565b610f7680620001d06000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80634e69d5601461006757806384de1ea514610085578063877db67d146100a3578063b422a51e146100bf578063de292789146100db578063fb32c508146100f9575b600080fd5b61006f610117565b60405161007c919061073b565b60405180910390f35b61008d61012e565b60405161009a9190610797565b60405180910390f35b6100bd60048036038101906100b8919061090c565b610157565b005b6100d960048036038101906100d4919061090c565b6102c6565b005b6100e3610504565b6040516100f091906109d4565b60405180910390f35b61010161069a565b60405161010e9190610797565b60405180910390f35b6000600160149054906101000a900460ff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146101e7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101de90610a42565b60405180910390fd5b600060028111156101fb576101fa6106c4565b5b600160149054906101000a900460ff16600281111561021d5761021c6106c4565b5b1461025d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161025490610aae565b60405180910390fd5b60018060146101000a81548160ff02191690836002811115610282576102816106c4565b5b021790555080600290816102969190610ce4565b507f7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b60405160405180910390a150565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610356576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034d90610a42565b60405180910390fd5b6000600281111561036a576103696106c4565b5b600160149054906101000a900460ff16600281111561038c5761038b6106c4565b5b146103cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c390610aae565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff164760405161041390610de7565b60006040518083038185875af1925050503d8060008114610450576040519150601f19603f3d011682016040523d82523d6000602084013e610455565b606091505b5050905080610499576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049090610e48565b60405180910390fd5b6002600160146101000a81548160ff021916908360028111156104bf576104be6106c4565b5b021790555081600290816104d39190610ce4565b507ff6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e60405160405180910390a15050565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614610594576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161058b90610eb4565b60405180910390fd5b600060028111156105a8576105a76106c4565b5b600160149054906101000a900460ff1660028111156105ca576105c96106c4565b5b0361060a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060190610f20565b60405180910390fd5b6002805461061790610afd565b80601f016020809104026020016040519081016040528092919081815260200182805461064390610afd565b80156106905780601f1061066557610100808354040283529160200191610690565b820191906000526020600020905b81548152906001019060200180831161067357829003601f168201915b5050505050905090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110610704576107036106c4565b5b50565b6000819050610715826106f3565b919050565b600061072582610707565b9050919050565b6107358161071a565b82525050565b6000602082019050610750600083018461072c565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061078182610756565b9050919050565b61079181610776565b82525050565b60006020820190506107ac6000830184610788565b92915050565b6000604051905090565b600080fd5b600080fd5b600080fd5b600080fd5b6000601f19601f8301169050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610819826107d0565b810181811067ffffffffffffffff82111715610838576108376107e1565b5b80604052505050565b600061084b6107b2565b90506108578282610810565b919050565b600067ffffffffffffffff821115610877576108766107e1565b5b610880826107d0565b9050602081019050919050565b82818337600083830152505050565b60006108af6108aa8461085c565b610841565b9050828152602081018484840111156108cb576108ca6107cb565b5b6108d684828561088d565b509392505050565b600082601f8301126108f3576108f26107c6565b5b813561090384826020860161089c565b91505092915050565b600060208284031215610922576109216107bc565b5b600082013567ffffffffffffffff8111156109405761093f6107c1565b5b61094c848285016108de565b91505092915050565b600081519050919050565b600082825260208201905092915050565b60005b8381101561098f578082015181840152602081019050610974565b60008484015250505050565b60006109a682610955565b6109b08185610960565b93506109c0818560208601610971565b6109c9816107d0565b840191505092915050565b600060208201905081810360008301526109ee818461099b565b905092915050565b7f43616c6c6572206973206e6f7420627269646765000000000000000000000000600082015250565b6000610a2c601483610960565b9150610a37826109f6565b602082019050919050565b60006020820190508181036000830152610a5b81610a1f565b9050919050565b7f4a6f62206973206e6f742070656e64696e670000000000000000000000000000600082015250565b6000610a98601283610960565b9150610aa382610a62565b602082019050919050565b60006020820190508181036000830152610ac781610a8b565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b60006002820490506001821680610b1557607f821691505b602082108103610b2857610b27610ace565b5b50919050565b60008190508160005260206000209050919050565b60006020601f8301049050919050565b600082821b905092915050565b600060088302610b907fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82610b53565b610b9a8683610b53565b95508019841693508086168417925050509392505050565b6000819050919050565b6000819050919050565b6000610be1610bdc610bd784610bb2565b610bbc565b610bb2565b9050919050565b6000819050919050565b610bfb83610bc6565b610c0f610c0782610be8565b848454610b60565b825550505050565b600090565b610c24610c17565b610c2f818484610bf2565b505050565b5b81811015610c5357610c48600082610c1c565b600181019050610c35565b5050565b601f821115610c9857610c6981610b2e565b610c7284610b43565b81016020851015610c81578190505b610c95610c8d85610b43565b830182610c34565b50505b505050565b600082821c905092915050565b6000610cbb60001984600802610c9d565b1980831691505092915050565b6000610cd48383610caa565b9150826002028217905092915050565b610ced82610955565b67ffffffffffffffff811115610d0657610d056107e1565b5b610d108254610afd565b610d1b828285610c57565b600060209050601f831160018114610d4e5760008415610d3c578287015190505b610d468582610cc8565b865550610dae565b601f198416610d5c86610b2e565b60005b82811015610d8457848901518255600182019150602085019450602081019050610d5f565b86831015610da15784890151610d9d601f891682610caa565b8355505b6001600288020188555050505b505050505050565b600081905092915050565b50565b6000610dd1600083610db6565b9150610ddc82610dc1565b600082019050919050565b6000610df282610dc4565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610e32601483610960565b9150610e3d82610dfc565b602082019050919050565b60006020820190508181036000830152610e6181610e25565b9050919050565b7f43616c6c6572206973206e6f7420696e69746961746f72000000000000000000600082015250565b6000610e9e601783610960565b9150610ea982610e68565b602082019050919050565b60006020820190508181036000830152610ecd81610e91565b9050919050565b7f4a6f62206973206e6f7420636f6d706c65746564000000000000000000000000600082015250565b6000610f0a601483610960565b9150610f1582610ed4565b602082019050919050565b60006020820190508181036000830152610f3981610efd565b905091905056fea2646970667358221220c6e1d6395916b7526155d5a58e6dab059a8428f8b12edca0e99dddb705f8088964736f6c63430008110033a2646970667358221220bdf3e29a9be9e15f8fd0dccc795e15c395572ad3d9885ce8d49e2220c7ac0d5964736f6c63430008110033",
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

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address)
func (_Contracts *ContractsCaller) GetBridgeAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "getBridgeAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address)
func (_Contracts *ContractsSession) GetBridgeAddress() (common.Address, error) {
	return _Contracts.Contract.GetBridgeAddress(&_Contracts.CallOpts)
}

// GetBridgeAddress is a free data retrieval call binding the contract method 0xfb32c508.
//
// Solidity: function getBridgeAddress() view returns(address)
func (_Contracts *ContractsCallerSession) GetBridgeAddress() (common.Address, error) {
	return _Contracts.Contract.GetBridgeAddress(&_Contracts.CallOpts)
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

// FailAndRefund is a paid mutator transaction binding the contract method 0xf947bead.
//
// Solidity: function failAndRefund(address job, string reason) returns()
func (_Contracts *ContractsTransactor) FailAndRefund(opts *bind.TransactOpts, job common.Address, reason string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "failAndRefund", job, reason)
}

// FailAndRefund is a paid mutator transaction binding the contract method 0xf947bead.
//
// Solidity: function failAndRefund(address job, string reason) returns()
func (_Contracts *ContractsSession) FailAndRefund(job common.Address, reason string) (*types.Transaction, error) {
	return _Contracts.Contract.FailAndRefund(&_Contracts.TransactOpts, job, reason)
}

// FailAndRefund is a paid mutator transaction binding the contract method 0xf947bead.
//
// Solidity: function failAndRefund(address job, string reason) returns()
func (_Contracts *ContractsTransactorSession) FailAndRefund(job common.Address, reason string) (*types.Transaction, error) {
	return _Contracts.Contract.FailAndRefund(&_Contracts.TransactOpts, job, reason)
}

// Request is a paid mutator transaction binding the contract method 0x2c199889.
//
// Solidity: function request(string cid) payable returns()
func (_Contracts *ContractsTransactor) Request(opts *bind.TransactOpts, cid string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "request", cid)
}

// Request is a paid mutator transaction binding the contract method 0x2c199889.
//
// Solidity: function request(string cid) payable returns()
func (_Contracts *ContractsSession) Request(cid string) (*types.Transaction, error) {
	return _Contracts.Contract.Request(&_Contracts.TransactOpts, cid)
}

// Request is a paid mutator transaction binding the contract method 0x2c199889.
//
// Solidity: function request(string cid) payable returns()
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

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address bridgeAddress) returns()
func (_Contracts *ContractsTransactor) SetBridge(opts *bind.TransactOpts, bridgeAddress common.Address) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "setBridge", bridgeAddress)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address bridgeAddress) returns()
func (_Contracts *ContractsSession) SetBridge(bridgeAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetBridge(&_Contracts.TransactOpts, bridgeAddress)
}

// SetBridge is a paid mutator transaction binding the contract method 0x8dd14802.
//
// Solidity: function setBridge(address bridgeAddress) returns()
func (_Contracts *ContractsTransactorSession) SetBridge(bridgeAddress common.Address) (*types.Transaction, error) {
	return _Contracts.Contract.SetBridge(&_Contracts.TransactOpts, bridgeAddress)
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
