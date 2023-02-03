package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"sapoBridge/hardhat/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// generate creates a new wallet and prints the address to stdout
// The wallet is stored in the ./wallets directory
// Be aware that this is only for experimental purposes
func generate() {
	ks := keystore.NewKeyStore("./", keystore.StandardScryptN, keystore.StandardScryptP)
	password := "croaked!"
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account.Address.Hex())
}

// test_call is a test function to call the smart contract from the bridge
func test_call() {
	wallet_file := "wallet file here"
	addr := common.HexToAddress("bridge contract address here")

	client, err := ethclient.Dial("wss://ws-filecoin-hyperspace.chainstacklabs.com/rpc/v0")
	if err != nil {
		log.Fatal(err)
	}

	contract, err := contracts.NewContracts(addr, client)
	if err != nil {
		log.Fatal(err)
	}

	jsonBytes, err := ioutil.ReadFile(wallet_file)
	if err != nil {
		log.Fatal(err)
	}

	password := "croaked!"
	key, err := keystore.DecryptKey(jsonBytes, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(key.Address)

	reader, err := os.Open(wallet_file)
	if err != nil {
		log.Fatal(err)
	}

	opts, err := bind.NewTransactorWithChainID(reader, password, big.NewInt(3141))
	if err != nil {
		log.Fatal(err)
	}

	tx, err := contract.Request(opts, "cid here")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx: %v\n", tx)
}

func main() {
	// test_call()
}
