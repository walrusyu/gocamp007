package tracing

import (
	"context"
	"github.com/walrusyu/gocamp007/demo/internal/middleware"
	"github.com/walrusyu/gocamp007/demo/internal/tracing"
	"github.com/walrusyu/gocamp007/demo/internal/transport"
	"go.opentelemetry.io/otel/trace"
)

func TracingMiddlewareOnServer(handlers ...tracing.OptionHandler) middleware.Middleware {
	tracer := tracing.NewTracer(trace.SpanKindServer, handlers...)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				var span trace.Span
				ctx, span = tracer.Start(ctx, tr.Operation(), tr.RequestHeader())
				tracing.SetServerSpan(ctx, span, req)
				defer func() { tracer.End(ctx, span, reply, err) }()
			}
			return handler(ctx, req)
		}
	}
}

func TracingMiddlewareOnClient(handlers ...tracing.OptionHandler) middleware.Middleware {
	tracer := tracing.NewTracer(trace.SpanKindClient, handlers...)
	return func(handler middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req interface{}) (reply interface{}, err error) {
			if tr, ok := transport.FromClientContext(ctx); ok {
				var span trace.Span
				ctx, span = tracer.Start(ctx, tr.Operation(), tr.RequestHeader())
				tracing.SetClientSpan(ctx, span, req)
				defer func() { tracer.End(ctx, span, reply, err) }()
			}
			return handler(ctx, req)
		}
	}
}
