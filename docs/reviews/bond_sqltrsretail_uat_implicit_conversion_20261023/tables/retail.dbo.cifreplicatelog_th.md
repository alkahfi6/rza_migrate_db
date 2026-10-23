# Retail UAT.dbo.CIFReplicateLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 2 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifno"></a>
## CIFNo

**Tipe dasar:** `DECIMAL(19,0)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CIFReplicateLog_TH.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004748 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMxSyncCIFOnlineInsert:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMxSyncCIFOnlineInsert.sql#L23) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.cifreplicatelog_th.CIFNo` : `DECIMAL(19,0)` | `@pcCIFNo` : `varchar(19)` | `@pcCIFNo` → `DECIMAL(19,0)` | `retail.dbo.cifreplicatelog_th.CIFNo <- @pcCIFNo` |

<a id="kolom-guid"></a>
## Guid

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CIFReplicateLog_TH.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004749 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMxSyncCIFOnlineInsert:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMxSyncCIFOnlineInsert.sql#L23) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.cifreplicatelog_th.Guid` : `VARCHAR(40)` | `@uGuid` : `uniqueidentifier` | `@uGuid` → `VARCHAR(40)` | `retail.dbo.cifreplicatelog_th.Guid <- @uGuid` |

<a id="kolom-processnik"></a>
## ProcessNIK

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CIFReplicateLog_TH.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004751 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMxSyncCIFProcess:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMxSyncCIFProcess.sql#L131) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.cifreplicatelog_th.ProcessNIK` : `INT` | `'777'` : `varchar` | `'777'` → `INT` | `retail.dbo.cifreplicatelog_th.ProcessNIK <- '777'` |
