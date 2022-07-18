package multisig

import (
	"fmt"
	"strings"

	"github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types/tx/signing"
)




type AminoMultisignature struct {
	BitArray *types.CompactBitArray
	Sigs     [][]byte
}


func NewMultisig(n int) *signing.MultiSignatureData {
	return &signing.MultiSignatureData{
		BitArray:   types.NewCompactBitArray(n),
		Signatures: make([]signing.SignatureData, 0, n),
	}
}


func getIndex(pk types.PubKey, keys []types.PubKey) int {
	for i := 0; i < len(keys); i++ {
		if pk.Equals(keys[i]) {
			return i
		}
	}
	return -1
}




func AddSignature(mSig *signing.MultiSignatureData, sig signing.SignatureData, index int) {
	newSigIndex := mSig.BitArray.NumTrueBitsBefore(index)

	if mSig.BitArray.GetIndex(index) {
		mSig.Signatures[newSigIndex] = sig
		return
	}
	mSig.BitArray.SetIndex(index, true)

	if newSigIndex == len(mSig.Signatures) {
		mSig.Signatures = append(mSig.Signatures, sig)
		return
	}


	mSig.Signatures = append(mSig.Signatures, &signing.SingleSignatureData{})
	copy(mSig.Signatures[newSigIndex+1:], mSig.Signatures[newSigIndex:])
	mSig.Signatures[newSigIndex] = sig
}



func AddSignatureFromPubKey(mSig *signing.MultiSignatureData, sig signing.SignatureData, pubkey types.PubKey, keys []types.PubKey) error {
	if mSig == nil {
		return fmt.Errorf("value of mSig is nil %v", mSig)
	}
	if sig == nil {
		return fmt.Errorf("value of sig is nil %v", sig)
	}

	if pubkey == nil || keys == nil {
		return fmt.Errorf("pubkey or keys can't be nil %v %v", pubkey, keys)
	}
	index := getIndex(pubkey, keys)
	if index == -1 {
		keysStr := make([]string, len(keys))
		for i, k := range keys {
			keysStr[i] = fmt.Sprintf("%X", k.Bytes())
		}

		return fmt.Errorf("provided key %X doesn't exist in pubkeys: \n%s", pubkey.Bytes(), strings.Join(keysStr, "\n"))
	}

	AddSignature(mSig, sig, index)
	return nil
}

func AddSignatureV2(mSig *signing.MultiSignatureData, sig signing.SignatureV2, keys []types.PubKey) error {
	return AddSignatureFromPubKey(mSig, sig.Data, sig.PubKey, keys)
}
