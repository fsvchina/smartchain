package keeper_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

func newMonikerValidator(t testing.TB, operator sdk.ValAddress, pubKey cryptotypes.PubKey, moniker string) types.Validator {
	v, err := types.NewValidator(operator, pubKey, types.Description{Moniker: moniker})
	require.NoError(t, err)
	return v
}

func bootstrapValidatorTest(t testing.TB, power int64, numAddrs int) (*simapp.SimApp, sdk.Context, []sdk.AccAddress, []sdk.ValAddress) {
	_, app, ctx := createTestInput()

	addrDels, addrVals := generateAddresses(app, ctx, numAddrs)

	amt := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
	totalSupply := sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), amt.MulRaw(int64(len(addrDels)))))

	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)

	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), totalSupply))

	return app, ctx, addrDels, addrVals
}

func initValidators(t testing.TB, power int64, numAddrs int, powers []int64) (*simapp.SimApp, sdk.Context, []sdk.AccAddress, []sdk.ValAddress, []types.Validator) {
	app, ctx, addrs, valAddrs := bootstrapValidatorTest(t, power, numAddrs)
	pks := simapp.CreateTestPubKeys(numAddrs)

	vs := make([]types.Validator, len(powers))
	for i, power := range powers {
		vs[i] = teststaking.NewValidator(t, sdk.ValAddress(addrs[i]), pks[i])
		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		vs[i], _ = vs[i].AddTokensFromDel(tokens)
	}
	return app, ctx, addrs, valAddrs, vs
}

func TestSetValidator(t *testing.T) {
	app, ctx, _, _ := bootstrapValidatorTest(t, 10, 100)

	valPubKey := PKs[0]
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())
	valTokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 10)


	validator := teststaking.NewValidator(t, valAddr, valPubKey)
	validator, _ = validator.AddTokensFromDel(valTokens)
	require.Equal(t, types.Unbonded, validator.Status)
	assert.Equal(t, valTokens, validator.Tokens)
	assert.Equal(t, valTokens, validator.DelegatorShares.RoundInt())
	app.StakingKeeper.SetValidator(ctx, validator)
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validator)


	updates := applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 1)
	validator, found := app.StakingKeeper.GetValidator(ctx, valAddr)
	require.True(t, found)
	require.Equal(t, validator.ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])


	require.Equal(t, types.Bonded, validator.Status)
	assert.Equal(t, valTokens, validator.Tokens)
	assert.Equal(t, valTokens, validator.DelegatorShares.RoundInt())


	resVal, found := app.StakingKeeper.GetValidator(ctx, valAddr)
	assert.True(ValEq(t, validator, resVal))
	require.True(t, found)

	resVals := app.StakingKeeper.GetLastValidators(ctx)
	require.Equal(t, 1, len(resVals))
	assert.True(ValEq(t, validator, resVals[0]))

	resVals = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, 1, len(resVals))
	require.True(ValEq(t, validator, resVals[0]))

	resVals = app.StakingKeeper.GetValidators(ctx, 1)
	require.Equal(t, 1, len(resVals))
	require.True(ValEq(t, validator, resVals[0]))

	resVals = app.StakingKeeper.GetValidators(ctx, 10)
	require.Equal(t, 1, len(resVals))
	require.True(ValEq(t, validator, resVals[0]))

	allVals := app.StakingKeeper.GetAllValidators(ctx)
	require.Equal(t, 1, len(allVals))
}

func TestUpdateValidatorByPowerIndex(t *testing.T) {
	app, ctx, _, _ := bootstrapValidatorTest(t, 0, 100)
	_, addrVals := generateAddresses(app, ctx, 1)

	bondedPool := app.StakingKeeper.GetBondedPool(ctx)
	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, bondedPool.GetName(), sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), app.StakingKeeper.TokensFromConsensusPower(ctx, 1234)))))
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), app.StakingKeeper.TokensFromConsensusPower(ctx, 10000)))))

	app.AccountKeeper.SetModuleAccount(ctx, bondedPool)
	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)


	validator := teststaking.NewValidator(t, addrVals[0], PKs[0])
	validator, delSharesCreated := validator.AddTokensFromDel(app.StakingKeeper.TokensFromConsensusPower(ctx, 100))
	require.Equal(t, types.Unbonded, validator.Status)
	require.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 100), validator.Tokens)
	keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validator, true)
	validator, found := app.StakingKeeper.GetValidator(ctx, addrVals[0])
	require.True(t, found)
	require.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 100), validator.Tokens)

	power := types.GetValidatorsByPowerIndexKey(validator, app.StakingKeeper.PowerReduction(ctx))
	require.True(t, keeper.ValidatorByPowerIndexExists(ctx, app.StakingKeeper, power))


	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validator)
	validator, burned := validator.RemoveDelShares(delSharesCreated.Quo(sdk.NewDec(2)))
	require.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 50), burned)
	keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validator, true)
	require.False(t, keeper.ValidatorByPowerIndexExists(ctx, app.StakingKeeper, power))

	validator, found = app.StakingKeeper.GetValidator(ctx, addrVals[0])
	require.True(t, found)

	power = types.GetValidatorsByPowerIndexKey(validator, app.StakingKeeper.PowerReduction(ctx))
	require.True(t, keeper.ValidatorByPowerIndexExists(ctx, app.StakingKeeper, power))
}

