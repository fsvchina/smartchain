package v043

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/x/gov/types"
)

const proposalIDLen = 8





func migratePrefixProposalAddress(store sdk.KVStore, prefixBz []byte) {
	oldStore := prefix.NewStore(store, prefixBz)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		proposalID := oldStoreIter.Key()[:proposalIDLen]
		addr := oldStoreIter.Key()[proposalIDLen:]
		newStoreKey := append(append(prefixBz, proposalID...), address.MustLengthPrefix(addr)...)


		store.Set(newStoreKey, oldStoreIter.Value())
		oldStore.Delete(oldStoreIter.Key())
	}
}




func migrateVote(oldVote types.Vote) types.Vote {
	return types.Vote{
		ProposalId: oldVote.ProposalId,
		Voter:      oldVote.Voter,
		Options:    types.NewNonSplitVoteOption(oldVote.Option),
	}
}


func migrateStoreWeightedVotes(store sdk.KVStore, cdc codec.BinaryCodec) error {
	iterator := sdk.KVStorePrefixIterator(store, types.VotesKeyPrefix)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		var oldVote types.Vote
		err := cdc.Unmarshal(iterator.Value(), &oldVote)
		if err != nil {
			return err
		}

		newVote := migrateVote(oldVote)
		fmt.Println("migrateStoreWeightedVotes newVote=", newVote)
		bz, err := cdc.Marshal(&newVote)
		if err != nil {
			return err
		}

		store.Set(iterator.Key(), bz)
	}

	return nil
}



//

func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	migratePrefixProposalAddress(store, types.DepositsKeyPrefix)
	migratePrefixProposalAddress(store, types.VotesKeyPrefix)
	return migrateStoreWeightedVotes(store, cdc)
}
