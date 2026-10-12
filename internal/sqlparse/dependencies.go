package sqlparse

import "strings"

// ObjectReferenceKind identifies the SQL context that referenced an object.
// It lets callers resolve a name only against object kinds valid in that context.
type ObjectReferenceKind string

const (
	ObjectReferenceData      ObjectReferenceKind = "data"
	ObjectReferenceFunction  ObjectReferenceKind = "function"
	ObjectReferenceProcedure ObjectReferenceKind = "procedure"
	ObjectReferenceType      ObjectReferenceKind = "type"
)

// ObjectReference is a schema-qualified object reference found outside comments
// and string literals. Name is normalized to an unbracketed schema.object form.
type ObjectReference struct {
	Name string
	Kind ObjectReferenceKind
}

// ExtractObjectReferences finds explicit, schema-qualified object references
// used by SQL Server modules. Dynamic SQL is intentionally ignored because its
// target cannot be resolved deterministically from source text.
func ExtractObjectReferences(sql string) []ObjectReference {
	sourced := sourceTokens(sql)
	tokens := make([]string, len(sourced))
	for i, token := range sourced {
		tokens[i] = token.text
		if token.quoted {
			tokens[i] = sql[token.start:token.end]
		}
	}
	seen := make(map[string]struct{})
	out := make([]ObjectReference, 0)
	add := func(name string, kind ObjectReferenceKind) {
		if name == "" {
			return
		}
		key := string(kind) + "|" + name
		if _, exists := seen[key]; exists {
			return
		}
		seen[key] = struct{}{}
		out = append(out, ObjectReference{Name: name, Kind: kind})
	}
	nonRoutineEnds := map[int]bool{}
	for _, access := range ExtractSQLAccesses(sql) {
		if access.Operation != "READ" || access.NonPersisted {
			nonRoutineEnds[access.End] = true
		}
		if access.NonPersisted || access.Ambiguous {
			continue
		}
		name, ok := dependencyCatalogParts(access.NameParts)
		if !ok {
			continue
		}
		kind := ObjectReferenceData
		if access.Function {
			kind = ObjectReferenceFunction
		}
		add(name, kind)
	}
	for _, call := range ExtractSQLCalls(sql) {
		if name, ok := dependencyCatalogParts(call.NameParts); ok {
			add(name, ObjectReferenceProcedure)
		}
	}

	for i := 0; i < len(tokens); i++ {
		word := strings.ToUpper(tokens[i])
		switch word {
		case "REFERENCES":
			if name, next, ok := dependencyDataName(tokens, i+1); ok {
				kind := ObjectReferenceData
				// In a FOREIGN KEY clause, the parenthesis after REFERENCES
				// starts the referenced column list and must remain a table
				// dependency, including the compact form dbo.Parent([Id]).
				_ = next
				add(name, kind)
			}
		}

		if dependencyRoutineCallAllowed(tokens, i) {
			if name, next, ok := dependencyRoutineName(tokens, i); ok && next < len(tokens) && tokens[next] == "(" && !nonRoutineEnds[sourced[next-1].end] && !dependencyCTEColumnList(sourced, next) {
				add(name, ObjectReferenceFunction)
			}
		}

		// Table-valued parameters are declared as
		// "@items dbo.OrderItemType READONLY". Scalar alias types can use the
		// same qualified form without READONLY, so a qualified name immediately
		// following a parameter token is always a type dependency.
		if strings.HasPrefix(tokens[i], "@") {
			if name, ok := dependencyParameterType(tokens, i+1); ok {
				add(name, ObjectReferenceType)
			}
		}
		if strings.EqualFold(tokens[i], "RETURNS") {
			if name, _, ok := dependencyQualifiedName(tokens, i+1); ok {
				add(name, ObjectReferenceType)
			}
		}
	}
	return out
}

func dependencyCatalogParts(parts []string) (string, bool) {
	if len(parts) == 0 {
		return "", false
	}
	if len(parts) >= 3 {
		return "", false
	}
	if strings.HasPrefix(parts[len(parts)-1], "@") || strings.HasPrefix(parts[len(parts)-1], "#") {
		return "", false
	}
	if len(parts) == 1 {
		name := parts[0]
		if strings.HasPrefix(name, "@") {
			return "", false
		}
		return "dbo." + name, true
	}
	return strings.Join(parts, "."), true
}

func dependencyCTEColumnList(tokens []sourceToken, parens int) bool {
	next := sourceSkipTokenParens(tokens, parens)
	return next < len(tokens) && !tokens[next].quoted && tokens[next].upper == "AS"
}

