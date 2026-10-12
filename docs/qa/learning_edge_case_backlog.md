# Learning Edge Case Backlog

Dokumen ini mencatat gap skenario `examples/rules_2x_learning` yang sudah ditambah dan backlog lanjutan.

## Ditambahkan pada iterasi ini

1. `32_pass_index_include_identity`
- Fokus: `2.1`, `7.6`, `17.1`, `17.2`.
- Tujuan:
  - validasi PASS untuk table `IDENTITY(1,1)`,
  - validasi PASS untuk multi index per file,
  - validasi PASS untuk `CREATE [UNIQUE] INDEX ... INCLUDE (...)`.

2. `33_pass_index_clustered_keyword`
- Fokus: `2.1`, `7.6`, `17.2`.
- Tujuan: contoh PASS untuk `CLUSTERED/NONCLUSTERED` source index dan multi-index dalam satu file.

3. `34_pass_index_nonclustered_keyword`
- Fokus: `2.1`, `7.6`, `17.2`.
- Tujuan: contoh PASS untuk `NONCLUSTERED`, `UNIQUE NONCLUSTERED`, dan `INCLUDE`.

4. `35_pass_index_filtered_where`
- Fokus: `2.1`, `7.6`, `17.2`.
- Tujuan: contoh PASS untuk filtered index (`WHERE`) dalam source index.

5. `36_fail_forbidden_sql_token`
- Fokus: `4`.
- Tujuan: contoh FAIL untuk forbidden SQL token (`USE master`).

6. `37_fail_wrapper_throw_required`
- Fokus: `9`, `12`.
- Tujuan: contoh FAIL saat wrapper CATCH tidak melakukan `THROW`.

7. `38_fail_core_audit_contract_drift`
- Fokus: `10`, `11`, `16`.
- Tujuan: contoh FAIL untuk drift core audit init dan lock contract.

8. `39_warn_missing_constraints_naming`
- Fokus: `19`.
- Tujuan: contoh warning non-blocking saat `constraints_naming.md` hilang.

9. `40_fail_repo_hygiene_content_missing`
- Fokus: `20`.
- Tujuan: contoh FAIL saat `.gitattributes` tidak memenuhi token wajib.

10. `41_fail_go_batch_separator`
- Fokus: `4`.
- Tujuan: contoh FAIL eksplisit untuk batch separator `GO`.

11. `42_pass_warning_traceability_bundle`
- Fokus: `8`, `13`, `15`, `24`.
- Tujuan: contoh PASS-with-warning yang meng-assert `pass.reason` dan pola SQL generated secara bersamaan.

## Snapshot coverage rules.md -> learning examples (top-level)

Semua top-level section `1..25` sekarang sudah memiliki coverage eksplisit di `learning_matrix.json`.

- Coverage lifecycle/command: `3`, `21`, `25`.
- Coverage safety/quality gates: `4`, `9`, `10`, `11`, `16`, `19`, `20`.
- Coverage doc/config/reference sections: `1`, `22`, `23`, `24`.
- Coverage contract utama migration/source/patch/wrapper: `2`, `5`, `6`, `7`, `8`, `12`, `13`, `14`, `15`, `17`, `18`.

## Rekomendasi backlog berikutnya

1. `pass_alter_column_safe_matrix`
- Tambah skenario alter-column aman (widen type/nullable relaxation) untuk melengkapi precheck map.

## Resolved report-analysis edge cases (2026-09-10)

2026-10-12: JOIN identifier review tidak lagi bergantung pada index/IDENTITY/name-ID atau ketiadaan PK lain. Dedicated listing menyimpan mapping/source provenance dan metadata evidence; generic Code JOIN operand tetap signal, bukan unique row ID. FK-only tidak meningkatkan observed usage; self-role dihitung sekali. Filtered/composite boundaries dan no-write/idempotence diuji.

