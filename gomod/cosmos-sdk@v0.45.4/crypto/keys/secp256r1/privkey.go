package secp256r1

import (
	"github.com/cosmos/cosmos-sdk/crypto/keys/internal/ecdsa"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)


func GenPrivKey() (*PrivKey, error) {
	key, err := ecdsa.GenPrivKey(secp256r1)
	return &PrivKey{&ecdsaSK{key}}, err
}


func (m *PrivKey) PubKey() cryptotypes.PubKey {
	return &PubKey{&ecdsaPK{m.Secret.PubKey()}}
}


func (m *PrivKey) String() string {
	return m.Secret.String(name)
}


func (m *PrivKey) Type() string {
	return name
}


func (m *PrivKey) Sign(msg []byte) ([]byte, error) {
	return m.Secret.Sign(msg)
}


func (m *PrivKey) Bytes() []byte {
	if m == nil {
		return nil
	}
	return m.Secret.Bytes()
}


func (m *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	sk2, ok := other.(*PrivKey)
	if !ok {
		return false
	}
	return m.Secret.Equal(&sk2.Secret.PrivateKey)
}

type ecdsaSK struct {
	ecdsa.PrivKey
}


func (sk *ecdsaSK) Size() int {
	if sk == nil {
		return 0
	}
	return fieldSize
}


func (sk *ecdsaSK) Unmarshal(bz []byte) error {
	return sk.PrivKey.Unmarshal(bz, secp256r1, fieldSize)
}
