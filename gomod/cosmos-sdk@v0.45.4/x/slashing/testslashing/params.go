package testslashing

import (
	"github.com/cosmos/cosmos-sdk/x/slashing/types"
)




func TestParams() types.Params {
	params := types.DefaultParams()
	params.SignedBlocksWindow = 1000
	params.DowntimeJailDuration = 60 * 60

	return params
}
