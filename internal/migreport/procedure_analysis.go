package migreport

import (
	"fmt"
	"regexp"
	"sort"
	"strings"

	"rza_migrate_db/internal/sqlparse"
)

const (
	ResolutionStatic         = "RESOLVED_STATIC"
	ResolutionDynamic        = "RESOLVED_DYNAMIC"
	ResolutionFromCaller     = "RESOLVED_FROM_CALLER"
	ResolutionExternal       = "EXTERNAL_DATABASE"
	ResolutionUnresolved     = "UNRESOLVED_DYNAMIC"
	ResolutionNotPersisted   = "NOT_PERSISTED"
	ResolutionNotFound       = "NOT_FOUND"
	ResolutionColumnNotFound = "COLUMN_NOT_FOUND"
	ResolutionAmbiguous      = "AMBIGUOUS"
	ResolutionSystem         = "SYSTEM_OBJECT"
	procedureDiagramMaxNodes = 20
)

type ProcedureProfile struct {
	Name              string
	SourcePath        string
	Parameters        []ProcedureParameter
	Accesses          []ProcedureDataAccess
	Mappings          []ProcedureWriteMapping
	Calls             []ProcedureCall
	Functions         []ProcedureDependency
	Views             []ProcedureDependency
	DynamicSQL        []ProcedureDynamicSQL
	ControlSteps      []ProcedureControlStep
	ReturnCodes       []string
	HasRaiseError     bool
	HasThrow          bool
	HasTransaction    bool
	HasTryCatch       bool
	HasDirectResult   bool
	ExtractionNotices []string
}

type ProcedureParameter struct {
	Name, SQLType, Direction, Default string
}

type ProcedureDataAccess struct {
	Ref, Operation, Object, Resolution, Source, Guard string
	Line                                              int
	statementStart, statementEnd                      int
	offset, targetEnd                                 int
	queryDepth                                        int
}

type ProcedureWriteMapping struct {
	Ref, Object, Column, Expression, Filter, Resolution string
	Line                                                int
}

type ProcedureCall struct {
	Object, Arguments, Resolution, Source, Guard string
	Line                                         int
	offset                                       int
}

type ProcedureControlStep struct {
	Event, Evidence string
	Line            int
}

type ProcedureDependency struct {
	Object, Resolution string
}

type ProcedureDynamicSQL struct {
	Line, Expression, RuntimeSource, Resolution string
	ResolvedObjects                             []string
}

type procedureCatalog struct {
	byKind map[string]map[string][]string
	tables map[string]*DocumentedTable
}

var (
	returnPattern = regexp.MustCompile(`(?i)\bRETURN\s+(-?[0-9]+)\b`)
)

func buildProcedureProfiles(modules []moduleSource, tables map[string]*DocumentedTable, views, functions, procedures map[string]struct{}) []ProcedureProfile {
	catalog := newProcedureCatalog(tables, views, functions, procedures)
	profiles := make([]ProcedureProfile, 0)
	for _, module := range modules {
		if module.Kind != "procedure" {
			continue
		}
		profile := analyzeProcedure(module, catalog)
		profile.SourcePath = "db_object/procedures/" + module.Name + ".sql"
		profiles = append(profiles, profile)
	}
	sort.Slice(profiles, func(i, j int) bool { return strings.ToLower(profiles[i].Name) < strings.ToLower(profiles[j].Name) })
	applyCallerResolution(profiles, catalog)
	return profiles
}

func newProcedureCatalog(tables map[string]*DocumentedTable, views, functions, procedures map[string]struct{}) procedureCatalog {
	c := procedureCatalog{tables: tables, byKind: map[string]map[string][]string{"table": {}, "view": {}, "function": {}, "procedure": {}}}
	for name := range tables {
		c.add("table", name)
	}
	for name := range views {
		c.add("view", name)
	}
	for name := range functions {
		c.add("function", name)
	}
	for name := range procedures {
		c.add("procedure", name)
	}
	return c
}

func (c procedureCatalog) add(kind, name string) {
	key := strings.ToLower(name)
	c.byKind[kind][key] = append(c.byKind[kind][key], name)
}

