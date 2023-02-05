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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sapoJob\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cid1\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cid2\",\"type\":\"bytes32\"}],\"name\":\"JobExecutionRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"bytes64ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"}],\"name\":\"failAndRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cid1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid2\",\"type\":\"bytes32\"}],\"name\":\"request\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"result1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"result2\",\"type\":\"bytes32\"}],\"name\":\"saveResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200257538038062002575833981810160405281019062000037919062000129565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200015b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000f182620000c4565b9050919050565b6200010381620000e4565b81146200010f57600080fd5b50565b6000815190506200012381620000f8565b92915050565b600060208284031215620001425762000141620000bf565b5b6000620001528482850162000112565b91505092915050565b61240a806200016b6000396000f3fe6080604052600436106200007a5760003560e01c80631d90e28711620000555780631d90e287146200013557806351bffe8414620001635780638dd148021462000183578063fb32c50814620001b1576200007a565b80630ae523e5146200007f5780631453d75614620000ad5780631a3cbef414620000f1575b600080fd5b3480156200008c57600080fd5b50620000ab6004803603810190620000a5919062000aeb565b620001e1565b005b348015620000ba57600080fd5b50620000d96004803603810190620000d3919062000b58565b620002d9565b604051620000e8919062000c39565b60405180910390f35b348015620000fe57600080fd5b506200011d600480360381019062000117919062000aeb565b62000609565b6040516200012c919062000d2b565b60405180910390f35b3480156200014257600080fd5b506200016160048036038101906200015b919062000d4f565b620006d8565b005b6200018160048036038101906200017b919062000b58565b620007e1565b005b3480156200019057600080fd5b50620001af6004803603810190620001a9919062000aeb565b62000974565b005b348015620001be57600080fd5b50620001c962000a49565b604051620001d8919062000dbc565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000274576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200026b9062000e4f565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16638c45c8656040518163ffffffff1660e01b8152600401600060405180830381600087803b158015620002bd57600080fd5b505af1158015620002d2573d6000803e3d6000fd5b5050505050565b60606000604067ffffffffffffffff811115620002fb57620002fa62000e71565b5b6040519080825280601f01601f1916602001820160405280156200032e5781602001600182028036833780820191505090505b5090506000807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60001b90506000805b60208110156200042657600082848a16901b9050600060f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614620003f75780868681518110620003b757620003b662000ea0565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508480620003f39062000f08565b9550505b60088362000406919062000f55565b9250600884901c93505080806200041d9062000f08565b9150506200035e565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60001b91506000905060005b60208110156200051c57600082848916901b9050600060f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614620004ed5780868681518110620004ad57620004ac62000ea0565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508480620004e99062000f08565b9550505b600883620004fc919062000f55565b9250600884901c9350508080620005139062000f08565b91505062000454565b5060008367ffffffffffffffff8111156200053c576200053b62000e71565b5b6040519080825280601f01601f1916602001820160405280156200056f5781602001600182028036833780820191505090505b50905060005b84811015620005fa5785818151811062000594576200059362000ea0565b5b602001015160f81c60f81b828281518110620005b557620005b462000ea0565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080620005f19062000f08565b91505062000575565b50809550505050505092915050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015620006cc57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162000681575b50505050509050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146200076b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620007629062000e4f565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663836fbaca83836040518363ffffffff1660e01b8152600401620007a892919062000fa1565b600060405180830381600087803b158015620007c357600080fd5b505af1158015620007d8573d6000803e3d6000fd5b50505050505050565b67016345785d8a00003410156200082f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000826906200101e565b60405180910390fd5b60003433600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051620008649062000a73565b6200087192919062001040565b6040518091039082f09050801580156200088f573d6000803e3d6000fd5b5090507ff0ebf3718a236c7f981d1ac8a3a920707b028c465a25f4f590b39ba49b24a14b818484604051620008c7939291906200106d565b60405180910390a1600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000a05576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620009fc9062001120565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611292806200114383390190565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000ab38262000a86565b9050919050565b62000ac58162000aa6565b811462000ad157600080fd5b50565b60008135905062000ae58162000aba565b92915050565b60006020828403121562000b045762000b0362000a81565b5b600062000b148482850162000ad4565b91505092915050565b6000819050919050565b62000b328162000b1d565b811462000b3e57600080fd5b50565b60008135905062000b528162000b27565b92915050565b6000806040838503121562000b725762000b7162000a81565b5b600062000b828582860162000b41565b925050602062000b958582860162000b41565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b8381101562000bdb57808201518184015260208101905062000bbe565b60008484015250505050565b6000601f19601f8301169050919050565b600062000c058262000b9f565b62000c11818562000baa565b935062000c2381856020860162000bbb565b62000c2e8162000be7565b840191505092915050565b6000602082019050818103600083015262000c55818462000bf8565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b62000c948162000aa6565b82525050565b600062000ca8838362000c89565b60208301905092915050565b6000602082019050919050565b600062000cce8262000c5d565b62000cda818562000c68565b935062000ce78362000c79565b8060005b8381101562000d1e57815162000d02888262000c9a565b975062000d0f8362000cb4565b92505060018101905062000ceb565b5085935050505092915050565b6000602082019050818103600083015262000d47818462000cc1565b905092915050565b60008060006060848603121562000d6b5762000d6a62000a81565b5b600062000d7b8682870162000ad4565b935050602062000d8e8682870162000b41565b925050604062000da18682870162000b41565b9150509250925092565b62000db68162000aa6565b82525050565b600060208201905062000dd3600083018462000dab565b92915050565b7f4f6e6c79204272696467652063616e2063616c6c20746869732066756e63746960008201527f6f6e000000000000000000000000000000000000000000000000000000000000602082015250565b600062000e3760228362000baa565b915062000e448262000dd9565b604082019050919050565b6000602082019050818103600083015262000e6a8162000e28565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b600062000f158262000efe565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362000f4a5762000f4962000ecf565b5b600182019050919050565b600062000f628262000efe565b915062000f6f8362000efe565b925082820190508082111562000f8a5762000f8962000ecf565b5b92915050565b62000f9b8162000b1d565b82525050565b600060408201905062000fb8600083018562000f90565b62000fc7602083018462000f90565b9392505050565b7f4d696e696d756d20636f6c61746572616c20697320302e315446494c00000000600082015250565b600062001006601c8362000baa565b9150620010138262000fce565b602082019050919050565b60006020820190508181036000830152620010398162000ff7565b9050919050565b600060408201905062001057600083018562000dab565b62001066602083018462000dab565b9392505050565b600060608201905062001084600083018662000dab565b62001093602083018562000f90565b620010a2604083018462000f90565b949350505050565b7f4f6e6c79204f776e65722063616e2063616c6c20746869732066756e6374696f60008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b60006200110860218362000baa565b91506200111582620010aa565b604082019050919050565b600060208201905081810360008301526200113b81620010f9565b905091905056fe60806040526000600260146101000a81548160ff021916908360028111156200002d576200002c62000121565b5b021790555060405162001292380380620012928339818101604052810190620000579190620001ba565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000201565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001828262000155565b9050919050565b620001948162000175565b8114620001a057600080fd5b50565b600081519050620001b48162000189565b92915050565b60008060408385031215620001d457620001d362000150565b5b6000620001e485828601620001a3565b9250506020620001f785828601620001a3565b9150509250929050565b61108180620002116000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80634e69d56014610067578063836fbaca1461008557806384de1ea5146100a15780638c45c865146100bf578063de292789146100c9578063fb32c508146100e7575b600080fd5b61006f610105565b60405161007c9190610991565b60405180910390f35b61009f600480360381019061009a91906109f6565b61011c565b005b6100a96103b2565b6040516100b69190610a77565b60405180910390f35b6100c76103db565b005b6100d1610732565b6040516100de9190610b22565b60405180910390f35b6100ef6108f0565b6040516100fc9190610a77565b60405180910390f35b6000600260149054906101000a900460ff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806101c55750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610204576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fb90610b90565b60405180910390fd5b600060028111156102185761021761091a565b5b600260149054906101000a900460ff16600281111561023a5761023961091a565b5b1461027a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027190610bfc565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16476040516102c290610c4d565b60006040518083038185875af1925050503d80600081146102ff576040519150601f19603f3d011682016040523d82523d6000602084013e610304565b606091505b5050905080610348576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033f90610cae565b60405180910390fd5b6001600260146101000a81548160ff0219169083600281111561036e5761036d61091a565b5b021790555082600381905550816004819055507f7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b60405160405180910390a1505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806104845750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b6104c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104ba90610b90565b60405180910390fd5b600060028111156104d7576104d661091a565b5b600260149054906101000a900460ff1660028111156104f9576104f861091a565b5b14610539576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161053090610bfc565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1666b1a2bc2ec5000060405161058790610c4d565b60006040518083038185875af1925050503d80600081146105c4576040519150601f19603f3d011682016040523d82523d6000602084013e6105c9565b606091505b505090508061060d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161060490610d1a565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff164760405161065190610c4d565b60006040518083038185875af1925050503d806000811461068e576040519150601f19603f3d011682016040523d82523d6000602084013e610693565b606091505b505080915050806106d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106d090610dac565b60405180910390fd5b60028060146101000a81548160ff021916908360028111156106fe576106fd61091a565b5b02179055507ff6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e60405160405180910390a150565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146107c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016107b990610e18565b60405180910390fd5b600060028111156107d6576107d561091a565b5b600260149054906101000a900460ff1660028111156107f8576107f761091a565b5b03610838576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161082f90610e84565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff16631453d7566003546004546040518363ffffffff1660e01b81526004016108a0929190610eb3565b600060405180830381865afa1580156108bd573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906108e69190611002565b9050809250505090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6003811061095a5761095961091a565b5b50565b600081905061096b82610949565b919050565b600061097b8261095d565b9050919050565b61098b81610970565b82525050565b60006020820190506109a66000830184610982565b92915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6109d3816109c0565b81146109de57600080fd5b50565b6000813590506109f0816109ca565b92915050565b60008060408385031215610a0d57610a0c6109b6565b5b6000610a1b858286016109e1565b9250506020610a2c858286016109e1565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610a6182610a36565b9050919050565b610a7181610a56565b82525050565b6000602082019050610a8c6000830184610a68565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b83811015610acc578082015181840152602081019050610ab1565b60008484015250505050565b6000601f19601f8301169050919050565b6000610af482610a92565b610afe8185610a9d565b9350610b0e818560208601610aae565b610b1781610ad8565b840191505092915050565b60006020820190508181036000830152610b3c8184610ae9565b905092915050565b7f43616c6c6572206973206e6f7420627269646765000000000000000000000000600082015250565b6000610b7a601483610a9d565b9150610b8582610b44565b602082019050919050565b60006020820190508181036000830152610ba981610b6d565b9050919050565b7f4a6f62206973206e6f742070656e64696e670000000000000000000000000000600082015250565b6000610be6601283610a9d565b9150610bf182610bb0565b602082019050919050565b60006020820190508181036000830152610c1581610bd9565b9050919050565b600081905092915050565b50565b6000610c37600083610c1c565b9150610c4282610c27565b600082019050919050565b6000610c5882610c2a565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610c98601483610a9d565b9150610ca382610c62565b602082019050919050565b60006020820190508181036000830152610cc781610c8b565b9050919050565b7f4661696c656420746f2073656e6420457468657220746f206272696467650000600082015250565b6000610d04601e83610a9d565b9150610d0f82610cce565b602082019050919050565b60006020820190508181036000830152610d3381610cf7565b9050919050565b7f4661696c656420746f2073656e6420457468657220746f20696e69746961746f60008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b6000610d96602183610a9d565b9150610da182610d3a565b604082019050919050565b60006020820190508181036000830152610dc581610d89565b9050919050565b7f43616c6c6572206973206e6f7420696e69746961746f72000000000000000000600082015250565b6000610e02601783610a9d565b9150610e0d82610dcc565b602082019050919050565b60006020820190508181036000830152610e3181610df5565b9050919050565b7f4a6f62206973206e6f7420636f6d706c65746564000000000000000000000000600082015250565b6000610e6e601483610a9d565b9150610e7982610e38565b602082019050919050565b60006020820190508181036000830152610e9d81610e61565b9050919050565b610ead816109c0565b82525050565b6000604082019050610ec86000830185610ea4565b610ed56020830184610ea4565b9392505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610f1e82610ad8565b810181811067ffffffffffffffff82111715610f3d57610f3c610ee6565b5b80604052505050565b6000610f506109ac565b9050610f5c8282610f15565b919050565b600067ffffffffffffffff821115610f7c57610f7b610ee6565b5b610f8582610ad8565b9050602081019050919050565b6000610fa5610fa084610f61565b610f46565b905082815260208101848484011115610fc157610fc0610ee1565b5b610fcc848285610aae565b509392505050565b600082601f830112610fe957610fe8610edc565b5b8151610ff9848260208601610f92565b91505092915050565b600060208284031215611018576110176109b6565b5b600082015167ffffffffffffffff811115611036576110356109bb565b5b61104284828501610fd4565b9150509291505056fea26469706673582212208c68a6b9766ac716d1d319f5d0f6fd3e30628b0b9ea0079a24e6aea5521fb0e264736f6c63430008110033a2646970667358221220a92c1a9874046b0716b6b4970377ba0eb480de14af1a1168b38905792e74b54864736f6c63430008110033",
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

