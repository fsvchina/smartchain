package staking_test

import (
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"
	tmtypes "github.com/tendermint/tendermint/types"

	"github.com/golang/protobuf/proto"

	cryptocodec "github.com/cosmos/cosmos-sdk/crypto/codec"
	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/keeper"
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)

func bootstrapHandlerGenesisTest(t *testing.T, power int64, numAddrs int, accAmount sdk.Int) (*simapp.SimApp, sdk.Context, []sdk.AccAddress, []sdk.ValAddress) {
	_, app, ctx := getBaseSimappWithCustomKeeper()

	addrDels, addrVals := generateAddresses(app, ctx, numAddrs, accAmount)

	amt := app.StakingKeeper.TokensFromConsensusPower(ctx, power)
	totalSupply := sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.BondDenom(ctx), amt.MulRaw(int64(len(addrDels)))))

	notBondedPool := app.StakingKeeper.GetNotBondedPool(ctx)


	app.AccountKeeper.SetModuleAccount(ctx, notBondedPool)
	require.NoError(t, simapp.FundModuleAccount(app.BankKeeper, ctx, notBondedPool.GetName(), totalSupply))
	return app, ctx, addrDels, addrVals
}

func TestValidatorByPowerIndex(t *testing.T) {
	initPower := int64(1000000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 10, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	validatorAddr, validatorAddr3 := valAddrs[0], valAddrs[1]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	initBond := tstaking.CreateValidatorWithValPower(validatorAddr, PKs[0], initPower, true)


	updates, err := app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(updates))


	bond, found := app.StakingKeeper.GetDelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr)
	require.True(t, found)
	gotBond := bond.Shares.RoundInt()
	require.Equal(t, initBond, gotBond)


	validator, found := app.StakingKeeper.GetValidator(ctx, validatorAddr)
	require.True(t, found)
	power := types.GetValidatorsByPowerIndexKey(validator, app.StakingKeeper.PowerReduction(ctx))
	require.True(t, keeper.ValidatorByPowerIndexExists(ctx, app.StakingKeeper, power))


	tstaking.CreateValidatorWithValPower(validatorAddr3, PKs[2], initPower, true)


	updates, err = app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(updates))


	consAddr0 := sdk.ConsAddress(PKs[0].Address())
	app.StakingKeeper.Slash(ctx, consAddr0, 0, initPower, sdk.NewDecWithPrec(5, 1))
	app.StakingKeeper.Jail(ctx, consAddr0)
	app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)

	validator, found = app.StakingKeeper.GetValidator(ctx, validatorAddr)
	require.True(t, found)
	require.Equal(t, types.Unbonding, validator.Status)
	require.Equal(t, initBond.QuoRaw(2), validator.Tokens)
	app.StakingKeeper.Unjail(ctx, consAddr0)


	require.False(t, keeper.ValidatorByPowerIndexExists(ctx, app.StakingKeeper, power))


	validator, found = app.StakingKeeper.GetValidator(ctx, validatorAddr)
	require.True(t, found)
	power2 := types.GetValidatorsByPowerIndexKey(validator, app.StakingKeeper.PowerReduction(ctx))
	require.True(t, keeper.ValidatorByPowerIndexExists(ctx, app.StakingKeeper, power2))


	power3 := types.GetValidatorsByPowerIndexKey(validator, app.StakingKeeper.PowerReduction(ctx))
	require.Equal(t, power2, power3)


	totalBond := validator.TokensFromShares(bond.GetShares()).TruncateInt()
	res := tstaking.Undelegate(sdk.AccAddress(validatorAddr), validatorAddr, totalBond, true)

	var resData types.MsgUndelegateResponse
	err = proto.Unmarshal(res.Data, &resData)
	require.NoError(t, err)

	ctx = ctx.WithBlockTime(resData.CompletionTime)
	staking.EndBlocker(ctx, app.StakingKeeper)
	staking.EndBlocker(ctx, app.StakingKeeper)


	_, found = app.StakingKeeper.GetValidator(ctx, validatorAddr)
	require.False(t, found)
	require.False(t, keeper.ValidatorByPowerIndexExists(ctx, app.StakingKeeper, power3))
}

func TestDuplicatesMsgCreateValidator(t *testing.T) {
	initPower := int64(1000000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 10, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))

	addr1, addr2 := valAddrs[0], valAddrs[1]
	pk1, pk2 := PKs[0], PKs[1]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	valTokens := tstaking.CreateValidatorWithValPower(addr1, pk1, 10, true)
	app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)

	validator := tstaking.CheckValidator(addr1, types.Bonded, false)
	assert.Equal(t, addr1.String(), validator.OperatorAddress)
	consKey, err := validator.TmConsPublicKey()
	require.NoError(t, err)
	tmPk1, err := cryptocodec.ToTmProtoPublicKey(pk1)
	require.NoError(t, err)
	assert.Equal(t, tmPk1, consKey)
	assert.Equal(t, valTokens, validator.BondedTokens())
	assert.Equal(t, valTokens.ToDec(), validator.DelegatorShares)
	assert.Equal(t, types.Description{}, validator.Description)


	tstaking.CreateValidator(addr1, pk2, valTokens, false)


	tstaking.CreateValidator(addr2, pk1, valTokens, false)


	tstaking.CreateValidator(addr2, pk2, valTokens, true)


	updates, err := app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(updates))

	validator = tstaking.CheckValidator(addr2, types.Bonded, false)
	assert.Equal(t, addr2.String(), validator.OperatorAddress)
	consPk, err := validator.TmConsPublicKey()
	require.NoError(t, err)
	tmPk2, err := cryptocodec.ToTmProtoPublicKey(pk2)
	require.NoError(t, err)
	assert.Equal(t, tmPk2, consPk)
	assert.True(sdk.IntEq(t, valTokens, validator.Tokens))
	assert.True(sdk.DecEq(t, valTokens.ToDec(), validator.DelegatorShares))
	assert.Equal(t, types.Description{}, validator.Description)
}

