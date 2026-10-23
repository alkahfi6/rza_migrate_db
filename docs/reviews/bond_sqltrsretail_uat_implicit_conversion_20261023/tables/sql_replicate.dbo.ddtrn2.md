# SQL_REPLICATE.dbo.DDTRN2 — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 12 lokasi source, 1 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 8 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dorc"></a>
## DORC

**Tipe dasar:** `VARCHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDTRN2.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008211 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:273](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L273) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `DORC` : `VARCHAR(1)` | `#tempDDHIST.TRDORC` : `CHAR(1)` | `DORC` → `CHAR(1)` | `#tempDDHIST.TRDORC <- DORC` |
| IC-008223 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:318](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L318) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `DORC` : `VARCHAR(1)` | `#tempDDHIST.TRDORC` : `CHAR(1)` | `DORC` → `CHAR(1)` | `#tempDDHIST.TRDORC <- DORC` |
| IC-008226 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:328](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L328) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `DORC` : `VARCHAR(1)` | `#tempDDHIST.TRDORC` : `CHAR(1)` | `DORC` → `CHAR(1)` | `#tempDDHIST.TRDORC <- DORC` |
| IC-008229 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:338](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L338) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `DORC` : `VARCHAR(1)` | `#tempDDHIST.TRDORC` : `CHAR(1)` | `DORC` → `CHAR(1)` | `#tempDDHIST.TRDORC <- DORC` |

<a id="kolom-tracct"></a>
## TRACCT

**Tipe dasar:** `DECIMAL(18,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDTRN2.sql#L8). **Existing leading index:** idx_tracct_tratyp.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008213 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:279](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L279) | WHERE / P11 / COMPARISON | `TRACCT` : `DECIMAL(18,0)` | `@nNorekGNCTWO` : `BIGINT` | `@nNorekGNCTWO` → `DECIMAL` | `TRACCT = @nNorekGNCTWO` |
| IC-008225 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:324](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L324) | WHERE / P11 / COMPARISON | `TRACCT` : `DECIMAL(18,0)` | `@nNorekGNCTWO` : `BIGINT` | `@nNorekGNCTWO` → `DECIMAL` | `TRACCT = @nNorekGNCTWO` |
| IC-008228 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:334](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L334) | WHERE / P11 / COMPARISON | `TRACCT` : `DECIMAL(18,0)` | `@nNorekGNCTWO` : `BIGINT` | `@nNorekGNCTWO` → `DECIMAL` | `TRACCT = @nNorekGNCTWO` |
| IC-008231 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:344](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L344) | WHERE / P11 / COMPARISON | `TRACCT` : `DECIMAL(18,0)` | `@nNorekGNCTWO` : `BIGINT` | `@nNorekGNCTWO` → `DECIMAL` | `TRACCT = @nNorekGNCTWO` |

<a id="kolom-trremk"></a>
## TRREMK

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDTRN2.sql#L27). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008212 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:273](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L273) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `TRREMK` : `VARCHAR(40)` | `#tempDDHIST.TRREMK` : `CHAR(40)` | `TRREMK` → `CHAR(40)` | `#tempDDHIST.TRREMK <- TRREMK` |
| IC-008224 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:318](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L318) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `TRREMK` : `VARCHAR(40)` | `#tempDDHIST.TRREMK` : `CHAR(40)` | `TRREMK` → `CHAR(40)` | `#tempDDHIST.TRREMK <- TRREMK` |
| IC-008227 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:328](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L328) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `TRREMK` : `VARCHAR(40)` | `#tempDDHIST.TRREMK` : `CHAR(40)` | `TRREMK` → `CHAR(40)` | `#tempDDHIST.TRREMK <- TRREMK` |
| IC-008230 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:338](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L338) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `TRREMK` : `VARCHAR(40)` | `#tempDDHIST.TRREMK` : `CHAR(40)` | `TRREMK` → `CHAR(40)` | `#tempDDHIST.TRREMK <- TRREMK` |
