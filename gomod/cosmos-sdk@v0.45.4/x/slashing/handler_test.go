package slashing_test

import (
	"errors"
	"strings"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/cosmos/cosmos-sdk/simapp"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/slashing"
	"github.com/cosmos/cosmos-sdk/x/slashing/keeper"
	"github.com/cosmos/cosmos-sdk/x/slashing/testslashing"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
	"github.com/cosmos/cosmos-sdk/x/staking"
	"github.com/cosmos/cosmos-sdk/x/staking/teststaking"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
)

func TestCannotUnjailUnlessJailed(t *testing.T) {

	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	pks := simapp.CreateTestPubKeys(1)
	simapp.AddTestAddrsFromPubKeys(app, ctx, pks, app.StakingKeeper.TokensFromConsensusPower(ctx, 200))

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	slh := slashing.NewHandler(app.SlashingKeeper)
	addr, val := sdk.ValAddress(pks[0].Address()), pks[0]

	amt := tstaking.CreateValidatorWithValPower(addr, val, 100, true)
	staking.EndBlocker(ctx, app.StakingKeeper)
	require.Equal(
		t, app.BankKeeper.GetAllBalances(ctx, sdk.AccAddress(addr)),
		sdk.Coins{sdk.NewCoin(app.StakingKeeper.GetParams(ctx).BondDenom, InitTokens.Sub(amt))},
	)
	require.Equal(t, amt, app.StakingKeeper.Validator(ctx, addr).GetBondedTokens())


	res, err := slh(ctx, types.NewMsgUnjail(addr))
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, errors.Is(types.ErrValidatorNotJailed, err))
}

func TestCannotUnjailUnlessMeetMinSelfDelegation(t *testing.T) {

	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	pks := simapp.CreateTestPubKeys(1)
	simapp.AddTestAddrsFromPubKeys(app, ctx, pks, app.StakingKeeper.TokensFromConsensusPower(ctx, 200))

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	slh := slashing.NewHandler(app.SlashingKeeper)
	addr, val := sdk.ValAddress(pks[0].Address()), pks[0]
	amt := app.StakingKeeper.TokensFromConsensusPower(ctx, 100)
	msg := tstaking.CreateValidatorMsg(addr, val, amt)
	msg.MinSelfDelegation = amt
	tstaking.Handle(msg, true)

	staking.EndBlocker(ctx, app.StakingKeeper)
	require.Equal(
		t, app.BankKeeper.GetAllBalances(ctx, sdk.AccAddress(addr)),
		sdk.Coins{sdk.NewCoin(app.StakingKeeper.GetParams(ctx).BondDenom, InitTokens.Sub(amt))},
	)

	tstaking.Undelegate(sdk.AccAddress(addr), addr, sdk.OneInt(), true)
	require.True(t, app.StakingKeeper.Validator(ctx, addr).IsJailed())


	res, err := slh(ctx, types.NewMsgUnjail(addr))
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, errors.Is(types.ErrSelfDelegationTooLowToUnjail, err))
}

func TestJailedValidatorDelegations(t *testing.T) {

	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Unix(0, 0)})
	pks := simapp.CreateTestPubKeys(3)

	simapp.AddTestAddrsFromPubKeys(app, ctx, pks, app.StakingKeeper.TokensFromConsensusPower(ctx, 20))
	app.SlashingKeeper.SetParams(ctx, testslashing.TestParams())

	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)
	stakingParams := app.StakingKeeper.GetParams(ctx)
	app.StakingKeeper.SetParams(ctx, stakingParams)
	valAddr, consAddr := sdk.ValAddress(pks[1].Address()), sdk.ConsAddress(pks[0].Address())

	amt := tstaking.CreateValidatorWithValPower(valAddr, pks[1], 10, true)
	staking.EndBlocker(ctx, app.StakingKeeper)


	newInfo := types.NewValidatorSigningInfo(consAddr, 0, 0, time.Unix(0, 0), false, 0)
	app.SlashingKeeper.SetValidatorSigningInfo(ctx, consAddr, newInfo)


	delAddr := sdk.AccAddress(pks[2].Address())
	tstaking.Delegate(delAddr, valAddr, amt)


	valAcc := sdk.AccAddress(valAddr)
	tstaking.Undelegate(valAcc, valAddr, amt, true)
	_, err := app.StakingKeeper.CompleteUnbonding(ctx, sdk.AccAddress(valAddr), valAddr)
	require.Nil(t, err, "expected complete unbonding validator to be ok, got: %v", err)


	validator, found := app.StakingKeeper.GetValidator(ctx, valAddr)
	require.True(t, found)
	require.True(t, validator.IsJailed())


	res, err := slashing.NewHandler(app.SlashingKeeper)(ctx, types.NewMsgUnjail(valAddr))
	require.Error(t, err)
	require.Nil(t, res)


	tstaking.Delegate(valAcc, valAddr, amt)


	res, err = slashing.NewHandler(app.SlashingKeeper)(ctx, types.NewMsgUnjail(valAddr))
	require.NoError(t, err)
	require.NotNil(t, res)
}

