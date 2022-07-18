package types

import (
	fmt "fmt"
	"sort"
	"strings"

	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

type (
	PruningOptions = types.PruningOptions
)

type (
	Store                     = types.Store
	Committer                 = types.Committer
	CommitStore               = types.CommitStore
	Queryable                 = types.Queryable
	MultiStore                = types.MultiStore
	CacheMultiStore           = types.CacheMultiStore
	CommitMultiStore          = types.CommitMultiStore
	MultiStorePersistentCache = types.MultiStorePersistentCache
	KVStore                   = types.KVStore
	Iterator                  = types.Iterator
)



type StoreDecoderRegistry map[string]func(kvA, kvB kv.Pair) string


func KVStorePrefixIterator(kvs KVStore, prefix []byte) Iterator {
	return types.KVStorePrefixIterator(kvs, prefix)
}


func KVStoreReversePrefixIterator(kvs KVStore, prefix []byte) Iterator {
	return types.KVStoreReversePrefixIterator(kvs, prefix)
}



func KVStorePrefixIteratorPaginated(kvs KVStore, prefix []byte, page, limit uint) Iterator {
	return types.KVStorePrefixIteratorPaginated(kvs, prefix, page, limit)
}



func KVStoreReversePrefixIteratorPaginated(kvs KVStore, prefix []byte, page, limit uint) Iterator {
	return types.KVStoreReversePrefixIteratorPaginated(kvs, prefix, page, limit)
}



func DiffKVStores(a KVStore, b KVStore, prefixesToSkip [][]byte) (kvAs, kvBs []kv.Pair) {
	return types.DiffKVStores(a, b, prefixesToSkip)
}

type (
	CacheKVStore  = types.CacheKVStore
	CommitKVStore = types.CommitKVStore
	CacheWrap     = types.CacheWrap
	CacheWrapper  = types.CacheWrapper
	CommitID      = types.CommitID
)

type StoreType = types.StoreType

const (
	StoreTypeMulti     = types.StoreTypeMulti
	StoreTypeDB        = types.StoreTypeDB
	StoreTypeIAVL      = types.StoreTypeIAVL
	StoreTypeTransient = types.StoreTypeTransient
	StoreTypeMemory    = types.StoreTypeMemory
)

type (
	StoreKey          = types.StoreKey
	CapabilityKey     = types.CapabilityKey
	KVStoreKey        = types.KVStoreKey
	TransientStoreKey = types.TransientStoreKey
	MemoryStoreKey    = types.MemoryStoreKey
)



func assertNoPrefix(keys []string) {
	sorted := make([]string, len(keys))
	copy(sorted, keys)
	sort.Strings(sorted)
	for i := 1; i < len(sorted); i++ {
		if strings.HasPrefix(sorted[i], sorted[i-1]) {
			panic(fmt.Sprint("Potential key collision between KVStores:", sorted[i], " - ", sorted[i-1]))
		}
	}
}


func NewKVStoreKey(name string) *KVStoreKey {
	return types.NewKVStoreKey(name)
}




func NewKVStoreKeys(names ...string) map[string]*KVStoreKey {
	assertNoPrefix(names)
	keys := make(map[string]*KVStoreKey, len(names))
	for _, n := range names {
		keys[n] = NewKVStoreKey(n)
	}

	return keys
}



func NewTransientStoreKey(name string) *TransientStoreKey {
	return types.NewTransientStoreKey(name)
}





func NewTransientStoreKeys(names ...string) map[string]*TransientStoreKey {
	assertNoPrefix(names)
	keys := make(map[string]*TransientStoreKey)
	for _, n := range names {
		keys[n] = NewTransientStoreKey(n)
	}

	return keys
}





func NewMemoryStoreKeys(names ...string) map[string]*MemoryStoreKey {
	assertNoPrefix(names)
	keys := make(map[string]*MemoryStoreKey)
	for _, n := range names {
		keys[n] = types.NewMemoryStoreKey(n)
	}

	return keys
}




func PrefixEndBytes(prefix []byte) []byte {
	return types.PrefixEndBytes(prefix)
}



func InclusiveEndBytes(inclusiveBytes []byte) (exclusiveBytes []byte) {
	return types.InclusiveEndBytes(inclusiveBytes)
}




type KVPair = types.KVPair





type TraceContext = types.TraceContext



type (
	Gas       = types.Gas
	GasMeter  = types.GasMeter
	GasConfig = types.GasConfig
)

func NewGasMeter(limit Gas) GasMeter {
	return types.NewGasMeter(limit)
}

type (
	ErrorOutOfGas    = types.ErrorOutOfGas
	ErrorGasOverflow = types.ErrorGasOverflow
)

func NewInfiniteGasMeter() GasMeter {
	return types.NewInfiniteGasMeter()
}
