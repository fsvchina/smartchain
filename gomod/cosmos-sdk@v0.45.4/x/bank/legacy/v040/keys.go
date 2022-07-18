

package v040

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v040auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v040"
)

const (

	ModuleName = "bank"


	StoreKey = ModuleName


	RouterKey = ModuleName


	QuerierRoute = ModuleName
)


var (
	BalancesPrefix      = []byte("balances")
	SupplyKey           = []byte{0x00}
	DenomMetadataPrefix = []byte{0x1}
)


func DenomMetadataKey(denom string) []byte {
	d := []byte(denom)
	return append(DenomMetadataPrefix, d...)
}




func AddressFromBalancesStore(key []byte) sdk.AccAddress {
	addr := key[:v040auth.AddrLen]
	if len(addr) != v040auth.AddrLen {
		panic(fmt.Sprintf("unexpected account address key length; got: %d, expected: %d", len(addr), v040auth.AddrLen))
	}

	return sdk.AccAddress(addr)
}
