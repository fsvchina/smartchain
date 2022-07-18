package cache

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/types"

	lru "github.com/hashicorp/golang-lru"
)

var (
	_ types.CommitKVStore             = (*CommitKVStoreCache)(nil)
	_ types.MultiStorePersistentCache = (*CommitKVStoreCacheManager)(nil)



	DefaultCommitKVStoreCacheSize uint = 1000
)

type (






	CommitKVStoreCache struct {
		types.CommitKVStore
		cache *lru.ARCCache
	}





	CommitKVStoreCacheManager struct {
		cacheSize uint
		caches    map[string]types.CommitKVStore
	}
)

func NewCommitKVStoreCache(store types.CommitKVStore, size uint) *CommitKVStoreCache {
	cache, err := lru.NewARC(int(size))
	if err != nil {
		panic(fmt.Errorf("failed to create KVStore cache: %s", err))
	}

	return &CommitKVStoreCache{
		CommitKVStore: store,
		cache:         cache,
	}
}

func NewCommitKVStoreCacheManager(size uint) *CommitKVStoreCacheManager {
	return &CommitKVStoreCacheManager{
		cacheSize: size,
		caches:    make(map[string]types.CommitKVStore),
	}
}




func (cmgr *CommitKVStoreCacheManager) GetStoreCache(key types.StoreKey, store types.CommitKVStore) types.CommitKVStore {
	if cmgr.caches[key.Name()] == nil {
		cmgr.caches[key.Name()] = NewCommitKVStoreCache(store, cmgr.cacheSize)
	}

	return cmgr.caches[key.Name()]
}


func (cmgr *CommitKVStoreCacheManager) Unwrap(key types.StoreKey) types.CommitKVStore {
	if ckv, ok := cmgr.caches[key.Name()]; ok {
		return ckv.(*CommitKVStoreCache).CommitKVStore
	}

	return nil
}


func (cmgr *CommitKVStoreCacheManager) Reset() {



	for key := range cmgr.caches {
		delete(cmgr.caches, key)
	}
}


func (ckv *CommitKVStoreCache) CacheWrap() types.CacheWrap {
	return cachekv.NewStore(ckv)
}




func (ckv *CommitKVStoreCache) Get(key []byte) []byte {
	types.AssertValidKey(key)

	keyStr := string(key)
	valueI, ok := ckv.cache.Get(keyStr)
	if ok {

		return valueI.([]byte)
	}


	value := ckv.CommitKVStore.Get(key)
	ckv.cache.Add(keyStr, value)

	return value
}



func (ckv *CommitKVStoreCache) Set(key, value []byte) {
	types.AssertValidKey(key)
	types.AssertValidValue(value)

	ckv.cache.Add(string(key), value)
	ckv.CommitKVStore.Set(key, value)
}



func (ckv *CommitKVStoreCache) Delete(key []byte) {
	ckv.cache.Remove(string(key))
	ckv.CommitKVStore.Delete(key)
}
