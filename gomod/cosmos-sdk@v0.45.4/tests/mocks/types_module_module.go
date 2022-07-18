



package mocks

import (
	json "encoding/json"
	reflect "reflect"

	client "github.com/cosmos/cosmos-sdk/client"
	codec "github.com/cosmos/cosmos-sdk/codec"
	types "github.com/cosmos/cosmos-sdk/codec/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	module "github.com/cosmos/cosmos-sdk/types/module"
	gomock "github.com/golang/mock/gomock"
	mux "github.com/gorilla/mux"
	runtime "github.com/grpc-ecosystem/grpc-gateway/runtime"
	cobra "github.com/spf13/cobra"
	types1 "github.com/tendermint/tendermint/abci/types"
)


type MockAppModuleBasic struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleBasicMockRecorder
}


type MockAppModuleBasicMockRecorder struct {
	mock *MockAppModuleBasic
}


func NewMockAppModuleBasic(ctrl *gomock.Controller) *MockAppModuleBasic {
	mock := &MockAppModuleBasic{ctrl: ctrl}
	mock.recorder = &MockAppModuleBasicMockRecorder{mock}
	return mock
}


func (m *MockAppModuleBasic) EXPECT() *MockAppModuleBasicMockRecorder {
	return m.recorder
}


func (m *MockAppModuleBasic) DefaultGenesis(arg0 codec.JSONCodec) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultGenesis", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}


func (mr *MockAppModuleBasicMockRecorder) DefaultGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultGenesis", reflect.TypeOf((*MockAppModuleBasic)(nil).DefaultGenesis), arg0)
}


func (m *MockAppModuleBasic) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}


func (mr *MockAppModuleBasicMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockAppModuleBasic)(nil).GetQueryCmd))
}


func (m *MockAppModuleBasic) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}


func (mr *MockAppModuleBasicMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockAppModuleBasic)(nil).GetTxCmd))
}


func (m *MockAppModuleBasic) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}


func (mr *MockAppModuleBasicMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppModuleBasic)(nil).Name))
}


func (m *MockAppModuleBasic) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}


func (mr *MockAppModuleBasicMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}


func (m *MockAppModuleBasic) RegisterInterfaces(arg0 types.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}


func (mr *MockAppModuleBasicMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterInterfaces), arg0)
}


func (m *MockAppModuleBasic) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}


func (mr *MockAppModuleBasicMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterLegacyAminoCodec), arg0)
}


func (m *MockAppModuleBasic) RegisterRESTRoutes(arg0 client.Context, arg1 *mux.Router) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRESTRoutes", arg0, arg1)
}


func (mr *MockAppModuleBasicMockRecorder) RegisterRESTRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRESTRoutes", reflect.TypeOf((*MockAppModuleBasic)(nil).RegisterRESTRoutes), arg0, arg1)
}


func (m *MockAppModuleBasic) ValidateGenesis(arg0 codec.JSONCodec, arg1 client.TxEncodingConfig, arg2 json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockAppModuleBasicMockRecorder) ValidateGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGenesis", reflect.TypeOf((*MockAppModuleBasic)(nil).ValidateGenesis), arg0, arg1, arg2)
}


type MockAppModuleGenesis struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleGenesisMockRecorder
}


type MockAppModuleGenesisMockRecorder struct {
	mock *MockAppModuleGenesis
}


func NewMockAppModuleGenesis(ctrl *gomock.Controller) *MockAppModuleGenesis {
	mock := &MockAppModuleGenesis{ctrl: ctrl}
	mock.recorder = &MockAppModuleGenesisMockRecorder{mock}
	return mock
}


func (m *MockAppModuleGenesis) EXPECT() *MockAppModuleGenesisMockRecorder {
	return m.recorder
}


func (m *MockAppModuleGenesis) DefaultGenesis(arg0 codec.JSONCodec) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultGenesis", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}


func (mr *MockAppModuleGenesisMockRecorder) DefaultGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).DefaultGenesis), arg0)
}


func (m *MockAppModuleGenesis) ExportGenesis(arg0 types0.Context, arg1 codec.JSONCodec) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportGenesis", arg0, arg1)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}


func (mr *MockAppModuleGenesisMockRecorder) ExportGenesis(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).ExportGenesis), arg0, arg1)
}


func (m *MockAppModuleGenesis) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}


func (mr *MockAppModuleGenesisMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockAppModuleGenesis)(nil).GetQueryCmd))
}


