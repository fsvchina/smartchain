


package v036

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/codec"
	v036gov "github.com/cosmos/cosmos-sdk/x/gov/legacy/v036"
)

const (

	ModuleName = "params"


	RouterKey = "params"
)

const (

	ProposalTypeChange = "ParameterChange"
)


const (
	DefaultCodespace = "params"

	CodeUnknownSubspace  = 1
	CodeSettingParameter = 2
	CodeEmptyData        = 3
)


var _ v036gov.Content = ParameterChangeProposal{}



type ParameterChangeProposal struct {
	Title       string        `json:"title" yaml:"title"`
	Description string        `json:"description" yaml:"description"`
	Changes     []ParamChange `json:"changes" yaml:"changes"`
}

func NewParameterChangeProposal(title, description string, changes []ParamChange) ParameterChangeProposal {
	return ParameterChangeProposal{title, description, changes}
}


func (pcp ParameterChangeProposal) GetTitle() string { return pcp.Title }


func (pcp ParameterChangeProposal) GetDescription() string { return pcp.Description }


func (pcp ParameterChangeProposal) ProposalRoute() string { return RouterKey }


func (pcp ParameterChangeProposal) ProposalType() string { return ProposalTypeChange }


func (pcp ParameterChangeProposal) ValidateBasic() error {
	err := v036gov.ValidateAbstract(pcp)
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
      Subkey:   %X
      Value:    %X
`, pc.Subspace, pc.Key, pc.Subkey, pc.Value))
	}

	return b.String()
}


type ParamChange struct {
	Subspace string `json:"subspace" yaml:"subspace"`
	Key      string `json:"key" yaml:"key"`
	Subkey   string `json:"subkey,omitempty" yaml:"subkey,omitempty"`
	Value    string `json:"value" yaml:"value"`
}

func NewParamChange(subspace, key, value string) ParamChange {
	return ParamChange{subspace, key, "", value}
}

func NewParamChangeWithSubkey(subspace, key, subkey, value string) ParamChange {
	return ParamChange{subspace, key, subkey, value}
}


func (pc ParamChange) String() string {
	return fmt.Sprintf(`Param Change:
  Subspace: %s
  Key:      %s
  Subkey:   %X
  Value:    %X
`, pc.Subspace, pc.Key, pc.Subkey, pc.Value)
}



func ValidateChanges(changes []ParamChange) error {
	if len(changes) == 0 {
		return ErrEmptyChanges(DefaultCodespace)
	}

	for _, pc := range changes {
		if len(pc.Subspace) == 0 {
			return ErrEmptySubspace(DefaultCodespace)
		}
		if len(pc.Key) == 0 {
			return ErrEmptyKey(DefaultCodespace)
		}
		if len(pc.Value) == 0 {
			return ErrEmptyValue(DefaultCodespace)
		}
	}

	return nil
}


func ErrUnknownSubspace(codespace string, space string) error {
	return fmt.Errorf("unknown subspace %s", space)
}


func ErrSettingParameter(codespace string, key, subkey, value, msg string) error {
	return fmt.Errorf("error setting parameter %s on %s (%s): %s", value, key, subkey, msg)
}


func ErrEmptyChanges(codespace string) error {
	return fmt.Errorf("submitted parameter changes are empty")
}


func ErrEmptySubspace(codespace string) error {
	return fmt.Errorf("parameter subspace is empty")
}


func ErrEmptyKey(codespace string) error {
	return fmt.Errorf("parameter key is empty")
}


func ErrEmptyValue(codespace string) error {
	return fmt.Errorf("parameter value is empty")
}

func RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	cdc.RegisterConcrete(ParameterChangeProposal{}, "cosmos-sdk/ParameterChangeProposal", nil)
}
