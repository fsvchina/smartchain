



package mocks

import (
	reflect "reflect"

	client "github.com/cosmos/cosmos-sdk/client"
	types "github.com/cosmos/cosmos-sdk/crypto/types"
	types0 "github.com/cosmos/cosmos-sdk/types"
	gomock "github.com/golang/mock/gomock"
)


type MockAccount struct {
	ctrl     *gomock.Controller
	recorder *MockAccountMockRecorder
}


type MockAccountMockRecorder struct {
	mock *MockAccount
}


func NewMockAccount(ctrl *gomock.Controller) *MockAccount {
	mock := &MockAccount{ctrl: ctrl}
	mock.recorder = &MockAccountMockRecorder{mock}
	return mock
}


func (m *MockAccount) EXPECT() *MockAccountMockRecorder {
	return m.recorder
}


func (m *MockAccount) GetAccountNumber() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountNumber")
	ret0, _ := ret[0].(uint64)
	return ret0
}


func (mr *MockAccountMockRecorder) GetAccountNumber() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountNumber", reflect.TypeOf((*MockAccount)(nil).GetAccountNumber))
}


func (m *MockAccount) GetAddress() types0.AccAddress {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAddress")
	ret0, _ := ret[0].(types0.AccAddress)
	return ret0
}


func (mr *MockAccountMockRecorder) GetAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAddress", reflect.TypeOf((*MockAccount)(nil).GetAddress))
}


func (m *MockAccount) GetPubKey() types.PubKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPubKey")
	ret0, _ := ret[0].(types.PubKey)
	return ret0
}


func (mr *MockAccountMockRecorder) GetPubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPubKey", reflect.TypeOf((*MockAccount)(nil).GetPubKey))
}


func (m *MockAccount) GetSequence() uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSequence")
	ret0, _ := ret[0].(uint64)
	return ret0
}


func (mr *MockAccountMockRecorder) GetSequence() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSequence", reflect.TypeOf((*MockAccount)(nil).GetSequence))
}


type MockAccountRetriever struct {
	ctrl     *gomock.Controller
	recorder *MockAccountRetrieverMockRecorder
}


type MockAccountRetrieverMockRecorder struct {
	mock *MockAccountRetriever
}


func NewMockAccountRetriever(ctrl *gomock.Controller) *MockAccountRetriever {
	mock := &MockAccountRetriever{ctrl: ctrl}
	mock.recorder = &MockAccountRetrieverMockRecorder{mock}
	return mock
}


func (m *MockAccountRetriever) EXPECT() *MockAccountRetrieverMockRecorder {
	return m.recorder
}


func (m *MockAccountRetriever) EnsureExists(clientCtx client.Context, addr types0.AccAddress) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnsureExists", clientCtx, addr)
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockAccountRetrieverMockRecorder) EnsureExists(clientCtx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnsureExists", reflect.TypeOf((*MockAccountRetriever)(nil).EnsureExists), clientCtx, addr)
}


func (m *MockAccountRetriever) GetAccount(clientCtx client.Context, addr types0.AccAddress) (client.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccount", clientCtx, addr)
	ret0, _ := ret[0].(client.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}


func (mr *MockAccountRetrieverMockRecorder) GetAccount(clientCtx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccount", reflect.TypeOf((*MockAccountRetriever)(nil).GetAccount), clientCtx, addr)
}


func (m *MockAccountRetriever) GetAccountNumberSequence(clientCtx client.Context, addr types0.AccAddress) (uint64, uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountNumberSequence", clientCtx, addr)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(uint64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}


func (mr *MockAccountRetrieverMockRecorder) GetAccountNumberSequence(clientCtx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountNumberSequence", reflect.TypeOf((*MockAccountRetriever)(nil).GetAccountNumberSequence), clientCtx, addr)
}


func (m *MockAccountRetriever) GetAccountWithHeight(clientCtx client.Context, addr types0.AccAddress) (client.Account, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAccountWithHeight", clientCtx, addr)
	ret0, _ := ret[0].(client.Account)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}


func (mr *MockAccountRetrieverMockRecorder) GetAccountWithHeight(clientCtx, addr interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAccountWithHeight", reflect.TypeOf((*MockAccountRetriever)(nil).GetAccountWithHeight), clientCtx, addr)
}
