package migreport

import (
	"fmt"
	"net/url"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"rza_migrate_db/internal/sqlparse"
)

const procedureDocsDirectory = "procedures"

func writeProcedureDocs(docsDir string, docs DesiredStateDocs) error {
	root := filepath.Join(docsDir, procedureDocsDirectory)
	profilesDir := filepath.Join(root, "profiles")
	if err := os.MkdirAll(profilesDir, 0755); err != nil {
		return fmt.Errorf("create procedure documentation directory: %w", err)
	}
	if err := writeGeneratedMarkdown(filepath.Join(root, "README.md"), renderProcedureIndex(docs)+"\n[Database documentation](../README.md)\n"); err != nil {
		return err
	}
	if err := writeGeneratedMarkdown(filepath.Join(root, "procedure_impact_catalog.md"), renderProcedureCatalog(docs)+"\n[Database documentation](../README.md)\n"); err != nil {
		return err
	}
	for i := range docs.ProcedureProfiles {
		name := docs.ProcedureProfiles[i].Name + ".md"
		if err := writeGeneratedMarkdown(filepath.Join(profilesDir, name), renderProcedureProfile(docs, &docs.ProcedureProfiles[i])+"\n[Database documentation](../../README.md)\n"); err != nil {
			return err
		}
	}
	return nil
}

func renderProcedureIndex(docs DesiredStateDocs) string {
	var b strings.Builder
	b.WriteString("# Procedure Object & Process Profiles\n\n")
	b.WriteString("Document ini adalah entry point untuk memahami object, operasi data, procedure call, recursive impact, transaction/error signal, dan dynamic SQL yang terlihat pada source `db_object/procedures`. Informasi ini merupakan static evidence; bukan bukti branch runtime, jumlah row, atau business intent yang tidak tertulis.\n\n")
	b.WriteString(fmt.Sprintf("- Procedure profiles: **%d**\n- Procedures dengan dynamic SQL: **%d**\n- Source database access: **none** — analyzer hanya membaca `db_object`\n\n", len(docs.ProcedureProfiles), countProfilesWithDynamic(docs.ProcedureProfiles)))
	b.WriteString("## Reading order\n\n1. Cari procedure pada index.\n2. Baca `Procedure contract`, `At a glance`, dan `Execution signals`.\n3. Gunakan `Local execution flow` untuk pola source dan `Local object map` untuk inventory object yang direferensikan langsung.\n4. Gunakan `Direct object access` serta `Exact write mapping` untuk bukti statement root.\n5. Gunakan `Recursive procedure map` dan `Recursive data impact` untuk efek child procedure.\n6. Jangan menganggap item `UNRESOLVED_DYNAMIC`, `NOT_FOUND`, atau `AMBIGUOUS` sebagai object yang sudah terbukti.\n\n")
	b.WriteString("## Procedure index\n\n| Procedure | Direct data objects | Direct calls | Dynamic SQL | Profile |\n| --- | ---: | ---: | ---: | --- |\n")
	for _, p := range docs.ProcedureProfiles {
		b.WriteString(fmt.Sprintf("| `%s` | %d | %d | %d | [Open profile](profiles/%s.md) |\n", md(p.Name), len(uniqueAccessObjects(p.Accesses)), len(p.Calls), len(p.DynamicSQL), procedureLinkName(p.Name)))
	}
	b.WriteString(procedureResolutionReference())
	return b.String()
}

func renderProcedureCatalog(docs DesiredStateDocs) string {
	var b strings.Builder
	b.WriteString("# Procedure Impact Catalog\n\nCatalog ini adalah audit table lintas procedure. Gunakan profile per procedure untuk urutan statement, parameter, exact update mapping, dan recursive call path.\n\n")
	b.WriteString("| Procedure | Source | Direct reads | Direct writes | Direct calls | Functions | Views | Dynamic SQL | Unresolved references |\n| --- | --- | ---: | ---: | ---: | ---: | ---: | ---: | ---: |\n")
	for _, p := range docs.ProcedureProfiles {
		reads, writes, unresolved := 0, 0, 0
		for _, a := range p.Accesses {
			if a.Operation == "READ" {
				reads++
			} else {
				writes++
			}
			if !resolvedForImpact(a.Resolution) {
				unresolved++
			}
		}
		for _, c := range p.Calls {
			if !resolvedForImpact(c.Resolution) {
				unresolved++
			}
		}
		for _, dynamic := range p.DynamicSQL {
			if dynamic.Resolution == ResolutionUnresolved || dynamic.Resolution == ResolutionNotFound || dynamic.Resolution == ResolutionAmbiguous {
				unresolved++
			}
		}
		b.WriteString(fmt.Sprintf("| [`%s`](profiles/%s.md) | `%s` | %d | %d | %d | %d | %d | %d | %d |\n", md(p.Name), procedureLinkName(p.Name), md(p.SourcePath), reads, writes, len(p.Calls), len(p.Functions), len(p.Views), len(p.DynamicSQL), unresolved))
	}
	b.WriteString("\nCounts adalah jumlah static statement/reference yang ditemukan, bukan jumlah execution atau row database.\n")
	b.WriteString(procedureResolutionReference())
	return b.String()
}

