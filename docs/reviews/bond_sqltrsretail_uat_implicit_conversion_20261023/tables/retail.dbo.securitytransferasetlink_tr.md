# Retail UAT.dbo.SecurityTransferAsetLink_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 21 lokasi source, 5 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 9 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAsetLink_TR.sql#L2). **Existing leading index:** SecurityTransferAsetLink_TR_CN1.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006850 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.DealId` : `INT` | `retail.dbo.securitytransaction_tt.PODealId` : `BIGINT` | `tr.DealId` → `BIGINT` | `retail.dbo.securitytransaction_tt.PODealId <- tr.DealId` |
| IC-006859 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L85) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.DealId` : `INT` | `retail.dbo.securitytransaction_tt.PODealId` : `BIGINT` | `tr.DealId` → `BIGINT` | `retail.dbo.securitytransaction_tt.PODealId <- tr.DealId` |
| IC-006869 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.DealId` : `INT` | `retail.dbo.securitytransaction_tt.PODealId` : `BIGINT` | `tr.DealId` → `BIGINT` | `retail.dbo.securitytransaction_tt.PODealId <- tr.DealId` |
| IC-006447 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateTransactionBalance:270](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransactionBalance.sql#L270) | WHERE / P11 / COMPARISON | `b.DealId` : `INT` | `@pnDealIdFOP` : `BIGINT` | `b.DealId` → `BIGINT` | `b.DealId != @pnDealIdFOP` |

<a id="kolom-dealidbankjual"></a>
## DealIdBankJual

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAsetLink_TR.sql#L3). **Existing leading index:** SecurityTransferAsetLink_TR_N1.

**Frekuensi source:** 10 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006867 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.DealIdBankJual` : `INT` | `retail.dbo.securitytransaction_tt.FOPOriDealId` : `BIGINT` | `tr.DealIdBankJual` → `BIGINT` | `retail.dbo.securitytransaction_tt.FOPOriDealId <- tr.DealIdBankJual` |
| IC-009947 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L53) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `stal.DealIdBankJual` : `INT` | `#tempResult.DealIdChild` : `bigint` | `stal.DealIdBankJual` → `bigint` | `#tempResult.DealIdChild <- stal.DealIdBankJual` |
| IC-009957 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `stal.DealIdBankJual` : `INT` | `#tempResult.DealIdChild` : `bigint` | `stal.DealIdBankJual` → `bigint` | `#tempResult.DealIdChild <- stal.DealIdBankJual` |
| IC-005425 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateFOPTransferAsset:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFOPTransferAsset.sql#L127) | JOIN_ON / P11 / COMPARISON | `tr.DealIdBankJual` : `INT` | `li.DealId` : `BIGINT` | `tr.DealIdBankJual` → `BIGINT` | `li.DealId = tr.DealIdBankJual` |
| IC-005427 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateFOPTransferAsset:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateFOPTransferAsset.sql#L147) | JOIN_ON / P11 / COMPARISON | `tr.DealIdBankJual` : `INT` | `li.DealId` : `BIGINT` | `tr.DealIdBankJual` → `BIGINT` | `li.DealId = tr.DealIdBankJual` |
| IC-006446 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateTransactionBalance:264](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransactionBalance.sql#L264) | JOIN_ON / P11 / COMPARISON | `b.DealIdBankJual` : `INT` | `a.DealId` : `BIGINT` | `b.DealIdBankJual` → `BIGINT` | `a.DealId = b.DealIdBankJual` |
| IC-006854 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessFOPTransferAsset:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L68) | JOIN_ON / P11 / COMPARISON | `tr.DealIdBankJual` : `INT` | `tt.DealId` : `BIGINT` | `tr.DealIdBankJual` → `BIGINT` | `tt.DealId = tr.DealIdBankJual` |
| IC-006862 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessFOPTransferAsset:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L110) | JOIN_ON / P11 / COMPARISON | `tr.DealIdBankJual` : `INT` | `tt.DealId` : `BIGINT` | `tr.DealIdBankJual` → `BIGINT` | `tt.DealId = tr.DealIdBankJual` |
| IC-006883 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessFOPTransferAsset:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L158) | JOIN_ON / P11 / COMPARISON | `tr.DealIdBankJual` : `INT` | `rtt.FOPOriDealId` : `BIGINT` | `tr.DealIdBankJual` → `BIGINT` | `tr.DealIdBankJual = rtt.FOPOriDealId` |
| IC-006890 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessFOPTransferAsset:300](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L300) | JOIN_ON / P11 / COMPARISON | `tr.DealIdBankJual` : `INT` | `rtt.FOPOriDealId` : `BIGINT` | `tr.DealIdBankJual` → `BIGINT` | `tr.DealIdBankJual = rtt.FOPOriDealId` |

<a id="kolom-facevalue"></a>
## FaceValue

**Tipe dasar:** `DECIMAL(25,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAsetLink_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006849 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `(tr.FaceValue*-1)` : `DECIMAL (combined result; facets unknown)` | `retail.dbo.securitytransaction_tt.FaceValue` : `MONEY` | `(tr.FaceValue*-1)` → `MONEY` | `retail.dbo.securitytransaction_tt.FaceValue <- (tr.FaceValue*-1)` |
| IC-006858 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L85) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `(tr.FaceValue)` : `DECIMAL(25,2)` | `retail.dbo.securitytransaction_tt.FaceValue` : `MONEY` | `(tr.FaceValue)` → `MONEY` | `retail.dbo.securitytransaction_tt.FaceValue <- (tr.FaceValue)` |
| IC-006868 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `(tr.FaceValue)` : `DECIMAL(25,2)` | `retail.dbo.securitytransaction_tt.FaceValue` : `MONEY` | `(tr.FaceValue)` → `MONEY` | `retail.dbo.securitytransaction_tt.FaceValue <- (tr.FaceValue)` |
| IC-006853 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessFOPTransferAsset:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L57) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `tr.FaceValue` : `DECIMAL(25,2)` | `-1` : `int` | `-1` → `DECIMAL` | `tr.FaceValue*-1` |
| IC-013046 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_GetDataTransaksiAktifitas:296](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetDataTransaksiAktifitas.sql#L296) | CASE / P05 / BINARY_Multiply | `tc.FaceValue` : `DECIMAL(25,2)` | `-1` : `int` | `-1` → `DECIMAL` | `tc.FaceValue*-1` |
| IC-013047 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_GetDataTransaksiAktifitas:304](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetDataTransaksiAktifitas.sql#L304) | CASE / P05 / BINARY_Multiply | `tc.FaceValue` : `DECIMAL(25,2)` | `-1` : `int` | `-1` → `DECIMAL` | `tc.FaceValue*-1` |
| IC-013048 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_GetDataTransaksiAktifitas:311](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetDataTransaksiAktifitas.sql#L311) | CASE / P05 / BINARY_Multiply | `tc.FaceValue` : `DECIMAL(25,2)` | `-1` : `int` | `-1` → `DECIMAL` | `tc.FaceValue*-1` |
