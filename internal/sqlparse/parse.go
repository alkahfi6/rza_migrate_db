package sqlparse

import (
	"strings"
	"unicode"

	"rza_migrate_db/internal/sqlfmt"
)

var defaultForbiddenSQLTokens = []string{
	"GO",
	"USE",
	"CREATE DATABASE",
	"ALTER DATABASE",
	"CREATE LOGIN",
	"ALTER LOGIN",
	"CREATE USER",
}

// CompactUpperSQL normalizes SQL then strips all whitespace for compact matching.
func CompactUpperSQL(input string) string {
	normalized := strings.ToUpper(sqlfmt.NormalizeSQLText(input, true, true))
	var b strings.Builder
	b.Grow(len(normalized))
	for _, ch := range normalized {
		switch ch {
		case ' ', '\t', '\n', '\r':
			continue
		default:
			b.WriteRune(ch)
		}
	}
	return b.String()
}

// ExtractNamedNVarcharParamLiterals extracts NVARCHAR literals assigned via named parameters.
func ExtractNamedNVarcharParamLiterals(upperSQL, param string) []string {
	out := make([]string, 0)
	cursor := 0
	for {
		valueStart, nextCursor, ok := findNamedParamAssignmentValueStart(upperSQL, param, cursor)
		if !ok {
			break
		}
		if value, _, ok := readNVarcharLiteral(upperSQL, valueStart); ok {
			out = append(out, value)
		}
		cursor = nextCursor
	}
	return out
}

// ExtractDeclaredNVarcharLiteral extracts a DECLARE assignment while matching
// SQL keywords case-insensitively and preserving the literal's exact spelling.
func ExtractDeclaredNVarcharLiteral(sqlText, varName string) (string, bool) {
	upperSQL := strings.ToUpper(sqlText)
	cursor := 0
	anchor := "DECLARE " + strings.ToUpper(varName)
	for {
		rel := strings.Index(upperSQL[cursor:], anchor)
		if rel < 0 {
			return "", false
		}
		start := cursor + rel + len(anchor)
		eq := strings.Index(upperSQL[start:], "=")
		if eq < 0 {
			cursor = start
			continue
		}
		valueStart := start + eq + 1
		if value, _, ok := readNVarcharLiteral(sqlText, valueStart); ok {
			return value, true
		}
		cursor = start
	}
}

// ExtractDeclaredIntegerLiteral extracts a DECLARE integer assignment using a
// case-insensitive variable declaration search.
func ExtractDeclaredIntegerLiteral(sqlText, varName string) (string, bool) {
	upperSQL := strings.ToUpper(sqlText)
	cursor := 0
	anchor := "DECLARE " + strings.ToUpper(varName)
	for {
		rel := strings.Index(upperSQL[cursor:], anchor)
		if rel < 0 {
			return "", false
		}
		start := cursor + rel + len(anchor)
		eq := strings.Index(upperSQL[start:], "=")
		if eq < 0 {
			cursor = start
			continue
		}
		i := start + eq + 1
		for i < len(upperSQL) && (upperSQL[i] == ' ' || upperSQL[i] == '\t' || upperSQL[i] == '\n' || upperSQL[i] == '\r') {
			i++
		}
		j := i
		for j < len(upperSQL) && upperSQL[j] >= '0' && upperSQL[j] <= '9' {
			j++
		}
		if j > i {
			return upperSQL[i:j], true
		}
		cursor = start
	}
}

// ExtractNamedIntParamLiterals extracts integer literals assigned via named parameters.
func ExtractNamedIntParamLiterals(upperSQL, param string) []int {
	out := make([]int, 0)
	cursor := 0
	for {
		valueStart, nextCursor, ok := findNamedParamAssignmentValueStart(upperSQL, param, cursor)
		if !ok {
			break
		}
		i := valueStart
		for i < len(upperSQL) && (upperSQL[i] == ' ' || upperSQL[i] == '\t' || upperSQL[i] == '\n' || upperSQL[i] == '\r') {
			i++
		}
		if i >= len(upperSQL) || (upperSQL[i] < '0' || upperSQL[i] > '9') {
			cursor = nextCursor
			continue
		}

		value := 0
		for i < len(upperSQL) {
			ch := upperSQL[i]
			if ch < '0' || ch > '9' {
				break
			}
			value = value*10 + int(ch-'0')
			i++
		}
		out = append(out, value)
		cursor = nextCursor
	}
	return out
}

