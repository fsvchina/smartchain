package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)


var _ sdk.Msg = &MsgVerifyInvariant{}



func NewMsgVerifyInvariant(sender sdk.AccAddress, invModeName, invRoute string) *MsgVerifyInvariant {
	return &MsgVerifyInvariant{
		Sender:              sender.String(),
		InvariantModuleName: invModeName,
		InvariantRoute:      invRoute,
	}
}

func (msg MsgVerifyInvariant) Route() string { return ModuleName }
func (msg MsgVerifyInvariant) Type() string  { return "verify_invariant" }


func (msg MsgVerifyInvariant) GetSigners() []sdk.AccAddress {
	sender, _ := sdk.AccAddressFromBech32(msg.Sender)
	return []sdk.AccAddress{sender}
}


func (msg MsgVerifyInvariant) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}


func (msg MsgVerifyInvariant) ValidateBasic() error {
	if msg.Sender == "" {
		return ErrNoSender
	}
	return nil
}


func (msg MsgVerifyInvariant) FullInvariantRoute() string {
	return msg.InvariantModuleName + "/" + msg.InvariantRoute
}