func analyzeProcedure(module moduleSource, catalog procedureCatalog) ProcedureProfile {
	p := ProcedureProfile{Name: module.Name, Parameters: parseProcedureParameters(module.SQL)}
	clean := sqlparse.StripSQLStringsAndComments(module.SQL)
	originalLines, cleanLines := strings.Split(module.SQL, "\n"), strings.Split(clean, "\n")
	p.Accesses = extractProcedureAccesses(module.SQL, originalLines, cleanLines, catalog)
	p.Mappings = extractProcedureWriteMappings(module.SQL, p.Accesses, catalog)
	validateWriteMappingColumns(p.Mappings, catalog)
	p.Calls = extractProcedureCalls(module.SQL, originalLines, cleanLines, catalog)
	p.DynamicSQL = extractProcedureDynamicSQL(module.SQL, originalLines, catalog)
	p.ControlSteps = extractProcedureControlSteps(originalLines, cleanLines)
	p.Functions, p.Views = extractRoutineDependencies(module.SQL, catalog)
	p.ReturnCodes = uniqueSortedMatches(returnPattern, clean)
	upper := strings.ToUpper(clean)
	p.HasRaiseError = strings.Contains(upper, "RAISERROR")
	p.HasThrow = containsSQLWord(upper, "THROW")
	p.HasTransaction = strings.Contains(upper, "BEGIN TRAN") || strings.Contains(upper, "BEGIN TRANSACTION")
	p.HasTryCatch = strings.Contains(upper, "BEGIN TRY") && strings.Contains(upper, "BEGIN CATCH")
	p.HasDirectResult = sqlparse.HasClientResultSet(module.SQL)
	p.ExtractionNotices = procedureExtractionNotices(p)
	for _, access := range p.Accesses {
		if access.Operation != "INSERT" || access.Resolution != ResolutionStatic {
			continue
		}
		if _, valid := insertTargetColumns(module.SQL, access); !valid {
			p.ExtractionNotices = append(p.ExtractionNotices, fmt.Sprintf("INSERT baris %d: daftar target column memakai qualifier/expression atau delimiter yang tidak didukung; exact mapping tidak diterbitkan. Periksa source SQL.", access.Line))
		}
	}
	return p
}

func parseProcedureParameters(sql string) []ProcedureParameter {
	tokens := sqlparse.SQLTokens(sql)
	header := -1
	for i, t := range tokens {
		if !t.Quoted && (t.Upper == "PROCEDURE" || t.Upper == "PROC") {
			header = i
			break
		}
	}
	if header < 0 {
		return nil
	}
	start, end := -1, -1
	for i := header + 1; i < len(tokens); i++ {
		t := tokens[i]
		if t.Depth != 0 || t.Quoted {
			continue
		}
		if t.Upper == "WITH" {
			end = t.Start
			break
		}
		if t.Upper == "AS" {
			end = t.Start
			break
		}
		if start < 0 && strings.HasPrefix(t.Text, "@") {
			start = t.Start
		}
	}
	if start < 0 || end <= start {
		return nil
	}
	var out []ProcedureParameter
	for _, part := range sqlparse.SplitTopLevelComma(sql[start:end]) {
		part = strings.TrimSpace(part)
		ts := sqlparse.SQLTokens(part)
		if len(ts) < 2 || !strings.HasPrefix(ts[0].Text, "@") {
			continue
		}
		direction := "INPUT"
		typeEnd, defaultStart, defaultEnd := len(part), -1, len(part)
		for _, t := range ts[1:] {
			if t.Depth != 0 || t.Quoted {
				continue
			}
			if t.Text == "=" {
				typeEnd = t.Start
				defaultStart = t.End
			}
			if t.Upper == "OUT" || t.Upper == "OUTPUT" {
				direction = "INPUT/OUTPUT"
				if defaultStart < 0 {
					typeEnd = t.Start
				} else {
					defaultEnd = t.Start
				}
			}
			if t.Upper == "READONLY" && defaultStart < 0 {
				typeEnd = t.Start
			}
		}
		def := ""
		if defaultStart >= 0 {
			def = strings.TrimSpace(part[defaultStart:defaultEnd])
		}
		out = append(out, ProcedureParameter{Name: ts[0].Text, SQLType: compactSQL(part[ts[0].End:typeEnd]), Direction: direction, Default: def})
	}
	return out
}

func topLevelEqual(s string) int {
	for _, token := range sqlparse.SQLTokens(s) {
		if token.Depth == 0 && !token.Quoted && token.Text == "=" {
			return token.Start
		}
	}
	return -1
}

func extractProcedureAccesses(sql string, original, clean []string, catalog procedureCatalog) []ProcedureDataAccess {
	scanned := sqlparse.ExtractSQLAccesses(sql)
	refs := map[string]int{"UPDATE": 0, "INSERT": 0, "DELETE": 0, "MERGE": 0, "READ": 0}
	out := make([]ProcedureDataAccess, 0, len(scanned))
	for _, access := range scanned {
		raw := access.Name
		if raw == "" || (access.Function && isSQLPseudoObject(raw)) {
			continue
		}
		if access.Operation == "READ" && access.Function {
			_, functionStatus := catalog.resolveParts(raw, access.NameParts, "function")
			if functionStatus == ResolutionStatic {
				continue
			}
		}
		object, resolution := raw, ResolutionNotPersisted
		if !access.NonPersisted {
			object, resolution = catalog.resolveParts(raw, access.NameParts, "data")
		}
		if access.Ambiguous || access.CTECaseAmbiguous {
			resolution = ResolutionAmbiguous
		}
		line := sourceLineNumber(sql, access.Offset)
		refs[access.Operation]++
		prefix := map[string]string{"UPDATE": "U", "INSERT": "I", "DELETE": "D", "MERGE": "M", "READ": "R"}[access.Operation]
		sourceEnd := access.StatementEnd
		if sourceEnd > len(sql) {
			sourceEnd = len(sql)
		}
		out = append(out, ProcedureDataAccess{
			Ref: fmt.Sprintf("%s%d", prefix, refs[access.Operation]), Operation: access.Operation,
			Object: object, Resolution: resolution, Source: compactSQL(sql[access.Offset:sourceEnd]),
			Guard: procedureGuardForLine(original, clean, line-1), Line: line,
			statementStart: access.StatementStart, statementEnd: access.StatementEnd,
			offset: access.Offset, targetEnd: access.End, queryDepth: access.QueryDepth,
		})
	}
	resolveStructuredDMLAliases(out, scanned, catalog)
	return out
}

