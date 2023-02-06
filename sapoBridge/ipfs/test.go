package main

import (
	"log"
	"sapoBridge/pkg/bridge"
	// ipfsapi "github.com/ipfs/go-ipfs-api"
)

func main() {
	// convertToBytes()

	// ipfsShell := ipfsapi.NewLocalShell()
	// out, _ := ipfsShell.Cat("bafybeic4tveejbvbm5frqohfhi64rzbvjv3iknq5xr4ng46laa3gsfdqgi")

	// log.Print(string(str))

}

func convertToBytes() {
	cid := "Qmbji4cLwj35C7pfhLWZTte2XX8J6obg6hHVKG6KtXvR5M"
	cid1, _ := bridge.PackToHex(cid[:32])
	cid2, _ := bridge.PackToHex(cid[32:])

	log.Print(cid1)
	log.Print(cid2)
}

func parseSpecsCid() {
	jobSpec, err := bridge.ParseSpecs("QmcFV6bjk7JdyyPDCCaSUBigggDP4sq77DtZ2V6s53nnsR")

	if err != nil {
		log.Print("ERROR HAS OCCURED")
		log.Fatal(err)
	}

	log.Print("ALLRIGHT")
	log.Print(jobSpec.Image)
	log.Print(jobSpec.InputsCid)
	log.Print(jobSpec.RunParams)
}
