package proposal

import (
	"fmt"
	"strings"

	yaml "gopkg.in/yaml.v2"

	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
)

const (

	ProposalTypeChange = "ParameterChange"
)


var _ govtypes.Content = &ParameterChangeProposal{}

func init() {
	govtypes.RegisterProposalType(ProposalTypeChange)
	govtypes.RegisterProposalTypeCodec(&ParameterChangeProposal{}, "cosmos-sdk/ParameterChangeProposal")
}

func NewParameterChangeProposal(title, description string, changes []ParamChange) *ParameterChangeProposal {
	return &ParameterChangeProposal{title, description, changes}
}


func (pcp *ParameterChangeProposal) GetTitle() string { return pcp.Title }


func (pcp *ParameterChangeProposal) GetDescription() string { return pcp.Description }


func (pcp *ParameterChangeProposal) ProposalRoute() string { return RouterKey }


func (pcp *ParameterChangeProposal) ProposalType() string { return ProposalTypeChange }


func (pcp *ParameterChangeProposal) ValidateBasic() error {
	err := govtypes.ValidateAbstract(pcp)
	if err != nil {
		return err
	}

	return ValidateChanges(pcp.Changes)
}


func (pcp ParameterChangeProposal) String() string {
	var b strings.Builder

	b.WriteString(fmt.Sprintf(`Parameter Change Proposal:
  Title:       %s
  Description: %s
  Changes:
`, pcp.Title, pcp.Description))

	for _, pc := range pcp.Changes {
		b.WriteString(fmt.Sprintf(`    Param Change:
      Subspace: %s
      Key:      %s
      Value:    %X
`, pc.Subspace, pc.Key, pc.Value))
	}

	return b.String()
}

func NewParamChange(subspace, key, value string) ParamChange {
	return ParamChange{subspace, key, value}
}


func (pc ParamChange) String() string {
	out, _ := yaml.Marshal(pc)
	return string(out)
}



func ValidateChanges(changes []ParamChange) error {
	if len(changes) == 0 {
		return ErrEmptyChanges
	}

	for _, pc := range changes {
		if len(pc.Subspace) == 0 {
			return ErrEmptySubspace
		}
		if len(pc.Key) == 0 {
			return ErrEmptyKey
		}
		if len(pc.Value) == 0 {
			return ErrEmptyValue
		}
	}

	return nil
}
