package migreport

import (
	"sort"
	"strings"

	"rza_migrate_db/internal/sqlparse"
)

type relationshipToken struct{ Text, Upper string }
type relationshipEndpoint struct{ Object, Alias, Column string }
type relationshipAggregate struct {
	ObservedRelationship
	joinTypes map[string]struct{}
	modules   map[string]struct{}
}

func analyzeModuleRelationships(modules []moduleSource, tables map[string]*DocumentedTable, views map[string]struct{}) ([]ObservedRelationship, []ModuleDependency, []string) {
	aggregates := map[string]*relationshipAggregate{}
	dependencySet := map[string]ModuleDependency{}
	var dynamic []string
	for _, module := range modules {
		if moduleUsesDynamicSQL(module.SQL) {
			dynamic = append(dynamic, module.Name)
		}
		for _, ref := range sqlparse.ExtractObjectReferences(module.SQL) {
			if ref.Kind != sqlparse.ObjectReferenceData {
				continue
			}
			key := strings.ToLower(module.Name + "|" + module.Kind + "|" + ref.Name)
			dependencySet[key] = ModuleDependency{Module: module.Name, Kind: module.Kind, Object: ref.Name}
		}
		for _, observed := range extractObservedJoins(module, tables, views) {
			key := relationshipKey(observed)
			agg := aggregates[key]
			if agg == nil {
				agg = &relationshipAggregate{ObservedRelationship: observed, joinTypes: map[string]struct{}{}, modules: map[string]struct{}{}}
				aggregates[key] = agg
			}
			agg.Occurrences++
			agg.joinTypes[observed.JoinTypes[0]] = struct{}{}
			agg.modules[module.Name] = struct{}{}
		}
	}

	relations := make([]ObservedRelationship, 0, len(aggregates))
	for _, agg := range aggregates {
		agg.JoinTypes = sortedSet(agg.joinTypes)
		agg.Modules = sortedSet(agg.modules)
		relations = append(relations, agg.ObservedRelationship)
	}
	sort.Slice(relations, func(i, j int) bool { return relationshipKey(relations[i]) < relationshipKey(relations[j]) })
	dependencies := make([]ModuleDependency, 0, len(dependencySet))
	for _, dep := range dependencySet {
		dependencies = append(dependencies, dep)
	}
	sort.Slice(dependencies, func(i, j int) bool {
		return dependencies[i].Module+dependencies[i].Kind+dependencies[i].Object < dependencies[j].Module+dependencies[j].Kind+dependencies[j].Object
	})
	sort.Strings(dynamic)
	return relations, dependencies, uniqueRelationshipStrings(dynamic)
}

