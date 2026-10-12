package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"sync"
	"unicode"

	"rza_migrate_db/internal/diff"
	"rza_migrate_db/internal/emit"
	"rza_migrate_db/internal/fsx"
	"rza_migrate_db/internal/history"
	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlparse"
)

func (e Engine) planGeneration(report *output.Report) ([]plannedMigration, map[string]struct{}, string, error) {
	historyState, existingFiles, maxVersion, err := e.reconstructHistory(report)
	if err != nil {
		return nil, nil, "", err
	}
	if report.Fail.Count > 0 {
		return nil, existingFiles, maxVersion, nil
	}

	desiredObjects, err := e.collectDesiredObjects(report)
	if err != nil {
		return nil, nil, "", err
	}
	if report.Fail.Count > 0 {
		return nil, existingFiles, maxVersion, nil
	}

	patchSources, err := e.collectPatchSources(report)
	if err != nil {
		return nil, nil, "", err
	}
	if report.Fail.Count > 0 {
		return nil, existingFiles, maxVersion, nil
	}

	return e.planGenerationWithInputs(report, historyState, existingFiles, maxVersion, desiredObjects, patchSources)
}

func (e Engine) planGenerationWithInputs(
	report *output.Report,
	historyState map[string]historyObject,
	existingFiles map[string]struct{},
	maxVersion string,
	desiredObjects []desiredObject,
	patchSources []patchSource,
) ([]plannedMigration, map[string]struct{}, string, error) {
	sortDesiredObjectsByOrder(desiredObjects, e.cfg.GenerationOrder)

	desiredByKey := make(map[string]desiredObject, len(desiredObjects))
	for _, d := range desiredObjects {
		if isEngineInfrastructureObject(d.Category, d.Object) {
			continue
		}
		desiredByKey[objectKey(d.Category, d.Object)] = d
	}
	unmanaged, err := e.loadUnmanagedObjectSet()
	if err != nil {
		return nil, existingFiles, maxVersion, fmt.Errorf("load object management manifest: %w", err)
	}

	ops := make([]plannedMigration, 0)
	for _, d := range desiredObjects {
		if isEngineInfrastructureObject(d.Category, d.Object) {
			continue
		}
		key := objectKey(d.Category, d.Object)
		last, hasHistory := historyState[key]
		if hasHistory && !last.Deleted && last.NormalizedHash == d.NormalizedHash {
			continue
		}

		nextVK := 1
		downSQL := ""
		upSQL := d.SemanticSQL
		reason := "CREATED"
		hasDestructiveDown := false
		removedColumns := []string{}
		if hasHistory {
			nextVK = last.LastVK + 1
		}
		if hasHistory && !last.Deleted {
			downSQL = last.SemanticSQL
			reason = "ALTERED"

			if d.Category == "types" {
				typeUp, typeDown, dependents, typeErr := buildTableTypeLifecycleSQL(
					d.Object,
					last.SemanticSQL,
					d.SemanticSQL,
					historyState,
					desiredByKey,
					e.cfg.GenerationOrder,
				)
				if typeErr != nil {
					report.AddFail(output.Item{
						Rule:       "TYPE_DEPENDENCY_PLAN_INVALID",
						Severity:   3,
						File:       d.SourceFile,
						Object:     d.Object,
						Message:    typeErr.Error(),
						Suggestion: "resolve circular module dependencies or provide a reviewed manual migration",
						Reason:     "REQUIRES_MANUAL",
						Action:     "fail",
					})
					continue
				}
				upSQL, downSQL = typeUp, typeDown
				report.AddPass(output.Item{
					Rule:     "TYPE_RECREATE_DEPENDENCIES",
					Severity: 1,
					File:     d.SourceFile,
					Object:   d.Object,
					Message:  fmt.Sprintf("table type change will recreate %d dependent function/procedure object(s): %s", len(dependents), strings.Join(dependents, ", ")),
					Reason:   "REPLACED",
					Action:   "run",
				})
			}
			if d.Category == "tables" {
				renameUp, renameDown, renameOK, renameErr := planColumnRenames(last.SemanticSQL, d.SemanticSQL, d.Object, d.RenameColumns, historyState)
				if renameErr != nil {
					report.AddFail(output.Item{Rule: "COLUMN_RENAME_PLAN_INVALID", Severity: 3, File: d.SourceFile, Object: d.Object, Message: renameErr.Error(), Suggestion: "keep rename-only changes in one migration and preserve the exact column type/modifiers", Reason: "REQUIRES_MANUAL", Action: "fail"})
					continue
				}
				if renameOK {
					upSQL, downSQL = renameUp, renameDown
					report.AddPass(output.Item{Rule: "TYPED_COLUMN_RENAME", Severity: 1, File: d.SourceFile, Object: d.Object, Message: "column rename planned with reversible sp_rename up/down", Reason: "OBJECT_EXISTS", Action: "run"})
				} else {
					risk, riskErr := assessTableTransition(last.SemanticSQL, d.SemanticSQL, d.Object)
					if riskErr != nil {
						report.AddFail(output.Item{Rule: "TABLE_TYPED_MODEL_PARSE_FAILED", Severity: 3, File: d.SourceFile, Object: d.Object, Message: riskErr.Error(), Reason: "REQUIRES_MANUAL", Action: "fail"})
						continue
					}
					if risk.Level == riskManual {
						report.AddFail(output.Item{
							Rule: "MANUAL_MIGRATION_REQUIRED", Severity: 3, File: d.SourceFile, Object: d.Object,
							Message:    strings.Join(risk.Reasons, "; "),
							Suggestion: "add a reviewed up/down pair directly to migrations/, then retain db_object as desired state",
							Reason:     "REQUIRES_MANUAL", Action: "fail",
						})
						continue
					}
					if risk.Level == riskGuarded {
						report.AddPass(output.Item{
							Rule: "GUARDED_AUTO_MIGRATION", Severity: 2, File: d.SourceFile, Object: d.Object,
							Message:    strings.Join(risk.Reasons, "; "),
							Suggestion: "run generated precheck against a production-like copy before deployment; it will fail atomically if data is incompatible",
							Reason:     "OBJECT_EXISTS", Action: "warn",
						})
					}
					colChanges, cerr := classifyTableColumnChanges(last.SemanticSQL, d.SemanticSQL)
					if cerr != nil {
						report.AddFail(output.Item{
							Rule:     "TABLE_DIFF_PARSE_FAILED",
							Severity: 3,
							File:     d.SourceFile,
							Object:   d.Object,
							Message:  cerr.Error(),
							Reason:   "REQUIRES_MANUAL",
							Action:   "fail",
						})
						continue
					}

					if len(colChanges.Removed) > 0 {
						if len(colChanges.Changed) > 0 {
							report.AddFail(output.Item{
								Rule:       "MIXED_TABLE_CHANGE_REQUIRES_MANUAL",
								Severity:   3,
								File:       d.SourceFile,
								Object:     d.Object,
								Message:    "mixed table change (drop with alter) requires manual migration",
								Suggestion: "split table evolution into explicit manual migration steps",
								Reason:     "REQUIRES_MANUAL",
								Action:     "fail",
							})
							continue
						}

						removedColumns = colChanges.Removed
						unnamedBlocked, unnamedErr := detectUnnamedConstraintDependencyBlocks(d.Object, removedColumns, last.SemanticSQL, historyState)
						if unnamedErr != nil {
							report.AddFail(output.Item{
								Rule:     "TABLE_DIFF_PARSE_FAILED",
								Severity: 3,
								File:     d.SourceFile,
								Object:   d.Object,
								Message:  unnamedErr.Error(),
								Reason:   "REQUIRES_MANUAL",
								Action:   "fail",
							})
							continue
						}
						if len(unnamedBlocked) > 0 {
							report.AddFail(output.Item{
								Rule:       "UNNAMED_CONSTRAINT_BLOCKS_DETERMINISTIC_PLAN",
								Severity:   3,
								File:       d.SourceFile,
								Object:     d.Object,
								Message:    "drop/alter plan blocked by unnamed dependency: " + strings.Join(unnamedBlocked, ", "),
								Suggestion: "rename impacted constraints/index dependencies explicitly, then rerun gen",
								Reason:     "REQUIRES_MANUAL",
								Action:     "fail",
							})
							continue
						}

						dropPlan, dropPlanErr := buildDropColumnDependencyPlan(d.Object, removedColumns, last.SemanticSQL, historyState)
						if dropPlanErr != nil {
							report.AddFail(output.Item{
								Rule:     "TABLE_DIFF_PARSE_FAILED",
								Severity: 3,
								File:     d.SourceFile,
								Object:   d.Object,
								Message:  dropPlanErr.Error(),
								Reason:   "REQUIRES_MANUAL",
								Action:   "fail",
							})
							continue
						}
						downRestoreSQL, downRestoreErr := buildDropColumnDownRestoreSQL(last.SemanticSQL, d.Object, removedColumns, dropPlan, historyState)
						if downRestoreErr != nil {
							report.AddFail(output.Item{
								Rule:     "TABLE_DIFF_PARSE_FAILED",
								Severity: 3,
								File:     d.SourceFile,
								Object:   d.Object,
								Message:  downRestoreErr.Error(),
								Reason:   "REQUIRES_MANUAL",
								Action:   "fail",
							})
							continue
						}
						dropUpSQL := buildDropColumnUpSQL(d.Object, removedColumns, dropPlan)
						if applyErr := applyConstraintDropsToHistoryState(historyState, dropPlan.Constraints); applyErr != nil {
							report.AddFail(output.Item{
								Rule:     "TABLE_DIFF_PARSE_FAILED",
								Severity: 3,
								File:     d.SourceFile,
								Object:   d.Object,
								Message:  applyErr.Error(),
								Reason:   "REQUIRES_MANUAL",
								Action:   "fail",
							})
							continue
						}
						if applyErr := applyIndexDropsToHistoryState(historyState, dropPlan.Indexes); applyErr != nil {
							report.AddFail(output.Item{
								Rule:     "TABLE_DIFF_PARSE_FAILED",
								Severity: 3,
								File:     d.SourceFile,
								Object:   d.Object,
								Message:  applyErr.Error(),
								Reason:   "REQUIRES_MANUAL",
								Action:   "fail",
							})
							continue
						}
						upSQL = dropUpSQL
						downSQL = downRestoreSQL
						if len(colChanges.Added) > 0 {
							addOnly := tableColumnChanges{
								Added: append([]string(nil), colChanges.Added...),
							}
							addUpSQL, addDownSQL, addDestructiveDown, addErr := buildTableAlterColumnsSQL(last.SemanticSQL, d.SemanticSQL, d.Object, addOnly)
							if addErr != nil {
								report.AddFail(output.Item{
									Rule:       "TABLE_ALTER_REQUIRES_MANUAL",
									Severity:   3,
									File:       d.SourceFile,
									Object:     d.Object,
									Message:    addErr.Error(),
									Suggestion: "split table evolution into explicit manual migration steps",
									Reason:     "REQUIRES_MANUAL",
									Action:     "fail",
								})
								continue
							}
							upSQL = concatSQLSegments(dropUpSQL, addUpSQL)
							downSQL = concatSQLSegments(addDownSQL, downRestoreSQL)
							if addDestructiveDown {
								hasDestructiveDown = true
							}
						}
						// The drop plan owns constraints that depend on removed columns.
						// Reconstruct its intermediate table state first, then plan only
						// the remaining named constraint delta. This preserves down order:
						// remove new constraints/columns before restoring old columns/FKs.
						intermediate, intermediateErr := semanticFromMigrationWithState(
							migrationMeta{Category: "tables", Object: d.Object}, upSQL, last.SemanticSQL, true,
						)
						if intermediateErr != nil {
							report.AddFail(output.Item{
								Rule: "TABLE_DIFF_PARSE_FAILED", Severity: 3, File: d.SourceFile, Object: d.Object,
								Message: intermediateErr.Error(), Reason: "REQUIRES_MANUAL", Action: "fail",
							})
							continue
						}
						constraintUp, constraintDown, constraintOK, constraintErr := buildNonDefaultConstraintLifecycleWithDependencies(intermediate, d.SemanticSQL, d.Object, historyState)
						if constraintErr != nil {
							report.AddFail(output.Item{
								Rule: "TABLE_CONSTRAINT_CHANGE_REQUIRES_MANUAL", Severity: 3, File: d.SourceFile, Object: d.Object,
								Message: constraintErr.Error(), Suggestion: "split the column and constraint changes or create a reviewed manual migration",
								Reason: "REQUIRES_MANUAL", Action: "fail",
							})
							continue
						}
						if constraintOK {
							constraintUpExecutable := constraintUp
							if len(colChanges.Added) > 0 {
								constraintUpExecutable = dynamicSQLBatch(constraintUp)
							}
							upSQL = concatSQLSegments(upSQL, constraintUpExecutable)
							downSQL = concatSQLSegments(constraintDown, downSQL)
						}
						hasDestructiveDown = true
						report.AddPass(output.Item{
							Rule:     "DESTRUCTIVE_DOWN_WARNING",
							Severity: 2,
							File:     d.SourceFile,
							Object:   d.Object,
							Message:  fmt.Sprintf("table change removes column(s): %s", strings.Join(removedColumns, ", ")),
							Reason:   "DATA_RESTORE_NOT_GUARANTEED",
							Action:   "warn",
						})
						report.AddPass(output.Item{
							Rule:     "INDEX_RECREATE_WARNING",
							Severity: 2,
							File:     d.SourceFile,
							Object:   d.Object,
							Message:  "drop column may remove impacted indexes; recreate indexes manually if needed",
							Reason:   "INDEX_RECREATE_REQUIRED",
							Action:   "warn",
						})
					} else {
						constraintOnlyPlanned := false
						if len(colChanges.Added) == 0 && len(colChanges.Changed) == 0 && len(colChanges.DefaultChanged) == 0 {
							constraintUpSQL, constraintDownSQL, constraintOK, constraintErr := buildConstraintLifecycleWithDependencies(last.SemanticSQL, d.SemanticSQL, d.Object, historyState)
							if constraintErr != nil {
								report.AddFail(output.Item{
									Rule:       "TABLE_CONSTRAINT_CHANGE_REQUIRES_MANUAL",
									Severity:   3,
									File:       d.SourceFile,
									Object:     d.Object,
									Message:    constraintErr.Error(),
									Suggestion: "create manual migration for constraint/index dependency updates",
									Reason:     "REQUIRES_MANUAL",
									Action:     "fail",
								})
								continue
							}
							if constraintOK {
								upSQL = constraintUpSQL
								downSQL = constraintDownSQL
								constraintOnlyPlanned = true
							} else {
								report.AddFail(output.Item{
									Rule:       "TABLE_CHANGE_REQUIRES_MANUAL",
									Severity:   3,
									File:       d.SourceFile,
									Object:     d.Object,
									Message:    "table change outside supported auto-plan (likely constraint-only change)",
									Suggestion: "create manual migration for constraint/index dependency updates",
									Reason:     "REQUIRES_MANUAL",
									Action:     "fail",
								})
								continue
							}
						}
						if constraintOnlyPlanned {
							// Constraint-only CHECK lifecycle was fully planned above.
						} else {
							unnamedBlocked, unnamedErr := detectUnnamedConstraintDependencyBlocks(d.Object, colChanges.Changed, last.SemanticSQL, historyState)
							if unnamedErr != nil {
								report.AddFail(output.Item{
									Rule:     "TABLE_DIFF_PARSE_FAILED",
									Severity: 3,
									File:     d.SourceFile,
									Object:   d.Object,
									Message:  unnamedErr.Error(),
									Reason:   "REQUIRES_MANUAL",
									Action:   "fail",
								})
								continue
							}
							if len(unnamedBlocked) > 0 {
								report.AddFail(output.Item{
									Rule:       "UNNAMED_CONSTRAINT_BLOCKS_DETERMINISTIC_PLAN",
									Severity:   3,
									File:       d.SourceFile,
									Object:     d.Object,
									Message:    "drop/alter plan blocked by unnamed dependency: " + strings.Join(unnamedBlocked, ", "),
									Suggestion: "rename impacted constraints/index dependencies explicitly, then rerun gen",
									Reason:     "REQUIRES_MANUAL",
									Action:     "fail",
								})
								continue
							}

							depPlan, depErr := buildDropColumnDependencyPlan(d.Object, colChanges.Changed, last.SemanticSQL, historyState)
							if depErr != nil {
								report.AddFail(output.Item{
									Rule:     "TABLE_DIFF_PARSE_FAILED",
									Severity: 3,
									File:     d.SourceFile,
									Object:   d.Object,
									Message:  depErr.Error(),
									Reason:   "REQUIRES_MANUAL",
									Action:   "fail",
								})
								continue
							}
							if hasUnsafeAlterColumnDependency(depPlan) {
								report.AddFail(output.Item{
									Rule:       "ALTER_COLUMN_DEPENDENCY_REQUIRES_MANUAL",
									Severity:   3,
									File:       d.SourceFile,
									Object:     d.Object,
									Message:    "alter column dependency is outside supported auto drop/recreate subset",
									Suggestion: "create manual migration to manage dependency drop/recreate explicitly",
									Reason:     "REQUIRES_MANUAL",
									Action:     "fail",
								})
								continue
							}
							advancedIndexBlocked, advancedIndexErr := findUnsafeAdvancedAlterIndexDependencies(depPlan, historyState)
							if advancedIndexErr != nil {
								report.AddFail(output.Item{
									Rule: "TABLE_DIFF_PARSE_FAILED", Severity: 3, File: d.SourceFile, Object: d.Object,
									Message: advancedIndexErr.Error(), Reason: "REQUIRES_MANUAL", Action: "fail",
								})
								continue
							}
							if len(advancedIndexBlocked) > 0 {
								report.AddFail(output.Item{
									Rule:       "ALTER_COLUMN_ADVANCED_INDEX_REQUIRES_MANUAL",
									Severity:   3,
									File:       d.SourceFile,
									Object:     d.Object,
									Message:    "advanced index dependency cannot be proven compatible with ALTER COLUMN: " + strings.Join(advancedIndexBlocked, ", "),
									Suggestion: "use a reviewed migration that drops/recreates the XML, spatial, or columnstore index with the target column type",
									Reason:     "REQUIRES_MANUAL",
									Action:     "fail",
								})
								continue
							}

							alterUpSQL, alterDownSQL, destructiveDown, alterErr := buildTableAlterColumnsSQL(last.SemanticSQL, d.SemanticSQL, d.Object, colChanges)
							if alterErr != nil {
								report.AddFail(output.Item{
									Rule:       "TABLE_ALTER_REQUIRES_MANUAL",
									Severity:   3,
									File:       d.SourceFile,
									Object:     d.Object,
									Message:    alterErr.Error(),
									Suggestion: "split table evolution into explicit manual migration steps",
									Reason:     "REQUIRES_MANUAL",
									Action:     "fail",
								})
								continue
							}
							// Reconstruct the state after dependency drops and column mutations,
							// then plan the remaining constraint delta once. This avoids restoring
							// an impacted CHECK/UQ twice when its definition also changed.
							depPrefix := buildDropDependencySQL(depPlan)
							baseUpSQL := concatSQLSegments(depPrefix, alterUpSQL)
							intermediate, intermediateErr := semanticFromMigrationWithState(
								migrationMeta{Category: "tables", Object: d.Object}, baseUpSQL, last.SemanticSQL, true,
							)
							if intermediateErr != nil {
								report.AddFail(output.Item{
									Rule: "TABLE_DIFF_PARSE_FAILED", Severity: 3, File: d.SourceFile, Object: d.Object,
									Message: intermediateErr.Error(), Reason: "REQUIRES_MANUAL", Action: "fail",
								})
								continue
							}
							constraintUp, constraintDown, _, constraintErr := buildNonDefaultConstraintLifecycleWithDependencies(intermediate, d.SemanticSQL, d.Object, historyState)
							if constraintErr != nil {
								report.AddFail(output.Item{
									Rule:       "TABLE_CONSTRAINT_CHANGE_REQUIRES_MANUAL",
									Severity:   3,
									File:       d.SourceFile,
									Object:     d.Object,
									Message:    constraintErr.Error(),
									Suggestion: "split the column and constraint changes or create a reviewed manual migration",
									Reason:     "REQUIRES_MANUAL",
									Action:     "fail",
								})
								continue
							}
							upIndexRestore, upIndexRestoreErr := buildReAddDroppedIndexesDownSQL(depPlan.Indexes, historyState)
							if upIndexRestoreErr != nil {
								report.AddFail(output.Item{Rule: "ALTER_COLUMN_DEPENDENCY_REQUIRES_MANUAL", Severity: 3, File: d.SourceFile, Object: d.Object, Message: upIndexRestoreErr.Error(), Reason: "REQUIRES_MANUAL", Action: "fail"})
								continue
							}
							defaultDependencies := make([]dropConstraintTarget, 0)
							nextConstraintDefinitions, nextConstraintDefinitionsErr := extractNamedConstraintDefinitionMap(d.SemanticSQL)
							if nextConstraintDefinitionsErr != nil {
								report.AddFail(output.Item{Rule: "TABLE_DIFF_PARSE_FAILED", Severity: 3, File: d.SourceFile, Object: d.Object, Message: nextConstraintDefinitionsErr.Error(), Reason: "REQUIRES_MANUAL", Action: "fail"})
								continue
							}
							for _, dependency := range depPlan.Constraints {
								if strings.EqualFold(strings.TrimSpace(dependency.Kind), "default") {
									if strings.TrimSpace(nextConstraintDefinitions[stripIdentifierBrackets(dependency.Name)]) != "" {
										defaultDependencies = append(defaultDependencies, dependency)
									}
								}
							}
							upDefaultRestore, upDefaultRestoreErr := buildReAddDroppedConstraintsDownSQL(d.SemanticSQL, d.Object, defaultDependencies, historyState)
							if upDefaultRestoreErr != nil {
								report.AddFail(output.Item{Rule: "ALTER_COLUMN_DEPENDENCY_REQUIRES_MANUAL", Severity: 3, File: d.SourceFile, Object: d.Object, Message: upDefaultRestoreErr.Error(), Reason: "REQUIRES_MANUAL", Action: "fail"})
								continue
							}
							downDepPlan, downDepPlanErr := excludeColumnLifecycleDefaultsFromDependencyRestore(depPlan, last.SemanticSQL, colChanges.DefaultChanged)
							if downDepPlanErr != nil {
								report.AddFail(output.Item{Rule: "TABLE_DIFF_PARSE_FAILED", Severity: 3, File: d.SourceFile, Object: d.Object, Message: downDepPlanErr.Error(), Reason: "REQUIRES_MANUAL", Action: "fail"})
								continue
							}
							depDownRestore, depDownRestoreErr := buildReAddDroppedDependenciesSQL(last.SemanticSQL, d.Object, downDepPlan, historyState)
							if depDownRestoreErr != nil {
								report.AddFail(output.Item{
									Rule:       "ALTER_COLUMN_DEPENDENCY_REQUIRES_MANUAL",
									Severity:   3,
									File:       d.SourceFile,
									Object:     d.Object,
									Message:    depDownRestoreErr.Error(),
									Suggestion: "create manual migration to manage dependency drop/recreate explicitly",
									Reason:     "REQUIRES_MANUAL",
									Action:     "fail",
								})
								continue
							}
							constraintUpExecutable := constraintUp
							if len(colChanges.Added) > 0 {
								constraintUpExecutable = dynamicSQLBatch(constraintUp)
							}
							upSQL = concatSQLSegments(baseUpSQL, constraintUpExecutable, upDefaultRestore, upIndexRestore)
							downSQL = concatSQLSegments(depPrefix, constraintDown, alterDownSQL, depDownRestore)
							if destructiveDown {
								hasDestructiveDown = true
								report.AddPass(output.Item{
									Rule:     "DESTRUCTIVE_DOWN_WARNING",
									Severity: 2,
									File:     d.SourceFile,
									Object:   d.Object,
									Message:  "down migration drops column(s) added by up migration",
									Reason:   "DATA_RESTORE_NOT_GUARANTEED",
									Action:   "warn",
								})
							}
						}
					}
				}
			}
			if d.Category == "indexes" {
				indexUpSQL, indexDownSQL, indexErr := buildIndexLifecycleSQL(last.SemanticSQL, d.SemanticSQL, d.Object)
				if indexErr != nil {
					report.AddFail(output.Item{
						Rule:     "INDEX_DIFF_PARSE_FAILED",
						Severity: 3,
						File:     d.SourceFile,
						Object:   d.Object,
						Message:  indexErr.Error(),
						Reason:   "REQUIRES_MANUAL",
						Action:   "fail",
					})
					continue
				}
				if strings.TrimSpace(indexUpSQL) == "" && strings.TrimSpace(indexDownSQL) == "" {
					continue
				}
				upSQL = indexUpSQL
				downSQL = indexDownSQL
			}
		} else if d.Category == "types" {
			downSQL = buildDropTypeSQL(d.Object)
		} else if d.Category == "tables" {
			downSQL = buildDropTableSQL(d.Object)
			hasDestructiveDown = true
			report.AddPass(output.Item{
				Rule:     "DESTRUCTIVE_DOWN_WARNING",
				Severity: 2,
				File:     d.SourceFile,
				Object:   d.Object,
				Message:  "down migration drops table created by up migration",
				Reason:   "DATA_RESTORE_NOT_GUARANTEED",
				Action:   "warn",
			})
		} else if d.Category == "indexes" {
			indexUpSQL, upErr := buildCreateIndexesSQLFromState(d.Object, d.SemanticSQL)
			if upErr != nil {
				report.AddFail(output.Item{
					Rule:     "INDEX_DIFF_PARSE_FAILED",
					Severity: 3,
					File:     d.SourceFile,
					Object:   d.Object,
					Message:  upErr.Error(),
					Reason:   "REQUIRES_MANUAL",
					Action:   "fail",
				})
				continue
			}
			indexDownSQL, indexErr := buildDropIndexesSQLFromState(d.Object, d.SemanticSQL)
			if indexErr != nil {
				report.AddFail(output.Item{
					Rule:     "INDEX_DIFF_PARSE_FAILED",
					Severity: 3,
					File:     d.SourceFile,
					Object:   d.Object,
					Message:  indexErr.Error(),
					Reason:   "REQUIRES_MANUAL",
					Action:   "fail",
				})
				continue
			}
			upSQL = indexUpSQL
			downSQL = indexDownSQL
		} else if isProgrammabilityCategory(d.Category) {
			// The first version of a programmable object must have a real
			// inverse. A comment-only down migration would let the migration
			// version move backwards while leaving the object installed.
			programmabilityDownSQL, downErr := buildDropProgrammabilitySQL(d.Category, d.Object)
			if downErr != nil {
				report.AddFail(output.Item{
					Rule:     "PROGRAMMABILITY_DOWN_PLAN_INVALID",
					Severity: 3,
					File:     d.SourceFile,
					Object:   d.Object,
					Message:  downErr.Error(),
					Reason:   "REQUIRES_MANUAL",
					Action:   "fail",
				})
				continue
			}
			downSQL = programmabilityDownSQL
		}
		ops = append(ops, plannedMigration{
			Category:           d.Category,
			Object:             d.Object,
			VK:                 nextVK,
			UpSemanticSQL:      upSQL,
			DownSemanticSQL:    downSQL,
			HasDestructiveDown: hasDestructiveDown,
			RemovedColumns:     removedColumns,
			SourceFile:         d.SourceFile,
		})

		report.AddPass(output.Item{
			Rule:     "PLAN_SCHEMA_MIGRATION",
			Severity: 1,
			File:     d.SourceFile,
			Object:   d.Object,
			Message:  fmt.Sprintf("planned %s migration for %s", strings.ToLower(reason), d.Object),
			Reason:   reason,
			Action:   "run",
		})
	}

	removedKeys := make([]string, 0)
	for key, st := range historyState {
		if st.Category == "patch" {
			continue
		}
		if isEngineInfrastructureObject(st.Category, st.Object) {
			continue
		}
		if _, ok := desiredByKey[key]; ok {
			continue
		}
		if _, ignored := unmanaged[key]; ignored {
			report.AddPass(output.Item{
				Rule: "UNMANAGED_OBJECT_REMOVAL_SKIPPED", Severity: 2,
				File: e.cfg.ManagementManifestFile, Object: st.Object,
				Message: "object is explicitly unmanaged because the extractor could not represent it losslessly; no removal migration was planned",
				Reason:  "REQUIRES_MANUAL", Action: "warn",
			})
			continue
		}
		removedKeys = append(removedKeys, key)
	}
	sort.Strings(removedKeys)
	for _, key := range removedKeys {
		st := historyState[key]
		if st.Deleted {
			continue
		}

		if st.Category == "tables" {
			nextVK := st.LastVK + 1
			if nextVK < 1 {
				nextVK = 1
			}

			ops = append(ops, plannedMigration{
				Category:        "tables",
				Object:          st.Object,
				VK:              nextVK,
				UpSemanticSQL:   buildDropTableSQL(st.Object),
				DownSemanticSQL: st.SemanticSQL,
				SourceFile:      e.cfg.DbObjectDir,
			})
			report.AddPass(output.Item{
				Rule:     "PLAN_SCHEMA_MIGRATION",
				Severity: 1,
				File:     e.cfg.DbObjectDir,
				Object:   st.Object,
				Message:  "planned guarded table removal; execution is blocked while data or dependencies remain",
				Reason:   "DROPPED",
				Action:   "run",
			})
			report.AddPass(output.Item{
				Rule:     "DESTRUCTIVE_UP_WARNING",
				Severity: 2,
				File:     e.cfg.DbObjectDir,
				Object:   st.Object,
				Message:  "table removal is destructive; generated prechecks require an empty, unreferenced table",
				Reason:   "DESTRUCTIVE_CHANGE",
				Action:   "warn",
			})
			continue
		}

		if st.Category == "indexes" {
			if strings.TrimSpace(st.SemanticSQL) == "" {
				report.AddPass(output.Item{
					Rule:     "INDEX_STATE_ALREADY_EMPTY",
					Severity: 0,
					File:     e.cfg.MigrationsDir,
					Object:   st.Object,
					Message:  "index state already empty; skip drop migration",
					Reason:   "OBJECT_NOT_EXISTS",
					Action:   "skip",
				})
				continue
			}

			nextVK := st.LastVK + 1
			if nextVK < 1 {
				nextVK = 1
			}
			dropUpSQL, dropErr := buildDropIndexesSQLFromState(st.Object, st.SemanticSQL)
			if dropErr != nil {
				report.AddFail(output.Item{
					Rule:     "INDEX_DIFF_PARSE_FAILED",
					Severity: 3,
					File:     e.cfg.DbObjectDir,
					Object:   st.Object,
					Message:  dropErr.Error(),
					Reason:   "REQUIRES_MANUAL",
					Action:   "fail",
				})
				continue
			}
			if strings.TrimSpace(dropUpSQL) == "" {
				report.AddPass(output.Item{
					Rule:     "INDEX_STATE_ALREADY_EMPTY",
					Severity: 0,
					File:     e.cfg.MigrationsDir,
					Object:   st.Object,
					Message:  "index state already empty; skip drop migration",
					Reason:   "OBJECT_NOT_EXISTS",
					Action:   "skip",
				})
				continue
			}
			restoreDownSQL, restoreErr := buildCreateIndexesSQLFromState(st.Object, st.SemanticSQL)
			if restoreErr != nil {
				report.AddFail(output.Item{
					Rule:     "INDEX_DIFF_PARSE_FAILED",
					Severity: 3,
					File:     e.cfg.DbObjectDir,
					Object:   st.Object,
					Message:  restoreErr.Error(),
					Reason:   "REQUIRES_MANUAL",
					Action:   "fail",
				})
				continue
			}

			ops = append(ops, plannedMigration{
				Category:        "indexes",
				Object:          st.Object,
				VK:              nextVK,
				UpSemanticSQL:   dropUpSQL,
				DownSemanticSQL: restoreDownSQL,
				SourceFile:      e.cfg.DbObjectDir,
			})
			report.AddPass(output.Item{
				Rule:     "PLAN_SCHEMA_MIGRATION",
				Severity: 1,
				File:     e.cfg.DbObjectDir,
				Object:   st.Object,
				Message:  fmt.Sprintf("planned dropped migration for %s", st.Object),
				Reason:   "DROPPED",
				Action:   "run",
			})
			continue
		}

		nextVK := st.LastVK + 1
		if nextVK < 1 {
			nextVK = 1
		}
		dropSQL := ""
		var dropErr error
		switch st.Category {
		case "types":
			dropSQL = buildGuardedDropTypeSQL(st.Object)
		case "views", "functions", "procedures":
			dropSQL, dropErr = buildGuardedDropProgrammabilitySQL(st.Category, st.Object)
		default:
			dropErr = fmt.Errorf("unsupported removal category %q", st.Category)
		}
		if dropErr != nil {
			report.AddFail(output.Item{
				Rule: "DROP_PLAN_INVALID", Severity: 3, File: e.cfg.DbObjectDir,
				Object: st.Object, Message: dropErr.Error(), Reason: "REQUIRES_MANUAL", Action: "fail",
			})
			continue
		}
		ops = append(ops, plannedMigration{
			Category: st.Category, Object: st.Object, VK: nextVK,
			UpSemanticSQL: dropSQL, DownSemanticSQL: st.SemanticSQL,
			SourceFile: e.cfg.DbObjectDir,
		})
		report.AddPass(output.Item{
			Rule: "PLAN_SCHEMA_MIGRATION", Severity: 1, File: e.cfg.DbObjectDir,
			Object: st.Object, Message: fmt.Sprintf("planned removal migration for %s", st.Object),
			Reason: "DROPPED", Action: "run",
		})
	}

	patchOps, err := e.planPatchMigrationsFromSources(report, historyState, patchSources)
	if err != nil {
		return nil, nil, "", err
	}
	ops = append(ops, patchOps...)

	ordered, orderErr := orderPlannedMigrations(ops, desiredByKey, e.cfg.GenerationOrder)
	if orderErr != nil {
		report.AddFail(output.Item{
			Rule:     "DEPENDENCY_CYCLE",
			Severity: 3,
			File:     e.cfg.DbObjectDir,
			Message:  orderErr.Error(),
			Reason:   "REQUIRES_MANUAL",
			Action:   "fail",
		})
		return nil, existingFiles, maxVersion, nil
	}
	report.Notes["dependency_order_enabled"] = true
	report.Notes["dependency_references_scanned"] = ordered.ReferencesScanned
	report.Notes["dependency_references_resolved"] = ordered.ResolvedReferences
	report.Notes["dependency_edges"] = ordered.Edges
	return ordered.Ops, existingFiles, maxVersion, nil
}

func (e Engine) reconstructHistory(report *output.Report) (map[string]historyObject, map[string]struct{}, string, error) {
	root := e.cfg.RepoPath(e.cfg.MigrationsDir)
	entries, err := os.ReadDir(root)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			report.AddFail(output.Item{
				Rule:     "MIGRATIONS_DIR_MISSING",
				Severity: 3,
				File:     e.cfg.MigrationsDir,
				Message:  "migrations directory is missing",
				Reason:   "OBJECT_NOT_EXISTS",
				Action:   "fail",
			})
			return map[string]historyObject{}, map[string]struct{}{}, "", nil
		}
		return nil, nil, "", err
	}

	snapshotHash, names, hashErr := e.migrationsSnapshotHash(root, entries)
	if hashErr != nil {
		return nil, nil, "", hashErr
	}
	if cached, hit, cacheErr := e.loadReconstructCache(snapshotHash); cacheErr != nil {
		report.Notes["reconstruct_cache_reset_reason"] = cacheErr.Error()
	} else if hit {
		existingFiles := make(map[string]struct{}, len(cached.Existing))
		for _, fileName := range cached.Existing {
			existingFiles[fileName] = struct{}{}
		}
		historyCopy := cloneHistoryStateMap(cached.History)
		e.applyReconstructionNotes(report, historyCopy, cached.MaxVersion)
		report.Notes["reconstruct_cache_hit"] = true
		report.Notes["reconstruct_cache_saved"] = true
		return historyCopy, existingFiles, cached.MaxVersion, nil
	}
	report.Notes["reconstruct_cache_hit"] = false

	conflictKeys := make(map[string]struct{})
	existingFiles := make(map[string]struct{}, len(names))
	history := make(map[string]historyObject)
	maxVersion := ""

	preloaded := e.preloadHistoryFiles(names)
	for _, pre := range preloaded {
		existingFiles[pre.Name] = struct{}{}
		if pre.Err != nil {
			return nil, nil, "", pre.Err
		}
		if pre.Issue != nil {
			report.AddFail(output.Item{
				Rule:     "HISTORY_FILENAME_PARSE_FAILED",
				Severity: 3,
				File:     filepath.ToSlash(filepath.Join(e.cfg.MigrationsDir, pre.Name)),
				Message:  pre.Issue.Message,
				Reason:   "REQUIRES_MANUAL",
				Action:   "fail",
			})
			continue
		}
		if pre.Skip {
			continue
		}

		meta := pre.Meta
		relPath := pre.RelPath
		normalized := pre.Normalized

		key := objectKey(meta.Category, meta.Object)
		current, hasCurrent := history[key]

		semantic, parseErr := semanticFromMigrationWithState(meta, normalized, current.SemanticSQL, hasCurrent)
		if parseErr != nil {
			report.AddFail(output.Item{
				Rule:     "UNSUPPORTED_SQL_SUBSET",
				Severity: 3,
				File:     relPath,
				Object:   meta.Object,
				Message:  parseErr.Error(),
				Reason:   "REQUIRES_MANUAL",
				Action:   "fail",
			})
			continue
		}

		conflictKey := fmt.Sprintf("%s|%s|%s|v%d", meta.Version, meta.Category, meta.Object, meta.VK)
		if _, seen := conflictKeys[conflictKey]; seen {
			report.AddFail(output.Item{
				Rule:     "HISTORY_CONFLICT_KEY",
				Severity: 3,
				File:     relPath,
				Object:   meta.Object,
				Message:  "duplicate history key detected",
				Reason:   "REQUIRES_MANUAL",
				Action:   "fail",
			})
			continue
		}
		conflictKeys[conflictKey] = struct{}{}

		next := historyObject{
			Category:       meta.Category,
			Object:         meta.Object,
			LastVersion:    meta.Version,
			LastVK:         meta.VK,
			SemanticSQL:    semantic,
			NormalizedHash: hashSemanticByCategory(meta.Category, semantic),
			Deleted:        migrationDropsObject(meta.Category, normalized),
			ManualDropOnly: meta.Category == "tables" && isManualDropTableRequestMigration(normalized),
		}
		if meta.Category == "functions" {
			next.Subtype, _ = classifyFunctionSubtype(semantic)
		}

		current, exists := history[key]
		if !exists || shouldReplaceHistory(current, next) {
			history[key] = next
		}

		if meta.Category == "tables" {
			droppedConstraints := extractDroppedConstraintsFromMigrationSQL(normalized)
			addedConstraints := extractAddedConstraintsFromMigrationSQL(normalized)
			for _, batch := range extractSpExecuteSQLBatches(normalized) {
				droppedConstraints = append(droppedConstraints, extractDroppedConstraintsFromMigrationSQL(batch)...)
				addedConstraints = append(addedConstraints, extractAddedConstraintsFromMigrationSQL(batch)...)
			}
			// semanticFromMigrationWithState already owns every constraint on the
			// migration's target table. Reapplying those transitions here can erase
			// a locally restored constraint when a guarded ADD is intentionally not
			// recognized by the cross-object scanner. This layer exists only to
			// propagate transitions such as incoming FKs on another table.
			crossObjectDrops := droppedConstraints[:0]
			for _, dependency := range droppedConstraints {
				if !sameQualifiedObject(dependency.Table, meta.Object) {
					crossObjectDrops = append(crossObjectDrops, dependency)
				}
			}
			droppedConstraints = crossObjectDrops
			crossObjectAdds := addedConstraints[:0]
			for _, dependency := range addedConstraints {
				if !sameQualifiedObject(dependency.Table, meta.Object) {
					crossObjectAdds = append(crossObjectAdds, dependency)
				}
			}
			addedConstraints = crossObjectAdds
			if applyErr := applyConstraintDropsToHistoryState(history, droppedConstraints); applyErr != nil {
				report.AddFail(output.Item{
					Rule:     "UNSUPPORTED_SQL_SUBSET",
					Severity: 3,
					File:     relPath,
					Object:   meta.Object,
					Message:  applyErr.Error(),
					Reason:   "REQUIRES_MANUAL",
					Action:   "fail",
				})
				continue
			}
			if applyErr := applyConstraintAddsToHistoryState(history, addedConstraints); applyErr != nil {
				report.AddFail(output.Item{
					Rule:     "UNSUPPORTED_SQL_SUBSET",
					Severity: 3,
					File:     relPath,
					Object:   meta.Object,
					Message:  applyErr.Error(),
					Reason:   "REQUIRES_MANUAL",
					Action:   "fail",
				})
				continue
			}
			droppedIndexes := extractDropIndexTargets(normalized)
			if applyErr := applyIndexDropsToHistoryState(history, droppedIndexes); applyErr != nil {
				report.AddFail(output.Item{
					Rule:     "UNSUPPORTED_SQL_SUBSET",
					Severity: 3,
					File:     relPath,
					Object:   meta.Object,
					Message:  applyErr.Error(),
					Reason:   "REQUIRES_MANUAL",
					Action:   "fail",
				})
				continue
			}
			createdIndexes := extractCreateIndexStatementsAllowEmpty(meta.Object, normalized)
			if applyErr := applyIndexCreatesToHistoryState(history, meta.Version, meta.Object, createdIndexes); applyErr != nil {
				report.AddFail(output.Item{
					Rule:     "UNSUPPORTED_SQL_SUBSET",
					Severity: 3,
					File:     relPath,
					Object:   meta.Object,
					Message:  applyErr.Error(),
					Reason:   "REQUIRES_MANUAL",
					Action:   "fail",
				})
				continue
			}
		}

		if len(meta.Version) == 17 && meta.Version > maxVersion {
			maxVersion = meta.Version
		}
	}

	e.applyReconstructionNotes(report, history, maxVersion)
	if report.Fail.Count == 0 {
		if err := e.saveReconstructCache(snapshotHash, history, existingFiles, maxVersion); err != nil {
			report.Notes["reconstruct_cache_write_error"] = err.Error()
			report.Notes["reconstruct_cache_saved"] = false
		} else {
			report.Notes["reconstruct_cache_saved"] = true
		}
	} else {
		report.Notes["reconstruct_cache_saved"] = false
	}
	return history, existingFiles, maxVersion, nil
}

func (e Engine) applyReconstructionNotes(report *output.Report, history map[string]historyObject, maxVersion string) {
	reconstruction := make(map[string]map[string]interface{})
	keys := make([]string, 0, len(history))
	for key := range history {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	const reconstructionSampleLimit = 300
	sampleCount := len(keys)
	if sampleCount > reconstructionSampleLimit {
		sampleCount = reconstructionSampleLimit
	}
	for i := 0; i < sampleCount; i++ {
		key := keys[i]
		st := history[key]
		reconstruction[key] = map[string]interface{}{
			"last_seen_version": st.LastVersion,
			"last_seen_vk":      st.LastVK,
			"normalized_hash":   st.NormalizedHash,
			"manual_drop_only":  st.ManualDropOnly,
		}
	}
	report.Notes["reconstruction"] = reconstruction
	report.Notes["reconstruction_total_objects"] = len(keys)
	if len(keys) > reconstructionSampleLimit {
		report.Notes["reconstruction_truncated"] = true
		report.Notes["reconstruction_sample_limit"] = reconstructionSampleLimit
	}
	if maxVersion != "" {
		report.Notes["last_seen_version"] = maxVersion
	}
}

type historyPreloadTask struct {
	Index int
	Name  string
}

type historyPreloadResult struct {
	Name       string
	Meta       migrationMeta
	Issue      *parseIssue
	RelPath    string
	Normalized string
	Skip       bool
	Err        error
}

func (e Engine) preloadHistoryFiles(names []string) []historyPreloadResult {
	if len(names) == 0 {
		return []historyPreloadResult{}
	}

	tasks := make([]historyPreloadTask, len(names))
	for i, name := range names {
		tasks[i] = historyPreloadTask{
			Index: i,
			Name:  name,
		}
	}

	results := make([]historyPreloadResult, len(tasks))
	workers := e.readParallelWorkers(len(tasks))
	if workers <= 1 {
		for _, task := range tasks {
			results[task.Index] = e.parseHistoryPreloadTask(task.Name)
		}
		return results
	}

	jobs := make(chan historyPreloadTask)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range jobs {
				results[task.Index] = e.parseHistoryPreloadTask(task.Name)
			}
		}()
	}

	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)
	wg.Wait()
	return results
}

func (e Engine) parseHistoryPreloadTask(name string) historyPreloadResult {
	meta, issue := e.parseMigrationFileName(name)
	if issue != nil {
		return historyPreloadResult{
			Name:  name,
			Issue: issue,
		}
	}

	result := historyPreloadResult{
		Name: name,
		Meta: meta,
	}
	if meta.IsCore || meta.Direction != "up" {
		result.Skip = true
		return result
	}

	relPath := filepath.ToSlash(filepath.Join(e.cfg.MigrationsDir, name))
	raw, err := e.readRepositoryFile(relPath)
	if err != nil {
		result.Err = err
		return result
	}

	result.RelPath = relPath
	result.Normalized = normalizeSQLText(string(raw), e.cfg.TrimTrailingSpaces, e.cfg.EnsureFinalNewline)
	return result
}

type desiredCollectTask struct {
	Index       int
	Category    string
	AbsFile     string
	RelFile     string
	Size        int64
	ModUnixNano int64
}

type desiredCollectResult struct {
	Object *desiredObject
	Fail   *output.Item
	Err    error
}

func (e Engine) collectDesiredObjects(report *output.Report) ([]desiredObject, error) {
	categories := []string{"types", "tables", "indexes", "views", "functions", "procedures"}

	tasks := make([]desiredCollectTask, 0)
	for _, category := range categories {
		files, err := fsx.ResolveDbObjectCategoryFiles(e.cfg.RepoPath, listFilesRecursive, e.cfg.DbObjectDir, category)
		if err != nil {
			return nil, err
		}
		for _, absFile := range files {
			if !strings.EqualFold(filepath.Ext(absFile), ".sql") {
				continue
			}
			relFile := e.toRelPath(absFile)
			tasks = append(tasks, desiredCollectTask{
				Index:    len(tasks),
				Category: category,
				AbsFile:  absFile,
				RelFile:  relFile,
			})
		}
	}

	snapshotHash, snapshotErr := e.desiredSnapshotHash(tasks)
	if snapshotErr != nil {
		return nil, snapshotErr
	}
	if cached, hit, cacheErr := e.loadDesiredCollectCache(snapshotHash); cacheErr != nil {
		report.Notes["desired_collect_cache_reset_reason"] = cacheErr.Error()
	} else if hit {
		report.Notes["desired_collect_cache_hit"] = true
		return cached, nil
	}
	report.Notes["desired_collect_cache_hit"] = false

	failBefore := report.Fail.Count
	results := e.runDesiredCollectTasks(tasks)
	out := make([]desiredObject, 0, len(results))
	for idx := range results {
		res := results[idx]
		if res.Err != nil {
			return nil, res.Err
		}
		if res.Fail != nil {
			report.AddFail(*res.Fail)
			continue
		}
		if res.Object != nil {
			out = append(out, *res.Object)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		if out[i].Category == out[j].Category {
			return out[i].Object < out[j].Object
		}
		return out[i].Category < out[j].Category
	})
	if report.Fail.Count == failBefore {
		if err := e.saveDesiredCollectCache(snapshotHash, out); err != nil {
			report.Notes["desired_collect_cache_write_error"] = err.Error()
			report.Notes["desired_collect_cache_saved"] = false
		} else {
			report.Notes["desired_collect_cache_saved"] = true
		}
	} else {
		report.Notes["desired_collect_cache_saved"] = false
	}
	return out, nil
}

func (e Engine) runDesiredCollectTasks(tasks []desiredCollectTask) []desiredCollectResult {
	if len(tasks) == 0 {
		return []desiredCollectResult{}
	}

	results := make([]desiredCollectResult, len(tasks))
	workers := e.readParallelWorkers(len(tasks))
	if workers <= 1 {
		for _, task := range tasks {
			results[task.Index] = e.parseDesiredCollectTask(task)
		}
		return results
	}

	jobs := make(chan desiredCollectTask)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range jobs {
				results[task.Index] = e.parseDesiredCollectTask(task)
			}
		}()
	}

	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)
	wg.Wait()
	return results
}

func (e Engine) parseDesiredCollectTask(task desiredCollectTask) desiredCollectResult {
	relFile := task.RelFile
	baseName := filepath.Base(task.AbsFile)
	object := strings.TrimSuffix(baseName, filepath.Ext(baseName))
	if !strings.HasPrefix(object, "dbo.") {
		fail := output.Item{
			Rule:     "INVALID_DB_OBJECT_FILENAME",
			Severity: 3,
			File:     relFile,
			Message:  "db_object file must use dbo.<Object>.sql",
			Reason:   "REQUIRES_MANUAL",
			Action:   "fail",
		}
		return desiredCollectResult{Fail: &fail}
	}

	raw, err := e.readRepositoryFile(task.RelFile)
	if err != nil {
		return desiredCollectResult{Err: err}
	}
	normalized := normalizeSQLText(string(raw), e.cfg.TrimTrailingSpaces, e.cfg.EnsureFinalNewline)
	if issue := assessSourceCapability(task.Category, normalized); issue != nil {
		fail := output.Item{
			Rule:       "UNSUPPORTED_SQL_FEATURE",
			Severity:   3,
			File:       relFile,
			Object:     object,
			Message:    issue.Error(),
			Suggestion: "write a reviewed migration pair directly in migrations/; do not add another source of truth",
			Reason:     "REQUIRES_MANUAL",
			Action:     "fail",
		}
		return desiredCollectResult{Fail: &fail}
	}

	meta := migrationMeta{
		Category: task.Category,
		Object:   object,
	}
	semantic, parseErr := semanticFromSource(meta, normalized)
	if parseErr != nil {
		rule := sourceParseFailureRule(meta.Category)
		fail := output.Item{
			Rule:     rule,
			Severity: 3,
			File:     relFile,
			Object:   object,
			Message:  parseErr.Error(),
			Reason:   "REQUIRES_MANUAL",
			Action:   "fail",
		}
		return desiredCollectResult{Fail: &fail}
	}

	obj := desiredObject{
		Category:       task.Category,
		Object:         object,
		SourceFile:     relFile,
		SemanticSQL:    semantic,
		NormalizedHash: hashSemanticByCategory(task.Category, semantic),
	}
	if task.Category == "functions" {
		subtype, subtypeErr := validateFunctionSource(semantic)
		if subtypeErr != nil {
			fail := output.Item{
				Rule:     "FUNCTION_SUBTYPE_INVALID",
				Severity: 3,
				File:     relFile,
				Object:   object,
				Message:  subtypeErr.Error(),
				Reason:   "REQUIRES_MANUAL",
				Action:   "fail",
			}
			return desiredCollectResult{Fail: &fail}
		}
		obj.Subtype = subtype
	}
	if task.Category == "tables" {
		renames, renameErr := parseColumnRenameDirectives(normalized)
		if renameErr != nil {
			fail := output.Item{Rule: "INVALID_COLUMN_RENAME_HINT", Severity: 3, File: relFile, Object: object, Message: renameErr.Error(), Suggestion: "use: -- rza:rename-column OldName -> NewName", Reason: "INVALID_SOURCE_LAYOUT", Action: "fail"}
			return desiredCollectResult{Fail: &fail}
		}
		obj.RenameColumns = renames
	}
	return desiredCollectResult{Object: &obj}
}

func sourceParseFailureRule(category string) string {
	switch strings.ToLower(strings.TrimSpace(category)) {
	case "types":
		return "TYPE_SOURCE_SYNTAX_INVALID"
	case "tables":
		return "TABLE_SOURCE_SYNTAX_INVALID"
	case "indexes":
		return "INDEX_SOURCE_SYNTAX_INVALID"
	case "views", "functions", "procedures":
		return "PROGRAMMABILITY_SOURCE_INVALID"
	default:
		return "UNSUPPORTED_SQL_SUBSET"
	}
}

func (e Engine) planPatchMigrations(report *output.Report, history map[string]historyObject) ([]plannedMigration, error) {
	patches, err := e.collectPatchSources(report)
	if err != nil {
		return nil, err
	}
	if report.Fail.Count > 0 {
		return nil, nil
	}
	return e.planPatchMigrationsFromSources(report, history, patches)
}

func (e Engine) planPatchMigrationsFromSources(report *output.Report, history map[string]historyObject, patches []patchSource) ([]plannedMigration, error) {
	byCanonical := make(map[string]patchSource, len(patches))
	for _, p := range patches {
		byCanonical[p.CanonicalName] = p
	}

	historyPatchNames := make([]string, 0)
	for _, st := range history {
		if st.Category != "patch" {
			continue
		}
		historyPatchNames = append(historyPatchNames, st.Object)
	}
	sort.Strings(historyPatchNames)
	for _, patchName := range historyPatchNames {
		if _, ok := byCanonical[patchName]; ok {
			continue
		}
		report.AddFail(output.Item{
			Rule:     "PATCH_SOURCE_MISSING",
			Severity: 3,
			File:     e.cfg.PatchDir,
			Object:   patchName,
			Message:  "existing patch migration has no matching patch source folder",
			Reason:   "DRIFT_DETECTED",
			Action:   "fail",
		})
	}
	if report.Fail.Count > 0 {
		return nil, nil
	}

	ops := make([]plannedMigration, 0)
	for _, patch := range patches {
		key := objectKey("patch", patch.CanonicalName)
		if _, exists := history[key]; exists {
			report.AddPass(output.Item{
				Rule:     "PATCH_DUPLICATION_SKIPPED",
				Severity: 1,
				File:     patch.RelDir,
				Object:   patch.CanonicalName,
				Message:  "patch migration already exists for canonical name, generation skipped",
				Reason:   "OBJECT_EXISTS",
				Action:   "skip",
			})
			continue
		}

		ops = append(ops, plannedMigration{
			Category:            "patch",
			Object:              patch.CanonicalName,
			VK:                  1,
			UpSemanticSQL:       patch.UpSQL,
			DownSemanticSQL:     patch.DownSQL,
			SourceFile:          patch.RelDir,
			PatchFolderOriginal: patch.FolderName,
			PatchNameCanonical:  patch.CanonicalName,
			PatchReadmePath:     patch.ReadmePath,
			PatchUpHash:         patch.UpHash,
			PatchDownHash:       patch.DownHash,
		})

		report.AddPass(output.Item{
			Rule:     "PLAN_PATCH_MIGRATION",
			Severity: 1,
			File:     patch.RelDir,
			Object:   patch.CanonicalName,
			Message:  "planned patch migration",
			Reason:   "CREATED",
			Action:   "run",
		})
	}

	return ops, nil
}

type patchCollectTask struct {
	Index      int
	Folder     string
	Canonical  string
	RelDir     string
	ReadmeAbs  string
	UpAbs      string
	DownAbs    string
	ReadmePath string
	UpPath     string
	DownPath   string
}

type patchCollectResult struct {
	Source *patchSource
	Fail   *output.Item
}

func (e Engine) collectPatchSources(report *output.Report) ([]patchSource, error) {
	failBefore := report.Fail.Count
	root := e.cfg.RepoPath(e.cfg.PatchDir)
	entries, err := os.ReadDir(root)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return []patchSource{}, nil
		}
		return nil, err
	}

	type namedEntry struct {
		Name string
		Dir  os.DirEntry
	}
	dirs := make([]namedEntry, 0)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		dirs = append(dirs, namedEntry{Name: entry.Name(), Dir: entry})
	}
	sort.Slice(dirs, func(i, j int) bool {
		return dirs[i].Name < dirs[j].Name
	})

	canonicalUsed := make(map[string]string)
	tasks := make([]patchCollectTask, 0, len(dirs))
	for _, dir := range dirs {
		folder := dir.Name
		canonical := canonicalizePatchName(folder)
		relDir := filepath.ToSlash(filepath.Join(e.cfg.PatchDir, folder))
		if canonical == "" {
			report.AddFail(output.Item{
				Rule:     "PATCH_CANONICAL_NAME_EMPTY",
				Severity: 3,
				File:     relDir,
				Message:  "patch canonical name is empty after normalization",
				Reason:   "REQUIRES_MANUAL",
				Action:   "fail",
			})
			continue
		}
		if prev, exists := canonicalUsed[canonical]; exists {
			report.AddFail(output.Item{
				Rule:     "PATCH_CANONICAL_NAME_DUPLICATE",
				Severity: 3,
				File:     relDir,
				Message:  "duplicate canonical patch name with " + prev,
				Reason:   "REQUIRES_MANUAL",
				Action:   "fail",
			})
			continue
		}
		canonicalUsed[canonical] = relDir
		resolved, complete, resolveErr := fsx.ResolveLogicalFilesCaseInsensitive(filepath.Join(root, folder), "README.md", "up.sql", "down.sql")
		if resolveErr != nil {
			return nil, resolveErr
		}
		if !complete {
			report.AddFail(output.Item{
				Rule:       "PATCH_FOLDER_INVALID_CONTENT",
				Severity:   3,
				File:       relDir,
				Message:    "patch folder must contain exactly one README.md, up.sql, and down.sql logical file",
				Suggestion: "remove extra/duplicate entries or restore the missing patch source file; filename casing may vary",
				Reason:     "REQUIRES_MANUAL",
				Action:     "fail",
			})
			continue
		}
		readmeAbs := resolved["readme.md"]
		upAbs := resolved["up.sql"]
		downAbs := resolved["down.sql"]

		tasks = append(tasks, patchCollectTask{
			Index:      len(tasks),
			Folder:     folder,
			Canonical:  canonical,
			RelDir:     relDir,
			ReadmeAbs:  readmeAbs,
			UpAbs:      upAbs,
			DownAbs:    downAbs,
			ReadmePath: filepath.ToSlash(filepath.Join(relDir, filepath.Base(readmeAbs))),
			UpPath:     filepath.ToSlash(filepath.Join(relDir, filepath.Base(upAbs))),
			DownPath:   filepath.ToSlash(filepath.Join(relDir, filepath.Base(downAbs))),
		})
	}

	if report.Fail.Count == failBefore {
		snapshotHash, hashErr := buildPatchCollectSnapshotHash(tasks)
		if hashErr != nil {
			return nil, hashErr
		}
		if cached, hit, cacheErr := e.loadPatchCollectCache(snapshotHash); cacheErr != nil {
			report.Notes["patch_collect_cache_reset_reason"] = cacheErr.Error()
		} else if hit {
			report.Notes["patch_collect_cache_hit"] = true
			report.Notes["patch_collect_cache_saved"] = true
			return cached, nil
		}
		report.Notes["patch_collect_cache_hit"] = false

		results := e.runPatchCollectTasks(tasks)
		out := make([]patchSource, 0, len(results))
		for idx := range results {
			res := results[idx]
			if res.Fail != nil {
				report.AddFail(*res.Fail)
				continue
			}
			if res.Source != nil {
				out = append(out, *res.Source)
			}
		}

		sort.Slice(out, func(i, j int) bool {
			return out[i].CanonicalName < out[j].CanonicalName
		})
		if report.Fail.Count == failBefore {
			if err := e.savePatchCollectCache(snapshotHash, out); err != nil {
				report.Notes["patch_collect_cache_write_error"] = err.Error()
				report.Notes["patch_collect_cache_saved"] = false
			} else {
				report.Notes["patch_collect_cache_saved"] = true
			}
		} else {
			report.Notes["patch_collect_cache_saved"] = false
		}
		return out, nil
	}

	results := e.runPatchCollectTasks(tasks)
	out := make([]patchSource, 0, len(results))
	for idx := range results {
		res := results[idx]
		if res.Fail != nil {
			report.AddFail(*res.Fail)
			continue
		}
		if res.Source != nil {
			out = append(out, *res.Source)
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].CanonicalName < out[j].CanonicalName
	})
	report.Notes["patch_collect_cache_saved"] = false
	return out, nil
}

func (e Engine) runPatchCollectTasks(tasks []patchCollectTask) []patchCollectResult {
	if len(tasks) == 0 {
		return []patchCollectResult{}
	}

	results := make([]patchCollectResult, len(tasks))
	workers := e.readParallelWorkers(len(tasks))
	if workers <= 1 {
		for _, task := range tasks {
			results[task.Index] = e.parsePatchCollectTask(task)
		}
		return results
	}

	jobs := make(chan patchCollectTask)
	var wg sync.WaitGroup
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for task := range jobs {
				results[task.Index] = e.parsePatchCollectTask(task)
			}
		}()
	}

	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)
	wg.Wait()
	return results
}

func (e Engine) parsePatchCollectTask(task patchCollectTask) patchCollectResult {
	if _, err := os.ReadFile(task.ReadmeAbs); err != nil {
		fail := output.Item{
			Rule:     "PATCH_FOLDER_INVALID_CONTENT",
			Severity: 3,
			File:     task.RelDir,
			Message:  "missing README.md in patch folder",
			Reason:   "REQUIRES_MANUAL",
			Action:   "fail",
		}
		return patchCollectResult{Fail: &fail}
	}

	upBytes, err := os.ReadFile(task.UpAbs)
	if err != nil {
		fail := output.Item{
			Rule:     "PATCH_FOLDER_INVALID_CONTENT",
			Severity: 3,
			File:     task.RelDir,
			Message:  "missing up.sql in patch folder",
			Reason:   "REQUIRES_MANUAL",
			Action:   "fail",
		}
		return patchCollectResult{Fail: &fail}
	}

	downBytes, err := os.ReadFile(task.DownAbs)
	if err != nil {
		fail := output.Item{
			Rule:     "PATCH_FOLDER_INVALID_CONTENT",
			Severity: 3,
			File:     task.RelDir,
			Message:  "missing down.sql in patch folder",
			Reason:   "REQUIRES_MANUAL",
			Action:   "fail",
		}
		return patchCollectResult{Fail: &fail}
	}

	src := patchSource{
		FolderName:    task.Folder,
		CanonicalName: task.Canonical,
		RelDir:        task.RelDir,
		ReadmePath:    task.ReadmePath,
		UpPath:        task.UpPath,
		DownPath:      task.DownPath,
		UpSQL:         normalizeSQLText(string(upBytes), e.cfg.TrimTrailingSpaces, e.cfg.EnsureFinalNewline),
		DownSQL:       normalizeSQLText(string(downBytes), e.cfg.TrimTrailingSpaces, e.cfg.EnsureFinalNewline),
		UpHash:        hashStableContent(string(upBytes)),
		DownHash:      hashStableContent(string(downBytes)),
	}
	return patchCollectResult{Source: &src}
}

func (e Engine) readParallelWorkers(taskCount int) int {
	if taskCount <= 1 {
		return 1
	}
	if !e.cfg.EnableParallelRead {
		return 1
	}
	workers := e.cfg.ParallelWorkers
	if workers < 1 {
		workers = 1
	}
	cpuCap := runtime.GOMAXPROCS(0)
	if cpuCap < 1 {
		cpuCap = 1
	}
	if workers > cpuCap {
		workers = cpuCap
	}
	if workers > taskCount {
		workers = taskCount
	}
	return workers
}

func (e Engine) allocateMigrationBatch(
	ops []plannedMigration,
	baseTimestamp string,
	startSeq int,
	reserved map[string]struct{},
) ([]allocatedMigration, int, int, error) {
	pairSeq := startSeq
	allocated := make([]allocatedMigration, 0, len(ops))
	for idx, op := range ops {
		version, nextSeq, err := e.allocateVersion(baseTimestamp, pairSeq, reserved, op)
		if err != nil {
			if idx == 0 {
				return nil, len(ops), pairSeq, err
			}
			return allocated, len(ops) - idx, pairSeq, nil
		}
		pairSeq = nextSeq

		upName, downName := migrationFileNames(version, op)
		reserved[upName] = struct{}{}
		reserved[downName] = struct{}{}
		emit.ReserveVersion(reserved, version)
		allocated = append(allocated, allocatedMigration{
			Op:       op,
			Version:  version,
			UpName:   upName,
			DownName: downName,
		})
	}
	return allocated, 0, pairSeq, nil
}

func (e Engine) allocateVersion(baseTimestamp string, startSeq int, reserved map[string]struct{}, op plannedMigration) (string, int, error) {
	return emit.AllocateVersion(baseTimestamp, startSeq, e.cfg.PairSeqWidth, e.cfg.MaxPairsPerRun, reserved, emit.PlannedMigrationRef{
		Category: op.Category,
		Object:   op.Object,
		VK:       op.VK,
	})
}

func migrationFileNames(version string, op plannedMigration) (string, string) {
	return emit.MigrationFileNames(version, emit.PlannedMigrationRef{
		Category: op.Category,
		Object:   op.Object,
		VK:       op.VK,
	})
}

func renderTxWrapper(version, category, fileName, target string, isUp bool, body string) string {
	direction := "down"
	okReason := "DROPPED"
	if isUp {
		direction = "up"
		okReason = "CREATED"
	}
	action := auditActionForWrapper(category, isUp, body)
	precheckBlock := buildWrapperPrecheckSQL(category, target, action, isUp, body)
	canonicalBody := strings.TrimSpace(compactGeneratedSQLWhitespace(body))
	contentHash := sha256.Sum256([]byte(canonicalBody))
	script := fmt.Sprintf(`/*
Severity mapping:
0=INFO (ok/skip normal)
1=NOTICE (schema change)
2=WARNING (follow-up/manual risk)
3=ERROR (stop)
*/
SET NOCOUNT ON;
SET XACT_ABORT ON;
DECLARE @run_key UNIQUEIDENTIFIER = NEWID();
DECLARE @operation_id UNIQUEIDENTIFIER = TRY_CONVERT(UNIQUEIDENTIFIER, SESSION_CONTEXT(N'mig_operation_id'));
IF @operation_id IS NULL SET @operation_id = @run_key;
DECLARE @version BIGINT = %s;
DECLARE @category NVARCHAR(16) = N'%s';
DECLARE @file NVARCHAR(260) = N'%s';
DECLARE @target NVARCHAR(256) = N'%s';
DECLARE @action NVARCHAR(64) = N'%s';
DECLARE @direction NVARCHAR(8) = N'%s';
DECLARE @content_sha256 CHAR(64) = N'%x';
DECLARE @started_at DATETIME2(7) = SYSUTCDATETIME();

EXEC sys.sp_set_session_context @key=N'mig_operation_id', @value=@operation_id;
EXEC sys.sp_set_session_context @key=N'mig_direction', @value=@direction;
EXEC sys.sp_set_session_context @key=N'mig_content_sha256', @value=@content_sha256;
EXEC sys.sp_set_session_context @key=N'mig_started_at', @value=@started_at;

IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NULL
BEGIN
  ;THROW 51000, N'Migration audit infrastructure is unavailable', 1;
END

%s

BEGIN TRY
  BEGIN TRAN;

  EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'migration run start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;

-- MIGRATION_BODY_BEGIN
%s
-- MIGRATION_BODY_END

  EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'ok',
      @severity = 1,
      @action = @action,
      @target = @target,
      @reason = N'%s',
      @message = N'migration step ok',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;

  COMMIT TRAN;
END TRY
BEGIN CATCH
  IF XACT_STATE() <> 0 ROLLBACK TRAN;
  DECLARE @err_message NVARCHAR(4000) = ERROR_MESSAGE();
  DECLARE @err_number INT = ERROR_NUMBER();
  DECLARE @err_line INT = ERROR_LINE();
  DECLARE @err_procedure NVARCHAR(128) = ERROR_PROCEDURE();

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'fail',
      @severity = 3,
      @action = N'migration_failed',
      @target = @target,
      @reason = N'SQL_ERROR',
      @message = N'migration failed',
      @detail = @err_message,
      @error_number = @err_number,
      @error_line = @err_line,
      @error_procedure = @err_procedure;
  END
  ;THROW;
END CATCH;
`, version, category, fileName, target, action, direction, contentHash, precheckBlock, canonicalBody, okReason)
	return compactGeneratedSQLWhitespace(script)
}

func buildWrapperPrecheckSQL(category, target, action string, isUp bool, body string) string {
	switch category {
	case "types":
		return `-- type existence and dependent modules are validated by the migration body`
	case "tables":
		if action == "create_table" && strings.Contains(strings.ToUpper(body), "CREATE TABLE") {
			return buildTableWrapperPrecheck(target, true, body)
		}
		if action == "drop_table" && strings.Contains(strings.ToUpper(body), "DROP TABLE") {
			return buildTableWrapperPrecheck(target, false, body)
		}
		return `-- precheck is handled per statement in the migration body
-- check table existence
-- check column existence
-- check data length before ALTER COLUMN
-- check data convertibility before type change`
	case "indexes":
		return `-- precheck is handled per statement in the migration body
-- check target table existence
-- check index existence`
	case "patch":
		return `-- precheck is defined by patch SQL body
-- engine wrapper still enforces transaction and audit logging`
	}
	return "-- precheck is not required for this category"
}

func buildTableMixedColumnWrapperPrecheck(target string, droppedColumns []string, addedColumns []string) string {
	safeTarget := escapeSQLLiteral(target)
	buildValueRows := func(columns []string) []string {
		out := make([]string, 0, len(columns))
		seen := make(map[string]struct{}, len(columns))
		for _, col := range columns {
			trimmed := strings.TrimSpace(col)
			if trimmed == "" {
				continue
			}
			key := trimmed
			if _, exists := seen[key]; exists {
				continue
			}
			seen[key] = struct{}{}
			out = append(out, fmt.Sprintf("(N'%s')", escapeSQLLiteral(trimmed)))
		}
		sort.Strings(out)
		return out
	}

	dropRows := buildValueRows(droppedColumns)
	addRows := buildValueRows(addedColumns)
	if len(dropRows) == 0 || len(addRows) == 0 {
		return `-- precheck skipped: mixed-column wrapper check requires both drop and add column lists`
	}

	return fmt.Sprintf(`IF OBJECT_ID(N'%s', N'U') IS NULL
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_NOT_EXISTS',
      @message = N'Precheck: table already absent',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END

IF OBJECT_ID(N'%s', N'U') IS NOT NULL
AND NOT EXISTS (
  SELECT 1
  FROM (VALUES %s) AS d(col_name)
  WHERE COL_LENGTH(N'%s', d.col_name) IS NOT NULL
)
AND NOT EXISTS (
  SELECT 1
  FROM (VALUES %s) AS a(col_name)
  WHERE COL_LENGTH(N'%s', a.col_name) IS NULL
)
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_EXISTS',
      @message = N'Precheck: mixed column state already applied',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END`, safeTarget, safeTarget, strings.Join(dropRows, ", "), safeTarget, strings.Join(addRows, ", "), safeTarget)
}

func buildTableColumnWrapperPrecheck(target string, columns []string, isAdd bool) string {
	safeTarget := escapeSQLLiteral(target)
	valueRows := make([]string, 0, len(columns))
	for _, col := range columns {
		trimmed := strings.TrimSpace(col)
		if trimmed == "" {
			continue
		}
		valueRows = append(valueRows, fmt.Sprintf("(N'%s')", escapeSQLLiteral(trimmed)))
	}
	if len(valueRows) == 0 {
		return `-- precheck skipped: no target columns found for wrapper-level check`
	}
	valuesSQL := strings.Join(valueRows, ", ")

	if isAdd {
		return fmt.Sprintf(`IF OBJECT_ID(N'%s', N'U') IS NOT NULL
AND NOT EXISTS (
  SELECT 1
  FROM (VALUES %s) AS c(col_name)
  WHERE COL_LENGTH(N'%s', c.col_name) IS NULL
)
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'COLUMN_EXISTS',
      @message = N'Precheck: target column(s) already exist',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END`, safeTarget, valuesSQL, safeTarget)
	}

	return fmt.Sprintf(`IF OBJECT_ID(N'%s', N'U') IS NULL
OR NOT EXISTS (
  SELECT 1
  FROM (VALUES %s) AS c(col_name)
  WHERE COL_LENGTH(N'%s', c.col_name) IS NOT NULL
)
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'COLUMN_NOT_EXISTS',
      @message = N'Precheck: target column(s) already absent',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END`, safeTarget, valuesSQL, safeTarget)
}

func buildTableWrapperPrecheck(target string, isCreate bool, createSQL string) string {
	safeTarget := escapeSQLLiteral(target)
	if isCreate {
		equivalent, supported := buildExistingTableEquivalentCondition(target, createSQL)
		if supported {
			skipLog := indentSQL(buildStepAuditLogSQL("skip", 0, "create_table", target, "OBJECT_EXISTS", "Precheck: bootstrap-compatible table already exists"), "    ")
			return fmt.Sprintf(`IF OBJECT_ID(N'%s', N'U') IS NOT NULL
BEGIN
  -- existing table match profile: bootstrap-compatible
  IF %s
  BEGIN
%s
    RETURN;
  END

  ;THROW 51000, N'Schema drift: existing table is missing a required column or has an incompatible base type, identity-kind, or computed-kind', 1;
END`, safeTarget, equivalent, skipLog)
		}
		return fmt.Sprintf(`IF OBJECT_ID(N'%s', N'U') IS NOT NULL
BEGIN
  ;THROW 51000, N'Schema drift: CREATE TABLE target already exists; equivalence is not proven', 1;
END`, safeTarget)
	}

	return fmt.Sprintf(`IF OBJECT_ID(N'%s', N'U') IS NULL
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_NOT_EXISTS',
      @message = N'Precheck: table already absent',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END`, safeTarget)
}

// buildExistingTableEquivalentCondition proves bootstrap compatibility for an
// existing CREATE TABLE target. It intentionally checks only the schema facts
// required for safe application startup: every requested column must exist
// with a compatible base type, identity-kind, and computed-kind. Length,
// precision/scale, nullability, collation, ordinal position, defaults,
// constraints, physical index options, and extra database columns are
// non-blocking because they are commonly managed by DBAs. Explicit ALTER,
// DROP, ADD CONSTRAINT, and CREATE INDEX migrations retain their own strict
// prechecks elsewhere in this generator.
func buildExistingTableEquivalentCondition(target, createSQL string) (string, bool) {
	items, err := extractCreateTableItems(createSQL)
	if err != nil {
		return "", false
	}
	requiredColumns := make([]tableColumnDef, 0, len(items))
	for _, item := range items {
		upper := strings.ToUpper(strings.TrimSpace(item))
		if isConstraintLikeTableItem(upper) || strings.HasPrefix(upper, "DEFAULT ") {
			continue
		}
		def, parseErr := parseTableColumnDefFragment(item)
		if parseErr != nil || def.Name == "" {
			return "", false
		}
		requiredColumns = append(requiredColumns, def)
	}
	if len(requiredColumns) == 0 {
		return "", false
	}
	conditions := make([]string, 0, len(requiredColumns))
	for _, def := range requiredColumns {
		conditions = append(conditions, buildBootstrapCompatibleColumnCondition(target, def))
	}
	return strings.Join(conditions, "\n    AND "), true
}

func buildBootstrapCompatibleColumnCondition(table string, def tableColumnDef) string {
	base := strings.ToUpper(strings.TrimSpace(def.BaseType))
	if base == "" {
		base, _, _ = parseTypeToken(firstSQLToken(def.RawType))
	}
	boolInt := func(value bool) int {
		if value {
			return 1
		}
		return 0
	}
	conditions := []string{
		fmt.Sprintf("c.object_id = OBJECT_ID(N'%s', N'U')", escapeSQLLiteral(table)),
		fmt.Sprintf("c.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2", escapeSQLLiteral(def.Name)),
		fmt.Sprintf("c.is_computed = %d", boolInt(def.IsComputed)),
		fmt.Sprintf("c.is_identity = %d", boolInt(def.IdentityDefinition != "")),
	}
	if !def.IsComputed {
		switch base {
		case "DECIMAL", "NUMERIC":
			conditions = append(conditions, "UPPER(TYPE_NAME(c.user_type_id)) IN (N'DECIMAL', N'NUMERIC')")
		case "ROWVERSION", "TIMESTAMP":
			conditions = append(conditions, "UPPER(TYPE_NAME(c.user_type_id)) IN (N'ROWVERSION', N'TIMESTAMP')")
		default:
			conditions = append(conditions, fmt.Sprintf("UPPER(TYPE_NAME(c.user_type_id)) = N'%s'", escapeSQLLiteral(base)))
		}
	}
	return "EXISTS (SELECT 1 FROM sys.columns AS c WHERE " + strings.Join(conditions, " AND ") + ")"
}

func buildDefaultConstraintEquivalentConditions(target string, defaults map[string]tableDefaultConstraintDef) []string {
	safeTarget := escapeSQLLiteral(target)
	conditions := []string{
		fmt.Sprintf(`(SELECT COUNT_BIG(*) FROM sys.default_constraints WHERE parent_object_id = OBJECT_ID(N'%s', N'U')) = %d`, safeTarget, len(defaults)),
	}
	columns := make([]string, 0, len(defaults))
	for column := range defaults {
		columns = append(columns, column)
	}
	sort.Strings(columns)
	for _, column := range columns {
		def := defaults[column]
		parts := []string{
			fmt.Sprintf("dc.parent_object_id = OBJECT_ID(N'%s', N'U')", safeTarget),
			fmt.Sprintf("dc.parent_column_id = COLUMNPROPERTY(OBJECT_ID(N'%s', N'U'), N'%s', 'ColumnId')", safeTarget, escapeSQLLiteral(column)),
			buildCatalogSQLTextSemanticCondition("dc.definition", def.Expression),
		}
		if strings.TrimSpace(def.ConstraintName) != "" {
			parts = append(parts, fmt.Sprintf("dc.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2", escapeSQLLiteral(def.ConstraintName)))
		}
		conditions = append(conditions, "EXISTS (SELECT 1 FROM sys.default_constraints AS dc WHERE "+strings.Join(parts, " AND ")+")")
	}
	return conditions
}

func buildTableConstraintEquivalentConditions(target string, items []string) ([]string, bool) {
	type constraintItem struct {
		kind       string
		definition string
	}
	constraints := make([]constraintItem, 0)
	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		if !isTableEquivalenceConstraintItem(line) {
			continue
		}
		kind := constraintKindForTableItem(line)
		if kind == "default" || kind == "constraint" {
			continue
		}
		constraints = append(constraints, constraintItem{kind: kind, definition: line})
	}

	counts := map[string]int{"primary_key": 0, "unique": 0, "check": 0, "foreign_key": 0}
	for _, item := range constraints {
		counts[item.kind]++
	}
	safeTarget := escapeSQLLiteral(target)
	conditions := []string{
		fmt.Sprintf(`(SELECT COUNT_BIG(*) FROM sys.key_constraints WHERE parent_object_id = OBJECT_ID(N'%s', N'U') AND type = N'PK') = %d`, safeTarget, counts["primary_key"]),
		fmt.Sprintf(`(SELECT COUNT_BIG(*) FROM sys.key_constraints WHERE parent_object_id = OBJECT_ID(N'%s', N'U') AND type = N'UQ') = %d`, safeTarget, counts["unique"]),
		fmt.Sprintf(`(SELECT COUNT_BIG(*) FROM sys.check_constraints WHERE parent_object_id = OBJECT_ID(N'%s', N'U')) = %d`, safeTarget, counts["check"]),
		fmt.Sprintf(`(SELECT COUNT_BIG(*) FROM sys.foreign_keys WHERE parent_object_id = OBJECT_ID(N'%s', N'U')) = %d`, safeTarget, counts["foreign_key"]),
	}
	for _, item := range constraints {
		var condition string
		var ok bool
		switch item.kind {
		case "primary_key", "unique":
			condition, ok = buildKeyConstraintEquivalentCondition(target, item.definition, item.kind)
		case "check":
			condition, ok = buildCheckConstraintEquivalentCondition(target, item.definition)
		case "foreign_key":
			condition, ok = buildForeignKeyEquivalentCondition(target, item.definition)
		default:
			return nil, false
		}
		if !ok {
			return nil, false
		}
		conditions = append(conditions, condition)
	}
	return conditions, true
}

func isTableEquivalenceConstraintItem(line string) bool {
	upper := strings.ToUpper(strings.TrimSpace(line))
	if isConstraintLikeTableItem(upper) {
		return true
	}
	for _, keyword := range []string{"PRIMARY KEY", "UNIQUE", "CHECK", "FOREIGN KEY", "REFERENCES"} {
		if len(findAllKeywordPositions(upper, keyword)) > 0 {
			return true
		}
	}
	return false
}

type tableKeyColumnDef struct {
	name       string
	descending bool
}

func buildKeyConstraintEquivalentCondition(target, definition, kind string) (string, bool) {
	columns, ok := extractKeyConstraintColumnDefinitions(definition, kind)
	if !ok || len(columns) == 0 {
		return "", false
	}
	name, named := extractNamedConstraintName(definition)
	name = stripIdentifierBrackets(name)
	expectedType := "UQ"
	if kind == "primary_key" {
		expectedType = "PK"
	}
	upper := strings.ToUpper(definition)
	expectedIndexType := 1
	if strings.Contains(upper, " NONCLUSTERED") {
		expectedIndexType = 2
	}
	safeTarget := escapeSQLLiteral(target)
	parts := []string{
		fmt.Sprintf("kc.parent_object_id = OBJECT_ID(N'%s', N'U')", safeTarget),
		fmt.Sprintf("kc.type = N'%s'", expectedType),
		fmt.Sprintf("i.type = %d", expectedIndexType),
		fmt.Sprintf("(SELECT COUNT_BIG(*) FROM sys.index_columns AS ic WHERE ic.object_id=kc.parent_object_id AND ic.index_id=kc.unique_index_id AND ic.key_ordinal > 0) = %d", len(columns)),
	}
	if named && name != "" {
		parts = append(parts, fmt.Sprintf("kc.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2", escapeSQLLiteral(name)))
	}
	for ordinal, column := range columns {
		desc := 0
		if column.descending {
			desc = 1
		}
		parts = append(parts, fmt.Sprintf(`EXISTS (SELECT 1 FROM sys.index_columns AS ic JOIN sys.columns AS c ON c.object_id=ic.object_id AND c.column_id=ic.column_id WHERE ic.object_id=kc.parent_object_id AND ic.index_id=kc.unique_index_id AND ic.key_ordinal=%d AND ic.is_descending_key=%d AND c.name COLLATE Latin1_General_100_BIN2=N'%s' COLLATE Latin1_General_100_BIN2)`, ordinal+1, desc, escapeSQLLiteral(column.name)))
	}
	return "EXISTS (SELECT 1 FROM sys.key_constraints AS kc JOIN sys.indexes AS i ON i.object_id=kc.parent_object_id AND i.index_id=kc.unique_index_id WHERE " + strings.Join(parts, " AND ") + ")", true
}

func extractKeyConstraintColumnDefinitions(definition, kind string) ([]tableKeyColumnDef, bool) {
	normalized := normalizeSQLText(definition, true, true)
	upper := strings.ToUpper(normalized)
	keyword := "UNIQUE"
	if kind == "primary_key" {
		keyword = "PRIMARY KEY"
	}
	positions := findAllKeywordPositions(upper, keyword)
	if len(positions) == 0 {
		return nil, false
	}
	after := strings.TrimSpace(normalized[positions[0]+len(keyword):])
	afterUpper := strings.ToUpper(after)
	if strings.HasPrefix(afterUpper, "CLUSTERED ") {
		after = strings.TrimSpace(after[len("CLUSTERED "):])
	} else if strings.HasPrefix(afterUpper, "NONCLUSTERED ") {
		after = strings.TrimSpace(after[len("NONCLUSTERED "):])
	}
	if !strings.HasPrefix(after, "(") {
		// Inline PRIMARY KEY/UNIQUE applies to the leading column.
		token := firstSQLIdentifierToken(normalized)
		name := columnIdentifierName(token)
		if name == "" {
			return nil, false
		}
		return []tableKeyColumnDef{{name: name}}, true
	}
	raw, ok := extractLeadingParenthesizedContent(after)
	if !ok {
		return nil, false
	}
	items := splitTopLevelByComma(raw)
	out := make([]tableKeyColumnDef, 0, len(items))
	for _, item := range items {
		token := firstSQLIdentifierToken(strings.TrimSpace(item))
		name := columnIdentifierName(token)
		if name == "" {
			return nil, false
		}
		remainder := strings.ToUpper(strings.TrimSpace(strings.TrimSpace(item)[len(token):]))
		out = append(out, tableKeyColumnDef{name: name, descending: strings.HasPrefix(remainder, "DESC")})
	}
	return out, len(out) > 0
}

func buildCheckConstraintEquivalentCondition(target, definition string) (string, bool) {
	expression, err := extractCheckConstraintExpression(definition)
	if err != nil {
		return "", false
	}
	name, named := extractNamedConstraintName(definition)
	name = stripIdentifierBrackets(name)
	safeTarget := escapeSQLLiteral(target)
	parts := []string{
		fmt.Sprintf("cc.parent_object_id = OBJECT_ID(N'%s', N'U')", safeTarget),
		"cc.is_disabled = 0",
		"cc.is_not_trusted = 0",
		buildCatalogSQLTextSemanticCondition("cc.definition", "("+strings.TrimSpace(expression)+")"),
	}
	if named && name != "" {
		parts = append(parts, fmt.Sprintf("cc.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2", escapeSQLLiteral(name)))
	}
	return "EXISTS (SELECT 1 FROM sys.check_constraints AS cc WHERE " + strings.Join(parts, " AND ") + ")", true
}

func buildForeignKeyEquivalentCondition(target, definition string) (string, bool) {
	fk, err := parseForeignKeyDefinition(definition)
	if err != nil || len(fk.LocalColumns) == 0 {
		return "", false
	}
	name, named := extractNamedConstraintName(definition)
	name = stripIdentifierBrackets(name)
	deleteAction := foreignKeyReferentialAction(definition, "DELETE")
	updateAction := foreignKeyReferentialAction(definition, "UPDATE")
	safeTarget := escapeSQLLiteral(target)
	safeRef := escapeSQLLiteral(fk.ReferencedTable)
	parts := []string{
		fmt.Sprintf("fk.parent_object_id = OBJECT_ID(N'%s', N'U')", safeTarget),
		fmt.Sprintf("fk.referenced_object_id = OBJECT_ID(N'%s', N'U')", safeRef),
		"fk.is_disabled = 0",
		"fk.is_not_trusted = 0",
		fmt.Sprintf("fk.is_not_for_replication = %d", boolToInt(strings.Contains(strings.ToUpper(definition), "NOT FOR REPLICATION"))),
		fmt.Sprintf("fk.delete_referential_action_desc = N'%s'", deleteAction),
		fmt.Sprintf("fk.update_referential_action_desc = N'%s'", updateAction),
		fmt.Sprintf("(SELECT COUNT_BIG(*) FROM sys.foreign_key_columns AS fkc WHERE fkc.constraint_object_id=fk.object_id) = %d", len(fk.LocalColumns)),
	}
	if named && name != "" {
		parts = append(parts, fmt.Sprintf("fk.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2", escapeSQLLiteral(name)))
	}
	for ordinal := range fk.LocalColumns {
		parts = append(parts, fmt.Sprintf(`EXISTS (SELECT 1 FROM sys.foreign_key_columns AS fkc JOIN sys.columns AS pc ON pc.object_id=fkc.parent_object_id AND pc.column_id=fkc.parent_column_id JOIN sys.columns AS rc ON rc.object_id=fkc.referenced_object_id AND rc.column_id=fkc.referenced_column_id WHERE fkc.constraint_object_id=fk.object_id AND fkc.constraint_column_id=%d AND pc.name COLLATE Latin1_General_100_BIN2=N'%s' COLLATE Latin1_General_100_BIN2 AND rc.name COLLATE Latin1_General_100_BIN2=N'%s' COLLATE Latin1_General_100_BIN2)`, ordinal+1, escapeSQLLiteral(fk.LocalColumns[ordinal]), escapeSQLLiteral(fk.ReferencedColumns[ordinal])))
	}
	return "EXISTS (SELECT 1 FROM sys.foreign_keys AS fk WHERE " + strings.Join(parts, " AND ") + ")", true
}

func foreignKeyReferentialAction(definition, action string) string {
	upper := strings.ToUpper(definition)
	prefix := "ON " + action + " "
	pos := strings.Index(upper, prefix)
	if pos < 0 {
		return "NO_ACTION"
	}
	after := strings.TrimSpace(upper[pos+len(prefix):])
	switch {
	case strings.HasPrefix(after, "CASCADE"):
		return "CASCADE"
	case strings.HasPrefix(after, "SET NULL"):
		return "SET_NULL"
	case strings.HasPrefix(after, "SET DEFAULT"):
		return "SET_DEFAULT"
	default:
		return "NO_ACTION"
	}
}

func buildCatalogSQLTextEquivalentCondition(catalogExpression, expected string) string {
	compact := removeSQLWhitespace(expected)
	return fmt.Sprintf(`REPLACE(REPLACE(REPLACE(REPLACE(CONVERT(NVARCHAR(MAX),%s),N' ',N''),NCHAR(9),N''),NCHAR(10),N''),NCHAR(13),N'') COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2`, catalogExpression, escapeSQLLiteral(compact))
}

func buildCatalogSQLTextSemanticCondition(catalogExpression, expected string) string {
	compactCatalog := fmt.Sprintf("REPLACE(REPLACE(REPLACE(REPLACE(CONVERT(NVARCHAR(MAX),%s),N' ',N''),NCHAR(9),N''),NCHAR(10),N''),NCHAR(13),N'')", catalogExpression)
	conditions := []string{
		fmt.Sprintf("LOWER(%s) COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2", compactCatalog, escapeSQLLiteral(strings.ToLower(removeSQLWhitespace(expected)))),
	}
	for _, token := range exactSQLTokensForSemanticComparison(expected) {
		conditions = append(conditions, fmt.Sprintf("CHARINDEX(N'%s' COLLATE Latin1_General_100_BIN2, CONVERT(NVARCHAR(MAX),%s) COLLATE Latin1_General_100_BIN2) > 0", escapeSQLLiteral(token), catalogExpression))
	}
	return "(" + strings.Join(conditions, " AND ") + ")"
}

func exactSQLTokensForSemanticComparison(input string) []string {
	tokens := make([]string, 0)
	seen := map[string]struct{}{}
	add := func(token string) {
		if token == "" {
			return
		}
		if _, exists := seen[token]; exists {
			return
		}
		seen[token] = struct{}{}
		tokens = append(tokens, token)
	}
	for index := 0; index < len(input); {
		switch input[index] {
		case '[':
			start := index
			index++
			for index < len(input) {
				if input[index] == ']' {
					if index+1 < len(input) && input[index+1] == ']' {
						index += 2
						continue
					}
					index++
					add(input[start:index])
					break
				}
				index++
			}
		case '\'':
			start := index
			if start > 0 && (input[start-1] == 'N' || input[start-1] == 'n') {
				start--
			}
			index++
			for index < len(input) {
				if input[index] == '\'' {
					if index+1 < len(input) && input[index+1] == '\'' {
						index += 2
						continue
					}
					index++
					add(input[start:index])
					break
				}
				index++
			}
		default:
			index++
		}
	}
	sort.Strings(tokens)
	return tokens
}

func removeSQLWhitespace(input string) string {
	replacer := strings.NewReplacer(" ", "", "\t", "", "\r", "", "\n", "")
	return replacer.Replace(strings.TrimSpace(input))
}

func boolToInt(value bool) int {
	if value {
		return 1
	}
	return 0
}

func buildIndexWrapperPrecheck(tableName, indexName string, isCreate bool) string {
	safeTable := escapeSQLLiteral(tableName)
	safeIndex := escapeSQLLiteral(indexName)
	targetIndex := tableName + "." + indexName
	safeTargetIndex := escapeSQLLiteral(targetIndex)

	if isCreate {
		return fmt.Sprintf(`IF EXISTS (
  SELECT 1
  FROM sys.indexes
  WHERE object_id = OBJECT_ID(N'%s')
    AND name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
)
BEGIN
  ;THROW 51000, N'Schema drift: CREATE INDEX target already exists; equivalence is not proven', 1;
END`, safeTable, safeIndex)
	}

	return fmt.Sprintf(`IF NOT EXISTS (
  SELECT 1
  FROM sys.indexes
  WHERE object_id = OBJECT_ID(N'%s')
    AND name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
)
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = N'%s',
      @reason = N'OBJECT_NOT_EXISTS',
      @message = N'Precheck: index already absent',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END`, safeTable, safeIndex, safeTargetIndex)
}

func buildIndexWrapperPrecheckMany(tableName string, indexNames []string, isCreate bool) string {
	safeTable := escapeSQLLiteral(tableName)
	normalized := make([]string, 0, len(indexNames))
	seen := make(map[string]struct{}, len(indexNames))
	for _, name := range indexNames {
		clean := stripIdentifierBrackets(cleanIdentifierToken(name))
		if clean == "" {
			continue
		}
		key := clean
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		normalized = append(normalized, clean)
	}
	if len(normalized) == 0 {
		return `-- precheck skipped: no target indexes found for wrapper-level check`
	}
	sort.Strings(normalized)

	valueRows := make([]string, 0, len(normalized))
	for _, name := range normalized {
		valueRows = append(valueRows, fmt.Sprintf("(N'%s')", escapeSQLLiteral(name)))
	}
	valuesSQL := strings.Join(valueRows, ", ")

	if isCreate {
		return fmt.Sprintf(`IF OBJECT_ID(N'%s', N'U') IS NOT NULL
AND NOT EXISTS (
  SELECT 1
  FROM (VALUES %s) AS i(index_name)
  WHERE NOT EXISTS (
    SELECT 1
    FROM sys.indexes s
    WHERE s.object_id = OBJECT_ID(N'%s')
      AND s.name = i.index_name
  )
)
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'INDEX_EXISTS',
      @message = N'Precheck: index(es) already exist',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END`, safeTable, valuesSQL, safeTable)
	}

	return fmt.Sprintf(`IF OBJECT_ID(N'%s', N'U') IS NULL
OR NOT EXISTS (
  SELECT 1
  FROM (VALUES %s) AS i(index_name)
  WHERE EXISTS (
    SELECT 1
    FROM sys.indexes s
    WHERE s.object_id = OBJECT_ID(N'%s')
      AND s.name = i.index_name
  )
)
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = @action,
      @target = @target,
      @reason = N'OBJECT_NOT_EXISTS',
      @message = N'Precheck: index(es) already absent',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  RETURN;
END`, safeTable, valuesSQL, safeTable)
}

func escapeSQLLiteral(input string) string {
	return strings.ReplaceAll(input, "'", "''")
}

func buildStepAuditLogSQL(decision string, severity int, action, target, reason, message string) string {
	return normalizeSQLText(fmt.Sprintf(`IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
BEGIN
  EXEC dbo.usp_mig_log_event
    @run_key = @run_key,
    @version = @version,
    @category = @category,
    @file_name = @file,
    @decision = N'%s',
    @severity = %d,
    @action = N'%s',
    @target = N'%s',
    @reason = N'%s',
    @message = N'%s',
    @detail = @direction,
    @error_number = NULL,
    @error_line = NULL,
    @error_procedure = NULL;
END`, escapeSQLLiteral(decision), severity, escapeSQLLiteral(action), escapeSQLLiteral(target), escapeSQLLiteral(reason), escapeSQLLiteral(message)), true, true)
}

func auditActionForWrapper(category string, isUp bool, body string) string {
	if isManualDropTableRequestMigration(body) {
		return "followup_required"
	}

	upperBody := strings.ToUpper(body)
	switch category {
	case "types":
		if strings.Contains(upperBody, "DROP TYPE") && strings.Contains(upperBody, "CREATE TYPE") {
			return "recreate_type"
		}
		if strings.Contains(upperBody, "DROP TYPE") {
			return "drop_type"
		}
		return "create_type"
	case "tables":
		if strings.Contains(upperBody, "N'ADD_DEFAULT'") ||
			(strings.Contains(upperBody, "ADD CONSTRAINT") && strings.Contains(upperBody, " DEFAULT ")) {
			return "add_default"
		}
		if strings.Contains(upperBody, "N'DROP_DEFAULT'") ||
			(strings.Contains(upperBody, "FROM SYS.DEFAULT_CONSTRAINTS") && strings.Contains(upperBody, "DROP CONSTRAINT")) {
			return "drop_default"
		}
		if strings.Contains(upperBody, "DROP COLUMN") {
			return "drop_column"
		}
		if strings.Contains(upperBody, "CREATE TABLE") {
			return "create_table"
		}
		if strings.Contains(upperBody, "ALTER TABLE") {
			return "alter_column"
		}
		if strings.Contains(upperBody, "DROP TABLE") {
			return "drop_table"
		}
		return "alter_column"
	case "indexes":
		if strings.Contains(upperBody, "DROP INDEX") {
			return "drop_index"
		}
		return "add_index"
	case "views":
		if strings.Contains(upperBody, "DROP VIEW") {
			return "drop_view"
		}
		return "create_or_alter_view"
	case "functions":
		if strings.Contains(upperBody, "DROP FUNCTION") {
			return "drop_function"
		}
		return "create_or_alter_function"
	case "procedures":
		if strings.Contains(upperBody, "DROP PROCEDURE") {
			return "drop_procedure"
		}
		return "create_or_alter_procedure"
	case "patch":
		return "backfill"
	default:
		return "followup_required"
	}
}

func indentSQL(sql, prefix string) string {
	lines := strings.Split(normalizeSQLText(sql, true, true), "\n")
	for i, line := range lines {
		if strings.TrimSpace(line) == "" {
			lines[i] = prefix
			continue
		}
		lines[i] = prefix + line
	}
	return strings.Join(lines, "\n")
}

func compactGeneratedSQLWhitespace(sql string) string {
	normalized := normalizeSQLText(sql, true, false)
	lines := strings.Split(normalized, "\n")
	out := make([]string, 0, len(lines))
	previousBlank := false
	for _, line := range lines {
		blank := strings.TrimSpace(line) == ""
		if blank && previousBlank {
			continue
		}
		out = append(out, line)
		previousBlank = blank
	}
	for len(out) > 0 && strings.TrimSpace(out[0]) == "" {
		out = out[1:]
	}
	for len(out) > 0 && strings.TrimSpace(out[len(out)-1]) == "" {
		out = out[:len(out)-1]
	}
	if len(out) == 0 {
		return ""
	}
	return strings.Join(out, "\n") + "\n"
}

// legacySemanticFromMigration is retained temporarily while lower-level table
// mutation helpers are progressively moved out of this legacy generator file.
func legacySemanticFromMigration(meta migrationMeta, sql string) (string, error) {
	return semanticFromMigrationWithState(meta, sql, "", false)
}

func legacySemanticFromMigrationWithState(meta migrationMeta, sql, currentSemantic string, hasCurrent bool) (string, error) {
	switch meta.Category {
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

		working := currentSemantic
		applied := false

		removedCols, dropErr := extractDroppedColumns(meta.Object, sql)
		if dropErr == nil && len(removedCols) > 0 {
			if !hasCurrent || strings.TrimSpace(currentSemantic) == "" {
				return "", fmt.Errorf("drop-column migration requires prior table state for %s", meta.Object)
			}
			updated, applyErr := applyDropColumnsToCreateTable(working, removedCols)
			if applyErr != nil {
				return "", applyErr
			}
			working = updated
			applied = true
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
			working = updated
			applied = true
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

func legacySemanticFromSource(meta migrationMeta, sql string) (string, error) {
	switch meta.Category {
	case "tables":
		stmt, err := extractCreateTableStatement(meta.Object, sql)
		if err != nil {
			return "", err
		}
		stmt, _ = repairCreateTableCommaSyntax(stmt)
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
		return normalizeSQLText(stmt, true, true), nil
	default:
		return "", fmt.Errorf("unsupported source category %q", meta.Category)
	}
}

func canonicalizeCreateTableSourceStatement(stmt string) (string, error) {
	items, err := extractCreateTableItems(stmt)
	if err != nil {
		return "", err
	}

	columnItems := make([]string, 0, len(items))
	constraintItems := make([]string, 0, len(items))
	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		leadingTrivia, sqlItem := splitLeadingTableItemComments(line)
		if strings.TrimSpace(sqlItem) == "" {
			return "", fmt.Errorf("CREATE TABLE item contains comments but no column or constraint definition")
		}

		upper := strings.ToUpper(strings.TrimSpace(sqlparse.StripSQLStringsAndComments(sqlItem)))
		if isConstraintLikeTableItem(upper) || strings.HasPrefix(upper, "DEFAULT ") {
			constraintItems = append(constraintItems, attachLeadingTableItemComments(leadingTrivia, strings.TrimSpace(sqlItem)))
			continue
		}

		columnItems = append(columnItems, attachLeadingTableItemComments(leadingTrivia, canonicalizeTableColumnItemStyle(sqlItem)))
	}

	rebuiltItems := append(columnItems, constraintItems...)
	return rebuildCreateTableStatementWithItems(stmt, rebuiltItems)
}

// splitLeadingTableItemComments separates comments attached to the next table
// item from the SQL definition itself. Without this distinction a line such as
// "-- id remark\nid INT" is parsed as a column named "--".
func splitLeadingTableItemComments(input string) (string, string) {
	i := 0
	consumedComment := false
	for i < len(input) {
		for i < len(input) {
			switch input[i] {
			case ' ', '\t', '\r', '\n':
				i++
			default:
				goto trivia
			}
		}

	trivia:
		if i+1 >= len(input) {
			break
		}
		switch {
		case input[i] == '-' && input[i+1] == '-':
			consumedComment = true
			i += 2
			for i < len(input) && input[i] != '\r' && input[i] != '\n' {
				i++
			}
		case input[i] == '/' && input[i+1] == '*':
			consumedComment = true
			depth := 1
			i += 2
			for i < len(input) && depth > 0 {
				if i+1 < len(input) && input[i] == '/' && input[i+1] == '*' {
					depth++
					i += 2
					continue
				}
				if i+1 < len(input) && input[i] == '*' && input[i+1] == '/' {
					depth--
					i += 2
					continue
				}
				i++
			}
		default:
			if !consumedComment {
				return "", input
			}
			return strings.TrimSpace(input[:i]), strings.TrimSpace(input[i:])
		}
	}
	if !consumedComment {
		return "", input
	}
	return strings.TrimSpace(input), ""
}

func attachLeadingTableItemComments(comments, sqlItem string) string {
	comments = strings.TrimSpace(comments)
	sqlItem = strings.TrimSpace(sqlItem)
	if comments == "" {
		return sqlItem
	}
	if sqlItem == "" {
		return comments
	}
	// rebuildCreateTableStatementWithItems adds indentation to the first line.
	// Indent continuation lines here so comments and their SQL item remain
	// aligned in the generated CREATE TABLE statement.
	comments = strings.ReplaceAll(comments, "\n", "\n  ")
	return comments + "\n  " + sqlItem
}

func canonicalizeTableColumnItemStyle(line string) string {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return line
	}

	nameToken, consumed := consumeLeadingColumnIdentifierToken(trimmed)
	if nameToken == "" {
		return line
	}
	// Preserve whitespace inside bracketed identifiers, including a single-space
	// column name (`[ ]`), which is valid SQL Server syntax.
	name := columnIdentifierName(nameToken)
	if name == "" {
		return line
	}

	rest := strings.TrimSpace(trimmed[consumed:])
	if rest == "" {
		return quoteIdentifier(name) + " NULL"
	}
	if hasKeywordAtWordBoundary(strings.ToUpper(rest), 0, "AS") {
		// Computed columns must retain their original suffix; appending NULL is invalid.
		return strings.TrimSpace(quoteIdentifier(name) + " " + rest)
	}

	upperRest := strings.ToUpper(rest)
	if len(findAllKeywordPositions(upperRest, "NULL")) == 0 {
		rest = strings.TrimSpace(rest) + " NULL"
	}
	return strings.TrimSpace(quoteIdentifier(name) + " " + rest)
}

func consumeLeadingColumnIdentifierToken(input string) (string, int) {
	if input == "" {
		return "", 0
	}
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return "", 0
	}
	if trimmed[0] != '[' {
		token := firstSQLToken(trimmed)
		return token, len(token)
	}

	i := 1
	for i < len(trimmed) {
		if trimmed[i] == ']' {
			if i+1 < len(trimmed) && trimmed[i+1] == ']' {
				i += 2
				continue
			}
			return trimmed[:i+1], i + 1
		}
		i++
	}
	return firstSQLToken(trimmed), len(firstSQLToken(trimmed))
}

func columnIdentifierName(token string) string {
	trimmed := strings.TrimSpace(token)
	if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") && len(trimmed) >= 2 {
		inner := trimmed[1 : len(trimmed)-1]
		return strings.ReplaceAll(inner, "]]", "]")
	}
	return strings.TrimSpace(cleanIdentifierToken(trimmed))
}

func extractCreateTableStatement(objectName, sql string) (string, error) {
	stmt, _, _, err := extractCreateTableStatementBounds(objectName, sql)
	return stmt, err
}

func extractCreateTableStatementBounds(objectName, sql string) (string, int, int, error) {
	normalized := normalizeSQLText(sql, true, true)
	upper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(normalized))
	const keyword = "CREATE TABLE "
	start := -1
	consumedTarget := 0
	for cursor := 0; cursor < len(normalized); {
		rel := strings.Index(upper[cursor:], keyword)
		if rel < 0 {
			break
		}
		candidateStart := cursor + rel
		target, consumed := consumeLeadingObjectToken(normalized[candidateStart+len(keyword):])
		if target != "" && sameQualifiedObject(target, objectName) {
			start = candidateStart
			consumedTarget = consumed
			break
		}
		cursor = candidateStart + len(keyword)
	}
	if start < 0 {
		return "", 0, 0, fmt.Errorf("table statement for %s not found", objectName)
	}

	// Opening "(" must appear right after CREATE TABLE <object> (ignoring whitespace).
	// This prevents false-positive parsing when invalid syntax uses "{...}" while
	// datatype tokens still contain parentheses (e.g. VARCHAR(2)).
	open := start + len(keyword) + consumedTarget
	for open < len(normalized) {
		ch := normalized[open]
		if ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
			open++
			continue
		}
		break
	}
	if open >= len(normalized) || normalized[open] != '(' {
		return "", 0, 0, fmt.Errorf("invalid CREATE TABLE statement for %s", objectName)
	}

	close, ok := sqlparse.FindMatchingParen(normalized, open)
	if !ok {
		return "", 0, 0, fmt.Errorf("unterminated CREATE TABLE statement for %s", objectName)
	}
	end := close + 1

	consumeTrailingWS := func(idx int) int {
		for idx < len(normalized) && (normalized[idx] == ' ' || normalized[idx] == '\t' || normalized[idx] == '\n' || normalized[idx] == '\r') {
			idx++
		}
		return idx
	}

	end = consumeTrailingWS(end)
	if end < len(normalized) && normalized[end] == ';' {
		end++
		// Also consume whitespace after semicolon so replacement does not keep
		// accumulating trailing blank lines on repeated fmt runs.
		end = consumeTrailingWS(end)
	}

	return strings.TrimSpace(normalized[start:end]), start, end, nil
}

func createTableAnchors(objectName string) []string {
	trimmed := strings.TrimSpace(objectName)
	if trimmed == "" {
		return []string{}
	}

	anchors := []string{"CREATE TABLE " + trimmed}
	parts := strings.Split(trimmed, ".")
	if len(parts) == 2 {
		schema := stripIdentifierBrackets(strings.TrimSpace(parts[0]))
		object := stripIdentifierBrackets(strings.TrimSpace(parts[1]))
		if schema != "" && object != "" {
			anchors = append(anchors,
				"CREATE TABLE ["+escapeBracketIdentifier(schema)+"].["+escapeBracketIdentifier(object)+"]",
				"CREATE TABLE "+schema+".["+escapeBracketIdentifier(object)+"]",
				"CREATE TABLE ["+escapeBracketIdentifier(schema)+"]."+object,
			)
		}
	}

	out := make([]string, 0, len(anchors))
	seen := make(map[string]struct{}, len(anchors))
	for _, anchor := range anchors {
		key := strings.ToUpper(anchor)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		out = append(out, anchor)
	}
	return out
}

func escapeBracketIdentifier(input string) string {
	return strings.ReplaceAll(input, "]", "]]")
}

func extractCreateIndexStatements(objectName, sql string) ([]string, error) {
	stmts := extractCreateIndexStatementsAllowEmpty(objectName, sql)
	if len(stmts) == 0 {
		return nil, fmt.Errorf("index statement for %s not found", objectName)
	}
	return stmts, nil
}

func extractCreateIndexStatementsAllowEmpty(objectName, sql string) []string {
	normalized := normalizeSQLText(sql, true, true)
	semantic := sqlparse.StripSQLStringsAndComments(normalized)
	upper := strings.ToUpper(semantic)
	expectedObject := stripIdentifierBrackets(objectName)

	positions := findCreateIndexStatementPositions(upper)
	if len(positions) == 0 {
		return []string{}
	}
	sort.Ints(positions)

	type seenPos struct {
		Position int
	}
	seen := make(map[int]seenPos)
	stmts := make([]string, 0)
	for positionIndex, pos := range positions {
		if _, ok := seen[pos]; ok {
			continue
		}
		seen[pos] = seenPos{Position: pos}
		end := sqlparse.FindStatementEnd(normalized, pos)
		// SSMS/extractor snapshots do not always include semicolons between
		// consecutive CREATE INDEX statements. Never let one extracted statement
		// consume the next statement and then serialize it twice.
		if positionIndex+1 < len(positions) && positions[positionIndex+1] < end {
			end = positions[positionIndex+1]
		}

		stmt := strings.TrimSpace(normalized[pos:end])
		if stmt == "" {
			continue
		}
		if !indexStatementTargetsObject(stmt, expectedObject) {
			continue
		}
		stmts = append(stmts, stmt)
	}
	return stmts
}

func validateIndexSourceTargets(objectName, sql string) error {
	normalized := normalizeSQLText(sql, true, true)
	semantic := sqlparse.StripSQLStringsAndComments(normalized)
	upperSemantic := strings.ToUpper(semantic)
	positions := findCreateIndexStatementPositions(upperSemantic)
	if len(positions) == 0 {
		return nil
	}

	expected := stripIdentifierBrackets(objectName)
	for positionIndex, pos := range positions {
		end := sqlparse.FindStatementEnd(normalized, pos)
		if positionIndex+1 < len(positions) && positions[positionIndex+1] < end {
			end = positions[positionIndex+1]
		}
		stmt := strings.TrimSpace(normalized[pos:end])
		if stmt == "" {
			continue
		}
		target, err := parseCreateIndexTargetTable(stmt)
		if err != nil {
			return fmt.Errorf("index source invalid for %s: %w", objectName, err)
		}
		if !indexTargetMatchesObject(target, expected) {
			return fmt.Errorf("index source invalid for %s: CREATE INDEX targets %s", objectName, target)
		}
	}
	return nil
}

func parseCreateIndexTargetTable(stmt string) (string, error) {
	normalized := normalizeSQLText(stmt, true, true)
	upper := strings.ToUpper(maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(normalized)))
	onPositions := findAllKeywordPositions(upper, "ON")
	if len(onPositions) == 0 {
		return "", fmt.Errorf("missing ON target in CREATE INDEX statement")
	}
	onPos := onPositions[0]
	afterOn := strings.TrimSpace(normalized[onPos+len("ON"):])
	targetToken, _ := consumeLeadingObjectToken(afterOn)
	target := stripIdentifierBrackets(cleanIdentifierToken(targetToken))
	if target == "" {
		return "", fmt.Errorf("invalid ON target in CREATE INDEX statement")
	}
	return target, nil
}

func indexStatementTargetsObject(stmt, expectedObject string) bool {
	target, err := parseCreateIndexTargetTable(stmt)
	return err == nil && indexTargetMatchesObject(target, expectedObject)
}

// Index snapshots are owned by dbo.<Table>.sql. SQL Server permits an
// unqualified table name, so treat it as dbo only for that snapshot contract.
func indexTargetMatchesObject(target, expected string) bool {
	target = stripIdentifierBrackets(target)
	expected = stripIdentifierBrackets(expected)
	if target == expected {
		return true
	}
	return strings.HasPrefix(expected, "dbo.") && target == strings.TrimPrefix(expected, "dbo.")
}

type indexStatementDef struct {
	Name      string
	Statement string
}

func parseIndexStateMap(objectName, sql string) (map[string]indexStatementDef, error) {
	out := make(map[string]indexStatementDef)
	stmts := extractCreateIndexStatementsAllowEmpty(objectName, sql)
	for _, stmt := range stmts {
		if err := validateIndexSubsetSupport(stmt); err != nil {
			return nil, fmt.Errorf("invalid CREATE INDEX for %s: %w", objectName, err)
		}
		indexName, err := parseCreateIndexName(stmt)
		if err != nil {
			return nil, err
		}
		if _, _, err := parseCreateIndexKeysAndFilter(stmt); err != nil {
			return nil, fmt.Errorf("invalid CREATE INDEX %s for %s: %w", indexName, objectName, err)
		}
		if _, err := parseCreateIndexIncludeColumns(stmt); err != nil {
			return nil, fmt.Errorf("invalid CREATE INDEX %s for %s: %w", indexName, objectName, err)
		}
		clean := stripIdentifierBrackets(cleanIdentifierToken(indexName))
		if clean == "" {
			return nil, fmt.Errorf("invalid CREATE INDEX statement: empty index name")
		}
		key := clean
		if existing, exists := out[key]; exists {
			return nil, fmt.Errorf("duplicate index name %s for %s (definitions: %s | %s)", clean, objectName, existing.Statement, stmt)
		}
		out[key] = indexStatementDef{
			Name:      clean,
			Statement: normalizeIndexStatement(qualifyIndexTargetToObject(stmt, objectName)),
		}
	}
	return out, nil
}

func normalizeIndexStatement(stmt string) string {
	normalized := strings.TrimSpace(normalizeSQLText(stmt, true, true))
	if normalized == "" {
		return ""
	}
	normalized = strings.TrimRight(normalized, " \t\r\n;")
	if normalized == "" {
		return ""
	}
	return normalized + ";"
}

func qualifyIndexTargetToObject(stmt, objectName string) string {
	normalized := normalizeSQLText(stmt, true, true)
	target, err := parseCreateIndexTargetTable(normalized)
	if err != nil || !indexTargetMatchesObject(target, objectName) || strings.Contains(target, ".") {
		return normalized
	}

	onPositions := findAllKeywordPositions(strings.ToUpper(maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(normalized))), "ON")
	if len(onPositions) == 0 {
		return normalized
	}
	afterOn := onPositions[0] + len("ON")
	targetStart := afterOn
	for targetStart < len(normalized) && (normalized[targetStart] == ' ' || normalized[targetStart] == '\t' || normalized[targetStart] == '\r' || normalized[targetStart] == '\n') {
		targetStart++
	}
	_, consumed := consumeLeadingObjectToken(normalized[targetStart:])
	if consumed == 0 {
		return normalized
	}
	return normalized[:targetStart] + objectName + normalized[targetStart+consumed:]
}

func serializeIndexState(indexes map[string]indexStatementDef) string {
	if len(indexes) == 0 {
		return ""
	}
	keys := make([]string, 0, len(indexes))
	for key := range indexes {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	lines := make([]string, 0, len(keys))
	for _, key := range keys {
		line := normalizeIndexStatement(indexes[key].Statement)
		if line == "" {
			continue
		}
		lines = append(lines, line)
	}
	return normalizeSQLText(strings.Join(lines, "\n"), true, true)
}

func sortIndexStateKeys(indexes map[string]indexStatementDef, keys []string, create bool) {
	rank := func(statement string) int {
		upper := strings.ToUpper(statement)
		switch {
		case strings.Contains(upper, "CREATE PRIMARY XML INDEX"):
			if create {
				return 0
			}
			return 2
		case strings.Contains(upper, "CREATE XML INDEX") && strings.Contains(upper, "USING XML INDEX"):
			if create {
				return 1
			}
			return 0
		default:
			return 1
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		ri, rj := rank(indexes[keys[i]].Statement), rank(indexes[keys[j]].Statement)
		if ri != rj {
			return ri < rj
		}
		return keys[i] < keys[j]
	})
}

func buildDropIndexesSQLFromState(objectName, semantic string) (string, error) {
	indexes, err := parseIndexStateMap(objectName, semantic)
	if err != nil {
		return "", err
	}
	if len(indexes) == 0 {
		return "", nil
	}

	keys := make([]string, 0, len(indexes))
	for key := range indexes {
		keys = append(keys, key)
	}
	sortIndexStateKeys(indexes, keys, false)

	lines := make([]string, 0, len(keys))
	for _, key := range keys {
		name := indexes[key].Name
		lines = append(lines, buildDropIndexStatement(objectName, name))
	}
	return normalizeSQLText(strings.Join(lines, "\n"), true, true), nil
}

func buildCreateIndexesSQLFromState(objectName, semantic string) (string, error) {
	indexes, err := parseIndexStateMap(objectName, semantic)
	if err != nil {
		return "", err
	}
	if len(indexes) == 0 {
		return "", nil
	}

	keys := make([]string, 0, len(indexes))
	for key := range indexes {
		keys = append(keys, key)
	}
	sortIndexStateKeys(indexes, keys, true)

	lines := make([]string, 0, len(keys))
	for _, key := range keys {
		def := indexes[key]
		lines = append(lines, buildCreateIndexIfNotExistsStatement(objectName, def.Name, def.Statement))
	}
	return normalizeSQLText(strings.Join(lines, "\n"), true, true), nil
}

func buildDropIndexStatement(objectName, indexName string) string {
	cleanName := stripIdentifierBrackets(cleanIdentifierToken(indexName))
	safeObject := escapeSQLLiteral(objectName)
	safeIndex := escapeSQLLiteral(cleanName)
	target := objectName + "." + cleanName
	runLog := indentSQL(buildStepAuditLogSQL("run", 1, "drop_index", target, "OBJECT_EXISTS", "drop index start"), "  ")
	okLog := indentSQL(buildStepAuditLogSQL("ok", 1, "drop_index", target, "DROPPED", "drop index ok"), "  ")
	skipLog := indentSQL(buildStepAuditLogSQL("skip", 0, "drop_index", target, "INDEX_NOT_EXISTS", "Precheck: index already absent"), "  ")
	return normalizeSQLText(fmt.Sprintf(`-- check index exists before DROP INDEX
IF EXISTS (
  SELECT 1
  FROM sys.indexes
  WHERE object_id = OBJECT_ID(N'%s')
    AND name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
)
BEGIN
%s
  DROP INDEX IF EXISTS %s ON %s;
%s
END
ELSE
BEGIN
%s
END`, safeObject, safeIndex,
		runLog,
		quoteIdentifier(cleanName), quoteQualifiedIdentifier(objectName),
		okLog,
		skipLog), true, true)
}

func buildIndexLifecycleSQL(previousSemantic, nextSemantic, objectName string) (string, string, error) {
	prevMap, err := parseIndexStateMap(objectName, previousSemantic)
	if err != nil {
		return "", "", err
	}
	nextMap, err := parseIndexStateMap(objectName, nextSemantic)
	if err != nil {
		return "", "", err
	}

	upDrops := make([]string, 0)
	upCreates := make([]string, 0)
	downDrops := make([]string, 0)
	downCreates := make([]string, 0)

	for key, prev := range prevMap {
		next, ok := nextMap[key]
		if !ok || normalizeForSemanticHash(prev.Statement) != normalizeForSemanticHash(next.Statement) {
			upDrops = append(upDrops, key)
			downCreates = append(downCreates, key)
		}
	}
	for key, next := range nextMap {
		prev, ok := prevMap[key]
		if !ok || normalizeForSemanticHash(prev.Statement) != normalizeForSemanticHash(next.Statement) {
			upCreates = append(upCreates, key)
			downDrops = append(downDrops, key)
		}
	}

	sortIndexStateKeys(prevMap, upDrops, false)
	sortIndexStateKeys(nextMap, upCreates, true)
	sortIndexStateKeys(nextMap, downDrops, false)
	sortIndexStateKeys(prevMap, downCreates, true)

	upLines := make([]string, 0, len(upDrops)+len(upCreates))
	for _, key := range upDrops {
		upLines = append(upLines, buildDropIndexStatement(objectName, prevMap[key].Name))
	}
	for _, key := range upCreates {
		upLines = append(upLines, buildCreateIndexIfNotExistsStatement(objectName, nextMap[key].Name, nextMap[key].Statement))
	}

	downLines := make([]string, 0, len(downDrops)+len(downCreates))
	for _, key := range downDrops {
		downLines = append(downLines, buildDropIndexStatement(objectName, nextMap[key].Name))
	}
	for _, key := range downCreates {
		downLines = append(downLines, buildCreateIndexIfNotExistsStatement(objectName, prevMap[key].Name, prevMap[key].Statement))
	}

	upSQL := normalizeSQLText(strings.Join(upLines, "\n"), true, true)
	downSQL := normalizeSQLText(strings.Join(downLines, "\n"), true, true)
	return upSQL, downSQL, nil
}

func buildCreateIndexIfNotExistsStatement(tableName, indexName, createStmt string) string {
	safeTable := escapeSQLLiteral(tableName)
	cleanIndex := stripIdentifierBrackets(cleanIdentifierToken(indexName))
	safeIndex := escapeSQLLiteral(cleanIndex)
	target := tableName + "." + cleanIndex
	stmt := normalizeIndexStatement(createStmt)
	if strings.TrimSpace(stmt) == "" || cleanIndex == "" {
		return ""
	}
	runLog := indentSQL(buildStepAuditLogSQL("run", 1, "add_index", target, "OBJECT_EXISTS", "create index start"), "  ")
	okLog := indentSQL(buildStepAuditLogSQL("ok", 1, "add_index", target, "CREATED", "create index ok"), "  ")
	indentedStmt := indentSQL(stmt, "  ")
	columns, _, _ := parseCreateIndexKeysAndFilter(stmt)
	includeColumns, _ := parseCreateIndexIncludeColumns(stmt)
	columns = append(columns, includeColumns...)
	uniquePrecheck := buildUniqueIndexDataPrecheck(tableName, cleanIndex, stmt)
	equivalentPredicate, equivalenceSupported := buildExistingIndexEquivalentPredicate(tableName, cleanIndex, stmt)
	existingIndexSQL := `;THROW 51000, N'Schema drift: existing index definition was not proven equivalent', 1;`
	if equivalenceSupported {
		skipLog := indentSQL(buildStepAuditLogSQL("skip", 0, "add_index", target, "INDEX_EXISTS", "Precheck: equivalent index already exists"), "  ")
		existingIndexSQL = fmt.Sprintf(`IF %s
BEGIN
%s
END
ELSE
BEGIN
  ;THROW 51000, N'Schema drift: existing index differs from the requested definition', 1;
END`, equivalentPredicate, skipLog)
	}
	return normalizeSQLText(fmt.Sprintf(`%s

%s

-- check index does not exist before CREATE INDEX
IF NOT EXISTS (
  SELECT 1
  FROM sys.indexes
  WHERE object_id = OBJECT_ID(N'%s')
    AND name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
)
BEGIN
%s
%s
%s
%s
END
ELSE
BEGIN
%s
END`, buildRequireTableExistsSQL(tableName, "CREATE INDEX"), buildRequireColumnsExistSQL(tableName, columns, "CREATE INDEX"), safeTable, safeIndex,
		runLog,
		indentSQL(uniquePrecheck, "  "),
		indentedStmt,
		okLog,
		indentSQL(existingIndexSQL, "  ")), true, true)
}

type indexKeyDefinition struct {
	Name       string
	Descending bool
}

func buildExistingIndexEquivalentPredicate(tableName, indexName, stmt string) (string, bool) {
	upper := strings.ToUpper(strings.TrimSpace(stmt))
	// XML, spatial and columnstore indexes expose different catalog semantics.
	// Until each is modelled completely, an existing object remains a hard
	// drift failure instead of an unsafe name-only skip.
	if strings.Contains(upper, " XML INDEX ") ||
		strings.HasPrefix(upper, "CREATE XML INDEX") ||
		strings.Contains(upper, " SPATIAL INDEX ") ||
		strings.Contains(upper, " COLUMNSTORE INDEX ") {
		return "", false
	}
	// Physical placement is operationally significant. Until data-space and
	// partition-scheme catalog comparison is modelled, never accept an
	// existing index with an explicit placement clause as equivalent.
	if hasExplicitIndexDataSpaceClause(stmt) {
		return "", false
	}

	keys, err := parseCreateIndexKeyDefinitions(stmt)
	if err != nil || len(keys) == 0 {
		return "", false
	}
	includes, err := parseCreateIndexIncludeColumns(stmt)
	if err != nil {
		return "", false
	}
	_, filter, err := parseCreateIndexKeysAndFilter(stmt)
	if err != nil {
		return "", false
	}

	indexType := 2 // NONCLUSTERED
	if strings.Contains(upper, " CLUSTERED INDEX ") && !strings.Contains(upper, " NONCLUSTERED INDEX ") {
		indexType = 1
	}
	isUnique := 0
	if strings.HasPrefix(upper, "CREATE UNIQUE ") {
		isUnique = 1
	}
	optionConditions, optionsSupported := buildIndexPersistedOptionConditions(stmt)
	if !optionsSupported {
		return "", false
	}

	keyRows := make([]string, 0, len(keys))
	for ordinal, key := range keys {
		desc := 0
		if key.Descending {
			desc = 1
		}
		keyRows = append(keyRows, fmt.Sprintf("(%d, N'%s', %d)", ordinal+1, escapeSQLLiteral(key.Name), desc))
	}
	includeRows := make([]string, 0, len(includes))
	for _, include := range includes {
		includeRows = append(includeRows, fmt.Sprintf("(N'%s')", escapeSQLLiteral(include)))
	}

	filterCheck := "i.has_filter = 0"
	if strings.TrimSpace(filter) != "" {
		filterCheck = fmt.Sprintf(`i.has_filter = 1
    AND REPLACE(REPLACE(REPLACE(REPLACE(i.filter_definition, N' ', N''), CHAR(9), N''), CHAR(10), N''), CHAR(13), N'')
        COLLATE Latin1_General_100_BIN2 =
        REPLACE(REPLACE(REPLACE(REPLACE(N'%s', N' ', N''), CHAR(9), N''), CHAR(10), N''), CHAR(13), N'')
        COLLATE Latin1_General_100_BIN2`, escapeSQLLiteral(filter))
	}

	includeCheck := `AND NOT EXISTS (
      SELECT 1
      FROM sys.index_columns ic
      WHERE ic.object_id = i.object_id
        AND ic.index_id = i.index_id
        AND ic.is_included_column = 1
    )`
	if len(includeRows) > 0 {
		includeCheck = fmt.Sprintf(`AND (
      SELECT COUNT_BIG(*)
      FROM sys.index_columns ic
      WHERE ic.object_id = i.object_id
        AND ic.index_id = i.index_id
        AND ic.is_included_column = 1
    ) = %d
    AND NOT EXISTS (
      SELECT 1
      FROM (VALUES %s) expected(column_name)
      WHERE NOT EXISTS (
        SELECT 1
        FROM sys.index_columns ic
        JOIN sys.columns c
          ON c.object_id = ic.object_id
         AND c.column_id = ic.column_id
        WHERE ic.object_id = i.object_id
          AND ic.index_id = i.index_id
          AND ic.is_included_column = 1
          AND c.name COLLATE Latin1_General_100_BIN2 = expected.column_name COLLATE Latin1_General_100_BIN2
      )
    )`, len(includeRows), strings.Join(includeRows, ", "))
	}

	return fmt.Sprintf(`EXISTS (
  SELECT 1
  FROM sys.indexes i
  WHERE i.object_id = OBJECT_ID(N'%s')
    AND i.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
    AND i.type = %d
    AND i.is_unique = %d
    AND i.is_disabled = 0
    AND i.is_hypothetical = 0
    AND %s
    AND %s
    AND (
      SELECT COUNT_BIG(*)
      FROM sys.index_columns ic
      WHERE ic.object_id = i.object_id
        AND ic.index_id = i.index_id
        AND ic.key_ordinal > 0
    ) = %d
    AND NOT EXISTS (
      SELECT 1
      FROM (VALUES %s) expected(key_ordinal, column_name, is_descending_key)
      WHERE NOT EXISTS (
        SELECT 1
        FROM sys.index_columns ic
        JOIN sys.columns c
          ON c.object_id = ic.object_id
         AND c.column_id = ic.column_id
        WHERE ic.object_id = i.object_id
          AND ic.index_id = i.index_id
          AND ic.key_ordinal = expected.key_ordinal
          AND ic.is_descending_key = expected.is_descending_key
          AND c.name COLLATE Latin1_General_100_BIN2 = expected.column_name COLLATE Latin1_General_100_BIN2
      )
    )
    %s
)`, escapeSQLLiteral(tableName), escapeSQLLiteral(indexName), indexType, isUnique, optionConditions, filterCheck,
		len(keys), strings.Join(keyRows, ", "), includeCheck), true
}

func hasExplicitIndexDataSpaceClause(stmt string) bool {
	normalized := normalizeIndexStatement(stmt)
	semantic := maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(normalized))
	onPositions := findAllKeywordPositions(strings.ToUpper(semantic), "ON")
	if len(onPositions) == 0 {
		return false
	}
	afterOn := strings.TrimSpace(normalized[onPositions[0]+len("ON"):])
	_, consumed := consumeLeadingObjectToken(afterOn)
	if consumed == 0 {
		return true
	}
	afterTargetOffset := onPositions[0] + len("ON") +
		(strings.Index(normalized[onPositions[0]+len("ON"):], afterOn)) + consumed
	afterTarget := normalized[afterTargetOffset:]
	open := strings.Index(afterTarget, "(")
	if open < 0 {
		return true
	}
	closePos, ok := sqlparse.FindMatchingParen(afterTarget, open)
	if !ok {
		return true
	}
	tail := maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(afterTarget[closePos+1:]))
	upperTail := strings.ToUpper(tail)
	depth := 0
	for i := 0; i < len(upperTail); i++ {
		switch upperTail[i] {
		case '(':
			depth++
		case ')':
			if depth > 0 {
				depth--
			}
		default:
			if depth == 0 && hasKeywordAtWordBoundary(upperTail, i, "ON") {
				return true
			}
		}
	}
	return false
}

func buildIndexPersistedOptionConditions(stmt string) (string, bool) {
	expected := map[string]string{
		"FILLFACTOR":       "0",
		"PAD_INDEX":        "OFF",
		"IGNORE_DUP_KEY":   "OFF",
		"ALLOW_ROW_LOCKS":  "ON",
		"ALLOW_PAGE_LOCKS": "ON",
	}
	normalized := normalizeIndexStatement(stmt)
	semantic := maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(normalized))
	withPositions := findAllKeywordPositions(strings.ToUpper(semantic), "WITH")
	if len(withPositions) > 1 {
		return "", false
	}
	if len(withPositions) == 1 {
		afterWith := strings.TrimSpace(normalized[withPositions[0]+len("WITH"):])
		if !strings.HasPrefix(afterWith, "(") {
			return "", false
		}
		closePos, ok := sqlparse.FindMatchingParen(afterWith, 0)
		if !ok {
			return "", false
		}
		for _, raw := range splitTopLevelByComma(afterWith[1:closePos]) {
			parts := strings.SplitN(raw, "=", 2)
			if len(parts) != 2 {
				return "", false
			}
			name := strings.ToUpper(strings.TrimSpace(parts[0]))
			value := strings.ToUpper(strings.TrimSpace(parts[1]))
			switch name {
			case "FILLFACTOR":
				if _, err := strconv.Atoi(value); err != nil {
					return "", false
				}
				expected[name] = value
			case "PAD_INDEX", "IGNORE_DUP_KEY", "ALLOW_ROW_LOCKS", "ALLOW_PAGE_LOCKS":
				if value != "ON" && value != "OFF" {
					return "", false
				}
				expected[name] = value
			case "SORT_IN_TEMPDB", "ONLINE", "MAXDOP":
				// Execution-only build options do not change index semantics.
			default:
				// Persisted options not represented here must fail closed.
				return "", false
			}
		}
	}
	bit := func(value string) int {
		if value == "ON" {
			return 1
		}
		return 0
	}
	return fmt.Sprintf(
		"i.fill_factor = %s AND i.is_padded = %d AND i.ignore_dup_key = %d AND i.allow_row_locks = %d AND i.allow_page_locks = %d",
		expected["FILLFACTOR"], bit(expected["PAD_INDEX"]), bit(expected["IGNORE_DUP_KEY"]),
		bit(expected["ALLOW_ROW_LOCKS"]), bit(expected["ALLOW_PAGE_LOCKS"]),
	), true
}

func parseCreateIndexKeyDefinitions(stmt string) ([]indexKeyDefinition, error) {
	normalized := normalizeIndexStatement(stmt)
	semantic := maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(normalized))
	onPositions := findAllKeywordPositions(strings.ToUpper(semantic), "ON")
	if len(onPositions) == 0 {
		return nil, fmt.Errorf("CREATE INDEX ON target not found")
	}
	afterOn := strings.TrimSpace(normalized[onPositions[0]+len("ON"):])
	_, consumed := consumeLeadingObjectToken(afterOn)
	if consumed == 0 {
		return nil, fmt.Errorf("CREATE INDEX target is malformed")
	}
	afterTarget := strings.TrimSpace(afterOn[consumed:])
	if !strings.HasPrefix(afterTarget, "(") {
		return nil, fmt.Errorf("CREATE INDEX key column list is missing")
	}
	closePos, ok := sqlparse.FindMatchingParen(afterTarget, 0)
	if !ok {
		return nil, fmt.Errorf("CREATE INDEX key column list is malformed")
	}
	parts := splitTopLevelByComma(afterTarget[1:closePos])
	out := make([]indexKeyDefinition, 0, len(parts))
	for _, part := range parts {
		clean := strings.TrimSpace(part)
		token := firstSQLIdentifierToken(clean)
		name := stripIdentifierBrackets(cleanIdentifierToken(token))
		remainder := strings.ToUpper(strings.TrimSpace(clean[len(token):]))
		if name == "" || (remainder != "" && remainder != "ASC" && remainder != "DESC") {
			return nil, fmt.Errorf("unsupported CREATE INDEX key definition %q", part)
		}
		out = append(out, indexKeyDefinition{Name: name, Descending: remainder == "DESC"})
	}
	return out, nil
}

func parseCreateIndexIncludeColumns(stmt string) ([]string, error) {
	normalized := normalizeIndexStatement(stmt)
	semantic := maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(normalized))
	positions := findAllKeywordPositions(strings.ToUpper(semantic), "INCLUDE")
	if len(positions) == 0 {
		return nil, nil
	}
	if len(positions) > 1 {
		return nil, fmt.Errorf("CREATE INDEX contains multiple INCLUDE clauses")
	}

	afterInclude := strings.TrimSpace(normalized[positions[0]+len("INCLUDE"):])
	if !strings.HasPrefix(afterInclude, "(") {
		return nil, fmt.Errorf("CREATE INDEX INCLUDE column list is missing")
	}
	closePos, ok := sqlparse.FindMatchingParen(afterInclude, 0)
	if !ok {
		return nil, fmt.Errorf("CREATE INDEX INCLUDE column list is malformed")
	}
	rawList := afterInclude[1:closePos]
	parts := splitTopLevelByComma(rawList)
	columns := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			return nil, fmt.Errorf("CREATE INDEX INCLUDE contains an empty column")
		}
		token := firstSQLIdentifierToken(trimmed)
		name := stripIdentifierBrackets(cleanIdentifierToken(token))
		if name == "" || strings.TrimSpace(trimmed[len(token):]) != "" {
			return nil, fmt.Errorf("CREATE INDEX INCLUDE column %q is malformed", trimmed)
		}
		columns = append(columns, name)
	}
	if len(columns) == 0 {
		return nil, fmt.Errorf("CREATE INDEX INCLUDE column list is empty")
	}
	return columns, nil
}

func parseCreateIndexKeysAndFilter(stmt string) ([]string, string, error) {
	normalized := normalizeIndexStatement(stmt)
	semantic := maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(normalized))
	onPositions := findAllKeywordPositions(strings.ToUpper(semantic), "ON")
	if len(onPositions) == 0 {
		return nil, "", fmt.Errorf("CREATE INDEX ON target not found")
	}
	afterOn := strings.TrimSpace(normalized[onPositions[0]+len("ON"):])
	_, consumed := consumeLeadingObjectToken(afterOn)
	if consumed == 0 {
		return nil, "", fmt.Errorf("CREATE INDEX target is malformed")
	}
	afterTarget := strings.TrimSpace(afterOn[consumed:])
	if !strings.HasPrefix(afterTarget, "(") {
		upper := strings.ToUpper(semantic)
		if strings.Contains(upper, "CLUSTERED COLUMNSTORE INDEX") && !strings.Contains(upper, "NONCLUSTERED COLUMNSTORE INDEX") {
			return nil, "", nil
		}
		return nil, "", fmt.Errorf("CREATE INDEX key column list is missing")
	}
	closePos, ok := sqlparse.FindMatchingParen(afterTarget, 0)
	if !ok {
		return nil, "", fmt.Errorf("CREATE INDEX key column list is malformed")
	}
	columns := parseIdentifierList(afterTarget[1:closePos])
	if len(columns) == 0 {
		return nil, "", fmt.Errorf("CREATE INDEX key column list is empty")
	}

	remainder := strings.TrimSpace(afterTarget[closePos+1:])
	remainderSemantic := maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(remainder))
	wherePositions := findAllKeywordPositions(strings.ToUpper(remainderSemantic), "WHERE")
	if len(wherePositions) == 0 {
		return columns, "", nil
	}
	wherePos := wherePositions[0]
	predicate := strings.TrimSpace(remainder[wherePos+len("WHERE"):])
	predicateSemantic := maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(predicate))
	if withPositions := findAllKeywordPositions(strings.ToUpper(predicateSemantic), "WITH"); len(withPositions) > 0 {
		predicate = strings.TrimSpace(predicate[:withPositions[0]])
	}
	predicate = strings.TrimSpace(strings.TrimRight(predicate, ";"))
	if predicate == "" {
		return nil, "", fmt.Errorf("filtered CREATE INDEX predicate is empty")
	}
	return columns, predicate, nil
}

func buildUniqueIndexDataPrecheck(tableName, indexName, stmt string) string {
	upper := strings.ToUpper(strings.TrimSpace(stmt))
	if !strings.HasPrefix(upper, "CREATE UNIQUE ") {
		return ""
	}
	columns, predicate, err := parseCreateIndexKeysAndFilter(stmt)
	if err != nil {
		return fmt.Sprintf(";THROW 51000, N'Precheck failed: could not parse UNIQUE INDEX %s', 1;", escapeSQLLiteral(indexName))
	}
	quoted := make([]string, 0, len(columns))
	for _, column := range columns {
		quoted = append(quoted, quoteIdentifier(column))
	}
	whereSQL := ""
	if predicate != "" {
		whereSQL = "\n  WHERE " + predicate
	}
	return normalizeSQLText(fmt.Sprintf(`-- check duplicate keys before CREATE UNIQUE INDEX
IF EXISTS (
  SELECT 1
  FROM %s%s
  GROUP BY %s
  HAVING COUNT_BIG(*) > 1
)
BEGIN
  ;THROW 51000, N'Precheck blocked: UNIQUE INDEX %s has duplicate key values', 1;
END`, quoteQualifiedIdentifier(tableName), whereSQL, strings.Join(quoted, ", "), escapeSQLLiteral(indexName)), true, true)
}

func extractDropIndexTargets(sql string) []dropIndexTarget {
	normalized := normalizeSQLText(sql, true, true)
	semantic := sqlparse.StripSQLStringsAndComments(normalized)
	upper := strings.ToUpper(semantic)
	anchor := "DROP INDEX "

	seen := make(map[string]dropIndexTarget)
	cursor := 0
	for {
		rel := strings.Index(upper[cursor:], anchor)
		if rel < 0 {
			break
		}
		start := cursor + rel
		end := sqlparse.FindStatementEnd(normalized, start)

		stmt := strings.TrimSpace(normalized[start:end])
		stmtUpper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(stmt))
		if strings.HasPrefix(stmtUpper, "DROP INDEX ") {
			afterDrop := strings.TrimSpace(stmt[len("DROP INDEX "):])
			upperAfterDrop := strings.ToUpper(afterDrop)
			if strings.HasPrefix(upperAfterDrop, "IF EXISTS ") {
				afterDrop = strings.TrimSpace(afterDrop[len("IF EXISTS "):])
			}
			onIdx := strings.Index(strings.ToUpper(afterDrop), " ON ")
			if onIdx > 0 {
				nameToken := strings.TrimSpace(afterDrop[:onIdx])
				targetToken, _ := consumeLeadingObjectToken(strings.TrimSpace(afterDrop[onIdx+len(" ON "):]))
				name := stripIdentifierBrackets(cleanIdentifierToken(nameToken))
				target := stripIdentifierBrackets(cleanIdentifierToken(targetToken))
				if name != "" && target != "" {
					key := target + "|" + name
					seen[key] = dropIndexTarget{
						Table: target,
						Name:  name,
					}
				}
			}
		}

		cursor = end
		if cursor >= len(upper) {
			break
		}
	}

	out := make([]dropIndexTarget, 0, len(seen))
	for _, item := range seen {
		out = append(out, item)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Table != out[j].Table {
			return out[i].Table < out[j].Table
		}
		return out[i].Name < out[j].Name
	})
	return out
}

func applyIndexStateTransition(objectName, migrationSQL, currentSemantic string, hasCurrent bool) (string, error) {
	current := make(map[string]indexStatementDef)
	var err error
	if hasCurrent && strings.TrimSpace(currentSemantic) != "" {
		current, err = parseIndexStateMap(objectName, currentSemantic)
		if err != nil {
			return "", err
		}
	}

	createMap, err := parseIndexStateMap(objectName, migrationSQL)
	if err != nil {
		return "", err
	}
	drops := extractDropIndexTargets(migrationSQL)
	targetedDrops := make([]dropIndexTarget, 0)
	for _, drop := range drops {
		if sameQualifiedObject(drop.Table, objectName) {
			targetedDrops = append(targetedDrops, drop)
		}
	}

	if len(targetedDrops) == 0 && len(createMap) == 0 {
		return "", fmt.Errorf("no CREATE/DROP INDEX statement found for %s", objectName)
	}
	if len(targetedDrops) > 0 && !hasCurrent && strings.TrimSpace(currentSemantic) == "" && len(createMap) == 0 {
		return "", fmt.Errorf("drop-index migration requires prior index state for %s", objectName)
	}

	for _, drop := range targetedDrops {
		delete(current, stripIdentifierBrackets(drop.Name))
	}
	for key, def := range createMap {
		current[key] = def
	}

	return serializeIndexState(current), nil
}

func extractProgrammabilityStatement(category, objectName, sql string) (string, error) {
	if dynamicBatch, ok := extractSpExecuteSQLBatch(sql); ok {
		sql = dynamicBatch
	}
	normalized := normalizeSQLText(sql, true, true)
	upper := strings.ToUpper(normalized)

	kind := ""
	switch category {
	case "views":
		kind = "VIEW"
	case "functions":
		kind = "FUNCTION"
	case "procedures":
		kind = "PROCEDURE"
	default:
		return "", fmt.Errorf("unsupported programmability category %q", category)
	}

	for _, match := range findCreateOrAlterKindMatches(upper, kind) {
		start := match.Start
		rest := strings.TrimSpace(normalized[match.AfterKind:])
		token, _ := consumeLeadingObjectToken(rest)
		token = stripIdentifierBrackets(strings.TrimSpace(strings.TrimRight(token, ",;")))
		if token != "" && sameQualifiedObject(token, objectName) {
			return strings.TrimSpace(normalized[start:]), nil
		}
	}
	return "", fmt.Errorf("CREATE OR ALTER %s statement for %s not found", kind, objectName)
}

func extractSpExecuteSQLBatch(sql string) (string, bool) {
	batches := extractSpExecuteSQLBatches(sql)
	if len(batches) == 0 {
		return "", false
	}
	return batches[0], true
}

func extractSpExecuteSQLBatches(sql string) []string {
	upper := strings.ToUpper(sql)
	marker := "EXEC SYS.SP_EXECUTESQL N'"
	batches := make([]string, 0)
	for searchFrom := 0; searchFrom < len(sql); {
		rel := strings.Index(upper[searchFrom:], marker)
		if rel < 0 {
			break
		}
		cursor := searchFrom + rel + len(marker)
		var decoded strings.Builder
		closed := false
		for cursor < len(sql) {
			if sql[cursor] != '\'' {
				decoded.WriteByte(sql[cursor])
				cursor++
				continue
			}
			if cursor+1 < len(sql) && sql[cursor+1] == '\'' {
				decoded.WriteByte('\'')
				cursor += 2
				continue
			}
			closed = true
			cursor++
			break
		}
		if closed && strings.TrimSpace(decoded.String()) != "" {
			batches = append(batches, strings.TrimSpace(decoded.String()))
		}
		searchFrom = cursor
	}
	return batches
}

// buildDropProgrammabilitySQL renders the exact inverse of an initial
// CREATE OR ALTER module migration. SQL Server supports DROP ... IF EXISTS
// on every supported server version that also supports CREATE OR ALTER.
func buildDropProgrammabilitySQL(category, objectName string) (string, error) {
	kind := ""
	switch strings.ToLower(strings.TrimSpace(category)) {
	case "views":
		kind = "VIEW"
	case "functions":
		kind = "FUNCTION"
	case "procedures":
		kind = "PROCEDURE"
	default:
		return "", fmt.Errorf("unsupported programmability category %q", category)
	}

	target := quoteQualifiedIdentifier(objectName)
	if strings.TrimSpace(target) == "" {
		return "", fmt.Errorf("programmability object name must not be empty")
	}
	return fmt.Sprintf("DROP %s IF EXISTS %s;\n", kind, target), nil
}

func buildGuardedDropProgrammabilitySQL(category, objectName string) (string, error) {
	dropSQL, err := buildDropProgrammabilitySQL(category, objectName)
	if err != nil {
		return "", err
	}
	safeObject := escapeSQLLiteral(objectName)
	textReferenceCondition := buildModuleTextReferenceCondition(objectName)
	return normalizeSQLText(fmt.Sprintf(`-- Do not remove a module while another SQL object still references it.
IF OBJECT_ID(N'%s') IS NOT NULL
AND EXISTS (
  SELECT 1
  FROM sys.sql_expression_dependencies
  WHERE referenced_id = OBJECT_ID(N'%s')
    AND referencing_id <> referenced_id
)
BEGIN
  ;THROW 51000, N'Precheck blocked: object is referenced by another SQL object', 1;
END

IF OBJECT_ID(N'%s') IS NOT NULL
AND EXISTS (
  SELECT 1
  FROM sys.sql_modules
  WHERE object_id <> OBJECT_ID(N'%s')
    AND (
      %s
    )
)
BEGIN
  ;THROW 51000, N'Precheck blocked: object name is used in module text or dynamic SQL', 1;
END

%s`, safeObject, safeObject, safeObject, safeObject, textReferenceCondition, strings.TrimSpace(dropSQL)), true, true), nil
}

func buildGuardedDropTypeSQL(objectName string) string {
	safeObject := escapeSQLLiteral(objectName)
	return normalizeSQLText(fmt.Sprintf(`-- Do not remove a user-defined table type while columns or parameters still use it.
IF TYPE_ID(N'%s') IS NOT NULL
AND (
  EXISTS (
    SELECT 1
    FROM sys.columns
    WHERE user_type_id = TYPE_ID(N'%s')
  )
  OR EXISTS (
    SELECT 1
    FROM sys.parameters
    WHERE user_type_id = TYPE_ID(N'%s')
  )
)
BEGIN
  ;THROW 51000, N'Precheck blocked: user-defined type is still in use', 1;
END

DROP TYPE IF EXISTS %s;`, safeObject, safeObject, safeObject, quoteQualifiedIdentifier(objectName)), true, true)
}

type createOrAlterKindMatch struct {
	Start     int
	AfterKind int
}

func findCreateOrAlterKindMatches(inputUpper, kind string) []createOrAlterKindMatch {
	matches := make([]createOrAlterKindMatch, 0)
	for i := 0; i < len(inputUpper); i++ {
		if !hasKeywordAtWordBoundary(inputUpper, i, "CREATE") {
			continue
		}
		next, tok := readNextWordUpper(inputUpper, i+len("CREATE"))
		if tok != "OR" {
			continue
		}
		next, tok = readNextWordUpper(inputUpper, next)
		if tok != "ALTER" {
			continue
		}
		next, tok = readNextWordUpper(inputUpper, next)
		if tok != kind {
			continue
		}
		matches = append(matches, createOrAlterKindMatch{Start: i, AfterKind: next})
	}
	return matches
}

func validateTableSubsetSupport(stmt string) error {
	// Feature detection must inspect SQL tokens, not words that happen to
	// occur in comments, string literals, or bracketed identifiers.
	upper := strings.ToUpper(maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(stmt)))
	for _, token := range []string{
		"FILESTREAM",
		"COLUMN_SET",
		"MASKED WITH",
		"ENCRYPTED WITH",
		"PERIOD FOR SYSTEM_TIME",
		"FOR SYSTEM_TIME",
		"GENERATED ALWAYS",
		"SYSTEM_VERSIONING",
		" FILEGROUP ",
		" PARTITION ",
		"MEMORY_OPTIMIZED",
		"LEDGER",
		" AS NODE",
		" AS EDGE",
		"FILESTREAM_ON",
	} {
		if strings.Contains(upper, token) {
			return fmt.Errorf("unsupported table subset feature detected: %s", strings.TrimSpace(token))
		}
	}
	items, err := extractCreateTableItems(stmt)
	if err != nil {
		return err
	}
	for _, item := range items {
		if err := validateComputedColumnItem(item); err != nil {
			return err
		}
	}
	return nil
}

func validateComputedColumnItem(item string) error {
	trimmed := strings.TrimSpace(item)
	nameToken, consumed := consumeLeadingColumnIdentifierToken(trimmed)
	if nameToken == "" {
		return nil
	}
	rest := strings.TrimSpace(trimmed[consumed:])
	if !hasKeywordAtWordBoundary(strings.ToUpper(rest), 0, "AS") {
		return nil
	}
	if strings.HasSuffix(strings.ToUpper(rest), " PERSISTED NULL") {
		return fmt.Errorf("computed column %s must not use PERSISTED NULL", columnIdentifierName(nameToken))
	}
	return nil
}

func isComputedColumnTableItem(item string) bool {
	trimmed := strings.TrimSpace(item)
	if trimmed == "" {
		return false
	}
	upper := strings.ToUpper(trimmed)
	if strings.HasPrefix(upper, "CONSTRAINT ") ||
		strings.HasPrefix(upper, "PRIMARY KEY") ||
		strings.HasPrefix(upper, "UNIQUE") ||
		strings.HasPrefix(upper, "CHECK") ||
		strings.HasPrefix(upper, "FOREIGN KEY") {
		return false
	}
	for i := 0; i < len(upper); i++ {
		if hasKeywordAtWordBoundary(upper, i, "AS") {
			return true
		}
	}
	return false
}

func validateIndexSubsetSupport(stmt string) error {
	upper := strings.ToUpper(maskBracketedIdentifiers(sqlparse.StripSQLStringsAndComments(stmt)))
	for _, token := range []string{
		"RESUMABLE =",
		"ON PARTITIONS",
		"DROP_EXISTING = ON",
		"CREATE FULLTEXT INDEX",
		"CREATE SELECTIVE XML INDEX",
		"CREATE JSON INDEX",
		"BUCKET_COUNT",
	} {
		if strings.Contains(upper, token) {
			return fmt.Errorf("unsupported index subset feature detected: %s", strings.TrimSpace(token))
		}
	}
	if strings.Contains(upper, "COLUMNSTORE INDEX") && len(findAllKeywordPositions(upper, "ORDER")) > 0 {
		return fmt.Errorf("unsupported index subset feature detected: ordered columnstore")
	}
	return nil
}

func maskBracketedIdentifiers(input string) string {
	if !strings.Contains(input, "[") {
		return input
	}
	out := []byte(input)
	for i := 0; i < len(out); i++ {
		if out[i] != '[' {
			continue
		}
		out[i] = ' '
		for i++; i < len(out); i++ {
			if out[i] == ']' {
				if i+1 < len(out) && out[i+1] == ']' {
					out[i], out[i+1] = ' ', ' '
					i++
					continue
				}
				out[i] = ' '
				break
			}
			out[i] = ' '
		}
	}
	return string(out)
}

func findAllKeywordPositions(inputUpper, keywordUpper string) []int {
	if keywordUpper == "" || len(inputUpper) < len(keywordUpper) {
		return []int{}
	}
	pos := make([]int, 0)
	limit := len(inputUpper) - len(keywordUpper)
	for i := 0; i <= limit; i++ {
		if inputUpper[i:i+len(keywordUpper)] != keywordUpper {
			continue
		}
		leftOK := i == 0 || !isWordChar(rune(inputUpper[i-1]))
		rightIdx := i + len(keywordUpper)
		rightOK := rightIdx >= len(inputUpper) || !isWordChar(rune(inputUpper[rightIdx]))
		if leftOK && rightOK {
			pos = append(pos, i)
		}
	}
	return pos
}

func findCreateIndexStatementPositions(inputUpper string) []int {
	positions := make([]int, 0)
	seen := make(map[int]struct{})

	for i := 0; i < len(inputUpper); i++ {
		if !hasKeywordAtWordBoundary(inputUpper, i, "CREATE") {
			continue
		}

		next, tok := readNextWordUpper(inputUpper, i+len("CREATE"))
		if tok == "" {
			continue
		}

		matched := false
		switch tok {
		case "INDEX":
			matched = true
		case "UNIQUE":
			next2, tok2 := readNextWordUpper(inputUpper, next)
			switch tok2 {
			case "INDEX":
				matched = true
			case "NONCLUSTERED", "CLUSTERED":
				_, tok3 := readNextWordUpper(inputUpper, next2)
				if tok3 == "INDEX" {
					matched = true
				}
			}
		case "NONCLUSTERED", "CLUSTERED":
			next2, tok2 := readNextWordUpper(inputUpper, next)
			if tok2 == "INDEX" {
				matched = true
			} else if tok2 == "COLUMNSTORE" {
				_, tok3 := readNextWordUpper(inputUpper, next2)
				matched = tok3 == "INDEX"
			}
		case "COLUMNSTORE", "XML", "SPATIAL":
			_, tok2 := readNextWordUpper(inputUpper, next)
			matched = tok2 == "INDEX"
		case "PRIMARY":
			next2, tok2 := readNextWordUpper(inputUpper, next)
			if tok2 == "XML" {
				_, tok3 := readNextWordUpper(inputUpper, next2)
				matched = tok3 == "INDEX"
			}
		}
		if matched {
			if _, ok := seen[i]; ok {
				continue
			}
			seen[i] = struct{}{}
			positions = append(positions, i)
		}
	}
	sort.Ints(positions)
	return positions
}

func hasKeywordAtWordBoundary(inputUpper string, pos int, keyword string) bool {
	if pos < 0 || pos+len(keyword) > len(inputUpper) {
		return false
	}
	if inputUpper[pos:pos+len(keyword)] != keyword {
		return false
	}
	leftOK := pos == 0 || !isWordChar(rune(inputUpper[pos-1]))
	rightIdx := pos + len(keyword)
	rightOK := rightIdx >= len(inputUpper) || !isWordChar(rune(inputUpper[rightIdx]))
	return leftOK && rightOK
}

func readNextWordUpper(inputUpper string, start int) (int, string) {
	i := start
	for i < len(inputUpper) {
		switch inputUpper[i] {
		case ' ', '\t', '\r', '\n':
			i++
		default:
			goto tokenStart
		}
	}
	return i, ""

tokenStart:
	j := i
	for j < len(inputUpper) && isWordChar(rune(inputUpper[j])) {
		j++
	}
	if j == i {
		return i + 1, ""
	}
	return j, inputUpper[i:j]
}

func sortDesiredObjectsByOrder(items []desiredObject, order []string) {
	orderIndex := diff.BuildOrderIndex(order)
	sort.Slice(items, func(i, j int) bool {
		return diff.LessCategoryObject(orderIndex, len(order), items[i].Category, items[i].Object, items[j].Category, items[j].Object)
	})
}

func sortPlannedOps(ops []plannedMigration, order []string) {
	orderIndex := diff.BuildOrderIndex(order)
	sort.Slice(ops, func(i, j int) bool {
		return diff.LessPlanned(orderIndex, len(order), ops[i].Category, ops[i].Object, ops[i].VK, ops[j].Category, ops[j].Object, ops[j].VK)
	})
}

func removeFiles(paths []string) error {
	var removeErr error
	for _, path := range paths {
		if err := os.Remove(path); err != nil && !errors.Is(err, os.ErrNotExist) && removeErr == nil {
			removeErr = err
		}
	}
	return removeErr
}

func (e Engine) copyFailedEmitFiles(paths []string) ([]string, error) {
	if len(paths) == 0 {
		return []string{}, nil
	}
	targetDir := filepath.Join(e.cfg.RepoPath(e.cfg.OutDir), "failed_emit")
	if err := os.MkdirAll(targetDir, 0o755); err != nil {
		return nil, err
	}
	copied := make([]string, 0, len(paths))
	for _, src := range paths {
		base := filepath.Base(src)
		dst := filepath.Join(targetDir, base)
		data, err := os.ReadFile(src)
		if err != nil {
			return copied, err
		}
		if err := os.WriteFile(dst, data, 0o644); err != nil {
			return copied, err
		}
		copied = append(copied, filepath.ToSlash(filepath.Join(e.cfg.OutDir, "failed_emit", base)))
	}
	return copied, nil
}

func objectKey(category, object string) string {
	return history.ObjectKey(category, object)
}

func shouldReplaceHistory(current, next historyObject) bool {
	return history.ShouldReplace(current.LastVersion, current.LastVK, next.LastVersion, next.LastVK)
}

func hashNormalized(input string) string {
	sum := sha256.Sum256([]byte(normalizeForSemanticHash(input)))
	return hex.EncodeToString(sum[:])
}

func hashSemanticByCategory(category, input string) string {
	normalized := normalizeForSemanticHash(input)
	if category == "tables" {
		normalized = normalizeCreateTableForHash(input)
	} else if category == "indexes" {
		normalized = normalizeCreateIndexesForHash(input)
	} else if category == "views" || category == "functions" || category == "procedures" {
		normalized = normalizeProgrammabilityForHash(input)
	}
	sum := sha256.Sum256([]byte(normalized))
	return hex.EncodeToString(sum[:])
}

func hashStableContent(input string) string {
	normalized := normalizeSQLText(input, true, true)
	sum := sha256.Sum256([]byte(normalized))
	return hex.EncodeToString(sum[:])
}

func canonicalizePatchName(input string) string {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return ""
	}

	var b strings.Builder
	lastUnderscore := false
	for _, ch := range strings.ToLower(trimmed) {
		switch {
		case unicode.IsSpace(ch):
			if !lastUnderscore {
				b.WriteRune('_')
				lastUnderscore = true
			}
		case (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '-':
			if ch == '_' {
				if lastUnderscore {
					continue
				}
				lastUnderscore = true
			} else {
				lastUnderscore = false
			}
			b.WriteRune(ch)
		default:
			if !lastUnderscore {
				b.WriteRune('_')
				lastUnderscore = true
			}
		}
	}

	out := b.String()
	out = strings.Trim(out, "_-")
	for strings.Contains(out, "__") {
		out = strings.ReplaceAll(out, "__", "_")
	}
	return out
}

func normalizeCreateTableForHash(input string) string {
	items, err := extractCreateTableItems(input)
	if err != nil || len(items) == 0 {
		return normalizeCreateTableTerminatorForHash(input)
	}

	defaultsByColumn, defaultsErr := extractTableDefaultConstraintMap(input)
	if defaultsErr != nil {
		defaultsByColumn = map[string]tableDefaultConstraintDef{}
	}

	columnItems := make([]string, 0, len(items))
	constraintItems := make([]string, 0, len(items))
	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		if isStandaloneDefaultConstraintItemForHash(line) {
			continue
		}
		if canonicalColumn, ok := canonicalizeColumnItemForHash(line); ok {
			columnItems = append(columnItems, canonicalColumn)
			continue
		}
		upperLine := strings.ToUpper(strings.TrimSpace(line))
		if isConstraintLikeTableItem(upperLine) {
			constraintItems = append(constraintItems, strings.Join(strings.Fields(line), " "))
			continue
		}
		columnItems = append(columnItems, line)
	}
	sort.Strings(columnItems)
	sort.Strings(constraintItems)

	normalizedItems := make([]string, 0, len(columnItems)+len(constraintItems)+len(defaultsByColumn))
	normalizedItems = append(normalizedItems, columnItems...)
	normalizedItems = append(normalizedItems, constraintItems...)
	if len(defaultsByColumn) > 0 {
		cols := make([]string, 0, len(defaultsByColumn))
		for col := range defaultsByColumn {
			cols = append(cols, col)
		}
		sort.Strings(cols)
		for _, col := range cols {
			signature := canonicalDefaultConstraintSignatureForHash(defaultsByColumn[col])
			if strings.TrimSpace(signature) != "" {
				normalizedItems = append(normalizedItems, signature)
			}
		}
	}
	rebuilt, err := rebuildCreateTableStatementWithItems(input, normalizedItems)
	if err != nil {
		return normalizeCreateTableTerminatorForHash(input)
	}
	return normalizeCreateTableTerminatorForHash(rebuilt)
}

func normalizeCreateTableTerminatorForHash(input string) string {
	normalized := strings.TrimSpace(normalizeForSemanticHash(input))
	normalized = strings.TrimSuffix(normalized, ";")
	return strings.TrimSpace(normalized)
}

func normalizeCreateIndexesForHash(input string) string {
	normalized := normalizeSQLText(input, true, true)
	if strings.TrimSpace(normalized) == "" {
		return normalizeForSemanticHash(normalized)
	}

	upper := strings.ToUpper(normalized)
	positions := findCreateIndexStatementPositions(upper)
	if len(positions) == 0 {
		return normalizeForSemanticHash(normalized)
	}
	sort.Ints(positions)

	type indexDef struct {
		Name string
		Stmt string
	}
	byKey := make(map[string]indexDef)
	for _, pos := range positions {
		end := sqlparse.FindStatementEnd(normalized, pos)
		stmt := strings.TrimSpace(normalized[pos:end])
		if stmt == "" {
			continue
		}
		name, err := parseCreateIndexName(stmt)
		if err != nil {
			// Keep backward compatibility: if parser cannot identify one statement,
			// fallback to generic normalization instead of forcing false drift.
			return normalizeForSemanticHash(normalized)
		}
		cleanName := stripIdentifierBrackets(cleanIdentifierToken(name))
		if cleanName == "" {
			return normalizeForSemanticHash(normalized)
		}
		key := cleanName
		byKey[key] = indexDef{
			Name: cleanName,
			Stmt: normalizeIndexStatement(stmt),
		}
	}

	if len(byKey) == 0 {
		return normalizeForSemanticHash(normalized)
	}

	keys := make([]string, 0, len(byKey))
	for key := range byKey {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	lines := make([]string, 0, len(keys))
	for _, key := range keys {
		stmt := strings.TrimSpace(byKey[key].Stmt)
		if stmt == "" {
			continue
		}
		lines = append(lines, stmt)
	}
	if len(lines) == 0 {
		return normalizeForSemanticHash(normalized)
	}
	return normalizeForSemanticHash(strings.Join(lines, "\n"))
}

func canonicalizeColumnItemForHash(item string) (string, bool) {
	upper := strings.ToUpper(strings.TrimSpace(item))
	if upper == "" {
		return "", false
	}
	if strings.HasPrefix(upper, "CONSTRAINT ") ||
		strings.HasPrefix(upper, "PRIMARY KEY") ||
		strings.HasPrefix(upper, "UNIQUE ") ||
		strings.HasPrefix(upper, "FOREIGN KEY") ||
		strings.HasPrefix(upper, "CHECK ") {
		return "", false
	}

	nameToken := firstSQLToken(item)
	if nameToken == "" {
		return "", false
	}
	name := cleanIdentifierToken(nameToken)
	if name == "" {
		return "", false
	}
	rest := strings.TrimSpace(item[len(firstSQLToken(item)):])
	if hasKeywordAtWordBoundary(strings.ToUpper(rest), 0, "AS") {
		return fmt.Sprintf("[%s] %s", name, normalizeForSemanticHash(rest)), true
	}
	typeToken := firstSQLToken(rest)
	if typeToken == "" {
		return "", false
	}
	rawType := parseColumnRawTypeSQL(rest, typeToken)
	nullable := true
	restUpper := strings.ToUpper(rest)
	if strings.Contains(restUpper, " NOT NULL") {
		nullable = false
	} else if strings.Contains(restUpper, " NULL") {
		nullable = true
	}
	nullability := "NULL"
	if !nullable {
		nullability = "NOT NULL"
	}
	modifiers := make([]string, 0, 4)
	if identity := extractIdentityDefinition(rest); identity != "" {
		modifiers = append(modifiers, normalizeForSemanticHash(identity))
	}
	if strings.Contains(restUpper, " ROWGUIDCOL") {
		modifiers = append(modifiers, "ROWGUIDCOL")
	}
	if strings.Contains(restUpper, " SPARSE") {
		modifiers = append(modifiers, "SPARSE")
	}
	if kind := constraintKindForTableItem(item); kind != "constraint" && kind != "default" {
		if positions := findAllKeywordPositions(restUpper, "CONSTRAINT"); len(positions) > 0 {
			inline := strings.TrimSpace(rest[positions[0]:])
			modifiers = append(modifiers, "INLINE="+normalizeForSemanticHash(inline))
		}
	}
	modifierSQL := ""
	if len(modifiers) > 0 {
		modifierSQL = " " + strings.Join(modifiers, " ")
	}
	return fmt.Sprintf("[%s] %s%s %s", name, normalizeForSemanticHash(rawType), modifierSQL, nullability), true
}

func isStandaloneDefaultConstraintItemForHash(line string) bool {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return false
	}
	upper := strings.ToUpper(trimmed)
	if !strings.HasPrefix(upper, "CONSTRAINT ") && !strings.HasPrefix(upper, "DEFAULT ") {
		return false
	}
	_, ok := parseTableLevelDefaultConstraint(trimmed, upper)
	return ok
}

func canonicalDefaultConstraintSignatureForHash(def tableDefaultConstraintDef) string {
	column := strings.TrimSpace(def.Column)
	expression := normalizeDefaultExpressionForCompare(def.Expression)
	if column == "" || expression == "" {
		return ""
	}
	name := strings.TrimSpace(def.ConstraintName)
	return fmt.Sprintf("DEFAULT_SIGNATURE[%s][%s][%s]", column, name, expression)
}

func normalizeForSemanticHash(input string) string {
	return sqlparse.CanonicalSemanticSQL(normalizeSQLText(input, true, true))
}

func normalizeProgrammabilityForHash(input string) string {
	return sqlparse.CanonicalSemanticSQL(normalizeSQLText(input, true, true))
}

type tableColumnChanges struct {
	Added          []string
	Removed        []string
	Changed        []string
	DefaultChanged []string
}

type dropConstraintTarget struct {
	Table string
	Name  string
	Kind  string
}

type dropIndexTarget struct {
	Table string
	Name  string
}

type dropColumnDependencyPlan struct {
	Constraints []dropConstraintTarget
	Indexes     []dropIndexTarget
}

type foreignKeyConstraintReference struct {
	Name              string
	ReferencedTable   string
	ReferencedColumns []string
}

func classifyTableColumnChanges(previousCreateTable, nextCreateTable string) (tableColumnChanges, error) {
	prevCols, err := parseTableColumns(previousCreateTable)
	if err != nil {
		return tableColumnChanges{}, err
	}
	nextCols, err := parseTableColumns(nextCreateTable)
	if err != nil {
		return tableColumnChanges{}, err
	}

	added := make([]string, 0)
	removed := make([]string, 0)
	changed := make([]string, 0)
	defaultChanged := make([]string, 0)

	for name, prev := range prevCols {
		if _, ok := nextCols[name]; !ok {
			removed = append(removed, prev.Name)
		}
	}
	for name, next := range nextCols {
		if _, ok := prevCols[name]; !ok {
			added = append(added, next.Name)
		}
	}
	for name, prev := range prevCols {
		next, ok := nextCols[name]
		if !ok {
			continue
		}
		if !columnTypeEquivalent(prev, next) || prev.Nullable != next.Nullable || prev.IdentityDefinition != next.IdentityDefinition || prev.RowGuidCol != next.RowGuidCol || prev.Sparse != next.Sparse ||
			prev.IsComputed != next.IsComputed || prev.ComputedDefinition != next.ComputedDefinition {
			changed = append(changed, next.Name)
		}
	}

	prevDefaultsByColumn, err := extractTableDefaultConstraintMap(previousCreateTable)
	if err != nil {
		return tableColumnChanges{}, err
	}
	nextDefaultsByColumn, err := extractTableDefaultConstraintMap(nextCreateTable)
	if err != nil {
		return tableColumnChanges{}, err
	}
	defaultCandidateCols := make(map[string]struct{}, len(prevDefaultsByColumn)+len(nextDefaultsByColumn))
	for col := range prevDefaultsByColumn {
		defaultCandidateCols[col] = struct{}{}
	}
	for col := range nextDefaultsByColumn {
		defaultCandidateCols[col] = struct{}{}
	}
	for col := range defaultCandidateCols {
		// Default evolution is only relevant for existing columns.
		if _, inPrev := prevCols[col]; !inPrev {
			continue
		}
		if _, inNext := nextCols[col]; !inNext {
			continue
		}
		prevDef, hasPrev := prevDefaultsByColumn[col]
		nextDef, hasNext := nextDefaultsByColumn[col]
		if !hasPrev || !hasNext || !defaultConstraintEquivalent(prevDef, nextDef) {
			if nextDef.Column != "" {
				defaultChanged = append(defaultChanged, nextDef.Column)
			} else if prevDef.Column != "" {
				defaultChanged = append(defaultChanged, prevDef.Column)
			} else {
				defaultChanged = append(defaultChanged, col)
			}
		}
	}

	sort.Strings(added)
	sort.Strings(removed)
	sort.Strings(changed)
	sort.Strings(defaultChanged)
	return tableColumnChanges{
		Added:          added,
		Removed:        removed,
		Changed:        changed,
		DefaultChanged: defaultChanged,
	}, nil
}

func buildDropColumnDependencyPlan(object string, removedColumns []string, previousCreateTable string, historyState map[string]historyObject) (dropColumnDependencyPlan, error) {
	if len(removedColumns) == 0 {
		return dropColumnDependencyPlan{}, nil
	}

	removeSet := make(map[string]struct{}, len(removedColumns))
	for _, col := range removedColumns {
		removeSet[strings.TrimSpace(col)] = struct{}{}
	}

	localConstraints, err := extractNamedConstraintDropsForRemovedColumns(object, previousCreateTable, removeSet)
	if err != nil {
		return dropColumnDependencyPlan{}, err
	}
	incomingForeignKeys, err := extractIncomingForeignKeyDropsFromHistory(object, removeSet, historyState)
	if err != nil {
		return dropColumnDependencyPlan{}, err
	}

	constraintByKey := make(map[string]dropConstraintTarget)
	mergeConstraint := func(item dropConstraintTarget) {
		key := item.Table + "|" + item.Name
		current, exists := constraintByKey[key]
		if !exists || constraintDropPriority(item.Kind) < constraintDropPriority(current.Kind) {
			constraintByKey[key] = item
		}
	}
	for _, item := range localConstraints {
		mergeConstraint(item)
	}
	for _, item := range incomingForeignKeys {
		mergeConstraint(item)
	}

	constraints := make([]dropConstraintTarget, 0, len(constraintByKey))
	for _, item := range constraintByKey {
		constraints = append(constraints, item)
	}

	indexes := make([]dropIndexTarget, 0)
	indexKey := objectKey("indexes", object)
	if indexState, ok := historyState[indexKey]; ok && strings.TrimSpace(indexState.SemanticSQL) != "" {
		indexes, err = extractImpactedIndexDropsForRemovedColumns(object, indexState.SemanticSQL, removeSet)
		if err != nil {
			return dropColumnDependencyPlan{}, err
		}
	}

	sort.Slice(constraints, func(i, j int) bool {
		pi := constraintDropPriority(constraints[i].Kind)
		pj := constraintDropPriority(constraints[j].Kind)
		if pi != pj {
			return pi < pj
		}
		if constraints[i].Table != constraints[j].Table {
			return constraints[i].Table < constraints[j].Table
		}
		return constraints[i].Name < constraints[j].Name
	})
	sort.Slice(indexes, func(i, j int) bool {
		if indexes[i].Table != indexes[j].Table {
			return indexes[i].Table < indexes[j].Table
		}
		return indexes[i].Name < indexes[j].Name
	})

	return dropColumnDependencyPlan{
		Constraints: constraints,
		Indexes:     indexes,
	}, nil
}

func buildAlterColumnDependencyBlockList(object string, changedColumns []string, previousCreateTable string, historyState map[string]historyObject) ([]string, error) {
	if len(changedColumns) == 0 {
		return []string{}, nil
	}
	plan, err := buildDropColumnDependencyPlan(object, changedColumns, previousCreateTable, historyState)
	if err != nil {
		return nil, err
	}
	blocked := make([]string, 0, len(plan.Constraints)+len(plan.Indexes))
	for _, dep := range plan.Constraints {
		// Named DEFAULT constraints on the same table are auto-managed
		// in buildTableAlterColumnsSQL via default-constraint delta steps.
		if strings.EqualFold(strings.TrimSpace(dep.Kind), "default") && sameQualifiedObject(dep.Table, object) {
			continue
		}
		blocked = append(blocked, fmt.Sprintf("constraint:%s.%s", dep.Table, dep.Name))
	}
	for _, dep := range plan.Indexes {
		blocked = append(blocked, fmt.Sprintf("index:%s.%s", dep.Table, dep.Name))
	}
	sort.Strings(blocked)
	return blocked, nil
}

func detectUnnamedConstraintDependencyBlocks(object string, columns []string, previousCreateTable string, historyState map[string]historyObject) ([]string, error) {
	if len(columns) == 0 {
		return []string{}, nil
	}

	removeSet := make(map[string]struct{}, len(columns))
	for _, col := range columns {
		name := strings.TrimSpace(col)
		if name == "" {
			continue
		}
		removeSet[name] = struct{}{}
	}
	if len(removeSet) == 0 {
		return []string{}, nil
	}

	localUnnamed, err := extractUnnamedConstraintBlocksForColumns(object, previousCreateTable, removeSet)
	if err != nil {
		return nil, err
	}
	incomingUnnamed, err := extractIncomingUnnamedForeignKeyBlocksFromHistory(object, removeSet, historyState)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]struct{}, len(localUnnamed)+len(incomingUnnamed))
	out := make([]string, 0, len(localUnnamed)+len(incomingUnnamed))
	for _, item := range append(localUnnamed, incomingUnnamed...) {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}
		if _, exists := seen[item]; exists {
			continue
		}
		seen[item] = struct{}{}
		out = append(out, item)
	}
	sort.Strings(out)
	return out, nil
}

func extractUnnamedConstraintBlocksForColumns(object string, createTableStmt string, removeSet map[string]struct{}) ([]string, error) {
	items, err := extractCreateTableItems(createTableStmt)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]struct{})
	out := make([]string, 0)
	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		upper := strings.ToUpper(line)
		hasConstraintKeyword := strings.Contains(upper, "CONSTRAINT ")
		if !hasConstraintKeyword && !isConstraintLikeTableItem(upper) {
			continue
		}

		impacted := false
		for col := range removeSet {
			if containsColumnReference(line, col) {
				impacted = true
				break
			}
		}
		if !impacted {
			continue
		}

		if _, named := extractNamedConstraintName(line); named {
			continue
		}

		kind := constraintKindForTableItem(line)
		descriptor := fmt.Sprintf("%s:%s.<unnamed:%s>", kind, object, kind)
		if _, exists := seen[descriptor]; exists {
			continue
		}
		seen[descriptor] = struct{}{}
		out = append(out, descriptor)
	}

	sort.Strings(out)
	return out, nil
}

func extractIncomingUnnamedForeignKeyBlocksFromHistory(object string, removeSet map[string]struct{}, historyState map[string]historyObject) ([]string, error) {
	keys := make([]string, 0, len(historyState))
	for key := range historyState {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	seen := make(map[string]struct{})
	out := make([]string, 0)
	for _, key := range keys {
		st := historyState[key]
		if st.Category != "tables" || strings.TrimSpace(st.SemanticSQL) == "" {
			continue
		}

		items, err := extractCreateTableItems(st.SemanticSQL)
		if err != nil {
			return nil, err
		}
		for _, raw := range items {
			line := strings.TrimSpace(raw)
			if line == "" {
				continue
			}
			upper := strings.ToUpper(line)
			if !strings.Contains(upper, "FOREIGN KEY") || !strings.Contains(upper, "REFERENCES ") {
				continue
			}
			if _, named := extractNamedConstraintName(line); named {
				continue
			}

			refTable, refCols, parseErr := parseForeignKeyConstraintReference(line)
			if parseErr != nil {
				return nil, parseErr
			}
			if !sameQualifiedObject(refTable, object) {
				continue
			}

			impacted := false
			for _, col := range refCols {
				if _, ok := removeSet[strings.TrimSpace(col)]; ok {
					impacted = true
					break
				}
			}
			if !impacted {
				continue
			}

			descriptor := fmt.Sprintf("foreign_key:%s.<unnamed:foreign_key>", st.Object)
			if _, exists := seen[descriptor]; exists {
				continue
			}
			seen[descriptor] = struct{}{}
			out = append(out, descriptor)
		}
	}

	sort.Strings(out)
	return out, nil
}

func buildDropColumnDownRestoreSQL(previousCreateTable, object string, removedColumns []string, depPlan dropColumnDependencyPlan, historyState map[string]historyObject) (string, error) {
	reAddColumnsSQL, err := buildReAddColumnsDownSQL(previousCreateTable, object, removedColumns)
	if err != nil {
		return "", err
	}
	reAddConstraintsSQL, err := buildReAddDroppedConstraintsDownSQL(previousCreateTable, object, depPlan.Constraints, historyState)
	if err != nil {
		return "", err
	}
	reAddIndexesSQL, err := buildReAddDroppedIndexesDownSQL(depPlan.Indexes, historyState)
	if err != nil {
		return "", err
	}

	segments := make([]string, 0, 2)
	if strings.TrimSpace(reAddColumnsSQL) != "" {
		segments = append(segments, strings.TrimSpace(reAddColumnsSQL))
	}
	restoreDependencies := concatSQLSegments(reAddConstraintsSQL, reAddIndexesSQL)
	if strings.TrimSpace(restoreDependencies) != "" {
		// SQL Server compiles a batch before executing it. A FK/CHECK/index that
		// references a column restored earlier in the same batch can therefore
		// fail with "Invalid column name". Compile dependency restoration only
		// after ADD COLUMN has executed.
		if strings.TrimSpace(reAddColumnsSQL) != "" {
			restoreDependencies = dynamicSQLBatch(restoreDependencies)
		}
		segments = append(segments, strings.TrimSpace(restoreDependencies))
	}
	if len(segments) == 0 {
		return "", nil
	}
	return normalizeSQLText(strings.Join(segments, "\n\n"), true, true), nil
}

func buildDropDependencySQL(depPlan dropColumnDependencyPlan) string {
	segments := make([]string, 0, len(depPlan.Constraints)+len(depPlan.Indexes))
	for _, dep := range depPlan.Constraints {
		segments = append(segments, buildDropConstraintIfExistsStatement(dep.Table, dep.Name, dep.Kind))
	}
	for _, dep := range depPlan.Indexes {
		segments = append(segments, buildDropIndexDependencyStatement(dep.Table, dep.Name))
	}
	return concatSQLSegments(segments...)
}

func buildReAddDroppedDependenciesSQL(previousCreateTable, object string, depPlan dropColumnDependencyPlan, historyState map[string]historyObject) (string, error) {
	reAddConstraintsSQL, err := buildReAddDroppedConstraintsDownSQL(previousCreateTable, object, depPlan.Constraints, historyState)
	if err != nil {
		return "", err
	}
	reAddIndexesSQL, err := buildReAddDroppedIndexesDownSQL(depPlan.Indexes, historyState)
	if err != nil {
		return "", err
	}
	return concatSQLSegments(reAddConstraintsSQL, reAddIndexesSQL), nil
}

func filterAutoManagedAlterColumnConstraints(drops []dropConstraintTarget, object string) []dropConstraintTarget {
	out := make([]dropConstraintTarget, 0, len(drops))
	for _, dep := range drops {
		if strings.EqualFold(strings.TrimSpace(dep.Kind), "default") && sameQualifiedObject(dep.Table, object) {
			continue
		}
		out = append(out, dep)
	}
	return out
}

// excludeColumnLifecycleDefaultsFromDependencyRestore prevents a changed
// DEFAULT from being restored twice on down. buildTableAlterColumnsSQL owns
// DEFAULT changes, while the dependency plan owns unchanged defaults that only
// had to be removed temporarily to permit ALTER COLUMN.
func excludeColumnLifecycleDefaultsFromDependencyRestore(plan dropColumnDependencyPlan, previousCreateTable string, defaultChangedColumns []string) (dropColumnDependencyPlan, error) {
	if len(defaultChangedColumns) == 0 || len(plan.Constraints) == 0 {
		return plan, nil
	}
	defaults, err := extractTableDefaultConstraintMap(previousCreateTable)
	if err != nil {
		return dropColumnDependencyPlan{}, err
	}
	changed := make(map[string]struct{}, len(defaultChangedColumns))
	for _, column := range defaultChangedColumns {
		changed[strings.TrimSpace(column)] = struct{}{}
	}
	ownedNames := make(map[string]struct{})
	for column, def := range defaults {
		if _, ok := changed[strings.TrimSpace(column)]; !ok {
			continue
		}
		name := stripIdentifierBrackets(strings.TrimSpace(def.ConstraintName))
		if name != "" {
			ownedNames[name] = struct{}{}
		}
	}
	filtered := make([]dropConstraintTarget, 0, len(plan.Constraints))
	for _, dependency := range plan.Constraints {
		if strings.EqualFold(strings.TrimSpace(dependency.Kind), "default") {
			if _, owned := ownedNames[stripIdentifierBrackets(strings.TrimSpace(dependency.Name))]; owned {
				continue
			}
		}
		filtered = append(filtered, dependency)
	}
	plan.Constraints = filtered
	return plan, nil
}

func hasUnsafeAlterColumnDependency(depPlan dropColumnDependencyPlan) bool {
	for _, dep := range depPlan.Constraints {
		switch strings.ToLower(strings.TrimSpace(dep.Kind)) {
		case "check", "default", "unique":
			continue
		default:
			return true
		}
	}
	return false
}

func findUnsafeAdvancedAlterIndexDependencies(depPlan dropColumnDependencyPlan, historyState map[string]historyObject) ([]string, error) {
	blocked := make([]string, 0)
	for _, dep := range depPlan.Indexes {
		key, ok := findHistoryIndexKey(historyState, dep.Table)
		if !ok {
			return nil, fmt.Errorf("index state for dependency not found: %s.%s", dep.Table, dep.Name)
		}
		state, err := parseIndexStateMap(historyState[key].Object, historyState[key].SemanticSQL)
		if err != nil {
			return nil, err
		}
		definition, ok := state[stripIdentifierBrackets(dep.Name)]
		if !ok {
			return nil, fmt.Errorf("index definition for dependency not found: %s.%s", dep.Table, dep.Name)
		}
		upper := strings.ToUpper(definition.Statement)
		if strings.Contains(upper, " XML INDEX ") || strings.Contains(upper, " SPATIAL INDEX ") || strings.Contains(upper, " COLUMNSTORE INDEX ") {
			blocked = append(blocked, dep.Table+"."+dep.Name)
		}
	}
	sort.Strings(blocked)
	return blocked, nil
}

func buildReAddDroppedConstraintsDownSQL(previousCreateTable, object string, drops []dropConstraintTarget, historyState map[string]historyObject) (string, error) {
	if len(drops) == 0 {
		return "", nil
	}

	localMap, err := extractNamedConstraintDefinitionMap(previousCreateTable)
	if err != nil {
		return "", err
	}

	restore := make([]string, 0, len(drops))
	seen := make(map[string]struct{}, len(drops))
	for i := len(drops) - 1; i >= 0; i-- {
		dep := drops[i]
		key := stripIdentifierBrackets(dep.Table) + "|" + stripIdentifierBrackets(dep.Name)
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}

		var definition string
		if sameQualifiedObject(dep.Table, object) {
			if strings.EqualFold(strings.TrimSpace(dep.Kind), "default") {
				defaults, defaultErr := extractTableDefaultConstraintMap(previousCreateTable)
				if defaultErr != nil {
					return "", defaultErr
				}
				for _, candidate := range defaults {
					if stripIdentifierBrackets(candidate.ConstraintName) != stripIdentifierBrackets(dep.Name) {
						continue
					}
					definition, defaultErr = buildDefaultConstraintDefinitionSQL(candidate)
					if defaultErr != nil {
						return "", defaultErr
					}
					break
				}
			} else {
				definition = localMap[stripIdentifierBrackets(dep.Name)]
			}
		} else {
			tableKey, ok := findHistoryTableKey(historyState, dep.Table)
			if !ok {
				return "", fmt.Errorf("table state for dependency restore not found: %s", dep.Table)
			}
			tableState := historyState[tableKey]
			depMap, depErr := extractNamedConstraintDefinitionMap(tableState.SemanticSQL)
			if depErr != nil {
				return "", depErr
			}
			definition = depMap[stripIdentifierBrackets(dep.Name)]
		}
		if strings.TrimSpace(definition) == "" {
			return "", fmt.Errorf("constraint definition not found for restore: %s.%s", dep.Table, dep.Name)
		}
		addSQL, addErr := buildAddConstraintWithDataPrecheckStatement(dep.Table, dep.Name, dep.Kind, definition)
		if addErr != nil {
			return "", fmt.Errorf("build guarded restore for %s.%s: %w", dep.Table, dep.Name, addErr)
		}
		restore = append(restore, addSQL)
	}
	return normalizeSQLText(strings.Join(restore, "\n"), true, true), nil
}

func buildReAddDroppedIndexesDownSQL(drops []dropIndexTarget, historyState map[string]historyObject) (string, error) {
	if len(drops) == 0 {
		return "", nil
	}

	restore := make([]string, 0, len(drops))
	seen := make(map[string]struct{}, len(drops))
	for _, dep := range drops {
		key := stripIdentifierBrackets(dep.Table) + "|" + stripIdentifierBrackets(dep.Name)
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}

		indexKey, ok := findHistoryIndexKey(historyState, dep.Table)
		if !ok {
			return "", fmt.Errorf("index state for restore not found: %s", dep.Table)
		}
		indexState := historyState[indexKey]
		stateMap, err := parseIndexStateMap(indexState.Object, indexState.SemanticSQL)
		if err != nil {
			return "", err
		}
		def, ok := stateMap[stripIdentifierBrackets(dep.Name)]
		if !ok {
			return "", fmt.Errorf("index definition not found for restore: %s.%s", dep.Table, dep.Name)
		}

		restore = append(restore, buildCreateIndexIfNotExistsStatement(dep.Table, dep.Name, def.Statement))
	}
	return normalizeSQLText(strings.Join(restore, "\n"), true, true), nil
}

func constraintDropPriority(kind string) int {
	switch strings.ToLower(strings.TrimSpace(kind)) {
	case "foreign_key":
		return 0
	case "default":
		return 1
	case "check":
		return 2
	case "unique":
		return 3
	case "primary_key":
		return 4
	default:
		return 5
	}
}

func extractIncomingForeignKeyDropsFromHistory(object string, removeSet map[string]struct{}, historyState map[string]historyObject) ([]dropConstraintTarget, error) {
	keys := make([]string, 0, len(historyState))
	for key := range historyState {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	seen := make(map[string]dropConstraintTarget)
	for _, key := range keys {
		st := historyState[key]
		if st.Category != "tables" || strings.TrimSpace(st.SemanticSQL) == "" {
			continue
		}

		refs, err := extractForeignKeyConstraints(st.Object, st.SemanticSQL)
		if err != nil {
			return nil, err
		}
		for _, ref := range refs {
			if !sameQualifiedObject(ref.ReferencedTable, object) {
				continue
			}
			impacted := false
			for _, col := range ref.ReferencedColumns {
				if _, ok := removeSet[strings.TrimSpace(col)]; ok {
					impacted = true
					break
				}
			}
			if !impacted {
				continue
			}

			seenKey := st.Object + "|" + ref.Name
			seen[seenKey] = dropConstraintTarget{
				Table: st.Object,
				Name:  ref.Name,
				Kind:  "foreign_key",
			}
		}
	}

	out := make([]dropConstraintTarget, 0, len(seen))
	for _, item := range seen {
		out = append(out, item)
	}
	return out, nil
}

func extractForeignKeyConstraints(object, createTableStmt string) ([]foreignKeyConstraintReference, error) {
	items, err := extractCreateTableItems(createTableStmt)
	if err != nil {
		return nil, err
	}

	out := make([]foreignKeyConstraintReference, 0)
	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		upper := strings.ToUpper(line)
		if !strings.Contains(upper, "FOREIGN KEY") || !strings.Contains(upper, "REFERENCES ") {
			continue
		}

		name, ok := extractNamedConstraintName(line)
		if !ok {
			return nil, fmt.Errorf("foreign key constraint must be named for deterministic drop in %s: %s", object, line)
		}
		refTable, refColumns, parseErr := parseForeignKeyConstraintReference(line)
		if parseErr != nil {
			return nil, parseErr
		}
		out = append(out, foreignKeyConstraintReference{
			Name:              name,
			ReferencedTable:   refTable,
			ReferencedColumns: refColumns,
		})
	}
	return out, nil
}

func parseForeignKeyConstraintReference(line string) (string, []string, error) {
	normalized := normalizeSQLText(line, true, true)
	upper := strings.ToUpper(normalized)
	refIdx := strings.Index(upper, "REFERENCES ")
	if refIdx < 0 {
		return "", nil, fmt.Errorf("invalid FOREIGN KEY constraint: REFERENCES target missing")
	}

	rest := strings.TrimSpace(normalized[refIdx+len("REFERENCES "):])
	if rest == "" {
		return "", nil, fmt.Errorf("invalid FOREIGN KEY constraint: REFERENCES target missing")
	}

	tableToken, consumed := consumeLeadingObjectToken(rest)
	if tableToken == "" {
		return "", nil, fmt.Errorf("invalid FOREIGN KEY constraint: referenced table missing")
	}
	refTable := stripIdentifierBrackets(strings.TrimSpace(strings.TrimRight(tableToken, ",;")))
	if refTable == "" {
		return "", nil, fmt.Errorf("invalid FOREIGN KEY constraint: referenced table missing")
	}

	afterTable := strings.TrimSpace(rest[consumed:])
	if afterTable == "" || !strings.HasPrefix(afterTable, "(") {
		return "", nil, fmt.Errorf("invalid FOREIGN KEY constraint: referenced columns missing for %s", refTable)
	}

	refColumnsRaw, ok := extractLeadingParenthesizedContent(afterTable)
	if !ok {
		return "", nil, fmt.Errorf("invalid FOREIGN KEY constraint: malformed referenced columns for %s", refTable)
	}
	refColumns := parseIdentifierList(refColumnsRaw)
	if len(refColumns) == 0 {
		return "", nil, fmt.Errorf("invalid FOREIGN KEY constraint: referenced columns missing for %s", refTable)
	}

	return refTable, refColumns, nil
}

func consumeLeadingObjectToken(input string) (string, int) {
	if input == "" {
		return "", 0
	}
	i := 0
	for i < len(input) {
		switch input[i] {
		case ' ', '\t', '\n', '\r', '(', ',', ';':
			return input[:i], i
		default:
			i++
		}
	}
	return input, len(input)
}

func extractLeadingParenthesizedContent(input string) (string, bool) {
	trimmed := strings.TrimSpace(input)
	if !strings.HasPrefix(trimmed, "(") {
		return "", false
	}

	close, ok := sqlparse.FindMatchingParen(trimmed, 0)
	if !ok {
		return "", false
	}
	return trimmed[1:close], true
}

func parseIdentifierList(raw string) []string {
	parts := splitTopLevelByComma(raw)
	out := make([]string, 0, len(parts))
	for _, part := range parts {
		token := firstSQLIdentifierToken(strings.TrimSpace(part))
		name := stripIdentifierBrackets(cleanIdentifierToken(token))
		if name == "" {
			continue
		}
		// Preserve the source spelling. Identifier comparison may be
		// case-insensitive, but emitted SQL must also work on databases whose
		// catalog collation is case-sensitive.
		out = append(out, name)
	}
	return out
}

func sameQualifiedObject(a, b string) bool {
	return stripIdentifierBrackets(a) == stripIdentifierBrackets(b)
}

func extractNamedConstraintDropsForRemovedColumns(object string, createTableStmt string, removeSet map[string]struct{}) ([]dropConstraintTarget, error) {
	items, err := extractCreateTableItems(createTableStmt)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]dropConstraintTarget)
	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		upper := strings.ToUpper(line)
		hasConstraintKeyword := strings.Contains(upper, "CONSTRAINT ")
		if !hasConstraintKeyword && !isConstraintLikeTableItem(upper) {
			continue
		}

		impacted := false
		for col := range removeSet {
			if containsColumnReference(line, col) {
				impacted = true
				break
			}
		}
		if !impacted {
			continue
		}

		name, ok := extractNamedConstraintName(line)
		if !ok {
			return nil, fmt.Errorf("drop-column dependency requires named constraint in %s: %s", object, line)
		}
		key := object + "|" + name
		seen[key] = dropConstraintTarget{
			Table: object,
			Name:  name,
			Kind:  constraintKindForTableItem(line),
		}
	}

	out := make([]dropConstraintTarget, 0, len(seen))
	for _, item := range seen {
		out = append(out, item)
	}
	return out, nil
}

func extractCreateTableItems(createTableStmt string) ([]string, error) {
	stmt := normalizeSQLText(createTableStmt, true, true)
	openIdx := strings.Index(stmt, "(")
	if openIdx < 0 {
		return nil, fmt.Errorf("invalid CREATE TABLE statement: no opening parenthesis")
	}

	closeIdx, ok := sqlparse.FindMatchingParen(stmt, openIdx)
	if !ok {
		return nil, fmt.Errorf("invalid CREATE TABLE statement: no closing parenthesis")
	}

	inside := stmt[openIdx+1 : closeIdx]
	items := splitTopLevelByComma(inside)
	if err := validateCreateTableCommaSeparators(items); err != nil {
		return nil, err
	}
	return items, nil
}

func extractNamedConstraintDefinitionMap(createTableStmt string) (map[string]string, error) {
	items, err := extractCreateTableItems(createTableStmt)
	if err != nil {
		return nil, err
	}

	out := make(map[string]string)
	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		name, ok := extractNamedConstraintName(line)
		if !ok {
			continue
		}
		key := stripIdentifierBrackets(name)
		if existing, exists := out[key]; exists {
			if strings.TrimSpace(existing) == strings.TrimSpace(line) {
				continue
			}
			return nil, fmt.Errorf("duplicate constraint definition detected for %s", stripIdentifierBrackets(name))
		}
		out[key] = line
	}
	return out, nil
}

func constraintKindForTableItem(line string) string {
	upper := strings.ToUpper(line)
	switch {
	case strings.Contains(upper, "FOREIGN KEY"):
		return "foreign_key"
	case strings.Contains(upper, "PRIMARY KEY"):
		return "primary_key"
	case strings.Contains(upper, "UNIQUE"):
		return "unique"
	case strings.Contains(upper, "CHECK"):
		return "check"
	case strings.Contains(upper, " DEFAULT ") || strings.Contains(upper, " DEFAULT("):
		return "default"
	default:
		return "constraint"
	}
}

func extractNamedConstraintName(line string) (string, bool) {
	trimmed := strings.TrimSpace(line)
	upper := strings.ToUpper(trimmed)
	idx := strings.Index(upper, "CONSTRAINT ")
	if idx < 0 {
		return "", false
	}
	rest := strings.TrimSpace(trimmed[idx+len("CONSTRAINT "):])
	if rest == "" {
		return "", false
	}

	nameToken := firstSQLToken(rest)
	name := cleanIdentifierToken(nameToken)
	if name == "" {
		return "", false
	}
	return name, true
}

func extractCheckConstraintExpression(definition string) (string, error) {
	normalized := normalizeSQLText(definition, true, true)
	upper := strings.ToUpper(normalized)
	checkPositions := findAllKeywordPositions(upper, "CHECK")
	if len(checkPositions) == 0 {
		return "", fmt.Errorf("CHECK constraint expression not found")
	}
	after := strings.TrimSpace(normalized[checkPositions[0]+len("CHECK"):])
	afterUpper := strings.ToUpper(after)
	if strings.HasPrefix(afterUpper, "NOT FOR REPLICATION") {
		after = strings.TrimSpace(after[len("NOT FOR REPLICATION"):])
	}
	if !strings.HasPrefix(after, "(") {
		return "", fmt.Errorf("CHECK constraint expression must be parenthesized")
	}
	expr, ok := extractLeadingParenthesizedContent(after)
	if !ok || strings.TrimSpace(expr) == "" {
		return "", fmt.Errorf("CHECK constraint expression is empty")
	}
	return strings.TrimSpace(expr), nil
}

func extractImpactedIndexDropsForRemovedColumns(object, indexSemantic string, removeSet map[string]struct{}) ([]dropIndexTarget, error) {
	if strings.TrimSpace(indexSemantic) == "" {
		return []dropIndexTarget{}, nil
	}

	stmts, err := extractCreateIndexStatements(object, indexSemantic)
	if err != nil {
		return nil, err
	}

	seen := make(map[string]dropIndexTarget)
	for _, stmt := range stmts {
		impacted := false
		for col := range removeSet {
			if containsColumnReference(stmt, col) {
				impacted = true
				break
			}
		}
		if !impacted {
			continue
		}

		indexName, parseErr := parseCreateIndexName(stmt)
		if parseErr != nil {
			return nil, parseErr
		}
		key := object + "|" + indexName
		seen[key] = dropIndexTarget{
			Table: object,
			Name:  indexName,
		}
	}

	out := make([]dropIndexTarget, 0, len(seen))
	for _, item := range seen {
		out = append(out, item)
	}
	return out, nil
}

func parseCreateIndexName(stmt string) (string, error) {
	normalized := normalizeSQLText(stmt, true, true)
	upper := strings.ToUpper(normalized)
	indexPos := strings.Index(upper, "INDEX ")
	if indexPos < 0 {
		return "", fmt.Errorf("invalid CREATE INDEX statement: INDEX keyword not found")
	}

	rest := strings.TrimSpace(normalized[indexPos+len("INDEX "):])
	if rest == "" {
		return "", fmt.Errorf("invalid CREATE INDEX statement: missing index name")
	}
	onPositions := findAllKeywordPositions(strings.ToUpper(maskBracketedIdentifiers(rest)), "ON")
	if len(onPositions) == 0 {
		return "", fmt.Errorf("invalid CREATE INDEX statement: ON target not found")
	}
	onPos := onPositions[0]

	indexName := cleanIdentifierToken(rest[:onPos])
	if indexName == "" {
		return "", fmt.Errorf("invalid CREATE INDEX statement: empty index name")
	}
	return indexName, nil
}

func cleanIdentifierToken(token string) string {
	trimmed := strings.TrimSpace(token)
	trimmed = strings.TrimRight(trimmed, ",);")
	if strings.HasPrefix(trimmed, "[") && strings.HasSuffix(trimmed, "]") && len(trimmed) >= 2 {
		trimmed = trimmed[1 : len(trimmed)-1]
		trimmed = strings.ReplaceAll(trimmed, "]]", "]")
	}
	return strings.TrimSpace(trimmed)
}

func stripIdentifierBrackets(input string) string {
	return strings.TrimSpace(strings.ReplaceAll(strings.ReplaceAll(input, "[", ""), "]", ""))
}

func quoteIdentifier(name string) string {
	outer := strings.TrimSpace(name)
	clean := outer
	if outer == "" && name != "" {
		clean = name // SQL Server permits delimited identifiers containing spaces.
	}
	if strings.HasPrefix(outer, "[") && strings.HasSuffix(outer, "]") && len(outer) >= 2 {
		clean = outer[1 : len(outer)-1]
		clean = strings.ReplaceAll(clean, "]]", "]")
	}
	return "[" + strings.ReplaceAll(clean, "]", "]]") + "]"
}

func quoteQualifiedIdentifier(name string) string {
	trimmed := strings.TrimSpace(name)
	if trimmed == "" {
		return ""
	}
	parts := make([]string, 0, 2)
	start := 0
	inBracket := false
	for i := 0; i < len(trimmed); i++ {
		switch trimmed[i] {
		case '[':
			inBracket = true
		case ']':
			if inBracket && i+1 < len(trimmed) && trimmed[i+1] == ']' {
				i++
			} else {
				inBracket = false
			}
		case '.':
			if !inBracket {
				parts = append(parts, trimmed[start:i])
				start = i + 1
			}
		}
	}
	parts = append(parts, trimmed[start:])
	for i := range parts {
		clean := cleanIdentifierToken(parts[i])
		if isSafeBareIdentifier(clean) {
			parts[i] = clean
		} else {
			parts[i] = quoteIdentifier(clean)
		}
	}
	return strings.Join(parts, ".")
}

func isSafeBareIdentifier(name string) bool {
	if name == "" || !((name[0] >= 'A' && name[0] <= 'Z') || (name[0] >= 'a' && name[0] <= 'z') || name[0] == '_') {
		return false
	}
	for i := 1; i < len(name); i++ {
		ch := name[i]
		if !((ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '$' || ch == '#') {
			return false
		}
	}
	switch strings.ToUpper(name) {
	case "ADD", "ALTER", "COLUMN", "CONSTRAINT", "CREATE", "DEFAULT", "DELETE", "DROP", "FROM", "FUNCTION", "GROUP", "INDEX", "INSERT", "KEY", "ORDER", "PRIMARY", "PROCEDURE", "REFERENCES", "SELECT", "TABLE", "UNIQUE", "UPDATE", "USER", "VALUES", "VIEW", "WHERE":
		return false
	}
	return true
}

func buildDropColumnUpSQL(object string, removedColumns []string, depPlan dropColumnDependencyPlan) string {
	if len(removedColumns) == 0 {
		return ""
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("IF OBJECT_ID(N'%s', N'U') IS NULL\nBEGIN\n", object))
	b.WriteString("  RETURN;\n")
	b.WriteString("END\n\n")

	for _, dep := range depPlan.Constraints {
		b.WriteString(buildDropConstraintIfExistsStatement(dep.Table, dep.Name, dep.Kind) + "\n")
	}
	for _, dep := range depPlan.Indexes {
		b.WriteString(buildDropIndexDependencyStatement(dep.Table, dep.Name) + "\n")
	}

	for _, col := range removedColumns {
		b.WriteString(buildDropColumnIfExistsStatement(object, col) + "\n")
	}

	return normalizeSQLText(b.String(), true, true)
}

func buildDropColumnIfExistsStatement(object, column string) string {
	target := object + "." + column
	runLog := indentSQL(buildStepAuditLogSQL("run", 1, "drop_column", target, "OBJECT_EXISTS", "drop column start"), "  ")
	okLog := indentSQL(buildStepAuditLogSQL("ok", 1, "drop_column", target, "DROPPED", "drop column ok"), "  ")
	skipLog := indentSQL(buildStepAuditLogSQL("skip", 0, "drop_column", target, "COLUMN_NOT_EXISTS", "Precheck: target column(s) already absent"), "  ")
	return normalizeSQLText(fmt.Sprintf(`-- check column exists before DROP COLUMN
IF (%s)
BEGIN
%s
  IF EXISTS (SELECT 1 FROM %s WHERE %s IS NOT NULL)
  BEGIN
    ;THROW 51000, N'Precheck blocked: DROP COLUMN target contains data', 1;
  END

  ALTER TABLE %s DROP COLUMN %s;
%s
END
ELSE
BEGIN
%s
END`, buildExactColumnExistsCondition(object, column),
		runLog,
		quoteQualifiedIdentifier(object), quoteIdentifier(column),
		quoteQualifiedIdentifier(object), quoteIdentifier(column),
		okLog,
		skipLog), true, true)
}

func buildDropDerivedColumnIfExistsStatement(object, column string) string {
	target := object + "." + column
	runLog := indentSQL(buildStepAuditLogSQL("run", 1, "drop_computed_column", target, "OBJECT_EXISTS", "drop computed column start"), "  ")
	okLog := indentSQL(buildStepAuditLogSQL("ok", 1, "drop_computed_column", target, "DROPPED", "drop computed column ok"), "  ")
	skipLog := indentSQL(buildStepAuditLogSQL("skip", 0, "drop_computed_column", target, "COLUMN_NOT_EXISTS", "Precheck: computed column already absent"), "  ")
	return normalizeSQLText(fmt.Sprintf(`-- computed values are derived; dependency checks are planned separately
IF (%s)
BEGIN
%s
  IF NOT EXISTS (
    SELECT 1
    FROM sys.computed_columns AS cc
    WHERE cc.object_id = OBJECT_ID(N'%s', N'U')
      AND cc.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
  )
  BEGIN
    ;THROW 51000, N'Schema drift: expected computed column is a regular column', 1;
  END

  ALTER TABLE %s DROP COLUMN %s;
%s
END
ELSE
BEGIN
%s
END`, buildExactColumnExistsCondition(object, column), runLog,
		escapeSQLLiteral(object), escapeSQLLiteral(column), quoteQualifiedIdentifier(object), quoteIdentifier(column), okLog, skipLog), true, true)
}

func buildDropIndexDependencyStatement(objectName, indexName string) string {
	cleanName := stripIdentifierBrackets(cleanIdentifierToken(indexName))
	if cleanName == "" {
		return ""
	}
	return buildDropIndexStatement(objectName, cleanName)
}

func buildAddColumnIfMissingStatement(object string, def tableColumnDef) string {
	target := object + "." + def.Name
	runLog := indentSQL(buildStepAuditLogSQL("run", 1, "add_column", target, "OBJECT_EXISTS", "add column start"), "  ")
	okLog := indentSQL(buildStepAuditLogSQL("ok", 1, "add_column", target, "CREATED", "add column ok"), "  ")
	skipLog := indentSQL(buildStepAuditLogSQL("skip", 0, "add_column", target, "COLUMN_EXISTS", "Precheck: target column(s) already exist"), "  ")
	return normalizeSQLText(fmt.Sprintf(`%s

-- check exact-case column does not exist before ADD COLUMN
IF NOT (%s)
BEGIN
%s
  ALTER TABLE %s ADD %s;
%s
END
ELSE
BEGIN
  IF NOT (%s)
  BEGIN
    ;THROW 51000, N'Schema drift: existing column does not match the requested definition', 1;
  END
%s
END`, buildRequireTableExistsSQL(object, "ADD COLUMN"), buildExactColumnExistsCondition(object, def.Name),
		runLog,
		quoteQualifiedIdentifier(object), columnDefinitionSQL(def),
		okLog,
		buildColumnEquivalentCondition(object, def),
		skipLog), true, true)
}

func buildDropConstraintIfExistsStatement(table, name, kind string) string {
	condition := buildConstraintExistsCondition(table, name, kind)
	return normalizeSQLText(fmt.Sprintf(`IF (%s)
BEGIN
  ALTER TABLE %s DROP CONSTRAINT %s;
END`, condition, quoteQualifiedIdentifier(table), quoteIdentifier(name)), true, true)
}

func buildAddConstraintIfMissingStatement(table, name, kind, definition string) string {
	condition := buildConstraintExistsCondition(table, name, kind)
	return normalizeSQLText(fmt.Sprintf(`%s

IF NOT (%s)
BEGIN
  ALTER TABLE %s ADD %s;
END
ELSE
BEGIN
  ;THROW 51000, N'Schema drift: existing constraint definition was not proven equivalent', 1;
END`, buildRequireTableExistsSQL(table, "ADD CONSTRAINT"), condition, quoteQualifiedIdentifier(table), definition), true, true)
}

// buildAddConstraintWithDataPrecheckStatement is the single dispatch point for
// restoring named constraints. Existence-only guards are enough for DEFAULT,
// while CHECK/key/FK constraints must also prove that current rows satisfy the
// constraint before ALTER TABLE is attempted.
func buildAddConstraintWithDataPrecheckStatement(table, name, kind, definition string) (string, error) {
	switch strings.ToLower(strings.TrimSpace(kind)) {
	case "default":
		return buildAddConstraintIfMissingStatement(table, name, kind, definition), nil
	case "check":
		return buildAddCheckConstraintWithPrecheckStatement(table, name, definition)
	case "unique", "primary_key":
		return buildAddKeyConstraintWithPrecheckStatement(table, name, strings.ToLower(strings.TrimSpace(kind)), definition)
	case "foreign_key":
		return buildAddForeignKeyWithPrecheckStatement(table, name, definition)
	default:
		return "", fmt.Errorf("unsupported constraint kind for guarded restore: %s", kind)
	}
}

func buildAddCheckConstraintWithPrecheckStatement(table, name, definition string) (string, error) {
	expr, err := extractCheckConstraintExpression(definition)
	if err != nil {
		return "", err
	}
	condition := buildConstraintExistsCondition(table, name, "check")
	return normalizeSQLText(fmt.Sprintf(`%s

IF NOT (%s)
BEGIN
  IF EXISTS (
    SELECT 1
    FROM %s
    WHERE NOT (%s)
  )
  BEGIN
    ;THROW 51000, 'CHECK constraint %s would fail existing data', 1;
  END

  ALTER TABLE %s ADD %s;
END
ELSE
BEGIN
  ;THROW 51000, N'Schema drift: existing CHECK constraint definition was not proven equivalent', 1;
END`, buildRequireTableExistsSQL(table, "ADD CHECK CONSTRAINT"), condition, quoteQualifiedIdentifier(table), expr, escapeSQLLiteral(stripIdentifierBrackets(name)), quoteQualifiedIdentifier(table), definition), true, true), nil
}

// buildAddKeyConstraintWithPrecheckStatement makes named UNIQUE and PRIMARY
// KEY transitions automatic only when current data can satisfy the key. The
// same check is emitted in down, so post-up writes cannot create silent drift.
func buildAddKeyConstraintWithPrecheckStatement(table, name, kind, definition string) (string, error) {
	columns, err := extractKeyConstraintColumns(definition, kind)
	if err != nil {
		return "", err
	}
	quoted := make([]string, 0, len(columns))
	for _, column := range columns {
		quoted = append(quoted, quoteIdentifier(column))
	}
	groupBy := strings.Join(quoted, ", ")
	condition := buildConstraintExistsCondition(table, name, kind)
	nullCheck := ""
	if kind == "primary_key" {
		nulls := make([]string, 0, len(quoted))
		for _, column := range quoted {
			nulls = append(nulls, column+" IS NULL")
		}
		nullCheck = fmt.Sprintf("\n  IF EXISTS (SELECT 1 FROM %s WHERE %s)\n  BEGIN\n    ;THROW 51000, 'PRIMARY KEY %s cannot be added because key column contains NULL', 1;\n  END", quoteQualifiedIdentifier(table), strings.Join(nulls, " OR "), escapeSQLLiteral(stripIdentifierBrackets(name)))
	}
	operation := "ADD " + strings.ToUpper(strings.ReplaceAll(kind, "_", " ")) + " CONSTRAINT"
	return normalizeSQLText(fmt.Sprintf(`%s

%s

IF NOT (%s)
BEGIN%s
  IF EXISTS (
    SELECT 1
    FROM %s
    GROUP BY %s
    HAVING COUNT_BIG(*) > 1
  )
  BEGIN
    ;THROW 51000, '%s constraint %s cannot be added because duplicate key values exist', 1;
  END

  ALTER TABLE %s ADD %s;
END
ELSE
BEGIN
  ;THROW 51000, N'Schema drift: existing key constraint definition was not proven equivalent', 1;
END`, buildRequireTableExistsSQL(table, operation), buildRequireColumnsExistSQL(table, columns, operation), condition, nullCheck, quoteQualifiedIdentifier(table), groupBy, strings.ToUpper(strings.ReplaceAll(kind, "_", " ")), escapeSQLLiteral(stripIdentifierBrackets(name)), quoteQualifiedIdentifier(table), definition), true, true), nil
}

func extractKeyConstraintColumns(definition, kind string) ([]string, error) {
	normalized := normalizeSQLText(definition, true, true)
	upper := strings.ToUpper(normalized)
	keyword := "UNIQUE"
	if kind == "primary_key" {
		keyword = "PRIMARY KEY"
	}
	positions := findAllKeywordPositions(upper, keyword)
	if len(positions) == 0 {
		return nil, fmt.Errorf("%s keyword not found in constraint definition", keyword)
	}
	after := strings.TrimSpace(normalized[positions[0]+len(keyword):])
	afterUpper := strings.ToUpper(after)
	if strings.HasPrefix(afterUpper, "CLUSTERED ") {
		after = strings.TrimSpace(after[len("CLUSTERED "):])
	} else if strings.HasPrefix(afterUpper, "NONCLUSTERED ") {
		after = strings.TrimSpace(after[len("NONCLUSTERED "):])
	}
	if !strings.HasPrefix(after, "(") {
		return nil, fmt.Errorf("%s constraint key column list is missing", keyword)
	}
	raw, ok := extractLeadingParenthesizedContent(after)
	if !ok {
		return nil, fmt.Errorf("%s constraint key column list is malformed", keyword)
	}
	columns := parseIdentifierList(raw)
	if len(columns) == 0 {
		return nil, fmt.Errorf("%s constraint key column list is empty", keyword)
	}
	return columns, nil
}

type foreignKeyDefinition struct {
	LocalColumns      []string
	ReferencedTable   string
	ReferencedColumns []string
}

// parseForeignKeyDefinition extracts both sides of a named FK. Cascade and
// trust options stay in the original definition and are replayed verbatim by
// the renderer; only the column mapping is used for an orphan-data precheck.
func parseForeignKeyDefinition(definition string) (foreignKeyDefinition, error) {
	normalized := normalizeSQLText(definition, true, true)
	upper := strings.ToUpper(normalized)
	pos := strings.Index(upper, "FOREIGN KEY")
	if pos < 0 {
		return foreignKeyDefinition{}, fmt.Errorf("FOREIGN KEY keyword not found")
	}
	after := strings.TrimSpace(normalized[pos+len("FOREIGN KEY"):])
	if !strings.HasPrefix(after, "(") {
		return foreignKeyDefinition{}, fmt.Errorf("FOREIGN KEY local column list is missing")
	}
	localRaw, ok := extractLeadingParenthesizedContent(after)
	if !ok {
		return foreignKeyDefinition{}, fmt.Errorf("FOREIGN KEY local column list is malformed")
	}
	local := parseIdentifierList(localRaw)
	if len(local) == 0 {
		return foreignKeyDefinition{}, fmt.Errorf("FOREIGN KEY local column list is empty")
	}
	refTable, refColumns, err := parseForeignKeyConstraintReference(normalized)
	if err != nil {
		return foreignKeyDefinition{}, err
	}
	if len(local) != len(refColumns) {
		return foreignKeyDefinition{}, fmt.Errorf("FOREIGN KEY local/referenced column count differs")
	}
	return foreignKeyDefinition{LocalColumns: local, ReferencedTable: refTable, ReferencedColumns: refColumns}, nil
}

func buildAddForeignKeyWithPrecheckStatement(table, name, definition string) (string, error) {
	fk, err := parseForeignKeyDefinition(definition)
	if err != nil {
		return "", err
	}
	condition := buildConstraintExistsCondition(table, name, "foreign_key")
	nonNull := make([]string, 0, len(fk.LocalColumns))
	joins := make([]string, 0, len(fk.LocalColumns))
	for i := range fk.LocalColumns {
		local := quoteIdentifier(fk.LocalColumns[i])
		ref := quoteIdentifier(fk.ReferencedColumns[i])
		nonNull = append(nonNull, "c."+local+" IS NOT NULL")
		joins = append(joins, "p."+ref+" = c."+local)
	}
	return normalizeSQLText(fmt.Sprintf(`%s

%s

%s

%s

IF NOT (%s)
BEGIN
  IF EXISTS (
    SELECT 1
    FROM %s AS c
    WHERE %s
      AND NOT EXISTS (
        SELECT 1
        FROM %s AS p
        WHERE %s
      )
  )
  BEGIN
    ;THROW 51000, 'FOREIGN KEY %s cannot be added because orphan rows exist', 1;
  END

  ALTER TABLE %s ADD %s;
END
ELSE
BEGIN
  ;THROW 51000, N'Schema drift: existing FOREIGN KEY definition was not proven equivalent', 1;
END`, buildRequireTableExistsSQL(table, "ADD FOREIGN KEY"), buildRequireTableExistsSQL(fk.ReferencedTable, "ADD FOREIGN KEY"), buildRequireColumnsExistSQL(table, fk.LocalColumns, "ADD FOREIGN KEY"), buildRequireColumnsExistSQL(fk.ReferencedTable, fk.ReferencedColumns, "ADD FOREIGN KEY"), condition, quoteQualifiedIdentifier(table), strings.Join(nonNull, " AND "), quoteQualifiedIdentifier(fk.ReferencedTable), strings.Join(joins, " AND "), escapeSQLLiteral(stripIdentifierBrackets(name)), quoteQualifiedIdentifier(table), definition), true, true), nil
}

func buildConstraintExistsCondition(table, name, kind string) string {
	cleanName := stripIdentifierBrackets(cleanIdentifierToken(name))
	if cleanName == "" {
		return "1 = 0"
	}
	fullName := cleanName
	if !strings.Contains(cleanName, ".") {
		schema := "dbo"
		cleanTable := stripIdentifierBrackets(strings.TrimSpace(table))
		if dot := strings.LastIndex(cleanTable, "."); dot > 0 {
			schema = strings.TrimSpace(cleanTable[:dot])
		}
		fullName = schema + "." + cleanName
	}
	safeName := escapeSQLLiteral(fullName)
	switch strings.ToLower(strings.TrimSpace(kind)) {
	case "foreign_key":
		return fmt.Sprintf("OBJECT_ID(N'%s', N'F') IS NOT NULL", safeName)
	case "primary_key":
		return fmt.Sprintf("OBJECT_ID(N'%s', N'PK') IS NOT NULL", safeName)
	case "unique":
		return fmt.Sprintf("OBJECT_ID(N'%s', N'UQ') IS NOT NULL", safeName)
	case "check":
		return fmt.Sprintf("OBJECT_ID(N'%s', N'C') IS NOT NULL", safeName)
	case "default":
		return fmt.Sprintf("OBJECT_ID(N'%s', N'D') IS NOT NULL", safeName)
	default:
		// Fallback for unknown constraint kind in legacy state.
		return fmt.Sprintf("(OBJECT_ID(N'%s', N'F') IS NOT NULL OR OBJECT_ID(N'%s', N'PK') IS NOT NULL OR OBJECT_ID(N'%s', N'UQ') IS NOT NULL OR OBJECT_ID(N'%s', N'C') IS NOT NULL OR OBJECT_ID(N'%s', N'D') IS NOT NULL)", safeName, safeName, safeName, safeName, safeName)
	}
}

func buildRequireTableExistsSQL(table, operation string) string {
	safeTable := escapeSQLLiteral(table)
	safeOperation := escapeSQLLiteral(operation)
	return normalizeSQLText(fmt.Sprintf(`-- check target table exists before %s
IF OBJECT_ID(N'%s', N'U') IS NULL
BEGIN
  ;THROW 51000, N'Precheck blocked: %s target table does not exist: %s', 1;
END`, operation, safeTable, safeOperation, safeTable), true, true)
}

func buildRequireColumnsExistSQL(table string, columns []string, operation string) string {
	parts := make([]string, 0, len(columns))
	seen := make(map[string]struct{}, len(columns))
	for _, column := range columns {
		clean := stripIdentifierBrackets(strings.TrimSpace(column))
		if clean == "" {
			continue
		}
		key := clean
		if _, exists := seen[key]; exists {
			continue
		}
		seen[key] = struct{}{}
		parts = append(parts, "NOT ("+buildExactColumnExistsCondition(table, clean)+")")
	}
	if len(parts) == 0 {
		return ""
	}
	sort.Strings(parts)
	return normalizeSQLText(fmt.Sprintf(`-- check target columns exist before %s
IF %s
BEGIN
  ;THROW 51000, N'Precheck blocked: %s target columns do not exist', 1;
END`, operation, strings.Join(parts, " OR "), escapeSQLLiteral(operation)), true, true)
}

func buildExactColumnExistsCondition(table, column string) string {
	return fmt.Sprintf(`EXISTS (
  SELECT 1
  FROM sys.columns AS c
  WHERE c.object_id = OBJECT_ID(N'%s', N'U')
    AND c.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
)`, escapeSQLLiteral(table), escapeSQLLiteral(column))
}

func buildColumnEquivalentCondition(table string, def tableColumnDef, expectedOrdinal ...int) string {
	base := strings.ToUpper(strings.TrimSpace(def.BaseType))
	if base == "" {
		base, _, _ = parseTypeToken(firstSQLToken(def.RawType))
	}
	conditions := []string{
		fmt.Sprintf("c.object_id = OBJECT_ID(N'%s', N'U')", escapeSQLLiteral(table)),
		fmt.Sprintf("c.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2", escapeSQLLiteral(def.Name)),
	}
	if !def.IsComputed {
		conditions = append(conditions, fmt.Sprintf("UPPER(TYPE_NAME(c.user_type_id)) = N'%s'", escapeSQLLiteral(base)))
	}
	if len(expectedOrdinal) > 0 && expectedOrdinal[0] > 0 {
		conditions = append(conditions, fmt.Sprintf("c.column_id = %d", expectedOrdinal[0]))
	}
	boolInt := func(value bool) int {
		if value {
			return 1
		}
		return 0
	}
	conditions = append(conditions,
		fmt.Sprintf("c.is_computed = %d", boolInt(def.IsComputed)),
		fmt.Sprintf("c.is_identity = %d", boolInt(def.IdentityDefinition != "")),
		fmt.Sprintf("c.is_rowguidcol = %d", boolInt(def.RowGuidCol)),
		fmt.Sprintf("c.is_sparse = %d", boolInt(def.Sparse)),
	)
	if def.IsComputed {
		expression, persisted, ok := parseComputedColumnDefinition(def.ComputedDefinition)
		if !ok {
			conditions = append(conditions, "1 = 0")
		} else {
			conditions = append(conditions,
				fmt.Sprintf(`EXISTS (SELECT 1 FROM sys.computed_columns AS cc WHERE cc.object_id=c.object_id AND cc.column_id=c.column_id AND cc.is_persisted=%d AND %s)`, boolInt(persisted), buildCatalogSQLTextSemanticCondition("cc.definition", expression)),
			)
		}
	} else if isStringType(base) || base == "VARBINARY" || base == "BINARY" {
		conditions = append(conditions, fmt.Sprintf("c.is_nullable = %d", boolInt(def.Nullable)))
		expected := def.Length
		if expected == 0 && (base == "CHAR" || base == "VARCHAR" || base == "NCHAR" || base == "NVARCHAR" || base == "BINARY" || base == "VARBINARY") {
			expected = 1
		}
		if expected > 0 && def.IsUnicode {
			expected *= 2
		}
		if expected != 0 {
			conditions = append(conditions, fmt.Sprintf("c.max_length = %d", expected))
		}
	}
	if !def.IsComputed && !(isStringType(base) || base == "VARBINARY" || base == "BINARY") {
		conditions = append(conditions, fmt.Sprintf("c.is_nullable = %d", boolInt(def.Nullable)))
	}
	if base == "DECIMAL" || base == "NUMERIC" {
		if strings.Contains(def.RawType, "(") {
			if precision, scale, ok := parsePrecisionScale(def.RawType); ok {
				conditions = append(conditions, fmt.Sprintf("c.precision = %d", precision), fmt.Sprintf("c.scale = %d", scale))
			}
		} else {
			conditions = append(conditions, "c.precision = 18", "c.scale = 0")
		}
	}
	if base == "TIME" || base == "DATETIME2" || base == "DATETIMEOFFSET" {
		if strings.Contains(def.RawType, "(") {
			if _, scale, ok := parsePrecisionScale(def.RawType); ok {
				conditions = append(conditions, fmt.Sprintf("c.scale = %d", scale))
			}
		} else {
			conditions = append(conditions, "c.scale = 7")
		}
	}
	if pos := strings.Index(strings.ToUpper(def.RawType), " COLLATE "); pos >= 0 {
		collation := strings.TrimSpace(def.RawType[pos+len(" COLLATE "):])
		conditions = append(conditions, fmt.Sprintf("c.collation_name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2", escapeSQLLiteral(collation)))
	} else if isStringType(base) {
		conditions = append(conditions, "c.collation_name = CONVERT(sysname, DATABASEPROPERTYEX(DB_NAME(), 'Collation'))")
	}
	if def.IdentityDefinition != "" {
		seed, increment, ok := parseIdentitySeedIncrement(def.IdentityDefinition)
		if ok {
			conditions = append(conditions, fmt.Sprintf("EXISTS (SELECT 1 FROM sys.identity_columns AS ic WHERE ic.object_id=c.object_id AND ic.column_id=c.column_id AND CONVERT(NVARCHAR(100),ic.seed_value)=N'%s' AND CONVERT(NVARCHAR(100),ic.increment_value)=N'%s')", escapeSQLLiteral(seed), escapeSQLLiteral(increment)))
		}
	}
	return "EXISTS (SELECT 1 FROM sys.columns AS c WHERE " + strings.Join(conditions, " AND ") + ")"
}

func parseComputedColumnDefinition(definition string) (string, bool, bool) {
	trimmed := strings.TrimSpace(definition)
	upper := strings.ToUpper(trimmed)
	if !hasKeywordAtWordBoundary(upper, 0, "AS") {
		return "", false, false
	}
	body := strings.TrimSpace(trimmed[len("AS"):])
	persisted := false
	bodyUpper := strings.ToUpper(body)
	if positions := findAllKeywordPositions(bodyUpper, "PERSISTED"); len(positions) > 0 {
		pos := positions[len(positions)-1]
		if strings.TrimSpace(bodyUpper[pos:]) == "PERSISTED" {
			body = strings.TrimSpace(body[:pos])
			persisted = true
		}
	}
	if body == "" {
		return "", false, false
	}
	return body, persisted, true
}

func parsePrecisionScale(raw string) (int, int, bool) {
	open := strings.Index(raw, "(")
	close := strings.Index(raw, ")")
	if open < 0 || close <= open {
		return 0, 0, false
	}
	parts := strings.Split(raw[open+1:close], ",")
	first, err := strconv.Atoi(strings.TrimSpace(parts[0]))
	if err != nil {
		return 0, 0, false
	}
	if len(parts) == 1 {
		return first, first, true
	}
	second, err := strconv.Atoi(strings.TrimSpace(parts[1]))
	return first, second, err == nil
}

func parseIdentitySeedIncrement(identity string) (string, string, bool) {
	open := strings.Index(identity, "(")
	close := strings.LastIndex(identity, ")")
	if open < 0 || close <= open {
		return "", "", false
	}
	parts := strings.Split(identity[open+1:close], ",")
	if len(parts) != 2 {
		return "", "", false
	}
	return strings.TrimSpace(parts[0]), strings.TrimSpace(parts[1]), true
}

func concatSQLSegments(parts ...string) string {
	segments := make([]string, 0, len(parts))
	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed == "" {
			continue
		}
		segments = append(segments, trimmed)
	}
	if len(segments) == 0 {
		return ""
	}
	return normalizeSQLText(strings.Join(segments, "\n\n"), true, true)
}

func applyConstraintDropsToHistoryState(historyState map[string]historyObject, drops []dropConstraintTarget) error {
	if len(drops) == 0 {
		return nil
	}

	for _, dep := range drops {
		key, ok := findHistoryTableKey(historyState, dep.Table)
		if !ok {
			continue
		}
		st := historyState[key]
		if strings.TrimSpace(st.SemanticSQL) == "" {
			continue
		}

		updated, err := removeNamedConstraintFromCreateTable(st.SemanticSQL, dep.Name)
		if err != nil {
			return err
		}
		st.SemanticSQL = normalizeSQLText(updated, true, true)
		st.NormalizedHash = hashSemanticByCategory(st.Category, st.SemanticSQL)
		historyState[key] = st
	}

	return nil
}

func applyIndexDropsToHistoryState(historyState map[string]historyObject, drops []dropIndexTarget) error {
	if len(drops) == 0 {
		return nil
	}

	byTable := make(map[string][]dropIndexTarget)
	for _, dep := range drops {
		tableKey := stripIdentifierBrackets(dep.Table)
		byTable[tableKey] = append(byTable[tableKey], dep)
	}

	for _, tableDrops := range byTable {
		tableName := tableDrops[0].Table
		key, ok := findHistoryIndexKey(historyState, tableName)
		if !ok {
			continue
		}
		st := historyState[key]
		if strings.TrimSpace(st.SemanticSQL) == "" {
			continue
		}

		stateMap, err := parseIndexStateMap(st.Object, st.SemanticSQL)
		if err != nil {
			return err
		}
		for _, dep := range tableDrops {
			dropKey := stripIdentifierBrackets(dep.Name)
			delete(stateMap, dropKey)
		}

		st.SemanticSQL = serializeIndexState(stateMap)
		st.NormalizedHash = hashSemanticByCategory(st.Category, st.SemanticSQL)
		historyState[key] = st
	}

	return nil
}

type addedConstraintTarget struct {
	Table      string
	Name       string
	Kind       string
	Definition string
}

func applyConstraintAddsToHistoryState(historyState map[string]historyObject, adds []addedConstraintTarget) error {
	for _, add := range adds {
		key, ok := findHistoryTableKey(historyState, add.Table)
		if !ok {
			continue
		}
		st := historyState[key]
		if strings.TrimSpace(st.SemanticSQL) == "" {
			continue
		}
		updated, err := addOrReplaceNamedConstraintInCreateTable(st.SemanticSQL, add.Name, add.Definition)
		if err != nil {
			return err
		}
		st.SemanticSQL = normalizeSQLText(updated, true, true)
		st.NormalizedHash = hashSemanticByCategory(st.Category, st.SemanticSQL)
		historyState[key] = st
	}
	return nil
}

func applyIndexCreatesToHistoryState(historyState map[string]historyObject, version, object string, stmts []string) error {
	if len(stmts) == 0 {
		return nil
	}
	key, ok := findHistoryIndexKey(historyState, object)
	st := historyObject{
		Category:    "indexes",
		Object:      object,
		LastVersion: version,
		LastVK:      1,
	}
	if ok {
		st = historyState[key]
	} else {
		key = objectKey("indexes", object)
	}
	stateMap := map[string]indexStatementDef{}
	if strings.TrimSpace(st.SemanticSQL) != "" {
		parsed, err := parseIndexStateMap(object, st.SemanticSQL)
		if err != nil {
			return err
		}
		stateMap = parsed
	}
	for _, stmt := range stmts {
		name, err := parseCreateIndexName(stmt)
		if err != nil {
			return err
		}
		cleanName := stripIdentifierBrackets(name)
		stateMap[cleanName] = indexStatementDef{Name: cleanName, Statement: normalizeSQLText(stmt, true, true)}
	}
	st.SemanticSQL = serializeIndexState(stateMap)
	st.NormalizedHash = hashSemanticByCategory(st.Category, st.SemanticSQL)
	historyState[key] = st
	return nil
}

func findHistoryTableKey(historyState map[string]historyObject, object string) (string, bool) {
	exactKey := objectKey("tables", object)
	if _, ok := historyState[exactKey]; ok {
		return exactKey, true
	}
	for key, st := range historyState {
		if st.Category != "tables" {
			continue
		}
		if sameQualifiedObject(st.Object, object) {
			return key, true
		}
	}
	return "", false
}

func findHistoryIndexKey(historyState map[string]historyObject, object string) (string, bool) {
	exactKey := objectKey("indexes", object)
	if _, ok := historyState[exactKey]; ok {
		return exactKey, true
	}
	for key, st := range historyState {
		if st.Category != "indexes" {
			continue
		}
		if sameQualifiedObject(st.Object, object) {
			return key, true
		}
	}
	return "", false
}

func extractDroppedConstraintsFromMigrationSQL(sql string) []dropConstraintTarget {
	normalized := normalizeSQLText(sql, true, true)
	// Search only executable SQL. Dynamic SQL text is decoded and processed by
	// the caller; scanning its escaped string representation here would record
	// definitions containing doubled quotes as real schema state.
	upper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(normalized))
	anchor := "ALTER TABLE "

	seen := make(map[string]dropConstraintTarget)
	cursor := 0
	for {
		rel := strings.Index(upper[cursor:], anchor)
		if rel < 0 {
			break
		}
		start := cursor + rel
		end := len(normalized)
		if semi := strings.Index(normalized[start:], ";"); semi >= 0 {
			end = start + semi + 1
		}

		stmt := strings.TrimSpace(normalized[start:end])
		stmtUpper := strings.ToUpper(stmt)
		if strings.HasPrefix(stmtUpper, "ALTER TABLE ") && strings.Contains(stmtUpper, " DROP CONSTRAINT ") {
			afterAlter := strings.TrimSpace(stmt[len("ALTER TABLE "):])
			tableToken := firstSQLToken(afterAlter)
			afterTable := strings.TrimSpace(afterAlter[len(tableToken):])
			if tableToken != "" && strings.HasPrefix(strings.ToUpper(afterTable), "DROP CONSTRAINT ") {
				nameToken := firstSQLToken(strings.TrimSpace(afterTable[len("DROP CONSTRAINT "):]))
				name := stripIdentifierBrackets(cleanIdentifierToken(nameToken))
				table := stripIdentifierBrackets(strings.TrimSpace(strings.TrimRight(tableToken, ",;")))
				if name != "" && table != "" {
					key := table + "|" + name
					seen[key] = dropConstraintTarget{
						Table: table,
						Name:  name,
						Kind:  "constraint",
					}
				}
			}
		}

		cursor = end
		if cursor >= len(upper) {
			break
		}
	}

	out := make([]dropConstraintTarget, 0, len(seen))
	for _, item := range seen {
		out = append(out, item)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Table != out[j].Table {
			return out[i].Table < out[j].Table
		}
		return out[i].Name < out[j].Name
	})
	return out
}

func extractAddedConstraintsFromMigrationSQL(sql string) []addedConstraintTarget {
	normalized := normalizeSQLText(sql, true, true)
	// See extractDroppedConstraintsFromMigrationSQL: string literals are not
	// executable at this level and must not be interpreted as ALTER TABLE.
	upper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(normalized))
	anchor := "ALTER TABLE "

	seen := make(map[string]addedConstraintTarget)
	cursor := 0
	for {
		rel := strings.Index(upper[cursor:], anchor)
		if rel < 0 {
			break
		}
		start := cursor + rel
		end := sqlparse.FindStatementEnd(normalized, start)
		stmt := strings.TrimSpace(normalized[start:end])
		stmtUpper := strings.ToUpper(stmt)
		if strings.HasPrefix(stmtUpper, "ALTER TABLE ") && strings.Contains(stmtUpper, " ADD CONSTRAINT ") {
			afterAlter := strings.TrimSpace(stmt[len("ALTER TABLE "):])
			tableToken := firstSQLToken(afterAlter)
			afterTable := strings.TrimSpace(afterAlter[len(tableToken):])
			if tableToken != "" && strings.HasPrefix(strings.ToUpper(afterTable), "ADD CONSTRAINT ") {
				definition := strings.TrimSpace(afterTable[len("ADD "):])
				name, ok := extractNamedConstraintName(definition)
				table := stripIdentifierBrackets(strings.TrimSpace(strings.TrimRight(tableToken, ",;")))
				if ok && table != "" {
					cleanName := stripIdentifierBrackets(name)
					key := table + "|" + cleanName
					seen[key] = addedConstraintTarget{
						Table:      table,
						Name:       cleanName,
						Kind:       constraintKindForTableItem(definition),
						Definition: strings.TrimRight(strings.TrimSpace(definition), ";"),
					}
				}
			}
		}
		cursor = end
		if cursor >= len(upper) {
			break
		}
	}

	out := make([]addedConstraintTarget, 0, len(seen))
	for _, item := range seen {
		out = append(out, item)
	}
	sort.Slice(out, func(i, j int) bool {
		if out[i].Table != out[j].Table {
			return out[i].Table < out[j].Table
		}
		return out[i].Name < out[j].Name
	})
	return out
}

func removeNamedConstraintFromCreateTable(currentCreateTable, constraintName string) (string, error) {
	items, err := extractCreateTableItems(currentCreateTable)
	if err != nil {
		return "", err
	}

	target := stripIdentifierBrackets(constraintName)
	if target == "" {
		return normalizeSQLText(currentCreateTable, true, true), nil
	}

	filtered := make([]string, 0, len(items))
	removed := false
	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}

		name, ok := extractNamedConstraintName(line)
		if ok && stripIdentifierBrackets(name) == target {
			if rewritten, inlineOK := removeInlineDefaultConstraintFromColumnItem(line, constraintName); inlineOK {
				filtered = append(filtered, rewritten)
				removed = true
				continue
			}
			removed = true
			continue
		}
		filtered = append(filtered, line)
	}

	if !removed {
		return normalizeSQLText(currentCreateTable, true, true), nil
	}
	if len(filtered) == 0 {
		return "", fmt.Errorf("constraint drop removed all table items for constraint %s", constraintName)
	}
	return rebuildCreateTableStatementWithItems(currentCreateTable, filtered)
}

func addOrReplaceNamedConstraintInCreateTable(currentCreateTable, constraintName, definition string) (string, error) {
	without, err := removeNamedConstraintFromCreateTable(currentCreateTable, constraintName)
	if err != nil {
		return "", err
	}
	items, err := extractCreateTableItems(without)
	if err != nil {
		return "", err
	}
	definition = strings.TrimRight(strings.TrimSpace(definition), ";")
	if definition == "" {
		return normalizeSQLText(without, true, true), nil
	}
	items = append(items, definition)
	return rebuildCreateTableStatementWithItems(without, items)
}

func removeInlineDefaultConstraintFromColumnItem(line, constraintName string) (string, bool) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return "", false
	}
	upper := strings.ToUpper(trimmed)
	if strings.HasPrefix(upper, "CONSTRAINT ") ||
		strings.HasPrefix(upper, "PRIMARY KEY") ||
		strings.HasPrefix(upper, "UNIQUE ") ||
		strings.HasPrefix(upper, "FOREIGN KEY") ||
		strings.HasPrefix(upper, "CHECK ") {
		return "", false
	}

	target := stripIdentifierBrackets(constraintName)
	if target == "" {
		return "", false
	}

	defaultPositions := findAllKeywordPositions(upper, "DEFAULT")
	for _, defaultPos := range defaultPositions {
		expr, exprEnd, ok := consumeDefaultExpression(trimmed, defaultPos+len("DEFAULT"))
		_ = expr
		if !ok {
			continue
		}

		prefix := trimmed[:defaultPos]
		prefixUpper := strings.ToUpper(prefix)
		constraintPositions := findAllKeywordPositions(prefixUpper, "CONSTRAINT")
		if len(constraintPositions) == 0 {
			continue
		}
		constraintPos := constraintPositions[len(constraintPositions)-1]
		candidate := strings.TrimSpace(prefix[constraintPos+len("CONSTRAINT"):])
		inlineName := stripIdentifierBrackets(cleanIdentifierToken(firstSQLToken(candidate)))
		if inlineName == "" || inlineName != target {
			continue
		}

		head := strings.TrimRight(trimmed[:constraintPos], " \t")
		tail := strings.TrimSpace(trimmed[exprEnd:])
		switch tail {
		case "":
			return head, true
		case ",":
			return head + ",", true
		default:
			if strings.HasPrefix(tail, ",") {
				return head + tail, true
			}
			return head + " " + tail, true
		}
	}
	return "", false
}

func rebuildCreateTableStatementWithItems(createTableStmt string, items []string) (string, error) {
	stmt := normalizeSQLText(createTableStmt, true, true)
	openIdx := strings.Index(stmt, "(")
	if openIdx < 0 {
		return "", fmt.Errorf("invalid CREATE TABLE statement: no opening parenthesis")
	}

	closeIdx, ok := sqlparse.FindMatchingParen(stmt, openIdx)
	if !ok {
		return "", fmt.Errorf("invalid CREATE TABLE statement: no closing parenthesis")
	}
	if len(items) == 0 {
		return "", fmt.Errorf("cannot rebuild CREATE TABLE with empty items")
	}

	rebuiltInside := "  " + strings.Join(items, ",\n  ") + "\n"
	rebuilt := stmt[:openIdx+1] + "\n" + rebuiltInside + stmt[closeIdx:]
	return normalizeSQLText(rebuilt, true, true), nil
}

func buildReAddColumnsDownSQL(previousCreateTable, object string, removedColumns []string) (string, error) {
	prevCols, err := parseTableColumns(previousCreateTable)
	if err != nil {
		return "", err
	}
	if len(removedColumns) == 0 {
		return "", nil
	}

	lines := make([]string, 0, len(removedColumns))
	for _, col := range removedColumns {
		def, ok := prevCols[col]
		if !ok {
			return "", fmt.Errorf("column %s not found in previous table state for %s", col, object)
		}
		lines = append(lines, buildAddColumnIfMissingStatement(object, def))
	}
	return normalizeSQLText(strings.Join(lines, "\n"), true, true), nil
}

func buildTableAlterColumnsSQL(previousCreateTable, nextCreateTable, object string, changes tableColumnChanges) (string, string, bool, error) {
	prevCols, err := parseTableColumns(previousCreateTable)
	if err != nil {
		return "", "", false, err
	}
	nextCols, err := parseTableColumns(nextCreateTable)
	if err != nil {
		return "", "", false, err
	}
	for _, changedName := range changes.Changed {
		prevChanged, prevOK := prevCols[changedName]
		nextChanged, nextOK := nextCols[changedName]
		if !prevOK || !nextOK || (!prevChanged.IsComputed && !nextChanged.IsComputed) {
			continue
		}
		for _, state := range []map[string]tableColumnDef{prevCols, nextCols} {
			for otherName, other := range state {
				if otherName == changedName || !other.IsComputed {
					continue
				}
				if containsColumnReference(other.ComputedDefinition, changedName) {
					return "", "", false, fmt.Errorf("computed dependency requires ordered manual migration: %s.%s depends on %s", object, other.Name, changedName)
				}
			}
		}
	}
	prevDefaultsByColumn, err := extractTableDefaultConstraintMap(previousCreateTable)
	if err != nil {
		return "", "", false, err
	}
	nextDefaultsByColumn, err := extractTableDefaultConstraintMap(nextCreateTable)
	if err != nil {
		return "", "", false, err
	}

	type step struct {
		Up   string
		Down string
	}
	steps := make([]step, 0, len(changes.Added)+len(changes.Changed))

	changed := append([]string(nil), changes.Changed...)
	sort.Strings(changed)
	changedSet := make(map[string]struct{}, len(changed))
	for _, name := range changed {
		changedSet[name] = struct{}{}
		key := name
		prevDef, okPrev := prevCols[key]
		nextDef, okNext := nextCols[key]
		if !okPrev || !okNext {
			return "", "", false, fmt.Errorf("alter column state missing for %s.%s", object, name)
		}
		if prevDef.IdentityDefinition != nextDef.IdentityDefinition {
			return "", "", false, fmt.Errorf("changing IDENTITY definition requires manual table rebuild: %s.%s", object, name)
		}
		if prevDef.RowGuidCol != nextDef.RowGuidCol {
			return "", "", false, fmt.Errorf("changing ROWGUIDCOL requires manual migration: %s.%s", object, name)
		}
		if prevDef.IsComputed != nextDef.IsComputed {
			return "", "", false, fmt.Errorf("computed/non-computed transition requires manual data migration: %s.%s", object, name)
		}
		if prevDef.IsComputed && nextDef.IsComputed {
			steps = append(steps, step{
				Up:   concatSQLSegments(buildDropDerivedColumnIfExistsStatement(object, prevDef.Name), buildAddColumnIfMissingStatement(object, nextDef)),
				Down: concatSQLSegments(buildDropDerivedColumnIfExistsStatement(object, nextDef.Name), buildAddColumnIfMissingStatement(object, prevDef)),
			})
			continue
		}
		if !isSupportedAutomaticColumnTypeTransition(prevDef, nextDef) {
			return "", "", false, fmt.Errorf("unproven column type transition requires manual migration: %s.%s (%s -> %s)", object, name, prevDef.RawType, nextDef.RawType)
		}

		alterUp := buildAlterColumnGuardedSQL(object, prevDef, nextDef)
		alterDown := buildAlterColumnGuardedSQL(object, nextDef, prevDef)
		steps = append(steps, step{
			Up:   alterUp,
			Down: alterDown,
		})

		defaultUp, defaultDown, hasDefaultDelta, defaultErr := buildDefaultConstraintDeltaStep(object, name, prevDefaultsByColumn, nextDefaultsByColumn)
		if defaultErr != nil {
			return "", "", false, defaultErr
		}
		if hasDefaultDelta {
			steps = append(steps, step{
				Up:   defaultUp,
				Down: defaultDown,
			})
		}
	}

	defaultOnly := append([]string(nil), changes.DefaultChanged...)
	sort.Strings(defaultOnly)
	for _, name := range defaultOnly {
		if _, alreadyHandled := changedSet[name]; alreadyHandled {
			continue
		}
		defaultUp, defaultDown, hasDefaultDelta, defaultErr := buildDefaultConstraintDeltaStep(object, name, prevDefaultsByColumn, nextDefaultsByColumn)
		if defaultErr != nil {
			return "", "", false, defaultErr
		}
		if hasDefaultDelta {
			steps = append(steps, step{
				Up:   defaultUp,
				Down: defaultDown,
			})
		}
	}

	added := append([]string(nil), changes.Added...)
	sort.Strings(added)
	for _, name := range added {
		nextDef, ok := nextCols[name]
		if !ok {
			return "", "", false, fmt.Errorf("added column state missing for %s.%s", object, name)
		}
		if nextDef.IsComputed {
			steps = append(steps, step{
				Up:   buildAddColumnIfMissingStatement(object, nextDef),
				Down: buildDropDerivedColumnIfExistsStatement(object, nextDef.Name),
			})
			continue
		}
		if nextDef.IdentityDefinition != "" {
			steps = append(steps, step{
				Up:   buildAddColumnIfMissingStatement(object, nextDef),
				Down: buildDropColumnIfExistsStatement(object, nextDef.Name),
			})
			continue
		}
		if !nextDef.Nullable {
			defaultDef, hasDefault := nextDefaultsByColumn[name]
			if !hasDefault {
				return "", "", false, fmt.Errorf("adding NOT NULL column requires default constraint for auto migration: %s.%s", object, name)
			}
			dropColumnDown := buildDropColumnIfExistsStatement(object, nextDef.Name)
			defaultDown, defaultErr := buildDropDefaultConstraintIfExistsStatement(object, defaultDef)
			if defaultErr != nil {
				return "", "", false, defaultErr
			}
			steps = append(steps, step{
				Up:   buildAddNotNullColumnWithDefaultSQL(object, nextDef, defaultDef),
				Down: normalizeSQLText(defaultDown+"\n"+dropColumnDown, true, true),
			})
			continue
		}
		addStep := step{
			Up:   buildAddColumnIfMissingStatement(object, nextDef),
			Down: buildDropColumnIfExistsStatement(object, nextDef.Name),
		}
		if defaultDef, hasDefault := nextDefaultsByColumn[name]; hasDefault {
			defaultUp, defaultErr := buildAddDefaultConstraintIfMissingStatement(object, defaultDef)
			if defaultErr != nil {
				return "", "", false, defaultErr
			}
			defaultDown, defaultErr := buildDropDefaultConstraintIfExistsStatement(object, defaultDef)
			if defaultErr != nil {
				return "", "", false, defaultErr
			}
			addStep.Up = normalizeSQLText(addStep.Up+"\n"+defaultUp, true, true)
			addStep.Down = normalizeSQLText(defaultDown+"\n"+addStep.Down, true, true)
		}
		steps = append(steps, addStep)
	}

	if len(steps) == 0 {
		return "", "", false, fmt.Errorf("no auto alter step generated for %s", object)
	}

	upLines := make([]string, 0, len(steps))
	downLines := make([]string, 0, len(steps))
	hasDestructiveDown := false
	for _, s := range steps {
		upLines = append(upLines, s.Up)
	}
	for i := len(steps) - 1; i >= 0; i-- {
		downLines = append(downLines, steps[i].Down)
		if strings.Contains(strings.ToUpper(steps[i].Down), "DROP COLUMN") {
			hasDestructiveDown = true
		}
	}

	upSQL := normalizeSQLText(strings.Join(upLines, "\n"), true, true)
	downSQL := normalizeSQLText(strings.Join(downLines, "\n"), true, true)
	return upSQL, downSQL, hasDestructiveDown, nil
}

func columnDefinitionSQL(def tableColumnDef) string {
	if def.IsComputed {
		return fmt.Sprintf("%s %s", quoteIdentifier(def.Name), def.ComputedDefinition)
	}
	nullability := "NULL"
	if !def.Nullable {
		nullability = "NOT NULL"
	}
	modifiers := ""
	if def.IdentityDefinition != "" {
		modifiers += " " + def.IdentityDefinition
	}
	if def.RowGuidCol {
		modifiers += " ROWGUIDCOL"
	}
	if def.Sparse {
		modifiers += " SPARSE"
	}
	return fmt.Sprintf("%s %s%s %s", quoteIdentifier(def.Name), def.RawType, modifiers, nullability)
}

func buildAddNotNullColumnWithDefaultSQL(object string, def tableColumnDef, defaultDef tableDefaultConstraintDef) string {
	var b strings.Builder
	b.WriteString(buildRequireTableExistsSQL(object, "ADD COLUMN"))
	b.WriteString("\n\n")
	b.WriteString(fmt.Sprintf("IF NOT (%s)\nBEGIN\n", buildExactColumnExistsCondition(object, def.Name)))
	b.WriteString(fmt.Sprintf("  ALTER TABLE %s ADD %s", quoteQualifiedIdentifier(object), columnDefinitionSQL(def)))
	if strings.TrimSpace(defaultDef.ConstraintName) != "" {
		b.WriteString(" CONSTRAINT ")
		b.WriteString(quoteIdentifier(defaultDef.ConstraintName))
	}
	b.WriteString(" DEFAULT ")
	b.WriteString(strings.TrimSpace(defaultDef.Expression))
	b.WriteString(" WITH VALUES;\n")
	b.WriteString("END\nELSE IF NOT (")
	b.WriteString(buildColumnEquivalentCondition(object, def))
	b.WriteString(")\nBEGIN\n  ;THROW 51000, N'Schema drift: existing column does not match the requested definition', 1;\nEND")
	return normalizeSQLText(b.String(), true, true)
}

func buildAlterColumnGuardedSQL(object string, prevDef, nextDef tableColumnDef) string {
	blocks := make([]string, 0, 6)
	columnRef := quoteIdentifier(nextDef.Name)
	target := object + "." + nextDef.Name
	varSuffix := safeSQLVarSuffix(nextDef.Name)
	skipVar := "@skip_alter_" + varSuffix

	blocks = append(blocks, "-- check alter-column safety per target column")
	blocks = append(blocks, fmt.Sprintf("DECLARE %s BIT = 0;", skipVar))
	blocks = append(blocks, buildRequireTableExistsSQL(object, "ALTER COLUMN"))

	// Even a widening/nullability-preserving ALTER must verify that its target
	// column exists.  The data-safety checks below are intentionally conditional,
	// so without this baseline guard those otherwise-safe changes were emitted as
	// a bare ALTER TABLE statement and rejected by strict verification.
	blocks = append(blocks, normalizeSQLText(fmt.Sprintf(`-- check target column exists before ALTER COLUMN
IF NOT (%s)
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 2,
      @action = N'alter_column',
      @target = N'%s',
      @reason = N'COLUMN_NOT_EXISTS',
      @message = N'Precheck: ALTER COLUMN target does not exist',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  SET %s = 1;
END`, buildExactColumnExistsCondition(object, nextDef.Name), target, skipVar), true, true))

	if prevDef.Nullable && !nextDef.Nullable {
		nullVar := "@null_cnt_" + varSuffix
		detailVar := "@detail_null_cnt_" + varSuffix
		blocks = append(blocks, "-- check NULL values before changing to NOT NULL")
		blocks = append(blocks, normalizeSQLText(fmt.Sprintf(`DECLARE %s BIGINT;
DECLARE %s NVARCHAR(128);
SELECT %s = COUNT_BIG(1)
FROM %s
WHERE %s IS NULL;

IF (%s > 0)
BEGIN
  SET %s = N'null_cnt=' + CONVERT(NVARCHAR(20), %s);
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 2,
      @action = N'alter_column',
      @target = N'%s',
      @reason = N'REQUIRES_MANUAL',
      @message = N'Precheck: ALTER COLUMN NOT NULL blocked',
      @detail = %s,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  SET %s = 1;
END`, nullVar, detailVar, nullVar, quoteQualifiedIdentifier(object), columnRef, nullVar, detailVar, nullVar, target, detailVar, skipVar), true, true))
	}

	if ((isStringType(prevDef.BaseType) && isStringType(nextDef.BaseType)) || (isBinaryType(prevDef.BaseType) && isBinaryType(nextDef.BaseType))) && nextDef.Length > 0 {
		tooLongVar := "@too_long_" + varSuffix
		detailVar := "@detail_too_long_" + varSuffix
		maxExpr := fmt.Sprintf("%d", nextDef.Length)
		if nextDef.IsUnicode {
			maxExpr = fmt.Sprintf("(%d * 2)", nextDef.Length)
		}
		reason := "DATA_TOO_LONG_FOR_TARGET"
		message := "Precheck: target length blocked"
		if prevDef.Length > 0 && nextDef.Length < prevDef.Length {
			reason = "NARROWING_WOULD_TRUNCATE"
			message = "Precheck: SHRINK length blocked"
		}
		blocks = append(blocks, "-- check data length fits target column type/length")
		blocks = append(blocks, normalizeSQLText(fmt.Sprintf(`DECLARE %s BIGINT;
DECLARE %s NVARCHAR(128);
SELECT %s = COUNT_BIG(1)
FROM %s
WHERE %s IS NOT NULL
  AND DATALENGTH(%s) > %s;

IF (%s > 0)
BEGIN
  SET %s = N'too_long=' + CONVERT(NVARCHAR(20), %s);
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 2,
      @action = N'alter_column',
      @target = N'%s',
      @reason = N'%s',
      @message = N'%s',
      @detail = %s,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  SET %s = 1;
END`, tooLongVar, detailVar, tooLongVar, quoteQualifiedIdentifier(object), columnRef, columnRef, maxExpr, tooLongVar, detailVar, tooLongVar, target, reason, message, detailVar, skipVar), true, true))
	}

	guardedNumericOrText := (isNumericType(prevDef.BaseType) || isStringType(prevDef.BaseType)) && (isNumericType(nextDef.BaseType) || isStringType(nextDef.BaseType)) && (isNumericType(prevDef.BaseType) || isNumericType(nextDef.BaseType))
	if (guardedNumericOrText || isDateTimeType(prevDef.BaseType) && isDateTimeType(nextDef.BaseType)) && !columnTypeEquivalent(prevDef, nextDef) {
		badVar := "@bad_" + varSuffix
		detailVar := "@detail_bad_" + varSuffix
		prevConvertType := conversionTypeSQL(prevDef.RawType)
		nextConvertType := conversionTypeSQL(nextDef.RawType)
		blocks = append(blocks, "-- check data convertibility before changing column type")
		blocks = append(blocks, normalizeSQLText(fmt.Sprintf(`DECLARE %s BIGINT;
DECLARE %s NVARCHAR(128);
SELECT %s = COUNT_BIG(1)
FROM %s
WHERE %s IS NOT NULL
  AND (
    TRY_CONVERT(%s, %s) IS NULL
    OR TRY_CONVERT(%s, TRY_CONVERT(%s, %s)) <> %s
  );

IF (%s > 0)
BEGIN
  SET %s = N'bad_cnt=' + CONVERT(NVARCHAR(20), %s);
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 2,
      @action = N'alter_column',
      @target = N'%s',
      @reason = N'REQUIRES_MANUAL',
      @message = N'Precheck: TYPE change blocked (non-convertible)',
      @detail = %s,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  SET %s = 1;
END`, badVar, detailVar, badVar, quoteQualifiedIdentifier(object), columnRef, nextConvertType, columnRef, prevConvertType, nextConvertType, columnRef, columnRef, badVar, detailVar, badVar, target, detailVar, skipVar), true, true))
	}

	if prevDef.BaseType == "NVARCHAR" && prevDef.Length == -1 && nextDef.BaseType == "XML" && strings.EqualFold(strings.TrimSpace(nextDef.RawType), "XML") {
		badXMLVar := "@bad_xml_" + varSuffix
		detailVar := "@detail_bad_xml_" + varSuffix
		blocks = append(blocks, "-- check all text values are well-formed XML before restoring XML type")
		blocks = append(blocks, normalizeSQLText(fmt.Sprintf(`DECLARE %s BIGINT;
DECLARE %s NVARCHAR(128);
SELECT %s = COUNT_BIG(1)
FROM %s
WHERE %s IS NOT NULL
  AND TRY_CONVERT(XML, %s) IS NULL;

IF (%s > 0)
BEGIN
  SET %s = N'invalid_xml_cnt=' + CONVERT(NVARCHAR(20), %s);
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 2,
      @action = N'alter_column',
      @target = N'%s',
      @reason = N'REQUIRES_MANUAL',
      @message = N'Precheck: NVARCHAR(MAX) to XML blocked by invalid XML text',
      @detail = %s,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  SET %s = 1;
END`, badXMLVar, detailVar, badXMLVar, quoteQualifiedIdentifier(object), columnRef, columnRef, badXMLVar, detailVar, badXMLVar, target, detailVar, skipVar), true, true))
	}

	// A precheck that finds incompatible data must abort the outer transaction.
	// Silently skipping ALTER COLUMN would mark the migration applied while the
	// schema remains old, which breaks db_object/history consistency.
	blocks = append(blocks, fmt.Sprintf(`IF (%s <> 0)
BEGIN
  ;THROW 51000, N'Precheck blocked: ALTER COLUMN schema was not changed', 1;
END`, skipVar))

	blocks = append(blocks, normalizeSQLText(fmt.Sprintf(`IF (%s = 0)
BEGIN
  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'run',
      @severity = 1,
      @action = N'alter_column',
      @target = N'%s',
      @reason = N'OBJECT_EXISTS',
      @message = N'alter column start',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END

  ALTER TABLE %s ALTER COLUMN %s;

  IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'ok',
      @severity = 1,
      @action = N'alter_column',
      @target = N'%s',
      @reason = N'ALTERED',
      @message = N'alter column ok',
      @detail = @direction,
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
END`, skipVar, target, quoteQualifiedIdentifier(object), columnDefinitionSQL(nextDef), target), true, true))

	return normalizeSQLText(strings.Join(blocks, "\n\n"), true, true)
}

// isSupportedAutomaticColumnTypeTransition is deliberately conservative. An
// ALTER is emitted only for families whose losslessness can be proved by the
// generated runtime guards. Everything else is rejected before a migration is
// written, so SQL Server never becomes the first place an unsafe transition is
// discovered.
func isSupportedAutomaticColumnTypeTransition(before, after tableColumnDef) bool {
	beforeBase := strings.ToUpper(strings.TrimSpace(before.BaseType))
	afterBase := strings.ToUpper(strings.TrimSpace(after.BaseType))
	if beforeBase == afterBase {
		return true
	}
	guardedNumericOrText := (isLosslessGuardableNumericType(beforeBase) || isStringType(beforeBase)) &&
		(isLosslessGuardableNumericType(afterBase) || isStringType(afterBase)) &&
		(isLosslessGuardableNumericType(beforeBase) || isLosslessGuardableNumericType(afterBase))
	if guardedNumericOrText || isDateTimeType(beforeBase) && isDateTimeType(afterBase) {
		return true
	}
	return (beforeBase == "XML" && afterBase == "NVARCHAR" && after.Length == -1) ||
		(beforeBase == "NVARCHAR" && before.Length == -1 && afterBase == "XML")
}

func buildDefaultConstraintDeltaStep(
	object, column string,
	prevByColumn, nextByColumn map[string]tableDefaultConstraintDef,
) (string, string, bool, error) {
	key := strings.TrimSpace(column)
	prevDef, hasPrev := prevByColumn[key]
	nextDef, hasNext := nextByColumn[key]
	if !hasPrev && !hasNext {
		return "", "", false, nil
	}

	if hasPrev && hasNext && defaultConstraintEquivalent(prevDef, nextDef) {
		return "", "", false, nil
	}

	upParts := make([]string, 0, 2)
	downParts := make([]string, 0, 2)

	if hasPrev {
		upDrop, err := buildDropDefaultConstraintIfExistsStatement(object, prevDef)
		if err != nil {
			return "", "", false, err
		}
		upParts = append(upParts, upDrop)
	}
	if hasNext {
		upAdd, err := buildAddDefaultConstraintIfMissingStatement(object, nextDef)
		if err != nil {
			return "", "", false, err
		}
		upParts = append(upParts, upAdd)
	}

	if hasNext {
		downDrop, err := buildDropDefaultConstraintIfExistsStatement(object, nextDef)
		if err != nil {
			return "", "", false, err
		}
		downParts = append(downParts, downDrop)
	}
	if hasPrev {
		downAdd, err := buildAddDefaultConstraintIfMissingStatement(object, prevDef)
		if err != nil {
			return "", "", false, err
		}
		downParts = append(downParts, downAdd)
	}

	return normalizeSQLText(strings.Join(upParts, "\n"), true, true),
		normalizeSQLText(strings.Join(downParts, "\n"), true, true),
		true, nil
}

func defaultConstraintEquivalent(a, b tableDefaultConstraintDef) bool {
	nameA := strings.TrimSpace(a.ConstraintName)
	nameB := strings.TrimSpace(b.ConstraintName)
	exprA := normalizeDefaultExpressionForCompare(a.Expression)
	exprB := normalizeDefaultExpressionForCompare(b.Expression)
	return nameA == nameB && exprA == exprB
}

func buildCheckConstraintOnlyLifecycleSQL(previousCreateTable, nextCreateTable, object string) (string, string, bool, error) {
	return buildNamedConstraintLifecycleSQL(previousCreateTable, nextCreateTable, object, false)
}

func buildNamedConstraintLifecycleSQL(previousCreateTable, nextCreateTable, object string, skipDefaults bool) (string, string, bool, error) {
	prevMap, err := extractNamedConstraintDefinitionMap(previousCreateTable)
	if err != nil {
		return "", "", false, err
	}
	// An inline named DEFAULT is represented by the complete column item in the
	// generic named-constraint map. Do not rely on generic text classification
	// to identify it: data types or other column clauses can make that
	// classification ambiguous. The dedicated DEFAULT parser knows both inline
	// and table-level forms and is therefore authoritative when callers ask to
	// plan non-default constraints only.
	defaultNames := make(map[string]struct{})
	if skipDefaults {
		for _, createTable := range []string{previousCreateTable, nextCreateTable} {
			defaults, defaultErr := extractTableDefaultConstraintMap(createTable)
			if defaultErr != nil {
				return "", "", false, defaultErr
			}
			for _, def := range defaults {
				name := stripIdentifierBrackets(strings.TrimSpace(def.ConstraintName))
				if name != "" {
					defaultNames[name] = struct{}{}
				}
			}
		}
	}
	nextMap, err := extractNamedConstraintDefinitionMap(nextCreateTable)
	if err != nil {
		return "", "", false, err
	}

	keys := make(map[string]struct{}, len(prevMap)+len(nextMap))
	for key := range prevMap {
		keys[key] = struct{}{}
	}
	for key := range nextMap {
		keys[key] = struct{}{}
	}

	orderedKeys := make([]string, 0, len(keys))
	for key := range keys {
		orderedKeys = append(orderedKeys, key)
	}
	sort.Strings(orderedKeys)
	upParts := make([]string, 0)
	downParts := make([]string, 0)
	for _, key := range orderedKeys {
		if _, isDefault := defaultNames[stripIdentifierBrackets(key)]; isDefault {
			continue
		}
		prevDef, hasPrev := prevMap[key]
		nextDef, hasNext := nextMap[key]
		if hasPrev && hasNext && constraintDefinitionEquivalent(prevDef, nextDef) {
			continue
		}

		prevKind := ""
		nextKind := ""
		if hasPrev {
			prevKind = constraintKindForTableItem(prevDef)
		}
		if hasNext {
			nextKind = constraintKindForTableItem(nextDef)
		}
		if hasPrev && hasNext && prevKind != nextKind {
			return "", "", false, fmt.Errorf("constraint-only auto-plan cannot change constraint kind: %s", key)
		}

		kind := prevKind
		if kind == "" {
			kind = nextKind
		}
		if kind != "check" && kind != "default" && kind != "unique" && kind != "primary_key" && kind != "foreign_key" {
			return "", "", false, fmt.Errorf("constraint-only auto-plan supports named CHECK/DEFAULT/UNIQUE/PRIMARY KEY/FOREIGN KEY constraints only: %s", key)
		}
		if skipDefaults && kind == "default" {
			continue
		}

		if kind == "default" {
			var prevDefault tableDefaultConstraintDef
			var nextDefault tableDefaultConstraintDef
			if hasPrev {
				var ok bool
				prevDefault, ok = parseDefaultConstraintFromCreateTableItem(prevDef)
				if !ok {
					return "", "", false, fmt.Errorf("DEFAULT constraint parse failed: %s", key)
				}
			}
			if hasNext {
				var ok bool
				nextDefault, ok = parseDefaultConstraintFromCreateTableItem(nextDef)
				if !ok {
					return "", "", false, fmt.Errorf("DEFAULT constraint parse failed: %s", key)
				}
			}
			if hasPrev && hasNext && defaultConstraintEquivalent(prevDefault, nextDefault) {
				continue
			}
			if hasPrev {
				upDrop, err := buildDropDefaultConstraintIfExistsStatement(object, prevDefault)
				if err != nil {
					return "", "", false, err
				}
				upParts = append(upParts, upDrop)
			}
			if hasNext {
				upAdd, err := buildAddDefaultConstraintIfMissingStatement(object, nextDefault)
				if err != nil {
					return "", "", false, err
				}
				upParts = append(upParts, upAdd)
			}
			if hasNext {
				downDrop, err := buildDropDefaultConstraintIfExistsStatement(object, nextDefault)
				if err != nil {
					return "", "", false, err
				}
				downParts = append(downParts, downDrop)
			}
			if hasPrev {
				downAdd, err := buildAddDefaultConstraintIfMissingStatement(object, prevDefault)
				if err != nil {
					return "", "", false, err
				}
				downParts = append(downParts, downAdd)
			}
			continue
		}

		if kind == "unique" || kind == "primary_key" {
			name := key
			if hasPrev {
				if parsed, ok := extractNamedConstraintName(prevDef); ok {
					name = stripIdentifierBrackets(parsed)
				}
				upParts = append(upParts, buildDropConstraintIfExistsStatement(object, name, kind))
			}
			if hasNext {
				if parsed, ok := extractNamedConstraintName(nextDef); ok {
					name = stripIdentifierBrackets(parsed)
				}
				add, addErr := buildAddKeyConstraintWithPrecheckStatement(object, name, kind, nextDef)
				if addErr != nil {
					return "", "", false, addErr
				}
				upParts = append(upParts, add)
			}
			if hasNext {
				if parsed, ok := extractNamedConstraintName(nextDef); ok {
					name = stripIdentifierBrackets(parsed)
				}
				downParts = append(downParts, buildDropConstraintIfExistsStatement(object, name, kind))
			}
			if hasPrev {
				if parsed, ok := extractNamedConstraintName(prevDef); ok {
					name = stripIdentifierBrackets(parsed)
				}
				add, addErr := buildAddKeyConstraintWithPrecheckStatement(object, name, kind, prevDef)
				if addErr != nil {
					return "", "", false, addErr
				}
				downParts = append(downParts, add)
			}
			continue
		}

		if kind == "foreign_key" {
			name := key
			if hasPrev {
				if parsed, ok := extractNamedConstraintName(prevDef); ok {
					name = stripIdentifierBrackets(parsed)
				}
				upParts = append(upParts, buildDropConstraintIfExistsStatement(object, name, kind))
			}
			if hasNext {
				if parsed, ok := extractNamedConstraintName(nextDef); ok {
					name = stripIdentifierBrackets(parsed)
				}
				add, addErr := buildAddForeignKeyWithPrecheckStatement(object, name, nextDef)
				if addErr != nil {
					return "", "", false, addErr
				}
				upParts = append(upParts, add)
			}
			if hasNext {
				if parsed, ok := extractNamedConstraintName(nextDef); ok {
					name = stripIdentifierBrackets(parsed)
				}
				downParts = append(downParts, buildDropConstraintIfExistsStatement(object, name, kind))
			}
			if hasPrev {
				if parsed, ok := extractNamedConstraintName(prevDef); ok {
					name = stripIdentifierBrackets(parsed)
				}
				add, addErr := buildAddForeignKeyWithPrecheckStatement(object, name, prevDef)
				if addErr != nil {
					return "", "", false, addErr
				}
				downParts = append(downParts, add)
			}
			continue
		}

		if (hasPrev && !strings.HasPrefix(strings.ToUpper(strings.TrimSpace(prevDef)), "CONSTRAINT ")) ||
			(hasNext && !strings.HasPrefix(strings.ToUpper(strings.TrimSpace(nextDef)), "CONSTRAINT ")) {
			return "", "", false, fmt.Errorf("inline CHECK constraint change requires manual migration: %s", key)
		}

		name := key
		if hasPrev {
			if parsed, ok := extractNamedConstraintName(prevDef); ok {
				name = stripIdentifierBrackets(parsed)
			}
			upParts = append(upParts, buildDropConstraintIfExistsStatement(object, name, "check"))
		}
		if hasNext {
			if parsed, ok := extractNamedConstraintName(nextDef); ok {
				name = stripIdentifierBrackets(parsed)
			}
			add, addErr := buildAddCheckConstraintWithPrecheckStatement(object, name, nextDef)
			if addErr != nil {
				return "", "", false, addErr
			}
			upParts = append(upParts, add)
		}

		if hasNext {
			if parsed, ok := extractNamedConstraintName(nextDef); ok {
				name = stripIdentifierBrackets(parsed)
			}
			downParts = append(downParts, buildDropConstraintIfExistsStatement(object, name, "check"))
		}
		if hasPrev {
			if parsed, ok := extractNamedConstraintName(prevDef); ok {
				name = stripIdentifierBrackets(parsed)
			}
			add, addErr := buildAddCheckConstraintWithPrecheckStatement(object, name, prevDef)
			if addErr != nil {
				return "", "", false, addErr
			}
			downParts = append(downParts, add)
		}
	}

	if len(upParts) == 0 && len(downParts) == 0 {
		return "", "", false, nil
	}
	return concatSQLSegments(upParts...), concatSQLSegments(downParts...), true, nil
}

type incomingForeignKeyDefinition struct {
	Table      string
	Name       string
	Definition string
}

// buildConstraintLifecycleWithDependencies protects PK/UQ evolution. A
// referenced key cannot be dropped while incoming FKs exist, so known incoming
// FKs from migration history are temporarily removed and recreated with the
// same orphan precheck in both directions.
func buildConstraintLifecycleWithDependencies(previousCreateTable, nextCreateTable, object string, historyState map[string]historyObject) (string, string, bool, error) {
	return buildConstraintLifecycleWithDependenciesMode(previousCreateTable, nextCreateTable, object, historyState, false)
}

func buildNonDefaultConstraintLifecycleWithDependencies(previousCreateTable, nextCreateTable, object string, historyState map[string]historyObject) (string, string, bool, error) {
	return buildConstraintLifecycleWithDependenciesMode(previousCreateTable, nextCreateTable, object, historyState, true)
}

func buildConstraintLifecycleWithDependenciesMode(previousCreateTable, nextCreateTable, object string, historyState map[string]historyObject, skipDefaults bool) (string, string, bool, error) {
	up, down, ok, err := buildNamedConstraintLifecycleSQL(previousCreateTable, nextCreateTable, object, skipDefaults)
	if err != nil || !ok {
		return up, down, ok, err
	}
	if !hasChangedKeyConstraint(previousCreateTable, nextCreateTable) {
		return up, down, ok, nil
	}
	incoming, err := collectIncomingForeignKeyDefinitions(object, historyState)
	if err != nil {
		return "", "", false, err
	}
	if len(incoming) == 0 {
		return up, down, ok, nil
	}
	dropParts := make([]string, 0, len(incoming))
	addParts := make([]string, 0, len(incoming))
	for _, fk := range incoming {
		dropParts = append(dropParts, buildDropConstraintIfExistsStatement(fk.Table, fk.Name, "foreign_key"))
		add, addErr := buildAddForeignKeyWithPrecheckStatement(fk.Table, fk.Name, fk.Definition)
		if addErr != nil {
			return "", "", false, addErr
		}
		addParts = append(addParts, add)
	}
	drops := normalizeSQLText(strings.Join(dropParts, "\n"), true, true)
	adds := normalizeSQLText(strings.Join(addParts, "\n"), true, true)
	return concatSQLSegments(drops, up, adds), concatSQLSegments(drops, down, adds), true, nil
}

func hasChangedKeyConstraint(previousCreateTable, nextCreateTable string) bool {
	previous, prevErr := extractNamedConstraintDefinitionMap(previousCreateTable)
	next, nextErr := extractNamedConstraintDefinitionMap(nextCreateTable)
	if prevErr != nil || nextErr != nil {
		return false
	}
	keys := map[string]struct{}{}
	for key := range previous {
		keys[key] = struct{}{}
	}
	for key := range next {
		keys[key] = struct{}{}
	}
	for key := range keys {
		before, hasBefore := previous[key]
		after, hasAfter := next[key]
		kind := ""
		if hasBefore {
			kind = constraintKindForTableItem(before)
		}
		if kind == "" && hasAfter {
			kind = constraintKindForTableItem(after)
		}
		if (kind == "primary_key" || kind == "unique") && (!hasBefore || !hasAfter || !constraintDefinitionEquivalent(before, after)) {
			return true
		}
	}
	return false
}

func collectIncomingForeignKeyDefinitions(object string, historyState map[string]historyObject) ([]incomingForeignKeyDefinition, error) {
	keys := make([]string, 0, len(historyState))
	for key := range historyState {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	result := make([]incomingForeignKeyDefinition, 0)
	for _, key := range keys {
		state := historyState[key]
		if state.Category != "tables" || strings.TrimSpace(state.SemanticSQL) == "" {
			continue
		}
		definitions, err := extractNamedConstraintDefinitionMap(state.SemanticSQL)
		if err != nil {
			return nil, err
		}
		for _, definition := range definitions {
			if constraintKindForTableItem(definition) != "foreign_key" {
				continue
			}
			parsed, err := parseForeignKeyDefinition(definition)
			if err != nil {
				return nil, err
			}
			if !sameQualifiedObject(parsed.ReferencedTable, object) {
				continue
			}
			name, ok := extractNamedConstraintName(definition)
			if !ok {
				return nil, fmt.Errorf("incoming foreign key must be named: %s", definition)
			}
			result = append(result, incomingForeignKeyDefinition{Table: state.Object, Name: stripIdentifierBrackets(name), Definition: definition})
		}
	}
	sort.Slice(result, func(i, j int) bool {
		if result[i].Table != result[j].Table {
			return result[i].Table < result[j].Table
		}
		return result[i].Name < result[j].Name
	})
	return result, nil
}

func constraintDefinitionEquivalent(a, b string) bool {
	return normalizeForSemanticHash(a) == normalizeForSemanticHash(b)
}

func normalizeDefaultExpressionForCompare(expr string) string {
	if strings.TrimSpace(expr) == "" {
		return ""
	}
	return normalizeForSemanticHash(expr)
}

func buildDefaultConstraintDefinitionSQL(def tableDefaultConstraintDef) (string, error) {
	column := strings.TrimSpace(def.Column)
	expr := strings.TrimSpace(def.Expression)
	if column == "" || expr == "" {
		return "", fmt.Errorf("invalid default constraint definition: column/expression required")
	}

	var b strings.Builder
	if strings.TrimSpace(def.ConstraintName) != "" {
		b.WriteString("CONSTRAINT ")
		b.WriteString(quoteIdentifier(def.ConstraintName))
		b.WriteString(" ")
	}
	b.WriteString("DEFAULT ")
	b.WriteString(expr)
	b.WriteString(" FOR ")
	b.WriteString(quoteIdentifier(column))
	return b.String(), nil
}

func buildAddDefaultConstraintIfMissingStatement(object string, def tableDefaultConstraintDef) (string, error) {
	defSQL, err := buildDefaultConstraintDefinitionSQL(def)
	if err != nil {
		return "", err
	}
	column := strings.TrimSpace(def.Column)
	target := object + "." + column
	runLog := indentSQL(buildStepAuditLogSQL("run", 1, "add_default", target, "OBJECT_EXISTS", "add default start"), "  ")
	okLog := indentSQL(buildStepAuditLogSQL("ok", 1, "add_default", target, "CREATED", "add default ok"), "  ")

	constraintName := strings.TrimSpace(def.ConstraintName)
	if constraintName == "" {
		return "", fmt.Errorf("default constraint must be named for deterministic migration: %s.%s", object, column)
	}
	safeConstraintName := escapeSQLLiteral(stripIdentifierBrackets(constraintName))
	safeObject := escapeSQLLiteral(object)
	safeColumn := escapeSQLLiteral(stripIdentifierBrackets(column))
	return normalizeSQLText(fmt.Sprintf(`%s

%s

-- check default constraint does not exist on target table/column before ADD DEFAULT
IF NOT EXISTS (
  SELECT 1
  FROM sys.default_constraints dc
  JOIN sys.columns c
    ON c.object_id = dc.parent_object_id
   AND c.column_id = dc.parent_column_id
  WHERE dc.parent_object_id = OBJECT_ID(N'%s')
    AND dc.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
    AND c.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
)
BEGIN
%s
  ALTER TABLE %s ADD %s;
%s
END
ELSE
BEGIN
  ;THROW 51000, N'Schema drift: existing DEFAULT constraint definition was not proven equivalent', 1;
END`, buildRequireTableExistsSQL(object, "ADD DEFAULT CONSTRAINT"), buildRequireColumnsExistSQL(object, []string{column}, "ADD DEFAULT CONSTRAINT"), safeObject, safeConstraintName, safeColumn, runLog, quoteQualifiedIdentifier(object), defSQL, okLog), true, true), nil
}

func buildDropDefaultConstraintIfExistsStatement(object string, def tableDefaultConstraintDef) (string, error) {
	constraintName := strings.TrimSpace(def.ConstraintName)
	if constraintName == "" {
		return "", fmt.Errorf("default constraint must be named for deterministic migration: %s.%s", object, strings.TrimSpace(def.Column))
	}
	column := strings.TrimSpace(def.Column)
	target := object + "." + column
	runLog := indentSQL(buildStepAuditLogSQL("run", 1, "drop_default", target, "OBJECT_EXISTS", "drop default start"), "  ")
	okLog := indentSQL(buildStepAuditLogSQL("ok", 1, "drop_default", target, "DROPPED", "drop default ok"), "  ")
	skipLog := indentSQL(buildStepAuditLogSQL("skip", 0, "drop_default", target, "DF_NOT_EXISTS", "Precheck: default constraint already absent"), "  ")
	safeConstraintName := escapeSQLLiteral(stripIdentifierBrackets(constraintName))
	safeObject := escapeSQLLiteral(object)
	safeColumn := escapeSQLLiteral(stripIdentifierBrackets(column))
	return normalizeSQLText(fmt.Sprintf(`-- check default constraint exists on target table/column before DROP DEFAULT
IF EXISTS (
  SELECT 1
  FROM sys.default_constraints dc
  JOIN sys.columns c
    ON c.object_id = dc.parent_object_id
   AND c.column_id = dc.parent_column_id
  WHERE dc.parent_object_id = OBJECT_ID(N'%s')
    AND dc.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
    AND c.name COLLATE Latin1_General_100_BIN2 = N'%s' COLLATE Latin1_General_100_BIN2
)
BEGIN
%s
  ALTER TABLE %s DROP CONSTRAINT %s;
%s
END
ELSE
BEGIN
%s
END`, safeObject, safeConstraintName, safeColumn, runLog, quoteQualifiedIdentifier(object), quoteIdentifier(constraintName), okLog, skipLog), true, true), nil
}

func safeSQLVarSuffix(name string) string {
	sum := sha256.Sum256([]byte(strings.TrimSpace(name)))
	return "col_" + hex.EncodeToString(sum[:])
}

func conversionTypeSQL(rawType string) string {
	normalized := strings.TrimSpace(rawType)
	upper := strings.ToUpper(normalized)
	if pos := strings.Index(upper, " COLLATE "); pos >= 0 {
		return strings.TrimSpace(normalized[:pos])
	}
	return normalized
}

func buildDropTableSQL(object string) string {
	safeObject := escapeSQLLiteral(object)
	textReferenceCondition := buildModuleTextReferenceCondition(object)
	quotedObject := quoteQualifiedIdentifier(object)
	return normalizeSQLText(fmt.Sprintf(`-- A desired-state removal is allowed only for an empty, unreferenced table.
IF OBJECT_ID(N'%s', N'U') IS NOT NULL
AND EXISTS (SELECT 1 FROM %s WITH (TABLOCKX, HOLDLOCK))
BEGIN
  ;THROW 51000, N'Precheck blocked: DROP TABLE target contains data', 1;
END

IF OBJECT_ID(N'%s', N'U') IS NOT NULL
AND EXISTS (
  SELECT 1
  FROM sys.foreign_keys
  WHERE referenced_object_id = OBJECT_ID(N'%s')
    AND parent_object_id <> referenced_object_id
)
BEGIN
  ;THROW 51000, N'Precheck blocked: DROP TABLE target is referenced by a foreign key', 1;
END

IF OBJECT_ID(N'%s', N'U') IS NOT NULL
AND EXISTS (
  SELECT 1
  FROM sys.sql_expression_dependencies
  WHERE referenced_id = OBJECT_ID(N'%s')
    AND referencing_id <> referenced_id
    AND NOT EXISTS (
      SELECT 1
      FROM sys.objects AS own
      WHERE own.object_id = referencing_id
        AND own.parent_object_id = OBJECT_ID(N'%s')
    )
)
BEGIN
  ;THROW 51000, N'Precheck blocked: DROP TABLE target is referenced by another SQL object', 1;
END

IF OBJECT_ID(N'%s', N'U') IS NOT NULL
AND EXISTS (
  SELECT 1
  FROM sys.sql_modules
  WHERE object_id <> OBJECT_ID(N'%s')
    AND (
      %s
    )
)
BEGIN
  ;THROW 51000, N'Precheck blocked: DROP TABLE target name is used in module text or dynamic SQL', 1;
END

IF OBJECT_ID(N'%s', N'U') IS NOT NULL
AND EXISTS (
  SELECT 1
  FROM sys.triggers
  WHERE parent_id = OBJECT_ID(N'%s')
    AND is_ms_shipped = 0
)
BEGIN
  ;THROW 51000, N'Precheck blocked: DROP TABLE target still has a trigger', 1;
END

DROP TABLE IF EXISTS %s;`,
		safeObject, quotedObject,
		safeObject, safeObject, safeObject,
		safeObject, safeObject,
		safeObject, safeObject, textReferenceCondition,
		safeObject, safeObject,
		quotedObject), true, true)
}

func forceBracketQualifiedIdentifier(object string) string {
	parts := strings.Split(strings.TrimSpace(object), ".")
	if len(parts) != 2 {
		return quoteQualifiedIdentifier(object)
	}
	return quoteIdentifier(stripIdentifierBrackets(parts[0])) + "." + quoteIdentifier(stripIdentifierBrackets(parts[1]))
}

func buildModuleTextReferenceCondition(object string) string {
	parts := strings.Split(strings.TrimSpace(object), ".")
	variants := []string{strings.TrimSpace(object), forceBracketQualifiedIdentifier(object)}
	if len(parts) == 2 {
		schema := stripIdentifierBrackets(parts[0])
		name := stripIdentifierBrackets(parts[1])
		variants = append(variants,
			schema+"."+quoteIdentifier(name),
			quoteIdentifier(schema)+"."+name,
		)
	}
	seen := make(map[string]struct{}, len(variants))
	conditions := make([]string, 0, len(variants))
	for _, variant := range variants {
		variant = strings.TrimSpace(variant)
		if variant == "" {
			continue
		}
		if _, exists := seen[variant]; exists {
			continue
		}
		seen[variant] = struct{}{}
		conditions = append(conditions, fmt.Sprintf(
			"CHARINDEX(N'%s', definition COLLATE Latin1_General_100_BIN2) > 0",
			escapeSQLLiteral(variant),
		))
	}
	return strings.Join(conditions, "\n      OR ")
}

func isEngineInfrastructureObject(category, object string) bool {
	category = strings.ToLower(strings.TrimSpace(category))
	object = strings.TrimSpace(object)
	switch category {
	case "procedures":
		return object == "dbo.usp_mig_log_event" || object == "dbo.usp_mig_try_get_lock"
	case "tables", "indexes":
		return object == "dbo.mig_event" ||
			object == "dbo.mig_operation_journal" ||
			object == "dbo.schema_migrations"
	}
	return false
}

func buildManualDropTableRequestUpSQL(object string) string {
	body := fmt.Sprintf(`DECLARE @rza_manual_drop_table_request BIT = 1;

IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
BEGIN
  IF OBJECT_ID(N'%s', N'U') IS NULL
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'skip',
      @severity = 0,
      @action = N'followup_required',
      @target = N'%s',
      @reason = N'TABLE_NOT_EXISTS',
      @message = N'[MANUAL_ONLY] table already absent',
      @detail = N'No automatic DROP TABLE executed.',
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
  ELSE
  BEGIN
    EXEC dbo.usp_mig_log_event
      @run_key = @run_key,
      @version = @version,
      @category = @category,
      @file_name = @file,
      @decision = N'warn',
      @severity = 2,
      @action = N'followup_required',
      @target = N'%s',
      @reason = N'REQUIRES_MANUAL',
      @message = N'[MANUAL_ONLY] drop table requires manual execution',
      @detail = N'Engine emitted manual-only migration without DROP TABLE.',
      @error_number = NULL,
      @error_line = NULL,
      @error_procedure = NULL;
  END
END
`, object, object, object)
	return normalizeSQLText(body, true, true)
}

func buildManualDropTableRequestDownSQL(object string) string {
	body := fmt.Sprintf(`DECLARE @rza_manual_drop_table_request BIT = 1;

IF OBJECT_ID(N'dbo.usp_mig_log_event', N'P') IS NOT NULL
BEGIN
  EXEC dbo.usp_mig_log_event
    @run_key = @run_key,
    @version = @version,
    @category = @category,
    @file_name = @file,
    @decision = N'warn',
    @severity = 2,
    @action = N'followup_required',
    @target = N'%s',
    @reason = N'DATA_RESTORE_NOT_GUARANTEED',
    @message = N'[MANUAL_ONLY] down cannot guarantee data restore',
    @detail = N'Manual follow-up required for restore path.',
    @error_number = NULL,
    @error_line = NULL,
    @error_procedure = NULL;
END
`, object)
	return normalizeSQLText(body, true, true)
}

func isManualDropTableRequestMigration(sql string) bool {
	return strings.Contains(compactUpperSQL(sql), "DECLARE@RZA_MANUAL_DROP_TABLE_REQUESTBIT=1;")
}

func extractTableAlterColumnDefs(object, sql string) (map[string]tableColumnDef, map[string]tableColumnDef, error) {
	added := make(map[string]tableColumnDef)
	altered := make(map[string]tableColumnDef)
	for _, chunk := range extractAlterTableChunksForObject(object, sql) {
		upper := strings.ToUpper(chunk)
		var definition string
		var out map[string]tableColumnDef
		switch {
		case strings.HasPrefix(upper, "ADD "):
			definition = strings.TrimSpace(chunk[len("ADD "):])
			definitionUpper := strings.ToUpper(definition)
			if strings.HasPrefix(definitionUpper, "CONSTRAINT ") || strings.HasPrefix(definitionUpper, "DEFAULT ") {
				continue
			}
			out = added
		case strings.HasPrefix(upper, "ALTER COLUMN "):
			definition = strings.TrimSpace(chunk[len("ALTER COLUMN "):])
			out = altered
		default:
			continue
		}
		if definition == "" {
			return nil, nil, fmt.Errorf("empty column definition in ALTER TABLE statement for %s", object)
		}
		def, err := parseTableColumnDefFragment(definition)
		if err != nil {
			return nil, nil, err
		}
		out[strings.TrimSpace(def.Name)] = def
	}
	if len(added) == 0 && len(altered) == 0 {
		return nil, nil, fmt.Errorf("no ALTER TABLE ADD/ALTER COLUMN statement found for %s", object)
	}
	return added, altered, nil
}

func extractTableDefaultConstraintTransitions(object, sql string) ([]string, []tableDefaultConstraintDef) {
	dropped := make([]string, 0)
	added := make([]tableDefaultConstraintDef, 0)
	seenDrop := make(map[string]struct{})
	seenAdd := make(map[string]struct{})

	for _, chunk := range extractAlterTableChunksForObject(object, sql) {
		chunkUpper := strings.ToUpper(chunk)
		if strings.HasPrefix(chunkUpper, "DROP CONSTRAINT ") {
			nameToken := firstSQLToken(strings.TrimSpace(chunk[len("DROP CONSTRAINT "):]))
			name := stripIdentifierBrackets(cleanIdentifierToken(nameToken))
			if name != "" {
				key := name
				if _, ok := seenDrop[key]; !ok {
					seenDrop[key] = struct{}{}
					dropped = append(dropped, name)
				}
			}
		} else if strings.HasPrefix(chunkUpper, "ADD ") {
			addChunk := strings.TrimSpace(chunk[len("ADD "):])
			if def, ok := parseDefaultConstraintFromCreateTableItem(addChunk); ok {
				if strings.TrimSpace(def.Column) != "" && strings.TrimSpace(def.Expression) != "" {
					key := strings.TrimSpace(def.Column) + "|" +
						strings.TrimSpace(def.ConstraintName) + "|" +
						normalizeDefaultExpressionForCompare(def.Expression)
					if _, exists := seenAdd[key]; !exists {
						seenAdd[key] = struct{}{}
						added = append(added, def)
					}
				}
			}
		}
	}

	sort.Strings(dropped)
	sort.Slice(added, func(i, j int) bool {
		ci := strings.TrimSpace(added[i].Column)
		cj := strings.TrimSpace(added[j].Column)
		if ci != cj {
			return ci < cj
		}
		ni := strings.TrimSpace(added[i].ConstraintName)
		nj := strings.TrimSpace(added[j].ConstraintName)
		return ni < nj
	})
	return dropped, added
}

func extractTableCheckConstraintTransitions(object, sql string) ([]string, []tableCheckConstraintDef) {
	dropped := make([]string, 0)
	added := make([]tableCheckConstraintDef, 0)
	seenDrop := make(map[string]struct{})
	seenAdd := make(map[string]struct{})

	for _, chunk := range extractAlterTableChunksForObject(object, sql) {
		chunkUpper := strings.ToUpper(chunk)
		if strings.HasPrefix(chunkUpper, "DROP CONSTRAINT ") {
			nameToken := firstSQLToken(strings.TrimSpace(chunk[len("DROP CONSTRAINT "):]))
			name := stripIdentifierBrackets(cleanIdentifierToken(nameToken))
			if name != "" {
				key := name
				if _, ok := seenDrop[key]; !ok {
					seenDrop[key] = struct{}{}
					dropped = append(dropped, name)
				}
			}
		} else if strings.HasPrefix(chunkUpper, "ADD ") {
			addChunk := strings.TrimSpace(chunk[len("ADD "):])
			if def, ok := parseCheckConstraintFromCreateTableItem(addChunk); ok {
				key := strings.TrimSpace(def.Name) + "|" +
					strings.Join(strings.Fields(normalizeForSemanticHash(def.Definition)), " ")
				if _, exists := seenAdd[key]; !exists {
					seenAdd[key] = struct{}{}
					added = append(added, def)
				}
			}
		}
	}

	sort.Strings(dropped)
	sort.Slice(added, func(i, j int) bool {
		return strings.TrimSpace(added[i].Name) < strings.TrimSpace(added[j].Name)
	})
	return dropped, added
}

func extractAlterTableChunksForObject(object, sql string) []string {
	normalized := normalizeSQLText(sql, true, true)
	semanticUpper := strings.ToUpper(sqlparse.StripSQLStringsAndComments(normalized))
	const anchor = "ALTER TABLE "
	out := make([]string, 0)
	for cursor := 0; cursor < len(normalized); {
		rel := strings.Index(semanticUpper[cursor:], anchor)
		if rel < 0 {
			break
		}
		start := cursor + rel
		afterStart := start + len(anchor)
		after := normalized[afterStart:]
		tableToken, consumed := consumeLeadingObjectToken(after)
		if tableToken == "" || consumed <= 0 {
			cursor = afterStart
			continue
		}
		end := sqlparse.FindStatementEnd(normalized, start)
		if end <= afterStart+consumed {
			cursor = afterStart + consumed
			continue
		}
		target := stripIdentifierBrackets(cleanIdentifierToken(tableToken))
		if sameQualifiedObject(target, object) {
			out = append(out, strings.TrimSpace(strings.TrimSuffix(normalized[afterStart+consumed:end], ";")))
		}
		cursor = end
	}
	for _, batch := range extractSpExecuteSQLBatches(sql) {
		out = append(out, extractAlterTableChunksForObject(object, batch)...)
	}
	return out
}

func parseCheckConstraintFromCreateTableItem(line string) (tableCheckConstraintDef, bool) {
	trimmed := strings.TrimRight(strings.TrimSpace(line), ";")
	if trimmed == "" || constraintKindForTableItem(trimmed) != "check" {
		return tableCheckConstraintDef{}, false
	}
	name, ok := extractNamedConstraintName(trimmed)
	if !ok {
		return tableCheckConstraintDef{}, false
	}
	name = stripIdentifierBrackets(name)
	if name == "" {
		return tableCheckConstraintDef{}, false
	}
	return tableCheckConstraintDef{
		Name:       name,
		Definition: trimmed,
	}, true
}

func applyAddDefaultConstraintsToCreateTable(currentCreateTable string, defs []tableDefaultConstraintDef) (string, error) {
	working := normalizeSQLText(currentCreateTable, true, true)
	for _, def := range defs {
		column := strings.TrimSpace(def.Column)
		expr := strings.TrimSpace(def.Expression)
		if column == "" || expr == "" {
			continue
		}
		if strings.TrimSpace(def.ConstraintName) == "" {
			return "", fmt.Errorf("default constraint must be named for deterministic state reconstruction: column=%s", column)
		}

		items, err := extractCreateTableItems(working)
		if err != nil {
			return "", err
		}

		filtered := make([]string, 0, len(items))
		targetCol := column
		targetName := stripIdentifierBrackets(def.ConstraintName)
		for _, raw := range items {
			line := strings.TrimSpace(raw)
			if line == "" {
				continue
			}
			if named, ok := extractNamedConstraintName(line); ok &&
				stripIdentifierBrackets(named) == targetName {
				continue
			}
			if existingDef, ok := parseDefaultConstraintFromCreateTableItem(line); ok {
				if strings.TrimSpace(existingDef.Column) == targetCol {
					continue
				}
			}
			filtered = append(filtered, line)
		}

		defLine, err := buildDefaultConstraintDefinitionSQL(def)
		if err != nil {
			return "", err
		}

		insertIdx := len(filtered)
		for idx, line := range filtered {
			upperLine := strings.ToUpper(strings.TrimSpace(line))
			if isConstraintLikeTableItem(upperLine) || strings.HasPrefix(upperLine, "DEFAULT ") {
				insertIdx = idx
				break
			}
		}
		withInsert := append([]string{}, filtered[:insertIdx]...)
		withInsert = append(withInsert, defLine)
		withInsert = append(withInsert, filtered[insertIdx:]...)

		updated, err := rebuildCreateTableStatementWithItems(working, withInsert)
		if err != nil {
			return "", err
		}
		working = updated
	}
	return normalizeSQLText(working, true, true), nil
}

func parseTableColumnDefFragment(fragment string) (tableColumnDef, error) {
	trimmed := strings.TrimSpace(fragment)
	if trimmed == "" {
		return tableColumnDef{}, fmt.Errorf("empty column definition fragment")
	}
	columns, err := parseTableColumns("CREATE TABLE dbo.__mig_column_parse (\n  " + trimmed + "\n);")
	if err != nil {
		return tableColumnDef{}, err
	}
	if len(columns) != 1 {
		return tableColumnDef{}, fmt.Errorf("expected one column definition, got %d: %s", len(columns), fragment)
	}
	for _, def := range columns {
		return def, nil
	}
	return tableColumnDef{}, fmt.Errorf("column definition could not be parsed: %s", fragment)
}

func applyAddAlterColumnsToCreateTable(currentCreateTable string, added, altered map[string]tableColumnDef) (string, error) {
	stmt := normalizeSQLText(currentCreateTable, true, true)
	openIdx := strings.Index(stmt, "(")
	if openIdx < 0 {
		return "", fmt.Errorf("invalid CREATE TABLE state: no opening parenthesis")
	}

	closeIdx, ok := sqlparse.FindMatchingParen(stmt, openIdx)
	if !ok {
		return "", fmt.Errorf("invalid CREATE TABLE state: no closing parenthesis")
	}

	inside := stmt[openIdx+1 : closeIdx]
	items := splitTopLevelByComma(inside)
	updatedItems := make([]string, 0, len(items)+len(added))
	existingCols := make(map[string]struct{})
	remainingAltered := make(map[string]tableColumnDef, len(altered))
	for name, def := range altered {
		remainingAltered[name] = def
	}

	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		upper := strings.ToUpper(line)
		if isConstraintLikeTableItem(upper) {
			updatedItems = append(updatedItems, line)
			continue
		}

		nameToken := firstSQLToken(line)
		colName := columnIdentifierName(nameToken)
		if colName == "" {
			updatedItems = append(updatedItems, line)
			continue
		}
		existingCols[colName] = struct{}{}

		if def, ok := remainingAltered[colName]; ok {
			updatedItems = append(updatedItems, columnDefinitionSQL(def))
			delete(remainingAltered, colName)
			continue
		}
		updatedItems = append(updatedItems, line)
	}

	if len(remainingAltered) > 0 {
		keys := make([]string, 0, len(remainingAltered))
		for name := range remainingAltered {
			keys = append(keys, name)
		}
		sort.Strings(keys)
		return "", fmt.Errorf("alter column target(s) not found in current state: %s", strings.Join(keys, ", "))
	}

	addNames := make([]string, 0, len(added))
	for name := range added {
		if _, exists := existingCols[name]; exists {
			return "", fmt.Errorf("cannot add column %s because it already exists", name)
		}
		addNames = append(addNames, name)
	}
	sort.Strings(addNames)

	insertIdx := len(updatedItems)
	for idx, item := range updatedItems {
		if isConstraintLikeTableItem(strings.ToUpper(item)) {
			insertIdx = idx
			break
		}
	}

	addLines := make([]string, 0, len(addNames))
	for _, name := range addNames {
		addLines = append(addLines, columnDefinitionSQL(added[name]))
	}

	if len(addLines) > 0 {
		before := append([]string{}, updatedItems[:insertIdx]...)
		after := append([]string{}, updatedItems[insertIdx:]...)
		updatedItems = append(before, addLines...)
		updatedItems = append(updatedItems, after...)
	}

	if len(updatedItems) == 0 {
		return "", fmt.Errorf("table definition became empty after add/alter apply")
	}

	rebuiltInside := "  " + strings.Join(updatedItems, ",\n  ") + "\n"
	rebuilt := stmt[:openIdx+1] + "\n" + rebuiltInside + stmt[closeIdx:]
	return normalizeSQLText(rebuilt, true, true), nil
}

func isConstraintLikeTableItem(upperLine string) bool {
	return strings.HasPrefix(upperLine, "CONSTRAINT ") ||
		strings.HasPrefix(upperLine, "PRIMARY KEY") ||
		strings.HasPrefix(upperLine, "UNIQUE ") ||
		strings.HasPrefix(upperLine, "FOREIGN KEY") ||
		strings.HasPrefix(upperLine, "CHECK ")
}

func extractDroppedColumns(object, sql string) ([]string, error) {
	type set map[string]struct{}
	removed := make(set)
	found := false
	for _, chunk := range extractAlterTableChunksForObject(object, sql) {
		upper := strings.ToUpper(chunk)
		if !strings.HasPrefix(upper, "DROP COLUMN ") {
			continue
		}
		list := strings.TrimSpace(chunk[len("DROP COLUMN "):])
		if list == "" {
			return nil, fmt.Errorf("empty DROP COLUMN list for %s", object)
		}
		for _, token := range splitTopLevelByComma(list) {
			col := strings.TrimSpace(token)
			if col == "" {
				continue
			}
			if strings.HasPrefix(col, "[") && strings.HasSuffix(col, "]") && len(col) >= 2 {
				col = col[1 : len(col)-1]
			}
			col = strings.TrimSpace(col)
			if col == "" {
				continue
			}
			removed[col] = struct{}{}
			found = true
		}
	}

	if !found {
		return nil, fmt.Errorf("no ALTER TABLE ... DROP COLUMN statement found for %s", object)
	}

	out := make([]string, 0, len(removed))
	for col := range removed {
		out = append(out, col)
	}
	sort.Strings(out)
	return out, nil
}

func applyDropColumnsToCreateTable(currentCreateTable string, removedColumns []string) (string, error) {
	stmt := normalizeSQLText(currentCreateTable, true, true)
	openIdx := strings.Index(stmt, "(")
	if openIdx < 0 {
		return "", fmt.Errorf("invalid CREATE TABLE state: no opening parenthesis")
	}

	closeIdx, ok := sqlparse.FindMatchingParen(stmt, openIdx)
	if !ok {
		return "", fmt.Errorf("invalid CREATE TABLE state: no closing parenthesis")
	}

	inside := stmt[openIdx+1 : closeIdx]
	items := splitTopLevelByComma(inside)
	removeSet := make(map[string]struct{}, len(removedColumns))
	for _, col := range removedColumns {
		removeSet[strings.TrimSpace(col)] = struct{}{}
	}

	filtered := make([]string, 0, len(items))
	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		if shouldRemoveTableItemForDroppedColumn(line, removeSet) {
			continue
		}
		filtered = append(filtered, line)
	}

	if len(filtered) == 0 {
		return "", fmt.Errorf("drop-column result produced empty table definition")
	}

	rebuiltInside := "  " + strings.Join(filtered, ",\n  ") + "\n"
	rebuilt := stmt[:openIdx+1] + "\n" + rebuiltInside + stmt[closeIdx:]
	return normalizeSQLText(rebuilt, true, true), nil
}

func shouldRemoveTableItemForDroppedColumn(item string, removeSet map[string]struct{}) bool {
	name := firstSQLToken(item)
	nameTrimmed := columnIdentifierName(name)
	if _, ok := removeSet[nameTrimmed]; ok {
		return true
	}

	upper := strings.ToUpper(item)
	if strings.HasPrefix(upper, "CONSTRAINT ") ||
		strings.HasPrefix(upper, "PRIMARY KEY") ||
		strings.HasPrefix(upper, "UNIQUE ") ||
		strings.HasPrefix(upper, "FOREIGN KEY") ||
		strings.HasPrefix(upper, "CHECK ") {
		for col := range removeSet {
			if containsColumnReference(item, col) {
				return true
			}
		}
	}
	return false
}

func containsColumnReference(input, column string) bool {
	return sqlparse.ContainsIdentifierExact(input, strings.TrimSpace(column))
}

func splitTopLevelByComma(input string) []string {
	return sqlparse.SplitTopLevelComma(input)
}

func validateCreateTableStructure(createTableStmt, objectName string) error {
	items, err := extractCreateTableItems(createTableStmt)
	if err != nil {
		if objectName == "" {
			return fmt.Errorf("invalid CREATE TABLE structure: %w", err)
		}
		return fmt.Errorf("invalid CREATE TABLE structure for %s: %w", objectName, err)
	}
	seenColumns := make(map[string]struct{}, len(items))
	columnCount := 0
	for _, item := range items {
		trimmed := trimLeadingSQLComments(item)
		if trimmed == "" || isConstraintLikeTableItem(strings.ToUpper(trimmed)) {
			continue
		}
		def, parseErr := parseTableColumnDefFragment(trimmed)
		if parseErr != nil {
			continue // the subset validator reports the more specific syntax error
		}
		name := stripIdentifierBrackets(strings.TrimSpace(def.Name))
		if name == "" {
			continue
		}
		if _, duplicate := seenColumns[name]; duplicate {
			if objectName == "" {
				return fmt.Errorf("duplicate exact-case column name %q in CREATE TABLE", name)
			}
			return fmt.Errorf("duplicate exact-case column name %q in CREATE TABLE for %s", name, objectName)
		}
		seenColumns[name] = struct{}{}
		columnCount++
	}
	if columnCount == 0 {
		if objectName == "" {
			return fmt.Errorf("CREATE TABLE must contain at least one column")
		}
		return fmt.Errorf("CREATE TABLE for %s must contain at least one column", objectName)
	}
	return nil
}

func trimLeadingSQLComments(input string) string {
	remaining := strings.TrimSpace(input)
	for remaining != "" {
		switch {
		case strings.HasPrefix(remaining, "--"):
			lineEnd := strings.IndexAny(remaining, "\r\n")
			if lineEnd < 0 {
				return ""
			}
			remaining = strings.TrimSpace(remaining[lineEnd:])
		case strings.HasPrefix(remaining, "/*"):
			blockEnd := strings.Index(remaining[2:], "*/")
			if blockEnd < 0 {
				return remaining
			}
			remaining = strings.TrimSpace(remaining[2+blockEnd+2:])
		default:
			return remaining
		}
	}
	return remaining
}

func validateCreateTableCommaSeparators(items []string) error {
	for idx, raw := range items {
		if strings.TrimSpace(raw) == "" {
			if idx == len(items)-1 {
				return fmt.Errorf("trailing comma before closing ')' in CREATE TABLE item list")
			}
			return fmt.Errorf("empty table item detected (possible duplicate comma)")
		}
	}
	for _, raw := range items {
		starts := collectCreateTableItemStarts(raw)
		if len(starts) <= 1 {
			continue
		}
		return fmt.Errorf("missing comma between table items near %q and %q", starts[0], starts[1])
	}
	return nil
}

func repairCreateTableCommaSyntax(createTableStmt string) (string, bool) {
	stmt := normalizeSQLText(createTableStmt, true, true)
	openIdx := strings.Index(stmt, "(")
	if openIdx < 0 {
		return createTableStmt, false
	}
	closeIdx, ok := sqlparse.FindMatchingParen(stmt, openIdx)
	if !ok {
		return createTableStmt, false
	}

	inside := stmt[openIdx+1 : closeIdx]
	commaParts := sqlparse.SplitTopLevelComma(inside)
	recovered := make([]string, 0, len(commaParts))
	changed := false
	for _, rawPart := range commaParts {
		part := strings.TrimSpace(rawPart)
		if part == "" {
			changed = true // duplicate comma or trailing comma
			continue
		}

		offsets := append([]int{0}, sqlparse.TopLevelLineOffsets(part)...)
		splitAt := make([]int, 0, len(offsets))
		for _, offset := range offsets {
			if offset < 0 || offset >= len(part) {
				continue
			}
			lineEnd := strings.IndexAny(part[offset:], "\r\n")
			if lineEnd < 0 {
				lineEnd = len(part) - offset
			}
			line := strings.TrimSpace(part[offset : offset+lineEnd])
			if line != "" && isLikelyCreateTableItemStart(line) {
				splitAt = append(splitAt, offset)
			}
		}
		if len(splitAt) <= 1 {
			recovered = append(recovered, part)
			continue
		}
		changed = true
		for i, start := range splitAt {
			end := len(part)
			if i+1 < len(splitAt) {
				end = splitAt[i+1]
			}
			item := strings.TrimSpace(part[start:end])
			if item != "" {
				recovered = append(recovered, item)
			}
		}
	}
	if !changed || len(recovered) == 0 {
		return createTableStmt, false
	}
	rebuilt, err := rebuildCreateTableStatementWithItems(stmt, recovered)
	if err != nil {
		return createTableStmt, false
	}
	if err := validateCreateTableStructure(rebuilt, ""); err != nil {
		return createTableStmt, false
	}
	return rebuilt, true
}

func collectCreateTableItemStarts(part string) []string {
	semantic := sqlparse.StripSQLStringsAndComments(part)
	lines := strings.Split(semantic, "\n")
	out := make([]string, 0, 2)
	for _, ln := range lines {
		trimmed := strings.TrimSpace(ln)
		if trimmed == "" {
			continue
		}
		if !isLikelyCreateTableItemStart(trimmed) {
			continue
		}
		out = append(out, compactSQLPreview(trimmed))
	}
	return out
}

func isLikelyCreateTableItemStart(line string) bool {
	// A standalone remark can contain ordinary words that resemble
	// "<column> <type>". Classify the line from its SQL-bearing portion so a
	// comment above a column is never recovered as a separate table item.
	semanticLine := strings.TrimSpace(sqlparse.StripSQLStringsAndComments(line))
	if semanticLine == "" {
		return false
	}

	upper := strings.ToUpper(semanticLine)
	if upper == "" {
		return false
	}
	if strings.HasPrefix(upper, "CONSTRAINT ") ||
		strings.HasPrefix(upper, "PRIMARY KEY") ||
		strings.HasPrefix(upper, "UNIQUE ") ||
		strings.HasPrefix(upper, "FOREIGN KEY") ||
		strings.HasPrefix(upper, "CHECK ") {
		return true
	}

	first := firstSQLToken(semanticLine)
	if first == "" {
		return false
	}
	firstName := cleanIdentifierToken(first)
	if firstName == "" {
		return false
	}

	rest := strings.TrimSpace(semanticLine[len(first):])
	if rest == "" {
		return false
	}
	second := strings.ToUpper(strings.TrimSpace(firstSQLToken(rest)))
	if second == "" {
		return false
	}
	if second[0] < 'A' || second[0] > 'Z' {
		return false
	}
	switch second {
	case "ON", "REFERENCES", "ASC", "DESC", "AND", "OR", "ELSE", "THEN", "END":
		return false
	}
	return true
}

func compactSQLPreview(line string) string {
	trimmed := strings.TrimSpace(line)
	if len(trimmed) <= 80 {
		return trimmed
	}
	return strings.TrimSpace(trimmed[:77]) + "..."
}

func firstSQLToken(input string) string {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" {
		return ""
	}
	for i := 0; i < len(trimmed); i++ {
		if trimmed[i] == ' ' || trimmed[i] == '\t' || trimmed[i] == '\n' || trimmed[i] == '\r' {
			return trimmed[:i]
		}
	}
	return trimmed
}

func firstSQLIdentifierToken(input string) string {
	trimmed := strings.TrimSpace(input)
	if trimmed == "" || trimmed[0] != '[' {
		return firstSQLToken(trimmed)
	}
	for index := 1; index < len(trimmed); index++ {
		if trimmed[index] != ']' {
			continue
		}
		if index+1 < len(trimmed) && trimmed[index+1] == ']' {
			index++
			continue
		}
		return trimmed[:index+1]
	}
	return trimmed
}

type tableColumnDef struct {
	Name               string
	BaseType           string
	RawType            string
	Length             int
	Nullable           bool
	IsUnicode          bool
	IsComputed         bool
	ComputedDefinition string
	IdentityDefinition string
	RowGuidCol         bool
	Sparse             bool
}

type tableDefaultConstraintDef struct {
	Column         string
	ConstraintName string
	Expression     string
}

type tableCheckConstraintDef struct {
	Name       string
	Definition string
}

func detectUnsafeAlterColumn(previousCreateTable, nextCreateTable string) ([]string, error) {
	prevCols, err := parseTableColumns(previousCreateTable)
	if err != nil {
		return nil, err
	}
	nextCols, err := parseTableColumns(nextCreateTable)
	if err != nil {
		return nil, err
	}

	issues := make([]string, 0)
	for name, prev := range prevCols {
		next, ok := nextCols[name]
		if !ok {
			continue
		}

		if prev.Nullable && !next.Nullable {
			issues = append(issues, fmt.Sprintf("NULL_TO_NOT_NULL(%s)", name))
		}

		if isStringType(prev.BaseType) && isStringType(next.BaseType) {
			if prev.Length > 0 && next.Length > 0 && next.Length < prev.Length {
				if next.IsUnicode {
					issues = append(issues, fmt.Sprintf("SHRINK_NSTRING(%s:%d->%d)", name, prev.Length, next.Length))
				} else {
					issues = append(issues, fmt.Sprintf("SHRINK_STRING(%s:%d->%d)", name, prev.Length, next.Length))
				}
			}
		}

		if prev.BaseType != next.BaseType && isNumericType(next.BaseType) {
			issues = append(issues, fmt.Sprintf("TYPE_TO_NUMERIC(%s:%s->%s)", name, prev.BaseType, next.BaseType))
		}
	}

	sort.Strings(issues)
	return issues, nil
}

func parseTableColumns(createTableStmt string) (map[string]tableColumnDef, error) {
	stmt := normalizeSQLText(createTableStmt, true, true)
	openIdx := strings.Index(stmt, "(")
	if openIdx < 0 {
		return nil, fmt.Errorf("invalid CREATE TABLE statement: no opening parenthesis")
	}

	closeIdx, ok := sqlparse.FindMatchingParen(stmt, openIdx)
	if !ok {
		return nil, fmt.Errorf("invalid CREATE TABLE statement: no closing parenthesis")
	}

	inside := stmt[openIdx+1 : closeIdx]
	items := splitTopLevelByComma(inside)
	out := make(map[string]tableColumnDef)

	for _, item := range items {
		line := strings.TrimSpace(item)
		if line == "" {
			continue
		}
		upper := strings.ToUpper(line)
		if strings.HasPrefix(upper, "CONSTRAINT ") ||
			strings.HasPrefix(upper, "PRIMARY KEY") ||
			strings.HasPrefix(upper, "UNIQUE ") ||
			strings.HasPrefix(upper, "FOREIGN KEY") ||
			strings.HasPrefix(upper, "CHECK ") {
			continue
		}

		nameToken, consumed := consumeLeadingColumnIdentifierToken(line)
		if nameToken == "" {
			continue
		}
		name := columnIdentifierName(nameToken)
		if name == "" {
			continue
		}
		key := name

		rest := strings.TrimSpace(line[consumed:])
		if hasKeywordAtWordBoundary(strings.ToUpper(rest), 0, "AS") {
			definition := strings.TrimSpace(rest)
			if len(definition) == len("AS") {
				return nil, fmt.Errorf("computed column %s has no expression", name)
			}
			if existing, exists := out[key]; exists {
				if existing.IsComputed && existing.ComputedDefinition == definition {
					continue
				}
				return nil, fmt.Errorf("duplicate column definition detected for %s", name)
			}
			out[key] = tableColumnDef{
				Name:               name,
				BaseType:           "COMPUTED",
				Nullable:           true,
				IsComputed:         true,
				ComputedDefinition: definition,
			}
			continue
		}
		typeToken := firstSQLToken(rest)
		if typeToken == "" {
			continue
		}
		rawType := parseColumnRawTypeSQL(rest, typeToken)

		baseType, length, isUnicode := parseTypeToken(rawType)
		if strings.Contains(strings.ToUpper(rawType), " COLLATE ") && isKnownNonCollatableType(baseType) {
			return nil, fmt.Errorf("column %s uses COLLATE with non-character type %s", name, baseType)
		}
		nullable := true
		restUpper := strings.ToUpper(rest)
		if strings.Contains(restUpper, " NOT NULL") {
			nullable = false
		} else if strings.Contains(restUpper, " NULL") {
			nullable = true
		}

		if existing, exists := out[key]; exists {
			if strings.EqualFold(existing.RawType, rawType) && existing.Nullable == nullable {
				continue
			}
			return nil, fmt.Errorf("duplicate column definition detected for %s", name)
		}
		identity := extractIdentityDefinition(rest)
		rowGuid := strings.Contains(restUpper, " ROWGUIDCOL")
		if rowGuid && baseType != "UNIQUEIDENTIFIER" {
			return nil, fmt.Errorf("ROWGUIDCOL column %s must use UNIQUEIDENTIFIER", name)
		}
		sparse := strings.Contains(restUpper, " SPARSE")
		if sparse && !nullable {
			return nil, fmt.Errorf("SPARSE column %s must be nullable", name)
		}
		out[key] = tableColumnDef{
			Name:               name,
			BaseType:           baseType,
			RawType:            rawType,
			Length:             length,
			Nullable:           nullable,
			IsUnicode:          isUnicode,
			IdentityDefinition: identity,
			RowGuidCol:         rowGuid,
			Sparse:             sparse,
		}
	}

	return out, nil
}

func extractIdentityDefinition(rest string) string {
	upper := strings.ToUpper(rest)
	pos := strings.Index(upper, "IDENTITY")
	if pos < 0 {
		return ""
	}
	after := strings.TrimSpace(rest[pos+len("IDENTITY"):])
	if !strings.HasPrefix(after, "(") {
		return "IDENTITY"
	}
	inside, ok := extractLeadingParenthesizedContent(after)
	if !ok {
		return ""
	}
	return "IDENTITY(" + strings.TrimSpace(inside) + ")"
}

func extractTableDefaultConstraintMap(createTableStmt string) (map[string]tableDefaultConstraintDef, error) {
	items, err := extractCreateTableItems(createTableStmt)
	if err != nil {
		return nil, err
	}

	out := make(map[string]tableDefaultConstraintDef)
	for _, raw := range items {
		line := strings.TrimSpace(raw)
		if line == "" {
			continue
		}
		def, ok := parseDefaultConstraintFromCreateTableItem(line)
		if !ok {
			continue
		}
		if def.Column == "" || strings.TrimSpace(def.Expression) == "" {
			continue
		}
		key := strings.TrimSpace(def.Column)
		if existing, exists := out[key]; exists {
			_ = existing
			return nil, fmt.Errorf("multiple DEFAULT constraints detected for column %s", key)
		}
		out[key] = def
	}
	return out, nil
}

func parseDefaultConstraintFromCreateTableItem(line string) (tableDefaultConstraintDef, bool) {
	trimmed := strings.TrimSpace(line)
	if trimmed == "" {
		return tableDefaultConstraintDef{}, false
	}
	upper := strings.ToUpper(trimmed)
	if !strings.Contains(upper, "DEFAULT") {
		return tableDefaultConstraintDef{}, false
	}

	if tableDef, ok := parseTableLevelDefaultConstraint(trimmed, upper); ok {
		return tableDef, true
	}
	if inlineDef, ok := parseInlineColumnDefaultConstraint(trimmed, upper); ok {
		return inlineDef, true
	}
	return tableDefaultConstraintDef{}, false
}

func parseTableLevelDefaultConstraint(line, upper string) (tableDefaultConstraintDef, bool) {
	defaultPositions := findAllKeywordPositions(upper, "DEFAULT")
	if len(defaultPositions) == 0 {
		return tableDefaultConstraintDef{}, false
	}

	for _, defaultPos := range defaultPositions {
		expr, exprEnd, ok := consumeDefaultExpression(line, defaultPos+len("DEFAULT"))
		if !ok {
			continue
		}
		restUpper := upper[exprEnd:]
		forPositions := findAllKeywordPositions(restUpper, "FOR")
		if len(forPositions) == 0 {
			continue
		}
		forPos := exprEnd + forPositions[0]
		afterFor := strings.TrimSpace(line[forPos+len("FOR"):])
		columnToken := firstSQLToken(afterFor)
		columnName := stripIdentifierBrackets(cleanIdentifierToken(columnToken))
		if columnName == "" {
			continue
		}

		constraintName, _ := extractNamedConstraintName(line)
		constraintName = stripIdentifierBrackets(constraintName)
		return tableDefaultConstraintDef{
			Column:         columnName,
			ConstraintName: constraintName,
			Expression:     strings.TrimSpace(expr),
		}, true
	}
	return tableDefaultConstraintDef{}, false
}

func parseInlineColumnDefaultConstraint(line, upper string) (tableDefaultConstraintDef, bool) {
	if strings.HasPrefix(upper, "CONSTRAINT ") ||
		strings.HasPrefix(upper, "PRIMARY KEY") ||
		strings.HasPrefix(upper, "UNIQUE ") ||
		strings.HasPrefix(upper, "FOREIGN KEY") ||
		strings.HasPrefix(upper, "CHECK ") {
		return tableDefaultConstraintDef{}, false
	}

	nameToken := firstSQLToken(line)
	if nameToken == "" {
		return tableDefaultConstraintDef{}, false
	}
	columnName := stripIdentifierBrackets(cleanIdentifierToken(nameToken))
	if columnName == "" {
		return tableDefaultConstraintDef{}, false
	}

	rest := strings.TrimSpace(line[len(nameToken):])
	if rest == "" {
		return tableDefaultConstraintDef{}, false
	}
	restUpper := strings.ToUpper(rest)
	defaultPositions := findAllKeywordPositions(restUpper, "DEFAULT")
	if len(defaultPositions) == 0 {
		return tableDefaultConstraintDef{}, false
	}

	absOffset := len(line) - len(rest)
	for _, defaultPos := range defaultPositions {
		defaultAbs := absOffset + defaultPos
		expr, _, ok := consumeDefaultExpression(line, defaultAbs+len("DEFAULT"))
		if !ok {
			continue
		}

		prefix := rest[:defaultPos]
		constraintName := ""
		prefixUpper := strings.ToUpper(prefix)
		if positions := findAllKeywordPositions(prefixUpper, "CONSTRAINT"); len(positions) > 0 {
			idx := positions[len(positions)-1]
			candidate := strings.TrimSpace(prefix[idx+len("CONSTRAINT"):])
			constraintName = stripIdentifierBrackets(cleanIdentifierToken(firstSQLToken(candidate)))
		}

		return tableDefaultConstraintDef{
			Column:         columnName,
			ConstraintName: constraintName,
			Expression:     strings.TrimSpace(expr),
		}, true
	}
	return tableDefaultConstraintDef{}, false
}

func consumeDefaultExpression(input string, start int) (string, int, bool) {
	if start < 0 || start >= len(input) {
		return "", start, false
	}

	i := start
	for i < len(input) && unicode.IsSpace(rune(input[i])) {
		i++
	}
	if i >= len(input) {
		return "", start, false
	}

	if input[i] == '(' {
		depth := 0
		inString := false
		for j := i; j < len(input); j++ {
			ch := input[j]
			if inString {
				if ch == '\'' {
					if j+1 < len(input) && input[j+1] == '\'' {
						j++
						continue
					}
					inString = false
				}
				continue
			}
			switch ch {
			case '\'':
				inString = true
			case '(':
				depth++
			case ')':
				depth--
				if depth == 0 {
					return strings.TrimSpace(input[i : j+1]), j + 1, true
				}
			}
		}
		return "", start, false
	}

	j := i
	for j < len(input) {
		ch := input[j]
		if unicode.IsSpace(rune(ch)) || ch == ',' || ch == ';' {
			break
		}
		j++
	}
	expr := strings.TrimSpace(input[i:j])
	if expr == "" {
		return "", start, false
	}
	return expr, j, true
}

func parseTypeToken(typeToken string) (baseType string, length int, isUnicode bool) {
	token := strings.ToUpper(strings.TrimSpace(typeToken))
	if collatePos := strings.Index(token, " COLLATE "); collatePos >= 0 {
		token = strings.TrimSpace(token[:collatePos])
	}
	openIdx := strings.Index(token, "(")
	if openIdx < 0 {
		base := normalizeColumnBaseType(token)
		return base, 0, base == "NCHAR" || base == "NVARCHAR"
	}

	base := normalizeColumnBaseType(token[:openIdx])
	inner := ""
	if closeIdx := strings.LastIndex(token, ")"); closeIdx > openIdx {
		inner = strings.TrimSpace(token[openIdx+1 : closeIdx])
	}
	if strings.EqualFold(inner, "MAX") {
		return base, -1, base == "NCHAR" || base == "NVARCHAR"
	}
	if inner == "" {
		return base, 0, base == "NCHAR" || base == "NVARCHAR"
	}

	lengthValue := 0
	for i := 0; i < len(inner); i++ {
		ch := inner[i]
		if ch < '0' || ch > '9' {
			if i == 0 {
				lengthValue = 0
			}
			break
		}
		lengthValue = lengthValue*10 + int(ch-'0')
	}
	return base, lengthValue, base == "NCHAR" || base == "NVARCHAR"
}

// normalizeColumnBaseType canonicalizes only a single delimited type name.
// Extracted SQL Server definitions commonly spell built-ins as [DECIMAL] or
// [VARCHAR]; those forms must participate in the same diff/risk rules as their
// unquoted equivalents. Schema-qualified alias/CLR types remain qualified and
// therefore retain the conservative manual-transition policy.
func normalizeColumnBaseType(baseType string) string {
	token := strings.ToUpper(strings.TrimSpace(baseType))
	if token == "" || strings.Contains(token, ".") {
		return token
	}
	if unquoted, ok := unquoteSingleTypeIdentifier(token); ok {
		return strings.ToUpper(strings.TrimSpace(unquoted))
	}
	return token
}

func normalizeColumnTypeDeclaration(rawType string) string {
	trimmed := strings.ToUpper(strings.TrimSpace(rawType))
	typeName, cursor := consumeColumnTypeName(trimmed)
	if typeName == "" {
		return trimmed
	}
	suffix := strings.TrimSpace(trimmed[cursor:])
	if suffix == "" || strings.HasPrefix(suffix, "(") {
		return normalizeColumnBaseType(typeName) + suffix
	}
	return normalizeColumnBaseType(typeName) + " " + suffix
}

func columnTypeEquivalent(before, after tableColumnDef) bool {
	return normalizeColumnBaseType(before.BaseType) == normalizeColumnBaseType(after.BaseType) &&
		normalizeColumnTypeDeclaration(before.RawType) == normalizeColumnTypeDeclaration(after.RawType)
}

func unquoteSingleTypeIdentifier(token string) (string, bool) {
	if len(token) < 2 {
		return "", false
	}
	var delimiter, escaped byte
	switch token[0] {
	case '[':
		delimiter, escaped = ']', ']'
	case '"':
		delimiter, escaped = '"', '"'
	default:
		return "", false
	}

	var builder strings.Builder
	for cursor := 1; cursor < len(token); cursor++ {
		if token[cursor] != delimiter {
			builder.WriteByte(token[cursor])
			continue
		}
		if cursor+1 < len(token) && token[cursor+1] == escaped {
			builder.WriteByte(delimiter)
			cursor++
			continue
		}
		return builder.String(), cursor == len(token)-1
	}
	return "", false
}

func parseColumnRawTypeSQL(rest, typeToken string) string {
	trimmed := strings.TrimSpace(rest)
	typeName, cursor := consumeColumnTypeName(trimmed)
	if typeName == "" {
		return strings.ToUpper(strings.TrimSpace(typeToken))
	}

	raw := typeName
	cursor = skipSQLWhitespace(trimmed, cursor)
	if cursor < len(trimmed) && trimmed[cursor] == '(' {
		if closeIdx, ok := sqlparse.FindMatchingParen(trimmed, cursor); ok {
			parameters := normalizeColumnTypeParameters(trimmed[cursor+1 : closeIdx])
			raw += "(" + parameters + ")"
			cursor = closeIdx + 1
		}
	}

	cursor = skipSQLWhitespace(trimmed, cursor)
	if cursor >= len(trimmed) || !hasKeywordAtWordBoundary(strings.ToUpper(trimmed), cursor, "COLLATE") {
		return raw
	}
	cursor = skipSQLWhitespace(trimmed, cursor+len("COLLATE"))
	collateName := firstSQLIdentifierToken(trimmed[cursor:])
	if collateName == "" {
		return raw
	}
	return raw + " COLLATE " + collateName
}

// consumeColumnTypeName reads a built-in or schema-qualified SQL type without
// consuming its optional parameter list. SQL Server permits whitespace both
// around the qualification dot and before the opening parenthesis, so a plain
// whitespace token is not sufficient for column-diff semantics.
func consumeColumnTypeName(input string) (string, int) {
	parts := make([]string, 0, 2)
	cursor := 0
	for {
		cursor = skipSQLWhitespace(input, cursor)
		if cursor >= len(input) {
			break
		}

		start := cursor
		switch input[cursor] {
		case '[':
			cursor++
			for cursor < len(input) {
				if input[cursor] != ']' {
					cursor++
					continue
				}
				cursor++
				if cursor < len(input) && input[cursor] == ']' {
					cursor++
					continue
				}
				break
			}
		case '"':
			cursor++
			for cursor < len(input) {
				if input[cursor] != '"' {
					cursor++
					continue
				}
				cursor++
				if cursor < len(input) && input[cursor] == '"' {
					cursor++
					continue
				}
				break
			}
		default:
			for cursor < len(input) && !isSQLWhitespace(input[cursor]) && input[cursor] != '.' && input[cursor] != '(' {
				cursor++
			}
		}

		part := strings.TrimSpace(input[start:cursor])
		if part == "" {
			break
		}
		parts = append(parts, strings.ToUpper(part))

		next := skipSQLWhitespace(input, cursor)
		if next >= len(input) || input[next] != '.' {
			return strings.Join(parts, "."), cursor
		}
		cursor = next + 1
	}
	return strings.Join(parts, "."), cursor
}

func skipSQLWhitespace(input string, cursor int) int {
	for cursor < len(input) && isSQLWhitespace(input[cursor]) {
		cursor++
	}
	return cursor
}

func isSQLWhitespace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n'
}

func normalizeColumnTypeParameters(input string) string {
	canonical := strings.Join(strings.Fields(strings.TrimSpace(input)), " ")
	for _, replacement := range []struct{ old, new string }{
		{" ,", ","}, {", ", ","}, {"( ", "("}, {" )", ")"},
	} {
		canonical = strings.ReplaceAll(canonical, replacement.old, replacement.new)
	}
	return strings.ToUpper(canonical)
}

func isStringType(baseType string) bool {
	switch normalizeColumnBaseType(baseType) {
	case "CHAR", "VARCHAR", "NCHAR", "NVARCHAR":
		return true
	default:
		return false
	}
}

func isKnownNonCollatableType(baseType string) bool {
	switch normalizeColumnBaseType(baseType) {
	case "BIGINT", "BINARY", "BIT", "DATE", "DATETIME", "DATETIME2", "DATETIMEOFFSET",
		"DECIMAL", "FLOAT", "GEOGRAPHY", "GEOMETRY", "HIERARCHYID", "IMAGE", "INT",
		"MONEY", "NUMERIC", "REAL", "ROWVERSION", "SMALLDATETIME", "SMALLINT",
		"SMALLMONEY", "SQL_VARIANT", "TIME", "TIMESTAMP", "TINYINT", "UNIQUEIDENTIFIER",
		"VARBINARY", "XML":
		return true
	default:
		// Unknown names may be alias types backed by a collatable character type.
		return false
	}
}

func isBinaryType(baseType string) bool {
	switch normalizeColumnBaseType(baseType) {
	case "BINARY", "VARBINARY":
		return true
	default:
		return false
	}
}

func isNumericType(baseType string) bool {
	switch normalizeColumnBaseType(baseType) {
	case "INT", "BIGINT", "SMALLINT", "TINYINT", "BIT", "DECIMAL", "NUMERIC", "MONEY", "SMALLMONEY", "FLOAT", "REAL":
		return true
	default:
		return false
	}
}

func isDateTimeType(baseType string) bool {
	switch normalizeColumnBaseType(baseType) {
	case "DATE", "TIME", "DATETIME", "DATETIME2", "SMALLDATETIME", "DATETIMEOFFSET":
		return true
	default:
		return false
	}
}
