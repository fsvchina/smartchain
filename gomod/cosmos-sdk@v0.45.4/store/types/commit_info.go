package types

import (
	fmt "fmt"

	ics23 "github.com/confio/ics23/go"
	tmcrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"

	sdkmaps "github.com/cosmos/cosmos-sdk/store/internal/maps"
	sdkproofs "github.com/cosmos/cosmos-sdk/store/internal/proofs"
)



//




func (si StoreInfo) GetHash() []byte {
	return si.CommitId.Hash
}

func (ci CommitInfo) toMap() map[string][]byte {
	m := make(map[string][]byte, len(ci.StoreInfos))
	for _, storeInfo := range ci.StoreInfos {
		m[storeInfo.Name] = storeInfo.GetHash()
	}

	return m
}


func (ci CommitInfo) Hash() []byte {

	if len(ci.StoreInfos) == 0 {
		return nil
	}

	rootHash, _, _ := sdkmaps.ProofsFromMap(ci.toMap())
	return rootHash
}

func (ci CommitInfo) ProofOp(storeName string) tmcrypto.ProofOp {
	cmap := ci.toMap()
	_, proofs, _ := sdkmaps.ProofsFromMap(cmap)

	proof := proofs[storeName]
	if proof == nil {
		panic(fmt.Sprintf("ProofOp for %s but not registered store name", storeName))
	}


	existProof, err := sdkproofs.ConvertExistenceProof(proof, []byte(storeName), cmap[storeName])
	if err != nil {
		panic(fmt.Errorf("could not convert simple proof to existence proof: %w", err))
	}

	commitmentProof := &ics23.CommitmentProof{
		Proof: &ics23.CommitmentProof_Exist{
			Exist: existProof,
		},
	}

	return NewSimpleMerkleCommitmentOp([]byte(storeName), commitmentProof).ProofOp()
}

func (ci CommitInfo) CommitID() CommitID {
	return CommitID{
		Version: ci.Version,
		Hash:    ci.Hash(),
	}
}
