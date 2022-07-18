package types

import (
	"bytes"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"

	"github.com/hashicorp/golang-lru/simplelru"
	yaml "gopkg.in/yaml.v2"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	"github.com/cosmos/cosmos-sdk/internal/conv"
	"github.com/cosmos/cosmos-sdk/types/address"
	"github.com/cosmos/cosmos-sdk/types/bech32"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (



	//









	Bech32MainPrefix = "cosmos"


	Purpose = 44


	CoinType = 118



	FullFundraiserPath = "m/44'/118'/0'/0/0"


	PrefixAccount = "acc"

	PrefixValidator = "val"

	PrefixConsensus = "cons"

	PrefixPublic = "pub"

	PrefixOperator = "oper"


	PrefixAddress = "addr"


	Bech32PrefixAccAddr = Bech32MainPrefix

	Bech32PrefixAccPub = Bech32MainPrefix + PrefixPublic

	Bech32PrefixValAddr = Bech32MainPrefix + PrefixValidator + PrefixOperator

	Bech32PrefixValPub = Bech32MainPrefix + PrefixValidator + PrefixOperator + PrefixPublic

	Bech32PrefixConsAddr = Bech32MainPrefix + PrefixValidator + PrefixConsensus

	Bech32PrefixConsPub = Bech32MainPrefix + PrefixValidator + PrefixConsensus + PrefixPublic
)


var (


	accAddrMu     sync.Mutex
	accAddrCache  *simplelru.LRU
	consAddrMu    sync.Mutex
	consAddrCache *simplelru.LRU
	valAddrMu     sync.Mutex
	valAddrCache  *simplelru.LRU
)

func init() {
	var err error


	if accAddrCache, err = simplelru.NewLRU(60000, nil); err != nil {
		panic(err)
	}
	if consAddrCache, err = simplelru.NewLRU(500, nil); err != nil {
		panic(err)
	}
	if valAddrCache, err = simplelru.NewLRU(500, nil); err != nil {
		panic(err)
	}
}


type Address interface {
	Equals(Address) bool
	Empty() bool
	Marshal() ([]byte, error)
	MarshalJSON() ([]byte, error)
	Bytes() []byte
	String() string
	Format(s fmt.State, verb rune)
}


var _ Address = AccAddress{}
var _ Address = ValAddress{}
var _ Address = ConsAddress{}

var _ yaml.Marshaler = AccAddress{}
var _ yaml.Marshaler = ValAddress{}
var _ yaml.Marshaler = ConsAddress{}







type AccAddress []byte


func AccAddressFromHex(address string) (addr AccAddress, err error) {
	bz, err := addressBytesFromHexString(address)
	return AccAddress(bz), err
}




func VerifyAddressFormat(bz []byte) error {
	verifier := GetConfig().GetAddressVerifier()
	if verifier != nil {
		return verifier(bz)
	}

	if len(bz) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrUnknownAddress, "addresses cannot be empty")
	}

	if len(bz) > address.MaxAddrLen {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownAddress, "address max length is %d, got %d", address.MaxAddrLen, len(bz))
	}

	return nil
}


func AccAddressFromBech32(address string) (addr AccAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return AccAddress{}, errors.New("empty address string is not allowed")
	}

	bech32PrefixAccAddr := GetConfig().GetBech32AccountAddrPrefix()

	bz, err := GetFromBech32(address, bech32PrefixAccAddr)
	if err != nil {
		return nil, err
	}

	err = VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return AccAddress(bz), nil
}


func (aa AccAddress) Equals(aa2 Address) bool {
	if aa.Empty() && aa2.Empty() {
		return true
	}

	return bytes.Equal(aa.Bytes(), aa2.Bytes())
}


func (aa AccAddress) Empty() bool {
	return aa == nil || len(aa) == 0
}



func (aa AccAddress) Marshal() ([]byte, error) {
	return aa, nil
}



func (aa *AccAddress) Unmarshal(data []byte) error {
	*aa = data
	return nil
}


func (aa AccAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(aa.String())
}


func (aa AccAddress) MarshalYAML() (interface{}, error) {
	return aa.String(), nil
}


