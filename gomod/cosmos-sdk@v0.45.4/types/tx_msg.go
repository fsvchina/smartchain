package types

import (
	"github.com/gogo/protobuf/proto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

type (

	Msg interface {
		proto.Message



		ValidateBasic() error




		GetSigners() []AccAddress
	}



	Fee interface {
		GetGas() uint64
		GetAmount() Coins
	}



	Signature interface {
		GetPubKey() cryptotypes.PubKey
		GetSignature() []byte
	}


	Tx interface {

		GetMsgs() []Msg



		ValidateBasic() error
	}


	FeeTx interface {
		Tx
		GetGas() uint64
		GetFee() Coins
		FeePayer() AccAddress
		FeeGranter() AccAddress
	}


	TxWithMemo interface {
		Tx
		GetMemo() string
	}



	TxWithTimeoutHeight interface {
		Tx

		GetTimeoutHeight() uint64
	}
)


type TxDecoder func(txBytes []byte) (Tx, error)


type TxEncoder func(tx Tx) ([]byte, error)


func MsgTypeURL(msg Msg) string {
	return "/" + proto.MessageName(msg)
}