func renderProcedureProfile(docs DesiredStateDocs, p *ProcedureProfile) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("# `%s`\n\n", md(p.Name)))
	b.WriteString(fmt.Sprintf("**Source:** `%s`  \n**Evidence boundary:** static source dari `db_object`; tidak ada koneksi atau eksekusi SQL.\n\n", md(p.SourcePath)))
	b.WriteString("**Quick navigation:** [Local execution flow](#local-execution-flow) · [Local object map](#local-object-map) · [Direct object access](#direct-object-access) · [Exact write mapping](#exact-write-mapping) · [Recursive procedure map](#recursive-procedure-map) · [Recursive data impact](#recursive-data-impact) · [Dynamic SQL](#dynamic-sql-analysis)\n\n")
	b.WriteString("## Procedure contract\n\n")
	if len(p.Parameters) == 0 {
		b.WriteString("Procedure tidak memiliki parameter yang terdeteksi.\n\n")
	} else {
		b.WriteString("| Parameter | SQL type | Direction | Default |\n| --- | --- | --- | --- |\n")
		for _, parameter := range p.Parameters {
			b.WriteString(fmt.Sprintf("| `%s` | `%s` | `%s` | %s |\n", md(parameter.Name), md(parameter.SQLType), parameter.Direction, codeOrDash(parameter.Default)))
		}
		b.WriteString("\n")
	}
	b.WriteString("| Output/error signal | Observed result |\n| --- | --- |\n")
	b.WriteString(fmt.Sprintf("| Literal return codes | %s |\n", codeList(p.ReturnCodes)))
	b.WriteString(fmt.Sprintf("| Direct result set | %s |\n", yesNo(p.HasDirectResult)))
	b.WriteString(fmt.Sprintf("| `RAISERROR` | %s |\n| `THROW` | %s |\n\n", yesNo(p.HasRaiseError), yesNo(p.HasThrow)))

	b.WriteString("## At a glance\n\n| Direct objects | Reads | Writes | Direct calls | Functions | Views | Dynamic SQL |\n| ---: | ---: | ---: | ---: | ---: | ---: | ---: |\n")
	reads, writes := 0, 0
	for _, a := range p.Accesses {
		if a.Operation == "READ" {
			reads++
		} else {
			writes++
		}
	}
	b.WriteString(fmt.Sprintf("| %d | %d | %d | %d | %d | %d | %d |\n\n", len(uniqueAccessObjects(p.Accesses)), reads, writes, len(p.Calls), len(p.Functions), len(p.Views), len(p.DynamicSQL)))
	b.WriteString("## Execution signals\n\n| Signal | Detected | Meaning for extraction |\n| --- | --- | --- |\n")
	b.WriteString(fmt.Sprintf("| Explicit transaction | %s | Pertahankan posisi commit/rollback; jangan memperluas atomicity tanpa review. |\n", yesNo(p.HasTransaction)))
	b.WriteString(fmt.Sprintf("| TRY/CATCH | %s | Error flow harus dibaca bersama `RAISERROR`, `THROW`, dan return code. |\n", yesNo(p.HasTryCatch)))
	b.WriteString(fmt.Sprintf("| Dynamic SQL | %s | Object dynamic hanya valid bila resolution terbukti. |\n\n", yesNo(len(p.DynamicSQL) > 0)))

	b.WriteString("## Local execution flow\n\nDiagram ini membaca pola procedure dari atas ke bawah. Scope transaction, pilihan conditional, titik bergabung kembali, child procedure call, dan error path hanya ditampilkan bila buktinya terlihat pada source. Cabang dari decision node adalah kemungkinan path, bukan urutan yang semuanya dijalankan. Guard diagram diringkas tanpa mengubah nilai pembeda; bentuk lengkap dan source line tetap tersedia pada `Observed process sequence` dan `Direct procedure calls`. Child procedure hanya menjadi call node dan tidak diekspansi di sini.\n\n")
	b.WriteString(renderProcedureMermaid(p))
	b.WriteString("### Diagram legend\n\n| Color/type | Meaning |\n| --- | --- |\n| Abu-abu | Read atau condition check terhadap table/view. |\n| Biru | Transaction scope atau perubahan data (`INSERT`, `UPDATE`, `DELETE`, `MERGE`). |\n| Ungu | Child procedure call; isi child dibaca pada recursive sections. |\n| Cokelat | Decision/conditional routing atau control statement. |\n| Merah | Error handler. Garis putus-putus menuju node ini berarti error transfer yang terlihat pada source. |\n\n")
	b.WriteString("## Local object map\n\nPeta ini menjawab object apa saja yang direferensikan langsung oleh procedure root. Panah menyatakan jenis akses (`READ`, `INSERT`, `UPDATE`, `DELETE`, `MERGE`, `EXEC`, atau function call), bukan arah foreign key dan bukan urutan eksekusi. Isi child procedure tidak dicampurkan; dampaknya tersedia pada `Recursive data impact`.\n\n")
	b.WriteString(renderLocalObjectMap(p))
	b.WriteString(renderObservedProcessSequence(p))

	b.WriteString("## Direct object access\n\n| Ref | Operation | Object | Guard | Resolution | Source line | Evidence |\n| --- | --- | --- | --- | --- | ---: | --- |\n")
	if len(p.Accesses) == 0 {
		b.WriteString("| — | — | — | — | — | — | Tidak ada persisted/static data access yang terdeteksi. |\n")
	}
	for _, a := range p.Accesses {
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | `%s` | %s | `%s` | %d | `%s` |\n", a.Ref, a.Operation, md(a.Object), codeOrDash(a.Guard), a.Resolution, a.Line, md(compactEvidence(a.Source))))
	}

	b.WriteString("\n## Exact write mapping\n\nTarget column dan source expression dipisahkan secara konservatif. `COLUMN_NOT_FOUND` atau `AMBIGUOUS` mempertahankan expression sebagai evidence untuk review; bukan mapping kolom terverifikasi. Untuk `UPDATE`, filter ditampilkan bila ada; `—` berarti tidak ada filter yang dipetakan, termasuk UPDATE tanpa WHERE. Statement yang tidak aman diparse tetap ada di Direct object access untuk manual review.\n\n| Ref | Target | Target column | Source expression | Filter | Resolution |\n| --- | --- | --- | --- | --- | --- |\n")
	if len(p.Mappings) == 0 {
		b.WriteString("| — | — | — | — | — | Tidak ada mapping INSERT/UPDATE yang dapat dibuktikan secara konservatif. |\n")
	}
	for _, m := range p.Mappings {
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | `%s` | `%s` | `%s` | `%s` |\n", m.Ref, md(m.Object), md(m.Column), md(m.Expression), md(emptyAsDash(m.Filter)), m.Resolution))
	}
	b.WriteString(renderProcedureKeyContext(docs, p))

	b.WriteString("\n## Direct procedure calls\n\nSetiap call ditampilkan sebagai blok agar parameter mapping dan guard SQL tetap terbaca tanpa tabel horizontal yang terlalu lebar.\n\n")
	if len(p.Calls) == 0 {
		b.WriteString("Tidak ada static procedure call yang terdeteksi.\n")
	}
	for index, call := range p.Calls {
		b.WriteString(fmt.Sprintf("### Call %d — `%s`\n\n", index+1, md(call.Object)))
		b.WriteString(fmt.Sprintf("- **Parameter mapping:** %s\n", codeOrDash(formatCallArgumentMapping(docs, call))))
		b.WriteString(fmt.Sprintf("- **Guard:** %s\n", codeOrDash(call.Guard)))
		b.WriteString(fmt.Sprintf("- **Resolution:** `%s`\n", call.Resolution))
		b.WriteString(fmt.Sprintf("- **Source:** line %d — `%s`\n\n", call.Line, md(compactEvidence(call.Source))))
	}

	b.WriteString("\n## Direct functions and views\n\n| Type | Object | Resolution |\n| --- | --- | --- |\n")
	if len(p.Functions)+len(p.Views) == 0 {
		b.WriteString("| — | — | Tidak ada function/view yang terdeteksi langsung. |\n")
	}
	for _, d := range p.Functions {
		b.WriteString(fmt.Sprintf("| FUNCTION | `%s` | `%s` |\n", md(d.Object), d.Resolution))
	}
	for _, d := range p.Views {
		b.WriteString(fmt.Sprintf("| VIEW | `%s` | `%s` |\n", md(d.Object), d.Resolution))
	}

	b.WriteString("\n## Recursive procedure map\n\nPanah berarti `caller → callee`, bukan urutan eksekusi antarsibling. Edge label menampilkan guard static yang terlihat pada caller; bila guard tidak terdeteksi, source line tetap ditampilkan. Diagram dibatasi untuk renderability, sedangkan tabel call path setelahnya tetap menjadi audit trail lengkap.\n\n")
	b.WriteString(renderRecursiveProcedureMermaid(docs, p.Name))

	b.WriteString("## Recursive call paths\n\nSetiap path mempertahankan arah `caller → callee`. `CYCLE` menghentikan recursion, sedangkan `ALREADY_MAPPED` berarti callee sudah diperluas melalui path sebelumnya sehingga isi yang sama tidak diduplikasi. Semua direct call edge tetap dicatat.\n\n| Call path | Status |\n| --- | --- |\n")
	paths := procedureCallPaths(docs, p.Name)
	if len(paths) == 0 {
		b.WriteString("| — | Tidak ada resolved child procedure. |\n")
	}
	for _, path := range paths {
		b.WriteString(fmt.Sprintf("| %s | `%s` |\n", codeCallPath(path.Names), path.Status))
	}

	b.WriteString("\n## Recursive data impact\n\nObject di bawah disentuh oleh child procedure, bukan oleh root secara langsung. Call path mencegah attribution yang keliru.\n\n| Operation | Object | Call path | Resolution | Source line |\n| --- | --- | --- | --- | ---: |\n")
	impacts := recursiveProcedureImpacts(docs, p.Name)
	if len(impacts) == 0 {
		b.WriteString("| — | — | — | — | — |\n")
	}
	for _, impact := range impacts {
		b.WriteString(fmt.Sprintf("| `%s` | `%s` | %s | `%s` | %d |\n", impact.Access.Operation, md(impact.Access.Object), codeCallPath(impact.Path), impact.Access.Resolution, impact.Access.Line))
	}

	b.WriteString("\n## Dynamic SQL analysis\n\n| Source line | Expression/evidence | Runtime source | Resolution | Resolved objects |\n| ---: | --- | --- | --- | --- |\n")
	if len(p.DynamicSQL) == 0 {
		b.WriteString("| — | Tidak ditemukan `sp_executesql`, `EXEC(<expression>)`, atau dynamic procedure variable. | — | `NONE_DETECTED` | — |\n")
	}
	for _, d := range p.DynamicSQL {
		b.WriteString(fmt.Sprintf("| %s | `%s` | %s | `%s` | %s |\n", d.Line, md(d.Expression), codeOrDash(d.RuntimeSource), d.Resolution, codeList(d.ResolvedObjects)))
	}

	b.WriteString("\n## Extraction review\n\n")
	if len(p.ExtractionNotices) == 0 {
		b.WriteString("Tidak ada automated warning khusus. Tetap validasi branch runtime, data cardinality, locking, retry, idempotency, dan caller di luar `db_object`.\n")
	}
	for _, notice := range p.ExtractionNotices {
		b.WriteString("- " + notice + "\n")
	}
	b.WriteString("\nSource-only analysis tidak membuktikan jumlah row, runtime branch, data quality, caller dari application/job, atau behavior object external yang snapshot-nya tidak tersedia.\n")
	b.WriteString(procedureResolutionReference())
	return b.String()
}

