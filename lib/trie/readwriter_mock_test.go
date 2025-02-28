// Code generated by MockGen. DO NOT EDIT.
// Source: io (interfaces: ReadWriter)

// Package trie is a generated GoMock package.
package trie

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockReadWriter is a mock of ReadWriter interface.
type MockReadWriter struct {
	ctrl     *gomock.Controller
	recorder *MockReadWriterMockRecorder
}

// MockReadWriterMockRecorder is the mock recorder for MockReadWriter.
type MockReadWriterMockRecorder struct {
	mock *MockReadWriter
}

// NewMockReadWriter creates a new mock instance.
func NewMockReadWriter(ctrl *gomock.Controller) *MockReadWriter {
	mock := &MockReadWriter{ctrl: ctrl}
	mock.recorder = &MockReadWriterMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockReadWriter) EXPECT() *MockReadWriterMockRecorder {
	return m.recorder
}

// Read mocks base method.
func (m *MockReadWriter) Read(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Read", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Read indicates an expected call of Read.
func (mr *MockReadWriterMockRecorder) Read(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Read", reflect.TypeOf((*MockReadWriter)(nil).Read), arg0)
}

// Write mocks base method.
func (m *MockReadWriter) Write(arg0 []byte) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Write indicates an expected call of Write.
func (mr *MockReadWriterMockRecorder) Write(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockReadWriter)(nil).Write), arg0)
}
