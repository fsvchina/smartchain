package web3

import (
	"github.com/tharsis/ethermint/version"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)


type PublicAPI struct{}


func NewPublicAPI() *PublicAPI {
	return &PublicAPI{}
}


func (a *PublicAPI) ClientVersion() string {
	return version.Version()
}


func (a *PublicAPI) Sha3(input string) hexutil.Bytes {
	return crypto.Keccak256(hexutil.Bytes(input))
}
