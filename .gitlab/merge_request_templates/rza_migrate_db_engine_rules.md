# MR Template - rza_migrate_db Engine Rules Gate

Gunakan template ini untuk semua MR yang mengubah engine `rza_migrate_db`, rule validation, parser, diff, emit, atau policy migration.

## Ringkasan Perubahan

- Scope perubahan:
- Modul yang terdampak:
- Rule utama yang terdampak:
- Risiko utama:

## Hard Lock Acknowledgement

- [ ] Saya memahami fail rule: jika severity=3 atau `out/output.json.fail.count > 0`, proses wajib exit 1 dan tidak ada migration baru yang boleh tersisa.
- [ ] Saya memastikan perubahan ini tidak memperkenalkan bypass terhadap fail-hard.
- [ ] Saya memastikan tidak ada parser utama berbasis regex untuk grammar filename/SQL subset yang dikunci.

## Rule Checklist (RC-01 s.d. RC-20)

- [ ] RC-01 Hard fail rule severity=3/no emit.
- [ ] RC-02 Folder contracts `db_object/patch/migrations`.
- [ ] RC-03 No `GO` dan forbidden SQL.
- [ ] RC-04 Pairing/version/filename grammar.
- [ ] RC-05 Ordering emit dan aturan `v{k}`.
- [ ] RC-06 Reconstruction fail-hard.
- [ ] RC-07 Data master contracts.
- [ ] RC-08 Destructive policy + skip-vs-stop.
- [ ] RC-09 Transaction + THROW strategy.
- [ ] RC-10 Audit schema + logger contract.
- [ ] RC-11 Wrapper standards by category.
- [ ] RC-12 Down migration policy.
- [ ] RC-13 Patch policy end-to-end.
- [ ] RC-14 `out/output.json` schema and semantics.
- [ ] RC-15 Locking enforcement (`usp_mig_try_get_lock` dipanggil sebelum `migrate up`).
- [ ] RC-16 MVP subset restrictions.
- [ ] RC-17 ALTER COLUMN precheck templates.
- [ ] RC-18 `constraints_naming.md` + repo hygiene.
- [ ] RC-19 CLI contract + exit code.
- [ ] RC-20 Verify integrity final.
- [ ] Cross-tool behavior tetap linear pada extractor, checker apply, checker `-check`, generator, Linux tool, dan `build_all.bat` sesuai scope.
- [ ] Semua fail item dapat didiagnosis dari console CI tanpa wajib membuka `out/output.json`.
- [ ] Patch filename/path behavior telah diuji pada filesystem case-sensitive bila patch contract terdampak.

## Bukti Validasi

- Command yang dijalankan:
  - `rza_migrate_db fmt`
  - `rza_migrate_db verify`
  - `rza_migrate_db gen`
- Ringkasan hasil:
- Link/log test unit:
- Link/log test integration:
- Bukti deterministic output (fixture/golden):
- Dokumen yang diperbarui:
- Dokumen yang direview tetapi tidak berubah (beserta alasan):
- Bukti checksum/build identity/source digest binary Linux:
- Jika `rza_migrate_report` berubah, bukti smoke test Windows/Linux, parity dataset, round-trip status, dan hash `migrations/` sebelum/sesudah:

## Dampak Output dan Migration

- [ ] Tidak ada perubahan output yang non-deterministic.
- [ ] Tidak ada migration baru pada skenario fail.
- [ ] Naming migration tetap patuh format 17 digit version + pairing up/down.
- [ ] Wrapper dan audit call minimal tetap valid untuk tables/indexes/patch.

## Reviewer Sign-off

- [ ] Reviewer engineering setuju.
- [ ] Reviewer QA setuju.
- [ ] Jika ada exception, exception dicatat dengan alasan teknis dan follow-up issue.
