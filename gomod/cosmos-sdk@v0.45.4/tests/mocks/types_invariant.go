



package mocks

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)


type MockInvariantRegistry struct {
	ctrl     *gomock.Controller
	recorder *MockInvariantRegistryMockRecorder
}


type MockInvariantRegistryMockRecorder struct {
	mock *MockInvariantRegistry
}


func NewMockInvariantRegistry(ctrl *gomock.Controller) *MockInvariantRegistry {
	mock := &MockInvariantRegistry{ctrl: ctrl}
	mock.recorder = &MockInvariantRegistryMockRecorder{mock}
	return mock
}


func (m *MockInvariantRegistry) EXPECT() *MockInvariantRegistryMockRecorder {
	return m.recorder
}


func (m *MockInvariantRegistry) RegisterRoute(moduleName, route string, invar types.Invariant) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterRoute", moduleName, route, invar)
}


func (mr *MockInvariantRegistryMockRecorder) RegisterRoute(moduleName, route, invar interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterRoute", reflect.TypeOf((*MockInvariantRegistry)(nil).RegisterRoute), moduleName, route, invar)
}
