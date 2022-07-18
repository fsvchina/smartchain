package types

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (

	ModuleName = "capability"


	StoreKey = ModuleName


	MemStoreKey = "mem_capability"
)

var (


	KeyIndex = []byte("index")



	KeyPrefixIndexCapability = []byte("capability_index")


	KeyMemInitialized = []byte("mem_initialized")
)



func RevCapabilityKey(module, name string) []byte {
	return []byte(fmt.Sprintf("%s/rev/%s", module, name))
}



func FwdCapabilityKey(module string, cap *Capability) []byte {
	return []byte(fmt.Sprintf("%s/fwd/%p", module, cap))
}


func IndexToKey(index uint64) []byte {
	return sdk.Uint64ToBigEndian(index)
}



func IndexFromKey(key []byte) uint64 {
	return sdk.BigEndianToUint64(key)
}
