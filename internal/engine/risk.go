package engine

import (
	"fmt"
	"strings"
)

type migrationRisk string

const (
	riskSafe    migrationRisk = "safe"
	riskGuarded migrationRisk = "guarded_auto"
	riskManual  migrationRisk = "manual_required"
)

type riskAssessment struct {
	Level   migrationRisk
	Reasons []string
}

// assessTableTransition rejects only transitions for which the current engine
// cannot prove a lossless and deterministic automatic migration. The decision
// is derived from the two authoritative states, never stored separately.
func assessTableTransition(beforeSQL, afterSQL, object string) (riskAssessment, error) {
	before, err := BuildTableModel(object, beforeSQL)
	if err != nil {
		return riskAssessment{}, err
	}
	after, err := BuildTableModel(object, afterSQL)
	if err != nil {
		return riskAssessment{}, err
	}
	diff := DiffTableModels(before, after)
	manualReasons := make([]string, 0)
	guardedReasons := make([]string, 0)
	if len(diff.Removed) > 0 {
		for _, col := range diff.Removed {
			guardedReasons = append(guardedReasons, "column removal runs only when the column contains no non-NULL data: "+col.Name)
		}
	}
	for _, change := range diff.Changed {
		from, to := change.Before, change.After
		if from.Identity != to.Identity || from.RowGuidCol != to.RowGuidCol {
			manualReasons = append(manualReasons, "IDENTITY/ROWGUIDCOL modifier cannot be altered in place: "+to.Name)
			continue
		}
		if from.Computed || to.Computed {
			if from.Computed && to.Computed && isNonPersistedComputed(from) && isNonPersistedComputed(to) {
				guardedReasons = append(guardedReasons, "non-persisted computed expression will be rebuilt; dependency checks must pass: "+to.Name)
			} else {
				manualReasons = append(manualReasons, "computed-column transition requires explicit manual migration: "+to.Name)
			}
			continue
		}
		if from.Nullable && !to.Nullable {
			guardedReasons = append(guardedReasons, "NULL to NOT NULL will run only when no existing NULL values: "+to.Name)
		}
		if isLossyTypeTransition(from.Type, to.Type) {
			if isGuardableTypeTransition(from.Type, to.Type) {
				guardedReasons = append(guardedReasons, fmt.Sprintf("lossless conversion precheck required: %s (%s -> %s)", to.Name, from.Type.Declaration, to.Type.Declaration))
			} else {
				manualReasons = append(manualReasons, fmt.Sprintf("lossy or unproven type transition requires explicit manual migration: %s (%s -> %s)", to.Name, from.Type.Declaration, to.Type.Declaration))
			}
		}
	}
	if len(manualReasons) > 0 {
		return riskAssessment{Level: riskManual, Reasons: manualReasons}, nil
	}
	if len(guardedReasons) > 0 {
		return riskAssessment{Level: riskGuarded, Reasons: guardedReasons}, nil
	}
	if len(manualReasons) == 0 {
		return riskAssessment{Level: riskSafe}, nil
	}
	return riskAssessment{Level: riskManual, Reasons: manualReasons}, nil
}

func isNonPersistedComputed(column ColumnModel) bool {
	return column.Computed && !strings.Contains(strings.ToUpper(column.Expression), " PERSISTED")
}

func isGuardableTypeTransition(before, after SQLType) bool {
	beforeBase := normalizeColumnBaseType(before.Base)
	afterBase := normalizeColumnBaseType(after.Base)
	if isGuardableXMLNVarCharMaxTransition(before, after) {
		return true
	}
	if isStringType(before.Base) && isStringType(after.Base) {
		if !strings.EqualFold(sqlTypeCollation(before), sqlTypeCollation(after)) {
			return false
		}
		return beforeBase == afterBase && before.Unicode == after.Unicode
	}
	if isBinaryType(before.Base) && isBinaryType(after.Base) {
		return beforeBase == afterBase
	}
	if isStringType(before.Base) && isLosslessGuardableNumericType(after.Base) {
		return true
	}
	if isLosslessGuardableNumericType(before.Base) && isStringType(after.Base) {
		return true
	}
	if isLosslessGuardableNumericType(before.Base) && isLosslessGuardableNumericType(after.Base) {
		return true
	}
	return isDateTimeType(before.Base) && isDateTimeType(after.Base)
}

// XML values can be serialized losslessly to NVARCHAR(MAX). The reverse
// direction remains automatic only with a runtime well-formedness precheck,
// because new text written after the up migration may not be valid XML.
func isGuardableXMLNVarCharMaxTransition(before, after SQLType) bool {
	beforeBase := normalizeColumnBaseType(before.Base)
	afterBase := normalizeColumnBaseType(after.Base)
	return beforeBase == "XML" && isUntypedXMLSQLType(before) && afterBase == "NVARCHAR" && after.Length == -1 ||
		beforeBase == "NVARCHAR" && before.Length == -1 && afterBase == "XML" && isUntypedXMLSQLType(after)
}

func isUntypedXMLSQLType(sqlType SQLType) bool {
	return strings.EqualFold(strings.TrimSpace(sqlType.Declaration), "XML")
}

func sqlTypeCollation(sqlType SQLType) string {
	upper := strings.ToUpper(strings.TrimSpace(sqlType.Declaration))
	marker := " COLLATE "
	pos := strings.Index(upper, marker)
	if pos < 0 {
		return ""
	}
	rest := strings.TrimSpace(upper[pos+len(marker):])
	if fields := strings.Fields(rest); len(fields) > 0 {
		return fields[0]
	}
	return ""
}

func isLosslessGuardableNumericType(base string) bool {
	switch normalizeColumnBaseType(base) {
	case "BIT", "TINYINT", "SMALLINT", "INT", "BIGINT", "DECIMAL", "NUMERIC":
		return true
	default:
		// FLOAT/REAL are deliberately excluded: SQL Server binary floating point
		// equality is not a reliable lossless round-trip predicate. MONEY types
		// are also excluded until an exact scale-aware model is implemented.
		return false
	}
}

func isLossyTypeTransition(before, after SQLType) bool {
	beforeBase := normalizeColumnBaseType(before.Base)
	afterBase := normalizeColumnBaseType(after.Base)
	if beforeBase != afterBase {
		return true
	}
	if isStringType(before.Base) {
		if before.Unicode != after.Unicode {
			return true
		}
		if !strings.EqualFold(sqlTypeCollation(before), sqlTypeCollation(after)) {
			return true
		}
		return before.Length == -1 && after.Length != -1 || before.Length > after.Length && after.Length >= 0
	}
	// Decimal/numeric, datetime/time and binary declarations retain important
	// precision/scale semantics in Declaration; automatic alteration is denied.
	return before.Declaration != after.Declaration && (strings.Contains(beforeBase, "DECIMAL") || strings.Contains(beforeBase, "NUMERIC") || isDateTimeType(beforeBase) || strings.Contains(beforeBase, "BINARY"))
}
