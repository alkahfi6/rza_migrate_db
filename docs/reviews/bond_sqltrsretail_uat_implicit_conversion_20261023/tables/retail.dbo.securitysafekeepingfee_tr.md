# Retail UAT.dbo.SecuritySafeKeepingFee_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 94 lokasi source, 5 procedure, 14 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 6 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-type1minimumkompensasi"></a>
## Type1MinimumKompensasi

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004005 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGenerateSafeKeepingFeeCustody:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L114) | CASE / P05 / BINARY_Multiply | `sskftr.Type1MinimumKompensasi` : `MONEY` | `(100 + sskftr.Type1PPNKompensasi)` : `DECIMAL (combined result; facets unknown)` | `sskftr.Type1MinimumKompensasi` → `DECIMAL` | `sskftr.Type1MinimumKompensasi * (100 + sskftr.Type1PPNKompensasi)` |
| IC-004007 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGenerateSafeKeepingFeeCustody:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L114) | CASE / P11 / COMPARISON | `sskftr.Type1MinimumKompensasi` : `MONEY` | `tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio` : `FLOAT (combined result; facets unknown)` | `sskftr.Type1MinimumKompensasi` → `FLOAT` | `tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio < sskftr.Type1MinimumKompensasi` |
| IC-004043 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGenerateSafeKeepingFeeCustody_New:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L114) | CASE / P05 / BINARY_Multiply | `sskftr.Type1MinimumKompensasi` : `MONEY` | `(100 + sskftr.Type1PPNKompensasi)` : `DECIMAL (combined result; facets unknown)` | `sskftr.Type1MinimumKompensasi` → `DECIMAL` | `sskftr.Type1MinimumKompensasi * (100 + sskftr.Type1PPNKompensasi)` |
| IC-004045 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGenerateSafeKeepingFeeCustody_New:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L114) | CASE / P11 / COMPARISON | `sskftr.Type1MinimumKompensasi` : `MONEY` | `tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio` : `FLOAT (combined result; facets unknown)` | `sskftr.Type1MinimumKompensasi` → `FLOAT` | `tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio < sskftr.Type1MinimumKompensasi` |
| IC-002724 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:222](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L222) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1MinimumKompensasi` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type1MinimumKompensasi, 0)` |
| IC-013305 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:309](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L309) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1MinimumKompensasi` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type1MinimumKompensasi,0)` |
| IC-013346 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L330) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1MinimumKompensasi` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type1MinimumKompensasi,0)` |

<a id="kolom-type1ppn"></a>
## Type1PPN

