package migreport

import (
	"sort"
	"strings"

	"rza_migrate_db/internal/sqlparse"
)

type relationshipToken = sqlparse.SQLToken
type relationshipEndpoint struct {
	Object, Alias, Column string
	External              bool
}
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
		catalog := newProcedureCatalog(tables, views, nil, nil)
		for _, ref := range sqlparse.ExtractSQLAccesses(module.SQL) {
			if ref.Function {
				continue
			}
			object, status := catalog.resolveParts(ref.Name, ref.NameParts, "data")
			if ref.NonPersisted {
				object, status = ref.Name, ResolutionNotPersisted
			}
			if ref.Ambiguous || ref.CTECaseAmbiguous {
				object, status = ref.Name, ResolutionAmbiguous
			}
			key := strings.ToLower(module.Name + "|" + module.Kind + "|" + object)
			if _, exists := dependencySet[key]; !exists {
				dependencySet[key] = ModuleDependency{Module: module.Name, Kind: module.Kind, Object: object, Resolution: status}
			}
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
	tokens := relationshipTokens(module.SQL)
	accesses := sqlparse.ExtractSQLAccesses(module.SQL)
	catalog := newProcedureCatalog(tables, views, nil, nil)
	byOffset := map[int]sqlparse.SQLAccess{}
	for _, access := range accesses {
		if access.Operation == "READ" {
			byOffset[access.Offset] = access
		}
	}
	var out []ObservedRelationship
	for from, token := range tokens {
		if token.Quoted || token.Upper != "FROM" {
			continue
		}
		seed, ok := byOffset[token.Start]
		if !ok {
			continue
		} // Cursor and non-query FROM are not join scopes.
		end := len(tokens)
		for j := from + 1; j < len(tokens); j++ {
			t := tokens[j]
			if t.Depth < token.Depth || t.Start >= seed.StatementEnd {
				end = j
				break
			}
			if t.Depth == token.Depth && !t.Quoted {
				if t.Upper == "MERGE" && j+1 < len(tokens) && !tokens[j+1].Quoted && tokens[j+1].Upper == "JOIN" {
					continue
				}
				switch t.Upper {
				case "SELECT", "FROM", "UNION", "EXCEPT", "INTERSECT", "INSERT", "UPDATE", "DELETE", "MERGE", "EXEC", "EXECUTE", "RETURN", "GO", ";":
					end = j
				}
				if end == j {
					break
				}
			}
		}
		aliases := map[string]string{}
		duplicate := map[string]bool{}
		commaJoin := false
		for _, access := range accesses {
			if access.Operation != "READ" || access.QueryDepth != token.Depth || access.Offset < token.Start || (end < len(tokens) && access.Offset >= tokens[end].Start) {
				continue
			}
			alias := access.Alias
			if alias == "" && len(access.NameParts) > 0 {
				alias = access.NameParts[len(access.NameParts)-1]
			}
			key := strings.ToLower(alias)
			object, status := catalog.resolveParts(access.Name, access.NameParts, "data")
			if access.NonPersisted || access.Ambiguous || access.CTECaseAmbiguous || access.Function || (status != ResolutionStatic && status != ResolutionExternal) {
				object = ""
			}
			if _, exists := aliases[key]; exists {
				duplicate[key] = true
				object = ""
			}
			if duplicate[key] {
				object = ""
			}
			aliases[key] = object
			if access.Alias == "" && len(access.NameParts) > 1 {
				aliases[strings.ToLower(access.Name)] = object
			}
		}
		for j := from + 1; j < end; j++ {
			t := tokens[j]
			if t.Depth != token.Depth || t.Quoted {
				continue
			}
			if t.Text == "," {
				if _, ok := byOffset[t.Start]; ok {
					commaJoin = true
				}
				continue
			}
			if t.Upper == "WHERE" {
				if commaJoin {
					stop := relationshipPredicateEnd(tokens, j+1, end, token.Depth)
					out = append(out, observedRelationsInRange(tokens, j+1, stop, aliases, tables, views, "IMPLICIT", module.Name)...)
				}
				continue
			}
			if t.Upper != "JOIN" {
				continue
			}
			joinType := "INNER"
			prev := j - 1
			if prev >= from && !tokens[prev].Quoted {
				switch tokens[prev].Upper {
				case "HASH", "LOOP", "MERGE", "REMOTE":
					prev--
				}
			}
			if prev >= from && !tokens[prev].Quoted && tokens[prev].Upper == "OUTER" {
				prev--
			}
			if prev >= from && !tokens[prev].Quoted {
				switch tokens[prev].Upper {
				case "LEFT", "RIGHT", "FULL", "CROSS", "INNER":
					joinType = tokens[prev].Upper
				}
			}
			on := j + 1
			for on < end && !(tokens[on].Depth == token.Depth && !tokens[on].Quoted && (tokens[on].Upper == "ON" || relationshipBoundary(tokens[on].Upper))) {
				on++
			}
			if on >= end || tokens[on].Upper != "ON" || tokens[on].Quoted {
				continue
			}
			stop := relationshipPredicateEnd(tokens, on+1, end, token.Depth)
			out = append(out, observedRelationsInRange(tokens, on+1, stop, aliases, tables, views, joinType, module.Name)...)
		}
	}
	return out
}

