package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"go.opentelemetry.io/otel/trace"
)

func TestWithTraceIDHeaderAlwaysSetsTraceAndRequestID(t *testing.T) {
	t.Parallel()

	mw := WithTraceIDHeader(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	rec := httptest.NewRecorder()
	mw.ServeHTTP(rec, req)

	traceID := rec.Header().Get("trace-id")
	if traceID == "" {
		t.Fatalf("trace-id header must not be empty")
	}
	requestID := rec.Header().Get("request-id")
	if requestID == "" {
		t.Fatalf("request-id header must not be empty")
	}
	if rec.Header().Get("X-Trace-Id") != traceID {
		t.Fatalf("X-Trace-Id must match trace-id")
	}
	if rec.Header().Get("X-Request-Id") != requestID {
		t.Fatalf("X-Request-Id must match request-id")
	}
}

func TestWithTraceIDHeaderPropagatesIncomingRequestIDAndNIK(t *testing.T) {
	t.Parallel()

	var gotReqID string
	var gotNIK string
	mw := WithTraceIDHeader(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotReqID = RequestIDFromContext(r.Context())
		gotNIK = NIKFromContext(r.Context())
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	req.Header.Set("request-id", "req-123")
	req.Header.Set("nik", "EMP-7788")
	rec := httptest.NewRecorder()
	mw.ServeHTTP(rec, req)

	if gotReqID != "req-123" {
		t.Fatalf("request-id context mismatch: got=%q", gotReqID)
	}
	if gotNIK != "EMP-7788" {
		t.Fatalf("nik context mismatch: got=%q", gotNIK)
	}
	if rec.Header().Get("request-id") != "req-123" {
		t.Fatalf("request-id response mismatch")
	}
	if rec.Header().Get("nik") != "EMP-7788" {
		t.Fatalf("nik response mismatch")
	}
}

func TestWithTraceIDHeaderUsesSpanTraceIDWhenAvailable(t *testing.T) {
	t.Parallel()

	traceID, err := trace.TraceIDFromHex("1234567890abcdef1234567890abcdef")
	if err != nil {
		t.Fatalf("invalid trace id test fixture")
	}
	spanID, err := trace.SpanIDFromHex("1234567890abcdef")
	if err != nil {
		t.Fatalf("invalid span id test fixture")
	}
	sc := trace.NewSpanContext(trace.SpanContextConfig{
		TraceID:    traceID,
		SpanID:     spanID,
		TraceFlags: trace.FlagsSampled,
		Remote:     true,
	})

	var gotTraceID string
	mw := WithTraceIDHeader(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		gotTraceID = TraceIDFromContext(r.Context())
		w.WriteHeader(http.StatusOK)
	}))

	req := httptest.NewRequest(http.MethodGet, "/healthz", nil)
	req.Header.Set("trace-id", "override-me")
	req = req.WithContext(trace.ContextWithSpanContext(context.Background(), sc))
	rec := httptest.NewRecorder()
	mw.ServeHTTP(rec, req)

	want := "1234567890abcdef1234567890abcdef"
	if gotTraceID != want {
		t.Fatalf("trace-id context mismatch: got=%q want=%q", gotTraceID, want)
	}
	if rec.Header().Get("trace-id") != want {
		t.Fatalf("trace-id response mismatch: got=%q want=%q", rec.Header().Get("trace-id"), want)
	}
}
