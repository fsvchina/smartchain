package cachemulti

import (
	"fmt"
	"io"

	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/dbadapter"
	"github.com/cosmos/cosmos-sdk/store/types"
)








type Store struct {
	db     types.CacheKVStore
	stores map[types.StoreKey]types.CacheWrap
	keys   map[string]types.StoreKey

	traceWriter  io.Writer
	traceContext types.TraceContext

	listeners map[types.StoreKey][]types.WriteListener
}

var _ types.CacheMultiStore = Store{}




func NewFromKVStore(
	store types.KVStore, stores map[types.StoreKey]types.CacheWrapper,
	keys map[string]types.StoreKey, traceWriter io.Writer, traceContext types.TraceContext,
	listeners map[types.StoreKey][]types.WriteListener,
) Store {
	cms := Store{
		db:           cachekv.NewStore(store),
		stores:       make(map[types.StoreKey]types.CacheWrap, len(stores)),
		keys:         keys,
		traceWriter:  traceWriter,
		traceContext: traceContext,
		listeners:    listeners,
	}

	for key, store := range stores {
		var cacheWrapped types.CacheWrap
		if cms.TracingEnabled() {
			cacheWrapped = store.CacheWrapWithTrace(cms.traceWriter, cms.traceContext)
		} else {
			cacheWrapped = store.CacheWrap()
		}
		if cms.ListeningEnabled(key) {
			cms.stores[key] = cacheWrapped.CacheWrapWithListeners(key, cms.listeners[key])
		} else {
			cms.stores[key] = cacheWrapped
		}
	}

	return cms
}



func NewStore(
	db dbm.DB, stores map[types.StoreKey]types.CacheWrapper, keys map[string]types.StoreKey,
	traceWriter io.Writer, traceContext types.TraceContext, listeners map[types.StoreKey][]types.WriteListener,
) Store {

	return NewFromKVStore(dbadapter.Store{DB: db}, stores, keys, traceWriter, traceContext, listeners)
}

func newCacheMultiStoreFromCMS(cms Store) Store {
	stores := make(map[types.StoreKey]types.CacheWrapper)
	for k, v := range cms.stores {
		stores[k] = v
	}

	return NewFromKVStore(cms.db, stores, nil, cms.traceWriter, cms.traceContext, cms.listeners)
}



func (cms Store) SetTracer(w io.Writer) types.MultiStore {
	cms.traceWriter = w
	return cms
}





func (cms Store) SetTracingContext(tc types.TraceContext) types.MultiStore {
	if cms.traceContext != nil {
		for k, v := range tc {
			cms.traceContext[k] = v
		}
	} else {
		cms.traceContext = tc
	}

	return cms
}


func (cms Store) TracingEnabled() bool {
	return cms.traceWriter != nil
}


func (cms Store) AddListeners(key types.StoreKey, listeners []types.WriteListener) {
	if ls, ok := cms.listeners[key]; ok {
		cms.listeners[key] = append(ls, listeners...)
	} else {
		cms.listeners[key] = listeners
	}
}


func (cms Store) ListeningEnabled(key types.StoreKey) bool {
	if ls, ok := cms.listeners[key]; ok {
		return len(ls) != 0
	}
	return false
}


func (cms Store) GetStoreType() types.StoreType {
	return types.StoreTypeMulti
}


func (cms Store) Write() {
	cms.db.Write()
	for _, store := range cms.stores {
		store.Write()
	}
}


func (cms Store) CacheWrap() types.CacheWrap {
	return cms.CacheMultiStore().(types.CacheWrap)
}


func (cms Store) CacheWrapWithTrace(_ io.Writer, _ types.TraceContext) types.CacheWrap {
	return cms.CacheWrap()
}


func (cms Store) CacheWrapWithListeners(_ types.StoreKey, _ []types.WriteListener) types.CacheWrap {
	return cms.CacheWrap()
}


func (cms Store) CacheMultiStore() types.CacheMultiStore {
	return newCacheMultiStoreFromCMS(cms)
}



//


func (cms Store) CacheMultiStoreWithVersion(_ int64) (types.CacheMultiStore, error) {
	panic("cannot branch cached multi-store with a version")
}


func (cms Store) GetStore(key types.StoreKey) types.Store {
	s := cms.stores[key]
	if key == nil || s == nil {
		panic(fmt.Sprintf("kv store with key %v has not been registered in stores", key))
	}
	return s.(types.Store)
}


func (cms Store) GetKVStore(key types.StoreKey) types.KVStore {
	store := cms.stores[key]
	if key == nil || store == nil {
		panic(fmt.Sprintf("kv store with key %v has not been registered in stores", key))
	}
	return store.(types.KVStore)
}
