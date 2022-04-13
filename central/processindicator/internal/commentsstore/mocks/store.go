// Code generated by MockGen. DO NOT EDIT.
// Source: store.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	analystnotes "github.com/stackrox/stackrox/central/analystnotes"
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

// AddProcessComment mocks base method.
func (m *MockStore) AddProcessComment(key *analystnotes.ProcessNoteKey, comment *storage.Comment) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddProcessComment", key, comment)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddProcessComment indicates an expected call of AddProcessComment.
func (mr *MockStoreMockRecorder) AddProcessComment(key, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddProcessComment", reflect.TypeOf((*MockStore)(nil).AddProcessComment), key, comment)
}

// GetComment mocks base method.
func (m *MockStore) GetComment(key *analystnotes.ProcessNoteKey, commentID string) (*storage.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComment", key, commentID)
	ret0, _ := ret[0].(*storage.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComment indicates an expected call of GetComment.
func (mr *MockStoreMockRecorder) GetComment(key, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComment", reflect.TypeOf((*MockStore)(nil).GetComment), key, commentID)
}

// GetComments mocks base method.
func (m *MockStore) GetComments(key *analystnotes.ProcessNoteKey) ([]*storage.Comment, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetComments", key)
	ret0, _ := ret[0].([]*storage.Comment)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetComments indicates an expected call of GetComments.
func (mr *MockStoreMockRecorder) GetComments(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetComments", reflect.TypeOf((*MockStore)(nil).GetComments), key)
}

// GetCommentsCount mocks base method.
func (m *MockStore) GetCommentsCount(key *analystnotes.ProcessNoteKey) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetCommentsCount", key)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetCommentsCount indicates an expected call of GetCommentsCount.
func (mr *MockStoreMockRecorder) GetCommentsCount(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCommentsCount", reflect.TypeOf((*MockStore)(nil).GetCommentsCount), key)
}

// RemoveAllProcessComments mocks base method.
func (m *MockStore) RemoveAllProcessComments(key *analystnotes.ProcessNoteKey) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveAllProcessComments", key)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveAllProcessComments indicates an expected call of RemoveAllProcessComments.
func (mr *MockStoreMockRecorder) RemoveAllProcessComments(key interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveAllProcessComments", reflect.TypeOf((*MockStore)(nil).RemoveAllProcessComments), key)
}

// RemoveProcessComment mocks base method.
func (m *MockStore) RemoveProcessComment(key *analystnotes.ProcessNoteKey, commentID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveProcessComment", key, commentID)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveProcessComment indicates an expected call of RemoveProcessComment.
func (mr *MockStoreMockRecorder) RemoveProcessComment(key, commentID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveProcessComment", reflect.TypeOf((*MockStore)(nil).RemoveProcessComment), key, commentID)
}

// UpdateProcessComment mocks base method.
func (m *MockStore) UpdateProcessComment(key *analystnotes.ProcessNoteKey, comment *storage.Comment) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateProcessComment", key, comment)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateProcessComment indicates an expected call of UpdateProcessComment.
func (mr *MockStoreMockRecorder) UpdateProcessComment(key, comment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateProcessComment", reflect.TypeOf((*MockStore)(nil).UpdateProcessComment), key, comment)
}
