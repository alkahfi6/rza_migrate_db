package sqlfmt

import "strings"

// NormalizeSQLText applies repository SQL text normalization policy.
func NormalizeSQLText(input string, trimTrailingSpaces bool, ensureFinalNewline bool) string {
	normalized := strings.ReplaceAll(input, "\r\n", "\n")
	normalized = strings.ReplaceAll(normalized, "\r", "\n")
	normalized = strings.TrimPrefix(normalized, "\uFEFF")
	// SQL copied from spreadsheets or legacy exports may use non-breaking spaces
	// between identifiers and data types. T-SQL tokenization does not reliably
	// treat those bytes as separators, so canonicalize them before parsing.
	normalized = strings.ReplaceAll(normalized, "\u00C2\u00A0", " ")
	normalized = strings.ReplaceAll(normalized, "\u00A0", " ")

	lines := strings.Split(normalized, "\n")
	if trimTrailingSpaces {
		for i := range lines {
			lines[i] = strings.TrimRight(lines[i], " \t")
		}
	}

	out := strings.Join(lines, "\n")
	if ensureFinalNewline && out != "" && !strings.HasSuffix(out, "\n") {
		out += "\n"
	}
	return out
}
