package sqlparse

import "strings"

// SQLAccess is a source-positioned object access found in a SQL module.
// NonPersisted identifies a scoped CTE or derived source. Catalog resolution
// remains the caller's responsibility; NameParts preserves quoted separators.
type SQLAccess struct {
	Operation        string
	Name             string
	Alias            string
	Function         bool
	Offset           int
	End              int
	StatementStart   int
	StatementEnd     int
	NonPersisted     bool
	QueryDepth       int
	Ambiguous        bool
	NameParts        []string
	CTECaseAmbiguous bool
}

type sourceToken struct {
	text, upper string
	start, end  int
	depth       int
	identifier  bool
	quoted      bool
}

type cteScope struct {
	name       string
	start, end int
}

// FindTopLevelKeyword ignores quoted identifiers, strings, comments, and nested
// expressions. Offsets refer to the original SQL; -1 means no match.
func FindTopLevelKeyword(sql, keyword string) (int, int) {
	for _, token := range sourceTokens(sql) {
		if token.depth == 0 && !token.quoted && token.upper == strings.ToUpper(keyword) {
			return token.start, token.end
		}
	}
	return -1, -1
}

// SQLCall describes an explicit static EXEC target. Variable/literal dynamic
// execution and EXECUTE AS security clauses are intentionally excluded.
type SQLCall struct {
	Name                            string
	NameParts                       []string
	Offset, TargetEnd, StatementEnd int
}

func ExtractSQLCalls(sql string) []SQLCall {
	tokens := sourceTokens(sql)
	var out []SQLCall
	for i, token := range tokens {
		if token.quoted || (token.upper != "EXEC" && token.upper != "EXECUTE") {
			continue
		}
		pos := i + 1
		if pos+1 < len(tokens) && !tokens[pos].quoted && strings.HasPrefix(tokens[pos].text, "@") && tokens[pos+1].text == "=" {
			pos += 2
		}
		name, parts, next, ok := sourceObjectName(tokens, pos)
		if !ok || strings.HasPrefix(parts[0], "@") {
			continue
		}
		_, end := sourceStatementBounds(tokens, i, len(sql))
		out = append(out, SQLCall{Name: name, NameParts: parts, Offset: token.start, TargetEnd: tokens[next-1].end, StatementEnd: end})
	}
	return out
}

