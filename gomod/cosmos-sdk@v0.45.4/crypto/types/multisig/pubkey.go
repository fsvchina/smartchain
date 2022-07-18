package multisig

import (
	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)



type PubKey interface {
	types.PubKey



	VerifyMultisignature(getSignBytes GetSignBytesFunc, sig *signing.MultiSignatureData) error


	GetPubKeys() []types.PubKey


	GetThreshold() uint
}




type GetSignBytesFunc func(mode signing.SignMode) ([]byte, error)
