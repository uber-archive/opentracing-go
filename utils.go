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
	"errors"
	"golang.org/x/net/context"
)

// GetSpanFromHeader creates a top-level RPC server-side span. If the provided header value can be parsed
// as span ID, meaning the client is also instrumented for a compatible tracing system and it booked an RPC
// span, the server joins that span. Otherwise it creates a new root span. If the header contains a value
// that cannot be parsed as span ID, this method returns an error.
func GetSpanFromHeader(header string, tracer Tracer, spanName string, endpoint *Endpoint, options *BeginOptions) (Span, error) {
	if header == "" {
		return tracer.BeginTrace(spanName, endpoint, options), nil
	}
	spanID, err := tracer.GetStringPickler().FromString(header)
	if err != nil {
		return nil, err
	}
	if spanID == nil {
		return tracer.BeginTrace(spanName, endpoint, options), nil
	} else {
		return tracer.JoinTrace(spanName, endpoint, spanID, options), nil
	}
}

const (
	CurrentSpanContextKey = "tracing.current_span"
)

var (
	NoCurrentSpanError  = errors.New("No tracing span found in the context")
	BadCurrentSpanError = errors.New("Tracing span found in the context is of the wrong type")
)

// ContextWithSpan creates a child context that stores the current span
func ContextWithSpan(ctx context.Context, span Span) context.Context {
	return context.WithValue(ctx, CurrentSpanContextKey, span)
}

// GetSpanFromContext retrieves the current span from the context
func GetSpanFromContext(ctx context.Context) (Span, error) {
	if val := ctx.Value(CurrentSpanContextKey); val == nil {
		return nil, NoCurrentSpanError
	} else if span, ok := val.(Span); !ok {
		return nil, BadCurrentSpanError
	} else {
		return span, nil
	}
}
