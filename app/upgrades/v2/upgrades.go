package v2

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"

	claimstypes "fs.video/smartchain/x/claims/types"
	erc20types "fs.video/smartchain/x/erc20/types"
)


func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {






		vm[claimstypes.ModuleName] = 1
		vm[erc20types.ModuleName] = 1

		return mm.RunMigrations(ctx, configurator, vm)
	}
}
