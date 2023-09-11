// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/gardener/test-infra/pkg/util/elasticsearch (interfaces: Client)

// Package mock_elasticsearch is a generated GoMock package.
package mock_elasticsearch

import (
	context "context"
	io "io"
	reflect "reflect"

	gomock "go.uber.org/mock/gomock"
)

// MockClient is a mock of Client interface.
type MockClient struct {
	ctrl     *gomock.Controller
	recorder *MockClientMockRecorder
}

// MockClientMockRecorder is the mock recorder for MockClient.
type MockClientMockRecorder struct {
	mock *MockClient
}

// NewMockClient creates a new mock instance.
func NewMockClient(ctrl *gomock.Controller) *MockClient {
	mock := &MockClient{ctrl: ctrl}
	mock.recorder = &MockClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockClient) EXPECT() *MockClientMockRecorder {
	return m.recorder
}

// Bulk mocks base method.
func (m *MockClient) Bulk(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Bulk", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Bulk indicates an expected call of Bulk.
func (mr *MockClientMockRecorder) Bulk(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Bulk", reflect.TypeOf((*MockClient)(nil).Bulk), arg0)
}

// BulkFromFile mocks base method.
func (m *MockClient) BulkFromFile(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BulkFromFile", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// BulkFromFile indicates an expected call of BulkFromFile.
func (mr *MockClientMockRecorder) BulkFromFile(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BulkFromFile", reflect.TypeOf((*MockClient)(nil).BulkFromFile), arg0)
}

// Request mocks base method.
func (m *MockClient) Request(arg0, arg1 string, arg2 io.Reader) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Request", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Request indicates an expected call of Request.
func (mr *MockClientMockRecorder) Request(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Request", reflect.TypeOf((*MockClient)(nil).Request), arg0, arg1, arg2)
}

// RequestWithCtx mocks base method.
func (m *MockClient) RequestWithCtx(arg0 context.Context, arg1, arg2 string, arg3 io.Reader) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RequestWithCtx", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// RequestWithCtx indicates an expected call of RequestWithCtx.
func (mr *MockClientMockRecorder) RequestWithCtx(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RequestWithCtx", reflect.TypeOf((*MockClient)(nil).RequestWithCtx), arg0, arg1, arg2, arg3)
}
