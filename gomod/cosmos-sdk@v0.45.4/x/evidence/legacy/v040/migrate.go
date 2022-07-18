package v040

import (
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	v038evidence "github.com/cosmos/cosmos-sdk/x/evidence/legacy/v038"
	v040evidence "github.com/cosmos/cosmos-sdk/x/evidence/types"
)

func migrateEvidence(oldEvidence v038evidence.Evidence) *codectypes.Any {
	switch oldEvidence := oldEvidence.(type) {
	case v038evidence.Equivocation:
		{
			newEquivocation := &v040evidence.Equivocation{
				Height:           oldEvidence.Height,
				Time:             oldEvidence.Time,
				Power:            oldEvidence.Power,
				ConsensusAddress: oldEvidence.ConsensusAddress.String(),
			}
			any, err := codectypes.NewAnyWithValue(newEquivocation)
			if err != nil {
				panic(err)
			}

			return any
		}
	default:
		panic(fmt.Errorf("'%T' is not a valid evidence type", oldEvidence))
	}
}



//



func Migrate(evidenceState v038evidence.GenesisState) *v040evidence.GenesisState {
	var newEvidences = make([]*codectypes.Any, len(evidenceState.Evidence))
	for i, oldEvidence := range evidenceState.Evidence {
		newEvidences[i] = migrateEvidence(oldEvidence)
	}

	return &v040evidence.GenesisState{
		Evidence: newEvidences,
	}
}
