package baseapp

import (
	"testing"

	"github.com/stretchr/testify/require"

	abci "github.com/tendermint/tendermint/abci/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

var testQuerier = func(_ sdk.Context, _ []string, _ abci.RequestQuery) ([]byte, error) {
	return nil, nil
}

func TestQueryRouter(t *testing.T) {
	qr := NewQueryRouter()


	require.Panics(t, func() {
		qr.AddRoute("*", testQuerier)
	})

	qr.AddRoute("testRoute", testQuerier)
	q := qr.Route("testRoute")
	require.NotNil(t, q)


	require.Panics(t, func() {
		qr.AddRoute("testRoute", testQuerier)
	})
}
