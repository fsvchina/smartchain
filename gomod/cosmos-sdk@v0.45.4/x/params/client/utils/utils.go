package utils

import (
	"encoding/json"
	"io/ioutil"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/rest"
	"github.com/cosmos/cosmos-sdk/x/params/types/proposal"
)

type (


	ParamChangesJSON []ParamChangeJSON



	ParamChangeJSON struct {
		Subspace string          `json:"subspace" yaml:"subspace"`
		Key      string          `json:"key" yaml:"key"`
		Value    json.RawMessage `json:"value" yaml:"value"`
	}



	ParamChangeProposalJSON struct {
		Title       string           `json:"title" yaml:"title"`
		Description string           `json:"description" yaml:"description"`
		Changes     ParamChangesJSON `json:"changes" yaml:"changes"`
		Deposit     string           `json:"deposit" yaml:"deposit"`
	}


	ParamChangeProposalReq struct {
		BaseReq rest.BaseReq `json:"base_req" yaml:"base_req"`

		Title       string           `json:"title" yaml:"title"`
		Description string           `json:"description" yaml:"description"`
		Changes     ParamChangesJSON `json:"changes" yaml:"changes"`
		Proposer    sdk.AccAddress   `json:"proposer" yaml:"proposer"`
		Deposit     sdk.Coins        `json:"deposit" yaml:"deposit"`
	}
)

func NewParamChangeJSON(subspace, key string, value json.RawMessage) ParamChangeJSON {
	return ParamChangeJSON{subspace, key, value}
}


func (pcj ParamChangeJSON) ToParamChange() proposal.ParamChange {
	return proposal.NewParamChange(pcj.Subspace, pcj.Key, string(pcj.Value))
}



func (pcj ParamChangesJSON) ToParamChanges() []proposal.ParamChange {
	res := make([]proposal.ParamChange, len(pcj))
	for i, pc := range pcj {
		res[i] = pc.ToParamChange()
	}
	return res
}



func ParseParamChangeProposalJSON(cdc *codec.LegacyAmino, proposalFile string) (ParamChangeProposalJSON, error) {
	proposal := ParamChangeProposalJSON{}

	contents, err := ioutil.ReadFile(proposalFile)
	if err != nil {
		return proposal, err
	}

	if err := cdc.UnmarshalJSON(contents, &proposal); err != nil {
		return proposal, err
	}

	return proposal, nil
}
