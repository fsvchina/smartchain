package simulation

import (
	"fmt"
	"math/rand"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
)

func min(a int, b int) int {
	if a <= b {
		return a
	}
	return b
}




func SimulateParamChangeProposalContent(paramChangePool []simulation.ParamChange) simulation.ContentSimulatorFn {
	numProposals := 0

	maxSimultaneousParamChanges := min(len(paramChangePool), 1000)
	if maxSimultaneousParamChanges == 0 {
		panic("param changes array is empty")
	}

	return func(r *rand.Rand, _ sdk.Context, _ []simulation.Account) simulation.Content {
		numChanges := simulation.RandIntBetween(r, 1, maxSimultaneousParamChanges)
		paramChanges := make([]proposal.ParamChange, numChanges)


		paramChoices := r.Perm(len(paramChangePool))

		for i := 0; i < numChanges; i++ {
			spc := paramChangePool[paramChoices[i]]

			paramChanges[i] = proposal.NewParamChange(spc.Subspace(), spc.Key(), spc.SimValue()(r))
		}

		title := fmt.Sprintf("title from SimulateParamChangeProposalContent-%d", numProposals)
		desc := fmt.Sprintf("desc from SimulateParamChangeProposalContent-%d. Random short desc: %s",
			numProposals, simulation.RandStringOfLength(r, 20))
		numProposals++
		return proposal.NewParameterChangeProposal(
			title,
			desc,
			paramChanges,
		)
	}
}
