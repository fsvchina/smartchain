package keyring

import (
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
)




type Language int

const (


	English Language = iota + 1

	Japanese

	Korean

	Spanish

	ChineseSimplified

	ChineseTraditional

	French

	Italian
)

const (

	DefaultBIP39Passphrase = ""


	defaultEntropySize = 256
	addressSuffix      = "address"
	infoSuffix         = "info"
)


type KeyType uint


const (
	TypeLocal   KeyType = 0
	TypeLedger  KeyType = 1
	TypeOffline KeyType = 2
	TypeMulti   KeyType = 3
)

var keyTypes = map[KeyType]string{
	TypeLocal:   "local",
	TypeLedger:  "ledger",
	TypeOffline: "offline",
	TypeMulti:   "multi",
}


func (kt KeyType) String() string {
	return keyTypes[kt]
}

type (

	DeriveKeyFunc func(mnemonic string, bip39Passphrase, hdPath string, algo hd.PubKeyType) ([]byte, error)

	PrivKeyGenFunc func(bz []byte, algo hd.PubKeyType) (cryptotypes.PrivKey, error)
)
