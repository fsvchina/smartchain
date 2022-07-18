package feegrant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (

	ModuleName = "feegrant"


	StoreKey = ModuleName


	RouterKey = ModuleName


	QuerierRoute = ModuleName
)

var (

	FeeAllowanceKeyPrefix = []byte{0x00}
)



func FeeAllowanceKey(granter sdk.AccAddress, grantee sdk.AccAddress) []byte {
	return append(FeeAllowancePrefixByGrantee(grantee), address.MustLengthPrefix(granter.Bytes())...)
}


func FeeAllowancePrefixByGrantee(grantee sdk.AccAddress) []byte {
	return append(FeeAllowanceKeyPrefix, address.MustLengthPrefix(grantee.Bytes())...)
}
