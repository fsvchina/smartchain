package crypto

import (
	"encoding/hex"
	"fmt"

	"github.com/tendermint/crypto/bcrypt"
	"github.com/tendermint/tendermint/crypto"
	"github.com/tendermint/tendermint/crypto/armor"
	"github.com/tendermint/tendermint/crypto/xsalsa20symmetric"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	blockTypePrivKey = "TENDERMINT PRIVATE KEY"
	blockTypeKeyInfo = "TENDERMINT KEY INFO"
	blockTypePubKey  = "TENDERMINT PUBLIC KEY"

	defaultAlgo = "secp256k1"

	headerVersion = "version"
	headerType    = "type"
)














var BcryptSecurityParameter = 12





func ArmorInfoBytes(bz []byte) string {
	header := map[string]string{
		headerType:    "Info",
		headerVersion: "0.0.0",
	}

	return armor.EncodeArmor(blockTypeKeyInfo, header, bz)
}


func ArmorPubKeyBytes(bz []byte, algo string) string {
	header := map[string]string{
		headerVersion: "0.0.1",
	}
	if algo != "" {
		header[headerType] = algo
	}

	return armor.EncodeArmor(blockTypePubKey, header, bz)
}





func UnarmorInfoBytes(armorStr string) ([]byte, error) {
	bz, header, err := unarmorBytes(armorStr, blockTypeKeyInfo)
	if err != nil {
		return nil, err
	}

	if header[headerVersion] != "0.0.0" {
		return nil, fmt.Errorf("unrecognized version: %v", header[headerVersion])
	}

	return bz, nil
}


func UnarmorPubKeyBytes(armorStr string) (bz []byte, algo string, err error) {
	bz, header, err := unarmorBytes(armorStr, blockTypePubKey)
	if err != nil {
		return nil, "", fmt.Errorf("couldn't unarmor bytes: %v", err)
	}

	switch header[headerVersion] {
	case "0.0.0":
		return bz, defaultAlgo, err
	case "0.0.1":
		if header[headerType] == "" {
			header[headerType] = defaultAlgo
		}

		return bz, header[headerType], err
	case "":
		return nil, "", fmt.Errorf("header's version field is empty")
	default:
		err = fmt.Errorf("unrecognized version: %v", header[headerVersion])
		return nil, "", err
	}
}

func unarmorBytes(armorStr, blockType string) (bz []byte, header map[string]string, err error) {
	bType, header, bz, err := armor.DecodeArmor(armorStr)
	if err != nil {
		return
	}

	if bType != blockType {
		err = fmt.Errorf("unrecognized armor type %q, expected: %q", bType, blockType)
		return
	}

	return
}





func EncryptArmorPrivKey(privKey cryptotypes.PrivKey, passphrase string, algo string) string {
	saltBytes, encBytes := encryptPrivKey(privKey, passphrase)
	header := map[string]string{
		"kdf":  "bcrypt",
		"salt": fmt.Sprintf("%X", saltBytes),
	}

	if algo != "" {
		header[headerType] = algo
	}

	armorStr := armor.EncodeArmor(blockTypePrivKey, header, encBytes)

	return armorStr
}




func encryptPrivKey(privKey cryptotypes.PrivKey, passphrase string) (saltBytes []byte, encBytes []byte) {
	saltBytes = crypto.CRandBytes(16)
	key, err := bcrypt.GenerateFromPassword(saltBytes, []byte(passphrase), BcryptSecurityParameter)

	if err != nil {
		panic(sdkerrors.Wrap(err, "error generating bcrypt key from passphrase"))
	}

	key = crypto.Sha256(key)
	privKeyBytes := legacy.Cdc.MustMarshal(privKey)

	return saltBytes, xsalsa20symmetric.EncryptSymmetric(privKeyBytes, key)
}


func UnarmorDecryptPrivKey(armorStr string, passphrase string) (privKey cryptotypes.PrivKey, algo string, err error) {
	blockType, header, encBytes, err := armor.DecodeArmor(armorStr)
	if err != nil {
		return privKey, "", err
	}

	if blockType != blockTypePrivKey {
		return privKey, "", fmt.Errorf("unrecognized armor type: %v", blockType)
	}

	if header["kdf"] != "bcrypt" {
		return privKey, "", fmt.Errorf("unrecognized KDF type: %v", header["kdf"])
	}

	if header["salt"] == "" {
		return privKey, "", fmt.Errorf("missing salt bytes")
	}

	saltBytes, err := hex.DecodeString(header["salt"])
	if err != nil {
		return privKey, "", fmt.Errorf("error decoding salt: %v", err.Error())
	}

	privKey, err = decryptPrivKey(saltBytes, encBytes, passphrase)

	if header[headerType] == "" {
		header[headerType] = defaultAlgo
	}

	return privKey, header[headerType], err
}

func decryptPrivKey(saltBytes []byte, encBytes []byte, passphrase string) (privKey cryptotypes.PrivKey, err error) {
	key, err := bcrypt.GenerateFromPassword(saltBytes, []byte(passphrase), BcryptSecurityParameter)
	if err != nil {
		return privKey, sdkerrors.Wrap(err, "error generating bcrypt key from passphrase")
	}

	key = crypto.Sha256(key)

	privKeyBytes, err := xsalsa20symmetric.DecryptSymmetric(encBytes, key)
	if err != nil && err.Error() == "Ciphertext decryption failed" {
		return privKey, sdkerrors.ErrWrongPassword
	} else if err != nil {
		return privKey, err
	}

	return legacy.PrivKeyFromBytes(privKeyBytes)
}
