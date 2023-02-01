package bridge

import (
	"context"
	"encoding/json"

	"github.com/filecoin-project/bacalhau/pkg/model"
)

func debugJob(ctx context.Context, out chan<- ContractSubmittedEvent) error {
	image := "quintenbons/bacalhau-inout"
	param := []string{}
	cid := "QmZ1fH7nFpxN3jhtb5wSrEnSNxb435ShyYtZhudwKksDXy"

	spec, err := json.Marshal(model.Spec{
		Engine:    model.EngineDocker,
		Verifier:  model.VerifierNoop,
		Publisher: model.PublisherIpfs,
		Docker: model.JobSpecDocker{
			Image:      image,
			Entrypoint: param,
		},
		Resources: model.ResourceUsageConfig{
			GPU: "1",
		},
		Inputs: []model.StorageSpec{
			{
				Name: "inputs",
				Path: "/inputs",
				CID:  cid,
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

	out <- &event{
		orderId: []byte{}, // dummy
		state:   OrderStateSubmitted,
		jobSpec: spec,
	}

	if err != nil {
		return err
	}

	return nil
}
