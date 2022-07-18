package types

import (
	"fmt"

	"github.com/gogo/protobuf/proto"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/evidence/exported"
)


const (
	TypeMsgSubmitEvidence = "submit_evidence"
)

var (
	_ sdk.Msg                       = &MsgSubmitEvidence{}
	_ types.UnpackInterfacesMessage = MsgSubmitEvidence{}
	_ exported.MsgSubmitEvidenceI   = &MsgSubmitEvidence{}
)



func NewMsgSubmitEvidence(s sdk.AccAddress, evi exported.Evidence) (*MsgSubmitEvidence, error) {
	msg, ok := evi.(proto.Message)
	if !ok {
		return nil, fmt.Errorf("cannot proto marshal %T", evi)
	}
	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return nil, err
	}
	return &MsgSubmitEvidence{Submitter: s.String(), Evidence: any}, nil
}


func (m MsgSubmitEvidence) Route() string { return RouterKey }


func (m MsgSubmitEvidence) Type() string { return TypeMsgSubmitEvidence }


func (m MsgSubmitEvidence) ValidateBasic() error {
	if m.Submitter == "" {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, m.Submitter)
	}

	evi := m.GetEvidence()
	if evi == nil {
		return sdkerrors.Wrap(ErrInvalidEvidence, "missing evidence")
	}
	if err := evi.ValidateBasic(); err != nil {
		return err
	}

	return nil
}



func (m MsgSubmitEvidence) GetSignBytes() []byte {
	return sdk.MustSortJSON(ModuleCdc.MustMarshalJSON(&m))
}


func (m MsgSubmitEvidence) GetSigners() []sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(m.Submitter)
	if err != nil {
		return nil
	}

	return []sdk.AccAddress{accAddr}
}

func (m MsgSubmitEvidence) GetEvidence() exported.Evidence {
	evi, ok := m.Evidence.GetCachedValue().(exported.Evidence)
	if !ok {
		return nil
	}
	return evi
}

func (m MsgSubmitEvidence) GetSubmitter() sdk.AccAddress {
	accAddr, err := sdk.AccAddressFromBech32(m.Submitter)
	if err != nil {
		return nil
	}
	return accAddr
}

func (m MsgSubmitEvidence) UnpackInterfaces(ctx types.AnyUnpacker) error {
	var evi exported.Evidence
	return ctx.UnpackAny(m.Evidence, &evi)
}
