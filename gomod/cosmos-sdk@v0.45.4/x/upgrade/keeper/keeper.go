package keeper

import (
	"encoding/binary"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"sort"

	"github.com/tendermint/tendermint/libs/log"
	tmos "github.com/tendermint/tendermint/libs/os"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	store "github.com/cosmos/cosmos-sdk/store/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	xp "github.com/cosmos/cosmos-sdk/x/upgrade/exported"
	"github.com/cosmos/cosmos-sdk/x/upgrade/types"
)


const UpgradeInfoFileName string = "upgrade-info.json"

type Keeper struct {
	homePath           string
	skipUpgradeHeights map[int64]bool
	storeKey           sdk.StoreKey
	cdc                codec.BinaryCodec
	upgradeHandlers    map[string]types.UpgradeHandler
	versionSetter      xp.ProtocolVersionSetter
	downgradeVerified  bool
}







func NewKeeper(skipUpgradeHeights map[int64]bool, storeKey sdk.StoreKey, cdc codec.BinaryCodec, homePath string, vs xp.ProtocolVersionSetter) Keeper {
	return Keeper{
		homePath:           homePath,
		skipUpgradeHeights: skipUpgradeHeights,
		storeKey:           storeKey,
		cdc:                cdc,
		upgradeHandlers:    map[string]types.UpgradeHandler{},
		versionSetter:      vs,
	}
}




func (k Keeper) SetUpgradeHandler(name string, upgradeHandler types.UpgradeHandler) {
	k.upgradeHandlers[name] = upgradeHandler
}


func (k Keeper) setProtocolVersion(ctx sdk.Context, v uint64) {
	store := ctx.KVStore(k.storeKey)
	versionBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(versionBytes, v)
	store.Set([]byte{types.ProtocolVersionByte}, versionBytes)
}


func (k Keeper) getProtocolVersion(ctx sdk.Context) uint64 {
	store := ctx.KVStore(k.storeKey)
	ok := store.Has([]byte{types.ProtocolVersionByte})
	if ok {
		pvBytes := store.Get([]byte{types.ProtocolVersionByte})
		protocolVersion := binary.BigEndian.Uint64(pvBytes)

		return protocolVersion
	}

	return 0
}


func (k Keeper) SetModuleVersionMap(ctx sdk.Context, vm module.VersionMap) {
	if len(vm) > 0 {
		store := ctx.KVStore(k.storeKey)
		versionStore := prefix.NewStore(store, []byte{types.VersionMapByte})



		sortedModNames := make([]string, 0, len(vm))

		for key := range vm {
			sortedModNames = append(sortedModNames, key)
		}
		sort.Strings(sortedModNames)

		for _, modName := range sortedModNames {
			ver := vm[modName]
			nameBytes := []byte(modName)
			verBytes := make([]byte, 8)
			binary.BigEndian.PutUint64(verBytes, ver)
			versionStore.Set(nameBytes, verBytes)
		}
	}
}



func (k Keeper) GetModuleVersionMap(ctx sdk.Context) module.VersionMap {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, []byte{types.VersionMapByte})

	vm := make(module.VersionMap)
	defer it.Close()
	for ; it.Valid(); it.Next() {
		moduleBytes := it.Key()

		name := string(moduleBytes[1:])
		moduleVersion := binary.BigEndian.Uint64(it.Value())
		vm[name] = moduleVersion
	}

	return vm
}


func (k Keeper) GetModuleVersions(ctx sdk.Context) []*types.ModuleVersion {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, []byte{types.VersionMapByte})
	defer it.Close()

	mv := make([]*types.ModuleVersion, 0)
	for ; it.Valid(); it.Next() {
		moduleBytes := it.Key()
		name := string(moduleBytes[1:])
		moduleVersion := binary.BigEndian.Uint64(it.Value())
		mv = append(mv, &types.ModuleVersion{
			Name:    name,
			Version: moduleVersion,
		})
	}
	return mv
}


func (k Keeper) getModuleVersion(ctx sdk.Context, name string) (uint64, bool) {
	store := ctx.KVStore(k.storeKey)
	it := sdk.KVStorePrefixIterator(store, []byte{types.VersionMapByte})
	defer it.Close()

	for ; it.Valid(); it.Next() {
		moduleName := string(it.Key()[1:])
		if moduleName == name {
			version := binary.BigEndian.Uint64(it.Value())
			return version, true
		}
	}
	return 0, false
}






func (k Keeper) ScheduleUpgrade(ctx sdk.Context, plan types.Plan) error {
	if err := plan.ValidateBasic(); err != nil {
		return err
	}



	if plan.Height < ctx.BlockHeight() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "upgrade cannot be scheduled in the past")
	}

	if k.GetDoneHeight(ctx, plan.Name) != 0 {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "upgrade with name %s has already been completed", plan.Name)
	}

	store := ctx.KVStore(k.storeKey)


	oldPlan, found := k.GetUpgradePlan(ctx)
	if found {
		k.ClearIBCState(ctx, oldPlan.Height)
	}

	bz := k.cdc.MustMarshal(&plan)
	store.Set(types.PlanKey(), bz)

	return nil
}


func (k Keeper) SetUpgradedClient(ctx sdk.Context, planHeight int64, bz []byte) error {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.UpgradedClientKey(planHeight), bz)
	return nil
}


func (k Keeper) GetUpgradedClient(ctx sdk.Context, height int64) ([]byte, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.UpgradedClientKey(height))
	if len(bz) == 0 {
		return nil, false
	}

	return bz, true
}



