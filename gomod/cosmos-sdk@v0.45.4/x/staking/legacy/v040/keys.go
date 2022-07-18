

package v040

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"strconv"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	v040auth "github.com/cosmos/cosmos-sdk/x/auth/legacy/v040"
	"github.com/cosmos/cosmos-sdk/x/staking/types"
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
	return append(ValidatorsKey, operatorAddr.Bytes()...)
}



func GetValidatorByConsAddrKey(addr sdk.ConsAddress) []byte {
	return append(ValidatorsByConsAddrKey, addr.Bytes()...)
}


func AddressFromLastValidatorPowerKey(key []byte) []byte {
	return key[1:]
}





func GetValidatorsByPowerIndexKey(validator types.Validator) []byte {



	consensusPower := sdk.TokensToConsensusPower(validator.Tokens, sdk.DefaultPowerReduction)
	consensusPowerBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(consensusPowerBytes, uint64(consensusPower))

	powerBytes := consensusPowerBytes
	powerBytesLen := len(powerBytes)


	key := make([]byte, 1+powerBytesLen+v040auth.AddrLen)

	key[0] = ValidatorsByPowerIndexKey[0]
	copy(key[1:powerBytesLen+1], powerBytes)
	addr, err := sdk.ValAddressFromBech32(validator.OperatorAddress)
	if err != nil {
		panic(err)
	}
	operAddrInvr := sdk.CopyBytes(addr)

	for i, b := range operAddrInvr {
		operAddrInvr[i] = ^b
	}

	copy(key[powerBytesLen+1:], operAddrInvr)

	return key
}


func GetLastValidatorPowerKey(operator sdk.ValAddress) []byte {
	return append(LastValidatorPowerKey, operator...)
}


func ParseValidatorPowerRankKey(key []byte) (operAddr []byte) {
	powerBytesLen := 8
	if len(key) != 1+powerBytesLen+v040auth.AddrLen {
		panic("Invalid validator power rank key length")
	}

	operAddr = sdk.CopyBytes(key[powerBytesLen+1:])

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
	return append(GetDelegationsKey(delAddr), valAddr.Bytes()...)
}


func GetDelegationsKey(delAddr sdk.AccAddress) []byte {
	return append(DelegationKey, delAddr.Bytes()...)
}



func GetUBDKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(
		GetUBDsKey(delAddr.Bytes()),
		valAddr.Bytes()...)
}



func GetUBDByValIndexKey(delAddr sdk.AccAddress, valAddr sdk.ValAddress) []byte {
	return append(GetUBDsByValIndexKey(valAddr), delAddr.Bytes()...)
}


func GetUBDKeyFromValIndexKey(indexKey []byte) []byte {
	addrs := indexKey[1:]
	if len(addrs) != 2*v040auth.AddrLen {
		panic("unexpected key length")
	}

	valAddr := addrs[:v040auth.AddrLen]
	delAddr := addrs[v040auth.AddrLen:]

	return GetUBDKey(delAddr, valAddr)
}


func GetUBDsKey(delAddr sdk.AccAddress) []byte {
	return append(UnbondingDelegationKey, delAddr.Bytes()...)
}


func GetUBDsByValIndexKey(valAddr sdk.ValAddress) []byte {
	return append(UnbondingDelegationByValIndexKey, valAddr.Bytes()...)
}


func GetUnbondingDelegationTimeKey(timestamp time.Time) []byte {
	bz := sdk.FormatTimeBytes(timestamp)
	return append(UnbondingQueueKey, bz...)
}



func GetREDKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) []byte {
	key := make([]byte, 1+v040auth.AddrLen*3)

	copy(key[0:v040auth.AddrLen+1], GetREDsKey(delAddr.Bytes()))
	copy(key[v040auth.AddrLen+1:2*v040auth.AddrLen+1], valSrcAddr.Bytes())
	copy(key[2*v040auth.AddrLen+1:3*v040auth.AddrLen+1], valDstAddr.Bytes())

	return key
}



func GetREDByValSrcIndexKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) []byte {
	REDSFromValsSrcKey := GetREDsFromValSrcIndexKey(valSrcAddr)
	offset := len(REDSFromValsSrcKey)


	key := make([]byte, len(REDSFromValsSrcKey)+2*v040auth.AddrLen)
	copy(key[0:offset], REDSFromValsSrcKey)
	copy(key[offset:offset+v040auth.AddrLen], delAddr.Bytes())
	copy(key[offset+v040auth.AddrLen:offset+2*v040auth.AddrLen], valDstAddr.Bytes())

	return key
}



func GetREDByValDstIndexKey(delAddr sdk.AccAddress, valSrcAddr, valDstAddr sdk.ValAddress) []byte {
	REDSToValsDstKey := GetREDsToValDstIndexKey(valDstAddr)
	offset := len(REDSToValsDstKey)


	key := make([]byte, len(REDSToValsDstKey)+2*v040auth.AddrLen)
	copy(key[0:offset], REDSToValsDstKey)
	copy(key[offset:offset+v040auth.AddrLen], delAddr.Bytes())
	copy(key[offset+v040auth.AddrLen:offset+2*v040auth.AddrLen], valSrcAddr.Bytes())

	return key
}


func GetREDKeyFromValSrcIndexKey(indexKey []byte) []byte {

	if len(indexKey) != 3*v040auth.AddrLen+1 {
		panic("unexpected key length")
	}

	valSrcAddr := indexKey[1 : v040auth.AddrLen+1]
	delAddr := indexKey[v040auth.AddrLen+1 : 2*v040auth.AddrLen+1]
	valDstAddr := indexKey[2*v040auth.AddrLen+1 : 3*v040auth.AddrLen+1]

	return GetREDKey(delAddr, valSrcAddr, valDstAddr)
}


func GetREDKeyFromValDstIndexKey(indexKey []byte) []byte {

	if len(indexKey) != 3*v040auth.AddrLen+1 {
		panic("unexpected key length")
	}

	valDstAddr := indexKey[1 : v040auth.AddrLen+1]
	delAddr := indexKey[v040auth.AddrLen+1 : 2*v040auth.AddrLen+1]
	valSrcAddr := indexKey[2*v040auth.AddrLen+1 : 3*v040auth.AddrLen+1]

	return GetREDKey(delAddr, valSrcAddr, valDstAddr)
}



func GetRedelegationTimeKey(timestamp time.Time) []byte {
	bz := sdk.FormatTimeBytes(timestamp)
	return append(RedelegationQueueKey, bz...)
}



func GetREDsKey(delAddr sdk.AccAddress) []byte {
	return append(RedelegationKey, delAddr.Bytes()...)
}



func GetREDsFromValSrcIndexKey(valSrcAddr sdk.ValAddress) []byte {
	return append(RedelegationByValSrcIndexKey, valSrcAddr.Bytes()...)
}



func GetREDsToValDstIndexKey(valDstAddr sdk.ValAddress) []byte {
	return append(RedelegationByValDstIndexKey, valDstAddr.Bytes()...)
}



func GetREDsByDelToValDstIndexKey(delAddr sdk.AccAddress, valDstAddr sdk.ValAddress) []byte {
	return append(GetREDsToValDstIndexKey(valDstAddr), delAddr.Bytes()...)
}


func GetHistoricalInfoKey(height int64) []byte {
	return append(HistoricalInfoKey, []byte(strconv.FormatInt(height, 10))...)
}
