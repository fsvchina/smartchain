package filters

import (
	"context"
	"encoding/binary"
	"math/big"

	"github.com/tharsis/ethermint/rpc/types"

	"github.com/pkg/errors"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/eth/filters"
)



type BloomIV struct {
	I [3]uint
	V [3]byte
}


type Filter struct {
	logger   log.Logger
	backend  Backend
	criteria filters.FilterCriteria

	bloomFilters [][]BloomIV
}



func NewBlockFilter(logger log.Logger, backend Backend, criteria filters.FilterCriteria) *Filter {

	return newFilter(logger, backend, criteria, nil)
}



func NewRangeFilter(logger log.Logger, backend Backend, begin, end int64, addresses []common.Address, topics [][]common.Hash) *Filter {



	var filtersBz [][][]byte
	if len(addresses) > 0 {
		filter := make([][]byte, len(addresses))
		for i, address := range addresses {
			filter[i] = address.Bytes()
		}
		filtersBz = append(filtersBz, filter)
	}

	for _, topicList := range topics {
		filter := make([][]byte, len(topicList))
		for i, topic := range topicList {
			filter[i] = topic.Bytes()
		}
		filtersBz = append(filtersBz, filter)
	}


	criteria := filters.FilterCriteria{
		FromBlock: big.NewInt(begin),
		ToBlock:   big.NewInt(end),
		Addresses: addresses,
		Topics:    topics,
	}

	return newFilter(logger, backend, criteria, createBloomFilters(filtersBz, logger))
}


func newFilter(logger log.Logger, backend Backend, criteria filters.FilterCriteria, bloomFilters [][]BloomIV) *Filter {
	return &Filter{
		logger:       logger,
		backend:      backend,
		criteria:     criteria,
		bloomFilters: bloomFilters,
	}
}

const (
	maxToOverhang = 600
)



func (f *Filter) Logs(_ context.Context, logLimit int, blockLimit int64) ([]*ethtypes.Log, error) {
	logs := []*ethtypes.Log{}
	var err error


	if f.criteria.BlockHash != nil && *f.criteria.BlockHash != (common.Hash{}) {
		header, err := f.backend.HeaderByHash(*f.criteria.BlockHash)
		if err != nil {
			return nil, errors.Wrap(err, "failed to fetch header by hash")
		}

		if header == nil {
			return nil, errors.Errorf("unknown block header %s", f.criteria.BlockHash.String())
		}

		return f.blockLogs(header.Number.Int64(), header.Bloom)
	}


	header, err := f.backend.HeaderByNumber(types.EthLatestBlockNumber)
	if err != nil {
		return nil, errors.Wrap(err, "failed to fetch header by number (latest)")
	}

	if header == nil || header.Number == nil {
		f.logger.Debug("header not found or has no number")
		return nil, nil
	}

	head := header.Number.Int64()
	if f.criteria.FromBlock.Int64() < 0 {
		f.criteria.FromBlock = big.NewInt(head)
	} else if f.criteria.FromBlock.Int64() == 0 {
		f.criteria.FromBlock = big.NewInt(1)
	}
	if f.criteria.ToBlock.Int64() < 0 {
		f.criteria.ToBlock = big.NewInt(head)
	} else if f.criteria.ToBlock.Int64() == 0 {
		f.criteria.ToBlock = big.NewInt(1)
	}

	if f.criteria.ToBlock.Int64()-f.criteria.FromBlock.Int64() > blockLimit {
		return nil, errors.Errorf("maximum [from, to] blocks distance: %d", blockLimit)
	}


	if f.criteria.FromBlock.Int64() > head {
		return []*ethtypes.Log{}, nil
	} else if f.criteria.ToBlock.Int64() > head+maxToOverhang {
		f.criteria.ToBlock = big.NewInt(head + maxToOverhang)
	}

	from := f.criteria.FromBlock.Int64()
	to := f.criteria.ToBlock.Int64()

	for height := from; height <= to; height++ {
		bloom, err := f.backend.BlockBloom(&height)
		if err != nil {
			return nil, err
		}

		filtered, err := f.blockLogs(height, bloom)
		if err != nil {
			return nil, errors.Wrapf(err, "failed to fetch block by number %d", height)
		}


		if len(logs)+len(filtered) > logLimit {
			return nil, errors.Errorf("query returned more than %d results", logLimit)
		}
		logs = append(logs, filtered...)
	}
	return logs, nil
}


func (f *Filter) blockLogs(height int64, bloom ethtypes.Bloom) ([]*ethtypes.Log, error) {
	if !bloomFilter(bloom, f.criteria.Addresses, f.criteria.Topics) {
		return []*ethtypes.Log{}, nil
	}



	logsList, err := f.backend.GetLogsByNumber(types.BlockNumber(height))
	if err != nil {
		return []*ethtypes.Log{}, errors.Wrapf(err, "failed to fetch logs block number %d", height)
	}

	unfiltered := make([]*ethtypes.Log, 0)
	for _, logs := range logsList {
		unfiltered = append(unfiltered, logs...)
	}

	logs := FilterLogs(unfiltered, nil, nil, f.criteria.Addresses, f.criteria.Topics)
	if len(logs) == 0 {
		return []*ethtypes.Log{}, nil
	}

	return logs, nil
}

func createBloomFilters(filters [][][]byte, logger log.Logger) [][]BloomIV {
	bloomFilters := make([][]BloomIV, 0)
	for _, filter := range filters {

		if len(filter) == 0 {
			continue
		}
		bloomIVs := make([]BloomIV, len(filter))




		for i, clause := range filter {
			if clause == nil {
				bloomIVs = nil
				break
			}

			iv, err := calcBloomIVs(clause)
			if err != nil {
				bloomIVs = nil
				logger.Error("calcBloomIVs error: %v", err)
				break
			}

			bloomIVs[i] = iv
		}

		if bloomIVs != nil {
			bloomFilters = append(bloomFilters, bloomIVs)
		}
	}
	return bloomFilters
}



func calcBloomIVs(data []byte) (BloomIV, error) {
	hashbuf := make([]byte, 6)
	biv := BloomIV{}

	sha := crypto.NewKeccakState()
	sha.Reset()
	if _, err := sha.Write(data); err != nil {
		return BloomIV{}, err
	}
	if _, err := sha.Read(hashbuf); err != nil {
		return BloomIV{}, err
	}


	biv.V[0] = byte(1 << (hashbuf[1] & 0x7))
	biv.V[1] = byte(1 << (hashbuf[3] & 0x7))
	biv.V[2] = byte(1 << (hashbuf[5] & 0x7))

	biv.I[0] = ethtypes.BloomByteLength - uint((binary.BigEndian.Uint16(hashbuf)&0x7ff)>>3) - 1
	biv.I[1] = ethtypes.BloomByteLength - uint((binary.BigEndian.Uint16(hashbuf[2:])&0x7ff)>>3) - 1
	biv.I[2] = ethtypes.BloomByteLength - uint((binary.BigEndian.Uint16(hashbuf[4:])&0x7ff)>>3) - 1

	return biv, nil
}
