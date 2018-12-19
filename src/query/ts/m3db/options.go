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

package m3db

import (
	"errors"
	"fmt"
	"io"
	"time"

	"github.com/m3db/m3/src/dbnode/encoding"
	"github.com/m3db/m3/src/dbnode/encoding/m3tsz"
	"github.com/m3db/m3/src/query/models"
	"github.com/m3db/m3/src/query/pools"
	"github.com/m3db/m3/src/query/ts/m3db/consolidators"
)

var (
	defaultLookbackDuration = time.Duration(0)
	defaultConsolidationFn  = consolidators.TakeLast
	defaultIterAlloc        = func(r io.Reader) encoding.ReaderIterator {
		return m3tsz.NewReaderIterator(r, m3tsz.DefaultIntOptimizationEnabled, encoding.NewOptions())
	}
)

type encodedBlockOptions struct {
	lookbackDuration time.Duration
	consolidationFn  consolidators.ConsolidationFunc
	tagOptions       models.TagOptions
	bounds           *models.Bounds
	iterAlloc        encoding.ReaderIteratorAllocate
	pools            encoding.IteratorPools
}

// NewOptions creates a default encoded block options which dictates how
// encoded blocks are generated.
func NewOptions() Options {
	return &encodedBlockOptions{
		lookbackDuration: defaultLookbackDuration,
		consolidationFn:  defaultConsolidationFn,
		tagOptions:       models.NewTagOptions(),
		iterAlloc:        defaultIterAlloc,
		pools:            pools.BuildIteratorPools(),
	}
}

func (o *encodedBlockOptions) SetLookbackDuration(
	lookback time.Duration,
) Options {
	opts := *o
	opts.lookbackDuration = lookback
	return &opts
}

func (o *encodedBlockOptions) GetLookbackDuration() time.Duration {
	return o.lookbackDuration
}

func (o *encodedBlockOptions) SetConsolidationFunc(
	fn consolidators.ConsolidationFunc,
) Options {
	opts := *o
	opts.consolidationFn = fn
	return &opts
}

func (o *encodedBlockOptions) GetConsolidationFunc() consolidators.ConsolidationFunc {
	return o.consolidationFn
}

func (o *encodedBlockOptions) SetTagOptions(tagOptions models.TagOptions) Options {
	opts := *o
	opts.tagOptions = tagOptions
	return &opts
}

func (o *encodedBlockOptions) GetTagOptions() models.TagOptions {
	return o.tagOptions
}

func (o *encodedBlockOptions) SetBounds(b models.Bounds) Options {
	opts := *o
	opts.bounds = &b
	return &opts
}

func (o *encodedBlockOptions) GetBounds() models.Bounds {
	return *o.bounds
}

func (o *encodedBlockOptions) SetIterAlloc(ia encoding.ReaderIteratorAllocate) Options {
	opts := *o
	opts.iterAlloc = ia
	return &opts
}

func (o *encodedBlockOptions) GetIterAlloc() encoding.ReaderIteratorAllocate {
	return o.iterAlloc
}

func (o *encodedBlockOptions) SetIteratorPools(p encoding.IteratorPools) Options {
	opts := *o
	opts.pools = p
	return &opts
}

func (o *encodedBlockOptions) GetIteratorPools() encoding.IteratorPools {
	return o.pools
}

func (o *encodedBlockOptions) Validate() error {
	if err := o.tagOptions.Validate(); err != nil {
		return fmt.Errorf("unable to validate tag options, err: %v", err)
	}

	if o.bounds == nil {
		return errors.New("unable to validate block options; bounds not set")
	}

	return nil
}