func TestUpdateBondedValidatorsDecreaseCliff(t *testing.T) {
	numVals := 10
	maxVals := 5


	app, ctx, _, valAddrs := bootstrapValidatorTest(t, 0, 100)

	bondedPool := app.StakingKeeper.GetBondedPool(ctx)
	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)


	params := app.StakingKeeper.GetParams(ctx)
	params.MaxValidators = uint32(maxVals)
	app.StakingKeeper.SetParams(ctx, params)


	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, bondedPool.GetName(), sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), app.StakingKeeper.TokensFromConsensusPower(ctx, 1234)))))
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), app.StakingKeeper.TokensFromConsensusPower(ctx, 10000)))))

	app.AccountKeeper.SetModuleAccount(ctx, bondedPool)
	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)

	validators := make([]types.Validator, numVals)
	for i := 0; i < len(validators); i++ {
		moniker := fmt.Sprintf("val#%d", int64(i))
		val := newMonikerValidator(t, valAddrs[i], PKs[i], moniker)
		delTokens := app.StakingKeeper.TokensFromConsensusPower(ctx, int64((i+1)*10))
		val, _ = val.AddTokensFromDel(delTokens)

		val = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, val, true)
		validators[i] = val
	}

	nextCliffVal := validators[numVals-maxVals+1]



	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, nextCliffVal)
	shares := app.StakingKeeper.TokensFromConsensusPower(ctx, 21)
	nextCliffVal, _ = nextCliffVal.RemoveDelShares(shares.ToDec())
	nextCliffVal = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, nextCliffVal, true)

	expectedValStatus := map[int]types.BondStatus{
		9: types.Bonded, 8: types.Bonded, 7: types.Bonded, 5: types.Bonded, 4: types.Bonded,
		0: types.Unbonding, 1: types.Unbonding, 2: types.Unbonding, 3: types.Unbonding, 6: types.Unbonding,
	}


	for valIdx, status := range expectedValStatus {
		valAddr := validators[valIdx].OperatorAddress
		addr, err := sdk.ValAddressFromBech32(valAddr)
		assert.NoError(t, err)
		val, _ := app.StakingKeeper.GetValidator(ctx, addr)

		assert.Equal(
			t, status, val.GetStatus(),
			fmt.Sprintf("expected validator at index %v to have status: %s", valIdx, status),
		)
	}
}

func TestSlashToZeroPowerRemoved(t *testing.T) {

	app, ctx, _, addrVals := bootstrapValidatorTest(t, 100, 20)


	validator := teststaking.NewValidator(t, addrVals[0], PKs[0])
	valTokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 100)

	bondedPool := app.StakingKeeper.GetBondedPool(ctx)

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, bondedPool.GetName(), sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), valTokens))))

	app.AccountKeeper.SetModuleAccount(ctx, bondedPool)

	validator, _ = validator.AddTokensFromDel(valTokens)
	require.Equal(t, types.Unbonded, validator.Status)
	require.Equal(t, valTokens, validator.Tokens)
	app.StakingKeeper.SetValidatorByConsAddr(ctx, validator)
	validator = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validator, true)
	require.Equal(t, valTokens, validator.Tokens, "\nvalidator %v\npool %v", validator, valTokens)


	app.StakingKeeper.Slash(ctx, sdk.ConsAddress(PKs[0].Address()), 0, 100, sdk.OneDec())

	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, -1)

	validator, _ = app.StakingKeeper.GetValidator(ctx, addrVals[0])
	require.Equal(t, validator.GetStatus(), types.Unbonding)
}


