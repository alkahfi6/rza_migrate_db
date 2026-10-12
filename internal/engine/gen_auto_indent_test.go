package engine

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
)

func TestAutoIndentChangedModuleSourcesPreservesComments(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	cfg := config.Default(root)
	e := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)

	rel := filepath.ToSlash(filepath.Join("db_object", "procedures", "dbo.ProcA.sql"))
	abs := cfg.RepoPath(rel)
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	before := "CREATE OR ALTER PROCEDURE dbo.ProcA\nAS\nBEGIN\n-- keep this comment\nIF EXISTS (SELECT 1)\nSELECT 1;\nEND\n"
	if err := os.WriteFile(abs, []byte(before), 0o644); err != nil {
		t.Fatalf("write source failed: %v", err)
	}

	changed, files, err := e.autoIndentChangedModuleSources(&report, []desiredObject{
		{
			Category:   "procedures",
			Object:     "dbo.ProcA",
			SourceFile: rel,
		},
	})
	if err != nil {
		t.Fatalf("autoIndentChangedModuleSources failed: %v", err)
	}
	if changed != 1 {
		t.Fatalf("expected 1 changed file, got %d", changed)
	}
	if len(files) != 1 || files[0] != rel {
		t.Fatalf("unexpected changed files: %#v", files)
	}

	afterRaw, err := os.ReadFile(abs)
	if err != nil {
		t.Fatalf("read source failed: %v", err)
	}
	after := string(afterRaw)
	if !strings.Contains(after, "-- keep this comment") {
		t.Fatalf("comment must be preserved:\n%s", after)
	}
	if !strings.Contains(after, "IF EXISTS (SELECT 1)\n    SELECT 1;") {
		t.Fatalf("expected indented IF body:\n%s", after)
	}
}

func TestAutoIndentChangedModuleSourcesPreservesBlockCommentIndentExactly(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	cfg := config.Default(root)
	e := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	rel := filepath.ToSlash(filepath.Join("db_object", "procedures", "dbo.ProcComment.sql"))
	abs := cfg.RepoPath(rel)
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatal(err)
	}
	comment := "/*\n               data saya ...\n       BEGIN CASE END\n*/"
	before := "CREATE OR ALTER PROCEDURE dbo.ProcComment\nAS\nBEGIN\n" + comment + "\nIF EXISTS (SELECT 1)\nSELECT 1;\nEND\n"
	if err := os.WriteFile(abs, []byte(before), 0o644); err != nil {
		t.Fatal(err)
	}
	changed, _, err := e.autoIndentChangedModuleSources(&report, []desiredObject{{Category: "procedures", Object: "dbo.ProcComment", SourceFile: rel}})
	if err != nil {
		t.Fatal(err)
	}
	if changed != 1 || report.Fail.Count != 0 {
		t.Fatalf("unexpected formatting result changed=%d fail=%#v", changed, report.Fail.Items)
	}
	after, err := os.ReadFile(abs)
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(after), comment) {
		t.Fatalf("block comment indentation changed:\n%s", after)
	}
	if !strings.Contains(string(after), "*/\n  IF EXISTS (SELECT 1)\n    SELECT 1;") {
		t.Fatalf("comment keywords affected following SQL:\n%s", after)
	}
}

