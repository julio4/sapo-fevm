package main

import (
	"log"
	"sapoBridge/pkg/bridge"
)

func main() {
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
