package rosetta

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"time"

	"github.com/cosmos/cosmos-sdk/version"

	abcitypes "github.com/tendermint/tendermint/abci/types"

	rosettatypes "github.com/coinbase/rosetta-sdk-go/types"
	"google.golang.org/grpc/metadata"

	"github.com/tendermint/tendermint/rpc/client/http"
	"google.golang.org/grpc"

	crgerrs "github.com/cosmos/cosmos-sdk/server/rosetta/lib/errors"
	crgtypes "github.com/cosmos/cosmos-sdk/server/rosetta/lib/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
	grpctypes "github.com/cosmos/cosmos-sdk/types/grpc"
	authtx "github.com/cosmos/cosmos-sdk/x/auth/tx"
	auth "github.com/cosmos/cosmos-sdk/x/auth/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"

	tmrpc "github.com/tendermint/tendermint/rpc/client"
)


var _ crgtypes.Client = (*Client)(nil)

const tmWebsocketPath = "/websocket"
const defaultNodeTimeout = 15 * time.Second


type Client struct {
	supportedOperations []string

	config *Config

	auth  auth.QueryClient
	bank  bank.QueryClient
	tmRPC tmrpc.Client

	version string

	converter Converter
}


func NewClient(cfg *Config) (*Client, error) {
	info := version.NewInfo()

	v := info.Version
	if v == "" {
		v = "unknown"
	}

	txConfig := authtx.NewTxConfig(cfg.Codec, authtx.DefaultSignModes)

	var supportedOperations []string
	for _, ii := range cfg.InterfaceRegistry.ListImplementations(sdk.MsgInterfaceProtoName) {
		resolvedMsg, err := cfg.InterfaceRegistry.Resolve(ii)
		if err != nil {
			continue
		}

		if _, ok := resolvedMsg.(sdk.Msg); ok {
			supportedOperations = append(supportedOperations, ii)
		}
	}

	supportedOperations = append(
		supportedOperations,
		bank.EventTypeCoinSpent,
		bank.EventTypeCoinReceived,
		bank.EventTypeCoinBurn,
	)

	return &Client{
		supportedOperations: supportedOperations,
		config:              cfg,
		auth:                nil,
		bank:                nil,
		tmRPC:               nil,
		version:             fmt.Sprintf("%s/%s", info.AppName, v),
		converter:           NewConverter(cfg.Codec, cfg.InterfaceRegistry, txConfig),
	}, nil
}




func (c *Client) Bootstrap() error {
	grpcConn, err := grpc.Dial(c.config.GRPCEndpoint, grpc.WithInsecure())
	if err != nil {
		return err
	}

	tmRPC, err := http.New(c.config.TendermintRPC, tmWebsocketPath)
	if err != nil {
		return err
	}

	authClient := auth.NewQueryClient(grpcConn)
	bankClient := bank.NewQueryClient(grpcConn)

	c.auth = authClient
	c.bank = bankClient
	c.tmRPC = tmRPC

	return nil
}


func (c *Client) Ready() error {
	ctx, cancel := context.WithTimeout(context.Background(), defaultNodeTimeout)
	defer cancel()
	_, err := c.tmRPC.Health(ctx)
	if err != nil {
		return err
	}
	_, err = c.bank.TotalSupply(ctx, &bank.QueryTotalSupplyRequest{})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) accountInfo(ctx context.Context, addr string, height *int64) (*SignerData, error) {
	if height != nil {
		strHeight := strconv.FormatInt(*height, 10)
		ctx = metadata.AppendToOutgoingContext(ctx, grpctypes.GRPCBlockHeightHeader, strHeight)
	}

	accountInfo, err := c.auth.Account(ctx, &auth.QueryAccountRequest{
		Address: addr,
	})
	if err != nil {
		return nil, crgerrs.FromGRPCToRosettaError(err)
	}

	signerData, err := c.converter.ToRosetta().SignerData(accountInfo.Account)
	if err != nil {
		return nil, err
	}
	return signerData, nil
}

