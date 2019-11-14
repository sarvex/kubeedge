// Code generated by MockGen. DO NOT EDIT.
// Source: staging/src/github.com/kubeedge/beehive/pkg/core/module.go

// Package mock_core is a generated GoMock package.
package beehive

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockModule is a mock of Module interface
type MockModule struct {
	ctrl     *gomock.Controller
	recorder *MockModuleMockRecorder
}

// MockModuleMockRecorder is the mock recorder for MockModule
type MockModuleMockRecorder struct {
	mock *MockModule
}

// NewMockModule creates a new mock instance
func NewMockModule(ctrl *gomock.Controller) *MockModule {
	mock := &MockModule{ctrl: ctrl}
	mock.recorder = &MockModuleMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModule) EXPECT() *MockModuleMockRecorder {
	return m.recorder
}

// Name mocks base method
func (m *MockModule) Name() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(string)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockModuleMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockModule)(nil).Name))
}

// Group mocks base method
func (m *MockModule) Group() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Group")
	ret0, _ := ret[0].(string)
	return ret0
}

// Group indicates an expected call of Group
func (mr *MockModuleMockRecorder) Group() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Group", reflect.TypeOf((*MockModule)(nil).Group))
}

// Start mocks base method
func (m *MockModule) Start() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Start")
}

// Start indicates an expected call of Start
func (mr *MockModuleMockRecorder) Start() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Start", reflect.TypeOf((*MockModule)(nil).Start))
}

// Cleanup mocks base method
func (m *MockModule) Cleanup() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Cleanup")
}

// Cleanup indicates an expected call of Cleanup
func (mr *MockModuleMockRecorder) Cleanup() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Cleanup", reflect.TypeOf((*MockModule)(nil).Cleanup))
}
