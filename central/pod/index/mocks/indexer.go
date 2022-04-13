// Code generated by MockGen. DO NOT EDIT.
// Source: indexer.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	storage "github.com/stackrox/stackrox/generated/storage"
	search "github.com/stackrox/stackrox/pkg/search"
	blevesearch "github.com/stackrox/stackrox/pkg/search/blevesearch"
)

// MockIndexer is a mock of Indexer interface.
type MockIndexer struct {
	ctrl     *gomock.Controller
	recorder *MockIndexerMockRecorder
}

// MockIndexerMockRecorder is the mock recorder for MockIndexer.
type MockIndexerMockRecorder struct {
	mock *MockIndexer
}

// NewMockIndexer creates a new mock instance.
func NewMockIndexer(ctrl *gomock.Controller) *MockIndexer {
	mock := &MockIndexer{ctrl: ctrl}
	mock.recorder = &MockIndexerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockIndexer) EXPECT() *MockIndexerMockRecorder {
	return m.recorder
}

// AddPod mocks base method.
func (m *MockIndexer) AddPod(pod *storage.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPod", pod)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPod indicates an expected call of AddPod.
func (mr *MockIndexerMockRecorder) AddPod(pod interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPod", reflect.TypeOf((*MockIndexer)(nil).AddPod), pod)
}

// AddPods mocks base method.
func (m *MockIndexer) AddPods(pods []*storage.Pod) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPods", pods)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddPods indicates an expected call of AddPods.
func (mr *MockIndexerMockRecorder) AddPods(pods interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPods", reflect.TypeOf((*MockIndexer)(nil).AddPods), pods)
}

// Count mocks base method.
func (m *MockIndexer) Count(q *v1.Query, opts ...blevesearch.SearchOption) (int, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{q}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Count", varargs...)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockIndexerMockRecorder) Count(q interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{q}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockIndexer)(nil).Count), varargs...)
}

// DeletePod mocks base method.
func (m *MockIndexer) DeletePod(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePod", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePod indicates an expected call of DeletePod.
func (mr *MockIndexerMockRecorder) DeletePod(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePod", reflect.TypeOf((*MockIndexer)(nil).DeletePod), id)
}

// DeletePods mocks base method.
func (m *MockIndexer) DeletePods(ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePods", ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePods indicates an expected call of DeletePods.
func (mr *MockIndexerMockRecorder) DeletePods(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePods", reflect.TypeOf((*MockIndexer)(nil).DeletePods), ids)
}

// MarkInitialIndexingComplete mocks base method.
func (m *MockIndexer) MarkInitialIndexingComplete() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "MarkInitialIndexingComplete")
	ret0, _ := ret[0].(error)
	return ret0
}

// MarkInitialIndexingComplete indicates an expected call of MarkInitialIndexingComplete.
func (mr *MockIndexerMockRecorder) MarkInitialIndexingComplete() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "MarkInitialIndexingComplete", reflect.TypeOf((*MockIndexer)(nil).MarkInitialIndexingComplete))
}

// NeedsInitialIndexing mocks base method.
func (m *MockIndexer) NeedsInitialIndexing() (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NeedsInitialIndexing")
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NeedsInitialIndexing indicates an expected call of NeedsInitialIndexing.
func (mr *MockIndexerMockRecorder) NeedsInitialIndexing() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NeedsInitialIndexing", reflect.TypeOf((*MockIndexer)(nil).NeedsInitialIndexing))
}

// Search mocks base method.
func (m *MockIndexer) Search(q *v1.Query, opts ...blevesearch.SearchOption) ([]search.Result, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{q}
	for _, a := range opts {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Search", varargs...)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockIndexerMockRecorder) Search(q interface{}, opts ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{q}, opts...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockIndexer)(nil).Search), varargs...)
}
