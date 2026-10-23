# SQL_REPLICATE.dbo.Mx_BondTran_TMP — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 1 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-m-creator"></a>
## M_CREATOR

**Tipe dasar:** `[bigint]`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.Mx_BondTran_TMP.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003638 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCopyDataTrxBONDFromMx:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxBONDFromMx.sql#L86) | JOIN_ON / P11 / COMPARISON | `M_CREATOR` : `[bigint]` | `0` : `int` | `0` → `BIGINT` | `M_CREATOR!=0` |

<a id="kolom-m-label"></a>
## M_LABEL

**Tipe dasar:** `[nvarchar](max)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.Mx_BondTran_TMP.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003636 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCopyDataTrxBONDFromMx:4](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxBONDFromMx.sql#L4) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.[M_LABEL]` : `[nvarchar](max)` | `retail.dbo.mxbondtran.M_LABEL` : `VARCHAR(MAX)` | `tmp.[M_LABEL]` → `VARCHAR(MAX)` | `retail.dbo.mxbondtran.M_LABEL <- tmp.[M_LABEL]` |

<a id="kolom-m-name"></a>
## M_NAME

**Tipe dasar:** `[nvarchar](max)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.Mx_BondTran_TMP.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003637 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCopyDataTrxBONDFromMx:4](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxBONDFromMx.sql#L4) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tmp.[M_NAME]` : `[nvarchar](max)` | `retail.dbo.mxbondtran.M_NAME` : `VARCHAR(MAX)` | `tmp.[M_NAME]` → `VARCHAR(MAX)` | `retail.dbo.mxbondtran.M_NAME <- tmp.[M_NAME]` |
