package simulation



import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)


const (
	unbondingTime     = "unbonding_time"
	maxValidators     = "max_validators"
	historicalEntries = "historical_entries"
)


func genUnbondingTime(r *rand.Rand) (ubdTime time.Duration) {
	return time.Duration(simulation.RandIntBetween(r, 60, 60*60*24*3*2)) * time.Second
}


func genMaxValidators(r *rand.Rand) (maxValidators uint32) {
	return uint32(r.Intn(250) + 1)
}


func getHistEntries(r *rand.Rand) uint32 {
	return uint32(r.Intn(int(types.DefaultHistoricalEntries + 1)))
}


func RandomizedGenState(simState *module.SimulationState) {

	var (
		unbondTime  time.Duration
		maxVals     uint32
		histEntries uint32
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, unbondingTime, &unbondTime, simState.Rand,
		func(r *rand.Rand) { unbondTime = genUnbondingTime(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, maxValidators, &maxVals, simState.Rand,
		func(r *rand.Rand) { maxVals = genMaxValidators(r) },
	)

	simState.AppParams.GetOrGenerate(
		simState.Cdc, historicalEntries, &histEntries, simState.Rand,
		func(r *rand.Rand) { histEntries = getHistEntries(r) },
	)



	simState.UnbondTime = unbondTime
	params := types.NewParams(simState.UnbondTime, maxVals, 7, histEntries, sdk.DefaultBondDenom)


	var (
		validators  []types.Validator
		delegations []types.Delegation
	)

	valAddrs := make([]sdk.ValAddress, simState.NumBonded)

	for i := 0; i < int(simState.NumBonded); i++ {
		valAddr := sdk.ValAddress(simState.Accounts[i].Address)
		valAddrs[i] = valAddr

		maxCommission := sdk.NewDecWithPrec(int64(simulation.RandIntBetween(simState.Rand, 1, 100)), 2)
		commission := types.NewCommission(
			simulation.RandomDecAmount(simState.Rand, maxCommission),
			maxCommission,
			simulation.RandomDecAmount(simState.Rand, maxCommission),
		)

		validator, err := types.NewValidator(valAddr, simState.Accounts[i].ConsKey.PubKey(), types.Description{})
		if err != nil {
			panic(err)
		}
		validator.Tokens = sdk.NewInt(simState.InitialStake)
		validator.DelegatorShares = sdk.NewDec(simState.InitialStake)
		validator.Commission = commission

		delegation := types.NewDelegation(simState.Accounts[i].Address, valAddr, sdk.NewDec(simState.InitialStake))

		validators = append(validators, validator)
		delegations = append(delegations, delegation)
	}

	stakingGenesis := types.NewGenesisState(params, validators, delegations)

	bz, err := json.MarshalIndent(&stakingGenesis.Params, "", " ")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Selected randomly generated staking parameters:\n%s\n", bz)
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(stakingGenesis)
}
