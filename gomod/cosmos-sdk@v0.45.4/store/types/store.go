package types

import (
	"fmt"
	"io"

	abci "github.com/tendermint/tendermint/abci/types"
	tmstrings "github.com/tendermint/tendermint/libs/strings"
	dbm "github.com/tendermint/tm-db"

	snapshottypes "github.com/cosmos/cosmos-sdk/snapshots/types"
	"github.com/cosmos/cosmos-sdk/types/kv"
)

type Store interface {
	GetStoreType() StoreType
	CacheWrapper
}


type Committer interface {
	Commit() CommitID
	LastCommitID() CommitID

	SetPruning(PruningOptions)
	GetPruning() PruningOptions
}


type CommitStore interface {
	Committer
	Store
}



//

type Queryable interface {
	Query(abci.RequestQuery) abci.ResponseQuery
}





type StoreUpgrades struct {
	Added   []string      `json:"added"`
	Renamed []StoreRename `json:"renamed"`
	Deleted []string      `json:"deleted"`
}



type UpgradeInfo struct {
	Name   string `json:"name"`
	Height int64  `json:"height"`
}




type StoreRename struct {
	OldKey string `json:"old_key"`
	NewKey string `json:"new_key"`
}


func (s *StoreUpgrades) IsAdded(key string) bool {
	if s == nil {
		return false
	}
	return tmstrings.StringInSlice(key, s.Added)
}


func (s *StoreUpgrades) IsDeleted(key string) bool {
	if s == nil {
		return false
	}
	for _, d := range s.Deleted {
		if d == key {
			return true
		}
	}
	return false
}



func (s *StoreUpgrades) RenamedFrom(key string) string {
	if s == nil {
		return ""
	}
	for _, re := range s.Renamed {
		if re.NewKey == key {
			return re.OldKey
		}
	}
	return ""

}

type MultiStore interface {
	Store




	CacheMultiStore() CacheMultiStore



	CacheMultiStoreWithVersion(version int64) (CacheMultiStore, error)



	GetStore(StoreKey) Store
	GetKVStore(StoreKey) KVStore


	TracingEnabled() bool




	SetTracer(w io.Writer) MultiStore




	SetTracingContext(TraceContext) MultiStore


	ListeningEnabled(key StoreKey) bool



	AddListeners(key StoreKey, listeners []WriteListener)
}


type CacheMultiStore interface {
	MultiStore
	Write()
}


type CommitMultiStore interface {
	Committer
	MultiStore
	snapshottypes.Snapshotter



	MountStoreWithDB(key StoreKey, typ StoreType, db dbm.DB)


	GetCommitStore(key StoreKey) CommitStore


	GetCommitKVStore(key StoreKey) CommitKVStore



	LoadLatestVersion() error




	LoadLatestVersionAndUpgrade(upgrades *StoreUpgrades) error




	LoadVersionAndUpgrade(ver int64, upgrades *StoreUpgrades) error





	LoadVersion(ver int64) error



	SetInterBlockCache(MultiStorePersistentCache)



	SetInitialVersion(version int64) error


	SetIAVLCacheSize(size int)
}





type KVStore interface {
	Store


	Get(key []byte) []byte


	Has(key []byte) bool


	Set(key, value []byte)


	Delete(key []byte)







	Iterator(start, end []byte) Iterator






	ReverseIterator(start, end []byte) Iterator
}


type Iterator = dbm.Iterator




type CacheKVStore interface {
	KVStore


	Write()
}


type CommitKVStore interface {
	Committer
	KVStore
}








type CacheWrap interface {

	Write()


	CacheWrap() CacheWrap


	CacheWrapWithTrace(w io.Writer, tc TraceContext) CacheWrap


	CacheWrapWithListeners(storeKey StoreKey, listeners []WriteListener) CacheWrap
}

type CacheWrapper interface {

	CacheWrap() CacheWrap


	CacheWrapWithTrace(w io.Writer, tc TraceContext) CacheWrap


	CacheWrapWithListeners(storeKey StoreKey, listeners []WriteListener) CacheWrap
}

func (cid CommitID) IsZero() bool {
	return cid.Version == 0 && len(cid.Hash) == 0
}

func (cid CommitID) String() string {
	return fmt.Sprintf("CommitID{%v:%X}", cid.Hash, cid.Version)
}





type StoreType int

const (
	StoreTypeMulti StoreType = iota
	StoreTypeDB
	StoreTypeIAVL
	StoreTypeTransient
	StoreTypeMemory
)

func (st StoreType) String() string {
	switch st {
	case StoreTypeMulti:
		return "StoreTypeMulti"

	case StoreTypeDB:
		return "StoreTypeDB"

	case StoreTypeIAVL:
		return "StoreTypeIAVL"

	case StoreTypeTransient:
		return "StoreTypeTransient"

	case StoreTypeMemory:
		return "StoreTypeMemory"
	}

	return "unknown store type"
}





type StoreKey interface {
	Name() string
	String() string
}



type CapabilityKey StoreKey



type KVStoreKey struct {
	name string
}



func NewKVStoreKey(name string) *KVStoreKey {
	if name == "" {
		panic("empty key name not allowed")
	}
	return &KVStoreKey{
		name: name,
	}
}

func (key *KVStoreKey) Name() string {
	return key.name
}

func (key *KVStoreKey) String() string {
	return fmt.Sprintf("KVStoreKey{%p, %s}", key, key.name)
}


type TransientStoreKey struct {
	name string
}



func NewTransientStoreKey(name string) *TransientStoreKey {
	return &TransientStoreKey{
		name: name,
	}
}


func (key *TransientStoreKey) Name() string {
	return key.name
}


func (key *TransientStoreKey) String() string {
	return fmt.Sprintf("TransientStoreKey{%p, %s}", key, key.name)
}


type MemoryStoreKey struct {
	name string
}

func NewMemoryStoreKey(name string) *MemoryStoreKey {
	return &MemoryStoreKey{name: name}
}


func (key *MemoryStoreKey) Name() string {
	return key.name
}


func (key *MemoryStoreKey) String() string {
	return fmt.Sprintf("MemoryStoreKey{%p, %s}", key, key.name)
}




type KVPair kv.Pair





type TraceContext map[string]interface{}



type MultiStorePersistentCache interface {


	GetStoreCache(key StoreKey, store CommitKVStore) CommitKVStore


	Unwrap(key StoreKey) CommitKVStore


	Reset()
}



type StoreWithInitialVersion interface {


	SetInitialVersion(version int64)
}
