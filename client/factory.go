package client

import (
	"bytes"
	"fs.video/smartchain/app"
	"fs.video/smartchain/core"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/cosmos/cosmos-sdk/simapp/params"
	authType "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/spf13/pflag"
	rpchttp "github.com/tendermint/tendermint/rpc/client/http"
	"github.com/tharsis/ethermint/encoding"
)

var clientCtx client.Context

var clientFactory tx.Factory

var encodingConfig params.EncodingConfig

func NewEvmClient() EvmClient {
	return EvmClient{core.EvmRpcURL}
}

func NewTxClient() TxClient {
	return TxClient{core.ServerURL, "TxClient"}
}

func NewBlockClient() BlockClient {
	return BlockClient{core.ServerURL, "sc-BlockClient"}
}

func init() {
	encodingConfig = encoding.MakeConfig(app.ModuleBasics)

	rpcClient, err := rpchttp.New(core.RpcURL, "/websocket")
	if err != nil {
		panic("start ctx client error.")
	}

	clientCtx = client.Context{}.
		WithChainID(core.ChainID).
		WithCodec(encodingConfig.Marshaler).
		WithTxConfig(encodingConfig.TxConfig).
		WithLegacyAmino(encodingConfig.Amino).
		WithOffline(true).
		WithNodeURI(core.RpcURL).
		WithClient(rpcClient).
		WithAccountRetriever(authType.AccountRetriever{})





	flags := pflag.NewFlagSet("chat", pflag.ContinueOnError)

	flagErrorBuf := new(bytes.Buffer)

	flags.SetOutput(flagErrorBuf)


	clientFactory = tx.NewFactoryCLI(clientCtx, flags)
	clientFactory.WithChainID(core.ChainID).
		WithAccountRetriever(clientCtx.AccountRetriever).
		WithTxConfig(clientCtx.TxConfig)
}
