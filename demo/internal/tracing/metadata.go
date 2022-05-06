package tracing

import (
	"context"
	"go.opentelemetry.io/otel/propagation"
)

const serviceHeader = "service-name"

var _ propagation.TextMapPropagator = &MetadataPropagator{}

type MetadataPropagator struct{}

// Inject sets metadata key-values from ctx into the carrier.
func (b MetadataPropagator) Inject(ctx context.Context, carrier propagation.TextMapCarrier) {
	appInfo, ok := FromContext(ctx)
	if ok {
		carrier.Set(serviceHeader, appInfo.AppName)
	}
}

// Extract returns a copy of parent with the metadata from the carrier added.
func (b MetadataPropagator) Extract(parent context.Context, carrier propagation.TextMapCarrier) context.Context {
	name := carrier.Get(serviceHeader)
	if name != "" {
		if appInfo, ok := FromContext(parent); !ok {
			appInfo.MetaData[serviceHeader] = name
		} else {
			appInfo = AppInfo{}
			appInfo.MetaData[serviceHeader] = name
			parent = NewContext(parent, appInfo)
		}
	}

	return parent
}

func (b MetadataPropagator) Fields() []string {
	return []string{serviceHeader}
}
