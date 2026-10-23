# Retail UAT.dbo.REPOApprovalParameterProduct_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-approvestatus"></a>
## ApproveStatus

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOApprovalParameterProduct_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003309 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApprovalParamRepo:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApprovalParamRepo.sql#L97) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.ApproveStatus` : `VARCHAR(10)` | `@pcApprStatus` : `CHAR(1)` | `@pcApprStatus` → `VARCHAR(10)` | `tm.ApproveStatus = @pcApprStatus` |

<a id="kolom-nikinput"></a>
## NIKInput

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOApprovalParameterProduct_TM.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005931 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParamRepo:142](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamRepo.sql#L142) | WHERE / P04 / COMPARISON | `tm.NIKInput` : `INT` | `@pcParam1` : `VARCHAR(20)` | `@pcParam1` → `INT` | `tm.NIKInput <> @pcParam1` |
