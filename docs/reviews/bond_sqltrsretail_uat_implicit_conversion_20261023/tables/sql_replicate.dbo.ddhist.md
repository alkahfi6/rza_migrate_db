# SQL_REPLICATE.dbo.DDHIST — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 12 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-tracct"></a>
## TRACCT

**Tipe dasar:** `DECIMAL(18,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDHIST.sql#L3). **Existing leading index:** idx_tracct_tratyp_trdate_serial.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008208 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:266](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L266) | WHERE / P11 / COMPARISON | `TRACCT` : `DECIMAL(18,0)` | `@nNorekGNCTWO` : `BIGINT` | `@nNorekGNCTWO` → `DECIMAL` | `TRACCT = @nNorekGNCTWO` |
| IC-008214 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:290](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L290) | WHERE / P11 / COMPARISON | `TRACCT` : `DECIMAL(18,0)` | `@nNorekGNCTWO` : `BIGINT` | `@nNorekGNCTWO` → `DECIMAL` | `TRACCT = @nNorekGNCTWO` |
| IC-008217 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:302](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L302) | WHERE / P11 / COMPARISON | `TRACCT` : `DECIMAL(18,0)` | `@nNorekGNCTWO` : `BIGINT` | `@nNorekGNCTWO` → `DECIMAL` | `TRACCT = @nNorekGNCTWO` |
| IC-008220 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:313](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L313) | WHERE / P11 / COMPARISON | `TRACCT` : `DECIMAL(18,0)` | `@nNorekGNCTWO` : `BIGINT` | `@nNorekGNCTWO` → `DECIMAL` | `TRACCT = @nNorekGNCTWO` |

<a id="kolom-trdate"></a>
## TRDATE

**Tipe dasar:** `DECIMAL(18,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDHIST.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008209 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:270](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L270) | WHERE / P11 / COMPARISON | `a.TRDATE` : `DECIMAL(18,0)` | `@nJulianDateTo` : `INT` | `@nJulianDateTo` → `DECIMAL` | `a.TRDATE <= @nJulianDateTo` |
| IC-008210 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:271](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L271) | WHERE / P11 / COMPARISON | `a.TRDATE` : `DECIMAL(18,0)` | `@nJulianDateFrom` : `INT` | `@nJulianDateFrom` → `DECIMAL` | `a.TRDATE >= @nJulianDateFrom` |
| IC-008215 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:293](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L293) | WHERE / P11 / COMPARISON | `a.TRDATE` : `DECIMAL(18,0)` | `@nJulianDateTo` : `INT` | `@nJulianDateTo` → `DECIMAL` | `a.TRDATE <= @nJulianDateTo` |
| IC-008216 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:294](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L294) | WHERE / P11 / COMPARISON | `a.TRDATE` : `DECIMAL(18,0)` | `@nJulianDateFrom` : `INT` | `@nJulianDateFrom` → `DECIMAL` | `a.TRDATE >= @nJulianDateFrom` |
| IC-008218 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:304](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L304) | WHERE / P11 / COMPARISON | `a.TRDATE` : `DECIMAL(18,0)` | `@nJulianDateTo` : `INT` | `@nJulianDateTo` → `DECIMAL` | `a.TRDATE <= @nJulianDateTo` |
| IC-008219 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:305](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L305) | WHERE / P11 / COMPARISON | `a.TRDATE` : `DECIMAL(18,0)` | `@nJulianDateFrom` : `INT` | `@nJulianDateFrom` → `DECIMAL` | `a.TRDATE >= @nJulianDateFrom` |
| IC-008221 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:315](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L315) | WHERE / P11 / COMPARISON | `a.TRDATE` : `DECIMAL(18,0)` | `@nJulianDateTo` : `INT` | `@nJulianDateTo` → `DECIMAL` | `a.TRDATE <= @nJulianDateTo` |
| IC-008222 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSReportRekonMaterai:316](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportRekonMaterai.sql#L316) | WHERE / P11 / COMPARISON | `a.TRDATE` : `DECIMAL(18,0)` | `@nJulianDateFrom` : `INT` | `@nJulianDateFrom` → `DECIMAL` | `a.TRDATE >= @nJulianDateFrom` |
