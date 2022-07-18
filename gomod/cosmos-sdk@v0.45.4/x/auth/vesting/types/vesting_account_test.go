package types_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tmtime "github.com/tendermint/tendermint/types/time"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/testutil/testdata"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/auth/vesting/types"
)

var (
	stakeDenom = "stake"
	feeDenom   = "fee"
)

func TestGetVestedCoinsContVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)

	bacc, origCoins := initBaseAccount()
	cva := types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())


	vestedCoins := cva.GetVestedCoins(now)
	require.Nil(t, vestedCoins)


	vestedCoins = cva.GetVestedCoins(endTime)
	require.Equal(t, origCoins, vestedCoins)


	vestedCoins = cva.GetVestedCoins(now.Add(12 * time.Hour))
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}, vestedCoins)


	vestedCoins = cva.GetVestedCoins(now.Add(48 * time.Hour))
	require.Equal(t, origCoins, vestedCoins)
}

func TestGetVestingCoinsContVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)

	bacc, origCoins := initBaseAccount()
	cva := types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())


	vestingCoins := cva.GetVestingCoins(now)
	require.Equal(t, origCoins, vestingCoins)


	vestingCoins = cva.GetVestingCoins(endTime)
	require.Nil(t, vestingCoins)


	vestingCoins = cva.GetVestingCoins(now.Add(12 * time.Hour))
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}, vestingCoins)
}

func TestSpendableCoinsContVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)

	bacc, origCoins := initBaseAccount()
	cva := types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())



	lockedCoins := cva.LockedCoins(now)
	require.Equal(t, origCoins, lockedCoins)


	lockedCoins = cva.LockedCoins(endTime)
	require.Equal(t, sdk.NewCoins(), lockedCoins)


	lockedCoins = cva.LockedCoins(now.Add(12 * time.Hour))
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}, lockedCoins)
}

func TestTrackDelegationContVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)

	bacc, origCoins := initBaseAccount()


	cva := types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())
	cva.TrackDelegation(now, origCoins, origCoins)
	require.Equal(t, origCoins, cva.DelegatedVesting)
	require.Nil(t, cva.DelegatedFree)


	cva = types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())
	cva.TrackDelegation(endTime, origCoins, origCoins)
	require.Nil(t, cva.DelegatedVesting)
	require.Equal(t, origCoins, cva.DelegatedFree)


	cva = types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())
	cva.TrackDelegation(now.Add(12*time.Hour), origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)}, cva.DelegatedVesting)
	require.Nil(t, cva.DelegatedFree)

	cva.TrackDelegation(now.Add(12*time.Hour), origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)}, cva.DelegatedVesting)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)}, cva.DelegatedFree)


	cva = types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())
	require.Panics(t, func() {
		cva.TrackDelegation(endTime, origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 1000000)})
	})
	require.Nil(t, cva.DelegatedVesting)
	require.Nil(t, cva.DelegatedFree)
}

func TestTrackUndelegationContVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)

	bacc, origCoins := initBaseAccount()


	cva := types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())
	cva.TrackDelegation(now, origCoins, origCoins)
	cva.TrackUndelegation(origCoins)
	require.Nil(t, cva.DelegatedFree)
	require.Nil(t, cva.DelegatedVesting)


	cva = types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())

	cva.TrackDelegation(endTime, origCoins, origCoins)
	cva.TrackUndelegation(origCoins)
	require.Nil(t, cva.DelegatedFree)
	require.Nil(t, cva.DelegatedVesting)


	cva = types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())

	require.Panics(t, func() {
		cva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 0)})
	})
	require.Nil(t, cva.DelegatedFree)
	require.Nil(t, cva.DelegatedVesting)


	cva = types.NewContinuousVestingAccount(bacc, origCoins, now.Unix(), endTime.Unix())
	cva.TrackDelegation(now.Add(12*time.Hour), origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	cva.TrackDelegation(now.Add(12*time.Hour), origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})


	cva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 25)})
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 25)}, cva.DelegatedFree)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)}, cva.DelegatedVesting)


	cva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	require.Nil(t, cva.DelegatedFree)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 25)}, cva.DelegatedVesting)
}

func TestGetVestedCoinsDelVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)

	bacc, origCoins := initBaseAccount()


	dva := types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())
	vestedCoins := dva.GetVestedCoins(now)
	require.Nil(t, vestedCoins)


	vestedCoins = dva.GetVestedCoins(endTime)
	require.Equal(t, origCoins, vestedCoins)
}

func TestGetVestingCoinsDelVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)

	bacc, origCoins := initBaseAccount()


	dva := types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())
	vestingCoins := dva.GetVestingCoins(now)
	require.Equal(t, origCoins, vestingCoins)


	vestingCoins = dva.GetVestingCoins(endTime)
	require.Nil(t, vestingCoins)
}

func TestSpendableCoinsDelVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)

	bacc, origCoins := initBaseAccount()



	dva := types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())
	lockedCoins := dva.LockedCoins(now)
	require.True(t, lockedCoins.IsEqual(origCoins))



	lockedCoins = dva.LockedCoins(endTime)
	require.Equal(t, sdk.NewCoins(), lockedCoins)


	lockedCoins = dva.LockedCoins(now.Add(12 * time.Hour))
	require.True(t, lockedCoins.IsEqual(origCoins))



	delegatedAmount := sdk.NewCoins(sdk.NewInt64Coin(stakeDenom, 50))
	dva.TrackDelegation(now.Add(12*time.Hour), origCoins, delegatedAmount)
	lockedCoins = dva.LockedCoins(now.Add(12 * time.Hour))
	require.True(t, lockedCoins.IsEqual(origCoins.Sub(delegatedAmount)))
}

func TestTrackDelegationDelVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)

	bacc, origCoins := initBaseAccount()


	dva := types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())
	dva.TrackDelegation(now, origCoins, origCoins)
	require.Equal(t, origCoins, dva.DelegatedVesting)
	require.Nil(t, dva.DelegatedFree)


	dva = types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())
	dva.TrackDelegation(endTime, origCoins, origCoins)
	require.Nil(t, dva.DelegatedVesting)
	require.Equal(t, origCoins, dva.DelegatedFree)



	dva = types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())
	dva.TrackDelegation(now.Add(12*time.Hour), origCoins, origCoins)
	require.Equal(t, origCoins, dva.DelegatedVesting)
	require.Nil(t, dva.DelegatedFree)


	dva = types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())

	require.Panics(t, func() {
		dva.TrackDelegation(endTime, origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 1000000)})
	})
	require.Nil(t, dva.DelegatedVesting)
	require.Nil(t, dva.DelegatedFree)
}

func TestTrackUndelegationDelVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)

	bacc, origCoins := initBaseAccount()


	dva := types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())
	dva.TrackDelegation(now, origCoins, origCoins)
	dva.TrackUndelegation(origCoins)
	require.Nil(t, dva.DelegatedFree)
	require.Nil(t, dva.DelegatedVesting)


	dva = types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())
	dva.TrackDelegation(endTime, origCoins, origCoins)
	dva.TrackUndelegation(origCoins)
	require.Nil(t, dva.DelegatedFree)
	require.Nil(t, dva.DelegatedVesting)


	dva = types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())

	require.Panics(t, func() {
		dva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 0)})
	})
	require.Nil(t, dva.DelegatedFree)
	require.Nil(t, dva.DelegatedVesting)


	dva = types.NewDelayedVestingAccount(bacc, origCoins, endTime.Unix())
	dva.TrackDelegation(now.Add(12*time.Hour), origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	dva.TrackDelegation(now.Add(12*time.Hour), origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})


	dva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 25)})

	require.Nil(t, dva.DelegatedFree)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 75)}, dva.DelegatedVesting)


	dva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	require.Nil(t, dva.DelegatedFree)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 25)}, dva.DelegatedVesting)
}

func TestGetVestedCoinsPeriodicVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)
	periods := types.Periods{
		types.Period{Length: int64(12 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}},
		types.Period{Length: int64(6 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}},
		types.Period{Length: int64(6 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}},
	}

	bacc, origCoins := initBaseAccount()
	pva := types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)


	vestedCoins := pva.GetVestedCoins(now)
	require.Nil(t, vestedCoins)


	vestedCoins = pva.GetVestedCoins(endTime)
	require.Equal(t, origCoins, vestedCoins)


	vestedCoins = pva.GetVestedCoins(now.Add(6 * time.Hour))
	require.Nil(t, vestedCoins)


	vestedCoins = pva.GetVestedCoins(now.Add(12 * time.Hour))
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}, vestedCoins)


	vestedCoins = pva.GetVestedCoins(now.Add(15 * time.Hour))
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}, vestedCoins)


	vestedCoins = pva.GetVestedCoins(now.Add(18 * time.Hour))
	require.Equal(t,
		sdk.Coins{
			sdk.NewInt64Coin(feeDenom, 750), sdk.NewInt64Coin(stakeDenom, 75)}, vestedCoins)


	vestedCoins = pva.GetVestedCoins(now.Add(48 * time.Hour))
	require.Equal(t, origCoins, vestedCoins)
}

func TestGetVestingCoinsPeriodicVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)
	periods := types.Periods{
		types.Period{Length: int64(12 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}},
		types.Period{Length: int64(6 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}},
		types.Period{Length: int64(6 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}},
	}

	bacc, origCoins := initBaseAccount()
	pva := types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)


	vestingCoins := pva.GetVestingCoins(now)
	require.Equal(t, origCoins, vestingCoins)


	vestingCoins = pva.GetVestingCoins(endTime)
	require.Nil(t, vestingCoins)


	vestingCoins = pva.GetVestingCoins(now.Add(12 * time.Hour))
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}, vestingCoins)


	vestingCoins = pva.GetVestingCoins(now.Add(15 * time.Hour))
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}, vestingCoins)


	vestingCoins = pva.GetVestingCoins(now.Add(18 * time.Hour))
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}, vestingCoins)


	vestingCoins = pva.GetVestingCoins(now.Add(48 * time.Hour))
	require.Nil(t, vestingCoins)
}

func TestSpendableCoinsPeriodicVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)
	periods := types.Periods{
		types.Period{Length: int64(12 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}},
		types.Period{Length: int64(6 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}},
		types.Period{Length: int64(6 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}},
	}

	bacc, origCoins := initBaseAccount()
	pva := types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)



	lockedCoins := pva.LockedCoins(now)
	require.Equal(t, origCoins, lockedCoins)



	lockedCoins = pva.LockedCoins(endTime)
	require.Equal(t, sdk.NewCoins(), lockedCoins)


	lockedCoins = pva.LockedCoins(now.Add(12 * time.Hour))
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}, lockedCoins)
}

func TestTrackDelegationPeriodicVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)
	periods := types.Periods{
		types.Period{Length: int64(12 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}},
		types.Period{Length: int64(6 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}},
		types.Period{Length: int64(6 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}},
	}

	bacc, origCoins := initBaseAccount()


	pva := types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)
	pva.TrackDelegation(now, origCoins, origCoins)
	require.Equal(t, origCoins, pva.DelegatedVesting)
	require.Nil(t, pva.DelegatedFree)


	pva = types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)
	pva.TrackDelegation(endTime, origCoins, origCoins)
	require.Nil(t, pva.DelegatedVesting)
	require.Equal(t, origCoins, pva.DelegatedFree)


	pva = types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)
	pva.TrackDelegation(now, origCoins, periods[0].Amount)

	require.Equal(t, pva.DelegatedVesting, periods[0].Amount)
	require.Nil(t, pva.DelegatedFree)


	pva = types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)
	pva.TrackDelegation(now.Add(12*time.Hour), origCoins, periods[0].Amount.Add(periods[1].Amount...))

	require.Equal(t, pva.DelegatedFree, periods[1].Amount)
	require.Equal(t, pva.DelegatedVesting, periods[0].Amount)


	pva = types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)
	pva.TrackDelegation(now.Add(12*time.Hour), origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)}, pva.DelegatedVesting)
	require.Nil(t, pva.DelegatedFree)

	pva.TrackDelegation(now.Add(12*time.Hour), origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)}, pva.DelegatedVesting)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)}, pva.DelegatedFree)


	pva = types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)
	require.Panics(t, func() {
		pva.TrackDelegation(endTime, origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 1000000)})
	})
	require.Nil(t, pva.DelegatedVesting)
	require.Nil(t, pva.DelegatedFree)
}

func TestTrackUndelegationPeriodicVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(24 * time.Hour)
	periods := types.Periods{
		types.Period{Length: int64(12 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 500), sdk.NewInt64Coin(stakeDenom, 50)}},
		types.Period{Length: int64(6 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}},
		types.Period{Length: int64(6 * 60 * 60), Amount: sdk.Coins{sdk.NewInt64Coin(feeDenom, 250), sdk.NewInt64Coin(stakeDenom, 25)}},
	}

	bacc, origCoins := initBaseAccount()


	pva := types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)
	pva.TrackDelegation(now, origCoins, origCoins)
	pva.TrackUndelegation(origCoins)
	require.Nil(t, pva.DelegatedFree)
	require.Nil(t, pva.DelegatedVesting)


	pva = types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)

	pva.TrackDelegation(endTime, origCoins, origCoins)
	pva.TrackUndelegation(origCoins)
	require.Nil(t, pva.DelegatedFree)
	require.Nil(t, pva.DelegatedVesting)


	pva = types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)
	pva.TrackDelegation(endTime, origCoins, periods[0].Amount)
	pva.TrackUndelegation(periods[0].Amount)
	require.Nil(t, pva.DelegatedFree)
	require.Nil(t, pva.DelegatedVesting)


	pva = types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)

	require.Panics(t, func() {
		pva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 0)})
	})
	require.Nil(t, pva.DelegatedFree)
	require.Nil(t, pva.DelegatedVesting)


	pva = types.NewPeriodicVestingAccount(bacc, origCoins, now.Unix(), periods)
	pva.TrackDelegation(now.Add(12*time.Hour), origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	pva.TrackDelegation(now.Add(12*time.Hour), origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})


	pva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 25)})
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 25)}, pva.DelegatedFree)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)}, pva.DelegatedVesting)


	pva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	require.Nil(t, pva.DelegatedFree)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 25)}, pva.DelegatedVesting)
}

func TestGetVestedCoinsPermLockedVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(1000 * 24 * time.Hour)

	bacc, origCoins := initBaseAccount()


	plva := types.NewPermanentLockedAccount(bacc, origCoins)
	vestedCoins := plva.GetVestedCoins(now)
	require.Nil(t, vestedCoins)


	vestedCoins = plva.GetVestedCoins(endTime)
	require.Nil(t, vestedCoins)
}

func TestGetVestingCoinsPermLockedVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(1000 * 24 * time.Hour)

	bacc, origCoins := initBaseAccount()


	plva := types.NewPermanentLockedAccount(bacc, origCoins)
	vestingCoins := plva.GetVestingCoins(now)
	require.Equal(t, origCoins, vestingCoins)


	vestingCoins = plva.GetVestingCoins(endTime)
	require.Equal(t, origCoins, vestingCoins)
}

func TestSpendableCoinsPermLockedVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(1000 * 24 * time.Hour)

	bacc, origCoins := initBaseAccount()



	plva := types.NewPermanentLockedAccount(bacc, origCoins)
	lockedCoins := plva.LockedCoins(now)
	require.True(t, lockedCoins.IsEqual(origCoins))


	lockedCoins = plva.LockedCoins(endTime)
	require.True(t, lockedCoins.IsEqual(origCoins))



	delegatedAmount := sdk.NewCoins(sdk.NewInt64Coin(stakeDenom, 50))
	plva.TrackDelegation(now.Add(12*time.Hour), origCoins, delegatedAmount)
	lockedCoins = plva.LockedCoins(now.Add(12 * time.Hour))
	require.True(t, lockedCoins.IsEqual(origCoins.Sub(delegatedAmount)))
}

func TestTrackDelegationPermLockedVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(1000 * 24 * time.Hour)

	bacc, origCoins := initBaseAccount()


	plva := types.NewPermanentLockedAccount(bacc, origCoins)
	plva.TrackDelegation(now, origCoins, origCoins)
	require.Equal(t, origCoins, plva.DelegatedVesting)
	require.Nil(t, plva.DelegatedFree)


	plva = types.NewPermanentLockedAccount(bacc, origCoins)
	plva.TrackDelegation(endTime, origCoins, origCoins)
	require.Equal(t, origCoins, plva.DelegatedVesting)
	require.Nil(t, plva.DelegatedFree)


	plva = types.NewPermanentLockedAccount(bacc, origCoins)

	require.Panics(t, func() {
		plva.TrackDelegation(endTime, origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 1000000)})
	})
	require.Nil(t, plva.DelegatedVesting)
	require.Nil(t, plva.DelegatedFree)
}

func TestTrackUndelegationPermLockedVestingAcc(t *testing.T) {
	now := tmtime.Now()
	endTime := now.Add(1000 * 24 * time.Hour)

	bacc, origCoins := initBaseAccount()


	plva := types.NewPermanentLockedAccount(bacc, origCoins)
	plva.TrackDelegation(now, origCoins, origCoins)
	plva.TrackUndelegation(origCoins)
	require.Nil(t, plva.DelegatedFree)
	require.Nil(t, plva.DelegatedVesting)


	plva = types.NewPermanentLockedAccount(bacc, origCoins)
	plva.TrackDelegation(endTime, origCoins, origCoins)
	plva.TrackUndelegation(origCoins)
	require.Nil(t, plva.DelegatedFree)
	require.Nil(t, plva.DelegatedVesting)


	plva = types.NewPermanentLockedAccount(bacc, origCoins)
	require.Panics(t, func() {
		plva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 0)})
	})
	require.Nil(t, plva.DelegatedFree)
	require.Nil(t, plva.DelegatedVesting)


	plva = types.NewPermanentLockedAccount(bacc, origCoins)
	plva.TrackDelegation(now, origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	plva.TrackDelegation(now, origCoins, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})


	plva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 25)})

	require.Nil(t, plva.DelegatedFree)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 75)}, plva.DelegatedVesting)


	plva.TrackUndelegation(sdk.Coins{sdk.NewInt64Coin(stakeDenom, 50)})
	require.Nil(t, plva.DelegatedFree)
	require.Equal(t, sdk.Coins{sdk.NewInt64Coin(stakeDenom, 25)}, plva.DelegatedVesting)
}

func TestGenesisAccountValidate(t *testing.T) {
	pubkey := secp256k1.GenPrivKey().PubKey()
	addr := sdk.AccAddress(pubkey.Address())
	baseAcc := authtypes.NewBaseAccount(addr, pubkey, 0, 0)
	initialVesting := sdk.NewCoins(sdk.NewInt64Coin(sdk.DefaultBondDenom, 50))
	baseVestingWithCoins := types.NewBaseVestingAccount(baseAcc, initialVesting, 100)
	tests := []struct {
		name   string
		acc    authtypes.GenesisAccount
		expErr bool
	}{
		{
			"valid base account",
			baseAcc,
			false,
		},
		{
			"invalid base valid account",
			authtypes.NewBaseAccount(addr, secp256k1.GenPrivKey().PubKey(), 0, 0),
			true,
		},
		{
			"valid base vesting account",
			baseVestingWithCoins,
			false,
		},
		{
			"valid continuous vesting account",
			types.NewContinuousVestingAccount(baseAcc, initialVesting, 100, 200),
			false,
		},
		{
			"invalid vesting times",
			types.NewContinuousVestingAccount(baseAcc, initialVesting, 1654668078, 1554668078),
			true,
		},
		{
			"valid periodic vesting account",
			types.NewPeriodicVestingAccount(baseAcc, initialVesting, 0, types.Periods{types.Period{Length: int64(100), Amount: sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 50)}}}),
			false,
		},
		{
			"invalid vesting period lengths",
			types.NewPeriodicVestingAccountRaw(
				baseVestingWithCoins,
				0, types.Periods{types.Period{Length: int64(50), Amount: sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 50)}}}),
			true,
		},
		{
			"invalid vesting period amounts",
			types.NewPeriodicVestingAccountRaw(
				baseVestingWithCoins,
				0, types.Periods{types.Period{Length: int64(100), Amount: sdk.Coins{sdk.NewInt64Coin(sdk.DefaultBondDenom, 25)}}}),
			true,
		},
		{
			"valid permanent locked vesting account",
			types.NewPermanentLockedAccount(baseAcc, initialVesting),
			false,
		},
		{
			"invalid positive end time for permanently locked vest account",
			&types.PermanentLockedAccount{BaseVestingAccount: baseVestingWithCoins},
			true,
		},
	}

	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			require.Equal(t, tt.expErr, tt.acc.Validate() != nil)
		})
	}
}

