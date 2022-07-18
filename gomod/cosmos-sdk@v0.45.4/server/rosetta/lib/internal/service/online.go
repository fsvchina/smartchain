package service

import (
	"context"
	"time"

	"github.com/coinbase/rosetta-sdk-go/types"

	crgerrs "github.com/cosmos/cosmos-sdk/server/rosetta/lib/errors"
	crgtypes "github.com/cosmos/cosmos-sdk/server/rosetta/lib/types"
)


const genesisBlockFetchTimeout = 15 * time.Second



func NewOnlineNetwork(network *types.NetworkIdentifier, client crgtypes.Client) (crgtypes.API, error) {
	ctx, cancel := context.WithTimeout(context.Background(), genesisBlockFetchTimeout)
	defer cancel()

	var genesisHeight int64 = -1
	block, err := client.BlockByHeight(ctx, &genesisHeight)
	if err != nil {
		return OnlineNetwork{}, err
	}

	return OnlineNetwork{
		client:                 client,
		network:                network,
		networkOptions:         networkOptionsFromClient(client, block.Block),
		genesisBlockIdentifier: block.Block,
	}, nil
}


type OnlineNetwork struct {
	client crgtypes.Client

	network        *types.NetworkIdentifier
	networkOptions *types.NetworkOptionsResponse

	genesisBlockIdentifier *types.BlockIdentifier
}



func (o OnlineNetwork) AccountCoins(_ context.Context, _ *types.AccountCoinsRequest) (*types.AccountCoinsResponse, *types.Error) {
	return nil, crgerrs.ToRosetta(crgerrs.ErrOffline)
}


func networkOptionsFromClient(client crgtypes.Client, genesisBlock *types.BlockIdentifier) *types.NetworkOptionsResponse {
	var tsi *int64
	if genesisBlock != nil {
		tsi = &(genesisBlock.Index)
	}
	return &types.NetworkOptionsResponse{
		Version: &types.Version{
			RosettaVersion: crgtypes.SpecVersion,
			NodeVersion:    client.Version(),
		},
		Allow: &types.Allow{
			OperationStatuses:       client.OperationStatuses(),
			OperationTypes:          client.SupportedOperations(),
			Errors:                  crgerrs.SealAndListErrors(),
			HistoricalBalanceLookup: true,
			TimestampStartIndex:     tsi,
		},
	}
}
