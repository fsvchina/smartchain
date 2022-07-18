package cachekv

import (
	"bytes"

	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/types"
)




type memIterator struct {
	types.Iterator

	lastKey []byte
	deleted map[string]struct{}
}

func newMemIterator(start, end []byte, items *dbm.MemDB, deleted map[string]struct{}, ascending bool) *memIterator {
	var iter types.Iterator
	var err error

	if ascending {
		iter, err = items.Iterator(start, end)
	} else {
		iter, err = items.ReverseIterator(start, end)
	}

	if err != nil {
		panic(err)
	}

	return &memIterator{
		Iterator: iter,

		lastKey: nil,
		deleted: deleted,
	}
}

func (mi *memIterator) Value() []byte {
	key := mi.Iterator.Key()





	reCallingOnOldLastKey := (mi.lastKey != nil) && bytes.Equal(key, mi.lastKey)
	if _, ok := mi.deleted[string(key)]; ok && !reCallingOnOldLastKey {
		return nil
	}
	mi.lastKey = key
	return mi.Iterator.Value()
}