// ExtractSQLAccesses scans statement tokens instead of physical lines. It is
// intentionally conservative: it reports syntactically explicit data targets
// and leaves object-kind and ambiguity decisions to the catalog-backed caller.
func ExtractSQLAccesses(sql string) []SQLAccess {
	tokens := sourceTokens(sql)
	scopes := sourceCTEScopes(tokens, len(sql))
	consumedReadTarget := map[int]bool{}
	fromList := map[int]bool{}
	out := make([]SQLAccess, 0)
	for i := 0; i < len(tokens); i++ {
		if tokens[i].quoted || (i > 0 && tokens[i-1].text == ".") {
			continue
		}
		depth := tokens[i].depth
		switch tokens[i].upper {
		case "SELECT", "WHERE", "GROUP", "ORDER", "HAVING", "SET", "UNION", "EXCEPT", "INTERSECT", ";":
			fromList[depth] = false
		}
		if tokens[i].text == ")" {
			for d := range fromList {
				if d > depth {
					delete(fromList, d)
				}
			}
		}
		op, nameAt := "", i+1
		switch tokens[i].upper {
		case "UPDATE":
			if sourceMergeAction(tokens, i) {
				continue
			}
			op = "UPDATE"
			nameAt = sourceSkipDMLTop(tokens, nameAt)
		case "INSERT":
			if sourceMergeAction(tokens, i) {
				continue
			}
			op = "INSERT"
			nameAt = sourceSkipDMLTop(tokens, nameAt)
			if nameAt < len(tokens) && !tokens[nameAt].quoted && tokens[nameAt].upper == "INTO" {
				nameAt++
			}
		case "DELETE":
			if sourceMergeAction(tokens, i) {
				continue
			}
			op = "DELETE"
			nameAt = sourceSkipDMLTop(tokens, nameAt)
			if nameAt < len(tokens) && !tokens[nameAt].quoted && tokens[nameAt].upper == "FROM" {
				consumedReadTarget[nameAt] = true
				nameAt++
			}
		case "MERGE":
			op = "MERGE"
			if nameAt < len(tokens) && !tokens[nameAt].quoted && tokens[nameAt].upper == "INTO" {
				nameAt++
			}
		case "FROM", "JOIN", "APPLY", "USING":
			if tokens[i].upper == "USING" && !sourceInMerge(tokens, i) {
				continue
			}
			if consumedReadTarget[i] || sourceFromIsCursor(tokens, i) {
				continue
			}
			fromList[depth] = true
			op = "READ"
		case ",":
			if !fromList[depth] {
				continue
			}
			op = "READ"
		case "INTO":
			// INSERT/MERGE INTO were already consumed as DML targets.
			head := sourceStatementHead(tokens, i)
			if head < 0 || (tokens[head].upper != "SELECT" && !sourceHasKeywordBetween(tokens, head, i, "OUTPUT")) {
				continue
			}
			op = "INSERT"
		default:
			continue
		}

		name, parts, next, ok := sourceObjectName(tokens, nameAt)
		derived := false
		if !ok && op == "READ" && nameAt < len(tokens) && tokens[nameAt].text == "(" {
			aliasAt := sourceSkipTokenParens(tokens, nameAt)
			if aliasAt < len(tokens) && !tokens[aliasAt].quoted && tokens[aliasAt].upper == "AS" {
				aliasAt++
			}
			name, parts, next, ok = sourceObjectName(tokens, aliasAt)
			derived = ok
		}
		if !ok {
			continue
		}
		start, end := sourceStatementBounds(tokens, i, len(sql))
		alias := sourceObjectAlias(tokens, next)
		out = append(out, SQLAccess{
			Operation:        op,
			Name:             name,
			Alias:            alias,
			Function:         op == "READ" && !derived && next < len(tokens) && tokens[next].text == "(",
			Offset:           tokens[i].start,
			End:              tokens[next-1].end,
			StatementStart:   start,
			StatementEnd:     end,
			NonPersisted:     derived || sourceNameIsScopedCTE(parts, tokens[i].start, scopes),
			QueryDepth:       tokens[i].depth,
			NameParts:        parts,
			CTECaseAmbiguous: sourceNameHasCaseAmbiguousCTE(parts, tokens[i].start, scopes),
		})
		if derived {
			out[len(out)-1].Alias = name
		}
	}
	// Resolve syntactically proven DML aliases before catalog lookup, so every
	// consumer (including migration ordering) sees the same physical target.
	for i := range out {
		if (out[i].Operation != "UPDATE" && out[i].Operation != "DELETE") || len(out[i].NameParts) != 1 {
			continue
		}
		var candidates []SQLAccess
		for _, candidate := range out {
			if candidate.Operation == "READ" && candidate.Alias == out[i].Name &&
				candidate.QueryDepth == out[i].QueryDepth && candidate.StatementStart == out[i].StatementStart && candidate.StatementEnd == out[i].StatementEnd {
				candidates = append(candidates, candidate)
			}
		}
		if len(candidates) == 1 {
			out[i].Name, out[i].NonPersisted = candidates[0].Name, candidates[0].NonPersisted
			out[i].NameParts = candidates[0].NameParts
			out[i].CTECaseAmbiguous = candidates[0].CTECaseAmbiguous
		} else if len(candidates) > 1 {
			out[i].Ambiguous = true
		}
	}
	return out
}