func TestInvalidPubKeyTypeMsgCreateValidator(t *testing.T) {
	initPower := int64(1000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 1, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	ctx = ctx.WithConsensusParams(&abci.ConsensusParams{
		Validator: &tmproto.ValidatorParams{PubKeyTypes: []string{tmtypes.ABCIPubKeyTypeEd25519}},
	})

	addr := valAddrs[0]
	invalidPk := secp256k1.GenPrivKey().PubKey()
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	tstaking.CreateValidator(addr, invalidPk, sdk.NewInt(10), false)
}

func TestBothPubKeyTypesMsgCreateValidator(t *testing.T) {
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, 1000, 2, sdk.NewInt(1000))
	ctx = ctx.WithConsensusParams(&abci.ConsensusParams{
		Validator: &tmproto.ValidatorParams{PubKeyTypes: []string{tmtypes.ABCIPubKeyTypeEd25519, tmtypes.ABCIPubKeyTypeSecp256k1}},
	})

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	testCases := []struct {
		name string
		addr sdk.ValAddress
		pk   cryptotypes.PubKey
	}{
		{
			"can create a validator with ed25519 pubkey",
			valAddrs[0],
			ed25519.GenPrivKey().PubKey(),
		},
		{
			"can create a validator with secp256k1 pubkey",
			valAddrs[1],
			secp256k1.GenPrivKey().PubKey(),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(*testing.T) {
			tstaking.CreateValidator(tc.addr, tc.pk, sdk.NewInt(10), true)
		})
	}
}

func TestLegacyValidatorDelegations(t *testing.T) {
	initPower := int64(1000)
	app, ctx, delAddrs, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 2, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	valAddr := valAddrs[0]
	valConsPubKey, valConsAddr := PKs[0], sdk.ConsAddress(PKs[0].Address())
	delAddr := delAddrs[1]


	bondAmount := tstaking.CreateValidatorWithValPower(valAddr, valConsPubKey, 10, true)


	updates, err := app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(updates))


	validator := tstaking.CheckValidator(valAddr, types.Bonded, false)
	require.Equal(t, bondAmount, validator.DelegatorShares.RoundInt())
	require.Equal(t, bondAmount, validator.BondedTokens())


	tstaking.Delegate(delAddr, valAddr, bondAmount)


	validator = tstaking.CheckValidator(valAddr, types.Bonded, false)
	require.Equal(t, bondAmount.MulRaw(2), validator.DelegatorShares.RoundInt())
	require.Equal(t, bondAmount.MulRaw(2), validator.BondedTokens())


	res := tstaking.Undelegate(sdk.AccAddress(valAddr), valAddr, bondAmount, true)

	var resData types.MsgUndelegateResponse
	err = proto.Unmarshal(res.Data, &resData)
	require.NoError(t, err)

	ctx = ctx.WithBlockTime(resData.CompletionTime)
	tstaking.Ctx = ctx
	staking.EndBlocker(ctx, app.StakingKeeper)


	validator = tstaking.CheckValidator(valAddr, -1, true)
	require.Equal(t, bondAmount, validator.Tokens)


	bond, found := app.StakingKeeper.GetDelegation(ctx, delAddr, valAddr)
	require.True(t, found)
	require.Equal(t, bondAmount, bond.Shares.RoundInt())
	require.Equal(t, bondAmount, validator.DelegatorShares.RoundInt())


	tstaking.Delegate(sdk.AccAddress(valAddr), valAddr, bondAmount)


	validator, found = app.StakingKeeper.GetValidator(ctx, valAddr)
	require.True(t, found)
	require.Equal(t, bondAmount.MulRaw(2), validator.DelegatorShares.RoundInt())
	require.Equal(t, bondAmount.MulRaw(2), validator.Tokens)


	app.StakingKeeper.Unjail(ctx, valConsAddr)


	tstaking.Delegate(delAddr, valAddr, bondAmount)


	validator, found = app.StakingKeeper.GetValidator(ctx, valAddr)
	require.True(t, found)
	require.Equal(t, bondAmount.MulRaw(3), validator.DelegatorShares.RoundInt())
	require.Equal(t, bondAmount.MulRaw(3), validator.Tokens)


	bond, found = app.StakingKeeper.GetDelegation(ctx, delAddr, valAddr)
	require.True(t, found)
	require.Equal(t, bondAmount.MulRaw(2), bond.Shares.RoundInt())
	require.Equal(t, bondAmount.MulRaw(3), validator.DelegatorShares.RoundInt())
}

