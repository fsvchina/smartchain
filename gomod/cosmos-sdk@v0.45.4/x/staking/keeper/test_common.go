package keeper

import (
	"bytes"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)


func ValidatorByPowerIndexExists(ctx sdk.Context, keeper Keeper, power []byte) bool {
	store := ctx.KVStore(keeper.storeKey)
	return store.Has(power)
}


func TestingUpdateValidator(keeper Keeper, ctx sdk.Context, validator types.Validator, apply bool) types.Validator {
	keeper.SetValidator(ctx, validator)


	store := ctx.KVStore(keeper.storeKey)
	deleted := false

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsByPowerIndexKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		valAddr := types.ParseValidatorPowerRankKey(iterator.Key())
		if bytes.Equal(valAddr, validator.GetOperator()) {
			if deleted {
				panic("found duplicate power index key")
			} else {
				deleted = true
			}

			store.Delete(iterator.Key())
		}
	}

	keeper.SetValidatorByPowerIndex(ctx, validator)

	if !apply {
		ctx, _ = ctx.CacheContext()
	}
	_, err := keeper.ApplyAndReturnValidatorSetUpdates(ctx)
	if err != nil {
		panic(err)
	}

	validator, found := keeper.GetValidator(ctx, validator.GetOperator())
	if !found {
		panic("validator expected but not found")
	}

	return validator
}


func RandomValidator(r *rand.Rand, keeper Keeper, ctx sdk.Context) (val types.Validator, ok bool) {
	vals := keeper.GetAllValidators(ctx)
	if len(vals) == 0 {
		return types.Validator{}, false
	}

	i := r.Intn(len(vals))

	return vals[i], true
}
