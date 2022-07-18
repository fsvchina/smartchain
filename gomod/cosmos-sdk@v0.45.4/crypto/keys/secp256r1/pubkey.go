package secp256r1

import (
	"github.com/gogo/protobuf/proto"
	tmcrypto "github.com/tendermint/tendermint/crypto"

	ecdsa "github.com/cosmos/cosmos-sdk/crypto/keys/internal/ecdsa"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)


func (m *PubKey) String() string {
	return m.Key.String(name)
}


func (m *PubKey) Bytes() []byte {
	if m == nil {
		return nil
	}
	return m.Key.Bytes()
}


func (m *PubKey) Equals(other cryptotypes.PubKey) bool {
	pk2, ok := other.(*PubKey)
	if !ok {
		return false
	}
	return m.Key.Equal(&pk2.Key.PublicKey)
}


func (m *PubKey) Address() tmcrypto.Address {
	return m.Key.Address(proto.MessageName(m))
}


func (m *PubKey) Type() string {
	return name
}


func (m *PubKey) VerifySignature(msg []byte, sig []byte) bool {
	return m.Key.VerifySignature(msg, sig)
}

type ecdsaPK struct {
	ecdsa.PubKey
}


func (pk *ecdsaPK) Size() int {
	if pk == nil {
		return 0
	}
	return pubKeySize
}


func (pk *ecdsaPK) Unmarshal(bz []byte) error {
	return pk.PubKey.Unmarshal(bz, secp256r1, pubKeySize)
}
