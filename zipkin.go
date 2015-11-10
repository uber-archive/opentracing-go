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

// ZipkinCompatibleTracer is a tracer that represents trace ID as a 4-tuple similar to Zipkin.
type ZipkinCompatibleTracer interface {
	// CreateSpanID instantiates ZipkinSpanID from 4 values. It is not meant for creating brand new IDs
	// externally, but for constructing an ID based on the 4 values read from the incoming request.
	// For example, TChannel protocol explicitly records these 4 values in its frame, so one cannot use
	// a more abstract factory like StringPickler.
	CreateSpanID(traceID, spanID, parentID int64, flags byte) ZipkinSpanID
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
