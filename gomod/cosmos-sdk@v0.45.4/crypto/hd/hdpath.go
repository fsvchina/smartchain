package hd

import (
	"crypto/hmac"
	"crypto/sha512"
	"encoding/binary"
	"fmt"
	"math/big"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/btcsuite/btcd/btcec"
)





type BIP44Params struct {
	Purpose      uint32 `json:"purpose"`
	CoinType     uint32 `json:"coinType"`
	Account      uint32 `json:"account"`
	Change       bool   `json:"change"`
	AddressIndex uint32 `json:"addressIndex"`
}



func NewParams(purpose, coinType, account uint32, change bool, addressIdx uint32) *BIP44Params {
	return &BIP44Params{
		Purpose:      purpose,
		CoinType:     coinType,
		Account:      account,
		Change:       change,
		AddressIndex: addressIdx,
	}
}



func NewParamsFromPath(path string) (*BIP44Params, error) {
	spl := strings.Split(path, "/")


	switch {
	case spl[0] == path:
		return nil, fmt.Errorf("path %s doesn't contain '/' separators", path)

	case strings.TrimSpace(spl[0]) == "":
		return nil, fmt.Errorf("ambiguous path %s: use 'm/' prefix for absolute paths, or no leading '/' for relative ones", path)

	case strings.TrimSpace(spl[0]) == "m":
		spl = spl[1:]
	}

	if len(spl) != 5 {
		return nil, fmt.Errorf("invalid path length %s", path)
	}


	purpose, err := hardenedInt(spl[0])
	if err != nil {
		return nil, fmt.Errorf("invalid HD path purpose %s: %w", spl[0], err)
	}

	coinType, err := hardenedInt(spl[1])
	if err != nil {
		return nil, fmt.Errorf("invalid HD path coin type %s: %w", spl[1], err)
	}

	account, err := hardenedInt(spl[2])
	if err != nil {
		return nil, fmt.Errorf("invalid HD path account %s: %w", spl[2], err)
	}

	change, err := hardenedInt(spl[3])
	if err != nil {
		return nil, fmt.Errorf("invalid HD path change %s: %w", spl[3], err)
	}

	addressIdx, err := hardenedInt(spl[4])
	if err != nil {
		return nil, fmt.Errorf("invalid HD path address index %s: %w", spl[4], err)
	}


	if spl[0] != "44'" {
		return nil, fmt.Errorf("first field in path must be 44', got %s", spl[0])
	}

	if !isHardened(spl[1]) || !isHardened(spl[2]) {
		return nil,
			fmt.Errorf("second and third field in path must be hardened (ie. contain the suffix ', got %s and %s", spl[1], spl[2])
	}

	if isHardened(spl[3]) || isHardened(spl[4]) {
		return nil,
			fmt.Errorf("fourth and fifth field in path must not be hardened (ie. not contain the suffix ', got %s and %s", spl[3], spl[4])
	}

	if !(change == 0 || change == 1) {
		return nil, fmt.Errorf("change field can only be 0 or 1")
	}

	return &BIP44Params{
		Purpose:      purpose,
		CoinType:     coinType,
		Account:      account,
		Change:       change > 0,
		AddressIndex: addressIdx,
	}, nil
}

func hardenedInt(field string) (uint32, error) {
	field = strings.TrimSuffix(field, "'")

	i, err := strconv.ParseUint(field, 10, 32)
	if err != nil {
		return 0, err
	}

	return uint32(i), nil
}

func isHardened(field string) bool {
	return strings.HasSuffix(field, "'")
}




func NewFundraiserParams(account, coinType, addressIdx uint32) *BIP44Params {
	return NewParams(44, coinType, account, false, addressIdx)
}


func (p BIP44Params) DerivationPath() []uint32 {
	change := uint32(0)
	if p.Change {
		change = 1
	}

	return []uint32{
		p.Purpose,
		p.CoinType,
		p.Account,
		change,
		p.AddressIndex,
	}
}



func (p BIP44Params) String() string {
	var changeStr string
	if p.Change {
		changeStr = "1"
	} else {
		changeStr = "0"
	}
	return fmt.Sprintf("m/%d'/%d'/%d'/%s/%d",
		p.Purpose,
		p.CoinType,
		p.Account,
		changeStr,
		p.AddressIndex)
}


func ComputeMastersFromSeed(seed []byte) (secret [32]byte, chainCode [32]byte) {
	curveIdentifier := []byte("Bitcoin seed")
	secret, chainCode = i64(curveIdentifier, seed)

	return
}



func DerivePrivateKeyForPath(privKeyBytes, chainCode [32]byte, path string) ([]byte, error) {


	path = strings.TrimRightFunc(path, func(r rune) bool { return r == filepath.Separator })
	data := privKeyBytes
	parts := strings.Split(path, "/")

	switch {
	case parts[0] == path:
		return nil, fmt.Errorf("path '%s' doesn't contain '/' separators", path)
	case strings.TrimSpace(parts[0]) == "m":
		parts = parts[1:]
	}

	for i, part := range parts {
		if part == "" {
			return nil, fmt.Errorf("path %q with split element #%d is an empty string", part, i)
		}

		harden := part[len(part)-1:] == "'"

		if harden {
			part = part[:len(part)-1]
		}




		idx, err := strconv.ParseUint(part, 10, 31)
		if err != nil {
			return []byte{}, fmt.Errorf("invalid BIP 32 path %s: %w", path, err)
		}

		data, chainCode = derivePrivateKey(data, chainCode, uint32(idx), harden)
	}

	derivedKey := make([]byte, 32)
	n := copy(derivedKey, data[:])

	if n != 32 || len(data) != 32 {
		return []byte{}, fmt.Errorf("expected a key of length 32, got length: %d", len(data))
	}

	return derivedKey, nil
}






func derivePrivateKey(privKeyBytes [32]byte, chainCode [32]byte, index uint32, harden bool) ([32]byte, [32]byte) {
	var data []byte

	if harden {
		index |= 0x80000000

		data = append([]byte{byte(0)}, privKeyBytes[:]...)
	} else {

		_, ecPub := btcec.PrivKeyFromBytes(btcec.S256(), privKeyBytes[:])
		pubkeyBytes := ecPub.SerializeCompressed()
		data = pubkeyBytes

		crypto/secp256k1
		pubkey := secp256k1.PrivKeySecp256k1(privKeyBytes).PubKey()
		public := pubkey.(secp256k1.PubKeySecp256k1)
		data = public[:]
		*/
	}

	data = append(data, uint32ToBytes(index)...)
	data2, chainCode2 := i64(chainCode[:], data)
	x := addScalars(privKeyBytes[:], data2[:])

	return x, chainCode2
}


func addScalars(a []byte, b []byte) [32]byte {
	aInt := new(big.Int).SetBytes(a)
	bInt := new(big.Int).SetBytes(b)
	sInt := new(big.Int).Add(aInt, bInt)
	x := sInt.Mod(sInt, btcec.S256().N).Bytes()
	x2 := [32]byte{}
	copy(x2[32-len(x):], x)

	return x2
}

func uint32ToBytes(i uint32) []byte {
	b := [4]byte{}
	binary.BigEndian.PutUint32(b[:], i)

	return b[:]
}


func i64(key []byte, data []byte) (il [32]byte, ir [32]byte) {
	mac := hmac.New(sha512.New, key)

	_, _ = mac.Write(data)

	I := mac.Sum(nil)
	copy(il[:], I[:32])
	copy(ir[:], I[32:])

	return
}


func CreateHDPath(coinType, account, index uint32) *BIP44Params {
	return NewFundraiserParams(account, coinType, index)
}
