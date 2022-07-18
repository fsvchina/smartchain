package types

import (
	tmprotocrypto "github.com/tendermint/tendermint/proto/tendermint/crypto"

	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


type DelegationI interface {
	GetDelegatorAddr() sdk.AccAddress
	GetValidatorAddr() sdk.ValAddress
	GetShares() sdk.Dec
}


type ValidatorI interface {
	IsJailed() bool
	GetMoniker() string
	GetStatus() BondStatus
	IsBonded() bool
	IsUnbonded() bool
	IsUnbonding() bool
	GetOperator() sdk.ValAddress
	ConsPubKey() (cryptotypes.PubKey, error)
	TmConsPublicKey() (tmprotocrypto.PublicKey, error)
	GetConsAddr() (sdk.ConsAddress, error)
	GetTokens() sdk.Int
	GetBondedTokens() sdk.Int
	GetConsensusPower(sdk.Int) int64
	GetCommission() sdk.Dec
	GetMinSelfDelegation() sdk.Int
	GetDelegatorShares() sdk.Dec
	TokensFromShares(sdk.Dec) sdk.Dec
	TokensFromSharesTruncated(sdk.Dec) sdk.Dec
	TokensFromSharesRoundUp(sdk.Dec) sdk.Dec
	SharesFromTokens(amt sdk.Int) (sdk.Dec, error)
	SharesFromTokensTruncated(amt sdk.Int) (sdk.Dec, error)
}
