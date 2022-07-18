

package v040

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v040auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v040"
)

const (

	ModuleName = "slashing"


	StoreKey = ModuleName


	RouterKey = ModuleName


	QuerierRoute = ModuleName
)



//

//

//

var (
	ValidatorSigningInfoKeyPrefix         = []byte{0x01}
	ValidatorMissedBlockBitArrayKeyPrefix = []byte{0x02}
	AddrPubkeyRelationKeyPrefix           = []byte{0x03}
)


func ValidatorSigningInfoKey(v sdk.ConsAddress) []byte {
	return append(ValidatorSigningInfoKeyPrefix, v.Bytes()...)
}


func ValidatorSigningInfoAddress(key []byte) (v sdk.ConsAddress) {
	addr := key[1:]
	if len(addr) != v040auth.AddrLen {
		panic("unexpected key length")
	}
	return sdk.ConsAddress(addr)
}


func ValidatorMissedBlockBitArrayPrefixKey(v sdk.ConsAddress) []byte {
	return append(ValidatorMissedBlockBitArrayKeyPrefix, v.Bytes()...)
}


func ValidatorMissedBlockBitArrayKey(v sdk.ConsAddress, i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))
	return append(ValidatorMissedBlockBitArrayPrefixKey(v), b...)
}


func AddrPubkeyRelationKey(address []byte) []byte {
	return append(AddrPubkeyRelationKeyPrefix, address...)
}
