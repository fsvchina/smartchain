





package mocks

import (
	reflect "reflect"

	types "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)


type MockAnteDecorator struct {
	ctrl     *gomock.Controller
	recorder *MockAnteDecoratorMockRecorder
}


type MockAnteDecoratorMockRecorder struct {
	mock *MockAnteDecorator
}


func NewMockAnteDecorator(ctrl *gomock.Controller) *MockAnteDecorator {
	mock := &MockAnteDecorator{ctrl: ctrl}
	mock.recorder = &MockAnteDecoratorMockRecorder{mock}
	return mock
}


func (m *MockAnteDecorator) EXPECT() *MockAnteDecoratorMockRecorder {
	return m.recorder
}


func (m *MockAnteDecorator) AnteHandle(ctx types.Context, tx types.Tx, simulate bool, next types.AnteHandler) (types.Context, error) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "AnteHandle", ctx, tx, simulate, next)

	return next(ctx, tx, simulate)
}


func (mr *MockAnteDecoratorMockRecorder) AnteHandle(ctx, tx, simulate, next interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AnteHandle", reflect.TypeOf((*MockAnteDecorator)(nil).AnteHandle), ctx, tx, simulate, next)
}
