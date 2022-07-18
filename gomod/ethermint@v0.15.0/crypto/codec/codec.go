package codec

import (
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
)


func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &ethsecp256k1.PubKey{})
}