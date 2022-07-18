package authz

import (
	"github.com/gogo/protobuf/proto"

	sdk "github.com/cosmos/cosmos-sdk/types"
)



type Authorization interface {
	proto.Message



	MsgTypeURL() string



	Accept(ctx sdk.Context, msg sdk.Msg) (AcceptResponse, error)



	ValidateBasic() error
}



type AcceptResponse struct {

	Accept bool


	Delete bool


	Updated Authorization
}
