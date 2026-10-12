package sqlparse

import "strings"

// CanonicalSemanticSQL returns a comment/whitespace-insensitive token stream.
// SQL keywords are case-insensitive, while identifiers and string literals
// retain their exact spelling. Bracketed and bare identifiers share the same
// representation so [dbo].[Thing] and dbo.Thing compare equally without making
// Thing and thing equivalent.
func CanonicalSemanticSQL(input string) string {
	var tokens []string
	for i := 0; i < len(input); {
		ch := input[i]
		if isSQLSpace(ch) {
			i++
			continue
		}
		if ch == '-' && i+1 < len(input) && input[i+1] == '-' {
			i += 2
			for i < len(input) && input[i] != '\n' && input[i] != '\r' {
				i++
			}
			continue
		}
		if ch == '/' && i+1 < len(input) && input[i+1] == '*' {
			i = skipNestedBlockComment(input, i+2)
			continue
		}
		if ch == '\'' {
			start := i
			i++
			for i < len(input) {
				if input[i] != '\'' {
					i++
					continue
				}
				if i+1 < len(input) && input[i+1] == '\'' {
					i += 2
					continue
				}
				i++
				break
			}
			tokens = append(tokens, "S:"+input[start:i])
			continue
		}
		if ch == '[' {
			var value strings.Builder
			i++
			for i < len(input) {
				if input[i] != ']' {
					value.WriteByte(input[i])
					i++
					continue
				}
				if i+1 < len(input) && input[i+1] == ']' {
					value.WriteByte(']')
					i += 2
					continue
				}
				i++
				break
			}
			tokens = append(tokens, "I:"+value.String())
			continue
		}
		if ch == '"' {
			var value strings.Builder
			i++
			for i < len(input) {
				if input[i] != '"' {
					value.WriteByte(input[i])
					i++
					continue
				}
				if i+1 < len(input) && input[i+1] == '"' {
					value.WriteByte('"')
					i += 2
					continue
				}
				i++
				break
			}
			tokens = append(tokens, "I:"+value.String())
			continue
		}
		if isSemanticWordByte(ch) {
			start := i
			for i < len(input) && isSemanticWordByte(input[i]) {
				i++
			}
			word := input[start:i]
			upper := strings.ToUpper(word)
			if _, ok := semanticSQLKeywords[upper]; ok {
				tokens = append(tokens, "K:"+upper)
			} else {
				tokens = append(tokens, "I:"+word)
			}
			continue
		}

		// Keep operators and punctuation as distinct tokens. A trailing statement
		// terminator does not alter object semantics.
		if ch != ';' {
			tokens = append(tokens, "P:"+string(ch))
		}
		i++
	}
	return strings.Join(tokens, " ")
}

func isSQLSpace(ch byte) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' || ch == '\f'
}

func isSemanticWordByte(ch byte) bool {
	return isIdentifierByte(ch) || (ch >= '0' && ch <= '9') || ch == '@' || ch == '#' || ch == '$'
}

func skipNestedBlockComment(input string, i int) int {
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
	return i
}

var semanticSQLKeywords = map[string]struct{}{
	"ADD": {}, "ALL": {}, "ALTER": {}, "AND": {}, "AS": {}, "ASC": {}, "BEGIN": {},
	"BY": {}, "CASE": {}, "CHECK": {}, "CLUSTERED": {}, "COLLATE": {}, "COLUMN": {},
	"CONSTRAINT": {}, "CREATE": {}, "DEFAULT": {}, "DELETE": {}, "DESC": {}, "DISTINCT": {},
	"DROP": {}, "ELSE": {}, "END": {}, "EXEC": {}, "EXECUTE": {}, "EXISTS": {}, "FILTER": {},
	"FOR": {}, "FOREIGN": {}, "FROM": {}, "FULL": {}, "FUNCTION": {}, "GO": {}, "GROUP": {},
	"HAVING": {}, "IF": {}, "IN": {}, "INCLUDE": {}, "INDEX": {}, "INNER": {}, "INSERT": {},
	"INTO": {}, "IS": {}, "JOIN": {}, "KEY": {}, "LEFT": {}, "MERGE": {}, "N": {}, "NO": {},
	"NONCLUSTERED": {}, "NOT": {}, "NULL": {}, "ON": {}, "OR": {}, "ORDER": {}, "OUTER": {},
	"PRIMARY": {}, "PROCEDURE": {}, "REFERENCES": {}, "RETURN": {}, "RETURNS": {}, "RIGHT": {},
	"SCHEMABINDING": {}, "SELECT": {}, "SET": {}, "TABLE": {}, "THEN": {}, "THROW": {}, "TOP": {},
	"TRIGGER": {}, "UNION": {}, "UNIQUE": {}, "UPDATE": {}, "VALUES": {}, "VIEW": {}, "WHEN": {},
	"WHERE": {}, "WITH": {},
	// Built-in types are case-insensitive language tokens. User-defined type
	// names are intentionally absent and therefore remain exact-case.
	"BIGINT": {}, "BINARY": {}, "BIT": {}, "CHAR": {}, "DATE": {}, "DATETIME": {},
	"DATETIME2": {}, "DATETIMEOFFSET": {}, "DECIMAL": {}, "FLOAT": {}, "GEOGRAPHY": {},
	"GEOMETRY": {}, "HIERARCHYID": {}, "IMAGE": {}, "INT": {}, "MONEY": {}, "NCHAR": {},
	"NTEXT": {}, "NUMERIC": {}, "NVARCHAR": {}, "REAL": {}, "ROWVERSION": {}, "SMALLDATETIME": {},
	"SMALLINT": {}, "SMALLMONEY": {}, "SQL_VARIANT": {}, "TEXT": {}, "TIME": {}, "TIMESTAMP": {},
	"TINYINT": {}, "UNIQUEIDENTIFIER": {}, "VARBINARY": {}, "VARCHAR": {}, "XML": {},
}
