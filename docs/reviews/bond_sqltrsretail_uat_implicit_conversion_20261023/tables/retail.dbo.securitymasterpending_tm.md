# Retail UAT.dbo.SecurityMasterPending_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 57 lokasi source, 4 procedure, 19 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 5 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-couponrate"></a>
## CouponRate

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007326 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSecurityMasterAuthorization:397](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L397) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smptm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(smptm.CouponRate, 0)` |
| IC-007330 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSProcessSecurityMasterAuthorization:426](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L426) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `smptm.CouponRate` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(smptm.CouponRate, 0)` |

<a id="kolom-flagcashback"></a>
## FlagCashBack

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L84). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006006 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L123) | CASE / P11 / COMPARISON | `smptm.FlagCashBack` : `BIT` | `1` : `int` | `smptm.FlagCashBack` → `INT` | `smptm.FlagCashBack = 1` |
| IC-013297 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L94) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagCashBack` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagCashBack,0)` |

<a id="kolom-flagechannel"></a>
## FlagEChannel

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L89). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013300 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:113](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L113) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagEChannel` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagEChannel, 0)` |

<a id="kolom-flagperdana"></a>
## FlagPerdana

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L71). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006002 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L66) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `smptm.FlagPerdana` : `BIT` | `1` : `int` | `smptm.FlagPerdana` → `INT` | `WHEN 1 THEN 'YES'` |
| IC-006003 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L67) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `smptm.FlagPerdana` : `BIT` | `0` : `int` | `smptm.FlagPerdana` → `INT` | `WHEN 0 THEN 'NO'` |

<a id="kolom-flagsingleaccount"></a>
## FlagSingleAccount

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L58). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006000 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L21) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `smptm.FlagSingleAccount` : `BIT` | `1` : `int` | `smptm.FlagSingleAccount` → `INT` | `WHEN 1 THEN 'YES'` |
| IC-006001 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L22) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `smptm.FlagSingleAccount` : `BIT` | `0` : `int` | `smptm.FlagSingleAccount` → `INT` | `WHEN 0 THEN 'NO'` |

<a id="kolom-flagsyariah"></a>
## FlagSyariah

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L88). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013299 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L112) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagSyariah` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagSyariah, 0)` |

<a id="kolom-flagtransaksibeli"></a>
## FlagTransaksiBeli

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L83). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006005 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L122) | CASE / P11 / COMPARISON | `smptm.FlagTransaksiBeli` : `BIT` | `1` : `int` | `smptm.FlagTransaksiBeli` → `INT` | `smptm.FlagTransaksiBeli = 1` |
| IC-013296 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L93) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagTransaksiBeli` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagTransaksiBeli,0)` |

<a id="kolom-flagtransaksibeliandcashback"></a>
## FlagTransaksiBeliAndCashBack

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L85). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006007 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L124) | CASE / P11 / COMPARISON | `smptm.FlagTransaksiBeliAndCashBack` : `BIT` | `1` : `int` | `smptm.FlagTransaksiBeliAndCashBack` → `INT` | `smptm.FlagTransaksiBeliAndCashBack = 1` |
| IC-013298 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L95) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `FlagTransaksiBeliAndCashBack` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(FlagTransaksiBeliAndCashBack,0)` |

<a id="kolom-flagtransaksibeliorcashback"></a>
## FlagTransaksiBeliOrCashBack

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L66). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006004 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L121) | CASE / P11 / COMPARISON | `smptm.FlagTransaksiBeliOrCashBack` : `BIT` | `1` : `int` | `smptm.FlagTransaksiBeliOrCashBack` → `INT` | `smptm.FlagTransaksiBeliOrCashBack = 1` |

<a id="kolom-imbalhasil"></a>
## ImbalHasil

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L102). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006010 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSecurityMasterAuthorization:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L192) | JOIN_ON / P01 / COMPARISON | `smptm.ImbalHasil` : `SMALLINT` | `pr.Value` : `VARCHAR(255)` | `pr.Value` → `SMALLINT` | `smptm.ImbalHasil = pr.Value` |

<a id="kolom-isusefortransaksi"></a>
## IsUseForTransaksi

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L87). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013295 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L75) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IsUseForTransaksi` : `BIT` | `0` : `int` | `0` → `BIT` | `isnull(IsUseForTransaksi,0)` |

