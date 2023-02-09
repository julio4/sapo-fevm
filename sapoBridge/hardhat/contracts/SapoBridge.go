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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sapoJob\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cid1\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"cid2\",\"type\":\"bytes32\"}],\"name\":\"JobExecutionRequest\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"x\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"y\",\"type\":\"bytes32\"}],\"name\":\"bytes64ToString\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"}],\"name\":\"failAndRefund\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBridgeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getJobs\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"cid1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid2\",\"type\":\"bytes32\"}],\"name\":\"request\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"job\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"result1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"result2\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid1\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"cid2\",\"type\":\"bytes32\"}],\"name\":\"saveResult\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"bridgeAddress\",\"type\":\"address\"}],\"name\":\"setBridge\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156200001157600080fd5b506040516200266d3803806200266d833981810160405281019062000037919062000129565b336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550506200015b565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620000f182620000c4565b9050919050565b6200010381620000e4565b81146200010f57600080fd5b50565b6000815190506200012381620000f8565b92915050565b600060208284031215620001425762000141620000bf565b5b6000620001528482850162000112565b91505092915050565b612502806200016b6000396000f3fe6080604052600436106200007a5760003560e01c806351bffe84116200005557806351bffe8414620001355780636b118e9e14620001555780638dd148021462000183578063fb32c50814620001b1576200007a565b80630ae523e5146200007f5780631453d75614620000ad5780631a3cbef414620000f1575b600080fd5b3480156200008c57600080fd5b50620000ab6004803603810190620000a5919062000af1565b620001e1565b005b348015620000ba57600080fd5b50620000d96004803603810190620000d3919062000b5e565b620002d9565b604051620000e8919062000c3f565b60405180910390f35b348015620000fe57600080fd5b506200011d600480360381019062000117919062000af1565b62000609565b6040516200012c919062000d31565b60405180910390f35b6200015360048036038101906200014d919062000b5e565b620006d8565b005b3480156200016257600080fd5b506200018160048036038101906200017b919062000d55565b6200086b565b005b3480156200019057600080fd5b50620001af6004803603810190620001a9919062000af1565b6200097a565b005b348015620001be57600080fd5b50620001c962000a4f565b604051620001d8919062000dee565b60405180910390f35b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000274576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200026b9062000e81565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16638c45c8656040518163ffffffff1660e01b8152600401600060405180830381600087803b158015620002bd57600080fd5b505af1158015620002d2573d6000803e3d6000fd5b5050505050565b60606000604067ffffffffffffffff811115620002fb57620002fa62000ea3565b5b6040519080825280601f01601f1916602001820160405280156200032e5781602001600182028036833780820191505090505b5090506000807fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60001b90506000805b60208110156200042657600082848a16901b9050600060f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614620003f75780868681518110620003b757620003b662000ed2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508480620003f39062000f3a565b9550505b60088362000406919062000f87565b9250600884901c93505080806200041d9062000f3a565b9150506200035e565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff60001b91506000905060005b60208110156200051c57600082848916901b9050600060f81b817effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614620004ed5780868681518110620004ad57620004ac62000ed2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508480620004e99062000f3a565b9550505b600883620004fc919062000f87565b9250600884901c9350508080620005139062000f3a565b91505062000454565b5060008367ffffffffffffffff8111156200053c576200053b62000ea3565b5b6040519080825280601f01601f1916602001820160405280156200056f5781602001600182028036833780820191505090505b50905060005b84811015620005fa5785818151811062000594576200059362000ed2565b5b602001015160f81c60f81b828281518110620005b557620005b462000ed2565b5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053508080620005f19062000f3a565b91505062000575565b50809550505050505092915050565b6060600260008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020805480602002602001604051908101604052809291908181526020018280548015620006cc57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831162000681575b50505050509050919050565b67016345785d8a000034101562000726576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016200071d9062001012565b60405180910390fd5b60003433600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166040516200075b9062000a79565b6200076892919062001034565b6040518091039082f090508015801562000786573d6000803e3d6000fd5b5090507ff0ebf3718a236c7f981d1ac8a3a920707b028c465a25f4f590b39ba49b24a14b818484604051620007be9392919062001072565b60405180910390a1600260003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819080600181540180825580915050600190039060005260206000200160009091909190916101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614620008fe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620008f59062000e81565b60405180910390fd5b8473ffffffffffffffffffffffffffffffffffffffff166333c303ad858585856040518563ffffffff1660e01b81526004016200093f9493929190620010af565b600060405180830381600087803b1580156200095a57600080fd5b505af11580156200096f573d6000803e3d6000fd5b505050505050505050565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161462000a0b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040162000a029062001172565b60405180910390fd5b80600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b611338806200119583390190565b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b600062000ab98262000a8c565b9050919050565b62000acb8162000aac565b811462000ad757600080fd5b50565b60008135905062000aeb8162000ac0565b92915050565b60006020828403121562000b0a5762000b0962000a87565b5b600062000b1a8482850162000ada565b91505092915050565b6000819050919050565b62000b388162000b23565b811462000b4457600080fd5b50565b60008135905062000b588162000b2d565b92915050565b6000806040838503121562000b785762000b7762000a87565b5b600062000b888582860162000b47565b925050602062000b9b8582860162000b47565b9150509250929050565b600081519050919050565b600082825260208201905092915050565b60005b8381101562000be157808201518184015260208101905062000bc4565b60008484015250505050565b6000601f19601f8301169050919050565b600062000c0b8262000ba5565b62000c17818562000bb0565b935062000c2981856020860162000bc1565b62000c348162000bed565b840191505092915050565b6000602082019050818103600083015262000c5b818462000bfe565b905092915050565b600081519050919050565b600082825260208201905092915050565b6000819050602082019050919050565b62000c9a8162000aac565b82525050565b600062000cae838362000c8f565b60208301905092915050565b6000602082019050919050565b600062000cd48262000c63565b62000ce0818562000c6e565b935062000ced8362000c7f565b8060005b8381101562000d2457815162000d08888262000ca0565b975062000d158362000cba565b92505060018101905062000cf1565b5085935050505092915050565b6000602082019050818103600083015262000d4d818462000cc7565b905092915050565b600080600080600060a0868803121562000d745762000d7362000a87565b5b600062000d848882890162000ada565b955050602062000d978882890162000b47565b945050604062000daa8882890162000b47565b935050606062000dbd8882890162000b47565b925050608062000dd08882890162000b47565b9150509295509295909350565b62000de88162000aac565b82525050565b600060208201905062000e05600083018462000ddd565b92915050565b7f4f6e6c79204272696467652063616e2063616c6c20746869732066756e63746960008201527f6f6e000000000000000000000000000000000000000000000000000000000000602082015250565b600062000e6960228362000bb0565b915062000e768262000e0b565b604082019050919050565b6000602082019050818103600083015262000e9c8162000e5a565b9050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000819050919050565b600062000f478262000f30565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff820362000f7c5762000f7b62000f01565b5b600182019050919050565b600062000f948262000f30565b915062000fa18362000f30565b925082820190508082111562000fbc5762000fbb62000f01565b5b92915050565b7f4d696e696d756d20636f6c61746572616c20697320302e315446494c00000000600082015250565b600062000ffa601c8362000bb0565b9150620010078262000fc2565b602082019050919050565b600060208201905081810360008301526200102d8162000feb565b9050919050565b60006040820190506200104b600083018562000ddd565b6200105a602083018462000ddd565b9392505050565b6200106c8162000b23565b82525050565b600060608201905062001089600083018662000ddd565b62001098602083018562001061565b620010a7604083018462001061565b949350505050565b6000608082019050620010c6600083018762001061565b620010d5602083018662001061565b620010e4604083018562001061565b620010f3606083018462001061565b95945050505050565b7f4f6e6c79204f776e65722063616e2063616c6c20746869732066756e6374696f60008201527f6e00000000000000000000000000000000000000000000000000000000000000602082015250565b60006200115a60218362000bb0565b91506200116782620010fc565b604082019050919050565b600060208201905081810360008301526200118d816200114b565b905091905056fe60806040526000600260146101000a81548160ff021916908360028111156200002d576200002c62000121565b5b021790555060405162001338380380620013388339818101604052810190620000579190620001ba565b816000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555033600260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505062000201565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b600080fd5b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000620001828262000155565b9050919050565b620001948162000175565b8114620001a057600080fd5b50565b600081519050620001b48162000189565b92915050565b60008060408385031215620001d457620001d362000150565b5b6000620001e485828601620001a3565b9250506020620001f785828601620001a3565b9150509250929050565b61112780620002116000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c806384de1ea51161005b57806384de1ea5146100da5780638c45c865146100f8578063de29278914610102578063fb32c508146101205761007d565b806310a0ebb01461008257806333c303ad146100a05780634e69d560146100bc575b600080fd5b61008a61013e565b6040516100979190610a95565b60405180910390f35b6100ba60048036038101906100b59190610b01565b61026e565b005b6100c4610514565b6040516100d19190610bdf565b60405180910390f35b6100e261052b565b6040516100ef9190610c3b565b60405180910390f35b610100610554565b005b61010a6108ab565b6040516101179190610a95565b60405180910390f35b6101286109db565b6040516101359190610c3b565b60405180910390f35b60606000600281111561015457610153610b68565b5b600260149054906101000a900460ff16600281111561017657610175610b68565b5b036101b6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101ad90610ca2565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff16631453d7566005546006546040518363ffffffff1660e01b815260040161021e929190610cd1565b600060405180830381865afa15801561023b573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906102649190610e20565b9050809250505090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806103175750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b610356576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161034d90610eb5565b60405180910390fd5b6000600281111561036a57610369610b68565b5b600260149054906101000a900460ff16600281111561038c5761038b610b68565b5b146103cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103c390610f21565b60405180910390fd5b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff164760405161041490610f72565b60006040518083038185875af1925050503d8060008114610451576040519150601f19603f3d011682016040523d82523d6000602084013e610456565b606091505b505090508061049a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161049190610fd3565b60405180910390fd5b6001600260146101000a81548160ff021916908360028111156104c0576104bf610b68565b5b0217905550846003819055508360048190555082600581905550816006819055507f7d6e73dea3a1d0a15642cc83127e32ad9e9fffb137c1567d352780788aa54e2b60405160405180910390a15050505050565b6000600260149054906101000a900460ff16905090565b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806105fd5750600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b61063c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161063390610eb5565b60405180910390fd5b600060028111156106505761064f610b68565b5b600260149054906101000a900460ff16600281111561067257610671610b68565b5b146106b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016106a990610f21565b60405180910390fd5b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1666b1a2bc2ec5000060405161070090610f72565b60006040518083038185875af1925050503d806000811461073d576040519150601f19603f3d011682016040523d82523d6000602084013e610742565b606091505b5050905080610786576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161077d9061103f565b60405180910390fd5b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16476040516107ca90610f72565b60006040518083038185875af1925050503d8060008114610807576040519150601f19603f3d011682016040523d82523d6000602084013e61080c565b606091505b50508091505080610852576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610849906110d1565b60405180910390fd5b60028060146101000a81548160ff0219169083600281111561087757610876610b68565b5b02179055507ff6da0b6aaf5c5b2aa1b43736ef16a25b4809e181c3e882b5e55d00b5bd6ac30e60405160405180910390a150565b6060600060028111156108c1576108c0610b68565b5b600260149054906101000a900460ff1660028111156108e3576108e2610b68565b5b03610923576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161091a90610ca2565b60405180910390fd5b6000600260009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff16631453d7566003546004546040518363ffffffff1660e01b815260040161098b929190610cd1565b600060405180830381865afa1580156109a8573d6000803e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906109d19190610e20565b9050809250505090565b6000600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905090565b600081519050919050565b600082825260208201905092915050565b60005b83811015610a3f578082015181840152602081019050610a24565b60008484015250505050565b6000601f19601f8301169050919050565b6000610a6782610a05565b610a718185610a10565b9350610a81818560208601610a21565b610a8a81610a4b565b840191505092915050565b60006020820190508181036000830152610aaf8184610a5c565b905092915050565b6000604051905090565b600080fd5b600080fd5b6000819050919050565b610ade81610acb565b8114610ae957600080fd5b50565b600081359050610afb81610ad5565b92915050565b60008060008060808587031215610b1b57610b1a610ac1565b5b6000610b2987828801610aec565b9450506020610b3a87828801610aec565b9350506040610b4b87828801610aec565b9250506060610b5c87828801610aec565b91505092959194509250565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b60038110610ba857610ba7610b68565b5b50565b6000819050610bb982610b97565b919050565b6000610bc982610bab565b9050919050565b610bd981610bbe565b82525050565b6000602082019050610bf46000830184610bd0565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000610c2582610bfa565b9050919050565b610c3581610c1a565b82525050565b6000602082019050610c506000830184610c2c565b92915050565b7f4a6f62206973206e6f7420636f6d706c65746564000000000000000000000000600082015250565b6000610c8c601483610a10565b9150610c9782610c56565b602082019050919050565b60006020820190508181036000830152610cbb81610c7f565b9050919050565b610ccb81610acb565b82525050565b6000604082019050610ce66000830185610cc2565b610cf36020830184610cc2565b9392505050565b600080fd5b600080fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b610d3c82610a4b565b810181811067ffffffffffffffff82111715610d5b57610d5a610d04565b5b80604052505050565b6000610d6e610ab7565b9050610d7a8282610d33565b919050565b600067ffffffffffffffff821115610d9a57610d99610d04565b5b610da382610a4b565b9050602081019050919050565b6000610dc3610dbe84610d7f565b610d64565b905082815260208101848484011115610ddf57610dde610cff565b5b610dea848285610a21565b509392505050565b600082601f830112610e0757610e06610cfa565b5b8151610e17848260208601610db0565b91505092915050565b600060208284031215610e3657610e35610ac1565b5b600082015167ffffffffffffffff811115610e5457610e53610ac6565b5b610e6084828501610df2565b91505092915050565b7f43616c6c6572206973206e6f7420627269646765000000000000000000000000600082015250565b6000610e9f601483610a10565b9150610eaa82610e69565b602082019050919050565b60006020820190508181036000830152610ece81610e92565b9050919050565b7f4a6f62206973206e6f742070656e64696e670000000000000000000000000000600082015250565b6000610f0b601283610a10565b9150610f1682610ed5565b602082019050919050565b60006020820190508181036000830152610f3a81610efe565b9050919050565b600081905092915050565b50565b6000610f5c600083610f41565b9150610f6782610f4c565b600082019050919050565b6000610f7d82610f4f565b9150819050919050565b7f4661696c656420746f2073656e64204574686572000000000000000000000000600082015250565b6000610fbd601483610a10565b9150610fc882610f87565b602082019050919050565b60006020820190508181036000830152610fec81610fb0565b9050919050565b7f4661696c656420746f2073656e6420457468657220746f206272696467650000600082015250565b6000611029601e83610a10565b915061103482610ff3565b602082019050919050565b600060208201905081810360008301526110588161101c565b9050919050565b7f4661696c656420746f2073656e6420457468657220746f20696e69746961746f60008201527f7200000000000000000000000000000000000000000000000000000000000000602082015250565b60006110bb602183610a10565b91506110c68261105f565b604082019050919050565b600060208201905081810360008301526110ea816110ae565b905091905056fea2646970667358221220196167c624193ca8d7523208441f6a2fcf72556f1d0f1c0ab34c62e8055e265364736f6c63430008110033a2646970667358221220d439a3af3c4e2fb6bbccc806f6792983653871061c3643eb8cf77768798b7bba64736f6c63430008110033",
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

