// Code generated by MockGen. DO NOT EDIT.
// Source: internal/exporter/exporter_context_interface.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExporterContextInterface is a mock of ExporterContextInterface interface.
type MockExporterContextInterface struct {
	ctrl     *gomock.Controller
	recorder *MockExporterContextInterfaceMockRecorder
}

// MockExporterContextInterfaceMockRecorder is the mock recorder for MockExporterContextInterface.
type MockExporterContextInterfaceMockRecorder struct {
	mock *MockExporterContextInterface
}

// NewMockExporterContextInterface creates a new mock instance.
func NewMockExporterContextInterface(ctrl *gomock.Controller) *MockExporterContextInterface {
	mock := &MockExporterContextInterface{ctrl: ctrl}
	mock.recorder = &MockExporterContextInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExporterContextInterface) EXPECT() *MockExporterContextInterfaceMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockExporterContextInterface) Save(fileName, exporterName string, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", fileName, exporterName, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockExporterContextInterfaceMockRecorder) Save(fileName, exporterName, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockExporterContextInterface)(nil).Save), fileName, exporterName, data)
}
