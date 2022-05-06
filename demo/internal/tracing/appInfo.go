package tracing

import "context"

type AppInfo struct {
	AppName  string
	MetaData map[string]string
}

type appKey struct{}

func NewContext(ctx context.Context, s AppInfo) context.Context {
	return context.WithValue(ctx, appKey{}, s)
}

func FromContext(ctx context.Context) (s AppInfo, ok bool) {
	s, ok = ctx.Value(appKey{}).(AppInfo)
	return
}
