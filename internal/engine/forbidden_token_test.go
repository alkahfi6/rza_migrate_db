package engine

import (
	"testing"

	"rza_migrate_db/internal/sqlparse"
)

func TestDetectForbiddenTokensIgnoresCommentsAndStrings(t *testing.T) {
	t.Parallel()

	sql := `
-- USE master;
/* CREATE LOGIN app_login */
SELECT 'GO';
SELECT 'CREATE USER app_user';
`

	got := sqlparse.DetectForbiddenTokens(sql)
	if len(got) != 0 {
		t.Fatalf("expected no forbidden tokens, got: %#v", got)
	}
}

func TestDetectForbiddenTokensDetectsRealTokens(t *testing.T) {
	t.Parallel()

	sql := `
USE master;
CREATE USER app_user FOR LOGIN app_login;
GO
`

	got := sqlparse.DetectForbiddenTokens(sql)
	expected := map[string]bool{
		"USE":         true,
		"CREATE USER": true,
		"GO":          true,
	}
	if len(got) != len(expected) {
		t.Fatalf("unexpected token count: got=%d want=%d tokens=%#v", len(got), len(expected), got)
	}
	for _, token := range got {
		if !expected[token] {
			t.Fatalf("unexpected token detected: %s", token)
		}
	}
}
