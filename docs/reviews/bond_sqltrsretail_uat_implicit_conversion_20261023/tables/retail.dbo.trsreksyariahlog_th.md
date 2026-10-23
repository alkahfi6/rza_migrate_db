# Retail UAT.dbo.TRSRekSyariahLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSRekSyariahLog_TH.sql#L3). **Existing leading index:** TRSRekSyariahLog_TH_CN1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002894 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:913](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L913) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsreksyariahlog_th.CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `retail.dbo.trsreksyariahlog_th.CIFId <- @nCIFId` |

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSRekSyariahLog_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002876 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:427](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L427) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsreksyariahlog_th.OperationCode` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.trsreksyariahlog_th.OperationCode <- 'O'` |
| IC-002892 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:891](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L891) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsreksyariahlog_th.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.trsreksyariahlog_th.OperationCode <- 'U'` |
| IC-002895 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:913](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L913) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsreksyariahlog_th.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.trsreksyariahlog_th.OperationCode <- 'I'` |
