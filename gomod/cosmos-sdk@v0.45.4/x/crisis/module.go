package crisis

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	"github.com/cosmos/cosmos-sdk/telemetry"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/cosmos/cosmos-sdk/x/crisis/client/cli"
	"github.com/cosmos/cosmos-sdk/x/crisis/keeper"
	"github.com/cosmos/cosmos-sdk/x/crisis/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)


const (
	FlagSkipGenesisInvariants = "x-crisis-skip-assert-invariants"
)


type AppModuleBasic struct{}


func (AppModuleBasic) Name() string {
	return types.ModuleName
}


func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterLegacyAminoCodec(cdc)
}



func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	return cdc.MustMarshalJSON(types.DefaultGenesisState())
}


func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, config client.TxEncodingConfig, bz json.RawMessage) error {
	var data types.GenesisState
	if err := cdc.UnmarshalJSON(bz, &data); err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}

	return types.ValidateGenesis(&data)
}


func (AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}


func (AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {}


func (b AppModuleBasic) GetTxCmd() *cobra.Command {
	return cli.NewTxCmd()
}


func (AppModuleBasic) GetQueryCmd() *cobra.Command { return nil }



func (AppModuleBasic) RegisterInterfaces(registry codectypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}


type AppModule struct {
	AppModuleBasic




	keeper *keeper.Keeper

	skipGenesisInvariants bool
}





func NewAppModule(keeper *keeper.Keeper, skipGenesisInvariants bool) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         keeper,

		skipGenesisInvariants: skipGenesisInvariants,
	}
}


func AddModuleInitFlags(startCmd *cobra.Command) {
	startCmd.Flags().Bool(FlagSkipGenesisInvariants, false, "Skip x/crisis invariants check on startup")
}


func (AppModule) Name() string {
	return types.ModuleName
}


func (AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}


func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(RouterKey, NewHandler(*am.keeper))
}


func (AppModule) QuerierRoute() string { return "" }


func (AppModule) LegacyQuerierHandler(*codec.LegacyAmino) sdk.Querier { return nil }


func (am AppModule) RegisterServices(cfg module.Configurator) {
	types.RegisterMsgServer(cfg.MsgServer(), am.keeper)
}



func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	start := time.Now()
	var genesisState types.GenesisState
	cdc.MustUnmarshalJSON(data, &genesisState)
	telemetry.MeasureSince(start, "InitGenesis", "crisis", "unmarshal")

	am.keeper.InitGenesis(ctx, &genesisState)
	if !am.skipGenesisInvariants {
		am.keeper.AssertInvariants(ctx)
	}
	return []abci.ValidatorUpdate{}
}



func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	gs := am.keeper.ExportGenesis(ctx)
	return cdc.MustMarshalJSON(gs)
}


func (AppModule) ConsensusVersion() uint64 { return 1 }


func (AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}



func (am AppModule) EndBlock(ctx sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	EndBlocker(ctx, *am.keeper)
	return []abci.ValidatorUpdate{}
}