<a id="kolom-jumlahunit"></a>
## JumlahUnit

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L54). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013293 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L72) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `JumlahUnit` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(JumlahUnit, 0)` |

<a id="kolom-lastprocesscode"></a>
## LastProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L53). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007286 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMaster:488](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L488) | JOIN_ON / P12 / COMPARISON | `smptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `smptm.LastProcessCode` → `VARCHAR` | `smptm.LastProcessCode = 'P'` |
| IC-007291 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:520](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L520) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterpending_tm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securitymasterpending_tm.LastProcessCode <- 'P'` |
| IC-007315 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:1130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1130) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterpending_tm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securitymasterpending_tm.LastProcessCode <- 'P'` |
| IC-007319 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:1690](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1690) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterpending_tm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.securitymasterpending_tm.LastProcessCode <- 'P'` |

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L50). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 24 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007325 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:393](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L393) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'U'` |
| IC-007329 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:422](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L422) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'U'` |
| IC-007332 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:659](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L659) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'I'` |
| IC-007339 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:684](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L684) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'I'` |
| IC-007342 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:721](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L721) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'I'` |
| IC-007343 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:765](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L765) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'I'` |
| IC-007345 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1008](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1008) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'I'` |
| IC-007347 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1041](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1041) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'I'` |
| IC-007348 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1060](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1060) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'I'` |
| IC-007350 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1087](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1087) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'I'` |
| IC-007351 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1106) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'I'` |
| IC-007354 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1163) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'I'` |
| IC-007355 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1278](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1278) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'U'` |
| IC-007356 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1286) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'U'` |
| IC-007357 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1297](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1297) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'U'` |
| IC-007358 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1344](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1344) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'U'` |
| IC-007360 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1584](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1584) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'U'` |
| IC-007361 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1590](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1590) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'U'` |
| IC-007367 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1897](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1897) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'D'` |
| IC-007368 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1905](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1905) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'D'` |
| IC-007369 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSecurityMasterAuthorization:1911](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L1911) | WHERE / P12 / COMPARISON | `smptm.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `smptm.OperationCode` → `VARCHAR` | `smptm.OperationCode = 'D'` |
| IC-007292 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:520](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L520) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterpending_tm.OperationCode` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.securitymasterpending_tm.OperationCode <- 'I'` |
| IC-007316 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:1130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1130) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterpending_tm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `'U'` → `CHAR(1)` | `retail.dbo.securitymasterpending_tm.OperationCode <- 'U'` |
| IC-007320 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityMaster:1690](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMaster.sql#L1690) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securitymasterpending_tm.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.securitymasterpending_tm.OperationCode <- 'D'` |

<a id="kolom-safekeepingfeetarif"></a>
## SafeKeepingFeeTarif

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013292 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListSecurityMasterPending_TM:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L67) | CASE / P11 / COMPARISON | `SafeKeepingFeeTarif` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `SafeKeepingFeeTarif = 0` |

<a id="kolom-secccy"></a>
## SecCcy

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_ARGUMENT.

**Pola dan arah review:**

- **P10 — Argumen function mengikuti signature:** Konversi mengikuti parameter function yang dideklarasikan. Cocokkan signature dan pemanggil; function per row tetap membutuhkan review terpisah, tidak selalu masalah index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007336 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSProcessSecurityMasterAuthorization:671](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L671) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `smptm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `smptm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('B',smptm.NextCouponDate,2,smptm.SecCcy)` |
| IC-007337 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSProcessSecurityMasterAuthorization:672](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L672) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `smptm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `smptm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('B',smptm.NextCouponDate,1,smptm.SecCcy)` |
| IC-007338 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSProcessSecurityMasterAuthorization:674](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L674) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `smptm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `smptm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('F', getdate(), tst.SettlementDate , smptm.SecCcy)` |
| IC-007340 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSProcessSecurityMasterAuthorization:702](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L702) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `smptm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `smptm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('F', getdate(), case when tr2.ChannelCode = 'CB' THEN psb.SettlementDateBankBuyCabang ELSE psb.SettlementDateBankBuyOM END,smptm.SecCcy)` |
| IC-007341 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_ARGUMENT | [dbo.TRSProcessSecurityMasterAuthorization:707](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityMasterAuthorization.sql#L707) | FUNCTION_CALL / P10 / UDF_ARGUMENT | `smptm.SecCcy` : `VARCHAR(3)` | `retail.dbo.fngetmxworkingdate.@pcCcy` : `char(3)` | `smptm.SecCcy` → `char(3)` | `dbo.fnGetMxWorkingDate('F', getdate(), case when tr2.ChannelCode = 'CB' THEN psb.SettlementDateBankSellCabang ELSE psb.SettlementDateBankSellOM END,smptm.SecCcy)` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L2). **Existing leading index:** SecurityMasterPending_TM_CN1&#124;SecurityMasterPending_TM_NU1.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006009 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSecurityMasterAuthorization:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L181) | WHERE / P01 / COMPARISON | `smptm.SecId` : `INT` | `g.process_key` : `VARCHAR(100)` | `g.process_key` → `INT` | `g.process_key = smptm.SecId` |
| IC-013302 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_ListSecurityMasterPending_TM:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L141) | WHERE / P01 / COMPARISON | `smp.SecId` : `INT` | `g.process_key` : `VARCHAR(100)` | `g.process_key` → `INT` | `g.process_key = smp.SecId` |

<a id="kolom-sectype"></a>
## SecType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005998 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L15) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `smptm.SecType` : `SMALLINT` | `1` : `int` | `smptm.SecType` → `INT` | `WHEN 1 THEN 'Obligasi'` |
| IC-005999 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSecurityMasterAuthorization:16](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSecurityMasterAuthorization.sql#L16) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `smptm.SecType` : `SMALLINT` | `2` : `int` | `smptm.SecType` → `INT` | `WHEN 2 THEN 'SBI'` |

<a id="kolom-taxwnawithoutnpwp"></a>
## TaxWNAWithoutNPWP

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityMasterPending_TM.sql#L55). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013294 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.trs_ListSecurityMasterPending_TM:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListSecurityMasterPending_TM.sql#L73) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `TaxWNAWithoutNPWP` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(TaxWNAWithoutNPWP,0)` |
