// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/stackrox/generated/storage"
)

// MockStore is a mock of Store interface.
type MockStore struct {
	ctrl     *gomock.Controller
	recorder *MockStoreMockRecorder
}

// MockStoreMockRecorder is the mock recorder for MockStore.
type MockStoreMockRecorder struct {
	mock *MockStore
}

// NewMockStore creates a new mock instance.
func NewMockStore(ctrl *gomock.Controller) *MockStore {
	mock := &MockStore{ctrl: ctrl}
	mock.recorder = &MockStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockStore) EXPECT() *MockStoreMockRecorder {
	return m.recorder
}

// GetBackup mocks base method.
func (m *MockStore) GetBackup(id string) (*storage.ExternalBackup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBackup", id)
	ret0, _ := ret[0].(*storage.ExternalBackup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetBackup indicates an expected call of GetBackup.
func (mr *MockStoreMockRecorder) GetBackup(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBackup", reflect.TypeOf((*MockStore)(nil).GetBackup), id)
}

// ListBackups mocks base method.
func (m *MockStore) ListBackups() ([]*storage.ExternalBackup, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListBackups")
	ret0, _ := ret[0].([]*storage.ExternalBackup)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListBackups indicates an expected call of ListBackups.
func (mr *MockStoreMockRecorder) ListBackups() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListBackups", reflect.TypeOf((*MockStore)(nil).ListBackups))
}

// RemoveBackup mocks base method.
func (m *MockStore) RemoveBackup(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveBackup", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveBackup indicates an expected call of RemoveBackup.
func (mr *MockStoreMockRecorder) RemoveBackup(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveBackup", reflect.TypeOf((*MockStore)(nil).RemoveBackup), id)
}

// UpsertBackup mocks base method.
func (m *MockStore) UpsertBackup(backup *storage.ExternalBackup) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertBackup", backup)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertBackup indicates an expected call of UpsertBackup.
func (mr *MockStoreMockRecorder) UpsertBackup(backup interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertBackup", reflect.TypeOf((*MockStore)(nil).UpsertBackup), backup)
}
