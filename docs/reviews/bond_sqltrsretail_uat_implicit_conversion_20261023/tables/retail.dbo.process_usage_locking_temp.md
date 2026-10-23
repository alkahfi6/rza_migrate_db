# Retail UAT.dbo.process_usage_locking_temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 1 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-last-update-date"></a>
## last_update_date

**Tipe dasar:** `DATETIME2(7)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.process_usage_locking_temp.sql#L5). **Existing leading index:** process_usage_locking_temp_02.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004931 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSObligasiHousekeepProcess:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiHousekeepProcess.sql#L135) | WHERE / P13 / COMPARISON | `[last_update_date]` : `DATETIME2(7)` | `@dDateClear` : `datetime` | `@dDateClear` → `DATETIME2` | `[last_update_date] < @dDateClear` |
