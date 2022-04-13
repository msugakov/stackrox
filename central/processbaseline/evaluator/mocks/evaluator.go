// Code generated by MockGen. DO NOT EDIT.
// Source: evaluator.go

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	storage "github.com/stackrox/stackrox/generated/storage"
)

// MockEvaluator is a mock of Evaluator interface.
type MockEvaluator struct {
	ctrl     *gomock.Controller
	recorder *MockEvaluatorMockRecorder
}

// MockEvaluatorMockRecorder is the mock recorder for MockEvaluator.
type MockEvaluatorMockRecorder struct {
	mock *MockEvaluator
}

// NewMockEvaluator creates a new mock instance.
func NewMockEvaluator(ctrl *gomock.Controller) *MockEvaluator {
	mock := &MockEvaluator{ctrl: ctrl}
	mock.recorder = &MockEvaluatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockEvaluator) EXPECT() *MockEvaluatorMockRecorder {
	return m.recorder
}

// EvaluateBaselinesAndPersistResult mocks base method.
func (m *MockEvaluator) EvaluateBaselinesAndPersistResult(deployment *storage.Deployment) ([]*storage.ProcessIndicator, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EvaluateBaselinesAndPersistResult", deployment)
	ret0, _ := ret[0].([]*storage.ProcessIndicator)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EvaluateBaselinesAndPersistResult indicates an expected call of EvaluateBaselinesAndPersistResult.
func (mr *MockEvaluatorMockRecorder) EvaluateBaselinesAndPersistResult(deployment interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EvaluateBaselinesAndPersistResult", reflect.TypeOf((*MockEvaluator)(nil).EvaluateBaselinesAndPersistResult), deployment)
}