func (c *Client) Balances(ctx context.Context, addr string, height *int64) ([]*rosettatypes.Amount, error) {
	if height != nil {
		strHeight := strconv.FormatInt(*height, 10)
		ctx = metadata.AppendToOutgoingContext(ctx, grpctypes.GRPCBlockHeightHeader, strHeight)
	}

	balance, err := c.bank.AllBalances(ctx, &bank.QueryAllBalancesRequest{
		Address: addr,
	})
	if err != nil {
		return nil, crgerrs.FromGRPCToRosettaError(err)
	}

	availableCoins, err := c.coins(ctx)
	if err != nil {
		return nil, err
	}

	return c.converter.ToRosetta().Amounts(balance.Balances, availableCoins), nil
}

func (c *Client) BlockByHash(ctx context.Context, hash string) (crgtypes.BlockResponse, error) {
	bHash, err := hex.DecodeString(hash)
	if err != nil {
		return crgtypes.BlockResponse{}, fmt.Errorf("invalid block hash: %s", err)
	}

	block, err := c.tmRPC.BlockByHash(ctx, bHash)
	if err != nil {
		return crgtypes.BlockResponse{}, crgerrs.WrapError(crgerrs.ErrBadGateway, err.Error())
	}

	return c.converter.ToRosetta().BlockResponse(block), nil
}

func (c *Client) BlockByHeight(ctx context.Context, height *int64) (crgtypes.BlockResponse, error) {
	height, err := c.getHeight(ctx, height)
	if err != nil {
		return crgtypes.BlockResponse{}, crgerrs.WrapError(crgerrs.ErrBadGateway, err.Error())
	}
	block, err := c.tmRPC.Block(ctx, height)
	if err != nil {
		return crgtypes.BlockResponse{}, crgerrs.WrapError(crgerrs.ErrBadGateway, err.Error())
	}

	return c.converter.ToRosetta().BlockResponse(block), nil
}

func (c *Client) BlockTransactionsByHash(ctx context.Context, hash string) (crgtypes.BlockTransactionsResponse, error) {

	blockResp, err := c.BlockByHash(ctx, hash)
	if err != nil {
		return crgtypes.BlockTransactionsResponse{}, err
	}

	return c.blockTxs(ctx, &blockResp.Block.Index)
}

func (c *Client) BlockTransactionsByHeight(ctx context.Context, height *int64) (crgtypes.BlockTransactionsResponse, error) {
	height, err := c.getHeight(ctx, height)
	if err != nil {
		return crgtypes.BlockTransactionsResponse{}, crgerrs.WrapError(crgerrs.ErrBadGateway, err.Error())
	}
	blockTxResp, err := c.blockTxs(ctx, height)
	if err != nil {
		return crgtypes.BlockTransactionsResponse{}, err
	}
	return blockTxResp, nil
}


func (c *Client) coins(ctx context.Context) (sdk.Coins, error) {
	supply, err := c.bank.TotalSupply(ctx, &bank.QueryTotalSupplyRequest{})
	if err != nil {
		return nil, crgerrs.FromGRPCToRosettaError(err)
	}
	return supply.Supply, nil
}

func (c *Client) TxOperationsAndSignersAccountIdentifiers(signed bool, txBytes []byte) (ops []*rosettatypes.Operation, signers []*rosettatypes.AccountIdentifier, err error) {
	switch signed {
	case false:
		rosTx, err := c.converter.ToRosetta().Tx(txBytes, nil)
		if err != nil {
			return nil, nil, err
		}
		return rosTx.Operations, nil, err
	default:
		ops, signers, err = c.converter.ToRosetta().OpsAndSigners(txBytes)
		return
	}
}



