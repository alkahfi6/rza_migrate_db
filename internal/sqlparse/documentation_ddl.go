package sqlparse

import "strings"

// DDLMetadata is an observational source projection. It does not validate a
// migration or authorize corrections. Keywords, names and statement boundaries
// use the same lexer as access and dependency discovery.
type DDLMetadata struct {
	Name        string
	Columns     []DDLColumn
	Keys        []DDLKey
	ForeignKeys []DDLForeignKey
}
type DDLColumn struct {
	Name, DataType     string
	Nullable, Identity bool
}
type DDLKey struct {
	Name, Evidence string
	Columns        []string
	Filtered       bool
}
type DDLForeignKey struct {
	Name, Target           string
	Columns, TargetColumns []string
}
type DDLIndex struct {
	Name, Table      string
	Columns          []string
	Unique, Filtered bool
}

func ddlKeyword(t sourceToken, keyword string) bool { return !t.quoted && t.upper == keyword }

// ModuleDeclarationName reads the declared identity, not names mentioned in a
// body/comment/literal. It is not a module grammar validator.
func ModuleDeclarationName(sql, kind string) string {
	tokens := sourceTokens(sql)
	for i := 0; i+2 < len(tokens); i++ {
		if !ddlKeyword(tokens[i], "CREATE") && !ddlKeyword(tokens[i], "ALTER") {
			continue
		}
		j := i + 1
		if j+2 < len(tokens) && ddlKeyword(tokens[j], "OR") && ddlKeyword(tokens[j+1], "ALTER") {
			j += 2
		}
		if j >= len(tokens) {
			continue
		}
		matches := ddlKeyword(tokens[j], strings.ToUpper(kind)) || (kind == "procedure" && ddlKeyword(tokens[j], "PROC"))
		if !matches {
			continue
		}
		name, parts, _, ok := sourceObjectName(tokens, j+1)
		if ok {
			return ddlCatalogName(name, parts)
		}
	}
	return ""
}

func ReadTableMetadata(sql string) DDLMetadata {
	tokens := sourceTokens(sql)
	var out DDLMetadata
	for i := 0; i+2 < len(tokens); i++ {
		if !ddlKeyword(tokens[i], "CREATE") || !ddlKeyword(tokens[i+1], "TABLE") {
			continue
		}
		name, nameParts, next, ok := sourceObjectName(tokens, i+2)
		if !ok || next >= len(tokens) || tokens[next].text != "(" {
			continue
		}
		end := sourceSkipTokenParens(tokens, next)
		if end <= next+1 || end > len(tokens) || tokens[end-1].text != ")" {
			return out
		}
		out.Name = ddlCatalogName(name, nameParts)
		for _, item := range SplitTopLevelComma(sql[tokens[next].end:tokens[end-1].start]) {
			readDDLTableItem(item, &out)
		}
		// Extractor dumps may declare foreign keys after the CREATE body.
		for j := end; j+2 < len(tokens); j++ {
			if !ddlKeyword(tokens[j], "ALTER") || !ddlKeyword(tokens[j+1], "TABLE") {
				continue
			}
			target, targetParts, after, valid := sourceObjectName(tokens, j+2)
			if !valid || !strings.EqualFold(ddlCatalogName(target, targetParts), out.Name) {
				continue
			}
			stop := ddlStatementEnd(tokens, j, len(sql))
			for after < len(tokens) && tokens[after].start < stop && !ddlKeyword(tokens[after], "ADD") {
				after++
			}
			if after < len(tokens) && tokens[after].start < stop {
				for _, item := range SplitTopLevelComma(sql[tokens[after].end:stop]) {
					readDDLTableItem(item, &out)
				}
			}
		}
		return out
	}
	return out
}

