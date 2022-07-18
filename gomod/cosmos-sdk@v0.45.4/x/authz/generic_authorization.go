package authz

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var (
	_ Authorization = &GenericAuthorization{}
)


func NewGenericAuthorization(msgTypeURL string) *GenericAuthorization {
	return &GenericAuthorization{
		Msg: msgTypeURL,
	}
}


func (a GenericAuthorization) MsgTypeURL() string {
	return a.Msg
}


func (a GenericAuthorization) Accept(ctx sdk.Context, msg sdk.Msg) (AcceptResponse, error) {
	return AcceptResponse{Accept: true}, nil
}


func (a GenericAuthorization) ValidateBasic() error {
	return nil
}