func (aa *AccAddress) UnmarshalJSON(data []byte) error {
	var s string
	err := json.Unmarshal(data, &s)

	if err != nil {
		return err
	}
	if s == "" {
		*aa = AccAddress{}
		return nil
	}

	aa2, err := AccAddressFromBech32(s)
	if err != nil {
		return err
	}

	*aa = aa2
	return nil
}


func (aa *AccAddress) UnmarshalYAML(data []byte) error {
	var s string
	err := yaml.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*aa = AccAddress{}
		return nil
	}

	aa2, err := AccAddressFromBech32(s)
	if err != nil {
		return err
	}

	*aa = aa2
	return nil
}


func (aa AccAddress) Bytes() []byte {
	return aa
}


func (aa AccAddress) String() string {
	if aa.Empty() {
		return ""
	}

	var key = conv.UnsafeBytesToStr(aa)
	accAddrMu.Lock()
	defer accAddrMu.Unlock()
	addr, ok := accAddrCache.Get(key)
	if ok {
		return addr.(string)
	}
	return cacheBech32Addr(GetConfig().GetBech32AccountAddrPrefix(), aa, accAddrCache, key)
}



func (aa AccAddress) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(aa.String()))
	case 'p':
		s.Write([]byte(fmt.Sprintf("%p", aa)))
	default:
		s.Write([]byte(fmt.Sprintf("%X", []byte(aa))))
	}
}







type ValAddress []byte


func ValAddressFromHex(address string) (addr ValAddress, err error) {
	bz, err := addressBytesFromHexString(address)
	return ValAddress(bz), err
}


func ValAddressFromBech32(address string) (addr ValAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return ValAddress{}, errors.New("empty address string is not allowed")
	}

	bech32PrefixValAddr := GetConfig().GetBech32ValidatorAddrPrefix()

	bz, err := GetFromBech32(address, bech32PrefixValAddr)
	if err != nil {
		return nil, err
	}

	err = VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return ValAddress(bz), nil
}


func (va ValAddress) Equals(va2 Address) bool {
	if va.Empty() && va2.Empty() {
		return true
	}

	return bytes.Equal(va.Bytes(), va2.Bytes())
}


func (va ValAddress) Empty() bool {
	return va == nil || len(va) == 0
}



func (va ValAddress) Marshal() ([]byte, error) {
	return va, nil
}



func (va *ValAddress) Unmarshal(data []byte) error {
	*va = data
	return nil
}


func (va ValAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(va.String())
}


func (va ValAddress) MarshalYAML() (interface{}, error) {
	return va.String(), nil
}


func (va *ValAddress) UnmarshalJSON(data []byte) error {
	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*va = ValAddress{}
		return nil
	}

	va2, err := ValAddressFromBech32(s)
	if err != nil {
		return err
	}

	*va = va2
	return nil
}


func (va *ValAddress) UnmarshalYAML(data []byte) error {
	var s string

	err := yaml.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*va = ValAddress{}
		return nil
	}

	va2, err := ValAddressFromBech32(s)
	if err != nil {
		return err
	}

	*va = va2
	return nil
}


func (va ValAddress) Bytes() []byte {
	return va
}


func (va ValAddress) String() string {
	if va.Empty() {
		return ""
	}

	var key = conv.UnsafeBytesToStr(va)
	valAddrMu.Lock()
	defer valAddrMu.Unlock()
	addr, ok := valAddrCache.Get(key)
	if ok {
		return addr.(string)
	}
	return cacheBech32Addr(GetConfig().GetBech32ValidatorAddrPrefix(), va, valAddrCache, key)
}



func (va ValAddress) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(va.String()))
	case 'p':
		s.Write([]byte(fmt.Sprintf("%p", va)))
	default:
		s.Write([]byte(fmt.Sprintf("%X", []byte(va))))
	}
}







type ConsAddress []byte


func ConsAddressFromHex(address string) (addr ConsAddress, err error) {
	bz, err := addressBytesFromHexString(address)
	return ConsAddress(bz), err
}


