package bridge

import (
	"context"
	"crypto/rand"
	"encoding/json"
	"math/big"
	"os"
	"time"

	"sapoBridge/hardhat/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/filecoin-project/bacalhau/pkg/requester/publicapi"
	"github.com/go-co-op/gocron"
	"github.com/rs/zerolog/log"
)

type SmartContract interface {
	Listen(context.Context, chan<- ContractSubmittedEvent) error

	Complete(context.Context, BacalhauJobCompletedEvent, publicapi.RequesterAPIClient) (ContractPaidEvent, error)

	Refund(context.Context, ContractFailedEvent) (ContractRefundedEvent, error)
}

type RealContract struct {
	contract *contracts.Contracts
	transact *bind.TransactOpts

	maxSeenBlock uint64
}

// Complete implements SmartContract
func (r *RealContract) Complete(ctx context.Context, event BacalhauJobCompletedEvent, client publicapi.RequesterAPIClient) (ContractPaidEvent, error) {
	timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	jobs, err := client.GetResults(timeoutCtx, "2cb7f3ed-efa3-4cff-a8e9-621c14abef07")

	if err != nil {
		log.Ctx(ctx).Error().Err(err).Send()
		return event, err
	}

	outCID := jobs[0].Data.CID

	jobId1, err1 := Pack(event.JobID()[:32])
	jobId2, err2 := Pack(event.JobID()[32:])

	CID1, err3 := Pack(outCID[:32])
	CID2, err4 := Pack(outCID[32:])

	localCtx := log.Ctx(ctx).With().
		Str("jobAddress", event.Addr().Hex()).
		Bytes("jobId1", jobId1[:]).
		Bytes("jobId2", jobId2[:]).
		Bytes("CID1", CID1[:]).
		Bytes("CID2", CID2[:]).
		Logger().WithContext(ctx)

	if err1 != nil || err2 != nil || err3 != nil || err4 != nil {
		log.Ctx(localCtx).Debug().
			Err(err1).Err(err2).Err(err3).Err(err4).
			Msg("Result saving of completed job has failed")
		return event, context.Canceled
	}

	tx, err := r.contract.SaveResult(r.transact, event.Addr(), jobId1, jobId2, CID1, CID2)

	if err != nil {
		log.Ctx(localCtx).Err(err).Msg("Error in saveResult transaction")
		return event, context.Canceled
	}

	log.Ctx(localCtx).Debug().
		Stringer("transactionId", tx.Hash()).
		Msg("Successfully completed saveResult transaction")

	return event.Paid(), nil
}

// Listen implements SmartContract
func (r *RealContract) Listen(ctx context.Context, out chan<- ContractSubmittedEvent) error {
	scheduler := gocron.NewScheduler(time.UTC)
	_, err := scheduler.Every(15*time.Second).SingletonMode().Do(r.ReadLogs, ctx, out)
	if err != nil {
		return err
	}

	scheduler.StartAsync()
	defer scheduler.Stop()

	<-ctx.Done()
	return nil
}

func (r *RealContract) ReadLogs(ctx context.Context, out chan<- ContractSubmittedEvent) {
	log.Ctx(ctx).Debug().Uint64("fromBlock", r.maxSeenBlock+1).Msg("Polling for smart contract events")

	opts := bind.FilterOpts{Start: uint64(r.maxSeenBlock + 1), Context: ctx}
	logs, err := r.contract.ContractsFilterer.FilterJobExecutionRequest(&opts)
	if err != nil {
		log.Ctx(ctx).Error().Err(err).Send()
		return
	}
	defer logs.Close()

	for logs.Next() {
		recvEvent := logs.Event

		cid := Unpack(recvEvent.Cid1) + Unpack(recvEvent.Cid2)

		log.Ctx(ctx).Debug().
			Stringer("txn", recvEvent.Raw.TxHash).
			Uint64("block#", recvEvent.Raw.BlockNumber).
			Str("job contract", recvEvent.SapoJob.String()).
			Str("job spec cid", cid).
			Bool("removed", recvEvent.Raw.Removed).
			Msg("Event")

		if recvEvent.Raw.Removed {
			continue
		}

		// Get specs from cid
		lightSpecs, err := ParseSpecs(ctx, cid)
		// lightSpecs, err := DummySpecs(cid) // TODO: delete

		if err != nil { // retry?
			log.Ctx(ctx).Error().Err(err).Msg("Input CID parsing failed")
			continue
		}

		var spec []byte

		if len(lightSpecs.InputsCid) > 5 {
			spec, err = json.Marshal(model.Spec{
				Engine:    model.EngineDocker,
				Verifier:  model.VerifierNoop,
				Publisher: model.PublisherIpfs,
				Docker: model.JobSpecDocker{
					Image:      lightSpecs.Image,
					Entrypoint: lightSpecs.RunParams,
				},
				Inputs: []model.StorageSpec{
					{
						StorageSource: model.StorageSourceIPFS,
						Name:          "inputs",
						Path:          "/inputs",
						CID:           lightSpecs.InputsCid,
					},
				},
				Outputs: []model.StorageSpec{
					{
						Name: "outputs",
						Path: "/outputs",
					},
				},
				Deal: model.Deal{
					Concurrency: 1,
				},
			})
		} else {
			spec, err = json.Marshal(model.Spec{
				Engine:    model.EngineDocker,
				Verifier:  model.VerifierNoop,
				Publisher: model.PublisherIpfs,
				Docker: model.JobSpecDocker{
					Image:      lightSpecs.Image,
					Entrypoint: lightSpecs.RunParams,
				},
				Outputs: []model.StorageSpec{
					{
						Name: "outputs",
						Path: "/outputs",
					},
				},
				Deal: model.Deal{
					Concurrency: 1,
				},
			})
		}

		if err != nil {
			log.Ctx(ctx).Error().Err(err).Send()
			continue
		}

		out <- &event{
			orderId: recvEvent.Raw.TxHash.Bytes(),
			state:   OrderStateSubmitted,
			jobSpec: spec,
			jobAddr: recvEvent.SapoJob,
		}

		r.maxSeenBlock = recvEvent.Raw.BlockNumber
	}
}

