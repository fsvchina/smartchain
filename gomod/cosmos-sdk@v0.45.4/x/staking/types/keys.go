package types

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (

	ModuleName = "staking"


	StoreKey = ModuleName


	QuerierRoute = ModuleName


	RouterKey = ModuleName
)

var (


	LastValidatorPowerKey = []byte{0x11}
	LastTotalPowerKey     = []byte{0x12}

	ValidatorsKey             = []byte{0x21}
	ValidatorsByConsAddrKey   = []byte{0x22}
	ValidatorsByPowerIndexKey = []byte{0x23}

	DelegationKey                    = []byte{0x31}
	UnbondingDelegationKey           = []byte{0x32}
	UnbondingDelegationByValIndexKey = []byte{0x33}
	RedelegationKey                  = []byte{0x34}
	RedelegationByValSrcIndexKey     = []byte{0x35}
	RedelegationByValDstIndexKey     = []byte{0x36}

	UnbondingQueueKey    = []byte{0x41}
	RedelegationQueueKey = []byte{0x42}
	ValidatorQueueKey    = []byte{0x43}

	HistoricalInfoKey = []byte{0x50}
)



func GetValidatorKey(operatorAddr sdk.ValAddress) []byte {
	return append(ValidatorsKey, address.MustLengthPrefix(operatorAddr)...)
}



func GetValidatorByConsAddrKey(addr sdk.ConsAddress) []byte {
	return append(ValidatorsByConsAddrKey, address.MustLengthPrefix(addr)...)
}


func AddressFromValidatorsKey(key []byte) []byte {
	return key[2:]
}


func AddressFromLastValidatorPowerKey(key []byte) []byte {
	return key[2:]
}





func GetValidatorsByPowerIndexKey(validator Validator, powerReduction sdk.Int) []byte {



	consensusPower := sdk.TokensToConsensusPower(validator.Tokens, powerReduction)
	consensusPowerBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(consensusPowerBytes, uint64(consensusPower))

	powerBytes := consensusPowerBytes
	powerBytesLen := len(powerBytes)

	addr, err := sdk.ValAddressFromBech32(validator.OperatorAddress)
	if err != nil {
		panic(err)
	}
	operAddrInvr := sdk.CopyBytes(addr)
	addrLen := len(operAddrInvr)

	for i, b := range operAddrInvr {
		operAddrInvr[i] = ^b
	}


	key := make([]byte, 1+powerBytesLen+1+addrLen)

	key[0] = ValidatorsByPowerIndexKey[0]
	copy(key[1:powerBytesLen+1], powerBytes)
	key[powerBytesLen+1] = byte(addrLen)
	copy(key[powerBytesLen+2:], operAddrInvr)

	return key
}


func GetLastValidatorPowerKey(operator sdk.ValAddress) []byte {
	return append(LastValidatorPowerKey, address.MustLengthPrefix(operator)...)
}


func ParseValidatorPowerRankKey(key []byte) (operAddr []byte) {
	powerBytesLen := 8


	operAddr = sdk.CopyBytes(key[powerBytesLen+2:])

	for i, b := range operAddr {
		operAddr[i] = ^b
	}

	return operAddr
}



func GetValidatorQueueKey(timestamp time.Time, height int64) []byte {
	heightBz := sdk.Uint64ToBigEndian(uint64(height))
	timeBz := sdk.FormatTimeBytes(timestamp)
	timeBzL := len(timeBz)
	prefixL := len(ValidatorQueueKey)

	bz := make([]byte, prefixL+8+timeBzL+8)


	copy(bz[:prefixL], ValidatorQueueKey)


	copy(bz[prefixL:prefixL+8], sdk.Uint64ToBigEndian(uint64(timeBzL)))


	copy(bz[prefixL+8:prefixL+8+timeBzL], timeBz)


	copy(bz[prefixL+8+timeBzL:], heightBz)

	return bz
}



func ParseValidatorQueueKey(bz []byte) (time.Time, int64, error) {
	prefixL := len(ValidatorQueueKey)
	if prefix := bz[:prefixL]; !bytes.Equal(prefix, ValidatorQueueKey) {
		return time.Time{}, 0, fmt.Errorf("invalid prefix; expected: %X, got: %X", ValidatorQueueKey, prefix)
	}

	timeBzL := sdk.BigEndianToUint64(bz[prefixL : prefixL+8])
	ts, err := sdk.ParseTimeBytes(bz[prefixL+8 : prefixL+8+int(timeBzL)])
	if err != nil {
		return time.Time{}, 0, err
	}

	height := sdk.BigEndianToUint64(bz[prefixL+8+int(timeBzL):])

	return ts, int64(height), nil
}



func GetDelegationKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(GetDelegationsKey(delAddr), address.MustLengthPrefix(valAddr)...)
}


func GetDelegationsKey(delAddr sdk.AccAddress) []byte {
	return append(DelegationKey, address.MustLengthPrefix(delAddr)...)
}



func GetUBDKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(GetUBDsKey(delAddr.Bytes()), address.MustLengthPrefix(valAddr)...)
}



func GetUBDByValIndexKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(GetUBDsByValIndexKey(valAddr), address.MustLengthPrefix(delAddr)...)
}


