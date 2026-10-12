package engine

import "rza_migrate_db/internal/output"

// buildMigrationPlan is the orchestration boundary between reconstructed
// history, desired db_object state, and rendering. All inputs are derived from
// the two authoritative directories and remain in memory for the run.
func (e Engine) buildMigrationPlan(report *output.Report, history map[string]historyObject, existing map[string]struct{}, maxVersion string, desired []desiredObject, patches []patchSource) ([]plannedMigration, map[string]struct{}, string, error) {
	return e.planGenerationWithInputs(report, history, existing, maxVersion, desired, patches)
}
