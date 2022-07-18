package types

import (
	"fmt"
)



func KVStorePrefixIteratorPaginated(kvs KVStore, prefix []byte, page, limit uint) Iterator {
	pi := &PaginatedIterator{
		Iterator: KVStorePrefixIterator(kvs, prefix),
		page:     page,
		limit:    limit,
	}
	pi.skip()
	return pi
}



func KVStoreReversePrefixIteratorPaginated(kvs KVStore, prefix []byte, page, limit uint) Iterator {
	pi := &PaginatedIterator{
		Iterator: KVStoreReversePrefixIterator(kvs, prefix),
		page:     page,
		limit:    limit,
	}
	pi.skip()
	return pi
}


type PaginatedIterator struct {
	Iterator

	page, limit uint
	iterated    uint

}

func (pi *PaginatedIterator) skip() {
	for i := (pi.page - 1) * pi.limit; i > 0 && pi.Iterator.Valid(); i-- {
		pi.Iterator.Next()
	}
}


func (pi *PaginatedIterator) Next() {
	if !pi.Valid() {
		panic(fmt.Sprintf("PaginatedIterator reached limit %d", pi.limit))
	}
	pi.Iterator.Next()
	pi.iterated++
}


func (pi *PaginatedIterator) Valid() bool {
	if pi.iterated >= pi.limit {
		return false
	}
	return pi.Iterator.Valid()
}
