# SQL_REPLICATE.dbo.JHFXDT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 23 lokasi source, 4 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 3 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-jfxdcr"></a>
## JFXDCR

**Tipe dasar:** `DECIMAL(13,7)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.JHFXDT.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 18 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003535 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCalculateMateraiFee:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L15) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `JFXDCR` : `DECIMAL(13,7)` | `@nRate_TAX` : `MONEY` | `JFXDCR` → `MONEY` | `@nRate_TAX = JFXDCR` |
| IC-008155 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L72) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `jh.JFXDCR` : `DECIMAL(13,7)` | `#tempResult.RateTax` : `MONEY` | `jh.JFXDCR` → `MONEY` | `#tempResult.RateTax <- jh.JFXDCR` |
| IC-008195 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSReportRekonMaterai:202](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L202) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `jh.JFXDCR` : `DECIMAL(13,7)` | `#tempResult.RateTax` : `MONEY` | `jh.JFXDCR` → `MONEY` | `#tempResult.RateTax <- jh.JFXDCR` |
| IC-007839 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L278) | EXPRESSION / P09 / NULLIF_COMPARE | `jh.JFXDCR` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `nullif(jh.JFXDCR, 0)` |
| IC-007840 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:279](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L279) | EXPRESSION / P09 / NULLIF_COMPARE | `jh.JFXDCR` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `nullif(jh.JFXDCR, 0)` |
| IC-007841 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L281) | EXPRESSION / P09 / NULLIF_COMPARE | `jh.JFXDCR` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `nullif(jh.JFXDCR, 0)` |
| IC-008161 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L97) | EXPRESSION / P09 / NULLIF_COMPARE | `jh.JFXDCR` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `NULLIF(jh.JFXDCR, 0)` |
| IC-008163 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L98) | EXPRESSION / P09 / NULLIF_COMPARE | `jh.JFXDCR` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `NULLIF(jh.JFXDCR, 0)` |
| IC-008165 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L99) | EXPRESSION / P09 / NULLIF_COMPARE | `jh.JFXDCR` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `NULLIF(jh.JFXDCR, 0)` |
| IC-008197 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L220) | CASE / P11 / COMPARISON | `isnull(nullif(jh.JFXDCR, 0), 0)` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `isnull(nullif(jh.JFXDCR, 0), 0) = 0` |
| IC-008198 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportRekonMaterai:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L220) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `nullif(jh.JFXDCR, 0)` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL(13,7)` | `isnull(nullif(jh.JFXDCR, 0), 0)` |
| IC-008199 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L220) | EXPRESSION / P09 / NULLIF_COMPARE | `jh.JFXDCR` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `nullif(jh.JFXDCR, 0)` |
| IC-008200 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L225) | CASE / P11 / COMPARISON | `isnull(nullif(jh.JFXDCR, 0), 0)` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `isnull(nullif(jh.JFXDCR, 0), 0) = 0` |
| IC-008201 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportRekonMaterai:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L225) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `nullif(jh.JFXDCR, 0)` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL(13,7)` | `isnull(nullif(jh.JFXDCR, 0), 0)` |
| IC-008202 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L225) | EXPRESSION / P09 / NULLIF_COMPARE | `jh.JFXDCR` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `nullif(jh.JFXDCR, 0)` |
| IC-008203 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L231) | CASE / P11 / COMPARISON | `isnull(nullif(jh.JFXDCR, 0), 0)` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `isnull(nullif(jh.JFXDCR, 0), 0) = 0` |
| IC-008204 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSReportRekonMaterai:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L231) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `nullif(jh.JFXDCR, 0)` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL(13,7)` | `isnull(nullif(jh.JFXDCR, 0), 0)` |
| IC-008205 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L231) | EXPRESSION / P09 / NULLIF_COMPARE | `jh.JFXDCR` : `DECIMAL(13,7)` | `0` : `int` | `0` → `DECIMAL` | `nullif(jh.JFXDCR, 0)` |

<a id="kolom-jfxdrn"></a>
## JFXDRN

**Tipe dasar:** `DECIMAL(5,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.JHFXDT.sql#L4). **Existing leading index:** idx_jhfxdt_jfxdrn.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003114 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SecurityProcessLelangToTransaction:422](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessLelangToTransaction.sql#L422) | JOIN_ON / P11 / COMPARISON | `jh.JFXDRN` : `DECIMAL(5,0)` | `17` : `int` | `17` → `DECIMAL` | `jh.JFXDRN = 17` |
| IC-003534 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateMateraiFee:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L15) | WHERE / P11 / COMPARISON | `JFXDRN` : `DECIMAL(5,0)` | `17` : `int` | `17` → `DECIMAL` | `JFXDRN = 17` |
| IC-007842 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L291) | JOIN_ON / P11 / COMPARISON | `JFXDRN` : `DECIMAL(5,0)` | `17` : `int` | `17` → `DECIMAL` | `JFXDRN = 17` |
| IC-008166 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L111) | JOIN_ON / P11 / COMPARISON | `JFXDRN` : `DECIMAL(5,0)` | `17` : `int` | `17` → `DECIMAL` | `JFXDRN = 17` |
| IC-008206 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:246](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L246) | JOIN_ON / P11 / COMPARISON | `JFXDRN` : `DECIMAL(5,0)` | `17` : `int` | `17` → `DECIMAL` | `JFXDRN = 17` |