func extractObservedJoins(module moduleSource, tables map[string]*DocumentedTable, views map[string]struct{}) []ObservedRelationship {
	tokens := relationshipTokens(sqlparse.StripSQLStringsAndComments(module.SQL))
	aliases := map[string]string{}
	var out []ObservedRelationship
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Text == ";" {
			aliases = map[string]string{}
			continue
		}
		keyword := tokens[i].Upper
		if keyword != "FROM" && keyword != "JOIN" {
			continue
		}
		// A FROM starts a new query scope. Keeping aliases from an earlier
		// statement/query block is more dangerous than conservatively missing a
		// correlated-subquery relationship because SQL files commonly omit ';'.
		if keyword == "FROM" {
			aliases = map[string]string{}
		}
		object, next, ok := relationshipObject(tokens, i+1)
		if !ok {
			continue
		}
		alias, after := relationshipAlias(tokens, next, object)
		aliases[strings.ToLower(alias)] = object
		if keyword == "FROM" {
			// Support the legacy implicit-join form: FROM A a, B b WHERE ...
			cursor := after
			for cursor < len(tokens) && tokens[cursor].Text == "," {
				commaObject, commaNext, commaOK := relationshipObject(tokens, cursor+1)
				if !commaOK {
					break
				}
				commaAlias, commaAfter := relationshipAlias(tokens, commaNext, commaObject)
				aliases[strings.ToLower(commaAlias)] = commaObject
				cursor = commaAfter
			}
			i = after - 1
			continue
		}
		joinType := "INNER"
		if i > 0 {
			switch tokens[i-1].Upper {
			case "LEFT", "RIGHT", "FULL", "CROSS", "INNER":
				joinType = tokens[i-1].Upper
			}
		}
		on := after
		for on < len(tokens) && tokens[on].Upper != "ON" && !relationshipBoundary(tokens[on].Upper) {
			on++
		}
		if on >= len(tokens) || tokens[on].Upper != "ON" {
			i = after - 1
			continue
		}
		end := on + 1
		for end < len(tokens) && !relationshipBoundary(tokens[end].Upper) && tokens[end].Text != ";" {
			end++
		}
		out = append(out, observedRelationsInRange(tokens, on+1, end, aliases, tables, views, joinType, module.Name)...)
		i = end - 1
	}
	// A second pass captures legacy comma joins. Alias collection above remains
	// scoped conservatively by each FROM, and WHERE predicates are bounded.
	aliases = map[string]string{}
	for i := 0; i < len(tokens); i++ {
		if tokens[i].Upper == "FROM" {
			aliases = map[string]string{}
			cursor := i + 1
			for {
				object, next, ok := relationshipObject(tokens, cursor)
				if !ok {
					break
				}
				alias, after := relationshipAlias(tokens, next, object)
				aliases[strings.ToLower(alias)] = object
				cursor = after
				if cursor >= len(tokens) || tokens[cursor].Text != "," {
					break
				}
				cursor++
			}
		}
		if tokens[i].Upper != "WHERE" || len(aliases) < 2 {
			continue
		}
		end := i + 1
		for end < len(tokens) && !relationshipBoundary(tokens[end].Upper) && tokens[end].Text != ";" {
			end++
		}
		out = append(out, observedRelationsInRange(tokens, i+1, end, aliases, tables, views, "IMPLICIT", module.Name)...)
		i = end - 1
	}
	return out
}

func observedRelationsInRange(tokens []relationshipToken, start, end int, aliases map[string]string, tables map[string]*DocumentedTable, views map[string]struct{}, joinType, module string) []ObservedRelationship {
	var out []ObservedRelationship
	for eq := start; eq < end; eq++ {
		if !relationshipEquality(tokens, eq) {
			continue
		}
		leftStart := relationshipExpressionStart(tokens, start, eq)
		rightEnd := relationshipExpressionEnd(tokens, eq+1, end)
		left, leftOK := endpointInExpression(tokens, leftStart, eq, aliases)
		right, rightOK := endpointInExpression(tokens, eq+1, rightEnd, aliases)
		if !leftOK || !rightOK || strings.EqualFold(left.Alias, right.Alias) {
			continue
		}
		if !knownRelationshipEndpoint(left, tables, views) || !knownRelationshipEndpoint(right, tables, views) {
			continue
		}
		out = append(out, canonicalObservedRelationship(left, right, joinType, module))
	}
	return out
}

func relationshipTokens(sql string) []relationshipToken {
	var out []relationshipToken
	for i := 0; i < len(sql); {
		if sql[i] == ' ' || sql[i] == '\t' || sql[i] == '\r' || sql[i] == '\n' {
			i++
			continue
		}
		if sql[i] == '[' {
			start := i
			i++
			for i < len(sql) {
				if sql[i] == ']' {
					i++
					break
				}
				i++
			}
			value := strings.TrimSuffix(strings.TrimPrefix(sql[start:i], "["), "]")
			out = append(out, relationshipToken{Text: value, Upper: strings.ToUpper(value)})
			continue
		}
		if isRelationshipWord(sql[i]) {
			start := i
			for i < len(sql) && isRelationshipWord(sql[i]) {
				i++
			}
			value := sql[start:i]
			out = append(out, relationshipToken{Text: value, Upper: strings.ToUpper(value)})
			continue
		}
		out = append(out, relationshipToken{Text: string(sql[i]), Upper: string(sql[i])})
		i++
	}
	return out
}