func TestValidatorBasics(t *testing.T) {
	app, ctx, _, addrVals := bootstrapValidatorTest(t, 1000, 20)


	var validators [3]types.Validator
	powers := []int64{9, 8, 7}
	for i, power := range powers {
		validators[i] = teststaking.NewValidator(t, addrVals[i], PKs[i])
		validators[i].Status = types.Unbonded
		validators[i].Tokens = sdk.ZeroInt()
		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)

		validators[i], _ = validators[i].AddTokensFromDel(tokens)
	}
	assert.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 9), validators[0].Tokens)
	assert.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 8), validators[1].Tokens)
	assert.Equal(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 7), validators[2].Tokens)


	_, found := app.StakingKeeper.GetValidator(ctx, addrVals[0])
	require.False(t, found)
	resVals := app.StakingKeeper.GetLastValidators(ctx)
	require.Zero(t, len(resVals))

	resVals = app.StakingKeeper.GetValidators(ctx, 2)
	require.Zero(t, len(resVals))


	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], true)
	app.StakingKeeper.SetValidatorByConsAddr(ctx, validators[0])
	resVal, found := app.StakingKeeper.GetValidator(ctx, addrVals[0])
	require.True(t, found)
	assert.True(ValEq(t, validators[0], resVal))


	resVal, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.ConsAddress(PKs[0].Address()))
	require.True(t, found)
	assert.True(ValEq(t, validators[0], resVal))
	resVal, found = app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(PKs[0]))
	require.True(t, found)
	assert.True(ValEq(t, validators[0], resVal))

	resVals = app.StakingKeeper.GetLastValidators(ctx)
	require.Equal(t, 1, len(resVals))
	assert.True(ValEq(t, validators[0], resVals[0]))
	assert.Equal(t, types.Bonded, validators[0].Status)
	assert.True(sdk.IntEq(t, app.StakingKeeper.TokensFromConsensusPower(ctx, 9), validators[0].BondedTokens()))


	validators[0].Status = types.Bonded
	validators[0].Tokens = app.StakingKeeper.TokensFromConsensusPower(ctx, 10)
	validators[0].DelegatorShares = validators[0].Tokens.ToDec()
	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], true)
	resVal, found = app.StakingKeeper.GetValidator(ctx, addrVals[0])
	require.True(t, found)
	assert.True(ValEq(t, validators[0], resVal))

	resVals = app.StakingKeeper.GetLastValidators(ctx)
	require.Equal(t, 1, len(resVals))
	assert.True(ValEq(t, validators[0], resVals[0]))


	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], true)
	validators[2] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[2], true)
	resVal, found = app.StakingKeeper.GetValidator(ctx, addrVals[1])
	require.True(t, found)
	assert.True(ValEq(t, validators[1], resVal))
	resVal, found = app.StakingKeeper.GetValidator(ctx, addrVals[2])
	require.True(t, found)
	assert.True(ValEq(t, validators[2], resVal))

	resVals = app.StakingKeeper.GetLastValidators(ctx)
	require.Equal(t, 3, len(resVals))
	assert.True(ValEq(t, validators[0], resVals[0]))
	assert.True(ValEq(t, validators[1], resVals[1]))
	assert.True(ValEq(t, validators[2], resVals[2]))




	assert.PanicsWithValue(t,
		"cannot call RemoveValidator on bonded or unbonding validators",
		func() { app.StakingKeeper.RemoveValidator(ctx, validators[1].GetOperator()) })


	validators[1].Status = types.Unbonded
	app.StakingKeeper.SetValidator(ctx, validators[1])
	assert.PanicsWithValue(t,
		"attempting to remove a validator which still contains tokens",
		func() { app.StakingKeeper.RemoveValidator(ctx, validators[1].GetOperator()) })

	validators[1].Tokens = sdk.ZeroInt()
	app.StakingKeeper.SetValidator(ctx, validators[1])
	app.StakingKeeper.RemoveValidator(ctx, validators[1].GetOperator())
	_, found = app.StakingKeeper.GetValidator(ctx, addrVals[1])
	require.False(t, found)
}


func TestGetValidatorSortingUnmixed(t *testing.T) {
	app, ctx, addrs, _ := bootstrapValidatorTest(t, 1000, 20)


	amts := []sdk.Int{
		sdk.NewIntFromUint64(0),
		app.StakingKeeper.PowerReduction(ctx).MulRaw(100),
		app.StakingKeeper.PowerReduction(ctx),
		app.StakingKeeper.PowerReduction(ctx).MulRaw(400),
		app.StakingKeeper.PowerReduction(ctx).MulRaw(200)}
	n := len(amts)
	var validators [5]types.Validator
	for i, amt := range amts {
		validators[i] = teststaking.NewValidator(t, sdk.ValAddress(addrs[i]), PKs[i])
		validators[i].Status = types.Bonded
		validators[i].Tokens = amt
		validators[i].DelegatorShares = sdk.NewDecFromInt(amt)
		keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[i], true)
	}


	resValidators := app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	assert.Equal(t, n, len(resValidators))
	assert.Equal(t, sdk.NewInt(400).Mul(app.StakingKeeper.PowerReduction(ctx)), resValidators[0].BondedTokens(), "%v", resValidators)
	assert.Equal(t, sdk.NewInt(200).Mul(app.StakingKeeper.PowerReduction(ctx)), resValidators[1].BondedTokens(), "%v", resValidators)
	assert.Equal(t, sdk.NewInt(100).Mul(app.StakingKeeper.PowerReduction(ctx)), resValidators[2].BondedTokens(), "%v", resValidators)
	assert.Equal(t, sdk.NewInt(1).Mul(app.StakingKeeper.PowerReduction(ctx)), resValidators[3].BondedTokens(), "%v", resValidators)
	assert.Equal(t, sdk.NewInt(0), resValidators[4].BondedTokens(), "%v", resValidators)
	assert.Equal(t, validators[3].OperatorAddress, resValidators[0].OperatorAddress, "%v", resValidators)
	assert.Equal(t, validators[4].OperatorAddress, resValidators[1].OperatorAddress, "%v", resValidators)
	assert.Equal(t, validators[1].OperatorAddress, resValidators[2].OperatorAddress, "%v", resValidators)
	assert.Equal(t, validators[2].OperatorAddress, resValidators[3].OperatorAddress, "%v", resValidators)
	assert.Equal(t, validators[0].OperatorAddress, resValidators[4].OperatorAddress, "%v", resValidators)


	validators[3].Tokens = sdk.NewInt(500).Mul(app.StakingKeeper.PowerReduction(ctx))
	keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[3], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, len(resValidators), n)
	assert.True(ValEq(t, validators[3], resValidators[0]))


	validators[3].Tokens = sdk.NewInt(300).Mul(app.StakingKeeper.PowerReduction(ctx))
	keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[3], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, len(resValidators), n)
	assert.True(ValEq(t, validators[3], resValidators[0]))
	assert.True(ValEq(t, validators[4], resValidators[1]))


	validators[3].Tokens = sdk.NewInt(200).Mul(app.StakingKeeper.PowerReduction(ctx))
	ctx = ctx.WithBlockHeight(10)
	keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[3], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, len(resValidators), n)
	assert.True(ValEq(t, validators[3], resValidators[0]))
	assert.True(ValEq(t, validators[4], resValidators[1]))


	ctx = ctx.WithBlockHeight(20)
	keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[4], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, len(resValidators), n)
	assert.True(ValEq(t, validators[3], resValidators[0]))
	assert.True(ValEq(t, validators[4], resValidators[1]))


	validators[3].Tokens = sdk.NewInt(300).Mul(app.StakingKeeper.PowerReduction(ctx))
	validators[4].Tokens = sdk.NewInt(300).Mul(app.StakingKeeper.PowerReduction(ctx))
	keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[3], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, len(resValidators), n)
	ctx = ctx.WithBlockHeight(30)
	keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[4], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, len(resValidators), n, "%v", resValidators)
	assert.True(ValEq(t, validators[3], resValidators[0]))
	assert.True(ValEq(t, validators[4], resValidators[1]))
}

