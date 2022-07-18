package client

import (
	"context"
	"fmt"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	abci "github.com/tendermint/tendermint/abci/types"
	tmbytes "github.com/tendermint/tendermint/libs/bytes"
	rpcclient "github.com/tendermint/tendermint/rpc/client"

	"github.com/cosmos/cosmos-sdk/store/rootmulti"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)



func (ctx Context) GetNode() (rpcclient.Client, error) {
	if ctx.Client == nil {
		return nil, errors.New("no RPC client is defined in offline mode")
	}

	return ctx.Client, nil
}




func (ctx Context) Query(path string) ([]byte, int64, error) {
	return ctx.query(path, nil)
}




func (ctx Context) QueryWithData(path string, data []byte) ([]byte, int64, error) {
	return ctx.query(path, data)
}




func (ctx Context) QueryStore(key tmbytes.HexBytes, storeName string) ([]byte, int64, error) {
	return ctx.queryStore(key, storeName, "key")
}





func (ctx Context) QueryABCI(req abci.RequestQuery) (abci.ResponseQuery, error) {
	return ctx.queryABCI(req)
}


func (ctx Context) GetFromAddress() sdk.AccAddress {
	return ctx.FromAddress
}


func (ctx Context) GetFeeGranterAddress() sdk.AccAddress {
	return ctx.FeeGranter
}


func (ctx Context) GetFromName() string {
	return ctx.FromName
}

func (ctx Context) queryABCI(req abci.RequestQuery) (abci.ResponseQuery, error) {
	node, err := ctx.GetNode()
	if err != nil {
		return abci.ResponseQuery{}, err
	}

	var queryHeight int64
	if req.Height != 0 {
		queryHeight = req.Height
	} else {

		queryHeight = ctx.Height
	}

	opts := rpcclient.ABCIQueryOptions{
		Height: queryHeight,
		Prove:  req.Prove,
	}

	result, err := node.ABCIQueryWithOptions(context.Background(), req.Path, req.Data, opts)
	if err != nil {
		return abci.ResponseQuery{}, err
	}

	if !result.Response.IsOK() {
		return abci.ResponseQuery{}, sdkErrorToGRPCError(result.Response)
	}


	if !opts.Prove || !isQueryStoreWithProof(req.Path) {
		return result.Response, nil
	}

	return result.Response, nil
}

func sdkErrorToGRPCError(resp abci.ResponseQuery) error {
	switch resp.Code {
	case sdkerrors.ErrInvalidRequest.ABCICode():
		return status.Error(codes.InvalidArgument, resp.Log)
	case sdkerrors.ErrUnauthorized.ABCICode():
		return status.Error(codes.Unauthenticated, resp.Log)
	case sdkerrors.ErrKeyNotFound.ABCICode():
		return status.Error(codes.NotFound, resp.Log)
	default:
		return status.Error(codes.Unknown, resp.Log)
	}
}




func (ctx Context) query(path string, key tmbytes.HexBytes) ([]byte, int64, error) {
	resp, err := ctx.queryABCI(abci.RequestQuery{
		Path:   path,
		Data:   key,
		Height: ctx.Height,
	})
	if err != nil {
		return nil, 0, err
	}

	return resp.Value, resp.Height, nil
}




func (ctx Context) queryStore(key tmbytes.HexBytes, storeName, endPath string) ([]byte, int64, error) {
	path := fmt.Sprintf("/store/%s/%s", storeName, endPath)
	return ctx.query(path, key)
}



func isQueryStoreWithProof(path string) bool {
	if !strings.HasPrefix(path, "/") {
		return false
	}

	paths := strings.SplitN(path[1:], "/", 3)

	switch {
	case len(paths) != 3:
		return false
	case paths[0] != "store":
		return false
	case rootmulti.RequireProof("/" + paths[2]):
		return true
	}

	return false
}
