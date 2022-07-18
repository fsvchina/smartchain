


package secp256k1

import (
	"github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1/internal/secp256k1"
)


func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	rsv, err := secp256k1.Sign(crypto.Sha256(msg), privKey.Key)
	if err != nil {
		return nil, err
	}

	rs := rsv[:len(rsv)-1]
	return rs, nil
}



func (pubKey *PubKey) VerifySignature(msg []byte, sigStr []byte) bool {
	return secp256k1.VerifySignature(pubKey.Bytes(), crypto.Sha256(msg), sigStr)
}
