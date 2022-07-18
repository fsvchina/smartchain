package types

import (
	"bytes"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

var (
	_ authtypes.AccountI                 = (*EthAccount)(nil)
	_ EthAccountI                        = (*EthAccount)(nil)
	_ authtypes.GenesisAccount           = (*EthAccount)(nil)
	_ codectypes.UnpackInterfacesMessage = (*EthAccount)(nil)
)

var emptyCodeHash = crypto.Keccak256(nil)

const (

	AccountTypeEOA = int8(iota + 1)

	AccountTypeContract
)


type EthAccountI interface {
	authtypes.AccountI

	EthAddress() common.Address

	GetCodeHash() common.Hash

	SetCodeHash(code common.Hash) error

	Type() int8
}







func ProtoAccount() authtypes.AccountI {
	return &EthAccount{
		BaseAccount: &authtypes.BaseAccount{},
		CodeHash:    common.BytesToHash(emptyCodeHash).String(),
	}
}


func (acc EthAccount) EthAddress() common.Address {
	return common.BytesToAddress(acc.GetAddress().Bytes())
}


func (acc EthAccount) GetCodeHash() common.Hash {
	return common.HexToHash(acc.CodeHash)
}


func (acc *EthAccount) SetCodeHash(codeHash common.Hash) error {
	acc.CodeHash = codeHash.Hex()
	return nil
}


func (acc EthAccount) Type() int8 {
	if bytes.Equal(emptyCodeHash, common.Hex2Bytes(acc.CodeHash)) {
		return AccountTypeEOA
	}
	return AccountTypeContract
}
