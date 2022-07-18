package acc

import (
	"errors"
	"fs.video/smartchain/cmd/config"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"strings"
)


func AccAddressFromBech32(address string) (addr sdk.AccAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return sdk.AccAddress{}, errors.New("empty address string is not allowed")
	}

	bz, err := sdk.GetFromBech32(address, config.Bech32Prefix)
	if err != nil {
		return nil, err
	}

	err = verifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return sdk.AccAddress(bz), nil
}

func verifyAddressFormat(bz []byte) error {
	if len(bz) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "addresses cannot be empty")
	}

	if len(bz) > address.MaxAddrLen {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "address max length is %d, got %d", address.MaxAddrLen, len(bz))
	}

	return nil
}

func AccAddressToString(addr sdk.AccAddress) string {
	if addr.Empty() {
		return ""
	}
	bech32Addr, err := bech32.ConvertAndEncode(config.Bech32Prefix, addr)
	if err != nil {
		panic(err)
	}
	return bech32Addr
}