func TestAutoIndentChangedModuleSourcesDoesNotPartiallyWriteOnSafetyFailure(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	cfg := config.Default(root)
	e := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
	safeRel := filepath.ToSlash(filepath.Join("db_object", "procedures", "dbo.Safe.sql"))
	unsafeRel := filepath.ToSlash(filepath.Join("db_object", "procedures", "dbo.Unsafe.sql"))
	if err := os.MkdirAll(filepath.Dir(cfg.RepoPath(safeRel)), 0o755); err != nil {
		t.Fatal(err)
	}
	safeBefore := "CREATE OR ALTER PROCEDURE dbo.Safe\nAS\nBEGIN\nIF EXISTS (SELECT 1)\nSELECT 1;\nEND\n"
	// The spaces before the newline are part of the multiline string literal.
	// Repository trailing-space normalization would alter the literal, so the
	// semantic gate must reject the batch before Safe.sql is written.
	unsafeBefore := "CREATE OR ALTER PROCEDURE dbo.Unsafe\nAS\nBEGIN\nDECLARE @sql nvarchar(max) = N'SELECT 1   \nFROM dbo.T';\nEND\n"
	if err := os.WriteFile(cfg.RepoPath(safeRel), []byte(safeBefore), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(cfg.RepoPath(unsafeRel), []byte(unsafeBefore), 0o644); err != nil {
		t.Fatal(err)
	}
	changed, files, err := e.autoIndentChangedModuleSources(&report, []desiredObject{
		{Category: "procedures", Object: "dbo.Safe", SourceFile: safeRel},
		{Category: "procedures", Object: "dbo.Unsafe", SourceFile: unsafeRel},
	})
	if err != nil {
		t.Fatal(err)
	}
	if changed != 0 || len(files) != 0 || !hasFailRuleWithMessage(report.Fail.Items, "FORMAT_SAFETY_BLOCKED", "") {
		t.Fatalf("expected atomic safety rejection, changed=%d files=%#v fail=%#v", changed, files, report.Fail.Items)
	}
	after, err := os.ReadFile(cfg.RepoPath(safeRel))
	if err != nil {
		t.Fatal(err)
	}
	if string(after) != safeBefore {
		t.Fatalf("safe file was partially written despite another file failing safety:\n%s", after)
	}
}

func TestAutoIndentChangedModuleSourcesFormatsTablesAndIndexes(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	cfg := config.Default(root)
	e := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)

	tableRel := filepath.ToSlash(filepath.Join("db_object", "tables", "dbo.TableA.sql"))
	tableAbs := cfg.RepoPath(tableRel)
	if err := os.MkdirAll(filepath.Dir(tableAbs), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	tableBefore := "CREATE TABLE dbo.TableA (\n  id INT NOT NULL\n  ,status VARCHAR(10) NULL\n);\n"
	if err := os.WriteFile(tableAbs, []byte(tableBefore), 0o644); err != nil {
		t.Fatalf("write source failed: %v", err)
	}

	indexRel := filepath.ToSlash(filepath.Join("db_object", "indexes", "dbo.TableA.sql"))
	indexAbs := cfg.RepoPath(indexRel)
	if err := os.MkdirAll(filepath.Dir(indexAbs), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	indexBefore := "CREATE INDEX IX_TableA_B ON dbo.TableA (ColB) ;\nCREATE INDEX IX_TableA_A ON dbo.TableA (ColA)\n"
	if err := os.WriteFile(indexAbs, []byte(indexBefore), 0o644); err != nil {
		t.Fatalf("write source failed: %v", err)
	}

	changed, files, err := e.autoIndentChangedModuleSources(&report, []desiredObject{
		{
			Category:   "tables",
			Object:     "dbo.TableA",
			SourceFile: tableRel,
		},
		{
			Category:   "indexes",
			Object:     "dbo.TableA",
			SourceFile: indexRel,
		},
	})
	if err != nil {
		t.Fatalf("autoIndentChangedModuleSources failed: %v", err)
	}
	if changed != 2 {
		t.Fatalf("expected 2 changed files, got %d", changed)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 changed files list, got %#v", files)
	}

	tableAfterRaw, err := os.ReadFile(tableAbs)
	if err != nil {
		t.Fatalf("read source failed: %v", err)
	}
	tableAfter := string(tableAfterRaw)
	if strings.Contains(tableAfter, "\n  ,") {
		t.Fatalf("expected canonical trailing comma formatting for table:\n%s", tableAfter)
	}
	if !strings.Contains(tableAfter, "id INT NOT NULL,\n  status VARCHAR(10) NULL") {
		t.Fatalf("table canonical formatting mismatch:\n%s", tableAfter)
	}

	indexAfterRaw, err := os.ReadFile(indexAbs)
	if err != nil {
		t.Fatalf("read source failed: %v", err)
	}
	indexAfter := string(indexAfterRaw)
	if !strings.Contains(indexAfter, "CREATE INDEX IX_TableA_A ON dbo.TableA (ColA);") {
		t.Fatalf("expected normalized/sorted first index statement:\n%s", indexAfter)
	}
	if !strings.Contains(indexAfter, "CREATE INDEX IX_TableA_B ON dbo.TableA (ColB);") {
		t.Fatalf("expected normalized second index statement:\n%s", indexAfter)
	}
	if strings.Index(indexAfter, "IX_TableA_A") > strings.Index(indexAfter, "IX_TableA_B") {
		t.Fatalf("expected indexes sorted by name after canonical formatting:\n%s", indexAfter)
	}
}

func TestCanonicalizeIndexSnapshotFormattingPreservesUnsupportedIndexSource(t *testing.T) {
	t.Parallel()

	sql := `CREATE INDEX IX_TableA_Code ON dbo.TableA(Code);
CREATE COLUMNSTORE INDEX IX_TableA_BadColumnstore ON dbo.TableA(Code);`
	got := canonicalizeIndexSnapshotFormatting(sql, "dbo.TableA")
	if !strings.Contains(got, "CREATE COLUMNSTORE INDEX IX_TableA_BadColumnstore") {
		t.Fatalf("unsupported index statement must be preserved for validator, got:\n%s", got)
	}
}

func TestCanonicalizeTableAndIndexFormattingPreserveCommentLayout(t *testing.T) {
	t.Parallel()
	table := "CREATE TABLE dbo.TableA (\n/*\n               data saya ...\n*/\n  Id INT NOT NULL\n);\n"
	if got := canonicalizeTableSnapshotFormatting(table, "dbo.TableA"); got != table {
		t.Fatalf("comment-bearing table must retain exact layout\nwant:\n%s\ngot:\n%s", table, got)
	}
	index := "/*\n               index remark\n*/\nCREATE INDEX IX_TableA_Id ON dbo.TableA(Id);\n"
	if got := canonicalizeIndexSnapshotFormatting(index, "dbo.TableA"); got != index {
		t.Fatalf("comment-bearing index must retain exact layout\nwant:\n%s\ngot:\n%s", index, got)
	}
}

func TestFormattingSafetyIssueProtectsCommentsLiteralsAndIdentifierCase(t *testing.T) {
	t.Parallel()
	if issue := formattingSafetyIssue("procedures", "SELECT  Code FROM dbo.T;\n", "SELECT Code\nFROM dbo.T;\n"); issue != "" {
		t.Fatalf("whitespace-only formatting should be safe: %s", issue)
	}
	if issue := formattingSafetyIssue("procedures", "  -- keep\nSELECT 1;\n", "-- keep\nSELECT 1;\n"); issue == "" {
		t.Fatal("comment indentation change must be blocked")
	}
	if issue := formattingSafetyIssue("tables", "  -- field info   \nId INT;\n", "  -- field info\nId INT;\n"); issue != "" {
		t.Fatalf("trailing comment whitespace normalization should be safe: %s", issue)
	}
	if issue := formattingSafetyIssue("procedures", "SELECT N'a  b';\n", "SELECT N'a b';\n"); issue == "" {
		t.Fatal("literal content change must be blocked")
	}
	if issue := formattingSafetyIssue("procedures", "SELECT Code FROM dbo.T;\n", "SELECT code FROM dbo.T;\n"); issue == "" {
		t.Fatal("case-sensitive identifier change must be blocked")
	}
}

func TestAutoIndentChangedModuleSourcesSkipsUnsupportedCategory(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	cfg := config.Default(root)
	e := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)

	rel := filepath.ToSlash(filepath.Join("db_object", "patch_like", "x.sql"))
	abs := cfg.RepoPath(rel)
	if err := os.MkdirAll(filepath.Dir(abs), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}
	original := "SELECT 1;\n"
	if err := os.WriteFile(abs, []byte(original), 0o644); err != nil {
		t.Fatalf("write source failed: %v", err)
	}

	changed, files, err := e.autoIndentChangedModuleSources(&report, []desiredObject{
		{
			Category:   "patch",
			Object:     "patch_like",
			SourceFile: rel,
		},
	})
	if err != nil {
		t.Fatalf("autoIndentChangedModuleSources failed: %v", err)
	}
	if changed != 0 {
		t.Fatalf("expected 0 changed file, got %d", changed)
	}
	if len(files) != 0 {
		t.Fatalf("expected empty changed files, got %#v", files)
	}

	afterRaw, err := os.ReadFile(abs)
	if err != nil {
		t.Fatalf("read source failed: %v", err)
	}
	if string(afterRaw) != original {
		t.Fatalf("unsupported category source should not be mutated")
	}
}

func TestAutoIndentPendingPatchSourcesPreservesComments(t *testing.T) {
	t.Parallel()

	root := t.TempDir()
	cfg := config.Default(root)
	e := New(cfg)
	report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)

	relDir := filepath.ToSlash(filepath.Join("patch", "Patch A"))
	upRel := filepath.ToSlash(filepath.Join(relDir, "up.sql"))
	downRel := filepath.ToSlash(filepath.Join(relDir, "down.sql"))
	upAbs := cfg.RepoPath(upRel)
	downAbs := cfg.RepoPath(downRel)
	if err := os.MkdirAll(filepath.Dir(upAbs), 0o755); err != nil {
		t.Fatalf("mkdir failed: %v", err)
	}

	upSQL := "-- keep up comment\nIF EXISTS (SELECT 1)\nUPDATE dbo.TableA SET status='DONE' WHERE status='NEW';\n"
	downSQL := "-- keep down comment\nIF EXISTS (SELECT 1)\nUPDATE dbo.TableA SET status='NEW' WHERE status='DONE';\n"
	if err := os.WriteFile(upAbs, []byte(upSQL), 0o644); err != nil {
		t.Fatalf("write up.sql failed: %v", err)
	}
	if err := os.WriteFile(downAbs, []byte(downSQL), 0o644); err != nil {
		t.Fatalf("write down.sql failed: %v", err)
	}

	changed, files, err := e.autoIndentPendingPatchSources(&report, []patchSource{
		{
			RelDir: relDir,
		},
	}, "safe")
	if err != nil {
		t.Fatalf("autoIndentPendingPatchSources failed: %v", err)
	}
	if changed != 2 {
		t.Fatalf("expected 2 changed files, got %d", changed)
	}
	if len(files) != 2 {
		t.Fatalf("expected 2 file entries, got %#v", files)
	}

	upAfterRaw, err := os.ReadFile(upAbs)
	if err != nil {
		t.Fatalf("read up.sql failed: %v", err)
	}
	upAfter := string(upAfterRaw)
	if !strings.Contains(upAfter, "-- keep up comment") {
		t.Fatalf("up.sql comment must be preserved:\n%s", upAfter)
	}
	if !strings.Contains(upAfter, "IF EXISTS (SELECT 1)\n  UPDATE") {
		t.Fatalf("up.sql should be reindented:\n%s", upAfter)
	}

	downAfterRaw, err := os.ReadFile(downAbs)
	if err != nil {
		t.Fatalf("read down.sql failed: %v", err)
	}
	downAfter := string(downAfterRaw)
	if !strings.Contains(downAfter, "-- keep down comment") {
		t.Fatalf("down.sql comment must be preserved:\n%s", downAfter)
	}
	if !strings.Contains(downAfter, "IF EXISTS (SELECT 1)\n  UPDATE") {
		t.Fatalf("down.sql should be reindented:\n%s", downAfter)
	}
}