func dependencyParameterType(tokens []string, start int) (string, bool) {
	if name, _, ok := dependencyQualifiedName(tokens, start); ok {
		return name, true
	}
	// An unqualified user-defined table type is only unambiguous when followed
	// by READONLY. Built-in scalar parameter types do not use READONLY.
	if start+1 >= len(tokens) || !isDependencyIdentifier(tokens[start]) || !strings.EqualFold(tokens[start+1], "READONLY") {
		return "", false
	}
	return "dbo." + normalizeDependencyIdentifier(tokens[start]), true
}

func dependencyRoutineCallAllowed(tokens []string, start int) bool {
	if start > 0 && tokens[start-1] == "." {
		// The schema-qualified name is evaluated from its schema token. Do not
		// also treat the object component as an unqualified function.
		return false
	}
	if start == 0 {
		return true
	}
	switch strings.ToUpper(tokens[start-1]) {
	case "TABLE", "REFERENCES", "CONSTRAINT", "KEY", "FUNCTION", "PROCEDURE", "PROC", "TYPE", "INDEX":
		// Parentheses in these contexts contain table definitions, referenced
		// columns, or key columns rather than function arguments.
		return false
	default:
		return true
	}
}

func dependencyQualifiedName(tokens []string, start int) (string, int, bool) {
	if start+2 >= len(tokens) || !isDependencyIdentifier(tokens[start]) || tokens[start+1] != "." || !isDependencyIdentifier(tokens[start+2]) {
		return "", start, false
	}
	if start+3 < len(tokens) && tokens[start+3] == "." {
		// Cross-database/server references are outside this repository's dbo graph.
		return "", start, false
	}
	return normalizeDependencyIdentifier(tokens[start]) + "." + normalizeDependencyIdentifier(tokens[start+2]), start + 3, true
}

func dependencyDataName(tokens []string, start int) (string, int, bool) {
	if name, next, ok := dependencyQualifiedName(tokens, start); ok {
		return name, next, true
	}
	if start+1 < len(tokens) && tokens[start+1] == "." {
		return "", start, false
	}
	if start >= len(tokens) || !isDependencyIdentifier(tokens[start]) || strings.HasPrefix(tokens[start], "@") {
		return "", start, false
	}
	return "dbo." + normalizeDependencyIdentifier(tokens[start]), start + 1, true
}

func dependencyRoutineName(tokens []string, start int) (string, int, bool) {
	if name, next, ok := dependencyQualifiedName(tokens, start); ok {
		return name, next, true
	}
	if start+1 < len(tokens) && tokens[start+1] == "." {
		return "", start, false
	}
	if start >= len(tokens) || !isDependencyIdentifier(tokens[start]) || strings.HasPrefix(tokens[start], "@") || isDependencyKeyword(tokens[start]) {
		return "", start, false
	}
	return "dbo." + normalizeDependencyIdentifier(tokens[start]), start + 1, true
}

// Keywords can be followed by parentheses in valid SQL (notably CTE `AS (`).
// They are never unqualified routine names and must not create false graph
// edges such as dbo.AS.
func isDependencyKeyword(token string) bool {
	switch strings.ToUpper(strings.TrimSpace(token)) {
	case "AS", "SELECT", "FROM", "JOIN", "WHERE", "ON", "WITH", "INTO", "UPDATE", "INSERT", "DELETE", "MERGE", "VALUES", "CASE", "WHEN", "THEN", "ELSE", "END", "OVER", "PARTITION", "ORDER", "BY", "GROUP", "HAVING", "RETURN", "RETURNS", "BEGIN", "CREATE", "ALTER", "DROP", "TABLE", "VIEW", "FUNCTION", "PROCEDURE", "EXEC", "EXECUTE", "CONSTRAINT", "PRIMARY", "FOREIGN", "KEY", "REFERENCES", "EXISTS", "IN", "NOT", "AND", "OR", "CHECK", "DEFAULT", "TOP", "DISTINCT":
		return true
	default:
		return false
	}
}

func isDependencyIdentifier(token string) bool {
	if token == "" || token == "." || token == "(" || token == ")" || token == "=" || token == "," || token == ";" {
		return false
	}
	first := token[0]
	return first == '[' || first == '"' || first == '_' || first == '#' || first == '@' || first >= 128 || (first >= 'a' && first <= 'z') || (first >= 'A' && first <= 'Z')
}

func normalizeDependencyIdentifier(token string) string {
	token = strings.TrimSpace(token)
	if len(token) >= 2 && token[0] == '[' && token[len(token)-1] == ']' {
		token = strings.ReplaceAll(token[1:len(token)-1], "]]", "]")
	} else if len(token) >= 2 && token[0] == '"' && token[len(token)-1] == '"' {
		token = strings.ReplaceAll(token[1:len(token)-1], `""`, `"`)
	}
	return token
}

func isDependencyWordByte(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '#' || ch == '@' || ch == '$'
}
