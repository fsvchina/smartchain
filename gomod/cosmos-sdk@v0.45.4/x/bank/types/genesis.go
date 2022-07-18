package types

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)



func (gs GenesisState) Validate() error {
	if err := gs.Params.Validate(); err != nil {
		return err
	}

	seenBalances := make(map[string]bool)
	seenMetadatas := make(map[string]bool)

	totalSupply := sdk.Coins{}

	for _, balance := range gs.Balances {
		if seenBalances[balance.Address] {
			return fmt.Errorf("duplicate balance for address %s", balance.Address)
		}

		if err := balance.Validate(); err != nil {
			return err
		}

		seenBalances[balance.Address] = true

		totalSupply = totalSupply.Add(balance.Coins...)
	}

	for _, metadata := range gs.DenomMetadata {
		if seenMetadatas[metadata.Base] {
			return fmt.Errorf("duplicate client metadata for denom %s", metadata.Base)
		}

		if err := metadata.Validate(); err != nil {
			return err
		}

		seenMetadatas[metadata.Base] = true
	}

	if !gs.Supply.Empty() {

		err := gs.Supply.Validate()
		if err != nil {
			return err
		}

		if !gs.Supply.IsEqual(totalSupply) {
			return fmt.Errorf("genesis supply is incorrect, expected %v, got %v", gs.Supply, totalSupply)
		}
	}

	return nil
}


func NewGenesisState(params Params, balances []Balance, supply sdk.Coins, denomMetaData []Metadata) *GenesisState {
	return &GenesisState{
		Params:        params,
		Balances:      balances,
		Supply:        supply,
		DenomMetadata: denomMetaData,
	}
}


func DefaultGenesisState() *GenesisState {
	return NewGenesisState(DefaultParams(), []Balance{}, sdk.Coins{}, []Metadata{})
}



func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}
