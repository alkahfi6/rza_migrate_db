package migreport

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// WriteReports writes migration datasets and JSON into the audit directory.
func WriteReports(outputDir string, res *RunResult) error {
	if err := os.MkdirAll(filepath.Join(outputDir, "datasets"), 0755); err != nil {
		return fmt.Errorf("failed to create output directory %s: %w", outputDir, err)
	}

	// 1. report_summary.csv
	if err := writeReportSummary(filepath.Join(outputDir, "datasets", "report_summary.csv"), res.Summary); err != nil {
		return err
	}

	// 2. migration_summary.csv
	if err := writeMigrationSummary(filepath.Join(outputDir, "datasets", "migration_summary.csv"), res.Migrations); err != nil {
		return err
	}

	// 3. object_history.csv
	if err := writeObjectHistory(filepath.Join(outputDir, "datasets", "object_history.csv"), res.ObjectHistory); err != nil {
		return err
	}

	// 4. table_changes.csv
	if err := writeTableChanges(filepath.Join(outputDir, "datasets", "table_changes.csv"), res.TableChanges); err != nil {
		return err
	}

	// 5. index_changes.csv
	if err := writeIndexChanges(filepath.Join(outputDir, "datasets", "index_changes.csv"), res.IndexChanges); err != nil {
		return err
	}

	// 6. programmable_changes.csv
	if err := writeProgChanges(filepath.Join(outputDir, "datasets", "programmable_changes.csv"), res.ProgChanges); err != nil {
		return err
	}

	// 7. type_changes.csv
	if err := writeTypeChanges(filepath.Join(outputDir, "datasets", "type_changes.csv"), res.TypeChanges); err != nil {
		return err
	}

	// 8. report.json
	if err := writeJSONReport(filepath.Join(outputDir, "report.json"), res.JSONReport); err != nil {
		return err
	}

	return nil
}

func writeReportSummary(path string, s ReportSummary) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{
		"run_at", "repository", "migrations_dir", "total_files", "total_pairs",
		"excluded_pairs", "unpaired_files", "total_objects", "total_versions", "total_operations",
		"tables", "indexes", "procedures", "views", "functions", "types",
		"warnings", "errors", "status",
	}
	if err := w.Write(header); err != nil {
		return err
	}

	row := []string{
		s.RunAt.Format(time.RFC3339),
		s.Repository,
		s.MigrationsDir,
		strconv.Itoa(s.TotalFiles),
		strconv.Itoa(s.TotalPairs),
		strconv.Itoa(s.ExcludedPairs),
		strconv.Itoa(s.UnpairedFiles),
		strconv.Itoa(s.TotalObjects),
		strconv.Itoa(s.TotalVersions),
		strconv.Itoa(s.TotalOperations),
		strconv.Itoa(s.Tables),
		strconv.Itoa(s.Indexes),
		strconv.Itoa(s.Procedures),
		strconv.Itoa(s.Views),
		strconv.Itoa(s.Functions),
		strconv.Itoa(s.Types),
		strconv.Itoa(s.Warnings),
		strconv.Itoa(s.Errors),
		s.Status,
	}
	return w.Write(row)
}

