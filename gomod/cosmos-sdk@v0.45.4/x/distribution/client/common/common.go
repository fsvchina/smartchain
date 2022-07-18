package common

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/distribution/types"
)



func QueryDelegationRewards(clientCtx client.Context, delAddr, valAddr string) ([]byte, int64, error) {
	delegatorAddr, err := sdk.AccAddressFromBech32(delAddr)
	if err != nil {
		return nil, 0, err
	}

	validatorAddr, err := sdk.ValAddressFromBech32(valAddr)
	if err != nil {
		return nil, 0, err
	}

	params := types.NewQueryDelegationRewardsParams(delegatorAddr, validatorAddr)
	bz, err := clientCtx.LegacyAmino.MarshalJSON(params)
	if err != nil {
		return nil, 0, fmt.Errorf("failed to marshal params: %w", err)
	}

	route := fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDelegationRewards)
	return clientCtx.QueryWithData(route, bz)
}



func QueryDelegatorValidators(clientCtx client.Context, delegatorAddr sdk.AccAddress) ([]byte, error) {
	res, _, err := clientCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryDelegatorValidators),
		clientCtx.LegacyAmino.MustMarshalJSON(types.NewQueryDelegatorParams(delegatorAddr)),
	)
	return res, err
}


func QueryValidatorCommission(clientCtx client.Context, validatorAddr sdk.ValAddress) ([]byte, error) {
	res, _, err := clientCtx.QueryWithData(
		fmt.Sprintf("custom/%s/%s", types.QuerierRoute, types.QueryValidatorCommission),
		clientCtx.LegacyAmino.MustMarshalJSON(types.NewQueryValidatorCommissionParams(validatorAddr)),
	)
	return res, err
}



func WithdrawAllDelegatorRewards(clientCtx client.Context, delegatorAddr sdk.AccAddress) ([]sdk.Msg, error) {


	bz, err := QueryDelegatorValidators(clientCtx, delegatorAddr)
	if err != nil {
		return nil, err
	}

	var validators []sdk.ValAddress
	if err := clientCtx.LegacyAmino.UnmarshalJSON(bz, &validators); err != nil {
		return nil, err
	}


	msgs := make([]sdk.Msg, 0, len(validators))
	for _, valAddr := range validators {
		msg := types.NewMsgWithdrawDelegatorReward(delegatorAddr, valAddr)
		if err := msg.ValidateBasic(); err != nil {
			return nil, err
		}
		msgs = append(msgs, msg)
	}

	return msgs, nil
}



func WithdrawValidatorRewardsAndCommission(validatorAddr sdk.ValAddress) ([]sdk.Msg, error) {
	commissionMsg := types.NewMsgWithdrawValidatorCommission(validatorAddr)
	if err := commissionMsg.ValidateBasic(); err != nil {
		return nil, err
	}


	rewardMsg := types.NewMsgWithdrawDelegatorReward(sdk.AccAddress(validatorAddr.Bytes()), validatorAddr)
	if err := rewardMsg.ValidateBasic(); err != nil {
		return nil, err
	}

	return []sdk.Msg{commissionMsg, rewardMsg}, nil
}
