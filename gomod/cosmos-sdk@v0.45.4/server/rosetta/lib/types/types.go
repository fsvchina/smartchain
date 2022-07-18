package types

import (
	"context"

	"github.com/coinbase/rosetta-sdk-go/server"
	"github.com/coinbase/rosetta-sdk-go/types"
)


const SpecVersion = ""



type NetworkInformationProvider interface {

	SupportedOperations() []string

	OperationStatuses() []*types.OperationStatus

	Version() string
}


type Client interface {

	Bootstrap() error




	Ready() error






	Balances(ctx context.Context, addr string, height *int64) ([]*types.Amount, error)

	BlockByHash(ctx context.Context, hash string) (BlockResponse, error)

	BlockByHeight(ctx context.Context, height *int64) (BlockResponse, error)


	BlockTransactionsByHash(ctx context.Context, hash string) (BlockTransactionsResponse, error)


	BlockTransactionsByHeight(ctx context.Context, height *int64) (BlockTransactionsResponse, error)

	GetTx(ctx context.Context, hash string) (*types.Transaction, error)


	GetUnconfirmedTx(ctx context.Context, hash string) (*types.Transaction, error)

	Mempool(ctx context.Context) ([]*types.TransactionIdentifier, error)

	Peers(ctx context.Context) ([]*types.Peer, error)

	Status(ctx context.Context) (*types.SyncStatus, error)





	PostTx(txBytes []byte) (res *types.TransactionIdentifier, meta map[string]interface{}, err error)

	ConstructionMetadataFromOptions(ctx context.Context, options map[string]interface{}) (meta map[string]interface{}, err error)
	OfflineClient
}


type OfflineClient interface {
	NetworkInformationProvider

	SignedTx(ctx context.Context, txBytes []byte, sigs []*types.Signature) (signedTxBytes []byte, err error)


	TxOperationsAndSignersAccountIdentifiers(signed bool, hexBytes []byte) (ops []*types.Operation, signers []*types.AccountIdentifier, err error)

	ConstructionPayload(ctx context.Context, req *types.ConstructionPayloadsRequest) (resp *types.ConstructionPayloadsResponse, err error)

	PreprocessOperationsToOptions(ctx context.Context, req *types.ConstructionPreprocessRequest) (resp *types.ConstructionPreprocessResponse, err error)

	AccountIdentifierFromPublicKey(pubKey *types.PublicKey) (*types.AccountIdentifier, error)
}

type BlockTransactionsResponse struct {
	BlockResponse
	Transactions []*types.Transaction
}

type BlockResponse struct {
	Block                *types.BlockIdentifier
	ParentBlock          *types.BlockIdentifier
	MillisecondTimestamp int64
	TxCount              int64
}



type API interface {
	DataAPI
	ConstructionAPI
}


type DataAPI interface {
	server.NetworkAPIServicer
	server.AccountAPIServicer
	server.BlockAPIServicer
	server.MempoolAPIServicer
}

var _ server.ConstructionAPIServicer = ConstructionAPI(nil)



type ConstructionAPI interface {
	ConstructionOnlineAPI
	ConstructionOfflineAPI
}



type ConstructionOnlineAPI interface {
	ConstructionMetadata(
		context.Context,
		*types.ConstructionMetadataRequest,
	) (*types.ConstructionMetadataResponse, *types.Error)
	ConstructionSubmit(
		context.Context,
		*types.ConstructionSubmitRequest,
	) (*types.TransactionIdentifierResponse, *types.Error)
}



type ConstructionOfflineAPI interface {
	ConstructionCombine(
		context.Context,
		*types.ConstructionCombineRequest,
	) (*types.ConstructionCombineResponse, *types.Error)
	ConstructionDerive(
		context.Context,
		*types.ConstructionDeriveRequest,
	) (*types.ConstructionDeriveResponse, *types.Error)
	ConstructionHash(
		context.Context,
		*types.ConstructionHashRequest,
	) (*types.TransactionIdentifierResponse, *types.Error)
	ConstructionParse(
		context.Context,
		*types.ConstructionParseRequest,
	) (*types.ConstructionParseResponse, *types.Error)
	ConstructionPayloads(
		context.Context,
		*types.ConstructionPayloadsRequest,
	) (*types.ConstructionPayloadsResponse, *types.Error)
	ConstructionPreprocess(
		context.Context,
		*types.ConstructionPreprocessRequest,
	) (*types.ConstructionPreprocessResponse, *types.Error)
}
