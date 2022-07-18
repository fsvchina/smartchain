package keeper

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/libs/log"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/crisis/types"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)


type Keeper struct {
	routes         []types.InvarRoute
	paramSpace     paramtypes.Subspace
	invCheckPeriod uint

	supplyKeeper types.SupplyKeeper

	feeCollectorName string
}


func NewKeeper(
	paramSpace paramtypes.Subspace, invCheckPeriod uint, supplyKeeper types.SupplyKeeper,
	feeCollectorName string,
) Keeper {


	if !paramSpace.HasKeyTable() {
		paramSpace = paramSpace.WithKeyTable(types.ParamKeyTable())
	}

	return Keeper{
		routes:           make([]types.InvarRoute, 0),
		paramSpace:       paramSpace,
		invCheckPeriod:   invCheckPeriod,
		supplyKeeper:     supplyKeeper,
		feeCollectorName: feeCollectorName,
	}
}


func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", "x/"+types.ModuleName)
}


func (k *Keeper) RegisterRoute(moduleName, route string, invar sdk.Invariant) {
	invarRoute := types.NewInvarRoute(moduleName, route, invar)
	k.routes = append(k.routes, invarRoute)
}


func (k Keeper) Routes() []types.InvarRoute {
	return k.routes
}


func (k Keeper) Invariants() []sdk.Invariant {
	invars := make([]sdk.Invariant, len(k.routes))
	for i, route := range k.routes {
		invars[i] = route.Invar
	}
	return invars
}



func (k Keeper) AssertInvariants(ctx sdk.Context) {
	logger := k.Logger(ctx)

	start := time.Now()
	invarRoutes := k.Routes()
	n := len(invarRoutes)
	for i, ir := range invarRoutes {
		logger.Info("asserting crisis invariants", "inv", fmt.Sprint(i, "/", n), "name", ir.FullRoute())
		if res, stop := ir.Invar(ctx); stop {


			panic(fmt.Errorf("invariant broken: %s\n"+
				"\tCRITICAL please submit the following transaction:\n"+
				"\t\t tx crisis invariant-broken %s %s", res, ir.ModuleName, ir.Route))
		}
	}

	diff := time.Since(start)
	logger.Info("asserted all invariants", "duration", diff, "height", ctx.BlockHeight())
}


func (k Keeper) InvCheckPeriod() uint { return k.invCheckPeriod }


func (k Keeper) SendCoinsFromAccountToFeeCollector(ctx sdk.Context, senderAddr sdk.AccAddress, amt sdk.Coins) error {
	return k.supplyKeeper.SendCoinsFromAccountToModule(ctx, senderAddr, k.feeCollectorName, amt)
}