// DetectForbiddenTokens scans SQL for restricted tokens while ignoring comments/strings.
func DetectForbiddenTokens(sql string) []string {
	return DetectForbiddenTokensWithList(sql, nil)
}

// DetectForbiddenTokensWithList scans SQL for restricted tokens from caller-supplied list.
// Nil/empty list falls back to engine default token set.
func DetectForbiddenTokensWithList(sql string, tokens []string) []string {
	orderedTokens, hasGO := normalizeForbiddenTokenList(tokens)
	upper := strings.ToUpper(StripSQLStringsAndComments(sql))
	seen := map[string]bool{}
	out := []string{}

	if hasGO {
		lines := strings.Split(strings.ReplaceAll(strings.ReplaceAll(upper, "\r\n", "\n"), "\r", "\n"), "\n")
		for _, line := range lines {
			if strings.TrimSpace(line) == "GO" && !seen["GO"] {
				seen["GO"] = true
				out = append(out, "GO")
			}
		}
	}

	for _, token := range orderedTokens {
		if token == "GO" {
			continue
		}

		matched := false
		switch token {
		case "USE":
			matched = ContainsStandaloneWord(upper, token)
		default:
			matched = strings.Contains(upper, token)
		}
		if matched && !seen[token] {
			seen[token] = true
			out = append(out, token)
		}
	}
	return out
}

func normalizeForbiddenTokenList(tokens []string) ([]string, bool) {
	source := tokens
	if len(source) == 0 {
		source = defaultForbiddenSQLTokens
	}

	unique := make(map[string]struct{}, len(source))
	ordered := make([]string, 0, len(source))
	for _, raw := range source {
		token := strings.ToUpper(strings.TrimSpace(raw))
		if token == "" {
			continue
		}
		if _, exists := unique[token]; exists {
			continue
		}
		unique[token] = struct{}{}
		ordered = append(ordered, token)
	}
	hasGO := false
	for _, token := range ordered {
		if token == "GO" {
			hasGO = true
			break
		}
	}
	return ordered, hasGO
}

// FindStatementEnd returns the end position (exclusive) of the first statement starting at `start`.
// Semicolons inside string literals/comments are ignored.
func FindStatementEnd(sql string, start int) int {
	if start < 0 {
		start = 0
	}
	if start >= len(sql) {
		return len(sql)
	}

	const (
		stateNormal = iota
		stateSingleQuotedString
		stateLineComment
		stateBlockComment
	)

	state := stateNormal
	for i := start; i < len(sql); i++ {
		ch := sql[i]

		switch state {
		case stateNormal:
			if ch == '\'' {
				state = stateSingleQuotedString
				continue
			}
			if ch == '-' && i+1 < len(sql) && sql[i+1] == '-' {
				state = stateLineComment
				i++
				continue
			}
			if ch == '/' && i+1 < len(sql) && sql[i+1] == '*' {
				state = stateBlockComment
				i++
				continue
			}
			if ch == ';' {
				return i + 1
			}

		case stateSingleQuotedString:
			if ch == '\'' {
				if i+1 < len(sql) && sql[i+1] == '\'' {
					i++
					continue
				}
				state = stateNormal
			}

		case stateLineComment:
			if ch == '\n' {
				state = stateNormal
			}

		case stateBlockComment:
			if ch == '*' && i+1 < len(sql) && sql[i+1] == '/' {
				state = stateNormal
				i++
			}
		}
	}
	return len(sql)
}

