// Code generated by MockGen. DO NOT EDIT.
// Source: internal/repository/service_repository_interface.go

// Package pdf is a generated GoMock package.
package pdf

import (
	reflect "reflect"

	model "fuegobyp-billing.com/internal/model"
	gomock "github.com/golang/mock/gomock"
)

// MockServiceRepositoryInterface is a mock of ServiceRepositoryInterface interface.
type MockServiceRepositoryInterface struct {
	ctrl     *gomock.Controller
	recorder *MockServiceRepositoryInterfaceMockRecorder
}

// MockServiceRepositoryInterfaceMockRecorder is the mock recorder for MockServiceRepositoryInterface.
type MockServiceRepositoryInterfaceMockRecorder struct {
	mock *MockServiceRepositoryInterface
}

// NewMockServiceRepositoryInterface creates a new mock instance.
func NewMockServiceRepositoryInterface(ctrl *gomock.Controller) *MockServiceRepositoryInterface {
	mock := &MockServiceRepositoryInterface{ctrl: ctrl}
	mock.recorder = &MockServiceRepositoryInterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockServiceRepositoryInterface) EXPECT() *MockServiceRepositoryInterfaceMockRecorder {
	return m.recorder
}

// Request mocks base method.
func (m *MockServiceRepositoryInterface) Request(id string) (model.Service, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", id)
	ret0, _ := ret[0].(model.Service)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockServiceRepositoryInterfaceMockRecorder) Request(id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockServiceRepositoryInterface)(nil).Request), id)
}