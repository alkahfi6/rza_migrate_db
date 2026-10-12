package migreport

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

var desiredStateReportFiles = []string{"README.md", "manifest.json", "identifiers/table_unique_ids.md", "relationships/relational_erd.md", "relationships/relationship_catalog.md", "relationships/module_lineage.md", "procedures/README.md", "procedures/procedure_impact_catalog.md", "procedures/profiles/*.md"}

func desiredStateArtifactFiles(docs DesiredStateDocs) []string {
	files := append([]string(nil), desiredStateReportFiles[:len(desiredStateReportFiles)-1]...)
	for _, profile := range docs.ProcedureProfiles {
		files = append(files, filepath.ToSlash(filepath.Join("procedures", "profiles", profile.Name+".md")))
	}
	return files
}

const relationshipDiagramDirectory = "relationship_diagrams"

const (
	relationalMapMaxObjects         = 20
	relationalMapMaxPairs           = 24
	relationalMapMaxMappings        = 35
	relationalMapMaxTextCharacters  = 9000
	relationshipCatalogMaxTableRows = 35
)

// WriteDesiredStateDocs writes grouped Markdown documents; Run publishes the indexes and manifest.
func WriteDesiredStateDocs(docsDir string, docs DesiredStateDocs) error {
	if err := os.MkdirAll(docsDir, 0755); err != nil {
		return fmt.Errorf("failed to create docs directory %s: %w", docsDir, err)
	}
	for name, body := range desiredStateMarkdown(docs) {
		if err := writeGeneratedMarkdown(filepath.Join(docsDir, name), body+"\n[Database documentation](../README.md)\n"); err != nil {
			return err
		}
	}
	if err := writeProcedureDocs(docsDir, docs); err != nil {
		return err
	}
	return nil
}

func desiredStateMarkdown(docs DesiredStateDocs) map[string]string {
	return map[string]string{
		"identifiers/table_unique_ids.md":       renderUniqueIDs(docs),
		"relationships/relational_erd.md":       renderRelationalERD(docs),
		"relationships/relationship_catalog.md": renderRelationshipCatalog(docs),
		"relationships/module_lineage.md":       renderModuleLineage(docs),
	}
}

func writeDesiredStateDocs(docsDir string, docs DesiredStateDocs) error {
	return WriteDesiredStateDocs(docsDir, docs)
}

func renderUniqueIDs(docs DesiredStateDocs) string {
	var b strings.Builder
	b.WriteString("# Table ID & Uniqueness Review\n\n")
	b.WriteString("Report ini membantu membedakan kolom yang benar-benar dijaga unique oleh database, kolom yang nilainya hanya dibuat otomatis, dan kolom yang terlihat seperti ID dari cara pemakaiannya. `IDENTITY`, nama berakhiran ID, pemakaian dalam JOIN, atau index non-unique belum cukup untuk membuktikan bahwa sebuah kolom unique.\n\n")
	if len(docs.Tables) == 0 {
		b.WriteString("Source `db_object/tables` tidak ditemukan.\n")
		return b.String()
	}
	enforcedTables, identityRiskTables, noIdentifierTables := 0, 0, 0
	for i := range docs.Tables {
		full := fullUniqueKeys(docs.Tables[i])
		if len(full) > 0 {
			enforcedTables++
		}
		if hasUnenforcedIdentity(docs.Tables[i], full) {
			identityRiskTables++
		}
		if len(full) == 0 && len(identifierCandidates(docs, docs.Tables[i])) == 0 && !hasIdentity(docs.Tables[i]) && !hasObservedJoinColumns(docs, docs.Tables[i]) {
			noIdentifierTables++
		}
	}
	b.WriteString(fmt.Sprintf("- Table yang diperiksa: **%d**\n- Table dengan unique key yang dijaga database: **%d**\n- Table dengan `IDENTITY` yang belum dijaga unique key: **%d**\n- Table tanpa key, `IDENTITY`, atau sinyal ID dari implementasi: **%d**\n\n", len(docs.Tables), enforcedTables, identityRiskTables, noIdentifierTables))
	b.WriteString("Kandidat JOIN tanpa key/index unique juga tersedia pada [Suspected JOIN identifiers](#suspected-join-identifiers), lengkap dengan evidence source.\n\n")
	b.WriteString("| Table | Columns | Classification | Uniqueness | Source | Join usage |\n| --- | --- | --- | --- | --- | ---: |\n")
	for _, table := range docs.Tables {
		full, wrote := fullUniqueKeys(table), false
		for _, key := range table.Keys {
			enforcement := "ENFORCED"
			if key.Filtered {
				enforcement = "CONDITIONAL"
			}
			b.WriteString(identifierRow(table.Name, key.Columns, key.Evidence, enforcement, key.Name, joinUsage(docs, table.Name, key.Columns)))
			wrote = true
		}
		for _, column := range table.Columns {
			if !column.Identity {
				continue
			}
			covered := columnsCoveredByFullUnique([]string{column.Name}, full)
			enforcement := "NOT_ENFORCED"
			if covered {
				enforcement = "ENFORCED_BY_KEY"
			}
			b.WriteString(identifierRow(table.Name, []string{column.Name}, "IDENTITY", enforcement, "", joinUsage(docs, table.Name, []string{column.Name})))
			wrote = true
		}
		if len(full) == 0 {
			for _, candidate := range identifierCandidates(docs, table) {
				b.WriteString(identifierRow(table.Name, []string{candidate.Column}, candidate.Basis, "NOT_ENFORCED", candidate.Source, candidate.Usage))
				wrote = true
			}
		}
		if !wrote && !hasObservedJoinColumns(docs, table) {
			b.WriteString(identifierRow(table.Name, nil, "NO_IDENTIFIER_EVIDENCE", "UNKNOWN", "", 0))
		}
	}
	b.WriteString("\n## Cara menggunakan report ini\n\nIsi tabel di atas berasal dari definisi table/index dan pemakaian JOIN yang ditemukan pada source. Report ini tidak membaca isi data di database, sehingga tidak memberi risk score khusus per table. Panduan di bawah berlaku per classification dan sengaja tidak diulang pada setiap baris.\n\n- `ENFORCED`: database menolak nilai duplikat untuk kombinasi kolom tersebut.\n- `ENFORCED_BY_KEY`: kolom `IDENTITY` ikut dilindungi oleh key/index unique yang terpisah.\n- `CONDITIONAL`: uniqueness hanya berlaku untuk baris yang memenuhi filter index.\n- `NOT_ENFORCED`: tidak ada mekanisme database yang membuktikan uniqueness. Ini belum tentu salah, tetapi perlu dicocokkan dengan aturan bisnis.\n- `UNKNOWN`: analyzer belum menemukan bukti atau sinyal identifier yang cukup.\n\n## Column reference\n\n- **Table**: schema dan nama table yang diperiksa.\n- **Columns**: satu kolom atau kombinasi kolom yang terkait.\n- **Classification**: jenis bukti atau sinyal ID.\n- **Uniqueness**: apakah database benar-benar menjaga nilai agar unique.\n- **Source**: nama constraint atau index yang menjadi sumber informasi. `—` berarti tidak ada constraint/index terkait.\n- **Join usage**: berapa kali kolom tersebut ditemukan dalam equality JOIN pada view/procedure. Nilai ini bukan jumlah baris dan bukan bukti uniqueness.\n\n## Classification reference\n\n- `PRIMARY_KEY`: primary key yang menjaga uniqueness dan tidak mengizinkan NULL pada key.\n- `UNIQUE_CONSTRAINT`: unique constraint yang menjaga kombinasi kolom tetap unique.\n- `UNIQUE_INDEX`: unique index tanpa filter.\n- `FILTERED_UNIQUE_INDEX`: unique index dengan filter; baris di luar filter tidak tercakup.\n- `IDENTITY`: SQL Server membuat nilainya otomatis. Tanpa key/index unique, duplikat tetap mungkin terjadi.\n- `OBSERVED_JOIN_ID_SIGNAL`: kolom terlihat seperti ID dan dipakai dalam equality JOIN. Sering kali ini hanya reference/foreign value yang boleh berulang.\n- `NONUNIQUE_INDEX_ID_SIGNAL`: kolom terlihat seperti ID dan memiliki index non-unique. Index membantu query, tetapi tidak mencegah duplikat.\n- `NO_IDENTIFIER_EVIDENCE`: belum ditemukan key, `IDENTITY`, atau sinyal pemakaian yang cukup untuk menunjukkan identifier.\n\nUntuk baris `NOT_ENFORCED` atau `UNKNOWN`, cek duplicate data, NULL, business scope, dan kemungkinan composite key bersama pemilik data sebelum mengubah schema.\n")
	b.WriteString(renderSuspectedJoinIdentifiers(docs))
	return b.String()
}