func isRelationshipWord(ch byte) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z') || (ch >= '0' && ch <= '9') || ch == '_' || ch == '#' || ch == '@' || ch == '$'
}

func relationshipObject(tokens []relationshipToken, start int) (string, int, bool) {
	if start >= len(tokens) || tokens[start].Text == "(" || strings.HasPrefix(tokens[start].Text, "#") || strings.HasPrefix(tokens[start].Text, "@") {
		return "", start, false
	}
	parts := []string{tokens[start].Text}
	next := start + 1
	for next+1 < len(tokens) && tokens[next].Text == "." && len(parts) < 3 {
		parts = append(parts, tokens[next+1].Text)
		next += 2
	}
	if len(parts) == 1 {
		return "dbo." + parts[0], next, true
	}
	if len(parts) == 2 {
		return parts[0] + "." + parts[1], next, true
	}
	return parts[0] + "." + parts[1] + "." + parts[2], next, true
}

func relationshipAlias(tokens []relationshipToken, start int, object string) (string, int) {
	i := start
	if i < len(tokens) && tokens[i].Upper == "WITH" && i+1 < len(tokens) && tokens[i+1].Text == "(" {
		i = skipRelationshipParen(tokens, i+1)
	}
	if i < len(tokens) && tokens[i].Upper == "AS" {
		i++
	}
	if i < len(tokens) && relationshipAliasAllowed(tokens[i]) {
		return tokens[i].Text, i + 1
	}
	parts := strings.Split(object, ".")
	return parts[len(parts)-1], i
}

func relationshipAliasAllowed(token relationshipToken) bool {
	if token.Text == "," || token.Text == ";" || token.Text == "(" || token.Text == ")" {
		return false
	}
	switch token.Upper {
	case "ON", "JOIN", "LEFT", "RIGHT", "FULL", "INNER", "CROSS", "WHERE", "GROUP", "ORDER", "HAVING", "UNION", "WITH":
		return false
	}
	return true
}

func skipRelationshipParen(tokens []relationshipToken, open int) int {
	depth := 0
	for i := open; i < len(tokens); i++ {
		if tokens[i].Text == "(" {
			depth++
		}
		if tokens[i].Text == ")" {
			depth--
			if depth == 0 {
				return i + 1
			}
		}
	}
	return len(tokens)
}

func relationshipBoundary(upper string) bool {
	switch upper {
	case "JOIN", "WHERE", "GROUP", "ORDER", "HAVING", "UNION", "EXCEPT", "INTERSECT", "RETURN", "OPTION":
		return true
	}
	return false
}

func relationshipEquality(tokens []relationshipToken, eq int) bool {
	if eq < 0 || eq >= len(tokens) || tokens[eq].Text != "=" {
		return false
	}
	// The tokenizer intentionally keeps punctuation simple, so reject the '='
	// portion of >=, <=, !=, and == here.
	if eq > 0 && (tokens[eq-1].Text == ">" || tokens[eq-1].Text == "<" || tokens[eq-1].Text == "!" || tokens[eq-1].Text == "=") {
		return false
	}
	return eq+1 >= len(tokens) || tokens[eq+1].Text != "="
}

func relationshipExpressionStart(tokens []relationshipToken, lower, eq int) int {
	for i := eq - 1; i >= lower; i-- {
		if tokens[i].Upper == "AND" || tokens[i].Upper == "OR" || tokens[i].Upper == "ON" || tokens[i].Text == "," {
			return i + 1
		}
	}
	return lower
}

func relationshipExpressionEnd(tokens []relationshipToken, start, upper int) int {
	for i := start; i < upper; i++ {
		if tokens[i].Upper == "AND" || tokens[i].Upper == "OR" || tokens[i].Text == "," {
			return i
		}
	}
	return upper
}

