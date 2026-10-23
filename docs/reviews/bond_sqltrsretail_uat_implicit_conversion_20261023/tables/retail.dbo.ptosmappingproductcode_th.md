# Retail UAT.dbo.PTOSMappingProductCode_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-nikupdate"></a>
## NIKUpdate

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSMappingProductCode_TH.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001695 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSSubmitMappingProduk:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMappingProduk.sql#L46) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosmappingproductcode_th.NIKUpdate` : `BIGINT` | `@pnNIK` : `int` | `@pnNIK` → `BIGINT` | `retail.dbo.ptosmappingproductcode_th.NIKUpdate <- @pnNIK` |
| IC-001696 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSSubmitMappingProduk:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMappingProduk.sql#L58) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosmappingproductcode_th.NIKUpdate` : `BIGINT` | `@pnNIK` : `int` | `@pnNIK` → `BIGINT` | `retail.dbo.ptosmappingproductcode_th.NIKUpdate <- @pnNIK` |
| IC-001699 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSSubmitMappingProduk:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMappingProduk.sql#L68) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosmappingproductcode_th.NIKUpdate` : `BIGINT` | `@pnNIK` : `int` | `@pnNIK` → `BIGINT` | `retail.dbo.ptosmappingproductcode_th.NIKUpdate <- @pnNIK` |
| IC-001701 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSSubmitMappingProduk:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMappingProduk.sql#L73) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosmappingproductcode_th.NIKUpdate` : `BIGINT` | `@pnNIK` : `int` | `@pnNIK` → `BIGINT` | `retail.dbo.ptosmappingproductcode_th.NIKUpdate <- @pnNIK` |
