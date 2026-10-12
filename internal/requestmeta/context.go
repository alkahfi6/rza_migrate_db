package requestmeta

import "context"

type contextKey string

const (
	requestIDKey contextKey = "request_id"
	traceIDKey   contextKey = "trace_id"
	nikKey       contextKey = "nik"
)

// WithRequestID stores request-id in context.
func WithRequestID(ctx context.Context, value string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, requestIDKey, value)
}

// WithTraceID stores trace-id in context.
func WithTraceID(ctx context.Context, value string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, traceIDKey, value)
}

// WithNIK stores nik in context.
func WithNIK(ctx context.Context, value string) context.Context {
	if ctx == nil {
		ctx = context.Background()
	}
	return context.WithValue(ctx, nikKey, value)
}

// RequestID returns request-id from context.
func RequestID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	v, _ := ctx.Value(requestIDKey).(string)
	return v
}

// TraceID returns trace-id from context.
func TraceID(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	v, _ := ctx.Value(traceIDKey).(string)
	return v
}

// NIK returns nik from context.
func NIK(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	v, _ := ctx.Value(nikKey).(string)
	return v
}
