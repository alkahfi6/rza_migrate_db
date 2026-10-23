# Retail UAT.dbo.TRSParamBookBld_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 52 lokasi source, 12 procedure, 6 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bookbldflag"></a>
## BookbldFlag

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamBookBld_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000731 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMSubmitTransactionSekunder:347](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L347) | CASE / P11 / COMPARISON | `tpbb.BookbldFlag` : `BIT` | `1` : `int` | `tpbb.BookbldFlag` → `INT` | `tpbb.BookbldFlag = 1` |
| IC-000819 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.OMValidateBalance:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateBalance.sql#L248) | WHERE / P11 / COMPARISON | `isnull(bld.BookbldFlag,0)` : `BIT` | `0` : `int` | `isnull(bld.BookbldFlag,0)` → `INT` | `isnull(bld.BookbldFlag,0) = 0` |
| IC-000820 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.OMValidateBalance:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateBalance.sql#L248) | WHERE / P11 / COMPARISON | `isnull(bld.BookbldFlag,0)` : `BIT` | `1` : `int` | `isnull(bld.BookbldFlag,0)` → `INT` | `isnull(bld.BookbldFlag,0) = 1` |
| IC-006918 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessInsertSecurityTransactionCommit:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L312) | WHERE / P11 / COMPARISON | `isnull(bld.BookbldFlag,0)` : `BIT` | `0` : `int` | `isnull(bld.BookbldFlag,0)` → `INT` | `isnull(bld.BookbldFlag,0) = 0` |
| IC-006919 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessInsertSecurityTransactionCommit:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L312) | WHERE / P11 / COMPARISON | `isnull(bld.BookbldFlag,0)` : `BIT` | `1` : `int` | `isnull(bld.BookbldFlag,0)` → `INT` | `isnull(bld.BookbldFlag,0) = 1` |
| IC-006935 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessInsertSecurityTransactionCommit:621](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L621) | CASE / P11 / COMPARISON | `tpbb.BookbldFlag` : `BIT` | `1` : `int` | `tpbb.BookbldFlag` → `INT` | `tpbb.BookbldFlag = 1` |
| IC-013724 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_StatusSecurityTransaction_TT:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql#L81) | CASE / P11 / COMPARISON | `tpbb.BookbldFlag` : `BIT` | `1` : `int` | `tpbb.BookbldFlag` → `INT` | `tpbb.BookbldFlag = 1` |
| IC-013742 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_StatusSecurityTransaction_TT:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql#L235) | CASE / P11 / COMPARISON | `tpbb.BookbldFlag` : `BIT` | `1` : `int` | `tpbb.BookbldFlag` → `INT` | `tpbb.BookbldFlag = 1` |
| IC-000821 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMValidateBalance:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateBalance.sql#L248) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bld.BookbldFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(bld.BookbldFlag,0)` |
| IC-000822 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMValidateBalance:248](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateBalance.sql#L248) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bld.BookbldFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(bld.BookbldFlag,0)` |
| IC-004448 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamBookbld:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamBookbld.sql#L52) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `BookbldFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `BookbldFlag = 0` |
| IC-006920 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertSecurityTransactionCommit:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L312) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bld.BookbldFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(bld.BookbldFlag,0)` |
| IC-006921 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertSecurityTransactionCommit:312](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L312) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bld.BookbldFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(bld.BookbldFlag,0)` |

<a id="kolom-hargamodal"></a>
## HargaModal

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamBookBld_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004444 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamBookbld:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamBookbld.sql#L47) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `HargaModal` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `HargaModal = 0` |
| IC-008590 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSResponseBondFromMX:497](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L497) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pb.HargaModal` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(pb.HargaModal,0)` |

<a id="kolom-nominalbeli"></a>
## NominalBeli

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamBookBld_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004447 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamBookbld:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamBookbld.sql#L51) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `NominalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `NominalBeli = 0` |
| IC-012529 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L50) | DML_CONDITION / P05 / BINARY_Add | `isnull(NominalBeli,0)` : `DECIMAL(20,5)` | `@pmNominal` : `money` | `@pmNominal` → `DECIMAL` | `isnull(NominalBeli,0) + @pmNominal` |
| IC-012530 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L50) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(NominalBeli,0)` |
| IC-012533 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L65) | DML_CONDITION / P05 / BINARY_Add | `isnull(NominalBeli,0)` : `DECIMAL(20,5)` | `(@pmNominal - @pmNominalBefore)` : `MONEY (combined result; facets unknown)` | `(@pmNominal - @pmNominalBefore)` → `DECIMAL` | `isnull(NominalBeli,0) + (@pmNominal - @pmNominalBefore)` |
| IC-012534 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L65) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(NominalBeli,0)` |
| IC-012537 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L80) | DML_CONDITION / P05 / BINARY_Subtract | `isnull(NominalBeli,0)` : `DECIMAL(20,5)` | `@pmNominal` : `money` | `@pmNominal` → `DECIMAL` | `isnull(NominalBeli,0) - @pmNominal` |
| IC-012538 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L80) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalBeli` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(NominalBeli,0)` |