func sourceLineNumber(sql string, offset int) int {
	if offset < 0 {
		return 1
	}
	if offset > len(sql) {
		offset = len(sql)
	}
	return strings.Count(sql[:offset], "\n") + 1
}

func resolveStructuredDMLAliases(accesses []ProcedureDataAccess, scanned []sqlparse.SQLAccess, catalog procedureCatalog) {
	for i := range accesses {
		if accesses[i].Operation == "READ" || accesses[i].Resolution == ResolutionStatic || accesses[i].Resolution == ResolutionExternal || accesses[i].Resolution == ResolutionNotPersisted || accesses[i].Resolution == ResolutionAmbiguous || strings.Contains(accesses[i].Object, ".") {
			continue
		}
		parts := strings.Split(accesses[i].Object, ".")
		alias := parts[len(parts)-1]
		type aliasMatch struct{ object, status string }
		matches := map[string]aliasMatch{}
		for _, candidate := range scanned {
			if candidate.Operation != "READ" || !strings.EqualFold(candidate.Alias, alias) ||
				candidate.QueryDepth != accesses[i].queryDepth || candidate.StatementStart != accesses[i].statementStart || candidate.StatementEnd != accesses[i].statementEnd {
				continue
			}
			object, status := catalog.resolveParts(candidate.Name, candidate.NameParts, "data")
			if candidate.NonPersisted {
				status = ResolutionNotPersisted
			}
			matches[object] = aliasMatch{object, status}
		}
		if len(matches) == 1 {
			for _, candidate := range matches {
				accesses[i].Object, accesses[i].Resolution = candidate.object, candidate.status
			}
			continue
		}
		if len(matches) > 1 {
			accesses[i].Resolution = ResolutionAmbiguous
		}
	}
}

func extractProcedureCalls(sql string, original, clean []string, catalog procedureCatalog) []ProcedureCall {
	var out []ProcedureCall
	for _, call := range sqlparse.ExtractSQLCalls(sql) {
		raw := call.Name
		if strings.HasPrefix(raw, "@") || strings.EqualFold(raw, "sp_executesql") || strings.EqualFold(raw, "dbo.sp_executesql") {
			continue
		}
		object, resolution := catalog.resolveParts(raw, call.NameParts, "procedure")
		arguments := strings.TrimSpace(strings.TrimSuffix(strings.TrimSpace(sql[call.TargetEnd:call.StatementEnd]), ";"))
		line := sourceLineNumber(sql, call.Offset)
		out = append(out, ProcedureCall{Object: object, Arguments: arguments, Resolution: resolution, Source: compactSQL(sql[call.Offset:call.StatementEnd]), Guard: procedureGuardForLine(original, clean, line-1), Line: line, offset: call.Offset})
	}
	return out
}

func extractProcedureControlSteps(original, clean []string) []ProcedureControlStep {
	var out []ProcedureControlStep
	labelPattern := regexp.MustCompile(`(?i)^\s*([A-Za-z_][A-Za-z0-9_$#]*)\s*:\s*$`)
	for i, line := range clean {
		trim := strings.TrimSpace(line)
		upper := strings.ToUpper(trim)
		event := ""
		switch {
		case strings.HasPrefix(upper, "BEGIN TRANSACTION") || strings.HasPrefix(upper, "BEGIN TRAN"):
			event = "TRANSACTION START"
		case strings.HasPrefix(upper, "COMMIT TRANSACTION") || strings.HasPrefix(upper, "COMMIT TRAN"):
			event = "COMMIT"
		case strings.HasPrefix(upper, "ROLLBACK TRANSACTION") || strings.HasPrefix(upper, "ROLLBACK TRAN"):
			event = "ROLLBACK"
		case strings.HasPrefix(upper, "GOTO "):
			event = "GOTO"
		case strings.HasPrefix(upper, "RETURN"):
			event = "RETURN"
		case strings.HasPrefix(upper, "RAISERROR"):
			event = "RAISE ERROR"
		case containsSQLWord(upper, "THROW"):
			event = "THROW"
		case labelPattern.MatchString(trim):
			event = "LABEL"
		}
		if event != "" {
			out = append(out, ProcedureControlStep{Event: event, Evidence: compactSQL(strings.TrimSpace(lineAt(original, i))), Line: i + 1})
		}
	}
	return out
}

