package keeper

import (
	"fmt"

	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/evidence/exported"
	"github.com/cosmos/cosmos-sdk/x/evidence/types"
)




type Keeper struct {
	cdc            codec.BinaryCodec
	storeKey       sdk.StoreKey
	router         types.Router
	stakingKeeper  types.StakingKeeper
	slashingKeeper types.SlashingKeeper
}

func NewKeeper(
	cdc codec.BinaryCodec, storeKey sdk.StoreKey, stakingKeeper types.StakingKeeper,
	slashingKeeper types.SlashingKeeper,
) *Keeper {

	return &Keeper{
		cdc:            cdc,
		storeKey:       storeKey,
		stakingKeeper:  stakingKeeper,
		slashingKeeper: slashingKeeper,
	}
}


func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}





func (k *Keeper) SetRouter(rtr types.Router) {



	if !rtr.Sealed() {
		rtr.Seal()
	}
	if k.router != nil {
		panic(fmt.Sprintf("attempting to reset router on x/%s", types.ModuleName))
	}

	k.router = rtr
}



func (k Keeper) GetEvidenceHandler(evidenceRoute string) (types.Handler, error) {
	if !k.router.HasRoute(evidenceRoute) {
		return nil, sdkerrors.Wrap(types.ErrNoEvidenceHandlerExists, evidenceRoute)
	}

	return k.router.GetRoute(evidenceRoute), nil
}





func (k Keeper) SubmitEvidence(ctx sdk.Context, evidence exported.Evidence) error {
	if _, ok := k.GetEvidence(ctx, evidence.Hash()); ok {
		return sdkerrors.Wrap(types.ErrEvidenceExists, evidence.Hash().String())
	}
	if !k.router.HasRoute(evidence.Route()) {
		return sdkerrors.Wrap(types.ErrNoEvidenceHandlerExists, evidence.Route())
	}

	handler := k.router.GetRoute(evidence.Route())
	if err := handler(ctx, evidence); err != nil {
		return sdkerrors.Wrap(types.ErrInvalidEvidence, err.Error())
	}

	ctx.EventManager().EmitEvent(
		sdk.NewEvent(
			types.EventTypeSubmitEvidence,
			sdk.NewAttribute(types.AttributeKeyEvidenceHash, evidence.Hash().String()),
		),
	)

	k.SetEvidence(ctx, evidence)
	return nil
}


func (k Keeper) SetEvidence(ctx sdk.Context, evidence exported.Evidence) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEvidence)
	store.Set(evidence.Hash(), k.MustMarshalEvidence(evidence))
}



func (k Keeper) GetEvidence(ctx sdk.Context, hash tmbytes.HexBytes) (exported.Evidence, bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEvidence)

	bz := store.Get(hash)
	if len(bz) == 0 {
		return nil, false
	}

	return k.MustUnmarshalEvidence(bz), true
}




func (k Keeper) IterateEvidence(ctx sdk.Context, cb func(exported.Evidence) bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefixEvidence)
	iterator := sdk.KVStorePrefixIterator(store, nil)

	defer iterator.Close()
	for ; iterator.Valid(); iterator.Next() {
		evidence := k.MustUnmarshalEvidence(iterator.Value())

		if cb(evidence) {
			break
		}
	}
}


func (k Keeper) GetAllEvidence(ctx sdk.Context) (evidence []exported.Evidence) {
	k.IterateEvidence(ctx, func(e exported.Evidence) bool {
		evidence = append(evidence, e)
		return false
	})

	return evidence
}



func (k Keeper) MustUnmarshalEvidence(bz []byte) exported.Evidence {
	evidence, err := k.UnmarshalEvidence(bz)
	if err != nil {
		panic(fmt.Errorf("failed to decode evidence: %w", err))
	}

	return evidence
}



func (k Keeper) MustMarshalEvidence(evidence exported.Evidence) []byte {
	bz, err := k.MarshalEvidence(evidence)
	if err != nil {
		panic(fmt.Errorf("failed to encode evidence: %w", err))
	}

	return bz
}


func (k Keeper) MarshalEvidence(evidenceI exported.Evidence) ([]byte, error) {
	return k.cdc.MarshalInterface(evidenceI)
}



func (k Keeper) UnmarshalEvidence(bz []byte) (exported.Evidence, error) {
	var evi exported.Evidence
	return evi, k.cdc.UnmarshalInterface(bz, &evi)
}
