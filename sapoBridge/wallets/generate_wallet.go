package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"
	"sapoBridge/hardhat/contracts"
	"sapoBridge/pkg/bridge"

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
	log.Print("Instantiating environment")

	wallet_file := "./bridgeaccount"
	addr := common.HexToAddress("0x21d4659bc8b766CEA18Fa47206f9BA6Df091de68")

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

	log.Print("Preparing transaction")

	jobAddr := common.HexToAddress("0x49e6dce29e388e2d8944e644d63754cF4A5cCEea")
	jobId := "aaaabaaacaaadaaaeaaafaaagaaahaaaiaaajaaakaaal"
	cid1, err1 := bridge.Pack(jobId[:32])
	cid2, err2 := bridge.Pack(jobId[32:])

	if err1 != nil || err2 != nil {
		log.Print(err1)
		log.Fatal(err2)
	}

	log.Print("Launching transaction")
	log.Print("jobAddr ", jobAddr)
	log.Print("cid1 ", cid1)
	log.Print("cid2 ", cid2)

	tx, err := contract.SaveResult(opts, jobAddr, cid1, cid2, cid1, cid2)

	if err != nil {
		log.Print("Error has occurred during transaction ", tx)
		log.Fatal(err)
	}

	fmt.Printf("tx: %v\n", tx)
}

func main() {
	test_call()
}