func TestGetValidatorSortingMixed(t *testing.T) {
	app, ctx, addrs, _ := bootstrapValidatorTest(t, 1000, 20)
	bondedPool := app.StakingKeeper.GetBondedPool(ctx)
	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, bondedPool.GetName(), sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), app.StakingKeeper.TokensFromConsensusPower(ctx, 501)))))
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), app.StakingKeeper.TokensFromConsensusPower(ctx, 0)))))

	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)
	app.AccountKeeper.SetModuleAccount(ctx, bondedPool)


	params := app.StakingKeeper.GetParams(ctx)
	params.MaxValidators = 2
	app.StakingKeeper.SetParams(ctx, params)


	amts := []sdk.Int{
		sdk.NewIntFromUint64(0),
		app.StakingKeeper.PowerReduction(ctx).MulRaw(100),
		app.StakingKeeper.PowerReduction(ctx),
		app.StakingKeeper.PowerReduction(ctx).MulRaw(400),
		app.StakingKeeper.PowerReduction(ctx).MulRaw(200)}

	var validators [5]types.Validator
	for i, amt := range amts {
		validators[i] = teststaking.NewValidator(t, sdk.ValAddress(addrs[i]), PKs[i])
		validators[i].DelegatorShares = sdk.NewDecFromInt(amt)
		validators[i].Status = types.Bonded
		validators[i].Tokens = amt
		keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[i], true)
	}

	val0, found := app.StakingKeeper.GetValidator(ctx, sdk.ValAddress(addrs[0]))
	require.True(t, found)
	val1, found := app.StakingKeeper.GetValidator(ctx, sdk.ValAddress(addrs[1]))
	require.True(t, found)
	val2, found := app.StakingKeeper.GetValidator(ctx, sdk.ValAddress(addrs[2]))
	require.True(t, found)
	val3, found := app.StakingKeeper.GetValidator(ctx, sdk.ValAddress(addrs[3]))
	require.True(t, found)
	val4, found := app.StakingKeeper.GetValidator(ctx, sdk.ValAddress(addrs[4]))
	require.True(t, found)
	require.Equal(t, types.Bonded, val0.Status)
	require.Equal(t, types.Unbonding, val1.Status)
	require.Equal(t, types.Unbonding, val2.Status)
	require.Equal(t, types.Bonded, val3.Status)
	require.Equal(t, types.Bonded, val4.Status)


	resValidators := app.StakingKeeper.GetBondedValidatorsByPower(ctx)

	assert.Equal(t, 2, len(resValidators))
	assert.Equal(t, sdk.NewInt(400).Mul(app.StakingKeeper.PowerReduction(ctx)), resValidators[0].BondedTokens(), "%v", resValidators)
	assert.Equal(t, sdk.NewInt(200).Mul(app.StakingKeeper.PowerReduction(ctx)), resValidators[1].BondedTokens(), "%v", resValidators)
	assert.Equal(t, validators[3].OperatorAddress, resValidators[0].OperatorAddress, "%v", resValidators)
	assert.Equal(t, validators[4].OperatorAddress, resValidators[1].OperatorAddress, "%v", resValidators)
}


