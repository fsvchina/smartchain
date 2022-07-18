package client

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/tendermint/tendermint/libs/cli"

	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)



const ClientContextKey = sdk.ContextKey("client.context")



func SetCmdClientContextHandler(clientCtx Context, cmd *cobra.Command) (err error) {
	clientCtx, err = ReadPersistentCommandFlags(clientCtx, cmd.Flags())
	if err != nil {
		return err
	}

	return SetCmdClientContext(cmd, clientCtx)
}


func ValidateCmd(cmd *cobra.Command, args []string) error {
	var unknownCmd string
	var skipNext bool

	for _, arg := range args {

		if arg == "--help" || arg == "-h" {
			return cmd.Help()
		}


		switch {
		case len(arg) > 0 && (arg[0] == '-'):


			if !strings.Contains(arg, "=") {
				skipNext = true
			} else {
				skipNext = false
			}
		case skipNext:

			skipNext = false
		case unknownCmd == "":


			unknownCmd = arg
		}
	}


	if unknownCmd != "" {
		err := fmt.Sprintf("unknown command \"%s\" for \"%s\"", unknownCmd, cmd.CalledAs())


		if suggestions := cmd.SuggestionsFor(unknownCmd); len(suggestions) > 0 {
			err += "\n\nDid you mean this?\n"
			for _, s := range suggestions {
				err += fmt.Sprintf("\t%v\n", s)
			}
		}
		return errors.New(err)
	}

	return cmd.Help()
}



//


//




func ReadPersistentCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {
	if clientCtx.OutputFormat == "" || flagSet.Changed(cli.OutputFlag) {
		output, _ := flagSet.GetString(cli.OutputFlag)
		clientCtx = clientCtx.WithOutputFormat(output)
	}

	if clientCtx.HomeDir == "" || flagSet.Changed(flags.FlagHome) {
		homeDir, _ := flagSet.GetString(flags.FlagHome)
		clientCtx = clientCtx.WithHomeDir(homeDir)
	}

	if !clientCtx.Simulate || flagSet.Changed(flags.FlagDryRun) {
		dryRun, _ := flagSet.GetBool(flags.FlagDryRun)
		clientCtx = clientCtx.WithSimulation(dryRun)
	}

	if clientCtx.KeyringDir == "" || flagSet.Changed(flags.FlagKeyringDir) {
		keyringDir, _ := flagSet.GetString(flags.FlagKeyringDir)



		if keyringDir == "" {
			keyringDir = clientCtx.HomeDir
		}

		clientCtx = clientCtx.WithKeyringDir(keyringDir)
	}

	if clientCtx.ChainID == "" || flagSet.Changed(flags.FlagChainID) {
		chainID, _ := flagSet.GetString(flags.FlagChainID)
		clientCtx = clientCtx.WithChainID(chainID)
	}

	if clientCtx.Keyring == nil || flagSet.Changed(flags.FlagKeyringBackend) {
		keyringBackend, _ := flagSet.GetString(flags.FlagKeyringBackend)

		if keyringBackend != "" {
			kr, err := NewKeyringFromBackend(clientCtx, keyringBackend)
			if err != nil {
				return clientCtx, err
			}

			clientCtx = clientCtx.WithKeyring(kr)
		}
	}

	if clientCtx.Client == nil || flagSet.Changed(flags.FlagNode) {
		rpcURI, _ := flagSet.GetString(flags.FlagNode)
		if rpcURI != "" {
			clientCtx = clientCtx.WithNodeURI(rpcURI)

			client, err := NewClientFromNode(rpcURI)
			if err != nil {
				return clientCtx, err
			}

			clientCtx = clientCtx.WithClient(client)
		}
	}

	return clientCtx, nil
}



//


//




func readQueryCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {
	if clientCtx.Height == 0 || flagSet.Changed(flags.FlagHeight) {
		height, _ := flagSet.GetInt64(flags.FlagHeight)
		clientCtx = clientCtx.WithHeight(height)
	}

	if !clientCtx.UseLedger || flagSet.Changed(flags.FlagUseLedger) {
		useLedger, _ := flagSet.GetBool(flags.FlagUseLedger)
		clientCtx = clientCtx.WithUseLedger(useLedger)
	}

	return ReadPersistentCommandFlags(clientCtx, flagSet)
}