func procedureGuardForLine(original, clean []string, index int) string {
	if index <= 0 || index >= len(clean) {
		return ""
	}
	targetIndent := leadingSQLIndent(clean[index])
	var guards []string
	for i := index - 1; i >= 0 && index-i <= 120; i-- {
		trim := strings.TrimSpace(clean[i])
		if trim == "" {
			continue
		}
		indent := leadingSQLIndent(clean[i])
		upper := strings.ToUpper(trim)
		if strings.HasPrefix(upper, "ELSE") && indent < targetIndent {
			if prior := previousIfAtIndent(original, clean, i, indent); prior != "" {
				guards = append(guards, "ELSE of "+prior)
			} else {
				guards = append(guards, "ELSE branch")
			}
			targetIndent = indent
			continue
		}
		if strings.HasPrefix(upper, "IF ") || strings.HasPrefix(upper, "IF(") {
			if indent < targetIndent {
				guards = append(guards, procedureIfGuardText(original, i))
				targetIndent = indent
			}
		}
		if targetIndent <= 0 {
			break
		}
	}
	for left, right := 0, len(guards)-1; left < right; left, right = left+1, right-1 {
		guards[left], guards[right] = guards[right], guards[left]
	}
	return strings.Join(uniqueStrings(guards), " AND ")
}

func previousIfAtIndent(original, clean []string, before, indent int) string {
	for i := before - 1; i >= 0 && before-i <= 40; i-- {
		trim := strings.TrimSpace(clean[i])
		if trim == "" {
			continue
		}
		lineIndent := leadingSQLIndent(clean[i])
		upper := strings.ToUpper(trim)
		if lineIndent == indent && (strings.HasPrefix(upper, "IF ") || strings.HasPrefix(upper, "IF(")) {
			return procedureIfGuardText(original, i)
		}
		if lineIndent < indent {
			return ""
		}
	}
	return ""
}

func leadingSQLIndent(line string) int {
	indent := 0
	for _, r := range line {
		switch r {
		case ' ':
			indent++
		case '\t':
			indent += 2
		default:
			return indent
		}
	}
	return indent
}

func procedureIfGuardText(lines []string, start int) string {
	if start < 0 || start >= len(lines) {
		return "IF"
	}
	parts := []string{strings.TrimSpace(lines[start])}
	balance, opened := sqlParenBalance(lines[start])
	for i := start + 1; opened && balance > 0 && i < len(lines) && i-start <= 30; i++ {
		trim := strings.TrimSpace(lines[i])
		if trim != "" {
			parts = append(parts, trim)
		}
		delta, hasParen := sqlParenBalance(lines[i])
		balance += delta
		opened = opened || hasParen
	}
	return compactSQL(strings.Join(parts, " "))
}

func sqlParenBalance(line string) (int, bool) {
	balance := 0
	opened := false
	inString := false
	for i := 0; i < len(line); i++ {
		switch line[i] {
		case '\'':
			if inString && i+1 < len(line) && line[i+1] == '\'' {
				i++
				continue
			}
			inString = !inString
		case '(':
			if !inString {
				balance++
				opened = true
			}
		case ')':
			if !inString {
				balance--
			}
		}
	}
	return balance, opened
}

func extractRoutineDependencies(sql string, catalog procedureCatalog) ([]ProcedureDependency, []ProcedureDependency) {
	seenFunction, seenView := map[string]bool{}, map[string]bool{}
	var functions, views []ProcedureDependency
	for _, ref := range sqlparse.ExtractObjectReferences(sql) {
		switch ref.Kind {
		case sqlparse.ObjectReferenceFunction:
			object, resolution := catalog.resolve(ref.Name, "function")
			// The shared dependency scanner intentionally over-collects
			// unqualified name(...) expressions. Only a catalog-backed UDF is
			// safe to present as a function dependency here.
			if resolution == ResolutionStatic && !seenFunction[strings.ToLower(object)] {
				functions = append(functions, ProcedureDependency{object, resolution})
				seenFunction[strings.ToLower(object)] = true
			}
		case sqlparse.ObjectReferenceData:
			object, resolution := catalog.resolve(ref.Name, "data")
			if resolution == ResolutionStatic && catalog.isKind(object, "view") && !seenView[strings.ToLower(object)] {
				views = append(views, ProcedureDependency{object, resolution})
				seenView[strings.ToLower(object)] = true
			}
		}
	}
	sort.Slice(functions, func(i, j int) bool { return functions[i].Object < functions[j].Object })
	sort.Slice(views, func(i, j int) bool { return views[i].Object < views[j].Object })
	return functions, views
}

