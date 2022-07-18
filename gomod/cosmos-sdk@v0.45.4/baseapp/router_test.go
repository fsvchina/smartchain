package baseapp

import (
	"testing"

	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var testHandler = func(_ sdk.Context, _ sdk.Msg) (*sdk.Result, error) {
	return &sdk.Result{}, nil
}

func TestRouter(t *testing.T) {
	rtr := NewRouter()


	require.Panics(t, func() {
		rtr.AddRoute(sdk.NewRoute("*", testHandler))
	})

	rtr.AddRoute(sdk.NewRoute("testRoute", testHandler))
	h := rtr.Route(sdk.Context{}, "testRoute")
	require.NotNil(t, h)


	require.Panics(t, func() {
		rtr.AddRoute(sdk.NewRoute("testRoute", testHandler))
	})
}
