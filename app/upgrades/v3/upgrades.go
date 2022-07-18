package v3

import (
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/x/genutil/types"

	feemarketv010 "github.com/tharsis/ethermint/x/feemarket/migrations/v010"
	feemarketv09types "github.com/tharsis/ethermint/x/feemarket/migrations/v09/types"
	feemarkettypes "github.com/tharsis/ethermint/x/feemarket/types"
)



func MigrateGenesis(appState types.AppMap, clientCtx client.Context) types.AppMap {

	if appState[feemarkettypes.ModuleName] == nil {
		return appState
	}


	var oldFeeMarketState feemarketv09types.GenesisState
	if err := clientCtx.Codec.UnmarshalJSON(appState[feemarkettypes.ModuleName], &oldFeeMarketState); err != nil {
		return appState
	}


	delete(appState, feemarkettypes.ModuleName)



	newFeeMarketState := feemarketv010.MigrateJSON(oldFeeMarketState)

	feeMarketBz, err := clientCtx.Codec.MarshalJSON(&newFeeMarketState)
	if err != nil {
		return appState
	}

	appState[feemarkettypes.ModuleName] = feeMarketBz

	return appState
}