type procedureSequenceEvent struct {
	line, rank                     int
	event, target, evidence, guard string
	resolution                     string
	offset                         int
}

func procedureEventLess(a, b procedureSequenceEvent) bool {
	if a.line != b.line {
		return a.line < b.line
	}
	// Static access/call events carry exact offsets. Control/dynamic evidence
	// still uses the line-based classifier and retains its established rank.
	position := func(e procedureSequenceEvent) int {
		if e.offset >= 0 {
			return e.offset
		}
		if e.rank == 0 {
			return -1
		}
		return int(^uint(0) >> 1)
	}
	if position(a) != position(b) {
		return position(a) < position(b)
	}
	return a.rank < b.rank
}

func renderObservedProcessSequence(p *ProcedureProfile) string {
	var events []procedureSequenceEvent
	for _, access := range p.Accesses {
		events = append(events, procedureSequenceEvent{access.Line, 1, access.Operation, access.Object, compactEvidence(access.Source), access.Guard, access.Resolution, access.offset})
	}
	for _, call := range p.Calls {
		events = append(events, procedureSequenceEvent{call.Line, 2, "EXEC", call.Object, compactEvidence(call.Source), call.Guard, call.Resolution, call.offset})
	}
	for _, dynamic := range p.DynamicSQL {
		events = append(events, procedureSequenceEvent{dynamicLineNumber(dynamic.Line), 3, "DYNAMIC EXEC", emptyAsDash(strings.Join(dynamic.ResolvedObjects, "; ")), compactEvidence(dynamic.Expression), "", dynamic.Resolution, -1})
	}
	for _, control := range p.ControlSteps {
		events = append(events, procedureSequenceEvent{control.Line, 0, control.Event, control.Evidence, control.Evidence, "", "SOURCE_CONTROL", -1})
	}
	sort.SliceStable(events, func(i, j int) bool {
		return procedureEventLess(events[i], events[j])
	})
	var b strings.Builder
	b.WriteString("## Observed process sequence\n\nUrutan ini menggabungkan transaction/control signal, data access, dan procedure execution berdasarkan source line. Ini membantu membaca alur statis, tetapi tidak menyatakan bahwa setiap branch selalu dijalankan.\n\n| Step | Source line | Event | Target | Guard | Resolution | Evidence |\n| ---: | ---: | --- | --- | --- | --- | --- |\n")
	if len(events) == 0 {
		b.WriteString("| — | — | — | — | — | — | Tidak ada process event yang terdeteksi. |\n\n")
		return b.String()
	}
	for i, event := range events {
		b.WriteString(fmt.Sprintf("| %d | %d | `%s` | `%s` | %s | `%s` | `%s` |\n", i+1, event.line, event.event, md(event.target), codeOrDash(event.guard), event.resolution, md(event.evidence)))
	}
	b.WriteString("\n")
	return b.String()
}

func formatCallArgumentMapping(docs DesiredStateDocs, call ProcedureCall) string {
	if call.Resolution != ResolutionStatic || strings.TrimSpace(call.Arguments) == "" {
		return emptyAsDash(call.Arguments)
	}
	lookup := profileLookup(docs.ProcedureProfiles)
	callee := lookup[strings.ToLower(call.Object)]
	if callee == nil {
		return call.Arguments
	}
	args := sqlparse.SplitTopLevelComma(strings.TrimSpace(strings.TrimRight(call.Arguments, ";")))
	var mappings []string
	position := 0
	for _, arg := range args {
		arg = strings.TrimSpace(arg)
		if arg == "" {
			continue
		}
		if eq := topLevelEqual(arg); eq >= 0 && strings.HasPrefix(strings.TrimSpace(arg[:eq]), "@") {
			mappings = append(mappings, strings.TrimSpace(arg[:eq])+" ← "+strings.TrimSpace(arg[eq+1:]))
			continue
		}
		if position >= len(callee.Parameters) {
			mappings = append(mappings, "? ← "+arg)
			position++
			continue
		}
		mappings = append(mappings, callee.Parameters[position].Name+" ← "+arg)
		position++
	}
	if len(mappings) == 0 {
		return call.Arguments
	}
	return strings.Join(mappings, "; ")
}

