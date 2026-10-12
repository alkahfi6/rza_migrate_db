package verify

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"rza_migrate_db/internal/fsx"
	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlparse"
)

// TableInlineDefaultOptions holds runtime context for inline DEFAULT validation.
type TableInlineDefaultOptions struct {
	RepoPath           func(string) string
	DbObjectDir        string
	ToRelPath          func(string) string
	ListFilesRecursive func(string) ([]string, error)
	ReadFile           func(string) ([]byte, error)
}

// ValidateTableInlineDefault checks db_object table snapshot DEFAULT determinism.
// Inline and table-level DEFAULT ... FOR forms are accepted, but must be named.
func ValidateTableInlineDefault(report *output.Report, opts TableInlineDefaultOptions) error {
	listFiles := opts.ListFilesRecursive
	if listFiles == nil {
		listFiles = fsx.ListFilesRecursive
	}

	files, err := fsx.ResolveDbObjectCategoryFiles(opts.RepoPath, listFiles, opts.DbObjectDir, "tables")
	if err != nil {
		return err
	}

	for _, file := range files {
		if !strings.EqualFold(filepath.Ext(file), ".sql") {
			continue
		}

		readFile := opts.ReadFile
		if readFile == nil {
			readFile = os.ReadFile
		}
		raw, err := readFile(file)
		if err != nil {
			return err
		}
		relPath := file
		if opts.ToRelPath != nil {
			relPath = opts.ToRelPath(file)
		}

		if hasInvalidInlineDefaultInCreateTable(string(raw)) {
			report.AddFail(output.Item{
				Rule:     "INLINE_DEFAULT_NOT_ALLOWED",
				Severity: 3,
				File:     relPath,
				Message:  "inline DEFAULT must be named with CONSTRAINT <name> in db_object table snapshot",
				Reason:   "INLINE_DEFAULT_NOT_ALLOWED",
				Action:   "FAIL",
			})
		}
		unnamedViolations := findUnnamedConstraintViolationsInCreateTable(string(raw))
		if len(unnamedViolations) > 0 {
			detailsLimit := len(unnamedViolations)
			if detailsLimit > 3 {
				detailsLimit = 3
			}
			details := make([]string, 0, detailsLimit)
			for _, violation := range unnamedViolations[:detailsLimit] {
				details = append(details, fmt.Sprintf("%s near %s", violation.Kind, violation.Preview))
			}
			message := "all constraints in CREATE TABLE must be named with CONSTRAINT <name>: " + strings.Join(details, "; ")
			if len(unnamedViolations) > detailsLimit {
				message += fmt.Sprintf(" (+%d more)", len(unnamedViolations)-detailsLimit)
			}
			report.AddFail(output.Item{
				Rule:       "UNNAMED_CONSTRAINT_NOT_ALLOWED",
				Severity:   3,
				File:       relPath,
				Message:    message,
				Suggestion: "use CONSTRAINT <name> for DEFAULT/CHECK/PRIMARY KEY/UNIQUE/FOREIGN KEY (including inline column constraints)",
				Reason:     "REQUIRES_MANUAL",
				Action:     "FAIL",
			})
		}
	}
	return nil
}

func hasInvalidInlineDefaultInCreateTable(sql string) bool {
	items := extractCreateTableItemsUpper(sql)
	if len(items) == 0 {
		return false
	}
	for _, item := range items {
		itemUpper := strings.ToUpper(strings.TrimSpace(item))
		if itemUpper == "" {
			continue
		}
		if !sqlparse.ContainsStandaloneWord(itemUpper, "DEFAULT") {
			continue
		}
		if keywordIsFirstTokenIdentifier(itemUpper, "DEFAULT") {
			continue
		}
		// Table-level default constraint is allowed in snapshot:
		// CONSTRAINT <name> DEFAULT (...) FOR <column>
		if strings.HasPrefix(itemUpper, "CONSTRAINT ") &&
			sqlparse.ContainsStandaloneWord(itemUpper, "FOR") {
			continue
		}

		// Inline column DEFAULT is allowed only if named:
		// <column> <type> CONSTRAINT <name> DEFAULT (...)
		defaultIdx := strings.Index(itemUpper, "DEFAULT")
		if defaultIdx < 0 {
			continue
		}
		prefix := itemUpper[:defaultIdx]
		if sqlparse.ContainsStandaloneWord(prefix, "CONSTRAINT") {
			continue
		}
		return true
	}
	return false
}

