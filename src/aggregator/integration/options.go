// Copyright (c) 2016 Uber Technologies, Inc.
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

package integration

import "time"

const (
	defaultServerStateChangeTimeout = 5 * time.Second
	defaultClientBatchSize          = 1440
	defaultClientConnectTimeout     = time.Second
	defaultWorkerPoolSize           = 4
	defaultInstanceID               = "localhost:6000"
	defaultNumShards                = 1024
	defaultPlacementKVKey           = "placement"
)

type testOptions interface {
	// SetMsgpackAddr sets the msgpack server address.
	SetMsgpackAddr(value string) testOptions

	// MsgpackAddr returns the msgpack server address.
	MsgpackAddr() string

	// SetHTTPAddr sets the http server address.
	SetHTTPAddr(value string) testOptions

	// HTTPAddr returns the http server address.
	HTTPAddr() string

	// SetInstanceID sets the instance id.
	SetInstanceID(value string) testOptions

	// InstanceID returns the instance id.
	InstanceID() string

	// SetNumShards sets the number of shards.
	SetNumShards(value int) testOptions

	// NumShards returns the number of shards.
	NumShards() int

	// SetPlacementKVKey sets the placement kv key.
	SetPlacementKVKey(value string) testOptions

	// PlacementKVKey returns the placement kv key.
	PlacementKVKey() string

	// SetClientBatchSize sets the client-side batch size.
	SetClientBatchSize(value int) testOptions

	// ClientBatchSize returns the client-side batch size.
	ClientBatchSize() int

	// SetClientConnectTimeout sets the client-side connect timeout.
	SetClientConnectTimeout(value time.Duration) testOptions

	// ClientConnectTimeout returns the client-side connect timeout.
	ClientConnectTimeout() time.Duration

	// SetServerStateChangeTimeout sets the client connect timeout.
	SetServerStateChangeTimeout(value time.Duration) testOptions

	// ServerStateChangeTimeout returns the client connect timeout.
	ServerStateChangeTimeout() time.Duration

	// SetWorkerPoolSize sets the number of workers in the worker pool.
	SetWorkerPoolSize(value int) testOptions

	// WorkerPoolSize returns the number of workers in the worker pool.
	WorkerPoolSize() int
}

type options struct {
	msgpackAddr              string
	httpAddr                 string
	instanceID               string
	numShards                int
	placementKVKey           string
	serverStateChangeTimeout time.Duration
	workerPoolSize           int
	clientBatchSize          int
	clientConnectTimeout     time.Duration
}

func newTestOptions() testOptions {
	return &options{
		instanceID:               defaultInstanceID,
		numShards:                defaultNumShards,
		placementKVKey:           defaultPlacementKVKey,
		serverStateChangeTimeout: defaultServerStateChangeTimeout,
		workerPoolSize:           defaultWorkerPoolSize,
		clientBatchSize:          defaultClientBatchSize,
		clientConnectTimeout:     defaultClientConnectTimeout,
	}
}

func (o *options) SetMsgpackAddr(value string) testOptions {
	opts := *o
	opts.msgpackAddr = value
	return &opts
}

func (o *options) MsgpackAddr() string {
	return o.msgpackAddr
}

func (o *options) SetHTTPAddr(value string) testOptions {
	opts := *o
	opts.httpAddr = value
	return &opts
}

func (o *options) HTTPAddr() string {
	return o.httpAddr
}

func (o *options) SetInstanceID(value string) testOptions {
	opts := *o
	opts.instanceID = value
	return &opts
}

func (o *options) InstanceID() string {
	return o.instanceID
}

func (o *options) SetNumShards(value int) testOptions {
	opts := *o
	opts.numShards = value
	return &opts
}

func (o *options) NumShards() int {
	return o.numShards
}

func (o *options) SetPlacementKVKey(value string) testOptions {
	opts := *o
	opts.placementKVKey = value
	return &opts
}

func (o *options) PlacementKVKey() string {
	return o.placementKVKey
}

func (o *options) SetClientBatchSize(value int) testOptions {
	opts := *o
	opts.clientBatchSize = value
	return &opts
}

func (o *options) ClientBatchSize() int {
	return o.clientBatchSize
}

func (o *options) SetClientConnectTimeout(value time.Duration) testOptions {
	opts := *o
	opts.clientConnectTimeout = value
	return &opts
}

func (o *options) ClientConnectTimeout() time.Duration {
	return o.clientConnectTimeout
}

func (o *options) SetServerStateChangeTimeout(value time.Duration) testOptions {
	opts := *o
	opts.serverStateChangeTimeout = value
	return &opts
}

func (o *options) ServerStateChangeTimeout() time.Duration {
	return o.serverStateChangeTimeout
}

func (o *options) SetWorkerPoolSize(value int) testOptions {
	opts := *o
	opts.workerPoolSize = value
	return &opts
}

func (o *options) WorkerPoolSize() int {
	return o.workerPoolSize
}