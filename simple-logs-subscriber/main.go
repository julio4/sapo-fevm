package main

import (
	"context"
	"fmt"
	"log"
  "math/big"
	"strings"
  "time"

  "github.com/ethereum/go-ethereum"
  "github.com/ethereum/go-ethereum/accounts/abi"
  "github.com/ethereum/go-ethereum/common"
  "github.com/ethereum/go-ethereum/ethclient"

  events "bacalhau-bridge/events"
)

func getCurrentBlock(client *ethclient.Client) (*big.Int, error) {
    block, err := client.BlockNumber(context.Background())
    if err != nil {
      return nil, err
    }
    return big.NewInt(int64(block)), nil
}

func processLogs(client *ethclient.Client, contractAddress common.Address, lastBlock *big.Int) (*big.Int, error) {
    currentBlock, err := getCurrentBlock(client)
    if err != nil {
      return nil, err
    }

    // assert lastBlock < currentBlock
    if lastBlock.Cmp(currentBlock) >= 0 {
      fmt.Println("DEBUG:: skipping (" + lastBlock.String() + " < " + currentBlock.String() +")")
      return nil, nil
    }

    fmt.Println("Processing blocks #" + lastBlock.String() + " - #" + currentBlock.String())
    query := ethereum.FilterQuery{
      FromBlock: lastBlock,
      ToBlock:   currentBlock,
      Addresses: []common.Address{
        contractAddress,
      },
    }

    logs, err := client.FilterLogs(context.Background(), query)
    if err != nil {
      return nil, err
    }

    contractAbi, err := abi.JSON(strings.NewReader(string(events.EventsABI)))
    if err != nil {
      return nil, err
    }

    for _, vLog := range logs {
      fmt.Println("> Log found:")
      //fmt.Println(vLog.BlockHash.Hex())
      //fmt.Println(vLog.BlockNumber)
      //fmt.Println(vLog.TxHash.Hex())

      event := struct {
        Number *big.Int
      }{}
      err := contractAbi.UnpackIntoInterface(&event, "NumberEvent", vLog.Data)
      if err != nil {
        log.Fatal(err)
      }

      fmt.Println("> " + event.Number.String())
    }

    // update lastBlock to maximum processed block
    return currentBlock, nil
}

func main() {
  // set-up Client
  contract := "0x9F8865559f2b22F3883e162bEbe80F6069Dd9Dc9"
  contractAddress := common.HexToAddress(contract)
	client, err := ethclient.Dial("https://api.hyperspace.node.glif.io/rpc/v1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

  currentBlock, err := getCurrentBlock(client)
	if err != nil {
		log.Fatal(err)
	}

  fmt.Println("Listen for logs for actor " + contractAddress.String() + " at block #" + currentBlock.String())

  oneConstant := big.NewInt(1)
  lastBlock := new(big.Int)
  lastBlock.Sub(currentBlock, oneConstant)

  // poll and process logs in a loop
  ticker := time.NewTicker(10 * time.Second)
  for range ticker.C {
    lastProcessedBlock, err := processLogs(client, contractAddress, lastBlock)
    if err != nil {
      log.Fatal(err)
    }
    if lastProcessedBlock != nil {
      // ensure we don't process again this block
      lastBlock.Add(lastProcessedBlock, oneConstant)
      fmt.Println("DEBUG:: waiting for next block " + lastBlock.String())
    }
  }
}
