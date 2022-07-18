package types

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"

	"github.com/spf13/cast"
	"google.golang.org/grpc/metadata"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"

	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"

	ethermint "github.com/tharsis/ethermint/types"
)


type BlockNumber int64

const (
	EthPendingBlockNumber  = BlockNumber(-2)
	EthLatestBlockNumber   = BlockNumber(-1)
	EthEarliestBlockNumber = BlockNumber(0)
)

const (
	BlockParamEarliest = "earliest"
	BlockParamLatest   = "latest"
	BlockParamPending  = "pending"
)


func NewBlockNumber(n *big.Int) BlockNumber {
	if !n.IsInt64() {

		return EthLatestBlockNumber
	}

	return BlockNumber(n.Int64())
}




func ContextWithHeight(height int64) context.Context {
	if height == 0 {
		return context.Background()
	}

	return metadata.AppendToOutgoingContext(context.Background(), grpctypes.GRPCBlockHeightHeader, fmt.Sprintf("%d", height))
}







func (bn *BlockNumber) UnmarshalJSON(data []byte) error {
	input := strings.TrimSpace(string(data))
	if len(input) >= 2 && input[0] == '"' && input[len(input)-1] == '"' {
		input = input[1 : len(input)-1]
	}

	switch input {
	case BlockParamEarliest:
		*bn = EthEarliestBlockNumber
		return nil
	case BlockParamLatest:
		*bn = EthLatestBlockNumber
		return nil
	case BlockParamPending:
		*bn = EthPendingBlockNumber
		return nil
	}

	blckNum, err := hexutil.DecodeUint64(input)
	if errors.Is(err, hexutil.ErrMissingPrefix) {
		blckNum = cast.ToUint64(input)
	} else if err != nil {
		return err
	}

	if blckNum > math.MaxInt64 {
		return fmt.Errorf("block number larger than int64")
	}
	*bn = BlockNumber(blckNum)

	return nil
}


func (bn BlockNumber) Int64() int64 {
	if bn < 0 {
		return 0
	} else if bn == 0 {
		return 1
	}

	return int64(bn)
}




func (bn BlockNumber) TmHeight() *int64 {
	if bn < 0 {
		return nil
	}

	height := bn.Int64()
	return &height
}


type BlockNumberOrHash struct {
	BlockNumber *BlockNumber `json:"blockNumber,omitempty"`
	BlockHash   *common.Hash `json:"blockHash,omitempty"`
}

func (bnh *BlockNumberOrHash) UnmarshalJSON(data []byte) error {
	type erased BlockNumberOrHash
	e := erased{}
	err := json.Unmarshal(data, &e)
	if err == nil {
		return bnh.checkUnmarshal(BlockNumberOrHash(e))
	}
	var input string
	err = json.Unmarshal(data, &input)
	if err != nil {
		return err
	}
	err = bnh.decodeFromString(input)
	if err != nil {
		return err
	}

	return nil
}

func (bnh *BlockNumberOrHash) checkUnmarshal(e BlockNumberOrHash) error {
	if e.BlockNumber != nil && e.BlockHash != nil {
		return fmt.Errorf("cannot specify both BlockHash and BlockNumber, choose one or the other")
	}
	bnh.BlockNumber = e.BlockNumber
	bnh.BlockHash = e.BlockHash
	return nil
}

func (bnh *BlockNumberOrHash) decodeFromString(input string) error {
	switch input {
	case BlockParamEarliest:
		bn := EthEarliestBlockNumber
		bnh.BlockNumber = &bn
	case BlockParamLatest:
		bn := EthLatestBlockNumber
		bnh.BlockNumber = &bn
	case BlockParamPending:
		bn := EthPendingBlockNumber
		bnh.BlockNumber = &bn
	default:

		if len(input) == 66 {
			hash := common.Hash{}
			err := hash.UnmarshalText([]byte(input))
			if err != nil {
				return err
			}
			bnh.BlockHash = &hash
			break
		}

		blockNumber, err := hexutil.DecodeUint64(input)
		if err != nil {
			return err
		}

		bnInt, err := ethermint.SafeInt64(blockNumber)
		if err != nil {
			return err
		}

		bn := BlockNumber(bnInt)
		bnh.BlockNumber = &bn
	}
	return nil
}
