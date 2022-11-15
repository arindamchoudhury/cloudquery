// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/authorization/v1beta1 (interfaces: AuthorizationV1beta1Interface)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta1 "k8s.io/client-go/kubernetes/typed/authorization/v1beta1"
	rest "k8s.io/client-go/rest"
)

// MockAuthorizationV1beta1Interface is a mock of AuthorizationV1beta1Interface interface.
type MockAuthorizationV1beta1Interface struct {
	ctrl     *gomock.Controller
	recorder *MockAuthorizationV1beta1InterfaceMockRecorder
}

// MockAuthorizationV1beta1InterfaceMockRecorder is the mock recorder for MockAuthorizationV1beta1Interface.
type MockAuthorizationV1beta1InterfaceMockRecorder struct {
	mock *MockAuthorizationV1beta1Interface
}

// NewMockAuthorizationV1beta1Interface creates a new mock instance.
func NewMockAuthorizationV1beta1Interface(ctrl *gomock.Controller) *MockAuthorizationV1beta1Interface {
	mock := &MockAuthorizationV1beta1Interface{ctrl: ctrl}
	mock.recorder = &MockAuthorizationV1beta1InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuthorizationV1beta1Interface) EXPECT() *MockAuthorizationV1beta1InterfaceMockRecorder {
	return m.recorder
}

// LocalSubjectAccessReviews mocks base method.
func (m *MockAuthorizationV1beta1Interface) LocalSubjectAccessReviews(arg0 string) v1beta1.LocalSubjectAccessReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LocalSubjectAccessReviews", arg0)
	ret0, _ := ret[0].(v1beta1.LocalSubjectAccessReviewInterface)
	return ret0
}

// LocalSubjectAccessReviews indicates an expected call of LocalSubjectAccessReviews.
func (mr *MockAuthorizationV1beta1InterfaceMockRecorder) LocalSubjectAccessReviews(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LocalSubjectAccessReviews", reflect.TypeOf((*MockAuthorizationV1beta1Interface)(nil).LocalSubjectAccessReviews), arg0)
}

// RESTClient mocks base method.
func (m *MockAuthorizationV1beta1Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockAuthorizationV1beta1InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockAuthorizationV1beta1Interface)(nil).RESTClient))
}

// SelfSubjectAccessReviews mocks base method.
func (m *MockAuthorizationV1beta1Interface) SelfSubjectAccessReviews() v1beta1.SelfSubjectAccessReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelfSubjectAccessReviews")
	ret0, _ := ret[0].(v1beta1.SelfSubjectAccessReviewInterface)
	return ret0
}

// SelfSubjectAccessReviews indicates an expected call of SelfSubjectAccessReviews.
func (mr *MockAuthorizationV1beta1InterfaceMockRecorder) SelfSubjectAccessReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelfSubjectAccessReviews", reflect.TypeOf((*MockAuthorizationV1beta1Interface)(nil).SelfSubjectAccessReviews))
}

// SelfSubjectRulesReviews mocks base method.
func (m *MockAuthorizationV1beta1Interface) SelfSubjectRulesReviews() v1beta1.SelfSubjectRulesReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelfSubjectRulesReviews")
	ret0, _ := ret[0].(v1beta1.SelfSubjectRulesReviewInterface)
	return ret0
}

// SelfSubjectRulesReviews indicates an expected call of SelfSubjectRulesReviews.
func (mr *MockAuthorizationV1beta1InterfaceMockRecorder) SelfSubjectRulesReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelfSubjectRulesReviews", reflect.TypeOf((*MockAuthorizationV1beta1Interface)(nil).SelfSubjectRulesReviews))
}

// SubjectAccessReviews mocks base method.
func (m *MockAuthorizationV1beta1Interface) SubjectAccessReviews() v1beta1.SubjectAccessReviewInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubjectAccessReviews")
	ret0, _ := ret[0].(v1beta1.SubjectAccessReviewInterface)
	return ret0
}

// SubjectAccessReviews indicates an expected call of SubjectAccessReviews.
func (mr *MockAuthorizationV1beta1InterfaceMockRecorder) SubjectAccessReviews() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubjectAccessReviews", reflect.TypeOf((*MockAuthorizationV1beta1Interface)(nil).SubjectAccessReviews))
}