package v040

import (
	v039mint "github.com/cosmos/cosmos-sdk/x/mint/legacy/v039"
	v040mint "github.com/cosmos/cosmos-sdk/x/mint/types"
)



//

func Migrate(mintGenState v039mint.GenesisState) *v040mint.GenesisState {
	return &v040mint.GenesisState{
		Minter: v040mint.Minter{
			Inflation:        mintGenState.Minter.Inflation,
			AnnualProvisions: mintGenState.Minter.AnnualProvisions,
		},
		Params: v040mint.Params{
			MintDenom:           mintGenState.Params.MintDenom,
			InflationRateChange: mintGenState.Params.InflationRateChange,
			InflationMax:        mintGenState.Params.InflationMax,
			InflationMin:        mintGenState.Params.InflationMin,
			GoalBonded:          mintGenState.Params.GoalBonded,
			BlocksPerYear:       mintGenState.Params.BlocksPerYear,
		},
	}
}
