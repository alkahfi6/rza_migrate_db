package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestBootstrapSeedsCoreAuditPairWithoutOverwrite(t *testing.T) {
	root := t.TempDir()
	cfg := config.Default(root)
	cfg.BootstrapMode = true
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	engine := New(cfg)
	if err := engine.seedCoreAuditForBootstrap(&report); err != nil {
		t.Fatal(err)
	}
	for _, name := range []string{coreAuditUpName, coreAuditDownName} {
		raw, err := os.ReadFile(filepath.Join(root, "migrations", name))
		if err != nil || len(raw) == 0 {
			t.Fatalf("seeded %s invalid: bytes=%d err=%v", name, len(raw), err)
		}
	}
	seededUp, _ := os.ReadFile(filepath.Join(root, "migrations", coreAuditUpName))
	for _, required := range []string{
		"dbo.mig_operation_journal",
		"EXEC sys.sp_executesql N'CREATE OR ALTER PROCEDURE dbo.usp_mig_log_event",
		"EXEC sys.sp_executesql N'CREATE OR ALTER PROCEDURE dbo.usp_mig_try_get_lock",
	} {
		if !strings.Contains(string(seededUp), required) {
			t.Fatalf("seeded core audit is missing %q", required)
		}
	}
	upPath := filepath.Join(root, "migrations", coreAuditUpName)
	before, _ := os.ReadFile(upPath)
	if err := engine.seedCoreAuditForBootstrap(&report); err != nil {
		t.Fatal(err)
	}
	after, _ := os.ReadFile(upPath)
	if string(after) != string(before) {
		t.Fatal("existing reviewed core pair was overwritten")
	}
}

func TestBootstrapRejectsPartialCoreAuditPair(t *testing.T) {
	root := t.TempDir()
	cfg := config.Default(root)
	cfg.BootstrapMode = true
	if err := os.MkdirAll(filepath.Join(root, "migrations"), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "migrations", coreAuditUpName), []byte("reviewed"), 0o600); err != nil {
		t.Fatal(err)
	}
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	if err := New(cfg).seedCoreAuditForBootstrap(&report); err != nil {
		t.Fatal(err)
	}
	if !hasRule(report.Fail.Items, "CORE_AUDIT_PARTIAL") {
		t.Fatalf("partial core pair not rejected: %#v", report.Fail.Items)
	}
}

func TestBootstrapFmtFormatsSafeFiles(t *testing.T) {
	root := t.TempDir()
	cfg := config.Default(root)
	cfg.BootstrapMode = true
	unsafePath := filepath.Join(root, "db_object", "indexes", "dbo.JHCLDR.sql")
	safePath := filepath.Join(root, "db_object", "procedures", "dbo.Safe.sql")
	if err := os.MkdirAll(filepath.Dir(unsafePath), 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(filepath.Dir(safePath), 0o755); err != nil {
		t.Fatal(err)
	}
	unsafe := "CREATE CLUSTERED INDEX IX_JHCLDR_CJDATE\nON JHCLDR (CJDATE);\nCREATE NONCLUSTERED INDEX IX_JHCLDR_CBRNBR\nON JHCLDR (CBRNBR);\n"
	safe := "CREATE OR ALTER PROCEDURE dbo.Safe AS\r\nBEGIN\r\nSELECT 1;   \r\nEND\r\n"
	if err := os.WriteFile(unsafePath, []byte(unsafe), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(safePath, []byte(safe), 0o600); err != nil {
		t.Fatal(err)
	}
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	if err := New(cfg).runFmt(&report); err != nil {
		t.Fatal(err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("bootstrap formatting result: pass=%#v fail=%#v", report.Pass.Items, report.Fail.Items)
	}
	gotUnsafe, _ := os.ReadFile(unsafePath)
	if string(gotUnsafe) == unsafe {
		t.Fatal("canonical index file was not formatted")
	}
	gotSafe, _ := os.ReadFile(safePath)
	if string(gotSafe) == safe {
		t.Fatal("safe file was not normalized")
	}
}

func TestFmtSafelyQualifiesIndexTargetFromFileContract(t *testing.T) {
	root := t.TempDir()
	cfg := config.Default(root)
	path := filepath.Join(root, "db_object", "indexes", "dbo.JHCLDR.sql")
	if err := os.MkdirAll(filepath.Dir(path), 0o755); err != nil {
		t.Fatal(err)
	}
	input := "CREATE CLUSTERED INDEX IX_JHCLDR_CJDATE\nON JHCLDR (CJDATE);\n"
	if err := os.WriteFile(path, []byte(input), 0o600); err != nil {
		t.Fatal(err)
	}
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	if err := New(cfg).runFmt(&report); err != nil {
		t.Fatal(err)
	}
	if report.Fail.Count != 0 {
		t.Fatalf("known schema qualification was rejected: %#v", report.Fail.Items)
	}
	raw, _ := os.ReadFile(path)
	if string(raw) == input {
		t.Fatalf("index target was not canonicalized: %s", raw)
	}
}
