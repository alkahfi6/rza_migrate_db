# Retail UAT.dbo.SecurityFeeNasabahPending_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-lastprocesscode"></a>
## LastProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabahPending_TM.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007267 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:326](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L326) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabahpending_tm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securityfeenasabahpending_tm.LastProcessCode <- 'P'` |
| IC-007273 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:450](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L450) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabahpending_tm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securityfeenasabahpending_tm.LastProcessCode <- 'P'` |

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabahPending_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007281 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityFeeAuthorization:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFeeAuthorization.sql#L106) | WHERE / P12 / IN_VALUE_TYPE | `sfnptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `sfnptm.OperationCode` → `VARCHAR` | `sfnptm.OperationCode IN ('I', 'U')` |
| IC-007282 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityFeeAuthorization:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFeeAuthorization.sql#L163) | WHERE / P12 / IN_VALUE_TYPE | `sfnptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `sfnptm.OperationCode` → `VARCHAR` | `sfnptm.OperationCode IN ('I', 'U')` |
| IC-007268 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:326](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L326) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabahpending_tm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.securityfeenasabahpending_tm.OperationCode <- 'I'` |
| IC-007274 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:450](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L450) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabahpending_tm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.securityfeenasabahpending_tm.OperationCode <- 'U'` |
