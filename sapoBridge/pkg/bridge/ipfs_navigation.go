package bridge

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	ipfsapi "github.com/ipfs/go-ipfs-api"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
)

type LightJobSpec struct {
	Image     string   `json:"image"`
	InputsCid string   `json:"inputsCid"`
	RunParams []string `json:"runParams"`
}

func ParseSpecsTimeout(cid string, timeout time.Duration) (*LightJobSpec, error) {
	ipfsShell := ipfsapi.NewLocalShell()

	if ipfsShell == nil {
		return nil, errors.New("Could not instanciate ipfs shell")
	}

	ipfsShell.SetTimeout(timeout)
	out, err := ipfsShell.Cat(cid)

	if err != nil {
		return nil, errors.Wrap(err, "Could not read CID")
	}

	jobSpecs := LightJobSpec{}
	decoder := json.NewDecoder(out)

	err = decoder.Decode(&jobSpecs)

	if err != nil {
		return nil, errors.Wrap(err, "Could not parse JSON contents of CID")
	}

	return &jobSpecs, nil
}

func ParseSpecsIPFS(ctx context.Context, cid string, tries int) (*LightJobSpec, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}
	out, err := client.Get("https://ipfs.io/ipfs/" + cid)

	if err != nil {
		if tries >= 10 {
			log.Ctx(ctx).Error().Err(err).Msg("http get failed. Abandonning")
			return nil, err
		}

		log.Ctx(ctx).Debug().Err(err).Msg(fmt.Sprintf("http get failed. Retrying in 3 seconds (%d tries)", tries))
		time.Sleep(5 * time.Second)
		return ParseSpecsIPFS(ctx, cid, tries+1)
	}

	defer out.Body.Close()

	jobSpecs := LightJobSpec{}
	decoder := json.NewDecoder(out.Body)

	err = decoder.Decode(&jobSpecs)

	log.Ctx(ctx).Info().Err(err).Msg("Retrieved cid " + cid + " from ipfs.io")

	return &jobSpecs, nil
}

// Parses specs from cid. Cid should reference a folder containing two files:
// {inputsFileName}: The /inputs file (also on ipfs, so we can just get the cid)
// {specFileName}: The specs file (1st line image, following parameters)
func ParseSpecs(ctx context.Context, cid string) (*LightJobSpec, error) {
	log.Ctx(ctx).Debug().Msg("Getting from ipfs with go-ipfs")
	out, err := ParseSpecsTimeout(cid, 10*time.Second)

	if err != nil {
		log.Ctx(ctx).Debug().Err(err).Msg("Get failed. Retrying with ipfs.io")
		out, err = ParseSpecsIPFS(ctx, cid, 0)
	}

	return out, err
}

// dummy specs for debugging
func DummySpecs(cid string) (*LightJobSpec, error) {
	jobSpec := LightJobSpec{
		Image:     "quintenbons/bacalhau-inout:0.2",
		RunParams: []string{},
		InputsCid: "QmUt1CYoJwDTknfeTZVdrjzGuHnhnBVAFzzEcMBp2MFMCh",
	}

	return &jobSpec, nil
}