func GetUBDKeyFromValIndexKey(indexKey []byte) []byte {
	addrs := indexKey[1:]

	valAddrLen := addrs[0]
	valAddr := addrs[1 : 1+valAddrLen]
	delAddr := addrs[valAddrLen+2:]

	return GetUBDKey(delAddr, valAddr)
}


func GetUBDsKey(delAddr sdk.AccAddress) []byte {
	return append(UnbondingDelegationKey, address.MustLengthPrefix(delAddr)...)
}


func GetUBDsByValIndexKey(valAddr sdk.ValAddress) []byte {
	return append(UnbondingDelegationByValIndexKey, address.MustLengthPrefix(valAddr)...)
}


func GetUnbondingDelegationTimeKey(timestamp time.Time) []byte {
	bz := sdk.FormatTimeBytes(timestamp)
	return append(UnbondingQueueKey, bz...)
}



func GetREDKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) []byte {

	key := make([]byte, 1+3+len(delAddr)+len(valSrcAddr)+len(valDstAddr))

	copy(key[0:2+len(delAddr)], GetREDsKey(delAddr.Bytes()))
	key[2+len(delAddr)] = byte(len(valSrcAddr))
	copy(key[3+len(delAddr):3+len(delAddr)+len(valSrcAddr)], valSrcAddr.Bytes())
	key[3+len(delAddr)+len(valSrcAddr)] = byte(len(valDstAddr))
	copy(key[4+len(delAddr)+len(valSrcAddr):], valDstAddr.Bytes())

	return key
}



func GetREDByValSrcIndexKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) []byte {
	REDSFromValsSrcKey := GetREDsFromValSrcIndexKey(valSrcAddr)
	offset := len(REDSFromValsSrcKey)


	key := make([]byte, offset+2+len(delAddr)+len(valDstAddr))
	copy(key[0:offset], REDSFromValsSrcKey)
	key[offset] = byte(len(delAddr))
	copy(key[offset+1:offset+1+len(delAddr)], delAddr.Bytes())
	key[offset+1+len(delAddr)] = byte(len(valDstAddr))
	copy(key[offset+2+len(delAddr):], valDstAddr.Bytes())

	return key
}



func GetREDByValDstIndexKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) []byte {
	REDSToValsDstKey := GetREDsToValDstIndexKey(valDstAddr)
	offset := len(REDSToValsDstKey)


	key := make([]byte, offset+2+len(delAddr)+len(valSrcAddr))
	copy(key[0:offset], REDSToValsDstKey)
	key[offset] = byte(len(delAddr))
	copy(key[offset+1:offset+1+len(delAddr)], delAddr.Bytes())
	key[offset+1+len(delAddr)] = byte(len(valSrcAddr))
	copy(key[offset+2+len(delAddr):], valSrcAddr.Bytes())

	return key
}


func GetREDKeyFromValSrcIndexKey(indexKey []byte) []byte {

	addrs := indexKey[1:]

	valSrcAddrLen := addrs[0]
	valSrcAddr := addrs[1 : valSrcAddrLen+1]
	delAddrLen := addrs[valSrcAddrLen+1]
	delAddr := addrs[valSrcAddrLen+2 : valSrcAddrLen+2+delAddrLen]
	valDstAddr := addrs[valSrcAddrLen+delAddrLen+3:]

	return GetREDKey(delAddr, valSrcAddr, valDstAddr)
}


func GetREDKeyFromValDstIndexKey(indexKey []byte) []byte {

	addrs := indexKey[1:]

	valDstAddrLen := addrs[0]
	valDstAddr := addrs[1 : valDstAddrLen+1]
	delAddrLen := addrs[valDstAddrLen+1]
	delAddr := addrs[valDstAddrLen+2 : valDstAddrLen+2+delAddrLen]
	valSrcAddr := addrs[valDstAddrLen+delAddrLen+3:]

	return GetREDKey(delAddr, valSrcAddr, valDstAddr)
}



func GetRedelegationTimeKey(timestamp time.Time) []byte {
	bz := sdk.FormatTimeBytes(timestamp)
	return append(RedelegationQueueKey, bz...)
}



func GetREDsKey(delAddr sdk.AccAddress) []byte {
	return append(RedelegationKey, address.MustLengthPrefix(delAddr)...)
}



func GetREDsFromValSrcIndexKey(valSrcAddr sdk.ValAddress) []byte {
	return append(RedelegationByValSrcIndexKey, address.MustLengthPrefix(valSrcAddr)...)
}



func GetREDsToValDstIndexKey(valDstAddr sdk.ValAddress) []byte {
	return append(RedelegationByValDstIndexKey, address.MustLengthPrefix(valDstAddr)...)
}



func GetREDsByDelToValDstIndexKey(delAddr sdk.AccAddress, valDstAddr sdk.ValAddress) []byte {
	return append(GetREDsToValDstIndexKey(valDstAddr), address.MustLengthPrefix(delAddr)...)
}


func GetHistoricalInfoKey(height int64) []byte {
	return append(HistoricalInfoKey, []byte(strconv.FormatInt(height, 10))...)
}
