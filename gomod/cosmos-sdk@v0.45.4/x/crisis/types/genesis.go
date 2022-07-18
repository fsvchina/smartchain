package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)


func NewGenesisState(constantFee sdk.Coin) *GenesisState {
	return &GenesisState{
		ConstantFee: constantFee,
	}
}


func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		ConstantFee: sdk.NewCoin(sdk.DefaultBondDenom, sdk.NewInt(1000)),
	}
}


func ValidateGenesis(data *GenesisState) error {
	if !data.ConstantFee.IsPositive() {
		return fmt.Errorf("constant fee must be positive: %s", data.ConstantFee)
	}
	return nil
}
