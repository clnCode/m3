// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/m3db/m3/src/aggregator/tools/deploy/validator.go

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

// Package deploy is a generated GoMock package.
package deploy

import (
	"reflect"

	"github.com/golang/mock/gomock"
)

// MockvalidatorFactory is a mock of validatorFactory interface
type MockvalidatorFactory struct {
	ctrl     *gomock.Controller
	recorder *MockvalidatorFactoryMockRecorder
}

// MockvalidatorFactoryMockRecorder is the mock recorder for MockvalidatorFactory
type MockvalidatorFactoryMockRecorder struct {
	mock *MockvalidatorFactory
}

// NewMockvalidatorFactory creates a new mock instance
func NewMockvalidatorFactory(ctrl *gomock.Controller) *MockvalidatorFactory {
	mock := &MockvalidatorFactory{ctrl: ctrl}
	mock.recorder = &MockvalidatorFactoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockvalidatorFactory) EXPECT() *MockvalidatorFactoryMockRecorder {
	return m.recorder
}

// ValidatorFor mocks base method
func (m *MockvalidatorFactory) ValidatorFor(instance instanceMetadata, group *instanceGroup, targetType targetType) validator {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ValidatorFor", instance, group, targetType)
	ret0, _ := ret[0].(validator)
	return ret0
}

// ValidatorFor indicates an expected call of ValidatorFor
func (mr *MockvalidatorFactoryMockRecorder) ValidatorFor(instance, group, targetType interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ValidatorFor", reflect.TypeOf((*MockvalidatorFactory)(nil).ValidatorFor), instance, group, targetType)
}
