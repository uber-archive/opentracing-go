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

package tracing_test

import (
	"code.uber.internal/infra/jaeger-client-go"
	"github.com/stretchr/testify/suite"
	"testing"
)

type tracerSuite struct {
	suite.Suite
	tracer tracing.Tracer
}

func TestNoopTracer(t *testing.T) {
	suite.Run(t, new(tracerSuite))
}

func (s *tracerSuite) SetupTest() {
	s.tracer = tracing.NewNoopTracer()
	s.NotNil(s.tracer)
}

func (s *tracerSuite) TearDownSuite() {
	s.tracer = nil
}

func (s *tracerSuite) TestRootSpan() {
	span := s.tracer.BeginRootSpan("test", nil, nil)
	s.NotNil(span.SpanID())

	span.AddAttribute("key", "value")
	span.AddEvent("event", nil)
	span.End(nil)
}

func (s *tracerSuite) TestServerSpan() {
	pickler := s.tracer.GetStringPickler()
	spanID, err := pickler.FromString("")
	s.NoError(err)
	s.Equal("tracing-disabled", spanID.String())
	s.Equal("", pickler.ToString(spanID))

	span := s.tracer.BeginSpan("test", nil, spanID, nil)
	s.Equal(spanID, span.SpanID())
	span.End(nil)
}

func (s *tracerSuite) TestClientSpan() {
	span := s.tracer.BeginRootSpan("test", nil, nil)
	s.NotNil(span.SpanID())

	child := span.BeginChildSpan("child", nil)
	s.Equal(span.SpanID(), child.SpanID())
	child.End(nil)

	span.End(nil)
}