**Tipe dasar:** `DECIMAL(13,10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004010 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSGenerateSafeKeepingFeeCustody:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L115) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(100 + sskftr.Type1PPN)` : `DECIMAL (combined result; facets unknown)` | `(tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio)` : `FLOAT (combined result; facets unknown)` | `(100 + sskftr.Type1PPN)` → `FLOAT` | `(tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type1PPN)` |
| IC-004048 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSGenerateSafeKeepingFeeCustody_New:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L115) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(100 + sskftr.Type1PPN)` : `DECIMAL (combined result; facets unknown)` | `(tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio)` : `FLOAT (combined result; facets unknown)` | `(100 + sskftr.Type1PPN)` → `FLOAT` | `(tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type1PPN)` |
| IC-002723 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:221](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L221) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1PPN` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type1PPN, 0)` |
| IC-004008 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L115) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `sskftr.Type1PPN` : `DECIMAL(13,10)` | `100` : `int` | `100` → `DECIMAL` | `100 + sskftr.Type1PPN` |
| IC-004009 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L115) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `(tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type1PPN)` : `FLOAT (combined result; facets unknown)` | `100` : `int` | `100` → `FLOAT` | `(tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type1PPN) / 100` |
| IC-004046 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L115) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `sskftr.Type1PPN` : `DECIMAL(13,10)` | `100` : `int` | `100` → `DECIMAL` | `100 + sskftr.Type1PPN` |
| IC-004047 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L115) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `(tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type1PPN)` : `FLOAT (combined result; facets unknown)` | `100` : `int` | `100` → `FLOAT` | `(tskfcth.Outstanding * sskftr.Type1SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type1PPN) / 100` |
| IC-013304 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:308](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L308) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1PPN` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type1PPN,0)` |
| IC-013345 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:329](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L329) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1PPN` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type1PPN,0)` |

<a id="kolom-type1ppnkompensasi"></a>
## Type1PPNKompensasi

**Tipe dasar:** `DECIMAL(13,10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002725 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L223) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1PPNKompensasi` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type1PPNKompensasi, 0)` |
| IC-004004 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L114) | CASE / P05 / BINARY_Add | `sskftr.Type1PPNKompensasi` : `DECIMAL(13,10)` | `100` : `int` | `100` → `DECIMAL` | `100 + sskftr.Type1PPNKompensasi` |
| IC-004005 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L114) | CASE / P05 / BINARY_Multiply | `(100 + sskftr.Type1PPNKompensasi)` : `DECIMAL (combined result; facets unknown)` | `sskftr.Type1MinimumKompensasi` : `MONEY` | `sskftr.Type1MinimumKompensasi` → `DECIMAL` | `sskftr.Type1MinimumKompensasi * (100 + sskftr.Type1PPNKompensasi)` |
| IC-004042 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L114) | CASE / P05 / BINARY_Add | `sskftr.Type1PPNKompensasi` : `DECIMAL(13,10)` | `100` : `int` | `100` → `DECIMAL` | `100 + sskftr.Type1PPNKompensasi` |
| IC-004043 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L114) | CASE / P05 / BINARY_Multiply | `(100 + sskftr.Type1PPNKompensasi)` : `DECIMAL (combined result; facets unknown)` | `sskftr.Type1MinimumKompensasi` : `MONEY` | `sskftr.Type1MinimumKompensasi` → `DECIMAL` | `sskftr.Type1MinimumKompensasi * (100 + sskftr.Type1PPNKompensasi)` |
| IC-013306 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:310](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L310) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1PPNKompensasi` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type1PPNKompensasi,0)` |
| IC-013347 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:331](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L331) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1PPNKompensasi` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type1PPNKompensasi,0)` |

<a id="kolom-type1safekeepingfee"></a>
## Type1SafeKeepingFee

