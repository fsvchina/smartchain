package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/evidence/exported"
)

type (




	Handler func(sdk.Context, exported.Evidence) error



	Router interface {
		AddRoute(r string, h Handler) Router
		HasRoute(r string) bool
		GetRoute(path string) Handler
		Seal()
		Sealed() bool
	}

	router struct {
		routes map[string]Handler
		sealed bool
	}
)

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


func (rtr router) Sealed() bool {
	return rtr.sealed
}



func (rtr *router) AddRoute(path string, h Handler) Router {
	if rtr.sealed {
		panic(fmt.Sprintf("router sealed; cannot register %s route handler", path))
	}
	if !sdk.IsAlphaNumeric(path) {
		panic("route expressions can only contain alphanumeric characters")
	}
	if rtr.HasRoute(path) {
		panic(fmt.Sprintf("route %s has already been registered", path))
	}

	rtr.routes[path] = h
	return rtr
}


func (rtr *router) HasRoute(path string) bool {
	return rtr.routes[path] != nil
}


func (rtr *router) GetRoute(path string) Handler {
	if !rtr.HasRoute(path) {
		panic(fmt.Sprintf("route does not exist for path %s", path))
	}
	return rtr.routes[path]
}
