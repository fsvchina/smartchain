package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


var (
	ErrNoSender         = sdkerrors.Register(ModuleName, 2, "sender address is empty")
	ErrUnknownInvariant = sdkerrors.Register(ModuleName, 3, "unknown invariant")
)
