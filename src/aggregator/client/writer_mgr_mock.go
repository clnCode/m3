// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/aggregator/client/writer_mgr.go

// Copyright (c) 2019 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

// Package client is a generated GoMock package.
package client

import (
	"reflect"

	"github.com/m3db/m3/src/cluster/placement"

	"github.com/golang/mock/gomock"
)

// MockinstanceWriterManager is a mock of instanceWriterManager interface
type MockinstanceWriterManager struct {
	ctrl     *gomock.Controller
	recorder *MockinstanceWriterManagerMockRecorder
}

// MockinstanceWriterManagerMockRecorder is the mock recorder for MockinstanceWriterManager
type MockinstanceWriterManagerMockRecorder struct {
	mock *MockinstanceWriterManager
}

// NewMockinstanceWriterManager creates a new mock instance
func NewMockinstanceWriterManager(ctrl *gomock.Controller) *MockinstanceWriterManager {
	mock := &MockinstanceWriterManager{ctrl: ctrl}
	mock.recorder = &MockinstanceWriterManagerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockinstanceWriterManager) EXPECT() *MockinstanceWriterManagerMockRecorder {
	return m.recorder
}

// AddInstances mocks base method
func (m *MockinstanceWriterManager) AddInstances(instances []placement.Instance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddInstances", instances)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddInstances indicates an expected call of AddInstances
func (mr *MockinstanceWriterManagerMockRecorder) AddInstances(instances interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddInstances", reflect.TypeOf((*MockinstanceWriterManager)(nil).AddInstances), instances)
}

// RemoveInstances mocks base method
func (m *MockinstanceWriterManager) RemoveInstances(instances []placement.Instance) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveInstances", instances)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveInstances indicates an expected call of RemoveInstances
func (mr *MockinstanceWriterManagerMockRecorder) RemoveInstances(instances interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveInstances", reflect.TypeOf((*MockinstanceWriterManager)(nil).RemoveInstances), instances)
}

// Write mocks base method
func (m *MockinstanceWriterManager) Write(instance placement.Instance, shardID uint32, payload payloadUnion) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Write", instance, shardID, payload)
	ret0, _ := ret[0].(error)
	return ret0
}

// Write indicates an expected call of Write
func (mr *MockinstanceWriterManagerMockRecorder) Write(instance, shardID, payload interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Write", reflect.TypeOf((*MockinstanceWriterManager)(nil).Write), instance, shardID, payload)
}

// Flush mocks base method
func (m *MockinstanceWriterManager) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockinstanceWriterManagerMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockinstanceWriterManager)(nil).Flush))
}

// Close mocks base method
func (m *MockinstanceWriterManager) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockinstanceWriterManagerMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockinstanceWriterManager)(nil).Close))
}
