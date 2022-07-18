package cachekv_test

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	tmrand "github.com/tendermint/tendermint/libs/rand"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/dbadapter"
	"github.com/cosmos/cosmos-sdk/store/types"
)

func newCacheKVStore() types.CacheKVStore {
	mem := dbadapter.Store{DB: dbm.NewMemDB()}
	return cachekv.NewStore(mem)
}

func keyFmt(i int) []byte { return bz(fmt.Sprintf("key%0.8d", i)) }
func valFmt(i int) []byte { return bz(fmt.Sprintf("value%0.8d", i)) }

func TestCacheKVStore(t *testing.T) {
	mem := dbadapter.Store{DB: dbm.NewMemDB()}
	st := cachekv.NewStore(mem)

	require.Empty(t, st.Get(keyFmt(1)), "Expected `key1` to be empty")


	mem.Set(keyFmt(1), valFmt(1))
	st.Set(keyFmt(1), valFmt(1))
	require.Equal(t, valFmt(1), st.Get(keyFmt(1)))


	st.Set(keyFmt(1), valFmt(2))
	require.Equal(t, valFmt(2), st.Get(keyFmt(1)))
	require.Equal(t, valFmt(1), mem.Get(keyFmt(1)))


	st.Write()
	require.Equal(t, valFmt(2), mem.Get(keyFmt(1)))
	require.Equal(t, valFmt(2), st.Get(keyFmt(1)))


	st.Write()
	st.Write()
	require.Equal(t, valFmt(2), mem.Get(keyFmt(1)))
	require.Equal(t, valFmt(2), st.Get(keyFmt(1)))


	st = cachekv.NewStore(mem)
	require.Equal(t, valFmt(2), st.Get(keyFmt(1)))


	st = cachekv.NewStore(mem)
	st.Delete(keyFmt(1))
	require.Empty(t, st.Get(keyFmt(1)))
	require.Equal(t, mem.Get(keyFmt(1)), valFmt(2))


	st.Write()
	require.Empty(t, st.Get(keyFmt(1)), "Expected `key1` to be empty")
	require.Empty(t, mem.Get(keyFmt(1)), "Expected `key1` to be empty")
}

func TestCacheKVStoreNoNilSet(t *testing.T) {
	mem := dbadapter.Store{DB: dbm.NewMemDB()}
	st := cachekv.NewStore(mem)
	require.Panics(t, func() { st.Set([]byte("key"), nil) }, "setting a nil value should panic")
	require.Panics(t, func() { st.Set(nil, []byte("value")) }, "setting a nil key should panic")
	require.Panics(t, func() { st.Set([]byte(""), []byte("value")) }, "setting an empty key should panic")
}

func TestCacheKVStoreNested(t *testing.T) {
	mem := dbadapter.Store{DB: dbm.NewMemDB()}
	st := cachekv.NewStore(mem)


	st.Set(keyFmt(1), valFmt(1))
	require.Empty(t, mem.Get(keyFmt(1)))
	require.Equal(t, valFmt(1), st.Get(keyFmt(1)))


	st2 := cachekv.NewStore(st)
	require.Equal(t, valFmt(1), st2.Get(keyFmt(1)))


	st2.Set(keyFmt(1), valFmt(3))
	require.Equal(t, []byte(nil), mem.Get(keyFmt(1)))
	require.Equal(t, valFmt(1), st.Get(keyFmt(1)))
	require.Equal(t, valFmt(3), st2.Get(keyFmt(1)))


	st2.Write()
	require.Equal(t, []byte(nil), mem.Get(keyFmt(1)))
	require.Equal(t, valFmt(3), st.Get(keyFmt(1)))


	st.Write()
	require.Equal(t, valFmt(3), mem.Get(keyFmt(1)))
}

func TestCacheKVIteratorBounds(t *testing.T) {
	st := newCacheKVStore()


	nItems := 5
	for i := 0; i < nItems; i++ {
		st.Set(keyFmt(i), valFmt(i))
	}


	itr := st.Iterator(nil, nil)
	var i = 0
	for ; itr.Valid(); itr.Next() {
		k, v := itr.Key(), itr.Value()
		require.Equal(t, keyFmt(i), k)
		require.Equal(t, valFmt(i), v)
		i++
	}
	require.Equal(t, nItems, i)


	itr = st.Iterator(bz("money"), nil)
	i = 0
	for ; itr.Valid(); itr.Next() {
		i++
	}
	require.Equal(t, 0, i)


	itr = st.Iterator(keyFmt(0), keyFmt(3))
	i = 0
	for ; itr.Valid(); itr.Next() {
		k, v := itr.Key(), itr.Value()
		require.Equal(t, keyFmt(i), k)
		require.Equal(t, valFmt(i), v)
		i++
	}
	require.Equal(t, 3, i)


	itr = st.Iterator(keyFmt(2), keyFmt(4))
	i = 2
	for ; itr.Valid(); itr.Next() {
		k, v := itr.Key(), itr.Value()
		require.Equal(t, keyFmt(i), k)
		require.Equal(t, valFmt(i), v)
		i++
	}
	require.Equal(t, 4, i)
}

