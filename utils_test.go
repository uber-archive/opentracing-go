// Copyright (c) 2015 Uber Technologies, Inc.
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
package tracing_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/uber-common/opentracing-go"
	"golang.org/x/net/context"
	"testing"
)

var endpoint = &tracing.Endpoint{ServiceName: "test-service"}

func TestGetSpanFromHeader(t *testing.T) {
	tracer := tracing.NewNoopTracer()

	// empty header causes a start of new trace
	span, err := tracing.GetSpanFromHeader("", tracer, "test-span", endpoint, nil)
	assert.NoError(t, err)
	assert.NotNil(t, span)

	// should recognize span ID and join the trace
	span, err = tracing.GetSpanFromHeader("x", tracer, "test-span", endpoint, nil)
	assert.NoError(t, err)
	assert.NotNil(t, span)

	// should not recognize span ID and start a new trace
	span, err = tracing.GetSpanFromHeader("y", tracer, "test-span", endpoint, nil)
	assert.NoError(t, err)
	assert.NotNil(t, span)

	// simulates header parsing error
	span, err = tracing.GetSpanFromHeader("error", tracer, "test-span", endpoint, nil)
	assert.Error(t, err)
	assert.Nil(t, span)
}

func TestContextOps(t *testing.T) {
	tracer := tracing.NewNoopTracer()
	span := tracer.BeginTrace("test-span", endpoint, nil)
	bgCtx := context.Background()

	_, err := tracing.GetSpanFromContext(bgCtx)
	assert.Equal(t, tracing.NoCurrentSpanError, err)

	badCtx := context.WithValue(bgCtx, tracing.CurrentSpanContextKey, tracer)
	_, err = tracing.GetSpanFromContext(badCtx)
	assert.Equal(t, tracing.BadCurrentSpanError, err)

	ctx := tracing.ContextWithSpan(bgCtx, span)
	span2, err := tracing.GetSpanFromContext(ctx)
	assert.NoError(t, err)
	assert.Equal(t, span, span2)
}