func extractProcedureDynamicSQL(sql string, lines []string, catalog procedureCatalog) []ProcedureDynamicSQL {
	tokens := sqlparse.SQLTokens(sql)
	calls := map[int]sqlparse.SQLCall{}
	for _, call := range sqlparse.ExtractSQLCalls(sql) {
		calls[call.Offset] = call
	}
	var out []ProcedureDynamicSQL
	for i, t := range tokens {
		if t.Quoted || (t.Upper != "EXEC" && t.Upper != "EXECUTE") || i+1 >= len(tokens) {
			continue
		}
		next := tokens[i+1]
		expression := ""
		literalAllowed := false
		if call, ok := calls[t.Start]; ok {
			if len(call.NameParts) == 0 || !strings.EqualFold(call.NameParts[len(call.NameParts)-1], "sp_executesql") {
				continue
			}
			expression = sql[call.TargetEnd:call.StatementEnd]
			literalAllowed = true
			// sp_executesql's first argument is the complete SQL expression; parameter
			// definitions/values after the first top-level comma are not part of it.
			parts := sqlparse.SplitTopLevelComma(expression)
			if len(parts) > 0 {
				expression = parts[0]
			}
			ts := sqlparse.SQLTokens(expression)
			if len(ts) >= 2 && strings.HasPrefix(ts[0].Text, "@") && ts[1].Text == "=" {
				expression = expression[ts[1].End:]
			}
		} else {
			if !next.Quoted && strings.HasPrefix(next.Text, "@") && i+2 < len(tokens) && tokens[i+2].Text == "=" {
				if i+3 >= len(tokens) || !strings.HasPrefix(tokens[i+3].Text, "@") {
					continue
				}
				next = tokens[i+3]
			}
			if next.Text == "(" {
				close, ok := sqlparse.FindMatchingParen(sql, next.Start)
				if !ok {
					expression = sql[next.End:]
				} else {
					expression = sql[next.End:close]
				}
				literalAllowed = ok
			} else if strings.HasPrefix(next.Text, "'") || next.Upper == "N" {
				stop := len(sql)
				for _, tail := range tokens[i+1:] {
					if tail.Depth == t.Depth && tail.Text == ";" {
						stop = tail.Start
						break
					}
				}
				expression = sql[next.Start:stop]
				literalAllowed = true
			} else if !next.Quoted && strings.HasPrefix(next.Text, "@") {
				expression = next.Text
			} else {
				continue
			}
		}
		line := sourceLineNumber(sql, t.Start)
		item := ProcedureDynamicSQL{Line: fmt.Sprintf("%d", line), Expression: compactSQL(strings.TrimSpace(expression)), Resolution: ResolutionUnresolved}
		if literal, ok := completeDynamicLiteral(expression); ok && literalAllowed {
			item.Expression = compactSQL(literal)
			item.RuntimeSource = "SQL literal"
			item.ResolvedObjects, item.Resolution = resolveDynamicLiteral(literal, catalog)
		} else {
			_, item.RuntimeSource = dynamicRuntimeEvidence(lines, line-1, item.Expression)
		}
		out = append(out, item)
	}
	return out
}

func completeDynamicLiteral(expression string) (string, bool) {
	tokens := sqlparse.SQLTokens(expression)
	if len(tokens) > 0 && tokens[len(tokens)-1].Text == ";" {
		tokens = tokens[:len(tokens)-1]
	}
	if len(tokens) == 2 && !tokens[0].Quoted && tokens[0].Upper == "N" {
		tokens = tokens[1:]
	}
	if len(tokens) != 1 || !strings.HasPrefix(tokens[0].Text, "'") || !strings.HasSuffix(tokens[0].Text, "'") || len(tokens[0].Text) < 2 {
		return "", false
	}
	raw := tokens[0].Text
	closed := false
	for i := 1; i < len(raw); i++ {
		if raw[i] != '\'' {
			continue
		}
		if i+1 < len(raw) && raw[i+1] == '\'' {
			i++
			continue
		}
		closed = i == len(raw)-1
		break
	}
	if !closed {
		return "", false
	}
	return strings.ReplaceAll(raw[1:len(raw)-1], "''", "'"), true
}

func resolveDynamicLiteral(sql string, catalog procedureCatalog) ([]string, string) {
	objects := []string{}
	resolution := ResolutionDynamic
	add := func(raw string, parts []string, kind string) {
		object, status := catalog.resolveParts(raw, parts, kind)
		objects = append(objects, object)
		if status != ResolutionStatic {
			resolution = status
		}
	}
	for _, ref := range sqlparse.ExtractSQLAccesses(sql) {
		if ref.NonPersisted {
			continue
		}
		if ref.Ambiguous || ref.CTECaseAmbiguous {
			objects = append(objects, ref.Name)
			resolution = ResolutionAmbiguous
			continue
		}
		kind := "data"
		if ref.Function {
			kind = "function"
		}
		add(ref.Name, ref.NameParts, kind)
	}
	for _, call := range sqlparse.ExtractSQLCalls(sql) {
		add(call.Name, call.NameParts, "procedure")
	}
	// Scalar function calls are not SQLAccess data targets.
	for _, ref := range sqlparse.ExtractObjectReferences(sql) {
		if ref.Kind == sqlparse.ObjectReferenceFunction {
			object, status := catalog.resolve(ref.Name, "function")
			objects = append(objects, object)
			if status != ResolutionStatic {
				resolution = status
			}
		}
	}
	if len(objects) == 0 {
		return nil, ResolutionUnresolved
	}
	return uniqueStrings(objects), resolution
}

func (c procedureCatalog) resolve(raw, expected string) (string, string) {
	raw = normalizeObjectToken(raw)
	return c.resolveParts(raw, strings.Split(raw, "."), expected)
}

