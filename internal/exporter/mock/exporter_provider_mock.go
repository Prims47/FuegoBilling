// Code generated by MockGen. DO NOT EDIT.
// Source: internal/exporter/exporter_provider_interface.go

// Package exporter is a generated GoMock package.
package exporter

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockExporterProviderInterface is a mock of ExporterProviderInterface interface.
type MockExporterProviderInterface struct {
	ctrl     *gomock.Controller
	recorder *MockExporterProviderInterfaceMockRecorder
}

// MockExporterProviderInterfaceMockRecorder is the mock recorder for MockExporterProviderInterface.
type MockExporterProviderInterfaceMockRecorder struct {
	mock *MockExporterProviderInterface
}

// NewMockExporterProviderInterface creates a new mock instance.
func NewMockExporterProviderInterface(ctrl *gomock.Controller) *MockExporterProviderInterface {
	mock := &MockExporterProviderInterface{ctrl: ctrl}
	mock.recorder = &MockExporterProviderInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockExporterProviderInterface) EXPECT() *MockExporterProviderInterfaceMockRecorder {
	return m.recorder
}

// Save mocks base method.
func (m *MockExporterProviderInterface) Save(fileName string, data []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Save", fileName, data)
	ret0, _ := ret[0].(error)
	return ret0
}

// Save indicates an expected call of Save.
func (mr *MockExporterProviderInterfaceMockRecorder) Save(fileName, data interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Save", reflect.TypeOf((*MockExporterProviderInterface)(nil).Save), fileName, data)
}

// CanSave mocks base method.
func (m *MockExporterProviderInterface) CanSave(exporterProviderName string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CanSave", exporterProviderName)
	ret0, _ := ret[0].(bool)
	return ret0
}

// CanSave indicates an expected call of CanSave.
func (mr *MockExporterProviderInterfaceMockRecorder) CanSave(exporterProviderName interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CanSave", reflect.TypeOf((*MockExporterProviderInterface)(nil).CanSave), exporterProviderName)
}