func TestGetValidatorsEdgeCases(t *testing.T) {
	app, ctx, addrs, _ := bootstrapValidatorTest(t, 1000, 20)


	params := app.StakingKeeper.GetParams(ctx)
	nMax := uint32(2)
	params.MaxValidators = nMax
	app.StakingKeeper.SetParams(ctx, params)


	powers := []int64{0, 100, 400, 400}
	var validators [4]types.Validator
	for i, power := range powers {
		moniker := fmt.Sprintf("val#%d", int64(i))
		validators[i] = newMonikerValidator(t, sdk.ValAddress(addrs[i]), PKs[i], moniker)

		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		validators[i], _ = validators[i].AddTokensFromDel(tokens)

		notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)
		require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), sdk.NewCoins(sdk.NewCoin(params.BondDenom, tokens))))
		app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)
		validators[i] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[i], true)
	}


	resValidators := app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, nMax, uint32(len(resValidators)))
	assert.True(ValEq(t, validators[2], resValidators[0]))
	assert.True(ValEq(t, validators[3], resValidators[1]))


	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validators[0])
	delTokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 500)
	validators[0], _ = validators[0].AddTokensFromDel(delTokens)
	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)

	newTokens := sdk.NewCoins()

	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), newTokens))
	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)




	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, nMax, uint32(len(resValidators)))
	assert.True(ValEq(t, validators[0], resValidators[0]))
	assert.True(ValEq(t, validators[2], resValidators[1]))




	//






	ctx = ctx.WithBlockHeight(40)

	var found bool
	validators[3], found = app.StakingKeeper.GetValidator(ctx, validators[3].GetOperator())
	assert.True(t, found)
	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validators[3])
	validators[3], _ = validators[3].AddTokensFromDel(app.StakingKeeper.TokensFromConsensusPower(ctx, 1))

	notBondedPool = app.StakingKeeper.GetNotBondedPool(ctx)
	newTokens = sdk.NewCoins(sdk.NewCoin(params.BondDenom, app.StakingKeeper.TokensFromConsensusPower(ctx, 1)))
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), newTokens))
	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)

	validators[3] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[3], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, nMax, uint32(len(resValidators)))
	assert.True(ValEq(t, validators[0], resValidators[0]))
	assert.True(ValEq(t, validators[3], resValidators[1]))


	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validators[3])
	rmTokens := validators[3].TokensFromShares(sdk.NewDec(201)).TruncateInt()
	validators[3], _ = validators[3].RemoveDelShares(sdk.NewDec(201))

	bondedPool := app.StakingKeeper.GetBondedPool(ctx)
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, bondedPool.GetName(), sdk.NewCoins(sdk.NewCoin(params.BondDenom, rmTokens))))
	app.AccountKeeper.SetModuleAccount(ctx, bondedPool)

	validators[3] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[3], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, nMax, uint32(len(resValidators)))
	assert.True(ValEq(t, validators[0], resValidators[0]))
	assert.True(ValEq(t, validators[2], resValidators[1]))


	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validators[3])
	validators[3], _ = validators[3].AddTokensFromDel(sdk.NewInt(200))

	notBondedPool = app.StakingKeeper.GetNotBondedPool(ctx)
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), sdk.NewCoins(sdk.NewCoin(params.BondDenom, sdk.NewInt(200)))))
	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)

	validators[3] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[3], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, nMax, uint32(len(resValidators)))
	assert.True(ValEq(t, validators[0], resValidators[0]))
	assert.True(ValEq(t, validators[2], resValidators[1]))
	_, exists := app.StakingKeeper.GetValidator(ctx, validators[3].GetOperator())
	require.True(t, exists)
}

func TestValidatorBondHeight(t *testing.T) {
	app, ctx, addrs, _ := bootstrapValidatorTest(t, 1000, 20)


	params := app.StakingKeeper.GetParams(ctx)
	params.MaxValidators = 2
	app.StakingKeeper.SetParams(ctx, params)


	var validators [3]types.Validator
	validators[0] = teststaking.NewValidator(t, sdk.ValAddress(PKs[0].Address().Bytes()), PKs[0])
	validators[1] = teststaking.NewValidator(t, sdk.ValAddress(addrs[1]), PKs[1])
	validators[2] = teststaking.NewValidator(t, sdk.ValAddress(addrs[2]), PKs[2])

	tokens0 := app.StakingKeeper.TokensFromConsensusPower(ctx, 200)
	tokens1 := app.StakingKeeper.TokensFromConsensusPower(ctx, 100)
	tokens2 := app.StakingKeeper.TokensFromConsensusPower(ctx, 100)
	validators[0], _ = validators[0].AddTokensFromDel(tokens0)
	validators[1], _ = validators[1].AddTokensFromDel(tokens1)
	validators[2], _ = validators[2].AddTokensFromDel(tokens2)

	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], true)




	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], true)
	validators[2] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[2], true)

	resValidators := app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, uint32(len(resValidators)), params.MaxValidators)

	assert.True(ValEq(t, validators[0], resValidators[0]))
	assert.True(ValEq(t, validators[1], resValidators[1]))
	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validators[1])
	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validators[2])
	delTokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 50)
	validators[1], _ = validators[1].AddTokensFromDel(delTokens)
	validators[2], _ = validators[2].AddTokensFromDel(delTokens)
	validators[2] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[2], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	require.Equal(t, params.MaxValidators, uint32(len(resValidators)))
	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], true)
	assert.True(ValEq(t, validators[0], resValidators[0]))
	assert.True(ValEq(t, validators[2], resValidators[1]))
}

func TestFullValidatorSetPowerChange(t *testing.T) {
	app, ctx, addrs, _ := bootstrapValidatorTest(t, 1000, 20)
	params := app.StakingKeeper.GetParams(ctx)
	max := 2
	params.MaxValidators = uint32(2)
	app.StakingKeeper.SetParams(ctx, params)


	powers := []int64{0, 100, 400, 400, 200}
	var validators [5]types.Validator
	for i, power := range powers {
		validators[i] = teststaking.NewValidator(t, sdk.ValAddress(addrs[i]), PKs[i])
		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		validators[i], _ = validators[i].AddTokensFromDel(tokens)
		keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[i], true)
	}
	for i := range powers {
		var found bool
		validators[i], found = app.StakingKeeper.GetValidator(ctx, validators[i].GetOperator())
		require.True(t, found)
	}
	assert.Equal(t, types.Unbonded, validators[0].Status)
	assert.Equal(t, types.Unbonding, validators[1].Status)
	assert.Equal(t, types.Bonded, validators[2].Status)
	assert.Equal(t, types.Bonded, validators[3].Status)
	assert.Equal(t, types.Unbonded, validators[4].Status)
	resValidators := app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	assert.Equal(t, max, len(resValidators))
	assert.True(ValEq(t, validators[2], resValidators[0]))
	assert.True(ValEq(t, validators[3], resValidators[1]))



	tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 600)
	validators[0], _ = validators[0].AddTokensFromDel(tokens)
	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], true)
	resValidators = app.StakingKeeper.GetBondedValidatorsByPower(ctx)
	assert.Equal(t, max, len(resValidators))
	assert.True(ValEq(t, validators[0], resValidators[0]))
	assert.True(ValEq(t, validators[2], resValidators[1]))
}

