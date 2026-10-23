# Retail UAT.dbo.MxBondTran — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 2 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-m-comment-bs"></a>
## M_COMMENT_BS

**Tipe dasar:** `NVARCHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MxBondTran.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006026 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L135) | WHERE / P04 / COMPARISON | `M_COMMENT_BS` : `NVARCHAR(1)` | `@pcTrxType` : `varchar(1)` | `@pcTrxType` → `NVARCHAR` | `M_COMMENT_BS = @pcTrxType` |
| IC-006027 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L135) | WHERE / P04 / IN_VALUE_TYPE | `M_COMMENT_BS` : `NVARCHAR(1)` | `'S'` : `varchar` | `'S'` → `NVARCHAR` | `M_COMMENT_BS in ('S','B')` |

<a id="kolom-m-label"></a>
## M_LABEL

**Tipe dasar:** `VARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MxBondTran.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003636 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCopyDataTrxBONDFromMx:4](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxBONDFromMx.sql#L4) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.mxbondtran.M_LABEL` : `VARCHAR(MAX)` | `tmp.[M_LABEL]` : `[nvarchar](max)` | `tmp.[M_LABEL]` → `VARCHAR(MAX)` | `retail.dbo.mxbondtran.M_LABEL <- tmp.[M_LABEL]` |

<a id="kolom-m-name"></a>
## M_NAME

**Tipe dasar:** `VARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MxBondTran.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003637 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCopyDataTrxBONDFromMx:4](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCopyDataTrxBONDFromMx.sql#L4) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.mxbondtran.M_NAME` : `VARCHAR(MAX)` | `tmp.[M_NAME]` : `[nvarchar](max)` | `tmp.[M_NAME]` → `VARCHAR(MAX)` | `retail.dbo.mxbondtran.M_NAME <- tmp.[M_NAME]` |

<a id="kolom-m-status"></a>
## M_STATUS

**Tipe dasar:** `NVARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.MxBondTran.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006028 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L137) | WHERE / P04 / COMPARISON | `M_STATUS` : `NVARCHAR(3)` | `'INS'` : `varchar` | `'INS'` → `NVARCHAR` | `M_STATUS = 'INS'` |