func TestCacheKVMergeIteratorBasics(t *testing.T) {
	st := newCacheKVStore()


	k, v := keyFmt(0), valFmt(0)
	st.Set(k, v)
	st.Delete(k)
	assertIterateDomain(t, st, 0)


	st.Set(k, v)
	assertIterateDomain(t, st, 1)


	st.Write()
	assertIterateDomain(t, st, 1)


	st.Delete(k)
	assertIterateDomain(t, st, 0)


	st.Write()
	assertIterateDomain(t, st, 0)


	k1, v1 := keyFmt(1), valFmt(1)
	st.Set(k, v)
	st.Set(k1, v1)
	assertIterateDomain(t, st, 2)


	st.Write()
	assertIterateDomain(t, st, 2)


	st.Delete(k1)
	assertIterateDomain(t, st, 1)


	st.Write()
	assertIterateDomain(t, st, 1)


	st.Delete(k)
	assertIterateDomain(t, st, 0)
}

func TestCacheKVMergeIteratorDeleteLast(t *testing.T) {
	st := newCacheKVStore()


	nItems := 5
	for i := 0; i < nItems; i++ {
		st.Set(keyFmt(i), valFmt(i))
	}
	st.Write()


	for i := nItems; i < nItems*2; i++ {
		st.Set(keyFmt(i), valFmt(i))
	}


	assertIterateDomain(t, st, nItems*2)


	for i := 0; i < nItems*2; i++ {
		last := nItems*2 - 1 - i
		st.Delete(keyFmt(last))
		assertIterateDomain(t, st, last)
	}
}

func TestCacheKVMergeIteratorDeletes(t *testing.T) {
	st := newCacheKVStore()
	truth := dbm.NewMemDB()


	nItems := 10
	for i := 0; i < nItems; i++ {
		doOp(t, st, truth, opSet, i)
	}
	st.Write()


	for i := 0; i < nItems; i += 2 {
		doOp(t, st, truth, opDel, i)
		assertIterateDomainCompare(t, st, truth)
	}


	st = newCacheKVStore()
	truth = dbm.NewMemDB()


	for i := 0; i < nItems; i++ {
		doOp(t, st, truth, opSet, i)
	}
	st.Write()


	for i := 1; i < nItems; i += 2 {
		doOp(t, st, truth, opDel, i)
		assertIterateDomainCompare(t, st, truth)
	}
}

func TestCacheKVMergeIteratorChunks(t *testing.T) {
	st := newCacheKVStore()


	truth := dbm.NewMemDB()


	setRange(t, st, truth, 0, 20)
	setRange(t, st, truth, 40, 60)
	st.Write()


	setRange(t, st, truth, 20, 40)
	setRange(t, st, truth, 60, 80)
	assertIterateDomainCheck(t, st, truth, []keyRange{{0, 80}})


	deleteRange(t, st, truth, 15, 25)
	assertIterateDomainCheck(t, st, truth, []keyRange{{0, 15}, {25, 80}})


	deleteRange(t, st, truth, 35, 45)
	assertIterateDomainCheck(t, st, truth, []keyRange{{0, 15}, {25, 35}, {45, 80}})


	st.Write()
	setRange(t, st, truth, 38, 42)
	deleteRange(t, st, truth, 40, 43)
	assertIterateDomainCheck(t, st, truth, []keyRange{{0, 15}, {25, 35}, {38, 40}, {45, 80}})
}

func TestCacheKVMergeIteratorRandom(t *testing.T) {
	st := newCacheKVStore()
	truth := dbm.NewMemDB()

	start, end := 25, 975
	max := 1000
	setRange(t, st, truth, start, end)


	for i := 0; i < 2000; i++ {
		doRandomOp(t, st, truth, max)
		assertIterateDomainCompare(t, st, truth)
	}
}




const (
	opSet      = 0
	opSetRange = 1
	opDel      = 2
	opDelRange = 3
	opWrite    = 4

	totalOps = 5
)

func randInt(n int) int {
	return tmrand.NewRand().Int() % n
}


func doOp(t *testing.T, st types.CacheKVStore, truth dbm.DB, op int, args ...int) {
	switch op {
	case opSet:
		k := args[0]
		st.Set(keyFmt(k), valFmt(k))
		err := truth.Set(keyFmt(k), valFmt(k))
		require.NoError(t, err)
	case opSetRange:
		start := args[0]
		end := args[1]
		setRange(t, st, truth, start, end)
	case opDel:
		k := args[0]
		st.Delete(keyFmt(k))
		err := truth.Delete(keyFmt(k))
		require.NoError(t, err)
	case opDelRange:
		start := args[0]
		end := args[1]
		deleteRange(t, st, truth, start, end)
	case opWrite:
		st.Write()
	}
}