func TestInvalidMsg(t *testing.T) {
	k := keeper.Keeper{}
	h := slashing.NewHandler(k)

	res, err := h(sdk.NewContext(nil, tmproto.Header{}, false, nil), testdata.NewTestMsg())
	require.Error(t, err)
	require.Nil(t, res)
	require.True(t, strings.Contains(err.Error(), "unrecognized slashing message type"))
}



func TestHandleAbsentValidator(t *testing.T) {

	app := simapp.Setup(false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Unix(0, 0)})
	pks := simapp.CreateTestPubKeys(1)
	simapp.AddTestAddrsFromPubKeys(app, ctx, pks, app.StakingKeeper.TokensFromConsensusPower(ctx, 200))
	app.SlashingKeeper.SetParams(ctx, testslashing.TestParams())

	power := int64(100)
	addr, val := sdk.ValAddress(pks[0].Address()), pks[0]
	slh := slashing.NewHandler(app.SlashingKeeper)
	tstaking := teststaking.NewHelper(t, ctx, app.StakingKeeper)

	amt := tstaking.CreateValidatorWithValPower(addr, val, power, true)
	staking.EndBlocker(ctx, app.StakingKeeper)

	require.Equal(
		t, app.BankKeeper.GetAllBalances(ctx, sdk.AccAddress(addr)),
		sdk.NewCoins(sdk.NewCoin(app.StakingKeeper.GetParams(ctx).BondDenom, InitTokens.Sub(amt))),
	)
	require.Equal(t, amt, app.StakingKeeper.Validator(ctx, addr).GetBondedTokens())


	info, found := app.SlashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(val.Address()))
	require.True(t, found)
	require.Equal(t, int64(0), info.StartHeight)
	require.Equal(t, int64(0), info.IndexOffset)
	require.Equal(t, int64(0), info.MissedBlocksCounter)
	require.Equal(t, time.Unix(0, 0).UTC(), info.JailedUntil)
	height := int64(0)


	for ; height < app.SlashingKeeper.SignedBlocksWindow(ctx); height++ {
		ctx = ctx.WithBlockHeight(height)
		app.SlashingKeeper.HandleValidatorSignature(ctx, val.Address(), power, true)
	}
	info, found = app.SlashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(val.Address()))
	require.True(t, found)
	require.Equal(t, int64(0), info.StartHeight)
	require.Equal(t, int64(0), info.MissedBlocksCounter)


	for ; height < app.SlashingKeeper.SignedBlocksWindow(ctx)+(app.SlashingKeeper.SignedBlocksWindow(ctx)-app.SlashingKeeper.MinSignedPerWindow(ctx)); height++ {
		ctx = ctx.WithBlockHeight(height)
		app.SlashingKeeper.HandleValidatorSignature(ctx, val.Address(), power, false)
	}
	info, found = app.SlashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(val.Address()))
	require.True(t, found)
	require.Equal(t, int64(0), info.StartHeight)
	require.Equal(t, app.SlashingKeeper.SignedBlocksWindow(ctx)-app.SlashingKeeper.MinSignedPerWindow(ctx), info.MissedBlocksCounter)


	validator, _ := app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(val))
	require.Equal(t, stakingtypes.Bonded, validator.GetStatus())

	bondPool := app.StakingKeeper.GetBondedPool(ctx)
	require.True(sdk.IntEq(t, amt, app.BankKeeper.GetBalance(ctx, bondPool.GetAddress(), app.StakingKeeper.BondDenom(ctx)).Amount))


	ctx = ctx.WithBlockHeight(height)
	app.SlashingKeeper.HandleValidatorSignature(ctx, val.Address(), power, false)
	info, found = app.SlashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(val.Address()))
	require.True(t, found)
	require.Equal(t, int64(0), info.StartHeight)

	require.Equal(t, int64(0), info.MissedBlocksCounter)


	staking.EndBlocker(ctx, app.StakingKeeper)


	validator, _ = app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(val))
	require.Equal(t, stakingtypes.Unbonding, validator.GetStatus())

	slashAmt := amt.ToDec().Mul(app.SlashingKeeper.SlashFractionDowntime(ctx)).RoundInt()


	require.True(t, amt.Sub(slashAmt).Equal(validator.GetTokens()))


	height++
	ctx = ctx.WithBlockHeight(height)
	app.SlashingKeeper.HandleValidatorSignature(ctx, val.Address(), power, false)
	info, found = app.SlashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(val.Address()))
	require.True(t, found)
	require.Equal(t, int64(0), info.StartHeight)
	require.Equal(t, int64(1), info.MissedBlocksCounter)


	staking.EndBlocker(ctx, app.StakingKeeper)


	validator, _ = app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(val))
	require.True(t, amt.Sub(slashAmt).Equal(validator.GetTokens()))


	res, err := slh(ctx, types.NewMsgUnjail(addr))
	require.Error(t, err)
	require.Nil(t, res)


	ctx = ctx.WithBlockHeader(tmproto.Header{Time: time.Unix(1, 0).Add(app.SlashingKeeper.DowntimeJailDuration(ctx))})
	res, err = slh(ctx, types.NewMsgUnjail(addr))
	require.NoError(t, err)
	require.NotNil(t, res)


	staking.EndBlocker(ctx, app.StakingKeeper)


	validator, _ = app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(val))
	require.Equal(t, stakingtypes.Bonded, validator.GetStatus())


	require.True(t, amt.Sub(slashAmt).Equal(app.BankKeeper.GetBalance(ctx, bondPool.GetAddress(), app.StakingKeeper.BondDenom(ctx)).Amount))


	info, found = app.SlashingKeeper.GetValidatorSigningInfo(ctx, sdk.ConsAddress(val.Address()))
	require.True(t, found)
	require.Equal(t, int64(0), info.StartHeight)

	require.Equal(t, int64(1), info.MissedBlocksCounter)


	height++
	ctx = ctx.WithBlockHeight(height)
	app.SlashingKeeper.HandleValidatorSignature(ctx, val.Address(), power, false)
	validator, _ = app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(val))
	require.Equal(t, stakingtypes.Bonded, validator.GetStatus())


	nextHeight := height + app.SlashingKeeper.MinSignedPerWindow(ctx) + 1
	for ; height < nextHeight; height++ {
		ctx = ctx.WithBlockHeight(height)
		app.SlashingKeeper.HandleValidatorSignature(ctx, val.Address(), power, false)
	}


	staking.EndBlocker(ctx, app.StakingKeeper)


	nextHeight = height + app.SlashingKeeper.MinSignedPerWindow(ctx) + 1
	for ; height <= nextHeight; height++ {
		ctx = ctx.WithBlockHeight(height)
		app.SlashingKeeper.HandleValidatorSignature(ctx, val.Address(), power, false)
	}


	staking.EndBlocker(ctx, app.StakingKeeper)

	validator, _ = app.StakingKeeper.GetValidatorByConsAddr(ctx, sdk.GetConsAddress(val))
	require.Equal(t, stakingtypes.Unbonding, validator.GetStatus())
}
