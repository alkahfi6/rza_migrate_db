# Retail UAT.dbo.TRSSimanisTransaction_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 3 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSimanisTransaction_TT.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004409 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertSimanis:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSimanis.sql#L11) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssimanistransaction_tt.DealId` : `BIGINT` | `@pcDealId` : `varchar (15)` | `@pcDealId` → `BIGINT` | `retail.dbo.trssimanistransaction_tt.DealId <- @pcDealId` |
| IC-006392 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSimanis:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSimanis.sql#L6) | WHERE / P04 / COMPARISON | `DealId` : `BIGINT` | `@pcDealId` : `varchar (15)` | `@pcDealId` → `BIGINT` | `DealId = @pcDealId` |
| IC-006393 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSimanis:8](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSimanis.sql#L8) | WHERE / P04 / COMPARISON | `DealId` : `BIGINT` | `@pcDealId` : `varchar (15)` | `@pcDealId` → `BIGINT` | `DealId = @pcDealId` |
| IC-012635 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateSimanis:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateSimanis.sql#L9) | WHERE / P04 / COMPARISON | `DealId` : `BIGINT` | `@pcDealId` : `varchar (15)` | `@pcDealId` → `BIGINT` | `DealId = @pcDealId` |
| IC-012636 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateSimanis:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateSimanis.sql#L13) | WHERE / P04 / COMPARISON | `DealId` : `BIGINT` | `@pcDealId` : `varchar (15)` | `@pcDealId` → `BIGINT` | `DealId = @pcDealId` |
| IC-012638 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateSimanis:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateSimanis.sql#L28) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssimanistransaction_tt.DealId` : `BIGINT` | `@pcDealId` : `varchar (15)` | `@pcDealId` → `BIGINT` | `retail.dbo.trssimanistransaction_tt.DealId <- @pcDealId` |
