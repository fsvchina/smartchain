package module

import (
	"context"
	"encoding/json"
	"math/rand"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"

	abci "github.com/tendermint/tendermint/abci/types"

	sdkclient "github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/feegrant"
	"github.com/cosmos/cosmos-sdk/x/feegrant/client/cli"
	"github.com/cosmos/cosmos-sdk/x/feegrant/keeper"
	"github.com/cosmos/cosmos-sdk/x/feegrant/simulation"
)

var (
	_ module.AppModule           = AppModule{}
	_ module.AppModuleBasic      = AppModuleBasic{}
	_ module.AppModuleSimulation = AppModule{}
)






type AppModuleBasic struct {
	cdc codec.Codec
}


func (AppModuleBasic) Name() string {
	return feegrant.ModuleName
}



func (am AppModule) RegisterServices(cfg module.Configurator) {
	feegrant.RegisterMsgServer(cfg.MsgServer(), keeper.NewMsgServerImpl(am.keeper))
	feegrant.RegisterQueryServer(cfg.QueryServer(), am.keeper)
}


func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
}


func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	feegrant.RegisterInterfaces(registry)
}


func (am AppModule) LegacyQuerierHandler(legacyQuerierCdc *codec.LegacyAmino) sdk.Querier {
	return nil
}



func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(feegrant.DefaultGenesisState())
}


func (a AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config sdkclient.TxEncodingConfig, bz json.RawMessage) error {
	var data feegrant.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		sdkerrors.Wrapf(err, "failed to unmarshal %s genesis state", feegrant.ModuleName)
	}

	return feegrant.ValidateGenesis(data)
}


func (AppModuleBasic) RegisterRESTRoutes(ctx sdkclient.Context, rtr *mux.Router) {}


func (AppModuleBasic) RegisterGRPCGatewayRoutes(clientCtx sdkclient.Context, mux *runtime.ServeMux) {
	feegrant.RegisterQueryHandlerClient(context.Background(), mux, feegrant.NewQueryClient(clientCtx))
}


func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.GetTxCmd()
}


func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return cli.GetQueryCmd()
}






type AppModule struct {
	AppModuleBasic
	keeper        keeper.Keeper
	accountKeeper feegrant.AccountKeeper
	bankKeeper    feegrant.BankKeeper
	registry      cdctypes.InterfaceRegistry
}


func NewAppModule(cdc codec.Codec, ak feegrant.AccountKeeper, bk feegrant.BankKeeper, keeper keeper.Keeper, registry cdctypes.InterfaceRegistry) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{cdc: cdc},
		keeper:         keeper,
		accountKeeper:  ak,
		bankKeeper:     bk,
		registry:       registry,
	}
}


func (AppModule) Name() string {
	return feegrant.ModuleName
}


func (am AppModule) RegisterInvariants(ir sdk.InvariantRegistry) {}


func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(feegrant.RouterKey, nil)
}


func (am AppModule) NewHandler() sdk.Handler {
	return nil
}


func (AppModule) QuerierRoute() string {
	return ""
}



func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, bz json.RawMessage) []abci.ValidatorUpdate {
	var gs feegrant.GenesisState
	cdc.MustUnmarshalJSON(bz, &gs)

	err := am.keeper.InitGenesis(ctx, &gs)
	if err != nil {
		panic(err)
	}
	return []abci.ValidatorUpdate{}
}



func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs, err := am.keeper.ExportGenesis(ctx)
	if err != nil {
		panic(err)
	}

	return cdc.MustMarshalJSON(gs)
}


func (AppModule) ConsensusVersion() uint64 { return 1 }


func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}



func (AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}




func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	simulation.RandomizedGenState(simState)
}



func (AppModule) ProposalContents(simState module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}


func (AppModule) RandomizedParams(r *rand.Rand) []simtypes.ParamChange {
	return nil
}


func (am AppModule) RegisterStoreDecoder(sdr sdk.StoreDecoderRegistry) {
	sdr[feegrant.StoreKey] = simulation.NewDecodeStore(am.cdc)
}


func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	return simulation.WeightedOperations(
		simState.AppParams, simState.Cdc, am.accountKeeper, am.bankKeeper, am.keeper,
	)
}
