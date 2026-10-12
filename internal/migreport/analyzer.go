package migreport

import (
	"fmt"
	"sort"
	"strings"
	"time"
)

// TableState tracks known table state in memory across versions.
type TableState struct {
	Columns            map[string]ParsedColumn     // colName -> def
	Constraints        map[string]ParsedConstraint // conName -> def
	DroppedConstraints map[string]ParsedConstraint // conName -> def
}

// IndexState tracks known indexes in memory.
type IndexState struct {
	Indexes        map[string]ParsedIndex // idxName -> def
	DroppedIndexes map[string]ParsedIndex // idxName -> previous def within current history
}

// ProgrammableState tracks known programmables in memory.
type ProgrammableState struct {
	NormalizedHash string
	RawDefinition  string
}

// TypeState tracks known types in memory.
type TypeState struct {
	NormalizedHash string
	RawDefinition  string
}

// ObjectVersionTracker tracks details of a specific version of an object.
type ObjectVersionTracker struct {
	Version          string
	FirstMigrationID string
	LastMigrationID  string
	FirstUpFile      string
	LastUpFile       string
	Status           string // INITIAL, UPDATED, DELETED, REINTRODUCED
	ChangeCount      int
}

// ObjectTracking tracks version history of an object.
type ObjectTracking struct {
	Category           string
	SchemaName         string
	ObjectName         string
	OrderedVersions    []string
	VersionMap         map[string]*ObjectVersionTracker
	FirstMigrationID   string
	LastMigrationID    string
	FirstUpFile        string
	LastUpFile         string
	IsCurrentlyDeleted bool
}

func getOrCreateTableState(tableStates map[string]*TableState, schema, table string) *TableState {
	key := strings.ToLower(fmt.Sprintf("%s.%s", schema, table))
	ts := tableStates[key]
	if ts == nil {
		ts = &TableState{
			Columns:            make(map[string]ParsedColumn),
			Constraints:        make(map[string]ParsedConstraint),
			DroppedConstraints: make(map[string]ParsedConstraint),
		}
		tableStates[key] = ts
	}
	return ts
}

func getOrCreateIndexState(indexStates map[string]*IndexState, schema, table string) *IndexState {
	key := strings.ToLower(fmt.Sprintf("%s.%s", schema, table))
	state := indexStates[key]
	if state == nil {
		state = &IndexState{
			Indexes:        make(map[string]ParsedIndex),
			DroppedIndexes: make(map[string]ParsedIndex),
		}
		indexStates[key] = state
	}
	return state
}