// Refund implements SmartContract
func (r *RealContract) Refund(ctx context.Context, e ContractFailedEvent) (ContractRefundedEvent, error) {
	// TODO refund
	tx, err := r.contract.FailAndRefund(r.transact, e.Addr())

	localCtx := log.Ctx(ctx).With().
		Str("jobAddress", e.Addr().Hex()).
		Logger().WithContext(ctx)

	if err != nil {
		log.Ctx(localCtx).Err(err).Msg("Error in failAndRefund transaction")
		return e, context.Canceled
	}

	log.Ctx(localCtx).Debug().
		Stringer("transactionId", tx.Hash()).
		Msg("Job failed. Successfully refunded user")

	return e.Refunded(), nil
}

func NewContract(contractAddr common.Address, wallet_file string) (SmartContract, error) {
	client, err := ethclient.Dial("wss://ws-filecoin-hyperspace.chainstacklabs.com/rpc/v0")
	if err != nil {
		return nil, err
	}

	contract, err := contracts.NewContracts(contractAddr, client)
	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}

	password := "croaked!"

	reader, err := os.Open(wallet_file)
	if err != nil {
		return nil, err
	}

	opts, err := bind.NewTransactorWithChainID(reader, password, big.NewInt(3141))

	number, err := client.BlockNumber(context.Background())
	if err != nil {
		return nil, err
	}

	return &RealContract{contract, opts, number}, nil
}

type ContractCompleteHandler func(context.Context, BacalhauJobCompletedEvent) (ContractPaidEvent, error)
type ContractRefundHandler func(context.Context, ContractFailedEvent) (ContractRefundedEvent, error)
type ContractListenHandler func(ctx context.Context, c chan<- ContractSubmittedEvent) error

type mockContract struct {
	CompleteHandler ContractCompleteHandler
	RefundHandler   ContractRefundHandler
	ListenHandler   ContractListenHandler
}

// Complete implements SmartContract
func (mock mockContract) Complete(ctx context.Context, e BacalhauJobCompletedEvent, client publicapi.RequesterAPIClient) (ContractPaidEvent, error) {
	log.Ctx(ctx).Info().Stringer("id", e.OrderId()).Msg("Complete")
	if mock.CompleteHandler != nil {
		return mock.CompleteHandler(ctx, e)
	}
	return e.Paid(), nil
}

// Listen implements SmartContract
func (mock mockContract) Listen(ctx context.Context, out chan<- ContractSubmittedEvent) error {
	log.Ctx(ctx).Debug().Msg("Listening")
	defer log.Ctx(ctx).Debug().Msg("Stopping listening")

	if mock.ListenHandler != nil {
		return mock.ListenHandler(ctx, out)
	}
	return nil
}

func exampleEvent() ContractSubmittedEvent {
	spec, err := json.Marshal(model.Spec{
		Engine:    model.EngineDocker,
		Verifier:  model.VerifierNoop,
		Publisher: model.PublisherEstuary,
		Docker: model.JobSpecDocker{
			Image:      "ubuntu",
			Entrypoint: []string{"date"},
		},
		Deal: model.Deal{
			Concurrency: 1,
		},
	})
	if err != nil {
		panic(err)
	}
	e := &event{
		attempts:    0,
		lastAttempt: time.Time{},
		state:       OrderStateSubmitted,
		jobSpec:     spec,
	}
	id := make([]byte, 0, 32)
	_, err = rand.Read(id)
	if err != nil {
		panic(err)
	}
	copy(e.orderId[:], id[0:32])
	return e
}

// Refund implements SmartContract
func (mock mockContract) Refund(ctx context.Context, e ContractFailedEvent) (ContractRefundedEvent, error) {
	log.Ctx(ctx).Info().Stringer("id", e.OrderId()).Msg("Refunded")
	if mock.RefundHandler != nil {
		return mock.RefundHandler(ctx, e)
	}
	return e.Refunded(), nil
}

func TimerContract() SmartContract {
	return mockContract{
		ListenHandler: func(ctx context.Context, out chan<- ContractSubmittedEvent) error {
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
			return nil
		},
	}
}
