package bridge

import (
	"context"
	"encoding/json"
	"time"

	"github.com/filecoin-project/bacalhau/pkg/model"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type SmartContract interface {
	Listen(context.Context, chan<- ContractSubmittedEvent) error

	Complete(context.Context, BacalhauJobCompletedEvent) (ContractPaidEvent, error)

	Refund(context.Context, ContractFailedEvent) (ContractRefundedEvent, error)
}

type ContractListenHandler func(ctx context.Context, c chan<- ContractSubmittedEvent) error
type ContractCompleteHandler func(context.Context, BacalhauJobCompletedEvent) (ContractPaidEvent, error)
type ContractRefundHandler func(context.Context, ContractFailedEvent) (ContractRefundedEvent, error)

// Complete implements SmartContract
func (mock mockContract) Complete(ctx context.Context, e BacalhauJobCompletedEvent) (ContractPaidEvent, error) {
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

// Refund implements SmartContract
func (mock mockContract) Refund(ctx context.Context, e ContractFailedEvent) (ContractRefundedEvent, error) {
	log.Ctx(ctx).Info().Stringer("id", e.OrderId()).Msg("Refunded")
	if mock.RefundHandler != nil {
		return mock.RefundHandler(ctx, e)
	}
	return e.Refunded(), nil
}

type mockContract struct {
	ListenHandler   ContractListenHandler
	CompleteHandler ContractCompleteHandler
	RefundHandler   ContractRefundHandler
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
	return &event{
		orderId:     uuid.New(),
		attempts:    0,
		lastAttempt: time.Time{},
		state:       OrderStateSubmitted,
		jobSpec:     spec,
	}
}
