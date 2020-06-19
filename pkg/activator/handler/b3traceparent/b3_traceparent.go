/*
Copyright 2018 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package b3traceparent

import (
	"go.opencensus.io/plugin/ochttp/propagation/b3"
	"go.opencensus.io/plugin/ochttp/propagation/tracecontext"
	"go.opencensus.io/trace"
	"go.opencensus.io/trace/propagation"
	"net/http"
)

// The propagation.HTTPFormats that this is built on.
var (
	b3format    = &b3.HTTPFormat{}
	traceparent = &tracecontext.HTTPFormat{}
)

// HTTPFormat is a propagation.HTTPFormat that reads both b3 and traceparent tracing headers,
// preferring traceparent. It will write both formats.
type HTTPFormat struct{}

var _ propagation.HTTPFormat = (*HTTPFormat)(nil)

func (H *HTTPFormat) SpanContextFromRequest(req *http.Request) (trace.SpanContext, bool) {
	if sc, ok := traceparent.SpanContextFromRequest(req); ok {
		return sc, true
	}
	if sc, ok := b3format.SpanContextFromRequest(req); ok {
		return sc, true
	}
	return trace.SpanContext{}, false
}

func (H *HTTPFormat) SpanContextToRequest(sc trace.SpanContext, req *http.Request) {
	traceparent.SpanContextToRequest(sc, req)
	b3format.SpanContextToRequest(sc, req)
}
