package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ Router = (*router)(nil)


//

type Router interface {
	AddRoute(r string, h Handler) (rtr Router)
	HasRoute(r string) bool
	GetRoute(path string) (h Handler)
	Seal()
}

type router struct {
	routes map[string]Handler
	sealed bool
}


func NewRouter() Router {
	return &router{
		routes: make(map[string]Handler),
	}
}



func (rtr *router) Seal() {
	if rtr.sealed {
		panic("router already sealed")
	}
	rtr.sealed = true
}



func (rtr *router) AddRoute(path string, h Handler) Router {
	if rtr.sealed {
		panic("router sealed; cannot add route handler")
	}

	if !sdk.IsAlphaNumeric(path) {
		panic("route expressions can only contain alphanumeric characters")
	}
	if rtr.HasRoute(path) {
		panic(fmt.Sprintf("route %s has already been initialized", path))
	}

	rtr.routes[path] = h
	return rtr
}


func (rtr *router) HasRoute(path string) bool {
	return rtr.routes[path] != nil
}


func (rtr *router) GetRoute(path string) Handler {
	if !rtr.HasRoute(path) {
		panic(fmt.Sprintf("route \"%s\" does not exist", path))
	}

	return rtr.routes[path]
}
