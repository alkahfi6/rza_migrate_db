package handler

import (
	"context"
	"crypto/subtle"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"

	"rza_migrate_db/internal/migrest"
)

const maxBodyBytes = 1 << 20 // 1MB

// MigrationService is minimal service contract used by HTTP handler.
type MigrationService interface {
	CurrentVersion(ctx context.Context) (migrest.VersionState, error)
	Pending(ctx context.Context) (migrest.PendingState, error)
	Up(ctx context.Context, steps int) (migrest.OperationResult, error)
	Down(ctx context.Context, steps int) (migrest.OperationResult, error)
	RunSpecificFile(ctx context.Context, fileName string) (migrest.OperationResult, error)
	AcknowledgeManualFile(ctx context.Context, fileName string, reason string) (migrest.OperationResult, error)
}

// MigrationHandler exposes migration operations over REST.
type MigrationHandler struct {
	service      MigrationService
	serviceName  string
	enableAPIKey bool
	apiKey       string
	now          func() time.Time
}

// NewMigrationHandler creates migration REST handler.
func NewMigrationHandler(service MigrationService, security migrest.SecurityConfig, serviceName string) *MigrationHandler {
	serviceName = strings.TrimSpace(serviceName)
	if serviceName == "" {
		serviceName = "migration_rest_service"
	}
	return &MigrationHandler{
		service:      service,
		serviceName:  serviceName,
		enableAPIKey: security.EnableAPIKey,
		apiKey:       strings.TrimSpace(security.APIKey),
		now:          time.Now,
	}
}

// Router builds http.Handler with all routes.
func (h *MigrationHandler) Router() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/healthz", h.handleHealthz)
	mux.HandleFunc("/swagger.yaml", h.handleSwaggerSpec)
	// legacy alias for backward compatibility
	mux.HandleFunc("/openapi.yaml", h.handleSwaggerSpec)
	mux.HandleFunc("/swagger", h.handleSwaggerUI)
	mux.HandleFunc("/api/v1/migrations/version", h.withAuth(h.handleVersion))
	mux.HandleFunc("/api/v1/migrations/pending", h.withAuth(h.handlePending))
	mux.HandleFunc("/api/v1/migrations/up", h.withAuth(h.handleUp))
	mux.HandleFunc("/api/v1/migrations/down", h.withAuth(h.handleDown))
	mux.HandleFunc("/api/v1/migrations/run-file", h.withAuth(h.handleRunFile))
	mux.HandleFunc("/api/v1/migrations/ack-manual", h.withAuth(h.handleAcknowledgeManual))
	return mux
}

type apiResponse struct {
	Status    string    `json:"status"`
	Timestamp string    `json:"timestamp"`
	Data      any       `json:"data,omitempty"`
	Error     *apiError `json:"error,omitempty"`
}

type apiError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type versionPayload struct {
	Version *string `json:"version"`
	Dirty   bool    `json:"dirty"`
}

type operationPayload struct {
	Operation      string                         `json:"operation"`
	OperationID    string                         `json:"operation_id"`
	FileName       string                         `json:"file_name,omitempty"`
	RequestedSteps int                            `json:"requested_steps"`
	RecordedSteps  int                            `json:"recorded_steps"`
	NoChange       bool                           `json:"no_change"`
	Before         versionPayload                 `json:"before"`
	After          versionPayload                 `json:"after"`
	AppliedFiles   []migrest.AppliedMigrationFile `json:"applied_files"`
	ManifestSHA256 string                         `json:"manifest_sha256,omitempty"`
}

type pendingItemPayload struct {
	FileName string `json:"file_name"`
	Version  string `json:"version"`
	Category string `json:"category"`
	Object   string `json:"object,omitempty"`
	VK       int    `json:"vk"`
	IsCore   bool   `json:"is_core"`
}

type pendingPayload struct {
	Current      versionPayload       `json:"current"`
	TotalPending int                  `json:"total_pending"`
	Files        []pendingItemPayload `json:"files"`
}

type stepsRequest struct {
	Steps int `json:"steps"`
}

type runFileRequest struct {
	File string `json:"file"`
}

type acknowledgeManualRequest struct {
	File   string `json:"file"`
	Reason string `json:"reason"`
}

func (h *MigrationHandler) handleHealthz(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeMethodNotAllowed(w)
		return
	}
	h.writeJSON(w, http.StatusOK, apiResponse{
		Status:    "ok",
		Timestamp: h.now().UTC().Format(time.RFC3339),
		Data: map[string]string{
			"service": h.serviceName,
		},
	})
}

