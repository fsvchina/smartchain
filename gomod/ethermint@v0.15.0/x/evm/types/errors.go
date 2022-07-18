package types

import (
	"errors"
	"fmt"

	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/ethereum/go-ethereum/common"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/vm"
)

const (
	codeErrInvalidState      = uint32(iota) + 2 
	codeErrExecutionReverted                    
	codeErrChainConfigNotFound
	codeErrInvalidChainConfig
	codeErrZeroAddress
	codeErrEmptyHash
	codeErrBloomNotFound
	codeErrTxReceiptNotFound
	codeErrCreateDisabled
	codeErrCallDisabled
	codeErrInvalidAmount
	codeErrInvalidGasPrice
	codeErrInvalidGasFee
	codeErrVMExecution
	codeErrInvalidRefund
	codeErrInconsistentGas
	codeErrInvalidGasCap
	codeErrInvalidBaseFee
	codeErrGasOverflow
	codeErrInvalidAccount
)

var ErrPostTxProcessing = errors.New("failed to execute post processing")

var (
	
	ErrInvalidState = sdkerrors.Register(ModuleName, codeErrInvalidState, "invalid storage state")

	
	ErrExecutionReverted = sdkerrors.Register(ModuleName, codeErrExecutionReverted, vm.ErrExecutionReverted.Error())

	
	ErrChainConfigNotFound = sdkerrors.Register(ModuleName, codeErrChainConfigNotFound, "chain configuration not found")

	
	ErrInvalidChainConfig = sdkerrors.Register(ModuleName, codeErrInvalidChainConfig, "invalid chain configuration")

	
	ErrZeroAddress = sdkerrors.Register(ModuleName, codeErrZeroAddress, "invalid zero address")

	
	ErrEmptyHash = sdkerrors.Register(ModuleName, codeErrEmptyHash, "empty hash")

	
	ErrBloomNotFound = sdkerrors.Register(ModuleName, codeErrBloomNotFound, "block bloom not found")

	
	ErrTxReceiptNotFound = sdkerrors.Register(ModuleName, codeErrTxReceiptNotFound, "transaction receipt not found")

	
	ErrCreateDisabled = sdkerrors.Register(ModuleName, codeErrCreateDisabled, "EVM Create operation is disabled")

	
	ErrCallDisabled = sdkerrors.Register(ModuleName, codeErrCallDisabled, "EVM Call operation is disabled")

	
	ErrInvalidAmount = sdkerrors.Register(ModuleName, codeErrInvalidAmount, "invalid transaction amount")

	
	ErrInvalidGasPrice = sdkerrors.Register(ModuleName, codeErrInvalidGasPrice, "invalid gas price")

	
	ErrInvalidGasFee = sdkerrors.Register(ModuleName, codeErrInvalidGasFee, "invalid gas fee")

	
	ErrVMExecution = sdkerrors.Register(ModuleName, codeErrVMExecution, "evm transaction execution failed")

	
	ErrInvalidRefund = sdkerrors.Register(ModuleName, codeErrInvalidRefund, "invalid gas refund amount")

	
	ErrInconsistentGas = sdkerrors.Register(ModuleName, codeErrInconsistentGas, "inconsistent gas")

	
	ErrInvalidGasCap = sdkerrors.Register(ModuleName, codeErrInvalidGasCap, "invalid gas cap")

	
	ErrInvalidBaseFee = sdkerrors.Register(ModuleName, codeErrInvalidBaseFee, "invalid base fee")

	
	ErrGasOverflow = sdkerrors.Register(ModuleName, codeErrGasOverflow, "gas computation overflow/underflow")

	
	ErrInvalidAccount = sdkerrors.Register(ModuleName, codeErrInvalidAccount, "account type is not a valid ethereum account")
)



func NewExecErrorWithReason(revertReason []byte) *RevertError {
	result := common.CopyBytes(revertReason)
	reason, errUnpack := abi.UnpackRevert(result)
	err := errors.New("execution reverted")
	if errUnpack == nil {
		err = fmt.Errorf("execution reverted: %v", reason)
	}
	return &RevertError{
		error:  err,
		reason: hexutil.Encode(result),
	}
}



type RevertError struct {
	error
	reason string 
}



func (e *RevertError) ErrorCode() int {
	return 3
}


func (e *RevertError) ErrorData() interface{} {
	return e.reason
}
