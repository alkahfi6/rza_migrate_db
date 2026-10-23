# Retail UAT.dbo.TRSMateraiTransaction_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-changeby"></a>
## ChangeBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TH.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001860 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:318](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L318) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_th.ChangeBy` : `INT` | `@pnUserNik` : `BIGINT` | `@pnUserNik` → `INT` | `retail.dbo.trsmateraitransaction_th.ChangeBy <- @pnUserNik` |
| IC-001865 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.RepoProsesCancelTrx:350](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoProsesCancelTrx.sql#L350) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_th.ChangeBy` : `INT` | `@pnUserNik` : `BIGINT` | `@pnUserNik` → `INT` | `retail.dbo.trsmateraitransaction_th.ChangeBy <- @pnUserNik` |

<a id="kolom-trxstatus"></a>
## TrxStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSMateraiTransaction_TH.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003017 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityLelangTrxSubmit:331](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityLelangTrxSubmit.sql#L331) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_th.TrxStatus` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.trsmateraitransaction_th.TrxStatus <- 0` |
| IC-004398 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:2190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L2190) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsmateraitransaction_th.TrxStatus` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.trsmateraitransaction_th.TrxStatus <- 0` |
