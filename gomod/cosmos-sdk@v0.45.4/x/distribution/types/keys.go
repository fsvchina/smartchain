package types

import (
	"encoding/binary"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/address"
)

const (

	ModuleName = "distribution"


	StoreKey = ModuleName


	RouterKey = ModuleName


	QuerierRoute = ModuleName
)



//

//

//

//

//

//

//

//

//

var (
	FeePoolKey                        = []byte{0x00}
	ProposerKey                       = []byte{0x01}
	ValidatorOutstandingRewardsPrefix = []byte{0x02}

	DelegatorWithdrawAddrPrefix          = []byte{0x03}
	DelegatorStartingInfoPrefix          = []byte{0x04}
	ValidatorHistoricalRewardsPrefix     = []byte{0x05}
	ValidatorCurrentRewardsPrefix        = []byte{0x06}
	ValidatorAccumulatedCommissionPrefix = []byte{0x07}
	ValidatorSlashEventPrefix            = []byte{0x08}
)


func GetValidatorOutstandingRewardsAddress(key []byte) (valAddr sdk.ValAddress) {




	addr := key[2:]
	if len(addr) != int(key[1]) {
		panic("unexpected key length")
	}

	return sdk.ValAddress(addr)
}


func GetDelegatorWithdrawInfoAddress(key []byte) (delAddr sdk.AccAddress) {




	addr := key[2:]
	if len(addr) != int(key[1]) {
		panic("unexpected key length")
	}

	return sdk.AccAddress(addr)
}


func GetDelegatorStartingInfoAddresses(key []byte) (valAddr sdk.ValAddress, delAddr sdk.AccAddress) {


	valAddrLen := int(key[1])
	valAddr = sdk.ValAddress(key[2 : 2+valAddrLen])
	delAddrLen := int(key[2+valAddrLen])
	delAddr = sdk.AccAddress(key[3+valAddrLen:])
	if len(delAddr.Bytes()) != delAddrLen {
		panic("unexpected key length")
	}

	return
}


func GetValidatorHistoricalRewardsAddressPeriod(key []byte) (valAddr sdk.ValAddress, period uint64) {


	valAddrLen := int(key[1])
	valAddr = sdk.ValAddress(key[2 : 2+valAddrLen])
	b := key[2+valAddrLen:]
	if len(b) != 8 {
		panic("unexpected key length")
	}
	period = binary.LittleEndian.Uint64(b)
	return
}


func GetValidatorCurrentRewardsAddress(key []byte) (valAddr sdk.ValAddress) {




	addr := key[2:]
	if len(addr) != int(key[1]) {
		panic("unexpected key length")
	}

	return sdk.ValAddress(addr)
}


func GetValidatorAccumulatedCommissionAddress(key []byte) (valAddr sdk.ValAddress) {




	addr := key[2:]
	if len(addr) != int(key[1]) {
		panic("unexpected key length")
	}

	return sdk.ValAddress(addr)
}


func GetValidatorSlashEventAddressHeight(key []byte) (valAddr sdk.ValAddress, height uint64) {


	valAddrLen := int(key[1])
	valAddr = key[2 : 2+valAddrLen]
	startB := 2 + valAddrLen
	b := key[startB : startB+8]
	height = binary.BigEndian.Uint64(b)
	return
}


func GetValidatorOutstandingRewardsKey(valAddr sdk.ValAddress) []byte {
	return append(ValidatorOutstandingRewardsPrefix, address.MustLengthPrefix(valAddr.Bytes())...)
}


func GetDelegatorWithdrawAddrKey(delAddr sdk.AccAddress) []byte {
	return append(DelegatorWithdrawAddrPrefix, address.MustLengthPrefix(delAddr.Bytes())...)
}


func GetDelegatorStartingInfoKey(v sdk.ValAddress, d sdk.AccAddress) []byte {
	return append(append(DelegatorStartingInfoPrefix, address.MustLengthPrefix(v.Bytes())...), address.MustLengthPrefix(d.Bytes())...)
}


func GetValidatorHistoricalRewardsPrefix(v sdk.ValAddress) []byte {
	return append(ValidatorHistoricalRewardsPrefix, address.MustLengthPrefix(v.Bytes())...)
}


func GetValidatorHistoricalRewardsKey(v sdk.ValAddress, k uint64) []byte {
	b := make([]byte, 8)
	binary.LittleEndian.PutUint64(b, k)
	return append(append(ValidatorHistoricalRewardsPrefix, address.MustLengthPrefix(v.Bytes())...), b...)
}


func GetValidatorCurrentRewardsKey(v sdk.ValAddress) []byte {
	return append(ValidatorCurrentRewardsPrefix, address.MustLengthPrefix(v.Bytes())...)
}


func GetValidatorAccumulatedCommissionKey(v sdk.ValAddress) []byte {
	return append(ValidatorAccumulatedCommissionPrefix, address.MustLengthPrefix(v.Bytes())...)
}


func GetValidatorSlashEventPrefix(v sdk.ValAddress) []byte {
	return append(ValidatorSlashEventPrefix, address.MustLengthPrefix(v.Bytes())...)
}


func GetValidatorSlashEventKeyPrefix(v sdk.ValAddress, height uint64) []byte {
	heightBz := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBz, height)

	return append(
		ValidatorSlashEventPrefix,
		append(address.MustLengthPrefix(v.Bytes()), heightBz...)...,
	)
}


func GetValidatorSlashEventKey(v sdk.ValAddress, height, period uint64) []byte {
	periodBz := make([]byte, 8)
	binary.BigEndian.PutUint64(periodBz, period)
	prefix := GetValidatorSlashEventKeyPrefix(v, height)

	return append(prefix, periodBz...)
}
