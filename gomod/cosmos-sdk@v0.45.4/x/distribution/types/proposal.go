package types

import (
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (

	ProposalTypeCommunityPoolSpend = "CommunityPoolSpend"
)


var _ govtypes.Content = &CommunityPoolSpendProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeCommunityPoolSpend)
	govtypes.RegisterProposalTypeCodec(&CommunityPoolSpendProposal{}, "cosmos-sdk/CommunityPoolSpendProposal")
}



func NewCommunityPoolSpendProposal(title, description string, recipient sdk.AccAddress, amount sdk.Coins) *CommunityPoolSpendProposal {
	return &CommunityPoolSpendProposal{title, description, recipient.String(), amount}
}


func (csp *CommunityPoolSpendProposal) GetTitle() string { return csp.Title }


func (csp *CommunityPoolSpendProposal) GetDescription() string { return csp.Description }


func (csp *CommunityPoolSpendProposal) ProposalRoute() string { return RouterKey }


func (csp *CommunityPoolSpendProposal) ProposalType() string { return ProposalTypeCommunityPoolSpend }


func (csp *CommunityPoolSpendProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(csp)
	if err != nil {
		return err
	}
	if !csp.Amount.IsValid() {
		return ErrInvalidProposalAmount
	}
	if csp.Recipient == "" {
		return ErrEmptyProposalRecipient
	}

	return nil
}


func (csp CommunityPoolSpendProposal) String() string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf(`Community Pool Spend Proposal:
  Title:       %s
  Description: %s
  Recipient:   %s
  Amount:      %s
`, csp.Title, csp.Description, csp.Recipient, csp.Amount))
	return b.String()
}
