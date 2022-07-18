package v043

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	v040auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v040"
	v043distribution "github.com/cosmos/cosmos-sdk/x/distribution/legacy/v043"
	v040staking "github.com/cosmos/cosmos-sdk/x/staking/legacy/v040"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
)





func migratePrefixAddressAddressAddress(store sdk.KVStore, prefixBz []byte) {
	oldStore := prefix.NewStore(store, prefixBz)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		addr1 := oldStoreIter.Key()[:v040auth.AddrLen]
		addr2 := oldStoreIter.Key()[v040auth.AddrLen : 2*v040auth.AddrLen]
		addr3 := oldStoreIter.Key()[2*v040auth.AddrLen:]
		newStoreKey := append(append(append(
			prefixBz,
			address.MustLengthPrefix(addr1)...), address.MustLengthPrefix(addr2)...), address.MustLengthPrefix(addr3)...,
		)


		store.Set(newStoreKey, oldStoreIter.Value())
		oldStore.Delete(oldStoreIter.Key())
	}
}

const powerBytesLen = 8

func migrateValidatorsByPowerIndexKey(store sdk.KVStore) {
	oldStore := prefix.NewStore(store, v040staking.ValidatorsByPowerIndexKey)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		powerBytes := oldStoreIter.Key()[:powerBytesLen]
		valAddr := oldStoreIter.Key()[powerBytesLen:]
		newStoreKey := append(append(types.ValidatorsByPowerIndexKey, powerBytes...), address.MustLengthPrefix(valAddr)...)


		store.Set(newStoreKey, oldStoreIter.Value())
		oldStore.Delete(oldStoreIter.Key())
	}
}



//

func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey) error {
	store := ctx.KVStore(storeKey)

	v043distribution.MigratePrefixAddress(store, v040staking.LastValidatorPowerKey)

	v043distribution.MigratePrefixAddress(store, v040staking.ValidatorsKey)
	v043distribution.MigratePrefixAddress(store, v040staking.ValidatorsByConsAddrKey)
	migrateValidatorsByPowerIndexKey(store)

	v043distribution.MigratePrefixAddressAddress(store, v040staking.DelegationKey)
	v043distribution.MigratePrefixAddressAddress(store, v040staking.UnbondingDelegationKey)
	v043distribution.MigratePrefixAddressAddress(store, v040staking.UnbondingDelegationByValIndexKey)
	migratePrefixAddressAddressAddress(store, v040staking.RedelegationKey)
	migratePrefixAddressAddressAddress(store, v040staking.RedelegationByValSrcIndexKey)
	migratePrefixAddressAddressAddress(store, v040staking.RedelegationByValDstIndexKey)

	return nil
}
