package migreport

import (
	"encoding/json"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func layoutRepository(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	for path, text := range map[string]string{
		"migrations/20261012000000001_tables_dbo.Parent_v1.up.sql":   "CREATE TABLE dbo.Parent (Id INT PRIMARY KEY);",
		"migrations/20261012000000001_tables_dbo.Parent_v1.down.sql": "DROP TABLE dbo.Parent;",
		"db_object/tables/dbo.Parent.sql":                            "CREATE TABLE dbo.Parent (Id INT PRIMARY KEY);",
		"db_object/procedures/dbo.Read#Parent.sql":                   "CREATE PROCEDURE [dbo].[Read#Parent] AS SELECT Id FROM dbo.Parent;",
	} {
		writeLayoutFile(t, filepath.Join(root, filepath.FromSlash(path)), []byte(text))
	}
	return root
}

func writeLayoutFile(t *testing.T, path string, content []byte) {
	t.Helper()
	if err := os.MkdirAll(filepath.Dir(path), 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, content, 0644); err != nil {
		t.Fatal(err)
	}
}

func runLayout(t *testing.T, opts Options) {
	t.Helper()
	if code, err := Run(opts); code != 0 || err != nil {
		t.Fatalf("report: code=%d error=%v", code, err)
	}
}

func layoutFiles(t *testing.T, root string) map[string]string {
	t.Helper()
	files := map[string]string{}
	if err := filepath.WalkDir(root, func(path string, entry os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if !entry.IsDir() {
			data, err := os.ReadFile(path)
			if err != nil {
				return err
			}
			files[path] = contentHash(data)
		}
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return files
}

func TestUnifiedDocumentationLayoutManifestAndLinks(t *testing.T) {
	for _, mode := range []string{"default", "custom-docs", "custom-roots"} {
		t.Run(mode, func(t *testing.T) {
			root := layoutRepository(t)
			before := layoutFiles(t, root)
			opts := Options{RepoRoot: root}
			docsDir := filepath.Join(root, "docs", "database")
			if mode != "default" {
				opts.DocsDir = "database docs #1"
				if mode == "custom-roots" {
					opts.OutputDir = "audit #1"
				}
				docsDir = filepath.Join(root, opts.DocsDir)
			}
			runLayout(t, opts)
			manifest, err := readLayoutManifest(filepath.Join(docsDir, "manifest.json"))
			if err != nil || manifest == nil {
				t.Fatalf("manifest: %+v %v", manifest, err)
			}
			if manifest.Schema != layoutSchema || len(manifest.Artifacts) != 18 {
				t.Fatalf("unexpected inventory: %+v", manifest)
			}
			linkPattern := regexp.MustCompile(`\]\(([^)]+)\)`)
			for _, a := range manifest.Artifacts {
				path := filepath.Join(manifest.Roots[a.Root], filepath.FromSlash(a.Path))
				data, err := os.ReadFile(path)
				if err != nil {
					t.Fatal(err)
				}
				if a.Path != "manifest.json" && (a.SHA256 == "" || contentHash(data) != a.SHA256) {
					t.Fatalf("bad artifact digest: %+v", a)
				}
				if strings.HasSuffix(path, ".md") {
					for _, match := range linkPattern.FindAllStringSubmatch(string(data), -1) {
						link, err := url.Parse(match[1])
						if err != nil {
							t.Fatal(err)
						}
						if link.Scheme != "" || link.Path == "" {
							continue
						}
						if _, err := os.Stat(filepath.Join(filepath.Dir(path), filepath.FromSlash(link.Path))); err != nil {
							t.Fatalf("broken link in %s: %s: %v", path, match[1], err)
						}
					}
				}
			}
			for path, hash := range before {
				data, err := os.ReadFile(path)
				if err != nil || contentHash(data) != hash {
					t.Fatalf("authored source changed: %s", path)
				}
			}
			if _, err := os.Stat(filepath.Join(root, "out")); !os.IsNotExist(err) {
				t.Fatalf("report created out/: %v", err)
			}
			if mode != "default" {
				if _, err := os.Stat(filepath.Join(root, "docs")); !os.IsNotExist(err) {
					t.Fatalf("custom -docs leaked default docs/: %v", err)
				}
			}
			// Exclude documented time metadata before comparing rerun bytes.
			first := layoutFiles(t, docsDir)
			runLayout(t, opts)
			second := layoutFiles(t, docsDir)
			for path, hash := range first {
				if strings.HasSuffix(path, "manifest.json") || strings.HasSuffix(path, "report.json") || strings.HasSuffix(path, "report_summary.csv") {
					continue
				}
				if second[path] != hash {
					t.Fatalf("unstable document/dataset: %s", path)
				}
			}
		})
	}
}

func TestOutputFileSymlinkAndAliasedRootsCannotEscapePreflight(t *testing.T) {
	for _, mode := range []string{"file", "root-alias"} {
		t.Run(mode, func(t *testing.T) {
			root := layoutRepository(t)
			docs := filepath.Join(root, "docs", "database")
			if err := os.MkdirAll(docs, 0755); err != nil {
				t.Fatal(err)
			}
			opts := Options{RepoRoot: root}
			manual := filepath.Join(root, "manual.md")
			writeLayoutFile(t, manual, []byte("manual"))
			link, target := filepath.Join(docs, "README.md"), manual
			if mode == "root-alias" {
				link, target = filepath.Join(root, "audit-alias"), docs
				opts.OutputDir = link
			}
			if err := os.Symlink(target, link); err != nil {
				t.Skipf("symlink unavailable: %v", err)
			}
			if code, err := Run(opts); code != 1 || err == nil {
				t.Fatalf("symlink/alias preflight must reject: code=%d err=%v", code, err)
			}
			if data, err := os.ReadFile(manual); err != nil || string(data) != "manual" {
				t.Fatal("file symlink changed manual content")
			}
			if _, err := os.Stat(filepath.Join(docs, "manifest.json")); !os.IsNotExist(err) {
				t.Fatal("symlink/alias rejection published partial output")
			}
		})
	}
}

func TestReportCleanupPreservesManualFilesAndModifiedStaleProfiles(t *testing.T) {
	for _, modified := range []bool{false, true} {
		t.Run(map[bool]string{false: "owned-stale", true: "modified-stale"}[modified], func(t *testing.T) {
			root := layoutRepository(t)
			docs := filepath.Join(root, "docs", "database")
			runLayout(t, Options{RepoRoot: root})
			profile := filepath.Join(docs, "procedures", "profiles", "dbo.Read#Parent.md")
			manual := filepath.Join(docs, "procedures", "profiles", "manual.md")
			writeLayoutFile(t, manual, []byte("manual notes"))
			if modified {
				f, err := os.OpenFile(profile, os.O_APPEND|os.O_WRONLY, 0644)
				if err != nil {
					t.Fatal(err)
				}
				_, _ = f.WriteString("\nManual annotation\n")
				_ = f.Close()
			}
			if err := os.Remove(filepath.Join(root, "db_object", "procedures", "dbo.Read#Parent.sql")); err != nil {
				t.Fatal(err)
			}
			runLayout(t, Options{RepoRoot: root})
			_, err := os.Stat(profile)
			if modified && err != nil || !modified && !os.IsNotExist(err) {
				t.Fatalf("stale cleanup modified=%v: %v", modified, err)
			}
			data, err := os.ReadFile(manual)
			if err != nil || string(data) != "manual notes" {
				t.Fatal("manual profile removed or modified")
			}
		})
	}
}

func TestManualCollisionAndEditedCurrentOutputBlockBeforeWriting(t *testing.T) {
	for _, owned := range []bool{false, true} {
		root := layoutRepository(t)
		docs := filepath.Join(root, "docs", "database")
		if owned {
			runLayout(t, Options{RepoRoot: root})
		}
		writeLayoutFile(t, filepath.Join(docs, "README.md"), []byte("manual entry point"))
		before := layoutFiles(t, root)
		code, err := Run(Options{RepoRoot: root})
		if code != 1 || err == nil || !strings.Contains(err.Error(), "preserving") {
			t.Fatalf("manual collision must block: code=%d err=%v", code, err)
		}
		after := layoutFiles(t, root)
		if len(before) != len(after) {
			t.Fatal("partial output after collision")
		}
		for path, hash := range before {
			if after[path] != hash {
				t.Fatalf("file changed despite collision: %s", path)
			}
		}
	}
}

func TestLegacyLayoutRemovesOnlyEquivalentGeneratedFiles(t *testing.T) {
	root := layoutRepository(t)
	docs := filepath.Join(root, "docs", "database")
	runLayout(t, Options{RepoRoot: root})
	current := filepath.Join(docs, "identifiers", "table_unique_ids.md")
	data, _ := os.ReadFile(current)
	writeLayoutFile(t, filepath.Join(docs, "table_unique_ids.md"), legacyMarkdownBody(data))
	writeLayoutFile(t, filepath.Join(docs, "module_lineage.md"), []byte("manual legacy notes"))
	writeLayoutFile(t, filepath.Join(docs, "table_relationship_tree.md"), []byte(renderRelationshipCompatibilityNotice()))
	legacy := filepath.Join(root, "out", "report")
	for _, path := range auditArtifactFiles() {
		if path == "README.md" {
			continue
		}
		data, _ := os.ReadFile(filepath.Join(docs, "migrations", filepath.FromSlash(path)))
		if strings.HasSuffix(path, "report_summary.csv") {
			continue // The timestamp differs on a later run and must be preserved.
		}
		writeLayoutFile(t, filepath.Join(legacy, filepath.Base(path)), data)
	}
	runLayout(t, Options{RepoRoot: root})
	for _, path := range []string{filepath.Join(docs, "table_unique_ids.md"), filepath.Join(docs, "table_relationship_tree.md"), legacy} {
		if _, err := os.Stat(path); !os.IsNotExist(err) {
			t.Fatalf("equivalent legacy output not cleaned: %s: %v", path, err)
		}
	}
	if data, err := os.ReadFile(filepath.Join(docs, "module_lineage.md")); err != nil || string(data) != "manual legacy notes" {
		t.Fatal("manual legacy document changed")
	}
}

func TestOutputOverlapAndNestedSymlinkFailWithoutSourceWrites(t *testing.T) {
	root := layoutRepository(t)
	docs := filepath.Join(root, "docs", "database")
	if code, err := Run(Options{RepoRoot: root, OutputDir: docs}); code != 1 || err == nil {
		t.Fatalf("overlap must fail: code=%d err=%v", code, err)
	}
	if _, err := os.Stat(docs); !os.IsNotExist(err) {
		t.Fatal("overlap created output")
	}
	if err := os.MkdirAll(docs, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(filepath.Join(root, "migrations"), filepath.Join(docs, "relationships")); err != nil {
		t.Skipf("symlink unavailable: %v", err)
	}
	before := layoutFiles(t, filepath.Join(root, "migrations"))
	if code, err := Run(Options{RepoRoot: root}); code != 1 || err == nil {
		t.Fatalf("nested symlink must fail: code=%d err=%v", code, err)
	}
	if _, err := os.Stat(filepath.Join(docs, "README.md")); !os.IsNotExist(err) {
		t.Fatal("nested symlink failure published partial output")
	}
	after := layoutFiles(t, filepath.Join(root, "migrations"))
	for path, hash := range before {
		if after[path] != hash {
			t.Fatal("symlink changed migration history")
		}
	}
}

func TestPreviousManifestCannotDeleteAnUnlistedManualPath(t *testing.T) {
	root := layoutRepository(t)
	docs := filepath.Join(root, "docs", "database")
	runLayout(t, Options{RepoRoot: root})
	manual := filepath.Join(docs, "manual.md")
	writeLayoutFile(t, manual, []byte("manual"))
	m, _ := readLayoutManifest(filepath.Join(docs, "manifest.json"))
	m.Artifacts = append(m.Artifacts, ReportArtifact{Root: "documentation_directory", Path: "manual.md", SHA256: contentHash([]byte("manual"))})
	data, _ := json.Marshal(m)
	writeLayoutFile(t, filepath.Join(docs, "manifest.json"), data)
	runLayout(t, Options{RepoRoot: root})
	if data, err := os.ReadFile(manual); err != nil || string(data) != "manual" {
		t.Fatal("unlisted manual filename deleted through manifest")
	}
}

func TestDocumentationOwnershipSurvivesRepositoryClone(t *testing.T) {
	for _, custom := range []bool{false, true} {
		t.Run(map[bool]string{false: "default", true: "retained-audit-root"}[custom], func(t *testing.T) {
			original := layoutRepository(t)
			opts := Options{RepoRoot: original}
			if custom {
				opts.DocsDir, opts.OutputDir = "database docs", "audit"
			}
			runLayout(t, opts)
			originalFiles := layoutFiles(t, original)
			clone := t.TempDir()
			for path := range originalFiles {
				relative, err := filepath.Rel(original, path)
				if err != nil {
					t.Fatal(err)
				}
				data, err := os.ReadFile(path)
				if err != nil {
					t.Fatal(err)
				}
				writeLayoutFile(t, filepath.Join(clone, relative), data)
			}
			writeLayoutFile(t, filepath.Join(clone, "db_object", "tables", "dbo.Parent.sql"), []byte("CREATE TABLE dbo.Parent (Id INT PRIMARY KEY, Code INT UNIQUE);"))
			opts.RepoRoot = clone
			if custom {
				opts.DocsDir = "renamed docs"
			}
			runLayout(t, opts)
			docs := filepath.Join(clone, "docs", "database")
			if custom {
				docs = filepath.Join(clone, opts.DocsDir)
			}
			ids, err := os.ReadFile(filepath.Join(docs, "identifiers", "table_unique_ids.md"))
			if err != nil || !strings.Contains(string(ids), "UNIQUE_CONSTRAINT") {
				t.Fatalf("clone could not regenerate source delta: %v", err)
			}
			for path, hash := range originalFiles {
				data, err := os.ReadFile(path)
				if err != nil || contentHash(data) != hash {
					t.Fatalf("clone modified the original repository: %s", path)
				}
			}
		})
	}
}