// HasClientResultSet identifies client SELECT and unsinked DML OUTPUT while
// excluding assignment, population, cursor, CTE-body, and nested SELECT forms.
func HasClientResultSet(sql string) bool {
	tokens := sourceTokens(sql)
	for i := 0; i < len(tokens); i++ {
		if !tokens[i].quoted && tokens[i].upper == "OUTPUT" && tokens[i].depth == 0 {
			head := sourceStatementHead(tokens, i)
			if head >= 0 {
				switch tokens[head].upper {
				case "INSERT", "UPDATE", "DELETE", "MERGE":
					end := sourceSelectEnd(tokens, i)
					if !sourceHasKeywordBetween(tokens, i, end, "INTO") {
						return true
					}
				}
			}
		}
		if tokens[i].quoted || tokens[i].upper != "SELECT" || tokens[i].depth != 0 || sourceSetContinuation(tokens, i) {
			continue
		}
		if sourceSelectBelongsToInsert(tokens, i) || sourceSelectIsCursor(tokens, i) {
			continue
		}
		end := sourceSelectEnd(tokens, i)
		hasInto, hasAssignment := false, false
		for pos := i + 1; pos < end; pos++ {
			if tokens[pos].depth != tokens[i].depth {
				continue
			}
			if tokens[pos].quoted {
				continue
			}
			switch tokens[pos].upper {
			case "FROM":
				pos = end
			case "INTO":
				hasInto = true
			case "=":
				if sourcePreviousProjectionIsVariable(tokens, i+1, pos) {
					hasAssignment = true
				}
			}
		}
		if !hasInto && !hasAssignment {
			return true
		}
	}
	return false
}

func sourceHasKeywordBetween(tokens []sourceToken, start, end int, keyword string) bool {
	for i := start + 1; i < end; i++ {
		if !tokens[i].quoted && tokens[i].depth == tokens[start].depth && tokens[i].upper == keyword {
			return true
		}
	}
	return false
}

func sourceTokens(sql string) []sourceToken {
	out := make([]sourceToken, 0)
	depth := 0
	for i := 0; i < len(sql); {
		if isSQLSpace(sql[i]) {
			i++
			continue
		}
		if sql[i] == '-' && i+1 < len(sql) && sql[i+1] == '-' {
			i += 2
			for i < len(sql) && sql[i] != '\n' && sql[i] != '\r' {
				i++
			}
			continue
		}
		if sql[i] == '/' && i+1 < len(sql) && sql[i+1] == '*' {
			i = skipNestedBlockComment(sql, i+2)
			continue
		}
		if sql[i] == '\'' {
			start := i
			i, _ = sourceSkipQuoted(sql, i, '\'', true)
			out = append(out, sourceToken{text: sql[start:i], start: start, end: i, depth: depth})
			continue
		}
		if sql[i] == '[' || sql[i] == '"' {
			start, quote := i, sql[i]
			close := byte(']')
			if quote == '"' {
				close = '"'
			}
			var closed bool
			i, closed = sourceSkipQuoted(sql, i, close, true)
			if !closed {
				break
			}
			value := sql[start+1 : i-1]
			value = strings.ReplaceAll(value, string(close)+string(close), string(close))
			out = append(out, sourceToken{text: value, upper: strings.ToUpper(value), start: start, end: i, depth: depth, identifier: true, quoted: true})
			continue
		}
		if isDependencyWordByte(sql[i]) || sql[i] >= 128 {
			start := i
			for i < len(sql) && (isDependencyWordByte(sql[i]) || sql[i] >= 128) {
				i++
			}
			value := sql[start:i]
			out = append(out, sourceToken{text: value, upper: strings.ToUpper(value), start: start, end: i, depth: depth, identifier: value[0] < '0' || value[0] > '9'})
			continue
		}
		start := i
		value := string(sql[i])
		if value == ")" && depth > 0 {
			depth--
		}
		out = append(out, sourceToken{text: value, upper: value, start: start, end: start + 1, depth: depth})
		i++
		if value == "(" {
			depth++
		}
	}
	return out
}

