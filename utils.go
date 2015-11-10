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

// GetSpanFromHeader creates a top-level RPC server-side span. If the provided header value can be parsed
// as span ID, meaning the client is also instrumented for a compatible tracing system and it booked an RPC
// span, the server joins that span. Otherwise it creates a new root span. If the header contains a value
// that cannot be parsed as span ID, this method returns an error.
func GetSpanFromHeader(header string, tracer Tracer, spanName string, endpoint *Endpoint, options *BeginOptions) (Span, error) {
	if header == "" {
		return tracer.BeginRootSpan(spanName, endpoint, options), nil
	}
	spanID, err := tracer.GetStringPickler().FromString(header)
	if err != nil {
		return nil, err
	}
	if spanID == nil {
		return tracer.BeginRootSpan(spanName, endpoint, options), nil
	} else {
		return tracer.BeginSpan(spanName, endpoint, spanID, options), nil
	}
}
