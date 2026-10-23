# Retail UAT.dbo.maturity_upload_temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 30 lokasi source, 2 procedure, 15 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-acquisitionprice"></a>
## AcquisitionPrice

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004635 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L286) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.AcquisitionPrice` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.AcquisitionPrice, 0)` |
| IC-004677 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L237) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.AcquisitionPrice` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.AcquisitionPrice, 0)` |

<a id="kolom-bonus"></a>
## Bonus

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004633 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:284](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L284) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.Bonus` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.Bonus, 0)` |
| IC-004675 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L235) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.Bonus` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.Bonus, 0)` |

<a id="kolom-capitalgain"></a>
## CapitalGain

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004636 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:287](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L287) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.CapitalGain` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.CapitalGain, 0)` |
| IC-004678 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L238) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.CapitalGain` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.CapitalGain, 0)` |

<a id="kolom-couponaftertax"></a>
## CouponAfterTax

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004630 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:275](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L275) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.CouponAfterTax` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.CouponAfterTax, 0)` |
| IC-004672 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L226) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.CouponAfterTax` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.CouponAfterTax, 0)` |

<a id="kolom-couponamount"></a>
## CouponAmount

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004638 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:289](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L289) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.CouponAmount` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.CouponAmount, 0)` |
| IC-004680 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:240](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L240) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.CouponAmount` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.CouponAmount, 0)` |

<a id="kolom-netprincipleamount"></a>
## NetPrincipleAmount

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004643 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:294](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L294) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.NetPrincipleAmount` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.NetPrincipleAmount, 0)` |
| IC-004685 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L245) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.NetPrincipleAmount` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.NetPrincipleAmount, 0)` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004632 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L282) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.Nominal` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.Nominal, 0)` |
| IC-004674 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:233](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L233) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.Nominal` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.Nominal, 0)` |

<a id="kolom-rounding"></a>
## Rounding

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004644 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L295) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.Rounding` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.Rounding, 0)` |
| IC-004686 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:246](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L246) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.Rounding` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.Rounding, 0)` |

<a id="kolom-safekeeping"></a>
## SafeKeeping

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004640 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:291](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L291) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.SafeKeeping` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.SafeKeeping, 0)` |
| IC-004682 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:242](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L242) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.SafeKeeping` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.SafeKeeping, 0)` |

<a id="kolom-safekeepingactual"></a>
## SafeKeepingActual

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004641 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:292](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L292) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.SafeKeepingActual` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.SafeKeepingActual, 0)` |
| IC-004683 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:243](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L243) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.SafeKeepingActual` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.SafeKeepingActual, 0)` |

<a id="kolom-safekeepingaftertax"></a>
## SafeKeepingAfterTax

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004631 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:276](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L276) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.SafeKeepingAfterTax` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.SafeKeepingAfterTax, 0)` |
| IC-004673 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:227](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L227) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.SafeKeepingAfterTax` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.SafeKeepingAfterTax, 0)` |

<a id="kolom-safekeepingppn"></a>
## SafeKeepingPPN

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004642 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:293](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L293) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.SafeKeepingPPN` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.SafeKeepingPPN, 0)` |
| IC-004684 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:244](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L244) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.SafeKeepingPPN` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.SafeKeepingPPN, 0)` |

<a id="kolom-soldprice"></a>
## SoldPrice

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004634 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:285](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L285) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.SoldPrice` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.SoldPrice, 0)` |
| IC-004676 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:236](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L236) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.SoldPrice` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.SoldPrice, 0)` |

<a id="kolom-taxoncapitalgain"></a>
## TaxOnCapitalGain

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004637 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:288](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L288) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.TaxOnCapitalGain` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.TaxOnCapitalGain, 0)` |
| IC-004679 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:239](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L239) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.TaxOnCapitalGain` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.TaxOnCapitalGain, 0)` |

<a id="kolom-taxoncoupon"></a>
## TaxOnCoupon

**Tipe dasar:** `[decimal](20, 5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.maturity_upload_temp.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004639 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity:290](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity.sql#L290) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.TaxOnCoupon` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.TaxOnCoupon, 0)` |
| IC-004681 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadMaturity_BAK:241](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadMaturity_BAK.sql#L241) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mut.TaxOnCoupon` : `[decimal](20, 5)` | `0` : `int` | `0` → `[decimal](20, 5)` | `ISNULL(mut.TaxOnCoupon, 0)` |
