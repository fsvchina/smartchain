package transient

import (
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/store/dbadapter"
	"github.com/cosmos/cosmos-sdk/store/types"
)

var _ types.Committer = (*Store)(nil)
var _ types.KVStore = (*Store)(nil)


type Store struct {
	dbadapter.Store
}


func NewStore() *Store {
	return &Store{Store: dbadapter.Store{DB: dbm.NewMemDB()}}
}



func (ts *Store) Commit() (id types.CommitID) {
	ts.Store = dbadapter.Store{DB: dbm.NewMemDB()}
	return
}

func (ts *Store) SetPruning(_ types.PruningOptions) {}



func (ts *Store) GetPruning() types.PruningOptions { return types.PruningOptions{} }


func (ts *Store) LastCommitID() (id types.CommitID) {
	return
}


func (ts *Store) GetStoreType() types.StoreType {
	return types.StoreTypeTransient
}
