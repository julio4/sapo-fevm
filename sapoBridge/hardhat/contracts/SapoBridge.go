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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sapoJob\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"JobExecutionRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"bytes64ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"}],\"name\":\"failAndRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"cid\",\"type\":\"string\"}],\"name\":\"request\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"result1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"result2\",\"type\":\"bytes32\"}],\"name\":\"saveResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200258138038062002581833981810160405281019062000037919062000129565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200015b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000f182620000c4565b9050919050565b6200010381620000e4565b81146200010f57600080fd5b50565b6000815190506200012381620000f8565b92915050565b600060208284031215620001425762000141620000bf565b5b6000620001528482850162000112565b91505092915050565b612416806200016b6000396000f3fe6080604052600436106200007a5760003560e01c80631d90e28711620000555780631d90e28714620001355780632c19988914620001635780638dd148021462000183578063fb32c50814620001b1576200007a565b80630ae523e5146200007f5780631453d75614620000ad5780631a3cbef414620000f1575b600080fd5b3480156200008c57600080fd5b50620000ab6004803603810190620000a5919062000b5f565b620001e1565b005b348015620000ba57600080fd5b50620000d96004803603810190620000d3919062000bcc565b620002d9565b604051620000e8919062000cad565b60405180910390f35b348015620000fe57600080fd5b506200011d600480360381019062000117919062000b5f565b62000609565b6040516200012c919062000d9f565b60405180910390f35b3480156200014257600080fd5b506200016160048036038101906200015b919062000dc3565b620006d8565b005b6200018160048036038101906200017b919062000f70565b620007e1565b005b3480156200019057600080fd5b50620001af6004803603810190620001a9919062000b5f565b620009d9565b005b348015620001be57600080fd5b50620001c962000aae565b604051620001d8919062000fd2565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000274576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200026b9062001065565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16638c45c8656040518163ffffffff1660e01b8152600401600060405180830381600087803b158015620002bd57600080fd5b505af1158015620002d2573d6000803e3d6000fd5b5050505050565b60606000604067ffffffffffffffff811115620002fb57620002fa62000e29565b5b6040519080825280601f01601f1916602001820160405280156200032e5781602001600182028036833780820191505090505b5090506000807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60001b90506000805b60208110156200042657600082848a16901b9050600060f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614620003f75780868681518110620003b757620003b662001087565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508480620003f390620010ef565b9550505b6008836200040691906200113c565b9250600884901c93505080806200041d90620010ef565b9150506200035e565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60001b91506000905060005b60208110156200051c57600082848916901b9050600060f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614620004ed5780868681518110620004ad57620004ac62001087565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508480620004e990620010ef565b9550505b600883620004fc91906200113c565b9250600884901c93505080806200051390620010ef565b91505062000454565b5060008367ffffffffffffffff8111156200053c576200053b62000e29565b5b6040519080825280601f01601f1916602001820160405280156200056f5781602001600182028036833780820191505090505b50905060005b84811015620005fa5785818151811062000594576200059362001087565b5b602001015160f81c60f81b828281518110620005b557620005b462001087565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080620005f190620010ef565b91505062000575565b50809550505050505092915050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015620006cc57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162000681575b50505050509050919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146200076b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620007629062001065565b60405180910390fd5b8273ffffffffffffffffffffffffffffffffffffffff1663836fbaca83836040518363ffffffff1660e01b8152600401620007a892919062001188565b600060405180830381600087803b158015620007c357600080fd5b505af1158015620007d8573d6000803e3d6000fd5b50505050505050565b67016345785d8a00003410156200082f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620008269062001205565b60405180910390fd5b600033600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16604051620008639062000ad8565b6200087092919062001227565b604051809103906000f0801580156200088d573d6000803e3d6000fd5b5090507f7843f28b8cbd00b16af4f1efc48e12641ca994dc385bff44528582ad78f240908183604051620008c392919062001254565b60405180910390a1600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015620009d4573d6000803e3d6000fd5b505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000a6a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000a6190620012fe565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b6110c0806200132183390190565b6000604051905090565b600080fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000b278262000afa565b9050919050565b62000b398162000b1a565b811462000b4557600080fd5b50565b60008135905062000b598162000b2e565b92915050565b60006020828403121562000b785762000b7762000af0565b5b600062000b888482850162000b48565b91505092915050565b6000819050919050565b62000ba68162000b91565b811462000bb257600080fd5b50565b60008135905062000bc68162000b9b565b92915050565b6000806040838503121562000be65762000be562000af0565b5b600062000bf68582860162000bb5565b925050602062000c098582860162000bb5565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b8381101562000c4f57808201518184015260208101905062000c32565b60008484015250505050565b6000601f19601f8301169050919050565b600062000c798262000c13565b62000c85818562000c1e565b935062000c9781856020860162000c2f565b62000ca28162000c5b565b840191505092915050565b6000602082019050818103600083015262000cc9818462000c6c565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b62000d088162000b1a565b82525050565b600062000d1c838362000cfd565b60208301905092915050565b6000602082019050919050565b600062000d428262000cd1565b62000d4e818562000cdc565b935062000d5b8362000ced565b8060005b8381101562000d9257815162000d76888262000d0e565b975062000d838362000d28565b92505060018101905062000d5f565b5085935050505092915050565b6000602082019050818103600083015262000dbb818462000d35565b905092915050565b60008060006060848603121562000ddf5762000dde62000af0565b5b600062000def8682870162000b48565b935050602062000e028682870162000bb5565b925050604062000e158682870162000bb5565b9150509250925092565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b62000e638262000c5b565b810181811067ffffffffffffffff8211171562000e855762000e8462000e29565b5b80604052505050565b600062000e9a62000ae6565b905062000ea8828262000e58565b919050565b600067ffffffffffffffff82111562000ecb5762000eca62000e29565b5b62000ed68262000c5b565b9050602081019050919050565b82818337600083830152505050565b600062000f0962000f038462000ead565b62000e8e565b90508281526020810184848401111562000f285762000f2762000e24565b5b62000f3584828562000ee3565b509392505050565b600082601f83011262000f555762000f5462000e1f565b5b813562000f6784826020860162000ef2565b91505092915050565b60006020828403121562000f895762000f8862000af0565b5b600082013567ffffffffffffffff81111562000faa5762000fa962000af5565b5b62000fb88482850162000f3d565b91505092915050565b62000fcc8162000b1a565b82525050565b600060208201905062000fe9600083018462000fc1565b92915050565b7f4f6e6c79204272696467652063616e2063616c6c20746869732066756e63746960008201527f6f6e000000000000000000000000000000000000000000000000000000000000602082015250565b60006200104d60228362000c1e565b91506200105a8262000fef565b604082019050919050565b6000602082019050818103600083015262001080816200103e565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b6000620010fc82620010e5565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620011315762001130620010b6565b5b600182019050919050565b60006200114982620010e5565b91506200115683620010e5565b9250828201905080821115620011715762001170620010b6565b5b92915050565b620011828162000b91565b82525050565b60006040820190506200119f600083018562001177565b620011ae602083018462001177565b9392505050565b7f4d696e696d756d20636f6c61746572616c20697320302e315446494c00000000600082015250565b6000620011ed601c8362000c1e565b9150620011fa82620011b5565b602082019050919050565b600060208201905081810360008301526200122081620011de565b9050919050565b60006040820190506200123e600083018562000fc1565b6200124d602083018462000fc1565b9392505050565b60006040820190506200126b600083018562000fc1565b81810360208301526200127f818462000c6c565b90509392505050565b7f4f6e6c79204f776e65722063616e2063616c6c20746869732066756e6374696f60008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b6000620012e660218362000c1e565b9150620012f38262001288565b604082019050919050565b600060208201905081810360008301526200131981620012d7565b905091905056fe60806040526000600260146101000a81548160ff021916908360028111156200002d576200002c62000121565b5b0217905550604051620010c0380380620010c08339818101604052810190620000579190620001ba565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000201565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001828262000155565b9050919050565b620001948162000175565b8114620001a057600080fd5b50565b600081519050620001b48162000189565b92915050565b60008060408385031215620001d457620001d362000150565b5b6000620001e485828601620001a3565b9250506020620001f785828601620001a3565b9150509250929050565b610eaf80620002116000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80634e69d56014610067578063836fbaca1461008557806384de1ea5146100a15780638c45c865146100bf578063de292789146100c9578063fb32c508146100e7575b600080fd5b61006f610105565b60405161007c91906108bd565b60405180910390f35b61009f600480360381019061009a9190610922565b61011c565b005b6100a96103b1565b6040516100b691906109a3565b60405180910390f35b6100c76103da565b005b6100d161065e565b6040516100de9190610a4e565b60405180910390f35b6100ef61081c565b6040516100fc91906109a3565b60405180910390f35b6000600260149054906101000a900460ff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806101c55750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610204576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101fb90610abc565b60405180910390fd5b6000600281111561021857610217610846565b5b600260149054906101000a900460ff16600281111561023a57610239610846565b5b1461027a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161027190610b28565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16476040516102c190610b79565b60006040518083038185875af1925050503d80600081146102fe576040519150601f19603f3d011682016040523d82523d6000602084013e610303565b606091505b5050905080610347576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161033e90610bda565b60405180910390fd5b6001600260146101000a81548160ff0219169083600281111561036d5761036c610846565b5b021790555082600381905550816004819055507f7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b60405160405180910390a1505050565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806104835750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b6104c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104b990610abc565b60405180910390fd5b600060028111156104d6576104d5610846565b5b600260149054906101000a900460ff1660028111156104f8576104f7610846565b5b14610538576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161052f90610b28565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff164760405161057f90610b79565b60006040518083038185875af1925050503d80600081146105bc576040519150601f19603f3d011682016040523d82523d6000602084013e6105c1565b606091505b5050905080610605576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016105fc90610bda565b60405180910390fd5b60028060146101000a81548160ff0219169083600281111561062a57610629610846565b5b02179055507ff6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e60405160405180910390a150565b606060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16146106ee576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106e590610c46565b60405180910390fd5b6000600281111561070257610701610846565b5b600260149054906101000a900460ff16600281111561072457610723610846565b5b03610764576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161075b90610cb2565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff16631453d7566003546004546040518363ffffffff1660e01b81526004016107cc929190610ce1565b600060405180830381865afa1580156107e9573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906108129190610e30565b9050809250505090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b6003811061088657610885610846565b5b50565b600081905061089782610875565b919050565b60006108a782610889565b9050919050565b6108b78161089c565b82525050565b60006020820190506108d260008301846108ae565b92915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b6108ff816108ec565b811461090a57600080fd5b50565b60008135905061091c816108f6565b92915050565b60008060408385031215610939576109386108e2565b5b60006109478582860161090d565b92505060206109588582860161090d565b9150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600061098d82610962565b9050919050565b61099d81610982565b82525050565b60006020820190506109b86000830184610994565b92915050565b600081519050919050565b600082825260208201905092915050565b60005b838110156109f85780820151818401526020810190506109dd565b60008484015250505050565b6000601f19601f8301169050919050565b6000610a20826109be565b610a2a81856109c9565b9350610a3a8185602086016109da565b610a4381610a04565b840191505092915050565b60006020820190508181036000830152610a688184610a15565b905092915050565b7f43616c6c6572206973206e6f7420627269646765000000000000000000000000600082015250565b6000610aa66014836109c9565b9150610ab182610a70565b602082019050919050565b60006020820190508181036000830152610ad581610a99565b9050919050565b7f4a6f62206973206e6f742070656e64696e670000000000000000000000000000600082015250565b6000610b126012836109c9565b9150610b1d82610adc565b602082019050919050565b60006020820190508181036000830152610b4181610b05565b9050919050565b600081905092915050565b50565b6000610b63600083610b48565b9150610b6e82610b53565b600082019050919050565b6000610b8482610b56565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610bc46014836109c9565b9150610bcf82610b8e565b602082019050919050565b60006020820190508181036000830152610bf381610bb7565b9050919050565b7f43616c6c6572206973206e6f7420696e69746961746f72000000000000000000600082015250565b6000610c306017836109c9565b9150610c3b82610bfa565b602082019050919050565b60006020820190508181036000830152610c5f81610c23565b9050919050565b7f4a6f62206973206e6f7420636f6d706c65746564000000000000000000000000600082015250565b6000610c9c6014836109c9565b9150610ca782610c66565b602082019050919050565b60006020820190508181036000830152610ccb81610c8f565b9050919050565b610cdb816108ec565b82525050565b6000604082019050610cf66000830185610cd2565b610d036020830184610cd2565b9392505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d4c82610a04565b810181811067ffffffffffffffff82111715610d6b57610d6a610d14565b5b80604052505050565b6000610d7e6108d8565b9050610d8a8282610d43565b919050565b600067ffffffffffffffff821115610daa57610da9610d14565b5b610db382610a04565b9050602081019050919050565b6000610dd3610dce84610d8f565b610d74565b905082815260208101848484011115610def57610dee610d0f565b5b610dfa8482856109da565b509392505050565b600082601f830112610e1757610e16610d0a565b5b8151610e27848260208601610dc0565b91505092915050565b600060208284031215610e4657610e456108e2565b5b600082015167ffffffffffffffff811115610e6457610e636108e7565b5b610e7084828501610e02565b9150509291505056fea2646970667358221220f9e080f7b4930658cf3405771fcbd9eff6e5f0cd065755560a342f2da2dc133864736f6c63430008110033a26469706673582212201e173ef6eb3d4c93bfeaead73b7effa674ed936f1908615b76405627ffb30da364736f6c63430008110033",
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
