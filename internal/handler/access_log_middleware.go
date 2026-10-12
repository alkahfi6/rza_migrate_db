package handler

import (
	"net/http"
	"strings"

	"github.com/felixge/httpsnoop"
	"go.uber.org/zap"

	"rza_migrate_db/internal/requestmeta"
)

// WithAccessLog emits one structured access log event per request.
func WithAccessLog(next http.Handler) http.Handler {
	if next == nil {
		return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		metrics := httpsnoop.CaptureMetrics(next, w, r)
		requestID, traceID, nik := requestCorrelationFromContextAndHeaders(r, w)

		fields := []zap.Field{
			zap.String("component", "http_access"),
			zap.String("method", strings.ToUpper(strings.TrimSpace(r.Method))),
			zap.String("path", r.URL.Path),
			zap.Int("status", metrics.Code),
			zap.Int64("response_bytes", metrics.Written),
			zap.Duration("duration", metrics.Duration),
		}
		if requestID != "" {
			fields = append(fields, zap.String("request_id", requestID))
		}
		if traceID != "" {
			fields = append(fields, zap.String("trace_id", traceID))
		}
		if nik != "" {
			fields = append(fields, zap.String("nik", nik))
		}

		logger := zap.L()
		switch {
		case metrics.Code >= 500:
			logger.Error("http_request", fields...)
		case metrics.Code >= 400:
			logger.Warn("http_request", fields...)
		default:
			logger.Info("http_request", fields...)
		}
	})
}

func requestCorrelationFromContextAndHeaders(r *http.Request, w http.ResponseWriter) (requestID string, traceID string, nik string) {
	requestID = strings.TrimSpace(requestmeta.RequestID(r.Context()))
	traceID = strings.TrimSpace(requestmeta.TraceID(r.Context()))
	nik = strings.TrimSpace(requestmeta.NIK(r.Context()))

	if requestID == "" {
		requestID = strings.TrimSpace(w.Header().Get(responseRequestIDHeader))
	}
	if traceID == "" {
		traceID = strings.TrimSpace(w.Header().Get(responseTraceIDHeader))
	}
	if nik == "" {
		nik = strings.TrimSpace(w.Header().Get(responseNIKHeader))
	}
	return requestID, traceID, nik
}
