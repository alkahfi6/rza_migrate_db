# Rules Traceability Matrix - rza_migrate_db

Rule 26 JOIN identifier refinement (2026-10-12): `identifier_join_review.go` → observed column mapping/module evidence tanpa index/identity prerequisite; `desired_state_writer.go` → FK-only exclusion dan self-role usage count. Controls: `identifier_join_review_test.go` untuk keyed-table/unindexed/generic/composite/filtered/self-join cases, `TestSuspectedJoinIdentifiersQA31GoldenReadOnlyAndIdempotent` untuk real source/golden, no source/history mutation, dan identical rerun. Report-only boundary: extractor/checker/generator tidak memperoleh schema action; release artifact parity tetap existing gate.


2026-10-01 refinement of rule 26: quoted-vs-keyword, malformed delimiters, TOP/MERGE/APPLY/comma/sink/static-call recognition map to `internal/sqlparse/access.go` and `dependencies.go`, covered by precision/quoted/reference/fuzz tests. Query-scoped aliases, per-occurrence evidence, and bounded mapping map to `internal/migreport/procedure_analysis.go` and `TestProcedureScannerQA31Golden`/negative regressions. `internal/engine/object_scan_contract_test.go` proves apply/read-only checker leaves authored SQL and migrations unchanged, generator retains only the new procedure up/down, and a second generation is idempotent. No extractor correction or new DDL category is introduced; released artifact/source parity stays enforced by the existing release contract.


Matriks ini memetakan setiap section `rules.md` ke dokumen turunan dan kontrol implementasi.

| Rule No | Section `rules.md` | `spec.md` | `fsd.md` | `plan.md` | Kontrol/Gate |
| --- | --- | --- | --- | --- | --- |
| 1 | Istilah & Konsep | 2 | 1, 4, 9 | 8 | Definisi konsisten di semua dokumen |
| 2 | Foldering Repo | 4 | 5, 7 | 6 (RC-02) | Verify folder contract |
| 3 | Tooling & Pipeline | 5 | 4 | 5, 6 (RC-19) | Command flow contract |
| 4 | Larangan & Safety Guard | 6 | 7 | 6 (RC-03) | Forbidden token checks |
| 5 | Pairing & Filename | 7 | 6.1, 7.3 | 6 (RC-04) | Pair-safe validation |
| 6 | Ordering, `v{k}`, Reconstruction | 8 | 8, 9, 10 | 6 (RC-05, RC-06) | Deterministic emit gate |
| 7 | Data Master Validation | 9 | 6.4, 7.1 | 6 (RC-07) | Source schema contract checks |
| 8 | Destructive & Decision Semantics | 10, 11 | 9.3 | 6 (RC-08) | Policy mapping checks |
| 9 | Transaction & Error Strategy | 12 | 10.3, 15 | 6 (RC-09) | Wrapper + THROW checks |
| 10 | Core Audit Init | 13.1 | 10, 14 | 5 (M10), 6 (RC-10) | Audit init pair checks |
| 11 | Audit Logging | 13.2, 13.3 | 10.3, 12 | 6 (RC-10) | Logger contract checks |
| 12 | Wrapper Standar | 14 | 10.3, 11 | 6 (RC-11) | Wrapper minimum validation |
| 13 | Rules Down Migration | 15 | 10.4 | 6 (RC-12) | Down behavior checks |
| 14 | Patch Policy | 16 | 7.2, 10.5 | 6 (RC-13) | Patch lifecycle + immutability hash checks + cross-platform case-insensitive actual-entry resolution |
| 15 | Output `out/output.json` | 17 | 3.3, 12 | 6 (RC-14) | Output schema checks + complete console failure diagnostics |
| 16 | Locking 1 Run per DB | 18 | 14 | 6 (RC-15) | Lock readiness check |
| 17 | MVP Supported Subset | 19 | 6.3, 16 | 6 (RC-16) | Unsupported feature fail-fast |
| 18 | ALTER COLUMN Precheck | 20 | 9.3 | 6 (RC-17) | Precheck templates + parameterized/bracketed-type parser/hash/diff regression (`TestColumnTypeParametersWithWhitespaceRemainSemantic`, `TestRiskClassifiesBracketedDecimalScaleChangeAsGuarded`, `TestBracketedAndBareBuiltInTypesAreSemanticallyEquivalent`, `TestGenDetectsParameterizedColumnChangesWithWhitespace`) |
| 19 | constraints_naming.md | 21 | 7.1, 7.4 | 6 (RC-18) | Optional guideline + deterministic unnamed constraint safe-fix |
| 20 | Repo Hygiene | 21 | 7.4 | 6 (RC-18) | Optional warning + bootstrap seed tanpa overwrite |
| 21 | CLI Contract | 22 | 4 | 6 (RC-19) | Command and exit checks |
| 22 | Engine Config & Verify Rules | 23, 24 | 2, 2.1, 3, 11 | 6 (RC-20) | Config + verify gates + modular package contract |
| 23 | Final Locks / Gaps Closed | 3, 25, 26 | 15, 17 | 2, 6, 8 | Hard lock governance |
| 24 | Fail/Pass Output Patterns | 27 | 12, 16.2 | 6 (RC-14), QA-18 | Output pattern checks |
| 25 | Runner Cheatsheet | 27 | 4.3, 15 | 6 (RC-19) | Runner behavior assumptions |
| 26 | Migration History Reconstruction & Report (`rza_migrate_report`) | 26 | 2.1, 4 | Migration report section | QA-30/QA-31; shared strict filename parser, source-positioned token access scanner (`internal/sqlparse`) with multiline identity, statement-scoped CTE/alias resolution and semantic client-result classification, excluded canonical core-audit bootstrap with explicit analyzed/excluded counters, ordered dense multi-item table reducer, wrapper-safe exactly-once `CREATE TABLE` scanner, uniqueness enforcement-vs-signal inventory, declared-FK/observed-equality graph with exact columns, Indonesian glossary, single-file connected Relational Maps bounded at 20 objects/24 lines/35 mappings/9.000 Mermaid characters without artificial Hub, node key/index context, object index/shared navigation, exact mapping edge labels/full pair+mapping coverage, separate self-join review/no raw HTML outside Mermaid code fences, separate module lineage/dynamic review with case-insensitive dependency dedup and static return-assignment recognition, catalog-backed procedure profiles with URL-safe links and source-preserved guards, top-down branch-aware local execution flow (transaction scope, conditional routing/rejoin, child calls, error transfer, Mermaid-rendered multi-line guards, diagram legend), separate local object inventory, full source sequence, exact write mapping, complete split recursive maps, cycle-safe recursive impact, dynamic resolution codes, repo-relative output resolution, concrete JSON artifact manifest, no same-column/operator/alias-scope/runtime-target inference, output-path immutability guard, unit/CLI/real-repo tests di `internal/sqlparse`, `internal/migreport`, dan `cmd/rza_migrate_report` |

