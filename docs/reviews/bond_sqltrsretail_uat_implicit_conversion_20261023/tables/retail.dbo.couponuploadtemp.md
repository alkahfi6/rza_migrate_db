# Retail UAT.dbo.CouponUploadTemp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 22 lokasi source, 2 procedure, 11 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bonus"></a>
## Bonus

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004531 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:211](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L211) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.Bonus` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.Bonus, 0)` |
| IC-004563 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L199) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.Bonus` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.Bonus, 0)` |

<a id="kolom-couponaftertax"></a>
## CouponAfterTax

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004539 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:219](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L219) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.CouponAfterTax` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.CouponAfterTax, 0)` |
| IC-004571 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:207](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L207) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.CouponAfterTax` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.CouponAfterTax, 0)` |

<a id="kolom-grossamount"></a>
## GrossAmount

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004532 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:212](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L212) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.GrossAmount` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.GrossAmount, 0)` |
| IC-004564 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:200](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L200) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.GrossAmount` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.GrossAmount, 0)` |

<a id="kolom-netamount"></a>
## NetAmount

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004537 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L217) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.NetAmount` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.NetAmount, 0)` |
| IC-004569 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L205) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.NetAmount` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.NetAmount, 0)` |

<a id="kolom-netamountround"></a>
## NetAmountRound

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004538 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:218](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L218) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.NetAmountRound` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.NetAmountRound, 0)` |
| IC-004570 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:206](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L206) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.NetAmountRound` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.NetAmountRound, 0)` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004530 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:210](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L210) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.Nominal` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.Nominal, 0)` |
| IC-004562 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L198) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.Nominal` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.Nominal, 0)` |

<a id="kolom-sfppn"></a>
## SFPPN

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004536 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L216) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.SFPPN` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.SFPPN, 0)` |
| IC-004568 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:204](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L204) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.SFPPN` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.SFPPN, 0)` |

<a id="kolom-safekeeping"></a>
## SafeKeeping

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004534 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:214](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L214) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.SafeKeeping` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.SafeKeeping, 0)` |
| IC-004566 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:202](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L202) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.SafeKeeping` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.SafeKeeping, 0)` |

<a id="kolom-safekeepingact"></a>
## SafeKeepingAct

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004535 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L215) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.SafeKeepingAct` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.SafeKeepingAct, 0)` |
| IC-004567 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:203](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L203) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.SafeKeepingAct` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.SafeKeepingAct, 0)` |

<a id="kolom-safekeepingaftertax"></a>
## SafeKeepingAfterTax

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004540 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:220](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L220) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.SafeKeepingAfterTax` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.SafeKeepingAfterTax, 0)` |
| IC-004572 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L208) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.SafeKeepingAfterTax` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.SafeKeepingAfterTax, 0)` |

<a id="kolom-taxonint"></a>
## TaxOnInt

**Tipe dasar:** `[money]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CouponUploadTemp.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004533 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon.sql#L213) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.TaxOnInt` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.TaxOnInt, 0)` |
| IC-004565 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMergeFileUploadKupon_BAK:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMergeFileUploadKupon_BAK.sql#L201) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `cut.TaxOnInt` : `[money]` | `0` : `int` | `0` → `[money]` | `ISNULL(cut.TaxOnInt, 0)` |