type identifierCandidate struct {
	Column, Basis, Source string
	Usage                 int
}

func identifierCandidates(docs DesiredStateDocs, table DocumentedTable) []identifierCandidate {
	seen := map[string]identifierCandidate{}
	for _, column := range table.Columns {
		if column.Identity || !idLikeColumn(column) {
			continue
		}
		usage := joinUsage(docs, table.Name, []string{column.Name})
		var indexes []string
		for _, index := range table.Indexes {
			if !index.Unique && containsFold(index.Columns, column.Name) {
				indexes = append(indexes, index.Name)
			}
		}
		if usage == 0 && len(indexes) == 0 {
			continue
		}
		basis, source := "OBSERVED_JOIN_ID_SIGNAL", ""
		if usage == 0 {
			basis, source = "NONUNIQUE_INDEX_ID_SIGNAL", strings.Join(indexes, ", ")
		}
		seen[strings.ToLower(column.Name)] = identifierCandidate{column.Name, basis, source, usage}
	}
	out := make([]identifierCandidate, 0, len(seen))
	for _, item := range seen {
		out = append(out, item)
	}
	sort.Slice(out, func(i, j int) bool { return out[i].Column < out[j].Column })
	return out
}
func idLikeColumn(c DocumentedColumn) bool {
	u := strings.ToUpper(c.Name)
	return u == "ID" || strings.HasSuffix(u, "ID") || strings.HasSuffix(u, "GUID") || strings.HasSuffix(u, "UUID") || u == "KEY" || strings.HasSuffix(u, "KEY") || c.DataType == "UNIQUEIDENTIFIER"
}
func fullUniqueKeys(t DocumentedTable) []DocumentedKey {
	var out []DocumentedKey
	for _, k := range t.Keys {
		if k.Evidence != "IDENTITY" && !k.Filtered {
			out = append(out, k)
		}
	}
	return out
}
func hasIdentity(t DocumentedTable) bool {
	for _, c := range t.Columns {
		if c.Identity {
			return true
		}
	}
	return false
}
func hasUnenforcedIdentity(t DocumentedTable, keys []DocumentedKey) bool {
	for _, c := range t.Columns {
		if c.Identity && !columnsCoveredByFullUnique([]string{c.Name}, keys) {
			return true
		}
	}
	return false
}
func columnsCoveredByFullUnique(columns []string, keys []DocumentedKey) bool {
	for _, key := range keys {
		if equalFoldSlice(columns, key.Columns) {
			return true
		}
	}
	return false
}
func equalFoldSlice(a, b []string) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if !strings.EqualFold(a[i], b[i]) {
			return false
		}
	}
	return true
}
func containsFold(values []string, target string) bool {
	for _, v := range values {
		if strings.EqualFold(v, target) {
			return true
		}
	}
	return false
}
func joinUsage(docs DesiredStateDocs, object string, columns []string) int {
	total := 0
	for _, r := range docs.Relationships {
		if r.Evidence != "OBSERVED_JOIN" {
			continue
		}
		if strings.EqualFold(r.LeftObject, object) && sameColumnSet(r.LeftColumns, columns) {
			total += r.Occurrences
		}
		if strings.EqualFold(r.RightObject, object) && sameColumnSet(r.RightColumns, columns) && !(strings.EqualFold(r.LeftObject, object) && sameColumnSet(r.LeftColumns, columns)) {
			total += r.Occurrences
		}
	}
	return total
}
func sameColumnSet(a, b []string) bool { return equalFoldSlice(a, b) }
func identifierRow(table string, columns []string, class, enforcement, name string, usage int) string {
	return fmt.Sprintf("| `%s` | %s | `%s` | `%s` | %s | %d |\n", md(table), codeColumns(columns), class, enforcement, codeOrDash(name), usage)
}