func readDDLTableItem(sql string, out *DDLMetadata) {
	tokens := sourceTokens(sql)
	if len(tokens) == 0 {
		return
	}
	column := ""
	constraint := ""
	start := 0
	if ddlKeyword(tokens[0], "CONSTRAINT") && len(tokens) > 1 {
		constraint = tokens[1].text
		start = 2
	}
	if start == 0 && sourceObjectIdentifier(tokens[0]) && !ddlKeyword(tokens[0], "PRIMARY") && !ddlKeyword(tokens[0], "UNIQUE") && !ddlKeyword(tokens[0], "FOREIGN") && !ddlKeyword(tokens[0], "CHECK") && !ddlKeyword(tokens[0], "INDEX") {
		column = tokens[0].text
		col := DDLColumn{Name: column, Nullable: true}
		if len(tokens) > 1 {
			if ddlKeyword(tokens[1], "AS") {
				col.DataType = "COMPUTED"
			} else {
				_, _, next, ok := sourceObjectName(tokens, 1)
				if ok {
					if next < len(tokens) && tokens[next].text == "(" {
						next = sourceSkipTokenParens(tokens, next)
					}
					if next > 1 {
						col.DataType = strings.ToUpper(strings.Join(strings.Fields(sql[tokens[1].start:tokens[next-1].end]), " "))
					}
				}
			}
		}
		for i := 1; i < len(tokens); i++ {
			if tokens[i].depth != 0 {
				continue
			}
			if ddlKeyword(tokens[i], "IDENTITY") {
				col.Identity = true
			}
			if ddlKeyword(tokens[i], "NOT") && i+1 < len(tokens) && ddlKeyword(tokens[i+1], "NULL") {
				col.Nullable = false
			}
			if ddlKeyword(tokens[i], "PRIMARY") && i+1 < len(tokens) && ddlKeyword(tokens[i+1], "KEY") {
				col.Nullable = false
			}
		}
		out.Columns = append(out.Columns, col)
		start = 1
	}
	for i := start; i < len(tokens); i++ {
		if tokens[i].depth != 0 || tokens[i].quoted {
			continue
		}
		if ddlKeyword(tokens[i], "CONSTRAINT") && i+1 < len(tokens) {
			constraint = tokens[i+1].text
			continue
		}
		if ddlKeyword(tokens[i], "PRIMARY") && i+1 < len(tokens) && ddlKeyword(tokens[i+1], "KEY") {
			cols := ddlColumnsAfter(tokens, i+2)
			if column != "" {
				cols = []string{column}
			}
			if len(cols) > 0 {
				out.Keys = append(out.Keys, DDLKey{Name: constraint, Columns: cols, Evidence: "PRIMARY_KEY"})
			}
			continue
		}
		if ddlKeyword(tokens[i], "UNIQUE") {
			cols := ddlColumnsAfter(tokens, i+1)
			if column != "" {
				cols = []string{column}
			}
			if len(cols) > 0 {
				out.Keys = append(out.Keys, DDLKey{Name: constraint, Columns: cols, Evidence: "UNIQUE_CONSTRAINT"})
			}
			continue
		}
		if ddlKeyword(tokens[i], "REFERENCES") {
			target, targetParts, next, ok := sourceObjectName(tokens, i+1)
			if !ok {
				return
			}
			cols := []string{}
			if column != "" {
				cols = []string{column}
			} else {
				for j := start; j < i; j++ {
					if ddlKeyword(tokens[j], "FOREIGN") {
						cols = ddlColumnsAfter(tokens, j+2)
						break
					}
				}
			}
			refCols := ddlColumnsAfter(tokens, next)
			if len(cols) > 0 && len(cols) == len(refCols) {
				out.ForeignKeys = append(out.ForeignKeys, DDLForeignKey{Name: constraint, Target: ddlCatalogName(target, targetParts), Columns: cols, TargetColumns: refCols})
			}
			return
		}
	}
}

func ddlColumnsAfter(tokens []sourceToken, start int) []string {
	for start < len(tokens) && tokens[start].text != "(" {
		start++
	}
	if start >= len(tokens) {
		return nil
	}
	end := sourceSkipTokenParens(tokens, start)
	if end <= start+1 || end > len(tokens) || tokens[end-1].text != ")" {
		return nil
	}
	var out []string
	expect := true
	ordered := false
	for i := start + 1; i < end-1; i++ {
		if tokens[i].depth != tokens[start].depth+1 {
			continue
		}
		if expect {
			if !sourceObjectIdentifier(tokens[i]) {
				return nil
			}
			out = append(out, tokens[i].text)
			expect = false
			ordered = false
		} else if tokens[i].text == "," {
			expect = true
		} else if !ordered && (ddlKeyword(tokens[i], "ASC") || ddlKeyword(tokens[i], "DESC")) {
			ordered = true
		} else {
			return nil
		}
	}
	if expect {
		return nil
	}
	return out
}

func ddlCatalogName(name string, parts []string) string {
	if len(parts) == 1 {
		return "dbo." + name
	}
	return name
}

func ddlStatementEnd(tokens []sourceToken, start, sqlLen int) int {
	for i := start + 1; i < len(tokens); i++ {
		if tokens[i].depth != tokens[start].depth || tokens[i].quoted {
			continue
		}
		if tokens[i].text == ";" || ddlKeyword(tokens[i], "GO") || ddlKeyword(tokens[i], "CREATE") || ddlKeyword(tokens[i], "ALTER") {
			return tokens[i].start
		}
	}
	return sqlLen
}

func ReadIndexMetadata(sql string) []DDLIndex {
	tokens := sourceTokens(sql)
	var out []DDLIndex
	for i := 0; i < len(tokens); i++ {
		if !ddlKeyword(tokens[i], "CREATE") {
			continue
		}
		j := i + 1
		unique := false
		for j < len(tokens) && !tokens[j].quoted && (tokens[j].upper == "UNIQUE" || tokens[j].upper == "CLUSTERED" || tokens[j].upper == "NONCLUSTERED" || tokens[j].upper == "COLUMNSTORE") {
			if tokens[j].upper == "UNIQUE" {
				unique = true
			}
			j++
		}
		if j+2 >= len(tokens) || !ddlKeyword(tokens[j], "INDEX") || !sourceObjectIdentifier(tokens[j+1]) || !ddlKeyword(tokens[j+2], "ON") {
			continue
		}
		name := tokens[j+1].text
		table, tableParts, next, ok := sourceObjectName(tokens, j+3)
		if !ok || next >= len(tokens) || tokens[next].text != "(" {
			continue
		}
		cols := ddlColumnsAfter(tokens, next)
		if len(cols) == 0 {
			continue
		}
		stop := ddlStatementEnd(tokens, i, len(sql))
		filtered := false
		for k := sourceSkipTokenParens(tokens, next); k < len(tokens) && tokens[k].start < stop; k++ {
			if tokens[k].depth == tokens[i].depth && ddlKeyword(tokens[k], "WHERE") {
				filtered = true
			}
		}
		out = append(out, DDLIndex{Name: name, Table: ddlCatalogName(table, tableParts), Columns: cols, Unique: unique, Filtered: filtered})
	}
	return out
}