func (m *MockAppModuleGenesis) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}


func (mr *MockAppModuleGenesisMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockAppModuleGenesis)(nil).GetTxCmd))
}


func (m *MockAppModuleGenesis) InitGenesis(arg0 types0.Context, arg1 codec.JSONCodec, arg2 json.RawMessage) []types1.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types1.ValidatorUpdate)
	return ret0
}


func (mr *MockAppModuleGenesisMockRecorder) InitGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).InitGenesis), arg0, arg1, arg2)
}


func (m *MockAppModuleGenesis) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}


func (mr *MockAppModuleGenesisMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppModuleGenesis)(nil).Name))
}


func (m *MockAppModuleGenesis) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}


func (mr *MockAppModuleGenesisMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}


func (m *MockAppModuleGenesis) RegisterInterfaces(arg0 types.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}


func (mr *MockAppModuleGenesisMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterInterfaces), arg0)
}


func (m *MockAppModuleGenesis) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}


func (mr *MockAppModuleGenesisMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterLegacyAminoCodec), arg0)
}


func (m *MockAppModuleGenesis) RegisterRESTRoutes(arg0 client.Context, arg1 *mux.Router) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRESTRoutes", arg0, arg1)
}


func (mr *MockAppModuleGenesisMockRecorder) RegisterRESTRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRESTRoutes", reflect.TypeOf((*MockAppModuleGenesis)(nil).RegisterRESTRoutes), arg0, arg1)
}


func (m *MockAppModuleGenesis) ValidateGenesis(arg0 codec.JSONCodec, arg1 client.TxEncodingConfig, arg2 json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockAppModuleGenesisMockRecorder) ValidateGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGenesis", reflect.TypeOf((*MockAppModuleGenesis)(nil).ValidateGenesis), arg0, arg1, arg2)
}


type MockAppModule struct {
	ctrl     *gomock.Controller
	recorder *MockAppModuleMockRecorder
}


type MockAppModuleMockRecorder struct {
	mock *MockAppModule
}


func NewMockAppModule(ctrl *gomock.Controller) *MockAppModule {
	mock := &MockAppModule{ctrl: ctrl}
	mock.recorder = &MockAppModuleMockRecorder{mock}
	return mock
}


func (m *MockAppModule) EXPECT() *MockAppModuleMockRecorder {
	return m.recorder
}


func (m *MockAppModule) BeginBlock(arg0 types0.Context, arg1 types1.RequestBeginBlock) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "BeginBlock", arg0, arg1)
}


func (mr *MockAppModuleMockRecorder) BeginBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginBlock", reflect.TypeOf((*MockAppModule)(nil).BeginBlock), arg0, arg1)
}


func (m *MockAppModule) ConsensusVersion() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ConsensusVersion")
	ret0, _ := ret[0].(uint64)
	return ret0
}


func (mr *MockAppModuleMockRecorder) ConsensusVersion() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ConsensusVersion", reflect.TypeOf((*MockAppModule)(nil).ConsensusVersion))
}


func (m *MockAppModule) DefaultGenesis(arg0 codec.JSONCodec) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DefaultGenesis", arg0)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}


func (mr *MockAppModuleMockRecorder) DefaultGenesis(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DefaultGenesis", reflect.TypeOf((*MockAppModule)(nil).DefaultGenesis), arg0)
}


func (m *MockAppModule) EndBlock(arg0 types0.Context, arg1 types1.RequestEndBlock) []types1.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndBlock", arg0, arg1)
	ret0, _ := ret[0].([]types1.ValidatorUpdate)
	return ret0
}


func (mr *MockAppModuleMockRecorder) EndBlock(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndBlock", reflect.TypeOf((*MockAppModule)(nil).EndBlock), arg0, arg1)
}


func (m *MockAppModule) ExportGenesis(arg0 types0.Context, arg1 codec.JSONCodec) json.RawMessage {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ExportGenesis", arg0, arg1)
	ret0, _ := ret[0].(json.RawMessage)
	return ret0
}


func (mr *MockAppModuleMockRecorder) ExportGenesis(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExportGenesis", reflect.TypeOf((*MockAppModule)(nil).ExportGenesis), arg0, arg1)
}


func (m *MockAppModule) GetQueryCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetQueryCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}


func (mr *MockAppModuleMockRecorder) GetQueryCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetQueryCmd", reflect.TypeOf((*MockAppModule)(nil).GetQueryCmd))
}


