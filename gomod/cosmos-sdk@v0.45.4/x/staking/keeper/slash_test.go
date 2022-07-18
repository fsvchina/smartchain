package keeper_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)


func bootstrapSlashTest(t *testing.T, power int64) (*simapp.SimApp, sdk.Context, []sdk.AccAddress, []sdk.ValAddress) {
	_, app, ctx := createTestInput()

	addrDels, addrVals := generateAddresses(app, ctx, 100)

	amt := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
	totalSupply := sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), amt.MulRaw(int64(len(addrDels)))))

	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), totalSupply))

	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)

	numVals := int64(3)
	bondedCoins := sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), amt.MulRaw(numVals)))
	bondedPool := app.StakingKeeper.GetBondedPool(ctx)


	app.AccountKeeper.SetModuleAccount(ctx, bondedPool)
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, bondedPool.GetName(), bondedCoins))

	for i := int64(0); i < numVals; i++ {
		validator := teststaking.NewValidator(t, addrVals[i], PKs[i])
		validator, _ = validator.AddTokensFromDel(amt)
		validator = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validator, true)
		app.StakingKeeper.SetValidatorByConsAddr(ctx, validator)
	}

	return app, ctx, addrDels, addrVals
}


func TestRevocation(t *testing.T) {
	app, ctx, _, addrVals := bootstrapSlashTest(t, 5)

	consAddr := sdk.ConsAddress(PKs[0].Address())


	val, found := app.StakingKeeper.GetValidator(ctx, addrVals[0])
	require.True(t, found)
	require.False(t, val.IsJailed())


	app.StakingKeeper.Jail(ctx, consAddr)
	val, found = app.StakingKeeper.GetValidator(ctx, addrVals[0])
	require.True(t, found)
	require.True(t, val.IsJailed())


	app.StakingKeeper.Unjail(ctx, consAddr)
	val, found = app.StakingKeeper.GetValidator(ctx, addrVals[0])
	require.True(t, found)
	require.False(t, val.IsJailed())
}


func TestSlashUnbondingDelegation(t *testing.T) {
	app, ctx, addrDels, addrVals := bootstrapSlashTest(t, 10)

	fraction := sdk.NewDecWithPrec(5, 1)



	ubd := types.NewUnbondingDelegation(addrDels[0], addrVals[0], 0,
		time.Unix(5, 0), sdk.NewInt(10))

	app.StakingKeeper.SetUnbondingDelegation(ctx, ubd)


	slashAmount := app.StakingKeeper.SlashUnbondingDelegation(ctx, ubd, 1, fraction)
	require.True(t, slashAmount.Equal(sdk.NewInt(0)))


	ctx = ctx.WithBlockHeader(tmproto.Header{Time: time.Unix(10, 0)})
	app.StakingKeeper.SetUnbondingDelegation(ctx, ubd)
	slashAmount = app.StakingKeeper.SlashUnbondingDelegation(ctx, ubd, 0, fraction)
	require.True(t, slashAmount.Equal(sdk.NewInt(0)))


	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)
	oldUnbondedPoolBalances := app.BankKeeper.GetAllBalances(ctx, notBondedPool.GetAddress())
	ctx = ctx.WithBlockHeader(tmproto.Header{Time: time.Unix(0, 0)})
	app.StakingKeeper.SetUnbondingDelegation(ctx, ubd)
	slashAmount = app.StakingKeeper.SlashUnbondingDelegation(ctx, ubd, 0, fraction)
	require.True(t, slashAmount.Equal(sdk.NewInt(5)))
	ubd, found := app.StakingKeeper.GetUnbondingDelegation(ctx, addrDels[0], addrVals[0])
	require.True(t, found)
	require.Len(t, ubd.Entries, 1)


	require.Equal(t, sdk.NewInt(10), ubd.Entries[0].InitialBalance)


	require.Equal(t, sdk.NewInt(5), ubd.Entries[0].Balance)
	newUnbondedPoolBalances := app.BankKeeper.GetAllBalances(ctx, notBondedPool.GetAddress())
	diffTokens := oldUnbondedPoolBalances.Sub(newUnbondedPoolBalances)
	require.True(t, diffTokens.AmountOf(app.StakingKeeper.BondDenom(ctx)).Equal(sdk.NewInt(5)))
}