func (k Keeper) SetUpgradedConsensusState(ctx sdk.Context, planHeight int64, bz []byte) error {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.UpgradedConsStateKey(planHeight), bz)
	return nil
}


func (k Keeper) GetUpgradedConsensusState(ctx sdk.Context, lastHeight int64) ([]byte, bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.UpgradedConsStateKey(lastHeight))
	if len(bz) == 0 {
		return nil, false
	}

	return bz, true
}


func (k Keeper) GetLastCompletedUpgrade(ctx sdk.Context) (string, int64) {
	iter := sdk.KVStoreReversePrefixIterator(ctx.KVStore(k.storeKey), []byte{types.DoneByte})
	defer iter.Close()
	if iter.Valid() {
		return parseDoneKey(iter.Key()), int64(binary.BigEndian.Uint64(iter.Value()))
	}

	return "", 0
}


func parseDoneKey(key []byte) string {
	if len(key) < 2 {
		panic(fmt.Sprintf("expected key of length at least %d, got %d", 2, len(key)))
	}

	return string(key[1:])
}


func (k Keeper) GetDoneHeight(ctx sdk.Context, name string) int64 {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{types.DoneByte})
	bz := store.Get([]byte(name))
	if len(bz) == 0 {
		return 0
	}

	return int64(binary.BigEndian.Uint64(bz))
}


func (k Keeper) ClearIBCState(ctx sdk.Context, lastHeight int64) {

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.UpgradedClientKey(lastHeight))
	store.Delete(types.UpgradedConsStateKey(lastHeight))
}


func (k Keeper) ClearUpgradePlan(ctx sdk.Context) {

	oldPlan, found := k.GetUpgradePlan(ctx)
	if found {
		k.ClearIBCState(ctx, oldPlan.Height)
	}

	store := ctx.KVStore(k.storeKey)
	store.Delete(types.PlanKey())
}


func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}



func (k Keeper) GetUpgradePlan(ctx sdk.Context) (plan types.Plan, havePlan bool) {
	store := ctx.KVStore(k.storeKey)
	bz := store.Get(types.PlanKey())
	if bz == nil {
		return plan, false
	}

	k.cdc.MustUnmarshal(bz, &plan)
	return plan, true
}


func (k Keeper) setDone(ctx sdk.Context, name string) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), []byte{types.DoneByte})
	bz := make([]byte, 8)
	binary.BigEndian.PutUint64(bz, uint64(ctx.BlockHeight()))
	store.Set([]byte(name), bz)
}


func (k Keeper) HasHandler(name string) bool {
	_, ok := k.upgradeHandlers[name]
	return ok
}


func (k Keeper) ApplyUpgrade(ctx sdk.Context, plan types.Plan) {
	handler := k.upgradeHandlers[plan.Name]
	if handler == nil {
		panic("ApplyUpgrade should never be called without first checking HasHandler")
	}

	updatedVM, err := handler(ctx, plan, k.GetModuleVersionMap(ctx))
	if err != nil {
		panic(err)
	}

	k.SetModuleVersionMap(ctx, updatedVM)


	nextProtocolVersion := k.getProtocolVersion(ctx) + 1
	k.setProtocolVersion(ctx, nextProtocolVersion)
	if k.versionSetter != nil {

		k.versionSetter.SetProtocolVersion(nextProtocolVersion)
	}



	k.ClearIBCState(ctx, plan.Height)
	k.ClearUpgradePlan(ctx)
	k.setDone(ctx, plan.Name)
}


func (k Keeper) IsSkipHeight(height int64) bool {
	return k.skipUpgradeHeights[height]
}





func (k Keeper) DumpUpgradeInfoToDisk(height int64, name string) error {
	return k.DumpUpgradeInfoWithInfoToDisk(height, name, "")
}





func (k Keeper) DumpUpgradeInfoWithInfoToDisk(height int64, name string, info string) error {
	upgradeInfoFilePath, err := k.GetUpgradeInfoPath()
	if err != nil {
		return err
	}

	upgradeInfo := upgradeInfo{
		Name:   name,
		Height: height,
		Info:   info,
	}
	bz, err := json.Marshal(upgradeInfo)
	if err != nil {
		return err
	}

	return os.WriteFile(upgradeInfoFilePath, bz, 0o600)
}


func (k Keeper) GetUpgradeInfoPath() (string, error) {
	upgradeInfoFileDir := path.Join(k.getHomeDir(), "data")
	err := tmos.EnsureDir(upgradeInfoFileDir, os.ModePerm)
	if err != nil {
		return "", err
	}

	return filepath.Join(upgradeInfoFileDir, UpgradeInfoFileName), nil
}


func (k Keeper) getHomeDir() string {
	return k.homePath
}





func (k Keeper) ReadUpgradeInfoFromDisk() (store.UpgradeInfo, error) {
	var upgradeInfo store.UpgradeInfo

	upgradeInfoPath, err := k.GetUpgradeInfoPath()
	if err != nil {
		return upgradeInfo, err
	}

	data, err := ioutil.ReadFile(upgradeInfoPath)
	if err != nil {

		if os.IsNotExist(err) {
			return upgradeInfo, nil
		}

		return upgradeInfo, err
	}

	if err := json.Unmarshal(data, &upgradeInfo); err != nil {
		return upgradeInfo, err
	}

	return upgradeInfo, nil
}


type upgradeInfo struct {

	Name string `json:"name,omitempty"`

	Height int64 `json:"height,omitempty"`

	Info string `json:"info,omitempty"`
}


func (k *Keeper) SetDowngradeVerified(v bool) {
	k.downgradeVerified = v
}


func (k Keeper) DowngradeVerified() bool {
	return k.downgradeVerified
}
