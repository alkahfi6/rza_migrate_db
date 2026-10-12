// Package sqlfix provides conservative, trivia-preserving repairs for SQL
// source files. Rewriters operate on token positions so comments, literals,
// and business expressions are never rewritten by broad text replacement.
package sqlfix

import (
	"crypto/sha256"
	"fmt"
	"regexp"
	"strings"
	"unicode"

	"rza_migrate_db/internal/sqlfmt"
	"rza_migrate_db/internal/sqlparse"
)

type Options struct {
	Category     string
	ObjectName   string
	Patch        bool
	SelfCatalog  string
	LocalObjects map[string]bool
}

type Change struct {
	Rule    string
	Message string
}

type Result struct {
	SQL      string
	Changes  []Change
	Warnings []string
}

type edit struct {
	start int
	end   int
	text  string
	rule  string
	msg   string
}

type tokenKind uint8

const (
	tokenWord tokenKind = iota
	tokenBracket
	tokenSymbol
)

type token struct {
	kind       tokenKind
	start, end int
	text       string
}

// RewriteSafe applies only repairs whose target can be proven from the file
// contract and token context. It is deliberately not a general SQL formatter.
func RewriteSafe(input string, opts Options) (Result, error) {
	sql := sqlfmt.NormalizeSQLText(input, true, true)
	var changes []Change

	var preambleChanges []Change
	sql, preambleChanges = stripBatchArtifacts(sql, opts)
	changes = append(changes, preambleChanges...)

	category := strings.ToLower(strings.TrimSpace(opts.Category))
	if opts.Patch {
		var changed bool
		sql, changed = canonicalizePatchProcAlias(sql)
		if changed {
			changes = append(changes, Change{Rule: "SAFE_PROC_ALIAS", Message: "normalized PROC keyword to PROCEDURE without changing patch operation semantics"})
		}
	} else if category == "procedures" || category == "views" || category == "functions" {
		var changed bool
		var msg string
		sql, changed, msg = canonicalizeModuleHeader(sql, category, opts.ObjectName)
		if changed {
			changes = append(changes, Change{Rule: "SAFE_MODULE_HEADER", Message: msg})
		}
	}

	if category == "tables" {
		var tableChanges []Change
		var err error
		sql, tableChanges, err = repairCreateTable(sql, opts.ObjectName)
		if err != nil {
			return Result{}, err
		}
		changes = append(changes, tableChanges...)
	}
	if strings.TrimSpace(opts.SelfCatalog) != "" && len(opts.LocalObjects) > 0 {
		var referenceChanges []Change
		sql, referenceChanges = rewriteProvenSelfCatalogReferences(sql, opts.SelfCatalog, opts.LocalObjects)
		changes = append(changes, referenceChanges...)
	}

	sql = sqlfmt.NormalizeSQLText(sql, true, true)
	return Result{SQL: sql, Changes: changes}, nil
}

func rewriteProvenSelfCatalogReferences(input, selfCatalog string, localObjects map[string]bool) (string, []Change) {
	tokens := scanTokens(input)
	edits := make([]edit, 0)
	for i := 0; i+4 < len(tokens); i++ {
		if (tokens[i].kind != tokenWord && tokens[i].kind != tokenBracket) ||
			!strings.EqualFold(identifierText(tokens[i]), strings.Trim(selfCatalog, "[]")) ||
			tokens[i+1].text != "." || tokens[i+3].text != "." {
			continue
		}
		schema := identifierText(tokens[i+2])
		object := identifierText(tokens[i+4])
		if schema == "" || object == "" || !localObjects[strings.ToLower(schema+"."+object)] {
			continue
		}
		edits = append(edits, edit{
			start: tokens[i].start, end: tokens[i+2].start,
			rule: "SAFE_REMOVE_SELF_CATALOG",
			msg:  "removed self-catalog qualifier from proven local object " + schema + "." + object,
		})
		i += 4
	}
	return applyEdits(input, edits)
}

func canonicalizePatchProcAlias(input string) (string, bool) {
	tokens := scanTokens(input)
	edits := make([]edit, 0)
	for i := 1; i < len(tokens); i++ {
		if !isWord(tokens[i], "PROC") {
			continue
		}
		previous := i - 1
		if !isWord(tokens[previous], "CREATE") && !isWord(tokens[previous], "ALTER") && !isWord(tokens[previous], "DROP") {
			continue
		}
		edits = append(edits, edit{start: tokens[i].start, end: tokens[i].end, text: "PROCEDURE"})
	}
	out, _ := applyEdits(input, edits)
	return out, len(edits) > 0
}

