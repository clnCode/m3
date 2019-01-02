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

// This file was automatically generated by genny.
// Any changes will be lost if this file is regenerated.
// see https://github.com/mauricelam/genny

package client

import (
	"fmt"

	"runtime/debug"

	"sync"

	"testing"

	"github.com/stretchr/testify/require"
)

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

// fetchTaggedOpEqualsFn allows users to override equality checks
// for `fetchTaggedOp` instances.
type fetchTaggedOpEqualsFn func(a, b *fetchTaggedOp) bool

// fetchTaggedOpGetHookFn allows users to override properties on items
// retrieved from the backing pools before returning in the Get()
// path.
type fetchTaggedOpGetHookFn func(*fetchTaggedOp) *fetchTaggedOp

// leakcheckFetchTaggedOpPoolOpts allows users to override default behaviour.
type leakcheckFetchTaggedOpPoolOpts struct {
	DisallowUntrackedPuts bool
	EqualsFn              fetchTaggedOpEqualsFn
	GetHookFn             fetchTaggedOpGetHookFn
}

// newLeakcheckFetchTaggedOpPool returns a new leakcheckFetchTaggedOpPool.
func newLeakcheckFetchTaggedOpPool(opts leakcheckFetchTaggedOpPoolOpts, backingPool fetchTaggedOpPool) *leakcheckFetchTaggedOpPool {
	if opts.EqualsFn == nil {
		// NB(prateek): fall-back to == in the worst case
		opts.EqualsFn = func(a, b *fetchTaggedOp) bool {
			return a == b
		}
	}
	return &leakcheckFetchTaggedOpPool{opts: opts, fetchTaggedOpPool: backingPool}
}

// leakcheckFetchTaggedOpPool wraps the underlying fetchTaggedOpPool to make it easier to
// track leaks/allocs.
type leakcheckFetchTaggedOpPool struct {
	sync.Mutex
	fetchTaggedOpPool
	NumGets      int
	NumPuts      int
	PendingItems []leakcheckFetchTaggedOp
	AllGetItems  []leakcheckFetchTaggedOp

	opts leakcheckFetchTaggedOpPoolOpts
}

// leakcheckFetchTaggedOp wraps `fetchTaggedOp` instances along with their last Get() paths.
type leakcheckFetchTaggedOp struct {
	Value         *fetchTaggedOp
	GetStacktrace []byte // GetStacktrace is the stacktrace for the Get() of this item
}

func (p *leakcheckFetchTaggedOpPool) Init() {
	p.Lock()
	defer p.Unlock()
	p.fetchTaggedOpPool.Init()
}

func (p *leakcheckFetchTaggedOpPool) Get() *fetchTaggedOp {
	p.Lock()
	defer p.Unlock()

	e := p.fetchTaggedOpPool.Get()
	if fn := p.opts.GetHookFn; fn != nil {
		e = fn(e)
	}

	p.NumGets++
	item := leakcheckFetchTaggedOp{
		Value:         e,
		GetStacktrace: debug.Stack(),
	}
	p.PendingItems = append(p.PendingItems, item)
	p.AllGetItems = append(p.AllGetItems, item)

	return e
}

func (p *leakcheckFetchTaggedOpPool) Put(value *fetchTaggedOp) {
	p.Lock()
	defer p.Unlock()

	idx := -1
	for i, item := range p.PendingItems {
		if p.opts.EqualsFn(item.Value, value) {
			idx = i
			break
		}
	}

	if idx == -1 && p.opts.DisallowUntrackedPuts {
		panic(fmt.Errorf("untracked object (%v) returned to pool", value))
	}

	if idx != -1 {
		// update slice
		p.PendingItems = append(p.PendingItems[:idx], p.PendingItems[idx+1:]...)
	}
	p.NumPuts++

	p.fetchTaggedOpPool.Put(value)
}

// Check ensures there are no leaks.
func (p *leakcheckFetchTaggedOpPool) Check(t *testing.T) {
	p.Lock()
	defer p.Unlock()

	require.Equal(t, p.NumGets, p.NumPuts)
	require.Empty(t, p.PendingItems)
}

type leakcheckFetchTaggedOpFn func(e leakcheckFetchTaggedOp)

// CheckExtended ensures there are no leaks, and executes the specified fn
func (p *leakcheckFetchTaggedOpPool) CheckExtended(t *testing.T, fn leakcheckFetchTaggedOpFn) {
	p.Check(t)
	p.Lock()
	defer p.Unlock()
	for _, e := range p.AllGetItems {
		fn(e)
	}
}