func hasTableLevelDefaultForConstraintInCreateTable(sql string) bool {
	items := extractCreateTableItemsUpper(sql)
	if len(items) == 0 {
		return false
	}
	for _, item := range items {
		itemUpper := strings.ToUpper(strings.TrimSpace(item))
		if itemUpper == "" {
			continue
		}
		if !strings.HasPrefix(itemUpper, "CONSTRAINT ") {
			continue
		}
		if !sqlparse.ContainsStandaloneWord(itemUpper, "DEFAULT") {
			continue
		}
		if sqlparse.ContainsStandaloneWord(itemUpper, "FOR") {
			return true
		}
	}
	return false
}

type unnamedConstraintViolation struct {
	Kind    string
	Preview string
}

func findUnnamedConstraintViolationsInCreateTable(sql string) []unnamedConstraintViolation {
	items := extractCreateTableItemsUpper(sql)
	if len(items) == 0 {
		return []unnamedConstraintViolation{}
	}

	violations := make([]unnamedConstraintViolation, 0)
	for _, raw := range items {
		item := strings.TrimSpace(raw)
		if item == "" {
			continue
		}
		itemUpper := strings.ToUpper(item)

		preview := compactTableItemPreview(itemUpper)
		if strings.HasPrefix(itemUpper, "CONSTRAINT ") {
			if !hasNamedConstraintDeclaration(itemUpper) {
				violations = append(violations, unnamedConstraintViolation{Kind: "CONSTRAINT", Preview: preview})
			}
			continue
		}

		switch {
		case startsWithKeywordToken(itemUpper, "PRIMARY KEY"):
			violations = append(violations, unnamedConstraintViolation{Kind: "PRIMARY KEY", Preview: preview})
			continue
		case startsWithKeywordToken(itemUpper, "UNIQUE"):
			violations = append(violations, unnamedConstraintViolation{Kind: "UNIQUE", Preview: preview})
			continue
		case startsWithKeywordToken(itemUpper, "CHECK"):
			violations = append(violations, unnamedConstraintViolation{Kind: "CHECK", Preview: preview})
			continue
		case startsWithKeywordToken(itemUpper, "FOREIGN KEY"):
			violations = append(violations, unnamedConstraintViolation{Kind: "FOREIGN KEY", Preview: preview})
			continue
		case startsWithKeywordToken(itemUpper, "DEFAULT"):
			violations = append(violations, unnamedConstraintViolation{Kind: "DEFAULT", Preview: preview})
			continue
		}

		if !keywordIsFirstTokenIdentifier(itemUpper, "DEFAULT") {
			if pos := firstStandaloneKeywordPos(itemUpper, "DEFAULT"); pos >= 0 && !hasConstraintBeforePos(itemUpper, pos) {
				violations = append(violations, unnamedConstraintViolation{Kind: "DEFAULT", Preview: preview})
				continue
			}
		}
		if !keywordIsFirstTokenIdentifier(itemUpper, "CHECK") {
			if pos := firstStandaloneKeywordPos(itemUpper, "CHECK"); pos >= 0 && !hasConstraintBeforePos(itemUpper, pos) {
				violations = append(violations, unnamedConstraintViolation{Kind: "CHECK", Preview: preview})
				continue
			}
		}
		if !keywordIsFirstTokenIdentifier(itemUpper, "PRIMARY") {
			if pos := firstPrimaryKeyPos(itemUpper); pos >= 0 && !hasConstraintBeforePos(itemUpper, pos) {
				violations = append(violations, unnamedConstraintViolation{Kind: "PRIMARY KEY", Preview: preview})
				continue
			}
		}
		if !keywordIsFirstTokenIdentifier(itemUpper, "UNIQUE") {
			if pos := firstStandaloneKeywordPos(itemUpper, "UNIQUE"); pos >= 0 && !hasConstraintBeforePos(itemUpper, pos) {
				violations = append(violations, unnamedConstraintViolation{Kind: "UNIQUE", Preview: preview})
				continue
			}
		}
		if !keywordIsFirstTokenIdentifier(itemUpper, "REFERENCES") {
			if pos := firstStandaloneKeywordPos(itemUpper, "REFERENCES"); pos >= 0 && !hasConstraintBeforePos(itemUpper, pos) {
				violations = append(violations, unnamedConstraintViolation{Kind: "FOREIGN KEY", Preview: preview})
				continue
			}
		}
	}
	return violations
}

