// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/stackrox/stackrox/pkg/sac (interfaces: ClusterGetter)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/stackrox/generated/storage"
	reflect "reflect"
)

// MockClusterGetter is a mock of ClusterGetter interface
type MockClusterGetter struct {
	ctrl     *gomock.Controller
	recorder *MockClusterGetterMockRecorder
}

// MockClusterGetterMockRecorder is the mock recorder for MockClusterGetter
type MockClusterGetterMockRecorder struct {
	mock *MockClusterGetter
}

// NewMockClusterGetter creates a new mock instance
func NewMockClusterGetter(ctrl *gomock.Controller) *MockClusterGetter {
	mock := &MockClusterGetter{ctrl: ctrl}
	mock.recorder = &MockClusterGetterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockClusterGetter) EXPECT() *MockClusterGetterMockRecorder {
	return m.recorder
}

// GetCluster mocks base method
func (m *MockClusterGetter) GetCluster(arg0 context.Context, arg1 string) (*storage.Cluster, bool, error) {
	ret := m.ctrl.Call(m, "GetCluster", arg0, arg1)
	ret0, _ := ret[0].(*storage.Cluster)
	ret1, _ := ret[1].(bool)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetCluster indicates an expected call of GetCluster
func (mr *MockClusterGetterMockRecorder) GetCluster(arg0, arg1 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetCluster", reflect.TypeOf((*MockClusterGetter)(nil).GetCluster), arg0, arg1)
}