func relationshipPredicateEnd(tokens []relationshipToken, start, end, depth int) int {
	for i := start; i < end; i++ {
		if tokens[i].Depth == depth && !tokens[i].Quoted && relationshipBoundary(tokens[i].Upper) {
			return i
		}
	}
	return end
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
		left, leftOK = canonicalRelationshipEndpoint(left, tables, views)
		right, rightOK = canonicalRelationshipEndpoint(right, tables, views)
		if !leftOK || !rightOK {
			continue
		}
		out = append(out, canonicalObservedRelationship(left, right, joinType, module))
	}
	return out
}

func relationshipTokens(sql string) []relationshipToken { return sqlparse.SQLTokens(sql) }

func relationshipBoundary(upper string) bool {
	switch upper {
	case "JOIN", "LEFT", "RIGHT", "FULL", "INNER", "CROSS", "OUTER", "WHERE", "GROUP", "ORDER", "HAVING", "UNION", "EXCEPT", "INTERSECT", "RETURN", "OPTION", "FOR":
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
		if !tokens[i].Quoted && (tokens[i].Upper == "AND" || tokens[i].Upper == "OR" || tokens[i].Upper == "ON") {
			return i + 1
		}
	}
	return lower
}

func relationshipExpressionEnd(tokens []relationshipToken, start, upper int) int {
	for i := start; i < upper; i++ {
		if !tokens[i].Quoted && (tokens[i].Upper == "AND" || tokens[i].Upper == "OR") {
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
	for _, token := range tokens[start:end] {
		if !token.Quoted && (token.Upper == "SELECT" || token.Upper == "EXISTS") {
			return found, false
		}
	}
	for i := start; i < end; i++ {
		if !tokens[i].Identifier {
			continue
		}
		parts := []string{tokens[i].Text}
		next := i + 1
		for next+1 < end && tokens[next].Text == "." && tokens[next+1].Identifier {
			parts = append(parts, tokens[next+1].Text)
			next += 2
		}
		if len(parts) < 2 {
			continue
		}
		if next < end && tokens[next].Text == "(" {
			i = next - 1
			continue
		} // qualified scalar wrapper
		alias := strings.Join(parts[:len(parts)-1], ".")
		object, ok := aliases[strings.ToLower(alias)]
		if !ok || object == "" {
			return found, false
		}
		candidate := relationshipEndpoint{Object: object, Alias: alias, Column: parts[len(parts)-1]}
		if hasFound && (!strings.EqualFold(found.Object, candidate.Object) || !strings.EqualFold(found.Alias, candidate.Alias) || !strings.EqualFold(found.Column, candidate.Column)) {
			return found, false
		}
		found = candidate
		hasFound = true
		i = next - 1
	}
	return found, hasFound
}

func canonicalRelationshipEndpoint(endpoint relationshipEndpoint, tables map[string]*DocumentedTable, views map[string]struct{}) (relationshipEndpoint, bool) {
	if table, ok := lookupDocumentedTable(tables, endpoint.Object); ok {
		endpoint.Object = table.Name
		var match string
		for _, column := range table.Columns {
			if strings.EqualFold(column.Name, endpoint.Column) {
				if match != "" {
					return endpoint, false
				}
				match = column.Name
			}
		}
		if match == "" {
			return endpoint, false
		}
		endpoint.Column = match
		return endpoint, true
	}
	matches := 0
	for view := range views {
		if strings.EqualFold(view, endpoint.Object) {
			endpoint.Object = view
			matches++
		}
	}
	endpoint.External = matches == 0 && strings.Count(endpoint.Object, ".") >= 2
	return endpoint, matches == 1 || endpoint.External
}

func canonicalObservedRelationship(left, right relationshipEndpoint, joinType, module string) ObservedRelationship {
	if left.Object+"|"+left.Column > right.Object+"|"+right.Column {
		left, right = right, left
	}
	confidence := "MEDIUM"
	if left.External || right.External {
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
	for _, call := range sqlparse.ExtractSQLCalls(sql) {
		if len(call.NameParts) > 0 && strings.EqualFold(call.NameParts[len(call.NameParts)-1], "sp_executesql") {
			return true
		}
	}
	tokens := relationshipTokens(sql)
	for i := 0; i+1 < len(tokens); i++ {
		if tokens[i].Quoted || (tokens[i].Upper != "EXEC" && tokens[i].Upper != "EXECUTE") {
			continue
		}
		next := tokens[i+1]
		if next.Text == "(" || strings.HasPrefix(next.Text, "'") || (next.Upper == "N" && i+2 < len(tokens) && strings.HasPrefix(tokens[i+2].Text, "'")) {
			return true
		}
		if strings.HasPrefix(next.Text, "@") && !next.Quoted {
			if i+2 < len(tokens) && tokens[i+2].Text == "=" {
				if i+3 < len(tokens) && strings.HasPrefix(tokens[i+3].Text, "@") && !tokens[i+3].Quoted {
					return true
				}
				continue
			}
			return true
		}
	}
	return false
}