func (c procedureCatalog) resolveParts(raw string, parts []string, expected string) (string, string) {
	if raw == "" {
		return raw, ResolutionNotFound
	}
	lastPart := parts[len(parts)-1]
	if strings.HasPrefix(lastPart, "#") || strings.HasPrefix(lastPart, "@") {
		return raw, ResolutionNotPersisted
	}
	if strings.EqualFold(parts[0], "sys") {
		return raw, ResolutionSystem
	}
	if len(parts) >= 3 {
		return raw, ResolutionExternal
	}
	kinds := []string{expected}
	if expected == "data" {
		kinds = []string{"table", "view"}
	}
	var matches []string
	for _, kind := range kinds {
		if len(parts) == 1 {
			for _, names := range c.byKind[kind] {
				for _, name := range names {
					separator := strings.Index(name, ".")
					if separator >= 0 && strings.EqualFold(name[separator+1:], raw) {
						matches = append(matches, name)
					}
				}
			}
		} else {
			matches = append(matches, c.byKind[kind][strings.ToLower(raw)]...)
		}
	}
	matches = uniqueStrings(matches)
	if len(matches) == 1 {
		return matches[0], ResolutionStatic
	}
	if len(matches) > 1 {
		return raw, ResolutionAmbiguous
	}
	if expected == "procedure" && (strings.HasPrefix(strings.ToLower(raw), "dbo.sp_") || strings.HasPrefix(strings.ToLower(raw), "sp_")) {
		return raw, ResolutionSystem
	}
	return raw, ResolutionNotFound
}

func (c procedureCatalog) isKind(object, kind string) bool {
	return len(c.byKind[kind][strings.ToLower(object)]) == 1
}

func extractProcedureWriteMappings(sql string, accesses []ProcedureDataAccess, catalog procedureCatalog) []ProcedureWriteMapping {
	var out []ProcedureWriteMapping
	for _, access := range accesses {
		if access.Resolution != ResolutionStatic {
			continue
		}
		if access.Operation == "INSERT" {
			out = append(out, extractInsertMappings(sql, access)...)
			continue
		}
		if access.Operation != "UPDATE" {
			continue
		}
		start, end := access.offset, access.statementEnd
		if start < 0 || start >= end || end > len(sql) {
			continue
		}
		chunk := sql[start:end]
		_, setEnd := sqlparse.FindTopLevelKeyword(chunk, "SET")
		if setEnd < 0 {
			continue
		}
		whereStart, whereEnd := sqlparse.FindTopLevelKeyword(chunk[setEnd:], "WHERE")
		assignmentEnd := len(chunk)
		filter := ""
		if whereStart >= 0 {
			assignmentEnd = setEnd + whereStart
			filter = firstSQLClause(strings.TrimSpace(chunk[setEnd+whereEnd:]))
		}
		if fromStart, _ := sqlparse.FindTopLevelKeyword(chunk[setEnd:assignmentEnd], "FROM"); fromStart >= 0 {
			assignmentEnd = setEnd + fromStart
		}
		for _, keyword := range []string{"OUTPUT", "OPTION"} {
			if pos, _ := sqlparse.FindTopLevelKeyword(chunk[setEnd:assignmentEnd], keyword); pos >= 0 {
				assignmentEnd = setEnd + pos
			}
		}
		assignments := chunk[setEnd:assignmentEnd]
		for _, assignment := range sqlparse.SplitTopLevelComma(assignments) {
			eq := topLevelEqual(assignment)
			if eq < 0 {
				continue
			}
			column := normalizeColumnToken(strings.TrimSpace(assignment[:eq]))
			expression := compactSQL(strings.TrimSuffix(strings.TrimSpace(assignment[eq+1:]), ";"))
			if column == "" || expression == "" {
				continue
			}
			out = append(out, ProcedureWriteMapping{Ref: access.Ref, Object: access.Object, Column: column, Expression: expression, Filter: compactSQL(filter), Resolution: access.Resolution, Line: access.Line})
		}
	}
	return out
}

func extractInsertMappings(sql string, access ProcedureDataAccess) []ProcedureWriteMapping {
	start, end := access.offset, access.statementEnd
	if start < 0 || start >= end || end > len(sql) {
		return nil
	}
	chunk := sql[start:end]
	afterTarget := access.targetEnd - start
	for afterTarget < len(chunk) && (chunk[afterTarget] == ' ' || chunk[afterTarget] == '\t' || chunk[afterTarget] == '\r' || chunk[afterTarget] == '\n') {
		afterTarget++
	}
	if afterTarget >= len(chunk) || chunk[afterTarget] != '(' {
		return nil
	}
	open := afterTarget
	close, ok := sqlparse.FindMatchingParen(chunk, open)
	if !ok {
		return nil
	}
	columns, valid := insertTargetColumns(sql, access)
	if !valid || len(columns) == 0 {
		return nil
	}
	tail := chunk[close+1:]
	selectStart, selectEnd := sqlparse.FindTopLevelKeyword(tail, "SELECT")
	valuesStart, valuesEnd := sqlparse.FindTopLevelKeyword(tail, "VALUES")
	var expressions []string
	if selectStart >= 0 && (valuesStart < 0 || selectStart < valuesStart) {
		after := tail[selectEnd:]
		// Projection modifiers and compound queries need a richer projection
		// model. Do not present them as exact positional expression mappings.
		if top, _ := sqlparse.FindTopLevelKeyword(after, "TOP"); top >= 0 {
			return nil
		}
		if distinct, _ := sqlparse.FindTopLevelKeyword(after, "DISTINCT"); distinct >= 0 {
			return nil
		}
		for _, word := range []string{"UNION", "EXCEPT", "INTERSECT"} {
			if pos, _ := sqlparse.FindTopLevelKeyword(after, word); pos >= 0 {
				return nil
			}
		}
		fromStart, _ := sqlparse.FindTopLevelKeyword(after, "FROM")
		if fromStart < 0 {
			return nil
		}
		expressions = sqlparse.SplitTopLevelComma(after[:fromStart])
	} else if valuesStart >= 0 {
		after := tail[valuesEnd:]
		valueOpen := strings.Index(after, "(")
		if valueOpen < 0 {
			return nil
		}
		valueClose, matched := sqlparse.FindMatchingParen(after, valueOpen)
		if !matched {
			return nil
		}
		if strings.HasPrefix(strings.TrimSpace(after[valueClose+1:]), ",") {
			return nil
		}
		expressions = sqlparse.SplitTopLevelComma(after[valueOpen+1 : valueClose])
	}
	if len(expressions) != len(columns) {
		return nil
	}
	out := make([]ProcedureWriteMapping, 0, len(columns))
	for i, column := range columns {
		out = append(out, ProcedureWriteMapping{Ref: access.Ref, Object: access.Object, Column: column, Expression: compactSQL(expressions[i]), Resolution: access.Resolution, Line: access.Line})
	}
	return out
}

