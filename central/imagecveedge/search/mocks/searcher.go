// Code generated by MockGen. DO NOT EDIT.
// Source: searcher.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/stackrox/generated/api/v1"
	storage "github.com/stackrox/stackrox/generated/storage"
	search "github.com/stackrox/stackrox/pkg/search"
)

// MockSearcher is a mock of Searcher interface.
type MockSearcher struct {
	ctrl     *gomock.Controller
	recorder *MockSearcherMockRecorder
}

// MockSearcherMockRecorder is the mock recorder for MockSearcher.
type MockSearcherMockRecorder struct {
	mock *MockSearcher
}

// NewMockSearcher creates a new mock instance.
func NewMockSearcher(ctrl *gomock.Controller) *MockSearcher {
	mock := &MockSearcher{ctrl: ctrl}
	mock.recorder = &MockSearcherMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSearcher) EXPECT() *MockSearcherMockRecorder {
	return m.recorder
}

// Count mocks base method.
func (m *MockSearcher) Count(ctx context.Context, query *v1.Query) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Count", ctx, query)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Count indicates an expected call of Count.
func (mr *MockSearcherMockRecorder) Count(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockSearcher)(nil).Count), ctx, query)
}

// Search mocks base method.
func (m *MockSearcher) Search(ctx context.Context, query *v1.Query) ([]search.Result, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, query)
	ret0, _ := ret[0].([]search.Result)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockSearcherMockRecorder) Search(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockSearcher)(nil).Search), ctx, query)
}

// SearchEdges mocks base method.
func (m *MockSearcher) SearchEdges(arg0 context.Context, arg1 *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchEdges", arg0, arg1)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchEdges indicates an expected call of SearchEdges.
func (mr *MockSearcherMockRecorder) SearchEdges(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchEdges", reflect.TypeOf((*MockSearcher)(nil).SearchEdges), arg0, arg1)
}

// SearchRawEdges mocks base method.
func (m *MockSearcher) SearchRawEdges(ctx context.Context, query *v1.Query) ([]*storage.ImageCVEEdge, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawEdges", ctx, query)
	ret0, _ := ret[0].([]*storage.ImageCVEEdge)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawEdges indicates an expected call of SearchRawEdges.
func (mr *MockSearcherMockRecorder) SearchRawEdges(ctx, query interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawEdges", reflect.TypeOf((*MockSearcher)(nil).SearchRawEdges), ctx, query)
}