// SaveResult is a paid mutator transaction binding the contract method 0x6b118e9e.
//
// Solidity: function saveResult(address job, bytes32 result1, bytes32 result2, bytes32 cid1, bytes32 cid2) returns()
func (_Contracts *ContractsTransactor) SaveResult(opts *bind.TransactOpts, job common.Address, result1 [32]byte, result2 [32]byte, cid1 [32]byte, cid2 [32]byte) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "saveResult", job, result1, result2, cid1, cid2)
}

// SaveResult is a paid mutator transaction binding the contract method 0x6b118e9e.
//
// Solidity: function saveResult(address job, bytes32 result1, bytes32 result2, bytes32 cid1, bytes32 cid2) returns()
func (_Contracts *ContractsSession) SaveResult(job common.Address, result1 [32]byte, result2 [32]byte, cid1 [32]byte, cid2 [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SaveResult(&_Contracts.TransactOpts, job, result1, result2, cid1, cid2)
}

// SaveResult is a paid mutator transaction binding the contract method 0x6b118e9e.
//
// Solidity: function saveResult(address job, bytes32 result1, bytes32 result2, bytes32 cid1, bytes32 cid2) returns()
func (_Contracts *ContractsTransactorSession) SaveResult(job common.Address, result1 [32]byte, result2 [32]byte, cid1 [32]byte, cid2 [32]byte) (*types.Transaction, error) {
	return _Contracts.Contract.SaveResult(&_Contracts.TransactOpts, job, result1, result2, cid1, cid2)
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
