package v040

import (
	v039crisis "github.com/cosmos/cosmos-sdk/x/crisis/legacy/v039"
	v040crisis "github.com/cosmos/cosmos-sdk/x/crisis/types"
)



//

func Migrate(crisisGenState v039crisis.GenesisState) *v040crisis.GenesisState {
	return &v040crisis.GenesisState{
		ConstantFee: crisisGenState.ConstantFee,
	}
}
