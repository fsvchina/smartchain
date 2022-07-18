package proofs

import (
	"sort"

	"github.com/tendermint/tendermint/libs/rand"
	tmcrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"

	sdkmaps "github.com/cosmos/cosmos-sdk/store/internal/maps"
)


type SimpleResult struct {
	Key      []byte
	Value    []byte
	Proof    *tmcrypto.Proof
	RootHash []byte
}


//

func GenerateRangeProof(size int, loc Where) *SimpleResult {
	data := BuildMap(size)
	root, proofs, allkeys := sdkmaps.ProofsFromMap(data)

	key := GetKey(allkeys, loc)
	proof := proofs[key]

	res := &SimpleResult{
		Key:      []byte(key),
		Value:    toValue(key),
		Proof:    proof,
		RootHash: root,
	}
	return res
}


type Where int

const (
	Left Where = iota
	Right
	Middle
)

func SortedKeys(data map[string][]byte) []string {
	keys := make([]string, len(data))
	i := 0
	for k := range data {
		keys[i] = k
		i++
	}
	sort.Strings(keys)
	return keys
}

func CalcRoot(data map[string][]byte) []byte {
	root, _, _ := sdkmaps.ProofsFromMap(data)
	return root
}


func GetKey(allkeys []string, loc Where) string {
	if loc == Left {
		return allkeys[0]
	}
	if loc == Right {
		return allkeys[len(allkeys)-1]
	}
	
	idx := rand.Int()%(len(allkeys)-2) + 1
	return allkeys[idx]
}


func GetNonKey(allkeys []string, loc Where) string {
	if loc == Left {
		return string([]byte{1, 1, 1, 1})
	}
	if loc == Right {
		return string([]byte{0xff, 0xff, 0xff, 0xff})
	}
	
	key := GetKey(allkeys, loc)
	key = key[:len(key)-2] + string([]byte{255, 255})
	return key
}

func toValue(key string) []byte {
	return []byte("value_for_" + key)
}



func BuildMap(size int) map[string][]byte {
	data := make(map[string][]byte)
	
	for i := 0; i < size; i++ {
		key := rand.Str(20)
		data[key] = toValue(key)
	}
	return data
}
