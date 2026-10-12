package migrest

import (
	"context"

	"go.uber.org/zap"

	"rza_migrate_db/internal/requestmeta"
)

func operationLogger(ctx context.Context, operation string, fileName string) *zap.Logger {
	fields := []zap.Field{
		zap.String("component", "migrest"),
		zap.String("operation", operation),
	}
	if fileName != "" {
		fields = append(fields, zap.String("file", fileName))
	}
	if requestID := requestmeta.RequestID(ctx); requestID != "" {
		fields = append(fields, zap.String("request_id", requestID))
	}
	if traceID := requestmeta.TraceID(ctx); traceID != "" {
		fields = append(fields, zap.String("trace_id", traceID))
	}
	if nik := requestmeta.NIK(ctx); nik != "" {
		fields = append(fields, zap.String("nik", nik))
	}
	return zap.L().With(fields...)
}

func appendVersionStateFields(fields []zap.Field, prefix string, state VersionState) []zap.Field {
	if state.Version == nil {
		fields = append(fields, zap.String(prefix+"_version", "nil"))
	} else {
		fields = append(fields, zap.Uint64(prefix+"_version", *state.Version))
	}
	fields = append(fields, zap.Bool(prefix+"_dirty", state.Dirty))
	return fields
}
