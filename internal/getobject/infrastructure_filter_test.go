package getobject

import "testing"

func TestExtractorInfrastructureTableFilterIsExactCase(t *testing.T) {
	t.Parallel()

	for _, name := range []string{
		"schema_migrations",
		"mig_event",
		"mig_operation_journal",
	} {
		if !isExtractorInfrastructureTable("dbo", name) {
			t.Fatalf("mandatory infrastructure table must be filtered: dbo.%s", name)
		}
	}
	for _, test := range []struct {
		schema string
		name   string
	}{
		{schema: "dbo", name: "Schema_Migrations"},
		{schema: "dbo", name: "Mig_Event"},
		{schema: "dbo", name: "Mig_Operation_Journal"},
		{schema: "audit", name: "schema_migrations"},
		{schema: "audit", name: "mig_event"},
		{schema: "dbo", name: "schema_migrations_archive"},
		{schema: "dbo", name: "mig_event_archive"},
	} {
		if isExtractorInfrastructureTable(test.schema, test.name) {
			t.Fatalf("case-distinct or user-owned table was filtered: %s.%s", test.schema, test.name)
		}
	}
}

func TestExtractorInfrastructureObjectFilter(t *testing.T) {
	t.Parallel()

	for _, test := range []struct {
		category string
		name     string
	}{
		{category: "tables", name: "mig_event"},
		{category: "indexes", name: "mig_event"},
		{category: "tables", name: "mig_operation_journal"},
		{category: "indexes", name: "mig_operation_journal"},
		{category: "procedures", name: "usp_mig_log_event"},
		{category: "procedures", name: "usp_mig_try_get_lock"},
	} {
		if !isExtractorInfrastructureObject(test.category, "dbo", test.name) {
			t.Fatalf("mandatory infrastructure object must be filtered: %s dbo.%s", test.category, test.name)
		}
	}

	for _, test := range []struct {
		category string
		schema   string
		name     string
	}{
		{category: "views", schema: "dbo", name: "mig_event"},
		{category: "procedures", schema: "dbo", name: "USP_MIG_LOG_EVENT"},
		{category: "procedures", schema: "audit", name: "usp_mig_log_event"},
		{category: "procedures", schema: "dbo", name: "usp_mig_log_event_archive"},
		{category: "tables", schema: "dbo", name: "Mig_Event"},
	} {
		if isExtractorInfrastructureObject(test.category, test.schema, test.name) {
			t.Fatalf("case-distinct or user-owned object was filtered: %s %s.%s", test.category, test.schema, test.name)
		}
	}
}
