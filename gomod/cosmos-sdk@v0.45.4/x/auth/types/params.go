package types

import (
	"fmt"

	yaml "gopkg.in/yaml.v2"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)


const (
	DefaultMaxMemoCharacters      uint64 = 256
	DefaultTxSigLimit             uint64 = 7
	DefaultTxSizeCostPerByte      uint64 = 10
	DefaultSigVerifyCostED25519   uint64 = 590
	DefaultSigVerifyCostSecp256k1 uint64 = 1000
)


var (
	KeyMaxMemoCharacters      = []byte("MaxMemoCharacters")
	KeyTxSigLimit             = []byte("TxSigLimit")
	KeyTxSizeCostPerByte      = []byte("TxSizeCostPerByte")
	KeySigVerifyCostED25519   = []byte("SigVerifyCostED25519")
	KeySigVerifyCostSecp256k1 = []byte("SigVerifyCostSecp256k1")
)

var _ paramtypes.ParamSet = &Params{}


func NewParams(
	maxMemoCharacters, txSigLimit, txSizeCostPerByte, sigVerifyCostED25519, sigVerifyCostSecp256k1 uint64,
) Params {
	return Params{
		MaxMemoCharacters:      maxMemoCharacters,
		TxSigLimit:             txSigLimit,
		TxSizeCostPerByte:      txSizeCostPerByte,
		SigVerifyCostED25519:   sigVerifyCostED25519,
		SigVerifyCostSecp256k1: sigVerifyCostSecp256k1,
	}
}


func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}



func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(KeyMaxMemoCharacters, &p.MaxMemoCharacters, validateMaxMemoCharacters),
		paramtypes.NewParamSetPair(KeyTxSigLimit, &p.TxSigLimit, validateTxSigLimit),
		paramtypes.NewParamSetPair(KeyTxSizeCostPerByte, &p.TxSizeCostPerByte, validateTxSizeCostPerByte),
		paramtypes.NewParamSetPair(KeySigVerifyCostED25519, &p.SigVerifyCostED25519, validateSigVerifyCostED25519),
		paramtypes.NewParamSetPair(KeySigVerifyCostSecp256k1, &p.SigVerifyCostSecp256k1, validateSigVerifyCostSecp256k1),
	}
}


func DefaultParams() Params {
	return Params{
		MaxMemoCharacters:      DefaultMaxMemoCharacters,
		TxSigLimit:             DefaultTxSigLimit,
		TxSizeCostPerByte:      DefaultTxSizeCostPerByte,
		SigVerifyCostED25519:   DefaultSigVerifyCostED25519,
		SigVerifyCostSecp256k1: DefaultSigVerifyCostSecp256k1,
	}
}







func (p Params) SigVerifyCostSecp256r1() uint64 {
	return p.SigVerifyCostSecp256k1 / 2
}


func (p Params) String() string {
	out, _ := yaml.Marshal(p)
	return string(out)
}

func validateTxSigLimit(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid tx signature limit: %d", v)
	}

	return nil
}

func validateSigVerifyCostED25519(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid ED25519 signature verification cost: %d", v)
	}

	return nil
}

func validateSigVerifyCostSecp256k1(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid SECK256k1 signature verification cost: %d", v)
	}

	return nil
}

func validateMaxMemoCharacters(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid max memo characters: %d", v)
	}

	return nil
}

func validateTxSizeCostPerByte(i interface{}) error {
	v, ok := i.(uint64)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	if v == 0 {
		return fmt.Errorf("invalid tx size cost per byte: %d", v)
	}

	return nil
}


func (p Params) Validate() error {
	if err := validateTxSigLimit(p.TxSigLimit); err != nil {
		return err
	}
	if err := validateSigVerifyCostED25519(p.SigVerifyCostED25519); err != nil {
		return err
	}
	if err := validateSigVerifyCostSecp256k1(p.SigVerifyCostSecp256k1); err != nil {
		return err
	}
	if err := validateMaxMemoCharacters(p.MaxMemoCharacters); err != nil {
		return err
	}
	if err := validateTxSizeCostPerByte(p.TxSizeCostPerByte); err != nil {
		return err
	}

	return nil
}
