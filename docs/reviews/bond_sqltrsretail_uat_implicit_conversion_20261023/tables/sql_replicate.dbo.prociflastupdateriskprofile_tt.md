# SQL_REPLICATE.dbo.ProCIFLastUpdateRiskProfile_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-lastupdateriskprofile"></a>
## LastUpdateRiskProfile

**Tipe dasar:** `VARCHAR(8)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.ProCIFLastUpdateRiskProfile_TT.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012718 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSValidateSecAccNo:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSValidateSecAccNo.sql#L23) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `LastUpdateRiskProfile` : `VARCHAR(8)` | `@dLastUpdateRisk` : `datetime` | `LastUpdateRiskProfile` → `datetime` | `@dLastUpdateRisk = LastUpdateRiskProfile` |

<a id="kolom-nocif"></a>
## NoCIF

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.ProCIFLastUpdateRiskProfile_TT.sql#L2). **Existing leading index:** ProCIFLastUpdateRiskProfile_TT_CN1.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013421 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusORI_Order_TT:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L66) | WHERE / P01 / COMPARISON | `NoCIF` : `BIGINT` | `tc.CIFNo` : `VARCHAR(19)` | `tc.CIFNo` → `BIGINT` | `NoCIF = tc.CIFNo` |
| IC-013428 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListStatusORI_Order_TT:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListStatusORI_Order_TT.sql#L143) | WHERE / P01 / COMPARISON | `NoCIF` : `BIGINT` | `tc.CIFNo` : `VARCHAR(19)` | `tc.CIFNo` → `BIGINT` | `NoCIF = tc.CIFNo` |
