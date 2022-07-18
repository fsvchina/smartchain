package rootmulti

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/codec"
	codecTypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/store/cachemulti"
	"github.com/cosmos/cosmos-sdk/store/iavl"
	sdkmaps "github.com/cosmos/cosmos-sdk/store/internal/maps"
	"github.com/cosmos/cosmos-sdk/store/listenkv"
	"github.com/cosmos/cosmos-sdk/store/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func TestStoreType(t *testing.T) {
	db := dbm.NewMemDB()
	store := NewStore(db)
	store.MountStoreWithDB(types.NewKVStoreKey("store1"), types.StoreTypeIAVL, db)
}

func TestGetCommitKVStore(t *testing.T) {
	var db dbm.DB = dbm.NewMemDB()
	ms := newMultiStoreWithMounts(db, types.PruneDefault)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	key := ms.keysByName["store1"]

	store1 := ms.GetCommitKVStore(key)
	require.NotNil(t, store1)
	require.IsType(t, &iavl.Store{}, store1)

	store2 := ms.GetCommitStore(key)
	require.NotNil(t, store2)
	require.IsType(t, &iavl.Store{}, store2)
}

func TestStoreMount(t *testing.T) {
	db := dbm.NewMemDB()
	store := NewStore(db)

	key1 := types.NewKVStoreKey("store1")
	key2 := types.NewKVStoreKey("store2")
	dup1 := types.NewKVStoreKey("store1")

	require.NotPanics(t, func() { store.MountStoreWithDB(key1, types.StoreTypeIAVL, db) })
	require.NotPanics(t, func() { store.MountStoreWithDB(key2, types.StoreTypeIAVL, db) })

	require.Panics(t, func() { store.MountStoreWithDB(key1, types.StoreTypeIAVL, db) })
	require.Panics(t, func() { store.MountStoreWithDB(nil, types.StoreTypeIAVL, db) })
	require.Panics(t, func() { store.MountStoreWithDB(dup1, types.StoreTypeIAVL, db) })
}

func TestCacheMultiStore(t *testing.T) {
	var db dbm.DB = dbm.NewMemDB()
	ms := newMultiStoreWithMounts(db, types.PruneNothing)

	cacheMulti := ms.CacheMultiStore()
	require.IsType(t, cachemulti.Store{}, cacheMulti)
}

func TestCacheMultiStoreWithVersion(t *testing.T) {
	var db dbm.DB = dbm.NewMemDB()
	ms := newMultiStoreWithMounts(db, types.PruneNothing)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	commitID := types.CommitID{}
	checkStore(t, ms, commitID, commitID)

	k, v := []byte("wind"), []byte("blows")

	store1 := ms.GetStoreByName("store1").(types.KVStore)
	store1.Set(k, v)

	cID := ms.Commit()
	require.Equal(t, int64(1), cID.Version)


	_, err = ms.CacheMultiStoreWithVersion(cID.Version + 1)
	require.NoError(t, err)


	cms, err := ms.CacheMultiStoreWithVersion(cID.Version)
	require.NoError(t, err)


	kvStore := cms.GetKVStore(ms.keysByName["store1"])
	require.NotNil(t, kvStore)
	require.Equal(t, kvStore.Get(k), v)


	require.Panics(t, func() {
		kvStore.Set(k, []byte("newValue"))
		cms.Write()
	})
}

func TestHashStableWithEmptyCommit(t *testing.T) {
	var db dbm.DB = dbm.NewMemDB()
	ms := newMultiStoreWithMounts(db, types.PruneNothing)
	err := ms.LoadLatestVersion()
	require.Nil(t, err)

	commitID := types.CommitID{}
	checkStore(t, ms, commitID, commitID)

	k, v := []byte("wind"), []byte("blows")

	store1 := ms.GetStoreByName("store1").(types.KVStore)
	store1.Set(k, v)

	cID := ms.Commit()
	require.Equal(t, int64(1), cID.Version)
	hash := cID.Hash


	cID = ms.Commit()
	require.Equal(t, int64(2), cID.Version)
	require.Equal(t, hash, cID.Hash)
}