func TestSlashRedelegation(t *testing.T) {
	app, ctx, addrDels, addrVals := bootstrapSlashTest(t, 10)
	fraction := sdk.NewDecWithPrec(5, 1)


	startCoins := sdk.NewCoins(sdk.NewInt64Coin(app.StakingKeeper.BondDenom(ctx), 15))
	bondedPool := app.StakingKeeper.GetBondedPool(ctx)
	balances := app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, bondedPool.GetName(), startCoins))
	app.AccountKeeper.SetModuleAccount(ctx, bondedPool)



	rd := types.NewRedelegation(addrDels[0], addrVals[0], addrVals[1], 0,
		time.Unix(5, 0), sdk.NewInt(10), sdk.NewDec(10))

	app.StakingKeeper.SetRedelegation(ctx, rd)


	del := types.NewDelegation(addrDels[0], addrVals[1], sdk.NewDec(10))
	app.StakingKeeper.SetDelegation(ctx, del)


	validator, found := app.StakingKeeper.GetValidator(ctx, addrVals[1])
	require.True(t, found)
	slashAmount := app.StakingKeeper.SlashRedelegation(ctx, validator, rd, 1, fraction)
	require.True(t, slashAmount.Equal(sdk.NewInt(0)))


	ctx = ctx.WithBlockHeader(tmproto.Header{Time: time.Unix(10, 0)})
	app.StakingKeeper.SetRedelegation(ctx, rd)
	validator, found = app.StakingKeeper.GetValidator(ctx, addrVals[1])
	require.True(t, found)
	slashAmount = app.StakingKeeper.SlashRedelegation(ctx, validator, rd, 0, fraction)
	require.True(t, slashAmount.Equal(sdk.NewInt(0)))

	balances = app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())


	ctx = ctx.WithBlockHeader(tmproto.Header{Time: time.Unix(0, 0)})
	app.StakingKeeper.SetRedelegation(ctx, rd)
	validator, found = app.StakingKeeper.GetValidator(ctx, addrVals[1])
	require.True(t, found)
	slashAmount = app.StakingKeeper.SlashRedelegation(ctx, validator, rd, 0, fraction)
	require.True(t, slashAmount.Equal(sdk.NewInt(5)))
	rd, found = app.StakingKeeper.GetRedelegation(ctx, addrDels[0], addrVals[0], addrVals[1])
	require.True(t, found)
	require.Len(t, rd.Entries, 1)


	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 1)


	require.Equal(t, sdk.NewInt(10), rd.Entries[0].InitialBalance)


	del, found = app.StakingKeeper.GetDelegation(ctx, addrDels[0], addrVals[1])
	require.True(t, found)
	require.Equal(t, int64(5), del.Shares.RoundInt64())


	burnedCoins := sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), slashAmount))
	require.Equal(t, balances.Sub(burnedCoins), app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress()))
}


func TestSlashAtFutureHeight(t *testing.T) {
	app, ctx, _, _ := bootstrapSlashTest(t, 10)

	consAddr := sdk.ConsAddress(PKs[0].Address())
	fraction := sdk.NewDecWithPrec(5, 1)
	require.Panics(t, func() { app.StakingKeeper.Slash(ctx, consAddr, 1, 10, fraction) })
}



func TestSlashAtNegativeHeight(t *testing.T) {
	app, ctx, _, _ := bootstrapSlashTest(t, 10)
	consAddr := sdk.ConsAddress(PKs[0].Address())
	fraction := sdk.NewDecWithPrec(5, 1)

	bondedPool := app.StakingKeeper.GetBondedPool(ctx)
	oldBondedPoolBalances := app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())

	validator, found := app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)
	app.StakingKeeper.Slash(ctx, consAddr, -2, 10, fraction)


	validator, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)


	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 1)

	validator, found = app.StakingKeeper.GetValidator(ctx, validator.GetOperator())
	require.True(t, found)

	require.Equal(t, int64(5), validator.GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))


	newBondedPoolBalances := app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())
	diffTokens := oldBondedPoolBalances.Sub(newBondedPoolBalances).AmountOf(app.StakingKeeper.BondDenom(ctx))
	require.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 5).String(), diffTokens.String())
}


