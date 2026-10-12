package engine

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"rza_migrate_db/internal/config"
	"rza_migrate_db/internal/output"
	"rza_migrate_db/internal/sqlparse"
)

type namedTableConstraint struct {
	Name      string
	Signature string
}

type constraintNameOccurrence struct {
	Name       string
	Start, End int
}

type constraintNameSourceWrite struct {
	Abs     string
	Rel     string
	Content []byte
	Count   int
}

// planDesiredTableConstraintNames computes trivia-preserving source rewrites
// without touching disk. Both interactive tools and the CI checker use this
// exact plan, preventing their auto-fix decisions from drifting apart.
func (e Engine) planDesiredTableConstraintNames(history map[string]historyObject, desired []desiredObject) ([]constraintNameSourceWrite, error) {
	writes := make([]constraintNameSourceWrite, 0)
	for _, item := range desired {
		if item.Category != "tables" {
			continue
		}
		previous, ok := history[objectKey(item.Category, item.Object)]
		if !ok || previous.Deleted || strings.TrimSpace(previous.SemanticSQL) == "" {
			continue
		}

		mapping, err := equivalentConstraintNameMapping(previous.SemanticSQL, item.SemanticSQL)
		if err != nil {
			return nil, fmt.Errorf("compare constraint names for %s: %w", item.Object, err)
		}
		if len(mapping) == 0 {
			continue
		}

		abs := e.cfg.RepoPath(filepath.FromSlash(item.SourceFile))
		raw, err := os.ReadFile(abs)
		if err != nil {
			return nil, err
		}
		rewritten, changed, err := replaceNamedConstraintsInTableSource(string(raw), item.Object, mapping)
		if err != nil {
			return nil, fmt.Errorf("reconcile constraint names in %s: %w", item.SourceFile, err)
		}
		if changed != len(mapping) {
			return nil, fmt.Errorf("reconcile constraint names in %s: changed %d of %d expected names", item.SourceFile, changed, len(mapping))
		}

		// Reparse before staging the write. This catches malformed replacements
		// and proves that the resulting constraint signatures match history.
		semantic, err := semanticFromSource(migrationMeta{Category: "tables", Object: item.Object}, rewritten)
		if err != nil {
			return nil, fmt.Errorf("validate reconciled source %s: %w", item.SourceFile, err)
		}
		remaining, err := equivalentConstraintNameMapping(previous.SemanticSQL, semantic)
		if err != nil {
			return nil, fmt.Errorf("validate reconciled constraints for %s: %w", item.Object, err)
		}
		if len(remaining) != 0 {
			return nil, fmt.Errorf("validate reconciled constraints for %s: %d equivalent name drift(s) remain", item.Object, len(remaining))
		}

		writes = append(writes, constraintNameSourceWrite{
			Abs: abs, Rel: item.SourceFile, Content: []byte(rewritten), Count: changed,
		})
	}

	sort.Slice(writes, func(i, j int) bool { return writes[i].Rel < writes[j].Rel })
	return writes, nil
}

func applyConstraintNameWrites(writes []constraintNameSourceWrite) (int, []string, error) {
	total := 0
	files := make([]string, 0, len(writes))
	for _, write := range writes {
		if err := atomicWriteSQLFile(write.Abs, write.Content); err != nil {
			return 0, nil, err
		}
		total += write.Count
		files = append(files, write.Rel)
	}
	return total, files, nil
}

// reconcileDesiredTableConstraintNames makes db_object retain the constraint
// names already established by the latest migration state. A rename is safe
// only when the complete constraint definition is otherwise identical.
func (e Engine) reconcileDesiredTableConstraintNames(history map[string]historyObject, desired []desiredObject) (int, []string, error) {
	writes, err := e.planDesiredTableConstraintNames(history, desired)
	if err != nil {
		return 0, nil, err
	}
	return applyConstraintNameWrites(writes)
}