// Analyze processes scanned migration pairs and reconstructs the full change history.
func Analyze(repoRoot string, pairs []MigrationPair, allFiles []FileInfo, initialWarnings []string, initialErrors []string) (*RunResult, error) {
	result := &RunResult{
		Warnings: append([]string{}, initialWarnings...),
		Errors:   append([]string{}, initialErrors...),
	}

	tableStates := make(map[string]*TableState) // schema.table -> state
	indexStates := make(map[string]*IndexState) // schema.table -> state
	progStates := make(map[string]*ProgrammableState)
	typeStates := make(map[string]*TypeState)
	objectMap := make(map[string]*ObjectTracking) // cat|schema|obj -> tracker

	var migrationSummaries []MigrationSummary
	var allTableChanges []TableChange
	var allIndexChanges []IndexChange
	var allProgChanges []ProgrammableChange
	var allTypeChanges []TypeChange

	totalVersionsSet := make(map[string]struct{})

	for _, pair := range pairs {
		migID := pair.MigrationID
		cat := pair.Category
		schema := pair.SchemaName
		obj := pair.ObjectName
		ver := pair.ObjectVersion

		upExists := "NO"
		downExists := "NO"
		upFileRel := ""
		downFileRel := ""
		var upSQL, downSQL string

		if pair.UpFile != nil {
			upExists = "YES"
			upFileRel = pair.UpFile.RelPath
			upSQL = string(pair.UpFile.RawBytes)
		}
		if pair.DownFile != nil {
			downExists = "YES"
			downFileRel = pair.DownFile.RelPath
			downSQL = string(pair.DownFile.RawBytes)
		}

		roundTrip := RoundTripPass
		migStatus := StatusOk
		if !pair.IsPaired {
			roundTrip = RoundTripUnpaired
			migStatus = StatusWarning
			result.Warnings = append(result.Warnings, fmt.Sprintf("unpaired migration %s (%s.%s)", migID, schema, obj))
		}

		// Track object version
		objKey := fmt.Sprintf("%s|%s|%s", cat, schema, obj)
		tracker, exists := objectMap[objKey]
		if !exists {
			tracker = &ObjectTracking{
				Category:         cat,
				SchemaName:       schema,
				ObjectName:       obj,
				OrderedVersions:  nil,
				VersionMap:       make(map[string]*ObjectVersionTracker),
				FirstMigrationID: migID,
				LastMigrationID:  migID,
				FirstUpFile:      upFileRel,
				LastUpFile:       upFileRel,
			}
			objectMap[objKey] = tracker
		}
		tracker.LastMigrationID = migID
		if upFileRel != "" {
			if tracker.FirstUpFile == "" {
				tracker.FirstUpFile = upFileRel
			}
			tracker.LastUpFile = upFileRel
		}

		verTracker, verExists := tracker.VersionMap[ver]
		if !verExists {
			tracker.OrderedVersions = append(tracker.OrderedVersions, ver)
			verStatus := HistoryUpdated
			if len(tracker.OrderedVersions) == 1 {
				verStatus = HistoryInitial
			}
			verTracker = &ObjectVersionTracker{
				Version:          ver,
				FirstMigrationID: migID,
				LastMigrationID:  migID,
				FirstUpFile:      upFileRel,
				LastUpFile:       upFileRel,
				Status:           verStatus,
			}
			tracker.VersionMap[ver] = verTracker
		} else {
			verTracker.LastMigrationID = migID
			if upFileRel != "" {
				if verTracker.FirstUpFile == "" {
					verTracker.FirstUpFile = upFileRel
				}
				verTracker.LastUpFile = upFileRel
			}
		}
		totalVersionsSet[fmt.Sprintf("%s|%s", objKey, ver)] = struct{}{}

		// State keys
		tblKey := fmt.Sprintf("%s.%s", schema, obj)

		currProgState := progStates[tblKey]
		if currProgState == nil {
			currProgState = &ProgrammableState{}
			progStates[tblKey] = currProgState
		}

		currTypeState := typeStates[tblKey]
		if currTypeState == nil {
			currTypeState = &TypeState{}
			typeStates[tblKey] = currTypeState
		}

		// Parse up SQL
		parsedTableStmts, parsedIndexes, parsedProg, parsedType := ParseStatementsFromSQL(cat, schema, obj, upSQL)
		recognizedUpSQL := len(parsedTableStmts) > 0 || len(parsedIndexes) > 0 || parsedProg != nil || parsedType != nil

		var opKinds []string
		opNo := 1
		tableChangesStart := len(allTableChanges)

		// 1. Process Table changes
		for _, stmt := range parsedTableStmts {
			reversible := ReversibleYes
			if stmt.Kind == OpDropTable {
				reversible = ReversibleNo
			}

			tSchema := stmt.SchemaName
			if tSchema == "" {
				tSchema = schema
			}
			tName := stmt.TableName
			if tName == "" {
				tName = obj
			}

			currTableState := getOrCreateTableState(tableStates, tSchema, tName)

			tc := TableChange{
				MigrationID:    migID,
				SchemaName:     tSchema,
				TableName:      tName,
				ObjectVersion:  ver,
				OperationNo:    opNo,
				OperationKind:  stmt.Kind,
				TargetName:     stmt.TargetName,
				OldValue:       stmt.OldValue,
				NewValue:       stmt.NewValue,
				DependencyName: stmt.DependencyName,
				UpFile:         upFileRel,
				DownFile:       downFileRel,
				Reversible:     reversible,
				Confidence:     ConfidenceHigh,
				Notes:          stmt.Notes,
			}

			emitChange := true

			// Apply to state and refine old_value/new_value
			switch stmt.Kind {
			case OpCreateTable:
				currTableState.Columns = make(map[string]ParsedColumn)
				currTableState.Constraints = make(map[string]ParsedConstraint)
				currTableState.DroppedConstraints = make(map[string]ParsedConstraint)
				indexState := getOrCreateIndexState(indexStates, tSchema, tName)
				indexState.Indexes = make(map[string]ParsedIndex)
				indexState.DroppedIndexes = make(map[string]ParsedIndex)
				if tracker.IsCurrentlyDeleted {
					verTracker.Status = HistoryReintroduced
					tracker.IsCurrentlyDeleted = false
				}
			case OpDropTable:
				currTableState.Columns = make(map[string]ParsedColumn)
				currTableState.Constraints = make(map[string]ParsedConstraint)
				currTableState.DroppedConstraints = make(map[string]ParsedConstraint)
				indexState := getOrCreateIndexState(indexStates, tSchema, tName)
				indexState.Indexes = make(map[string]ParsedIndex)
				indexState.DroppedIndexes = make(map[string]ParsedIndex)
				tracker.IsCurrentlyDeleted = true
				verTracker.Status = HistoryDeleted
			case OpAddColumn:
				if stmt.Column != nil {
					currTableState.Columns[strings.ToLower(stmt.TargetName)] = *stmt.Column
					tc.NewValue = stmt.Column.RawType
				}
			case OpDropColumn:
				if prevCol, ok := currTableState.Columns[strings.ToLower(stmt.TargetName)]; ok {
					tc.OldValue = prevCol.RawType
					delete(currTableState.Columns, strings.ToLower(stmt.TargetName))
				} else {
					tc.Confidence = ConfidenceLow
					tc.Notes = NoteManualReviewRequired
				}
			case OpRenameColumn:
				if prevCol, ok := currTableState.Columns[strings.ToLower(stmt.OldValue)]; ok {
					delete(currTableState.Columns, strings.ToLower(stmt.OldValue))
					prevCol.Name = stmt.NewValue
					currTableState.Columns[strings.ToLower(stmt.NewValue)] = prevCol
				} else {
					tc.Confidence = ConfidenceLow
					tc.Notes = NoteManualReviewRequired
				}
			case OpAlterColumnType:
				if prevCol, ok := currTableState.Columns[strings.ToLower(stmt.TargetName)]; ok {
					tc.OldValue = prevCol.RawType
					if equalSQLFragment(prevCol.RawType, stmt.NewValue) {
						emitChange = false
					}
					prevCol.RawType = stmt.NewValue
					currTableState.Columns[strings.ToLower(stmt.TargetName)] = prevCol
				} else {
					tc.Confidence = ConfidenceLow
					tc.Notes = NoteManualReviewRequired
				}
			case OpAlterNullability:
				if prevCol, ok := currTableState.Columns[strings.ToLower(stmt.TargetName)]; ok {
					if prevCol.Nullable {
						tc.OldValue = "NULL"
					} else {
						tc.OldValue = "NOT NULL"
					}
					newNullable := stmt.NewValue == "NULL"
					if prevCol.Nullable == newNullable {
						emitChange = false
					}
					prevCol.Nullable = newNullable
					currTableState.Columns[strings.ToLower(stmt.TargetName)] = prevCol
				} else {
					tc.Confidence = ConfidenceLow
					tc.Notes = NoteManualReviewRequired
				}
			case OpAddConstraint:
				if stmt.Constraint != nil {
					lowerName := strings.ToLower(stmt.TargetName)
					if prevCon, exists := currTableState.Constraints[lowerName]; exists {
						if prevCon.Type == "DEFAULT" || stmt.Constraint.Type == "DEFAULT" {
							tc.OperationKind = OpAlterDefault
							tc.OldValue = prevCon.RawDefinition
							tc.NewValue = stmt.Constraint.RawDefinition
							tc.Confidence = ConfidenceHigh
						} else if !CheckIdenticalConstraintDef(prevCon.RawDefinition, stmt.Constraint.RawDefinition) {
							tc.OperationKind = OpAlterConstraint
							tc.OldValue = prevCon.RawDefinition
							tc.NewValue = stmt.Constraint.RawDefinition
							tc.Confidence = ConfidenceMedium
							tc.Notes = "definition changed"
						}
					} else if droppedCon, wasDropped := currTableState.DroppedConstraints[lowerName]; wasDropped {
						if droppedCon.Type == "DEFAULT" || stmt.Constraint.Type == "DEFAULT" {
							tc.OperationKind = OpAlterDefault
							tc.OldValue = droppedCon.RawDefinition
							tc.NewValue = stmt.Constraint.RawDefinition
							tc.Confidence = ConfidenceHigh
						} else if !CheckIdenticalConstraintDef(droppedCon.RawDefinition, stmt.Constraint.RawDefinition) {
							tc.OperationKind = OpAlterConstraint
							tc.OldValue = droppedCon.RawDefinition
							tc.NewValue = stmt.Constraint.RawDefinition
							tc.Confidence = ConfidenceMedium
							tc.Notes = "definition changed"
						}
						delete(currTableState.DroppedConstraints, lowerName)
					} else {
						for prevName, prevCon := range currTableState.Constraints {
							if CheckIdenticalConstraintDef(prevCon.RawDefinition, stmt.Constraint.RawDefinition) && !strings.EqualFold(prevName, stmt.TargetName) {
								tc.Notes = "constraint renamed with identical definition"
								break
							}
						}
					}
					currTableState.Constraints[lowerName] = *stmt.Constraint
					tc.NewValue = stmt.Constraint.RawDefinition
				}
			case OpDropConstraint:
				lowerName := strings.ToLower(stmt.TargetName)
				if prevCon, ok := currTableState.Constraints[lowerName]; ok {
					tc.OldValue = prevCon.RawDefinition
					currTableState.DroppedConstraints[lowerName] = prevCon
					delete(currTableState.Constraints, lowerName)
				}
			case OpRenameConstraint:
				if prevCon, ok := currTableState.Constraints[strings.ToLower(stmt.OldValue)]; ok {
					delete(currTableState.Constraints, strings.ToLower(stmt.OldValue))
					prevCon.Name = stmt.NewValue
					currTableState.Constraints[strings.ToLower(stmt.NewValue)] = prevCon
					emitChange = false
				} else {
					tc.Confidence = ConfidenceLow
					tc.Notes = NoteManualReviewRequired
				}
			}

			if emitChange {
				allTableChanges = append(allTableChanges, tc)
				opNo++
			}
		}

		// A drop followed by an add under a different name but with an
		// identical definition is metadata-only constraint reconciliation.
		// It must not inflate the semantic change history.
		pairTableChanges := suppressConstraintRenameNoise(allTableChanges[tableChangesStart:])
		allTableChanges = append(allTableChanges[:tableChangesStart], pairTableChanges...)
		for i := tableChangesStart; i < len(allTableChanges); i++ {
			allTableChanges[i].OperationNo = i - tableChangesStart + 1
			opKinds = append(opKinds, allTableChanges[i].OperationKind)
		}
		opNo = len(pairTableChanges) + 1

		// 2. Process Index changes
		indexChangesStart := len(allIndexChanges)
		tableKindsCount := len(opKinds)
		indexOperationBase := opNo
		for _, idx := range parsedIndexes {
			indexSchema := schema
			operationIndexState := getOrCreateIndexState(indexStates, indexSchema, idx.TableName)
			if idx.IndexType == "RENAME" {
				opKind := OpRenameIndex
				oldIdxName := idx.RawDefinition
				oldColStr := ""
				oldIncStr := ""
				oldFilter := "NULL"
				oldUnique := "0"
				oldClustered := "0"

				if prevIdx, ok := operationIndexState.Indexes[strings.ToLower(oldIdxName)]; ok {
					oldColStr = strings.Join(prevIdx.Columns, ",")
					oldIncStr = strings.Join(prevIdx.IncludeColumns, ",")
					oldFilter = prevIdx.Filter
					if prevIdx.IsUnique {
						oldUnique = "1"
					}
					if prevIdx.IsClustered {
						oldClustered = "1"
					}
					delete(operationIndexState.Indexes, strings.ToLower(oldIdxName))
					operationIndexState.Indexes[strings.ToLower(idx.IndexName)] = prevIdx
				}
				notes := fmt.Sprintf("renamed from %s", oldIdxName)
				ic := IndexChange{
					MigrationID:       migID,
					SchemaName:        schema,
					TableName:         idx.TableName,
					IndexName:         idx.IndexName,
					ObjectVersion:     ver,
					OperationNo:       opNo,
					OperationKind:     opKind,
					IndexType:         "NONCLUSTERED",
					OldColumns:        oldColStr,
					NewColumns:        oldColStr,
					OldIncludeColumns: oldIncStr,
					NewIncludeColumns: oldIncStr,
					OldFilter:         oldFilter,
					NewFilter:         oldFilter,
					OldUnique:         oldUnique,
					NewUnique:         oldUnique,
					OldClustered:      oldClustered,
					NewClustered:      oldClustered,
					UpFile:            upFileRel,
					DownFile:          downFileRel,
					Reversible:        ReversibleYes,
					Confidence:        ConfidenceHigh,
					Notes:             notes,
				}
				allIndexChanges = append(allIndexChanges, ic)
				opKinds = append(opKinds, opKind)
				opNo++
				continue
			}

			opKind := OpAddIndex
			oldColStr := ""
			oldIncStr := ""
			oldFilter := "NULL"
			oldUnique := "0"
			oldClustered := "0"

			newColStr := strings.Join(idx.Columns, ",")
			newIncStr := strings.Join(idx.IncludeColumns, ",")
			newFilter := idx.Filter
			newUnique := "0"
			if idx.IsUnique {
				newUnique = "1"
			}
			newClustered := "0"
			if idx.IsClustered {
				newClustered = "1"
			}

			lowerIdxName := strings.ToLower(idx.IndexName)
			if idx.IndexType == "DROP" {
				opKind = OpDropIndex
				if prevIdx, ok := operationIndexState.Indexes[lowerIdxName]; ok {
					oldColStr = strings.Join(prevIdx.Columns, ",")
					oldIncStr = strings.Join(prevIdx.IncludeColumns, ",")
					oldFilter = prevIdx.Filter
					if prevIdx.IsUnique {
						oldUnique = "1"
					}
					if prevIdx.IsClustered {
						oldClustered = "1"
					}
					operationIndexState.DroppedIndexes[lowerIdxName] = prevIdx
					delete(operationIndexState.Indexes, lowerIdxName)
				}
			} else {
				if prevIdx, ok := operationIndexState.Indexes[lowerIdxName]; ok {
					opKind = OpAlterIndex
					oldColStr = strings.Join(prevIdx.Columns, ",")
					oldIncStr = strings.Join(prevIdx.IncludeColumns, ",")
					oldFilter = prevIdx.Filter
					if prevIdx.IsUnique {
						oldUnique = "1"
					}
					if prevIdx.IsClustered {
						oldClustered = "1"
					}
				} else if prevIdx, ok := operationIndexState.DroppedIndexes[lowerIdxName]; ok {
					opKind = OpAlterIndex
					oldColStr = strings.Join(prevIdx.Columns, ",")
					oldIncStr = strings.Join(prevIdx.IncludeColumns, ",")
					oldFilter = prevIdx.Filter
					if prevIdx.IsUnique {
						oldUnique = "1"
					}
					if prevIdx.IsClustered {
						oldClustered = "1"
					}
					delete(operationIndexState.DroppedIndexes, lowerIdxName)
				}
				operationIndexState.Indexes[lowerIdxName] = idx
			}

			ic := IndexChange{
				MigrationID:       migID,
				SchemaName:        schema,
				TableName:         idx.TableName,
				IndexName:         idx.IndexName,
				ObjectVersion:     ver,
				OperationNo:       opNo,
				OperationKind:     opKind,
				IndexType:         idx.IndexType,
				OldColumns:        oldColStr,
				NewColumns:        newColStr,
				OldIncludeColumns: oldIncStr,
				NewIncludeColumns: newIncStr,
				OldFilter:         oldFilter,
				NewFilter:         newFilter,
				OldUnique:         oldUnique,
				NewUnique:         newUnique,
				OldClustered:      oldClustered,
				NewClustered:      newClustered,
				UpFile:            upFileRel,
				DownFile:          downFileRel,
				Reversible:        ReversibleYes,
				Confidence:        ConfidenceHigh,
			}
			allIndexChanges = append(allIndexChanges, ic)
			opKinds = append(opKinds, opKind)
			opNo++
		}
		pairIndexChanges := suppressMechanicalIndexDrops(allIndexChanges[indexChangesStart:])
		allIndexChanges = append(allIndexChanges[:indexChangesStart], pairIndexChanges...)
		opKinds = opKinds[:tableKindsCount]
		for i := indexChangesStart; i < len(allIndexChanges); i++ {
			allIndexChanges[i].OperationNo = indexOperationBase + i - indexChangesStart
			opKinds = append(opKinds, allIndexChanges[i].OperationKind)
		}
		opNo = indexOperationBase + len(pairIndexChanges)

		// 3. Process Programmable changes
		if parsedProg != nil {
			oldHash := currProgState.NormalizedHash
			newHash := parsedProg.DefinitionHash

			// Check if semantic change or just formatting/noise
			if oldHash != "" && oldHash == newHash {
				// Pure formatting change, skip reporting as semantic operation
			} else {
				opKind := parsedProg.Kind
				if parsedProg.Kind == OpProgDrop {
					opKind = OpProgDrop
					tracker.IsCurrentlyDeleted = true
					verTracker.Status = HistoryDeleted
				} else if oldHash == "" {
					opKind = OpProgCreate
					if tracker.IsCurrentlyDeleted {
						verTracker.Status = HistoryReintroduced
						tracker.IsCurrentlyDeleted = false
					}
				} else {
					opKind = OpProgAlter
				}

				changedSec := DetectChangedSections(currProgState.RawDefinition, parsedProg.RawSQL)

				pc := ProgrammableChange{
					MigrationID:          migID,
					Category:             cat,
					SchemaName:           schema,
					ObjectName:           obj,
					ObjectVersion:        ver,
					OperationKind:        opKind,
					DefinitionHashBefore: oldHash,
					DefinitionHashAfter:  newHash,
					ChangedSections:      changedSec,
					UpFile:               upFileRel,
					DownFile:             downFileRel,
					Reversible:           ReversibleYes,
					Confidence:           ConfidenceHigh,
				}
				allProgChanges = append(allProgChanges, pc)
				opKinds = append(opKinds, opKind)
				opNo++

				if parsedProg.Kind == OpProgDrop {
					currProgState.NormalizedHash = ""
					currProgState.RawDefinition = ""
				} else {
					currProgState.NormalizedHash = newHash
					currProgState.RawDefinition = parsedProg.RawSQL
				}
			}
		}

		// 4. Process Type changes
		if parsedType != nil {
			oldDef := currTypeState.RawDefinition
			newDef := parsedType.RawDefinition
			oldHash := currTypeState.NormalizedHash
			newHash := ComputeNormalizedHash(newDef)

			if oldHash != "" && oldHash == newHash {
				// Formatting noise only
			} else {
				opKind := parsedType.Kind
				if parsedType.Kind == OpDropType {
					opKind = OpDropType
					tracker.IsCurrentlyDeleted = true
					verTracker.Status = HistoryDeleted
				} else if oldHash == "" {
					opKind = OpCreateType
					if tracker.IsCurrentlyDeleted {
						verTracker.Status = HistoryReintroduced
						tracker.IsCurrentlyDeleted = false
					}
				}

				tyc := TypeChange{
					MigrationID:   migID,
					SchemaName:    schema,
					TypeName:      obj,
					ObjectVersion: ver,
					OperationKind: opKind,
					OldDefinition: oldDef,
					NewDefinition: newDef,
					UpFile:        upFileRel,
					DownFile:      downFileRel,
					Reversible:    ReversibleYes,
					Confidence:    ConfidenceHigh,
				}
				allTypeChanges = append(allTypeChanges, tyc)
				opKinds = append(opKinds, opKind)
				opNo++

				if parsedType.Kind == OpDropType {
					currTypeState.NormalizedHash = ""
					currTypeState.RawDefinition = ""
				} else {
					currTypeState.NormalizedHash = newHash
					currTypeState.RawDefinition = newDef
				}
			}
		}

		// Check for ambiguous SQL or mixed objects
		if !recognizedUpSQL && pair.UpFile != nil && len(strings.TrimSpace(StripComments(upSQL))) > 0 {
			// File had SQL but no statements were recognized
			allTableChanges = append(allTableChanges, TableChange{
				MigrationID:   migID,
				SchemaName:    schema,
				TableName:     obj,
				ObjectVersion: ver,
				OperationNo:   1,
				OperationKind: "SQL_AMBIGUOUS",
				TargetName:    obj,
				UpFile:        upFileRel,
				DownFile:      downFileRel,
				Reversible:    ReversibleNo,
				Confidence:    ConfidenceLow,
				Notes:         NoteManualReviewRequired,
			})
			opKinds = append(opKinds, "SQL_AMBIGUOUS")
			opNo++
		}

		// Validate that the down script contains the semantic inverse of each
		// recognized top-level up operation. Child operations of CREATE TABLE
		// are covered by a matching DROP TABLE.
		if pair.IsPaired {
			roundTrip = evaluateRoundTrip(cat, schema, obj, upSQL, downSQL,
				parsedTableStmts, parsedIndexes, parsedProg, parsedType, recognizedUpSQL)
		}
		if roundTrip == RoundTripFail {
			migStatus = StatusWarning
			result.Warnings = append(result.Warnings, fmt.Sprintf("round-trip validation failed for migration %s (%s.%s)", migID, schema, obj))
		}

		// Unique operation kinds list
		uniqueKinds := uniqueStrings(opKinds)
		opKindsStr := strings.Join(uniqueKinds, "|")

		migrationSummaries = append(migrationSummaries, MigrationSummary{
			MigrationID:     migID,
			Category:        cat,
			SchemaName:      schema,
			ObjectName:      obj,
			ObjectVersion:   ver,
			OperationCount:  len(opKinds),
			OperationKinds:  opKindsStr,
			UpFile:          upFileRel,
			DownFile:        downFileRel,
			UpExists:        upExists,
			DownExists:      downExists,
			RoundTripStatus: roundTrip,
			Status:          migStatus,
		})
	}

	// Build ObjectHistory
	var objectHistories []ObjectHistory
	var sortedObjKeys []string
	for k := range objectMap {
		sortedObjKeys = append(sortedObjKeys, k)
	}
	sort.Strings(sortedObjKeys)

	for _, k := range sortedObjKeys {
		tracker := objectMap[k]
		for verIdx, v := range tracker.OrderedVersions {
			verTracker := tracker.VersionMap[v]
			prevVer := ""
			if verIdx > 0 {
				prevVer = tracker.OrderedVersions[verIdx-1]
			}

			// Count operations for this object and version
			chgCount := 0
			for _, tc := range allTableChanges {
				if tc.SchemaName == tracker.SchemaName && (tc.TableName == tracker.ObjectName || tracker.Category == "core") && tc.ObjectVersion == v {
					chgCount++
				}
			}
			for _, ic := range allIndexChanges {
				if ic.SchemaName == tracker.SchemaName && (ic.TableName == tracker.ObjectName || tracker.Category == "core") && ic.ObjectVersion == v {
					chgCount++
				}
			}
			for _, pc := range allProgChanges {
				if pc.SchemaName == tracker.SchemaName && pc.ObjectName == tracker.ObjectName && pc.ObjectVersion == v {
					chgCount++
				}
			}
			for _, ty := range allTypeChanges {
				if ty.SchemaName == tracker.SchemaName && ty.TypeName == tracker.ObjectName && ty.ObjectVersion == v {
					chgCount++
				}
			}
			if chgCount == 0 {
				for _, tc := range allTableChanges {
					if tc.MigrationID == verTracker.FirstMigrationID {
						chgCount++
					}
				}
				if chgCount == 0 {
					chgCount = 1
				}
			}

			objectHistories = append(objectHistories, ObjectHistory{
				Category:         tracker.Category,
				SchemaName:       tracker.SchemaName,
				ObjectName:       tracker.ObjectName,
				ObjectVersion:    v,
				FirstMigrationID: verTracker.FirstMigrationID,
				LastMigrationID:  verTracker.LastMigrationID,
				PreviousVersion:  prevVer,
				ChangeCount:      chgCount,
				FirstUpFile:      verTracker.FirstUpFile,
				LastUpFile:       verTracker.LastUpFile,
				HistoryStatus:    verTracker.Status,
			})
		}
	}

	// Sort datasets deterministically: migration_id, category, schema_name, object_name, operation_no
	sort.Slice(allTableChanges, func(i, j int) bool {
		if allTableChanges[i].MigrationID != allTableChanges[j].MigrationID {
			return allTableChanges[i].MigrationID < allTableChanges[j].MigrationID
		}
		if allTableChanges[i].SchemaName != allTableChanges[j].SchemaName {
			return allTableChanges[i].SchemaName < allTableChanges[j].SchemaName
		}
		if allTableChanges[i].TableName != allTableChanges[j].TableName {
			return allTableChanges[i].TableName < allTableChanges[j].TableName
		}
		return allTableChanges[i].OperationNo < allTableChanges[j].OperationNo
	})

	sort.Slice(allIndexChanges, func(i, j int) bool {
		if allIndexChanges[i].MigrationID != allIndexChanges[j].MigrationID {
			return allIndexChanges[i].MigrationID < allIndexChanges[j].MigrationID
		}
		if allIndexChanges[i].SchemaName != allIndexChanges[j].SchemaName {
			return allIndexChanges[i].SchemaName < allIndexChanges[j].SchemaName
		}
		if allIndexChanges[i].TableName != allIndexChanges[j].TableName {
			return allIndexChanges[i].TableName < allIndexChanges[j].TableName
		}
		return allIndexChanges[i].OperationNo < allIndexChanges[j].OperationNo
	})

	sort.Slice(allProgChanges, func(i, j int) bool {
		if allProgChanges[i].MigrationID != allProgChanges[j].MigrationID {
			return allProgChanges[i].MigrationID < allProgChanges[j].MigrationID
		}
		if allProgChanges[i].Category != allProgChanges[j].Category {
			return allProgChanges[i].Category < allProgChanges[j].Category
		}
		if allProgChanges[i].SchemaName != allProgChanges[j].SchemaName {
			return allProgChanges[i].SchemaName < allProgChanges[j].SchemaName
		}
		return allProgChanges[i].ObjectName < allProgChanges[j].ObjectName
	})

	sort.Slice(allTypeChanges, func(i, j int) bool {
		if allTypeChanges[i].MigrationID != allTypeChanges[j].MigrationID {
			return allTypeChanges[i].MigrationID < allTypeChanges[j].MigrationID
		}
		if allTypeChanges[i].SchemaName != allTypeChanges[j].SchemaName {
			return allTypeChanges[i].SchemaName < allTypeChanges[j].SchemaName
		}
		return allTypeChanges[i].TypeName < allTypeChanges[j].TypeName
	})

	// Count distinct objects per category
	distinctTables := make(map[string]bool)
	distinctIndexes := make(map[string]bool)
	distinctProcedures := make(map[string]bool)
	distinctViews := make(map[string]bool)
	distinctFunctions := make(map[string]bool)
	distinctTypes := make(map[string]bool)

	for _, tc := range allTableChanges {
		distinctTables[fmt.Sprintf("%s.%s", tc.SchemaName, tc.TableName)] = true
	}
	for _, ic := range allIndexChanges {
		distinctIndexes[fmt.Sprintf("%s.%s", ic.TableName, ic.IndexName)] = true
	}
	for _, pc := range allProgChanges {
		key := fmt.Sprintf("%s.%s", pc.SchemaName, pc.ObjectName)
		switch pc.Category {
		case "procedures":
			distinctProcedures[key] = true
		case "views":
			distinctViews[key] = true
		case "functions":
			distinctFunctions[key] = true
		}
	}
	for _, ty := range allTypeChanges {
		distinctTypes[fmt.Sprintf("%s.%s", ty.SchemaName, ty.TypeName)] = true
	}

	totalOps := len(allTableChanges) + len(allIndexChanges) + len(allProgChanges) + len(allTypeChanges)

	unpairedFiles := 0
	for _, p := range pairs {
		if !p.IsPaired {
			unpairedFiles++
		}
	}

	overallStatus := StatusOk
	if len(result.Errors) > 0 {
		overallStatus = StatusFailed
	} else if len(result.Warnings) > 0 {
		overallStatus = StatusWarning
	}

	result.Summary = ReportSummary{
		RunAt:           time.Now(),
		Repository:      repoRoot,
		MigrationsDir:   "migrations",
		TotalFiles:      len(allFiles),
		TotalPairs:      len(pairs),
		UnpairedFiles:   unpairedFiles,
		TotalObjects:    len(objectMap),
		TotalVersions:   len(totalVersionsSet),
		TotalOperations: totalOps,
		Tables:          len(distinctTables),
		Indexes:         len(distinctIndexes),
		Procedures:      len(distinctProcedures),
		Views:           len(distinctViews),
		Functions:       len(distinctFunctions),
		Types:           len(distinctTypes),
		Warnings:        len(result.Warnings),
		Errors:          len(result.Errors),
		Status:          overallStatus,
	}

	result.Migrations = migrationSummaries
	result.ObjectHistory = objectHistories
	result.TableChanges = allTableChanges
	result.IndexChanges = allIndexChanges
	result.ProgChanges = allProgChanges
	result.TypeChanges = allTypeChanges

	result.JSONReport = JSONReport{
		Schema:         "RZA_MIGRATION_REPORT_V1",
		GeneratedAt:    time.Now().Format(time.RFC3339),
		Repository:     repoRoot,
		Source:         "migrations-and-db_object",
		DatabaseAccess: false,
		SQLExecution:   false,
		FilesScanned:   len(allFiles),
		MigrationPairs: len(pairs),
		Objects:        len(objectMap),
		Operations:     totalOps,
		Reports: []string{
			"report_summary.csv",
			"migration_summary.csv",
			"object_history.csv",
			"table_changes.csv",
			"index_changes.csv",
			"programmable_changes.csv",
			"type_changes.csv",
			"report.json",
		},
		Warnings: result.Warnings,
		Errors:   result.Errors,
	}

	return result, nil
}