func TestIncrementsMsgDelegate(t *testing.T) {
	initPower := int64(1000)
	initBond := sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction)
	app, ctx, delAddrs, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 2, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))

	params := app.StakingKeeper.GetParams(ctx)
	validatorAddr, delegatorAddr := valAddrs[0], delAddrs[1]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	bondAmount := tstaking.CreateValidatorWithValPower(validatorAddr, PKs[0], 10, true)


	app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)

	validator := tstaking.CheckValidator(validatorAddr, types.Bonded, false)
	require.Equal(t, bondAmount, validator.DelegatorShares.RoundInt())
	require.Equal(t, bondAmount, validator.BondedTokens(), "validator: %v", validator)

	tstaking.CheckDelegator(delegatorAddr, validatorAddr, false)

	bond, found := app.StakingKeeper.GetDelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr)
	require.True(t, found)
	require.Equal(t, bondAmount, bond.Shares.RoundInt())

	bondedTokens := app.StakingKeeper.TotalBondedTokens(ctx)
	require.Equal(t, bondAmount, bondedTokens)

	for i := int64(0); i < 5; i++ {
		ctx = ctx.WithBlockHeight(i)
		tstaking.Ctx = ctx
		tstaking.Delegate(delegatorAddr, validatorAddr, bondAmount)


		validator, found := app.StakingKeeper.GetValidator(ctx, validatorAddr)
		require.True(t, found)
		bond, found := app.StakingKeeper.GetDelegation(ctx, delegatorAddr, validatorAddr)
		require.True(t, found)

		expBond := bondAmount.MulRaw(i + 1)
		expDelegatorShares := bondAmount.MulRaw(i + 2)
		expDelegatorAcc := initBond.Sub(expBond)

		gotBond := bond.Shares.RoundInt()
		gotDelegatorShares := validator.DelegatorShares.RoundInt()
		gotDelegatorAcc := app.BankKeeper.GetBalance(ctx, delegatorAddr, params.BondDenom).Amount

		require.Equal(t, expBond, gotBond,
			"i: %v\nexpBond: %v\ngotBond: %v\nvalidator: %v\nbond: %v\n",
			i, expBond, gotBond, validator, bond)
		require.Equal(t, expDelegatorShares, gotDelegatorShares,
			"i: %v\nexpDelegatorShares: %v\ngotDelegatorShares: %v\nvalidator: %v\nbond: %v\n",
			i, expDelegatorShares, gotDelegatorShares, validator, bond)
		require.Equal(t, expDelegatorAcc, gotDelegatorAcc,
			"i: %v\nexpDelegatorAcc: %v\ngotDelegatorAcc: %v\nvalidator: %v\nbond: %v\n",
			i, expDelegatorAcc, gotDelegatorAcc, validator, bond)
	}
}

func TestEditValidatorDecreaseMinSelfDelegation(t *testing.T) {
	initPower := int64(100)
	initBond := sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 1, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))

	validatorAddr := valAddrs[0]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	msgCreateValidator := tstaking.CreateValidatorMsg(validatorAddr, PKs[0], initBond)
	msgCreateValidator.MinSelfDelegation = sdk.NewInt(2)
	tstaking.Handle(msgCreateValidator, true)


	updates, err := app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(updates))


	bond, found := app.StakingKeeper.GetDelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr)
	require.True(t, found)
	gotBond := bond.Shares.RoundInt()
	require.Equal(t, initBond, gotBond,
		"initBond: %v\ngotBond: %v\nbond: %v\n",
		initBond, gotBond, bond)

	newMinSelfDelegation := sdk.OneInt()
	msgEditValidator := types.NewMsgEditValidator(validatorAddr, types.Description{}, nil, &newMinSelfDelegation)
	tstaking.Handle(msgEditValidator, false)
}

func TestEditValidatorIncreaseMinSelfDelegationBeyondCurrentBond(t *testing.T) {
	initPower := int64(100)
	initBond := sdk.TokensFromConsensusPower(100, sdk.DefaultPowerReduction)

	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 2, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	validatorAddr := valAddrs[0]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	msgCreateValidator := tstaking.CreateValidatorMsg(validatorAddr, PKs[0], initBond)
	msgCreateValidator.MinSelfDelegation = sdk.NewInt(2)
	tstaking.Handle(msgCreateValidator, true)


	updates, err := app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.NoError(t, err)
	require.Equal(t, 1, len(updates))


	bond, found := app.StakingKeeper.GetDelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr)
	require.True(t, found)
	gotBond := bond.Shares.RoundInt()
	require.Equal(t, initBond, gotBond,
		"initBond: %v\ngotBond: %v\nbond: %v\n",
		initBond, gotBond, bond)

	newMinSelfDelegation := initBond.Add(sdk.OneInt())
	msgEditValidator := types.NewMsgEditValidator(validatorAddr, types.Description{}, nil, &newMinSelfDelegation)
	tstaking.Handle(msgEditValidator, false)
}

func TestIncrementsMsgUnbond(t *testing.T) {
	initPower := int64(1000)

	app, ctx, delAddrs, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 2, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	params := app.StakingKeeper.GetParams(ctx)
	denom := params.BondDenom


	validatorAddr, delegatorAddr := valAddrs[0], delAddrs[1]
	initBond := tstaking.CreateValidatorWithValPower(validatorAddr, PKs[0], initPower, true)


	amt1 := app.BankKeeper.GetBalance(ctx, delegatorAddr, denom).Amount

	tstaking.Delegate(delegatorAddr, validatorAddr, initBond)


	amt2 := app.BankKeeper.GetBalance(ctx, delegatorAddr, denom).Amount
	require.True(sdk.IntEq(t, amt1.Sub(initBond), amt2))


	app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)

	validator, found := app.StakingKeeper.GetValidator(ctx, validatorAddr)
	require.True(t, found)
	require.Equal(t, initBond.MulRaw(2), validator.DelegatorShares.RoundInt())
	require.Equal(t, initBond.MulRaw(2), validator.BondedTokens())



	unbondAmt := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))
	msgUndelegate := types.NewMsgUndelegate(delegatorAddr, validatorAddr, unbondAmt)
	numUnbonds := int64(5)

	for i := int64(0); i < numUnbonds; i++ {
		res := tstaking.Handle(msgUndelegate, true)

		var resData types.MsgUndelegateResponse
		err := proto.Unmarshal(res.Data, &resData)
		require.NoError(t, err)

		ctx = ctx.WithBlockTime(resData.CompletionTime)
		tstaking.Ctx = ctx
		staking.EndBlocker(ctx, app.StakingKeeper)


		validator, found = app.StakingKeeper.GetValidator(ctx, validatorAddr)
		require.True(t, found)
		bond, found := app.StakingKeeper.GetDelegation(ctx, delegatorAddr, validatorAddr)
		require.True(t, found)

		expBond := initBond.Sub(unbondAmt.Amount.Mul(sdk.NewInt(i + 1)))
		expDelegatorShares := initBond.MulRaw(2).Sub(unbondAmt.Amount.Mul(sdk.NewInt(i + 1)))
		expDelegatorAcc := initBond.Sub(expBond)

		gotBond := bond.Shares.RoundInt()
		gotDelegatorShares := validator.DelegatorShares.RoundInt()
		gotDelegatorAcc := app.BankKeeper.GetBalance(ctx, delegatorAddr, params.BondDenom).Amount

		require.Equal(t, expBond, gotBond,
			"i: %v\nexpBond: %v\ngotBond: %v\nvalidator: %v\nbond: %v\n",
			i, expBond, gotBond, validator, bond)
		require.Equal(t, expDelegatorShares, gotDelegatorShares,
			"i: %v\nexpDelegatorShares: %v\ngotDelegatorShares: %v\nvalidator: %v\nbond: %v\n",
			i, expDelegatorShares, gotDelegatorShares, validator, bond)
		require.Equal(t, expDelegatorAcc, gotDelegatorAcc,
			"i: %v\nexpDelegatorAcc: %v\ngotDelegatorAcc: %v\nvalidator: %v\nbond: %v\n",
			i, expDelegatorAcc, gotDelegatorAcc, validator, bond)
	}


	errorCases := []sdk.Int{


		app.StakingKeeper.TokensFromConsensusPower(ctx, 1<<63-1),
		app.StakingKeeper.TokensFromConsensusPower(ctx, 1<<31),
		initBond,
	}

	for _, c := range errorCases {
		tstaking.Undelegate(delegatorAddr, validatorAddr, c, false)
	}


	leftBonded := initBond.Sub(unbondAmt.Amount.Mul(sdk.NewInt(numUnbonds)))
	tstaking.Undelegate(delegatorAddr, validatorAddr, leftBonded, true)
}

