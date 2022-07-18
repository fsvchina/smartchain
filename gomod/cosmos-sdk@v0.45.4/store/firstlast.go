package store

import (
	"bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkkv "github.com/cosmos/cosmos-sdk/types/kv"
)


func First(st KVStore, start, end []byte) (kv sdkkv.Pair, ok bool) {
	iter := st.Iterator(start, end)
	if !iter.Valid() {
		return kv, false
	}
	defer iter.Close()

	return sdkkv.Pair{Key: iter.Key(), Value: iter.Value()}, true
}


func Last(st KVStore, start, end []byte) (kv sdkkv.Pair, ok bool) {
	iter := st.ReverseIterator(end, start)
	if !iter.Valid() {
		if v := st.Get(start); v != nil {
			return sdkkv.Pair{Key: sdk.CopyBytes(start), Value: sdk.CopyBytes(v)}, true
		}
		return kv, false
	}
	defer iter.Close()

	if bytes.Equal(iter.Key(), end) {

		iter.Next()
		if !iter.Valid() {
			return kv, false
		}
	}

	return sdkkv.Pair{Key: iter.Key(), Value: iter.Value()}, true
}
