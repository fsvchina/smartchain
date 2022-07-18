package cachekv

import (
	"bytes"
	"errors"

	"github.com/cosmos/cosmos-sdk/store/types"
)






//

type cacheMergeIterator struct {
	parent    types.Iterator
	cache     types.Iterator
	ascending bool
}

var _ types.Iterator = (*cacheMergeIterator)(nil)

func newCacheMergeIterator(parent, cache types.Iterator, ascending bool) *cacheMergeIterator {
	iter := &cacheMergeIterator{
		parent:    parent,
		cache:     cache,
		ascending: ascending,
	}

	return iter
}



func (iter *cacheMergeIterator) Domain() (start, end []byte) {
	startP, endP := iter.parent.Domain()
	startC, endC := iter.cache.Domain()

	if iter.compare(startP, startC) < 0 {
		start = startP
	} else {
		start = startC
	}

	if iter.compare(endP, endC) < 0 {
		end = endC
	} else {
		end = endP
	}

	return start, end
}


func (iter *cacheMergeIterator) Valid() bool {
	return iter.skipUntilExistsOrInvalid()
}


func (iter *cacheMergeIterator) Next() {
	iter.skipUntilExistsOrInvalid()
	iter.assertValid()

	
	if !iter.parent.Valid() {
		iter.cache.Next()
		return
	}

	
	if !iter.cache.Valid() {
		iter.parent.Next()
		return
	}

	
	keyP, keyC := iter.parent.Key(), iter.cache.Key()
	switch iter.compare(keyP, keyC) {
	case -1: 
		iter.parent.Next()
	case 0: 
		iter.parent.Next()
		iter.cache.Next()
	case 1: 
		iter.cache.Next()
	}
}


func (iter *cacheMergeIterator) Key() []byte {
	iter.skipUntilExistsOrInvalid()
	iter.assertValid()

	
	if !iter.parent.Valid() {
		return iter.cache.Key()
	}

	
	if !iter.cache.Valid() {
		return iter.parent.Key()
	}

	
	keyP, keyC := iter.parent.Key(), iter.cache.Key()

	cmp := iter.compare(keyP, keyC)
	switch cmp {
	case -1: 
		return keyP
	case 0: 
		return keyP
	case 1: 
		return keyC
	default:
		panic("invalid compare result")
	}
}


func (iter *cacheMergeIterator) Value() []byte {
	iter.skipUntilExistsOrInvalid()
	iter.assertValid()

	
	if !iter.parent.Valid() {
		return iter.cache.Value()
	}

	
	if !iter.cache.Valid() {
		return iter.parent.Value()
	}

	
	keyP, keyC := iter.parent.Key(), iter.cache.Key()

	cmp := iter.compare(keyP, keyC)
	switch cmp {
	case -1: 
		return iter.parent.Value()
	case 0: 
		return iter.cache.Value()
	case 1: 
		return iter.cache.Value()
	default:
		panic("invalid comparison result")
	}
}


func (iter *cacheMergeIterator) Close() error {
	if err := iter.parent.Close(); err != nil {
		return err
	}

	return iter.cache.Close()
}



func (iter *cacheMergeIterator) Error() error {
	if !iter.Valid() {
		return errors.New("invalid cacheMergeIterator")
	}

	return nil
}



func (iter *cacheMergeIterator) assertValid() {
	if err := iter.Error(); err != nil {
		panic(err)
	}
}


func (iter *cacheMergeIterator) compare(a, b []byte) int {
	if iter.ascending {
		return bytes.Compare(a, b)
	}

	return bytes.Compare(a, b) * -1
}






func (iter *cacheMergeIterator) skipCacheDeletes(until []byte) {
	for iter.cache.Valid() &&
		iter.cache.Value() == nil &&
		(until == nil || iter.compare(iter.cache.Key(), until) < 0) {
		iter.cache.Next()
	}
}




func (iter *cacheMergeIterator) skipUntilExistsOrInvalid() bool {
	for {
		
		if !iter.parent.Valid() {
			iter.skipCacheDeletes(nil)
			return iter.cache.Valid()
		}
		

		if !iter.cache.Valid() {
			return true
		}
		

		
		keyP := iter.parent.Key()
		keyC := iter.cache.Key()

		switch iter.compare(keyP, keyC) {
		case -1: 
			return true

		case 0: 
			
			valueC := iter.cache.Value()
			if valueC == nil {
				iter.parent.Next()
				iter.cache.Next()

				continue
			}
			

			return true 
		case 1: 
			
			valueC := iter.cache.Value()
			if valueC == nil {
				iter.skipCacheDeletes(keyP)
				continue
			}
			

			return true 
		}
	}
}
