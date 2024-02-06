// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go
//
// Generated by this command:
//
//	mockgen -package mocks -destination mocks/datastore.go -source datastore.go
//

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
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

// Count mocks base method.
func (m *MockDataStore) Count(ctx context.Context, q *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, q)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockDataStoreMockRecorder) Count(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDataStore)(nil).Count), ctx, q)
}

// GetPod mocks base method.
func (m *MockDataStore) GetPod(ctx context.Context, id string) (*storage.Pod, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPod", ctx, id)
	ret0, _ := ret[0].(*storage.Pod)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPod indicates an expected call of GetPod.
func (mr *MockDataStoreMockRecorder) GetPod(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPod", reflect.TypeOf((*MockDataStore)(nil).GetPod), ctx, id)
}

// RemovePod mocks base method.
func (m *MockDataStore) RemovePod(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemovePod", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemovePod indicates an expected call of RemovePod.
func (mr *MockDataStoreMockRecorder) RemovePod(ctx, id any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemovePod", reflect.TypeOf((*MockDataStore)(nil).RemovePod), ctx, id)
}

// Search mocks base method.
func (m *MockDataStore) Search(ctx context.Context, q *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, q)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockDataStoreMockRecorder) Search(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchRawPods mocks base method.
func (m *MockDataStore) SearchRawPods(ctx context.Context, q *v1.Query) ([]*storage.Pod, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawPods", ctx, q)
	ret0, _ := ret[0].([]*storage.Pod)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawPods indicates an expected call of SearchRawPods.
func (mr *MockDataStoreMockRecorder) SearchRawPods(ctx, q any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawPods", reflect.TypeOf((*MockDataStore)(nil).SearchRawPods), ctx, q)
}

// UpsertPod mocks base method.
func (m *MockDataStore) UpsertPod(ctx context.Context, pod *storage.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertPod", ctx, pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertPod indicates an expected call of UpsertPod.
func (mr *MockDataStoreMockRecorder) UpsertPod(ctx, pod any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertPod", reflect.TypeOf((*MockDataStore)(nil).UpsertPod), ctx, pod)
}

// WalkByQuery mocks base method.
func (m *MockDataStore) WalkByQuery(ctx context.Context, q *v1.Query, fn func(*storage.Pod) error) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WalkByQuery", ctx, q, fn)
	ret0, _ := ret[0].(error)
	return ret0
}

// WalkByQuery indicates an expected call of WalkByQuery.
func (mr *MockDataStoreMockRecorder) WalkByQuery(ctx, q, fn any) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WalkByQuery", reflect.TypeOf((*MockDataStore)(nil).WalkByQuery), ctx, q, fn)
}
