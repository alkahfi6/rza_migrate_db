package history

import "testing"

func TestParseMigrationFileNameValid(t *testing.T) {
	t.Parallel()

	opts := ParseOptions{
		CoreAuditUpName:   "000000__core__mig__audit_init.up.sql",
		CoreAuditDownName: "000000__core__mig__audit_init.down.sql",
		AllowedCategories: map[string]struct{}{
			"tables":     {},
			"indexes":    {},
			"views":      {},
			"functions":  {},
			"procedures": {},
			"patch":      {},
		},
		AllowedObjectTokenRunes: allowedObjectRunes(),
	}

	meta, issue := ParseMigrationFileName("20260217073015000_tables_dbo.TableA_v1.up.sql", opts)
	if issue != nil {
		t.Fatalf("unexpected parse issue: %+v", issue)
	}
	if meta.Version != "20260217073015000" || meta.Category != "tables" || meta.Object != "dbo.TableA" || meta.Direction != "up" || meta.VK != 1 {
		t.Fatalf("unexpected parsed meta: %+v", meta)
	}
}

func TestParseMigrationFileNameValidObjectWithHash(t *testing.T) {
	t.Parallel()

	opts := ParseOptions{
		CoreAuditUpName:   "000000__core__mig__audit_init.up.sql",
		CoreAuditDownName: "000000__core__mig__audit_init.down.sql",
		AllowedCategories: map[string]struct{}{
			"tables": {},
		},
		AllowedObjectTokenRunes: allowedObjectRunes(),
	}

	meta, issue := ParseMigrationFileName("20260217073015000_tables_dbo.tmp_#data_param_v1.up.sql", opts)
	if issue != nil {
		t.Fatalf("unexpected parse issue: %+v", issue)
	}
	if meta.Object != "dbo.tmp_#data_param" {
		t.Fatalf("unexpected parsed object: %+v", meta)
	}
}

func TestParseMigrationFileNameRejectsNonDboObject(t *testing.T) {
	t.Parallel()

	opts := ParseOptions{
		CoreAuditUpName:   "000000__core__mig__audit_init.up.sql",
		CoreAuditDownName: "000000__core__mig__audit_init.down.sql",
		AllowedCategories: map[string]struct{}{
			"tables": {},
		},
		AllowedObjectTokenRunes: allowedObjectRunes(),
	}

	_, issue := ParseMigrationFileName("20260217073015000_tables_sales.TableA_v1.up.sql", opts)
	if issue == nil {
		t.Fatalf("expected parse issue for non-dbo migration object")
	}
	if issue.Rule != "MIGRATION_OBJECT_SCHEMA_NOT_DBO" {
		t.Fatalf("unexpected parse rule: %s", issue.Rule)
	}
}

func TestParseMigrationFileNameCorePair(t *testing.T) {
	t.Parallel()

	opts := ParseOptions{
		CoreAuditUpName:   "000000__core__mig__audit_init.up.sql",
		CoreAuditDownName: "000000__core__mig__audit_init.down.sql",
	}

	meta, issue := ParseMigrationFileName("000000__core__mig__audit_init.up.sql", opts)
	if issue != nil {
		t.Fatalf("unexpected parse issue for core up: %+v", issue)
	}
	if !meta.IsCore || meta.Direction != "up" {
		t.Fatalf("unexpected core up parse result: %+v", meta)
	}
}

func TestShouldReplace(t *testing.T) {
	t.Parallel()

	if !ShouldReplace("20260217073015000", 1, "20260217073015001", 1) {
		t.Fatalf("expected newer version to replace")
	}
	if ShouldReplace("20260217073015002", 1, "20260217073015001", 1) {
		t.Fatalf("did not expect older version to replace")
	}
	if !ShouldReplace("20260217073015002", 1, "20260217073015002", 2) {
		t.Fatalf("expected same version higher vk to replace")
	}
}

func allowedObjectRunes() map[rune]struct{} {
	m := make(map[rune]struct{})
	for ch := 'a'; ch <= 'z'; ch++ {
		m[ch] = struct{}{}
	}
	for ch := 'A'; ch <= 'Z'; ch++ {
		m[ch] = struct{}{}
	}
	for ch := '0'; ch <= '9'; ch++ {
		m[ch] = struct{}{}
	}
	m['_'] = struct{}{}
	m['#'] = struct{}{}
	m['.'] = struct{}{}
	return m
}
