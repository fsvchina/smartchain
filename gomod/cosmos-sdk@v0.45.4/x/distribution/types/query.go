package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
)



type QueryDelegatorTotalRewardsResponse struct {
	Rewards []DelegationDelegatorReward `json:"rewards" yaml:"rewards"`
	Total   sdk.DecCoins                `json:"total" yaml:"total"`
}


func NewQueryDelegatorTotalRewardsResponse(rewards []DelegationDelegatorReward,
	total sdk.DecCoins) QueryDelegatorTotalRewardsResponse {
	return QueryDelegatorTotalRewardsResponse{Rewards: rewards, Total: total}
}

func (res QueryDelegatorTotalRewardsResponse) String() string {
	out := "Delegator Total Rewards:\n"
	out += "  Rewards:"
	for _, reward := range res.Rewards {
		out += fmt.Sprintf(`  
	ValidatorAddress: %s
	Reward: %s`, reward.ValidatorAddress, reward.Reward)
	}
	out += fmt.Sprintf("\n  Total: %s\n", res.Total)
	return strings.TrimSpace(out)
}



func NewDelegationDelegatorReward(valAddr sdk.ValAddress,
	reward sdk.DecCoins) DelegationDelegatorReward {
	return DelegationDelegatorReward{ValidatorAddress: valAddr.String(), Reward: reward}
}
