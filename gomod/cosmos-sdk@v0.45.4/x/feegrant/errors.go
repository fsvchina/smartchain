package feegrant

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


const (
	DefaultCodespace = ModuleName
)

var (

	ErrFeeLimitExceeded = sdkerrors.Register(DefaultCodespace, 2, "fee limit exceeded")

	ErrFeeLimitExpired = sdkerrors.Register(DefaultCodespace, 3, "fee allowance expired")

	ErrInvalidDuration = sdkerrors.Register(DefaultCodespace, 4, "invalid duration")

	ErrNoAllowance = sdkerrors.Register(DefaultCodespace, 5, "no allowance")

	ErrNoMessages = sdkerrors.Register(DefaultCodespace, 6, "allowed messages are empty")

	ErrMessageNotAllowed = sdkerrors.Register(DefaultCodespace, 7, "message not allowed")
)
