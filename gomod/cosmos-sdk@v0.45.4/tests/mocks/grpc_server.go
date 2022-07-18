



package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
)


type MockServer struct {
	ctrl     *gomock.Controller
	recorder *MockServerMockRecorder
}


type MockServerMockRecorder struct {
	mock *MockServer
}


func NewMockServer(ctrl *gomock.Controller) *MockServer {
	mock := &MockServer{ctrl: ctrl}
	mock.recorder = &MockServerMockRecorder{mock}
	return mock
}


func (m *MockServer) EXPECT() *MockServerMockRecorder {
	return m.recorder
}


func (m *MockServer) RegisterService(arg0 *grpc.ServiceDesc, arg1 interface{}) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RegisterService", arg0, arg1)
}


func (mr *MockServerMockRecorder) RegisterService(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RegisterService", reflect.TypeOf((*MockServer)(nil).RegisterService), arg0, arg1)
}