func ConsAddressFromBech32(address string) (addr ConsAddress, err error) {
	if len(strings.TrimSpace(address)) == 0 {
		return ConsAddress{}, errors.New("empty address string is not allowed")
	}

	bech32PrefixConsAddr := GetConfig().GetBech32ConsensusAddrPrefix()

	bz, err := GetFromBech32(address, bech32PrefixConsAddr)
	if err != nil {
		return nil, err
	}

	err = VerifyAddressFormat(bz)
	if err != nil {
		return nil, err
	}

	return ConsAddress(bz), nil
}


func GetConsAddress(pubkey cryptotypes.PubKey) ConsAddress {
	return ConsAddress(pubkey.Address())
}


func (ca ConsAddress) Equals(ca2 Address) bool {
	if ca.Empty() && ca2.Empty() {
		return true
	}

	return bytes.Equal(ca.Bytes(), ca2.Bytes())
}


func (ca ConsAddress) Empty() bool {
	return ca == nil || len(ca) == 0
}



func (ca ConsAddress) Marshal() ([]byte, error) {
	return ca, nil
}



func (ca *ConsAddress) Unmarshal(data []byte) error {
	*ca = data
	return nil
}


func (ca ConsAddress) MarshalJSON() ([]byte, error) {
	return json.Marshal(ca.String())
}


func (ca ConsAddress) MarshalYAML() (interface{}, error) {
	return ca.String(), nil
}


func (ca *ConsAddress) UnmarshalJSON(data []byte) error {
	var s string

	err := json.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*ca = ConsAddress{}
		return nil
	}

	ca2, err := ConsAddressFromBech32(s)
	if err != nil {
		return err
	}

	*ca = ca2
	return nil
}


func (ca *ConsAddress) UnmarshalYAML(data []byte) error {
	var s string

	err := yaml.Unmarshal(data, &s)
	if err != nil {
		return err
	}
	if s == "" {
		*ca = ConsAddress{}
		return nil
	}

	ca2, err := ConsAddressFromBech32(s)
	if err != nil {
		return err
	}

	*ca = ca2
	return nil
}


func (ca ConsAddress) Bytes() []byte {
	return ca
}


func (ca ConsAddress) String() string {
	if ca.Empty() {
		return ""
	}

	var key = conv.UnsafeBytesToStr(ca)
	consAddrMu.Lock()
	defer consAddrMu.Unlock()
	addr, ok := consAddrCache.Get(key)
	if ok {
		return addr.(string)
	}
	return cacheBech32Addr(GetConfig().GetBech32ConsensusAddrPrefix(), ca, consAddrCache, key)
}




func Bech32ifyAddressBytes(prefix string, bs []byte) (string, error) {
	if len(bs) == 0 {
		return "", nil
	}
	if len(prefix) == 0 {
		return "", errors.New("prefix cannot be empty")
	}
	return bech32.ConvertAndEncode(prefix, bs)
}




func MustBech32ifyAddressBytes(prefix string, bs []byte) string {
	s, err := Bech32ifyAddressBytes(prefix, bs)
	if err != nil {
		panic(err)
	}
	return s
}



func (ca ConsAddress) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(ca.String()))
	case 'p':
		s.Write([]byte(fmt.Sprintf("%p", ca)))
	default:
		s.Write([]byte(fmt.Sprintf("%X", []byte(ca))))
	}
}





var errBech32EmptyAddress = errors.New("decoding Bech32 address failed: must provide a non empty address")


func GetFromBech32(bech32str, prefix string) ([]byte, error) {
	if len(bech32str) == 0 {
		return nil, errBech32EmptyAddress
	}

	hrp, bz, err := bech32.DecodeAndConvert(bech32str)
	if err != nil {
		return nil, err
	}

	if hrp != prefix {
		return nil, fmt.Errorf("invalid Bech32 prefix; expected %s, got %s", prefix, hrp)
	}

	return bz, nil
}

func addressBytesFromHexString(address string) ([]byte, error) {
	if len(address) == 0 {
		return nil, errors.New("decoding Bech32 address failed: must provide an address")
	}

	return hex.DecodeString(address)
}


func cacheBech32Addr(prefix string, addr []byte, cache *simplelru.LRU, cacheKey string) string {
	bech32Addr, err := bech32.ConvertAndEncode(prefix, addr)
	if err != nil {
		panic(err)
	}
	cache.Add(cacheKey, bech32Addr)
	return bech32Addr
}
