package migreport

import "time"

// Status constants for report summary and migration summary.
const (
	StatusOk      = "OK"
	StatusWarning = "WARNING"
	StatusFailed  = "FAILED"
)

// Round trip statuses.
const (
	RoundTripPass      = "PASS"
	RoundTripFail      = "FAIL"
	RoundTripNotTested = "NOT_TESTED"
	RoundTripUnpaired  = "UNPAIRED"
)

// History status constants.
const (
	HistoryInitial      = "INITIAL"
	HistoryUpdated      = "UPDATED"
	HistoryDeleted      = "DELETED"
	HistoryReintroduced = "REINTRODUCED"
	HistoryUnknown      = "UNKNOWN"
)

// Confidence levels.
const (
	ConfidenceHigh   = "HIGH"
	ConfidenceMedium = "MEDIUM"
	ConfidenceLow    = "LOW"
)

// Reversibility constants.
const (
	ReversibleYes = "YES"
	ReversibleNo  = "NO"
)

// Table operation kinds.
const (
	OpCreateTable      = "CREATE_TABLE"
	OpDropTable        = "DROP_TABLE"
	OpAddColumn        = "ADD_COLUMN"
	OpDropColumn       = "DROP_COLUMN"
	OpRenameColumn     = "RENAME_COLUMN"
	OpAlterColumnType  = "ALTER_COLUMN_TYPE"
	OpAlterNullability = "ALTER_NULLABILITY"
	OpAlterDefault     = "ALTER_DEFAULT"
	OpAddConstraint    = "ADD_CONSTRAINT"
	OpDropConstraint   = "DROP_CONSTRAINT"
	OpRenameConstraint = "RENAME_CONSTRAINT"
	OpAlterConstraint  = "ALTER_CONSTRAINT"
)

// Index operation kinds.
const (
	OpAddIndex    = "ADD_INDEX"
	OpDropIndex   = "DROP_INDEX"
	OpAlterIndex  = "ALTER_INDEX"
	OpRenameIndex = "RENAME_INDEX"
)

// Programmable operation kinds.
const (
	OpProgCreate            = "CREATE"
	OpProgAlter             = "ALTER"
	OpProgDrop              = "DROP"
	OpProgReplaceDefinition = "REPLACE_DEFINITION"
)

// Programmable changed sections.
const (
	SectionParameters = "PARAMETERS"
	SectionDeclare    = "DECLARE"
	SectionQuery      = "QUERY"
	SectionJoin       = "JOIN"
	SectionWhere      = "WHERE"
	SectionDML        = "DML"
	SectionReturnType = "RETURN_TYPE"
	SectionBody       = "BODY"
	SectionUnknown    = "UNKNOWN"
)

// Type operation kinds.
const (
	OpCreateType  = "CREATE_TYPE"
	OpDropType    = "DROP_TYPE"
	OpReplaceType = "REPLACE_TYPE"
	OpAlterType   = "ALTER_TYPE"
)

// Special notes.
const (
	NoteManualReviewRequired = "MANUAL_REVIEW_REQUIRED"
	NoteMixedObjectMigration = "MIXED_OBJECT_MIGRATION"
)

// ReportSummary corresponds to report_summary.csv.
type ReportSummary struct {
	RunAt           time.Time `json:"run_at"`
	Repository      string    `json:"repository"`
	MigrationsDir   string    `json:"migrations_dir"`
	TotalFiles      int       `json:"total_files"`
	TotalPairs      int       `json:"total_pairs"`
	ExcludedPairs   int       `json:"excluded_pairs"`
	UnpairedFiles   int       `json:"unpaired_files"`
	TotalObjects    int       `json:"total_objects"`
	TotalVersions   int       `json:"total_versions"`
	TotalOperations int       `json:"total_operations"`
	Tables          int       `json:"tables"`
	Indexes         int       `json:"indexes"`
	Procedures      int       `json:"procedures"`
	Views           int       `json:"views"`
	Functions       int       `json:"functions"`
	Types           int       `json:"types"`
	Warnings        int       `json:"warnings"`
	Errors          int       `json:"errors"`
	Status          string    `json:"status"`
}

// MigrationSummary corresponds to migration_summary.csv.
type MigrationSummary struct {
	MigrationID     string `json:"migration_id"`
	Category        string `json:"category"`
	SchemaName      string `json:"schema_name"`
	ObjectName      string `json:"object_name"`
	ObjectVersion   string `json:"object_version"`
	OperationCount  int    `json:"operation_count"`
	OperationKinds  string `json:"operation_kinds"`
	UpFile          string `json:"up_file"`
	DownFile        string `json:"down_file"`
	UpExists        string `json:"up_exists"`
	DownExists      string `json:"down_exists"`
	RoundTripStatus string `json:"round_trip_status"`
	Status          string `json:"status"`
}

// ObjectHistory corresponds to object_history.csv.
type ObjectHistory struct {
	Category         string `json:"category"`
	SchemaName       string `json:"schema_name"`
	ObjectName       string `json:"object_name"`
	ObjectVersion    string `json:"object_version"`
	FirstMigrationID string `json:"first_migration_id"`
	LastMigrationID  string `json:"last_migration_id"`
	PreviousVersion  string `json:"previous_version"`
	ChangeCount      int    `json:"change_count"`
	FirstUpFile      string `json:"first_up_file"`
	LastUpFile       string `json:"last_up_file"`
	HistoryStatus    string `json:"history_status"`
}

