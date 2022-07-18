package keeper

import (
	"fmt"

	gogotypes "github.com/gogo/protobuf/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/auth/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)


type AccountKeeperI interface {

	NewAccountWithAddress(sdk.Context, sdk.AccAddress) types.AccountI


	NewAccount(sdk.Context, types.AccountI) types.AccountI


	HasAccount(sdk.Context, sdk.AccAddress) bool


	GetAccount(sdk.Context, sdk.AccAddress) types.AccountI


	SetAccount(sdk.Context, types.AccountI)


	RemoveAccount(sdk.Context, types.AccountI)


	IterateAccounts(sdk.Context, func(types.AccountI) bool)


	GetPubKey(sdk.Context, sdk.AccAddress) (cryptotypes.PubKey, error)


	GetSequence(sdk.Context, sdk.AccAddress) (uint64, error)


	GetNextAccountNumber(sdk.Context) uint64
}



type AccountKeeper struct {
	key           sdk.StoreKey
	cdc           codec.BinaryCodec
	paramSubspace paramtypes.Subspace
	permAddrs     map[string]types.PermissionsForAddress


	proto func() types.AccountI
}

var _ AccountKeeperI = &AccountKeeper{}







func NewAccountKeeper(
	cdc codec.BinaryCodec, key sdk.StoreKey, paramstore paramtypes.Subspace, proto func() types.AccountI,
	maccPerms map[string][]string,
) AccountKeeper {


	if !paramstore.HasKeyTable() {
		paramstore = paramstore.WithKeyTable(types.ParamKeyTable())
	}

	permAddrs := make(map[string]types.PermissionsForAddress)
	for name, perms := range maccPerms {
		permAddrs[name] = types.NewPermissionsForAddress(name, perms)
	}

	return AccountKeeper{
		key:           key,
		proto:         proto,
		cdc:           cdc,
		paramSubspace: paramstore,
		permAddrs:     permAddrs,
	}
}


func (ak AccountKeeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}


func (ak AccountKeeper) GetPubKey(ctx sdk.Context, addr sdk.AccAddress) (cryptotypes.PubKey, error) {
	acc := ak.GetAccount(ctx, addr)
	if acc == nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "account %s does not exist", addr)
	}

	return acc.GetPubKey(), nil
}


func (ak AccountKeeper) GetSequence(ctx sdk.Context, addr sdk.AccAddress) (uint64, error) {
	acc := ak.GetAccount(ctx, addr)
	if acc == nil {
		return 0, sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "account %s does not exist", addr)
	}

	return acc.GetSequence(), nil
}



func (ak AccountKeeper) GetNextAccountNumber(ctx sdk.Context) uint64 {
	var accNumber uint64
	store := ctx.KVStore(ak.key)

	bz := store.Get(types.GlobalAccountNumberKey)
	if bz == nil {

		accNumber = 0
	} else {
		val := gogotypes.UInt64Value{}

		err := ak.cdc.Unmarshal(bz, &val)
		if err != nil {
			panic(err)
		}

		accNumber = val.GetValue()
	}

	bz = ak.cdc.MustMarshal(&gogotypes.UInt64Value{Value: accNumber + 1})
	store.Set(types.GlobalAccountNumberKey, bz)

	return accNumber
}



func (ak AccountKeeper) ValidatePermissions(macc types.ModuleAccountI) error {
	permAddr := ak.permAddrs[macc.GetName()]
	for _, perm := range macc.GetPermissions() {
		if !permAddr.HasPermission(perm) {
			return fmt.Errorf("invalid module permission %s", perm)
		}
	}

	return nil
}


func (ak AccountKeeper) GetModuleAddress(moduleName string) sdk.AccAddress {
	permAddr, ok := ak.permAddrs[moduleName]
	if !ok {
		return nil
	}

	return permAddr.GetAddress()
}


func (ak AccountKeeper) GetModuleAddressAndPermissions(moduleName string) (addr sdk.AccAddress, permissions []string) {
	permAddr, ok := ak.permAddrs[moduleName]
	if !ok {
		return addr, permissions
	}

	return permAddr.GetAddress(), permAddr.GetPermissions()
}



func (ak AccountKeeper) GetModuleAccountAndPermissions(ctx sdk.Context, moduleName string) (types.ModuleAccountI, []string) {
	addr, perms := ak.GetModuleAddressAndPermissions(moduleName)
	if addr == nil {
		return nil, []string{}
	}

	acc := ak.GetAccount(ctx, addr)
	if acc != nil {
		macc, ok := acc.(types.ModuleAccountI)
		if !ok {
			panic("account is not a module account")
		}
		return macc, perms
	}


	macc := types.NewEmptyModuleAccount(moduleName, perms...)
	maccI := (ak.NewAccount(ctx, macc)).(types.ModuleAccountI)
	ak.SetModuleAccount(ctx, maccI)

	return maccI, perms
}



func (ak AccountKeeper) GetModuleAccount(ctx sdk.Context, moduleName string) types.ModuleAccountI {
	acc, _ := ak.GetModuleAccountAndPermissions(ctx, moduleName)
	return acc
}


func (ak AccountKeeper) SetModuleAccount(ctx sdk.Context, macc types.ModuleAccountI) {
	ak.SetAccount(ctx, macc)
}

func (ak AccountKeeper) decodeAccount(bz []byte) types.AccountI {
	acc, err := ak.UnmarshalAccount(bz)
	if err != nil {
		panic(err)
	}

	return acc
}


func (ak AccountKeeper) MarshalAccount(accountI types.AccountI) ([]byte, error) {
	return ak.cdc.MarshalInterface(accountI)
}



func (ak AccountKeeper) UnmarshalAccount(bz []byte) (types.AccountI, error) {
	var acc types.AccountI
	return acc, ak.cdc.UnmarshalInterface(bz, &acc)
}


func (ak AccountKeeper) GetCodec() codec.BinaryCodec { return ak.cdc }
