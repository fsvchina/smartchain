package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (

	ModuleName = "auth"


	StoreKey = "acc"


	FeeCollectorName = "fee_collector"


	QuerierRoute = ModuleName
)

var (

	AddressStoreKeyPrefix = []byte{0x01}


	GlobalAccountNumberKey = []byte("globalAccountNumber")
)


func AddressStoreKey(addr sdk.AccAddress) []byte {
	return append(AddressStoreKeyPrefix, addr.Bytes()...)
}