func stripBatchArtifacts(input string, opts Options) (string, []Change) {
	tokens := scanTokens(input)
	lineStarts := []int{0}
	for i, ch := range []byte(input) {
		if ch == '\n' {
			lineStarts = append(lineStarts, i+1)
		}
	}
	lineEnds := make([]int, len(lineStarts))
	for i := range lineStarts {
		if i+1 < len(lineStarts) {
			lineEnds[i] = lineStarts[i+1]
		} else {
			lineEnds[i] = len(input)
		}
	}

	lineTokens := make([][]token, len(lineStarts))
	lineIndex := 0
	for _, tok := range tokens {
		for lineIndex+1 < len(lineStarts) && tok.start >= lineStarts[lineIndex+1] {
			lineIndex++
		}
		lineTokens[lineIndex] = append(lineTokens[lineIndex], tok)
	}

	definitionStart := len(input)
	for i := 0; i+1 < len(tokens); i++ {
		if isWord(tokens[i], "CREATE") || isWord(tokens[i], "ALTER") {
			definitionStart = tokens[i].start
			break
		}
	}

	seenPatchSQL := false
	removedUse := false
	edits := make([]edit, 0)
	for i, toks := range lineTokens {
		if len(toks) == 0 {
			continue
		}
		beforeDefinition := toks[0].start < definitionStart
		isGO := len(toks) == 1 && isWord(toks[0], "GO")
		isUSE := isUseLine(toks)
		isModuleSET := isModuleSetLine(toks)

		remove := false
		rule := ""
		msg := ""
		if opts.Patch {
			switch {
			case isUSE && !seenPatchSQL:
				remove, removedUse = true, true
				rule, msg = "SAFE_REMOVE_INITIAL_CATALOG", "removed leading USE catalog directive; migrations run in the configured target catalog"
			case isGO && !seenPatchSQL && removedUse:
				remove = true
				rule, msg = "SAFE_REMOVE_BATCH_SEPARATOR", "removed GO paired with the leading catalog directive"
			default:
				seenPatchSQL = true
			}
		} else {
			switch {
			case isUSE && beforeDefinition:
				remove = true
				rule, msg = "SAFE_REMOVE_INITIAL_CATALOG", "removed leading USE catalog directive from desired-state object"
			case isModuleSET && beforeDefinition:
				remove = true
				rule, msg = "SAFE_REMOVE_MODULE_SET", "removed SSMS module SET preamble"
			case isGO:
				remove = true
				rule, msg = "SAFE_REMOVE_BATCH_SEPARATOR", "removed GO batch separator from single-object desired state"
			}
		}
		if remove {
			edits = append(edits, edit{start: lineStarts[i], end: lineEnds[i], rule: rule, msg: msg})
		}
	}
	return applyEdits(input, edits)
}

func isUseLine(tokens []token) bool {
	if len(tokens) < 2 || !isWord(tokens[0], "USE") {
		return false
	}
	for _, tok := range tokens[1:] {
		if tok.kind == tokenWord || tok.kind == tokenBracket || tok.text == "." || tok.text == ";" {
			continue
		}
		return false
	}
	return true
}

func isModuleSetLine(tokens []token) bool {
	if len(tokens) < 3 || !isWord(tokens[0], "SET") {
		return false
	}
	if !isWord(tokens[1], "ANSI_NULLS") && !isWord(tokens[1], "QUOTED_IDENTIFIER") {
		return false
	}
	return isWord(tokens[2], "ON") || isWord(tokens[2], "OFF")
}

