package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
)



func (k Keeper) Unjail(ctx sdk.Context, validatorAddr sdk.ValAddress) error {
	validator := k.sk.Validator(ctx, validatorAddr)
	if validator == nil {
		return types.ErrNoValidatorForAddress
	}


	selfDel := k.sk.Delegation(ctx, sdk.AccAddress(validatorAddr), validatorAddr)
	if selfDel == nil {
		return types.ErrMissingSelfDelegation
	}

	tokens := validator.TokensFromShares(selfDel.GetShares()).TruncateInt()
	minSelfBond := validator.GetMinSelfDelegation()
	if tokens.LT(minSelfBond) {
		return sdkerrors.Wrapf(
			types.ErrSelfDelegationTooLowToUnjail, "%s less than %s", tokens, minSelfBond,
		)
	}


	if !validator.IsJailed() {
		return types.ErrValidatorNotJailed
	}

	consAddr, err := validator.GetConsAddr()
	if err != nil {
		return err
	}



	//




	info, found := k.GetValidatorSigningInfo(ctx, consAddr)
	if found {

		if info.Tombstoned {
			return types.ErrValidatorJailed
		}


		if ctx.BlockHeader().Time.Before(info.JailedUntil) {
			return types.ErrValidatorJailed
		}
	}

	k.sk.Unjail(ctx, consAddr)
	return nil
}
