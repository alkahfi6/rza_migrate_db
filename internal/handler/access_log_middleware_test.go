package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"go.uber.org/zap"
	"go.uber.org/zap/zaptest/observer"

	"rza_migrate_db/internal/requestmeta"
)

func TestWithAccessLogIncludesCorrelationFields(t *testing.T) {
	t.Parallel()

	core, observed := observer.New(zap.InfoLevel)
	logger := zap.New(core)
	restore := zap.ReplaceGlobals(logger)
	defer restore()

	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := requestmeta.WithRequestID(r.Context(), "req-01")
		ctx = requestmeta.WithTraceID(ctx, "trace-01")
		ctx = requestmeta.WithNIK(ctx, "nik-01")
		r = r.WithContext(ctx)
		w.Header().Set("request-id", "req-01")
		w.Header().Set("trace-id", "trace-01")
		w.Header().Set("nik", "nik-01")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(`{"ok":true}`))
	})
	handler := WithAccessLog(next)

	req := httptest.NewRequest(http.MethodPost, "/api/v1/migrations/up", nil)
	rec := httptest.NewRecorder()
	handler.ServeHTTP(rec, req)

	entries := observed.All()
	if len(entries) != 1 {
		t.Fatalf("log entries mismatch: got=%d want=1", len(entries))
	}

	entry := entries[0]
	if entry.Message != "http_request" {
		t.Fatalf("message mismatch: got=%s", entry.Message)
	}
	ctx := entry.ContextMap()
	if ctx["method"] != "POST" {
		t.Fatalf("method mismatch: got=%v", ctx["method"])
	}
	if ctx["path"] != "/api/v1/migrations/up" {
		t.Fatalf("path mismatch: got=%v", ctx["path"])
	}
	if status, ok := ctx["status"].(int64); !ok || status != 201 {
		t.Fatalf("status mismatch: got=%v", ctx["status"])
	}
	if ctx["request_id"] != "req-01" {
		t.Fatalf("request_id mismatch: got=%v", ctx["request_id"])
	}
	if ctx["trace_id"] != "trace-01" {
		t.Fatalf("trace_id mismatch: got=%v", ctx["trace_id"])
	}
	if ctx["nik"] != "nik-01" {
		t.Fatalf("nik mismatch: got=%v", ctx["nik"])
	}
}
