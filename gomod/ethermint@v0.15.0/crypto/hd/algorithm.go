package hd

import (
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	bip39 "github.com/tyler-smith/go-bip39"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"

	"github.com/tharsis/ethermint/crypto/ethsecp256k1"
)

const (

	EthSecp256k1Type = hd.PubKeyType(ethsecp256k1.KeyType)
)

var (



	SupportedAlgorithms = keyring.SigningAlgoList{EthSecp256k1, hd.Secp256k1}



	SupportedAlgorithmsLedger = keyring.SigningAlgoList{EthSecp256k1, hd.Secp256k1}
)



func EthSecp256k1Option() keyring.Option {
	return func(options *keyring.Options) {
		options.SupportedAlgos = SupportedAlgorithms
		options.SupportedAlgosLedger = SupportedAlgorithmsLedger
	}
}

var (
	_ keyring.SignatureAlgo = EthSecp256k1


	EthSecp256k1 = ethSecp256k1Algo{}
)

type ethSecp256k1Algo struct{}


func (s ethSecp256k1Algo) Name() hd.PubKeyType {
	return EthSecp256k1Type
}


func (s ethSecp256k1Algo) Derive() hd.DeriveFn {
	return func(mnemonic, bip39Passphrase, path string) ([]byte, error) {
		hdpath, err := accounts.ParseDerivationPath(path)
		if err != nil {
			return nil, err
		}

		seed, err := bip39.NewSeedWithErrorChecking(mnemonic, bip39Passphrase)
		if err != nil {
			return nil, err
		}


		masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
		if err != nil {
			return nil, err
		}

		key := masterKey
		for _, n := range hdpath {
			key, err = key.Derive(n)
			if err != nil {
				return nil, err
			}
		}


		privateKey, err := key.ECPrivKey()
		if err != nil {
			return nil, err
		}




		privateKeyECDSA := privateKey.ToECDSA()
		derivedKey := crypto.FromECDSA(privateKeyECDSA)

		return derivedKey, nil
	}
}


func (s ethSecp256k1Algo) Generate() hd.GenerateFn {
	return func(bz []byte) cryptotypes.PrivKey {
		bzArr := make([]byte, ethsecp256k1.PrivKeySize)
		copy(bzArr, bz)


		return &ethsecp256k1.PrivKey{
			Key: bzArr,
		}
	}
}
