// Code generated by MockGen. DO NOT EDIT.
// Source: types.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/rox/generated/storage"
)

// MockDeploymentStore is a mock of DeploymentStore interface.
type MockDeploymentStore struct {
	ctrl     *gomock.Controller
	recorder *MockDeploymentStoreMockRecorder
}

// MockDeploymentStoreMockRecorder is the mock recorder for MockDeploymentStore.
type MockDeploymentStoreMockRecorder struct {
	mock *MockDeploymentStore
}

// NewMockDeploymentStore creates a new mock instance.
func NewMockDeploymentStore(ctrl *gomock.Controller) *MockDeploymentStore {
	mock := &MockDeploymentStore{ctrl: ctrl}
	mock.recorder = &MockDeploymentStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeploymentStore) EXPECT() *MockDeploymentStoreMockRecorder {
	return m.recorder
}

// Get mocks base method.
func (m *MockDeploymentStore) Get(id string) *storage.Deployment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*storage.Deployment)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockDeploymentStoreMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockDeploymentStore)(nil).Get), id)
}

// GetAll mocks base method.
func (m *MockDeploymentStore) GetAll() []*storage.Deployment {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*storage.Deployment)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockDeploymentStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockDeploymentStore)(nil).GetAll))
}

// MockPodStore is a mock of PodStore interface.
type MockPodStore struct {
	ctrl     *gomock.Controller
	recorder *MockPodStoreMockRecorder
}

// MockPodStoreMockRecorder is the mock recorder for MockPodStore.
type MockPodStoreMockRecorder struct {
	mock *MockPodStore
}

// NewMockPodStore creates a new mock instance.
func NewMockPodStore(ctrl *gomock.Controller) *MockPodStore {
	mock := &MockPodStore{ctrl: ctrl}
	mock.recorder = &MockPodStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockPodStore) EXPECT() *MockPodStoreMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockPodStore) GetAll() []*storage.Pod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]*storage.Pod)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockPodStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockPodStore)(nil).GetAll))
}

// GetByName mocks base method.
func (m *MockPodStore) GetByName(podName, namespace string) *storage.Pod {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByName", podName, namespace)
	ret0, _ := ret[0].(*storage.Pod)
	return ret0
}

// GetByName indicates an expected call of GetByName.
func (mr *MockPodStoreMockRecorder) GetByName(podName, namespace interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByName", reflect.TypeOf((*MockPodStore)(nil).GetByName), podName, namespace)
}

// MockNetworkPolicyStore is a mock of NetworkPolicyStore interface.
type MockNetworkPolicyStore struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkPolicyStoreMockRecorder
}

// MockNetworkPolicyStoreMockRecorder is the mock recorder for MockNetworkPolicyStore.
type MockNetworkPolicyStoreMockRecorder struct {
	mock *MockNetworkPolicyStore
}

// NewMockNetworkPolicyStore creates a new mock instance.
func NewMockNetworkPolicyStore(ctrl *gomock.Controller) *MockNetworkPolicyStore {
	mock := &MockNetworkPolicyStore{ctrl: ctrl}
	mock.recorder = &MockNetworkPolicyStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkPolicyStore) EXPECT() *MockNetworkPolicyStoreMockRecorder {
	return m.recorder
}

// All mocks base method.
func (m *MockNetworkPolicyStore) All() map[string]*storage.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "All")
	ret0, _ := ret[0].(map[string]*storage.NetworkPolicy)
	return ret0
}

// All indicates an expected call of All.
func (mr *MockNetworkPolicyStoreMockRecorder) All() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "All", reflect.TypeOf((*MockNetworkPolicyStore)(nil).All))
}

// Delete mocks base method.
func (m *MockNetworkPolicyStore) Delete(ID, ns string) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Delete", ID, ns)
}

// Delete indicates an expected call of Delete.
func (mr *MockNetworkPolicyStoreMockRecorder) Delete(ID, ns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockNetworkPolicyStore)(nil).Delete), ID, ns)
}

// Find mocks base method.
func (m *MockNetworkPolicyStore) Find(namespace string, labels map[string]string) map[string]*storage.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Find", namespace, labels)
	ret0, _ := ret[0].(map[string]*storage.NetworkPolicy)
	return ret0
}

// Find indicates an expected call of Find.
func (mr *MockNetworkPolicyStoreMockRecorder) Find(namespace, labels interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Find", reflect.TypeOf((*MockNetworkPolicyStore)(nil).Find), namespace, labels)
}

// Get mocks base method.
func (m *MockNetworkPolicyStore) Get(id string) *storage.NetworkPolicy {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Get", id)
	ret0, _ := ret[0].(*storage.NetworkPolicy)
	return ret0
}

// Get indicates an expected call of Get.
func (mr *MockNetworkPolicyStoreMockRecorder) Get(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockNetworkPolicyStore)(nil).Get), id)
}

// Size mocks base method.
func (m *MockNetworkPolicyStore) Size() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Size")
	ret0, _ := ret[0].(int)
	return ret0
}

// Size indicates an expected call of Size.
func (mr *MockNetworkPolicyStoreMockRecorder) Size() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Size", reflect.TypeOf((*MockNetworkPolicyStore)(nil).Size))
}

// Upsert mocks base method.
func (m *MockNetworkPolicyStore) Upsert(ns *storage.NetworkPolicy) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Upsert", ns)
}

// Upsert indicates an expected call of Upsert.
func (mr *MockNetworkPolicyStoreMockRecorder) Upsert(ns interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockNetworkPolicyStore)(nil).Upsert), ns)
}

// MockServiceAccountStore is a mock of ServiceAccountStore interface.
type MockServiceAccountStore struct {
	ctrl     *gomock.Controller
	recorder *MockServiceAccountStoreMockRecorder
}

// MockServiceAccountStoreMockRecorder is the mock recorder for MockServiceAccountStore.
type MockServiceAccountStoreMockRecorder struct {
	mock *MockServiceAccountStore
}

// NewMockServiceAccountStore creates a new mock instance.
func NewMockServiceAccountStore(ctrl *gomock.Controller) *MockServiceAccountStore {
	mock := &MockServiceAccountStore{ctrl: ctrl}
	mock.recorder = &MockServiceAccountStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceAccountStore) EXPECT() *MockServiceAccountStoreMockRecorder {
	return m.recorder
}

// Add mocks base method.
func (m *MockServiceAccountStore) Add(sa *storage.ServiceAccount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", sa)
}

// Add indicates an expected call of Add.
func (mr *MockServiceAccountStoreMockRecorder) Add(sa interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockServiceAccountStore)(nil).Add), sa)
}

// GetImagePullSecrets mocks base method.
func (m *MockServiceAccountStore) GetImagePullSecrets(namespace, name string) []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetImagePullSecrets", namespace, name)
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetImagePullSecrets indicates an expected call of GetImagePullSecrets.
func (mr *MockServiceAccountStoreMockRecorder) GetImagePullSecrets(namespace, name interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetImagePullSecrets", reflect.TypeOf((*MockServiceAccountStore)(nil).GetImagePullSecrets), namespace, name)
}

// Remove mocks base method.
func (m *MockServiceAccountStore) Remove(sa *storage.ServiceAccount) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Remove", sa)
}

// Remove indicates an expected call of Remove.
func (mr *MockServiceAccountStoreMockRecorder) Remove(sa interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Remove", reflect.TypeOf((*MockServiceAccountStore)(nil).Remove), sa)
}