## Cara Pakai

1. Saat implementasi, tandai rule yang terdampak pada MR template.
2. Saat QA, jalankan checklist `docs/qa/qa_gate_engine_rules.md`.
3. Saat review dokumen, sinkronkan matrix ini jika ada perubahan section.

## Catatan Gap Aktual (2026-02-19)

Hasil review implementasi pada refinement M1-M5:

1. `gen` precheck kini menjalankan verify strict yang sama dengan `verify` (termasuk wrapper validation) sehingga legacy wrapper issue tidak bisa bypass no-diff run.
2. Wrapper patch di migration history sekarang wajib memenuhi contract wrapper (`WRAPPER_CONTRACT_MISSING` bila hilang metadata/transaction/audit contract).
3. Source contract `db_object/indexes` diperketat: file index hanya boleh berisi statement `CREATE [UNIQUE] INDEX` (statement lain => `INDEX_SOURCE_CONTRACT_VIOLATION`).
4. Source contract programmability diperketat: file harus tepat 1 object `CREATE OR ALTER` dan bebas DDL non-programmability (`PROGRAMMABILITY_SOURCE_CONTRACT_VIOLATION`).
5. Jalur `ALTER COLUMN` auto-plan menyisipkan precheck deterministic (`NULL->NOT NULL`, shrink string length, string->numeric conversion) dan log skip severity warning saat risiko data terdeteksi. Parser deklarasi tipe juga mempertahankan precision/scale/length/date-time precision di seluruh variasi whitespace serta menormalisasi delimiter tipe bawaan seperti `[DECIMAL]`, agar semantic hash/risk/diff tidak menghapus delta schema nyata atau menghasilkan false manual classification.
6. Orkestrasi urutan `validate*` untuk `verify`/`gen precheck` sudah dipindahkan ke `internal/verify` agar engine menjadi router tipis dan lifecycle report tetap konsisten.
7. Bridge typed callback dipisah ke `internal/engine/verify_callbacks.go` untuk menjaga type safety lintas package tanpa mengubah behavior validator existing.
8. Patch source resolver memakai nama directory entry aktual dengan logical-name comparison case-insensitive; regression test menjaga parity Windows/Linux dan mencegah false `PATCH_SOURCE_MISSING`.
9. `internal/output.PrintFailureDetails` dipakai checker dan generator agar semua fail item terlihat lengkap pada terminal dan GitLab job log.
10. `rza_migrate_report` adalah executable read-only independen untuk merekonstruksi histori dari `migrations/`, serta menghasilkan Markdown current-state dari `db_object/`, tanpa koneksi database atau mutasi source.

Residual risk aktif saat ini:
- backend formatter ANTLR masih fallback ke lite selama backend ANTLR belum diaktifkan penuh.
