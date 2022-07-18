package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
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
	return append(ValidatorSigningInfoKeyPrefix, address.MustLengthPrefix(v.Bytes())...)
}


func ValidatorSigningInfoAddress(key []byte) (v sdk.ConsAddress) {

	addr := key[2:]

	return sdk.ConsAddress(addr)
}


func ValidatorMissedBlockBitArrayPrefixKey(v sdk.ConsAddress) []byte {
	return append(ValidatorMissedBlockBitArrayKeyPrefix, address.MustLengthPrefix(v.Bytes())...)
}


func ValidatorMissedBlockBitArrayKey(v sdk.ConsAddress, i int64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, uint64(i))

	return append(ValidatorMissedBlockBitArrayPrefixKey(v), b...)
}


func AddrPubkeyRelationKey(addr []byte) []byte {
	return append(AddrPubkeyRelationKeyPrefix, address.MustLengthPrefix(addr)...)
}