//


//




func readTxCommandFlags(clientCtx Context, flagSet *pflag.FlagSet) (Context, error) {
	clientCtx, err := ReadPersistentCommandFlags(clientCtx, flagSet)
	if err != nil {
		return clientCtx, err
	}

	if !clientCtx.GenerateOnly || flagSet.Changed(flags.FlagGenerateOnly) {
		genOnly, _ := flagSet.GetBool(flags.FlagGenerateOnly)
		clientCtx = clientCtx.WithGenerateOnly(genOnly)
	}

	if !clientCtx.Offline || flagSet.Changed(flags.FlagOffline) {
		offline, _ := flagSet.GetBool(flags.FlagOffline)
		clientCtx = clientCtx.WithOffline(offline)
	}

	if !clientCtx.UseLedger || flagSet.Changed(flags.FlagUseLedger) {
		useLedger, _ := flagSet.GetBool(flags.FlagUseLedger)
		clientCtx = clientCtx.WithUseLedger(useLedger)
	}

	if clientCtx.BroadcastMode == "" || flagSet.Changed(flags.FlagBroadcastMode) {
		bMode, _ := flagSet.GetString(flags.FlagBroadcastMode)
		clientCtx = clientCtx.WithBroadcastMode(bMode)
	}

	if !clientCtx.SkipConfirm || flagSet.Changed(flags.FlagSkipConfirmation) {
		skipConfirm, _ := flagSet.GetBool(flags.FlagSkipConfirmation)
		clientCtx = clientCtx.WithSkipConfirmation(skipConfirm)
	}

	if clientCtx.SignModeStr == "" || flagSet.Changed(flags.FlagSignMode) {
		signModeStr, _ := flagSet.GetString(flags.FlagSignMode)
		clientCtx = clientCtx.WithSignModeStr(signModeStr)
	}

	if clientCtx.FeeGranter == nil || flagSet.Changed(flags.FlagFeeAccount) {
		granter, _ := flagSet.GetString(flags.FlagFeeAccount)

		if granter != "" {
			granterAcc, err := sdk.AccAddressFromBech32(granter)
			if err != nil {
				return clientCtx, err
			}

			clientCtx = clientCtx.WithFeeGranterAddress(granterAcc)
		}
	}

	if clientCtx.From == "" || flagSet.Changed(flags.FlagFrom) {
		from, _ := flagSet.GetString(flags.FlagFrom)
		fromAddr, fromName, keyType, err := GetFromFields(clientCtx.Keyring, from, clientCtx.GenerateOnly)
		if err != nil {
			return clientCtx, err
		}

		clientCtx = clientCtx.WithFrom(from).WithFromAddress(fromAddr).WithFromName(fromName)




		if keyType == keyring.TypeLedger && clientCtx.SignModeStr != flags.SignModeLegacyAminoJSON {
			fmt.Println("Default sign-mode 'direct' not supported by Ledger, using sign-mode 'amino-json'.")
			clientCtx = clientCtx.WithSignModeStr(flags.SignModeLegacyAminoJSON)
		}
	}
	return clientCtx, nil
}



//




func GetClientQueryContext(cmd *cobra.Command) (Context, error) {
	ctx := GetClientContextFromCmd(cmd)
	return readQueryCommandFlags(ctx, cmd.Flags())
}



//




func GetClientTxContext(cmd *cobra.Command) (Context, error) {
	ctx := GetClientContextFromCmd(cmd)
	return readTxCommandFlags(ctx, cmd.Flags())
}



func GetClientContextFromCmd(cmd *cobra.Command) Context {
	if v := cmd.Context().Value(ClientContextKey); v != nil {
		clientCtxPtr := v.(*Context)
		return *clientCtxPtr
	}

	return Context{}
}


func SetCmdClientContext(cmd *cobra.Command, clientCtx Context) error {
	v := cmd.Context().Value(ClientContextKey)
	if v == nil {
		return errors.New("client context not set")
	}

	clientCtxPtr := v.(*Context)
	*clientCtxPtr = clientCtx

	return nil
}
