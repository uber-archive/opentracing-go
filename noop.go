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

type noopTracer struct{}
type noopSpan struct{}
type noopSpanID struct{}
type noopStringPickler struct{}

var (
	defaultTracer       noopTracer
	defaultSpan         noopSpan
	defaultSpanID       noopSpanID
	defaultStringPicker noopStringPickler
)

// NewNoopTracer creates a tracer that does not perform any tracing
func NewNoopTracer() Tracer {
	return &defaultTracer
}

// BeginRootSpan implements BeginRootSpan() of tracing.Tracer
func (t *noopTracer) BeginRootSpan(spanName string, service *Endpoint, options *BeginOptions) Span {
	return &defaultSpan
}

// BeginSpan implements BeginSpan() of tracing.Tracer.
func (t *noopTracer) BeginSpan(spanName string, service *Endpoint, sID SpanID, options *BeginOptions) Span {
	return &defaultSpan
}

func (t *noopTracer) GetStringPickler() StringPickler {
	return &defaultStringPicker
}

func (t *noopTracer) Close() {
	// nothing to do
}

// String implements String() of tracing.SpanID
func (s *noopSpanID) String() string {
	return "tracing-disabled"
}

// SpanID implements SpanID() of tracing.Span
func (s *noopSpan) SpanID() SpanID {
	return &defaultSpanID
}

// BeginChildSpan implements BeginChildSpan() of tracing.Span
func (s *noopSpan) BeginChildSpan(name string, options *BeginOptions) Span {
	return &defaultSpan
}

// EndSpan implements EndSpan() of tracing.Span
func (s *noopSpan) End(options *EndOptions) {
	// noop
}

// AddAttribute implements AddAttribute() of tracing.Span
func (s *noopSpan) AddAttribute(name string, value interface{}) {
	// noop
}

// AddEvent implements AddEvent() of tracing.Span
func (s *noopSpan) AddEvent(name string, options *EventOptions) {
	// noop
}

// ToString implements ToString() of StringPickler
func (p *noopStringPickler) ToString(spanID SpanID) string {
	return ""
}

// FromString implements FromString() of StringPickler
func (p *noopStringPickler) FromString(value string) (SpanID, error) {
	return &defaultSpanID, nil
}