// runConstraintNameFix gives rza_migrate_check the same history-aware repair
// used by rza_migrate_db generation. EnableFmtCheck turns it into a read-only
// plan so MR pipelines fail until the interactive checker has applied it.
func (e Engine) runConstraintNameFix(report *output.Report) error {
	// With migration validation explicitly disabled there is no trusted history
	// contract to reconcile against. Keep this mode usable for object-only QA.
	if e.cfg.MigrationValidation == config.MigrationValidationOff {
		report.Notes["constraint_name_fix_count"] = 0
		report.Notes["constraint_name_fix_files"] = []string{}
		return nil
	}
	history, _, _, err := e.reconstructHistory(report)
	if err != nil || report.Fail.Count > 0 {
		return err
	}
	desired, err := e.collectDesiredObjects(report)
	if err != nil || report.Fail.Count > 0 {
		return err
	}
	writes, err := e.planDesiredTableConstraintNames(history, desired)
	if err != nil {
		return err
	}
	total := 0
	files := make([]string, 0, len(writes))
	for _, write := range writes {
		total += write.Count
		files = append(files, write.Rel)
	}
	report.Notes["constraint_name_fix_count"] = total
	report.Notes["constraint_name_fix_files"] = files
	if len(writes) == 0 {
		return nil
	}
	if e.cfg.EnableFmtCheck {
		report.AddFail(output.Item{
			Rule:       "CONSTRAINT_NAME_CHECK_CHANGES_REQUIRED",
			Severity:   2,
			File:       e.cfg.DbObjectDir,
			Message:    fmt.Sprintf("%d equivalent constraint name(s) across %d file(s) must be restored from migration history; no source files were written", total, len(files)),
			Suggestion: "run rza_migrate_check without -check (or double-click it) to apply the safe repair",
			Reason:     "REQUIRES_MANUAL",
			Action:     "FAIL",
		})
		return nil
	}
	applied, appliedFiles, err := applyConstraintNameWrites(writes)
	if err != nil {
		return err
	}
	e.resetRepositorySnapshot()
	report.AddPass(output.Item{
		Rule:     "CONSTRAINT_NAMES_RECONCILED",
		Severity: 0,
		File:     e.cfg.DbObjectDir,
		Message:  fmt.Sprintf("restored %d equivalent constraint name(s) from latest migration history across %d file(s)", applied, len(appliedFiles)),
		Reason:   "ALTERED",
		Action:   "ok",
	})
	return nil
}

func equivalentConstraintNameMapping(previousSQL, desiredSQL string) (map[string]string, error) {
	previous, err := namedTableConstraintsBySignature(previousSQL)
	if err != nil {
		return nil, err
	}
	next, err := namedTableConstraintsBySignature(desiredSQL)
	if err != nil {
		return nil, err
	}

	desiredNames := make(map[string][]string)
	for _, constraints := range next {
		for _, constraint := range constraints {
			key := strings.ToLower(constraint.Name)
			desiredNames[key] = append(desiredNames[key], constraint.Name)
		}
	}

	mapping := make(map[string]string)
	for signature, previousMatches := range previous {
		nextMatches := next[signature]
		// Identical duplicate definitions are ambiguous. Never guess which name
		// belongs to which constraint.
		if len(previousMatches) != 1 || len(nextMatches) != 1 {
			continue
		}
		oldName := previousMatches[0].Name
		newName := nextMatches[0].Name
		if oldName == newName {
			continue
		}
		collision := false
		for _, owner := range desiredNames[strings.ToLower(oldName)] {
			// The current target may differ from history by casing only. That is
			// safe to restore; another constraint occupying the same
			// case-insensitive namespace is not.
			if owner != newName {
				collision = true
				break
			}
		}
		if collision {
			continue
		}
		mapping[newName] = oldName
	}
	return mapping, nil
}

func namedTableConstraintsBySignature(createTableSQL string) (map[string][]namedTableConstraint, error) {
	items, err := extractCreateTableItems(createTableSQL)
	if err != nil {
		return nil, err
	}
	out := make(map[string][]namedTableConstraint)
	for _, item := range items {
		occurrences, err := constraintNameOccurrences(item)
		if err != nil {
			return nil, err
		}
		if len(occurrences) == 0 {
			continue
		}
		for target := range occurrences {
			normalized := item
			for i := len(occurrences) - 1; i >= 0; i-- {
				marker := "__RZA_OTHER_CONSTRAINT_NAME__"
				if i == target {
					marker = "__RZA_TARGET_CONSTRAINT_NAME__"
				}
				occurrence := occurrences[i]
				normalized = normalized[:occurrence.Start] + quoteIdentifier(marker) + normalized[occurrence.End:]
			}
			signature := sqlparse.CanonicalSemanticSQL(normalized)
			name := occurrences[target].Name
			out[signature] = append(out[signature], namedTableConstraint{Name: name, Signature: signature})
		}
	}
	return out, nil
}

