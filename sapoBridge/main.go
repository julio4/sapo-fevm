package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	"sapoBridge/pkg/bridge"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	ctx := log.Logger.WithContext(context.Background())
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt)
	defer cancel()

	repo, err := bridge.NewSQLiteRepository(ctx, "lilypad.sqlite")
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	addr := common.HexToAddress("0x8F25c6c517545303153731c19BE75D3064B2BDf4")
	wallet_filename := "./wallets/bridgeaccount"
	contract, err := bridge.NewContract(addr, wallet_filename)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		return
	}

	workflow := bridge.NewWorkflow(bridge.NewJobRunner(), contract, repo)

	err = workflow.Start(ctx)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
	}
}
