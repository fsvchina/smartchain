package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


func InitialFeePool() FeePool {
	return FeePool{
		CommunityPool: sdk.DecCoins{},
	}
}


func (f FeePool) ValidateGenesis() error {
	if f.CommunityPool.IsAnyNegative() {
		return fmt.Errorf("negative CommunityPool in distribution fee pool, is %v",
			f.CommunityPool)
	}

	return nil
}
