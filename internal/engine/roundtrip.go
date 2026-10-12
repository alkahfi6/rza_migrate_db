package engine

import (
	"fmt"
	"strings"
)

// validatePlannedRoundTrip validates the engine's own semantic reconstruction
// before files are emitted. It is in-memory only: db_object remains desired
// truth and migrations remain historical truth.
func validatePlannedRoundTrip(history map[string]historyObject, desired []desiredObject, ops []plannedMigration) error {
	desiredByKey := make(map[string]desiredObject, len(desired))
	for _, item := range desired {
		if isEngineInfrastructureObject(item.Category, item.Object) {
			continue
		}
		desiredByKey[objectKey(item.Category, item.Object)] = item
	}
	state := cloneHistoryStateMap(history)
	for _, op := range ops {
		if op.Category == "patch" {
			continue
		}
		key := objectKey(op.Category, op.Object)
		current, exists := state[key]
		meta := migrationMeta{Category: op.Category, Object: op.Object}
		next, err := semanticFromMigrationWithState(meta, op.UpSemanticSQL, current.SemanticSQL, exists)
		if err != nil {
			return fmt.Errorf("round-trip apply failed for %s: %w", key, err)
		}
		want, hasDesired := desiredByKey[key]
		if hasDesired && hashSemanticByCategory(op.Category, next) != want.NormalizedHash {
			return fmt.Errorf("round-trip mismatch for %s: generated up state differs from db_object desired state", key)
		}
		if !hasDesired && strings.TrimSpace(next) != "" {
			return fmt.Errorf("round-trip mismatch for %s: generated removal did not produce an empty object state", key)
		}
		state[key] = historyObject{
			Category: op.Category, Object: op.Object, SemanticSQL: next,
			NormalizedHash: hashSemanticByCategory(op.Category, next),
			Deleted:        !hasDesired && strings.TrimSpace(next) == "",
		}
	}

	for key, want := range desiredByKey {
		got, exists := state[key]
		if !exists || got.Deleted || hashSemanticByCategory(want.Category, got.SemanticSQL) != want.NormalizedHash {
			return fmt.Errorf("round-trip coverage failed for %s: final state does not match db_object", key)
		}
	}
	for key, got := range state {
		if got.Category == "patch" || isEngineInfrastructureObject(got.Category, got.Object) {
			continue
		}
		if _, exists := desiredByKey[key]; !exists && !got.Deleted && strings.TrimSpace(got.SemanticSQL) != "" {
			return fmt.Errorf("round-trip coverage failed for %s: history-only object has no removal operation", key)
		}
	}
	return nil
}
