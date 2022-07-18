package iavl

import (
	"errors"
	"fmt"
	"io"
	"time"

	ics23 "github.com/confio/ics23/go"
	"github.com/cosmos/iavl"
	abci "github.com/tendermint/tendermint/abci/types"
	tmcrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/listenkv"
	"github.com/cosmos/cosmos-sdk/store/tracekv"
	"github.com/cosmos/cosmos-sdk/store/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

const (
	DefaultIAVLCacheSize = 500000
)

var (
	_ types.KVStore                 = (*Store)(nil)
	_ types.CommitStore             = (*Store)(nil)
	_ types.CommitKVStore           = (*Store)(nil)
	_ types.Queryable               = (*Store)(nil)
	_ types.StoreWithInitialVersion = (*Store)(nil)
)


type Store struct {
	tree Tree
}




func LoadStore(db dbm.DB, id types.CommitID, lazyLoading bool, cacheSize int) (types.CommitKVStore, error) {
	return LoadStoreWithInitialVersion(db, id, lazyLoading, 0, cacheSize)
}





func LoadStoreWithInitialVersion(db dbm.DB, id types.CommitID, lazyLoading bool, initialVersion uint64, cacheSize int) (types.CommitKVStore, error) {
	tree, err := iavl.NewMutableTreeWithOpts(db, cacheSize, &iavl.Options{InitialVersion: initialVersion})
	if err != nil {
		return nil, err
	}

	if lazyLoading {
		_, err = tree.LazyLoadVersion(id.Version)
	} else {
		_, err = tree.LoadVersion(id.Version)
	}

	if err != nil {
		return nil, err
	}

	return &Store{
		tree: tree,
	}, nil
}



//



func UnsafeNewStore(tree *iavl.MutableTree) *Store {
	return &Store{
		tree: tree,
	}
}






func (st *Store) GetImmutable(version int64) (*Store, error) {
	if !st.VersionExists(version) {
		return &Store{tree: &immutableTree{&iavl.ImmutableTree{}}}, nil
	}

	iTree, err := st.tree.GetImmutable(version)
	if err != nil {
		return nil, err
	}

	return &Store{
		tree: &immutableTree{iTree},
	}, nil
}



func (st *Store) Commit() types.CommitID {
	defer telemetry.MeasureSince(time.Now(), "store", "iavl", "commit")

	hash, version, err := st.tree.SaveVersion()
	if err != nil {
		panic(err)
	}

	return types.CommitID{
		Version: version,
		Hash:    hash,
	}
}


func (st *Store) LastCommitID() types.CommitID {
	return types.CommitID{
		Version: st.tree.Version(),
		Hash:    st.tree.Hash(),
	}
}



func (st *Store) SetPruning(_ types.PruningOptions) {
	panic("cannot set pruning options on an initialized IAVL store")
}



func (st *Store) GetPruning() types.PruningOptions {
	panic("cannot get pruning options on an initialized IAVL store")
}


func (st *Store) VersionExists(version int64) bool {
	return st.tree.VersionExists(version)
}


func (st *Store) GetAllVersions() []int {
	return st.tree.(*iavl.MutableTree).AvailableVersions()
}


func (st *Store) GetStoreType() types.StoreType {
	return types.StoreTypeIAVL
}


func (st *Store) CacheWrap() types.CacheWrap {
	return cachekv.NewStore(st)
}


func (st *Store) CacheWrapWithTrace(w io.Writer, tc types.TraceContext) types.CacheWrap {
	return cachekv.NewStore(tracekv.NewStore(st, w, tc))
}


func (st *Store) CacheWrapWithListeners(storeKey types.StoreKey, listeners []types.WriteListener) types.CacheWrap {
	return cachekv.NewStore(listenkv.NewStore(st, storeKey, listeners))
}


func (st *Store) Set(key, value []byte) {
	types.AssertValidKey(key)
	types.AssertValidValue(value)
	st.tree.Set(key, value)
}


func (st *Store) Get(key []byte) []byte {
	defer telemetry.MeasureSince(time.Now(), "store", "iavl", "get")
	_, value := st.tree.Get(key)
	return value
}


func (st *Store) Has(key []byte) (exists bool) {
	defer telemetry.MeasureSince(time.Now(), "store", "iavl", "has")
	return st.tree.Has(key)
}


func (st *Store) Delete(key []byte) {
	defer telemetry.MeasureSince(time.Now(), "store", "iavl", "delete")
	st.tree.Remove(key)
}




func (st *Store) DeleteVersions(versions ...int64) error {
	return st.tree.DeleteVersions(versions...)
}


func (st *Store) Iterator(start, end []byte) types.Iterator {
	var iTree *iavl.ImmutableTree

	switch tree := st.tree.(type) {
	case *immutableTree:
		iTree = tree.ImmutableTree
	case *iavl.MutableTree:
		iTree = tree.ImmutableTree
	}

	return newIAVLIterator(iTree, start, end, true)
}


func (st *Store) ReverseIterator(start, end []byte) types.Iterator {
	var iTree *iavl.ImmutableTree

	switch tree := st.tree.(type) {
	case *immutableTree:
		iTree = tree.ImmutableTree
	case *iavl.MutableTree:
		iTree = tree.ImmutableTree
	}

	return newIAVLIterator(iTree, start, end, false)
}



func (st *Store) SetInitialVersion(version int64) {
	st.tree.SetInitialVersion(uint64(version))
}


func (st *Store) Export(version int64) (*iavl.Exporter, error) {
	istore, err := st.GetImmutable(version)
	if err != nil {
		return nil, fmt.Errorf("iavl export failed for version %v: %w", version, err)
	}
	tree, ok := istore.tree.(*immutableTree)
	if !ok || tree == nil {
		return nil, fmt.Errorf("iavl export failed: unable to fetch tree for version %v", version)
	}
	return tree.Export(), nil
}


func (st *Store) Import(version int64) (*iavl.Importer, error) {
	tree, ok := st.tree.(*iavl.MutableTree)
	if !ok {
		return nil, errors.New("iavl import failed: unable to find mutable tree")
	}
	return tree.Import(version)
}


func getHeight(tree Tree, req abci.RequestQuery) int64 {
	height := req.Height
	if height == 0 {
		latest := tree.Version()
		if tree.VersionExists(latest - 1) {
			height = latest - 1
		} else {
			height = latest
		}
	}
	return height
}


//





func (st *Store) Query(req abci.RequestQuery) (res abci.ResponseQuery) {
	defer telemetry.MeasureSince(time.Now(), "store", "iavl", "query")

	if len(req.Data) == 0 {
		return sdkerrors.QueryResult(sdkerrors.Wrap(sdkerrors.ErrTxDecode, "query cannot be zero length"))
	}

	tree := st.tree



	res.Height = getHeight(tree, req)

	switch req.Path {
	case "/key":
		key := req.Data

		res.Key = key
		if !st.VersionExists(res.Height) {
			res.Log = iavl.ErrVersionDoesNotExist.Error()
			break
		}

		_, res.Value = tree.GetVersioned(key, res.Height)
		if !req.Prove {
			break
		}



		iTree, err := tree.GetImmutable(res.Height)
		if err != nil {

			panic(fmt.Sprintf("version exists in store but could not retrieve corresponding versioned tree in store, %s", err.Error()))
		}
		mtree := &iavl.MutableTree{
			ImmutableTree: iTree,
		}


		res.ProofOps = getProofFromTree(mtree, req.Data, res.Value != nil)

	case "/subspace":
		pairs := kv.Pairs{
			Pairs: make([]kv.Pair, 0),
		}

		subspace := req.Data
		res.Key = subspace

		iterator := types.KVStorePrefixIterator(st, subspace)
		for ; iterator.Valid(); iterator.Next() {
			pairs.Pairs = append(pairs.Pairs, kv.Pair{Key: iterator.Key(), Value: iterator.Value()})
		}
		iterator.Close()

		bz, err := pairs.Marshal()
		if err != nil {
			panic(fmt.Errorf("failed to marshal KV pairs: %w", err))
		}

		res.Value = bz

	default:
		return sdkerrors.QueryResult(sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unexpected query path: %v", req.Path))
	}

	return res
}




func getProofFromTree(tree *iavl.MutableTree, key []byte, exists bool) *tmcrypto.ProofOps {
	var (
		commitmentProof *ics23.CommitmentProof
		err             error
	)

	if exists {

		commitmentProof, err = tree.GetMembershipProof(key)
		if err != nil {

			panic(fmt.Sprintf("unexpected value for empty proof: %s", err.Error()))
		}
	} else {

		commitmentProof, err = tree.GetNonMembershipProof(key)
		if err != nil {

			panic(fmt.Sprintf("unexpected error for nonexistence proof: %s", err.Error()))
		}
	}

	op := types.NewIavlCommitmentOp(key, commitmentProof)
	return &tmcrypto.ProofOps{Ops: []tmcrypto.ProofOp{op.ProofOp()}}
}




type iavlIterator struct {
	*iavl.Iterator
}

var _ types.Iterator = (*iavlIterator)(nil)




func newIAVLIterator(tree *iavl.ImmutableTree, start, end []byte, ascending bool) *iavlIterator {
	iter := &iavlIterator{
		Iterator: tree.Iterator(start, end, ascending),
	}
	return iter
}