func insertTargetColumns(sql string, access ProcedureDataAccess) ([]string, bool) {
	start, end := access.targetEnd, access.statementEnd
	if start < 0 || end > len(sql) || start >= end {
		return nil, false
	}
	tail := sql[start:end]
	tokens := sqlparse.SQLTokens(tail)
	if len(tokens) == 0 || tokens[0].Text != "(" {
		return nil, true
	} // No explicit column list: keep syntax-only access.
	close, ok := sqlparse.FindMatchingParen(tail, tokens[0].Start)
	if !ok {
		return nil, false
	}
	var columns []string
	for _, part := range sqlparse.SplitTopLevelComma(tail[tokens[0].End:close]) {
		items := sqlparse.SQLTokens(part)
		if len(items) != 1 || !items[0].Identifier {
			return nil, false
		}
		columns = append(columns, items[0].Text)
	}
	return columns, len(columns) > 0
}

func firstSQLClause(s string) string {
	lines := strings.Split(s, "\n")
	var kept []string
	for _, line := range lines {
		trim := strings.TrimSpace(line)
		upper := strings.ToUpper(trim)
		if trim == "" {
			if len(kept) > 0 {
				break
			}
			continue
		}
		if len(kept) > 0 && (strings.HasPrefix(upper, "IF(") || strings.HasPrefix(upper, "IF ") || strings.HasPrefix(upper, "END") || strings.HasPrefix(upper, "UPDATE ") || strings.HasPrefix(upper, "INSERT ") || strings.HasPrefix(upper, "EXEC ") || strings.HasPrefix(upper, "BEGIN ")) {
			break
		}
		kept = append(kept, trim)
	}
	return strings.Join(kept, " ")
}

func normalizeColumnToken(s string) string {
	tokens := sqlparse.SQLTokens(s)
	if len(tokens) == 0 {
		return ""
	}
	for i, t := range tokens {
		if i%2 == 0 {
			if !t.Identifier || strings.HasPrefix(t.Text, "@") {
				return ""
			}
		} else if t.Text != "." {
			return ""
		}
	}
	if len(tokens)%2 != 1 {
		return ""
	}
	return tokens[len(tokens)-1].Text
}

func validateWriteMappingColumns(mappings []ProcedureWriteMapping, catalog procedureCatalog) {
	for i := range mappings {
		table, ok := lookupDocumentedTable(catalog.tables, mappings[i].Object)
		if !ok || len(table.Columns) == 0 {
			continue
		} // Views and unavailable metadata remain syntax-only.
		matches := []string{}
		for _, c := range table.Columns {
			if strings.EqualFold(c.Name, mappings[i].Column) {
				matches = append(matches, c.Name)
			}
		}
		if len(matches) == 1 {
			mappings[i].Column = matches[0]
		} else if len(matches) == 0 {
			mappings[i].Resolution = ResolutionColumnNotFound
		} else {
			mappings[i].Resolution = ResolutionAmbiguous
		}
	}
}

func normalizeObjectToken(s string) string {
	s = strings.TrimSpace(strings.TrimRight(s, ";,"))
	parts := strings.Split(s, ".")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
		if len(parts[i]) >= 2 && strings.HasPrefix(parts[i], "[") && strings.HasSuffix(parts[i], "]") {
			parts[i] = strings.ReplaceAll(parts[i][1:len(parts[i])-1], "]]", "]")
		}
	}
	return strings.Join(parts, ".")
}

func isSQLPseudoObject(s string) bool {
	switch strings.ToUpper(strings.TrimSpace(s)) {
	case "SELECT", "INTO", "OPENXML", "OPENJSON", "STRING_SPLIT":
		return true
	default:
		return false
	}
}

