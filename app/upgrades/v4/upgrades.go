package v4

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
	ibcclientkeeper "github.com/cosmos/ibc-go/v3/modules/core/02-client/keeper"
	ibcclienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"

	"fs.video/smartchain/types"
)


func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	clientKeeper ibcclientkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger()




		if types.IsMainnet(ctx.ChainID()) {
			if err := UpdateIBCClients(ctx, clientKeeper); err != nil {

				logger.Error("FAILED TO UPDATE IBC CLIENTS", "error", err.Error())
			}
		}


		return mm.RunMigrations(ctx, configurator, vm)
	}
}


func UpdateIBCClients(ctx sdk.Context, k ibcclientkeeper.Keeper) error {
	proposalOsmosis := &ibcclienttypes.ClientUpdateProposal{
		Title:              "Update expired Osmosis IBC client",
		Description:        "Update the existing expired Cosmos Hub IBC client on Evmos (07-tendermint-0) in order to resume packet transfers between both chains.",
		SubjectClientId:    ExpiredOsmosisClient,
		SubstituteClientId: ActiveOsmosisClient,
	}

	proposalCosmosHub := &ibcclienttypes.ClientUpdateProposal{
		Title:              "Update expired Cosmos Hub IBC client",
		Description:        "Update the existing expired Cosmos Hub IBC client on Evmos (07-tendermint-3) in order to resume packet transfers between both chains.",
		SubjectClientId:    ExpiredCosmosHubClient,
		SubstituteClientId: ActiveCosmosHubClient,
	}

	if err := k.ClientUpdateProposal(ctx, proposalOsmosis); err != nil {
		return sdkerrors.Wrap(err, "failed to update Osmosis IBC client")
	}

	if err := k.ClientUpdateProposal(ctx, proposalCosmosHub); err != nil {
		return sdkerrors.Wrap(err, "failed to update Cosmos Hub IBC client")
	}

	return nil
}
