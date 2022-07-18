package core

import (
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	bankTypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibcTransferTypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
)


var ContractAddressFee = authtypes.NewModuleAddress(authtypes.FeeCollectorName)


var ContractAddressBank = authtypes.NewModuleAddress(bankTypes.ModuleName)


var ContractAddressDistribution = authtypes.NewModuleAddress(distrtypes.ModuleName)


var ContractAddressStakingBonded = authtypes.NewModuleAddress(stakingtypes.BondedPoolName)


var ContractAddressStakingNotBonded = authtypes.NewModuleAddress(stakingtypes.NotBondedPoolName)

var ContractAddressGov = authtypes.NewModuleAddress(govtypes.ModuleName)


var ContractAddressIbcTransfer = authtypes.NewModuleAddress(ibcTransferTypes.ModuleName)