func sourceSkipQuoted(sql string, start int, close byte, escaped bool) (int, bool) {
	for i := start + 1; i < len(sql); i++ {
		if sql[i] != close {
			continue
		}
		if escaped && i+1 < len(sql) && sql[i+1] == close {
			i++
			continue
		}
		return i + 1, true
	}
	return len(sql), false
}

func sourceObjectName(tokens []sourceToken, start int) (string, []string, int, bool) {
	if start >= len(tokens) || !sourceObjectIdentifier(tokens[start]) {
		return "", nil, start, false
	}
	parts := []string{tokens[start].text}
	pos := start + 1
	for pos < len(tokens) && tokens[pos].text == "." && len(parts) < 4 {
		if pos+1 < len(tokens) && tokens[pos+1].text == "." {
			if pos+2 >= len(tokens) || !sourceObjectIdentifier(tokens[pos+2]) {
				return "", nil, pos, false
			}
			parts = append(parts, "", tokens[pos+2].text)
			pos += 3
			continue
		}
		if pos+1 >= len(tokens) || !sourceObjectIdentifier(tokens[pos+1]) {
			return "", nil, pos, false
		}
		parts = append(parts, tokens[pos+1].text)
		pos += 2
	}
	if pos < len(tokens) && tokens[pos].text == "." {
		return "", nil, pos, false
	}
	return strings.Join(parts, "."), parts, pos, true
}

func sourceObjectIdentifier(token sourceToken) bool {
	if !token.identifier {
		return false
	}
	if token.quoted {
		return true
	}
	switch token.upper {
	case "SELECT", "FROM", "JOIN", "WHERE", "ON", "INTO", "SET", "VALUES", "BEGIN", "END", "AS", "WITH", "OUTPUT", "OPTION", "WHEN", "THEN", "ELSE", "INNER", "LEFT", "RIGHT", "FULL", "CROSS", "OUTER", "APPLY", "GROUP", "ORDER", "HAVING", "UNION", "ALL", "EXCEPT", "INTERSECT", "USING", "TOP", "PERCENT", "FOR", "OFFSET", "FETCH", "UPDATE", "INSERT", "DELETE", "MERGE":
		return false
	default:
		return true
	}
}

func sourceObjectAlias(tokens []sourceToken, pos int) string {
	if pos < len(tokens) && tokens[pos].text == "(" {
		pos = sourceSkipTokenParens(tokens, pos)
	}
	if pos < len(tokens) && !tokens[pos].quoted && tokens[pos].upper == "AS" {
		pos++
	}
	if pos >= len(tokens) || !tokens[pos].identifier || strings.HasPrefix(tokens[pos].text, "@") || !sourceObjectIdentifier(tokens[pos]) {
		return ""
	}
	return tokens[pos].text
}

func sourceFromIsCursor(tokens []sourceToken, from int) bool {
	for pos := from - 1; pos >= 0; pos-- {
		if tokens[pos].text == ";" {
			break
		}
		if !tokens[pos].quoted && tokens[pos].upper == "FETCH" {
			return true
		}
		if sourceStatementStarter(tokens, pos) {
			break
		}
	}
	return false
}

func sourceStatementBounds(tokens []sourceToken, at, sqlLen int) (int, int) {
	depth := tokens[at].depth
	start, end := tokens[at].start, sqlLen
	head := sourceStatementHead(tokens, at)
	if head >= 0 {
		start = tokens[head].start
	} else {
		head = at
	}
	for pos := at + 1; pos < len(tokens); pos++ {
		if tokens[pos].depth == depth && sourceControlBoundary(tokens, head, pos) {
			end = tokens[pos].start
			break
		}
		if tokens[pos].depth < depth {
			end = tokens[pos].start
			break
		}
		if tokens[pos].depth == depth && tokens[pos].text == ";" {
			end = tokens[pos].end
			break
		}
		if tokens[pos].depth == depth && sourceStatementStarter(tokens, pos) && !sourceStatementContinuation(tokens, pos, head) {
			end = tokens[pos].start
			break
		}
	}
	return start, end
}

