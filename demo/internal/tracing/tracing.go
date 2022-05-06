package tracing

import (
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

type OptionHandler func(*option)

type option struct {
	tracerProvider trace.TracerProvider
	propagator     propagation.TextMapPropagator
}

// WithPropagator with tracer propagator.
func WithPropagator(propagator propagation.TextMapPropagator) OptionHandler {
	return func(opts *option) {
		opts.propagator = propagator
	}
}

// WithTracerProvider with tracer provider.
// Deprecated: use otel.SetTracerProvider(provider) instead.
func WithTracerProvider(provider trace.TracerProvider) OptionHandler {
	return func(opts *option) {
		opts.tracerProvider = provider
	}
}

//
//// TraceID returns a traceid valuer.
//func TraceID() log.Valuer {
//	return func(ctx context.Context) interface{} {
//		if span := trace.SpanContextFromContext(ctx); span.HasTraceID() {
//			return span.TraceID().String()
//		}
//		return ""
//	}
//}
//
//// SpanID returns a spanid valuer.
//func SpanID() log.Valuer {
//	return func(ctx context.Context) interface{} {
//		if span := trace.SpanContextFromContext(ctx); span.HasSpanID() {
//			return span.SpanID().String()
//		}
//		return ""
//	}
//}
