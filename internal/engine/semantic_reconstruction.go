package engine

import (
	"fmt"
	"strings"
)

// semanticFromMigration reconstructs the in-memory historical state from the
// only historical authority: migration SQL. No reconstructed state is kept as
// user-maintained metadata.
func semanticFromMigration(meta migrationMeta, sql string) (string, error) {
	return semanticFromMigrationWithState(meta, sql, "", false)
}

func semanticFromMigrationWithState(meta migrationMeta, sql, currentSemantic string, hasCurrent bool) (string, error) {
	if migrationDropsObject(meta.Category, sql) {
		return "", nil
	}

	switch meta.Category {
	case "types":
		stmt, err := extractCreateTypeStatement(meta.Object, sql)
		if err != nil {
			return "", err
		}
		return normalizeSQLText(stmt, true, true), nil
	case "tables":
		stmt, err := extractCreateTableStatement(meta.Object, sql)
		if err == nil {
			if err := validateCreateTableStructure(stmt, meta.Object); err != nil {
				return "", err
			}
			if err := validateTableSubsetSupport(stmt); err != nil {
				return "", err
			}
			return normalizeSQLText(stmt, true, true), nil
		}
		if isManualDropTableRequestMigration(sql) {
			if !hasCurrent || strings.TrimSpace(currentSemantic) == "" {
				return "", fmt.Errorf("manual drop-table request requires prior table state for %s", meta.Object)
			}
			return normalizeSQLText(currentSemantic, true, true), nil
		}
		working, applied := currentSemantic, false
		renames := extractColumnRenames(sql)
		if len(renames) > 0 {
			if !hasCurrent || strings.TrimSpace(currentSemantic) == "" {
				return "", fmt.Errorf("column rename migration requires prior table state for %s", meta.Object)
			}
			updated, renameErr := applyColumnRenamesToCreateTable(working, renames)
			if renameErr != nil {
				return "", renameErr
			}
			working, applied = updated, true
		}
		removedCols, dropErr := extractDroppedColumns(meta.Object, sql)
		if dropErr == nil && len(removedCols) > 0 {
			if !hasCurrent || strings.TrimSpace(currentSemantic) == "" {
				return "", fmt.Errorf("drop-column migration requires prior table state for %s", meta.Object)
			}
			updated, applyErr := applyDropColumnsToCreateTable(working, removedCols)
			if applyErr != nil {
				return "", applyErr
			}
			working, applied = updated, true
		}
		addedCols, alteredCols, alterErr := extractTableAlterColumnDefs(meta.Object, sql)
		if alterErr == nil && (len(addedCols) > 0 || len(alteredCols) > 0) {
			if !hasCurrent || strings.TrimSpace(currentSemantic) == "" {
				return "", fmt.Errorf("alter-column migration requires prior table state for %s", meta.Object)
			}
			updated, applyErr := applyAddAlterColumnsToCreateTable(working, addedCols, alteredCols)
			if applyErr != nil {
				return "", applyErr
			}
			working, applied = updated, true
		}
		droppedDefaults, addedDefaults := extractTableDefaultConstraintTransitions(meta.Object, sql)
		if len(droppedDefaults) > 0 || len(addedDefaults) > 0 {
			if !hasCurrent || strings.TrimSpace(currentSemantic) == "" {
				return "", fmt.Errorf("default-constraint migration requires prior table state for %s", meta.Object)
			}
			for _, name := range droppedDefaults {
				updated, removeErr := removeNamedConstraintFromCreateTable(working, name)
				if removeErr != nil {
					return "", removeErr
				}
				working = updated
			}
			if len(addedDefaults) > 0 {
				updated, addErr := applyAddDefaultConstraintsToCreateTable(working, addedDefaults)
				if addErr != nil {
					return "", addErr
				}
				working = updated
			}
			applied = true
		}
		droppedChecks, addedChecks := extractTableCheckConstraintTransitions(meta.Object, sql)
		if len(droppedChecks) > 0 || len(addedChecks) > 0 {
			if !hasCurrent || strings.TrimSpace(currentSemantic) == "" {
				return "", fmt.Errorf("check-constraint migration requires prior table state for %s", meta.Object)
			}
			for _, name := range droppedChecks {
				updated, removeErr := removeNamedConstraintFromCreateTable(working, name)
				if removeErr != nil {
					return "", removeErr
				}
				working = updated
			}
			for _, def := range addedChecks {
				updated, addErr := addOrReplaceNamedConstraintInCreateTable(working, def.Name, def.Definition)
				if addErr != nil {
					return "", addErr
				}
				working = updated
			}
			applied = true
		}
		// Dependency-safe ALTER COLUMN plans can temporarily drop and restore
		// named UNIQUE/PRIMARY KEY/FOREIGN KEY constraints. The specialized
		// DEFAULT and CHECK reconstruction above predates those plans, so apply
		// the complete local named-constraint transition as the final semantic
		// view. Cross-table FK transitions are applied by the history-state layer.
		localConstraintTransition := false
		constraintDrops := extractDroppedConstraintsFromMigrationSQL(sql)
		constraintAdds := extractAddedConstraintsFromMigrationSQL(sql)
		for _, batch := range extractSpExecuteSQLBatches(sql) {
			constraintDrops = append(constraintDrops, extractDroppedConstraintsFromMigrationSQL(batch)...)
			constraintAdds = append(constraintAdds, extractAddedConstraintsFromMigrationSQL(batch)...)
		}
		for _, drop := range constraintDrops {
			if !sameQualifiedObject(drop.Table, meta.Object) {
				continue
			}
			updated, removeErr := removeNamedConstraintFromCreateTable(working, drop.Name)
			if removeErr != nil {
				return "", removeErr
			}
			working = updated
			localConstraintTransition = true
		}
		for _, add := range constraintAdds {
			if !sameQualifiedObject(add.Table, meta.Object) {
				continue
			}
			updated, addErr := addOrReplaceNamedConstraintInCreateTable(working, add.Name, add.Definition)
			if addErr != nil {
				return "", addErr
			}
			working = updated
			localConstraintTransition = true
		}
		if localConstraintTransition {
			applied = true
		}
		if applied {
			if err := validateCreateTableStructure(working, meta.Object); err != nil {
				return "", err
			}
			if err := validateTableSubsetSupport(working); err != nil {
				return "", err
			}
			return normalizeSQLText(working, true, true), nil
		}
		return "", err
	case "indexes":
		nextSemantic, err := applyIndexStateTransition(meta.Object, sql, currentSemantic, hasCurrent)
		if err != nil {
			return "", err
		}
		if strings.TrimSpace(nextSemantic) != "" {
			if err := validateIndexSubsetSupport(nextSemantic); err != nil {
				return "", err
			}
		}
		return normalizeSQLText(nextSemantic, true, true), nil
	case "views", "functions", "procedures":
		stmt, err := extractProgrammabilityStatement(meta.Category, meta.Object, sql)
		if err != nil {
			return "", err
		}
		return normalizeSQLText(stmt, true, true), nil
	case "patch":
		return normalizeSQLText(sql, true, true), nil
	default:
		return "", fmt.Errorf("unsupported category %q", meta.Category)
	}
}

