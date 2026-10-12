package main

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestLoadNFMCredentialsAndBuildURL(t *testing.T) {
	path := filepath.Join(t.TempDir(), "mssql_credential.nfm")
	content := "database_type=mssql\nhost=db.example.test,1444\nusername=test_user\npassword=secret\ninitial_catalog=MigrationTest\n"
	if err := os.WriteFile(path, []byte(content), 0o600); err != nil {
		t.Fatalf("write credentials: %v", err)
	}

	creds, err := loadNFMCredentials(path)
	if err != nil {
		t.Fatalf("loadNFMCredentials failed: %v", err)
	}
	dsn, err := creds.databaseURL()
	if err != nil {
		t.Fatalf("databaseURL failed: %v", err)
	}
	if !strings.Contains(dsn, "db.example.test:1444") || !strings.Contains(dsn, "database=MigrationTest") {
		t.Fatalf("unexpected DSN shape: %s", dsn)
	}
}

func TestResetGuards(t *testing.T) {
	if !isSystemDatabase("master") {
		t.Fatal("master must be guarded")
	}
	if isSystemDatabase("MigrationTest") {
		t.Fatal("application database must not be classified as system database")
	}
}

func TestValidateDatabaseName(t *testing.T) {
	for _, valid := range []string{"RZA_RT_BOND_REPLICATE", "db with space", "db]quoted"} {
		if err := validateDatabaseName(valid); err != nil {
			t.Fatalf("valid database %q rejected: %v", valid, err)
		}
	}
	for _, invalid := range []string{"", "bad\nname", strings.Repeat("x", 129)} {
		if err := validateDatabaseName(invalid); err == nil {
			t.Fatalf("invalid database %q accepted", invalid)
		}
	}
}

func TestDropStatementQuotesIdentifiers(t *testing.T) {
	statement, err := dropStatement(resetObject{Kind: "table", Schema: "dbo", Name: "audit]log"})
	if err != nil {
		t.Fatalf("dropStatement failed: %v", err)
	}
	if statement != "DROP TABLE [dbo].[audit]]log];" {
		t.Fatalf("unexpected statement: %s", statement)
	}
}
