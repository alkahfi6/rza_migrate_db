package migreport

import (
	"fmt"
	"sort"
	"strings"
)

// This is documentation evidence only. JOIN operands can be references or
// grouping values; they are not inferred unique keys or schema corrections.
func renderSuspectedJoinIdentifiers(docs DesiredStateDocs) string {
	var b strings.Builder
	b.WriteString("\n## Suspected JOIN identifiers\n\nListing ini menunjukkan kolom table yang dipakai dalam static equality JOIN, termasuk kolom tanpa index atau `IDENTITY` dan kolom pada table yang sudah memiliki primary key lain. `OBSERVED_JOIN_ID_SIGNAL` berarti namanya menyerupai ID/key; `OBSERVED_JOIN_COLUMN_SIGNAL` berarti pemakaian JOIN menjadi satu-satunya dasar kecurigaan. Keduanya adalah candidate join/reference, bukan bukti unique row ID. Kolom yang sudah memiliki single-column unfiltered unique key tersedia di inventory utama dan tidak diulang di sini.\n\n")
	b.WriteString("| Table | Column | Classification | Uniqueness | Join usage | Evidence |\n| --- | --- | --- | --- | ---: | --- |\n")
	tables := append([]DocumentedTable(nil), docs.Tables...)
	sort.Slice(tables, func(i, j int) bool { return tables[i].Name < tables[j].Name })
	wrote := false
	for _, table := range tables {
		columns := append([]DocumentedColumn(nil), table.Columns...)
		sort.Slice(columns, func(i, j int) bool { return columns[i].Name < columns[j].Name })
		for _, column := range columns {
			if columnsCoveredByFullUnique([]string{column.Name}, fullUniqueKeys(table)) {
				continue
			}
			usage, predicates, modules := observedColumnJoinEvidence(docs, table.Name, column.Name)
			if usage == 0 {
				continue
			}
			classification := "OBSERVED_JOIN_COLUMN_SIGNAL"
			if idLikeColumn(column) {
				classification = "OBSERVED_JOIN_ID_SIGNAL"
			}
			uniqueness := "NOT_ENFORCED"
			var indexNames, compositeKeys []string
			for _, key := range table.Keys {
				if !containsFold(key.Columns, column.Name) {
					continue
				}
				indexNames = append(indexNames, key.Name)
				if key.Filtered && len(key.Columns) == 1 {
					uniqueness = "CONDITIONAL"
				}
				if len(key.Columns) > 1 {
					compositeKeys = append(compositeKeys, "("+strings.Join(key.Columns, ", ")+")")
				}
			}
			for _, index := range table.Indexes {
				if containsFold(index.Columns, column.Name) {
					indexNames = append(indexNames, index.Name)
				}
			}
			indexEvidence := "Index/key: tidak ditemukan"
			if names := sortedEvidenceStrings(indexNames); len(names) > 0 {
				indexEvidence = "Index/key: " + strings.Join(names, "; ")
			}
			identityEvidence := "IDENTITY: tidak"
			if column.Identity {
				identityEvidence = "IDENTITY: ya"
			}
			evidence := "Mapping kolom pada equality JOIN: " + strings.Join(predicates, "; ") + ". Source modules: " + strings.Join(modules, "; ") + ". " + indexEvidence + ". " + identityEvidence + "."
			if len(compositeKeys) > 0 {
				evidence += " Bagian dari composite key " + strings.Join(sortedEvidenceStrings(compositeKeys), "; ") + "; uniqueness kolom sendiri tidak terbukti."
			}
			b.WriteString(fmt.Sprintf("| `%s` | `%s` | `%s` | `%s` | %d | %s |\n", md(table.Name), md(column.Name), classification, uniqueness, usage, md(evidence)))
			wrote = true
		}
	}
	if !wrote {
		b.WriteString("| — | — | — | — | 0 | Tidak ditemukan candidate JOIN identifier tambahan dari source. |\n")
	}
	b.WriteString("\nEvidence berasal dari view/procedure pada `db_object`, bukan isi row database. Declared FK saja, kesamaan nama, inequality, alias/column ambigu, temporary object, atau SQL runtime yang belum ter-resolve tidak menjadi observed JOIN evidence. Filtered uniqueness hanya berlaku untuk subset sesuai filter.\n")
	return b.String()
}

func observedColumnJoinEvidence(docs DesiredStateDocs, object, column string) (int, []string, []string) {
	usage := 0
	var predicates, modules []string
	for _, r := range docs.Relationships {
		if r.Evidence != "OBSERVED_JOIN" {
			continue
		}
		left := strings.EqualFold(r.LeftObject, object) && containsFold(r.LeftColumns, column)
		right := strings.EqualFold(r.RightObject, object) && containsFold(r.RightColumns, column)
		if !left && !right {
			continue
		}
		// A self-join involving the same column in both roles is one occurrence.
		usage += r.Occurrences
		if len(r.LeftColumns) == len(r.RightColumns) {
			for i := range r.LeftColumns {
				if (left && strings.EqualFold(r.LeftColumns[i], column)) || (right && strings.EqualFold(r.RightColumns[i], column)) {
					predicates = append(predicates, r.LeftObject+"."+r.LeftColumns[i]+" = "+r.RightObject+"."+r.RightColumns[i])
				}
			}
		}
		modules = append(modules, r.Modules...)
	}
	return usage, sortedEvidenceStrings(predicates), sortedEvidenceStrings(modules)
}

func hasObservedJoinColumns(docs DesiredStateDocs, table DocumentedTable) bool {
	for _, column := range table.Columns {
		if usage, _, _ := observedColumnJoinEvidence(docs, table.Name, column.Name); usage > 0 {
			return true
		}
	}
	return false
}

func sortedEvidenceStrings(values []string) []string {
	set := map[string]bool{}
	for _, value := range values {
		if value != "" {
			set[value] = true
		}
	}
	out := make([]string, 0, len(set))
	for value := range set {
		out = append(out, value)
	}
	sort.Strings(out)
	return out
}
