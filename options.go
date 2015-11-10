// Copyright (c) 2015 Uber Technologies, Inc.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package tracing

import "time"

// TimeOption can be used to provide externally captured time and duration to span methods, e.g. when recording
// spans produced by a component that could not emit them directly to the tracer, such as a mobile application.
type TimeOption struct {
	// Timestamp of the event being recorded, such as start time of the span recorded externally.
	// If nil, the tracer will capture the current time.  Microseconds precision.
	Timestamp *time.Time
}

// BeginOptions contains optional flags that can be passed to BeginChildSpan().
type BeginOptions struct {
	TimeOption

	// LocalComponent, marks the span as a local, in-process unit of work, such as a function call to a library.
	// When this field is empty string, the span is considered to be an RPC span.
	LocalComponent string

	// Async marks the span as async, non-blocking, indicating that the parent continues doing other work.
	// This can be used in calculation of a critical path through the trace.
	// By default spans are considered sync/blocking.
	Async bool

	// Peer identifies the peer endpoint of an RPC request. When a server creates a span to handle incoming
	// request, the Peer is the client that made the request, e.g. from http.Request.RemoteAddr. When a process
	// creates a child span in order to make an RPC request to another server, i.e. it acts as a client,
	// the Peer is the server it is about to call.
	Peer *Endpoint
}

// EndOptions contains optional flags that can be passed to span.End() method.
type EndOptions struct {
	// Duration of the span calculated externally. If not specified, the tracer will calculate it as endTs - startTs.
	// Microseconds precision.
	Duration *time.Duration

	// Error indicates that span execution finished with an error. This can be used by the tracers to treat the span
	// as an anomaly, rather than ignoring it, e.g. if it finished quickly.
	Error error
}

// EventOptions contains optional flags that can be passed to AppendEvent().
type EventOptions struct {
	TimeOption
}

