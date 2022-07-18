package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestCalculateRewardsBasic(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	addr := simapp.AddTestAddrs(app, ctx, 2, sdk.NewInt(1000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addr)


	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[0], valConsPk1, sdk.NewInt(100), true)


	staking.EndBlocker(ctx, app.StakingKeeper)
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)
	tstaking.Ctx = ctx


	val := app.StakingKeeper.Validator(ctx, valAddrs[0])
	del := app.StakingKeeper.Delegation(ctx, sdk.AccAddress(valAddrs[0]), valAddrs[0])


	require.Equal(t, uint64(2), app.DistrKeeper.GetValidatorHistoricalReferenceCount(ctx))


	endingPeriod := app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	require.Equal(t, uint64(2), app.DistrKeeper.GetValidatorHistoricalReferenceCount(ctx))


	rewards := app.DistrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)


	require.True(t, rewards.IsZero())


	initial := int64(10)
	tokens := sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial)}}
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	endingPeriod = app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial / 2)}}, rewards)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial / 2)}}, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission)
}

func TestCalculateRewardsAfterSlash(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addr := simapp.AddTestAddrs(app, ctx, 2, sdk.NewInt(100000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addr)
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	valPower := int64(100)
	tstaking.CreateValidatorWithValPower(valAddrs[0], valConsPk1, valPower, true)


	staking.EndBlocker(ctx, app.StakingKeeper)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	val := app.StakingKeeper.Validator(ctx, valAddrs[0])
	del := app.StakingKeeper.Delegation(ctx, sdk.AccAddress(valAddrs[0]), valAddrs[0])


	endingPeriod := app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards := app.DistrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)


	require.True(t, rewards.IsZero())


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)


	app.StakingKeeper.Slash(ctx, valConsAddr1, ctx.BlockHeight(), valPower, sdk.NewDecWithPrec(5, 1))


	val = app.StakingKeeper.Validator(ctx, valAddrs[0])


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)


	initial := app.StakingKeeper.TokensFromConsensusPower(ctx, 10)
	tokens := sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial.ToDec()}}
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	endingPeriod = app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial.QuoRaw(2).ToDec()}}, rewards)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial.QuoRaw(2).ToDec()}},
		app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission)
}

func TestCalculateRewardsAfterManySlashes(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	addr := simapp.AddTestAddrs(app, ctx, 2, sdk.NewInt(100000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addr)


	valPower := int64(100)
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	tstaking.CreateValidatorWithValPower(valAddrs[0], valConsPk1, valPower, true)


	staking.EndBlocker(ctx, app.StakingKeeper)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	val := app.StakingKeeper.Validator(ctx, valAddrs[0])
	del := app.StakingKeeper.Delegation(ctx, sdk.AccAddress(valAddrs[0]), valAddrs[0])


	endingPeriod := app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards := app.DistrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)


	require.True(t, rewards.IsZero())


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)


	app.StakingKeeper.Slash(ctx, valConsAddr1, ctx.BlockHeight(), valPower, sdk.NewDecWithPrec(5, 1))


	val = app.StakingKeeper.Validator(ctx, valAddrs[0])


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)


	initial := app.StakingKeeper.TokensFromConsensusPower(ctx, 10)
	tokens := sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial.ToDec()}}
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	app.StakingKeeper.Slash(ctx, valConsAddr1, ctx.BlockHeight(), valPower/2, sdk.NewDecWithPrec(5, 1))


	val = app.StakingKeeper.Validator(ctx, valAddrs[0])


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)


	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	endingPeriod = app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial.ToDec()}}, rewards)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial.ToDec()}},
		app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission)
}

func TestCalculateRewardsMultiDelegator(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	addr := simapp.AddTestAddrs(app, ctx, 2, sdk.NewInt(100000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addr)


	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[0], valConsPk1, sdk.NewInt(100), true)


	staking.EndBlocker(ctx, app.StakingKeeper)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	val := app.StakingKeeper.Validator(ctx, valAddrs[0])
	del1 := app.StakingKeeper.Delegation(ctx, sdk.AccAddress(valAddrs[0]), valAddrs[0])


	initial := int64(20)
	tokens := sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial)}}
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	tstaking.Ctx = ctx
	tstaking.Delegate(sdk.AccAddress(valAddrs[1]), valAddrs[0], sdk.NewInt(100))
	del2 := app.StakingKeeper.Delegation(ctx, sdk.AccAddress(valAddrs[1]), valAddrs[0])


	val = app.StakingKeeper.Validator(ctx, valAddrs[0])


	staking.EndBlocker(ctx, app.StakingKeeper)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	endingPeriod := app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards := app.DistrKeeper.CalculateDelegationRewards(ctx, val, del1, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial * 3 / 4)}}, rewards)


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del2, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial * 1 / 4)}}, rewards)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial)}}, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission)
}

