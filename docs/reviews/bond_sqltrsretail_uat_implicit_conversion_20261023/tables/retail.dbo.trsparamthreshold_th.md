# Retail UAT.dbo.TRSParamThreshold_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-equsdthreshold"></a>
## EqUSDThreshold

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamThreshold_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012034 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamThreshold:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L90) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamthreshold_th.EqUSDThreshold` : `BIGINT` | `tmp.EqUSDThreshold` : `decimal(25,5)` | `tmp.EqUSDThreshold` → `BIGINT` | `retail.dbo.trsparamthreshold_th.EqUSDThreshold <- tmp.EqUSDThreshold` |
| IC-012037 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamThreshold:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L118) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamthreshold_th.EqUSDThreshold` : `BIGINT` | `tmp.EqUSDThreshold` : `decimal(25,5)` | `tmp.EqUSDThreshold` → `BIGINT` | `retail.dbo.trsparamthreshold_th.EqUSDThreshold <- tmp.EqUSDThreshold` |

<a id="kolom-nikprocess"></a>
## NIKProcess

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamThreshold_TH.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012035 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamThreshold:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L90) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamthreshold_th.NIKProcess` : `BIGINT` | `@pnNIK` : `int` | `@pnNIK` → `BIGINT` | `retail.dbo.trsparamthreshold_th.NIKProcess <- @pnNIK` |
| IC-012038 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamThreshold:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L118) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamthreshold_th.NIKProcess` : `BIGINT` | `@pnNIK` : `int` | `@pnNIK` → `BIGINT` | `retail.dbo.trsparamthreshold_th.NIKProcess <- @pnNIK` |
| IC-012042 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamThreshold:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L163) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamthreshold_th.NIKProcess` : `BIGINT` | `@pnNIK` : `int` | `@pnNIK` → `BIGINT` | `retail.dbo.trsparamthreshold_th.NIKProcess <- @pnNIK` |
