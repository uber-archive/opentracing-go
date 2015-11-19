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

// Tracer is the entry point API between instrumentation code and the tracing implementation.
type Tracer interface {
	// BeginTrace starts a new trace and creates a new root span.
	// Used by any service that is instrumented for tracing, but did not receive trace ID from upstream.
	// The spanName should reflect the server's name of the end-point that received the request.
	// The domain of names must be limited, do not use UUIDs or entity IDs or timestamps as part of the name.
	// The service endpoint is mandatory.
	BeginTrace(spanName string, service *Endpoint, options *BeginOptions) Span

	// JoinTrace joins a trace started elsewhere and creates a span with the specified ID.
	// Used by services that receive trace ID from upstream.
	// The name should reflect the server's name of the end-point that received the request.
	// The domain of names must be limited, do not use UUIDs or entity IDs or timestamps as part of the name.
	JoinTrace(spanName string, service *Endpoint, spanID SpanID, options *BeginOptions) Span

	// GetStringPickler returns a pickler that can marshal SpanID to/from a string.
	// It can be used when transmitting SpanID across processes in a string protocol, e.g. in HTTP header.
	GetStringPickler() StringPickler

	// Close does a clean shutdown of the tracer, flushing any traces that may be buffered in memory.
	Close()
}

// Endpoint describes the service that participates in the trace.
type Endpoint struct {
	// ServiceName is mandatory name of the service represented by this end point
	ServiceName string

	// IPv4 is 4-byte IP v.4 address of the server represented by this endpoint
	IPv4 int32

	// Port number the server represented by this endpoint is listening to
	Port uint16
}

// StringPickler can marshall a SpanID to and from a string, e.g. for storing in HTTP header.
type StringPickler interface {
	// ToString() serializes a span ID to a string
	ToString(spanID SpanID) string

	// FromString deserialized a span ID from a string, or returns an error if the string value is malformed
	FromString(value string) (SpanID, error)
}
