package service

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/types"

	"github.com/cosmos/cosmos-sdk/server/rosetta/lib/errors"
	crgtypes "github.com/cosmos/cosmos-sdk/server/rosetta/lib/types"
)



func (on OnlineNetwork) AccountBalance(ctx context.Context, request *types.AccountBalanceRequest) (*types.AccountBalanceResponse, *types.Error) {
	var (
		height int64
		block  crgtypes.BlockResponse
		err    error
	)

	switch {
	case request.BlockIdentifier == nil:
		block, err = on.client.BlockByHeight(ctx, nil)
		if err != nil {
			return nil, errors.ToRosetta(err)
		}
	case request.BlockIdentifier.Hash != nil:
		block, err = on.client.BlockByHash(ctx, *request.BlockIdentifier.Hash)
		if err != nil {
			return nil, errors.ToRosetta(err)
		}
		height = block.Block.Index
	case request.BlockIdentifier.Index != nil:
		height = *request.BlockIdentifier.Index
		block, err = on.client.BlockByHeight(ctx, &height)
		if err != nil {
			return nil, errors.ToRosetta(err)
		}
	}

	accountCoins, err := on.client.Balances(ctx, request.AccountIdentifier.Address, &height)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.AccountBalanceResponse{
		BlockIdentifier: block.Block,
		Balances:        accountCoins,
		Metadata:        nil,
	}, nil
}


func (on OnlineNetwork) Block(ctx context.Context, request *types.BlockRequest) (*types.BlockResponse, *types.Error) {
	var (
		blockResponse crgtypes.BlockTransactionsResponse
		err           error
	)


	switch {
	case request.BlockIdentifier.Hash != nil:
		blockResponse, err = on.client.BlockTransactionsByHash(ctx, *request.BlockIdentifier.Hash)
		if err != nil {
			return nil, errors.ToRosetta(err)
		}
	case request.BlockIdentifier.Index != nil:
		blockResponse, err = on.client.BlockTransactionsByHeight(ctx, request.BlockIdentifier.Index)
		if err != nil {
			return nil, errors.ToRosetta(err)
		}
	default:
		err := errors.WrapError(errors.ErrBadArgument, "at least one of hash or index needs to be specified")
		return nil, errors.ToRosetta(err)
	}

	return &types.BlockResponse{
		Block: &types.Block{
			BlockIdentifier:       blockResponse.Block,
			ParentBlockIdentifier: blockResponse.ParentBlock,
			Timestamp:             blockResponse.MillisecondTimestamp,
			Transactions:          blockResponse.Transactions,
			Metadata:              nil,
		},
		OtherTransactions: nil,
	}, nil
}



func (on OnlineNetwork) BlockTransaction(ctx context.Context, request *types.BlockTransactionRequest) (*types.BlockTransactionResponse, *types.Error) {
	tx, err := on.client.GetTx(ctx, request.TransactionIdentifier.Hash)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.BlockTransactionResponse{
		Transaction: tx,
	}, nil
}


func (on OnlineNetwork) Mempool(ctx context.Context, _ *types.NetworkRequest) (*types.MempoolResponse, *types.Error) {
	txs, err := on.client.Mempool(ctx)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.MempoolResponse{
		TransactionIdentifiers: txs,
	}, nil
}



func (on OnlineNetwork) MempoolTransaction(ctx context.Context, request *types.MempoolTransactionRequest) (*types.MempoolTransactionResponse, *types.Error) {
	tx, err := on.client.GetUnconfirmedTx(ctx, request.TransactionIdentifier.Hash)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.MempoolTransactionResponse{
		Transaction: tx,
	}, nil
}

func (on OnlineNetwork) NetworkList(_ context.Context, _ *types.MetadataRequest) (*types.NetworkListResponse, *types.Error) {
	return &types.NetworkListResponse{NetworkIdentifiers: []*types.NetworkIdentifier{on.network}}, nil
}

func (on OnlineNetwork) NetworkOptions(_ context.Context, _ *types.NetworkRequest) (*types.NetworkOptionsResponse, *types.Error) {
	return on.networkOptions, nil
}

func (on OnlineNetwork) NetworkStatus(ctx context.Context, _ *types.NetworkRequest) (*types.NetworkStatusResponse, *types.Error) {
	block, err := on.client.BlockByHeight(ctx, nil)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	peers, err := on.client.Peers(ctx)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	syncStatus, err := on.client.Status(ctx)
	if err != nil {
		return nil, errors.ToRosetta(err)
	}

	return &types.NetworkStatusResponse{
		CurrentBlockIdentifier: block.Block,
		CurrentBlockTimestamp:  block.MillisecondTimestamp,
		GenesisBlockIdentifier: on.genesisBlockIdentifier,
		OldestBlockIdentifier:  nil,
		SyncStatus:             syncStatus,
		Peers:                  peers,
	}, nil
}
