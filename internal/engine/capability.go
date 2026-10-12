package engine

import (
	"fmt"
	"strings"

	"rza_migrate_db/internal/sqlparse"
)

// CapabilityIssue means the source is valid SQL Server SQL but outside the
// deliberately supported auto-generation subset. It is fail-closed by design.
type CapabilityIssue struct {
	Feature string
	Message string
}

func (i CapabilityIssue) Error() string { return i.Message }

// assessSourceCapability is intentionally conservative. Expanding this list
// requires a typed model, render support, round-trip test, and SQL Server test.
func assessSourceCapability(category, sql string) *CapabilityIssue {
	semantic := strings.ToUpper(maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(sql)))
	unsupported := map[string][]string{
		"tables":  {"FILESTREAM", "COLUMN_SET", "MASKED WITH", "ENCRYPTED WITH", "PERIOD FOR SYSTEM_TIME", "SYSTEM_VERSIONING", "GENERATED ALWAYS", "MEMORY_OPTIMIZED", "LEDGER", " AS NODE", " AS EDGE", "FILESTREAM_ON", "REMOTE_DATA_ARCHIVE"},
		"indexes": {"RESUMABLE =", "ON PARTITIONS", "DROP_EXISTING = ON", "CREATE FULLTEXT INDEX", "CREATE SELECTIVE XML INDEX", "CREATE JSON INDEX", "BUCKET_COUNT"},
	}
	for _, token := range unsupported[category] {
		if strings.Contains(semantic, token) {
			return &CapabilityIssue{Feature: strings.TrimSpace(token), Message: fmt.Sprintf("unsupported SQL Server feature for fail-closed auto generation: %s", strings.TrimSpace(token))}
		}
	}
	return nil
}
