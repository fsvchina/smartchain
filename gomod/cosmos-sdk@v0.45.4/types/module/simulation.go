package module

import (
	"encoding/json"

	"math/rand"
	"time"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
)



type AppModuleSimulation interface {

	GenerateGenesisState(input *SimulationState)


	ProposalContents(simState SimulationState) []simulation.WeightedProposalContent


	RandomizedParams(r *rand.Rand) []simulation.ParamChange


	RegisterStoreDecoder(sdk.StoreDecoderRegistry)


	WeightedOperations(simState SimulationState) []simulation.WeightedOperation
}



type SimulationManager struct {
	Modules       []AppModuleSimulation
	StoreDecoders sdk.StoreDecoderRegistry
}


//

func NewSimulationManager(modules ...AppModuleSimulation) *SimulationManager {
	return &SimulationManager{
		Modules:       modules,
		StoreDecoders: make(sdk.StoreDecoderRegistry),
	}
}



func (sm *SimulationManager) GetProposalContents(simState SimulationState) []simulation.WeightedProposalContent {
	wContents := make([]simulation.WeightedProposalContent, 0, len(sm.Modules))
	for _, module := range sm.Modules {
		wContents = append(wContents, module.ProposalContents(simState)...)
	}

	return wContents
}


func (sm *SimulationManager) RegisterStoreDecoders() {
	for _, module := range sm.Modules {
		module.RegisterStoreDecoder(sm.StoreDecoders)
	}
}



func (sm *SimulationManager) GenerateGenesisStates(simState *SimulationState) {
	for _, module := range sm.Modules {
		module.GenerateGenesisState(simState)
	}
}



func (sm *SimulationManager) GenerateParamChanges(seed int64) (paramChanges []simulation.ParamChange) {
	r := rand.New(rand.NewSource(seed))

	for _, module := range sm.Modules {
		paramChanges = append(paramChanges, module.RandomizedParams(r)...)
	}

	return
}


func (sm *SimulationManager) WeightedOperations(simState SimulationState) []simulation.WeightedOperation {
	wOps := make([]simulation.WeightedOperation, 0, len(sm.Modules))
	for _, module := range sm.Modules {
		wOps = append(wOps, module.WeightedOperations(simState)...)
	}

	return wOps
}



type SimulationState struct {
	AppParams    simulation.AppParams
	Cdc          codec.JSONCodec
	Rand         *rand.Rand
	GenState     map[string]json.RawMessage
	Accounts     []simulation.Account
	InitialStake int64
	NumBonded    int64
	GenTimestamp time.Time
	UnbondTime   time.Duration
	ParamChanges []simulation.ParamChange
	Contents     []simulation.WeightedProposalContent
}
