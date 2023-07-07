// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	datastore "github.com/stackrox/rox/central/processlisteningonport/datastore"
	storage "github.com/stackrox/rox/generated/storage"
	gomock "go.uber.org/mock/gomock"
)

// MockDataStore is a mock of DataStore interface.
type MockDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockDataStoreMockRecorder
}

// MockDataStoreMockRecorder is the mock recorder for MockDataStore.
type MockDataStoreMockRecorder struct {
	mock *MockDataStore
}

// NewMockDataStore creates a new mock instance.
func NewMockDataStore(ctrl *gomock.Controller) *MockDataStore {
	mock := &MockDataStore{ctrl: ctrl}
	mock.recorder = &MockDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDataStore) EXPECT() *MockDataStoreMockRecorder {
	return m.recorder
}

// AddProcessListeningOnPort mocks base method.
func (m *MockDataStore) AddProcessListeningOnPort(arg0 context.Context, arg1 ...*storage.ProcessListeningOnPortFromSensor) error {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0}
	for _, a := range arg1 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AddProcessListeningOnPort", varargs...)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddProcessListeningOnPort indicates an expected call of AddProcessListeningOnPort.
func (mr *MockDataStoreMockRecorder) AddProcessListeningOnPort(arg0 interface{}, arg1 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0}, arg1...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessListeningOnPort", reflect.TypeOf((*MockDataStore)(nil).AddProcessListeningOnPort), varargs...)
}

// GetProcessListeningOnPort mocks base method.
func (m *MockDataStore) GetProcessListeningOnPort(ctx context.Context, deployment string) ([]*storage.ProcessListeningOnPort, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProcessListeningOnPort", ctx, deployment)
	ret0, _ := ret[0].([]*storage.ProcessListeningOnPort)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProcessListeningOnPort indicates an expected call of GetProcessListeningOnPort.
func (mr *MockDataStoreMockRecorder) GetProcessListeningOnPort(ctx, deployment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProcessListeningOnPort", reflect.TypeOf((*MockDataStore)(nil).GetProcessListeningOnPort), ctx, deployment)
}

// RemoveProcessListeningOnPort mocks base method.
func (m *MockDataStore) RemoveProcessListeningOnPort(ctx context.Context, ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProcessListeningOnPort", ctx, ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessListeningOnPort indicates an expected call of RemoveProcessListeningOnPort.
func (mr *MockDataStoreMockRecorder) RemoveProcessListeningOnPort(ctx, ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessListeningOnPort", reflect.TypeOf((*MockDataStore)(nil).RemoveProcessListeningOnPort), ctx, ids)
}

// WalkAll mocks base method.
func (m *MockDataStore) WalkAll(ctx context.Context, fn datastore.WalkFn) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkAll", ctx, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkAll indicates an expected call of WalkAll.
func (mr *MockDataStoreMockRecorder) WalkAll(ctx, fn interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkAll", reflect.TypeOf((*MockDataStore)(nil).WalkAll), ctx, fn)
}