// StripSQLStringsAndComments strips SQL comments and string literals.
func StripSQLStringsAndComments(input string) string {
	sql := strings.ReplaceAll(strings.ReplaceAll(input, "\r\n", "\n"), "\r", "\n")
	var b strings.Builder

	const (
		stateNormal = iota
		stateSingleQuotedString
		stateLineComment
		stateBlockComment
	)

	state := stateNormal
	for i := 0; i < len(sql); i++ {
		ch := sql[i]

		switch state {
		case stateNormal:
			if ch == '\'' {
				b.WriteByte(' ')
				state = stateSingleQuotedString
				continue
			}
			if ch == '-' && i+1 < len(sql) && sql[i+1] == '-' {
				b.WriteByte(' ')
				b.WriteByte(' ')
				state = stateLineComment
				i++
				continue
			}
			if ch == '/' && i+1 < len(sql) && sql[i+1] == '*' {
				b.WriteByte(' ')
				b.WriteByte(' ')
				state = stateBlockComment
				i++
				continue
			}
			b.WriteByte(ch)

		case stateSingleQuotedString:
			if ch == '\'' {
				if i+1 < len(sql) && sql[i+1] == '\'' {
					b.WriteByte(' ')
					b.WriteByte(' ')
					i++
					continue
				}
				b.WriteByte(' ')
				state = stateNormal
				continue
			}
			if ch == '\n' {
				b.WriteByte('\n')
			} else {
				b.WriteByte(' ')
			}

		case stateLineComment:
			if ch == '\n' {
				b.WriteByte('\n')
				state = stateNormal
			} else {
				b.WriteByte(' ')
			}

		case stateBlockComment:
			if ch == '*' && i+1 < len(sql) && sql[i+1] == '/' {
				b.WriteByte(' ')
				b.WriteByte(' ')
				state = stateNormal
				i++
			} else if ch == '\n' {
				b.WriteByte('\n')
			} else {
				b.WriteByte(' ')
			}
		}
	}
	return b.String()
}

// ContainsStandaloneWord checks case-matching standalone token with identifier boundaries.
func ContainsStandaloneWord(inputUpper, wordUpper string) bool {
	wordLen := len(wordUpper)
	if wordLen == 0 || len(inputUpper) < wordLen {
		return false
	}
	for i := 0; i <= len(inputUpper)-wordLen; i++ {
		if inputUpper[i:i+wordLen] != wordUpper {
			continue
		}

		leftOK := i == 0 || !isWordChar(rune(inputUpper[i-1]))
		rightIdx := i + wordLen
		rightOK := rightIdx >= len(inputUpper) || !isWordChar(rune(inputUpper[rightIdx]))
		if leftOK && rightOK {
			return true
		}
	}
	return false
}

func findNamedParamAssignmentValueStart(upperSQL, param string, cursor int) (int, int, bool) {
	for {
		rel := strings.Index(upperSQL[cursor:], param)
		if rel < 0 {
			return 0, cursor, false
		}
		pos := cursor + rel
		end := pos + len(param)
		cursor = end

		if pos > 0 && isSQLIdentifierChar(upperSQL[pos-1]) {
			continue
		}
		if end < len(upperSQL) && isSQLIdentifierChar(upperSQL[end]) {
			continue
		}

		i := end
		for i < len(upperSQL) && (upperSQL[i] == ' ' || upperSQL[i] == '\t' || upperSQL[i] == '\n' || upperSQL[i] == '\r') {
			i++
		}
		if i >= len(upperSQL) || upperSQL[i] != '=' {
			continue
		}
		return i + 1, cursor, true
	}
}

func readNVarcharLiteral(upperSQL string, start int) (string, int, bool) {
	i := start
	for i < len(upperSQL) && (upperSQL[i] == ' ' || upperSQL[i] == '\t' || upperSQL[i] == '\n' || upperSQL[i] == '\r') {
		i++
	}
	if i >= len(upperSQL) {
		return "", i, false
	}

	if upperSQL[i] == 'N' || upperSQL[i] == 'n' {
		i++
	}
	if i >= len(upperSQL) || upperSQL[i] != '\'' {
		return "", i, false
	}
	i++

	var b strings.Builder
	for i < len(upperSQL) {
		ch := upperSQL[i]
		if ch == '\'' {
			if i+1 < len(upperSQL) && upperSQL[i+1] == '\'' {
				b.WriteByte('\'')
				i += 2
				continue
			}
			return b.String(), i + 1, true
		}
		b.WriteByte(ch)
		i++
	}
	return "", i, false
}

func isSQLIdentifierChar(ch byte) bool {
	return (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '@'
}

func isWordChar(ch rune) bool {
	return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ch == '_'
}
