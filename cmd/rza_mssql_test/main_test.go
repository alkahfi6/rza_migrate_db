package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadCredentialConfigBuildsEscapedSQLServerURL(t *testing.T) {
	path := filepath.Join(t.TempDir(), "credential.nfm")
	content := "database_type=mssql\nhost=127.0.0.1:1433\nusername=sa\npassword=P@ss word!\ninitial_catalog=MigrationTest\n"
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write config: %v", err)
	}

	cfg, err := loadCredentialConfig(path)
	if err != nil {
		t.Fatalf("load config: %v", err)
	}
	dsn, err := cfg.databaseURL()
	if err != nil {
		t.Fatalf("build URL: %v", err)
	}
	if !strings.HasPrefix(dsn, "sqlserver://sa:") || !strings.Contains(dsn, "database=MigrationTest") {
		t.Fatalf("unexpected database URL: %s", dsn)
	}
}

func TestResetSchemaUsesDependencyRetry(t *testing.T) {
	t.Parallel()

	// Keep the retry contract visible: a schema-bound dependent can fail in one
	// pass and become removable after its own dependent has been dropped.
	source, err := os.ReadFile("main.go")
	if err != nil {
		t.Fatalf("read main.go: %v", err)
	}
	text := string(source)
	for _, fragment := range []string{
		"WHILE @pass < 2000",
		"BEGIN TRY DROP",
		"SET @dropped = @dropped + 1",
		"reset could not drop all dbo modules",
	} {
		if !strings.Contains(text, fragment) {
			t.Fatalf("reset retry contract missing %q", fragment)
		}
	}
}

func TestLoadCredentialConfigRejectsIncompleteFile(t *testing.T) {
	path := filepath.Join(t.TempDir(), "credential.nfm")
	if err := os.WriteFile(path, []byte("host=localhost\n"), 0o600); err != nil {
		t.Fatalf("write config: %v", err)
	}
	if _, err := loadCredentialConfig(path); err == nil {
		t.Fatal("expected incomplete config error")
	}
}

func TestDatabaseURLSupportsNamedInstance(t *testing.T) {
	cfg := credentialConfig{
		DatabaseType: "mssql",
		Host:         `localhost\SQLEXPRESS`,
		Username:     "sa",
		Password:     "Password1!",
		Database:     "MigrationTest",
	}
	dsn, err := cfg.databaseURL()
	if err != nil {
		t.Fatalf("build URL: %v", err)
	}
	if !strings.Contains(dsn, "instance=SQLEXPRESS") {
		t.Fatalf("expected instance query parameter: %s", dsn)
	}
}
