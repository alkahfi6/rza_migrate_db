package engine

import (
	"strings"
	"testing"
)

func TestDetectUnsafeAlterColumn_NullToNotNull(t *testing.T) {
	t.Parallel()

	prev := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NULL
);
`
	next := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  status VARCHAR(10) NOT NULL
);
`

	issues, err := detectUnsafeAlterColumn(prev, next)
	if err != nil {
		t.Fatalf("detectUnsafeAlterColumn returned error: %v", err)
	}
	if !containsIssuePrefix(issues, "NULL_TO_NOT_NULL(status)") {
		t.Fatalf("expected NULL_TO_NOT_NULL issue, got %#v", issues)
	}
}

func TestDetectUnsafeAlterColumn_StringShrink(t *testing.T) {
	t.Parallel()

	prev := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  code VARCHAR(50) NULL
);
`
	next := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  code VARCHAR(10) NULL
);
`

	issues, err := detectUnsafeAlterColumn(prev, next)
	if err != nil {
		t.Fatalf("detectUnsafeAlterColumn returned error: %v", err)
	}
	if !containsIssuePrefix(issues, "SHRINK_STRING(code:50->10)") {
		t.Fatalf("expected SHRINK_STRING issue, got %#v", issues)
	}
}

func TestDetectUnsafeAlterColumn_TypeToNumeric(t *testing.T) {
	t.Parallel()

	prev := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  amount VARCHAR(20) NULL
);
`
	next := `
CREATE TABLE dbo.TableA (
  id INT NOT NULL,
  amount DECIMAL(12,2) NULL
);
`

	issues, err := detectUnsafeAlterColumn(prev, next)
	if err != nil {
		t.Fatalf("detectUnsafeAlterColumn returned error: %v", err)
	}
	if !containsIssuePrefix(issues, "TYPE_TO_NUMERIC(amount:VARCHAR->DECIMAL)") {
		t.Fatalf("expected TYPE_TO_NUMERIC issue, got %#v", issues)
	}
}

func containsIssuePrefix(issues []string, exact string) bool {
	for _, issue := range issues {
		if strings.EqualFold(issue, exact) {
			return true
		}
	}
	return false
}
