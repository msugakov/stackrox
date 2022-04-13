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

// AddNamespaceMetadata mocks base method.
func (m *MockIndexer) AddNamespaceMetadata(namespacemetadata *storage.NamespaceMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNamespaceMetadata", namespacemetadata)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNamespaceMetadata indicates an expected call of AddNamespaceMetadata.
func (mr *MockIndexerMockRecorder) AddNamespaceMetadata(namespacemetadata interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNamespaceMetadata", reflect.TypeOf((*MockIndexer)(nil).AddNamespaceMetadata), namespacemetadata)
}

// AddNamespaceMetadatas mocks base method.
func (m *MockIndexer) AddNamespaceMetadatas(namespacemetadatas []*storage.NamespaceMetadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNamespaceMetadatas", namespacemetadatas)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNamespaceMetadatas indicates an expected call of AddNamespaceMetadatas.
func (mr *MockIndexerMockRecorder) AddNamespaceMetadatas(namespacemetadatas interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNamespaceMetadatas", reflect.TypeOf((*MockIndexer)(nil).AddNamespaceMetadatas), namespacemetadatas)
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

// DeleteNamespaceMetadata mocks base method.
func (m *MockIndexer) DeleteNamespaceMetadata(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespaceMetadata", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNamespaceMetadata indicates an expected call of DeleteNamespaceMetadata.
func (mr *MockIndexerMockRecorder) DeleteNamespaceMetadata(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespaceMetadata", reflect.TypeOf((*MockIndexer)(nil).DeleteNamespaceMetadata), id)
}

// DeleteNamespaceMetadatas mocks base method.
func (m *MockIndexer) DeleteNamespaceMetadatas(ids []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteNamespaceMetadatas", ids)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteNamespaceMetadatas indicates an expected call of DeleteNamespaceMetadatas.
func (mr *MockIndexerMockRecorder) DeleteNamespaceMetadatas(ids interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteNamespaceMetadatas", reflect.TypeOf((*MockIndexer)(nil).DeleteNamespaceMetadatas), ids)
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
