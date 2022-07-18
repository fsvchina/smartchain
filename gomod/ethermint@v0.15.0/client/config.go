package client

import (
	"fmt"
	"os"
	"path"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"

	ethermint "github.com/tharsis/ethermint/types"
)


func InitConfig(cmd *cobra.Command) error {
	home, err := cmd.PersistentFlags().GetString(cli.HomeFlag)
	if err != nil {
		return err
	}

	configFile := path.Join(home, "config", "config.toml")
	_, err = os.Stat(configFile)
	if err != nil && !os.IsNotExist(err) {


		return err
	}
	if err == nil {
		viper.SetConfigFile(configFile)

		if err := viper.ReadInConfig(); err != nil {
			return err
		}
	}

	if err := viper.BindPFlag(flags.FlagChainID, cmd.PersistentFlags().Lookup(flags.FlagChainID)); err != nil {
		return err
	}

	if err := viper.BindPFlag(cli.EncodingFlag, cmd.PersistentFlags().Lookup(cli.EncodingFlag)); err != nil {
		return err
	}

	return viper.BindPFlag(cli.OutputFlag, cmd.PersistentFlags().Lookup(cli.OutputFlag))
}


func ValidateChainID(baseCmd *cobra.Command) *cobra.Command {

	baseRunE := baseCmd.RunE


	validateFn := func(cmd *cobra.Command, args []string) error {
		chainID, _ := cmd.Flags().GetString(flags.FlagChainID)

		if !ethermint.IsValidChainID(chainID) {
			return fmt.Errorf("invalid chain-id format: %s", chainID)
		}

		return baseRunE(cmd, args)
	}

	baseCmd.RunE = validateFn
	return baseCmd
}
