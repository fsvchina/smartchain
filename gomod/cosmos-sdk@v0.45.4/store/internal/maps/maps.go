package maps

import (
	"encoding/binary"

	"github.com/tendermint/tendermint/crypto/merkle"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmcrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/types/kv"
)



type merkleMap struct {
	kvs    kv.Pairs
	sorted bool
}

func newMerkleMap() *merkleMap {
	return &merkleMap{
		kvs:    kv.Pairs{},
		sorted: false,
	}
}




func (sm *merkleMap) set(key string, value []byte) {
	byteKey := []byte(key)
	assertValidKey(byteKey)

	sm.sorted = false



	vhash := tmhash.Sum(value)

	sm.kvs.Pairs = append(sm.kvs.Pairs, kv.Pair{
		Key:   byteKey,
		Value: vhash,
	})
}


func (sm *merkleMap) hash() []byte {
	sm.sort()
	return hashKVPairs(sm.kvs)
}

func (sm *merkleMap) sort() {
	if sm.sorted {
		return
	}

	sm.kvs.Sort()
	sm.sorted = true
}



func hashKVPairs(kvs kv.Pairs) []byte {
	kvsH := make([][]byte, len(kvs.Pairs))
	for i, kvp := range kvs.Pairs {
		kvsH[i] = KVPair(kvp).Bytes()
	}

	return merkle.HashFromByteSlices(kvsH)
}






type simpleMap struct {
	Kvs    kv.Pairs
	sorted bool
}

func newSimpleMap() *simpleMap {
	return &simpleMap{
		Kvs:    kv.Pairs{},
		sorted: false,
	}
}



func (sm *simpleMap) Set(key string, value []byte) {
	byteKey := []byte(key)
	assertValidKey(byteKey)
	sm.sorted = false




	vhash := tmhash.Sum(value)

	sm.Kvs.Pairs = append(sm.Kvs.Pairs, kv.Pair{
		Key:   byteKey,
		Value: vhash,
	})
}



func (sm *simpleMap) Hash() []byte {
	sm.Sort()
	return hashKVPairs(sm.Kvs)
}

func (sm *simpleMap) Sort() {
	if sm.sorted {
		return
	}
	sm.Kvs.Sort()
	sm.sorted = true
}



func (sm *simpleMap) KVPairs() kv.Pairs {
	sm.Sort()
	kvs := kv.Pairs{
		Pairs: make([]kv.Pair, len(sm.Kvs.Pairs)),
	}

	copy(kvs.Pairs, sm.Kvs.Pairs)
	return kvs
}






type KVPair kv.Pair



func NewKVPair(key, value []byte) KVPair {
	return KVPair(kv.Pair{
		Key:   key,
		Value: value,
	})
}



func (kv KVPair) Bytes() []byte {






	buf := make([]byte, 8+len(kv.Key)+8+len(kv.Value))


	nlk := binary.PutUvarint(buf, uint64(len(kv.Key)))
	nk := copy(buf[nlk:], kv.Key)


	nlv := binary.PutUvarint(buf[nlk+nk:], uint64(len(kv.Value)))
	nv := copy(buf[nlk+nk+nlv:], kv.Value)

	return buf[:nlk+nk+nlv+nv]
}



func HashFromMap(m map[string][]byte) []byte {
	mm := newMerkleMap()
	for k, v := range m {
		mm.set(k, v)
	}

	return mm.hash()
}




func ProofsFromMap(m map[string][]byte) ([]byte, map[string]*tmcrypto.Proof, []string) {
	sm := newSimpleMap()
	for k, v := range m {
		sm.Set(k, v)
	}

	sm.Sort()
	kvs := sm.Kvs
	kvsBytes := make([][]byte, len(kvs.Pairs))
	for i, kvp := range kvs.Pairs {
		kvsBytes[i] = KVPair(kvp).Bytes()
	}

	rootHash, proofList := merkle.ProofsFromByteSlices(kvsBytes)
	proofs := make(map[string]*tmcrypto.Proof)
	keys := make([]string, len(proofList))

	for i, kvp := range kvs.Pairs {
		proofs[string(kvp.Key)] = proofList[i].ToProto()
		keys[i] = string(kvp.Key)
	}

	return rootHash, proofs, keys
}

func assertValidKey(key []byte) {
	if len(key) == 0 {
		panic("key is nil")
	}
}