func renderProcedureKeyContext(docs DesiredStateDocs, p *ProcedureProfile) string {
	var b strings.Builder
	b.WriteString("\n## Data key and update-cardinality context\n\nKey/index berasal dari `db_object/tables` dan `db_object/indexes`. `NOT_PROVEN` tidak berarti update pasti multi-row; artinya filter static tidak membuktikan satu row melalui enforced unique key.\n\n")
	b.WriteString("| Table | Enforced unique keys | Non-unique indexes | Update filter coverage |\n| --- | --- | --- | --- |\n")
	tableLookup := map[string]DocumentedTable{}
	for _, table := range docs.Tables {
		tableLookup[strings.ToLower(table.Name)] = table
	}
	objects := uniqueAccessObjects(p.Accesses)
	wrote := false
	for _, object := range objects {
		table, ok := tableLookup[strings.ToLower(object)]
		if !ok {
			continue
		}
		wrote = true
		var keys, indexes []string
		for _, key := range table.Keys {
			keys = append(keys, fmt.Sprintf("%s: (%s)", key.Evidence, strings.Join(key.Columns, ", ")))
		}
		for _, index := range table.Indexes {
			if !index.Unique {
				indexes = append(indexes, fmt.Sprintf("%s: (%s)", index.Name, strings.Join(index.Columns, ", ")))
			}
		}
		coverage := updateFilterCoverage(p.Mappings, table)
		b.WriteString(fmt.Sprintf("| `%s` | %s | %s | `%s` |\n", md(table.Name), codeList(keys), codeList(indexes), coverage))
	}
	if !wrote {
		b.WriteString("| — | — | — | Tidak ada resolved table access. |\n")
	}
	return b.String()
}

func updateFilterCoverage(mappings []ProcedureWriteMapping, table DocumentedTable) string {
	var filters []string
	for _, mapping := range mappings {
		if strings.EqualFold(mapping.Object, table.Name) && mapping.Filter != "" {
			filters = append(filters, mapping.Filter)
		}
	}
	if len(filters) == 0 {
		return "NO_PARSED_UPDATE_FILTER"
	}
	if len(table.Keys) == 0 {
		return "NO_ENFORCED_UNIQUE_KEY"
	}
	for _, filter := range filters {
		for _, key := range table.Keys {
			if key.Filtered || len(key.Columns) == 0 {
				continue
			}
			covered := true
			for _, column := range key.Columns {
				if !regexpContainsIdentifier(filter, column) {
					covered = false
					break
				}
			}
			if covered {
				return "UNIQUE_KEY_COLUMNS_PRESENT"
			}
		}
	}
	return "NOT_PROVEN"
}

func regexpContainsIdentifier(text, identifier string) bool {
	upperText, upperIdentifier := strings.ToUpper(text), strings.ToUpper(identifier)
	for offset := 0; ; {
		pos := strings.Index(upperText[offset:], upperIdentifier)
		if pos < 0 {
			return false
		}
		pos += offset
		before := pos == 0 || !isIdentifierByte(upperText[pos-1])
		afterPos := pos + len(upperIdentifier)
		after := afterPos >= len(upperText) || !isIdentifierByte(upperText[afterPos])
		if before && after {
			return true
		}
		offset = afterPos
	}
}

func renderProcedureMermaid(p *ProcedureProfile) string {
	events := prepareProcedureFlowEvents(procedureProcessEvents(p))
	if len(events) == 0 {
		return "Tidak ada local process event yang dapat digambar.\n\n"
	}
	chunks := splitProcedureFlowEvents(events, 20)
	var b strings.Builder
	for part, chunk := range chunks {
		if len(chunks) > 1 {
			b.WriteString(fmt.Sprintf("**Local execution map %d of %d**\n\n", part+1, len(chunks)))
		}
		startLabel := "START"
		if part > 0 {
			startLabel = fmt.Sprintf("CONTINUE · line %d", chunk[0].line)
		}
		b.WriteString(renderProcedureFlowChunk(p.Name, startLabel, chunk))
	}
	return b.String()
}

func prepareProcedureFlowEvents(events []procedureSequenceEvent) []procedureSequenceEvent {
	var prepared []procedureSequenceEvent
	for index, event := range events {
		upperEvidence := strings.ToUpper(strings.TrimSpace(event.evidence))
		nextGuardContainsTarget := index+1 < len(events) && events[index+1].line-event.line <= 20 && strings.Contains(strings.ToLower(events[index+1].guard), strings.ToLower(event.target))
		if event.event == "READ" && (strings.HasPrefix(upperEvidence, "IF EXISTS") || strings.HasPrefix(upperEvidence, "AND EXISTS") || (strings.HasPrefix(upperEvidence, "FROM ") && ((event.guard != "" && strings.Contains(strings.ToLower(event.guard), strings.ToLower(event.target))) || nextGuardContainsTarget))) {
			event.event = "CHECK"
			event.evidence = "condition read"
		}
		if len(prepared) > 0 && event.event == "CHECK" {
			last := &prepared[len(prepared)-1]
			if last.event == "CHECK" && strings.EqualFold(last.target, event.target) && event.line-last.line <= 30 {
				if last.guard != event.guard {
					last.guard = ""
				}
				last.evidence = "multiple conditions"
				continue
			}
		}
		prepared = append(prepared, event)
	}
	prepared = collapseProcedureTransactions(prepared)
	return collapseProcedureReadChecks(prepared)
}

func collapseProcedureTransactions(events []procedureSequenceEvent) []procedureSequenceEvent {
	var out []procedureSequenceEvent
	for index := 0; index < len(events); {
		if events[index].event != "TRANSACTION START" {
			out = append(out, events[index])
			index++
			continue
		}
		end := index + 1
		for end < len(events) && events[end].event != "COMMIT" {
			end++
		}
		if end >= len(events) {
			out = append(out, events[index])
			index++
			continue
		}
		var actions []string
		hasErrorTransfer := false
		hasProcedureCall := false
		for _, event := range events[index+1 : end] {
			switch event.event {
			case "INSERT", "UPDATE", "DELETE", "MERGE":
				action := event.event + " " + event.target
				if !containsFold(actions, action) {
					actions = append(actions, action)
				}
			case "GOTO":
				if strings.Contains(strings.ToUpper(event.target), "ERROR") {
					hasErrorTransfer = true
				}
			case "EXEC", "DYNAMIC EXEC":
				hasProcedureCall = true
			}
		}
		if hasProcedureCall {
			out = append(out, events[index:end+1]...)
			index = end + 1
			continue
		}
		if len(actions) > 3 {
			actions = append(actions[:3], fmt.Sprintf("+%d other data changes", len(actions)-3))
		}
		target := fmt.Sprintf("lines %d–%d", events[index].line, events[end].line)
		if len(actions) > 0 {
			target += "\\n" + strings.Join(actions, "\\n")
		}
		target += "\\nCOMMIT"
		evidence := "transaction summary"
		if hasErrorTransfer {
			evidence += " · GOTO ERROR"
		}
		out = append(out, procedureSequenceEvent{line: events[index].line, rank: events[index].rank, event: "TRANSACTION", target: target, evidence: evidence, resolution: "SOURCE_CONTROL"})
		index = end + 1
	}
	return out
}

