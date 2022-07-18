package types

const (

	ModuleName = "feemarket"



	StoreKey = ModuleName


	RouterKey = ModuleName
)


const (
	prefixBlockGasUsed      = iota + 1
	deprecatedPrefixBaseFee
)


var (
	KeyPrefixBlockGasUsed = []byte{prefixBlockGasUsed}
)