func TestWithdrawDelegationRewardsBasic(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	balancePower := int64(1000)
	balanceTokens := app.StakingKeeper.TokensFromConsensusPower(ctx, balancePower)
	addr := simapp.AddTestAddrs(app, ctx, 1, sdk.NewInt(1000000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addr)
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	distrAcc := app.DistrKeeper.GetDistributionAccount(ctx)
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, distrAcc.GetName(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, balanceTokens))))
	app.AccountKeeper.SetModuleAccount(ctx, distrAcc)


	power := int64(100)
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	valTokens := tstaking.CreateValidatorWithValPower(valAddrs[0], valConsPk1, power, true)


	expTokens := balanceTokens.Sub(valTokens)
	require.Equal(t,
		sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, expTokens)},
		app.BankKeeper.GetAllBalances(ctx, sdk.AccAddress(valAddrs[0])),
	)


	staking.EndBlocker(ctx, app.StakingKeeper)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	val := app.StakingKeeper.Validator(ctx, valAddrs[0])


	initial := app.StakingKeeper.TokensFromConsensusPower(ctx, 10)
	tokens := sdk.DecCoins{sdk.NewDecCoin(sdk.DefaultBondDenom, initial)}

	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	require.Equal(t, uint64(2), app.DistrKeeper.GetValidatorHistoricalReferenceCount(ctx))


	_, err := app.DistrKeeper.WithdrawDelegationRewards(ctx, sdk.AccAddress(valAddrs[0]), valAddrs[0])
	require.Nil(t, err)


	require.Equal(t, uint64(2), app.DistrKeeper.GetValidatorHistoricalReferenceCount(ctx))


	exp := balanceTokens.Sub(valTokens).Add(initial.QuoRaw(2))
	require.Equal(t,
		sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, exp)},
		app.BankKeeper.GetAllBalances(ctx, sdk.AccAddress(valAddrs[0])),
	)


	_, err = app.DistrKeeper.WithdrawValidatorCommission(ctx, valAddrs[0])
	require.Nil(t, err)


	exp = balanceTokens.Sub(valTokens).Add(initial)
	require.Equal(t,
		sdk.Coins{sdk.NewCoin(sdk.DefaultBondDenom, exp)},
		app.BankKeeper.GetAllBalances(ctx, sdk.AccAddress(valAddrs[0])),
	)
}

func TestCalculateRewardsAfterManySlashesInSameBlock(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	addr := simapp.AddTestAddrs(app, ctx, 1, sdk.NewInt(1000000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addr)
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	valPower := int64(100)
	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	tstaking.CreateValidatorWithValPower(valAddrs[0], valConsPk1, valPower, true)


	staking.EndBlocker(ctx, app.StakingKeeper)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	val := app.StakingKeeper.Validator(ctx, valAddrs[0])
	del := app.StakingKeeper.Delegation(ctx, sdk.AccAddress(valAddrs[0]), valAddrs[0])


	endingPeriod := app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards := app.DistrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)


	require.True(t, rewards.IsZero())


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)


	initial := app.StakingKeeper.TokensFromConsensusPower(ctx, 10).ToDec()
	tokens := sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial}}
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	app.StakingKeeper.Slash(ctx, valConsAddr1, ctx.BlockHeight(), valPower, sdk.NewDecWithPrec(5, 1))


	app.StakingKeeper.Slash(ctx, valConsAddr1, ctx.BlockHeight(), valPower/2, sdk.NewDecWithPrec(5, 1))


	val = app.StakingKeeper.Validator(ctx, valAddrs[0])


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)


	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	endingPeriod = app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial}}, rewards)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial}}, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission)
}

func TestCalculateRewardsMultiDelegatorMultiSlash(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	addr := simapp.AddTestAddrs(app, ctx, 2, sdk.NewInt(1000000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addr)


	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	valPower := int64(100)
	tstaking.CreateValidatorWithValPower(valAddrs[0], valConsPk1, valPower, true)


	staking.EndBlocker(ctx, app.StakingKeeper)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	val := app.StakingKeeper.Validator(ctx, valAddrs[0])
	del1 := app.StakingKeeper.Delegation(ctx, sdk.AccAddress(valAddrs[0]), valAddrs[0])


	initial := app.StakingKeeper.TokensFromConsensusPower(ctx, 30).ToDec()
	tokens := sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial}}
	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)
	app.StakingKeeper.Slash(ctx, valConsAddr1, ctx.BlockHeight(), valPower, sdk.NewDecWithPrec(5, 1))
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)


	tstaking.DelegateWithPower(sdk.AccAddress(valAddrs[1]), valAddrs[0], 100)

	del2 := app.StakingKeeper.Delegation(ctx, sdk.AccAddress(valAddrs[1]), valAddrs[0])


	staking.EndBlocker(ctx, app.StakingKeeper)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)
	app.StakingKeeper.Slash(ctx, valConsAddr1, ctx.BlockHeight(), valPower, sdk.NewDecWithPrec(5, 1))
	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 3)


	val = app.StakingKeeper.Validator(ctx, valAddrs[0])


	endingPeriod := app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards := app.DistrKeeper.CalculateDelegationRewards(ctx, val, del1, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial.QuoInt64(2).Add(initial.QuoInt64(6))}}, rewards)


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del2, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial.QuoInt64(3)}}, rewards)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: initial}}, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission)
}

