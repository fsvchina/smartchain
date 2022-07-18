package v043

import (
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)


func migrateJSONWeightedVotes(oldVotes types.Votes) types.Votes {
	newVotes := make(types.Votes, len(oldVotes))
	for i, oldVote := range oldVotes {
		newVotes[i] = migrateVote(oldVote)
	}

	return newVotes
}



//

func MigrateJSON(oldState *types.GenesisState) *types.GenesisState {
	return &types.GenesisState{
		StartingProposalId: oldState.StartingProposalId,
		Deposits:           oldState.Deposits,
		Votes:              migrateJSONWeightedVotes(oldState.Votes),
		Proposals:          oldState.Proposals,
		DepositParams:      oldState.DepositParams,
		VotingParams:       oldState.VotingParams,
		TallyParams:        oldState.TallyParams,
	}
}
