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

// AddNodeComponentEdge mocks base method.
func (m *MockIndexer) AddNodeComponentEdge(nodecomponentedge *storage.NodeComponentEdge) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNodeComponentEdge", nodecomponentedge)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNodeComponentEdge indicates an expected call of AddNodeComponentEdge.
func (mr *MockIndexerMockRecorder) AddNodeComponentEdge(nodecomponentedge interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNodeComponentEdge", reflect.TypeOf((*MockIndexer)(nil).AddNodeComponentEdge), nodecomponentedge)
}

// AddNodeComponentEdges mocks base method.
func (m *MockIndexer) AddNodeComponentEdges(nodecomponentedges []*storage.NodeComponentEdge) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNodeComponentEdges", nodecomponentedges)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNodeComponentEdges indicates an expected call of AddNodeComponentEdges.
func (mr *MockIndexerMockRecorder) AddNodeComponentEdges(nodecomponentedges interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNodeComponentEdges", reflect.TypeOf((*MockIndexer)(nil).AddNodeComponentEdges), nodecomponentedges)
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

// DeleteNodeComponentEdge mocks base method.
func (m *MockIndexer) DeleteNodeComponentEdge(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNodeComponentEdge", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNodeComponentEdge indicates an expected call of DeleteNodeComponentEdge.
func (mr *MockIndexerMockRecorder) DeleteNodeComponentEdge(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNodeComponentEdge", reflect.TypeOf((*MockIndexer)(nil).DeleteNodeComponentEdge), id)
}

// DeleteNodeComponentEdges mocks base method.
func (m *MockIndexer) DeleteNodeComponentEdges(ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNodeComponentEdges", ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNodeComponentEdges indicates an expected call of DeleteNodeComponentEdges.
func (mr *MockIndexerMockRecorder) DeleteNodeComponentEdges(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNodeComponentEdges", reflect.TypeOf((*MockIndexer)(nil).DeleteNodeComponentEdges), ids)
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