func writeMigrationSummary(path string, items []MigrationSummary) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{
		"migration_id", "category", "schema_name", "object_name", "object_version",
		"operation_count", "operation_kinds", "up_file", "down_file",
		"up_exists", "down_exists", "round_trip_status", "status",
	}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, it := range items {
		// Clean up operation_kinds if already quoted
		opKinds := it.OperationKinds
		if len(opKinds) >= 2 && opKinds[0] == '"' && opKinds[len(opKinds)-1] == '"' {
			opKinds = opKinds[1 : len(opKinds)-1]
		}
		row := []string{
			it.MigrationID,
			it.Category,
			it.SchemaName,
			it.ObjectName,
			it.ObjectVersion,
			strconv.Itoa(it.OperationCount),
			opKinds,
			it.UpFile,
			it.DownFile,
			it.UpExists,
			it.DownExists,
			it.RoundTripStatus,
			it.Status,
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func writeObjectHistory(path string, items []ObjectHistory) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{
		"category", "schema_name", "object_name", "object_version",
		"first_migration_id", "last_migration_id",
		"previous_version", "change_count",
		"first_up_file", "last_up_file", "history_status",
	}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, it := range items {
		row := []string{
			it.Category,
			it.SchemaName,
			it.ObjectName,
			it.ObjectVersion,
			it.FirstMigrationID,
			it.LastMigrationID,
			it.PreviousVersion,
			strconv.Itoa(it.ChangeCount),
			it.FirstUpFile,
			it.LastUpFile,
			it.HistoryStatus,
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func writeTableChanges(path string, items []TableChange) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{
		"migration_id", "schema_name", "table_name", "object_version",
		"operation_no", "operation_kind", "target_name",
		"old_value", "new_value", "dependency_name",
		"up_file", "down_file", "reversible", "confidence", "notes",
	}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, it := range items {
		row := []string{
			it.MigrationID,
			it.SchemaName,
			it.TableName,
			it.ObjectVersion,
			strconv.Itoa(it.OperationNo),
			it.OperationKind,
			it.TargetName,
			it.OldValue,
			it.NewValue,
			it.DependencyName,
			it.UpFile,
			it.DownFile,
			it.Reversible,
			it.Confidence,
			it.Notes,
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func writeIndexChanges(path string, items []IndexChange) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{
		"migration_id", "schema_name", "table_name", "index_name",
		"object_version", "operation_no", "operation_kind",
		"index_type", "old_columns", "new_columns",
		"old_include_columns", "new_include_columns",
		"old_filter", "new_filter",
		"old_unique", "new_unique",
		"old_clustered", "new_clustered",
		"up_file", "down_file", "reversible", "confidence", "notes",
	}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, it := range items {
		row := []string{
			it.MigrationID,
			it.SchemaName,
			it.TableName,
			it.IndexName,
			it.ObjectVersion,
			strconv.Itoa(it.OperationNo),
			it.OperationKind,
			it.IndexType,
			it.OldColumns,
			it.NewColumns,
			it.OldIncludeColumns,
			it.NewIncludeColumns,
			it.OldFilter,
			it.NewFilter,
			it.OldUnique,
			it.NewUnique,
			it.OldClustered,
			it.NewClustered,
			it.UpFile,
			it.DownFile,
			it.Reversible,
			it.Confidence,
			it.Notes,
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func writeProgChanges(path string, items []ProgrammableChange) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{
		"migration_id", "category", "schema_name", "object_name",
		"object_version", "operation_kind", "definition_hash_before",
		"definition_hash_after", "changed_sections", "up_file", "down_file",
		"reversible", "confidence", "notes",
	}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, it := range items {
		row := []string{
			it.MigrationID,
			it.Category,
			it.SchemaName,
			it.ObjectName,
			it.ObjectVersion,
			it.OperationKind,
			it.DefinitionHashBefore,
			it.DefinitionHashAfter,
			it.ChangedSections,
			it.UpFile,
			it.DownFile,
			it.Reversible,
			it.Confidence,
			it.Notes,
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func writeTypeChanges(path string, items []TypeChange) error {
	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	w := csv.NewWriter(f)
	defer w.Flush()

	header := []string{
		"migration_id", "schema_name", "type_name", "object_version",
		"operation_kind", "old_definition", "new_definition",
		"up_file", "down_file", "reversible", "confidence", "notes",
	}
	if err := w.Write(header); err != nil {
		return err
	}

	for _, it := range items {
		row := []string{
			it.MigrationID,
			it.SchemaName,
			it.TypeName,
			it.ObjectVersion,
			it.OperationKind,
			it.OldDefinition,
			it.NewDefinition,
			it.UpFile,
			it.DownFile,
			it.Reversible,
			it.Confidence,
			it.Notes,
		}
		if err := w.Write(row); err != nil {
			return err
		}
	}
	return nil
}

func writeJSONReport(path string, r JSONReport) error {
	bytes, err := json.MarshalIndent(r, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(path, append(bytes, '\n'), 0644)
}

// PrintConsoleSummary prints the console banner and summary lines.
func PrintConsoleSummary(res *RunResult, outputDir, docsDir string) {
	fmt.Println("[rza_migrate_report] starting")
	fmt.Println("[source] migrations + db_object (read-only)")
	fmt.Println("[database] access=disabled")
	fmt.Println("[sql-execution] disabled")
	fmt.Printf("[scan] files=%d analyzed_pairs=%d excluded_pairs=%d\n", res.Summary.TotalFiles, res.Summary.TotalPairs, res.Summary.ExcludedPairs)
	fmt.Printf("[objects] total=%d operations=%d\n", res.Summary.TotalObjects, res.Summary.TotalOperations)

	reportFiles := auditArtifactFiles()
	cwd, _ := os.Getwd()
	for _, rf := range reportFiles {
		full := filepath.Join(outputDir, rf)
		displayPath := full
		if cwd != "" {
			if rel, err := filepath.Rel(cwd, full); err == nil && !strings.HasPrefix(rel, "..") {
				displayPath = rel
			}
		}
		fmt.Printf("[report] %s\n", filepath.ToSlash(displayPath))
	}
	for _, df := range desiredStateArtifactFiles(res.DesiredStateDocs) {
		full := filepath.Join(docsDir, df)
		displayPath := full
		if cwd != "" {
			if rel, err := filepath.Rel(cwd, full); err == nil && !strings.HasPrefix(rel, "..") {
				displayPath = rel
			}
		}
		fmt.Printf("[report] %s\n", filepath.ToSlash(displayPath))
	}

	fmt.Printf("[result] status=%s warnings=%d errors=%d\n", res.Summary.Status, res.Summary.Warnings, res.Summary.Errors)
}
