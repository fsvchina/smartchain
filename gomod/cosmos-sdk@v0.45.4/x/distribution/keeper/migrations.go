package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	v043 "github.com/cosmos/cosmos-sdk/x/distribution/legacy/v043"
)


type Migrator struct {
	keeper Keeper
}


func NewMigrator(keeper Keeper) Migrator {
	return Migrator{keeper: keeper}
}


func (m Migrator) Migrate1to2(ctx sdk.Context) error {
	return v043.MigrateStore(ctx, m.keeper.storeKey)
}
