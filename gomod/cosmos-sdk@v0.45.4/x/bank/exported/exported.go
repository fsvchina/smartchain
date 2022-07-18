package exported

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)



type GenesisBalance interface {
	GetAddress() sdk.AccAddress
	GetCoins() sdk.Coins
}