func collapseProcedureReadChecks(events []procedureSequenceEvent) []procedureSequenceEvent {
	var out []procedureSequenceEvent
	for index := 0; index < len(events); {
		if events[index].guard != "" || (events[index].event != "READ" && events[index].event != "CHECK") {
			out = append(out, events[index])
			index++
			continue
		}
		end := index + 1
		for end < len(events) && events[end].guard == "" && (events[end].event == "READ" || events[end].event == "CHECK") && events[end].line-events[end-1].line <= 30 {
			end++
		}
		if end-index == 1 {
			out = append(out, events[index])
			index++
			continue
		}
		var targets []string
		for _, event := range events[index:end] {
			if !containsFold(targets, event.target) {
				targets = append(targets, event.target)
			}
		}
		if len(targets) > 3 {
			targets = append(targets[:3], fmt.Sprintf("+%d other objects", len(targets)-3))
		}
		target := fmt.Sprintf("lines %d–%d\\n%s", events[index].line, events[end-1].line, strings.Join(targets, "\\n"))
		out = append(out, procedureSequenceEvent{line: events[index].line, rank: events[index].rank, event: "READ / CHECK", target: target, evidence: "grouped read conditions", resolution: "SOURCE_CONTROL"})
		index = end
	}
	return out
}

func splitProcedureFlowEvents(events []procedureSequenceEvent, maxNodes int) [][]procedureSequenceEvent {
	var chunks [][]procedureSequenceEvent
	start := 0
	for end := 1; end <= len(events); end++ {
		if estimatedProcedureFlowNodes(events[start:end]) <= maxNodes {
			continue
		}
		cut := end - 1
		if cut <= start {
			cut = end
		}
		chunks = append(chunks, events[start:cut])
		start = cut
	}
	if start < len(events) {
		chunks = append(chunks, events[start:])
	}
	return chunks
}

func estimatedProcedureFlowNodes(events []procedureSequenceEvent) int {
	nodes := 1
	inGuardRun := false
	inErrorHandler := false
	hasErrorPath := false
	for _, event := range events {
		if event.event == "LABEL" && strings.Contains(strings.ToUpper(event.target), "ERROR") {
			inErrorHandler = true
			inGuardRun = false
			hasErrorPath = true
			continue
		}
		if inErrorHandler {
			continue
		}
		if event.event == "GOTO" && strings.Contains(strings.ToUpper(event.target), "ERROR") {
			inGuardRun = false
			hasErrorPath = true
			continue
		}
		nodes++
		if event.guard != "" && !inGuardRun {
			nodes++
			inGuardRun = true
		} else if event.guard == "" {
			inGuardRun = false
		}
	}
	if hasErrorPath {
		nodes++
	}
	return nodes
}

func renderProcedureFlowChunk(procedureName, startLabel string, events []procedureSequenceEvent) string {
	var b strings.Builder
	b.WriteString("```mermaid\nflowchart TD\n")
	b.WriteString(fmt.Sprintf("    start[\"`**%s**<br/>%s`\"]\n", procedureMermaidMarkdown(procedureName), startLabel))
	previous := "start"
	nodeIndex, branchIndex := 0, 0
	classes := map[string][]string{}
	var errorSources []string
	errorSummary := procedureErrorSummary(events)
	emitEvent := func(event procedureSequenceEvent) string {
		id := fmt.Sprintf("n%d", nodeIndex)
		nodeIndex++
		label := procedureMermaidMarkdown(processEventLabel(event))
		if event.event == "RETURN" {
			b.WriteString(fmt.Sprintf("    %s([\"`%s`\"])\n", id, label))
		} else {
			b.WriteString(fmt.Sprintf("    %s[\"`%s`\"]\n", id, label))
		}
		classes[processEventClass(event.event)] = append(classes[processEventClass(event.event)], id)
		return id
	}
	for index := 0; index < len(events); {
		event := events[index]
		if event.event == "LABEL" && strings.Contains(strings.ToUpper(event.target), "ERROR") {
			break
		}
		if event.event == "GOTO" && strings.Contains(strings.ToUpper(event.target), "ERROR") {
			errorSources = append(errorSources, fmt.Sprintf("%s|GOTO ERROR · line %d", previous, event.line))
			index++
			continue
		}
		if event.guard == "" {
			id := emitEvent(event)
			b.WriteString(fmt.Sprintf("    %s --> %s\n", previous, id))
			if strings.Contains(strings.ToUpper(event.evidence), "GOTO ERROR") {
				errorSources = append(errorSources, id+"|GOTO ERROR inside transaction")
			}
			previous = id
			index++
			continue
		}
		end := index
		for end < len(events) && events[end].guard != "" && !(events[end].event == "GOTO" && strings.Contains(strings.ToUpper(events[end].target), "ERROR")) {
			end++
		}
		branchID := fmt.Sprintf("b%d", branchIndex)
		branchIndex++
		branchLabel := procedureBranchLabel(events[index:end])
		b.WriteString(fmt.Sprintf("    %s{\"%s<br/>lines %d–%d\"}\n", branchID, branchLabel, events[index].line, events[end-1].line))
		classes["decision"] = append(classes["decision"], branchID)
		b.WriteString(fmt.Sprintf("    %s --> %s\n", previous, branchID))
		lastByGuard := map[string]string{}
		var branchEnds []string
		for _, guarded := range events[index:end] {
			id := emitEvent(guarded)
			guardKey := guarded.guard
			if last := lastByGuard[guardKey]; last != "" {
				b.WriteString(fmt.Sprintf("    %s --> %s\n", last, id))
			} else {
				b.WriteString(fmt.Sprintf("    %s -->|\"%s\"| %s\n", branchID, procedureMermaid(compactGuardForDiagram(guarded.guard)), id))
			}
			lastByGuard[guardKey] = id
		}
		for _, id := range lastByGuard {
			branchEnds = append(branchEnds, id)
		}
		sort.Strings(branchEnds)
		nextIndex := end
		for nextIndex < len(events) && events[nextIndex].event == "GOTO" && strings.Contains(strings.ToUpper(events[nextIndex].target), "ERROR") {
			errorSources = append(errorSources, fmt.Sprintf("%s|GOTO ERROR · line %d", branchID, events[nextIndex].line))
			nextIndex++
		}
		if nextIndex < len(events) && !(events[nextIndex].event == "LABEL" && strings.Contains(strings.ToUpper(events[nextIndex].target), "ERROR")) {
			next := events[nextIndex]
			nextID := emitEvent(next)
			for _, branchEnd := range branchEnds {
				b.WriteString(fmt.Sprintf("    %s --> %s\n", branchEnd, nextID))
			}
			b.WriteString(fmt.Sprintf("    %s -.->|\"condition not met\"| %s\n", branchID, nextID))
			previous = nextID
			index = nextIndex + 1
		} else {
			previous = branchID
			index = nextIndex
		}
	}
	if errorSummary != "" || len(errorSources) > 0 {
		if errorSummary == "" {
			errorSummary = "ERROR PATH"
		}
		b.WriteString(fmt.Sprintf("    error[\"`%s`\"]\n", procedureMermaidMarkdown(errorSummary)))
		classes["failure"] = append(classes["failure"], "error")
		for _, source := range errorSources {
			parts := strings.SplitN(source, "|", 2)
			b.WriteString(fmt.Sprintf("    %s -.->|\"%s\"| error\n", parts[0], mermaid(parts[1])))
		}
	}
	b.WriteString("    classDef read fill:#343a40,stroke:#adb5bd,color:#ffffff\n    classDef write fill:#17365d,stroke:#6ea8dc,color:#ffffff\n    classDef procedureNode fill:#3f2a56,stroke:#b594d6,color:#ffffff\n    classDef control fill:#5a431d,stroke:#dfb35b,color:#ffffff\n    classDef decision fill:#5a431d,stroke:#dfb35b,color:#ffffff\n    classDef failure fill:#5a2525,stroke:#df7777,color:#ffffff\n")
	for _, className := range []string{"read", "write", "procedureNode", "control", "decision", "failure"} {
		if len(classes[className]) > 0 {
			b.WriteString("    class " + strings.Join(classes[className], ",") + " " + className + "\n")
		}
	}
	b.WriteString("```\n\n")
	return b.String()
}

