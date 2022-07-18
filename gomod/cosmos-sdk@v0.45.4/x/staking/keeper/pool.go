package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)


func (k Keeper) GetBondedPool(ctx sdk.Context) (bondedPool authtypes.ModuleAccountI) {
	return k.authKeeper.GetModuleAccount(ctx, types.BondedPoolName)
}


func (k Keeper) GetNotBondedPool(ctx sdk.Context) (notBondedPool authtypes.ModuleAccountI) {
	return k.authKeeper.GetModuleAccount(ctx, types.NotBondedPoolName)
}


func (k Keeper) bondedTokensToNotBonded(ctx sdk.Context, tokens sdk.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(k.BondDenom(ctx), tokens))
	if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.BondedPoolName, types.NotBondedPoolName, coins); err != nil {
		panic(err)
	}
}


func (k Keeper) notBondedTokensToBonded(ctx sdk.Context, tokens sdk.Int) {
	coins := sdk.NewCoins(sdk.NewCoin(k.BondDenom(ctx), tokens))
	if err := k.bankKeeper.SendCoinsFromModuleToModule(ctx, types.NotBondedPoolName, types.BondedPoolName, coins); err != nil {
		panic(err)
	}
}


func (k Keeper) burnBondedTokens(ctx sdk.Context, amt sdk.Int) error {
	if !amt.IsPositive() {

		return nil
	}

	coins := sdk.NewCoins(sdk.NewCoin(k.BondDenom(ctx), amt))

	return k.bankKeeper.BurnCoins(ctx, types.BondedPoolName, coins)
}


func (k Keeper) burnNotBondedTokens(ctx sdk.Context, amt sdk.Int) error {
	if !amt.IsPositive() {

		return nil
	}

	coins := sdk.NewCoins(sdk.NewCoin(k.BondDenom(ctx), amt))

	return k.bankKeeper.BurnCoins(ctx, types.NotBondedPoolName, coins)
}


func (k Keeper) TotalBondedTokens(ctx sdk.Context) sdk.Int {
	bondedPool := k.GetBondedPool(ctx)
	return k.bankKeeper.GetBalance(ctx, bondedPool.GetAddress(), k.BondDenom(ctx)).Amount
}


func (k Keeper) StakingTokenSupply(ctx sdk.Context) sdk.Int {
	return k.bankKeeper.GetSupply(ctx, k.BondDenom(ctx)).Amount
}


func (k Keeper) BondedRatio(ctx sdk.Context) sdk.Dec {
	stakeSupply := k.StakingTokenSupply(ctx)
	if stakeSupply.IsPositive() {
		return k.TotalBondedTokens(ctx).ToDec().QuoInt(stakeSupply)
	}

	return sdk.ZeroDec()
}
