package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
)



//




//








//

type UpgradeHandler func(ctx sdk.Context, plan Plan, fromVM module.VersionMap) (module.VersionMap, error)
