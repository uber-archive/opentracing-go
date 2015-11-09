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

import (
	"time"
)

// SpanID identifies the span. It contains information that must be passed between services as part of the
// instrumentation, e.g. in the form of HTTP headers or other protocol-specific data.
type SpanID interface {
	String() string
}

// ZipkinSpanID is a subtype of SpanID that can be exposed by Zipkin-compatible tracers.
type ZipkinSpanID interface {
	SpanID

	// TraceID represents globally unique ID of the trace. Usually generated as a random number.
	TraceID() int64

	// ID represents span ID. It must be unique within a given trace, but does not have to be globally unique.
	ID() int64

	// ParentID refers to the ID of the parent span. Should be 0 if the current span is a root span.
	ParentID() int64

	// IsSampled returns whether this trace was chosen for permanent storage by the sampling mechanism of the tracer.
	IsSampled() bool
}

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

	// Client identifies the client that sent the request causing creation of this (entry point) span.
	Client *Endpoint

	// Server can be used to identify the server that will be executing an RPC request represented
	// by the new (child) span.
	Server *Endpoint
}

// EndOptions contains optional flags that can be passed to EndSpan().
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

// Span represents a unit of work executed on behalf of a trace. Examples of spans include a remote procedure call,
// or a in-process method call to a sub-component.  A trace is required to have a single, top level "root" span,
// and zero or more children spans, which in turns can have their own children spans, thus forming a tree structure.
type Span interface {
	// SpanID returns the identifier of the span.
	SpanID() SpanID

	// BeginChildSpan denotes the beginning of a subordinate unit of work with a given name.
	BeginChildSpan(name string, options *BeginOptions) Span

	// End indicates that the work represented by this span has been completed or terminated.
	// If any attributes/events need to be added to the span, it should be done before calling End(),
	// otherwise they may be ignored.
	End(options *EndOptions)

	// AddAttribute attaches a key/value pair so the span. The same key may be repeated multiple times.
	// The storage allows spans and traces to be located both by key and by key=value combinations.
	// The set of supported value types is implementation specific. At minimum the following types
	// should be supported by the tracing system implementation:
	// * string
	// * int32
	// * int64
	// * float64
	// * bool
	// * []byte
	// * Endpoint
	// Other types may be optionally supported, e.g. JSON.
	AddAttribute(name string, value interface{})

	// AddEvent attaches a named marker with a timestamp to the span. The tracer will capture the timestamp.
	AddEvent(name string, options *EventOptions)
}