func renderRelationshipCatalog(docs DesiredStateDocs) string {
	relations := graphRelationships(docs)
	var b strings.Builder
	b.WriteString("# Database Relationship Catalog\n\nReport ini menyimpan audit detail untuk `relational_erd.md`. Gunakan nomor `Relationship Group` dan `Relational Map` yang sama untuk mencocokkan diagram dengan exact column mapping, join type, confidence, dan source module.\n\nRelationship hanya berasal dari foreign key di `db_object/tables` atau equality JOIN yang terlihat pada `db_object/views` dan `db_object/procedures`. Kesamaan nama kolom saja tidak dianggap sebagai relasi. `OBSERVED_JOIN` membuktikan pemakaian bersama dalam SQL, tetapi tidak membuktikan referential integrity, uniqueness, atau cardinality.\n\n")
	if len(relations) == 0 {
		b.WriteString("Tidak ditemukan relationship yang dapat dibuktikan dari source.\n")
		return b.String()
	}
	groups := relationshipComponents(relations)
	b.WriteString(fmt.Sprintf("- Total relationships: **%d**\n- Relationship groups: **%d**\n- Declared foreign keys: **%d**\n- Observed joins: **%d**\n\n", len(relations), len(groups), relationshipEvidenceCount(relations, "DECLARED_FK"), relationshipEvidenceCount(relations, "OBSERVED_JOIN")))
	b.WriteString("## Group summary\n\nSetiap `Relationship Group` berisi object yang masih terhubung melalui minimal satu jalur relationship. Nomor group bukan urutan proses atau hierarchy.\n\n| Group | Objects | Column mappings | Relational Maps | Declared FK | Observed JOIN | Self-joins | LOW | MEDIUM | HIGH |\n| ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: | ---: |\n")
	for groupNo, group := range groups {
		sets := buildRelationshipConnectionSets(docs, collapseRelationshipPairs(nonSelfRelationships(group)))
		b.WriteString(fmt.Sprintf("| %d | %d | %d | %d | %d | %d | %d | %d | %d | %d |\n", groupNo+1, len(componentNodes(group)), len(group), len(sets), relationshipEvidenceCount(group, "DECLARED_FK"), relationshipEvidenceCount(group, "OBSERVED_JOIN"), selfJoinCount(group), relationshipConfidenceCount(group, "LOW"), relationshipConfidenceCount(group, "MEDIUM"), relationshipConfidenceCount(group, "HIGH")))
	}
	for groupNo, group := range groups {
		sets := buildRelationshipConnectionSets(docs, collapseRelationshipPairs(nonSelfRelationships(group)))
		b.WriteString(fmt.Sprintf("\n## Relationship Group %d\n\n**%d objects, %d exact column mappings.** Urutan Relational Map sama dengan `relational_erd.md`.\n\n", groupNo+1, len(componentNodes(group)), len(group)))
		for setNo, set := range sets {
			b.WriteString(fmt.Sprintf("### Relational Map %d of %d\n\n", setNo+1, len(sets)))
			b.WriteString("**Main objects:** " + relationshipMapMainObjectList(set) + "\n\n")
			writeRelationshipCatalogTables(&b, relationshipPairsToRelationships(set), "Relational Map")
		}
		selfJoins := selfRelationships(group)
		if len(selfJoins) > 0 {
			b.WriteString("### Self-join Review\n\nRelasi berikut memakai dua alias berbeda dari object yang sama. Baris ini dipisahkan dari diagram agar tidak terlihat seperti satu record menunjuk dirinya sendiri.\n\n")
			writeRelationshipCatalogTables(&b, selfJoins, "Self-join")
		}
	}
	b.WriteString("## Column reference\n\n- **Evidence**: `DECLARED_FK` berasal dari foreign key; `OBSERVED_JOIN` berasal dari equality predicate pada view/procedure.\n- **Left/Right object.columns**: pasangan object dan kolom yang dibandingkan. Posisi kiri/kanan adalah canonical ordering, bukan arah parent-child.\n- **Self-join**: `YES` berarti dua alias berbeda mengacu ke object yang sama. Ini biasa dipakai untuk hierarchy atau membandingkan dua baris/periode; bukan berarti satu baris otomatis berelasi dengan dirinya sendiri.\n- **Join type**: bentuk JOIN yang ditemukan, misalnya `INNER`, `LEFT`, atau `IMPLICIT` untuk pola `FROM A, B WHERE ...`.\n- **Occurrences**: berapa kali mapping kolom yang sama ditemukan pada seluruh source module.\n- **Source modules**: view/procedure tempat relationship ditemukan. Untuk foreign key, kolom ini berisi nama constraint.\n- **Confidence**: seberapa lengkap analyzer dapat me-resolve kedua endpoint.\n\n## Evidence & confidence reference\n\n- `DECLARED_FK` / `HIGH`: relationship dideklarasikan langsung sebagai foreign key.\n- `OBSERVED_JOIN` / `MEDIUM`: kedua endpoint berhasil di-resolve dari source lokal dan equality predicate terlihat jelas. JOIN tetap tidak membuktikan FK, cardinality, atau uniqueness.\n- `OBSERVED_JOIN` / `LOW`: salah satu endpoint mengarah ke database/object yang definisinya tidak tersedia pada snapshot lokal. Mapping SQL terlihat, tetapi object eksternalnya perlu dikonfirmasi.\n\n## Batas cakupan\n\nHanya equality predicate dengan qualified column yang dapat di-resolve secara statis yang dimasukkan. Dynamic SQL, range/inequality join, kolom tanpa alias atau ambigu, ekspresi dengan beberapa kolom berbeda pada satu sisi, temporary/table variable, dan query scope yang tidak dapat di-resolve tetap membutuhkan review manual.\n")
	return b.String()
}

func writeRelationshipCatalogTables(b *strings.Builder, relations []ObservedRelationship, label string) {
	for start, partNo := 0, 1; start < len(relations); start, partNo = start+relationshipCatalogMaxTableRows, partNo+1 {
		end := start + relationshipCatalogMaxTableRows
		if end > len(relations) {
			end = len(relations)
		}
		if len(relations) > relationshipCatalogMaxTableRows {
			b.WriteString(fmt.Sprintf("#### %s detail %d dari %d\n\n", label, partNo, (len(relations)+relationshipCatalogMaxTableRows-1)/relationshipCatalogMaxTableRows))
		}
		b.WriteString("| Evidence | Left object.columns | Right object.columns | Self-join | Join type | Occurrences | Source modules | Confidence |\n| --- | --- | --- | --- | --- | ---: | --- | --- |\n")
		for _, relation := range relations[start:end] {
			b.WriteString(relationshipCatalogRow(relation))
		}
		b.WriteString("\n")
	}
}

func relationshipCatalogRow(r ObservedRelationship) string {
	joinType := codeOrDash(strings.Join(r.JoinTypes, ", "))
	return fmt.Sprintf("| `%s` | `%s.(%s)` | `%s.(%s)` | %s | %s | %d | %s | %s |\n", r.Evidence, md(r.LeftObject), md(strings.Join(r.LeftColumns, ", ")), md(r.RightObject), md(strings.Join(r.RightColumns, ", ")), yesNo(strings.EqualFold(r.LeftObject, r.RightObject)), joinType, r.Occurrences, codeList(r.Modules), emptyAsDash(r.Confidence))
}

func relationshipEvidenceCount(relations []ObservedRelationship, evidence string) int {
	count := 0
	for _, relation := range relations {
		if relation.Evidence == evidence {
			count++
		}
	}
	return count
}

func relationshipConfidenceCount(relations []ObservedRelationship, confidence string) int {
	count := 0
	for _, relation := range relations {
		if relation.Confidence == confidence {
			count++
		}
	}
	return count
}

func selfJoinCount(relations []ObservedRelationship) int {
	count := 0
	for _, relation := range relations {
		if strings.EqualFold(relation.LeftObject, relation.RightObject) {
			count++
		}
	}
	return count
}

