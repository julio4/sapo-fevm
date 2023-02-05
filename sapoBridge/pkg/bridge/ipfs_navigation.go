package bridge

import (
	"encoding/json"

	ipfsapi "github.com/ipfs/go-ipfs-api"
	"github.com/pkg/errors"
)

type LightJobSpec struct {
	Image     string   `json:"image"`
	InputsCid string   `json:"inputsCid"`
	RunParams []string `json:"runParams"`
}

// Parses specs from cid. Cid should reference a folder containing two files:
// {inputsFileName}: The /inputs file (also on ipfs, so we can just get the cid)
// {specFileName}: The specs file (1st line image, following parameters)
func ParseSpecs(cid string) (*LightJobSpec, error) {
	ipfsShell := ipfsapi.NewLocalShell()
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

// dummy specs for debugging
func DummySpecs(cid string) (*LightJobSpec, error) {
	jobSpec := LightJobSpec{
		Image:     "quintenbons/bacalhau-inout:0.2",
		RunParams: []string{},
		InputsCid: "QmUt1CYoJwDTknfeTZVdrjzGuHnhnBVAFzzEcMBp2MFMCh",
	}

	return &jobSpec, nil
}
