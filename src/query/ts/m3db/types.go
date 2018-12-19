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
	"time"

	"github.com/m3db/m3/src/dbnode/encoding"
	"github.com/m3db/m3/src/query/models"
	"github.com/m3db/m3/src/query/ts/m3db/consolidators"
)

// Options describes the options for encoded block converters.
type Options interface {
	// SetLookbackDuration sets the lookback duration.
	SetLookbackDuration(time.Duration) Options
	// GetLookbackDuration returns the lookback duration.
	GetLookbackDuration() time.Duration
	// SetLookbackDuration sets the consolidation function for the converter.
	SetConsolidationFunc(consolidators.ConsolidationFunc) Options
	// GetLookbackDuration returns the consolidation function.
	GetConsolidationFunc() consolidators.ConsolidationFunc
	// SetLookbackDuration sets the tag options for the converter.
	SetTagOptions(models.TagOptions) Options
	// GetTagOptions returns the tag options.
	GetTagOptions() models.TagOptions
	// SetBounds sets the bounds for the encoded block.
	SetBounds(models.Bounds) Options
	// GetBounds returns the bounds for the encoded block.
	GetBounds() models.Bounds
	// SetIterAlloc sets the iterator allocator.
	SetIterAlloc(encoding.ReaderIteratorAllocate) Options
	// GetIterAlloc returns the reader iterator allocator.
	GetIterAlloc() encoding.ReaderIteratorAllocate
	// SetIteratorPools sets the iterator pools for the converter.
	SetIteratorPools(encoding.IteratorPools) Options
	// GetIteratorPools returns the iterator pools for the converter.
	GetIteratorPools() encoding.IteratorPools

	// Validate ensures that the given block options are valid.
	Validate() error
}
