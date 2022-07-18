package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/ante"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
)



type Keeper struct {
	cdc        codec.BinaryCodec
	storeKey   sdk.StoreKey
	authKeeper feegrant.AccountKeeper
}

var _ ante.FeegrantKeeper = &Keeper{}


func NewKeeper(cdc codec.BinaryCodec, storeKey sdk.StoreKey, ak feegrant.AccountKeeper) Keeper {
	return Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		authKeeper: ak,
	}
}


func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", feegrant.ModuleName))
}


func (k Keeper) GrantAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress, feeAllowance feegrant.FeeAllowanceI) error {


	granteeAcc := k.authKeeper.GetAccount(ctx, grantee)
	if granteeAcc == nil {
		granteeAcc = k.authKeeper.NewAccountWithAddress(ctx, grantee)
		k.authKeeper.SetAccount(ctx, granteeAcc)
	}

	store := ctx.KVStore(k.storeKey)
	key := feegrant.FeeAllowanceKey(granter, grantee)
	grant, err := feegrant.NewGrant(granter, grantee, feeAllowance)
	if err != nil {
		return err
	}

	bz, err := k.cdc.Marshal(&grant)
	if err != nil {
		return err
	}

	store.Set(key, bz)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			feegrant.EventTypeSetFeeGrant,
			sdk.NewAttribute(feegrant.AttributeKeyGranter, grant.Granter),
			sdk.NewAttribute(feegrant.AttributeKeyGrantee, grant.Grantee),
		),
	)

	return nil
}


func (k Keeper) revokeAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress) error {
	_, err := k.getGrant(ctx, granter, grantee)
	if err != nil {
		return err
	}

	store := ctx.KVStore(k.storeKey)
	key := feegrant.FeeAllowanceKey(granter, grantee)
	store.Delete(key)

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			feegrant.EventTypeRevokeFeeGrant,
			sdk.NewAttribute(feegrant.AttributeKeyGranter, granter.String()),
			sdk.NewAttribute(feegrant.AttributeKeyGrantee, grantee.String()),
		),
	)
	return nil
}




func (k Keeper) GetAllowance(ctx sdk.Context, granter, grantee sdk.AccAddress) (feegrant.FeeAllowanceI, error) {
	grant, err := k.getGrant(ctx, granter, grantee)
	if err != nil {
		return nil, err
	}

	return grant.GetGrant()
}


func (k Keeper) getGrant(ctx sdk.Context, granter sdk.AccAddress, grantee sdk.AccAddress) (*feegrant.Grant, error) {
	store := ctx.KVStore(k.storeKey)
	key := feegrant.FeeAllowanceKey(granter, grantee)
	bz := store.Get(key)
	if len(bz) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnauthorized, "fee-grant not found")
	}

	var feegrant feegrant.Grant
	if err := k.cdc.Unmarshal(bz, &feegrant); err != nil {
		return nil, err
	}

	return &feegrant, nil
}




func (k Keeper) IterateAllFeeAllowances(ctx sdk.Context, cb func(grant feegrant.Grant) bool) error {
	store := ctx.KVStore(k.storeKey)
	iter := sdk.KVStorePrefixIterator(store, feegrant.FeeAllowanceKeyPrefix)
	defer iter.Close()

	stop := false
	for ; iter.Valid() && !stop; iter.Next() {
		bz := iter.Value()
		var feeGrant feegrant.Grant
		if err := k.cdc.Unmarshal(bz, &feeGrant); err != nil {
			return err
		}

		stop = cb(feeGrant)
	}

	return nil
}


func (k Keeper) UseGrantedFees(ctx sdk.Context, granter, grantee sdk.AccAddress, fee sdk.Coins, msgs []sdk.Msg) error {
	f, err := k.getGrant(ctx, granter, grantee)
	if err != nil {
		return err
	}

	grant, err := f.GetGrant()
	if err != nil {
		return err
	}

	remove, err := grant.Accept(ctx, fee, msgs)

	if remove {

		k.revokeAllowance(ctx, granter, grantee)
		if err != nil {
			return err
		}

		emitUseGrantEvent(ctx, granter.String(), grantee.String())

		return nil
	}

	if err != nil {
		return err
	}

	emitUseGrantEvent(ctx, granter.String(), grantee.String())


	return k.GrantAllowance(ctx, granter, grantee, grant)
}

func emitUseGrantEvent(ctx sdk.Context, granter, grantee string) {
	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			feegrant.EventTypeUseFeeGrant,
			sdk.NewAttribute(feegrant.AttributeKeyGranter, granter),
			sdk.NewAttribute(feegrant.AttributeKeyGrantee, grantee),
		),
	)
}


func (k Keeper) InitGenesis(ctx sdk.Context, data *feegrant.GenesisState) error {
	for _, f := range data.Allowances {
		granter, err := sdk.AccAddressFromBech32(f.Granter)
		if err != nil {
			return err
		}
		grantee, err := sdk.AccAddressFromBech32(f.Grantee)
		if err != nil {
			return err
		}

		grant, err := f.GetGrant()
		if err != nil {
			return err
		}

		err = k.GrantAllowance(ctx, granter, grantee, grant)
		if err != nil {
			return err
		}
	}
	return nil
}


func (k Keeper) ExportGenesis(ctx sdk.Context) (*feegrant.GenesisState, error) {
	var grants []feegrant.Grant

	err := k.IterateAllFeeAllowances(ctx, func(grant feegrant.Grant) bool {
		grants = append(grants, grant)
		return false
	})

	return &feegrant.GenesisState{
		Allowances: grants,
	}, err
}
