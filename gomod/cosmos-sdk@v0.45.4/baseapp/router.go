package baseapp

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Router struct {
	routes map[string]sdk.Handler
}

var _ sdk.Router = NewRouter()


func NewRouter() *Router {
	return &Router{
		routes: make(map[string]sdk.Handler),
	}
}



func (rtr *Router) AddRoute(route sdk.Route) sdk.Router {
	if !sdk.IsAlphaNumeric(route.Path()) {
		panic("route expressions can only contain alphanumeric characters")
	}
	if rtr.routes[route.Path()] != nil {
		panic(fmt.Sprintf("route %s has already been initialized", route.Path()))
	}

	rtr.routes[route.Path()] = route.Handler()
	return rtr
}


//

func (rtr *Router) Route(_ sdk.Context, path string) sdk.Handler {
	return rtr.routes[path]
}
