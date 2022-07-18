package cachekv_test

import (
	"testing"

	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/dbadapter"
)

var sink interface{}

const defaultValueSizeBz = 1 << 12


func benchmarkBlankParentIteratorNext(b *testing.B, keysize int) {
	mem := dbadapter.Store{DB: dbm.NewMemDB()}
	kvstore := cachekv.NewStore(mem)

	value := randSlice(defaultValueSizeBz)


	startKey := randSlice(32)


	keys := generateSequentialKeys(startKey, b.N+1)
	for _, k := range keys {
		kvstore.Set(k, value)
	}

	b.ReportAllocs()
	b.ResetTimer()

	iter := kvstore.Iterator(keys[0], keys[b.N])
	defer iter.Close()

	for _ = iter.Key(); iter.Valid(); iter.Next() {

		sink = iter
	}
}


func benchmarkBlankParentAppend(b *testing.B, keysize int) {
	mem := dbadapter.Store{DB: dbm.NewMemDB()}
	kvstore := cachekv.NewStore(mem)


	value := randSlice(32)


	startKey := randSlice(32)

	keys := generateSequentialKeys(startKey, b.N)

	b.ReportAllocs()
	b.ResetTimer()

	for _, k := range keys {
		kvstore.Set(k, value)
	}
}



func benchmarkRandomSet(b *testing.B, keysize int) {
	mem := dbadapter.Store{DB: dbm.NewMemDB()}
	kvstore := cachekv.NewStore(mem)


	value := randSlice(defaultValueSizeBz)
	keys := generateRandomKeys(keysize, b.N)

	b.ReportAllocs()
	b.ResetTimer()

	for _, k := range keys {
		kvstore.Set(k, value)
	}

	iter := kvstore.Iterator(keys[0], keys[b.N])
	defer iter.Close()

	for _ = iter.Key(); iter.Valid(); iter.Next() {

		sink = iter
	}
}





func benchmarkIteratorOnParentWithManyDeletes(b *testing.B, numDeletes int) {
	mem := dbadapter.Store{DB: dbm.NewMemDB()}


	value := randSlice(32)


	startKey := randSlice(32)
	keys := generateSequentialKeys(startKey, numDeletes)

	for _, k := range keys {
		mem.Set(k, value)
	}
	kvstore := cachekv.NewStore(mem)





	for _, k := range keys[1:] {
		kvstore.Delete(k)
	}

	b.ReportAllocs()
	b.ResetTimer()

	iter := kvstore.Iterator(keys[0], keys[b.N])
	defer iter.Close()

	for _ = iter.Key(); iter.Valid(); iter.Next() {

		sink = iter
	}
}

func BenchmarkBlankParentIteratorNextKeySize32(b *testing.B) {
	benchmarkBlankParentIteratorNext(b, 32)
}

func BenchmarkBlankParentAppendKeySize32(b *testing.B) {
	benchmarkBlankParentAppend(b, 32)
}

func BenchmarkSetKeySize32(b *testing.B) {
	benchmarkRandomSet(b, 32)
}

func BenchmarkIteratorOnParentWith1MDeletes(b *testing.B) {
	benchmarkIteratorOnParentWithManyDeletes(b, 1_000_000)
}