func TestApplyAndReturnValidatorSetUpdatesAllNone(t *testing.T) {
	app, ctx, _, _ := bootstrapValidatorTest(t, 1000, 20)

	powers := []int64{10, 20}
	var validators [2]types.Validator
	for i, power := range powers {
		valPubKey := PKs[i+1]
		valAddr := sdk.ValAddress(valPubKey.Address().Bytes())

		validators[i] = teststaking.NewValidator(t, valAddr, valPubKey)
		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		validators[i], _ = validators[i].AddTokensFromDel(tokens)
	}



	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 0)
	app.StakingKeeper.SetValidator(ctx, validators[0])
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[0])
	app.StakingKeeper.SetValidator(ctx, validators[1])
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[1])

	updates := applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)
	validators[0], _ = app.StakingKeeper.GetValidator(ctx, validators[0].GetOperator())
	validators[1], _ = app.StakingKeeper.GetValidator(ctx, validators[1].GetOperator())
	assert.Equal(t, validators[0].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[1])
	assert.Equal(t, validators[1].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])
}

func TestApplyAndReturnValidatorSetUpdatesIdentical(t *testing.T) {
	app, ctx, addrs, _ := bootstrapValidatorTest(t, 1000, 20)

	powers := []int64{10, 20}
	var validators [2]types.Validator
	for i, power := range powers {
		validators[i] = teststaking.NewValidator(t, sdk.ValAddress(addrs[i]), PKs[i])

		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		validators[i], _ = validators[i].AddTokensFromDel(tokens)

	}
	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], false)
	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], false)
	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)



	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], false)
	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], false)
	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 0)
}

func TestApplyAndReturnValidatorSetUpdatesSingleValueChange(t *testing.T) {
	app, ctx, addrs, _ := bootstrapValidatorTest(t, 1000, 20)

	powers := []int64{10, 20}
	var validators [2]types.Validator
	for i, power := range powers {
		validators[i] = teststaking.NewValidator(t, sdk.ValAddress(addrs[i]), PKs[i])

		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		validators[i], _ = validators[i].AddTokensFromDel(tokens)

	}
	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], false)
	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], false)
	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)



	validators[0].Status = types.Bonded
	validators[0].Tokens = app.StakingKeeper.TokensFromConsensusPower(ctx, 600)
	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], false)

	updates := applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 1)
	require.Equal(t, validators[0].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])
}

func TestApplyAndReturnValidatorSetUpdatesMultipleValueChange(t *testing.T) {
	powers := []int64{10, 20}

	app, ctx, _, _, validators := initValidators(t, 1000, 20, powers)

	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], false)
	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], false)
	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)



	delTokens1 := app.StakingKeeper.TokensFromConsensusPower(ctx, 190)
	delTokens2 := app.StakingKeeper.TokensFromConsensusPower(ctx, 80)
	validators[0], _ = validators[0].AddTokensFromDel(delTokens1)
	validators[1], _ = validators[1].AddTokensFromDel(delTokens2)
	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], false)
	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], false)

	updates := applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)
	require.Equal(t, validators[0].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])
	require.Equal(t, validators[1].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[1])
}

func TestApplyAndReturnValidatorSetUpdatesInserted(t *testing.T) {
	powers := []int64{10, 20, 5, 15, 25}
	app, ctx, _, _, validators := initValidators(t, 1000, 20, powers)

	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], false)
	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], false)
	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)



	app.StakingKeeper.SetValidator(ctx, validators[2])
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[2])
	updates := applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 1)
	validators[2], _ = app.StakingKeeper.GetValidator(ctx, validators[2].GetOperator())
	require.Equal(t, validators[2].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])



	app.StakingKeeper.SetValidator(ctx, validators[3])
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[3])
	updates = applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 1)
	validators[3], _ = app.StakingKeeper.GetValidator(ctx, validators[3].GetOperator())
	require.Equal(t, validators[3].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])



	app.StakingKeeper.SetValidator(ctx, validators[4])
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[4])
	updates = applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 1)
	validators[4], _ = app.StakingKeeper.GetValidator(ctx, validators[4].GetOperator())
	require.Equal(t, validators[4].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])
}

