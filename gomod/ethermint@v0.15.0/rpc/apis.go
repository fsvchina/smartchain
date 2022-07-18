

package rpc

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"

	"github.com/ethereum/go-ethereum/rpc"

	"github.com/tharsis/ethermint/rpc/backend"
	"github.com/tharsis/ethermint/rpc/namespaces/ethereum/debug"
	"github.com/tharsis/ethermint/rpc/namespaces/ethereum/eth"
	"github.com/tharsis/ethermint/rpc/namespaces/ethereum/eth/filters"
	"github.com/tharsis/ethermint/rpc/namespaces/ethereum/miner"
	"github.com/tharsis/ethermint/rpc/namespaces/ethereum/net"
	"github.com/tharsis/ethermint/rpc/namespaces/ethereum/personal"
	"github.com/tharsis/ethermint/rpc/namespaces/ethereum/txpool"
	"github.com/tharsis/ethermint/rpc/namespaces/ethereum/web3"
	"github.com/tharsis/ethermint/rpc/types"

	rpcclient "github.com/tendermint/tendermint/rpc/jsonrpc/client"
)


const (


	CosmosNamespace = "cosmos"



	Web3Namespace     = "web3"
	EthNamespace      = "eth"
	PersonalNamespace = "personal"
	NetNamespace      = "net"
	TxPoolNamespace   = "txpool"
	DebugNamespace    = "debug"
	MinerNamespace    = "miner"

	apiVersion = "1.0"
)


type APICreator = func(*server.Context, client.Context, *rpcclient.WSClient) []rpc.API


var apiCreators map[string]APICreator

func init() {
	apiCreators = map[string]APICreator{
		EthNamespace: func(ctx *server.Context, clientCtx client.Context, tmWSClient *rpcclient.WSClient) []rpc.API {
			nonceLock := new(types.AddrLocker)
			evmBackend := backend.NewBackend(ctx, ctx.Logger, clientCtx)
			return []rpc.API{
				{
					Namespace: EthNamespace,
					Version:   apiVersion,
					Service:   eth.NewPublicAPI(ctx.Logger, clientCtx, evmBackend, nonceLock),
					Public:    true,
				},
				{
					Namespace: EthNamespace,
					Version:   apiVersion,
					Service:   filters.NewPublicAPI(ctx.Logger, clientCtx, tmWSClient, evmBackend),
					Public:    true,
				},
			}
		},
		Web3Namespace: func(*server.Context, client.Context, *rpcclient.WSClient) []rpc.API {
			return []rpc.API{
				{
					Namespace: Web3Namespace,
					Version:   apiVersion,
					Service:   web3.NewPublicAPI(),
					Public:    true,
				},
			}
		},
		NetNamespace: func(_ *server.Context, clientCtx client.Context, _ *rpcclient.WSClient) []rpc.API {
			return []rpc.API{
				{
					Namespace: NetNamespace,
					Version:   apiVersion,
					Service:   net.NewPublicAPI(clientCtx),
					Public:    true,
				},
			}
		},
		PersonalNamespace: func(ctx *server.Context, clientCtx client.Context, _ *rpcclient.WSClient) []rpc.API {
			evmBackend := backend.NewBackend(ctx, ctx.Logger, clientCtx)
			return []rpc.API{
				{
					Namespace: PersonalNamespace,
					Version:   apiVersion,
					Service:   personal.NewAPI(ctx.Logger, clientCtx, evmBackend),
					Public:    false,
				},
			}
		},
		TxPoolNamespace: func(ctx *server.Context, _ client.Context, _ *rpcclient.WSClient) []rpc.API {
			return []rpc.API{
				{
					Namespace: TxPoolNamespace,
					Version:   apiVersion,
					Service:   txpool.NewPublicAPI(ctx.Logger),
					Public:    true,
				},
			}
		},
		DebugNamespace: func(ctx *server.Context, clientCtx client.Context, _ *rpcclient.WSClient) []rpc.API {
			evmBackend := backend.NewBackend(ctx, ctx.Logger, clientCtx)
			return []rpc.API{
				{
					Namespace: DebugNamespace,
					Version:   apiVersion,
					Service:   debug.NewAPI(ctx, evmBackend, clientCtx),
					Public:    true,
				},
			}
		},
		MinerNamespace: func(ctx *server.Context, clientCtx client.Context, _ *rpcclient.WSClient) []rpc.API {
			evmBackend := backend.NewBackend(ctx, ctx.Logger, clientCtx)
			return []rpc.API{
				{
					Namespace: MinerNamespace,
					Version:   apiVersion,
					Service:   miner.NewPrivateAPI(ctx, clientCtx, evmBackend),
					Public:    false,
				},
			}
		},
	}
}


func GetRPCAPIs(ctx *server.Context, clientCtx client.Context, tmWSClient *rpcclient.WSClient, selectedAPIs []string) []rpc.API {
	var apis []rpc.API

	for _, ns := range selectedAPIs {
		if creator, ok := apiCreators[ns]; ok {
			apis = append(apis, creator(ctx, clientCtx, tmWSClient)...)
		} else {
			ctx.Logger.Error("invalid namespace value", "namespace", ns)
		}
	}

	return apis
}



func RegisterAPINamespace(ns string, creator APICreator) error {
	if _, ok := apiCreators[ns]; ok {
		return fmt.Errorf("duplicated api namespace %s", ns)
	}
	apiCreators[ns] = creator
	return nil
}