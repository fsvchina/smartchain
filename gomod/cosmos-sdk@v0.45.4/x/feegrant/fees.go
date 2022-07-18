package feegrant

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)



type FeeAllowanceI interface {



	//



	//


	Accept(ctx sdk.Context, fee sdk.Coins, msgs []sdk.Msg) (remove bool, err error)



	ValidateBasic() error
}
