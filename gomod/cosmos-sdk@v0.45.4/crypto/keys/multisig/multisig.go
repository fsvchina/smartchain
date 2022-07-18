package multisig

import (
	fmt "fmt"

	tmcrypto "github.com/tendermint/tendermint/crypto"

	"github.com/cosmos/cosmos-sdk/codec/types"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	multisigtypes "github.com/cosmos/cosmos-sdk/crypto/types/multisig"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)

var _ multisigtypes.PubKey = &LegacyAminoPubKey{}
var _ types.UnpackInterfacesMessage = &LegacyAminoPubKey{}





func NewLegacyAminoPubKey(threshold int, pubKeys []cryptotypes.PubKey) *LegacyAminoPubKey {
	if threshold <= 0 {
		panic("threshold k of n multisignature: k <= 0")
	}
	if len(pubKeys) < threshold {
		panic("threshold k of n multisignature: len(pubKeys) < k")
	}
	anyPubKeys, err := packPubKeys(pubKeys)
	if err != nil {
		panic(err)
	}
	return &LegacyAminoPubKey{Threshold: uint32(threshold), PubKeys: anyPubKeys}
}


func (m *LegacyAminoPubKey) Address() cryptotypes.Address {
	return tmcrypto.AddressHash(m.Bytes())
}


func (m *LegacyAminoPubKey) Bytes() []byte {
	return AminoCdc.MustMarshal(m)
}






func (m *LegacyAminoPubKey) VerifyMultisignature(getSignBytes multisigtypes.GetSignBytesFunc, sig *signing.MultiSignatureData) error {
	bitarray := sig.BitArray
	sigs := sig.Signatures
	size := bitarray.Count()
	pubKeys := m.GetPubKeys()

	if len(pubKeys) != size {
		return fmt.Errorf("bit array size is incorrect, expecting: %d", len(pubKeys))
	}

	if len(sigs) < int(m.Threshold) || len(sigs) > size {
		return fmt.Errorf("signature size is incorrect %d", len(sigs))
	}

	if bitarray.NumTrueBitsBefore(size) < int(m.Threshold) {
		return fmt.Errorf("not enough signatures set, have %d, expected %d", bitarray.NumTrueBitsBefore(size), int(m.Threshold))
	}

	sigIndex := 0
	for i := 0; i < size; i++ {
		if bitarray.GetIndex(i) {
			si := sig.Signatures[sigIndex]
			switch si := si.(type) {
			case *signing.SingleSignatureData:
				msg, err := getSignBytes(si.SignMode)
				if err != nil {
					return err
				}
				if !pubKeys[i].VerifySignature(msg, si.Signature) {
					return fmt.Errorf("unable to verify signature at index %d", i)
				}
			case *signing.MultiSignatureData:
				nestedMultisigPk, ok := pubKeys[i].(multisigtypes.PubKey)
				if !ok {
					return fmt.Errorf("unable to parse pubkey of index %d", i)
				}
				if err := nestedMultisigPk.VerifyMultisignature(getSignBytes, si); err != nil {
					return err
				}
			default:
				return fmt.Errorf("improper signature data type for index %d", sigIndex)
			}
			sigIndex++
		}
	}
	return nil
}




func (m *LegacyAminoPubKey) VerifySignature(msg []byte, sig []byte) bool {
	panic("not implemented")
}


func (m *LegacyAminoPubKey) GetPubKeys() []cryptotypes.PubKey {
	if m != nil {
		pubKeys := make([]cryptotypes.PubKey, len(m.PubKeys))
		for i := 0; i < len(m.PubKeys); i++ {
			pubKeys[i] = m.PubKeys[i].GetCachedValue().(cryptotypes.PubKey)
		}
		return pubKeys
	}

	return nil
}



func (m *LegacyAminoPubKey) Equals(key cryptotypes.PubKey) bool {
	otherKey, ok := key.(multisigtypes.PubKey)
	if !ok {
		return false
	}
	pubKeys := m.GetPubKeys()
	otherPubKeys := otherKey.GetPubKeys()
	if m.GetThreshold() != otherKey.GetThreshold() || len(pubKeys) != len(otherPubKeys) {
		return false
	}

	for i := 0; i < len(pubKeys); i++ {
		if !pubKeys[i].Equals(otherPubKeys[i]) {
			return false
		}
	}
	return true
}


func (m *LegacyAminoPubKey) GetThreshold() uint {
	return uint(m.Threshold)
}


func (m *LegacyAminoPubKey) Type() string {
	return "PubKeyMultisigThreshold"
}


func (m *LegacyAminoPubKey) UnpackInterfaces(unpacker types.AnyUnpacker) error {
	for _, any := range m.PubKeys {
		var pk cryptotypes.PubKey
		err := unpacker.UnpackAny(any, &pk)
		if err != nil {
			return err
		}
	}
	return nil
}

func packPubKeys(pubKeys []cryptotypes.PubKey) ([]*types.Any, error) {
	anyPubKeys := make([]*types.Any, len(pubKeys))

	for i := 0; i < len(pubKeys); i++ {
		any, err := types.NewAnyWithValue(pubKeys[i])
		if err != nil {
			return nil, err
		}
		anyPubKeys[i] = any
	}
	return anyPubKeys, nil
}