func procedureExtractionNotices(p ProcedureProfile) []string {
	var out []string
	if p.HasTransaction {
		out = append(out, "Explicit transaction ditemukan; pertahankan boundary commit/rollback saat proses dipindahkan.")
	}
	if p.HasTryCatch || p.HasRaiseError || p.HasThrow {
		out = append(out, "Error contract ditemukan; review TRY/CATCH, RAISERROR/THROW, dan return code bersama-sama.")
	}
	if len(p.DynamicSQL) > 0 {
		out = append(out, "Dynamic SQL ditemukan; hanya object dengan resolution terbukti yang boleh dianggap dependency.")
	}
	for _, a := range p.Accesses {
		if a.Resolution != ResolutionStatic && a.Resolution != ResolutionNotPersisted {
			out = append(out, "Ada static reference yang belum ter-resolve ke db_object; review sebelum ekstraksi.")
			break
		}
	}
	for _, m := range p.Mappings {
		if m.Resolution == ResolutionColumnNotFound || m.Resolution == ResolutionAmbiguous {
			out = append(out, "Ada target column mapping yang tidak ditemukan atau ambigu pada snapshot table; expression dipertahankan sebagai evidence, bukan mapping terverifikasi.")
			break
		}
	}
	return uniqueStrings(out)
}

func applyCallerResolution(profiles []ProcedureProfile, catalog procedureCatalog) {
	byName := map[string]*ProcedureProfile{}
	for i := range profiles {
		byName[strings.ToLower(profiles[i].Name)] = &profiles[i]
	}
	for i := range profiles {
		for _, call := range profiles[i].Calls {
			callee := byName[strings.ToLower(call.Object)]
			if callee == nil {
				continue
			}
			args := sqlparse.SplitTopLevelComma(call.Arguments)
			for d := range callee.DynamicSQL {
				if callee.DynamicSQL[d].Resolution != ResolutionUnresolved && callee.DynamicSQL[d].Resolution != ResolutionFromCaller {
					continue
				}
				for paramIndex, param := range callee.Parameters {
					if paramIndex >= len(args) || (!strings.Contains(strings.ToLower(callee.DynamicSQL[d].Expression), strings.ToLower(param.Name)) && !strings.Contains(strings.ToLower(callee.DynamicSQL[d].RuntimeSource), strings.ToLower(param.Name))) {
						continue
					}
					literal := strings.TrimSpace(strings.TrimRight(args[paramIndex], ";"))
					if len(literal) < 2 || literal[0] != '\'' || literal[len(literal)-1] != '\'' {
						continue
					}
					candidate, status := catalog.resolve(strings.Trim(literal, "'"), "data")
					if status == ResolutionStatic {
						callee.DynamicSQL[d].Resolution = ResolutionFromCaller
						callee.DynamicSQL[d].ResolvedObjects = uniqueStrings(append(callee.DynamicSQL[d].ResolvedObjects, candidate))
						callerEvidence := "literal dari caller " + profiles[i].Name
						if !strings.Contains(callee.DynamicSQL[d].RuntimeSource, callerEvidence) {
							callee.DynamicSQL[d].RuntimeSource = strings.Trim(strings.TrimSpace(callee.DynamicSQL[d].RuntimeSource+"; "+callerEvidence), "; ")
						}
					}
				}
			}
		}
	}
}

func uniqueSortedMatches(re *regexp.Regexp, text string) []string {
	var out []string
	for _, m := range re.FindAllStringSubmatch(text, -1) {
		if len(m) > 1 {
			out = append(out, m[1])
		}
	}
	return uniqueStrings(out)
}

func dynamicLineNumber(value string) int {
	var n int
	_, _ = fmt.Sscanf(value, "%d", &n)
	return n
}
func variableNames(s string) []string {
	return uniqueStrings(regexp.MustCompile(`@[A-Za-z_][A-Za-z0-9_]*`).FindAllString(s, -1))
}

func dynamicRuntimeEvidence(lines []string, executionLine int, execution string) (string, string) {
	vars := variableNames(execution)
	expression := compactSQL(execution)
	allVars := append([]string{}, vars...)
	var definitions []string
	for _, variable := range vars {
		pattern := regexp.MustCompile(`(?i)\b(?:SET|SELECT)\s+` + regexp.QuoteMeta(variable) + `\s*=\s*(.+)$`)
		start := executionLine - 80
		if start < 0 {
			start = 0
		}
		for i := executionLine - 1; i >= start; i-- {
			match := pattern.FindStringSubmatch(lines[i])
			if len(match) == 0 {
				continue
			}
			definitions = append(definitions, "last assignment: "+compactSQL(variable+" = "+match[1]))
			allVars = append(allVars, variableNames(match[1])...)
			break
		}
	}
	runtimeSource := strings.Join(uniqueStrings(allVars), ", ")
	if len(definitions) > 0 {
		runtimeSource += "; " + strings.Join(definitions, "; ")
	}
	return expression, runtimeSource
}
func compactSQL(s string) string { return strings.Join(strings.Fields(s), " ") }
func lineAt(lines []string, index int) string {
	if index >= 0 && index < len(lines) {
		return lines[index]
	}
	return ""
}