func TestMultistoreCommitLoad(t *testing.T) {
	var db dbm.DB = dbm.NewMemDB()
	store := newMultiStoreWithMounts(db, types.PruneNothing)
	err := store.LoadLatestVersion()
	require.Nil(t, err)


	commitID := types.CommitID{}
	checkStore(t, store, commitID, commitID)


	s1 := store.GetStoreByName("store1")
	require.NotNil(t, s1)
	s3 := store.GetStoreByName("store3")
	require.NotNil(t, s3)
	s77 := store.GetStoreByName("store77")
	require.Nil(t, s77)


	nCommits := int64(3)
	for i := int64(0); i < nCommits; i++ {
		commitID = store.Commit()
		expectedCommitID := getExpectedCommitID(store, i+1)
		checkStore(t, store, expectedCommitID, commitID)
	}


	store = newMultiStoreWithMounts(db, types.PruneNothing)
	err = store.LoadLatestVersion()
	require.Nil(t, err)
	commitID = getExpectedCommitID(store, nCommits)
	checkStore(t, store, commitID, commitID)


	commitID = store.Commit()
	expectedCommitID := getExpectedCommitID(store, nCommits+1)
	checkStore(t, store, expectedCommitID, commitID)


	ver := nCommits - 1
	store = newMultiStoreWithMounts(db, types.PruneNothing)
	err = store.LoadVersion(ver)
	require.Nil(t, err)
	commitID = getExpectedCommitID(store, ver)
	checkStore(t, store, commitID, commitID)
}

func TestMultistoreLoadWithUpgrade(t *testing.T) {
	var db dbm.DB = dbm.NewMemDB()
	store := newMultiStoreWithMounts(db, types.PruneNothing)
	err := store.LoadLatestVersion()
	require.Nil(t, err)


	k1, v1 := []byte("first"), []byte("store")
	s1, _ := store.GetStoreByName("store1").(types.KVStore)
	require.NotNil(t, s1)
	s1.Set(k1, v1)

	k2, v2 := []byte("second"), []byte("restore")
	s2, _ := store.GetStoreByName("store2").(types.KVStore)
	require.NotNil(t, s2)
	s2.Set(k2, v2)

	k3, v3 := []byte("third"), []byte("dropped")
	s3, _ := store.GetStoreByName("store3").(types.KVStore)
	require.NotNil(t, s3)
	s3.Set(k3, v3)

	s4, _ := store.GetStoreByName("store4").(types.KVStore)
	require.Nil(t, s4)


	commitID := store.Commit()
	expectedCommitID := getExpectedCommitID(store, 1)
	checkStore(t, store, expectedCommitID, commitID)

	ci, err := getCommitInfo(db, 1)
	require.NoError(t, err)
	require.Equal(t, int64(1), ci.Version)
	require.Equal(t, 3, len(ci.StoreInfos))
	checkContains(t, ci.StoreInfos, []string{"store1", "store2", "store3"})


	store = newMultiStoreWithMounts(db, types.PruneNothing)

	err = store.LoadLatestVersion()
	require.Nil(t, err)
	commitID = getExpectedCommitID(store, 1)
	checkStore(t, store, commitID, commitID)


	s2, _ = store.GetStoreByName("store2").(types.KVStore)
	require.NotNil(t, s2)
	require.Equal(t, v2, s2.Get(k2))


	restore, upgrades := newMultiStoreWithModifiedMounts(db, types.PruneNothing)
	err = restore.LoadLatestVersionAndUpgrade(upgrades)
	require.Nil(t, err)


	s1, _ = restore.GetStoreByName("store1").(types.KVStore)
	require.NotNil(t, s1)
	require.Equal(t, v1, s1.Get(k1))


	s3, _ = restore.GetStoreByName("store3").(types.KVStore)
	require.NotNil(t, s3)
	require.Nil(t, s3.Get(k3))


	s4, _ = restore.GetStoreByName("store4").(types.KVStore)
	require.NotNil(t, s4)

	iterator := s4.Iterator(nil, nil)

	values := 0
	for ; iterator.Valid(); iterator.Next() {
		values += 1
	}
	require.Zero(t, values)

	require.NoError(t, iterator.Close())


	k4, v4 := []byte("fourth"), []byte("created")
	s4.Set(k4, v4)


	st2 := restore.GetStoreByName("store2")
	require.Nil(t, st2)


	rs2, _ := restore.GetStoreByName("restore2").(types.KVStore)
	require.NotNil(t, rs2)
	require.Equal(t, v2, rs2.Get(k2))


	migratedID := restore.Commit()
	require.Equal(t, migratedID.Version, int64(2))

	reload, _ := newMultiStoreWithModifiedMounts(db, types.PruneNothing)
	err = reload.LoadLatestVersion()
	require.Nil(t, err)
	require.Equal(t, migratedID, reload.LastCommitID())


	rl1, _ := reload.GetStoreByName("store1").(types.KVStore)
	require.NotNil(t, rl1)
	require.Equal(t, v1, rl1.Get(k1))

	rl2, _ := reload.GetStoreByName("restore2").(types.KVStore)
	require.NotNil(t, rl2)
	require.Equal(t, v2, rl2.Get(k2))

	rl4, _ := reload.GetStoreByName("store4").(types.KVStore)
	require.NotNil(t, rl4)
	require.Equal(t, v4, rl4.Get(k4))


	ci, err = getCommitInfo(db, 2)
	require.NoError(t, err)
	require.Equal(t, int64(2), ci.Version)
	require.Equal(t, 4, len(ci.StoreInfos), ci.StoreInfos)
	checkContains(t, ci.StoreInfos, []string{"store1", "restore2", "store3", "store4"})
}

