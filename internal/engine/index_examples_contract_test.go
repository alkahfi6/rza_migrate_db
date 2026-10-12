package engine

import (
	"path/filepath"
	"runtime"
	"testing"
)

func TestIndexSourceExamplesVariantSupport(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name          string
		sql           string
		expectedCount int
	}{
		{
			name: "include_and_identity",
			sql: `CREATE INDEX IX_TableA_Status ON dbo.TableA (status) INCLUDE (amount);
CREATE INDEX IX_TableA_StatusCopy ON dbo.TableA (status);
CREATE UNIQUE INDEX IX_TableA_Id ON dbo.TableA (id);`,
			expectedCount: 3,
		},
		{
			name: "clustered_keyword",
			sql: `CREATE CLUSTERED INDEX IX_TableA_Id ON dbo.TableA (id);
CREATE INDEX IX_TableA_Status ON dbo.TableA (status);`,
			expectedCount: 2,
		},
		{
			name: "nonclustered_keyword",
			sql: `CREATE NONCLUSTERED INDEX IX_TableA_Status ON dbo.TableA (status);
CREATE INDEX IX_TableA_Amount ON dbo.TableA (amount);`,
			expectedCount: 2,
		},
		{
			name: "filtered_where",
			sql: `CREATE INDEX IX_TableA_StatusActive ON dbo.TableA (status) WHERE status = 'ACTIVE';
CREATE INDEX IX_TableA_Amount ON dbo.TableA (amount);`,
			expectedCount: 2,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			const objectName = "dbo.TableA"
			stmts, err := extractCreateIndexStatements(objectName, tc.sql)
			if err != nil {
				t.Fatalf("extractCreateIndexStatements failed for %s: %v", tc.name, err)
			}
			if len(stmts) < tc.expectedCount {
				t.Fatalf("expected at least %d statements for %s, got %d", tc.expectedCount, tc.name, len(stmts))
			}

			meta := migrationMeta{
				Category: "indexes",
				Object:   objectName,
			}
			if _, err := semanticFromSource(meta, tc.sql); err != nil {
				t.Fatalf("semanticFromSource failed for %s: %v", tc.name, err)
			}
		})
	}
}

func repoRootFromEnginePackageDir(t *testing.T) string {
	t.Helper()
	_, file, _, ok := runtime.Caller(0)
	if !ok {
		t.Fatalf("runtime caller failed")
	}
	return filepath.Clean(filepath.Join(filepath.Dir(file), "..", ".."))
}