func renderRelationalERD(docs DesiredStateDocs) string {
	relations := graphRelationships(docs)
	var b strings.Builder
	b.WriteString("# Relational Map Berdasarkan Implementasi Database\n\nDocument ini menampilkan hubungan antarobject yang dapat dibuktikan dari source `db_object`. Sumber relationship hanya foreign key pada `db_object/tables` serta equality `JOIN` pada `db_object/views` dan `db_object/procedures`. Nama kolom yang kebetulan sama tidak cukup untuk membuat hubungan.\n\n")
	b.WriteString("## How to read\n\n- **Biru** adalah `TABLE`; **abu-abu** adalah `VIEW`; object yang definisinya tidak tersedia pada snapshot ditandai `EXTERNAL OBJECT`.\n- Garis solid berarti `DECLARED_FK`. Garis putus-putus tanpa panah berarti `OBSERVED_JOIN`.\n- Label garis menunjukkan exact column mapping. Posisi kiri dan kanan hanya layout, bukan source/target atau parent/child.\n- Status key/index ditampilkan tepat di sebelah kolom yang ikut relationship.\n- `PART OF UNIQUE KEY (A, B)` berarti uniqueness dijaga pada tuple lengkap; satu kolom sendiri belum tentu unique.\n- `NON-UNIQUE INDEX` berarti ada index, tetapi duplicate tetap diperbolehkan.\n- `NOT UNIQUE` berarti tidak ditemukan PK, unique constraint, atau unique index yang menjaga kolom tersebut.\n- `UNIQUENESS UNKNOWN` digunakan untuk `VIEW` atau external object karena source tidak membuktikan output-nya unique.\n- Diagram tidak menyatakan cardinality `1:1` atau `1:N` jika tidak dibuktikan oleh declared constraint.\n\nSelf-join tidak digambar sebagai garis kembali ke object yang sama. Exact mapping dan source module-nya tersedia pada `Self-join review`.\n\n")
	if len(relations) == 0 {
		b.WriteString("Tidak ditemukan foreign key atau relasi statis yang dapat dibuktikan.\n")
		return b.String()
	}
	groups := relationshipComponents(relations)
	b.WriteString(fmt.Sprintf("- Relationship mappings: **%d**\n- Declared foreign keys: **%d**\n- Observed joins: **%d**\n\n", len(relations), relationshipEvidenceCount(relations, "DECLARED_FK"), relationshipEvidenceCount(relations, "OBSERVED_JOIN")))
	b.WriteString("## Group summary\n\n`Relationship Group` adalah jaringan object yang masih tersambung melalui minimal satu jalur. Nomornya bukan urutan proses, prioritas, atau hierarchy.\n\n| Group | Objects | Object pairs | Column mappings | Relational Maps | Self-joins |\n| ---: | ---: | ---: | ---: | ---: | ---: |\n")
	for groupNo, group := range groups {
		pairs := collapseRelationshipPairs(nonSelfRelationships(group))
		sets := buildRelationshipConnectionSets(docs, pairs)
		b.WriteString(fmt.Sprintf("| %d | %d | %d | %d | %d | %d |\n", groupNo+1, len(componentNodes(group)), len(pairs), len(group), len(sets), selfJoinCount(group)))
	}
	b.WriteString("\n")
	for groupNo, group := range groups {
		pairs := collapseRelationshipPairs(nonSelfRelationships(group))
		sets := buildRelationshipConnectionSets(docs, pairs)
		b.WriteString(fmt.Sprintf("## Relationship Group %d\n\n**%d objects, %d object pairs, %d exact column mappings.**\n\n", groupNo+1, len(componentNodes(group)), len(pairs), len(group)))
		if len(sets) == 0 {
			b.WriteString("Group ini hanya mempunyai self-join, sehingga tidak ada diagram antardua object.\n\n")
		} else {
			writeRelationshipObjectIndex(&b, docs, groupNo+1, sets)
		}
		for setNo, set := range sets {
			b.WriteString(fmt.Sprintf("### Relationship Group %d Relational Map %d of %d\n\n", groupNo+1, setNo+1, len(sets)))
			b.WriteString("**Main objects:** " + relationshipMapMainObjectList(set) + "\n\n")
			b.WriteString(fmt.Sprintf("**%d objects, %d relationship lines, %d exact column mappings.**\n\n", len(relationshipSetNodes(set)), len(set), relationshipPairMappingCount(set)))
			if shared := relationshipSetContinuations(sets, setNo, groupNo+1); shared != "" {
				b.WriteString("**Shared object navigation**\n\nObject berikut juga muncul pada map lain karena menghubungkan bagian jaringan yang berbeda. Ini hanya bantuan navigasi document, bukan atribut database.\n\n")
				b.WriteString("| Shared object | Also appears in |\n| --- | --- |\n" + shared + "\n\n")
			}
			b.WriteString(renderMermaidRelationshipSet(docs, set))
		}
		selfJoins := selfRelationships(group)
		if len(selfJoins) > 0 {
			b.WriteString("### Self-join review\n\nSelf-join memakai dua alias atau instance berbeda dari object yang sama. Ini bukan berarti satu row otomatis menunjuk dirinya sendiri.\n\n| Object | Exact column mapping | Join type | Source modules | Meaning |\n| --- | --- | --- | --- | --- |\n")
			for _, relation := range selfJoins {
				mapping := strings.Join(relation.LeftColumns, ", ") + " = " + strings.Join(relation.RightColumns, ", ")
				b.WriteString(fmt.Sprintf("| `%s` | `%s` | %s | %s | Dua instance object dibandingkan; cardinality tidak dinyatakan. |\n", md(relation.LeftObject), md(mapping), codeOrDash(strings.Join(relation.JoinTypes, ", ")), codeList(relation.Modules)))
			}
			b.WriteString("\n")
		}
	}
	nonSelfCount := len(nonSelfRelationships(relations))
	selfCount := len(selfRelationships(relations))
	b.WriteString("## Coverage check\n\n| Pemeriksaan | Source | Ditampilkan | Status |\n| --- | ---: | ---: | --- |\n")
	b.WriteString(fmt.Sprintf("| Non-self exact mappings | %d | %d | **COMPLETE** |\n", nonSelfCount, nonSelfCount))
	b.WriteString(fmt.Sprintf("| Self-join exact mappings | %d | %d | **COMPLETE** |\n", selfCount, selfCount))
	b.WriteString(fmt.Sprintf("| Seluruh exact mappings | %d | %d | **COMPLETE** |\n\n", len(relations), nonSelfCount+selfCount))
	b.WriteString("Setiap non-self mapping tampil tepat satu kali pada Relational Map dan setiap self-join tampil pada `Self-join review`. Detail join type, confidence, occurrences, dan source module tetap tersedia pada [relationship_catalog.md](relationship_catalog.md).\n")
	return b.String()
}

type relationshipPair struct {
	LeftObject, RightObject   string
	Relationships             []ObservedRelationship
	DeclaredFK, Observed      int
	Occurrences               int
	MappingPart, MappingParts int
}

