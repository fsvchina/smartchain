package baseapp

import (
	"fmt"
	"runtime/debug"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)




type RecoveryHandler func(recoveryObj interface{}) error




type recoveryMiddleware func(recoveryObj interface{}) (recoveryMiddleware, error)



func processRecovery(recoveryObj interface{}, middleware recoveryMiddleware) error {
	if middleware == nil {
		return nil
	}

	next, err := middleware(recoveryObj)
	if err != nil {
		return err
	}

	return processRecovery(recoveryObj, next)
}


func newRecoveryMiddleware(handler RecoveryHandler, next recoveryMiddleware) recoveryMiddleware {
	return func(recoveryObj interface{}) (recoveryMiddleware, error) {
		if err := handler(recoveryObj); err != nil {
			return nil, err
		}

		return next, nil
	}
}


func newOutOfGasRecoveryMiddleware(gasWanted uint64, ctx sdk.Context, next recoveryMiddleware) recoveryMiddleware {
	handler := func(recoveryObj interface{}) error {
		err, ok := recoveryObj.(sdk.ErrorOutOfGas)
		if !ok {
			return nil
		}

		return sdkerrors.Wrap(
			sdkerrors.ErrOutOfGas, fmt.Sprintf(
				"out of gas in location: %v; gasWanted: %d, gasUsed: %d",
				err.Descriptor, gasWanted, ctx.GasMeter().GasConsumed(),
			),
		)
	}

	return newRecoveryMiddleware(handler, next)
}


func newDefaultRecoveryMiddleware() recoveryMiddleware {
	handler := func(recoveryObj interface{}) error {
		return sdkerrors.Wrap(
			sdkerrors.ErrPanic, fmt.Sprintf(
				"recovered: %v\nstack:\n%v", recoveryObj, string(debug.Stack()),
			),
		)
	}

	return newRecoveryMiddleware(handler, nil)
}