// Bytes64ToString is a free data retrieval call binding the contract method 0x1453d756.
//
// Solidity: function bytes64ToString(bytes32 x, bytes32 y) pure returns(string)
func (_Contracts *ContractsCaller) Bytes64ToString(opts *bind.CallOpts, x [32]byte, y [32]byte) (string, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "bytes64ToString", x, y)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Bytes64ToString is a free data retrieval call binding the contract method 0x1453d756.
//
// Solidity: function bytes64ToString(bytes32 x, bytes32 y) pure returns(string)
func (_Contracts *ContractsSession) Bytes64ToString(x [32]byte, y [32]byte) (string, error) {
	return _Contracts.Contract.Bytes64ToString(&_Contracts.CallOpts, x, y)
}

// Bytes64ToString is a free data retrieval call binding the contract method 0x1453d756.
//
// Solidity: function bytes64ToString(bytes32 x, bytes32 y) pure returns(string)
func (_Contracts *ContractsCallerSession) Bytes64ToString(x [32]byte, y [32]byte) (string, error) {
	return _Contracts.Contract.Bytes64ToString(&_Contracts.CallOpts, x, y)
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

// Request is a paid mutator transaction binding the contract method 0x51bffe84.
//
// Solidity: function request(bytes32 cid1, bytes32 cid2) payable returns()
func (_Contracts *ContractsTransactor) Request(opts *bind.TransactOpts, cid1 [32]byte, cid2 [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "request", cid1, cid2)
}

// Request is a paid mutator transaction binding the contract method 0x51bffe84.
//
// Solidity: function request(bytes32 cid1, bytes32 cid2) payable returns()
func (_Contracts *ContractsSession) Request(cid1 [32]byte, cid2 [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Request(&_Contracts.TransactOpts, cid1, cid2)
}

// Request is a paid mutator transaction binding the contract method 0x51bffe84.
//
// Solidity: function request(bytes32 cid1, bytes32 cid2) payable returns()
func (_Contracts *ContractsTransactorSession) Request(cid1 [32]byte, cid2 [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.Request(&_Contracts.TransactOpts, cid1, cid2)
}

// SaveResult is a paid mutator transaction binding the contract method 0x1d90e287.
//
// Solidity: function saveResult(address job, bytes32 result1, bytes32 result2) returns()
func (_Contracts *ContractsTransactor) SaveResult(opts *bind.TransactOpts, job common.Address, result1 [32]byte, result2 [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "saveResult", job, result1, result2)
}

// SaveResult is a paid mutator transaction binding the contract method 0x1d90e287.
//
// Solidity: function saveResult(address job, bytes32 result1, bytes32 result2) returns()
func (_Contracts *ContractsSession) SaveResult(job common.Address, result1 [32]byte, result2 [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SaveResult(&_Contracts.TransactOpts, job, result1, result2)
}

// SaveResult is a paid mutator transaction binding the contract method 0x1d90e287.
//
// Solidity: function saveResult(address job, bytes32 result1, bytes32 result2) returns()
func (_Contracts *ContractsTransactorSession) SaveResult(job common.Address, result1 [32]byte, result2 [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SaveResult(&_Contracts.TransactOpts, job, result1, result2)
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
	Cid1    [32]byte
	Cid2    [32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterJobExecutionRequest is a free log retrieval operation binding the contract event 0xf0ebf3718a236c7f981d1ac8a3a920707b028c465a25f4f590b39ba49b24a14b.
//
// Solidity: event JobExecutionRequest(address sapoJob, bytes32 cid1, bytes32 cid2)
func (_Contracts *ContractsFilterer) FilterJobExecutionRequest(opts *bind.FilterOpts) (*ContractsJobExecutionRequestIterator, error) {

	logs, sub, err := _Contracts.contract.FilterLogs(opts, "JobExecutionRequest")
	if err != nil {
		return nil, err
	}
	return &ContractsJobExecutionRequestIterator{contract: _Contracts.contract, event: "JobExecutionRequest", logs: logs, sub: sub}, nil
}

// WatchJobExecutionRequest is a free log subscription operation binding the contract event 0xf0ebf3718a236c7f981d1ac8a3a920707b028c465a25f4f590b39ba49b24a14b.
//
// Solidity: event JobExecutionRequest(address sapoJob, bytes32 cid1, bytes32 cid2)
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

// ParseJobExecutionRequest is a log parse operation binding the contract event 0xf0ebf3718a236c7f981d1ac8a3a920707b028c465a25f4f590b39ba49b24a14b.
//
// Solidity: event JobExecutionRequest(address sapoJob, bytes32 cid1, bytes32 cid2)
func (_Contracts *ContractsFilterer) ParseJobExecutionRequest(log types.Log) (*ContractsJobExecutionRequest, error) {
	event := new(ContractsJobExecutionRequest)
	if err := _Contracts.contract.UnpackLog(event, "JobExecutionRequest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
