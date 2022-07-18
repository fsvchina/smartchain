package signing

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/types/tx/signing"

	sdk "github.com/cosmos/cosmos-sdk/types"
)



type SignModeHandlerMap struct {
	defaultMode      signing.SignMode
	modes            []signing.SignMode
	signModeHandlers map[signing.SignMode]SignModeHandler
}

var _ SignModeHandler = SignModeHandlerMap{}


func NewSignModeHandlerMap(defaultMode signing.SignMode, handlers []SignModeHandler) SignModeHandlerMap {
	handlerMap := make(map[signing.SignMode]SignModeHandler)
	var modes []signing.SignMode

	for _, h := range handlers {
		for _, m := range h.Modes() {
			if _, have := handlerMap[m]; have {
				panic(fmt.Errorf("duplicate sign mode handler for mode %s", m))
			}
			handlerMap[m] = h
			modes = append(modes, m)
		}
	}

	return SignModeHandlerMap{
		defaultMode:      defaultMode,
		modes:            modes,
		signModeHandlers: handlerMap,
	}
}


func (h SignModeHandlerMap) DefaultMode() signing.SignMode {
	return h.defaultMode
}


func (h SignModeHandlerMap) Modes() []signing.SignMode {
	return h.modes
}


func (h SignModeHandlerMap) GetSignBytes(mode signing.SignMode, data SignerData, tx sdk.Tx) ([]byte, error) {
	handler, found := h.signModeHandlers[mode]
	if !found {
		return nil, fmt.Errorf("can't verify sign mode %s", mode.String())
	}
	return handler.GetSignBytes(mode, data, tx)
}
