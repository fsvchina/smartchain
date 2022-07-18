package rest

import (
	"github.com/gorilla/mux"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rest"
)


const (
	MethodGet = "GET"
)


func RegisterRoutes(clientCtx client.Context, rtr *mux.Router, storeName string) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	r.HandleFunc(
		"/auth/accounts/{address}", QueryAccountRequestHandlerFn(storeName, clientCtx),
	).Methods(MethodGet)

	r.HandleFunc(
		"/auth/params",
		queryParamsHandler(clientCtx),
	).Methods(MethodGet)
}


func RegisterTxRoutes(clientCtx client.Context, rtr *mux.Router) {
	r := rest.WithHTTPDeprecationHeaders(rtr)
	r.HandleFunc("/txs/{hash}", QueryTxRequestHandlerFn(clientCtx)).Methods("GET")
	r.HandleFunc("/txs", QueryTxsRequestHandlerFn(clientCtx)).Methods("GET")
	r.HandleFunc("/txs/decode", DecodeTxRequestHandlerFn(clientCtx)).Methods("POST")
}