func uniqueStrings(slice []string) []string {
	seen := make(map[string]bool)
	var out []string
	for _, s := range slice {
		if !seen[s] {
			seen[s] = true
			out = append(out, s)
		}
	}
	return out
}

func equalSQLFragment(left, right string) bool {
	left = strings.Join(strings.Fields(strings.ToUpper(strings.TrimSpace(left))), " ")
	right = strings.Join(strings.Fields(strings.ToUpper(strings.TrimSpace(right))), " ")
	return left == right
}

func suppressConstraintRenameNoise(changes []TableChange) []TableChange {
	suppressed := make(map[int]bool)
	// DROP + ADD of the same named constraint is represented by the analyzer
	// as one ALTER_* semantic transition. Remove the mechanical DROP row so
	// old_value/new_value stay together in a single operation.
	for dropIdx, drop := range changes {
		if drop.OperationKind != OpDropConstraint {
			continue
		}
		for _, alter := range changes {
			if alter.OperationKind != OpAlterConstraint && alter.OperationKind != OpAlterDefault {
				continue
			}
			if strings.EqualFold(drop.SchemaName, alter.SchemaName) &&
				strings.EqualFold(drop.TableName, alter.TableName) &&
				strings.EqualFold(drop.TargetName, alter.TargetName) {
				suppressed[dropIdx] = true
				break
			}
		}
	}

	for dropIdx, drop := range changes {
		if drop.OperationKind != OpDropConstraint || strings.TrimSpace(drop.OldValue) == "" {
			continue
		}
		for addIdx, add := range changes {
			if suppressed[addIdx] || add.OperationKind != OpAddConstraint || strings.TrimSpace(add.NewValue) == "" {
				continue
			}
			if !strings.EqualFold(drop.SchemaName, add.SchemaName) ||
				!strings.EqualFold(drop.TableName, add.TableName) ||
				strings.EqualFold(drop.TargetName, add.TargetName) {
				continue
			}
			if CheckIdenticalConstraintDef(drop.OldValue, add.NewValue) {
				suppressed[dropIdx] = true
				suppressed[addIdx] = true
				break
			}
		}
	}

	out := make([]TableChange, 0, len(changes)-len(suppressed))
	for idx, change := range changes {
		if !suppressed[idx] {
			out = append(out, change)
		}
	}
	return out
}

