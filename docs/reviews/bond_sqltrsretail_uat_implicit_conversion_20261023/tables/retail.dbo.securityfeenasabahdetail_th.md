# Retail UAT.dbo.SecurityFeeNasabahDetail_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 3 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-flagcode"></a>
## FlagCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabahDetail_TH.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009534 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterFeeNasabah:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterFeeNasabah.sql#L170) | WHERE / P12 / COMPARISON | `sfnd.FlagCode` : `CHAR(1)` | `'N'` : `varchar` | `sfnd.FlagCode` → `VARCHAR` | `sfnd.FlagCode = 'N'` |
| IC-009537 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptAuditTrailParameterFeeNasabah:253](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptAuditTrailParameterFeeNasabah.sql#L253) | JOIN_ON / P12 / COMPARISON | `sfnd.FlagCode` : `CHAR(1)` | `'O'` : `varchar` | `sfnd.FlagCode` → `VARCHAR` | `sfnd.FlagCode = 'O'` |
| IC-007277 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:512](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L512) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabahdetail_th.FlagCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.securityfeenasabahdetail_th.FlagCode <- 'O'` |
| IC-007283 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFeeAuthorization:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFeeAuthorization.sql#L165) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabahdetail_th.FlagCode` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.securityfeenasabahdetail_th.FlagCode <- 'N'` |