func TestMultipleMsgCreateValidator(t *testing.T) {
	initPower := int64(1000)
	initTokens := sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction)
	app, ctx, delAddrs, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 3, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))

	params := app.StakingKeeper.GetParams(ctx)
	blockTime := time.Now().UTC()
	ctx = ctx.WithBlockTime(blockTime)
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	validatorAddrs := []sdk.ValAddress{
		valAddrs[0],
		valAddrs[1],
		valAddrs[2],
	}
	delegatorAddrs := []sdk.AccAddress{
		delAddrs[0],
		delAddrs[1],
		delAddrs[2],
	}


	amt := app.StakingKeeper.TokensFromConsensusPower(ctx, 10)
	for i, validatorAddr := range validatorAddrs {
		tstaking.CreateValidator(validatorAddr, PKs[i], amt, true)

		validators := app.StakingKeeper.GetValidators(ctx, 100)
		require.Equal(t, (i + 1), len(validators))

		val := validators[i]
		balanceExpd := initTokens.Sub(amt)
		balanceGot := app.BankKeeper.GetBalance(ctx, delegatorAddrs[i], params.BondDenom).Amount

		require.Equal(t, i+1, len(validators), "expected %d validators got %d, validators: %v", i+1, len(validators), validators)
		require.Equal(t, amt, val.DelegatorShares.RoundInt(), "expected %d shares, got %d", amt, val.DelegatorShares)
		require.Equal(t, balanceExpd, balanceGot, "expected account to have %d, got %d", balanceExpd, balanceGot)
	}

	staking.EndBlocker(ctx, app.StakingKeeper)


	for i, validatorAddr := range validatorAddrs {
		_, found := app.StakingKeeper.GetValidator(ctx, validatorAddr)
		require.True(t, found)

		res := tstaking.Undelegate(delegatorAddrs[i], validatorAddr, amt, true)

		var resData types.MsgUndelegateResponse
		err := proto.Unmarshal(res.Data, &resData)
		require.NoError(t, err)


		staking.EndBlocker(ctx, app.StakingKeeper)


		staking.EndBlocker(ctx.WithBlockTime(blockTime.Add(params.UnbondingTime)), app.StakingKeeper)


		validators := app.StakingKeeper.GetValidators(ctx, 100)
		require.Equal(t, len(validatorAddrs)-(i+1), len(validators),
			"expected %d validators got %d", len(validatorAddrs)-(i+1), len(validators))

		_, found = app.StakingKeeper.GetValidator(ctx, validatorAddr)
		require.False(t, found)

		gotBalance := app.BankKeeper.GetBalance(ctx, delegatorAddrs[i], params.BondDenom).Amount
		require.Equal(t, initTokens, gotBalance, "expected account to have %d, got %d", initTokens, gotBalance)
	}
}

func TestMultipleMsgDelegate(t *testing.T) {
	initPower := int64(1000)
	app, ctx, delAddrs, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 50, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	validatorAddr, delegatorAddrs := valAddrs[0], delAddrs[1:]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	var amount int64 = 10


	tstaking.CreateValidator(validatorAddr, PKs[0], sdk.NewInt(amount), true)


	for _, delegatorAddr := range delegatorAddrs {
		tstaking.Delegate(delegatorAddr, validatorAddr, sdk.NewInt(10))
		tstaking.CheckDelegator(delegatorAddr, validatorAddr, true)
	}


	for _, delegatorAddr := range delegatorAddrs {
		res := tstaking.Undelegate(delegatorAddr, validatorAddr, sdk.NewInt(amount), true)

		var resData types.MsgUndelegateResponse
		err := proto.Unmarshal(res.Data, &resData)
		require.NoError(t, err)

		ctx = ctx.WithBlockTime(resData.CompletionTime)
		staking.EndBlocker(ctx, app.StakingKeeper)
		tstaking.Ctx = ctx


		_, found := app.StakingKeeper.GetDelegation(ctx, delegatorAddr, validatorAddr)
		require.False(t, found)
	}
}

