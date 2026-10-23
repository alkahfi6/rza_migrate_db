# Retail UAT.dbo.MxRepoTran — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 1 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-contract"></a>
## CONTRACT

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MxRepoTran.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003639 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L5) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.mxrepotran.CONTRACT` : `BIGINT` | `tmp.CONTRACT` : `DECIMAL(10, 0)` | `tmp.CONTRACT` → `BIGINT` | `retail.dbo.mxrepotran.CONTRACT <- tmp.CONTRACT` |

<a id="kolom-m-creator"></a>
## M_CREATOR

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MxRepoTran.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003640 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L5) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.mxrepotran.M_CREATOR` : `BIGINT` | `tmp.M_CREATOR` : `DECIMAL(10, 0)` | `tmp.M_CREATOR` → `BIGINT` | `retail.dbo.mxrepotran.M_CREATOR <- tmp.M_CREATOR` |

<a id="kolom-m-nb"></a>
## M_NB

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MxRepoTran.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003644 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L63) | JOIN_ON / P11 / COMPARISON | `mst.M_NB` : `BIGINT` | `tmp.M_NB` : `DECIMAL(10, 0)` | `mst.M_NB` → `DECIMAL` | `tmp.M_NB = mst.M_NB` |
| IC-003645 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L71) | JOIN_ON / P11 / COMPARISON | `mst.M_NB` : `BIGINT` | `tmp.M_NB` : `DECIMAL(10, 0)` | `mst.M_NB` → `DECIMAL` | `tmp.M_NB = mst.M_NB` |
| IC-003641 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L5) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.mxrepotran.M_NB` : `BIGINT` | `tmp.M_NB` : `DECIMAL(10, 0)` | `tmp.M_NB` → `BIGINT` | `retail.dbo.mxrepotran.M_NB <- tmp.M_NB` |

<a id="kolom-term"></a>
## Term

**Tipe dasar:** `VARCHAR(15)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MxRepoTran.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003642 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L5) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.mxrepotran.Term` : `VARCHAR(15)` | `tmp.Term` : `NVARCHAR(15)` | `tmp.Term` → `VARCHAR(15)` | `retail.dbo.mxrepotran.Term <- tmp.Term` |

<a id="kolom-x-m-nb"></a>
## X_M_NB

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MxRepoTran.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003643 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCopyDataTrxRepoFromMx:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L5) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.mxrepotran.X_M_NB` : `BIGINT` | `tmp.X_M_NB` : `DECIMAL(10, 0)` | `tmp.X_M_NB` → `BIGINT` | `retail.dbo.mxrepotran.X_M_NB <- tmp.X_M_NB` |
