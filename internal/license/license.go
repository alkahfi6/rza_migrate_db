package license

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

const (
	fileHeader = "RZA_LICENSE_V1"
	dateLayout = "2006-01-02"
)

var (
	// Embedded signing key for offline license verification.
	// Rotate only via new binary release.
	signingKey = []byte("RZA_LICENSE_SIGNING_KEY_V1_20260304")

	ErrExpired          = errors.New("license expired")
	ErrInvalidFormat    = errors.New("invalid license format")
	ErrInvalidSignature = errors.New("invalid license signature")
)

type Payload struct {
	Owner     string `json:"owner"`
	IssuedAt  string `json:"issued_at"`
	ExpiresAt string `json:"expires_at"`
}

type Info struct {
	Owner     string
	IssuedAt  time.Time
	ExpiresAt time.Time
}

func Generate(owner string, now time.Time, expiresAt time.Time) (string, error) {
	owner = strings.TrimSpace(owner)
	if owner == "" {
		return "", fmt.Errorf("owner must not be empty")
	}
	if !expiresAt.After(now) {
		return "", fmt.Errorf("expires_at must be after issued_at")
	}

	payload := Payload{
		Owner:     owner,
		IssuedAt:  now.Format(dateLayout),
		ExpiresAt: expiresAt.Format(dateLayout),
	}
	rawPayload, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}
	payloadEncoded := base64.StdEncoding.EncodeToString(rawPayload)
	signature := sign(payloadEncoded)

	var b strings.Builder
	b.WriteString(fileHeader)
	b.WriteByte('\n')
	b.WriteString(payloadEncoded)
	b.WriteByte('\n')
	b.WriteString(signature)
	b.WriteByte('\n')
	return b.String(), nil
}

func ValidateFile(path string, now time.Time) (Info, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return Info{}, err
	}
	return ValidateText(string(raw), now)
}

func ValidateText(text string, now time.Time) (Info, error) {
	lines := splitNonEmptyLines(text)
	if len(lines) != 3 {
		return Info{}, ErrInvalidFormat
	}
	if lines[0] != fileHeader {
		return Info{}, ErrInvalidFormat
	}

	payloadEncoded := lines[1]
	signatureHex := strings.ToLower(lines[2])
	if !hmac.Equal([]byte(signatureHex), []byte(sign(payloadEncoded))) {
		return Info{}, ErrInvalidSignature
	}

	payloadBytes, err := base64.StdEncoding.DecodeString(payloadEncoded)
	if err != nil {
		return Info{}, ErrInvalidFormat
	}

	var payload Payload
	if err := json.Unmarshal(payloadBytes, &payload); err != nil {
		return Info{}, ErrInvalidFormat
	}
	issuedAt, err := time.Parse(dateLayout, payload.IssuedAt)
	if err != nil {
		return Info{}, ErrInvalidFormat
	}
	expiresAt, err := time.Parse(dateLayout, payload.ExpiresAt)
	if err != nil {
		return Info{}, ErrInvalidFormat
	}
	if strings.TrimSpace(payload.Owner) == "" {
		return Info{}, ErrInvalidFormat
	}

	nowDate := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, time.UTC)
	if nowDate.After(expiresAt) {
		return Info{
			Owner:     payload.Owner,
			IssuedAt:  issuedAt,
			ExpiresAt: expiresAt,
		}, ErrExpired
	}

	return Info{
		Owner:     payload.Owner,
		IssuedAt:  issuedAt,
		ExpiresAt: expiresAt,
	}, nil
}

func sign(payloadEncoded string) string {
	m := hmac.New(sha256.New, signingKey)
	_, _ = m.Write([]byte(payloadEncoded))
	return hex.EncodeToString(m.Sum(nil))
}

func splitNonEmptyLines(text string) []string {
	rawLines := strings.Split(strings.ReplaceAll(text, "\r\n", "\n"), "\n")
	lines := make([]string, 0, len(rawLines))
	for _, line := range rawLines {
		trimmed := strings.TrimSpace(line)
		if trimmed == "" {
			continue
		}
		lines = append(lines, trimmed)
	}
	return lines
}