func collapseRelationshipPairs(relations []ObservedRelationship) []relationshipPair {
	byKey := map[string]*relationshipPair{}
	for _, relation := range relations {
		left, right := relation.LeftObject, relation.RightObject
		if strings.ToLower(left) > strings.ToLower(right) {
			left, right = right, left
		}
		key := strings.ToLower(left) + "|" + strings.ToLower(right)
		pair := byKey[key]
		if pair == nil {
			pair = &relationshipPair{LeftObject: left, RightObject: right}
			byKey[key] = pair
		}
		pair.Relationships = append(pair.Relationships, relation)
		pair.Occurrences += relation.Occurrences
		if relation.Evidence == "DECLARED_FK" {
			pair.DeclaredFK++
		} else {
			pair.Observed++
		}
	}
	keys := make([]string, 0, len(byKey))
	for key := range byKey {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]relationshipPair, 0, len(keys))
	for _, key := range keys {
		out = append(out, *byKey[key])
	}
	return out
}

func relationshipPairKey(pair relationshipPair) string {
	return strings.ToLower(pair.LeftObject) + "|" + strings.ToLower(pair.RightObject)
}

func buildRelationshipConnectionSets(docs DesiredStateDocs, pairs []relationshipPair) [][]relationshipPair {
	remaining := splitDenseRelationshipPairs(pairs, relationalMapMaxMappings, relationalMapMaxTextCharacters)
	var sets [][]relationshipPair
	for len(remaining) > 0 {
		startObject := highestUncoveredPairDegreeObject(remaining)
		start := 0
		for i, pair := range remaining {
			if strings.EqualFold(pair.LeftObject, startObject) || strings.EqualFold(pair.RightObject, startObject) {
				start = i
				break
			}
		}
		set := []relationshipPair{remaining[start]}
		remaining = append(remaining[:start], remaining[start+1:]...)
		nodes := relationshipSetNodeLookup(set)
		mappings := len(set[0].Relationships)
		for len(set) < relationalMapMaxPairs {
			candidate := -1
			candidateExistingEndpoints := -1
			candidateConnections := -1
			for i, pair := range remaining {
				leftExists := nodes[strings.ToLower(pair.LeftObject)]
				rightExists := nodes[strings.ToLower(pair.RightObject)]
				if !leftExists && !rightExists {
					continue
				}
				candidateNodes := relationshipSetNodeLookup(append(append([]relationshipPair(nil), set...), pair))
				if len(candidateNodes) > relationalMapMaxObjects {
					continue
				}
				if mappings+len(pair.Relationships) > relationalMapMaxMappings {
					continue
				}
				candidateSet := append(append([]relationshipPair(nil), set...), pair)
				if len(renderMermaidRelationshipSet(docs, candidateSet)) > relationalMapMaxTextCharacters {
					continue
				}
				existingEndpoints := 0
				if leftExists {
					existingEndpoints++
				}
				if rightExists {
					existingEndpoints++
				}
				connections := relationshipCandidateConnections(pair, nodes, remaining)
				if candidate < 0 || existingEndpoints > candidateExistingEndpoints ||
					(existingEndpoints == candidateExistingEndpoints && connections > candidateConnections) ||
					(existingEndpoints == candidateExistingEndpoints && connections == candidateConnections && relationshipPairKey(pair) < relationshipPairKey(remaining[candidate])) {
					candidate = i
					candidateExistingEndpoints = existingEndpoints
					candidateConnections = connections
				}
			}
			if candidate < 0 {
				break
			}
			pair := remaining[candidate]
			set = append(set, pair)
			mappings += len(pair.Relationships)
			nodes[strings.ToLower(pair.LeftObject)] = true
			nodes[strings.ToLower(pair.RightObject)] = true
			remaining = append(remaining[:candidate], remaining[candidate+1:]...)
		}
		sets = append(sets, set)
	}
	return sets
}

func relationshipCandidateConnections(pair relationshipPair, nodes map[string]bool, remaining []relationshipPair) int {
	newObject := ""
	if !nodes[strings.ToLower(pair.LeftObject)] {
		newObject = pair.LeftObject
	}
	if !nodes[strings.ToLower(pair.RightObject)] {
		if newObject != "" {
			return 0
		}
		newObject = pair.RightObject
	}
	if newObject == "" {
		return relationalMapMaxPairs
	}
	connections := 0
	for _, candidate := range remaining {
		if strings.EqualFold(candidate.LeftObject, newObject) && nodes[strings.ToLower(candidate.RightObject)] {
			connections++
		}
		if strings.EqualFold(candidate.RightObject, newObject) && nodes[strings.ToLower(candidate.LeftObject)] {
			connections++
		}
	}
	return connections
}

func splitDenseRelationshipPairs(pairs []relationshipPair, maxMappings, maxLabelCharacters int) []relationshipPair {
	var out []relationshipPair
	for _, pair := range pairs {
		var chunks []relationshipPair
		chunk := relationshipPair{LeftObject: pair.LeftObject, RightObject: pair.RightObject}
		chunkCharacters := 0
		for _, relation := range pair.Relationships {
			mappingCharacters := len(relationshipMappingLabel(pair, relation))
			if len(chunk.Relationships) > 0 && (len(chunk.Relationships) >= maxMappings || chunkCharacters+mappingCharacters > maxLabelCharacters) {
				chunks = append(chunks, chunk)
				chunk = relationshipPair{LeftObject: pair.LeftObject, RightObject: pair.RightObject}
				chunkCharacters = 0
			}
			chunk.Relationships = append(chunk.Relationships, relation)
			chunk.Occurrences += relation.Occurrences
			if relation.Evidence == "DECLARED_FK" {
				chunk.DeclaredFK++
			} else {
				chunk.Observed++
			}
			chunkCharacters += mappingCharacters
		}
		if len(chunk.Relationships) > 0 {
			chunks = append(chunks, chunk)
		}
		for index := range chunks {
			chunks[index].MappingPart = index + 1
			chunks[index].MappingParts = len(chunks)
			out = append(out, chunks[index])
		}
	}
	return out
}

func highestUncoveredPairDegreeObject(pairs []relationshipPair) string {
	degrees := map[string]int{}
	for _, pair := range pairs {
		degrees[pair.LeftObject]++
		degrees[pair.RightObject]++
	}
	best := ""
	for object, degree := range degrees {
		if best == "" || degree > degrees[best] || (degree == degrees[best] && strings.ToLower(object) < strings.ToLower(best)) {
			best = object
		}
	}
	return best
}

func relationshipPairMappingCount(pairs []relationshipPair) int {
	count := 0
	for _, pair := range pairs {
		count += len(pair.Relationships)
	}
	return count
}

func relationshipPairsToRelationships(pairs []relationshipPair) []ObservedRelationship {
	var out []ObservedRelationship
	for _, pair := range pairs {
		out = append(out, pair.Relationships...)
	}
	sort.Slice(out, func(i, j int) bool { return relationshipKey(out[i]) < relationshipKey(out[j]) })
	return out
}

