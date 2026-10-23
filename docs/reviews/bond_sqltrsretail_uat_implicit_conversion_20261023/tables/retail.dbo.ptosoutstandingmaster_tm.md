# Retail UAT.dbo.PTOSOutstandingMaster_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `DECIMAL(25,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingMaster_TM.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001053 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSJobMasterOutstanding:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSJobMasterOutstanding.sql#L91) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstandingmaster_tm.Nominal` : `DECIMAL(25,5)` | `Nominal` : `money` | `Nominal` → `DECIMAL(25,5)` | `retail.dbo.ptosoutstandingmaster_tm.Nominal <- Nominal` |
| IC-001716 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSSubmitMasterOutstanding:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L128) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstandingmaster_tm.Nominal` : `DECIMAL(25,5)` | `Nominal` : `money` | `Nominal` → `DECIMAL(25,5)` | `retail.dbo.ptosoutstandingmaster_tm.Nominal <- Nominal` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstandingMaster_TM.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001054 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSJobMasterOutstanding:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSJobMasterOutstanding.sql#L91) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstandingmaster_tm.Status` : `VARCHAR(10)` | `[Status]` : `int` | `[Status]` → `VARCHAR(10)` | `retail.dbo.ptosoutstandingmaster_tm.Status <- [Status]` |
| IC-001717 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSSubmitMasterOutstanding:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L128) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstandingmaster_tm.Status` : `VARCHAR(10)` | `[Status]` : `int` | `[Status]` → `VARCHAR(10)` | `retail.dbo.ptosoutstandingmaster_tm.Status <- [Status]` |
