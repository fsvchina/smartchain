package rootmulti

import (
	"github.com/tendermint/tendermint/crypto/merkle"

	storetypes "github.com/cosmos/cosmos-sdk/store/types"
)


func RequireProof(subpath string) bool {




	return subpath == "/key"
}





func DefaultProofRuntime() (prt *merkle.ProofRuntime) {
	prt = merkle.NewProofRuntime()
	prt.RegisterOpDecoder(storetypes.ProofOpIAVLCommitment, storetypes.CommitmentOpDecoder)
	prt.RegisterOpDecoder(storetypes.ProofOpSimpleMerkleCommitment, storetypes.CommitmentOpDecoder)
	return
}
