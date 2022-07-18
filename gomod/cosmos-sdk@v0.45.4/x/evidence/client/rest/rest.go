package rest

import (
	"net/http"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/rest"

	"github.com/gorilla/mux"
)


const (
	RestParamEvidenceHash = "evidence-hash"

	MethodGet = "GET"
)



type EvidenceRESTHandler struct {
	SubRoute string
	Handler  func(http.ResponseWriter, *http.Request)
}



func RegisterRoutes(clientCtx client.Context, rtr *mux.Router, handlers []EvidenceRESTHandler) {
	r := rest.WithHTTPDeprecationHeaders(rtr)

	registerQueryRoutes(clientCtx, r)
	registerTxRoutes(clientCtx, r, handlers)
}
