package signing

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)



type SignModeHandler interface {


	DefaultMode() signing.SignMode


	Modes() []signing.SignMode



	GetSignBytes(mode signing.SignMode, data SignerData, tx sdk.Tx) ([]byte, error)
}



type SignerData struct {

	ChainID string


	AccountNumber uint64





	Sequence uint64
}