<a id="kolom-nominaljual"></a>
## NominalJual

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamBookBld_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004446 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamBookbld:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamBookbld.sql#L50) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `NominalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `NominalJual = 0` |
| IC-012531 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L56) | DML_CONDITION / P05 / BINARY_Add | `isnull(NominalJual,0)` : `DECIMAL(20,5)` | `@pmNominal` : `money` | `@pmNominal` → `DECIMAL` | `isnull(NominalJual,0) + @pmNominal` |
| IC-012532 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L56) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(NominalJual,0)` |
| IC-012535 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L71) | DML_CONDITION / P05 / BINARY_Add | `isnull(NominalJual,0)` : `DECIMAL(20,5)` | `(@pmNominal - @pmNominalBefore)` : `MONEY (combined result; facets unknown)` | `(@pmNominal - @pmNominalBefore)` → `DECIMAL` | `isnull(NominalJual,0) + (@pmNominal - @pmNominalBefore)` |
| IC-012536 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L71) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(NominalJual,0)` |
| IC-012539 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L86) | DML_CONDITION / P05 / BINARY_Subtract | `isnull(NominalJual,0)` : `DECIMAL(20,5)` | `@pmNominal` : `money` | `@pmNominal` → `DECIMAL` | `isnull(NominalJual,0) - @pmNominal` |
| IC-012540 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateParamHargaModalFromTransaksi2ndORI.sql#L86) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `NominalJual` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(NominalJual,0)` |

