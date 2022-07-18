package keeper

import "github.com/cosmos/cosmos-sdk/x/gov/types"




func UnsafeSetHooks(k *Keeper, h types.GovHooks) {
	k.hooks = h
}
