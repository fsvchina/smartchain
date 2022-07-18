package ecdsa

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)





var p256Order = elliptic.P256().Params().N




var p256HalfOrder = new(big.Int).Rsh(p256Order, 1)



func IsSNormalized(sigS *big.Int) bool {
	return sigS.Cmp(p256HalfOrder) != 1
}



func NormalizeS(sigS *big.Int) *big.Int {

	if IsSNormalized(sigS) {
		return sigS
	}

	return new(big.Int).Sub(p256Order, sigS)
}




func signatureRaw(r *big.Int, s *big.Int) []byte {

	rBytes := r.Bytes()
	sBytes := s.Bytes()
	sigBytes := make([]byte, 64)

	copy(sigBytes[32-len(rBytes):32], rBytes)
	copy(sigBytes[64-len(sBytes):64], sBytes)
	return sigBytes
}



func GenPrivKey(curve elliptic.Curve) (PrivKey, error) {
	key, err := ecdsa.GenerateKey(curve, rand.Reader)
	if err != nil {
		return PrivKey{}, err
	}
	return PrivKey{*key}, nil
}

type PrivKey struct {
	ecdsa.PrivateKey
}


func (sk *PrivKey) PubKey() PubKey {
	return PubKey{sk.PublicKey, nil}
}


func (sk *PrivKey) Bytes() []byte {
	if sk == nil {
		return nil
	}
	fieldSize := (sk.Curve.Params().BitSize + 7) / 8
	bz := make([]byte, fieldSize)
	sk.D.FillBytes(bz)
	return bz
}









func (sk *PrivKey) Sign(msg []byte) ([]byte, error) {

	digest := sha256.Sum256(msg)
	r, s, err := ecdsa.Sign(rand.Reader, &sk.PrivateKey, digest[:])

	if err != nil {
		return nil, err
	}

	normS := NormalizeS(s)
	return signatureRaw(r, normS), nil
}


func (sk *PrivKey) String(name string) string {
	return name + "{-}"
}


func (sk *PrivKey) MarshalTo(dAtA []byte) (int, error) {
	bz := sk.Bytes()
	copy(dAtA, bz)
	return len(bz), nil
}


func (sk *PrivKey) Unmarshal(bz []byte, curve elliptic.Curve, expectedSize int) error {
	if len(bz) != expectedSize {
		return fmt.Errorf("wrong ECDSA SK bytes, expecting %d bytes", expectedSize)
	}

	sk.Curve = curve
	sk.D = new(big.Int).SetBytes(bz)
	sk.X, sk.Y = curve.ScalarBaseMult(bz)
	return nil
}
