package rpc

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec/legacy"
	"github.com/cosmos/cosmos-sdk/types/rest"
)


func BlockCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "block [height]",
		Short: "Get verified data for a the block at given height",
		Args:  cobra.MaximumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}
			var height *int64


			if len(args) > 0 {
				h, err := strconv.Atoi(args[0])
				if err != nil {
					return err
				}
				if h > 0 {
					tmp := int64(h)
					height = &tmp
				}
			}

			output, err := getBlock(clientCtx, height)
			if err != nil {
				return err
			}

			fmt.Println(string(output))
			return nil
		},
	}

	cmd.Flags().StringP(flags.FlagNode, "n", "tcp:

	return cmd
}

func getBlock(clientCtx client.Context, height *int64) ([]byte, error) {

	node, err := clientCtx.GetNode()
	if err != nil {
		return nil, err
	}




	res, err := node.Block(context.Background(), height)
	if err != nil {
		return nil, err
	}

	return legacy.Cdc.MarshalJSON(res)
}


func GetChainHeight(clientCtx client.Context) (int64, error) {
	node, err := clientCtx.GetNode()
	if err != nil {
		return -1, err
	}

	status, err := node.Status(context.Background())
	if err != nil {
		return -1, err
	}

	height := status.SyncInfo.LatestBlockHeight
	return height, nil
}


func BlockRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		height, err := strconv.ParseInt(vars["height"], 10, 64)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusBadRequest,
				"couldn't parse block height. Assumed format is '/block/{height}'.")
			return
		}

		chainHeight, err := GetChainHeight(clientCtx)
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, "failed to parse chain height")
			return
		}

		if height > chainHeight {
			rest.WriteErrorResponse(w, http.StatusNotFound, "requested block height is bigger then the chain length")
			return
		}

		output, err := getBlock(clientCtx, &height)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponseBare(w, clientCtx, output)
	}
}


func LatestBlockRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		output, err := getBlock(clientCtx, nil)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		rest.PostProcessResponseBare(w, clientCtx, output)
	}
}
