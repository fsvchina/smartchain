package baseapp

import (
	gocontext "context"
	"fmt"

	gogogrpc "github.com/gogo/protobuf/grpc"
	abci "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/grpc"

	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)





type QueryServiceTestHelper struct {
	*GRPCQueryRouter
	Ctx sdk.Context
}

var (
	_ gogogrpc.Server     = &QueryServiceTestHelper{}
	_ gogogrpc.ClientConn = &QueryServiceTestHelper{}
)



func NewQueryServerTestHelper(ctx sdk.Context, interfaceRegistry types.InterfaceRegistry) *QueryServiceTestHelper {
	qrt := NewGRPCQueryRouter()
	qrt.SetInterfaceRegistry(interfaceRegistry)
	return &QueryServiceTestHelper{GRPCQueryRouter: qrt, Ctx: ctx}
}


func (q *QueryServiceTestHelper) Invoke(_ gocontext.Context, method string, args, reply interface{}, _ ...grpc.CallOption) error {
	querier := q.Route(method)
	if querier == nil {
		return fmt.Errorf("handler not found for %s", method)
	}
	reqBz, err := protoCodec.Marshal(args)
	if err != nil {
		return err
	}

	res, err := querier(q.Ctx, abci.RequestQuery{Data: reqBz})
	if err != nil {
		return err
	}

	err = protoCodec.Unmarshal(res.Value, reply)
	if err != nil {
		return err
	}

	if q.interfaceRegistry != nil {
		return types.UnpackInterfaces(reply, q.interfaceRegistry)
	}

	return nil
}


func (q *QueryServiceTestHelper) NewStream(gocontext.Context, *grpc.StreamDesc, string, ...grpc.CallOption) (grpc.ClientStream, error) {
	return nil, fmt.Errorf("not supported")
}