func TestJailValidator(t *testing.T) {
	initPower := int64(1000)
	app, ctx, delAddrs, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 2, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	validatorAddr, delegatorAddr := valAddrs[0], delAddrs[1]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	var amt int64 = 10


	tstaking.CreateValidator(validatorAddr, PKs[0], sdk.NewInt(amt), true)
	tstaking.Delegate(delegatorAddr, validatorAddr, sdk.NewInt(amt))


	unamt := sdk.NewInt(amt)
	res := tstaking.Undelegate(sdk.AccAddress(validatorAddr), validatorAddr, unamt, true)

	var resData types.MsgUndelegateResponse
	err := proto.Unmarshal(res.Data, &resData)
	require.NoError(t, err)

	ctx = ctx.WithBlockTime(resData.CompletionTime)
	staking.EndBlocker(ctx, app.StakingKeeper)
	tstaking.Ctx = ctx

	tstaking.CheckValidator(validatorAddr, -1, true)


	tstaking.Undelegate(delegatorAddr, validatorAddr, unamt, true)

	err = proto.Unmarshal(res.Data, &resData)
	require.NoError(t, err)

	ctx = ctx.WithBlockTime(resData.CompletionTime)
	staking.EndBlocker(ctx, app.StakingKeeper)
	tstaking.Ctx = ctx


	tstaking.CreateValidator(validatorAddr, PKs[0], sdk.NewInt(amt), true)
}

func TestValidatorQueue(t *testing.T) {
	initPower := int64(1000)
	app, ctx, delAddrs, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 2, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	validatorAddr, delegatorAddr := valAddrs[0], delAddrs[1]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	params := app.StakingKeeper.GetParams(ctx)
	params.UnbondingTime = 7 * time.Second
	app.StakingKeeper.SetParams(ctx, params)


	amt := tstaking.CreateValidatorWithValPower(validatorAddr, PKs[0], 10, true)
	tstaking.Delegate(delegatorAddr, validatorAddr, amt)
	staking.EndBlocker(ctx, app.StakingKeeper)


	res := tstaking.Undelegate(sdk.AccAddress(validatorAddr), validatorAddr, amt, true)

	var resData types.MsgUndelegateResponse
	err := proto.Unmarshal(res.Data, &resData)
	require.NoError(t, err)

	finishTime := resData.CompletionTime

	ctx = tstaking.TurnBlock(finishTime)
	origHeader := ctx.BlockHeader()

	validator, found := app.StakingKeeper.GetValidator(ctx, validatorAddr)
	require.True(t, found)
	require.True(t, validator.IsUnbonding(), "%v", validator)


	ctx = tstaking.TurnBlock(origHeader.Time.Add(time.Second * 6))

	validator, found = app.StakingKeeper.GetValidator(ctx, validatorAddr)
	require.True(t, found)
	require.True(t, validator.IsUnbonding(), "%v", validator)


	ctx = tstaking.TurnBlock(origHeader.Time.Add(time.Second * 7))

	validator, found = app.StakingKeeper.GetValidator(ctx, validatorAddr)
	require.True(t, found)
	require.True(t, validator.IsUnbonded(), "%v", validator)
}

func TestUnbondingPeriod(t *testing.T) {
	initPower := int64(1000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 1, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	validatorAddr := valAddrs[0]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	params := app.StakingKeeper.GetParams(ctx)
	params.UnbondingTime = 7 * time.Second
	app.StakingKeeper.SetParams(ctx, params)


	amt := tstaking.CreateValidatorWithValPower(validatorAddr, PKs[0], 10, true)
	staking.EndBlocker(ctx, app.StakingKeeper)


	tstaking.Undelegate(sdk.AccAddress(validatorAddr), validatorAddr, amt, true)

	origHeader := ctx.BlockHeader()

	_, found := app.StakingKeeper.GetUnbondingDelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr)
	require.True(t, found, "should not have unbonded")


	staking.EndBlocker(ctx, app.StakingKeeper)
	_, found = app.StakingKeeper.GetUnbondingDelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr)
	require.True(t, found, "should not have unbonded")


	ctx = tstaking.TurnBlock(origHeader.Time.Add(time.Second * 6))
	_, found = app.StakingKeeper.GetUnbondingDelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr)
	require.True(t, found, "should not have unbonded")


	ctx = tstaking.TurnBlock(origHeader.Time.Add(time.Second * 7))
	_, found = app.StakingKeeper.GetUnbondingDelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr)
	require.False(t, found, "should have unbonded")
}

func TestUnbondingFromUnbondingValidator(t *testing.T) {
	initPower := int64(1000)
	app, ctx, delAddrs, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 2, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	validatorAddr, delegatorAddr := valAddrs[0], delAddrs[1]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	tstaking.CreateValidator(validatorAddr, PKs[0], sdk.NewInt(10), true)
	tstaking.Delegate(delegatorAddr, validatorAddr, sdk.NewInt(10))


	unbondAmt := sdk.NewInt(10)
	res := tstaking.Undelegate(sdk.AccAddress(validatorAddr), validatorAddr, unbondAmt, true)


	var resData types.MsgUndelegateResponse
	err := proto.Unmarshal(res.Data, &resData)
	require.NoError(t, err)

	ctx = ctx.WithBlockTime(resData.CompletionTime.Add(time.Second * -1))


	res = tstaking.Undelegate(delegatorAddr, validatorAddr, unbondAmt, true)

	ctx = tstaking.TurnBlockTimeDiff(app.StakingKeeper.UnbondingTime(ctx))
	tstaking.Ctx = ctx



	_, found := app.StakingKeeper.GetUnbondingDelegation(ctx, delegatorAddr, validatorAddr)
	require.False(t, found, "should be removed from state")
}