**Tipe dasar:** `DECIMAL(13,10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002722 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L220) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1SafeKeepingFee` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type1SafeKeepingFee, 0)` |
| IC-004003 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L111) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `sskftr.Type1SafeKeepingFee` : `DECIMAL(13,10)` | `tskfcth.Outstanding` : `money` | `tskfcth.Outstanding` → `DECIMAL` | `tskfcth.Outstanding * sskftr.Type1SafeKeepingFee` |
| IC-004006 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L114) | CASE / P05 / BINARY_Multiply | `sskftr.Type1SafeKeepingFee` : `DECIMAL(13,10)` | `tskfcth.Outstanding` : `money` | `tskfcth.Outstanding` → `DECIMAL` | `tskfcth.Outstanding * sskftr.Type1SafeKeepingFee` |
| IC-004011 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L115) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `sskftr.Type1SafeKeepingFee` : `DECIMAL(13,10)` | `tskfcth.Outstanding` : `money` | `tskfcth.Outstanding` → `DECIMAL` | `tskfcth.Outstanding * sskftr.Type1SafeKeepingFee` |
| IC-004041 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L111) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `sskftr.Type1SafeKeepingFee` : `DECIMAL(13,10)` | `tskfcth.Outstanding` : `money` | `tskfcth.Outstanding` → `DECIMAL` | `tskfcth.Outstanding * sskftr.Type1SafeKeepingFee` |
| IC-004044 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L114) | CASE / P05 / BINARY_Multiply | `sskftr.Type1SafeKeepingFee` : `DECIMAL(13,10)` | `tskfcth.Outstanding` : `money` | `tskfcth.Outstanding` → `DECIMAL` | `tskfcth.Outstanding * sskftr.Type1SafeKeepingFee` |
| IC-004049 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L115) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `sskftr.Type1SafeKeepingFee` : `DECIMAL(13,10)` | `tskfcth.Outstanding` : `money` | `tskfcth.Outstanding` → `DECIMAL` | `tskfcth.Outstanding * sskftr.Type1SafeKeepingFee` |
| IC-013303 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:307](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L307) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1SafeKeepingFee` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type1SafeKeepingFee,0)` |
| IC-013344 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:328](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L328) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type1SafeKeepingFee` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type1SafeKeepingFee,0)` |

<a id="kolom-type2fee"></a>
## Type2Fee

**Tipe dasar:** `DECIMAL(13,10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004013 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGenerateSafeKeepingFeeCustody:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L167) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `b.[Type2Fee]` : `DECIMAL(13,10)` | `a.[Fee]` : `money` | `b.[Type2Fee]` → `money` | `a.[Fee] = b.[Type2Fee]` |
| IC-004051 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGenerateSafeKeepingFeeCustody_New:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L167) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `b.[Type2Fee]` : `DECIMAL(13,10)` | `a.[Fee]` : `money` | `b.[Type2Fee]` → `money` | `a.[Fee] = b.[Type2Fee]` |
| IC-002728 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L226) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2Fee` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type2Fee, 0)` |
| IC-004029 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L265) | DML_CONDITION / P05 / BINARY_Divide | `b.Type2Fee * 1.000` : `DECIMAL (combined result; facets unknown)` | `100` : `int` | `100` → `DECIMAL` | `b.Type2Fee * 1.000 / 100` |
| IC-004031 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L265) | DML_CONDITION / P05 / BINARY_Multiply | `(b.Type2Fee * 1.000 / 100)` : `DECIMAL (combined result; facets unknown)` | `a.[SafeKeepingFeeAmount]` : `money` | `a.[SafeKeepingFeeAmount]` → `DECIMAL` | `a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100)` |
| IC-004067 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L265) | DML_CONDITION / P05 / BINARY_Divide | `b.Type2Fee * 1.000` : `DECIMAL (combined result; facets unknown)` | `100` : `int` | `100` → `DECIMAL` | `b.Type2Fee * 1.000 / 100` |
| IC-004069 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L265) | DML_CONDITION / P05 / BINARY_Multiply | `(b.Type2Fee * 1.000 / 100)` : `DECIMAL (combined result; facets unknown)` | `a.[SafeKeepingFeeAmount]` : `money` | `a.[SafeKeepingFeeAmount]` → `DECIMAL` | `a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100)` |
| IC-013309 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:313](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L313) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2Fee` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type2Fee,0)` |
| IC-013350 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:334](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L334) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2Fee` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type2Fee,0)` |

<a id="kolom-type2nominalfrom"></a>
## Type2NominalFrom

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002726 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:224](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L224) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2NominalFrom` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type2NominalFrom, 0)` |
| IC-013307 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:311](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L311) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2NominalFrom` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type2NominalFrom,0)` |
| IC-013348 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:332](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L332) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2NominalFrom` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type2NominalFrom,0)` |

<a id="kolom-type2nominalto"></a>
## Type2NominalTo

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002727 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L225) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2NominalTo` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type2NominalTo, 0)` |
| IC-013308 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L312) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2NominalTo` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type2NominalTo,0)` |
| IC-013349 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:333](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L333) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2NominalTo` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type2NominalTo,0)` |

<a id="kolom-type2ppn"></a>
## Type2PPN

