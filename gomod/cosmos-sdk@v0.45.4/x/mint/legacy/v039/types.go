package v039

import sdk "github.com/cosmos/cosmos-sdk/types"

const (
	ModuleName = "mint"
)

type (

	Minter struct {
		Inflation        sdk.Dec `json:"inflation" yaml:"inflation"`
		AnnualProvisions sdk.Dec `json:"annual_provisions" yaml:"annual_provisions"`
	}


	Params struct {
		MintDenom           string  `json:"mint_denom" yaml:"mint_denom"`
		InflationRateChange sdk.Dec `json:"inflation_rate_change" yaml:"inflation_rate_change"`
		InflationMax        sdk.Dec `json:"inflation_max" yaml:"inflation_max"`
		InflationMin        sdk.Dec `json:"inflation_min" yaml:"inflation_min"`
		GoalBonded          sdk.Dec `json:"goal_bonded" yaml:"goal_bonded"`
		BlocksPerYear       uint64  `json:"blocks_per_year" yaml:"blocks_per_year"`
	}


	GenesisState struct {
		Minter Minter `json:"minter" yaml:"minter"`
		Params Params `json:"params" yaml:"params"`
	}
)