func doRandomOp(t *testing.T, st types.CacheKVStore, truth dbm.DB, maxKey int) {
	r := randInt(totalOps)
	switch r {
	case opSet:
		k := randInt(maxKey)
		st.Set(keyFmt(k), valFmt(k))
		err := truth.Set(keyFmt(k), valFmt(k))
		require.NoError(t, err)
	case opSetRange:
		start := randInt(maxKey - 2)
		end := randInt(maxKey-start) + start
		setRange(t, st, truth, start, end)
	case opDel:
		k := randInt(maxKey)
		st.Delete(keyFmt(k))
		err := truth.Delete(keyFmt(k))
		require.NoError(t, err)
	case opDelRange:
		start := randInt(maxKey - 2)
		end := randInt(maxKey-start) + start
		deleteRange(t, st, truth, start, end)
	case opWrite:
		st.Write()
	}
}




func assertIterateDomain(t *testing.T, st types.KVStore, expectedN int) {
	itr := st.Iterator(nil, nil)
	var i = 0
	for ; itr.Valid(); itr.Next() {
		k, v := itr.Key(), itr.Value()
		require.Equal(t, keyFmt(i), k)
		require.Equal(t, valFmt(i), v)
		i++
	}
	require.Equal(t, expectedN, i)
}

func assertIterateDomainCheck(t *testing.T, st types.KVStore, mem dbm.DB, r []keyRange) {

	itr := st.Iterator(nil, nil)
	itr2, err := mem.Iterator(nil, nil)
	require.NoError(t, err)

	krc := newKeyRangeCounter(r)
	i := 0

	for ; krc.valid(); krc.next() {
		require.True(t, itr.Valid())
		require.True(t, itr2.Valid())


		k, v := itr.Key(), itr.Value()
		k2, v2 := itr2.Key(), itr2.Value()
		require.Equal(t, k, k2)
		require.Equal(t, v, v2)


		require.Equal(t, k, keyFmt(krc.key()))

		itr.Next()
		itr2.Next()
		i++
	}

	require.False(t, itr.Valid())
	require.False(t, itr2.Valid())
}

func assertIterateDomainCompare(t *testing.T, st types.KVStore, mem dbm.DB) {

	itr := st.Iterator(nil, nil)
	itr2, err := mem.Iterator(nil, nil)
	require.NoError(t, err)
	checkIterators(t, itr, itr2)
	checkIterators(t, itr2, itr)
}

func checkIterators(t *testing.T, itr, itr2 types.Iterator) {
	for ; itr.Valid(); itr.Next() {
		require.True(t, itr2.Valid())
		k, v := itr.Key(), itr.Value()
		k2, v2 := itr2.Key(), itr2.Value()
		require.Equal(t, k, k2)
		require.Equal(t, v, v2)
		itr2.Next()
	}
	require.False(t, itr.Valid())
	require.False(t, itr2.Valid())
}



func setRange(t *testing.T, st types.KVStore, mem dbm.DB, start, end int) {
	for i := start; i < end; i++ {
		st.Set(keyFmt(i), valFmt(i))
		err := mem.Set(keyFmt(i), valFmt(i))
		require.NoError(t, err)
	}
}

func deleteRange(t *testing.T, st types.KVStore, mem dbm.DB, start, end int) {
	for i := start; i < end; i++ {
		st.Delete(keyFmt(i))
		err := mem.Delete(keyFmt(i))
		require.NoError(t, err)
	}
}



type keyRange struct {
	start int
	end   int
}

func (kr keyRange) len() int {
	return kr.end - kr.start
}

func newKeyRangeCounter(kr []keyRange) *keyRangeCounter {
	return &keyRangeCounter{keyRanges: kr}
}


type keyRangeCounter struct {
	rangeIdx  int
	idx       int
	keyRanges []keyRange
}

func (krc *keyRangeCounter) valid() bool {
	maxRangeIdx := len(krc.keyRanges) - 1
	maxRange := krc.keyRanges[maxRangeIdx]


	if krc.rangeIdx <= maxRangeIdx &&
		krc.idx < maxRange.len() {
		return true
	}

	return false
}

func (krc *keyRangeCounter) next() {
	thisKeyRange := krc.keyRanges[krc.rangeIdx]
	if krc.idx == thisKeyRange.len()-1 {
		krc.rangeIdx++
		krc.idx = 0
	} else {
		krc.idx++
	}
}

func (krc *keyRangeCounter) key() int {
	thisKeyRange := krc.keyRanges[krc.rangeIdx]
	return thisKeyRange.start + krc.idx
}



func bz(s string) []byte { return []byte(s) }

func BenchmarkCacheKVStoreGetNoKeyFound(b *testing.B) {
	b.ReportAllocs()
	st := newCacheKVStore()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		st.Get([]byte{byte((i & 0xFF0000) >> 16), byte((i & 0xFF00) >> 8), byte(i & 0xFF)})
	}
}

func BenchmarkCacheKVStoreGetKeyFound(b *testing.B) {
	b.ReportAllocs()
	st := newCacheKVStore()
	for i := 0; i < b.N; i++ {
		arr := []byte{byte((i & 0xFF0000) >> 16), byte((i & 0xFF00) >> 8), byte(i & 0xFF)}
		st.Set(arr, arr)
	}
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		st.Get([]byte{byte((i & 0xFF0000) >> 16), byte((i & 0xFF00) >> 8), byte(i & 0xFF)})
	}
}