func TestApplyAndReturnValidatorSetUpdatesWithCliffValidator(t *testing.T) {
	app, ctx, addrs, _ := bootstrapValidatorTest(t, 1000, 20)
	params := types.DefaultParams()
	params.MaxValidators = 2
	app.StakingKeeper.SetParams(ctx, params)

	powers := []int64{10, 20, 5}
	var validators [5]types.Validator
	for i, power := range powers {
		validators[i] = teststaking.NewValidator(t, sdk.ValAddress(addrs[i]), PKs[i])
		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		validators[i], _ = validators[i].AddTokensFromDel(tokens)
	}
	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], false)
	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], false)
	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)



	keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[2], false)
	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 0)



	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 0)

	tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 10)
	validators[2], _ = validators[2].AddTokensFromDel(tokens)
	app.StakingKeeper.SetValidator(ctx, validators[2])
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[2])
	updates := applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)
	validators[2], _ = app.StakingKeeper.GetValidator(ctx, validators[2].GetOperator())
	require.Equal(t, validators[0].ABCIValidatorUpdateZero(), updates[1])
	require.Equal(t, validators[2].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])
}

func TestApplyAndReturnValidatorSetUpdatesPowerDecrease(t *testing.T) {
	app, ctx, addrs, _ := bootstrapValidatorTest(t, 1000, 20)

	powers := []int64{100, 100}
	var validators [2]types.Validator
	for i, power := range powers {
		validators[i] = teststaking.NewValidator(t, sdk.ValAddress(addrs[i]), PKs[i])
		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		validators[i], _ = validators[i].AddTokensFromDel(tokens)
	}
	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], false)
	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], false)
	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)


	require.Equal(t, int64(100), validators[0].GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))
	require.Equal(t, int64(100), validators[1].GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))



	delTokens1 := app.StakingKeeper.TokensFromConsensusPower(ctx, 20)
	delTokens2 := app.StakingKeeper.TokensFromConsensusPower(ctx, 30)
	validators[0], _ = validators[0].RemoveDelShares(delTokens1.ToDec())
	validators[1], _ = validators[1].RemoveDelShares(delTokens2.ToDec())
	validators[0] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[0], false)
	validators[1] = keeper.TestingUpdateValidator(app.StakingKeeper, ctx, validators[1], false)


	require.Equal(t, int64(80), validators[0].GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))
	require.Equal(t, int64(70), validators[1].GetConsensusPower(app.StakingKeeper.PowerReduction(ctx)))


	updates := applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)
	require.Equal(t, validators[0].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])
	require.Equal(t, validators[1].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[1])
}

func TestApplyAndReturnValidatorSetUpdatesNewValidator(t *testing.T) {
	app, ctx, _, _ := bootstrapValidatorTest(t, 1000, 20)
	params := app.StakingKeeper.GetParams(ctx)
	params.MaxValidators = uint32(3)

	app.StakingKeeper.SetParams(ctx, params)

	powers := []int64{100, 100}
	var validators [2]types.Validator


	for i, power := range powers {
		valPubKey := PKs[i+1]
		valAddr := sdk.ValAddress(valPubKey.Address().Bytes())

		validators[i] = teststaking.NewValidator(t, valAddr, valPubKey)
		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		validators[i], _ = validators[i].AddTokensFromDel(tokens)

		app.StakingKeeper.SetValidator(ctx, validators[i])
		app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[i])
	}


	updates := applyValidatorSetUpdates(t, ctx, app.StakingKeeper, len(validators))
	validators[0], _ = app.StakingKeeper.GetValidator(ctx, validators[0].GetOperator())
	validators[1], _ = app.StakingKeeper.GetValidator(ctx, validators[1].GetOperator())
	require.Equal(t, validators[0].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])
	require.Equal(t, validators[1].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[1])

	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 0)


	for i, power := range powers {

		app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validators[i])
		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		validators[i], _ = validators[i].AddTokensFromDel(tokens)

		app.StakingKeeper.SetValidator(ctx, validators[i])
		app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[i])
	}



	valPubKey := PKs[len(validators)+1]
	valAddr := sdk.ValAddress(valPubKey.Address().Bytes())
	amt := sdk.NewInt(100)

	validator := teststaking.NewValidator(t, valAddr, valPubKey)
	validator, _ = validator.AddTokensFromDel(amt)

	app.StakingKeeper.SetValidator(ctx, validator)

	validator, _ = validator.RemoveDelShares(amt.ToDec())
	app.StakingKeeper.SetValidator(ctx, validator)
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validator)


	valPubKey = PKs[len(validators)+2]
	valAddr = sdk.ValAddress(valPubKey.Address().Bytes())

	validator = teststaking.NewValidator(t, valAddr, valPubKey)
	tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 500)
	validator, _ = validator.AddTokensFromDel(tokens)
	app.StakingKeeper.SetValidator(ctx, validator)
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validator)


	updates = applyValidatorSetUpdates(t, ctx, app.StakingKeeper, len(validators)+1)
	validator, _ = app.StakingKeeper.GetValidator(ctx, validator.GetOperator())
	validators[0], _ = app.StakingKeeper.GetValidator(ctx, validators[0].GetOperator())
	validators[1], _ = app.StakingKeeper.GetValidator(ctx, validators[1].GetOperator())
	require.Equal(t, validator.ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])
	require.Equal(t, validators[0].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[1])
	require.Equal(t, validators[1].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[2])
}