- Desired-state relationship scanner tidak lagi membaca bagian `=` dari `>=`, `<=`, `!=`, atau `==` sebagai equality join.
- Alias di-reset pada query scope `FROM`; alias lama dari statement tanpa semicolon tidak boleh membentuk edge.
- Equality expression wajib mempunyai tepat satu resolved `alias.column` per sisi; wrapper satu-kolom didukung, ekspresi multi-column di-skip.
- Legacy comma join `FROM A, B WHERE A.Id = B.Id` tercakup; temp/table-variable endpoint dan dynamic SQL tetap manual review.
- Inline PK/UQ tercakup dan token `UNIQUE` pada tipe `UNIQUEIDENTIFIER` tidak boleh menjadi false unique constraint.
- Identifier-like JOIN/non-unique index diklasifikasikan sebagai review signal, bukan unique-key candidate.
- Repeated risk/action text tidak lagi ditampilkan per table row karena merupakan guidance berbasis classification, bukan data-specific risk score. ERD besar tidak dipotong alfabetis atau disebar ke banyak file: seluruh diagram berada pada satu document dan memakai connected Relational Map adaptif dengan hard limit 20 object, 24 relationship lines, 35 mappings, dan 9.000 karakter Mermaid tanpa artificial Hub ownership. Node memuat type serta key/index context; object index dan shared-object table menjaga navigasi. Coverage test memastikan semua non-self pairs dan exact mappings tetap lengkap; self-join dipisahkan.
- Procedure profile resolution kini catalog-backed ke `db_object` dan membedakan static, reconstructed literal dynamic, caller literal, runtime unresolved, external, temporary, missing, ambiguous, serta system objects. Case/schema collision tidak dipilih otomatis. TVF, built-in function, cursor, alias lintas statement, return-value EXEC assignment, dan internal SELECT forms mempunyai regression coverage agar tidak menjadi false object/result-set signal. Local execution flow top-down merangkum source menjadi transaction scope, conditional routing/rejoin, grouped checks, dan error transfer tanpa mengarang business phase; local object map memisahkan inventory direct dependency dari alur. Transaction yang memiliki child call tidak boleh diringkas. Guard multiline mempertahankan literal source pada audit detail, sementara diagram memakai line break yang dirender Mermaid agar `\n` tidak terlihat. Recursive traversal menghentikan cycle/subgraph berulang dan memecah diagram maksimal 20 procedure tanpa membuang direct call edge.

## Resolved report object-recognition edge cases (2026-09-30)

Refinement 2026-10-01: malformed identifier panic, quoted-keyword/space corruption, TOP target omission, phantom MERGE actions/missing USING, missing APPLY/comma/sink accesses, cross-subquery alias resolution, same-line deduplication, multiline EXEC, false INSERT/function-definition/CTE-column-list routine edges, and statement-crossing write mapping have focused regression coverage. OUTPUT/cursor/compound result contracts also covered. CTE case variants remain explicitly ambiguous in offline report.

Residual scope: scanner is not full T-SQL AST; statement/control guard reconstruction and complex projection/expression semantics remain conservative. Runtime SQL assignment/branch constant propagation, synonym resolution, and function-body lineage are not newly supported. INSERT multirow/TOP/DISTINCT/compound exact mappings are intentionally omitted rather than inferred. Mermaid node placement depends on renderer; source ordering tests do not prove visual positions identical.

- Access scanner tidak lagi bergantung pada physical line: keyword dan multipart object yang terpisah newline tetap menghasilkan source-positioned access.
- CTE identity dibatasi pada statement dan dibandingkan exact-case. CTE unqualified tidak menutupi schema-qualified persisted object pada statement berikutnya; case-distinct name tetap tersedia untuk catalog ambiguity policy.
- UPDATE/DELETE alias di-resolve hanya dari source alias pada statement yang sama. Temporary table dan table-variable identity dipertahankan sebagai `NOT_PERSISTED`.
- Client result detection membedakan direct SELECT dari assignment, `SELECT INTO`, subquery/`EXISTS`, dan `INSERT ... SELECT`.
- Relative `-out`/`-docs` sekarang repository-root based. `report.json` menyimpan resolved roots dan concrete artifact manifest, termasuk generated procedure profiles, sementara wildcard hanya bertahan pada compatibility field.
