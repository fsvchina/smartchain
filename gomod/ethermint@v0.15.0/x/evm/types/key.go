package types

import (
	"github.com/ethereum/go-ethereum/common"
)

const (

	ModuleName = "evm"




	StoreKey = ModuleName



	TransientKey = "transient_" + ModuleName


	RouterKey = ModuleName
)


const (
	prefixCode = iota + 1
	prefixStorage
)


const (
	prefixTransientBloom = iota + 1
	prefixTransientTxIndex
	prefixTransientLogSize
	prefixTransientGasUsed
)


var (
	KeyPrefixCode    = []byte{prefixCode}
	KeyPrefixStorage = []byte{prefixStorage}
)


var (
	KeyPrefixTransientBloom   = []byte{prefixTransientBloom}
	KeyPrefixTransientTxIndex = []byte{prefixTransientTxIndex}
	KeyPrefixTransientLogSize = []byte{prefixTransientLogSize}
	KeyPrefixTransientGasUsed = []byte{prefixTransientGasUsed}
)


func AddressStoragePrefix(address common.Address) []byte {
	return append(KeyPrefixStorage, address.Bytes()...)
}


func StateKey(address common.Address, key []byte) []byte {
	return append(AddressStoragePrefix(address), key...)
}
