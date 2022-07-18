package cli

import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/x/params/types"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
)


func NewQueryCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Querying commands for the params module",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(NewQuerySubspaceParamsCmd())

	return cmd
}



func NewQuerySubspaceParamsCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "subspace [subspace] [key]",
		Short: "Query for raw parameters by subspace and key",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			queryClient := proposal.NewQueryClient(clientCtx)

			params := proposal.QueryParamsRequest{Subspace: args[0], Key: args[1]}
			res, err := queryClient.Params(cmd.Context(), &params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&res.Param)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
