// Code generated by MockGen. DO NOT EDIT.
// Source: main.go

// Package mock_main is a generated GoMock package.
package mock_main

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockHttpRequester is a mock of HttpRequester interface.
type MockHttpRequester struct {
	ctrl     *gomock.Controller
	recorder *MockHttpRequesterMockRecorder
}

// MockHttpRequesterMockRecorder is the mock recorder for MockHttpRequester.
type MockHttpRequesterMockRecorder struct {
	mock *MockHttpRequester
}

// NewMockHttpRequester creates a new mock instance.
func NewMockHttpRequester(ctrl *gomock.Controller) *MockHttpRequester {
	mock := &MockHttpRequester{ctrl: ctrl}
	mock.recorder = &MockHttpRequesterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockHttpRequester) EXPECT() *MockHttpRequesterMockRecorder {
	return m.recorder
}

// httpGet mocks base method.
func (m *MockHttpRequester) httpGet(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "httpGet", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// httpGet indicates an expected call of httpGet.
func (mr *MockHttpRequesterMockRecorder) HttpGet(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "httpGet", reflect.TypeOf((*MockHttpRequester)(nil).httpGet), arg0)
}

// MockUsersService is a mock of UsersService interface.
type MockUsersService struct {
	ctrl     *gomock.Controller
	recorder *MockUsersServiceMockRecorder
}

// MockUsersServiceMockRecorder is the mock recorder for MockUsersService.
type MockUsersServiceMockRecorder struct {
	mock *MockUsersService
}

// NewMockUsersService creates a new mock instance.
func NewMockUsersService(ctrl *gomock.Controller) *MockUsersService {
	mock := &MockUsersService{ctrl: ctrl}
	mock.recorder = &MockUsersServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockUsersService) EXPECT() *MockUsersServiceMockRecorder {
	return m.recorder
}

// getUsers mocks base method.
func (m *MockUsersService) getUsers() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "getUsers")
}

// getUsers indicates an expected call of getUsers.
func (mr *MockUsersServiceMockRecorder) getUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "getUsers", reflect.TypeOf((*MockUsersService)(nil).getUsers))
}