func procedureBranchLabel(events []procedureSequenceEvent) string {
	for _, event := range events {
		if event.event == "EXEC" || event.event == "DYNAMIC EXEC" {
			return "PROCEDURE ROUTING"
		}
	}
	for _, event := range events {
		if event.event == "INSERT" || event.event == "UPDATE" || event.event == "DELETE" || event.event == "MERGE" {
			return "CONDITIONAL DATA CHANGE"
		}
	}
	return "CONDITIONAL PATH"
}

func procedureErrorSummary(events []procedureSequenceEvent) string {
	inside := false
	var actions []string
	for _, event := range events {
		if event.event == "LABEL" && strings.Contains(strings.ToUpper(event.target), "ERROR") {
			inside = true
			continue
		}
		if !inside {
			continue
		}
		action := event.event
		if event.event == "UPDATE" || event.event == "INSERT" || event.event == "DELETE" || event.event == "MERGE" {
			action += " " + event.target
		}
		if !containsFold(actions, action) {
			actions = append(actions, action)
		}
	}
	if !inside {
		return ""
	}
	if len(actions) == 0 {
		return "ERROR HANDLER"
	}
	return "ERROR HANDLER\\n" + strings.Join(actions, " · ")
}

type localObjectMapItem struct {
	name, kind, resolution string
	operations             []string
}

func renderLocalObjectMap(p *ProcedureProfile) string {
	itemsByKey := map[string]*localObjectMapItem{}
	add := func(name, kind, resolution, operation string) {
		if strings.TrimSpace(name) == "" || name == "—" {
			return
		}
		key := strings.ToLower(kind + "|" + name)
		item := itemsByKey[key]
		if item == nil {
			item = &localObjectMapItem{name: name, kind: kind, resolution: resolution}
			itemsByKey[key] = item
		}
		if !containsFold(item.operations, operation) {
			item.operations = append(item.operations, operation)
		}
	}
	for _, access := range p.Accesses {
		add(access.Object, "DATA OBJECT", access.Resolution, access.Operation)
	}
	for _, call := range p.Calls {
		add(call.Object, "PROCEDURE", call.Resolution, "EXEC")
	}
	for _, dependency := range p.Functions {
		add(dependency.Object, "FUNCTION", dependency.Resolution, "CALL")
	}
	for _, dependency := range p.Views {
		add(dependency.Object, "VIEW", dependency.Resolution, "READ")
	}
	for _, dynamic := range p.DynamicSQL {
		for _, object := range dynamic.ResolvedObjects {
			add(object, "DYNAMIC TARGET", dynamic.Resolution, "DYNAMIC EXEC")
		}
	}
	items := make([]localObjectMapItem, 0, len(itemsByKey))
	for _, item := range itemsByKey {
		sort.Strings(item.operations)
		items = append(items, *item)
	}
	sort.Slice(items, func(i, j int) bool {
		if items[i].kind != items[j].kind {
			return items[i].kind < items[j].kind
		}
		return strings.ToLower(items[i].name) < strings.ToLower(items[j].name)
	})
	if len(items) == 0 {
		return "Tidak ada catalog-backed local object yang terdeteksi.\n\n"
	}
	chunkSize := procedureDiagramMaxNodes - 1
	parts := (len(items) + chunkSize - 1) / chunkSize
	var b strings.Builder
	for part, offset := 0, 0; offset < len(items); part, offset = part+1, offset+chunkSize {
		end := offset + chunkSize
		if end > len(items) {
			end = len(items)
		}
		if parts > 1 {
			b.WriteString(fmt.Sprintf("**Local object map %d of %d**\n\n", part+1, parts))
		}
		b.WriteString("```mermaid\nflowchart LR\n")
		b.WriteString(fmt.Sprintf("    root[\"`**%s**<br/>PROCEDURE`\"]\n", procedureMermaidMarkdown(p.Name)))
		for index, item := range items[offset:end] {
			id := fmt.Sprintf("o%d", index)
			label := item.name + "\\n" + item.kind + " · " + item.resolution
			b.WriteString(fmt.Sprintf("    %s[\"`%s`\"]\n", id, procedureMermaidMarkdown(label)))
			b.WriteString(fmt.Sprintf("    root -->|\"%s\"| %s\n", mermaid(strings.Join(item.operations, " + ")), id))
		}
		b.WriteString("```\n\n")
	}
	return b.String()
}

