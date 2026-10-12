package migrest

import "fmt"

// ErrorKind classifies service errors for HTTP mapping.
type ErrorKind string

const (
	ErrorKindInvalidInput ErrorKind = "invalid_input"
	ErrorKindForbidden    ErrorKind = "forbidden"
	ErrorKindConflict     ErrorKind = "conflict"
	ErrorKindInternal     ErrorKind = "internal"
)

// ServiceError wraps domain errors with machine-readable kind.
type ServiceError struct {
	Kind    ErrorKind
	Message string
	Err     error
}

func (e *ServiceError) Error() string {
	if e == nil {
		return ""
	}
	if e.Err == nil {
		return e.Message
	}
	if e.Message == "" {
		return e.Err.Error()
	}
	return fmt.Sprintf("%s: %v", e.Message, e.Err)
}

func (e *ServiceError) Unwrap() error {
	if e == nil {
		return nil
	}
	return e.Err
}

func newServiceError(kind ErrorKind, message string, err error) error {
	return &ServiceError{
		Kind:    kind,
		Message: message,
		Err:     err,
	}
}
