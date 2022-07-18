package baseapp

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type state struct {
	ms  sdk.CacheMultiStore
	ctx sdk.Context
}



func (st *state) CacheMultiStore() sdk.CacheMultiStore {
	return st.ms.CacheMultiStore()
}


func (st *state) Context() sdk.Context {
	return st.ctx
}
