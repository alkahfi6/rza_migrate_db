package handler

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"net/http"
	"strings"

	"go.opentelemetry.io/otel/trace"

	"rza_migrate_db/internal/requestmeta"
)

const (
	responseTraceIDHeader   = "trace-id"
	responseRequestIDHeader = "request-id"
	responseNIKHeader       = "nik"
	legacyTraceIDHeader     = "X-Trace-Id"
	legacyRequestIDHeader   = "X-Request-Id"
)

// WithTraceIDHeader enriches response headers with request correlation values.
// Always sets:
// - trace-id (and X-Trace-Id compatibility header)
// - request-id (and X-Request-Id compatibility header)
// Also sets:
// - nik (when provided in request headers)
func WithTraceIDHeader(next http.Handler) http.Handler {
	if next == nil {
		return http.HandlerFunc(func(http.ResponseWriter, *http.Request) {})
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		traceID := resolveTraceID(r)
		requestID := resolveRequestID(r)
		nik := resolveNIK(r)

		w.Header().Set(responseTraceIDHeader, traceID)
		w.Header().Set(responseRequestIDHeader, requestID)
		// Compatibility with older client integrations.
		w.Header().Set(legacyTraceIDHeader, traceID)
		w.Header().Set(legacyRequestIDHeader, requestID)
		if nik != "" {
			w.Header().Set(responseNIKHeader, nik)
		}

		ctx := requestmeta.WithTraceID(r.Context(), traceID)
		ctx = requestmeta.WithRequestID(ctx, requestID)
		if nik != "" {
			ctx = requestmeta.WithNIK(ctx, nik)
		}
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func resolveTraceID(r *http.Request) string {
	if r == nil {
		return newCorrelationID()
	}
	if sc := trace.SpanFromContext(r.Context()).SpanContext(); sc.IsValid() {
		return sc.TraceID().String()
	}
	if incoming := firstHeaderValue(r, responseTraceIDHeader, legacyTraceIDHeader); incoming != "" {
		return incoming
	}
	return newCorrelationID()
}

func resolveRequestID(r *http.Request) string {
	if r == nil {
		return newCorrelationID()
	}
	if incoming := firstHeaderValue(r, responseRequestIDHeader, legacyRequestIDHeader); incoming != "" {
		return incoming
	}
	return newCorrelationID()
}

func resolveNIK(r *http.Request) string {
	if r == nil {
		return ""
	}
	return firstHeaderValue(r, responseNIKHeader, "X-NIK")
}

func firstHeaderValue(r *http.Request, keys ...string) string {
	for _, key := range keys {
		if key == "" {
			continue
		}
		if v := strings.TrimSpace(r.Header.Get(key)); v != "" {
			return sanitizeHeaderValue(v)
		}
	}
	return ""
}

func sanitizeHeaderValue(v string) string {
	trimmed := strings.TrimSpace(v)
	if len(trimmed) > 256 {
		trimmed = trimmed[:256]
	}
	return strings.Map(func(r rune) rune {
		// Drop control characters from propagated headers.
		if r < 32 || r == 127 {
			return -1
		}
		return r
	}, trimmed)
}

func newCorrelationID() string {
	var b [16]byte
	if _, err := rand.Read(b[:]); err != nil {
		// Constant fallback to keep behavior deterministic on entropy failure.
		return "00000000000000000000000000000000"
	}
	return hex.EncodeToString(b[:])
}

// RequestIDFromContext returns request-id propagated by middleware.
func RequestIDFromContext(ctx context.Context) string {
	return requestmeta.RequestID(ctx)
}

// TraceIDFromContext returns trace-id propagated by middleware.
func TraceIDFromContext(ctx context.Context) string {
	return requestmeta.TraceID(ctx)
}

// NIKFromContext returns nik propagated by middleware.
func NIKFromContext(ctx context.Context) string {
	return requestmeta.NIK(ctx)
}
