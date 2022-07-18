package baseapp

import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/grpc/reflection"

	gogogrpc "github.com/gogo/protobuf/grpc"
	abci "github.com/tendermint/tendermint/abci/types"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding"
	"google.golang.org/grpc/encoding/proto"

	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var protoCodec = encoding.GetCodec(proto.Name)


type GRPCQueryRouter struct {
	routes            map[string]GRPCQueryHandler
	interfaceRegistry codectypes.InterfaceRegistry
	serviceData       []serviceData
}


type serviceData struct {
	serviceDesc *grpc.ServiceDesc
	handler     interface{}
}

var _ gogogrpc.Server = &GRPCQueryRouter{}


func NewGRPCQueryRouter() *GRPCQueryRouter {
	return &GRPCQueryRouter{
		routes: map[string]GRPCQueryHandler{},
	}
}



type GRPCQueryHandler = func(ctx sdk.Context, req abci.RequestQuery) (abci.ResponseQuery, error)



func (qrt *GRPCQueryRouter) Route(path string) GRPCQueryHandler {
	handler, found := qrt.routes[path]
	if !found {
		return nil
	}
	return handler
}



//


func (qrt *GRPCQueryRouter) RegisterService(sd *grpc.ServiceDesc, handler interface{}) {

	for _, method := range sd.Methods {
		fqName := fmt.Sprintf("/%s/%s", sd.ServiceName, method.MethodName)
		methodHandler := method.Handler





		_, found := qrt.routes[fqName]
		if found {
			panic(
				fmt.Errorf(
					"gRPC query service %s has already been registered. Please make sure to only register each service once. "+
						"This usually means that there are conflicting modules registering the same gRPC query service",
					fqName,
				),
			)
		}

		qrt.routes[fqName] = func(ctx sdk.Context, req abci.RequestQuery) (abci.ResponseQuery, error) {


			res, err := methodHandler(handler, sdk.WrapSDKContext(ctx), func(i interface{}) error {
				err := protoCodec.Unmarshal(req.Data, i)
				if err != nil {
					return err
				}
				if qrt.interfaceRegistry != nil {
					return codectypes.UnpackInterfaces(i, qrt.interfaceRegistry)
				}
				return nil
			}, nil)
			if err != nil {
				return abci.ResponseQuery{}, err
			}


			resBytes, err := protoCodec.Marshal(res)
			if err != nil {
				return abci.ResponseQuery{}, err
			}


			return abci.ResponseQuery{
				Height: req.Height,
				Value:  resBytes,
			}, nil
		}
	}

	qrt.serviceData = append(qrt.serviceData, serviceData{
		serviceDesc: sd,
		handler:     handler,
	})
}



func (qrt *GRPCQueryRouter) SetInterfaceRegistry(interfaceRegistry codectypes.InterfaceRegistry) {
	qrt.interfaceRegistry = interfaceRegistry


	reflection.RegisterReflectionServiceServer(
		qrt,
		reflection.NewReflectionServiceServer(interfaceRegistry),
	)
}
