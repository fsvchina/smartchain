package keyring

import (
	"fmt"
	"strings"

	"github.com/pkg/errors"
	tmos "github.com/tendermint/tendermint/libs/os"
	dbm "github.com/tendermint/tm-db"

	"github.com/cosmos/cosmos-sdk/crypto"
	"github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)


type LegacyKeybase interface {
	List() ([]Info, error)
	Export(name string) (armor string, err error)
	ExportPrivKey(name, decryptPassphrase, encryptPassphrase string) (armor string, err error)
	ExportPubKey(name string) (armor string, err error)
	Close() error
}


func NewLegacy(name, dir string, opts ...KeybaseOption) (LegacyKeybase, error) {
	if err := tmos.EnsureDir(dir, 0700); err != nil {
		return nil, fmt.Errorf("failed to create Keybase directory: %s", err)
	}

	db, err := sdk.NewLevelDB(name, dir)
	if err != nil {
		return nil, err
	}

	return newDBKeybase(db), nil
}

var _ LegacyKeybase = dbKeybase{}



//

type dbKeybase struct {
	db dbm.DB
}



func newDBKeybase(db dbm.DB) dbKeybase {
	return dbKeybase{
		db: db,
	}
}


func (kb dbKeybase) List() ([]Info, error) {
	var res []Info

	iter, err := kb.db.Iterator(nil, nil)
	if err != nil {
		return nil, err
	}

	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		key := string(iter.Key())


		if strings.HasSuffix(key, infoSuffix) {
			info, err := unmarshalInfo(iter.Value())
			if err != nil {
				return nil, err
			}

			res = append(res, info)
		}
	}

	return res, nil
}


func (kb dbKeybase) Get(name string) (Info, error) {
	bs, err := kb.db.Get(infoKeyBz(name))
	if err != nil {
		return nil, err
	}

	if len(bs) == 0 {
		return nil, sdkerrors.Wrap(sdkerrors.ErrKeyNotFound, name)
	}

	return unmarshalInfo(bs)
}




func (kb dbKeybase) ExportPrivateKeyObject(name string, passphrase string) (types.PrivKey, error) {
	info, err := kb.Get(name)
	if err != nil {
		return nil, err
	}

	var priv types.PrivKey

	switch i := info.(type) {
	case localInfo:
		linfo := i
		if linfo.PrivKeyArmor == "" {
			err = fmt.Errorf("private key not available")
			return nil, err
		}

		priv, _, err = crypto.UnarmorDecryptPrivKey(linfo.PrivKeyArmor, passphrase)
		if err != nil {
			return nil, err
		}

	case ledgerInfo, offlineInfo, multiInfo:
		return nil, errors.New("only works on local private keys")
	}

	return priv, nil
}

func (kb dbKeybase) Export(name string) (armor string, err error) {
	bz, err := kb.db.Get(infoKeyBz(name))
	if err != nil {
		return "", err
	}

	if bz == nil {
		return "", fmt.Errorf("no key to export with name %s", name)
	}

	return crypto.ArmorInfoBytes(bz), nil
}



func (kb dbKeybase) ExportPubKey(name string) (armor string, err error) {
	bz, err := kb.db.Get(infoKeyBz(name))
	if err != nil {
		return "", err
	}

	if bz == nil {
		return "", fmt.Errorf("no key to export with name %s", name)
	}

	info, err := unmarshalInfo(bz)
	if err != nil {
		return
	}

	return crypto.ArmorPubKeyBytes(info.GetPubKey().Bytes(), string(info.GetAlgo())), nil
}




func (kb dbKeybase) ExportPrivKey(name string, decryptPassphrase string,
	encryptPassphrase string) (armor string, err error) {
	priv, err := kb.ExportPrivateKeyObject(name, decryptPassphrase)
	if err != nil {
		return "", err
	}

	info, err := kb.Get(name)
	if err != nil {
		return "", err
	}

	return crypto.EncryptArmorPrivKey(priv, encryptPassphrase, string(info.GetAlgo())), nil
}


func (kb dbKeybase) Close() error { return kb.db.Close() }

func infoKey(name string) string   { return fmt.Sprintf("%s.%s", name, infoSuffix) }
func infoKeyBz(name string) []byte { return []byte(infoKey(name)) }


type KeybaseOption func(*kbOptions)

type kbOptions struct {
}
