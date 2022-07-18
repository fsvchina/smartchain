package types

import (
	"fmt"
	"strings"
	"time"

	"github.com/gogo/protobuf/proto"
	yaml "gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


const DefaultStartingProposalID uint64 = 1


func NewProposal(content Content, id uint64, submitTime, depositEndTime time.Time) (Proposal, error) {
	msg, ok := content.(proto.Message)
	if !ok {
		return Proposal{}, fmt.Errorf("%T does not implement proto.Message", content)
	}

	any, err := types.NewAnyWithValue(msg)
	if err != nil {
		return Proposal{}, err
	}

	p := Proposal{
		Content:          any,
		ProposalId:       id,
		Status:           StatusDepositPeriod,
		FinalTallyResult: EmptyTallyResult(),
		TotalDeposit:     sdk.NewCoins(),
		SubmitTime:       submitTime,
		DepositEndTime:   depositEndTime,
	}

	return p, nil
}


func (p Proposal) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}


func (p Proposal) GetContent() Content {
	content, ok := p.Content.GetCachedValue().(Content)
	if !ok {
		return nil
	}
	return content
}

func (p Proposal) ProposalType() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.ProposalType()
}

func (p Proposal) ProposalRoute() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.ProposalRoute()
}

func (p Proposal) GetTitle() string {
	content := p.GetContent()
	if content == nil {
		return ""
	}
	return content.GetTitle()
}


func (p Proposal) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	var content Content
	return unpacker.UnpackAny(p.Content, &content)
}


type Proposals []Proposal

var _ types.UnpackInterfacesMessage = Proposals{}


func (p Proposals) Equal(other Proposals) bool {
	if len(p) != len(other) {
		return false
	}

	for i, proposal := range p {
		if !proposal.Equal(other[i]) {
			return false
		}
	}

	return true
}


func (p Proposals) String() string {
	out := "ID - (Status) [Type] Title\n"
	for _, prop := range p {
		out += fmt.Sprintf("%d - (%s) [%s] %s\n",
			prop.ProposalId, prop.Status,
			prop.ProposalType(), prop.GetTitle())
	}
	return strings.TrimSpace(out)
}


func (p Proposals) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, x := range p {
		err := x.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}

type (

	ProposalQueue []uint64
)


func ProposalStatusFromString(str string) (ProposalStatus, error) {
	num, ok := ProposalStatus_value[str]
	if !ok {
		return StatusNil, fmt.Errorf("'%s' is not a valid proposal status", str)
	}
	return ProposalStatus(num), nil
}



func ValidProposalStatus(status ProposalStatus) bool {
	if status == StatusDepositPeriod ||
		status == StatusVotingPeriod ||
		status == StatusPassed ||
		status == StatusRejected ||
		status == StatusFailed {
		return true
	}
	return false
}


func (status ProposalStatus) Marshal() ([]byte, error) {
	return []byte{byte(status)}, nil
}


func (status *ProposalStatus) Unmarshal(data []byte) error {
	*status = ProposalStatus(data[0])
	return nil
}



func (status ProposalStatus) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(status.String()))
	default:

		s.Write([]byte(fmt.Sprintf("%v", byte(status))))
	}
}


const (
	ProposalTypeText string = "Text"
)


var _ Content = &TextProposal{}


func NewTextProposal(title, description string) Content {
	return &TextProposal{title, description}
}


func (tp *TextProposal) GetTitle() string { return tp.Title }


func (tp *TextProposal) GetDescription() string { return tp.Description }


func (tp *TextProposal) ProposalRoute() string { return RouterKey }


func (tp *TextProposal) ProposalType() string { return ProposalTypeText }


func (tp *TextProposal) ValidateBasic() error { return ValidateAbstract(tp) }


func (tp TextProposal) String() string {
	out, _ := yaml.Marshal(tp)
	return string(out)
}

var validProposalTypes = map[string]struct{}{
	ProposalTypeText: {},
}



func RegisterProposalType(ty string) {
	if _, ok := validProposalTypes[ty]; ok {
		panic(fmt.Sprintf("already registered proposal type: %s", ty))
	}

	validProposalTypes[ty] = struct{}{}
}


func ContentFromProposalType(title, desc, ty string) Content {
	switch ty {
	case ProposalTypeText:
		return NewTextProposal(title, desc)

	default:
		return nil
	}
}



//

func IsValidProposalType(ty string) bool {
	_, ok := validProposalTypes[ty]
	return ok
}





func ProposalHandler(_ sdk.Context, c Content) error {
	switch c.ProposalType() {
	case ProposalTypeText:

		return nil

	default:
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "unrecognized gov proposal type: %s", c.ProposalType())
	}
}
