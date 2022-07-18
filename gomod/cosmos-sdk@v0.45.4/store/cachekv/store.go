package cachekv

import (
	"bytes"
	"io"
	"sort"
	"sync"
	"time"

	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/internal/conv"
	"github.com/cosmos/cosmos-sdk/store/listenkv"
	"github.com/cosmos/cosmos-sdk/store/tracekv"
	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	"github.com/cosmos/cosmos-sdk/types/kv"
)



type cValue struct {
	value []byte
	dirty bool
}


type Store struct {
	mtx           sync.Mutex
	cache         map[string]*cValue
	deleted       map[string]struct{}
	unsortedCache map[string]struct{}
	sortedCache   *dbm.MemDB
	parent        types.KVStore
}

var _ types.CacheKVStore = (*Store)(nil)


func NewStore(parent types.KVStore) *Store {
	return &Store{
		cache:         make(map[string]*cValue),
		deleted:       make(map[string]struct{}),
		unsortedCache: make(map[string]struct{}),
		sortedCache:   dbm.NewMemDB(),
		parent:        parent,
	}
}


func (store *Store) GetStoreType() types.StoreType {
	return store.parent.GetStoreType()
}


func (store *Store) Get(key []byte) (value []byte) {
	store.mtx.Lock()
	defer store.mtx.Unlock()

	types.AssertValidKey(key)

	cacheValue, ok := store.cache[conv.UnsafeBytesToStr(key)]
	if !ok {
		value = store.parent.Get(key)
		store.setCacheValue(key, value, false, false)
	} else {
		value = cacheValue.value
	}

	return value
}


func (store *Store) Set(key []byte, value []byte) {
	store.mtx.Lock()
	defer store.mtx.Unlock()

	types.AssertValidKey(key)
	types.AssertValidValue(value)

	store.setCacheValue(key, value, false, true)
}


func (store *Store) Has(key []byte) bool {
	value := store.Get(key)
	return value != nil
}


func (store *Store) Delete(key []byte) {
	store.mtx.Lock()
	defer store.mtx.Unlock()
	defer telemetry.MeasureSince(time.Now(), "store", "cachekv", "delete")

	types.AssertValidKey(key)
	store.setCacheValue(key, nil, true, true)
}