func (h *MigrationHandler) handleVersion(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeMethodNotAllowed(w)
		return
	}
	state, err := h.service.CurrentVersion(r.Context())
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	h.writeJSON(w, http.StatusOK, apiResponse{
		Status:    "ok",
		Timestamp: h.now().UTC().Format(time.RFC3339),
		Data: versionPayload{
			Version: toVersionStringPtr(state.Version),
			Dirty:   state.Dirty,
		},
	})
}

func (h *MigrationHandler) handlePending(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		h.writeMethodNotAllowed(w)
		return
	}
	state, err := h.service.Pending(r.Context())
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	files := make([]pendingItemPayload, 0, len(state.Files))
	for _, item := range state.Files {
		files = append(files, pendingItemPayload{
			FileName: item.FileName,
			Version:  strconv.FormatUint(item.Version, 10),
			Category: item.Category,
			Object:   item.Object,
			VK:       item.VK,
			IsCore:   item.IsCore,
		})
	}

	h.writeJSON(w, http.StatusOK, apiResponse{
		Status:    "ok",
		Timestamp: h.now().UTC().Format(time.RFC3339),
		Data: pendingPayload{
			Current: versionPayload{
				Version: toVersionStringPtr(state.Current.Version),
				Dirty:   state.Current.Dirty,
			},
			TotalPending: len(files),
			Files:        files,
		},
	})
}

func (h *MigrationHandler) handleUp(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeMethodNotAllowed(w)
		return
	}
	req, err := h.decodeOptionalStepsRequest(r)
	if err != nil {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: err.Error(),
			},
		})
		return
	}
	res, err := h.service.Up(r.Context(), req.Steps)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	h.writeJSON(w, http.StatusOK, apiResponse{
		Status:    "ok",
		Timestamp: h.now().UTC().Format(time.RFC3339),
		Data:      toOperationPayload(res),
	})
}

func (h *MigrationHandler) handleDown(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeMethodNotAllowed(w)
		return
	}
	req, err := h.decodeRequiredStepsRequest(r)
	if err != nil {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: err.Error(),
			},
		})
		return
	}
	res, err := h.service.Down(r.Context(), req.Steps)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	h.writeJSON(w, http.StatusOK, apiResponse{
		Status:    "ok",
		Timestamp: h.now().UTC().Format(time.RFC3339),
		Data:      toOperationPayload(res),
	})
}

func (h *MigrationHandler) handleRunFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeMethodNotAllowed(w)
		return
	}
	body, err := h.readJSONBody(r)
	if err != nil {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: err.Error(),
			},
		})
		return
	}
	var req runFileRequest
	if len(body) == 0 {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: "request body is required",
			},
		})
		return
	}
	if err := json.Unmarshal(body, &req); err != nil {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: "invalid json body",
			},
		})
		return
	}
	req.File = strings.TrimSpace(req.File)
	if req.File == "" {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: "file is required",
			},
		})
		return
	}

	res, err := h.service.RunSpecificFile(r.Context(), req.File)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	h.writeJSON(w, http.StatusOK, apiResponse{
		Status:    "ok",
		Timestamp: h.now().UTC().Format(time.RFC3339),
		Data:      toOperationPayload(res),
	})
}

func (h *MigrationHandler) handleAcknowledgeManual(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		h.writeMethodNotAllowed(w)
		return
	}
	body, err := h.readJSONBody(r)
	if err != nil {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: err.Error(),
			},
		})
		return
	}
	if len(body) == 0 {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: "request body is required",
			},
		})
		return
	}

	var req acknowledgeManualRequest
	if err := json.Unmarshal(body, &req); err != nil {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: "invalid json body",
			},
		})
		return
	}
	req.File = strings.TrimSpace(req.File)
	req.Reason = strings.TrimSpace(req.Reason)
	if req.File == "" {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: "file is required",
			},
		})
		return
	}
	if req.Reason == "" {
		h.writeJSON(w, http.StatusBadRequest, apiResponse{
			Status:    "fail",
			Timestamp: h.now().UTC().Format(time.RFC3339),
			Error: &apiError{
				Code:    "invalid_input",
				Message: "reason is required",
			},
		})
		return
	}

	res, err := h.service.AcknowledgeManualFile(r.Context(), req.File, req.Reason)
	if err != nil {
		h.writeServiceError(w, err)
		return
	}
	h.writeJSON(w, http.StatusOK, apiResponse{
		Status:    "ok",
		Timestamp: h.now().UTC().Format(time.RFC3339),
		Data:      toOperationPayload(res),
	})
}

