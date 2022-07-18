package v040

import (
	v039auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v039"
	v036supply "github.com/cosmos/cosmos-sdk/x/bank/legacy/v036"
	v038bank "github.com/cosmos/cosmos-sdk/x/bank/legacy/v038"
	"github.com/cosmos/cosmos-sdk/x/bank/types"
)



//



func Migrate(
	bankGenState v038bank.GenesisState,
	authGenState v039auth.GenesisState,
	supplyGenState v036supply.GenesisState,
) *types.GenesisState {
	balances := make([]types.Balance, len(authGenState.Accounts))
	for i, acc := range authGenState.Accounts {
		balances[i] = types.Balance{
			Address: acc.GetAddress().String(),
			Coins:   acc.GetCoins(),
		}
	}

	return &types.GenesisState{
		Params: types.Params{
			SendEnabled:        []*types.SendEnabled{},
			DefaultSendEnabled: bankGenState.SendEnabled,
		},
		Balances:      balances,
		Supply:        supplyGenState.Supply,
		DenomMetadata: []types.Metadata{},
	}
}
