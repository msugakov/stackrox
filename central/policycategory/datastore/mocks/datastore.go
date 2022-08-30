// Code generated by MockGen. DO NOT EDIT.
// Source: datastore.go

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1 "github.com/stackrox/rox/generated/api/v1"
	storage "github.com/stackrox/rox/generated/storage"
	search "github.com/stackrox/rox/pkg/search"
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

// AddPolicyCategory mocks base method.
func (m *MockDataStore) AddPolicyCategory(arg0 context.Context, arg1 *storage.PolicyCategory) (*storage.PolicyCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddPolicyCategory", arg0, arg1)
	ret0, _ := ret[0].(*storage.PolicyCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AddPolicyCategory indicates an expected call of AddPolicyCategory.
func (mr *MockDataStoreMockRecorder) AddPolicyCategory(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddPolicyCategory", reflect.TypeOf((*MockDataStore)(nil).AddPolicyCategory), arg0, arg1)
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
func (mr *MockDataStoreMockRecorder) Count(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Count", reflect.TypeOf((*MockDataStore)(nil).Count), ctx, q)
}

// DeletePolicyCategory mocks base method.
func (m *MockDataStore) DeletePolicyCategory(ctx context.Context, id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeletePolicyCategory", ctx, id)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeletePolicyCategory indicates an expected call of DeletePolicyCategory.
func (mr *MockDataStoreMockRecorder) DeletePolicyCategory(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeletePolicyCategory", reflect.TypeOf((*MockDataStore)(nil).DeletePolicyCategory), ctx, id)
}

// GetAllPolicyCategories mocks base method.
func (m *MockDataStore) GetAllPolicyCategories(ctx context.Context) ([]*storage.PolicyCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPolicyCategories", ctx)
	ret0, _ := ret[0].([]*storage.PolicyCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllPolicyCategories indicates an expected call of GetAllPolicyCategories.
func (mr *MockDataStoreMockRecorder) GetAllPolicyCategories(ctx interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPolicyCategories", reflect.TypeOf((*MockDataStore)(nil).GetAllPolicyCategories), ctx)
}

// GetPolicyCategoriesForPolicy mocks base method.
func (m *MockDataStore) GetPolicyCategoriesForPolicy(ctx context.Context, policyID string) ([]*storage.PolicyCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyCategoriesForPolicy", ctx, policyID)
	ret0, _ := ret[0].([]*storage.PolicyCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPolicyCategoriesForPolicy indicates an expected call of GetPolicyCategoriesForPolicy.
func (mr *MockDataStoreMockRecorder) GetPolicyCategoriesForPolicy(ctx, policyID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyCategoriesForPolicy", reflect.TypeOf((*MockDataStore)(nil).GetPolicyCategoriesForPolicy), ctx, policyID)
}

// GetPolicyCategory mocks base method.
func (m *MockDataStore) GetPolicyCategory(ctx context.Context, id string) (*storage.PolicyCategory, bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPolicyCategory", ctx, id)
	ret0, _ := ret[0].(*storage.PolicyCategory)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetPolicyCategory indicates an expected call of GetPolicyCategory.
func (mr *MockDataStoreMockRecorder) GetPolicyCategory(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPolicyCategory", reflect.TypeOf((*MockDataStore)(nil).GetPolicyCategory), ctx, id)
}

// RenamePolicyCategory mocks base method.
func (m *MockDataStore) RenamePolicyCategory(ctx context.Context, id, newName string) (*storage.PolicyCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RenamePolicyCategory", ctx, id, newName)
	ret0, _ := ret[0].(*storage.PolicyCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RenamePolicyCategory indicates an expected call of RenamePolicyCategory.
func (mr *MockDataStoreMockRecorder) RenamePolicyCategory(ctx, id, newName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RenamePolicyCategory", reflect.TypeOf((*MockDataStore)(nil).RenamePolicyCategory), ctx, id, newName)
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
func (mr *MockDataStoreMockRecorder) Search(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockDataStore)(nil).Search), ctx, q)
}

// SearchPolicyCategories mocks base method.
func (m *MockDataStore) SearchPolicyCategories(ctx context.Context, q *v1.Query) ([]*v1.SearchResult, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchPolicyCategories", ctx, q)
	ret0, _ := ret[0].([]*v1.SearchResult)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchPolicyCategories indicates an expected call of SearchPolicyCategories.
func (mr *MockDataStoreMockRecorder) SearchPolicyCategories(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchPolicyCategories", reflect.TypeOf((*MockDataStore)(nil).SearchPolicyCategories), ctx, q)
}

// SearchRawPolicyCategories mocks base method.
func (m *MockDataStore) SearchRawPolicyCategories(ctx context.Context, q *v1.Query) ([]*storage.PolicyCategory, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SearchRawPolicyCategories", ctx, q)
	ret0, _ := ret[0].([]*storage.PolicyCategory)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SearchRawPolicyCategories indicates an expected call of SearchRawPolicyCategories.
func (mr *MockDataStoreMockRecorder) SearchRawPolicyCategories(ctx, q interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SearchRawPolicyCategories", reflect.TypeOf((*MockDataStore)(nil).SearchRawPolicyCategories), ctx, q)
}

// SetPolicyCategoriesForPolicy mocks base method.
func (m *MockDataStore) SetPolicyCategoriesForPolicy(ctx context.Context, policyID string, categories []string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetPolicyCategoriesForPolicy", ctx, policyID, categories)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetPolicyCategoriesForPolicy indicates an expected call of SetPolicyCategoriesForPolicy.
func (mr *MockDataStoreMockRecorder) SetPolicyCategoriesForPolicy(ctx, policyID, categories interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetPolicyCategoriesForPolicy", reflect.TypeOf((*MockDataStore)(nil).SetPolicyCategoriesForPolicy), ctx, policyID, categories)
}
