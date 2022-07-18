package mem

import (
	"io"

	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/cachekv"
	"github.com/cosmos/cosmos-sdk/store/dbadapter"
	"github.com/cosmos/cosmos-sdk/store/listenkv"
	"github.com/cosmos/cosmos-sdk/store/tracekv"
	"github.com/cosmos/cosmos-sdk/store/types"
)

var (
	_ types.KVStore   = (*Store)(nil)
	_ types.Committer = (*Store)(nil)
)



type Store struct {
	dbadapter.Store
}

func NewStore() *Store {
	return NewStoreWithDB(dbm.NewMemDB())
}

func NewStoreWithDB(db *dbm.MemDB) *Store {
	return &Store{Store: dbadapter.Store{DB: db}}
}


func (s Store) GetStoreType() types.StoreType {
	return types.StoreTypeMemory
}


func (s Store) CacheWrap() types.CacheWrap {
	return cachekv.NewStore(s)
}


func (s Store) CacheWrapWithTrace(w io.Writer, tc types.TraceContext) types.CacheWrap {
	return cachekv.NewStore(tracekv.NewStore(s, w, tc))
}


func (s Store) CacheWrapWithListeners(storeKey types.StoreKey, listeners []types.WriteListener) types.CacheWrap {
	return cachekv.NewStore(listenkv.NewStore(s, storeKey, listeners))
}


func (s *Store) Commit() (id types.CommitID) { return }

func (s *Store) SetPruning(pruning types.PruningOptions) {}



func (s *Store) GetPruning() types.PruningOptions { return types.PruningOptions{} }

func (s Store) LastCommitID() (id types.CommitID) { return }