// endpointInExpression accepts an expression only when it contains exactly one
// distinct resolved alias.column endpoint. This supports harmless wrappers such
// as LTRIM(a.Code), but refuses arithmetic/CASE expressions that mention more
// than one column and therefore cannot prove a direct column mapping.
func endpointInExpression(tokens []relationshipToken, start, end int, aliases map[string]string) (relationshipEndpoint, bool) {
	var found relationshipEndpoint
	hasFound := false
	for i := start; i+2 < end; i++ {
		if tokens[i+1].Text != "." {
			continue
		}
		object, ok := aliases[strings.ToLower(tokens[i].Text)]
		if !ok {
			continue
		}
		candidate := relationshipEndpoint{Object: object, Alias: tokens[i].Text, Column: tokens[i+2].Text}
		if !hasFound {
			found = candidate
			hasFound = true
			continue
		}
		if !strings.EqualFold(found.Object, candidate.Object) || !strings.EqualFold(found.Alias, candidate.Alias) || !strings.EqualFold(found.Column, candidate.Column) {
			return relationshipEndpoint{}, false
		}
	}
	return found, hasFound
}

func knownRelationshipObject(object string, tables map[string]*DocumentedTable, views map[string]struct{}) bool {
	if _, ok := lookupDocumentedTable(tables, object); ok {
		return true
	}
	for view := range views {
		if strings.EqualFold(view, object) {
			return true
		}
	}
	return strings.Count(object, ".") == 2
}

func knownRelationshipEndpoint(endpoint relationshipEndpoint, tables map[string]*DocumentedTable, views map[string]struct{}) bool {
	if table, ok := lookupDocumentedTable(tables, endpoint.Object); ok {
		for _, column := range table.Columns {
			if strings.EqualFold(column.Name, endpoint.Column) {
				return true
			}
		}
		return false
	}
	return knownRelationshipObject(endpoint.Object, tables, views)
}

func canonicalObservedRelationship(left, right relationshipEndpoint, joinType, module string) ObservedRelationship {
	if left.Object+"|"+left.Column > right.Object+"|"+right.Column {
		left, right = right, left
	}
	confidence := "MEDIUM"
	if strings.Count(left.Object, ".") == 2 || strings.Count(right.Object, ".") == 2 {
		confidence = "LOW"
	}
	return ObservedRelationship{LeftObject: left.Object, LeftColumns: []string{left.Column}, RightObject: right.Object, RightColumns: []string{right.Column}, Evidence: "OBSERVED_JOIN", JoinTypes: []string{joinType}, Modules: []string{module}, Occurrences: 0, Confidence: confidence}
}

func relationshipKey(r ObservedRelationship) string {
	return r.LeftObject + "|" + strings.Join(r.LeftColumns, ",") + "|" + r.RightObject + "|" + strings.Join(r.RightColumns, ",") + "|" + r.Evidence
}
func sortedSet(values map[string]struct{}) []string {
	out := make([]string, 0, len(values))
	for value := range values {
		out = append(out, value)
	}
	sort.Strings(out)
	return out
}
func uniqueRelationshipStrings(values []string) []string {
	if len(values) == 0 {
		return nil
	}
	sort.Strings(values)
	out := values[:0]
	for _, value := range values {
		if len(out) == 0 || out[len(out)-1] != value {
			out = append(out, value)
		}
	}
	return out
}
func moduleUsesDynamicSQL(sql string) bool {
	clean := sqlparse.StripSQLStringsAndComments(sql)
	upper := strings.ToUpper(clean)
	if containsSQLWord(upper, "SP_EXECUTESQL") {
		return true
	}
	tokens := relationshipTokens(clean)
	for i := 0; i+1 < len(tokens); i++ {
		if tokens[i].Upper != "EXEC" && tokens[i].Upper != "EXECUTE" {
			continue
		}
		next := tokens[i+1].Text
		if next == "(" {
			return true
		}
		if strings.HasPrefix(next, "@") {
			// EXEC @return_code = dbo.StaticProcedure is a normal static call,
			// not execution of a procedure/SQL expression stored in @return_code.
			if i+2 < len(tokens) && tokens[i+2].Text == "=" {
				continue
			}
			return true
		}
	}
	return false
}
