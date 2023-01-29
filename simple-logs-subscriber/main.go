package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/contracts/token/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://api.hyperspace.node.glif.io/rpc/v1")
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	contractAddress := common.HexToAddress("0x773d8856dd7F78857490e5Eea65111D8d466A646")
	query := fmt.Sprintf("0x%s", abi.TokenABI)
	contractAbi, err := client.EthClient.ABIDecoder(query)
	if err != nil {
		log.Fatal(err)
	}

	logs := make(chan types.Log)
	sub, err := client.EthClient.SubscribeFilterLogs(context.Background(), types.FilterQuery{
		Addresses: []common.Address{contractAddress},
		Topics: [][]common.Hash{
			{common.HexToHash("0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef")},
		},
	}, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			event, err := contractAbi.Event(&vLog)
			if err != nil {
				log.Println("error decoding event:", err)
				continue
			}

			if event.Name == "Transfer" {
				var transferEvent abi.TokenTransfer
				if err := contractAbi.Unpack(&transferEvent, "Transfer", vLog.Data); err != nil {
					log.Println("error unpacking event data:", err)
					continue
				}

				fromAddress := transferEvent.From.Hex()
				toAddress := transferEvent.To.Hex()
				value := transferEvent.Value.String()

				fmt.Println("Event: Transfer")
				fmt.Println("From: ", strings.ToLower(fromAddress))
				fmt.Println("To: ", strings.ToLower(toAddress))
				fmt.Println("Value: ", value)
			}
		}
	}
}
