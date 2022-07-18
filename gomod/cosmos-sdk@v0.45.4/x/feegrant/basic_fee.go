package feegrant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ FeeAllowanceI = (*BasicAllowance)(nil)




//



//


func (a *BasicAllowance) Accept(ctx sdk.Context, fee sdk.Coins, _ []sdk.Msg) (bool, error) {
	if a.Expiration != nil && a.Expiration.Before(ctx.BlockTime()) {
		return true, sdkerrors.Wrap(ErrFeeLimitExpired, "basic allowance")
	}

	if a.SpendLimit != nil {
		left, invalid := a.SpendLimit.SafeSub(fee)
		if invalid {
			return false, sdkerrors.Wrap(ErrFeeLimitExceeded, "basic allowance")
		}

		a.SpendLimit = left
		return left.IsZero(), nil
	}

	return false, nil
}


func (a BasicAllowance) ValidateBasic() error {
	if a.SpendLimit != nil {
		if !a.SpendLimit.IsValid() {
			return sdkerrors.Wrapf(sdkerrors.ErrInvalidCoins, "send amount is invalid: %s", a.SpendLimit)
		}
		if !a.SpendLimit.IsAllPositive() {
			return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "spend limit must be positive")
		}
	}

	if a.Expiration != nil && a.Expiration.Unix() < 0 {
		return sdkerrors.Wrap(ErrInvalidDuration, "expiration time cannot be negative")
	}

	return nil
}