func TestParsePath(t *testing.T) {
	_, _, err := parsePath("foo")
	require.Error(t, err)

	store, subpath, err := parsePath("/foo")
	require.NoError(t, err)
	require.Equal(t, store, "foo")
	require.Equal(t, subpath, "")

	store, subpath, err = parsePath("/fizz/bang/baz")
	require.NoError(t, err)
	require.Equal(t, store, "fizz")
	require.Equal(t, subpath, "/bang/baz")

	substore, subsubpath, err := parsePath(subpath)
	require.NoError(t, err)
	require.Equal(t, substore, "bang")
	require.Equal(t, subsubpath, "/baz")

}

func TestMultiStoreRestart(t *testing.T) {
	db := dbm.NewMemDB()
	pruning := types.PruningOptions{
		KeepRecent: 2,
		KeepEvery:  3,
		Interval:   1,
	}
	multi := newMultiStoreWithMounts(db, pruning)
	err := multi.LoadLatestVersion()
	require.Nil(t, err)

	initCid := multi.LastCommitID()

	k, v := "wind", "blows"
	k2, v2 := "water", "flows"
	k3, v3 := "fire", "burns"

	for i := 1; i < 3; i++ {

		store1 := multi.GetStoreByName("store1").(types.KVStore)
		store1.Set([]byte(k), []byte(fmt.Sprintf("%s:%d", v, i)))


		store2 := multi.GetStoreByName("store2").(types.KVStore)
		store2.Set([]byte(k2), []byte(fmt.Sprintf("%s:%d", v2, i)))


		store3 := multi.GetStoreByName("store3").(types.KVStore)
		store3.Set([]byte(k3), []byte(fmt.Sprintf("%s:%d", v3, i)))

		multi.Commit()

		cinfo, err := getCommitInfo(multi.db, int64(i))
		require.NoError(t, err)
		require.Equal(t, int64(i), cinfo.Version)
	}


	store1 := multi.GetStoreByName("store1").(types.KVStore)
	store1.Set([]byte(k), []byte(fmt.Sprintf("%s:%d", v, 3)))


	store2 := multi.GetStoreByName("store2").(types.KVStore)
	store2.Set([]byte(k2), []byte(fmt.Sprintf("%s:%d", v2, 3)))

	multi.Commit()

	flushedCinfo, err := getCommitInfo(multi.db, 3)
	require.Nil(t, err)
	require.NotEqual(t, initCid, flushedCinfo, "CID is different after flush to disk")


	store3 := multi.GetStoreByName("store3").(types.KVStore)
	store3.Set([]byte(k3), []byte(fmt.Sprintf("%s:%d", v3, 3)))

	multi.Commit()

	postFlushCinfo, err := getCommitInfo(multi.db, 4)
	require.NoError(t, err)
	require.Equal(t, int64(4), postFlushCinfo.Version, "Commit changed after in-memory commit")

	multi = newMultiStoreWithMounts(db, pruning)
	err = multi.LoadLatestVersion()
	require.Nil(t, err)

	reloadedCid := multi.LastCommitID()
	require.Equal(t, int64(4), reloadedCid.Version, "Reloaded CID is not the same as last flushed CID")


	store1 = multi.GetStoreByName("store1").(types.KVStore)
	val := store1.Get([]byte(k))
	require.Equal(t, []byte(fmt.Sprintf("%s:%d", v, 3)), val, "Reloaded value not the same as last flushed value")

	store2 = multi.GetStoreByName("store2").(types.KVStore)
	val2 := store2.Get([]byte(k2))
	require.Equal(t, []byte(fmt.Sprintf("%s:%d", v2, 3)), val2, "Reloaded value not the same as last flushed value")


	store3 = multi.GetStoreByName("store3").(types.KVStore)
	val3 := store3.Get([]byte(k3))
	require.Equal(t, []byte(fmt.Sprintf("%s:%d", v3, 3)), val3, "Reloaded value not the same as last flushed value")
}