func canonicalizeModuleHeader(input, category, objectName string) (string, bool, string) {
	tokens := scanTokens(input)
	for i := 0; i < len(tokens); i++ {
		if !isWord(tokens[i], "CREATE") && !isWord(tokens[i], "ALTER") {
			continue
		}
		j := i + 1
		if j+1 < len(tokens) && isWord(tokens[j], "OR") && isWord(tokens[j+1], "ALTER") {
			j += 2
		}
		if j >= len(tokens) {
			continue
		}
		kind := ""
		switch {
		case isWord(tokens[j], "PROC") || isWord(tokens[j], "PROCEDURE"):
			kind = "PROCEDURE"
		case isWord(tokens[j], "VIEW"):
			kind = "VIEW"
		case isWord(tokens[j], "FUNCTION"):
			kind = "FUNCTION"
		default:
			continue
		}
		if category != "" && category != "patch" {
			expected := strings.TrimSuffix(strings.ToUpper(category), "S")
			if expected == "PROCEDURE" && kind != "PROCEDURE" || expected == "VIEW" && kind != "VIEW" || expected == "FUNCTION" && kind != "FUNCTION" {
				continue
			}
		}
		canonical := "CREATE OR ALTER " + kind
		edits := make([]edit, 0, 2)
		if input[tokens[i].start:tokens[j].end] != canonical {
			edits = append(edits, edit{start: tokens[i].start, end: tokens[j].end, text: canonical})
		}
		if j+1 < len(tokens) && (tokens[j+1].kind == tokenWord || tokens[j+1].kind == tokenBracket) {
			expectedBase := objectBaseName(objectName)
			actualBase := sanitizeName(identifierText(tokens[j+1]))
			qualified := j+2 < len(tokens) && tokens[j+2].text == "."
			if !qualified && expectedBase != "" && strings.EqualFold(expectedBase, actualBase) {
				schema := "dbo"
				if parts := strings.Split(objectName, "."); len(parts) > 1 && sanitizeName(parts[0]) != "" {
					schema = sanitizeName(parts[0])
				}
				edits = append(edits, edit{start: tokens[j+1].start, end: tokens[j+1].start, text: schema + "."})
			}
		}
		out, _ := applyEdits(input, edits)
		if len(edits) == 0 {
			return input, false, ""
		}
		return out, true, fmt.Sprintf("normalized module header to %s and repository-qualified object name", canonical)
	}
	return input, false, ""
}

func repairCreateTable(input, objectName string) (string, []Change, error) {
	tableName := objectBaseName(objectName)
	if tableName == "" {
		return input, nil, fmt.Errorf("safe table repair requires an object name")
	}

	// Fold the legacy extractor form into the desired-state CREATE TABLE before
	// naming defaults. This is intentionally limited to ADD DEFAULT ... FOR.
	var folded []Change
	input, folded = foldTrailingAlterDefault(input, tableName)

	tokens := scanTokens(input)
	open := -1
	headerEdits := make([]edit, 0, 1)
	for i := 0; i+2 < len(tokens); i++ {
		if isWord(tokens[i], "CREATE") && isWord(tokens[i+1], "TABLE") {
			if (tokens[i+2].kind == tokenWord || tokens[i+2].kind == tokenBracket) &&
				(i+3 >= len(tokens) || tokens[i+3].text != ".") &&
				strings.EqualFold(sanitizeName(identifierText(tokens[i+2])), tableName) {
				headerEdits = append(headerEdits, edit{
					start: tokens[i+2].start, end: tokens[i+2].start, text: "dbo.",
					rule: "SAFE_QUALIFY_OBJECT", msg: "qualified CREATE TABLE target from the repository filename contract",
				})
			}
			for j := i + 2; j < len(tokens); j++ {
				if tokens[j].text == "(" {
					open = tokens[j].start
					break
				}
			}
			break
		}
	}
	if open < 0 {
		return input, folded, nil
	}
	close, ok := sqlparse.FindMatchingParen(input, open)
	if !ok {
		return input, nil, fmt.Errorf("CREATE TABLE %s has an unbalanced item list", objectName)
	}

	edits := append([]edit{}, headerEdits...)
	insideTokens := tokensWithin(scanTokens(input), open+1, close)
	if len(insideTokens) > 0 && insideTokens[len(insideTokens)-1].text == "," {
		tok := insideTokens[len(insideTokens)-1]
		edits = append(edits, edit{start: tok.start, end: tok.end, rule: "SAFE_REMOVE_TRAILING_COMMA", msg: "removed trailing comma from CREATE TABLE item list"})
	}

	for _, span := range splitTableItemSpans(input, open+1, close) {
		itemTokens := tokensWithin(scanTokens(input[span.start:span.end]), 0, span.end-span.start)
		for i := range itemTokens {
			itemTokens[i].start += span.start
			itemTokens[i].end += span.start
		}
		if len(itemTokens) == 0 {
			continue
		}
		if len(itemTokens) > 2 && isKnownNonCollatableType(itemTokens[1]) {
			if collateIndex := wordIndex(itemTokens, "COLLATE"); collateIndex >= 0 && collateIndex+1 < len(itemTokens) {
				edits = append(edits, edit{
					start: itemTokens[collateIndex].start,
					end:   itemTokens[collateIndex+1].end,
					rule:  "SAFE_REMOVE_INVALID_COLLATION",
					msg:   "removed COLLATE from non-collatable " + strings.ToUpper(identifierText(itemTokens[1])) + " column " + identifierText(itemTokens[0]),
				})
			}
		}
		edits = append(edits, nameProvenUnnamedConstraint(input, itemTokens, tableName)...)
	}

	out, changes := applyEdits(input, edits)
	return out, append(folded, changes...), nil
}

