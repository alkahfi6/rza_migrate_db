# SQL_REPLICATE.dbo.MXHolidayDate — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-m-label"></a>
## M_LABEL

**Tipe dasar:** `NVARCHAR(30)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXHolidayDate.sql#L2). **Existing leading index:** MXHolidayDate_CN1.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001970 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoUploadValidation:293](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoUploadValidation.sql#L293) | WHERE / P04 / COMPARISON | `M_LABEL` : `NVARCHAR(30)` | `'JAKARTA'` : `varchar` | `'JAKARTA'` → `NVARCHAR` | `M_LABEL = 'JAKARTA'` |
| IC-001986 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Repo_ValidateField:30](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Repo_ValidateField.sql#L30) | WHERE / P04 / COMPARISON | `M_LABEL` : `NVARCHAR(30)` | `'JAKARTA'` : `varchar` | `'JAKARTA'` → `NVARCHAR` | `M_LABEL = 'JAKARTA'` |
