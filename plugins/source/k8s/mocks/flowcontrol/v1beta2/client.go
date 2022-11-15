// Code generated by MockGen. DO NOT EDIT.
// Source: k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2 (interfaces: FlowcontrolV1beta2Interface)

// Package v1 is a generated GoMock package.
package v1

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	v1beta2 "k8s.io/client-go/kubernetes/typed/flowcontrol/v1beta2"
	rest "k8s.io/client-go/rest"
)

// MockFlowcontrolV1beta2Interface is a mock of FlowcontrolV1beta2Interface interface.
type MockFlowcontrolV1beta2Interface struct {
	ctrl     *gomock.Controller
	recorder *MockFlowcontrolV1beta2InterfaceMockRecorder
}

// MockFlowcontrolV1beta2InterfaceMockRecorder is the mock recorder for MockFlowcontrolV1beta2Interface.
type MockFlowcontrolV1beta2InterfaceMockRecorder struct {
	mock *MockFlowcontrolV1beta2Interface
}

// NewMockFlowcontrolV1beta2Interface creates a new mock instance.
func NewMockFlowcontrolV1beta2Interface(ctrl *gomock.Controller) *MockFlowcontrolV1beta2Interface {
	mock := &MockFlowcontrolV1beta2Interface{ctrl: ctrl}
	mock.recorder = &MockFlowcontrolV1beta2InterfaceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockFlowcontrolV1beta2Interface) EXPECT() *MockFlowcontrolV1beta2InterfaceMockRecorder {
	return m.recorder
}

// FlowSchemas mocks base method.
func (m *MockFlowcontrolV1beta2Interface) FlowSchemas() v1beta2.FlowSchemaInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FlowSchemas")
	ret0, _ := ret[0].(v1beta2.FlowSchemaInterface)
	return ret0
}

// FlowSchemas indicates an expected call of FlowSchemas.
func (mr *MockFlowcontrolV1beta2InterfaceMockRecorder) FlowSchemas() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FlowSchemas", reflect.TypeOf((*MockFlowcontrolV1beta2Interface)(nil).FlowSchemas))
}

// PriorityLevelConfigurations mocks base method.
func (m *MockFlowcontrolV1beta2Interface) PriorityLevelConfigurations() v1beta2.PriorityLevelConfigurationInterface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PriorityLevelConfigurations")
	ret0, _ := ret[0].(v1beta2.PriorityLevelConfigurationInterface)
	return ret0
}

// PriorityLevelConfigurations indicates an expected call of PriorityLevelConfigurations.
func (mr *MockFlowcontrolV1beta2InterfaceMockRecorder) PriorityLevelConfigurations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PriorityLevelConfigurations", reflect.TypeOf((*MockFlowcontrolV1beta2Interface)(nil).PriorityLevelConfigurations))
}

// RESTClient mocks base method.
func (m *MockFlowcontrolV1beta2Interface) RESTClient() rest.Interface {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RESTClient")
	ret0, _ := ret[0].(rest.Interface)
	return ret0
}

// RESTClient indicates an expected call of RESTClient.
func (mr *MockFlowcontrolV1beta2InterfaceMockRecorder) RESTClient() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RESTClient", reflect.TypeOf((*MockFlowcontrolV1beta2Interface)(nil).RESTClient))
}