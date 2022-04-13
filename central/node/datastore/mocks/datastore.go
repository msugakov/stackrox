// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
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

// CountNodes mocks base method.
func (m *MockDataStore) CountNodes() (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CountNodes")
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CountNodes indicates an expected call of CountNodes.
func (mr *MockDataStoreMockRecorder) CountNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CountNodes", reflect.TypeOf((*MockDataStore)(nil).CountNodes))
}

// GetNode mocks base method.
func (m *MockDataStore) GetNode(id string) (*storage.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNode", id)
	ret0, _ := ret[0].(*storage.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNode indicates an expected call of GetNode.
func (mr *MockDataStoreMockRecorder) GetNode(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNode", reflect.TypeOf((*MockDataStore)(nil).GetNode), id)
}

// ListNodes mocks base method.
func (m *MockDataStore) ListNodes() ([]*storage.Node, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListNodes")
	ret0, _ := ret[0].([]*storage.Node)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListNodes indicates an expected call of ListNodes.
func (mr *MockDataStoreMockRecorder) ListNodes() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListNodes", reflect.TypeOf((*MockDataStore)(nil).ListNodes))
}

// RemoveNode mocks base method.
func (m *MockDataStore) RemoveNode(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveNode", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveNode indicates an expected call of RemoveNode.
func (mr *MockDataStoreMockRecorder) RemoveNode(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveNode", reflect.TypeOf((*MockDataStore)(nil).RemoveNode), id)
}

// UpsertNode mocks base method.
func (m *MockDataStore) UpsertNode(node *storage.Node) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpsertNode", node)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpsertNode indicates an expected call of UpsertNode.
func (mr *MockDataStoreMockRecorder) UpsertNode(node interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpsertNode", reflect.TypeOf((*MockDataStore)(nil).UpsertNode), node)
}
