package types

import (
	"fmt"
	"time"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/crypto/tmhash"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	"gopkg.in/yaml.v2"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/evidence/exported"
)


const (
	RouteEquivocation = "equivocation"
	TypeEquivocation  = "equivocation"
)

var _ exported.Evidence = &Equivocation{}


func (e *Equivocation) Route() string { return RouteEquivocation }


func (e *Equivocation) Type() string { return TypeEquivocation }

func (e *Equivocation) String() string {
	bz, _ := yaml.Marshal(e)
	return string(bz)
}


func (e *Equivocation) Hash() tmbytes.HexBytes {
	bz, err := e.Marshal()
	if err != nil {
		panic(err)
	}
	return tmhash.Sum(bz)
}


func (e *Equivocation) ValidateBasic() error {
	if e.Time.Unix() <= 0 {
		return fmt.Errorf("invalid equivocation time: %s", e.Time)
	}
	if e.Height < 1 {
		return fmt.Errorf("invalid equivocation height: %d", e.Height)
	}
	if e.Power < 1 {
		return fmt.Errorf("invalid equivocation validator power: %d", e.Power)
	}
	if e.ConsensusAddress == "" {
		return fmt.Errorf("invalid equivocation validator consensus address: %s", e.ConsensusAddress)
	}

	return nil
}



func (e Equivocation) GetConsensusAddress() sdk.ConsAddress {
	addr, _ := sdk.ConsAddressFromBech32(e.ConsensusAddress)
	return addr
}


func (e Equivocation) GetHeight() int64 {
	return e.Height
}


func (e Equivocation) GetTime() time.Time {
	return e.Time
}



func (e Equivocation) GetValidatorPower() int64 {
	return e.Power
}


func (e Equivocation) GetTotalPower() int64 { return 0 }



func FromABCIEvidence(e abci.Evidence) exported.Evidence {
	bech32PrefixConsAddr := sdk.GetConfig().GetBech32ConsensusAddrPrefix()
	consAddr, err := sdk.Bech32ifyAddressBytes(bech32PrefixConsAddr, e.Validator.Address)
	if err != nil {
		panic(err)
	}

	return &Equivocation{
		Height:           e.Height,
		Power:            e.Validator.Power,
		ConsensusAddress: consAddr,
		Time:             e.Time,
	}
}
