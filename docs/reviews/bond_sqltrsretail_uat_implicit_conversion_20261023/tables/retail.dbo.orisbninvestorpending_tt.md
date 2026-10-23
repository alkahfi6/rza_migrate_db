# Retail UAT.dbo.ORISBNInvestorPending_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 4 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorPending_TT.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002950 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPushMailNotifRegInvestor:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifRegInvestor.sql#L65) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[CIFId]` : `BIGINT` | `@nCIFId` : `int` | `[CIFId]` → `int` | `@nCIFId = [CIFId]` |

<a id="kolom-idsubregistry"></a>
## IdSubRegistry

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNInvestorPending_TT.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002821 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateRegInvestor:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L36) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `pnd.[IdSubRegistry]` : `VARCHAR(10)` | `@tmp_Data.IdSubRegistry` : `int` | `pnd.[IdSubRegistry]` → `int` | `@tmp_Data.IdSubRegistry <- pnd.[IdSubRegistry]` |
| IC-002822 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateRegInvestor:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L73) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `pnd.[IdSubRegistry]` : `VARCHAR(10)` | `@tmp_Data.IdSubRegistry` : `int` | `pnd.[IdSubRegistry]` → `int` | `@tmp_Data.IdSubRegistry <- pnd.[IdSubRegistry]` |
| IC-002823 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SBNSchedPopulateRegInvestor:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPopulateRegInvestor.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `pnd.[IdSubRegistry]` : `VARCHAR(10)` | `@tmp_Data.IdSubRegistry` : `int` | `pnd.[IdSubRegistry]` → `int` | `@tmp_Data.IdSubRegistry <- pnd.[IdSubRegistry]` |
