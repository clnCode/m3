// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/aggregator/client/queue.go

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

	"github.com/m3db/m3/src/metrics/encoding/protobuf"

	"github.com/golang/mock/gomock"
)

// MockinstanceQueue is a mock of instanceQueue interface
type MockinstanceQueue struct {
	ctrl     *gomock.Controller
	recorder *MockinstanceQueueMockRecorder
}

// MockinstanceQueueMockRecorder is the mock recorder for MockinstanceQueue
type MockinstanceQueueMockRecorder struct {
	mock *MockinstanceQueue
}

// NewMockinstanceQueue creates a new mock instance
func NewMockinstanceQueue(ctrl *gomock.Controller) *MockinstanceQueue {
	mock := &MockinstanceQueue{ctrl: ctrl}
	mock.recorder = &MockinstanceQueueMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockinstanceQueue) EXPECT() *MockinstanceQueueMockRecorder {
	return m.recorder
}

// Enqueue mocks base method
func (m *MockinstanceQueue) Enqueue(buf protobuf.Buffer) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Enqueue", buf)
	ret0, _ := ret[0].(error)
	return ret0
}

// Enqueue indicates an expected call of Enqueue
func (mr *MockinstanceQueueMockRecorder) Enqueue(buf interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Enqueue", reflect.TypeOf((*MockinstanceQueue)(nil).Enqueue), buf)
}

// Close mocks base method
func (m *MockinstanceQueue) Close() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (mr *MockinstanceQueueMockRecorder) Close() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Close", reflect.TypeOf((*MockinstanceQueue)(nil).Close))
}
