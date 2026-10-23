# Retail UAT.dbo.TRSRefundLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSRefundLog_TH.sql#L13). **Existing leading index:** TRSRefundLog_TH_NN1.

**Frekuensi source:** 8 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006532 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateUploadRefund:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateUploadRefund.sql#L67) | CASE / P12 / COMPARISON | `tRefund.Status` : `CHAR(1)` | `'P'` : `varchar` | `tRefund.Status` → `VARCHAR` | `tRefund.Status = 'P'` |
| IC-006533 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateUploadRefund:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateUploadRefund.sql#L68) | CASE / P12 / COMPARISON | `tRefund.Status` : `CHAR(1)` | `'F'` : `varchar` | `tRefund.Status` → `VARCHAR` | `tRefund.Status = 'F'` |
| IC-006534 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateUploadRefund:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateUploadRefund.sql#L69) | CASE / P12 / COMPARISON | `tRefund.Status` : `CHAR(1)` | `'S'` : `varchar` | `tRefund.Status` → `VARCHAR` | `tRefund.Status = 'S'` |
| IC-007745 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessUploadRefund:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessUploadRefund.sql#L122) | WHERE / P12 / COMPARISON | `tRefund.Status` : `CHAR(1)` | `'F'` : `varchar` | `tRefund.Status` → `VARCHAR` | `tRefund.Status <> 'F'` |
| IC-011522 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledRefundSafeKeepingFee:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledRefundSafeKeepingFee.sql#L24) | CASE / P12 / COMPARISON | `rlth.Status` : `CHAR(1)` | `'P'` : `varchar` | `rlth.Status` → `VARCHAR` | `rlth.Status = 'P'` |
| IC-011523 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledRefundSafeKeepingFee:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledRefundSafeKeepingFee.sql#L25) | CASE / P12 / COMPARISON | `rlth.Status` : `CHAR(1)` | `'S'` : `varchar` | `rlth.Status` → `VARCHAR` | `rlth.Status = 'S'` |
| IC-011524 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettledRefundSafeKeepingFee:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettledRefundSafeKeepingFee.sql#L26) | CASE / P12 / COMPARISON | `rlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `rlth.Status` → `VARCHAR` | `rlth.Status = 'F'` |
| IC-012548 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L130) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