func procedureProcessEvents(p *ProcedureProfile) []procedureSequenceEvent {
	var events []procedureSequenceEvent
	for _, access := range p.Accesses {
		events = append(events, procedureSequenceEvent{access.Line, 1, access.Operation, access.Object, access.Source, access.Guard, access.Resolution, access.offset})
	}
	for _, call := range p.Calls {
		events = append(events, procedureSequenceEvent{call.Line, 2, "EXEC", call.Object, call.Source, call.Guard, call.Resolution, call.offset})
	}
	for _, control := range p.ControlSteps {
		events = append(events, procedureSequenceEvent{control.Line, 0, control.Event, control.Evidence, control.Evidence, "", "SOURCE_CONTROL", -1})
	}
	for _, dynamic := range p.DynamicSQL {
		target := strings.Join(dynamic.ResolvedObjects, ", ")
		if target == "" {
			target = dynamic.Expression
		}
		events = append(events, procedureSequenceEvent{dynamicLineNumber(dynamic.Line), 3, "DYNAMIC EXEC", target, dynamic.Expression, "", dynamic.Resolution, -1})
	}
	sort.SliceStable(events, func(i, j int) bool {
		return procedureEventLess(events[i], events[j])
	})
	var collapsed []procedureSequenceEvent
	for _, event := range events {
		if len(collapsed) > 0 {
			last := &collapsed[len(collapsed)-1]
			if last.event == event.event && strings.EqualFold(last.target, event.target) && last.guard == event.guard {
				last.line = event.line
				continue
			}
		}
		collapsed = append(collapsed, event)
	}
	return collapsed
}

func processEventLabel(event procedureSequenceEvent) string {
	label := event.event
	if event.target != "" && event.target != event.event {
		if event.event == "TRANSACTION" || event.event == "READ / CHECK" {
			label += "\\n" + event.target
		} else {
			label += "\\n" + compactDiagramText(event.target, 70)
		}
	}
	if event.event == "CHECK" && event.evidence == "multiple conditions" {
		label += "\\nmultiple conditions"
	}
	return label
}

func compactGuardForDiagram(value string) string {
	value = strings.ReplaceAll(compactSQL(value), "|", "/")
	parts := strings.Split(value, " AND ")
	if len(parts) == 1 {
		return compactDiagramText(value, 68)
	}
	const maxParts = 3
	selected := parts
	if len(selected) > maxParts {
		selected = []string{parts[0], parts[1], parts[len(parts)-1]}
	}
	for index := range selected {
		selected[index] = compactDiagramText(strings.TrimSpace(selected[index]), 48)
	}
	return strings.Join(selected, "\\n")
}

func procedureMermaid(value string) string {
	value = mermaid(value)
	return strings.ReplaceAll(value, "\\n", "<br/>")
}

func procedureMermaidMarkdown(value string) string {
	value = mermaidMarkdown(value)
	value = strings.ReplaceAll(value, "\n", "<br/>")
	return strings.ReplaceAll(value, "\\n", "<br/>")
}

func processEventClass(event string) string {
	switch event {
	case "READ", "CHECK", "READ / CHECK":
		return "read"
	case "INSERT", "UPDATE", "DELETE", "MERGE", "TRANSACTION":
		return "write"
	case "EXEC", "DYNAMIC EXEC":
		return "procedureNode"
	case "RAISE ERROR", "THROW", "ROLLBACK":
		return "failure"
	default:
		return "control"
	}
}

func compactDiagramText(value string, max int) string {
	value = compactSQL(value)
	value = strings.ReplaceAll(value, "|", "/")
	if len(value) <= max {
		return value
	}
	if elseAt := strings.LastIndex(value, "ELSE of "); elseAt > 0 {
		prefixBudget := max / 3
		suffixBudget := max - prefixBudget - 5
		prefix := value[:elseAt]
		suffix := value[elseAt:]
		if len(prefix) > prefixBudget {
			prefix = prefix[:prefixBudget-3] + "..."
		}
		if len(suffix) > suffixBudget {
			head := suffixBudget / 2
			tail := suffixBudget - head - 5
			suffix = suffix[:head] + " ... " + suffix[len(suffix)-tail:]
		}
		return prefix + " · " + suffix
	}
	tail := (max * 2) / 5
	head := max - tail - 5
	return value[:head] + " ... " + value[len(value)-tail:]
}

type procedurePath struct {
	Names  []string
	Status string
}
type recursiveImpact struct {
	Path   []string
	Access ProcedureDataAccess
}

type recursiveCallEdge struct {
	caller, callee, guard string
	line                  int
}

func renderRecursiveProcedureMermaid(docs DesiredStateDocs, root string) string {
	lookup := profileLookup(docs.ProcedureProfiles)
	expanded := map[string]bool{}
	var edges []recursiveCallEdge
	var walk func(string)
	walk = func(name string) {
		key := strings.ToLower(name)
		if expanded[key] {
			return
		}
		expanded[key] = true
		profile := lookup[key]
		if profile == nil {
			return
		}
		for _, call := range profile.Calls {
			if call.Resolution != ResolutionStatic {
				continue
			}
			edges = append(edges, recursiveCallEdge{name, call.Object, call.Guard, call.Line})
			walk(call.Object)
		}
	}
	walk(root)
	if len(edges) == 0 {
		return "Tidak ada resolved recursive procedure call yang dapat digambar.\n\n"
	}
	var b strings.Builder
	var batches [][]recursiveCallEdge
	var batch []recursiveCallEdge
	batchNodes := map[string]bool{}
	for _, edge := range edges {
		newNodes := 0
		for _, name := range []string{edge.caller, edge.callee} {
			if !batchNodes[strings.ToLower(name)] {
				newNodes++
			}
		}
		if len(batch) > 0 && len(batchNodes)+newNodes > procedureDiagramMaxNodes {
			batches = append(batches, batch)
			batch = nil
			batchNodes = map[string]bool{}
		}
		batch = append(batch, edge)
		batchNodes[strings.ToLower(edge.caller)] = true
		batchNodes[strings.ToLower(edge.callee)] = true
	}
	if len(batch) > 0 {
		batches = append(batches, batch)
	}
	for index, current := range batches {
		if len(batches) > 1 {
			b.WriteString(fmt.Sprintf("**Recursive call map %d of %d**\n\n", index+1, len(batches)))
		}
		b.WriteString(renderRecursiveCallMap(current, root))
	}
	return b.String()
}