// nameProvenUnnamedConstraint names only constraint forms whose owning table,
// column list, and keyword position are structurally unambiguous. Malformed
// CONSTRAINT clauses and items with multiple constraint clauses are preserved
// for manual review instead of guessing business intent.
func nameProvenUnnamedConstraint(input string, tokens []token, tableName string) []edit {
	if len(tokens) == 0 || isWord(tokens[0], "CONSTRAINT") {
		return nil
	}

	tableLevel := isTableConstraintStart(tokens)
	column := ""
	if !tableLevel {
		column = identifierText(tokens[0])
		if column == "" {
			return nil
		}
	}

	type candidate struct {
		index  int
		prefix string
		kind   string
	}
	candidates := make([]candidate, 0, 1)
	for i := 0; i < len(tokens); i++ {
		if constraintNameAppliesTo(tokens, i) {
			continue
		}
		switch {
		case isWord(tokens[i], "DEFAULT"):
			candidates = append(candidates, candidate{i, "DF", "default"})
		case isWord(tokens[i], "CHECK"):
			candidates = append(candidates, candidate{i, "CK", "check"})
		case isWord(tokens[i], "UNIQUE"):
			candidates = append(candidates, candidate{i, "UQ", "unique"})
		case isWord(tokens[i], "PRIMARY") && i+1 < len(tokens) && isWord(tokens[i+1], "KEY"):
			candidates = append(candidates, candidate{i, "PK", "primary key"})
		case isWord(tokens[i], "FOREIGN") && i+1 < len(tokens) && isWord(tokens[i+1], "KEY"):
			candidates = append(candidates, candidate{i, "FK", "foreign key"})
		case isWord(tokens[i], "REFERENCES") && !tableLevel:
			candidates = append(candidates, candidate{i, "FK", "foreign key"})
		}
	}
	if len(candidates) != 1 {
		return nil
	}

	selected := candidates[0]
	detail := column
	if tableLevel {
		detail = strings.Join(constraintColumnList(tokens, selected.index), "_")
	}
	if selected.prefix == "DF" && tableLevel {
		detail = ""
		if forIndex := nextWordIndex(tokens, selected.index+1, "FOR"); forIndex >= 0 && forIndex+1 < len(tokens) {
			detail = identifierText(tokens[forIndex+1])
		}
	}
	if selected.prefix == "CK" {
		base := detail
		if base == "" {
			base = "rule"
		}
		raw := input[tokens[0].start:tokens[len(tokens)-1].end]
		digest := fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(strings.Fields(raw), " "))))[:8]
		detail = base + "_" + digest
	}
	if selected.prefix == "FK" {
		if target := referencedConstraintTarget(tokens, selected.index); target != "" {
			if detail == "" {
				detail = target
			} else {
				detail += "_" + target
			}
		}
		raw := input[tokens[0].start:tokens[len(tokens)-1].end]
		digest := fmt.Sprintf("%x", sha256.Sum256([]byte(strings.Join(strings.Fields(raw), " "))))[:8]
		detail += "_" + digest
	}
	if detail == "" && selected.prefix != "PK" {
		return nil
	}
	name := constraintName(selected.prefix, tableName, detail)
	if selected.prefix == "PK" {
		name = boundedConstraintName("PK_" + sanitizeName(tableName))
	}
	rule := "SAFE_NAME_" + strings.ReplaceAll(strings.ToUpper(selected.kind), " ", "_") + "_CONSTRAINT"
	return []edit{{
		start: tokens[selected.index].start,
		end:   tokens[selected.index].start,
		text:  "CONSTRAINT " + name + " ",
		rule:  rule,
		msg:   "named " + selected.kind + " constraint " + name,
	}}
}