func (c *Client) GetTx(ctx context.Context, hash string) (*rosettatypes.Transaction, error) {
	hashBytes, err := hex.DecodeString(hash)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrCodec, fmt.Sprintf("bad tx hash: %s", err))
	}


	txType, hashBytes := c.converter.ToSDK().HashToTxType(hashBytes)


	switch txType {

	case BeginBlockTx:

		block, err := c.tmRPC.BlockByHash(ctx, hashBytes)
		if err != nil {
			return nil, crgerrs.WrapError(crgerrs.ErrUnknown, err.Error())
		}


		fullBlock, err := c.blockTxs(ctx, &block.Block.Height)
		if err != nil {
			return nil, err
		}

		return fullBlock.Transactions[0], nil

	case DeliverTxTx:
		rawTx, err := c.tmRPC.Tx(ctx, hashBytes, true)
		if err != nil {
			return nil, crgerrs.WrapError(crgerrs.ErrUnknown, err.Error())
		}
		return c.converter.ToRosetta().Tx(rawTx.Tx, &rawTx.TxResult)

	case EndBlockTx:

		block, err := c.tmRPC.BlockByHash(ctx, hashBytes)
		if err != nil {
			return nil, crgerrs.WrapError(crgerrs.ErrUnknown, err.Error())
		}


		fullBlock, err := c.blockTxs(ctx, &block.Block.Height)
		if err != nil {
			return nil, err
		}


		return fullBlock.Transactions[len(fullBlock.Transactions)-1], nil

	default:
		return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, fmt.Sprintf("invalid tx hash provided: %s", hash))
	}
}


func (c *Client) GetUnconfirmedTx(ctx context.Context, hash string) (*rosettatypes.Transaction, error) {
	res, err := c.tmRPC.UnconfirmedTxs(ctx, nil)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrNotFound, "unconfirmed tx not found")
	}

	hashAsBytes, err := hex.DecodeString(hash)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrInterpreting, "invalid hash")
	}


	switch len(hashAsBytes) {
	default:
		return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, fmt.Sprintf("unrecognized tx size: %d", len(hashAsBytes)))
	case BeginEndBlockTxSize:
		return nil, crgerrs.WrapError(crgerrs.ErrBadArgument, "endblock and begin block txs cannot be unconfirmed")
	case DeliverTxSize:
		break
	}


	for _, unconfirmedTx := range res.Txs {
		if !bytes.Equal(unconfirmedTx.Hash(), hashAsBytes) {
			continue
		}

		return c.converter.ToRosetta().Tx(unconfirmedTx, nil)
	}
	return nil, crgerrs.WrapError(crgerrs.ErrNotFound, "transaction not found in mempool: "+hash)
}


func (c *Client) Mempool(ctx context.Context) ([]*rosettatypes.TransactionIdentifier, error) {
	txs, err := c.tmRPC.UnconfirmedTxs(ctx, nil)
	if err != nil {
		return nil, err
	}

	return c.converter.ToRosetta().TxIdentifiers(txs.Txs), nil
}


func (c *Client) Peers(ctx context.Context) ([]*rosettatypes.Peer, error) {
	netInfo, err := c.tmRPC.NetInfo(ctx)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrUnknown, err.Error())
	}
	return c.converter.ToRosetta().Peers(netInfo.Peers), nil
}

func (c *Client) Status(ctx context.Context) (*rosettatypes.SyncStatus, error) {
	status, err := c.tmRPC.Status(ctx)
	if err != nil {
		return nil, crgerrs.WrapError(crgerrs.ErrUnknown, err.Error())
	}
	return c.converter.ToRosetta().SyncStatus(status), err
}

func (c *Client) PostTx(txBytes []byte) (*rosettatypes.TransactionIdentifier, map[string]interface{}, error) {

	res, err := c.tmRPC.BroadcastTxSync(context.Background(), txBytes)
	if err != nil {
		return nil, nil, crgerrs.WrapError(crgerrs.ErrUnknown, err.Error())
	}

	if res.Code != abcitypes.CodeTypeOK {
		return nil, nil, crgerrs.WrapError(
			crgerrs.ErrUnknown,
			fmt.Sprintf("transaction broadcast failure: (%d) %s ", res.Code, res.Log),
		)
	}

	return &rosettatypes.TransactionIdentifier{
			Hash: fmt.Sprintf("%X", res.Hash),
		},
		map[string]interface{}{
			Log: res.Log,
		}, nil
}




