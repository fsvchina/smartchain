package capability

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/capability/keeper"
	"github.com/cosmos/cosmos-sdk/x/capability/types"
)



func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	if err := k.InitializeIndex(ctx, genState.Index); err != nil {
		panic(err)
	}


	for _, genOwner := range genState.Owners {
		k.SetOwners(ctx, genOwner.Index, genOwner.IndexOwners)
	}

	k.InitMemStore(ctx)
}


func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	index := k.GetLatestIndex(ctx)
	owners := []types.GenesisOwners{}

	for i := uint64(1); i < index; i++ {
		capabilityOwners, ok := k.GetOwners(ctx, i)
		if !ok || len(capabilityOwners.Owners) == 0 {
			continue
		}

		genOwner := types.GenesisOwners{
			Index:       i,
			IndexOwners: capabilityOwners,
		}
		owners = append(owners, genOwner)
	}

	return &types.GenesisState{
		Index:  index,
		Owners: owners,
	}
}
