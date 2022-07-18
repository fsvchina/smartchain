package main

import (
	"fs.video/smartchain/cmd/scd/cmd"
	"os"

	"github.com/cosmos/cosmos-sdk/server"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"fs.video/smartchain/app"
	cmdcfg "fs.video/smartchain/cmd/config"
)

func main() {
	setupConfig()
	cmdcfg.RegisterDenoms()

	rootCmd, _ := cmd.NewRootCmd()

	if err := svrcmd.Execute(rootCmd, app.DefaultNodeHome); err != nil {
		switch e := err.(type) {
		case server.ErrorCode:
			os.Exit(e.Code)

		default:
			os.Exit(1)
		}
	}
}

func setupConfig() {

	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)




	cmdcfg.SetBip44CoinType(config)
	config.Seal()
}
