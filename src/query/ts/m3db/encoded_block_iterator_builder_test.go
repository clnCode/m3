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

package m3db

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/m3db/m3/src/dbnode/encoding"
	"github.com/m3db/m3/src/query/models"
	"github.com/m3db/m3/src/query/test"
	"github.com/m3db/m3/src/query/ts/m3db/consolidators"
)

func buildCustomwhatever(
	start time.Time,
	stepSize time.Duration,
	dps [][]test.Datapoint,
) (
	encoding.SeriesIterator,
	models.Bounds,
) {
	i := rand.Int()
	iter, bounds, _ := test.BuildCustomIterator(
		dps,
		map[string]string{"a": "b", "c": fmt.Sprint(i)},
		fmt.Sprintf("abc%d", i), "namespace",
		start,
		blockSize, stepSize,
	)
	return iter, bounds
}

func genwhatever(
	stepSize time.Duration,
) (
	encoding.SeriesIterators,
	models.Bounds,
) {
	datapoints := [][][]test.Datapoint{
		{
			[]test.Datapoint{},
			[]test.Datapoint{
				{Value: 1, Offset: 0},
				{Value: 2, Offset: (time.Minute * 1) + time.Second},
				{Value: 3, Offset: (time.Minute * 2)},
				{Value: 4, Offset: (time.Minute * 3)},
				{Value: 5, Offset: (time.Minute * 4)},
				{Value: 6, Offset: (time.Minute * 5)},
			},
			[]test.Datapoint{
				{Value: 7, Offset: time.Minute * 0},
				{Value: 8, Offset: time.Minute * 5},
			},
			[]test.Datapoint{
				{Value: 9, Offset: time.Minute * 4},
			},
		},
		{
			[]test.Datapoint{
				{Value: 10, Offset: (time.Minute * 2)},
				{Value: 20, Offset: (time.Minute * 3)},
			},
			[]test.Datapoint{},
			[]test.Datapoint{
				{Value: 30, Offset: time.Minute},
			},
			[]test.Datapoint{
				{Value: 40, Offset: time.Second},
			},
		},
		{
			[]test.Datapoint{
				{Value: 100, Offset: (time.Minute * 3)},
			},
			[]test.Datapoint{
				{Value: 200, Offset: (time.Minute * 3)},
			},
			[]test.Datapoint{
				{Value: 300, Offset: (time.Minute * 3)},
			},
			[]test.Datapoint{
				{Value: 400, Offset: (time.Minute * 3)},
			},
			[]test.Datapoint{
				{Value: 500, Offset: 0},
			},
		},
	}

	iters := make([]encoding.SeriesIterator, len(datapoints))
	var (
		iter   encoding.SeriesIterator
		bounds models.Bounds
		start  = Start
	)
	for i, dps := range datapoints {
		iter, bounds = buildCustomwhatever(start, stepSize, dps)
		iters[i] = iter
	}

	return encoding.NewSeriesIterators(iters, nil), bounds
}

// BenchmarkBuilder-8   	 2000000	       903 ns/op	     512 B/op	       6 allocs/op
// BenchmarkBuilder-8   	 1000000	      1071 ns/op	     512 B/op	       6 allocs/op

//BenchmarkBuilder-8   	   30000	     47662 ns/op	   28311 B/op	     303 allocs/op
func BenchmarkBuilder(b *testing.B) {
	tb := newEncodedBlockBuilder(models.NewTagOptions(), consolidators.TakeLast)
	for i := 0; i < 100; i++ {
		iters, bounds := genwhatever(time.Minute)
		for _, iter := range iters.Iters() {
			tb.add(bounds, iter, true)
		}
	}

	for i := 0; i < b.N; i++ {
		tb.build()
	}
}
