package keyring

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/codec/legacy"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/crypto/hd"
	"github.com/cosmos/cosmos-sdk/crypto/keys/multisig"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/types"
)


type Info interface {

	GetType() KeyType

	GetName() string

	GetPubKey() cryptotypes.PubKey

	GetAddress() types.AccAddress

	GetPath() (*hd.BIP44Params, error)

	GetAlgo() hd.PubKeyType
}

var (
	_ Info = &localInfo{}
	_ Info = &ledgerInfo{}
	_ Info = &offlineInfo{}
	_ Info = &multiInfo{}
)



type localInfo struct {
	Name         string             `json:"name"`
	PubKey       cryptotypes.PubKey `json:"pubkey"`
	PrivKeyArmor string             `json:"privkey.armor"`
	Algo         hd.PubKeyType      `json:"algo"`
}

func newLocalInfo(name string, pub cryptotypes.PubKey, privArmor string, algo hd.PubKeyType) Info {
	return &localInfo{
		Name:         name,
		PubKey:       pub,
		PrivKeyArmor: privArmor,
		Algo:         algo,
	}
}


func (i localInfo) GetType() KeyType {
	return TypeLocal
}


func (i localInfo) GetName() string {
	return i.Name
}


func (i localInfo) GetPubKey() cryptotypes.PubKey {
	return i.PubKey
}


func (i localInfo) GetAddress() types.AccAddress {
	return i.PubKey.Address().Bytes()
}


func (i localInfo) GetAlgo() hd.PubKeyType {
	return i.Algo
}


func (i localInfo) GetPath() (*hd.BIP44Params, error) {
	return nil, fmt.Errorf("BIP44 Paths are not available for this type")
}



type ledgerInfo struct {
	Name   string             `json:"name"`
	PubKey cryptotypes.PubKey `json:"pubkey"`
	Path   hd.BIP44Params     `json:"path"`
	Algo   hd.PubKeyType      `json:"algo"`
}

func newLedgerInfo(name string, pub cryptotypes.PubKey, path hd.BIP44Params, algo hd.PubKeyType) Info {
	return &ledgerInfo{
		Name:   name,
		PubKey: pub,
		Path:   path,
		Algo:   algo,
	}
}


func (i ledgerInfo) GetType() KeyType {
	return TypeLedger
}


func (i ledgerInfo) GetName() string {
	return i.Name
}


func (i ledgerInfo) GetPubKey() cryptotypes.PubKey {
	return i.PubKey
}


func (i ledgerInfo) GetAddress() types.AccAddress {
	return i.PubKey.Address().Bytes()
}


func (i ledgerInfo) GetAlgo() hd.PubKeyType {
	return i.Algo
}


func (i ledgerInfo) GetPath() (*hd.BIP44Params, error) {
	tmp := i.Path
	return &tmp, nil
}



type offlineInfo struct {
	Name   string             `json:"name"`
	PubKey cryptotypes.PubKey `json:"pubkey"`
	Algo   hd.PubKeyType      `json:"algo"`
}

func newOfflineInfo(name string, pub cryptotypes.PubKey, algo hd.PubKeyType) Info {
	return &offlineInfo{
		Name:   name,
		PubKey: pub,
		Algo:   algo,
	}
}


func (i offlineInfo) GetType() KeyType {
	return TypeOffline
}


func (i offlineInfo) GetName() string {
	return i.Name
}


func (i offlineInfo) GetPubKey() cryptotypes.PubKey {
	return i.PubKey
}


func (i offlineInfo) GetAlgo() hd.PubKeyType {
	return i.Algo
}


func (i offlineInfo) GetAddress() types.AccAddress {
	return i.PubKey.Address().Bytes()
}


func (i offlineInfo) GetPath() (*hd.BIP44Params, error) {
	return nil, fmt.Errorf("BIP44 Paths are not available for this type")
}





type multisigPubKeyInfo struct {
	PubKey cryptotypes.PubKey `json:"pubkey"`
	Weight uint               `json:"weight"`
}


type multiInfo struct {
	Name      string               `json:"name"`
	PubKey    cryptotypes.PubKey   `json:"pubkey"`
	Threshold uint                 `json:"threshold"`
	PubKeys   []multisigPubKeyInfo `json:"pubkeys"`
}


func NewMultiInfo(name string, pub cryptotypes.PubKey) (Info, error) {
	if _, ok := pub.(*multisig.LegacyAminoPubKey); !ok {
		return nil, fmt.Errorf("MultiInfo supports only multisig.LegacyAminoPubKey, got  %T", pub)
	}
	return &multiInfo{
		Name:   name,
		PubKey: pub,
	}, nil
}


func (i multiInfo) GetType() KeyType {
	return TypeMulti
}


func (i multiInfo) GetName() string {
	return i.Name
}


func (i multiInfo) GetPubKey() cryptotypes.PubKey {
	return i.PubKey
}


func (i multiInfo) GetAddress() types.AccAddress {
	return i.PubKey.Address().Bytes()
}


func (i multiInfo) GetAlgo() hd.PubKeyType {
	return hd.MultiType
}


func (i multiInfo) GetPath() (*hd.BIP44Params, error) {
	return nil, fmt.Errorf("BIP44 Paths are not available for this type")
}


func (i multiInfo) UnpackInterfaces(unpacker codectypes.AnyUnpacker) error {
	multiPK := i.PubKey.(*multisig.LegacyAminoPubKey)

	return codectypes.UnpackInterfaces(multiPK, unpacker)
}


func marshalInfo(i Info) []byte {
	return legacy.Cdc.MustMarshalLengthPrefixed(i)
}


func unmarshalInfo(bz []byte) (info Info, err error) {
	err = legacy.Cdc.UnmarshalLengthPrefixed(bz, &info)
	if err != nil {
		return nil, err
	}





	//


	_, ok := info.(multiInfo)
	if ok {
		var multi multiInfo
		err = legacy.Cdc.UnmarshalLengthPrefixed(bz, &multi)

		return multi, err
	}

	return
}
