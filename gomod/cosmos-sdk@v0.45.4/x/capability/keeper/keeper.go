package keeper

import (
	"fmt"
	"strings"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/capability/types"
)

type (
	
	
	
	
	//
	
	
	
	
	//
	
	
	Keeper struct {
		cdc           codec.BinaryCodec
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		capMap        map[uint64]*types.Capability
		scopedModules map[string]struct{}
		sealed        bool
	}

	
	
	
	
	
	
	ScopedKeeper struct {
		cdc      codec.BinaryCodec
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
		capMap   map[uint64]*types.Capability
		module   string
	}
)



func NewKeeper(cdc codec.BinaryCodec, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		capMap:        make(map[uint64]*types.Capability),
		scopedModules: make(map[string]struct{}),
		sealed:        false,
	}
}




func (k *Keeper) ScopeToModule(moduleName string) ScopedKeeper {
	if k.sealed {
		panic("cannot scope to module via a sealed capability keeper")
	}
	if strings.TrimSpace(moduleName) == "" {
		panic("cannot scope to an empty module name")
	}

	if _, ok := k.scopedModules[moduleName]; ok {
		panic(fmt.Sprintf("cannot create multiple scoped keepers for the same module name: %s", moduleName))
	}

	k.scopedModules[moduleName] = struct{}{}

	return ScopedKeeper{
		cdc:      k.cdc,
		storeKey: k.storeKey,
		memKey:   k.memKey,
		capMap:   k.capMap,
		module:   moduleName,
	}
}



func (k *Keeper) Seal() {
	if k.sealed {
		panic("cannot initialize and seal an already sealed capability keeper")
	}

	k.sealed = true
}






func (k *Keeper) InitMemStore(ctx sdk.Context) {
	memStore := ctx.KVStore(k.memKey)
	memStoreType := memStore.GetStoreType()
	if memStoreType != sdk.StoreTypeMemory {
		panic(fmt.Sprintf("invalid memory store type; got %s, expected: %s", memStoreType, sdk.StoreTypeMemory))
	}

	
	noGasCtx := ctx.WithBlockGasMeter(sdk.NewInfiniteGasMeter())

	
	if !k.IsInitialized(noGasCtx) {
		prefixStore := prefix.NewStore(noGasCtx.KVStore(k.storeKey), types.KeyPrefixIndexCapability)
		iterator := sdk.KVStorePrefixIterator(prefixStore, nil)

		
		defer iterator.Close()

		for ; iterator.Valid(); iterator.Next() {
			index := types.IndexFromKey(iterator.Key())

			var capOwners types.CapabilityOwners

			k.cdc.MustUnmarshal(iterator.Value(), &capOwners)
			k.InitializeCapability(noGasCtx, index, capOwners)
		}

		
		memStore := noGasCtx.KVStore(k.memKey)
		memStore.Set(types.KeyMemInitialized, []byte{1})
	}
}


func (k *Keeper) IsInitialized(ctx sdk.Context) bool {
	memStore := ctx.KVStore(k.memKey)
	return memStore.Get(types.KeyMemInitialized) != nil
}




func (k Keeper) InitializeIndex(ctx sdk.Context, index uint64) error {
	if index == 0 {
		panic("SetIndex requires index > 0")
	}
	latest := k.GetLatestIndex(ctx)
	if latest > 0 {
		panic("SetIndex requires index to not be set")
	}

	
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyIndex, types.IndexToKey(index))
	return nil
}


func (k Keeper) GetLatestIndex(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	return types.IndexFromKey(store.Get(types.KeyIndex))
}


func (k Keeper) SetOwners(ctx sdk.Context, index uint64, owners types.CapabilityOwners) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixIndexCapability)
	indexKey := types.IndexToKey(index)

	
	prefixStore.Set(indexKey, k.cdc.MustMarshal(&owners))
}


func (k Keeper) GetOwners(ctx sdk.Context, index uint64) (types.CapabilityOwners, bool) {
	prefixStore := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixIndexCapability)
	indexKey := types.IndexToKey(index)

	
	ownerBytes := prefixStore.Get(indexKey)
	if ownerBytes == nil {
		return types.CapabilityOwners{}, false
	}
	var owners types.CapabilityOwners
	k.cdc.MustUnmarshal(ownerBytes, &owners)
	return owners, true
}




