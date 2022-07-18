package keeper

import (
	"math/big"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	authante "github.com/cosmos/cosmos-sdk/x/auth/ante"

	evmtypes "github.com/tharsis/ethermint/x/evm/types"

	"github.com/ethereum/go-ethereum/core"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)


func (k Keeper) DeductTxCostsFromUserBalance(
	ctx sdk.Context,
	msgEthTx evmtypes.MsgEthereumTx,
	txData evmtypes.TxData,
	denom string,
	homestead, istanbul, london bool,
) (sdk.Coins, error) {
	isContractCreation := txData.GetTo() == nil


	signerAcc, err := authante.GetSignerAcc(ctx, k.accountKeeper, msgEthTx.GetFrom())
	if err != nil {
		return nil, sdkerrors.Wrapf(err, "account not found for sender %s", msgEthTx.From)
	}

	gasLimit := txData.GetGas()

	var accessList ethtypes.AccessList
	if txData.GetAccessList() != nil {
		accessList = txData.GetAccessList()
	}

	intrinsicGas, err := core.IntrinsicGas(txData.GetData(), accessList, isContractCreation, homestead, istanbul)
	if err != nil {
		return nil, sdkerrors.Wrapf(
			err,
			"failed to retrieve intrinsic gas, contract creation = %t; homestead = %t, istanbul = %t",
			isContractCreation, homestead, istanbul,
		)
	}


	if ctx.IsCheckTx() && gasLimit < intrinsicGas {
		return nil, sdkerrors.Wrapf(
			sdkerrors.ErrOutOfGas,
			"gas limit too low: %d (gas limit) < %d (intrinsic gas)", gasLimit, intrinsicGas,
		)
	}

	var feeAmt *big.Int

	feeMktParams := k.feeMarketKeeper.GetParams(ctx)
	if london && !feeMktParams.NoBaseFee && txData.TxType() == ethtypes.DynamicFeeTxType {
		baseFee := k.feeMarketKeeper.GetBaseFee(ctx)
		if txData.GetGasFeeCap().Cmp(baseFee) < 0 {
			return nil, sdkerrors.Wrapf(sdkerrors.ErrInsufficientFee, "the tx gasfeecap is lower than the tx baseFee: %s (gasfeecap), %s (basefee) ", txData.GetGasFeeCap(), baseFee)
		}
		feeAmt = txData.EffectiveFee(baseFee)
	} else {
		feeAmt = txData.Fee()
	}

	if feeAmt.Sign() == 0 {

		return sdk.NewCoins(), nil
	}

	fees := sdk.Coins{sdk.NewCoin(denom, sdk.NewIntFromBigInt(feeAmt))}


	if err := authante.DeductFees(k.bankKeeper, ctx, signerAcc, fees); err != nil {
		return nil, sdkerrors.Wrapf(
			err,
			"failed to deduct full gas cost %s from the user %s balance",
			fees, msgEthTx.From,
		)
	}
	return fees, nil
}



func CheckSenderBalance(
	balance sdk.Int,
	txData evmtypes.TxData,
) error {
	cost := txData.Cost()

	if cost.Sign() < 0 {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInvalidCoins,
			"tx cost (%s) is negative and invalid", cost,
		)
	}

	if balance.IsNegative() || balance.BigInt().Cmp(cost) < 0 {
		return sdkerrors.Wrapf(
			sdkerrors.ErrInsufficientFunds,
			"sender balance < tx cost (%s < %s)", balance, txData.Cost(),
		)
	}
	return nil
}