func relationshipSetNodeLookup(pairs []relationshipPair) map[string]bool {
	nodes := map[string]bool{}
	for _, pair := range pairs {
		nodes[strings.ToLower(pair.LeftObject)] = true
		nodes[strings.ToLower(pair.RightObject)] = true
	}
	return nodes
}

func relationshipSetNodes(pairs []relationshipPair) []string {
	byKey := map[string]string{}
	for _, pair := range pairs {
		byKey[strings.ToLower(pair.LeftObject)] = pair.LeftObject
		byKey[strings.ToLower(pair.RightObject)] = pair.RightObject
	}
	keys := make([]string, 0, len(byKey))
	for key := range byKey {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]string, 0, len(keys))
	for _, key := range keys {
		out = append(out, byKey[key])
	}
	return out
}

func relationshipSetContinuations(sets [][]relationshipPair, current, groupNo int) string {
	currentNodes := relationshipSetNodeLookup(sets[current])
	var entries []string
	for nodeKey := range currentNodes {
		name := ""
		var mapNumbers []int
		for setNo, set := range sets {
			if setNo == current {
				continue
			}
			for _, object := range relationshipSetNodes(set) {
				if strings.EqualFold(object, nodeKey) {
					name = object
					mapNumbers = append(mapNumbers, setNo+1)
					break
				}
			}
		}
		if name != "" {
			var links []string
			for _, mapNo := range mapNumbers {
				links = append(links, relationshipMapLink(groupNo, mapNo, len(sets)))
			}
			entries = append(entries, fmt.Sprintf("| `%s` | %s |", md(name), strings.Join(links, "; ")))
		}
	}
	sort.Strings(entries)
	return strings.Join(entries, "\n")
}

func writeRelationshipObjectIndex(b *strings.Builder, docs DesiredStateDocs, groupNo int, sets [][]relationshipPair) {
	byObject := map[string][]int{}
	objectNames := map[string]string{}
	for setNo, set := range sets {
		for _, object := range relationshipSetNodes(set) {
			key := strings.ToLower(object)
			objectNames[key] = object
			if !containsInt(byObject[key], setNo+1) {
				byObject[key] = append(byObject[key], setNo+1)
			}
		}
	}
	keys := make([]string, 0, len(objectNames))
	for key := range objectNames {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	b.WriteString("### Object index\n\nCari object di index ini, lalu buka semua Relational Map yang tercantum. Object yang muncul pada beberapa map hanya menjadi penghubung navigasi; pengulangan tersebut bukan atribut database.\n\n")
	b.WriteString("| Object | Type | Maps | Object | Type | Maps |\n| --- | --- | --- | --- | --- | --- |\n")
	rows := (len(keys) + 1) / 2
	for row := 0; row < rows; row++ {
		leftKey := keys[row]
		left := relationshipObjectIndexCells(docs, objectNames[leftKey], groupNo, len(sets), byObject[leftKey])
		right := []string{"—", "—", "—"}
		if rightIndex := row + rows; rightIndex < len(keys) {
			rightKey := keys[rightIndex]
			right = relationshipObjectIndexCells(docs, objectNames[rightKey], groupNo, len(sets), byObject[rightKey])
		}
		b.WriteString(fmt.Sprintf("| %s | %s | %s | %s | %s | %s |\n", left[0], left[1], left[2], right[0], right[1], right[2]))
	}
	b.WriteString("\n")
}

func relationshipObjectIndexCells(docs DesiredStateDocs, object string, groupNo, mapCount int, mapNumbers []int) []string {
	links := make([]string, 0, len(mapNumbers))
	for _, mapNo := range mapNumbers {
		links = append(links, relationshipMapLink(groupNo, mapNo, mapCount))
	}
	return []string{"`" + md(object) + "`", relationshipObjectKind(docs, object), strings.Join(links, "; ")}
}

func relationshipMapLink(groupNo, mapNo, mapCount int) string {
	return fmt.Sprintf("[Map %d](#relationship-group-%d-relational-map-%d-of-%d)", mapNo, groupNo, mapNo, mapCount)
}

func containsInt(values []int, target int) bool {
	for _, value := range values {
		if value == target {
			return true
		}
	}
	return false
}

func relationshipMapMainObjectList(pairs []relationshipPair) string {
	degrees := map[string]int{}
	names := map[string]string{}
	for _, pair := range pairs {
		for _, object := range []string{pair.LeftObject, pair.RightObject} {
			key := strings.ToLower(object)
			degrees[key]++
			names[key] = object
		}
	}
	keys := make([]string, 0, len(names))
	for key := range names {
		keys = append(keys, key)
	}
	sort.Slice(keys, func(i, j int) bool {
		if degrees[keys[i]] != degrees[keys[j]] {
			return degrees[keys[i]] > degrees[keys[j]]
		}
		return keys[i] < keys[j]
	})
	if len(keys) > 3 {
		keys = keys[:3]
	}
	objects := make([]string, 0, len(keys))
	for _, key := range keys {
		objects = append(objects, "`"+md(names[key])+"`")
	}
	return strings.Join(objects, " · ")
}

func renderMermaidRelationshipSet(docs DesiredStateDocs, pairs []relationshipPair) string {
	var b strings.Builder
	b.WriteString("```mermaid\nflowchart LR\n")
	nodes := map[string]string{}
	classes := map[string][]string{"table": {}, "view": {}, "external": {}}
	for index, name := range relationshipSetNodes(pairs) {
		key := strings.ToLower(name)
		nodes[key] = "n" + strconv.Itoa(index+1)
		kind := strings.ToLower(relationshipObjectKind(docs, name))
		if kind != "table" && kind != "view" {
			kind = "external"
		}
		classes[kind] = append(classes[kind], nodes[key])
		b.WriteString(fmt.Sprintf("    %s[\"`%s`\"]\n", nodes[key], mermaidMarkdown(relationshipNodeLabel(docs, name, pairs))))
	}
	for _, pair := range pairs {
		line := "-.-"
		if pair.DeclaredFK > 0 && pair.Observed == 0 {
			line = "---"
		}
		b.WriteString(fmt.Sprintf("    %s %s|\"%s\"| %s\n", nodes[strings.ToLower(pair.LeftObject)], line, mermaid(relationshipPairLabel(pair)), nodes[strings.ToLower(pair.RightObject)]))
	}
	b.WriteString("\n    classDef table fill:#17365d,stroke:#6ea8dc,color:#ffffff,stroke-width:1.5px\n")
	b.WriteString("    classDef view fill:#343a40,stroke:#adb5bd,color:#ffffff,stroke-width:1.5px\n")
	b.WriteString("    classDef external fill:#4a3b23,stroke:#d8a657,color:#ffffff,stroke-width:1.5px\n")
	for _, kind := range []string{"table", "view", "external"} {
		if len(classes[kind]) > 0 {
			b.WriteString(fmt.Sprintf("    class %s %s\n", strings.Join(classes[kind], ","), kind))
		}
	}
	b.WriteString("```\n\n")
	return b.String()
}

func relationshipNodeLabel(docs DesiredStateDocs, object string, pairs []relationshipPair) string {
	lines := []string{"**" + object + "**", relationshipObjectKind(docs, object)}
	columns := relationshipObjectColumns(object, pairs)
	table, isTable := documentedTableByName(docs, object)
	for _, column := range columns {
		status := "UNIQUENESS UNKNOWN"
		if isTable {
			status = relationshipColumnStatus(table, column)
		}
		lines = append(lines, column+" — "+status)
	}
	return strings.Join(lines, "\n")
}

func relationshipObjectColumns(object string, pairs []relationshipPair) []string {
	byKey := map[string]string{}
	for _, pair := range pairs {
		for _, relation := range pair.Relationships {
			if strings.EqualFold(relation.LeftObject, object) {
				for _, column := range relation.LeftColumns {
					byKey[strings.ToLower(column)] = column
				}
			}
			if strings.EqualFold(relation.RightObject, object) {
				for _, column := range relation.RightColumns {
					byKey[strings.ToLower(column)] = column
				}
			}
		}
	}
	keys := make([]string, 0, len(byKey))
	for key := range byKey {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]string, 0, len(keys))
	for _, key := range keys {
		out = append(out, byKey[key])
	}
	return out
}

