// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	storage "github.com/stackrox/stackrox/generated/storage"
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

// AddNotifier mocks base method.
func (m *MockDataStore) AddNotifier(ctx context.Context, notifier *storage.Notifier) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNotifier", ctx, notifier)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddNotifier indicates an expected call of AddNotifier.
func (mr *MockDataStoreMockRecorder) AddNotifier(ctx, notifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNotifier", reflect.TypeOf((*MockDataStore)(nil).AddNotifier), ctx, notifier)
}

// GetNotifier mocks base method.
func (m *MockDataStore) GetNotifier(ctx context.Context, id string) (*storage.Notifier, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifier", ctx, id)
	ret0, _ := ret[0].(*storage.Notifier)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetNotifier indicates an expected call of GetNotifier.
func (mr *MockDataStoreMockRecorder) GetNotifier(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifier", reflect.TypeOf((*MockDataStore)(nil).GetNotifier), ctx, id)
}

// GetNotifiers mocks base method.
func (m *MockDataStore) GetNotifiers(ctx context.Context, request *v1.GetNotifiersRequest) ([]*storage.Notifier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNotifiers", ctx, request)
	ret0, _ := ret[0].([]*storage.Notifier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNotifiers indicates an expected call of GetNotifiers.
func (mr *MockDataStoreMockRecorder) GetNotifiers(ctx, request interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNotifiers", reflect.TypeOf((*MockDataStore)(nil).GetNotifiers), ctx, request)
}

// RemoveNotifier mocks base method.
func (m *MockDataStore) RemoveNotifier(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNotifier", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNotifier indicates an expected call of RemoveNotifier.
func (mr *MockDataStoreMockRecorder) RemoveNotifier(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNotifier", reflect.TypeOf((*MockDataStore)(nil).RemoveNotifier), ctx, id)
}

// UpdateNotifier mocks base method.
func (m *MockDataStore) UpdateNotifier(ctx context.Context, notifier *storage.Notifier) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateNotifier", ctx, notifier)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateNotifier indicates an expected call of UpdateNotifier.
func (mr *MockDataStoreMockRecorder) UpdateNotifier(ctx, notifier interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateNotifier", reflect.TypeOf((*MockDataStore)(nil).UpdateNotifier), ctx, notifier)
}
