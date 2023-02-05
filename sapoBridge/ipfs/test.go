package main

import (
	"log"
	"sapoBridge/pkg/bridge"
)

func main() {
	convertToBytes()

	// ipfsShell := ipfsapi.NewLocalShell()
	// out, _ := ipfsShell.Cat("QmPCVpJLVjTdDPEq3B5g7KVAS1ozJ8vT4p3toaxV8q5YVk")
	// a, _ := io.ReadAll(out)
	// log.Print(string(a))
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
