package server

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	sdk "github.com/cosmos/cosmos-sdk/types"
)



func GenerateCoinKey(algo keyring.SignatureAlgo) (sdk.AccAddress, string, error) {

	info, secret, err := keyring.NewInMemory().NewMnemonic(
		"name",
		keyring.English,
		sdk.GetConfig().GetFullBIP44Path(),
		keyring.DefaultBIP39Passphrase,
		algo,
	)
	if err != nil {
		return sdk.AccAddress{}, "", err
	}

	return sdk.AccAddress(info.GetPubKey().Address()), secret, nil
}






func GenerateSaveCoinKey(
	keybase keyring.Keyring,
	keyName string,
	overwrite bool,
	algo keyring.SignatureAlgo,
) (sdk.AccAddress, string, error) {
	exists := false
	_, err := keybase.Key(keyName)
	if err == nil {
		exists = true
	}


	if !overwrite && exists {
		return sdk.AccAddress{}, "", fmt.Errorf("key already exists, overwrite is disabled")
	}


	if exists {
		if err := keybase.Delete(keyName); err != nil {
			return sdk.AccAddress{}, "", fmt.Errorf("failed to overwrite key")
		}
	}

	k, mnemonic, err := keybase.NewMnemonic(keyName, keyring.English, sdk.GetConfig().GetFullBIP44Path(), keyring.DefaultBIP39Passphrase, algo)
	if err != nil {
		return sdk.AccAddress{}, "", err
	}

	return sdk.AccAddress(k.GetAddress()), mnemonic, nil
}
