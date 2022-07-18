package client

import (
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


type Account interface {
	GetAddress() sdk.AccAddress
	GetPubKey() cryptotypes.PubKey
	GetAccountNumber() uint64
	GetSequence() uint64
}




type AccountRetriever interface {
	GetAccount(clientCtx Context, addr sdk.AccAddress) (Account, error)
	GetAccountWithHeight(clientCtx Context, addr sdk.AccAddress) (Account, int64, error)
	EnsureExists(clientCtx Context, addr sdk.AccAddress) error
	GetAccountNumberSequence(clientCtx Context, addr sdk.AccAddress) (accNum uint64, accSeq uint64, err error)
}
