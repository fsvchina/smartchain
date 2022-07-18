package rest

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	clientrest "github.com/cosmos/cosmos-sdk/client/rest"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	genutilrest "github.com/cosmos/cosmos-sdk/x/genutil/client/rest"
)


func QueryAccountRequestHandlerFn(storeName string, clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		bech32addr := vars["address"]

		addr, err := sdk.AccAddressFromBech32(bech32addr)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		accGetter := types.AccountRetriever{}

		account, height, err := accGetter.GetAccountWithHeight(clientCtx, addr)
		if err != nil {


			if err := accGetter.EnsureExists(clientCtx, addr); err != nil {
				clientCtx = clientCtx.WithHeight(height)
				rest.PostProcessResponse(w, clientCtx, types.BaseAccount{})
				return
			}

			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, account)
	}
}




func QueryTxsRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			rest.WriteErrorResponse(
				w, http.StatusBadRequest,
				fmt.Sprintf("failed to parse query parameters: %s", err),
			)
			return
		}


		heightStr := r.FormValue("height")
		if heightStr != "" {
			if height, err := strconv.ParseInt(heightStr, 10, 64); err == nil && height == 0 {
				genutilrest.QueryGenesisTxs(clientCtx, w)
				return
			}
		}

		var (
			events      []string
			txs         []sdk.TxResponse
			page, limit int
		)

		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		if len(r.Form) == 0 {
			rest.PostProcessResponseBare(w, clientCtx, txs)
			return
		}

		events, page, limit, err = rest.ParseHTTPArgs(r)
		if rest.CheckBadRequestError(w, err) {
			return
		}

		searchResult, err := authtx.QueryTxsByEvents(clientCtx, events, page, limit, "")
		if rest.CheckInternalServerError(w, err) {
			return
		}

		for _, txRes := range searchResult.Txs {
			packStdTxResponse(w, clientCtx, txRes)
		}

		err = checkAminoMarshalError(clientCtx, searchResult, "/cosmos/tx/v1beta1/txs")
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())

			return
		}

		rest.PostProcessResponseBare(w, clientCtx, searchResult)
	}
}



func QueryTxRequestHandlerFn(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		hashHexStr := vars["hash"]

		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		output, err := authtx.QueryTx(clientCtx, hashHexStr)
		if err != nil {
			if strings.Contains(err.Error(), hashHexStr) {
				rest.WriteErrorResponse(w, http.StatusNotFound, err.Error())
				return
			}
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())
			return
		}

		err = packStdTxResponse(w, clientCtx, output)
		if err != nil {

			return
		}

		if output.Empty() {
			rest.WriteErrorResponse(w, http.StatusNotFound, fmt.Sprintf("no transaction found with hash %s", hashHexStr))
		}

		err = checkAminoMarshalError(clientCtx, output, "/cosmos/tx/v1beta1/txs/{txhash}")
		if err != nil {
			rest.WriteErrorResponse(w, http.StatusInternalServerError, err.Error())

			return
		}

		rest.PostProcessResponseBare(w, clientCtx, output)
	}
}

func queryParamsHandler(clientCtx client.Context) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		clientCtx, ok := rest.ParseQueryHeightOrReturnBadRequest(w, clientCtx, r)
		if !ok {
			return
		}

		route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryParams)
		res, height, err := clientCtx.QueryWithData(route, nil)
		if rest.CheckInternalServerError(w, err) {
			return
		}

		clientCtx = clientCtx.WithHeight(height)
		rest.PostProcessResponse(w, clientCtx, res)
	}
}




func packStdTxResponse(w http.ResponseWriter, clientCtx client.Context, txRes *sdk.TxResponse) error {


	txBytes := txRes.Tx.Value
	stdTx, err := convertToStdTx(w, clientCtx, txBytes)
	if err != nil {
		return err
	}


	txRes.Tx = codectypes.UnsafePackAny(stdTx)

	return nil
}



func checkAminoMarshalError(ctx client.Context, resp interface{}, grpcEndPoint string) error {

	marshaler := ctx.LegacyAmino

	_, err := marshaler.MarshalJSON(resp)
	if err != nil {



		return fmt.Errorf("this transaction cannot be displayed via legacy REST endpoints, because it does not support"+
			" Amino serialization. Please either use CLI, gRPC, gRPC-gateway, or directly query the Tendermint RPC"+
			" endpoint to query this transaction. The new REST endpoint (via gRPC-gateway) is %s. Please also see the"+
			"REST endpoints migration guide at %s for more info", grpcEndPoint, clientrest.DeprecationURL)

	}

	return nil
}