func sourceStatementStarter(tokens []sourceToken, pos int) bool {
	if tokens[pos].quoted || (pos > 0 && tokens[pos-1].text == ".") || sourceMergeAction(tokens, pos) {
		return false
	}
	switch tokens[pos].upper {
	case "SELECT", "INSERT", "UPDATE", "DELETE", "MERGE", "EXEC", "EXECUTE", "RETURN", "THROW", "RAISERROR":
		return true
	default:
		return false
	}
}

func sourceStatementContinuation(tokens []sourceToken, pos, current int) bool {
	if tokens[pos].upper != "SELECT" {
		return false
	}
	if sourceSetContinuation(tokens, pos) {
		return true
	}
	if tokens[current].upper != "INSERT" {
		return false
	}
	for i := current + 1; i < pos; i++ {
		if !tokens[i].quoted && tokens[i].upper == "SELECT" && tokens[i].depth == tokens[pos].depth {
			return false
		}
	}
	return true
}

func sourceCTEScopes(tokens []sourceToken, sqlLen int) []cteScope {
	var out []cteScope
	for i := 0; i < len(tokens); i++ {
		if tokens[i].quoted || tokens[i].upper != "WITH" {
			continue
		}
		pos := i + 1
		var names []string
		for pos < len(tokens) && sourceObjectIdentifier(tokens[pos]) {
			name := tokens[pos].text
			pos++
			if pos < len(tokens) && tokens[pos].text == "(" {
				pos = sourceSkipTokenParens(tokens, pos)
			}
			if pos >= len(tokens) || tokens[pos].upper != "AS" || pos+1 >= len(tokens) || tokens[pos+1].text != "(" {
				names = nil
				break
			}
			names = append(names, name)
			pos = sourceSkipTokenParens(tokens, pos+1)
			if pos >= len(tokens) || tokens[pos].text != "," {
				break
			}
			pos++
		}
		if len(names) == 0 || pos >= len(tokens) || !sourceStatementStarter(tokens, pos) {
			continue
		}
		_, end := sourceStatementBounds(tokens, pos, sqlLen)
		for _, name := range names {
			out = append(out, cteScope{name: name, start: tokens[i].start, end: end})
		}
	}
	return out
}

func sourceSkipTokenParens(tokens []sourceToken, start int) int {
	if start >= len(tokens) || tokens[start].text != "(" {
		return start
	}
	depth := tokens[start].depth
	for pos := start + 1; pos < len(tokens); pos++ {
		if tokens[pos].text == ")" && tokens[pos].depth == depth {
			return pos + 1
		}
	}
	return len(tokens)
}

func sourceNameIsScopedCTE(parts []string, offset int, scopes []cteScope) bool {
	if len(parts) != 1 {
		return false
	}
	for _, scope := range scopes {
		if offset >= scope.start && offset < scope.end && parts[0] == scope.name {
			return true
		}
	}
	return false
}

func sourceNameHasCaseAmbiguousCTE(parts []string, offset int, scopes []cteScope) bool {
	if len(parts) != 1 {
		return false
	}
	for _, scope := range scopes {
		if offset >= scope.start && offset < scope.end && parts[0] != scope.name && strings.EqualFold(parts[0], scope.name) {
			return true
		}
	}
	return false
}

func sourceSelectBelongsToInsert(tokens []sourceToken, selectAt int) bool {
	depth := tokens[selectAt].depth
	for pos := selectAt - 1; pos >= 0; pos-- {
		if tokens[pos].depth != depth {
			continue
		}
		if tokens[pos].text == ";" {
			return false
		}
		if sourceStatementStarter(tokens, pos) && !sourceSetContinuation(tokens, pos) {
			return tokens[pos].upper == "INSERT"
		}
	}
	return false
}