func TestRedelegationPeriod(t *testing.T) {
	initPower := int64(1000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 2, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	validatorAddr, validatorAddr2 := valAddrs[0], valAddrs[1]
	denom := app.StakingKeeper.GetParams(ctx).BondDenom
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	params := app.StakingKeeper.GetParams(ctx)
	params.UnbondingTime = 7 * time.Second
	app.StakingKeeper.SetParams(ctx, params)

	amt1 := app.BankKeeper.GetBalance(ctx, sdk.AccAddress(validatorAddr), denom).Amount


	tstaking.CreateValidator(validatorAddr, PKs[0], sdk.NewInt(10), true)


	amt2 := app.BankKeeper.GetBalance(ctx, sdk.AccAddress(validatorAddr), denom).Amount
	require.Equal(t, amt1.Sub(sdk.NewInt(10)), amt2, "expected coins to be subtracted")

	tstaking.CreateValidator(validatorAddr2, PKs[1], sdk.NewInt(10), true)
	bal1 := app.BankKeeper.GetAllBalances(ctx, sdk.AccAddress(validatorAddr))


	redAmt := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))
	msgBeginRedelegate := types.NewMsgBeginRedelegate(sdk.AccAddress(validatorAddr), validatorAddr, validatorAddr2, redAmt)
	tstaking.Handle(msgBeginRedelegate, true)


	bal2 := app.BankKeeper.GetAllBalances(ctx, sdk.AccAddress(validatorAddr))
	require.Equal(t, bal1, bal2)

	origHeader := ctx.BlockHeader()


	staking.EndBlocker(ctx, app.StakingKeeper)
	_, found := app.StakingKeeper.GetRedelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr, validatorAddr2)
	require.True(t, found, "should not have unbonded")


	ctx = tstaking.TurnBlock(origHeader.Time.Add(time.Second * 6))
	_, found = app.StakingKeeper.GetRedelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr, validatorAddr2)
	require.True(t, found, "should not have unbonded")


	ctx = tstaking.TurnBlock(origHeader.Time.Add(time.Second * 7))
	_, found = app.StakingKeeper.GetRedelegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr, validatorAddr2)
	require.False(t, found, "should have unbonded")
}

func TestTransitiveRedelegation(t *testing.T) {
	initPower := int64(1000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 3, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))

	val1, val2, val3 := valAddrs[0], valAddrs[1], valAddrs[2]
	blockTime := time.Now().UTC()
	ctx = ctx.WithBlockTime(blockTime)
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	tstaking.CreateValidator(val1, PKs[0], sdk.NewInt(10), true)
	tstaking.CreateValidator(val2, PKs[1], sdk.NewInt(10), true)
	tstaking.CreateValidator(val3, PKs[2], sdk.NewInt(10), true)


	redAmt := sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(10))
	msgBeginRedelegate := types.NewMsgBeginRedelegate(sdk.AccAddress(val1), val1, val2, redAmt)
	tstaking.Handle(msgBeginRedelegate, true)


	msgBeginRedelegate = types.NewMsgBeginRedelegate(sdk.AccAddress(val1), val2, val3, redAmt)
	tstaking.Handle(msgBeginRedelegate, false)

	params := app.StakingKeeper.GetParams(ctx)
	ctx = ctx.WithBlockTime(blockTime.Add(params.UnbondingTime))
	tstaking.Ctx = ctx


	staking.EndBlocker(ctx, app.StakingKeeper)


	tstaking.Handle(msgBeginRedelegate, true)
}

