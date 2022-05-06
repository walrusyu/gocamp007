package tracing

import (
	"context"
	"fmt"
	"github.com/walrusyu/gocamp007/demo/errors"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/protobuf/proto"
)

type tracer struct {
	tracer trace.Tracer
	kind   trace.SpanKind
	opt    *option
}

func NewTracer(kind trace.SpanKind, opts ...OptionHandler) *tracer {
	op := option{
		propagator: propagation.NewCompositeTextMapPropagator(MetadataPropagator{}, propagation.Baggage{}, propagation.TraceContext{}),
	}
	for _, o := range opts {
		o(&op)
	}
	if op.tracerProvider != nil {
		otel.SetTracerProvider(op.tracerProvider)
	}
	switch kind {
	case trace.SpanKindClient:
		return &tracer{tracer: otel.Tracer("github.com/walrusyu/gocamp007/demo"), kind: kind, opt: &op}
	case trace.SpanKindServer:
		return &tracer{tracer: otel.Tracer("github.com/walrusyu/gocamp007/demo"), kind: kind, opt: &op}
	default:
		panic(fmt.Sprintf("unsupported span kind: %v", kind))
	}
}

func (t *tracer) Start(ctx context.Context, operation string, carrier propagation.TextMapCarrier) (context.Context, trace.Span) {
	if t.kind == trace.SpanKindServer {
		ctx = t.opt.propagator.Extract(ctx, carrier)
	}
	ctx, span := t.tracer.Start(ctx,
		operation,
		trace.WithSpanKind(t.kind),
	)
	if t.kind == trace.SpanKindClient {
		t.opt.propagator.Inject(ctx, carrier)
	}
	return ctx, span
}

func (t *tracer) End(ctx context.Context, span trace.Span, m interface{}, err error) {
	if err != nil {
		span.RecordError(err)
		if e := errors.FromError(err); e != nil {
			span.SetAttributes(attribute.Key("rpc.status_code").Int64(int64(e.Code)))
		}
		span.SetStatus(codes.Error, err.Error())
	} else {
		span.SetStatus(codes.Ok, "OK")
	}

	if p, ok := m.(proto.Message); ok {
		if t.kind == trace.SpanKindServer {
			span.SetAttributes(attribute.Key("send_msg.size").Int(proto.Size(p)))
		} else {
			span.SetAttributes(attribute.Key("recv_msg.size").Int(proto.Size(p)))
		}
	}
	span.End()
}
