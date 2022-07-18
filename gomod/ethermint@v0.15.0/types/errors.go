package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (

	RootCodespace = "ethermint"
)



var (

	ErrInvalidValue = sdkerrors.Register(RootCodespace, 2, "invalid value")


	ErrInvalidChainID = sdkerrors.Register(RootCodespace, 3, "invalid chain ID")


	ErrMarshalBigInt = sdkerrors.Register(RootCodespace, 5, "cannot marshal big.Int to string")


	ErrUnmarshalBigInt = sdkerrors.Register(RootCodespace, 6, "cannot unmarshal big.Int from string")
)