func TestMultiStoreQuery(t *testing.T) {
	db := dbm.NewMemDB()
	multi := newMultiStoreWithMounts(db, types.PruneNothing)
	err := multi.LoadLatestVersion()
	require.Nil(t, err)

	k, v := []byte("wind"), []byte("blows")
	k2, v2 := []byte("water"), []byte("flows")


	cid := multi.Commit()


	garbage := multi.GetStoreByName("bad-name")
	require.Nil(t, garbage)


	store1 := multi.GetStoreByName("store1").(types.KVStore)
	store1.Set(k, v)


	store2 := multi.GetStoreByName("store2").(types.KVStore)
	store2.Set(k2, v2)


	cid = multi.Commit()
	ver := cid.Version


	multi = newMultiStoreWithMounts(db, types.PruneNothing)
	err = multi.LoadLatestVersion()
	require.Nil(t, err)


	query := abci.RequestQuery{Path: "/key", Data: k, Height: ver}
	qres := multi.Query(query)
	require.EqualValues(t, sdkerrors.ErrUnknownRequest.ABCICode(), qres.Code)
	require.EqualValues(t, sdkerrors.ErrUnknownRequest.Codespace(), qres.Codespace)

	query.Path = "h897fy32890rf63296r92"
	qres = multi.Query(query)
	require.EqualValues(t, sdkerrors.ErrUnknownRequest.ABCICode(), qres.Code)
	require.EqualValues(t, sdkerrors.ErrUnknownRequest.Codespace(), qres.Codespace)


	query.Path = "/garbage/key"
	qres = multi.Query(query)
	require.EqualValues(t, sdkerrors.ErrUnknownRequest.ABCICode(), qres.Code)
	require.EqualValues(t, sdkerrors.ErrUnknownRequest.Codespace(), qres.Codespace)


	query.Path = "/store1/key"
	qres = multi.Query(query)
	require.EqualValues(t, 0, qres.Code)
	require.Equal(t, v, qres.Value)


	query.Path = "/store2/key"
	query.Prove = true
	qres = multi.Query(query)
	require.EqualValues(t, 0, qres.Code)
	require.Nil(t, qres.Value)


	query.Data = k2
	qres = multi.Query(query)
	require.EqualValues(t, 0, qres.Code)
	require.Equal(t, v2, qres.Value)
}

