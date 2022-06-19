// Code generated by MockGen. DO NOT EDIT.
// Source: homeworks/homework19/ldapAuth.go

// Package homework19 is a generated GoMock package.
package homework19

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// Mockldap is a mock of ldap interface.
type Mockldap struct {
	ctrl     *gomock.Controller
	recorder *MockldapMockRecorder
}

// MockldapMockRecorder is the mock recorder for Mockldap.
type MockldapMockRecorder struct {
	mock *Mockldap
}

// NewMockldap creates a new mock instance.
func NewMockldap(ctrl *gomock.Controller) *Mockldap {
	mock := &Mockldap{ctrl: ctrl}
	mock.recorder = &MockldapMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *Mockldap) EXPECT() *MockldapMockRecorder {
	return m.recorder
}

// ldapCheck mocks base method.
func (m *Mockldap) ldapCheck(username string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ldapCheck", username)
	ret0, _ := ret[0].(bool)
	return ret0
}

// ldapCheck indicates an expected call of ldapCheck.
func (mr *MockldapMockRecorder) ldapCheck(username interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ldapCheck", reflect.TypeOf((*Mockldap)(nil).ldapCheck), username)
}
