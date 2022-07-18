package net

import (
	"context"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	rpcclient "github.com/tendermint/tendermint/rpc/client"
	ethermint "github.com/tharsis/ethermint/types"
)


type PublicAPI struct {
	networkVersion uint64
	tmClient       rpcclient.Client
}


func NewPublicAPI(clientCtx client.Context) *PublicAPI {

	chainIDEpoch, err := ethermint.ParseChainID(clientCtx.ChainID)
	if err != nil {
		panic(err)
	}

	return &PublicAPI{
		networkVersion: chainIDEpoch.Uint64(),
		tmClient:       clientCtx.Client,
	}
}


func (s *PublicAPI) Version() string {
	return fmt.Sprintf("%d", s.networkVersion)
}


func (s *PublicAPI) Listening() bool {
	ctx := context.Background()
	netInfo, err := s.tmClient.NetInfo(ctx)
	if err != nil {
		return false
	}
	return netInfo.Listening
}


func (s *PublicAPI) PeerCount() int {
	ctx := context.Background()
	netInfo, err := s.tmClient.NetInfo(ctx)
	if err != nil {
		return 0
	}
	return len(netInfo.Peers)
}