func TestPatchFormatModesPreserveUnsafeSource(t *testing.T) {
	t.Parallel()
	root := t.TempDir()
	cfg := config.Default(root)
	e := New(cfg)
	relDir := filepath.ToSlash(filepath.Join("patch", "Unsafe Patch"))
	upRel := filepath.ToSlash(filepath.Join(relDir, "up.sql"))
	downRel := filepath.ToSlash(filepath.Join(relDir, "down.sql"))
	if err := os.MkdirAll(filepath.Dir(cfg.RepoPath(upRel)), 0o755); err != nil {
		t.Fatal(err)
	}
	unsafe := "DECLARE @sql nvarchar(max) = N'SELECT 1   \nFROM dbo.T';\n"
	if err := os.WriteFile(cfg.RepoPath(upRel), []byte(unsafe), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(cfg.RepoPath(downRel), []byte("SELECT 1;\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	patches := []patchSource{{RelDir: relDir}}

	for _, mode := range []string{"off", "safe", "strict"} {
		report := output.NewReport(cfg.OutputNamingPattern, cfg.OutputNamingLineEndings)
		changed, _, err := e.autoIndentPendingPatchSources(&report, patches, mode)
		if err != nil {
			t.Fatalf("mode %s: %v", mode, err)
		}
		if changed != 0 {
			t.Fatalf("mode %s changed unsafe input", mode)
		}
		raw, readErr := os.ReadFile(cfg.RepoPath(upRel))
		if readErr != nil || string(raw) != unsafe {
			t.Fatalf("mode %s mutated unsafe source: %q (%v)", mode, raw, readErr)
		}
		switch mode {
		case "off":
			if report.Fail.Count != 0 {
				t.Fatalf("off mode unexpectedly failed: %#v", report.Fail.Items)
			}
		case "safe":
			if report.Fail.Count != 0 || !hasRule(report.Pass.Items, "PATCH_FORMAT_SKIPPED_UNSAFE") {
				t.Fatalf("safe mode should warn and continue: pass=%#v fail=%#v", report.Pass.Items, report.Fail.Items)
			}
		case "strict":
			if !hasRule(report.Fail.Items, "FORMAT_SAFETY_BLOCKED") {
				t.Fatalf("strict mode should fail: %#v", report.Fail.Items)
			}
		}
	}
}

func TestRunGenSkipsAutoIndentWhenDisabled(t *testing.T) {
	t.Parallel()

	repo := repoRootFromEnginePackageDir(t)
	srcScenario := filepath.Join(repo, "fixtures", "qa_rules", "07_vk_whitespace_only", "input")
	workRoot := t.TempDir()
	if err := copyDirForAutoIndentTest(srcScenario, workRoot); err != nil {
		t.Fatalf("copy scenario failed: %v", err)
	}

	viewRel := filepath.ToSlash(filepath.Join("db_object", "schema", "views", "dbo.ViewA.sql"))
	viewPath := filepath.Join(workRoot, filepath.FromSlash(viewRel))
	beforeRaw, err := os.ReadFile(viewPath)
	if err != nil {
		t.Fatalf("read source view failed: %v", err)
	}

	cfg := config.Default(workRoot)
	cfg.EnableGenAutoIndent = false
	cfg.EnableStrictMode = false
	eng := New(cfg)
	report, err := eng.Run("gen")
	if err != nil {
		t.Fatalf("gen failed: %v", err)
	}
	if report.Fail.Count > 0 {
		t.Fatalf("expected no fail, got: %+v", report.Fail.Items)
	}

	if enabled, ok := report.Notes["gen_auto_indent_enabled"]; !ok || enabled != false {
		t.Fatalf("expected note gen_auto_indent_enabled=false, got %#v", report.Notes["gen_auto_indent_enabled"])
	}
	if _, exists := report.Notes["gen_auto_indent_changed_files"]; exists {
		t.Fatalf("gen_auto_indent_changed_files should be absent when auto-indent disabled")
	}
	if _, exists := report.Notes["gen_auto_indent_patch_files"]; exists {
		t.Fatalf("gen_auto_indent_patch_files should be absent when auto-indent disabled")
	}
	for _, item := range report.Pass.Items {
		if item.Rule == "GEN_AUTO_INDENT_APPLIED" || item.Rule == "GEN_AUTO_INDENT_PATCH_APPLIED" {
			t.Fatalf("auto-indent pass rule should not appear when disabled: %+v", item)
		}
	}

	afterRaw, err := os.ReadFile(viewPath)
	if err != nil {
		t.Fatalf("read source view after gen failed: %v", err)
	}
	if !strings.Contains(string(afterRaw), "CREATE OR ALTER VIEW") {
		t.Fatalf("view source should keep a canonical CREATE OR ALTER VIEW header")
	}
	if strings.Contains(string(afterRaw), "  SELECT") && !strings.Contains(string(beforeRaw), "  SELECT") {
		t.Fatalf("view body appears auto-indented even though auto-indent is disabled")
	}
}

func copyDirForAutoIndentTest(src, dst string) error {
	return filepath.WalkDir(src, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		rel, err := filepath.Rel(src, path)
		if err != nil {
			return err
		}
		target := filepath.Join(dst, rel)
		if d.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		data, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, data, 0o644)
	})
}
