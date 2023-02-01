package bridge

import (
	"io"
	"strings"

	ipfsapi "github.com/ipfs/go-ipfs-api"
)

const specFileName string = "specs.bac"
const inputsFileName string = "inputs.bac"

type LightJobSpec struct {
	image     string
	params    []string
	inputsCid string
}

// Parses specs from cid. Cid should reference a folder containing two files:
// {inputsFileName}: The /inputs file (also on ipfs, so we can just get the cid)
// {specFileName}: The specs file (1st line image, following parameters)
func parseSpecs(cid string) (*LightJobSpec, error) {
	// TODO: better error managment
	ipfsShell := ipfsapi.NewLocalShell()
	inputCid, err := ipfsShell.ResolvePath(cid + "/" + inputsFileName)

	if err != nil {
		return nil, err
	}

	specCid, err := ipfsShell.ResolvePath(cid + "/" + specFileName)

	if err != nil {
		return nil, err
	}

	reader, err := ipfsShell.Cat(specCid)
	if err != nil {
		return nil, err
	}

	contents, err := io.ReadAll(reader)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(contents), "\n")

	jobSpec := LightJobSpec{
		image:     lines[0],
		params:    lines[1:],
		inputsCid: inputCid,
	}

	return &jobSpec, err
}

// dummy specs for debugging
func dummySpecs(cid string) (*LightJobSpec, error) {
	jobSpec := LightJobSpec{
		image:     "quintenbons/bacalhau-inout",
		params:    []string{},
		inputsCid: "QmZ1fH7nFpxN3jhtb5wSrEnSNxb435ShyYtZhudwKksDXy",
	}

	return &jobSpec, nil
}