// migrationDropsObject recognizes a terminal object-removal migration. It is
// deliberately category-scoped and refuses to classify a DROP+CREATE
// lifecycle as deletion.
func migrationDropsObject(category, sql string) bool {
	upper := compactUpperSQL(sql)
	var dropToken, createToken string
	switch category {
	case "types":
		dropToken, createToken = "DROPTYPE", "CREATETYPE"
	case "tables":
		dropToken, createToken = "DROPTABLE", "CREATETABLE"
	case "views":
		dropToken, createToken = "DROPVIEW", "CREATEORALTERVIEW"
	case "functions":
		dropToken, createToken = "DROPFUNCTION", "CREATEORALTERFUNCTION"
	case "procedures":
		dropToken, createToken = "DROPPROCEDURE", "CREATEORALTERPROCEDURE"
	default:
		return false
	}
	return strings.Contains(upper, dropToken) && !strings.Contains(upper, createToken)
}

// semanticFromSource constructs the normalized desired state directly from
// db_object. The returned string is a derived view, never a third source.
func semanticFromSource(meta migrationMeta, sql string) (string, error) {
	switch meta.Category {
	case "types":
		return validateTableTypeSource(meta.Object, sql)
	case "tables":
		stmt, err := extractCreateTableStatement(meta.Object, sql)
		if err != nil {
			return "", err
		}
		if err := validateCreateTableStructure(stmt, meta.Object); err != nil {
			return "", err
		}
		stmt, err = canonicalizeCreateTableSourceStatement(stmt)
		if err != nil {
			return "", err
		}
		if err := validateTableSubsetSupport(stmt); err != nil {
			return "", err
		}
		return normalizeSQLText(stmt, true, true), nil
	case "indexes":
		if err := validateIndexSubsetSupport(sql); err != nil {
			return "", err
		}
		if err := validateIndexSourceTargets(meta.Object, sql); err != nil {
			return "", err
		}
		stmts, err := extractCreateIndexStatements(meta.Object, sql)
		if err != nil {
			return "", err
		}
		joined := strings.Join(stmts, "\n")
		if err := validateIndexSubsetSupport(joined); err != nil {
			return "", err
		}
		return normalizeSQLText(joined, true, true), nil
	case "views", "functions", "procedures":
		stmt, err := extractProgrammabilityStatement(meta.Category, meta.Object, sql)
		if err != nil {
			return "", err
		}
		if meta.Category == "functions" {
			if _, err := validateFunctionSource(stmt); err != nil {
				return "", err
			}
		}
		return normalizeSQLText(stmt, true, true), nil
	default:
		return "", fmt.Errorf("unsupported source category %q", meta.Category)
	}
}
