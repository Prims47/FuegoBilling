// Code generated by MockGen. DO NOT EDIT.
// Source: internal/adapter/customer_adapter_interface.go

// Package repository is a generated GoMock package.
package repository

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	adapter "github.com/prims47/FuegoBilling/internal/adapter"
)

// MockCustomerAdapterInterface is a mock of CustomerAdapterInterface interface.
type MockCustomerAdapterInterface struct {
	ctrl     *gomock.Controller
	recorder *MockCustomerAdapterInterfaceMockRecorder
}

// MockCustomerAdapterInterfaceMockRecorder is the mock recorder for MockCustomerAdapterInterface.
type MockCustomerAdapterInterfaceMockRecorder struct {
	mock *MockCustomerAdapterInterface
}

// NewMockCustomerAdapterInterface creates a new mock instance.
func NewMockCustomerAdapterInterface(ctrl *gomock.Controller) *MockCustomerAdapterInterface {
	mock := &MockCustomerAdapterInterface{ctrl: ctrl}
	mock.recorder = &MockCustomerAdapterInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCustomerAdapterInterface) EXPECT() *MockCustomerAdapterInterfaceMockRecorder {
	return m.recorder
}

// Request mocks base method.
func (m *MockCustomerAdapterInterface) Request(id string) (adapter.CustomerAdapterResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", id)
	ret0, _ := ret[0].(adapter.CustomerAdapterResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockCustomerAdapterInterfaceMockRecorder) Request(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockCustomerAdapterInterface)(nil).Request), id)
}