func TestSlashValidatorAtCurrentHeight(t *testing.T) {
	app, ctx, _, _ := bootstrapSlashTest(t, 10)
	consAddr := sdk.ConsAddress(PKs[0].Address())
	fraction := sdk.NewDecWithPrec(5, 1)

	bondedPool := app.StakingKeeper.GetBondedPool(ctx)
	oldBondedPoolBalances := app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())

	validator, found := app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)
	app.StakingKeeper.Slash(ctx, consAddr, ctx.BlockHeight(), 10, fraction)


	validator, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)


	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 1)

	validator, found = app.StakingKeeper.GetValidator(ctx, validator.GetOperator())
	assert.True(t, found)

	require.Equal(t, int64(5), validator.GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))


	newBondedPoolBalances := app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())
	diffTokens := oldBondedPoolBalances.Sub(newBondedPoolBalances).AmountOf(app.StakingKeeper.BondDenom(ctx))
	require.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 5).String(), diffTokens.String())
}


func TestSlashWithUnbondingDelegation(t *testing.T) {
	app, ctx, addrDels, addrVals := bootstrapSlashTest(t, 10)

	consAddr := sdk.ConsAddress(PKs[0].Address())
	fraction := sdk.NewDecWithPrec(5, 1)



	ubdTokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 4)
	ubd := types.NewUnbondingDelegation(addrDels[0], addrVals[0], 11, time.Unix(0, 0), ubdTokens)
	app.StakingKeeper.SetUnbondingDelegation(ctx, ubd)


	ctx = ctx.WithBlockHeight(12)
	bondedPool := app.StakingKeeper.GetBondedPool(ctx)
	oldBondedPoolBalances := app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())

	validator, found := app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)
	app.StakingKeeper.Slash(ctx, consAddr, 10, 10, fraction)


	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 1)


	ubd, found = app.StakingKeeper.GetUnbondingDelegation(ctx, addrDels[0], addrVals[0])
	require.True(t, found)
	require.Len(t, ubd.Entries, 1)


	require.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 2), ubd.Entries[0].Balance)


	newBondedPoolBalances := app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())
	diffTokens := oldBondedPoolBalances.Sub(newBondedPoolBalances).AmountOf(app.StakingKeeper.BondDenom(ctx))
	require.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 3), diffTokens)


	validator, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)





	require.Equal(t, int64(7), validator.GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))


	ctx = ctx.WithBlockHeight(13)
	app.StakingKeeper.Slash(ctx, consAddr, 9, 10, fraction)

	ubd, found = app.StakingKeeper.GetUnbondingDelegation(ctx, addrDels[0], addrVals[0])
	require.True(t, found)
	require.Len(t, ubd.Entries, 1)


	require.Equal(t, sdk.NewInt(0), ubd.Entries[0].Balance)


	newBondedPoolBalances = app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())
	diffTokens = oldBondedPoolBalances.Sub(newBondedPoolBalances).AmountOf(app.StakingKeeper.BondDenom(ctx))
	require.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 6), diffTokens)


	validator, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)


	require.Equal(t, int64(4), validator.GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))





	ctx = ctx.WithBlockHeight(13)
	app.StakingKeeper.Slash(ctx, consAddr, 9, 10, fraction)

	ubd, found = app.StakingKeeper.GetUnbondingDelegation(ctx, addrDels[0], addrVals[0])
	require.True(t, found)
	require.Len(t, ubd.Entries, 1)


	require.Equal(t, sdk.NewInt(0), ubd.Entries[0].Balance)


	newBondedPoolBalances = app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())
	diffTokens = oldBondedPoolBalances.Sub(newBondedPoolBalances).AmountOf(app.StakingKeeper.BondDenom(ctx))
	require.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 9), diffTokens)


	validator, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)


	require.Equal(t, int64(1), validator.GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))





	ctx = ctx.WithBlockHeight(13)
	app.StakingKeeper.Slash(ctx, consAddr, 9, 10, fraction)

	ubd, found = app.StakingKeeper.GetUnbondingDelegation(ctx, addrDels[0], addrVals[0])
	require.True(t, found)
	require.Len(t, ubd.Entries, 1)


	require.Equal(t, sdk.NewInt(0), ubd.Entries[0].Balance)


	newBondedPoolBalances = app.BankKeeper.GetAllBalances(ctx, bondedPool.GetAddress())
	diffTokens = oldBondedPoolBalances.Sub(newBondedPoolBalances).AmountOf(app.StakingKeeper.BondDenom(ctx))
	require.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 10), diffTokens)


	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, -1)




	validator, _ = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.Equal(t, validator.GetStatus(), types.Unbonding)
}