func suppressMechanicalIndexDrops(changes []IndexChange) []IndexChange {
	suppressed := make(map[int]bool)
	for dropIdx, drop := range changes {
		if drop.OperationKind != OpDropIndex {
			continue
		}
		for _, replacement := range changes {
			if replacement.OperationKind == OpAlterIndex &&
				strings.EqualFold(drop.SchemaName, replacement.SchemaName) &&
				strings.EqualFold(drop.TableName, replacement.TableName) &&
				strings.EqualFold(drop.IndexName, replacement.IndexName) {
				suppressed[dropIdx] = true
				break
			}
		}
	}

	out := make([]IndexChange, 0, len(changes)-len(suppressed))
	for idx, change := range changes {
		if !suppressed[idx] {
			out = append(out, change)
		}
	}
	return out
}

func evaluateRoundTrip(
	category, schema, objectName, upSQL, downSQL string,
	upTables []ParsedTableStatement,
	upIndexes []ParsedIndex,
	upProg *ParsedProgrammableStatement,
	upType *ParsedTypeStatement,
	recognizedUp bool,
) string {
	upText := strings.TrimSpace(StripComments(upSQL))
	downText := strings.TrimSpace(StripComments(downSQL))
	if upText == "" && downText == "" {
		return RoundTripPass
	}
	if upText == "" || downText == "" {
		return RoundTripFail
	}
	if !recognizedUp {
		return RoundTripNotTested
	}

	downTables, downIndexes, downProg, downType := ParseStatementsFromSQL(category, schema, objectName, downSQL)
	recognizedDown := len(downTables) > 0 || len(downIndexes) > 0 || downProg != nil || downType != nil
	if !recognizedDown {
		return RoundTripNotTested
	}

	for _, up := range upTables {
		if !hasInverseTableOperation(up, upTables, downTables) {
			return RoundTripFail
		}
	}
	for _, up := range upIndexes {
		if !hasInverseIndexOperation(up, downIndexes, downTables) {
			return RoundTripFail
		}
	}
	if upProg != nil {
		if downProg == nil || !programmableKindsAreInverse(upProg.Kind, downProg.Kind) {
			return RoundTripFail
		}
	}
	if upType != nil {
		if downType == nil || !typeKindsAreInverse(upType.Kind, downType.Kind) {
			return RoundTripFail
		}
	}
	return RoundTripPass
}

