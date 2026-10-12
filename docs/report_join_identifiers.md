# Suspected JOIN identifiers

`rza_migrate_report` menghasilkan listing tambahan pada `table_unique_ids.md` untuk kolom yang terlihat sebagai join/reference identifier, walau tanpa index atau IDENTITY. PK pada kolom lain tidak menyembunyikan kandidat. Listing memakai observed static equality JOIN dari source view/procedure; tidak membaca isi row database atau menebak dari nama kolom saja.

Contoh dari QA-31:

| Table | Column | Classification | Uniqueness | Evidence |
| --- | --- | --- | --- | --- |
| dbo.JoinReference | ParentId | OBSERVED_JOIN_ID_SIGNAL | NOT_ENFORCED | Mapping dbo.JoinReference.ParentId = dbo.Parent.Id; source dbo.vJoinCandidates; index/key tidak ditemukan; IDENTITY tidak. |
| dbo.JoinReference | BusinessCode | OBSERVED_JOIN_COLUMN_SIGNAL | NOT_ENFORCED | Mapping dbo.JoinReference.BusinessCode = dbo.Parent.Code; source dbo.vJoinCandidates; index/key tidak ditemukan; IDENTITY tidak. |

Evidence menyatakan canonical column mapping, bukan salinan expression/alias SQL asli. Source modules serta occurrence count menunjukkan dasar pemakaian; bukan uniqueness, row count, cardinality, atau jaminan FK. ID-like name membedakan classification, tetapi generic column seperti Code tetap dapat menjadi signal bila source JOIN membuktikan pemakaiannya.

Single-column unfiltered unique key tetap di inventory enforcement utama dan tidak diulang dalam listing suspected. Kolom composite key tetap menjadi candidate dengan keterangan uniqueness tuple, bukan kolom sendiri. Filtered single-column unique memakai CONDITIONAL. IDENTITY tanpa single-column unique tetap tidak membuktikan uniqueness. Index/key dan IDENTITY metadata dicatat walaupun tidak ada.

Declared FK tanpa observed JOIN, inequality, temporary object, ambiguous mapping, name-only similarity, serta runtime SQL yang belum ter-resolve bukan observed evidence. FK tidak meningkatkan Join usage; self-join pada kolom yang sama di kedua role dihitung satu kali per occurrence. Output diurutkan/didedup deterministik.

Boundary: hanya renderer/report inventory yang berubah. Extractor tetap metadata/catalog based; checker interactive dan `-check` tidak menambah fixer atau migration; fmt/verify/generator tidak memakai candidate untuk schema change; Linux MR gate tetap read-only. Tidak ada sejarah migration yang ditulis ulang. Full Windows/Linux release build membawa perubahan report yang sama.

Proof: `identifier_join_review_test.go` mencakup unindexed/generic columns on keyed table, FK-only, filtered/composite key, self-join, dan summary consistency. `identifier_join_integration_test.go` mengeksekusi QA-31 golden, mempertahankan authored SQL/history hash, dan membuktikan identical second report. Final gates tetap full Go tests/hardening, complete build/checksum/source identity, GitLab Linux checker, serta Windows/Linux report parity.

Documentation updated: README, rules, spec, FSD, plan, walkthrough, capability matrix, traceability, QA gate, edge-case backlog, generated QA coverage, QA-31 README/input/golden, dan dokumen ini. Reviewed unchanged: engine docs guide, production readiness, GitLab checker/MR template, parallel-safe design, desired-state lifecycle, guarded migrations, bootstrap four repos, constraint naming, monitoring/REST/rebaseline/legacy wrapper policies. Semua boundary/check/release/data-safety contract tersebut tetap karena fitur ini report-only.