func TestApplyAndReturnValidatorSetUpdatesBondTransition(t *testing.T) {
	app, ctx, _, _ := bootstrapValidatorTest(t, 1000, 20)
	params := app.StakingKeeper.GetParams(ctx)
	params.MaxValidators = uint32(2)

	app.StakingKeeper.SetParams(ctx, params)

	powers := []int64{100, 200, 300}
	var validators [3]types.Validator


	for i, power := range powers {
		moniker := fmt.Sprintf("%d", i)
		valPubKey := PKs[i+1]
		valAddr := sdk.ValAddress(valPubKey.Address().Bytes())

		validators[i] = newMonikerValidator(t, valAddr, valPubKey, moniker)
		tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
		validators[i], _ = validators[i].AddTokensFromDel(tokens)
		app.StakingKeeper.SetValidator(ctx, validators[i])
		app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[i])
	}


	updates := applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 2)
	validators[2], _ = app.StakingKeeper.GetValidator(ctx, validators[2].GetOperator())
	validators[1], _ = app.StakingKeeper.GetValidator(ctx, validators[1].GetOperator())
	require.Equal(t, validators[2].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])
	require.Equal(t, validators[1].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[1])

	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 0)


	ctx = ctx.WithBlockHeight(1)

	var found bool
	validators[0], found = app.StakingKeeper.GetValidator(ctx, validators[0].GetOperator())
	require.True(t, found)

	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validators[0])
	tokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 1)
	validators[0], _ = validators[0].AddTokensFromDel(tokens)
	app.StakingKeeper.SetValidator(ctx, validators[0])
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[0])


	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 0)



	ctx = ctx.WithBlockHeight(2)

	validators[1], found = app.StakingKeeper.GetValidator(ctx, validators[1].GetOperator())
	require.True(t, found)

	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validators[0])
	validators[0], _ = validators[0].RemoveDelShares(validators[0].DelegatorShares)
	app.StakingKeeper.SetValidator(ctx, validators[0])
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[0])
	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 0)

	app.StakingKeeper.DeleteValidatorByPowerIndex(ctx, validators[1])
	tokens = app.StakingKeeper.TokensFromConsensusPower(ctx, 250)
	validators[1], _ = validators[1].AddTokensFromDel(tokens)
	app.StakingKeeper.SetValidator(ctx, validators[1])
	app.StakingKeeper.SetValidatorByPowerIndex(ctx, validators[1])


	updates = applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 1)
	require.Equal(t, validators[1].ABCIValidatorUpdate(app.StakingKeeper.PowerReduction(ctx)), updates[0])

	applyValidatorSetUpdates(t, ctx, app.StakingKeeper, 0)
}

func TestUpdateValidatorCommission(t *testing.T) {
	app, ctx, _, addrVals := bootstrapValidatorTest(t, 1000, 20)
	ctx = ctx.WithBlockHeader(tmproto.Header{Time: time.Now().UTC()})

	commission1 := types.NewCommissionWithTime(
		sdk.NewDecWithPrec(1, 1), sdk.NewDecWithPrec(3, 1),
		sdk.NewDecWithPrec(1, 1), time.Now().UTC().Add(time.Duration(-1)*time.Hour),
	)
	commission2 := types.NewCommission(sdk.NewDecWithPrec(1, 1), sdk.NewDecWithPrec(3, 1), sdk.NewDecWithPrec(1, 1))

	val1 := teststaking.NewValidator(t, addrVals[0], PKs[0])
	val2 := teststaking.NewValidator(t, addrVals[1], PKs[1])

	val1, _ = val1.SetInitialCommission(commission1)
	val2, _ = val2.SetInitialCommission(commission2)

	app.StakingKeeper.SetValidator(ctx, val1)
	app.StakingKeeper.SetValidator(ctx, val2)

	testCases := []struct {
		validator   types.Validator
		newRate     sdk.Dec
		expectedErr bool
	}{
		{val1, sdk.ZeroDec(), true},
		{val2, sdk.NewDecWithPrec(-1, 1), true},
		{val2, sdk.NewDecWithPrec(4, 1), true},
		{val2, sdk.NewDecWithPrec(3, 1), true},
		{val2, sdk.NewDecWithPrec(2, 1), false},
	}

	for i, tc := range testCases {
		commission, err := app.StakingKeeper.UpdateValidatorCommission(ctx, tc.validator, tc.newRate)

		if tc.expectedErr {
			require.Error(t, err, "expected error for test case #%d with rate: %s", i, tc.newRate)
		} else {
			tc.validator.Commission = commission
			app.StakingKeeper.SetValidator(ctx, tc.validator)
			val, found := app.StakingKeeper.GetValidator(ctx, tc.validator.GetOperator())

			require.True(t, found,
				"expected to find validator for test case #%d with rate: %s", i, tc.newRate,
			)
			require.NoError(t, err,
				"unexpected error for test case #%d with rate: %s", i, tc.newRate,
			)
			require.Equal(t, tc.newRate, val.Commission.Rate,
				"expected new validator commission rate for test case #%d with rate: %s", i, tc.newRate,
			)
			require.Equal(t, ctx.BlockHeader().Time, val.Commission.UpdateTime,
				"expected new validator commission update time for test case #%d with rate: %s", i, tc.newRate,
			)
		}
	}
}

func applyValidatorSetUpdates(t *testing.T, ctx sdk.Context, k keeper.Keeper, expectedUpdatesLen int) []abci.ValidatorUpdate {
	updates, err := k.ApplyAndReturnValidatorSetUpdates(ctx)
	require.NoError(t, err)
	if expectedUpdatesLen >= 0 {
		require.Equal(t, expectedUpdatesLen, len(updates), "%v", updates)
	}
	return updates
}