func extractCreateTableItemsUpper(sql string) []string {
	scanned := strings.ToUpper(sqlparse.StripSQLStringsAndComments(strings.ReplaceAll(strings.ReplaceAll(sql, "\r\n", "\n"), "\r", "\n")))
	createPos := strings.Index(scanned, "CREATE TABLE")
	if createPos < 0 {
		return []string{}
	}

	openIdx := strings.Index(scanned[createPos:], "(")
	if openIdx < 0 {
		return []string{}
	}
	openIdx += createPos

	depth := 0
	closeIdx := -1
	for i := openIdx; i < len(scanned); i++ {
		switch scanned[i] {
		case '(':
			depth++
		case ')':
			depth--
			if depth == 0 {
				closeIdx = i
				i = len(scanned)
			}
		}
	}
	if closeIdx <= openIdx {
		return []string{}
	}

	body := scanned[openIdx+1 : closeIdx]
	return splitTopLevelByComma(body)
}

func startsWithKeywordToken(lineUpper, keyword string) bool {
	if !strings.HasPrefix(lineUpper, keyword) {
		return false
	}
	if len(lineUpper) == len(keyword) {
		return true
	}
	next := rune(lineUpper[len(keyword)])
	return !isWordChar(next)
}

func firstStandaloneKeywordPos(lineUpper, keyword string) int {
	positions := findAllKeywordPositions(lineUpper, keyword)
	if len(positions) == 0 {
		return -1
	}
	return positions[0]
}

func firstPrimaryKeyPos(lineUpper string) int {
	positions := findAllKeywordPositions(lineUpper, "PRIMARY")
	for _, pos := range positions {
		rest := strings.TrimSpace(lineUpper[pos+len("PRIMARY"):])
		if strings.EqualFold(firstSQLToken(rest), "KEY") {
			return pos
		}
	}
	return -1
}

func hasConstraintBeforePos(lineUpper string, pos int) bool {
	if pos <= 0 {
		return false
	}
	prefix := lineUpper[:pos]
	positions := findAllKeywordPositions(prefix, "CONSTRAINT")
	if len(positions) == 0 {
		return false
	}
	rest := strings.TrimSpace(prefix[positions[len(positions)-1]+len("CONSTRAINT"):])
	name := firstSQLToken(rest)
	if name == "" {
		return false
	}
	rest = strings.TrimSpace(rest[len(name):])
	for _, keyword := range []string{"DEFAULT", "CHECK", "UNIQUE", "REFERENCES", "PRIMARY", "FOREIGN"} {
		if sqlparse.ContainsStandaloneWord(rest, keyword) {
			return false
		}
	}
	return true
}

func hasNamedConstraintDeclaration(itemUpper string) bool {
	if !strings.HasPrefix(itemUpper, "CONSTRAINT ") {
		return false
	}
	rest := strings.TrimSpace(itemUpper[len("CONSTRAINT "):])
	name := firstSQLToken(rest)
	if name == "" {
		return false
	}
	switch name {
	case "PRIMARY", "UNIQUE", "CHECK", "FOREIGN", "DEFAULT", "REFERENCES":
		return false
	}
	return true
}

func compactTableItemPreview(itemUpper string) string {
	compact := strings.Join(strings.Fields(strings.TrimSpace(itemUpper)), " ")
	if len(compact) <= 90 {
		return compact
	}
	return compact[:87] + "..."
}

func keywordIsFirstTokenIdentifier(lineUpper, keyword string) bool {
	token := strings.TrimSpace(firstSQLToken(lineUpper))
	if token == "" {
		return false
	}
	token = strings.Trim(token, "[]\"")
	return strings.EqualFold(token, keyword)
}

func splitTopLevelByComma(input string) []string {
	parts := make([]string, 0)
	start := 0
	depth := 0
	for i, ch := range input {
		switch ch {
		case '(':
			depth++
		case ')':
			if depth > 0 {
				depth--
			}
		case ',':
			if depth == 0 {
				parts = append(parts, input[start:i])
				start = i + 1
			}
		}
	}
	if start <= len(input) {
		parts = append(parts, input[start:])
	}
	return parts
}
