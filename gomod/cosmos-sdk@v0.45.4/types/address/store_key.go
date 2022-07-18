package address

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


const MaxAddrLen = 255



func LengthPrefix(bz []byte) ([]byte, error) {
	bzLen := len(bz)
	if bzLen == 0 {
		return bz, nil
	}

	if bzLen > MaxAddrLen {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "address length should be max %d bytes, got %d", MaxAddrLen, bzLen)
	}

	return append([]byte{byte(bzLen)}, bz...), nil
}


func MustLengthPrefix(bz []byte) []byte {
	res, err := LengthPrefix(bz)
	if err != nil {
		panic(err)
	}

	return res
}
