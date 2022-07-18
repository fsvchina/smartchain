



package mocks

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)


type MockRouter struct {
	ctrl     *gomock.Controller
	recorder *MockRouterMockRecorder
}


type MockRouterMockRecorder struct {
	mock *MockRouter
}


func NewMockRouter(ctrl *gomock.Controller) *MockRouter {
	mock := &MockRouter{ctrl: ctrl}
	mock.recorder = &MockRouterMockRecorder{mock}
	return mock
}


func (m *MockRouter) EXPECT() *MockRouterMockRecorder {
	return m.recorder
}


func (m *MockRouter) AddRoute(r types.Route) types.Router {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", r)
	ret0, _ := ret[0].(types.Router)
	return ret0
}


func (mr *MockRouterMockRecorder) AddRoute(r interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockRouter)(nil).AddRoute), r)
}


func (m *MockRouter) Route(ctx types.Context, path string) types.Handler {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", ctx, path)
	ret0, _ := ret[0].(types.Handler)
	return ret0
}


func (mr *MockRouterMockRecorder) Route(ctx, path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockRouter)(nil).Route), ctx, path)
}


type MockQueryRouter struct {
	ctrl     *gomock.Controller
	recorder *MockQueryRouterMockRecorder
}


type MockQueryRouterMockRecorder struct {
	mock *MockQueryRouter
}


func NewMockQueryRouter(ctrl *gomock.Controller) *MockQueryRouter {
	mock := &MockQueryRouter{ctrl: ctrl}
	mock.recorder = &MockQueryRouterMockRecorder{mock}
	return mock
}


func (m *MockQueryRouter) EXPECT() *MockQueryRouterMockRecorder {
	return m.recorder
}


func (m *MockQueryRouter) AddRoute(r string, h types.Querier) types.QueryRouter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddRoute", r, h)
	ret0, _ := ret[0].(types.QueryRouter)
	return ret0
}


func (mr *MockQueryRouterMockRecorder) AddRoute(r, h interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddRoute", reflect.TypeOf((*MockQueryRouter)(nil).AddRoute), r, h)
}


func (m *MockQueryRouter) Route(path string) types.Querier {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Route", path)
	ret0, _ := ret[0].(types.Querier)
	return ret0
}


func (mr *MockQueryRouterMockRecorder) Route(path interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Route", reflect.TypeOf((*MockQueryRouter)(nil).Route), path)
}
