# SQL_REPLICATE.dbo.Mx_RepoTran_TMP — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 9 lokasi source, 1 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 5 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-contract"></a>
## CONTRACT

**Tipe dasar:** `DECIMAL(10, 0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.Mx_RepoTran_TMP.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003639 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCopyDataTrxRepoFromMx:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L5) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.CONTRACT` : `DECIMAL(10, 0)` | `retail.dbo.mxrepotran.CONTRACT` : `BIGINT` | `tmp.CONTRACT` → `BIGINT` | `retail.dbo.mxrepotran.CONTRACT <- tmp.CONTRACT` |

<a id="kolom-m-creator"></a>
## M_CREATOR

**Tipe dasar:** `DECIMAL(10, 0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.Mx_RepoTran_TMP.sql#L3). **Existing leading index:** ix_mxrepotran_tmp_m_creator.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003640 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCopyDataTrxRepoFromMx:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L5) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.M_CREATOR` : `DECIMAL(10, 0)` | `retail.dbo.mxrepotran.M_CREATOR` : `BIGINT` | `tmp.M_CREATOR` → `BIGINT` | `retail.dbo.mxrepotran.M_CREATOR <- tmp.M_CREATOR` |
| IC-003646 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCopyDataTrxRepoFromMx:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L81) | JOIN_ON / P11 / COMPARISON | `M_CREATOR` : `DECIMAL(10, 0)` | `0` : `int` | `0` → `DECIMAL` | `M_CREATOR!=0` |
| IC-003647 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCopyDataTrxRepoFromMx:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L81) | JOIN_ON / P11 / COMPARISON | `tmp.M_CREATOR` : `DECIMAL(10, 0)` | `ctp.MurexNo` : `BIGINT` | `ctp.MurexNo` → `DECIMAL` | `tmp.M_CREATOR=ctp.MurexNo` |

<a id="kolom-m-nb"></a>
## M_NB

**Tipe dasar:** `DECIMAL(10, 0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.Mx_RepoTran_TMP.sql#L2). **Existing leading index:** ix_mxrepotran_tmp_m_nb_status&#124;ix_mxrepotran_tmp_m_nb.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003641 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCopyDataTrxRepoFromMx:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L5) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.M_NB` : `DECIMAL(10, 0)` | `retail.dbo.mxrepotran.M_NB` : `BIGINT` | `tmp.M_NB` → `BIGINT` | `retail.dbo.mxrepotran.M_NB <- tmp.M_NB` |
| IC-003644 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCopyDataTrxRepoFromMx:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L63) | JOIN_ON / P11 / COMPARISON | `tmp.M_NB` : `DECIMAL(10, 0)` | `mst.M_NB` : `BIGINT` | `mst.M_NB` → `DECIMAL` | `tmp.M_NB = mst.M_NB` |
| IC-003645 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCopyDataTrxRepoFromMx:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L71) | JOIN_ON / P11 / COMPARISON | `tmp.M_NB` : `DECIMAL(10, 0)` | `mst.M_NB` : `BIGINT` | `mst.M_NB` → `DECIMAL` | `tmp.M_NB = mst.M_NB` |

<a id="kolom-term"></a>
## Term

**Tipe dasar:** `NVARCHAR(15)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.Mx_RepoTran_TMP.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003642 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCopyDataTrxRepoFromMx:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L5) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.Term` : `NVARCHAR(15)` | `retail.dbo.mxrepotran.Term` : `VARCHAR(15)` | `tmp.Term` → `VARCHAR(15)` | `retail.dbo.mxrepotran.Term <- tmp.Term` |

<a id="kolom-x-m-nb"></a>
## X_M_NB

**Tipe dasar:** `DECIMAL(10, 0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.Mx_RepoTran_TMP.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003643 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCopyDataTrxRepoFromMx:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxRepoFromMx.sql#L5) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.X_M_NB` : `DECIMAL(10, 0)` | `retail.dbo.mxrepotran.X_M_NB` : `BIGINT` | `tmp.X_M_NB` → `BIGINT` | `retail.dbo.mxrepotran.X_M_NB <- tmp.X_M_NB` |