func TestMultipleRedelegationAtSameTime(t *testing.T) {
	initPower := int64(1000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 2, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	valAddr := valAddrs[0]
	valAddr2 := valAddrs[1]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	params := app.StakingKeeper.GetParams(ctx)
	params.UnbondingTime = 1 * time.Second
	app.StakingKeeper.SetParams(ctx, params)


	valTokens := tstaking.CreateValidatorWithValPower(valAddr, PKs[0], 10, true)
	tstaking.CreateValidator(valAddr2, PKs[1], valTokens, true)


	staking.EndBlocker(ctx, app.StakingKeeper)


	selfDelAddr := sdk.AccAddress(valAddr)
	redAmt := sdk.NewCoin(sdk.DefaultBondDenom, valTokens.QuoRaw(2))
	msgBeginRedelegate := types.NewMsgBeginRedelegate(selfDelAddr, valAddr, valAddr2, redAmt)
	tstaking.Handle(msgBeginRedelegate, true)


	rd, found := app.StakingKeeper.GetRedelegation(ctx, selfDelAddr, valAddr, valAddr2)
	require.True(t, found)
	require.Len(t, rd.Entries, 1)


	tstaking.Handle(msgBeginRedelegate, true)


	rd, found = app.StakingKeeper.GetRedelegation(ctx, selfDelAddr, valAddr, valAddr2)
	require.True(t, found)
	require.Len(t, rd.Entries, 2)


	ctx = tstaking.TurnBlockTimeDiff(1 * time.Second)
	rd, found = app.StakingKeeper.GetRedelegation(ctx, selfDelAddr, valAddr, valAddr2)
	require.False(t, found)
}

func TestMultipleRedelegationAtUniqueTimes(t *testing.T) {
	initPower := int64(1000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 2, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	valAddr := valAddrs[0]
	valAddr2 := valAddrs[1]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	params := app.StakingKeeper.GetParams(ctx)
	params.UnbondingTime = 10 * time.Second
	app.StakingKeeper.SetParams(ctx, params)


	valTokens := tstaking.CreateValidatorWithValPower(valAddr, PKs[0], 10, true)
	tstaking.CreateValidator(valAddr2, PKs[1], valTokens, true)


	staking.EndBlocker(ctx, app.StakingKeeper)


	selfDelAddr := sdk.AccAddress(valAddr)
	redAmt := sdk.NewCoin(sdk.DefaultBondDenom, valTokens.QuoRaw(2))
	msgBeginRedelegate := types.NewMsgBeginRedelegate(selfDelAddr, valAddr, valAddr2, redAmt)
	tstaking.Handle(msgBeginRedelegate, true)


	ctx = ctx.WithBlockTime(ctx.BlockHeader().Time.Add(5 * time.Second))
	tstaking.Ctx = ctx
	tstaking.Handle(msgBeginRedelegate, true)


	rd, found := app.StakingKeeper.GetRedelegation(ctx, selfDelAddr, valAddr, valAddr2)
	require.True(t, found)
	require.Len(t, rd.Entries, 2)


	ctx = tstaking.TurnBlockTimeDiff(5 * time.Second)
	rd, found = app.StakingKeeper.GetRedelegation(ctx, selfDelAddr, valAddr, valAddr2)
	require.True(t, found)
	require.Len(t, rd.Entries, 1)


	ctx = tstaking.TurnBlockTimeDiff(5 * time.Second)
	rd, found = app.StakingKeeper.GetRedelegation(ctx, selfDelAddr, valAddr, valAddr2)
	require.False(t, found)
}

func TestMultipleUnbondingDelegationAtSameTime(t *testing.T) {
	initPower := int64(1000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 1, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	valAddr := valAddrs[0]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	params := app.StakingKeeper.GetParams(ctx)
	params.UnbondingTime = 1 * time.Second
	app.StakingKeeper.SetParams(ctx, params)


	valTokens := tstaking.CreateValidatorWithValPower(valAddr, PKs[0], 10, true)


	staking.EndBlocker(ctx, app.StakingKeeper)


	selfDelAddr := sdk.AccAddress(valAddr)
	tstaking.Undelegate(selfDelAddr, valAddr, valTokens.QuoRaw(2), true)


	ubd, found := app.StakingKeeper.GetUnbondingDelegation(ctx, selfDelAddr, valAddr)
	require.True(t, found)
	require.Len(t, ubd.Entries, 1)


	tstaking.Undelegate(selfDelAddr, valAddr, valTokens.QuoRaw(2), true)


	ubd, found = app.StakingKeeper.GetUnbondingDelegation(ctx, selfDelAddr, valAddr)
	require.True(t, found)
	require.Len(t, ubd.Entries, 2)


	ctx = tstaking.TurnBlockTimeDiff(1 * time.Second)
	ubd, found = app.StakingKeeper.GetUnbondingDelegation(ctx, selfDelAddr, valAddr)
	require.False(t, found)
}

func TestMultipleUnbondingDelegationAtUniqueTimes(t *testing.T) {
	initPower := int64(1000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 1, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	valAddr := valAddrs[0]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	params := app.StakingKeeper.GetParams(ctx)
	params.UnbondingTime = 10 * time.Second
	app.StakingKeeper.SetParams(ctx, params)


	valTokens := tstaking.CreateValidatorWithValPower(valAddr, PKs[0], 10, true)


	staking.EndBlocker(ctx, app.StakingKeeper)


	selfDelAddr := sdk.AccAddress(valAddr)
	tstaking.Undelegate(selfDelAddr, valAddr, valTokens.QuoRaw(2), true)


	ubd, found := app.StakingKeeper.GetUnbondingDelegation(ctx, selfDelAddr, valAddr)
	require.True(t, found)
	require.Len(t, ubd.Entries, 1)


	ctx = ctx.WithBlockTime(ctx.BlockHeader().Time.Add(5 * time.Second))
	tstaking.Ctx = ctx
	tstaking.Undelegate(selfDelAddr, valAddr, valTokens.QuoRaw(2), true)


	ubd, found = app.StakingKeeper.GetUnbondingDelegation(ctx, selfDelAddr, valAddr)
	require.True(t, found)
	require.Len(t, ubd.Entries, 2)


	ctx = tstaking.TurnBlockTimeDiff(5 * time.Second)
	ubd, found = app.StakingKeeper.GetUnbondingDelegation(ctx, selfDelAddr, valAddr)
	require.True(t, found)
	require.Len(t, ubd.Entries, 1)


	ctx = tstaking.TurnBlockTimeDiff(5 * time.Second)
	ubd, found = app.StakingKeeper.GetUnbondingDelegation(ctx, selfDelAddr, valAddr)
	require.False(t, found)
}

func TestUnbondingWhenExcessValidators(t *testing.T) {
	initPower := int64(1000)
	app, ctx, _, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 3, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	val1 := valAddrs[0]
	val2 := valAddrs[1]
	val3 := valAddrs[2]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)


	params := app.StakingKeeper.GetParams(ctx)
	params.MaxValidators = 2
	app.StakingKeeper.SetParams(ctx, params)


	tstaking.CreateValidatorWithValPower(val1, PKs[0], 50, true)

	app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.Equal(t, 1, len(app.StakingKeeper.GetLastValidators(ctx)))

	valTokens2 := tstaking.CreateValidatorWithValPower(val2, PKs[1], 30, true)
	app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.Equal(t, 2, len(app.StakingKeeper.GetLastValidators(ctx)))

	tstaking.CreateValidatorWithValPower(val3, PKs[2], 10, true)
	app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.Equal(t, 2, len(app.StakingKeeper.GetLastValidators(ctx)))


	tstaking.Undelegate(sdk.AccAddress(val2), val2, valTokens2, true)

	app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)




	vals := app.StakingKeeper.GetLastValidators(ctx)
	require.Equal(t, 2, len(vals), "vals %v", vals)
	tstaking.CheckValidator(val1, types.Bonded, false)
}

func TestBondUnbondRedelegateSlashTwice(t *testing.T) {
	initPower := int64(1000)
	app, ctx, delAddrs, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 3, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	valA, valB, del := valAddrs[0], valAddrs[1], delAddrs[2]
	consAddr0 := sdk.ConsAddress(PKs[0].Address())
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	valTokens := tstaking.CreateValidatorWithValPower(valA, PKs[0], 10, true)
	tstaking.CreateValidator(valB, PKs[1], valTokens, true)


	tstaking.Delegate(del, valA, valTokens)


	updates, err := app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.NoError(t, err)
	require.Equal(t, 2, len(updates))


	ctx = ctx.WithBlockHeight(1)
	tstaking.Ctx = ctx


	unbondAmt := app.StakingKeeper.TokensFromConsensusPower(ctx, 4)
	tstaking.Undelegate(del, valA, unbondAmt, true)


	redAmt := sdk.NewCoin(sdk.DefaultBondDenom, app.StakingKeeper.TokensFromConsensusPower(ctx, 6))
	msgBeginRedelegate := types.NewMsgBeginRedelegate(del, valA, valB, redAmt)
	tstaking.Handle(msgBeginRedelegate, true)


	delegation, found := app.StakingKeeper.GetDelegation(ctx, del, valB)
	require.True(t, found)
	require.Equal(t, sdk.NewDecFromInt(redAmt.Amount), delegation.Shares)


	updates, err = app.StakingKeeper.ApplyAndReturnValidatorSetUpdates(ctx)
	require.NoError(t, err)
	require.Equal(t, 2, len(updates))


	app.StakingKeeper.Slash(ctx, consAddr0, 0, 20, sdk.NewDecWithPrec(5, 1))


	ubd, found := app.StakingKeeper.GetUnbondingDelegation(ctx, del, valA)
	require.True(t, found)
	require.Len(t, ubd.Entries, 1)
	require.Equal(t, unbondAmt.QuoRaw(2), ubd.Entries[0].Balance)


	redelegation, found := app.StakingKeeper.GetRedelegation(ctx, del, valA, valB)
	require.True(t, found)
	require.Len(t, redelegation.Entries, 1)


	delegation, found = app.StakingKeeper.GetDelegation(ctx, del, valB)
	require.True(t, found)
	require.Equal(t, sdk.NewDecFromInt(redAmt.Amount.QuoRaw(2)), delegation.Shares)


	validator, found := app.StakingKeeper.GetValidator(ctx, valA)
	require.True(t, found)
	require.Equal(t, valTokens.QuoRaw(2), validator.GetBondedTokens())


	ctx = ctx.WithBlockHeight(3)
	app.StakingKeeper.Slash(ctx, consAddr0, 2, 10, sdk.NewDecWithPrec(5, 1))
	tstaking.Ctx = ctx


	ubd, found = app.StakingKeeper.GetUnbondingDelegation(ctx, del, valA)
	require.True(t, found)
	require.Len(t, ubd.Entries, 1)
	require.Equal(t, unbondAmt.QuoRaw(2), ubd.Entries[0].Balance)


	redelegation, found = app.StakingKeeper.GetRedelegation(ctx, del, valA, valB)
	require.True(t, found)
	require.Len(t, redelegation.Entries, 1)


	delegation, found = app.StakingKeeper.GetDelegation(ctx, del, valB)
	require.True(t, found)
	require.Equal(t, sdk.NewDecFromInt(redAmt.Amount.QuoRaw(2)), delegation.Shares)


	staking.EndBlocker(ctx, app.StakingKeeper)



	validator, _ = app.StakingKeeper.GetValidator(ctx, valA)
	require.Equal(t, validator.GetStatus(), types.Unbonding)
}

func TestInvalidMsg(t *testing.T) {
	k := keeper.Keeper{}
	h := staking.NewHandler(k)

	res, err := h(sdk.NewContext(nil, tmproto.Header{}, false, nil), testdata.NewTestMsg())
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, strings.Contains(err.Error(), "unrecognized staking message type"))
}

func TestInvalidCoinDenom(t *testing.T) {
	initPower := int64(1000)
	app, ctx, delAddrs, valAddrs := bootstrapHandlerGenesisTest(t, initPower, 3, sdk.TokensFromConsensusPower(initPower, sdk.DefaultPowerReduction))
	valA, valB, delAddr := valAddrs[0], valAddrs[1], delAddrs[2]
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	valTokens := app.StakingKeeper.TokensFromConsensusPower(ctx, 100)
	invalidCoin := sdk.NewCoin("churros", valTokens)
	validCoin := sdk.NewCoin(sdk.DefaultBondDenom, valTokens)
	oneCoin := sdk.NewCoin(sdk.DefaultBondDenom, sdk.OneInt())

	commission := types.NewCommissionRates(sdk.OneDec(), sdk.OneDec(), sdk.ZeroDec())
	msgCreate, err := types.NewMsgCreateValidator(valA, PKs[0], invalidCoin, types.Description{}, commission, sdk.OneInt())
	require.NoError(t, err)
	tstaking.Handle(msgCreate, false)

	msgCreate, err = types.NewMsgCreateValidator(valA, PKs[0], validCoin, types.Description{}, commission, sdk.OneInt())
	require.NoError(t, err)
	tstaking.Handle(msgCreate, true)

	msgCreate, err = types.NewMsgCreateValidator(valB, PKs[1], validCoin, types.Description{}, commission, sdk.OneInt())
	require.NoError(t, err)
	tstaking.Handle(msgCreate, true)

	msgDelegate := types.NewMsgDelegate(delAddr, valA, invalidCoin)
	tstaking.Handle(msgDelegate, false)

	msgDelegate = types.NewMsgDelegate(delAddr, valA, validCoin)
	tstaking.Handle(msgDelegate, true)

	msgUndelegate := types.NewMsgUndelegate(delAddr, valA, invalidCoin)
	tstaking.Handle(msgUndelegate, false)

	msgUndelegate = types.NewMsgUndelegate(delAddr, valA, oneCoin)
	tstaking.Handle(msgUndelegate, true)

	msgRedelegate := types.NewMsgBeginRedelegate(delAddr, valA, valB, invalidCoin)
	tstaking.Handle(msgRedelegate, false)

	msgRedelegate = types.NewMsgBeginRedelegate(delAddr, valA, valB, oneCoin)
	tstaking.Handle(msgRedelegate, true)
}