func TestCalculateRewardsMultiDelegatorMultWithdraw(t *testing.T) {
	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	addr := simapp.AddTestAddrs(app, ctx, 2, sdk.NewInt(1000000000))
	valAddrs := simapp.ConvertAddrsToValAddrs(addr)
	initial := int64(20)


	distrAcc := app.DistrKeeper.GetDistributionAccount(ctx)
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, distrAcc.GetName(), sdk.NewCoins(sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)))))
	app.AccountKeeper.SetModuleAccount(ctx, distrAcc)

	tokens := sdk.DecCoins{sdk.NewDecCoinFromDec(sdk.DefaultBondDenom, sdk.NewDec(initial))}


	tstaking.Commission = stakingtypes.NewCommissionRates(sdk.NewDecWithPrec(5, 1), sdk.NewDecWithPrec(5, 1), sdk.NewDec(0))
	tstaking.CreateValidator(valAddrs[0], valConsPk1, sdk.NewInt(100), true)


	staking.EndBlocker(ctx, app.StakingKeeper)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	val := app.StakingKeeper.Validator(ctx, valAddrs[0])
	del1 := app.StakingKeeper.Delegation(ctx, sdk.AccAddress(valAddrs[0]), valAddrs[0])


	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	require.Equal(t, uint64(2), app.DistrKeeper.GetValidatorHistoricalReferenceCount(ctx))


	tstaking.Delegate(sdk.AccAddress(valAddrs[1]), valAddrs[0], sdk.NewInt(100))


	require.Equal(t, uint64(3), app.DistrKeeper.GetValidatorHistoricalReferenceCount(ctx))


	val = app.StakingKeeper.Validator(ctx, valAddrs[0])
	del2 := app.StakingKeeper.Delegation(ctx, sdk.AccAddress(valAddrs[1]), valAddrs[0])


	staking.EndBlocker(ctx, app.StakingKeeper)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	_, err := app.DistrKeeper.WithdrawDelegationRewards(ctx, sdk.AccAddress(valAddrs[0]), valAddrs[0])
	require.NoError(t, err)


	_, err = app.DistrKeeper.WithdrawDelegationRewards(ctx, sdk.AccAddress(valAddrs[1]), valAddrs[0])
	require.NoError(t, err)


	require.Equal(t, uint64(3), app.DistrKeeper.GetValidatorHistoricalReferenceCount(ctx))


	_, err = app.DistrKeeper.WithdrawValidatorCommission(ctx, valAddrs[0])
	require.NoError(t, err)


	endingPeriod := app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards := app.DistrKeeper.CalculateDelegationRewards(ctx, val, del1, endingPeriod)


	require.True(t, rewards.IsZero())


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del2, endingPeriod)


	require.True(t, rewards.IsZero())


	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission.IsZero())


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	_, err = app.DistrKeeper.WithdrawDelegationRewards(ctx, sdk.AccAddress(valAddrs[0]), valAddrs[0])
	require.NoError(t, err)


	endingPeriod = app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del1, endingPeriod)


	require.True(t, rewards.IsZero())


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del2, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial / 4)}}, rewards)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial / 2)}}, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission)


	ctx = ctx.WithBlockHeight(ctx.BlockHeight() + 1)


	app.DistrKeeper.AllocateTokensToValidator(ctx, val, tokens)


	_, err = app.DistrKeeper.WithdrawValidatorCommission(ctx, valAddrs[0])
	require.NoError(t, err)


	endingPeriod = app.DistrKeeper.IncrementValidatorPeriod(ctx, val)


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del1, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial / 4)}}, rewards)


	rewards = app.DistrKeeper.CalculateDelegationRewards(ctx, val, del2, endingPeriod)


	require.Equal(t, sdk.DecCoins{{Denom: sdk.DefaultBondDenom, Amount: sdk.NewDec(initial / 2)}}, rewards)


	require.True(t, app.DistrKeeper.GetValidatorAccumulatedCommission(ctx, valAddrs[0]).Commission.IsZero())
}
