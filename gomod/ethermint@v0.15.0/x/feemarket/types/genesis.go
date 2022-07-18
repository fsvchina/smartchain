package types


func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Params:   DefaultParams(),
		BlockGas: 0,
	}
}


func NewGenesisState(params Params, blockGas uint64) *GenesisState {
	return &GenesisState{
		Params:   params,
		BlockGas: blockGas,
	}
}



func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}
