package listenkv

import (
	"io"

	"github.com/cosmos/cosmos-sdk/store/types"
)

var _ types.KVStore = &Store{}




type Store struct {
	parent         types.KVStore
	listeners      []types.WriteListener
	parentStoreKey types.StoreKey
}



func NewStore(parent types.KVStore, parentStoreKey types.StoreKey, listeners []types.WriteListener) *Store {
	return &Store{parent: parent, listeners: listeners, parentStoreKey: parentStoreKey}
}



func (s *Store) Get(key []byte) []byte {
	value := s.parent.Get(key)
	return value
}



func (s *Store) Set(key []byte, value []byte) {
	types.AssertValidKey(key)
	s.parent.Set(key, value)
	s.onWrite(false, key, value)
}



func (s *Store) Delete(key []byte) {
	s.parent.Delete(key)
	s.onWrite(true, key, nil)
}



func (s *Store) Has(key []byte) bool {
	return s.parent.Has(key)
}



func (s *Store) Iterator(start, end []byte) types.Iterator {
	return s.iterator(start, end, true)
}



func (s *Store) ReverseIterator(start, end []byte) types.Iterator {
	return s.iterator(start, end, false)
}



func (s *Store) iterator(start, end []byte, ascending bool) types.Iterator {
	var parent types.Iterator

	if ascending {
		parent = s.parent.Iterator(start, end)
	} else {
		parent = s.parent.ReverseIterator(start, end)
	}

	return newTraceIterator(parent, s.listeners)
}

type listenIterator struct {
	parent    types.Iterator
	listeners []types.WriteListener
}

func newTraceIterator(parent types.Iterator, listeners []types.WriteListener) types.Iterator {
	return &listenIterator{parent: parent, listeners: listeners}
}


func (li *listenIterator) Domain() (start []byte, end []byte) {
	return li.parent.Domain()
}


func (li *listenIterator) Valid() bool {
	return li.parent.Valid()
}


func (li *listenIterator) Next() {
	li.parent.Next()
}


func (li *listenIterator) Key() []byte {
	key := li.parent.Key()
	return key
}


func (li *listenIterator) Value() []byte {
	value := li.parent.Value()
	return value
}


func (li *listenIterator) Close() error {
	return li.parent.Close()
}


func (li *listenIterator) Error() error {
	return li.parent.Error()
}



func (s *Store) GetStoreType() types.StoreType {
	return s.parent.GetStoreType()
}



func (s *Store) CacheWrap() types.CacheWrap {
	panic("cannot CacheWrap a ListenKVStore")
}



func (s *Store) CacheWrapWithTrace(_ io.Writer, _ types.TraceContext) types.CacheWrap {
	panic("cannot CacheWrapWithTrace a ListenKVStore")
}



func (s *Store) CacheWrapWithListeners(_ types.StoreKey, _ []types.WriteListener) types.CacheWrap {
	panic("cannot CacheWrapWithListeners a ListenKVStore")
}


func (s *Store) onWrite(delete bool, key, value []byte) {
	for _, l := range s.listeners {
		l.OnWrite(s.parentStoreKey, key, value, delete)
	}
}
