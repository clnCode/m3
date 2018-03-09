// Copyright (c) 2018 Uber Technologies, Inc.
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

// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3em/os/exec/types.go

package exec

import (
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockProcessListener is a mock of ProcessListener interface
type MockProcessListener struct {
	ctrl     *gomock.Controller
	recorder *MockProcessListenerMockRecorder
}

// MockProcessListenerMockRecorder is the mock recorder for MockProcessListener
type MockProcessListenerMockRecorder struct {
	mock *MockProcessListener
}

// NewMockProcessListener creates a new mock instance
func NewMockProcessListener(ctrl *gomock.Controller) *MockProcessListener {
	mock := &MockProcessListener{ctrl: ctrl}
	mock.recorder = &MockProcessListenerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockProcessListener) EXPECT() *MockProcessListenerMockRecorder {
	return _m.recorder
}

// OnComplete mocks base method
func (_m *MockProcessListener) OnComplete() {
	_m.ctrl.Call(_m, "OnComplete")
}

// OnComplete indicates an expected call of OnComplete
func (_mr *MockProcessListenerMockRecorder) OnComplete() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "OnComplete", reflect.TypeOf((*MockProcessListener)(nil).OnComplete))
}

// OnError mocks base method
func (_m *MockProcessListener) OnError(_param0 error) {
	_m.ctrl.Call(_m, "OnError", _param0)
}

// OnError indicates an expected call of OnError
func (_mr *MockProcessListenerMockRecorder) OnError(arg0 interface{}) *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "OnError", reflect.TypeOf((*MockProcessListener)(nil).OnError), arg0)
}

// MockProcessMonitor is a mock of ProcessMonitor interface
type MockProcessMonitor struct {
	ctrl     *gomock.Controller
	recorder *MockProcessMonitorMockRecorder
}

// MockProcessMonitorMockRecorder is the mock recorder for MockProcessMonitor
type MockProcessMonitorMockRecorder struct {
	mock *MockProcessMonitor
}

// NewMockProcessMonitor creates a new mock instance
func NewMockProcessMonitor(ctrl *gomock.Controller) *MockProcessMonitor {
	mock := &MockProcessMonitor{ctrl: ctrl}
	mock.recorder = &MockProcessMonitorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (_m *MockProcessMonitor) EXPECT() *MockProcessMonitorMockRecorder {
	return _m.recorder
}

// Start mocks base method
func (_m *MockProcessMonitor) Start() error {
	ret := _m.ctrl.Call(_m, "Start")
	ret0, _ := ret[0].(error)
	return ret0
}

// Start indicates an expected call of Start
func (_mr *MockProcessMonitorMockRecorder) Start() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Start", reflect.TypeOf((*MockProcessMonitor)(nil).Start))
}

// Stop mocks base method
func (_m *MockProcessMonitor) Stop() error {
	ret := _m.ctrl.Call(_m, "Stop")
	ret0, _ := ret[0].(error)
	return ret0
}

// Stop indicates an expected call of Stop
func (_mr *MockProcessMonitorMockRecorder) Stop() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Stop", reflect.TypeOf((*MockProcessMonitor)(nil).Stop))
}

// Running mocks base method
func (_m *MockProcessMonitor) Running() bool {
	ret := _m.ctrl.Call(_m, "Running")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Running indicates an expected call of Running
func (_mr *MockProcessMonitorMockRecorder) Running() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Running", reflect.TypeOf((*MockProcessMonitor)(nil).Running))
}

// Err mocks base method
func (_m *MockProcessMonitor) Err() error {
	ret := _m.ctrl.Call(_m, "Err")
	ret0, _ := ret[0].(error)
	return ret0
}

// Err indicates an expected call of Err
func (_mr *MockProcessMonitorMockRecorder) Err() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Err", reflect.TypeOf((*MockProcessMonitor)(nil).Err))
}

// Close mocks base method
func (_m *MockProcessMonitor) Close() error {
	ret := _m.ctrl.Call(_m, "Close")
	ret0, _ := ret[0].(error)
	return ret0
}

// Close indicates an expected call of Close
func (_mr *MockProcessMonitorMockRecorder) Close() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "Close", reflect.TypeOf((*MockProcessMonitor)(nil).Close))
}

// StdoutPath mocks base method
func (_m *MockProcessMonitor) StdoutPath() string {
	ret := _m.ctrl.Call(_m, "StdoutPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// StdoutPath indicates an expected call of StdoutPath
func (_mr *MockProcessMonitorMockRecorder) StdoutPath() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "StdoutPath", reflect.TypeOf((*MockProcessMonitor)(nil).StdoutPath))
}

// StderrPath mocks base method
func (_m *MockProcessMonitor) StderrPath() string {
	ret := _m.ctrl.Call(_m, "StderrPath")
	ret0, _ := ret[0].(string)
	return ret0
}

// StderrPath indicates an expected call of StderrPath
func (_mr *MockProcessMonitorMockRecorder) StderrPath() *gomock.Call {
	return _mr.mock.ctrl.RecordCallWithMethodType(_mr.mock, "StderrPath", reflect.TypeOf((*MockProcessMonitor)(nil).StderrPath))
}
