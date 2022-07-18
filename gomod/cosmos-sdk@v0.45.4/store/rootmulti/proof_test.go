package rootmulti

import (
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/iavl"
	"github.com/cosmos/cosmos-sdk/store/types"
)

func TestVerifyIAVLStoreQueryProof(t *testing.T) {

	db := dbm.NewMemDB()
	iStore, err := iavl.LoadStore(db, types.CommitID{}, false, iavl.DefaultIAVLCacheSize)
	store := iStore.(*iavl.Store)
	require.Nil(t, err)
	store.Set([]byte("MYKEY"), []byte("MYVALUE"))
	cid := store.Commit()


	res := store.Query(abci.RequestQuery{
		Path:  "/key",
		Data:  []byte("MYKEY"),
		Prove: true,
	})
	require.NotNil(t, res.ProofOps)


	prt := DefaultProofRuntime()
	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/MYKEY", []byte("MYVALUE"))
	require.Nil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/MYKEY_NOT", []byte("MYVALUE"))
	require.NotNil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/MYKEY/MYKEY", []byte("MYVALUE"))
	require.NotNil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "MYKEY", []byte("MYVALUE"))
	require.NotNil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/MYKEY", []byte("MYVALUE_NOT"))
	require.NotNil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/MYKEY", []byte(nil))
	require.NotNil(t, err)
}

func TestVerifyMultiStoreQueryProof(t *testing.T) {

	db := dbm.NewMemDB()
	store := NewStore(db)
	iavlStoreKey := types.NewKVStoreKey("iavlStoreKey")

	store.MountStoreWithDB(iavlStoreKey, types.StoreTypeIAVL, nil)
	require.NoError(t, store.LoadVersion(0))

	iavlStore := store.GetCommitStore(iavlStoreKey).(*iavl.Store)
	iavlStore.Set([]byte("MYKEY"), []byte("MYVALUE"))
	cid := store.Commit()


	res := store.Query(abci.RequestQuery{
		Path:  "/iavlStoreKey/key",
		Data:  []byte("MYKEY"),
		Prove: true,
	})
	require.NotNil(t, res.ProofOps)


	prt := DefaultProofRuntime()
	err := prt.VerifyValue(res.ProofOps, cid.Hash, "/iavlStoreKey/MYKEY", []byte("MYVALUE"))
	require.Nil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/iavlStoreKey/MYKEY", []byte("MYVALUE"))
	require.Nil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/iavlStoreKey/MYKEY_NOT", []byte("MYVALUE"))
	require.NotNil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/iavlStoreKey/MYKEY/MYKEY", []byte("MYVALUE"))
	require.NotNil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "iavlStoreKey/MYKEY", []byte("MYVALUE"))
	require.NotNil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/MYKEY", []byte("MYVALUE"))
	require.NotNil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/iavlStoreKey/MYKEY", []byte("MYVALUE_NOT"))
	require.NotNil(t, err)


	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/iavlStoreKey/MYKEY", []byte(nil))
	require.NotNil(t, err)
}

func TestVerifyMultiStoreQueryProofAbsence(t *testing.T) {

	db := dbm.NewMemDB()
	store := NewStore(db)
	iavlStoreKey := types.NewKVStoreKey("iavlStoreKey")

	store.MountStoreWithDB(iavlStoreKey, types.StoreTypeIAVL, nil)
	err := store.LoadVersion(0)
	require.NoError(t, err)

	iavlStore := store.GetCommitStore(iavlStoreKey).(*iavl.Store)
	iavlStore.Set([]byte("MYKEY"), []byte("MYVALUE"))
	cid := store.Commit()


	res := store.Query(abci.RequestQuery{
		Path:  "/iavlStoreKey/key",
		Data:  []byte("MYABSENTKEY"),
		Prove: true,
	})
	require.NotNil(t, res.ProofOps)


	prt := DefaultProofRuntime()
	err = prt.VerifyAbsence(res.ProofOps, cid.Hash, "/iavlStoreKey/MYABSENTKEY")
	require.Nil(t, err)


	prt = DefaultProofRuntime()
	err = prt.VerifyAbsence(res.ProofOps, cid.Hash, "/MYABSENTKEY")
	require.NotNil(t, err)


	prt = DefaultProofRuntime()
	err = prt.VerifyValue(res.ProofOps, cid.Hash, "/iavlStoreKey/MYABSENTKEY", []byte(""))
	require.NotNil(t, err)
}
