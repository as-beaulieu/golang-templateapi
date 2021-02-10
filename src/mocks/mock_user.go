// Code generated by MockGen. DO NOT EDIT.
// Source: src/service/user.go

// Package mocks is a generated GoMock package.
package mocks

import (
	models "TemplateApi/src/models"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockUserOperator is a mock of UserOperator interface
type MockUserOperator struct {
	ctrl     *gomock.Controller
	recorder *MockUserOperatorMockRecorder
}

// MockUserOperatorMockRecorder is the mock recorder for MockUserOperator
type MockUserOperatorMockRecorder struct {
	mock *MockUserOperator
}

// NewMockUserOperator creates a new mock instance
func NewMockUserOperator(ctrl *gomock.Controller) *MockUserOperator {
	mock := &MockUserOperator{ctrl: ctrl}
	mock.recorder = &MockUserOperatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockUserOperator) EXPECT() *MockUserOperatorMockRecorder {
	return m.recorder
}

// CreateUser mocks base method
func (m *MockUserOperator) CreateUser(user models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateUser", user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateUser indicates an expected call of CreateUser
func (mr *MockUserOperatorMockRecorder) CreateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateUser", reflect.TypeOf((*MockUserOperator)(nil).CreateUser), user)
}

// GetUsers mocks base method
func (m *MockUserOperator) GetUsers() ([]*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUsers")
	ret0, _ := ret[0].([]*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUsers indicates an expected call of GetUsers
func (mr *MockUserOperatorMockRecorder) GetUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUsers", reflect.TypeOf((*MockUserOperator)(nil).GetUsers))
}

// GetUserByID mocks base method
func (m *MockUserOperator) GetUserByID(userID string) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUserByID", userID)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUserByID indicates an expected call of GetUserByID
func (mr *MockUserOperatorMockRecorder) GetUserByID(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUserByID", reflect.TypeOf((*MockUserOperator)(nil).GetUserByID), userID)
}

// UpdateUser mocks base method
func (m *MockUserOperator) UpdateUser(user models.User) (*models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateUser", user)
	ret0, _ := ret[0].(*models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// UpdateUser indicates an expected call of UpdateUser
func (mr *MockUserOperatorMockRecorder) UpdateUser(user interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateUser", reflect.TypeOf((*MockUserOperator)(nil).UpdateUser), user)
}

// DeleteUser mocks base method
func (m *MockUserOperator) DeleteUser(userID string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeleteUser", userID)
	ret0, _ := ret[0].(error)
	return ret0
}

// DeleteUser indicates an expected call of DeleteUser
func (mr *MockUserOperatorMockRecorder) DeleteUser(userID interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteUser", reflect.TypeOf((*MockUserOperator)(nil).DeleteUser), userID)
}