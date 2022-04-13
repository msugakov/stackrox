// Code generated by MockGen. DO NOT EDIT.
// Source: health_reporter.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/stackrox/generated/storage"
)

// MockReporter is a mock of Reporter interface.
type MockReporter struct {
	ctrl     *gomock.Controller
	recorder *MockReporterMockRecorder
}

// MockReporterMockRecorder is the mock recorder for MockReporter.
type MockReporterMockRecorder struct {
	mock *MockReporter
}

// NewMockReporter creates a new mock instance.
func NewMockReporter(ctrl *gomock.Controller) *MockReporter {
	mock := &MockReporter{ctrl: ctrl}
	mock.recorder = &MockReporterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReporter) EXPECT() *MockReporterMockRecorder {
	return m.recorder
}

// Register mocks base method.
func (m *MockReporter) Register(id, name string, typ storage.IntegrationHealth_Type) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Register", id, name, typ)
	ret0, _ := ret[0].(error)
	return ret0
}

// Register indicates an expected call of Register.
func (mr *MockReporterMockRecorder) Register(id, name, typ interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Register", reflect.TypeOf((*MockReporter)(nil).Register), id, name, typ)
}

// RemoveIntegrationHealth mocks base method.
func (m *MockReporter) RemoveIntegrationHealth(id string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveIntegrationHealth", id)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveIntegrationHealth indicates an expected call of RemoveIntegrationHealth.
func (mr *MockReporterMockRecorder) RemoveIntegrationHealth(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveIntegrationHealth", reflect.TypeOf((*MockReporter)(nil).RemoveIntegrationHealth), id)
}

// UpdateIntegrationHealthAsync mocks base method.
func (m *MockReporter) UpdateIntegrationHealthAsync(arg0 *storage.IntegrationHealth) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "UpdateIntegrationHealthAsync", arg0)
}

// UpdateIntegrationHealthAsync indicates an expected call of UpdateIntegrationHealthAsync.
func (mr *MockReporterMockRecorder) UpdateIntegrationHealthAsync(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateIntegrationHealthAsync", reflect.TypeOf((*MockReporter)(nil).UpdateIntegrationHealthAsync), arg0)
}
