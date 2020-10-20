// Code generated by MockGen. DO NOT EDIT.
// Source: internal/services/format_int.go

// Package pdf is a generated GoMock package.
package pdf

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockFormatIntInterface is a mock of FormatIntInterface interface.
type MockFormatIntInterface struct {
	ctrl     *gomock.Controller
	recorder *MockFormatIntInterfaceMockRecorder
}

// MockFormatIntInterfaceMockRecorder is the mock recorder for MockFormatIntInterface.
type MockFormatIntInterfaceMockRecorder struct {
	mock *MockFormatIntInterface
}

// NewMockFormatIntInterface creates a new mock instance.
func NewMockFormatIntInterface(ctrl *gomock.Controller) *MockFormatIntInterface {
	mock := &MockFormatIntInterface{ctrl: ctrl}
	mock.recorder = &MockFormatIntInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFormatIntInterface) EXPECT() *MockFormatIntInterfaceMockRecorder {
	return m.recorder
}

// IntToStringFrenchFormat mocks base method.
func (m *MockFormatIntInterface) IntToStringFrenchFormat(value int) string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IntToStringFrenchFormat", value)
	ret0, _ := ret[0].(string)
	return ret0
}

// IntToStringFrenchFormat indicates an expected call of IntToStringFrenchFormat.
func (mr *MockFormatIntInterfaceMockRecorder) IntToStringFrenchFormat(value interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IntToStringFrenchFormat", reflect.TypeOf((*MockFormatIntInterface)(nil).IntToStringFrenchFormat), value)
}