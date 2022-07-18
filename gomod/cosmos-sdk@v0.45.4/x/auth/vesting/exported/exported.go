package exported

import (
	"time"

	"github.com/cosmos/cosmos-sdk/x/auth/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


type VestingAccount interface {
	types.AccountI



	//



	LockedCoins(blockTime time.Time) sdk.Coins





	TrackDelegation(blockTime time.Time, balance, amount sdk.Coins)



	TrackUndelegation(amount sdk.Coins)

	GetVestedCoins(blockTime time.Time) sdk.Coins
	GetVestingCoins(blockTime time.Time) sdk.Coins

	GetStartTime() int64
	GetEndTime() int64

	GetOriginalVesting() sdk.Coins
	GetDelegatedFree() sdk.Coins
	GetDelegatedVesting() sdk.Coins
}