func sourceSelectEnd(tokens []sourceToken, selectAt int) int {
	depth := tokens[selectAt].depth
	for pos := selectAt + 1; pos < len(tokens); pos++ {
		if tokens[pos].depth == depth && sourceControlBoundary(tokens, selectAt, pos) {
			return pos
		}
		if tokens[pos].depth == depth && tokens[pos].text == ";" {
			return pos
		}
		if tokens[pos].depth == depth && sourceStatementStarter(tokens, pos) && !sourceStatementContinuation(tokens, pos, selectAt) {
			return pos
		}
	}
	return len(tokens)
}

func sourceControlBoundary(tokens []sourceToken, head, pos int) bool {
	if tokens[pos].quoted {
		return false
	}
	switch tokens[pos].upper {
	case "END", "ELSE", "BEGIN":
	default:
		return false
	}
	caseDepth := 0
	for i := head; i < pos; i++ {
		if tokens[i].quoted || tokens[i].depth != tokens[pos].depth {
			continue
		}
		if tokens[i].upper == "CASE" {
			caseDepth++
		}
		if tokens[i].upper == "END" && caseDepth > 0 {
			caseDepth--
		}
	}
	return caseDepth == 0
}

func sourcePreviousProjectionIsVariable(tokens []sourceToken, start, equals int) bool {
	return equals > start && !tokens[equals-1].quoted && strings.HasPrefix(tokens[equals-1].text, "@")
}

func sourceSkipDMLTop(tokens []sourceToken, pos int) int {
	if pos >= len(tokens) || tokens[pos].quoted || tokens[pos].upper != "TOP" {
		return pos
	}
	pos++
	if pos < len(tokens) && tokens[pos].text == "(" {
		pos = sourceSkipTokenParens(tokens, pos)
	} else if pos < len(tokens) {
		pos++
	}
	if pos < len(tokens) && !tokens[pos].quoted && tokens[pos].upper == "PERCENT" {
		pos++
	}
	return pos
}

func sourceMergeAction(tokens []sourceToken, pos int) bool {
	if tokens[pos].quoted {
		return false
	}
	switch tokens[pos].upper {
	case "UPDATE", "INSERT", "DELETE":
	default:
		return false
	}
	return pos > 0 && !tokens[pos-1].quoted && tokens[pos-1].upper == "THEN" && sourceInMerge(tokens, pos)
}

func sourceInMerge(tokens []sourceToken, at int) bool {
	for pos := at - 1; pos >= 0; pos-- {
		if tokens[pos].depth < tokens[at].depth || tokens[pos].text == ";" {
			break
		}
		if tokens[pos].depth != tokens[at].depth || tokens[pos].quoted {
			continue
		}
		if tokens[pos].upper == "MERGE" {
			return true
		}
		if tokens[pos].upper == "SELECT" {
			break
		}
	}
	return false
}

func sourceSetContinuation(tokens []sourceToken, pos int) bool {
	if tokens[pos].quoted || tokens[pos].upper != "SELECT" {
		return false
	}
	pos--
	if pos >= 0 && !tokens[pos].quoted && tokens[pos].upper == "ALL" {
		pos--
	}
	return pos >= 0 && !tokens[pos].quoted && (tokens[pos].upper == "UNION" || tokens[pos].upper == "EXCEPT" || tokens[pos].upper == "INTERSECT")
}

func sourceStatementHead(tokens []sourceToken, at int) int {
	for pos := at; pos >= 0; pos-- {
		if tokens[pos].depth < tokens[at].depth || tokens[pos].text == ";" {
			break
		}
		if tokens[pos].depth != tokens[at].depth || !sourceStatementStarter(tokens, pos) || sourceSetContinuation(tokens, pos) {
			continue
		}
		if tokens[pos].upper == "SELECT" && sourceSelectBelongsToInsert(tokens, pos) {
			continue
		}
		return pos
	}
	return -1
}

func sourceSelectIsCursor(tokens []sourceToken, at int) bool {
	for pos := at - 1; pos >= 0; pos-- {
		if tokens[pos].text == ";" || sourceStatementStarter(tokens, pos) {
			break
		}
		if !tokens[pos].quoted && tokens[pos].upper == "CURSOR" {
			return true
		}
	}
	return false
}
