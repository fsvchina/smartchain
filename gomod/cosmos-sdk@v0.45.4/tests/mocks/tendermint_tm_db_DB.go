



package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	db "github.com/tendermint/tm-db"
)


type MockDB struct {
	ctrl     *gomock.Controller
	recorder *MockDBMockRecorder
}


type MockDBMockRecorder struct {
	mock *MockDB
}


func NewMockDB(ctrl *gomock.Controller) *MockDB {
	mock := &MockDB{ctrl: ctrl}
	mock.recorder = &MockDBMockRecorder{mock}
	return mock
}


func (m *MockDB) EXPECT() *MockDBMockRecorder {
	return m.recorder
}


func (m *MockDB) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockDBMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockDB)(nil).Close))
}


func (m *MockDB) Delete(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockDBMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockDB)(nil).Delete), arg0)
}


func (m *MockDB) DeleteSync(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteSync", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockDBMockRecorder) DeleteSync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteSync", reflect.TypeOf((*MockDB)(nil).DeleteSync), arg0)
}


func (m *MockDB) Get(arg0 []byte) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", arg0)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}


func (mr *MockDBMockRecorder) Get(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDB)(nil).Get), arg0)
}


func (m *MockDB) Has(arg0 []byte) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Has", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}


func (mr *MockDBMockRecorder) Has(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Has", reflect.TypeOf((*MockDB)(nil).Has), arg0)
}


func (m *MockDB) Iterator(arg0, arg1 []byte) (db.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Iterator", arg0, arg1)
	ret0, _ := ret[0].(db.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}


func (mr *MockDBMockRecorder) Iterator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Iterator", reflect.TypeOf((*MockDB)(nil).Iterator), arg0, arg1)
}


func (m *MockDB) NewBatch() db.Batch {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewBatch")
	ret0, _ := ret[0].(db.Batch)
	return ret0
}


func (mr *MockDBMockRecorder) NewBatch() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewBatch", reflect.TypeOf((*MockDB)(nil).NewBatch))
}


func (m *MockDB) Print() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Print")
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockDBMockRecorder) Print() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Print", reflect.TypeOf((*MockDB)(nil).Print))
}


func (m *MockDB) ReverseIterator(arg0, arg1 []byte) (db.Iterator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ReverseIterator", arg0, arg1)
	ret0, _ := ret[0].(db.Iterator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}


func (mr *MockDBMockRecorder) ReverseIterator(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ReverseIterator", reflect.TypeOf((*MockDB)(nil).ReverseIterator), arg0, arg1)
}


func (m *MockDB) Set(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Set", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockDBMockRecorder) Set(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Set", reflect.TypeOf((*MockDB)(nil).Set), arg0, arg1)
}


func (m *MockDB) SetSync(arg0, arg1 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetSync", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}


func (mr *MockDBMockRecorder) SetSync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetSync", reflect.TypeOf((*MockDB)(nil).SetSync), arg0, arg1)
}


func (m *MockDB) Stats() map[string]string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Stats")
	ret0, _ := ret[0].(map[string]string)
	return ret0
}


func (mr *MockDBMockRecorder) Stats() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stats", reflect.TypeOf((*MockDB)(nil).Stats))
}