<a id="kolom-releaseapproval"></a>
## ReleaseApproval

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamBookBld_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005281 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateApproveReleaseBookbld:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApproveReleaseBookbld.sql#L11) | WHERE / P11 / COMPARISON | `isnull(ReleaseApproval,0)` : `BIT` | `1` : `int` | `isnull(ReleaseApproval,0)` → `INT` | `isnull(ReleaseApproval,0) = 1` |
| IC-011669 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamBookbld:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamBookbld.sql#L135) | WHERE / P11 / COMPARISON | `b.ReleaseApproval` : `BIT` | `1` : `int` | `b.ReleaseApproval` → `INT` | `b.ReleaseApproval = 1` |
| IC-011671 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamBookbld:142](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamBookbld.sql#L142) | WHERE / P11 / COMPARISON | `b.ReleaseApproval` : `BIT` | `1` : `int` | `b.ReleaseApproval` → `INT` | `b.ReleaseApproval = 1` |
| IC-011677 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSSaveParamBookbld:174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamBookbld.sql#L174) | CASE / P11 / COMPARISON | `isnull(pb.ReleaseApproval,0)` : `BIT` | `0` : `int` | `isnull(pb.ReleaseApproval,0)` → `INT` | `isnull(pb.ReleaseApproval,0) = 0` |
| IC-005282 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateApproveReleaseBookbld:11](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApproveReleaseBookbld.sql#L11) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ReleaseApproval` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(ReleaseApproval,0)` |
| IC-011676 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamBookbld:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamBookbld.sql#L172) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pb.ReleaseApproval` : `BIT` | `( case when isnull(pb.ReleaseApproval,0) = 0 and tp.ReleaseFlag = 1 then 1 else 0 END )` : `INT (combined result; facets unknown)` | `( case when isnull(pb.ReleaseApproval,0) = 0 and tp.ReleaseFlag = 1 then 1 else 0 END )` → `BIT` | `pb.ReleaseApproval = ( case when isnull(pb.ReleaseApproval,0) = 0 and tp.ReleaseFlag = 1 then 1 else 0 END )` |
| IC-011679 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamBookbld:174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamBookbld.sql#L174) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pb.ReleaseApproval` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(pb.ReleaseApproval,0)` |

<a id="kolom-releaseflag"></a>
## ReleaseFlag

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamBookBld_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 16 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000732 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMSubmitTransactionSekunder:351](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L351) | CASE / P11 / COMPARISON | `tpbb.ReleaseFlag` : `BIT` | `1` : `int` | `tpbb.ReleaseFlag` → `INT` | `tpbb.ReleaseFlag = 1` |
| IC-000823 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.OMValidateBalance:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateBalance.sql#L249) | WHERE / P11 / COMPARISON | `isnull(bld.ReleaseFlag,0)` : `BIT` | `1` : `int` | `isnull(bld.ReleaseFlag,0)` → `INT` | `isnull(bld.ReleaseFlag,0) = 1` |
| IC-004887 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSObligasiBlotterDaily:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiBlotterDaily.sql#L75) | WHERE / P11 / COMPARISON | `pb.ReleaseFlag` : `BIT` | `1` : `int` | `pb.ReleaseFlag` → `INT` | `pb.ReleaseFlag = 1` |
| IC-004903 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSObligasiBlotterDaily:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSObligasiBlotterDaily.sql#L113) | WHERE / P11 / COMPARISON | `pb.ReleaseFlag` : `BIT` | `1` : `int` | `pb.ReleaseFlag` → `INT` | `pb.ReleaseFlag = 1` |
| IC-006922 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSProcessInsertSecurityTransactionCommit:313](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L313) | WHERE / P11 / COMPARISON | `isnull(bld.ReleaseFlag,0)` : `BIT` | `1` : `int` | `isnull(bld.ReleaseFlag,0)` → `INT` | `isnull(bld.ReleaseFlag,0) = 1` |
| IC-006936 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessInsertSecurityTransactionCommit:625](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L625) | CASE / P11 / COMPARISON | `tpbb.ReleaseFlag` : `BIT` | `1` : `int` | `tpbb.ReleaseFlag` → `INT` | `tpbb.ReleaseFlag = 1` |
| IC-009827 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:826](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L826) | CASE / P11 / COMPARISON | `pbb.ReleaseFlag` : `BIT` | `1` : `int` | `pbb.ReleaseFlag` → `INT` | `pbb.ReleaseFlag = 1` |
| IC-009828 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:827](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L827) | CASE / P11 / COMPARISON | `pbb.ReleaseFlag` : `BIT` | `0` : `int` | `pbb.ReleaseFlag` → `INT` | `pbb.ReleaseFlag = 0` |
| IC-011600 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptTransactionDetailWMTA:579](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L579) | CASE / P11 / COMPARISON | `pbb.ReleaseFlag` : `BIT` | `1` : `int` | `pbb.ReleaseFlag` → `INT` | `pbb.ReleaseFlag = 1` |
| IC-011601 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptTransactionDetailWMTA:580](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L580) | CASE / P11 / COMPARISON | `pbb.ReleaseFlag` : `BIT` | `0` : `int` | `pbb.ReleaseFlag` → `INT` | `pbb.ReleaseFlag = 0` |
| IC-013725 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_StatusSecurityTransaction_TT:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql#L85) | CASE / P11 / COMPARISON | `tpbb.ReleaseFlag` : `BIT` | `1` : `int` | `tpbb.ReleaseFlag` → `INT` | `tpbb.ReleaseFlag = 1` |
| IC-013743 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_StatusSecurityTransaction_TT:239](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction_TT.sql#L239) | CASE / P11 / COMPARISON | `tpbb.ReleaseFlag` : `BIT` | `1` : `int` | `tpbb.ReleaseFlag` → `INT` | `tpbb.ReleaseFlag = 1` |
| IC-000824 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OMValidateBalance:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMValidateBalance.sql#L249) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bld.ReleaseFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(bld.ReleaseFlag,0)` |
| IC-004445 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamBookbld:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamBookbld.sql#L48) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ReleaseFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `ReleaseFlag = 0` |
| IC-006923 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessInsertSecurityTransactionCommit:313](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessInsertSecurityTransactionCommit.sql#L313) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bld.ReleaseFlag` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(bld.ReleaseFlag,0)` |
| IC-011673 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamBookbld:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamBookbld.sql#L166) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pb.ReleaseFlag` : `BIT` | `( case when tp.ReleaseFlag = 1 then 0 else tp.ReleaseFlag END )` : `INT (combined result; facets unknown)` | `( case when tp.ReleaseFlag = 1 then 0 else tp.ReleaseFlag END )` → `BIT` | `pb.ReleaseFlag = ( case when tp.ReleaseFlag = 1 then 0 else tp.ReleaseFlag END )` |