func constraintNameAppliesTo(tokens []token, keywordIndex int) bool {
	constraintIndex := -1
	for i := 0; i < keywordIndex; i++ {
		if isWord(tokens[i], "CONSTRAINT") {
			constraintIndex = i
		}
	}
	if constraintIndex < 0 || constraintIndex+1 >= keywordIndex || identifierText(tokens[constraintIndex+1]) == "" {
		return false
	}
	for i := constraintIndex + 2; i < keywordIndex; i++ {
		if isConstraintKindKeyword(tokens, i) {
			return false
		}
	}
	return true
}

func isConstraintKindKeyword(tokens []token, index int) bool {
	return isWord(tokens[index], "DEFAULT") || isWord(tokens[index], "CHECK") || isWord(tokens[index], "UNIQUE") ||
		isWord(tokens[index], "REFERENCES") ||
		isWord(tokens[index], "PRIMARY") && index+1 < len(tokens) && isWord(tokens[index+1], "KEY") ||
		isWord(tokens[index], "FOREIGN") && index+1 < len(tokens) && isWord(tokens[index+1], "KEY")
}

func referencedConstraintTarget(tokens []token, start int) string {
	references := nextWordIndex(tokens, start, "REFERENCES")
	if references < 0 || references+1 >= len(tokens) {
		return ""
	}
	first := identifierText(tokens[references+1])
	if first == "" {
		return ""
	}
	if references+3 < len(tokens) && tokens[references+2].text == "." {
		if second := identifierText(tokens[references+3]); second != "" {
			return sanitizeName(second)
		}
	}
	return sanitizeName(first)
}

func constraintColumnList(tokens []token, keywordIndex int) []string {
	open := -1
	for i := keywordIndex + 1; i < len(tokens); i++ {
		if tokens[i].text == "(" {
			open = i
			break
		}
	}
	if open < 0 {
		return nil
	}
	columns := make([]string, 0, 2)
	depth := 0
	expectIdentifier := true
	for i := open; i < len(tokens); i++ {
		switch tokens[i].text {
		case "(":
			depth++
		case ")":
			depth--
			if depth == 0 {
				return columns
			}
		case ",":
			if depth == 1 {
				expectIdentifier = true
			}
		default:
			if depth == 1 && expectIdentifier {
				if value := identifierText(tokens[i]); value != "" {
					columns = append(columns, value)
					expectIdentifier = false
				}
			}
		}
	}
	return nil
}

func isKnownNonCollatableType(tok token) bool {
	if tok.kind != tokenWord && tok.kind != tokenBracket {
		return false
	}
	switch strings.ToUpper(identifierText(tok)) {
	case "XML", "BIGINT", "INT", "SMALLINT", "TINYINT", "BIT", "DECIMAL", "NUMERIC", "MONEY", "SMALLMONEY",
		"FLOAT", "REAL", "DATE", "TIME", "DATETIME", "DATETIME2", "SMALLDATETIME", "DATETIMEOFFSET",
		"UNIQUEIDENTIFIER", "BINARY", "VARBINARY", "IMAGE", "TIMESTAMP", "ROWVERSION", "GEOMETRY", "GEOGRAPHY":
		return true
	default:
		return false
	}
}

// DetectIndexTargets returns distinct schema-qualified table targets from
// top-level CREATE INDEX statements. Dynamic SQL and comments are ignored.
func DetectIndexTargets(input string) []string {
	tokens := scanTokens(input)
	seen := map[string]bool{}
	out := make([]string, 0)
	for i := 0; i < len(tokens); i++ {
		if !isWord(tokens[i], "CREATE") {
			continue
		}
		indexAt := nextWordIndex(tokens, i+1, "INDEX")
		if indexAt < 0 || indexAt > i+5 {
			continue
		}
		onAt := nextWordIndex(tokens, indexAt+1, "ON")
		if onAt < 0 || onAt > indexAt+4 || onAt+1 >= len(tokens) {
			continue
		}
		first := identifierText(tokens[onAt+1])
		if first == "" {
			continue
		}
		schema, name := "dbo", first
		if onAt+3 < len(tokens) && tokens[onAt+2].text == "." {
			schema = first
			name = identifierText(tokens[onAt+3])
		}
		if name == "" {
			continue
		}
		target := sanitizeName(schema) + "." + sanitizeName(name)
		key := strings.ToLower(target)
		if !seen[key] {
			seen[key] = true
			out = append(out, target)
		}
		i = onAt
	}
	return out
}

