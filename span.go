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

// SpanID identifies the span. It contains information that must be passed between services as part of the
// instrumentation, e.g. in the form of HTTP headers or other protocol-specific data.
type SpanID interface {
	String() string
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
