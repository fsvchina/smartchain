package simapp

import (
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/server/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)



type App interface {

	Name() string



	LegacyAmino() *codec.LegacyAmino


	BeginBlocker(ctx sdk.Context, req abci.RequestBeginBlock) abci.ResponseBeginBlock


	EndBlocker(ctx sdk.Context, req abci.RequestEndBlock) abci.ResponseEndBlock


	InitChainer(ctx sdk.Context, req abci.RequestInitChain) abci.ResponseInitChain


	LoadHeight(height int64) error


	ExportAppStateAndValidators(
		forZeroHeight bool, jailAllowedAddrs []string,
	) (types.ExportedApp, error)


	ModuleAccountAddrs() map[string]bool


	SimulationManager() *module.SimulationManager
}
