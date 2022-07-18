package main

import (
	"fs.video/log"
	cmdcfg "fs.video/smartchain/cmd/config"
	"fs.video/smartchain/cmd/node/cmd"
	scmd "fs.video/smartchain/cmd/scd/cmd"
	"fs.video/smartchain/core"
	svrcmd "github.com/cosmos/cosmos-sdk/server/cmd"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"time"
)


func Start(cosmosRepoPath string) {
	log := core.BuildLog(core.GetFuncName(), core.LmChainClient)
	var err error

	config := sdk.GetConfig()
	cmdcfg.SetBech32Prefixes(config)
	cmdcfg.SetBip44CoinType(config)
	config.Seal()
	cmdcfg.RegisterDenoms()

	log.Info("check config dir")

	if e, _ := checkSourceExist(filepath.Join(cosmosRepoPath, "config")); !e {
		log.Info("init chain")


		err = chainRun(core.CommandName, "init", "node1", "--chain-id", core.ChainID, "--home", cosmosRepoPath)
		if err != nil {
			panic(err)
		}

		err = replaceConfig(cosmosRepoPath)
		if err != nil {
			panic(err)
		}
	}
	log.Info("check genesis.json")

	err = checkGenesisFile(cosmosRepoPath)
	if err != nil {
		panic(err)
	}

	log.Info("check config.toml")

	err = checkConfigFile(cosmosRepoPath)
	if err != nil {
		panic(err)
	}


	err = checkAppToml(cosmosRepoPath)
	if err != nil {
		panic(err)
	}


	err = checkClientToml(cosmosRepoPath)
	if err != nil {
		panic(err)
	}


	err = checkValidatorStateJson(cosmosRepoPath)
	if err != nil {
		panic(err)
	}

	log.WithField("path", cosmosRepoPath).Info("chain repo")




	logLevel := "error"
	logLevelSet, ok := os.LookupEnv("SMART_CHAIN_LOGGING")
	if ok {
		logLevel = logLevelSet
	}

	log.Info("start chain")
	os.Args = []string{core.CommandName, "start", "--log_format", "json", "--log_level", logLevel, "--home", cosmosRepoPath}
	rootCmd, _ := scmd.NewRootCmd()
	rootCmd.SetErr(os.Stdout)
	rootCmd.SetOut(os.Stdout)
	if err := svrcmd.Execute(rootCmd, cosmosRepoPath); err != nil {

		os.Exit(1)
	}
}

func chainRun(arg ...string) error {
	os.Args = arg
	rootCmd, _ := scmd.NewRootCmd()
	if err := svrcmd.Execute(rootCmd, ""); err != nil {


		return err
	}
	return nil
}

func PathExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	return false
}

func startCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "start",
		Short: "start chain",
		Run: func(cmd *cobra.Command, args []string) {

			programPath, _ := filepath.Abs(os.Args[0])


			runtimePath, _ := filepath.Split(programPath)

			logPath := filepath.Join(runtimePath, "log")

			if !PathExists(logPath) {
				err := os.Mkdir(logPath, 0644)
				if err != nil {
					panic(err)
				}
			}

			daemonLogPath := filepath.Join(logPath, "chain.log")

			log.EnableLogStorage(daemonLogPath, time.Hour*24*7, time.Hour*24)

			pwd, err := os.Getwd()
			if err != nil {
				panic(err)
			}
			cosmosRepoPath := filepath.Join(pwd, ".scd")
			Start(cosmosRepoPath)
		},
	}
	return cmd
}

func main() {
	log.InitLogger(logrus.InfoLevel)

	rootCmd := &cobra.Command{
		Use:   core.CommandName,
		Short: "smart chain node",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	rootCmd.AddCommand(startCmd())
	rootCmd.AddCommand(cmd.VersionCmd())
	rootCmd.Execute()
}
