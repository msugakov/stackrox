// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
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

// DeleteLicenseKey mocks base method.
func (m *MockDataStore) DeleteLicenseKey(ctx context.Context, licenseID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteLicenseKey", ctx, licenseID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteLicenseKey indicates an expected call of DeleteLicenseKey.
func (mr *MockDataStoreMockRecorder) DeleteLicenseKey(ctx, licenseID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteLicenseKey", reflect.TypeOf((*MockDataStore)(nil).DeleteLicenseKey), ctx, licenseID)
}

// ListLicenseKeys mocks base method.
func (m *MockDataStore) ListLicenseKeys(ctx context.Context) ([]*storage.StoredLicenseKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListLicenseKeys", ctx)
	ret0, _ := ret[0].([]*storage.StoredLicenseKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListLicenseKeys indicates an expected call of ListLicenseKeys.
func (mr *MockDataStoreMockRecorder) ListLicenseKeys(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListLicenseKeys", reflect.TypeOf((*MockDataStore)(nil).ListLicenseKeys), ctx)
}

// UpsertLicenseKeys mocks base method.
func (m *MockDataStore) UpsertLicenseKeys(ctx context.Context, keys []*storage.StoredLicenseKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertLicenseKeys", ctx, keys)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertLicenseKeys indicates an expected call of UpsertLicenseKeys.
func (mr *MockDataStoreMockRecorder) UpsertLicenseKeys(ctx, keys interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertLicenseKeys", reflect.TypeOf((*MockDataStore)(nil).UpsertLicenseKeys), ctx, keys)
}
