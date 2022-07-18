package exported

import (
	"github.com/gogo/protobuf/proto"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"

	sdk "github.com/cosmos/cosmos-sdk/types"
)



type Evidence interface {
	proto.Message

	Route() string
	Type() string
	String() string
	Hash() tmbytes.HexBytes
	ValidateBasic() error


	GetHeight() int64
}



type ValidatorEvidence interface {
	Evidence


	GetConsensusAddress() sdk.ConsAddress


	GetValidatorPower() int64


	GetTotalPower() int64
}




type MsgSubmitEvidenceI interface {
	sdk.Msg

	GetEvidence() Evidence
	GetSubmitter() sdk.AccAddress
}