// TableChange corresponds to table_changes.csv.
type TableChange struct {
	MigrationID    string `json:"migration_id"`
	SchemaName     string `json:"schema_name"`
	TableName      string `json:"table_name"`
	ObjectVersion  string `json:"object_version"`
	OperationNo    int    `json:"operation_no"`
	OperationKind  string `json:"operation_kind"`
	TargetName     string `json:"target_name"`
	OldValue       string `json:"old_value"`
	NewValue       string `json:"new_value"`
	DependencyName string `json:"dependency_name"`
	UpFile         string `json:"up_file"`
	DownFile       string `json:"down_file"`
	Reversible     string `json:"reversible"`
	Confidence     string `json:"confidence"`
	Notes          string `json:"notes"`
}

// IndexChange corresponds to index_changes.csv.
type IndexChange struct {
	MigrationID       string `json:"migration_id"`
	SchemaName        string `json:"schema_name"`
	TableName         string `json:"table_name"`
	IndexName         string `json:"index_name"`
	ObjectVersion     string `json:"object_version"`
	OperationNo       int    `json:"operation_no"`
	OperationKind     string `json:"operation_kind"`
	IndexType         string `json:"index_type"`
	OldColumns        string `json:"old_columns"`
	NewColumns        string `json:"new_columns"`
	OldIncludeColumns string `json:"old_include_columns"`
	NewIncludeColumns string `json:"new_include_columns"`
	OldFilter         string `json:"old_filter"`
	NewFilter         string `json:"new_filter"`
	OldUnique         string `json:"old_unique"`
	NewUnique         string `json:"new_unique"`
	OldClustered      string `json:"old_clustered"`
	NewClustered      string `json:"new_clustered"`
	UpFile            string `json:"up_file"`
	DownFile          string `json:"down_file"`
	Reversible        string `json:"reversible"`
	Confidence        string `json:"confidence"`
	Notes             string `json:"notes"`
}

// ProgrammableChange corresponds to programmable_changes.csv.
type ProgrammableChange struct {
	MigrationID          string `json:"migration_id"`
	Category             string `json:"category"`
	SchemaName           string `json:"schema_name"`
	ObjectName           string `json:"object_name"`
	ObjectVersion        string `json:"object_version"`
	OperationKind        string `json:"operation_kind"`
	DefinitionHashBefore string `json:"definition_hash_before"`
	DefinitionHashAfter  string `json:"definition_hash_after"`
	ChangedSections      string `json:"changed_sections"`
	UpFile               string `json:"up_file"`
	DownFile             string `json:"down_file"`
	Reversible           string `json:"reversible"`
	Confidence           string `json:"confidence"`
	Notes                string `json:"notes"`
}

// TypeChange corresponds to type_changes.csv.
type TypeChange struct {
	MigrationID   string `json:"migration_id"`
	SchemaName    string `json:"schema_name"`
	TypeName      string `json:"type_name"`
	ObjectVersion string `json:"object_version"`
	OperationKind string `json:"operation_kind"`
	OldDefinition string `json:"old_definition"`
	NewDefinition string `json:"new_definition"`
	UpFile        string `json:"up_file"`
	DownFile      string `json:"down_file"`
	Reversible    string `json:"reversible"`
	Confidence    string `json:"confidence"`
	Notes         string `json:"notes"`
}

// JSONReport corresponds to report.json.
type JSONReport struct {
	Schema                 string           `json:"schema"`
	GeneratedAt            string           `json:"generated_at"`
	Repository             string           `json:"repository"`
	Source                 string           `json:"source"`
	DatabaseAccess         bool             `json:"database_access"`
	SQLExecution           bool             `json:"sql_execution"`
	FilesScanned           int              `json:"files_scanned"`
	MigrationPairs         int              `json:"migration_pairs"`
	ExcludedPairs          int              `json:"excluded_migration_pairs"`
	Objects                int              `json:"objects"`
	Operations             int              `json:"operations"`
	OutputDirectory        string           `json:"output_directory"`
	DocumentationDirectory string           `json:"documentation_directory"`
	DocumentationIndex     string           `json:"documentation_index"`
	Reports                []string         `json:"reports"`
	Artifacts              []ReportArtifact `json:"artifacts"`
	Warnings               []string         `json:"warnings"`
	Errors                 []string         `json:"errors"`
}

// ReportArtifact identifies one concrete generated file and its output root.
type ReportArtifact struct {
	Kind   string `json:"kind"`
	Root   string `json:"root"`
	Path   string `json:"path"`
	SHA256 string `json:"sha256,omitempty"`
}

// RunResult collects all output datasets before writing.
type RunResult struct {
	Summary          ReportSummary
	Migrations       []MigrationSummary
	ObjectHistory    []ObjectHistory
	TableChanges     []TableChange
	IndexChanges     []IndexChange
	ProgChanges      []ProgrammableChange
	TypeChanges      []TypeChange
	JSONReport       JSONReport
	Warnings         []string
	Errors           []string
	DesiredStateDocs DesiredStateDocs
}
