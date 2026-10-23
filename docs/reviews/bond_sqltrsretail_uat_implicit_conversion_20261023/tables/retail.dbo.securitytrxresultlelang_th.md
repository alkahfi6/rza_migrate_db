# Retail UAT.dbo.SecurityTrxResultLelang_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxResultLelang_TH.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003152 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityProcessResultLelang:669](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelang.sql#L669) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytrxresultlelang_th.SecId` : `INT` | `tt.SecId` : `BIGINT` | `tt.SecId` → `INT` | `retail.dbo.securitytrxresultlelang_th.SecId <- tt.SecId` |
| IC-003182 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityProcessResultLelangTA:506](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelangTA.sql#L506) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytrxresultlelang_th.SecId` : `INT` | `tt.SecId` : `BIGINT` | `tt.SecId` → `INT` | `retail.dbo.securitytrxresultlelang_th.SecId <- tt.SecId` |

<a id="kolom-trxid"></a>
## TrxId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTrxResultLelang_TH.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003153 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityProcessResultLelang:669](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelang.sql#L669) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytrxresultlelang_th.TrxId` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `retail.dbo.securitytrxresultlelang_th.TrxId <- 0` |
| IC-003183 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityProcessResultLelangTA:506](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityProcessResultLelangTA.sql#L506) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitytrxresultlelang_th.TrxId` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `retail.dbo.securitytrxresultlelang_th.TrxId <- 0` |