func TestSlashWithRedelegation(t *testing.T) {
	app, ctx, addrDels, addrVals := bootstrapSlashTest(t, 10)
	consAddr := sdk.ConsAddress(PKs[0].Address())
	fraction := sdk.NewDecWithPrec(5, 1)
	bondDenom := app.StakingKeeper.BondDenom(ctx)


	rdTokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 6)
	rd := types.NewRedelegation(addrDels[0], addrVals[0], addrVals[1], 11,
		time.Unix(0, 0), rdTokens, rdTokens.ToDec())
	app.StakingKeeper.SetRedelegation(ctx, rd)


	del := types.NewDelegation(addrDels[0], addrVals[1], rdTokens.ToDec())
	app.StakingKeeper.SetDelegation(ctx, del)


	bondedPool := app.StakingKeeper.GetBondedPool(ctx)
	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)
	rdCoins := sdk.NewCoins(sdk.NewCoin(bondDenom, rdTokens.MulRaw(2)))

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, bondedPool.GetName(), rdCoins))

	app.AccountKeeper.SetModuleAccount(ctx, bondedPool)

	oldBonded := app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount
	oldNotBonded := app.BankKeeper.GetBalance(ctx, notBondedPool.GetAddress(), bondDenom).Amount


	ctx = ctx.WithBlockHeight(12)
	validator, found := app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)

	require.NotPanics(t, func() { app.StakingKeeper.Slash(ctx, consAddr, 10, 10, fraction) })
	burnAmount := app.StakingKeeper.TokensFromConsensusPower(ctx, 10).ToDec().Mul(fraction).TruncateInt()

	bondedPool = app.StakingKeeper.GetBondedPool(ctx)
	notBondedPool = app.StakingKeeper.GetNotBondedPool(ctx)


	bondedPoolBalance := app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldBonded.Sub(burnAmount), bondedPoolBalance))

	notBondedPoolBalance := app.BankKeeper.GetBalance(ctx, notBondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldNotBonded, notBondedPoolBalance))
	oldBonded = app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount


	rd, found = app.StakingKeeper.GetRedelegation(ctx, addrDels[0], addrVals[0], addrVals[1])
	require.True(t, found)
	require.Len(t, rd.Entries, 1)

	validator, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)




	require.Equal(t, int64(8), validator.GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))


	validator, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)

	require.NotPanics(t, func() { app.StakingKeeper.Slash(ctx, consAddr, 10, 10, sdk.OneDec()) })
	burnAmount = app.StakingKeeper.TokensFromConsensusPower(ctx, 7)


	bondedPool = app.StakingKeeper.GetBondedPool(ctx)
	notBondedPool = app.StakingKeeper.GetNotBondedPool(ctx)


	bondedPoolBalance = app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldBonded.Sub(burnAmount), bondedPoolBalance))
	require.True(sdk.IntEq(t, oldNotBonded, notBondedPoolBalance))

	bondedPoolBalance = app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldBonded.Sub(burnAmount), bondedPoolBalance))

	notBondedPoolBalance = app.BankKeeper.GetBalance(ctx, notBondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldNotBonded, notBondedPoolBalance))
	oldBonded = app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount


	rd, found = app.StakingKeeper.GetRedelegation(ctx, addrDels[0], addrVals[0], addrVals[1])
	require.True(t, found)
	require.Len(t, rd.Entries, 1)

	validator, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)

	require.Equal(t, int64(4), validator.GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))


	ctx = ctx.WithBlockHeight(12)
	validator, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.True(t, found)

	require.NotPanics(t, func() { app.StakingKeeper.Slash(ctx, consAddr, 10, 10, sdk.OneDec()) })

	burnAmount = app.StakingKeeper.TokensFromConsensusPower(ctx, 10).ToDec().Mul(sdk.OneDec()).TruncateInt()
	burnAmount = burnAmount.Sub(sdk.OneDec().MulInt(rdTokens).TruncateInt())


	bondedPool = app.StakingKeeper.GetBondedPool(ctx)
	notBondedPool = app.StakingKeeper.GetNotBondedPool(ctx)

	bondedPoolBalance = app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldBonded.Sub(burnAmount), bondedPoolBalance))
	notBondedPoolBalance = app.BankKeeper.GetBalance(ctx, notBondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldNotBonded, notBondedPoolBalance))
	oldBonded = app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount


	rd, found = app.StakingKeeper.GetRedelegation(ctx, addrDels[0], addrVals[0], addrVals[1])
	require.True(t, found)
	require.Len(t, rd.Entries, 1)

	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, -1)


	validator, _ = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.Equal(t, validator.GetStatus(), types.Unbonding)



	ctx = ctx.WithBlockHeight(12)

	validator, _ = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.Equal(t, validator.GetStatus(), types.Unbonding)

	require.NotPanics(t, func() { app.StakingKeeper.Slash(ctx, consAddr, 10, 10, sdk.OneDec()) })


	bondedPool = app.StakingKeeper.GetBondedPool(ctx)
	notBondedPool = app.StakingKeeper.GetNotBondedPool(ctx)

	bondedPoolBalance = app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldBonded, bondedPoolBalance))
	notBondedPoolBalance = app.BankKeeper.GetBalance(ctx, notBondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldNotBonded, notBondedPoolBalance))


	rd, found = app.StakingKeeper.GetRedelegation(ctx, addrDels[0], addrVals[0], addrVals[1])
	require.True(t, found)
	require.Len(t, rd.Entries, 1)


	validator, _ = app.StakingKeeper.GetValidatorByConsAddr(ctx, consAddr)
	require.Equal(t, validator.GetStatus(), types.Unbonding)
}