func (h *MigrationHandler) writeServiceError(w http.ResponseWriter, err error) {
	status, code, message := mapServiceError(err)
	h.writeJSON(w, status, apiResponse{
		Status:    "fail",
		Timestamp: h.now().UTC().Format(time.RFC3339),
		Error: &apiError{
			Code:    code,
			Message: message,
		},
	})
}

func mapServiceError(err error) (status int, code string, message string) {
	var svcErr *migrest.ServiceError
	if errors.As(err, &svcErr) {
		switch svcErr.Kind {
		case migrest.ErrorKindInvalidInput:
			return http.StatusBadRequest, "invalid_input", svcErr.Error()
		case migrest.ErrorKindForbidden:
			return http.StatusForbidden, "forbidden", svcErr.Error()
		case migrest.ErrorKindConflict:
			return http.StatusConflict, "conflict", svcErr.Error()
		default:
			return http.StatusInternalServerError, "internal_error", svcErr.Error()
		}
	}
	return http.StatusInternalServerError, "internal_error", err.Error()
}

func (h *MigrationHandler) withAuth(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if h.enableAPIKey {
			given := strings.TrimSpace(r.Header.Get("X-API-Key"))
			if subtle.ConstantTimeCompare([]byte(given), []byte(h.apiKey)) != 1 {
				h.writeJSON(w, http.StatusUnauthorized, apiResponse{
					Status:    "fail",
					Timestamp: h.now().UTC().Format(time.RFC3339),
					Error: &apiError{
						Code:    "unauthorized",
						Message: "invalid api key",
					},
				})
				return
			}
		}
		next(w, r)
	}
}

func (h *MigrationHandler) decodeOptionalStepsRequest(r *http.Request) (stepsRequest, error) {
	body, err := h.readJSONBody(r)
	if err != nil {
		return stepsRequest{}, err
	}
	if len(body) == 0 {
		return stepsRequest{}, nil
	}

	var req stepsRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return stepsRequest{}, errors.New("invalid json body")
	}
	if req.Steps < 0 {
		return stepsRequest{}, errors.New("steps must be >= 0")
	}
	return req, nil
}

func (h *MigrationHandler) decodeRequiredStepsRequest(r *http.Request) (stepsRequest, error) {
	body, err := h.readJSONBody(r)
	if err != nil {
		return stepsRequest{}, err
	}
	if len(body) == 0 {
		return stepsRequest{}, errors.New("request body is required")
	}

	var req stepsRequest
	if err := json.Unmarshal(body, &req); err != nil {
		return stepsRequest{}, errors.New("invalid json body")
	}
	if req.Steps <= 0 {
		return stepsRequest{}, errors.New("steps must be > 0")
	}
	return req, nil
}

func (h *MigrationHandler) readJSONBody(r *http.Request) ([]byte, error) {
	if r.Body == nil {
		return nil, nil
	}
	defer r.Body.Close()
	limited := io.LimitReader(r.Body, maxBodyBytes)
	body, err := io.ReadAll(limited)
	if err != nil {
		return nil, err
	}
	trimmed := strings.TrimSpace(string(body))
	if trimmed == "" {
		return nil, nil
	}
	return []byte(trimmed), nil
}

func (h *MigrationHandler) writeMethodNotAllowed(w http.ResponseWriter) {
	h.writeJSON(w, http.StatusMethodNotAllowed, apiResponse{
		Status:    "fail",
		Timestamp: h.now().UTC().Format(time.RFC3339),
		Error: &apiError{
			Code:    "method_not_allowed",
			Message: "method not allowed",
		},
	})
}

func (h *MigrationHandler) writeJSON(w http.ResponseWriter, status int, payload apiResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	enc := json.NewEncoder(w)
	enc.SetEscapeHTML(false)
	_ = enc.Encode(payload)
}

func toOperationPayload(res migrest.OperationResult) operationPayload {
	return operationPayload{
		Operation:      res.Operation,
		OperationID:    res.OperationID,
		FileName:       res.FileName,
		RequestedSteps: res.RequestedSteps,
		RecordedSteps:  len(res.AppliedFiles),
		NoChange:       res.NoChange,
		AppliedFiles:   res.AppliedFiles,
		ManifestSHA256: res.ManifestSHA256,
		Before: versionPayload{
			Version: toVersionStringPtr(res.Before.Version),
			Dirty:   res.Before.Dirty,
		},
		After: versionPayload{
			Version: toVersionStringPtr(res.After.Version),
			Dirty:   res.After.Dirty,
		},
	}
}

func toVersionStringPtr(v *uint64) *string {
	if v == nil {
		return nil
	}
	out := strconv.FormatUint(*v, 10)
	return &out
}