func renderRecursiveCallMap(edges []recursiveCallEdge, root string) string {
	var nodeNames []string
	nodeSeen := map[string]bool{}
	for _, edge := range edges {
		for _, name := range []string{edge.caller, edge.callee} {
			key := strings.ToLower(name)
			if !nodeSeen[key] {
				nodeSeen[key] = true
				nodeNames = append(nodeNames, name)
			}
		}
	}
	ids := map[string]string{}
	var b strings.Builder
	b.WriteString("```mermaid\nflowchart TD\n")
	for i, name := range nodeNames {
		id := fmt.Sprintf("p%d", i)
		ids[strings.ToLower(name)] = id
		b.WriteString(fmt.Sprintf("    %s[\"`%s`\"]\n", id, procedureMermaidMarkdown(name)))
	}
	for _, edge := range edges {
		label := edge.guard
		if label == "" {
			label = "source line " + fmt.Sprint(edge.line)
		}
		b.WriteString(fmt.Sprintf("    %s -->|\"%s\"| %s\n", ids[strings.ToLower(edge.caller)], procedureMermaid(compactDiagramText(label, 100)), ids[strings.ToLower(edge.callee)]))
	}
	b.WriteString("    classDef rootNode fill:#17365d,stroke:#6ea8dc,color:#ffffff\n    classDef procedureNode fill:#3f2a56,stroke:#b594d6,color:#ffffff\n")
	var rootIDs, childIDs []string
	for _, name := range nodeNames {
		if strings.EqualFold(name, root) {
			rootIDs = append(rootIDs, ids[strings.ToLower(name)])
		} else {
			childIDs = append(childIDs, ids[strings.ToLower(name)])
		}
	}
	if len(rootIDs) > 0 {
		b.WriteString("    class " + strings.Join(rootIDs, ",") + " rootNode\n")
	}
	if len(childIDs) > 0 {
		b.WriteString("    class " + strings.Join(childIDs, ",") + " procedureNode\n")
	}
	b.WriteString("```\n\n")
	return b.String()
}

func procedureCallPaths(docs DesiredStateDocs, root string) []procedurePath {
	lookup := profileLookup(docs.ProcedureProfiles)
	var out []procedurePath
	expanded := map[string]bool{strings.ToLower(root): true}
	var walk func(string, []string, map[string]bool)
	walk = func(name string, path []string, visiting map[string]bool) {
		p := lookup[strings.ToLower(name)]
		if p == nil {
			return
		}
		for _, call := range p.Calls {
			if call.Resolution != ResolutionStatic {
				continue
			}
			next := append(append([]string{}, path...), call.Object)
			if visiting[strings.ToLower(call.Object)] {
				out = append(out, procedurePath{next, "CYCLE"})
				continue
			}
			if expanded[strings.ToLower(call.Object)] {
				out = append(out, procedurePath{next, "ALREADY_MAPPED"})
				continue
			}
			out = append(out, procedurePath{next, "RESOLVED"})
			expanded[strings.ToLower(call.Object)] = true
			copyVis := map[string]bool{}
			for k, v := range visiting {
				copyVis[k] = v
			}
			copyVis[strings.ToLower(call.Object)] = true
			walk(call.Object, next, copyVis)
		}
	}
	walk(root, []string{root}, map[string]bool{strings.ToLower(root): true})
	return uniqueProcedurePaths(out)
}

func recursiveProcedureImpacts(docs DesiredStateDocs, root string) []recursiveImpact {
	lookup := profileLookup(docs.ProcedureProfiles)
	paths := procedureCallPaths(docs, root)
	var out []recursiveImpact
	for _, path := range paths {
		if len(path.Names) < 2 {
			continue
		}
		child := lookup[strings.ToLower(path.Names[len(path.Names)-1])]
		if child == nil {
			continue
		}
		for _, a := range child.Accesses {
			if a.Operation == "READ" {
				continue
			}
			out = append(out, recursiveImpact{path.Names, a})
		}
	}
	sort.Slice(out, func(i, j int) bool {
		return strings.Join(out[i].Path, "|")+out[i].Access.Operation+out[i].Access.Object < strings.Join(out[j].Path, "|")+out[j].Access.Operation+out[j].Access.Object
	})
	return out
}

func profileLookup(profiles []ProcedureProfile) map[string]*ProcedureProfile {
	out := map[string]*ProcedureProfile{}
	for i := range profiles {
		out[strings.ToLower(profiles[i].Name)] = &profiles[i]
	}
	return out
}
func uniqueProcedurePaths(in []procedurePath) []procedurePath {
	seen := map[string]bool{}
	var out []procedurePath
	for _, p := range in {
		k := strings.Join(p.Names, "|") + p.Status
		if !seen[k] {
			seen[k] = true
			out = append(out, p)
		}
	}
	return out
}
func uniqueAccessObjects(in []ProcedureDataAccess) []string {
	var v []string
	for _, a := range in {
		if resolvedForImpact(a.Resolution) {
			v = append(v, a.Object)
		}
	}
	return uniqueStrings(v)
}
func resolvedForImpact(s string) bool {
	return s == ResolutionStatic || s == ResolutionDynamic || s == ResolutionFromCaller || s == ResolutionNotPersisted || s == ResolutionSystem
}
func countProfilesWithDynamic(in []ProcedureProfile) int {
	n := 0
	for _, p := range in {
		if len(p.DynamicSQL) > 0 {
			n++
		}
	}
	return n
}
func compactEvidence(s string) string {
	s = compactSQL(s)
	if len(s) > 180 {
		return s[:177] + "..."
	}
	return s
}
func procedureLinkName(name string) string { return url.PathEscape(name) }
func codeCallPath(path []string) string {
	var out []string
	for _, v := range path {
		out = append(out, "`"+md(v)+"`")
	}
	return strings.Join(out, " → ")
}

func procedureResolutionReference() string {
	return "\n## Resolution reference\n\n| Resolution | Meaning |\n| --- | --- |\n" +
		"| `RESOLVED_STATIC` | Nama object tertulis langsung dan cocok dengan object pada `db_object`. |\n" +
		"| `RESOLVED_DYNAMIC` | SQL literal dapat direkonstruksi, diparse ulang, dan object cocok dengan `db_object`. |\n" +
		"| `RESOLVED_FROM_CALLER` | Dynamic object parameter dapat dibuktikan dari literal pada static caller. |\n" +
		"| `EXTERNAL_DATABASE` | Three/four-part name menunjuk database/catalog lain; tidak dianggap object lokal. |\n" +
		"| `UNRESOLVED_DYNAMIC` | Target bergantung pada variable, parameter, metadata, configuration, atau runtime data. |\n" +
		"| `NOT_PERSISTED` | Temporary table atau table variable; bukan persisted database object. |\n" +
		"| `NOT_FOUND` | Nama terbaca tetapi tidak ditemukan pada snapshot `db_object` yang sesuai. |\n" +
		"| `COLUMN_NOT_FOUND` | Table ditemukan tetapi target column tidak ada pada snapshot; expression dipertahankan sebagai evidence untuk review. |\n" +
		"| `AMBIGUOUS` | Nama cocok ke lebih dari satu kind/object dan tidak aman dipilih otomatis. |\n" +
		"| `SYSTEM_OBJECT` | Built-in/system procedure; bukan application procedure dalam `db_object/procedures`. |\n\n" +
		"`db_object` digunakan untuk memvalidasi kandidat, bukan menebak nilai runtime.\n"
}