func TestSlashBoth(t *testing.T) {
	app, ctx, addrDels, addrVals := bootstrapSlashTest(t, 10)
	fraction := sdk.NewDecWithPrec(5, 1)
	bondDenom := app.StakingKeeper.BondDenom(ctx)



	rdATokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 6)
	rdA := types.NewRedelegation(addrDels[0], addrVals[0], addrVals[1], 11,
		time.Unix(0, 0), rdATokens,
		rdATokens.ToDec())
	app.StakingKeeper.SetRedelegation(ctx, rdA)


	delA := types.NewDelegation(addrDels[0], addrVals[1], rdATokens.ToDec())
	app.StakingKeeper.SetDelegation(ctx, delA)



	ubdATokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 4)
	ubdA := types.NewUnbondingDelegation(addrDels[0], addrVals[0], 11,
		time.Unix(0, 0), ubdATokens)
	app.StakingKeeper.SetUnbondingDelegation(ctx, ubdA)

	bondedCoins := sdk.NewCoins(sdk.NewCoin(bondDenom, rdATokens.MulRaw(2)))
	notBondedCoins := sdk.NewCoins(sdk.NewCoin(bondDenom, ubdATokens))


	bondedPool := app.StakingKeeper.GetBondedPool(ctx)
	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, bondedPool.GetName(), bondedCoins))
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), notBondedCoins))

	app.AccountKeeper.SetModuleAccount(ctx, bondedPool)
	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)

	oldBonded := app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount
	oldNotBonded := app.BankKeeper.GetBalance(ctx, notBondedPool.GetAddress(), bondDenom).Amount

	ctx = ctx.WithBlockHeight(12)
	validator, found := app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(PKs[0]))
	require.True(t, found)
	consAddr0 := sdk.ConsAddress(PKs[0].Address())
	app.StakingKeeper.Slash(ctx, consAddr0, 10, 10, fraction)

	burnedNotBondedAmount := fraction.MulInt(ubdATokens).TruncateInt()
	burnedBondAmount := app.StakingKeeper.TokensFromConsensusPower(ctx, 10).ToDec().Mul(fraction).TruncateInt()
	burnedBondAmount = burnedBondAmount.Sub(burnedNotBondedAmount)


	bondedPool = app.StakingKeeper.GetBondedPool(ctx)
	notBondedPool = app.StakingKeeper.GetNotBondedPool(ctx)

	bondedPoolBalance := app.BankKeeper.GetBalance(ctx, bondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldBonded.Sub(burnedBondAmount), bondedPoolBalance))

	notBondedPoolBalance := app.BankKeeper.GetBalance(ctx, notBondedPool.GetAddress(), bondDenom).Amount
	require.True(sdk.IntEq(t, oldNotBonded.Sub(burnedNotBondedAmount), notBondedPoolBalance))


	rdA, found = app.StakingKeeper.GetRedelegation(ctx, addrDels[0], addrVals[0], addrVals[1])
	require.True(t, found)
	require.Len(t, rdA.Entries, 1)

	validator, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(PKs[0]))
	require.True(t, found)

	require.Equal(t, int64(10), validator.GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))
}
