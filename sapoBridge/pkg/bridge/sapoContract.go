package bridge

import (
	"context"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"

	events "sapoBridge/pkg/contract"
)

func getCurrentBlock(client *ethclient.Client) (*big.Int, error) {
	block, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}
	return big.NewInt(int64(block)), nil
}

func processLogs(ctx context.Context, client *ethclient.Client, contractAddress common.Address, lastBlock *big.Int) (*big.Int, error) {
	currentBlock, err := getCurrentBlock(client)
	if err != nil {
		return nil, err
	}

	// assert lastBlock < currentBlock
	if lastBlock.Cmp(currentBlock) >= 0 {
		log.Ctx(ctx).Debug().Msg("Skipping processing blocks")
		return nil, nil
	}

	log.Ctx(ctx).Info().Int64("from", lastBlock.Int64()).Int64("to", currentBlock.Int64()).Msg("Processing blocks")
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
		//(vLog.BlockHash.Hex())
		//(vLog.BlockNumber)
		//(vLog.TxHash.Hex())

		event := struct {
			Number *big.Int
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "NumberEvent", vLog.Data)
		if err != nil {
			log.Fatal()
		}

		log.Ctx(ctx).Info().Stringer("number", event.Number).Msg("NumberEvent")
	}

	// update lastBlock to maximum processed block
	return currentBlock, nil
}

type SapoContract struct {
	address common.Address
	client  *ethclient.Client
}

func NewSapoContract(ctx context.Context, address string) *SapoContract {
	contract := new(SapoContract)
	contract.address = common.HexToAddress(address)
	client, err := ethclient.Dial("https://api.hyperspace.node.glif.io/rpc/v1")
	if err != nil {
		log.Ctx(ctx).Fatal().Err(err).Msg("Failed to connect to Gliph node")
	}
	contract.client = client

	return contract
}

func (c SapoContract) Listen(ctx context.Context, out chan<- ContractSubmittedEvent) error {
	log.Ctx(ctx).Debug().Msg("Listening")
	defer log.Ctx(ctx).Debug().Msg("Stopping listening")

	sched := gocron.NewScheduler(time.UTC)
	_, err := sched.WaitForSchedule().Every(10).Second().Do(func() {
		e := exampleEvent()
		log.Ctx(ctx).Info().Stringer("id", e.OrderId()).Msg("New order")
		out <- e
	})
	if err != nil {
		return err
	}

	sched.StartAsync()
	defer sched.Stop()
	<-ctx.Done()

	//currentBlock, err := getCurrentBlock(c.client)
	//if err != nil {
	//	log.Ctx(ctx).Fatal().Err(err).Msg("Failed to get current block")
	//}

	//oneConstant := big.NewInt(1)
	//lastBlock := new(big.Int)
	//lastBlock.Sub(currentBlock, oneConstant)

	//// poll and process logs in a loop
	//ticker := time.NewTicker(10 * time.Second)
	//for range ticker.C {
	//	lastProcessedBlock, err := processLogs(ctx, c.client, c.address, lastBlock)
	//	if err != nil {
	//		log.Ctx(ctx).Fatal().Err(err).Msg("Failed to process logs")
	//	}
	//	if lastProcessedBlock != nil {
	//		// ensure we don't process again this block
	//		lastBlock.Add(lastProcessedBlock, oneConstant)
	//		log.Ctx(ctx).Debug().Int64("block", lastBlock.Int64()).Msg("Waiting for next block")
	//	}
	//}

	return nil
}

// Complete implements SmartContract
func (c SapoContract) Complete(ctx context.Context, e BacalhauJobCompletedEvent) (ContractPaidEvent, error) {
	log.Ctx(ctx).Info().Stringer("id", e.OrderId()).Msg("Complete")

	return e.Paid(), nil
}

// Refund implements SmartContract
func (c SapoContract) Refund(ctx context.Context, e ContractFailedEvent) (ContractRefundedEvent, error) {
	log.Ctx(ctx).Info().Stringer("id", e.OrderId()).Msg("Refunded")

	return e.Refunded(), nil
}
