package authz

import (
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
)


func NewGenesisState(entries []GrantAuthorization) *GenesisState {
	return &GenesisState{
		Authorization: entries,
	}
}


func ValidateGenesis(data GenesisState) error {
	return nil
}


func DefaultGenesisState() *GenesisState {
	return &GenesisState{}
}

var _ cdctypes.UnpackInterfacesMessage = GenesisState{}


func (data GenesisState) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	for _, a := range data.Authorization {
		err := a.UnpackInterfaces(unpacker)
		if err != nil {
			return err
		}
	}
	return nil
}


func (msg GrantAuthorization) UnpackInterfaces(unpacker cdctypes.AnyUnpacker) error {
	var a Authorization
	return unpacker.UnpackAny(msg.Authorization, &a)
}