func (c *Client) ConstructionMetadataFromOptions(ctx context.Context, options map[string]interface{}) (meta map[string]interface{}, err error) {
	if len(options) == 0 {
		return nil, crgerrs.ErrBadArgument
	}

	constructionOptions := new(PreprocessOperationsOptionsResponse)

	err = constructionOptions.FromMetadata(options)
	if err != nil {
		return nil, err
	}

	signersData := make([]*SignerData, len(constructionOptions.ExpectedSigners))

	for i, signer := range constructionOptions.ExpectedSigners {
		accountInfo, err := c.accountInfo(ctx, signer, nil)
		if err != nil {
			return nil, err
		}

		signersData[i] = accountInfo
	}

	status, err := c.tmRPC.Status(ctx)
	if err != nil {
		return nil, err
	}

	metadataResp := ConstructionMetadata{
		ChainID:     status.NodeInfo.Network,
		SignersData: signersData,
		GasLimit:    constructionOptions.GasLimit,
		GasPrice:    constructionOptions.GasPrice,
		Memo:        constructionOptions.Memo,
	}

	return metadataResp.ToMetadata()
}

func (c *Client) blockTxs(ctx context.Context, height *int64) (crgtypes.BlockTransactionsResponse, error) {

	blockInfo, err := c.tmRPC.Block(ctx, height)
	if err != nil {
		return crgtypes.BlockTransactionsResponse{}, err
	}

	blockResults, err := c.tmRPC.BlockResults(ctx, height)
	if err != nil {
		return crgtypes.BlockTransactionsResponse{}, err
	}

	if len(blockResults.TxsResults) != len(blockInfo.Block.Txs) {

		panic("block results transactions do now match block transactions")
	}

	beginBlockTx := &rosettatypes.Transaction{
		TransactionIdentifier: &rosettatypes.TransactionIdentifier{Hash: c.converter.ToRosetta().BeginBlockTxHash(blockInfo.BlockID.Hash)},
		Operations: AddOperationIndexes(
			nil,
			c.converter.ToRosetta().BalanceOps(StatusTxSuccess, blockResults.BeginBlockEvents),
		),
	}

	endBlockTx := &rosettatypes.Transaction{
		TransactionIdentifier: &rosettatypes.TransactionIdentifier{Hash: c.converter.ToRosetta().EndBlockTxHash(blockInfo.BlockID.Hash)},
		Operations: AddOperationIndexes(
			nil,
			c.converter.ToRosetta().BalanceOps(StatusTxSuccess, blockResults.EndBlockEvents),
		),
	}

	deliverTx := make([]*rosettatypes.Transaction, len(blockInfo.Block.Txs))

	for i, tx := range blockInfo.Block.Txs {
		rosTx, err := c.converter.ToRosetta().Tx(tx, blockResults.TxsResults[i])
		if err != nil {
			return crgtypes.BlockTransactionsResponse{}, err
		}
		deliverTx[i] = rosTx
	}

	finalTxs := make([]*rosettatypes.Transaction, 0, 2+len(deliverTx))
	finalTxs = append(finalTxs, beginBlockTx)
	finalTxs = append(finalTxs, deliverTx...)
	finalTxs = append(finalTxs, endBlockTx)

	return crgtypes.BlockTransactionsResponse{
		BlockResponse: c.converter.ToRosetta().BlockResponse(blockInfo),
		Transactions:  finalTxs,
	}, nil
}

func (c *Client) getHeight(ctx context.Context, height *int64) (realHeight *int64, err error) {
	if height != nil && *height == -1 {
		genesisChunk, err := c.tmRPC.GenesisChunked(ctx, 0)
		if err != nil {
			return nil, err
		}

		heightNum, err := extractInitialHeightFromGenesisChunk(genesisChunk.Data)
		if err != nil {
			return nil, err
		}

		realHeight = &heightNum
	} else {
		realHeight = height
	}
	return
}

func extractInitialHeightFromGenesisChunk(genesisChunk string) (int64, error) {
	firstChunk, err := base64.StdEncoding.DecodeString(genesisChunk)
	if err != nil {
		return 0, err
	}

	re, err := regexp.Compile("\"initial_height\":\"(\\d+)\"")
	if err != nil {
		return 0, err
	}

	matches := re.FindStringSubmatch(string(firstChunk))
	if len(matches) != 2 {
		return 0, errors.New("failed to fetch initial_height")
	}

	heightStr := matches[1]
	return strconv.ParseInt(heightStr, 10, 64)
}