func hasInverseTableOperation(up ParsedTableStatement, allUp, down []ParsedTableStatement) bool {
	if up.Kind == OpAddColumn || up.Kind == OpAddConstraint {
		for _, parent := range allUp {
			if parent.Kind == OpCreateTable && sameTable(parent, up) && hasTableOperation(down, OpDropTable, up.SchemaName, up.TableName, "") {
				return true
			}
		}
	}

	switch up.Kind {
	case OpCreateTable:
		return hasTableOperation(down, OpDropTable, up.SchemaName, up.TableName, "")
	case OpDropTable:
		return hasTableOperation(down, OpCreateTable, up.SchemaName, up.TableName, "")
	case OpAddColumn:
		return hasTableOperation(down, OpDropColumn, up.SchemaName, up.TableName, up.TargetName)
	case OpDropColumn:
		return hasTableOperation(down, OpAddColumn, up.SchemaName, up.TableName, up.TargetName)
	case OpAddConstraint:
		return hasTableOperation(down, OpDropConstraint, up.SchemaName, up.TableName, up.TargetName)
	case OpDropConstraint:
		return hasTableOperation(down, OpAddConstraint, up.SchemaName, up.TableName, up.TargetName)
	case OpAlterColumnType, OpAlterNullability:
		return hasTableOperation(down, up.Kind, up.SchemaName, up.TableName, up.TargetName)
	case OpRenameColumn, OpRenameConstraint:
		for _, candidate := range down {
			if candidate.Kind == up.Kind && sameTable(candidate, up) &&
				strings.EqualFold(candidate.OldValue, up.NewValue) && strings.EqualFold(candidate.NewValue, up.OldValue) {
				return true
			}
		}
		return false
	default:
		return false
	}
}

