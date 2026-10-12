package engine

import (
	"fmt"
	"strings"

	"rza_migrate_db/internal/sqlparse"
)

// renderPlannedMigration is the rendering boundary: planning owns semantic SQL
// and this layer owns wrappers, metadata, and final file bytes.
func renderPlannedMigration(op plannedMigration, version, fileName string, isUp bool) string {
	body := op.UpSemanticSQL
	if !isUp {
		body = op.DownSemanticSQL
		if strings.TrimSpace(body) == "" {
			body = "-- WARNING: DATA_RESTORE_NOT_GUARANTEED\n-- Manual restore required.\n"
		} else if op.HasDestructiveDown {
			body = "-- WARNING: DATA_RESTORE_NOT_GUARANTEED\n-- Destructive change may not be fully restorable.\n\n" + body
		}
	}
	switch op.Category {
	case "types", "tables", "indexes":
		return renderTxWrapper(version, op.Category, fileName, op.Object, isUp, body)
	case "patch":
		header := fmt.Sprintf("/*\npatch_folder: %s\npatch_name: %s\nreadme: %s\npatch_up_sha256: %s\npatch_down_sha256: %s\n*/\n", op.PatchFolderOriginal, op.PatchNameCanonical, op.PatchReadmePath, op.PatchUpHash, op.PatchDownHash)
		return normalizeSQLText(header+renderTxWrapper(version, op.Category, fileName, op.Object, isUp, body), true, true)
	case "views", "functions", "procedures":
		return renderTxWrapper(version, op.Category, fileName, op.Object, isUp, renderProgrammabilityExecutableBody(body))
	default:
		return normalizeSQLText(body, true, true)
	}
}

// SQL Server requires CREATE/ALTER module statements to be the first statement
// in their batch. Executing that statement as a dedicated dynamic batch keeps
// the outer transaction, mandatory audit, checksum, and error wrapper intact.
func renderProgrammabilityExecutableBody(body string) string {
	normalized := normalizeSQLText(body, true, true)
	// Banner comments may legally precede a module definition. Detect the first
	// executable tokens rather than requiring CREATE OR ALTER at byte zero;
	// SQL Server still requires the module statement to run in its own batch.
	upper := strings.ToUpper(strings.TrimSpace(sqlparse.StripSQLStringsAndComments(normalized)))
	upper = strings.Join(strings.Fields(upper), " ")
	if strings.HasPrefix(upper, "CREATE OR ALTER VIEW ") ||
		strings.HasPrefix(upper, "CREATE OR ALTER FUNCTION ") ||
		strings.HasPrefix(upper, "CREATE OR ALTER PROCEDURE ") {
		return "EXEC sys.sp_executesql N'" + strings.ReplaceAll(strings.TrimSpace(normalized), "'", "''") + "';\n"
	}
	return normalized
}
