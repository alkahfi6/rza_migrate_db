# Retail UAT.dbo.CustodyUploadSfeeKupon_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 4 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-distributiondate"></a>
## DistributionDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CustodyUploadSfeeKupon_TR.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012853 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSValidateUploadSFeeKupon:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadSFeeKupon.sql#L155) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `DistributionDate` : `DATETIME` | `#tmpDataApprove.DistributionDate` : `VARCHAR(40)` | `DistributionDate` → `VARCHAR(40)` | `#tmpDataApprove.DistributionDate <- DistributionDate` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `TINYINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CustodyUploadSfeeKupon_TR.sql#L40). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005267 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalSFeeKupon:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalSFeeKupon.sql#L124) | WHERE / P11 / COMPARISON | `[Status]` : `TINYINT` | `0` : `int` | `[Status]` → `INT` | `[Status] = 0` |
| IC-006695 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovedUploadSFeeKupon:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSFeeKupon.sql#L27) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `TINYINT` | `1` : `int` | `1` → `TINYINT` | `[Status] = 1` |
| IC-006698 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovedUploadSFeeKupon:148](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSFeeKupon.sql#L148) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `TINYINT` | `99` : `int` | `99` → `TINYINT` | `[Status] = 99` |
| IC-012376 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSubmitUploadSfeeKupon:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitUploadSfeeKupon.sql#L155) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.custodyuploadsfeekupon_tr.Status` : `TINYINT` | `[Status]` : `INT` | `[Status]` → `TINYINT` | `retail.dbo.custodyuploadsfeekupon_tr.Status <- [Status]` |
| IC-012854 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSValidateUploadSFeeKupon:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateUploadSFeeKupon.sql#L167) | WHERE / P04 / COMPARISON | `[Status]` : `TINYINT` | `'0'` : `varchar` | `'0'` → `TINYINT` | `[Status] = '0'` |
