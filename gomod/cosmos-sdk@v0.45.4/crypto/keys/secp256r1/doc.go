

package secp256r1

import (
	"crypto/elliptic"
	"fmt"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)

const (

	fieldSize  = 32
	pubKeySize = fieldSize + 1

	name = "secp256r1"
)

var secp256r1 elliptic.Curve

func init() {
	secp256r1 = elliptic.P256()

	expected := (secp256r1.Params().BitSize + 7) / 8
	if expected != fieldSize {
		panic(fmt.Sprintf("Wrong secp256r1 curve fieldSize=%d, expecting=%d", fieldSize, expected))
	}
}


func RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	registry.RegisterImplementations((*cryptotypes.PubKey)(nil), &PubKey{})
}
