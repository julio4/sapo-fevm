package main

import (
	"context"
	"sapoBridge/pkg/bridge"

	"github.com/rs/zerolog/log"
	// ipfsapi "github.com/ipfs/go-ipfs-api"
)

func main() {
	convertToBytes()

	// ipfsShell := ipfsapi.NewLocalShell()
	// out, _ := ipfsShell.Cat("bafybeic4tveejbvbm5frqohfhi64rzbvjv3iknq5xr4ng46laa3gsfdqgi")

	// log.Print(string(str))

	// apiPort := 1234
	// apiHost := "35.245.115.191"
	// client := publicapi.NewRequesterAPIClient(fmt.Sprintf("http://%s:%d", apiHost, apiPort))
	// ctx := log.Logger.WithContext(context.Background())

	// jobs, _ := client.GetResults(ctx, "2cb7f3ed-efa3-4cff-a8e9-621c14abef07")

	// println("OK")
	// fmt.Printf("CID: %s DONE", jobs[0].Data.CID)
}

func convertToBytes() {
	cid := "QmQ2UTVYh8Z8YXpeBFtRcujpqppohSmQMR6YuiWZPiKPYR"
	cid1, _ := bridge.PackToHex(cid[:32])
	cid2, _ := bridge.PackToHex(cid[32:])

	println(cid1[:])
	println(cid2[:])
}

func parseSpecsCid(ctx context.Context) {
	jobSpec, err := bridge.ParseSpecs(ctx, "QmcFV6bjk7JdyyPDCCaSUBigggDP4sq77DtZ2V6s53nnsR")

	if err != nil {
		log.Print("ERROR HAS OCCURED")
		log.Err(err)
	}

	log.Print("ALLRIGHT")
	log.Print(jobSpec.Image)
	log.Print(jobSpec.InputsCid)
	log.Print(jobSpec.RunParams)
}