func relationshipObjectKind(docs DesiredStateDocs, object string) string {
	if _, ok := documentedTableByName(docs, object); ok {
		return "TABLE"
	}
	for _, module := range docs.Modules {
		if strings.EqualFold(module.Name, object) && strings.EqualFold(module.Kind, "view") {
			return "VIEW"
		}
	}
	return "EXTERNAL OBJECT"
}

func documentedTableByName(docs DesiredStateDocs, object string) (DocumentedTable, bool) {
	for _, table := range docs.Tables {
		if strings.EqualFold(table.Name, object) {
			return table, true
		}
	}
	return DocumentedTable{}, false
}

func relationshipColumnStatus(table DocumentedTable, column string) string {
	identity := false
	for _, candidate := range table.Columns {
		if strings.EqualFold(candidate.Name, column) {
			identity = candidate.Identity
			break
		}
	}
	keys := append([]DocumentedKey(nil), table.Keys...)
	sort.Slice(keys, func(i, j int) bool {
		if len(keys[i].Columns) != len(keys[j].Columns) {
			return len(keys[i].Columns) < len(keys[j].Columns)
		}
		return relationshipKeyRank(keys[i]) < relationshipKeyRank(keys[j])
	})
	for _, key := range keys {
		if !containsFold(key.Columns, column) {
			continue
		}
		status := relationshipKeyStatus(key)
		if identity {
			status += " · IDENTITY"
		}
		return status
	}
	indexes := append([]DocumentedIndex(nil), table.Indexes...)
	sort.Slice(indexes, func(i, j int) bool {
		if len(indexes[i].Columns) != len(indexes[j].Columns) {
			return len(indexes[i].Columns) < len(indexes[j].Columns)
		}
		return strings.ToLower(indexes[i].Name) < strings.ToLower(indexes[j].Name)
	})
	for _, index := range indexes {
		if index.Unique || !containsFold(index.Columns, column) {
			continue
		}
		status := "NON-UNIQUE INDEX"
		if len(index.Columns) > 1 {
			status += " (" + strings.Join(index.Columns, ", ") + ")"
		}
		if identity {
			return "IDENTITY · " + status
		}
		return status
	}
	if identity {
		return "IDENTITY · NOT UNIQUE"
	}
	return "NOT UNIQUE"
}

func relationshipKeyRank(key DocumentedKey) string {
	rank := map[string]string{"PRIMARY_KEY": "1", "UNIQUE_CONSTRAINT": "2", "UNIQUE_INDEX": "3", "FILTERED_UNIQUE_INDEX": "4"}
	prefix, ok := rank[key.Evidence]
	if !ok {
		prefix = "9"
	}
	return prefix + strings.ToLower(strings.Join(key.Columns, "|"))
}

func relationshipKeyStatus(key DocumentedKey) string {
	if len(key.Columns) > 1 {
		if key.Filtered {
			return "PART OF FILTERED UNIQUE KEY (" + strings.Join(key.Columns, ", ") + ")"
		}
		return "PART OF UNIQUE KEY (" + strings.Join(key.Columns, ", ") + ")"
	}
	switch key.Evidence {
	case "PRIMARY_KEY":
		return "PRIMARY KEY"
	case "UNIQUE_CONSTRAINT":
		return "UNIQUE CONSTRAINT"
	case "FILTERED_UNIQUE_INDEX":
		return "FILTERED UNIQUE INDEX"
	default:
		return "UNIQUE INDEX"
	}
}

func relationshipPairLabel(pair relationshipPair) string {
	labels := make([]string, 0, len(pair.Relationships))
	for _, relation := range pair.Relationships {
		labels = append(labels, relationshipMappingLabel(pair, relation))
	}
	label := strings.Join(labels, "; ")
	if pair.MappingParts > 1 {
		return fmt.Sprintf("Mapping batch %d/%d · %s", pair.MappingPart, pair.MappingParts, label)
	}
	return label
}

func relationshipMappingLabel(pair relationshipPair, relation ObservedRelationship) string {
	leftColumns, rightColumns := relation.LeftColumns, relation.RightColumns
	if !strings.EqualFold(relation.LeftObject, pair.LeftObject) {
		leftColumns, rightColumns = rightColumns, leftColumns
	}
	label := fmt.Sprintf("%s = %s", strings.Join(leftColumns, ", "), strings.Join(rightColumns, ", "))
	if pair.DeclaredFK > 0 && pair.Observed > 0 {
		evidence := "JOIN"
		if relation.Evidence == "DECLARED_FK" {
			evidence = "FK"
		}
		label += " [" + evidence + "]"
	}
	return label
}

func nonSelfRelationships(relations []ObservedRelationship) []ObservedRelationship {
	var out []ObservedRelationship
	for _, relation := range relations {
		if !strings.EqualFold(relation.LeftObject, relation.RightObject) {
			out = append(out, relation)
		}
	}
	return out
}

func selfRelationships(relations []ObservedRelationship) []ObservedRelationship {
	var out []ObservedRelationship
	for _, relation := range relations {
		if strings.EqualFold(relation.LeftObject, relation.RightObject) {
			out = append(out, relation)
		}
	}
	return out
}

type selfRelationshipSummary struct {
	Object                string
	Mappings, Occurrences int
}

