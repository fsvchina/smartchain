package simulation

import (
	"encoding/json"
	"fmt"
	"math/rand"

	"github.com/cosmos/cosmos-sdk/types/module"

	"github.com/tharsis/ethermint/x/evm/types"
)

const (
	extraEIPsKey = "extra_eips"
)


func GenExtraEIPs(r *rand.Rand) []int64 {
	var extraEIPs []int64

	if r.Intn(2) == 0 {
		extraEIPs = []int64{1344, 1884, 2200, 2929, 3198, 3529}
	}
	return extraEIPs
}


func GenEnableCreate(r *rand.Rand) bool {

	enableCreate := r.Intn(100) < 80
	return enableCreate
}


func GenEnableCall(r *rand.Rand) bool {

	enableCall := r.Intn(100) < 80
	return enableCall
}


func RandomizedGenState(simState *module.SimulationState) {

	var extraEIPs []int64

	simState.AppParams.GetOrGenerate(
		simState.Cdc, extraEIPsKey, &extraEIPs, simState.Rand,
		func(r *rand.Rand) { extraEIPs = GenExtraEIPs(r) },
	)

	params := types.NewParams(types.DefaultEVMDenom, true, true, types.DefaultChainConfig(), extraEIPs...)
	evmGenesis := types.NewGenesisState(params, []types.GenesisAccount{})

	bz, err := json.MarshalIndent(evmGenesis, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated %s parameters:\n%s\n", types.ModuleName, bz)

	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(evmGenesis)
}
