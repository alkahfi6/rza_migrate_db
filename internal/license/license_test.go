package license

import (
	"errors"
	"strings"
	"testing"
	"time"
)

func TestGenerateAndValidateText(t *testing.T) {
	t.Parallel()

	now := time.Date(2026, 3, 4, 0, 0, 0, 0, time.UTC)
	expires := time.Date(2026, 6, 4, 0, 0, 0, 0, time.UTC)
	text, err := Generate("template_repo", now, expires)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	info, err := ValidateText(text, now)
	if err != nil {
		t.Fatalf("ValidateText failed: %v", err)
	}
	if info.Owner != "template_repo" {
		t.Fatalf("owner mismatch: got %q", info.Owner)
	}
}

func TestValidateTextExpired(t *testing.T) {
	t.Parallel()

	issued := time.Date(2026, 3, 4, 0, 0, 0, 0, time.UTC)
	expires := time.Date(2026, 6, 4, 0, 0, 0, 0, time.UTC)
	text, err := Generate("template_repo", issued, expires)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}

	_, err = ValidateText(text, time.Date(2026, 6, 5, 0, 0, 0, 0, time.UTC))
	if !errors.Is(err, ErrExpired) {
		t.Fatalf("expected ErrExpired, got %v", err)
	}
}

func TestValidateTextInvalidSignature(t *testing.T) {
	t.Parallel()

	now := time.Date(2026, 3, 4, 0, 0, 0, 0, time.UTC)
	expires := time.Date(2026, 6, 4, 0, 0, 0, 0, time.UTC)
	text, err := Generate("template_repo", now, expires)
	if err != nil {
		t.Fatalf("Generate failed: %v", err)
	}
	lines := strings.Split(strings.TrimSpace(text), "\n")
	lines[1] = lines[1] + "tampered"
	tampered := strings.Join(lines, "\n") + "\n"

	_, err = ValidateText(tampered, now)
	if !errors.Is(err, ErrInvalidSignature) {
		t.Fatalf("expected ErrInvalidSignature, got %v", err)
	}
}
