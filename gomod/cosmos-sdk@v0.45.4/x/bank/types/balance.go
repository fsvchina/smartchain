package types

import (
	"bytes"
	"encoding/json"
	"fmt"
	"sort"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/bank/exported"
)

var _ exported.GenesisBalance = (*Balance)(nil)


func (b Balance) GetAddress() sdk.AccAddress {
	addr, err := sdk.AccAddressFromBech32(b.Address)
	if err != nil {
		panic(fmt.Errorf("couldn't convert %q to account address: %v", b.Address, err))
	}

	return addr
}


func (b Balance) GetCoins() sdk.Coins {
	return b.Coins
}


func (b Balance) Validate() error {
	if _, err := sdk.AccAddressFromBech32(b.Address); err != nil {
		return err
	}

	if err := b.Coins.Validate(); err != nil {
		return err
	}

	return nil
}

type balanceByAddress struct {
	addresses []sdk.AccAddress
	balances  []Balance
}

func (b balanceByAddress) Len() int { return len(b.addresses) }
func (b balanceByAddress) Less(i, j int) bool {
	return bytes.Compare(b.addresses[i], b.addresses[j]) < 0
}
func (b balanceByAddress) Swap(i, j int) {
	b.addresses[i], b.addresses[j] = b.addresses[j], b.addresses[i]
	b.balances[i], b.balances[j] = b.balances[j], b.balances[i]
}


func SanitizeGenesisBalances(balances []Balance) []Balance {










	addresses := make([]sdk.AccAddress, len(balances))
	for i := range balances {
		addr, _ := sdk.AccAddressFromBech32(balances[i].Address)
		addresses[i] = addr
	}


	sort.Sort(balanceByAddress{addresses: addresses, balances: balances})

	return balances
}


type GenesisBalancesIterator struct{}




func (GenesisBalancesIterator) IterateGenesisBalances(
	cdc codec.JSONCodec, appState map[string]json.RawMessage, cb func(exported.GenesisBalance) (stop bool),
) {
	for _, balance := range GetGenesisStateFromAppState(cdc, appState).Balances {
		if cb(balance) {
			break
		}
	}
}