**Tipe dasar:** `DECIMAL(13,10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004014 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGenerateSafeKeepingFeeCustody:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L168) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `b.[Type2PPN]` : `DECIMAL(13,10)` | `a.[PPN]` : `money` | `b.[Type2PPN]` → `money` | `a.[PPN] = b.[Type2PPN]` |
| IC-004026 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGenerateSafeKeepingFeeCustody:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L265) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `(a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100)) * @nRatio` : `FLOAT (combined result; facets unknown)` | `a.[TotalSKFProduct]` : `money` | `(a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100)) * @nRatio` → `money` | `a.[TotalSKFProduct] = (a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100)) * @nRatio` |
| IC-004052 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGenerateSafeKeepingFeeCustody_New:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L168) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `b.[Type2PPN]` : `DECIMAL(13,10)` | `a.[PPN]` : `money` | `b.[Type2PPN]` → `money` | `a.[PPN] = b.[Type2PPN]` |
| IC-004064 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGenerateSafeKeepingFeeCustody_New:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L265) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `(a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100)) * @nRatio` : `FLOAT (combined result; facets unknown)` | `a.[TotalSKFProduct]` : `money` | `(a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100)) * @nRatio` → `money` | `a.[TotalSKFProduct] = (a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100)) * @nRatio` |
| IC-004030 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSGenerateSafeKeepingFeeCustody:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L265) | DML_CONDITION / P05 / BINARY_Multiply | `(a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100))` : `DECIMAL (combined result; facets unknown)` | `@nRatio` : `float` | `(a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100))` → `FLOAT` | `(a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100)) * @nRatio` |
| IC-004068 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSGenerateSafeKeepingFeeCustody_New:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L265) | DML_CONDITION / P05 / BINARY_Multiply | `(a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100))` : `DECIMAL (combined result; facets unknown)` | `@nRatio` : `float` | `(a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100))` → `FLOAT` | `(a.[SafeKeepingFeeAmount] * (b.Type2Fee * 1.000 / 100) * ((100 + b.Type2PPN) / 100)) * @nRatio` |
| IC-002729 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:227](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L227) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2PPN` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type2PPN, 0)` |
| IC-004027 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L265) | DML_CONDITION / P05 / BINARY_Add | `b.Type2PPN` : `DECIMAL(13,10)` | `100` : `int` | `100` → `DECIMAL` | `100 + b.Type2PPN` |
| IC-004028 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L265) | DML_CONDITION / P05 / BINARY_Divide | `(100 + b.Type2PPN)` : `DECIMAL (combined result; facets unknown)` | `100` : `int` | `100` → `DECIMAL` | `(100 + b.Type2PPN) / 100` |
| IC-004065 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L265) | DML_CONDITION / P05 / BINARY_Add | `b.Type2PPN` : `DECIMAL(13,10)` | `100` : `int` | `100` → `DECIMAL` | `100 + b.Type2PPN` |
| IC-004066 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L265) | DML_CONDITION / P05 / BINARY_Divide | `(100 + b.Type2PPN)` : `DECIMAL (combined result; facets unknown)` | `100` : `int` | `100` → `DECIMAL` | `(100 + b.Type2PPN) / 100` |
| IC-013310 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:314](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L314) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2PPN` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type2PPN,0)` |
| IC-013351 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:335](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L335) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type2PPN` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type2PPN,0)` |

<a id="kolom-type3minimumkompensasi"></a>
## Type3MinimumKompensasi

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004017 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGenerateSafeKeepingFeeCustody:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L192) | CASE / P05 / BINARY_Multiply | `sskftr.Type3MinimumKompensasi` : `MONEY` | `(100 + sskftr.Type3PPNKompensasi)` : `DECIMAL (combined result; facets unknown)` | `sskftr.Type3MinimumKompensasi` → `DECIMAL` | `sskftr.Type3MinimumKompensasi * (100 + sskftr.Type3PPNKompensasi)` |
| IC-004019 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGenerateSafeKeepingFeeCustody:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L192) | CASE / P11 / COMPARISON | `sskftr.Type3MinimumKompensasi` : `MONEY` | `tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio` : `FLOAT (combined result; facets unknown)` | `sskftr.Type3MinimumKompensasi` → `FLOAT` | `tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio < sskftr.Type3MinimumKompensasi` |
| IC-004055 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGenerateSafeKeepingFeeCustody_New:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L192) | CASE / P05 / BINARY_Multiply | `sskftr.Type3MinimumKompensasi` : `MONEY` | `(100 + sskftr.Type3PPNKompensasi)` : `DECIMAL (combined result; facets unknown)` | `sskftr.Type3MinimumKompensasi` → `DECIMAL` | `sskftr.Type3MinimumKompensasi * (100 + sskftr.Type3PPNKompensasi)` |
| IC-004057 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGenerateSafeKeepingFeeCustody_New:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L192) | CASE / P11 / COMPARISON | `sskftr.Type3MinimumKompensasi` : `MONEY` | `tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio` : `FLOAT (combined result; facets unknown)` | `sskftr.Type3MinimumKompensasi` → `FLOAT` | `tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio < sskftr.Type3MinimumKompensasi` |
| IC-002734 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L232) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3MinimumKompensasi` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type3MinimumKompensasi, 0)` |
| IC-013315 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:319](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L319) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3MinimumKompensasi` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type3MinimumKompensasi,0)` |
| IC-013356 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:340](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L340) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3MinimumKompensasi` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type3MinimumKompensasi,0)` |

<a id="kolom-type3nominalfrom"></a>
## Type3NominalFrom

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002730 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L228) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3NominalFrom` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type3NominalFrom, 0)` |
| IC-013311 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:315](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L315) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3NominalFrom` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type3NominalFrom,0)` |
| IC-013352 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:336](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L336) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3NominalFrom` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type3NominalFrom,0)` |

<a id="kolom-type3nominalto"></a>
## Type3NominalTo

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002731 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:229](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L229) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3NominalTo` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type3NominalTo, 0)` |
| IC-013312 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:316](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L316) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3NominalTo` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type3NominalTo,0)` |
| IC-013353 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L337) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3NominalTo` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(sskftr.Type3NominalTo,0)` |

<a id="kolom-type3ppn"></a>
## Type3PPN

**Tipe dasar:** `DECIMAL(13,10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION, COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004022 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSGenerateSafeKeepingFeeCustody:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L193) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(100 + sskftr.Type3PPN)` : `DECIMAL (combined result; facets unknown)` | `(tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio)` : `FLOAT (combined result; facets unknown)` | `(100 + sskftr.Type3PPN)` → `FLOAT` | `(tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type3PPN)` |
| IC-004060 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSGenerateSafeKeepingFeeCustody_New:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L193) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `(100 + sskftr.Type3PPN)` : `DECIMAL (combined result; facets unknown)` | `(tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio)` : `FLOAT (combined result; facets unknown)` | `(100 + sskftr.Type3PPN)` → `FLOAT` | `(tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type3PPN)` |
| IC-002733 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:231](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L231) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3PPN` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type3PPN, 0)` |
| IC-004020 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L193) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `sskftr.Type3PPN` : `DECIMAL(13,10)` | `100` : `int` | `100` → `DECIMAL` | `100 + sskftr.Type3PPN` |
| IC-004021 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L193) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `(tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type3PPN)` : `FLOAT (combined result; facets unknown)` | `100` : `int` | `100` → `FLOAT` | `(tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type3PPN) / 100` |
| IC-004058 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L193) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `sskftr.Type3PPN` : `DECIMAL(13,10)` | `100` : `int` | `100` → `DECIMAL` | `100 + sskftr.Type3PPN` |
| IC-004059 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L193) | EXPRESSION_OR_CONTROL / P05 / BINARY_Divide | `(tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type3PPN)` : `FLOAT (combined result; facets unknown)` | `100` : `int` | `100` → `FLOAT` | `(tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee / 100 * @nRatio) * (100 + sskftr.Type3PPN) / 100` |
| IC-013314 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:318](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L318) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3PPN` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type3PPN,0)` |
| IC-013355 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:339](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L339) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3PPN` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type3PPN,0)` |

<a id="kolom-type3ppnkompensasi"></a>
## Type3PPNKompensasi

**Tipe dasar:** `DECIMAL(13,10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002735 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L233) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3PPNKompensasi` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type3PPNKompensasi, 0)` |
| IC-004016 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L192) | CASE / P05 / BINARY_Add | `sskftr.Type3PPNKompensasi` : `DECIMAL(13,10)` | `100` : `int` | `100` → `DECIMAL` | `100 + sskftr.Type3PPNKompensasi` |
| IC-004017 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L192) | CASE / P05 / BINARY_Multiply | `(100 + sskftr.Type3PPNKompensasi)` : `DECIMAL (combined result; facets unknown)` | `sskftr.Type3MinimumKompensasi` : `MONEY` | `sskftr.Type3MinimumKompensasi` → `DECIMAL` | `sskftr.Type3MinimumKompensasi * (100 + sskftr.Type3PPNKompensasi)` |
| IC-004054 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L192) | CASE / P05 / BINARY_Add | `sskftr.Type3PPNKompensasi` : `DECIMAL(13,10)` | `100` : `int` | `100` → `DECIMAL` | `100 + sskftr.Type3PPNKompensasi` |
| IC-004055 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L192) | CASE / P05 / BINARY_Multiply | `(100 + sskftr.Type3PPNKompensasi)` : `DECIMAL (combined result; facets unknown)` | `sskftr.Type3MinimumKompensasi` : `MONEY` | `sskftr.Type3MinimumKompensasi` → `DECIMAL` | `sskftr.Type3MinimumKompensasi * (100 + sskftr.Type3PPNKompensasi)` |
| IC-013316 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:320](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L320) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3PPNKompensasi` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type3PPNKompensasi,0)` |
| IC-013357 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:341](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L341) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3PPNKompensasi` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type3PPNKompensasi,0)` |