func TestMultiStore_Pruning(t *testing.T) {
	testCases := []struct {
		name        string
		numVersions int64
		po          types.PruningOptions
		deleted     []int64
		saved       []int64
	}{
		{"prune nothing", 10, types.PruneNothing, nil, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"prune everything", 10, types.PruneEverything, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9}, []int64{10}},
		{"prune some; no batch", 10, types.NewPruningOptions(2, 3, 1), []int64{1, 2, 4, 5, 7}, []int64{3, 6, 8, 9, 10}},
		{"prune some; small batch", 10, types.NewPruningOptions(2, 3, 3), []int64{1, 2, 4, 5}, []int64{3, 6, 7, 8, 9, 10}},
		{"prune some; large batch", 10, types.NewPruningOptions(2, 3, 11), nil, []int64{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
	}

	for _, tc := range testCases {
		tc := tc

		t.Run(tc.name, func(t *testing.T) {
			db := dbm.NewMemDB()
			ms := newMultiStoreWithMounts(db, tc.po)
			require.NoError(t, ms.LoadLatestVersion())

			for i := int64(0); i < tc.numVersions; i++ {
				ms.Commit()
			}

			for _, v := range tc.saved {
				_, err := ms.CacheMultiStoreWithVersion(v)
				require.NoError(t, err, "expected error when loading height: %d", v)
			}

			for _, v := range tc.deleted {
				_, err := ms.CacheMultiStoreWithVersion(v)
				require.NoError(t, err, "expected error when loading height: %d", v)
			}
		})
	}
}

func TestMultiStore_PruningRestart(t *testing.T) {
	db := dbm.NewMemDB()
	ms := newMultiStoreWithMounts(db, types.NewPruningOptions(2, 3, 11))
	require.NoError(t, ms.LoadLatestVersion())



	for i := int64(0); i < 10; i++ {
		ms.Commit()
	}

	pruneHeights := []int64{1, 2, 4, 5, 7}


	ph, err := getPruningHeights(ms.db)
	require.NoError(t, err)
	require.Equal(t, pruneHeights, ph)


	ms = newMultiStoreWithMounts(db, types.NewPruningOptions(2, 3, 11))
	err = ms.LoadLatestVersion()
	require.NoError(t, err)
	require.Equal(t, pruneHeights, ms.pruneHeights)


	ms.Commit()
	require.Empty(t, ms.pruneHeights)

	for _, v := range pruneHeights {
		_, err := ms.CacheMultiStoreWithVersion(v)
		require.NoError(t, err, "expected error when loading height: %d", v)
	}
}

func TestSetInitialVersion(t *testing.T) {
	db := dbm.NewMemDB()
	multi := newMultiStoreWithMounts(db, types.PruneNothing)

	require.NoError(t, multi.LoadLatestVersion())

	multi.SetInitialVersion(5)
	require.Equal(t, int64(5), multi.initialVersion)

	multi.Commit()
	require.Equal(t, int64(5), multi.LastCommitID().Version)

	ckvs := multi.GetCommitKVStore(multi.keysByName["store1"])
	iavlStore, ok := ckvs.(*iavl.Store)
	require.True(t, ok)
	require.True(t, iavlStore.VersionExists(5))
}

func TestAddListenersAndListeningEnabled(t *testing.T) {
	db := dbm.NewMemDB()
	multi := newMultiStoreWithMounts(db, types.PruneNothing)
	testKey := types.NewKVStoreKey("listening_test_key")
	enabled := multi.ListeningEnabled(testKey)
	require.False(t, enabled)

	multi.AddListeners(testKey, []types.WriteListener{})
	enabled = multi.ListeningEnabled(testKey)
	require.False(t, enabled)

	mockListener := types.NewStoreKVPairWriteListener(nil, nil)
	multi.AddListeners(testKey, []types.WriteListener{mockListener})
	wrongTestKey := types.NewKVStoreKey("wrong_listening_test_key")
	enabled = multi.ListeningEnabled(wrongTestKey)
	require.False(t, enabled)

	enabled = multi.ListeningEnabled(testKey)
	require.True(t, enabled)
}

var (
	interfaceRegistry = codecTypes.NewInterfaceRegistry()
	testMarshaller    = codec.NewProtoCodec(interfaceRegistry)
	testKey1          = []byte{1, 2, 3, 4, 5}
	testValue1        = []byte{5, 4, 3, 2, 1}
	testKey2          = []byte{2, 3, 4, 5, 6}
	testValue2        = []byte{6, 5, 4, 3, 2}
)

func TestGetListenWrappedKVStore(t *testing.T) {
	buf := new(bytes.Buffer)
	var db dbm.DB = dbm.NewMemDB()
	ms := newMultiStoreWithMounts(db, types.PruneNothing)
	ms.LoadLatestVersion()
	mockListeners := []types.WriteListener{types.NewStoreKVPairWriteListener(buf, testMarshaller)}
	ms.AddListeners(testStoreKey1, mockListeners)
	ms.AddListeners(testStoreKey2, mockListeners)

	listenWrappedStore1 := ms.GetKVStore(testStoreKey1)
	require.IsType(t, &listenkv.Store{}, listenWrappedStore1)

	listenWrappedStore1.Set(testKey1, testValue1)
	expectedOutputKVPairSet1, err := testMarshaller.MarshalLengthPrefixed(&types.StoreKVPair{
		Key:      testKey1,
		Value:    testValue1,
		StoreKey: testStoreKey1.Name(),
		Delete:   false,
	})
	require.Nil(t, err)
	kvPairSet1Bytes := buf.Bytes()
	buf.Reset()
	require.Equal(t, expectedOutputKVPairSet1, kvPairSet1Bytes)

	listenWrappedStore1.Delete(testKey1)
	expectedOutputKVPairDelete1, err := testMarshaller.MarshalLengthPrefixed(&types.StoreKVPair{
		Key:      testKey1,
		Value:    nil,
		StoreKey: testStoreKey1.Name(),
		Delete:   true,
	})
	require.Nil(t, err)
	kvPairDelete1Bytes := buf.Bytes()
	buf.Reset()
	require.Equal(t, expectedOutputKVPairDelete1, kvPairDelete1Bytes)

	listenWrappedStore2 := ms.GetKVStore(testStoreKey2)
	require.IsType(t, &listenkv.Store{}, listenWrappedStore2)

	listenWrappedStore2.Set(testKey2, testValue2)
	expectedOutputKVPairSet2, err := testMarshaller.MarshalLengthPrefixed(&types.StoreKVPair{
		Key:      testKey2,
		Value:    testValue2,
		StoreKey: testStoreKey2.Name(),
		Delete:   false,
	})
	kvPairSet2Bytes := buf.Bytes()
	buf.Reset()
	require.Equal(t, expectedOutputKVPairSet2, kvPairSet2Bytes)

	listenWrappedStore2.Delete(testKey2)
	expectedOutputKVPairDelete2, err := testMarshaller.MarshalLengthPrefixed(&types.StoreKVPair{
		Key:      testKey2,
		Value:    nil,
		StoreKey: testStoreKey2.Name(),
		Delete:   true,
	})
	kvPairDelete2Bytes := buf.Bytes()
	buf.Reset()
	require.Equal(t, expectedOutputKVPairDelete2, kvPairDelete2Bytes)

	unwrappedStore := ms.GetKVStore(testStoreKey3)
	require.IsType(t, &iavl.Store{}, unwrappedStore)

	unwrappedStore.Set(testKey2, testValue2)
	kvPairSet3Bytes := buf.Bytes()
	buf.Reset()
	require.Equal(t, []byte{}, kvPairSet3Bytes)

	unwrappedStore.Delete(testKey2)
	kvPairDelete3Bytes := buf.Bytes()
	buf.Reset()
	require.Equal(t, []byte{}, kvPairDelete3Bytes)
}

func TestCacheWraps(t *testing.T) {
	db := dbm.NewMemDB()
	multi := newMultiStoreWithMounts(db, types.PruneNothing)

	cacheWrapper := multi.CacheWrap()
	require.IsType(t, cachemulti.Store{}, cacheWrapper)

	cacheWrappedWithTrace := multi.CacheWrapWithTrace(nil, nil)
	require.IsType(t, cachemulti.Store{}, cacheWrappedWithTrace)

	cacheWrappedWithListeners := multi.CacheWrapWithListeners(nil, nil)
	require.IsType(t, cachemulti.Store{}, cacheWrappedWithListeners)
}

func TestTraceConcurrency(t *testing.T) {
	db := dbm.NewMemDB()
	multi := newMultiStoreWithMounts(db, types.PruneNothing)
	err := multi.LoadLatestVersion()
	require.NoError(t, err)

	b := &bytes.Buffer{}
	key := multi.keysByName["store1"]
	tc := types.TraceContext(map[string]interface{}{"blockHeight": 64})

	multi.SetTracer(b)
	multi.SetTracingContext(tc)

	cms := multi.CacheMultiStore()
	store1 := cms.GetKVStore(key)
	cw := store1.CacheWrapWithTrace(b, tc)
	_ = cw
	require.NotNil(t, store1)

	stop := make(chan struct{})
	stopW := make(chan struct{})

	go func(stop chan struct{}) {
		for {
			select {
			case <-stop:
				return
			default:
				store1.Set([]byte{1}, []byte{1})
				cms.Write()
			}
		}
	}(stop)

	go func(stop chan struct{}) {
		for {
			select {
			case <-stop:
				return
			default:
				multi.SetTracingContext(tc)
			}
		}
	}(stopW)

	time.Sleep(3 * time.Second)
	stop <- struct{}{}
	stopW <- struct{}{}
}




var (
	testStoreKey1 = types.NewKVStoreKey("store1")
	testStoreKey2 = types.NewKVStoreKey("store2")
	testStoreKey3 = types.NewKVStoreKey("store3")
)

func newMultiStoreWithMounts(db dbm.DB, pruningOpts types.PruningOptions) *Store {
	store := NewStore(db)
	store.pruningOpts = pruningOpts

	store.MountStoreWithDB(testStoreKey1, types.StoreTypeIAVL, nil)
	store.MountStoreWithDB(testStoreKey2, types.StoreTypeIAVL, nil)
	store.MountStoreWithDB(testStoreKey3, types.StoreTypeIAVL, nil)

	return store
}

func newMultiStoreWithModifiedMounts(db dbm.DB, pruningOpts types.PruningOptions) (*Store, *types.StoreUpgrades) {
	store := NewStore(db)
	store.pruningOpts = pruningOpts

	store.MountStoreWithDB(types.NewKVStoreKey("store1"), types.StoreTypeIAVL, nil)
	store.MountStoreWithDB(types.NewKVStoreKey("restore2"), types.StoreTypeIAVL, nil)
	store.MountStoreWithDB(types.NewKVStoreKey("store3"), types.StoreTypeIAVL, nil)
	store.MountStoreWithDB(types.NewKVStoreKey("store4"), types.StoreTypeIAVL, nil)

	upgrades := &types.StoreUpgrades{
		Added: []string{"store4"},
		Renamed: []types.StoreRename{{
			OldKey: "store2",
			NewKey: "restore2",
		}},
		Deleted: []string{"store3"},
	}

	return store, upgrades
}

func checkStore(t *testing.T, store *Store, expect, got types.CommitID) {
	require.Equal(t, expect, got)
	require.Equal(t, expect, store.LastCommitID())
}

func checkContains(t testing.TB, info []types.StoreInfo, wanted []string) {
	t.Helper()

	for _, want := range wanted {
		checkHas(t, info, want)
	}
}

func checkHas(t testing.TB, info []types.StoreInfo, want string) {
	t.Helper()
	for _, i := range info {
		if i.Name == want {
			return
		}
	}
	t.Fatalf("storeInfo doesn't contain %s", want)
}

func getExpectedCommitID(store *Store, ver int64) types.CommitID {
	return types.CommitID{
		Version: ver,
		Hash:    hashStores(store.stores),
	}
}

func hashStores(stores map[types.StoreKey]types.CommitKVStore) []byte {
	m := make(map[string][]byte, len(stores))
	for key, store := range stores {
		name := key.Name()
		m[name] = types.StoreInfo{
			Name:     name,
			CommitId: store.LastCommitID(),
		}.GetHash()
	}
	return sdkmaps.HashFromMap(m)
}
