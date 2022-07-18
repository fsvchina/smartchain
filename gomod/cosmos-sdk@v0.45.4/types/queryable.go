package types

import (
	abci "github.com/tendermint/tendermint/abci/types"
)



type Querier = func(ctx Context, path []string, req abci.RequestQuery) ([]byte, error)