<a id="kolom-type3safekeepingfee"></a>
## Type3SafeKeepingFee

**Tipe dasar:** `DECIMAL(13,10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecuritySafeKeepingFee_TR.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002732 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_trs_ListSecurityMaster_TM:230](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_trs_ListSecurityMaster_TM.sql#L230) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3SafeKeepingFee` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type3SafeKeepingFee, 0)` |
| IC-004015 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L189) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `sskftr.Type3SafeKeepingFee` : `DECIMAL(13,10)` | `tOutstanding.SumOutstanding` : `MONEY (aggregate; facets unknown)` | `tOutstanding.SumOutstanding` → `DECIMAL` | `tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee` |
| IC-004018 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L192) | CASE / P05 / BINARY_Multiply | `sskftr.Type3SafeKeepingFee` : `DECIMAL(13,10)` | `tOutstanding.SumOutstanding` : `MONEY (aggregate; facets unknown)` | `tOutstanding.SumOutstanding` → `DECIMAL` | `tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee` |
| IC-004023 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody.sql#L193) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `sskftr.Type3SafeKeepingFee` : `DECIMAL(13,10)` | `tOutstanding.SumOutstanding` : `MONEY (aggregate; facets unknown)` | `tOutstanding.SumOutstanding` → `DECIMAL` | `tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee` |
| IC-004053 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L189) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `sskftr.Type3SafeKeepingFee` : `DECIMAL(13,10)` | `tOutstanding.SumOutstanding` : `MONEY (aggregate; facets unknown)` | `tOutstanding.SumOutstanding` → `DECIMAL` | `tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee` |
| IC-004056 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L192) | CASE / P05 / BINARY_Multiply | `sskftr.Type3SafeKeepingFee` : `DECIMAL(13,10)` | `tOutstanding.SumOutstanding` : `MONEY (aggregate; facets unknown)` | `tOutstanding.SumOutstanding` → `DECIMAL` | `tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee` |
| IC-004061 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeCustody_New:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeCustody_New.sql#L193) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `sskftr.Type3SafeKeepingFee` : `DECIMAL(13,10)` | `tOutstanding.SumOutstanding` : `MONEY (aggregate; facets unknown)` | `tOutstanding.SumOutstanding` → `DECIMAL` | `tOutstanding.SumOutstanding * sskftr.Type3SafeKeepingFee` |
| IC-013313 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:317](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L317) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3SafeKeepingFee` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type3SafeKeepingFee,0)` |
| IC-013354 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMaster_TM:338](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMaster_TM.sql#L338) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sskftr.Type3SafeKeepingFee` : `DECIMAL(13,10)` | `0` : `int` | `0` → `DECIMAL(13,10)` | `ISNULL(sskftr.Type3SafeKeepingFee,0)` |
