package getobject

import (
	"strings"

	"rza_migrate_db/internal/tsqlast"
)

// ModuleFormatOptions configures formatter behavior for module SQL text.
type ModuleFormatOptions struct {
	StripComments bool
	Reindent      bool
	IndentSize    int
	Backend       string
}

// FormatModuleSQLWithOptions formats module SQL using the same formatter stack
// used by rza_get_object_db (ANTLR + lite fallback).
func FormatModuleSQLWithOptions(sqlText string, opts ModuleFormatOptions) string {
	indentSize := opts.IndentSize
	if indentSize <= 0 {
		indentSize = 2
	}

	backend := strings.TrimSpace(opts.Backend)
	if backend == "" {
		backend = "antlr"
	}

	out, _, err := tsqlast.FormatModuleSQL(sqlText, tsqlast.FormatConfig{
		StripComments:  opts.StripComments,
		Reindent:       opts.Reindent,
		IndentSize:     indentSize,
		Backend:        tsqlast.Backend(backend),
		FallbackToLite: true,
	})
	if err != nil {
		// defensive fallback to preserve formatter progress
		out, _, _ = tsqlast.FormatModuleSQL(sqlText, tsqlast.FormatConfig{
			StripComments:  opts.StripComments,
			Reindent:       opts.Reindent,
			IndentSize:     indentSize,
			Backend:        tsqlast.BackendLite,
			FallbackToLite: true,
		})
	}
	return normalizeSQL(out)
}

func formatModuleDefinition(sqlText string, cfg AppConfig) string {
	return FormatModuleSQLWithOptions(sqlText, ModuleFormatOptions{
		StripComments: cfg.StripModuleComments,
		Reindent:      cfg.ReindentModuleSQL,
		IndentSize:    cfg.IndentSize,
		Backend:       cfg.ModuleFormatterBackend,
	})
}

func stripSQLComments(input string) string {
	return tsqlast.StripSQLComments(input)
}

func reindentTSQL(input string, indentSize int) string {
	return tsqlast.ReindentTSQL(input, indentSize)
}

func expandTrailingControlBegin(line string) []string {
	return tsqlast.ExpandTrailingControlBegin(line)
}

func splitEndTrailingSuffix(line string) []string {
	return tsqlast.SplitEndTrailingSuffix(line)
}
