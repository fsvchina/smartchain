package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

	v010 "github.com/tharsis/ethermint/x/feemarket/migrations/v010"
)


type Migrator struct {
	keeper Keeper
}


func NewMigrator(keeper Keeper) Migrator {
	return Migrator{
		keeper: keeper,
	}
}


func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v010.MigrateStore(ctx, &m.keeper.paramSpace, m.keeper.storeKey)
}
