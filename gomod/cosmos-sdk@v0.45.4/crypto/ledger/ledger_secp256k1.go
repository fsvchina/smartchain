package ledger

import (
	"fmt"
	"os"

	"github.com/btcsuite/btcd/btcec"
	"github.com/pkg/errors"

	tmbtcec "github.com/tendermint/btcd/btcec"

	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	"github.com/cosmos/cosmos-sdk/crypto/types"
)

var (


	discoverLedger discoverLedgerFn
)

type (



	discoverLedgerFn func() (SECP256K1, error)


	SECP256K1 interface {
		Close() error

		GetPublicKeySECP256K1([]uint32) ([]byte, error)

		GetAddressPubKeySECP256K1([]uint32, string) ([]byte, string, error)

		SignSECP256K1([]uint32, []byte) ([]byte, error)
	}



	PrivKeyLedgerSecp256k1 struct {



		CachedPubKey types.PubKey
		Path         hd.BIP44Params
	}
)


//



func NewPrivKeySecp256k1Unsafe(path hd.BIP44Params) (types.LedgerPrivKey, error) {
	device, err := getDevice()
	if err != nil {
		return nil, err
	}
	defer warnIfErrors(device.Close)

	pubKey, err := getPubKeyUnsafe(device, path)
	if err != nil {
		return nil, err
	}

	return PrivKeyLedgerSecp256k1{pubKey, path}, nil
}



func NewPrivKeySecp256k1(path hd.BIP44Params, hrp string) (types.LedgerPrivKey, string, error) {
	device, err := getDevice()
	if err != nil {
		return nil, "", fmt.Errorf("failed to retrieve device: %w", err)
	}
	defer warnIfErrors(device.Close)

	pubKey, addr, err := getPubKeyAddrSafe(device, path, hrp)
	if err != nil {
		return nil, "", fmt.Errorf("failed to recover pubkey: %w", err)
	}

	return PrivKeyLedgerSecp256k1{pubKey, path}, addr, nil
}


func (pkl PrivKeyLedgerSecp256k1) PubKey() types.PubKey {
	return pkl.CachedPubKey
}


func (pkl PrivKeyLedgerSecp256k1) Sign(message []byte) ([]byte, error) {
	device, err := getDevice()
	if err != nil {
		return nil, err
	}
	defer warnIfErrors(device.Close)

	return sign(device, pkl, message)
}


func ShowAddress(path hd.BIP44Params, expectedPubKey types.PubKey,
	accountAddressPrefix string) error {
	device, err := getDevice()
	if err != nil {
		return err
	}
	defer warnIfErrors(device.Close)

	pubKey, err := getPubKeyUnsafe(device, path)
	if err != nil {
		return err
	}

	if !pubKey.Equals(expectedPubKey) {
		return fmt.Errorf("the key's pubkey does not match with the one retrieved from Ledger. Check that the HD path and device are the correct ones")
	}

	pubKey2, _, err := getPubKeyAddrSafe(device, path, accountAddressPrefix)
	if err != nil {
		return err
	}

	if !pubKey2.Equals(expectedPubKey) {
		return fmt.Errorf("the key's pubkey does not match with the one retrieved from Ledger. Check that the HD path and device are the correct ones")
	}

	return nil
}



func (pkl PrivKeyLedgerSecp256k1) ValidateKey() error {
	device, err := getDevice()
	if err != nil {
		return err
	}
	defer warnIfErrors(device.Close)

	return validateKey(device, pkl)
}


func (pkl *PrivKeyLedgerSecp256k1) AssertIsPrivKeyInner() {}



func (pkl PrivKeyLedgerSecp256k1) Bytes() []byte {
	return cdc.MustMarshal(pkl)
}



func (pkl PrivKeyLedgerSecp256k1) Equals(other types.LedgerPrivKey) bool {
	if otherKey, ok := other.(PrivKeyLedgerSecp256k1); ok {
		return pkl.CachedPubKey.Equals(otherKey.CachedPubKey)
	}
	return false
}

func (pkl PrivKeyLedgerSecp256k1) Type() string { return "PrivKeyLedgerSecp256k1" }



func warnIfErrors(f func() error) {
	if err := f(); err != nil {
		_, _ = fmt.Fprint(os.Stderr, "received error when closing ledger connection", err)
	}
}

func convertDERtoBER(signatureDER []byte) ([]byte, error) {
	sigDER, err := btcec.ParseDERSignature(signatureDER, btcec.S256())
	if err != nil {
		return nil, err
	}
	sigBER := tmbtcec.Signature{R: sigDER.R, S: sigDER.S}
	return sigBER.Serialize(), nil
}

func getDevice() (SECP256K1, error) {
	if discoverLedger == nil {
		return nil, errors.New("no Ledger discovery function defined")
	}

	device, err := discoverLedger()
	if err != nil {
		return nil, errors.Wrap(err, "ledger nano S")
	}

	return device, nil
}

func validateKey(device SECP256K1, pkl PrivKeyLedgerSecp256k1) error {
	pub, err := getPubKeyUnsafe(device, pkl.Path)
	if err != nil {
		return err
	}


	if !pub.Equals(pkl.CachedPubKey) {
		return fmt.Errorf("cached key does not match retrieved key")
	}

	return nil
}


//



func sign(device SECP256K1, pkl PrivKeyLedgerSecp256k1, msg []byte) ([]byte, error) {
	err := validateKey(device, pkl)
	if err != nil {
		return nil, err
	}

	sig, err := device.SignSECP256K1(pkl.Path.DerivationPath(), msg)
	if err != nil {
		return nil, err
	}

	return convertDERtoBER(sig)
}


//



//


func getPubKeyUnsafe(device SECP256K1, path hd.BIP44Params) (types.PubKey, error) {
	publicKey, err := device.GetPublicKeySECP256K1(path.DerivationPath())
	if err != nil {
		return nil, fmt.Errorf("please open Cosmos app on the Ledger device - error: %v", err)
	}


	cmp, err := btcec.ParsePubKey(publicKey, btcec.S256())
	if err != nil {
		return nil, fmt.Errorf("error parsing public key: %v", err)
	}

	compressedPublicKey := make([]byte, secp256k1.PubKeySize)
	copy(compressedPublicKey, cmp.SerializeCompressed())

	return &secp256k1.PubKey{Key: compressedPublicKey}, nil
}




//


func getPubKeyAddrSafe(device SECP256K1, path hd.BIP44Params, hrp string) (types.PubKey, string, error) {
	publicKey, addr, err := device.GetAddressPubKeySECP256K1(path.DerivationPath(), hrp)
	if err != nil {
		return nil, "", fmt.Errorf("%w: address rejected for path %s", err, path.String())
	}


	cmp, err := btcec.ParsePubKey(publicKey, btcec.S256())
	if err != nil {
		return nil, "", fmt.Errorf("error parsing public key: %v", err)
	}

	compressedPublicKey := make([]byte, secp256k1.PubKeySize)
	copy(compressedPublicKey, cmp.SerializeCompressed())

	return &secp256k1.PubKey{Key: compressedPublicKey}, addr, nil
}