type span struct{ start, end int }

func splitTableItemSpans(input string, start, end int) []span {
	tokens := tokensWithin(scanTokens(input), start, end)
	out := make([]span, 0)
	itemStart := start
	depth := 0
	for _, tok := range tokens {
		switch tok.text {
		case "(":
			depth++
		case ")":
			if depth > 0 {
				depth--
			}
		case ",":
			if depth == 0 {
				out = append(out, span{start: itemStart, end: tok.start})
				itemStart = tok.end
			}
		}
	}
	out = append(out, span{start: itemStart, end: end})
	return out
}

func foldTrailingAlterDefault(input, tableName string) (string, []Change) {
	tokens := scanTokens(input)
	for i := 0; i+5 < len(tokens); i++ {
		if !isWord(tokens[i], "ALTER") || !isWord(tokens[i+1], "TABLE") {
			continue
		}
		add := nextWordIndex(tokens, i+2, "ADD")
		if add < 0 || add > i+6 {
			continue
		}
		def := nextWordIndex(tokens, add+1, "DEFAULT")
		if def < 0 || def > add+2 {
			continue
		}
		forIndex := nextWordIndex(tokens, def+1, "FOR")
		if forIndex < 0 || forIndex+1 >= len(tokens) {
			continue
		}
		column := identifierText(tokens[forIndex+1])
		if column == "" {
			continue
		}
		stmtEnd := statementLineEnd(input, tokens[i].start)
		expression := strings.TrimSpace(input[tokens[def].end:tokens[forIndex].start])
		if expression == "" {
			continue
		}

		createOpen, createClose := createTableBounds(input)
		if createOpen < 0 || createClose < 0 {
			continue
		}
		insertAt := -1
		for _, item := range splitTableItemSpans(input, createOpen+1, createClose) {
			its := tokensWithin(scanTokens(input[item.start:item.end]), 0, item.end-item.start)
			if len(its) == 0 || !strings.EqualFold(identifierText(its[0]), column) {
				continue
			}
			if wordIndex(its, "DEFAULT") >= 0 {
				return input, nil
			}
			insertAt = trimSpaceEnd(input, item.start, item.end)
			break
		}
		if insertAt < 0 {
			continue
		}
		name := constraintName("DF", tableName, column)
		edits := []edit{
			{start: insertAt, end: insertAt, text: " CONSTRAINT " + name + " DEFAULT " + expression, rule: "SAFE_FOLD_DEFAULT_CONSTRAINT", msg: "folded and named ALTER TABLE default " + name},
			{start: tokens[i].start, end: stmtEnd, rule: "SAFE_FOLD_DEFAULT_CONSTRAINT", msg: "removed folded ALTER TABLE default statement"},
		}
		return applyEdits(input, edits)
	}
	return input, nil
}

func createTableBounds(input string) (int, int) {
	tokens := scanTokens(input)
	for i := 0; i+2 < len(tokens); i++ {
		if isWord(tokens[i], "CREATE") && isWord(tokens[i+1], "TABLE") {
			for j := i + 2; j < len(tokens); j++ {
				if tokens[j].text == "(" {
					if close, ok := sqlparse.FindMatchingParen(input, tokens[j].start); ok {
						return tokens[j].start, close
					}
				}
			}
		}
	}
	return -1, -1
}

func statementLineEnd(input string, start int) int {
	if semi := sqlparse.FindStatementEnd(input, start); semi < len(input) {
		return semi
	}
	if rel := strings.IndexByte(input[start:], '\n'); rel >= 0 {
		return start + rel + 1
	}
	return len(input)
}

func trimSpaceEnd(input string, start, end int) int {
	for end > start && unicode.IsSpace(rune(input[end-1])) {
		end--
	}
	return end
}

func isTableConstraintStart(tokens []token) bool {
	return isWord(tokens[0], "CONSTRAINT") || isWord(tokens[0], "PRIMARY") || isWord(tokens[0], "UNIQUE") || isWord(tokens[0], "CHECK") || isWord(tokens[0], "FOREIGN")
}

func constraintName(prefix, table, detail string) string {
	base := prefix + "_" + sanitizeName(table) + "__" + sanitizeName(detail)
	return boundedConstraintName(base)
}

