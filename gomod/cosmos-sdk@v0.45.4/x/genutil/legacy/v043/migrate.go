package v043

import (
	"github.com/cosmos/cosmos-sdk/client"
	v040bank "github.com/cosmos/cosmos-sdk/x/bank/legacy/v040"
	v043bank "github.com/cosmos/cosmos-sdk/x/bank/legacy/v043"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"
	v040gov "github.com/cosmos/cosmos-sdk/x/gov/legacy/v040"
	v043gov "github.com/cosmos/cosmos-sdk/x/gov/legacy/v043"
	gov "github.com/cosmos/cosmos-sdk/x/gov/types"
)


func Migrate(appState types.AppMap, clientCtx client.Context) types.AppMap {

	if appState[v040gov.ModuleName] != nil {

		var oldGovState gov.GenesisState
		clientCtx.Codec.MustUnmarshalJSON(appState[v040gov.ModuleName], &oldGovState)


		delete(appState, v040gov.ModuleName)



		appState[v043gov.ModuleName] = clientCtx.Codec.MustMarshalJSON(v043gov.MigrateJSON(&oldGovState))
	}

	if appState[v040bank.ModuleName] != nil {

		var oldBankState bank.GenesisState
		clientCtx.Codec.MustUnmarshalJSON(appState[v040bank.ModuleName], &oldBankState)


		delete(appState, v040bank.ModuleName)



		appState[v043bank.ModuleName] = clientCtx.Codec.MustMarshalJSON(v043bank.MigrateJSON(&oldBankState))
	}

	return appState
}
