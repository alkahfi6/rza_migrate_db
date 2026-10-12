package migreport_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"rza_migrate_db/internal/migreport"
)

func TestDesiredStateMarkdownDocumentsDeclaredKeysAndModuleUsage(t *testing.T) {
	tmp := t.TempDir()
	for _, dir := range []string{"migrations", "db_object/tables", "db_object/indexes", "db_object/views", "db_object/procedures"} {
		if err := os.MkdirAll(filepath.Join(tmp, dir), 0755); err != nil {
			t.Fatal(err)
		}
	}
	writeDesiredFixture(t, tmp, "db_object/tables/dbo.Parent.sql", `CREATE TABLE [dbo].[Parent] (
  [Id] INT IDENTITY(1,1) NOT NULL,
  [Code] NVARCHAR(20) NOT NULL,
  CONSTRAINT [PK_Parent] PRIMARY KEY ([Id]),
  CONSTRAINT [UQ_Parent_Code] UNIQUE ([Code])
);`)
	writeDesiredFixture(t, tmp, "db_object/tables/dbo.Child.sql", `CREATE TABLE dbo.Child (
  [Id] INT NOT NULL,
  [ParentId] INT NOT NULL,
  CONSTRAINT [PK_Child] PRIMARY KEY ([Id]),
  CONSTRAINT [FK_Child_Parent] FOREIGN KEY ([ParentId]) REFERENCES [dbo].[Parent] ([Id])
);`)
	writeDesiredFixture(t, tmp, "db_object/tables/dbo.Unrelated.sql", `CREATE TABLE dbo.Unrelated ([ParentId] INT NOT NULL);`)
	writeDesiredFixture(t, tmp, "db_object/tables/dbo.NotIdentity.sql", `CREATE TABLE dbo.NotIdentity ([M_IDENTITY] INT NOT NULL);`)
	writeDesiredFixture(t, tmp, "db_object/tables/dbo.LooseIdentity.sql", `CREATE TABLE dbo.LooseIdentity ([SequenceId] BIGINT IDENTITY(1,1) NOT NULL);`)
	writeDesiredFixture(t, tmp, "db_object/tables/dbo.InlineKey.sql", `CREATE TABLE dbo.InlineKey ([Id] INT IDENTITY(1,1) PRIMARY KEY, [ExternalCode] NVARCHAR(30) UNIQUE);`)
	writeDesiredFixture(t, tmp, "db_object/tables/dbo.GuidValue.sql", `CREATE TABLE dbo.GuidValue ([CorrelationGuid] UNIQUEIDENTIFIER NOT NULL);`)
	writeDesiredFixture(t, tmp, "db_object/tables/dbo.Hierarchy.sql", `CREATE TABLE dbo.Hierarchy ([Id] INT PRIMARY KEY, [ParentId] INT NULL);`)
	writeDesiredFixture(t, tmp, "db_object/tables/dbo.PDFEmail.sql", `CREATE TABLE dbo.PDFEmail ([InputDate] DATETIME NOT NULL, [AttachmentStatement] NVARCHAR(255) NULL);`)
	writeDesiredFixture(t, tmp, "db_object/tables/dbo.MailHistory.sql", `CREATE TABLE dbo.MailHistory ([FileNamePDF] NVARCHAR(255) NULL);`)
	writeDesiredFixture(t, tmp, "db_object/indexes/dbo.Parent.UX_Parent_Code.sql", `CREATE UNIQUE NONCLUSTERED INDEX [UX_Parent_Code] ON [dbo].[Parent] ([Code]);`)
	writeDesiredFixture(t, tmp, "db_object/views/dbo.vParent.sql", `CREATE OR ALTER VIEW dbo.vParent AS SELECT Id FROM dbo.Parent;`)
	writeDesiredFixture(t, tmp, "db_object/procedures/dbo.uspChild.sql", `CREATE OR ALTER PROCEDURE dbo.uspChild AS
SELECT c.Id FROM dbo.Child c JOIN dbo.Parent p ON c.ParentId = p.Id;`)
	writeDesiredFixture(t, tmp, "db_object/procedures/dbo.uspMail.sql", `CREATE OR ALTER PROCEDURE dbo.uspMail AS
SELECT pe.InputDate
FROM dbo.PDFEmail pe
JOIN dbo.MailHistory th
  ON pe.InputDate >= GETDATE()
 AND (LTRIM(pe.AttachmentStatement) = LTRIM(th.FileNamePDF));`)
	writeDesiredFixture(t, tmp, "db_object/procedures/dbo.uspLegacy.sql", `CREATE OR ALTER PROCEDURE dbo.uspLegacy AS
SELECT c.Id FROM dbo.Child c, dbo.Parent p WHERE c.ParentId = p.Id;
EXEC(@dynamic_sql);`)
	writeDesiredFixture(t, tmp, "db_object/procedures/dbo.uspHierarchy.sql", `CREATE OR ALTER PROCEDURE dbo.uspHierarchy AS
SELECT child.Id FROM dbo.Hierarchy child JOIN dbo.Hierarchy parent ON child.ParentId = parent.Id;`)

	out := filepath.Join(tmp, "out", "report")
	code, err := migreport.Run(migreport.Options{RepoRoot: tmp, OutputDir: out})
	if err != nil || code != 0 {
		t.Fatalf("report failed: code=%d err=%v", code, err)
	}
	docsDir := filepath.Join(tmp, "docs", "database")
	keys := readDesiredFixture(t, filepath.Join(docsDir, "identifiers/table_unique_ids.md"))
	if !strings.Contains(keys, "IDENTITY") || !strings.Contains(keys, "PRIMARY_KEY") || !strings.Contains(keys, "UNIQUE_CONSTRAINT") || !strings.Contains(keys, "UNIQUE_INDEX") {
		t.Fatalf("missing unique-id evidence: %s", keys)
	}
	relationships := readDesiredFixture(t, filepath.Join(docsDir, "relationships/relationship_catalog.md"))
	lineage := readDesiredFixture(t, filepath.Join(docsDir, "relationships/module_lineage.md"))
	for _, expected := range []string{"DECLARED_FK", "OBSERVED_JOIN", "dbo.Child.(ParentId)", "dbo.Parent.(Id)"} {
		if !strings.Contains(relationships, expected) {
			t.Fatalf("missing %q in relationship document: %s", expected, relationships)
		}
	}
	for _, expected := range []string{"dbo.vParent", "dbo.uspChild", "dbo.Parent", "dbo.Child"} {
		if !strings.Contains(lineage, expected) {
			t.Fatalf("missing %q in lineage: %s", expected, lineage)
		}
	}
	if !strings.Contains(lineage, "dbo.uspLegacy") || !strings.Contains(lineage, "Dynamic SQL yang harus diperiksa manual") {
		t.Fatalf("legacy lineage/dynamic SQL review is missing: %s", lineage)
	}
	if strings.Contains(relationships, "dbo.Unrelated") {
		t.Fatalf("same column name must not create a relationship: %s", relationships)
	}
	if !strings.Contains(relationships, "dbo.PDFEmail.(AttachmentStatement)") || !strings.Contains(relationships, "dbo.MailHistory.(FileNamePDF)") {
		t.Fatalf("wrapped equality relationship was not detected: %s", relationships)
	}
	if strings.Contains(relationships, "dbo.PDFEmail.(InputDate) ↔ dbo.PDFEmail") || strings.Contains(relationships, "dbo.PDFEmail.(InputDate) ↔ dbo.MailHistory") {
		t.Fatalf("the '=' token inside '>=' must not create a relationship: %s", relationships)
	}
	if strings.Contains(keys, "dbo.NotIdentity` | `M_IDENTITY` | IDENTITY") {
		t.Fatalf("identity-like column name must not become identity evidence: %s", keys)
	}
	if !strings.Contains(keys, "Tanpa key/index unique, duplikat tetap mungkin terjadi") {
		t.Fatalf("missing identity risk explanation: %s", keys)
	}
	if !strings.Contains(keys, "dbo.InlineKey` | `Id` | `PRIMARY_KEY` | `ENFORCED`") || !strings.Contains(keys, "dbo.InlineKey` | `ExternalCode` | `UNIQUE_CONSTRAINT` | `ENFORCED`") {
		t.Fatalf("inline key constraints must be included: %s", keys)
	}
	if strings.Contains(keys, "Risiko / tindakan") || strings.Contains(keys, "TINGGI:") || strings.Contains(keys, "PERIKSA:") {
		t.Fatalf("template risk guidance must not be repeated in table rows: %s", keys)
	}
	for _, heading := range []string{"Classification", "Uniqueness", "Source", "Join usage", "Classification reference"} {
		if !strings.Contains(keys, heading) {
			t.Fatalf("missing natural technical heading %q: %s", heading, keys)
		}
	}
	if strings.Contains(keys, "dbo.GuidValue` | `CorrelationGuid` | `UNIQUE_CONSTRAINT`") {
		t.Fatalf("UNIQUEIDENTIFIER data type must not become UNIQUE constraint evidence: %s", keys)
	}
	for _, name := range []string{"relationships/relational_erd.md", "relationships/relationship_catalog.md", "relationships/module_lineage.md"} {
		if _, err := os.Stat(filepath.Join(docsDir, name)); err != nil {
			t.Fatalf("missing report %s in docs: %v", name, err)
		}
		if _, err := os.Stat(filepath.Join(out, name)); !os.IsNotExist(err) {
			t.Fatalf("document %s must not exist in out/ directory: %v", name, err)
		}
	}
	procedureIndex := readDesiredFixture(t, filepath.Join(docsDir, "procedures", "README.md"))
	procedureCatalog := readDesiredFixture(t, filepath.Join(docsDir, "procedures", "procedure_impact_catalog.md"))
	procedureProfile := readDesiredFixture(t, filepath.Join(docsDir, "procedures", "profiles", "dbo.uspChild.md"))
	dynamicProfile := readDesiredFixture(t, filepath.Join(docsDir, "procedures", "profiles", "dbo.uspLegacy.md"))
	for _, expected := range []string{"Procedure Object & Process Profiles", "dbo.uspChild", "dbo.uspLegacy"} {
		if !strings.Contains(procedureIndex, expected) {
			t.Fatalf("procedure index missing %q: %s", expected, procedureIndex)
		}
	}
	if !strings.Contains(procedureCatalog, "Procedure Impact Catalog") || !strings.Contains(procedureProfile, "Direct object access") || !strings.Contains(procedureProfile, "dbo.Child") || !strings.Contains(procedureProfile, "dbo.Parent") {
		t.Fatalf("procedure catalog/profile is incomplete: catalog=%s profile=%s", procedureCatalog, procedureProfile)
	}
	if !strings.Contains(dynamicProfile, "UNRESOLVED_DYNAMIC") || !strings.Contains(dynamicProfile, "db_object` digunakan untuk memvalidasi kandidat") {
		t.Fatalf("dynamic SQL profile must stay conservative: %s", dynamicProfile)
	}
	erd := readDesiredFixture(t, filepath.Join(docsDir, "relationships/relational_erd.md"))
	if !strings.Contains(erd, "Relationship Group") || !strings.Contains(erd, "### Self-join review") || strings.Contains(erd, "peran A") || strings.Contains(erd, "peran B") {
		t.Fatalf("ERD must explain groups and separate self-joins from overview diagrams: %s", erd)
	}
	if !strings.Contains(erd, "ParentId = Id [FK]") || !strings.Contains(erd, "ParentId = Id [JOIN]") || strings.Contains(erd, " · usage ") || strings.Contains(erd, "### Hub:") {
		t.Fatalf("ERD must preserve exact column mappings without usage-only labels or artificial Hub ownership: %s", erd)
	}
	if strings.Contains(relationships, "<br>") || strings.Contains(erd, "<br>") {
		t.Fatal("generated Markdown must not contain raw HTML line breaks")
	}
	if !strings.Contains(relationships, "## Group summary") || !strings.Contains(relationships, "## Relationship Group") || !strings.Contains(relationships, "Evidence & confidence reference") {
		t.Fatalf("relationship catalog must be grouped and include a reference guide: %s", relationships)
	}
}

func writeDesiredFixture(t *testing.T, root, relative, value string) {
	t.Helper()
	if err := os.WriteFile(filepath.Join(root, relative), []byte(value), 0644); err != nil {
		t.Fatal(err)
	}
}

func readDesiredFixture(t *testing.T, path string) string {
	t.Helper()
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	return string(data)
}
