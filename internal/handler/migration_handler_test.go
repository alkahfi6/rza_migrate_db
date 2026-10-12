package handler

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"rza_migrate_db/internal/migrest"
)

func TestUpRequiresAPIKey(t *testing.T) {
	t.Parallel()

	h := NewMigrationHandler(stubService{}, migrest.SecurityConfig{
		EnableAPIKey: true,
		APIKey:       "secret",
	}, "test-service")
	router := h.Router()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/migrations/up", bytes.NewBufferString(`{"steps":1}`))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusUnauthorized {
		t.Fatalf("status mismatch: got=%d want=%d", rec.Code, http.StatusUnauthorized)
	}
}

func TestRunFileRequiresFilePayload(t *testing.T) {
	t.Parallel()

	h := NewMigrationHandler(stubService{}, migrest.SecurityConfig{
		EnableAPIKey: false,
	}, "test-service")
	router := h.Router()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/migrations/run-file", bytes.NewBufferString(`{}`))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status mismatch: got=%d want=%d", rec.Code, http.StatusBadRequest)
	}
}

func TestAcknowledgeManualRequiresPayload(t *testing.T) {
	t.Parallel()

	h := NewMigrationHandler(stubService{}, migrest.SecurityConfig{
		EnableAPIKey: false,
	}, "test-service")
	router := h.Router()

	req := httptest.NewRequest(http.MethodPost, "/api/v1/migrations/ack-manual", bytes.NewBufferString(`{}`))
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusBadRequest {
		t.Fatalf("status mismatch: got=%d want=%d", rec.Code, http.StatusBadRequest)
	}
}

func TestVersionSuccess(t *testing.T) {
	t.Parallel()

	v := uint64(20260217073015000)
	h := NewMigrationHandler(stubService{
		version: migrest.VersionState{
			Version: &v,
			Dirty:   false,
		},
	}, migrest.SecurityConfig{
		EnableAPIKey: false,
	}, "test-service")
	router := h.Router()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/migrations/version", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status mismatch: got=%d want=%d", rec.Code, http.StatusOK)
	}
	if body := rec.Body.String(); !bytes.Contains([]byte(body), []byte(`"version":"20260217073015000"`)) {
		t.Fatalf("expected response to contain version string, got: %s", body)
	}
}

func TestPendingSuccess(t *testing.T) {
	t.Parallel()

	v := uint64(20260217073015000)
	h := NewMigrationHandler(stubService{
		version: migrest.VersionState{
			Version: &v,
			Dirty:   false,
		},
	}, migrest.SecurityConfig{
		EnableAPIKey: false,
	}, "test-service")
	router := h.Router()

	req := httptest.NewRequest(http.MethodGet, "/api/v1/migrations/pending", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status mismatch: got=%d want=%d", rec.Code, http.StatusOK)
	}
	body := rec.Body.String()
	if !bytes.Contains([]byte(body), []byte(`"total_pending":1`)) {
		t.Fatalf("expected total_pending=1, got: %s", body)
	}
	if !bytes.Contains([]byte(body), []byte(`"file_name":"20260217073015001_indexes_dbo.TableA_v1.up.sql"`)) {
		t.Fatalf("expected pending file in response, got: %s", body)
	}
}