func (m *MockAppModule) GetTxCmd() *cobra.Command {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTxCmd")
	ret0, _ := ret[0].(*cobra.Command)
	return ret0
}


func (mr *MockAppModuleMockRecorder) GetTxCmd() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTxCmd", reflect.TypeOf((*MockAppModule)(nil).GetTxCmd))
}


func (m *MockAppModule) InitGenesis(arg0 types0.Context, arg1 codec.JSONCodec, arg2 json.RawMessage) []types1.ValidatorUpdate {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].([]types1.ValidatorUpdate)
	return ret0
}


func (mr *MockAppModuleMockRecorder) InitGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitGenesis", reflect.TypeOf((*MockAppModule)(nil).InitGenesis), arg0, arg1, arg2)
}


func (m *MockAppModule) LegacyQuerierHandler(arg0 *codec.LegacyAmino) types0.Querier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LegacyQuerierHandler", arg0)
	ret0, _ := ret[0].(types0.Querier)
	return ret0
}


func (mr *MockAppModuleMockRecorder) LegacyQuerierHandler(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LegacyQuerierHandler", reflect.TypeOf((*MockAppModule)(nil).LegacyQuerierHandler), arg0)
}


func (m *MockAppModule) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}


func (mr *MockAppModuleMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockAppModule)(nil).Name))
}


func (m *MockAppModule) QuerierRoute() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "QuerierRoute")
	ret0, _ := ret[0].(string)
	return ret0
}


func (mr *MockAppModuleMockRecorder) QuerierRoute() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "QuerierRoute", reflect.TypeOf((*MockAppModule)(nil).QuerierRoute))
}


func (m *MockAppModule) RegisterGRPCGatewayRoutes(arg0 client.Context, arg1 *runtime.ServeMux) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterGRPCGatewayRoutes", arg0, arg1)
}


func (mr *MockAppModuleMockRecorder) RegisterGRPCGatewayRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterGRPCGatewayRoutes", reflect.TypeOf((*MockAppModule)(nil).RegisterGRPCGatewayRoutes), arg0, arg1)
}


func (m *MockAppModule) RegisterInterfaces(arg0 types.InterfaceRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInterfaces", arg0)
}


func (mr *MockAppModuleMockRecorder) RegisterInterfaces(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInterfaces", reflect.TypeOf((*MockAppModule)(nil).RegisterInterfaces), arg0)
}


func (m *MockAppModule) RegisterInvariants(arg0 types0.InvariantRegistry) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterInvariants", arg0)
}


func (mr *MockAppModuleMockRecorder) RegisterInvariants(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterInvariants", reflect.TypeOf((*MockAppModule)(nil).RegisterInvariants), arg0)
}


func (m *MockAppModule) RegisterLegacyAminoCodec(arg0 *codec.LegacyAmino) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterLegacyAminoCodec", arg0)
}


func (mr *MockAppModuleMockRecorder) RegisterLegacyAminoCodec(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterLegacyAminoCodec", reflect.TypeOf((*MockAppModule)(nil).RegisterLegacyAminoCodec), arg0)
}


func (m *MockAppModule) RegisterRESTRoutes(arg0 client.Context, arg1 *mux.Router) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRESTRoutes", arg0, arg1)
}


func (mr *MockAppModuleMockRecorder) RegisterRESTRoutes(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRESTRoutes", reflect.TypeOf((*MockAppModule)(nil).RegisterRESTRoutes), arg0, arg1)
}


func (m *MockAppModule) RegisterServices(arg0 module.Configurator) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterServices", arg0)
}


func (mr *MockAppModuleMockRecorder) RegisterServices(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterServices", reflect.TypeOf((*MockAppModule)(nil).RegisterServices), arg0)
}


func (m *MockAppModule) Route() types0.Route {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route")
	ret0, _ := ret[0].(types0.Route)
	return ret0
}


func (mr *MockAppModuleMockRecorder) Route() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockAppModule)(nil).Route))
}


func (m *MockAppModule) ValidateGenesis(arg0 codec.JSONCodec, arg1 client.TxEncodingConfig, arg2 json.RawMessage) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidateGenesis", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockAppModuleMockRecorder) ValidateGenesis(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidateGenesis", reflect.TypeOf((*MockAppModule)(nil).ValidateGenesis), arg0, arg1, arg2)
}
