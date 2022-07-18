package v043

import (
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	v040auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v040"
	v040bank "github.com/cosmos/cosmos-sdk/x/bank/legacy/v040"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)




func migrateSupply(store sdk.KVStore, cdc codec.BinaryCodec) error {
	
	var oldSupplyI v040bank.SupplyI
	err := cdc.UnmarshalInterface(store.Get(v040bank.SupplyKey), &oldSupplyI)
	if err != nil {
		return err
	}

	
	store.Delete(v040bank.SupplyKey)

	if oldSupplyI == nil {
		return nil
	}

	
	supplyStore := prefix.NewStore(store, types.SupplyKey)

	
	
	oldSupply := oldSupplyI.(*types.Supply)
	for i := range oldSupply.Total {
		coin := oldSupply.Total[i]
		coinBz, err := coin.Amount.Marshal()
		if err != nil {
			return err
		}

		supplyStore.Set([]byte(coin.Denom), coinBz)
	}

	return nil
}



func migrateBalanceKeys(store sdk.KVStore) {
	
	
	
	
	oldStore := prefix.NewStore(store, v040bank.BalancesPrefix)

	oldStoreIter := oldStore.Iterator(nil, nil)
	defer oldStoreIter.Close()

	for ; oldStoreIter.Valid(); oldStoreIter.Next() {
		addr := v040bank.AddressFromBalancesStore(oldStoreIter.Key())
		denom := oldStoreIter.Key()[v040auth.AddrLen:]
		newStoreKey := append(types.CreateAccountBalancesPrefix(addr), denom...)

		
		store.Set(newStoreKey, oldStoreIter.Value())
		oldStore.Delete(oldStoreIter.Key())
	}
}



//




func MigrateStore(ctx sdk.Context, storeKey sdk.StoreKey, cdc codec.BinaryCodec) error {
	store := ctx.KVStore(storeKey)
	migrateBalanceKeys(store)

	if err := pruneZeroBalances(store, cdc); err != nil {
		return err
	}

	if err := migrateSupply(store, cdc); err != nil {
		return err
	}

	return pruneZeroSupply(store)
}


func pruneZeroBalances(store sdk.KVStore, cdc codec.BinaryCodec) error {
	balancesStore := prefix.NewStore(store, BalancesPrefix)
	iterator := balancesStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var balance sdk.Coin
		if err := cdc.Unmarshal(iterator.Value(), &balance); err != nil {
			return err
		}

		if balance.IsZero() {
			balancesStore.Delete(iterator.Key())
		}
	}
	return nil
}


func pruneZeroSupply(store sdk.KVStore) error {
	supplyStore := prefix.NewStore(store, SupplyKey)
	iterator := supplyStore.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var amount sdk.Int
		if err := amount.Unmarshal(iterator.Value()); err != nil {
			return err
		}

		if amount.IsZero() {
			supplyStore.Delete(iterator.Key())
		}
	}

	return nil
}