func (k Keeper) InitializeCapability(ctx sdk.Context, index uint64, owners types.CapabilityOwners) {

	memStore := ctx.KVStore(k.memKey)

	cap := types.NewCapability(index)
	for _, owner := range owners.Owners {
		
		
		memStore.Set(types.FwdCapabilityKey(owner.Module, cap), []byte(owner.Name))

		
		
		
		
		memStore.Set(types.RevCapabilityKey(owner.Module, owner.Name), sdk.Uint64ToBigEndian(index))

		
		k.capMap[index] = cap
	}

}







//


func (sk ScopedKeeper) NewCapability(ctx sdk.Context, name string) (*types.Capability, error) {
	if strings.TrimSpace(name) == "" {
		return nil, sdkerrors.Wrap(types.ErrInvalidCapabilityName, "capability name cannot be empty")
	}
	store := ctx.KVStore(sk.storeKey)

	if _, ok := sk.GetCapability(ctx, name); ok {
		return nil, sdkerrors.Wrapf(types.ErrCapabilityTaken, fmt.Sprintf("module: %s, name: %s", sk.module, name))
	}

	
	index := types.IndexFromKey(store.Get(types.KeyIndex))
	cap := types.NewCapability(index)

	
	if err := sk.addOwner(ctx, cap, name); err != nil {
		return nil, err
	}

	
	store.Set(types.KeyIndex, types.IndexToKey(index+1))

	memStore := ctx.KVStore(sk.memKey)

	
	
	memStore.Set(types.FwdCapabilityKey(sk.module, cap), []byte(name))

	
	
	
	
	memStore.Set(types.RevCapabilityKey(sk.module, name), sdk.Uint64ToBigEndian(index))

	
	sk.capMap[index] = cap

	logger(ctx).Info("created new capability", "module", sk.module, "name", name)

	return cap, nil
}






//


func (sk ScopedKeeper) AuthenticateCapability(ctx sdk.Context, cap *types.Capability, name string) bool {
	if strings.TrimSpace(name) == "" || cap == nil {
		return false
	}
	return sk.GetCapabilityName(ctx, cap) == name
}






func (sk ScopedKeeper) ClaimCapability(ctx sdk.Context, cap *types.Capability, name string) error {
	if cap == nil {
		return sdkerrors.Wrap(types.ErrNilCapability, "cannot claim nil capability")
	}
	if strings.TrimSpace(name) == "" {
		return sdkerrors.Wrap(types.ErrInvalidCapabilityName, "capability name cannot be empty")
	}
	
	if err := sk.addOwner(ctx, cap, name); err != nil {
		return err
	}

	memStore := ctx.KVStore(sk.memKey)

	
	
	memStore.Set(types.FwdCapabilityKey(sk.module, cap), []byte(name))

	
	
	
	
	memStore.Set(types.RevCapabilityKey(sk.module, name), sdk.Uint64ToBigEndian(cap.GetIndex()))

	logger(ctx).Info("claimed capability", "module", sk.module, "name", name, "capability", cap.GetIndex())

	return nil
}




func (sk ScopedKeeper) ReleaseCapability(ctx sdk.Context, cap *types.Capability) error {
	if cap == nil {
		return sdkerrors.Wrap(types.ErrNilCapability, "cannot release nil capability")
	}
	name := sk.GetCapabilityName(ctx, cap)
	if len(name) == 0 {
		return sdkerrors.Wrap(types.ErrCapabilityNotOwned, sk.module)
	}

	memStore := ctx.KVStore(sk.memKey)

	
	
	memStore.Delete(types.FwdCapabilityKey(sk.module, cap))

	
	
	memStore.Delete(types.RevCapabilityKey(sk.module, name))

	
	capOwners := sk.getOwners(ctx, cap)
	capOwners.Remove(types.NewOwner(sk.module, name))

	prefixStore := prefix.NewStore(ctx.KVStore(sk.storeKey), types.KeyPrefixIndexCapability)
	indexKey := types.IndexToKey(cap.GetIndex())

	if len(capOwners.Owners) == 0 {
		
		prefixStore.Delete(indexKey)
		
		delete(sk.capMap, cap.GetIndex())
	} else {
		
		prefixStore.Set(indexKey, sk.cdc.MustMarshal(capOwners))
	}

	return nil
}




