package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/sha256"
	"fmt"
	"math/big"

	tmcrypto "github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/errors"
)




func signatureFromBytes(sigStr []byte) *signature {
	return &signature{
		R: new(big.Int).SetBytes(sigStr[:32]),
		S: new(big.Int).SetBytes(sigStr[32:64]),
	}
}


type signature struct {
	R, S *big.Int
}

type PubKey struct {
	ecdsa.PublicKey


	address tmcrypto.Address
}




func (pk *PubKey) Address(protoName string) tmcrypto.Address {
	if pk.address == nil {
		pk.address = address.Hash(protoName, pk.Bytes())
	}
	return pk.address
}



func (pk *PubKey) Bytes() []byte {
	if pk == nil {
		return nil
	}
	return elliptic.MarshalCompressed(pk.Curve, pk.X, pk.Y)
}






func (pk *PubKey) VerifySignature(msg []byte, sig []byte) bool {






	if len(sig) != 64 {
		return false
	}

	s := signatureFromBytes(sig)
	if !IsSNormalized(s.S) {
		return false
	}

	h := sha256.Sum256(msg)
	return ecdsa.Verify(&pk.PublicKey, h[:], s.R, s.S)
}


func (pk *PubKey) String(curveName string) string {
	return fmt.Sprintf("%s{%X}", curveName, pk.Bytes())
}




func (pk *PubKey) MarshalTo(dAtA []byte) (int, error) {
	bz := pk.Bytes()
	copy(dAtA, bz)
	return len(bz), nil
}


func (pk *PubKey) Unmarshal(bz []byte, curve elliptic.Curve, expectedSize int) error {
	if len(bz) != expectedSize {
		return errors.Wrapf(errors.ErrInvalidPubKey, "wrong ECDSA PK bytes, expecting %d bytes, got %d", expectedSize, len(bz))
	}
	cpk := ecdsa.PublicKey{Curve: curve}
	cpk.X, cpk.Y = elliptic.UnmarshalCompressed(curve, bz)
	if cpk.X == nil || cpk.Y == nil {
		return errors.Wrapf(errors.ErrInvalidPubKey, "wrong ECDSA PK bytes, unknown curve type: %d", bz[0])
	}
	pk.PublicKey = cpk
	return nil
}
