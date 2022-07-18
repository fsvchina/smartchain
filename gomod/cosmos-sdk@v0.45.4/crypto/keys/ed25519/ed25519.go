package ed25519

import (
	"crypto/ed25519"
	"crypto/subtle"
	"fmt"
	"io"

	"github.com/hdevalence/ed25519consensus"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/tmhash"

	"github.com/cosmos/cosmos-sdk/codec"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/errors"
)



const (
	PrivKeyName = "tendermint/PrivKeyEd25519"
	PubKeyName  = "tendermint/PubKeyEd25519"

	PubKeySize = 32

	PrivKeySize = 64


	SignatureSize = 64


	SeedSize = 32

	keyType = "ed25519"
)

var _ cryptotypes.PrivKey = &PrivKey{}
var _ codec.AminoMarshaler = &PrivKey{}


func (privKey *PrivKey) Bytes() []byte {
	return privKey.Key
}








func (privKey *PrivKey) Sign(msg []byte) ([]byte, error) {
	return ed25519.Sign(privKey.Key, msg), nil
}


//

func (privKey *PrivKey) PubKey() cryptotypes.PubKey {


	initialized := false
	for _, v := range privKey.Key[32:] {
		if v != 0 {
			initialized = true
			break
		}
	}

	if !initialized {
		panic("Expected ed25519 PrivKey to include concatenated pubkey bytes")
	}

	pubkeyBytes := make([]byte, PubKeySize)
	copy(pubkeyBytes, privKey.Key[32:])
	return &PubKey{Key: pubkeyBytes}
}



func (privKey *PrivKey) Equals(other cryptotypes.LedgerPrivKey) bool {
	if privKey.Type() != other.Type() {
		return false
	}

	return subtle.ConstantTimeCompare(privKey.Bytes(), other.Bytes()) == 1
}

func (privKey *PrivKey) Type() string {
	return keyType
}


func (privKey PrivKey) MarshalAmino() ([]byte, error) {
	return privKey.Key, nil
}


func (privKey *PrivKey) UnmarshalAmino(bz []byte) error {
	if len(bz) != PrivKeySize {
		return fmt.Errorf("invalid privkey size")
	}
	privKey.Key = bz

	return nil
}


func (privKey PrivKey) MarshalAminoJSON() ([]byte, error) {


	return privKey.MarshalAmino()
}


func (privKey *PrivKey) UnmarshalAminoJSON(bz []byte) error {
	return privKey.UnmarshalAmino(bz)
}





func GenPrivKey() *PrivKey {
	return genPrivKey(crypto.CReader())
}


func genPrivKey(rand io.Reader) *PrivKey {
	seed := make([]byte, SeedSize)

	_, err := io.ReadFull(rand, seed)
	if err != nil {
		panic(err)
	}

	return &PrivKey{Key: ed25519.NewKeyFromSeed(seed)}
}






func GenPrivKeyFromSecret(secret []byte) *PrivKey {
	seed := crypto.Sha256(secret)

	return &PrivKey{Key: ed25519.NewKeyFromSeed(seed)}
}



var _ cryptotypes.PubKey = &PubKey{}
var _ codec.AminoMarshaler = &PubKey{}




func (pubKey *PubKey) Address() crypto.Address {
	if len(pubKey.Key) != PubKeySize {
		panic("pubkey is incorrect size")
	}


	return crypto.Address(tmhash.SumTruncated(pubKey.Key))
}


func (pubKey *PubKey) Bytes() []byte {
	return pubKey.Key
}

func (pubKey *PubKey) VerifySignature(msg []byte, sig []byte) bool {

	if len(sig) != SignatureSize {
		return false
	}


	return ed25519consensus.Verify(pubKey.Key, msg, sig)
}


func (pubKey *PubKey) String() string {
	return fmt.Sprintf("PubKeyEd25519{%X}", pubKey.Key)
}

func (pubKey *PubKey) Type() string {
	return keyType
}

func (pubKey *PubKey) Equals(other cryptotypes.PubKey) bool {
	if pubKey.Type() != other.Type() {
		return false
	}

	return subtle.ConstantTimeCompare(pubKey.Bytes(), other.Bytes()) == 1
}


func (pubKey PubKey) MarshalAmino() ([]byte, error) {
	return pubKey.Key, nil
}


func (pubKey *PubKey) UnmarshalAmino(bz []byte) error {
	if len(bz) != PubKeySize {
		return errors.Wrap(errors.ErrInvalidPubKey, "invalid pubkey size")
	}
	pubKey.Key = bz

	return nil
}


func (pubKey PubKey) MarshalAminoJSON() ([]byte, error) {


	return pubKey.MarshalAmino()
}


func (pubKey *PubKey) UnmarshalAminoJSON(bz []byte) error {
	return pubKey.UnmarshalAmino(bz)
}