func replaceNamedConstraintsInTableSource(sql, object string, mapping map[string]string) (string, int, error) {
	normalized := normalizeSQLText(sql, true, true)
	stmt, start, end, err := extractCreateTableStatementBounds(object, normalized)
	if err != nil {
		return "", 0, err
	}
	items, err := extractCreateTableItems(stmt)
	if err != nil {
		return "", 0, err
	}
	changed := 0
	for i, item := range items {
		occurrences, err := constraintNameOccurrences(item)
		if err != nil {
			return "", 0, err
		}
		if len(occurrences) == 0 {
			continue
		}
		rewritten := item
		for occurrenceIndex := len(occurrences) - 1; occurrenceIndex >= 0; occurrenceIndex-- {
			occurrence := occurrences[occurrenceIndex]
			oldName, ok := mapping[occurrence.Name]
			if !ok {
				continue
			}
			rewritten = rewritten[:occurrence.Start] + quoteIdentifier(oldName) + rewritten[occurrence.End:]
			changed++
		}
		items[i] = rewritten
	}
	if changed == 0 {
		return normalized, 0, nil
	}
	rebuilt, err := rebuildCreateTableStatementWithItems(stmt, items)
	if err != nil {
		return "", 0, err
	}
	return normalizeSQLText(normalized[:start]+rebuilt+normalized[end:], true, true), changed, nil
}

func constraintNameOccurrences(fragment string) ([]constraintNameOccurrence, error) {
	masked := maskSQLQuotedIdentifiers(sqlparse.StripSQLStringsAndComments(fragment))
	upper := strings.ToUpper(masked)
	const keyword = "CONSTRAINT"
	out := make([]constraintNameOccurrence, 0, 1)
	for cursor := 0; cursor < len(fragment); {
		rel := strings.Index(upper[cursor:], keyword)
		if rel < 0 {
			break
		}
		keywordStart := cursor + rel
		keywordEnd := keywordStart + len(keyword)
		if keywordStart > 0 {
			left := upper[keywordStart-1]
			if (left >= 'A' && left <= 'Z') || (left >= '0' && left <= '9') || left == '_' || left == '@' || left == '#' || left == '$' {
				cursor = keywordStart + len(keyword)
				continue
			}
		}
		if keywordEnd >= len(upper) || !isConstraintSQLSpace(upper[keywordEnd]) {
			cursor = keywordEnd
			continue
		}
		nameStart := keywordEnd
		for nameStart < len(fragment) && (fragment[nameStart] == ' ' || fragment[nameStart] == '\t' || fragment[nameStart] == '\r' || fragment[nameStart] == '\n') {
			nameStart++
		}
		token := firstSQLToken(fragment[nameStart:])
		if token == "" {
			return nil, fmt.Errorf("CONSTRAINT keyword has no name")
		}
		name := stripIdentifierBrackets(cleanIdentifierToken(token))
		if name == "" {
			return nil, fmt.Errorf("CONSTRAINT keyword has an invalid name token %q", token)
		}
		out = append(out, constraintNameOccurrence{Name: name, Start: nameStart, End: nameStart + len(token)})
		cursor = nameStart + len(token)
	}
	return out, nil
}

func isConstraintSQLSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\r' || ch == '\n' || ch == '\f'
}

func maskSQLQuotedIdentifiers(input string) string {
	raw := []byte(input)
	for i := 0; i < len(raw); {
		switch raw[i] {
		case '[':
			i++
			for i < len(raw) {
				if raw[i] != ']' {
					raw[i] = ' '
					i++
					continue
				}
				if i+1 < len(raw) && raw[i+1] == ']' {
					raw[i], raw[i+1] = ' ', ' '
					i += 2
					continue
				}
				i++
				break
			}
		case '"':
			i++
			for i < len(raw) {
				if raw[i] != '"' {
					raw[i] = ' '
					i++
					continue
				}
				if i+1 < len(raw) && raw[i+1] == '"' {
					raw[i], raw[i+1] = ' ', ' '
					i += 2
					continue
				}
				i++
				break
			}
		default:
			i++
		}
	}
	return string(raw)
}
