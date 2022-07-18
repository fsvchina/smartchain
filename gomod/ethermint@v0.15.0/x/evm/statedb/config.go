package statedb

import "github.com/ethereum/go-ethereum/common"


type TxConfig struct {
	BlockHash common.Hash
	TxHash    common.Hash
	TxIndex   uint
	LogIndex  uint
}


func NewTxConfig(bhash, thash common.Hash, txIndex, logIndex uint) TxConfig {
	return TxConfig{
		BlockHash: bhash,
		TxHash:    thash,
		TxIndex:   txIndex,
		LogIndex:  logIndex,
	}
}



func NewEmptyTxConfig(bhash common.Hash) TxConfig {
	return TxConfig{
		BlockHash: bhash,
		TxHash:    common.Hash{},
		TxIndex:   0,
		LogIndex:  0,
	}
}