func boundedConstraintName(base string) string {
	if len(base) <= 128 {
		return base
	}
	hash := fmt.Sprintf("%x", sha256.Sum256([]byte(base)))[:10]
	return base[:117] + "_" + hash
}

var invalidNameRun = regexp.MustCompile(`[^A-Za-z0-9_]+`)

func sanitizeName(value string) string {
	value = strings.Trim(value, "[]\"`")
	value = invalidNameRun.ReplaceAllString(value, "_")
	value = strings.Trim(value, "_")
	if value == "" {
		return "unnamed"
	}
	return value
}

func objectBaseName(value string) string {
	parts := strings.Split(value, ".")
	return sanitizeName(parts[len(parts)-1])
}

func identifierText(tok token) string {
	if tok.kind != tokenWord && tok.kind != tokenBracket {
		return ""
	}
	return strings.ReplaceAll(strings.Trim(tok.text, "[]"), "]]", "]")
}

func wordIndex(tokens []token, word string) int { return nextWordIndex(tokens, 0, word) }

func wordIndexBefore(tokens []token, word string, before int) int {
	for i := 0; i < before && i < len(tokens); i++ {
		if isWord(tokens[i], word) {
			return i
		}
	}
	return -1
}

func nextWordIndex(tokens []token, start int, word string) int {
	for i := start; i < len(tokens); i++ {
		if isWord(tokens[i], word) {
			return i
		}
	}
	return -1
}

func tokensWithin(tokens []token, start, end int) []token {
	out := make([]token, 0)
	for _, tok := range tokens {
		if tok.start >= start && tok.end <= end {
			out = append(out, tok)
		}
	}
	return out
}

func isWord(tok token, word string) bool {
	return tok.kind == tokenWord && strings.EqualFold(tok.text, word)
}

func applyEdits(input string, edits []edit) (string, []Change) {
	if len(edits) == 0 {
		return input, nil
	}
	// Stable insertion sort keeps this package dependency-free and edit lists
	// are tiny per SQL object.
	for i := 1; i < len(edits); i++ {
		for j := i; j > 0 && edits[j].start > edits[j-1].start; j-- {
			edits[j], edits[j-1] = edits[j-1], edits[j]
		}
	}
	out := input
	changes := make([]Change, 0, len(edits))
	seen := map[string]bool{}
	lastStart := len(input) + 1
	for _, e := range edits {
		if e.start < 0 || e.end < e.start || e.end > len(out) || e.end > lastStart {
			continue
		}
		out = out[:e.start] + e.text + out[e.end:]
		lastStart = e.start
		key := e.rule + "\x00" + e.msg
		if !seen[key] {
			seen[key] = true
			changes = append(changes, Change{Rule: e.rule, Message: e.msg})
		}
	}
	return out, changes
}

func scanTokens(input string) []token {
	out := make([]token, 0)
	for i := 0; i < len(input); {
		ch := input[i]
		switch {
		case unicode.IsSpace(rune(ch)):
			i++
		case ch == '-' && i+1 < len(input) && input[i+1] == '-':
			i += 2
			for i < len(input) && input[i] != '\n' {
				i++
			}
		case ch == '/' && i+1 < len(input) && input[i+1] == '*':
			i += 2
			depth := 1
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
		case ch == '\'':
			i++
			for i < len(input) {
				if input[i] == '\'' {
					if i+1 < len(input) && input[i+1] == '\'' {
						i += 2
						continue
					}
					i++
					break
				}
				i++
			}
		case ch == '[':
			start := i
			i++
			for i < len(input) {
				if input[i] == ']' {
					if i+1 < len(input) && input[i+1] == ']' {
						i += 2
						continue
					}
					i++
					break
				}
				i++
			}
			out = append(out, token{kind: tokenBracket, start: start, end: i, text: input[start:i]})
		case isWordByte(ch):
			start := i
			i++
			for i < len(input) && isWordByte(input[i]) {
				i++
			}
			out = append(out, token{kind: tokenWord, start: start, end: i, text: input[start:i]})
		default:
			out = append(out, token{kind: tokenSymbol, start: i, end: i + 1, text: input[i : i+1]})
			i++
		}
	}
	return out
}

func isWordByte(ch byte) bool {
	return ch >= 'a' && ch <= 'z' || ch >= 'A' && ch <= 'Z' || ch >= '0' && ch <= '9' || ch == '_' || ch == '@' || ch == '#' || ch == '$'
}