func TestContinuousVestingAccountMarshal(t *testing.T) {
	baseAcc, coins := initBaseAccount()
	baseVesting := types.NewBaseVestingAccount(baseAcc, coins, time.Now().Unix())
	acc := types.NewContinuousVestingAccountRaw(baseVesting, baseVesting.EndTime)

	bz, err := app.AccountKeeper.MarshalAccount(acc)
	require.Nil(t, err)

	acc2, err := app.AccountKeeper.UnmarshalAccount(bz)
	require.Nil(t, err)
	require.IsType(t, &types.ContinuousVestingAccount{}, acc2)
	require.Equal(t, acc.String(), acc2.String())


	_, err = app.AccountKeeper.UnmarshalAccount(bz[:len(bz)/2])
	require.NotNil(t, err)
}

func TestPeriodicVestingAccountMarshal(t *testing.T) {
	baseAcc, coins := initBaseAccount()
	acc := types.NewPeriodicVestingAccount(baseAcc, coins, time.Now().Unix(), types.Periods{types.Period{3600, coins}})

	bz, err := app.AccountKeeper.MarshalAccount(acc)
	require.Nil(t, err)

	acc2, err := app.AccountKeeper.UnmarshalAccount(bz)
	require.Nil(t, err)
	require.IsType(t, &types.PeriodicVestingAccount{}, acc2)
	require.Equal(t, acc.String(), acc2.String())


	_, err = app.AccountKeeper.UnmarshalAccount(bz[:len(bz)/2])
	require.NotNil(t, err)
}

func TestDelayedVestingAccountMarshal(t *testing.T) {
	baseAcc, coins := initBaseAccount()
	acc := types.NewDelayedVestingAccount(baseAcc, coins, time.Now().Unix())

	bz, err := app.AccountKeeper.MarshalAccount(acc)
	require.Nil(t, err)

	acc2, err := app.AccountKeeper.UnmarshalAccount(bz)
	require.Nil(t, err)
	require.IsType(t, &types.DelayedVestingAccount{}, acc2)
	require.Equal(t, acc.String(), acc2.String())


	_, err = app.AccountKeeper.UnmarshalAccount(bz[:len(bz)/2])
	require.NotNil(t, err)
}
func TestPermanentLockedAccountMarshal(t *testing.T) {
	baseAcc, coins := initBaseAccount()
	acc := types.NewPermanentLockedAccount(baseAcc, coins)

	bz, err := app.AccountKeeper.MarshalAccount(acc)
	require.Nil(t, err)

	acc2, err := app.AccountKeeper.UnmarshalAccount(bz)
	require.Nil(t, err)
	require.IsType(t, &types.PermanentLockedAccount{}, acc2)
	require.Equal(t, acc.String(), acc2.String())


	_, err = app.AccountKeeper.UnmarshalAccount(bz[:len(bz)/2])
	require.NotNil(t, err)
}

func initBaseAccount() (*authtypes.BaseAccount, sdk.Coins) {
	_, _, addr := testdata.KeyTestPubAddr()
	origCoins := sdk.Coins{sdk.NewInt64Coin(feeDenom, 1000), sdk.NewInt64Coin(stakeDenom, 100)}
	bacc := authtypes.NewBaseAccountWithAddress(addr)

	return bacc, origCoins
}
