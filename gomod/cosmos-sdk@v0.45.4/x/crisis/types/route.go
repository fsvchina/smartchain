package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)


type InvarRoute struct {
	ModuleName string
	Route      string
	Invar      sdk.Invariant
}


func NewInvarRoute(moduleName, route string, invar sdk.Invariant) InvarRoute {
	return InvarRoute{
		ModuleName: moduleName,
		Route:      route,
		Invar:      invar,
	}
}


func (i InvarRoute) FullRoute() string {
	return i.ModuleName + "/" + i.Route
}