func sameTable(left, right ParsedTableStatement) bool {
	return strings.EqualFold(left.SchemaName, right.SchemaName) && strings.EqualFold(left.TableName, right.TableName)
}

func hasTableOperation(items []ParsedTableStatement, kind, schema, table, target string) bool {
	for _, item := range items {
		if item.Kind == kind && strings.EqualFold(item.SchemaName, schema) && strings.EqualFold(item.TableName, table) &&
			(target == "" || strings.EqualFold(item.TargetName, target)) {
			return true
		}
	}
	return false
}

func hasInverseIndexOperation(up ParsedIndex, down []ParsedIndex, downTables []ParsedTableStatement) bool {
	for _, table := range downTables {
		if table.Kind == OpDropTable && strings.EqualFold(table.TableName, up.TableName) && up.IndexType != "DROP" {
			return true
		}
	}
	for _, candidate := range down {
		if up.IndexType == "DROP" && candidate.IndexType != "DROP" && strings.EqualFold(up.IndexName, candidate.IndexName) {
			return true
		}
		if up.IndexType != "DROP" && up.IndexType != "RENAME" && candidate.IndexType == "DROP" && strings.EqualFold(up.IndexName, candidate.IndexName) {
			return true
		}
		if up.IndexType == "RENAME" && candidate.IndexType == "RENAME" &&
			strings.EqualFold(up.IndexName, candidate.RawDefinition) && strings.EqualFold(up.RawDefinition, candidate.IndexName) {
			return true
		}
	}
	return false
}

func programmableKindsAreInverse(up, down string) bool {
	switch up {
	case OpProgDrop:
		return down == OpProgCreate || down == OpProgAlter
	case OpProgCreate:
		return down == OpProgDrop
	default:
		// CREATE OR ALTER is parsed as ALTER but is also the canonical create
		// form for a first version, whose inverse is DROP.
		return down == OpProgCreate || down == OpProgAlter || down == OpProgDrop
	}
}

func typeKindsAreInverse(up, down string) bool {
	switch up {
	case OpDropType:
		return down == OpCreateType
	case OpCreateType:
		return down == OpDropType
	default:
		return down == OpCreateType || down == OpAlterType || down == OpReplaceType
	}
}
