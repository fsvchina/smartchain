package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)


func (k Keeper) GetHistoricalInfo(ctx sdk.Context, height int64) (types.HistoricalInfo, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetHistoricalInfoKey(height)

	value := store.Get(key)
	if value == nil {
		return types.HistoricalInfo{}, false
	}

	return types.MustUnmarshalHistoricalInfo(k.cdc, value), true
}


func (k Keeper) SetHistoricalInfo(ctx sdk.Context, height int64, hi *types.HistoricalInfo) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetHistoricalInfoKey(height)
	value := k.cdc.MustMarshal(hi)
	store.Set(key, value)
}


func (k Keeper) DeleteHistoricalInfo(ctx sdk.Context, height int64) {
	store := ctx.KVStore(k.storeKey)
	key := types.GetHistoricalInfoKey(height)

	store.Delete(key)
}




func (k Keeper) IterateHistoricalInfo(ctx sdk.Context, cb func(types.HistoricalInfo) bool) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.HistoricalInfoKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		histInfo := types.MustUnmarshalHistoricalInfo(k.cdc, iterator.Value())
		if cb(histInfo) {
			break
		}
	}
}


func (k Keeper) GetAllHistoricalInfo(ctx sdk.Context) []types.HistoricalInfo {
	var infos []types.HistoricalInfo

	k.IterateHistoricalInfo(ctx, func(histInfo types.HistoricalInfo) bool {
		infos = append(infos, histInfo)
		return false
	})

	return infos
}



func (k Keeper) TrackHistoricalInfo(ctx sdk.Context) {
	entryNum := k.HistoricalEntries(ctx)








	for i := ctx.BlockHeight() - int64(entryNum); i >= 0; i-- {
		_, found := k.GetHistoricalInfo(ctx, i)
		if found {
			k.DeleteHistoricalInfo(ctx, i)
		} else {
			break
		}
	}


	if entryNum == 0 {
		return
	}


	lastVals := k.GetLastValidators(ctx)
	historicalEntry := types.NewHistoricalInfo(ctx.BlockHeader(), lastVals, k.PowerReduction(ctx))


	k.SetHistoricalInfo(ctx, ctx.BlockHeight(), &historicalEntry)
}