func (store *Store) Write() {
	store.mtx.Lock()
	defer store.mtx.Unlock()
	defer telemetry.MeasureSince(time.Now(), "store", "cachekv", "write")



	keys := make([]string, 0, len(store.cache))

	for key, dbValue := range store.cache {
		if dbValue.dirty {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)



	for _, key := range keys {
		if store.isDeleted(key) {




			store.parent.Delete([]byte(key))
			continue
		}

		cacheValue := store.cache[key]
		if cacheValue.value != nil {

			store.parent.Set([]byte(key), cacheValue.value)
		}
	}




	for key := range store.cache {
		delete(store.cache, key)
	}
	for key := range store.deleted {
		delete(store.deleted, key)
	}
	for key := range store.unsortedCache {
		delete(store.unsortedCache, key)
	}
	store.sortedCache = dbm.NewMemDB()
}


func (store *Store) CacheWrap() types.CacheWrap {
	return NewStore(store)
}


func (store *Store) CacheWrapWithTrace(w io.Writer, tc types.TraceContext) types.CacheWrap {
	return NewStore(tracekv.NewStore(store, w, tc))
}


func (store *Store) CacheWrapWithListeners(storeKey types.StoreKey, listeners []types.WriteListener) types.CacheWrap {
	return NewStore(listenkv.NewStore(store, storeKey, listeners))
}





func (store *Store) Iterator(start, end []byte) types.Iterator {
	return store.iterator(start, end, true)
}


func (store *Store) ReverseIterator(start, end []byte) types.Iterator {
	return store.iterator(start, end, false)
}

func (store *Store) iterator(start, end []byte, ascending bool) types.Iterator {
	store.mtx.Lock()
	defer store.mtx.Unlock()

	var parent, cache types.Iterator

	if ascending {
		parent = store.parent.Iterator(start, end)
	} else {
		parent = store.parent.ReverseIterator(start, end)
	}

	store.dirtyItems(start, end)
	cache = newMemIterator(start, end, store.sortedCache, store.deleted, ascending)

	return newCacheMergeIterator(parent, cache, ascending)
}

func findStartIndex(strL []string, startQ string) int {

	if len(strL) == 0 {
		return -1
	}

	var left, right, mid int
	right = len(strL) - 1
	for left <= right {
		mid = (left + right) >> 1
		midStr := strL[mid]
		if midStr == startQ {



			for i := mid - 1; i >= 0; i-- {
				if strL[i] != midStr {
					return i + 1
				}
			}
			return 0
		}
		if midStr < startQ {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	if left >= 0 && left < len(strL) && strL[left] >= startQ {
		return left
	}
	return -1
}

func findEndIndex(strL []string, endQ string) int {
	if len(strL) == 0 {
		return -1
	}


	var left, right, mid int
	right = len(strL) - 1
	for left <= right {
		mid = (left + right) >> 1
		midStr := strL[mid]
		if midStr == endQ {



			for i := mid - 1; i >= 0; i-- {
				if strL[i] < midStr {
					return i + 1
				}
			}
			return 0
		}
		if midStr < endQ {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}


	for i := right; i >= 0; i-- {
		if strL[i] < endQ {
			return i
		}
	}

	return -1
}

type sortState int

const (
	stateUnsorted sortState = iota
	stateAlreadySorted
)


func (store *Store) dirtyItems(start, end []byte) {
	startStr, endStr := conv.UnsafeBytesToStr(start), conv.UnsafeBytesToStr(end)
	if startStr > endStr {

		return
	}

	n := len(store.unsortedCache)
	unsorted := make([]*kv.Pair, 0)






	if n < 1024 {
		for key := range store.unsortedCache {
			if dbm.IsKeyInDomain(conv.UnsafeStrToBytes(key), start, end) {
				cacheValue := store.cache[key]
				unsorted = append(unsorted, &kv.Pair{Key: []byte(key), Value: cacheValue.value})
			}
		}
		store.clearUnsortedCacheSubset(unsorted, stateUnsorted)
		return
	}



	strL := make([]string, 0, n)
	for key := range store.unsortedCache {
		strL = append(strL, key)
	}
	sort.Strings(strL)



	startIndex := findStartIndex(strL, startStr)
	endIndex := findEndIndex(strL, endStr)

	if endIndex < 0 {
		endIndex = len(strL) - 1
	}
	if startIndex < 0 {
		startIndex = 0
	}

	kvL := make([]*kv.Pair, 0)
	for i := startIndex; i <= endIndex; i++ {
		key := strL[i]
		cacheValue := store.cache[key]
		kvL = append(kvL, &kv.Pair{Key: []byte(key), Value: cacheValue.value})
	}


	store.clearUnsortedCacheSubset(kvL, stateAlreadySorted)
}

func (store *Store) clearUnsortedCacheSubset(unsorted []*kv.Pair, sortState sortState) {
	n := len(store.unsortedCache)
	if len(unsorted) == n {
		for key := range store.unsortedCache {
			delete(store.unsortedCache, key)
		}
	} else {
		for _, kv := range unsorted {
			delete(store.unsortedCache, conv.UnsafeBytesToStr(kv.Key))
		}
	}

	if sortState == stateUnsorted {
		sort.Slice(unsorted, func(i, j int) bool {
			return bytes.Compare(unsorted[i].Key, unsorted[j].Key) < 0
		})
	}

	for _, item := range unsorted {
		if item.Value == nil {


			store.sortedCache.Set(item.Key, []byte{})
			continue
		}
		err := store.sortedCache.Set(item.Key, item.Value)
		if err != nil {
			panic(err)
		}
	}
}





func (store *Store) setCacheValue(key, value []byte, deleted bool, dirty bool) {
	keyStr := conv.UnsafeBytesToStr(key)
	store.cache[keyStr] = &cValue{
		value: value,
		dirty: dirty,
	}
	if deleted {
		store.deleted[keyStr] = struct{}{}
	} else {
		delete(store.deleted, keyStr)
	}
	if dirty {
		store.unsortedCache[conv.UnsafeBytesToStr(key)] = struct{}{}
	}
}

func (store *Store) isDeleted(key string) bool {
	_, ok := store.deleted[key]
	return ok
}