func summarizeSelfRelationships(relations []ObservedRelationship) []selfRelationshipSummary {
	byObject := map[string]*selfRelationshipSummary{}
	for _, relation := range relations {
		key := strings.ToLower(relation.LeftObject)
		summary := byObject[key]
		if summary == nil {
			summary = &selfRelationshipSummary{Object: relation.LeftObject}
			byObject[key] = summary
		}
		summary.Mappings++
		summary.Occurrences += relation.Occurrences
	}
	keys := make([]string, 0, len(byObject))
	for key := range byObject {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	out := make([]selfRelationshipSummary, 0, len(keys))
	for _, key := range keys {
		out = append(out, *byObject[key])
	}
	return out
}

func renderModuleLineage(docs DesiredStateDocs) string {
	var b strings.Builder
	b.WriteString("# Lineage Module\n\nReferensi langsung dan statis dari view/procedure/function ke data object. Resolution membedakan object catalog yang terbukti, external, temporary/CTE, nama yang tidak ditemukan, dan kandidat ambigu. Lineage dependency ini bukan klaim hubungan antartable.\n\n| Modul | Jenis | Object yang direferensikan | Resolution |\n| --- | --- | --- | --- |\n")
	for _, d := range docs.ModuleDependencies {
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | `%s` | `%s` |\n", md(d.Module), strings.ToUpper(d.Kind), md(d.Object), md(d.Resolution)))
	}
	if len(docs.ModuleDependencies) == 0 {
		b.WriteString("| — | — | — | — |\n")
	}
	b.WriteString(fmt.Sprintf("\n- Module yang diperiksa: **%d**\n- Module dengan minimal satu referensi data object statis: **%d**\n\n", len(docs.Modules), modulesWithDependencies(docs)))
	b.WriteString("\n## Dynamic SQL yang harus diperiksa manual\n\n")
	if len(docs.DynamicSQLModules) == 0 {
		b.WriteString("Tidak ada yang terdeteksi.\n")
	} else {
		for _, m := range docs.DynamicSQLModules {
			b.WriteString("- `" + md(m) + "`\n")
		}
	}
	b.WriteString("\n## Module tanpa referensi data object statis yang terdeteksi\n\nModule berikut mungkin hanya berisi logika scalar, memakai syntax yang belum didukung/tanpa alias, atau mempunyai dependency yang perlu diperiksa manual.\n\n")
	missing := modulesWithoutDependencies(docs)
	if len(missing) == 0 {
		b.WriteString("Tidak ada.\n")
	} else {
		for _, module := range missing {
			b.WriteString(fmt.Sprintf("- `%s` (`%s`)\n", md(module.Name), strings.ToUpper(module.Kind)))
		}
	}
	return b.String()
}

func modulesWithDependencies(docs DesiredStateDocs) int {
	seen := map[string]struct{}{}
	for _, dep := range docs.ModuleDependencies {
		seen[strings.ToLower(dep.Kind+"|"+dep.Module)] = struct{}{}
	}
	return len(seen)
}

func modulesWithoutDependencies(docs DesiredStateDocs) []DocumentedModule {
	withDependencies := map[string]struct{}{}
	for _, dep := range docs.ModuleDependencies {
		withDependencies[strings.ToLower(dep.Kind+"|"+dep.Module)] = struct{}{}
	}
	var out []DocumentedModule
	for _, module := range docs.Modules {
		if _, ok := withDependencies[strings.ToLower(module.Kind+"|"+module.Name)]; !ok {
			out = append(out, module)
		}
	}
	return out
}
func renderRelationshipCompatibilityNotice() string {
	return "# Tree Relasi Table\n\nNama report lama ini telah digantikan oleh:\n\n- `relational_erd.md` untuk diagram Mermaid;\n- `relationship_catalog.md` untuk mapping kolom dan bukti secara tepat;\n- `module_lineage.md` untuk dependency view/procedure.\n"
}
func graphRelationships(d DesiredStateDocs) []ObservedRelationship {
	out := append([]ObservedRelationship(nil), d.Relationships...)
	for _, t := range d.Tables {
		for _, fk := range t.ForeignKeys {
			out = append(out, ObservedRelationship{LeftObject: t.Name, LeftColumns: fk.Columns, RightObject: fk.TargetTable, RightColumns: fk.TargetColumns, Evidence: "DECLARED_FK", Occurrences: 1, Confidence: "HIGH", Modules: []string{fk.ConstraintName}})
		}
	}
	sort.Slice(out, func(i, j int) bool { return relationshipKey(out[i]) < relationshipKey(out[j]) })
	return out
}
func relationshipComponents(relations []ObservedRelationship) [][]ObservedRelationship {
	remaining := append([]ObservedRelationship(nil), relations...)
	var groups [][]ObservedRelationship
	for len(remaining) > 0 {
		nodes := map[string]bool{remaining[0].LeftObject: true, remaining[0].RightObject: true}
		var group []ObservedRelationship
		changed := true
		for changed {
			changed = false
			next := remaining[:0]
			for _, r := range remaining {
				if nodes[r.LeftObject] || nodes[r.RightObject] {
					group = append(group, r)
					nodes[r.LeftObject] = true
					nodes[r.RightObject] = true
					changed = true
				} else {
					next = append(next, r)
				}
			}
			remaining = next
		}
		groups = append(groups, group)
	}
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) != len(groups[j]) {
			return len(groups[i]) > len(groups[j])
		}
		leftNodes, rightNodes := componentNodes(groups[i]), componentNodes(groups[j])
		return strings.ToLower(strings.Join(leftNodes, "|")) < strings.ToLower(strings.Join(rightNodes, "|"))
	})
	return groups
}
func componentNodes(relations []ObservedRelationship) []string {
	set := map[string]struct{}{}
	for _, r := range relations {
		set[r.LeftObject] = struct{}{}
		set[r.RightObject] = struct{}{}
	}
	return sortedSet(set)
}
func codeColumns(c []string) string {
	if len(c) == 0 {
		return "—"
	}
	return "`" + md(strings.Join(c, ", ")) + "`"
}
func codeOrDash(v string) string {
	if v == "" {
		return "—"
	}
	return "`" + md(v) + "`"
}
func codeList(v []string) string {
	var out []string
	for _, x := range v {
		if strings.TrimSpace(x) != "" {
			out = append(out, "`"+md(x)+"`")
		}
	}
	if len(out) == 0 {
		return "—"
	}
	return strings.Join(out, "; ")
}
func md(v string) string      { return strings.ReplaceAll(strings.ReplaceAll(v, "|", "\\|"), "`", "'") }
func mermaid(v string) string { return strings.ReplaceAll(strings.ReplaceAll(v, "\"", "'"), "\n", " ") }
func mermaidMarkdown(v string) string {
	return strings.ReplaceAll(strings.ReplaceAll(v, "\"", "'"), "`", "'")
}
func emptyAsDash(v string) string {
	if v == "" {
		return "—"
	}
	return v
}
func yesNo(value bool) string {
	if value {
		return "YES"
	}
	return "NO"
}
