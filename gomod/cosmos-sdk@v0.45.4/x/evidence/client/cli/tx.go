package cli

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/x/evidence/types"

	"github.com/spf13/cobra"
)






func GetTxCmd(childCmds []*cobra.Command) *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Evidence transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	submitEvidenceCmd := SubmitEvidenceCmd()
	for _, childCmd := range childCmds {
		submitEvidenceCmd.AddCommand(childCmd)
	}



	return cmd
}




func SubmitEvidenceCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit",
		Short: "Submit arbitrary evidence of misbehavior",
	}

	return cmd
}