func (sk ScopedKeeper) GetCapability(ctx sdk.Context, name string) (*types.Capability, bool) {
	if strings.TrimSpace(name) == "" {
		return nil, false
	}
	memStore := ctx.KVStore(sk.memKey)

	key := types.RevCapabilityKey(sk.module, name)
	indexBytes := memStore.Get(key)
	index := sdk.BigEndianToUint64(indexBytes)

	if len(indexBytes) == 0 {
		
		
		
		
		
		

		return nil, false
	}

	cap := sk.capMap[index]
	if cap == nil {
		panic("capability found in memstore is missing from map")
	}

	return cap, true
}



func (sk ScopedKeeper) GetCapabilityName(ctx sdk.Context, cap *types.Capability) string {
	if cap == nil {
		return ""
	}
	memStore := ctx.KVStore(sk.memKey)

	return string(memStore.Get(types.FwdCapabilityKey(sk.module, cap)))
}



func (sk ScopedKeeper) GetOwners(ctx sdk.Context, name string) (*types.CapabilityOwners, bool) {
	if strings.TrimSpace(name) == "" {
		return nil, false
	}
	cap, ok := sk.GetCapability(ctx, name)
	if !ok {
		return nil, false
	}

	prefixStore := prefix.NewStore(ctx.KVStore(sk.storeKey), types.KeyPrefixIndexCapability)
	indexKey := types.IndexToKey(cap.GetIndex())

	var capOwners types.CapabilityOwners

	bz := prefixStore.Get(indexKey)
	if len(bz) == 0 {
		return nil, false
	}

	sk.cdc.MustUnmarshal(bz, &capOwners)

	return &capOwners, true
}





func (sk ScopedKeeper) LookupModules(ctx sdk.Context, name string) ([]string, *types.Capability, error) {
	if strings.TrimSpace(name) == "" {
		return nil, nil, sdkerrors.Wrap(types.ErrInvalidCapabilityName, "cannot lookup modules with empty capability name")
	}
	cap, ok := sk.GetCapability(ctx, name)
	if !ok {
		return nil, nil, sdkerrors.Wrap(types.ErrCapabilityNotFound, name)
	}

	capOwners, ok := sk.GetOwners(ctx, name)
	if !ok {
		return nil, nil, sdkerrors.Wrap(types.ErrCapabilityOwnersNotFound, name)
	}

	mods := make([]string, len(capOwners.Owners))
	for i, co := range capOwners.Owners {
		mods[i] = co.Module
	}

	return mods, cap, nil
}

func (sk ScopedKeeper) addOwner(ctx sdk.Context, cap *types.Capability, name string) error {
	prefixStore := prefix.NewStore(ctx.KVStore(sk.storeKey), types.KeyPrefixIndexCapability)
	indexKey := types.IndexToKey(cap.GetIndex())

	capOwners := sk.getOwners(ctx, cap)

	if err := capOwners.Set(types.NewOwner(sk.module, name)); err != nil {
		return err
	}

	
	prefixStore.Set(indexKey, sk.cdc.MustMarshal(capOwners))

	return nil
}

func (sk ScopedKeeper) getOwners(ctx sdk.Context, cap *types.Capability) *types.CapabilityOwners {
	prefixStore := prefix.NewStore(ctx.KVStore(sk.storeKey), types.KeyPrefixIndexCapability)
	indexKey := types.IndexToKey(cap.GetIndex())

	bz := prefixStore.Get(indexKey)

	if len(bz) == 0 {
		return types.NewCapabilityOwners()
	}

	var capOwners types.CapabilityOwners
	sk.cdc.MustUnmarshal(bz, &capOwners)
	return &capOwners
}

func logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}
