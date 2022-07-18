package importer

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	ethcons "github.com/ethereum/go-ethereum/consensus"
	ethstate "github.com/ethereum/go-ethereum/core/state"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethrpc "github.com/ethereum/go-ethereum/rpc"
)







//


type ChainContext struct {
	Coinbase        common.Address
	headersByNumber map[uint64]*ethtypes.Header
}



func NewChainContext() *ChainContext {
	return &ChainContext{
		headersByNumber: make(map[uint64]*ethtypes.Header),
	}
}



func (cc *ChainContext) Engine() ethcons.Engine {
	return cc
}



func (cc *ChainContext) SetHeader(number uint64, header *ethtypes.Header) {
	cc.headersByNumber[number] = header
}


//


func (cc *ChainContext) GetHeader(_ common.Hash, number uint64) *ethtypes.Header {
	if header, ok := cc.headersByNumber[number]; ok {
		return header
	}

	return nil
}




//


func (cc *ChainContext) Author(_ *ethtypes.Header) (common.Address, error) {
	return cc.Coinbase, nil
}



//


func (cc *ChainContext) APIs(_ ethcons.ChainHeaderReader) []ethrpc.API {
	return nil
}



func (cc *ChainContext) CalcDifficulty(_ ethcons.ChainHeaderReader, _ uint64, _ *ethtypes.Header) *big.Int {
	return nil
}



//

func (cc *ChainContext) Finalize(
	_ ethcons.ChainHeaderReader, _ *ethtypes.Header, _ *ethstate.StateDB,
	_ []*ethtypes.Transaction, _ []*ethtypes.Header) {
}



//



func (cc *ChainContext) FinalizeAndAssemble(_ ethcons.ChainHeaderReader, _ *ethtypes.Header, _ *ethstate.StateDB, _ []*ethtypes.Transaction,
	_ []*ethtypes.Header, _ []*ethtypes.Receipt,
) (*ethtypes.Block, error) {
	return nil, nil
}



//

func (cc *ChainContext) Prepare(_ ethcons.ChainHeaderReader, _ *ethtypes.Header) error {
	return nil
}



//

func (cc *ChainContext) Seal(_ ethcons.ChainHeaderReader, _ *ethtypes.Block, _ chan<- *ethtypes.Block, _ <-chan struct{}) error {
	return nil
}



func (cc *ChainContext) SealHash(header *ethtypes.Header) common.Hash {
	return common.Hash{}
}



//


func (cc *ChainContext) VerifyHeader(_ ethcons.ChainHeaderReader, _ *ethtypes.Header, _ bool) error {
	return nil
}



//


func (cc *ChainContext) VerifyHeaders(_ ethcons.ChainHeaderReader, _ []*ethtypes.Header, _ []bool) (chan<- struct{}, <-chan error) {
	return nil, nil
}



//


func (cc *ChainContext) VerifySeal(_ ethcons.ChainHeaderReader, _ *ethtypes.Header) error {
	return nil
}



func (cc *ChainContext) VerifyUncles(_ ethcons.ChainReader, _ *ethtypes.Block) error {
	return nil
}




func (cc *ChainContext) Close() error {
	return nil
}