func TestDownRequiresBodyAndPositiveSteps(t *testing.T) {
	t.Parallel()

	h := NewMigrationHandler(stubService{}, migrest.SecurityConfig{
		EnableAPIKey: false,
	}, "test-service")
	router := h.Router()

	reqNoBody := httptest.NewRequest(http.MethodPost, "/api/v1/migrations/down", nil)
	recNoBody := httptest.NewRecorder()
	router.ServeHTTP(recNoBody, reqNoBody)
	if recNoBody.Code != http.StatusBadRequest {
		t.Fatalf("status mismatch for empty body: got=%d want=%d", recNoBody.Code, http.StatusBadRequest)
	}

	reqZero := httptest.NewRequest(http.MethodPost, "/api/v1/migrations/down", bytes.NewBufferString(`{"steps":0}`))
	recZero := httptest.NewRecorder()
	router.ServeHTTP(recZero, reqZero)
	if recZero.Code != http.StatusBadRequest {
		t.Fatalf("status mismatch for steps=0: got=%d want=%d", recZero.Code, http.StatusBadRequest)
	}

	reqNeg := httptest.NewRequest(http.MethodPost, "/api/v1/migrations/down", bytes.NewBufferString(`{"steps":-1}`))
	recNeg := httptest.NewRecorder()
	router.ServeHTTP(recNeg, reqNeg)
	if recNeg.Code != http.StatusBadRequest {
		t.Fatalf("status mismatch for steps<0: got=%d want=%d", recNeg.Code, http.StatusBadRequest)
	}
}

func TestOpenAPIRoute(t *testing.T) {
	t.Parallel()

	h := NewMigrationHandler(stubService{}, migrest.SecurityConfig{EnableAPIKey: false}, "test-service")
	router := h.Router()

	req := httptest.NewRequest(http.MethodGet, "/swagger.yaml", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status mismatch: got=%d want=%d", rec.Code, http.StatusOK)
	}
	if body := rec.Body.String(); !bytes.Contains([]byte(body), []byte(`swagger: "2.0"`)) {
		t.Fatalf("expected openapi spec body, got: %s", body)
	}
}

func TestOpenAPILegacyAliasRoute(t *testing.T) {
	t.Parallel()

	h := NewMigrationHandler(stubService{}, migrest.SecurityConfig{EnableAPIKey: false}, "test-service")
	router := h.Router()

	req := httptest.NewRequest(http.MethodGet, "/openapi.yaml", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status mismatch: got=%d want=%d", rec.Code, http.StatusOK)
	}
}

func TestSwaggerRoute(t *testing.T) {
	t.Parallel()

	h := NewMigrationHandler(stubService{}, migrest.SecurityConfig{EnableAPIKey: false}, "test-service")
	router := h.Router()

	req := httptest.NewRequest(http.MethodGet, "/swagger", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Fatalf("status mismatch: got=%d want=%d", rec.Code, http.StatusOK)
	}
	if body := rec.Body.String(); !bytes.Contains([]byte(body), []byte("SwaggerUIBundle")) {
		t.Fatalf("expected swagger ui page, got: %s", body)
	}
}

type stubService struct {
	version migrest.VersionState
}

func (s stubService) CurrentVersion(context.Context) (migrest.VersionState, error) {
	return s.version, nil
}

func (s stubService) Pending(context.Context) (migrest.PendingState, error) {
	return migrest.PendingState{
		Current: s.version,
		Files: []migrest.PendingMigration{
			{
				FileName: "20260217073015001_indexes_dbo.TableA_v1.up.sql",
				Version:  20260217073015001,
				Category: "indexes",
				Object:   "dbo.TableA",
				VK:       1,
				IsCore:   false,
			},
		},
	}, nil
}

func (s stubService) Up(context.Context, int) (migrest.OperationResult, error) {
	return migrest.OperationResult{
		Operation: "up",
		Before:    s.version,
		After:     s.version,
	}, nil
}

func (s stubService) Down(context.Context, int) (migrest.OperationResult, error) {
	return migrest.OperationResult{
		Operation: "down",
		Before:    s.version,
		After:     s.version,
	}, nil
}

func (s stubService) RunSpecificFile(context.Context, string) (migrest.OperationResult, error) {
	return migrest.OperationResult{
		Operation: "run_file",
		Before:    s.version,
		After:     s.version,
	}, nil
}

func (s stubService) AcknowledgeManualFile(context.Context, string, string) (migrest.OperationResult, error) {
	return migrest.OperationResult{
		Operation: "ack_manual",
		Before:    s.version,
		After:     s.version,
	}, nil
}
