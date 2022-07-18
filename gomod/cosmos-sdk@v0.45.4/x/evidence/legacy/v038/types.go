


package v038

import (
	"fmt"
	"time"

	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"gopkg.in/yaml.v2"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


const (
	ModuleName            = "evidence"
	DefaultParamspace     = ModuleName
	DefaultMaxEvidenceAge = 60 * 2 * time.Second
)


const (
	RouteEquivocation = "equivocation"
	TypeEquivocation  = "equivocation"
)

var (
	amino = codec.NewLegacyAmino()




	//


	ModuleCdc = codec.NewAminoCodec(amino)
)



type Evidence interface {
	Route() string
	Type() string
	String() string
	Hash() tmbytes.HexBytes
	ValidateBasic() error


	GetHeight() int64
}


type Params struct {
	MaxEvidenceAge time.Duration `json:"max_evidence_age" yaml:"max_evidence_age"`
}


type GenesisState struct {
	Params   Params     `json:"params" yaml:"params"`
	Evidence []Evidence `json:"evidence" yaml:"evidence"`
}


var _ Evidence = Equivocation{}



type Equivocation struct {
	Height           int64           `json:"height" yaml:"height"`
	Time             time.Time       `json:"time" yaml:"time"`
	Power            int64           `json:"power" yaml:"power"`
	ConsensusAddress sdk.ConsAddress `json:"consensus_address" yaml:"consensus_address"`
}


func (e Equivocation) Route() string { return RouteEquivocation }


func (e Equivocation) Type() string { return TypeEquivocation }

func (e Equivocation) String() string {
	bz, _ := yaml.Marshal(e)
	return string(bz)
}


func (e Equivocation) Hash() tmbytes.HexBytes {
	return tmhash.Sum(ModuleCdc.LegacyAmino.MustMarshal(e))
}


func (e Equivocation) ValidateBasic() error {
	if e.Time.Unix() <= 0 {
		return fmt.Errorf("invalid equivocation time: %s", e.Time)
	}
	if e.Height < 1 {
		return fmt.Errorf("invalid equivocation height: %d", e.Height)
	}
	if e.Power < 1 {
		return fmt.Errorf("invalid equivocation validator power: %d", e.Power)
	}
	if e.ConsensusAddress.Empty() {
		return fmt.Errorf("invalid equivocation validator consensus address: %s", e.ConsensusAddress)
	}

	return nil
}


func (e Equivocation) GetHeight() int64 {
	return e.Height
}
