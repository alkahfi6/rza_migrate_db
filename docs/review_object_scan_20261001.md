# Review ulang object scan — 1 Oktober 2026

Perubahan sebelumnya belum sempurna. Review ini menemukan bug reproducible pada scanner dan report lalu menambahkan perbaikan di shared `internal/sqlparse`, consumer `internal/migreport`, serta regression lintas engine. Pola yang diserap dari `kahfigo/internal/sqlparser` adalah pembedaan jenis token, source position, scoped CTE/alias, serta pemisahan SELECT/OUTPUT population dari client result. Parser kahfigo tidak disalin menjadi migration planner.

## Temuan yang diperbaiki

| Temuan | Perbaikan dan bukti |
| --- | --- |
| `SELECT 1 FROM [` memicu panic slice bounds | Delimiter tidak lengkap tidak dijadikan identifier. Precision regression dan fuzz source bounds. |
| Quoted keyword/identifier, escaped bracket, spasi dan titik di dalam nama kehilangan identity | Keyword hanya unquoted; string opaque; `NameParts` mempertahankan komponen identity. Quoted-name/reference/catalog tests. |
| TOP menjadi target DML; MERGE USING tidak terbaca; action MERGE dianggap DML sendiri | Skip TOP expression, scan USING, dan bedakan WHEN/THEN action dari statement. Precision golden. |
| APPLY/comma sources dan SELECT/OUTPUT sink tidak tercatat | Source/sink accesses memakai scanner yang sama; OUTPUT tanpa sink menjadi client result. |
| Alias UPDATE/DELETE bercampur dengan nested query atau table bernama sama | Resolve alias sintaktis dahulu pada statement/query depth yang sama; kandidat ambigu tidak dibuang. Report dan dependency memakai target yang sama. |
| Occurrence pada baris yang sama terhapus atau urutan EXEC/UPDATE terbalik | Tidak dedup per line; audit/static flow menyortir source offsets. Same-line sequence test. |
| EXEC multiline kehilangan target/arguments; END ikut menjadi argument | Static call scanner shared dan source bounds yang membedakan block END dari CASE END. Call/block tests. |
| INSERT target/column list, function definition, EXISTS, dan CTE column list menjadi routine palsu; external prefix menjadi local edge | Dependency memakai lexer yang sama, membedakan context, dan menolak multipart external fallback. Reference negative tests. |
| Mapping melintasi statement, window panjang terpotong, atau FROM ikut expression | Mapping memakai source span dan top-level lexical clause. Multirow/TOP/DISTINCT/compound exact mappings belum disajikan. |
| Case-variant CTE dipilih sebagai table meskipun collation offline tidak diketahui | Report AMBIGUOUS; shared graph tetap mempertahankan kemungkinan physical dependency secara exact-case. |
| UNION ALL dari INSERT/SELECT INTO atau cursor SELECT menjadi client result | Compound/cursor owner classifier; negative result tests. |

## Batas komponen

| Jalur | Dampak / bukti |
| --- | --- |
| `rza_get_object_db` | Metadata extraction `internal/getobject` tetap sumber identity/schema. Tidak ada kategori baru, SQL rewrite, atau migration generation dari scanner ini. Extractor tests dijalankan lewat full suite. |
| Interactive `rza_migrate_check` | Safe corrections existing tetap; scanner tidak menambah fixer. Boundary test membuktikan SQL/history canonical tidak berubah dan tidak ada migration baru. |
| `rza_migrate_check -check` / Linux MR | Read-only plan/validation tetap; boundary test memeriksa hash source/history sebelum/sesudah, release smoke memakai command GitLab. |
| Formatter/verifier/generator | Tidak merewrite business/schema SQL karena access evidence. Shared reference discovery dipakai dependency ordering dan type/module dependency path; create/removal ordering dapat lebih tepat. Integration membuktikan hanya procedure baru yang menghasilkan pair, up mempertahankan body/constant, down DROP procedure, run kedua no-op. |
| Report | Offline catalog-backed inventory, evidence dan static-call ordering lebih presisi; source/history tidak dimutasi. Relative roots/concrete manifest dari perubahan sebelumnya tetap. |
| Release | `build_all.bat` membangun dan menerbitkan seluruh Windows/Linux artifacts dari source final; checksum, embedded identity, source digest, dan Linux checker wajib lolos. Validation evidence disimpan di `out/`, bukan migration history. |

## Bukti dan residual scope

Targeted `internal/sqlparse`, `internal/migreport`, dan `internal/engine` tests telah lolos. QA-31 sekarang punya `dbo.uspPrecision.sql` serta `procedure_scan` golden yang benar-benar dieksekusi, bukan daftar evidence pasif. Fuzz awal memproses 237.764 input selama 16 detik tanpa panic. Final validation mencakup fuzz setelah perubahan terakhir, full Go suite, hardening/repeatability, full build, dan Linux MR smoke; hasil aktual dilaporkan pada penutupan pekerjaan.

Ini belum full T-SQL AST atau jaminan seluruh grammar. Runtime SQL/constant propagation antar-assignment/branch, synonym metadata, function-body lineage, dan projection/mapping kompleks tetap pekerjaan lanjutan. Control/dynamic evidence masih memakai classifier per line; exact source offsets baru berlaku untuk static access/call. Diagram flow dapat merangkum event berulang, sedangkan audit table mempertahankan occurrence. Penempatan visual Mermaid bergantung renderer dan tidak dibuktikan hanya oleh deterministic source-order tests. Tidak ada live SQL Server integration baru yang diklaim oleh offline tests.

## Dokumentasi

Updated: README, rules, spec, FSD, plan, walkthrough, SQL Server capability matrix, rules traceability matrix, QA gate, learning edge-case backlog, QA-31 README/input/expected golden, previous kahfigo review pointer, review ini, dan generated QA coverage via script.

Reviewed intentionally unchanged: engine docs guide (reading/release workflow sama), production readiness dan `.gitlab/rza_migrate_check.gitlab-ci.yml` (release/MR contract sama), parallel-safe design (tidak ada concurrency/cache baru), desired-state lifecycle dan guarded migrations (history/diff/reversibility policy sama), bootstrap four repos (extractor/catalog workflow sama), constraints naming (tidak ada correction/naming baru), migration monitoring/REST/rebaseline/legacy wrapper policies (execution/service/history-repair/wrapper contract tidak berubah). MR templates tetap memakai gate release yang sama.
