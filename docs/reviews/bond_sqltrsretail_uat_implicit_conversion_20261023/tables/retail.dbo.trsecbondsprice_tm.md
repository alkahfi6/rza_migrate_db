# Retail UAT.dbo.TRSECBondsPrice_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 24 lokasi source, 4 procedure, 10 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bidprice"></a>
## Bidprice

**Tipe dasar:** `DECIMAL(20,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TM.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000267 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.EBWInvestmentInquiry#ORI:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EBWInvestmentInquiry%23ORI.sql#L47) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `isnull (tbp.Bidprice, 0.0)` : `DECIMAL(20,3)` | `tcs.Outstanding` : `MONEY (aggregate; facets unknown)` | `tcs.Outstanding` → `DECIMAL` | `tcs.Outstanding * isnull (tbp.Bidprice, 0.0)` |
| IC-000270 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.EBWInvestmentInquiry#ORI:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EBWInvestmentInquiry%23ORI.sql#L54) | WHERE / P11 / COMPARISON | `NAV` : `DECIMAL(20,3)` | `0` : `int` | `0` → `DECIMAL` | `NAV >= 0` |
| IC-000272 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.EBWInvestmentInquiry#ORI:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EBWInvestmentInquiry%23ORI.sql#L54) | WHERE / P11 / COMPARISON | `UnitNominal` : `DECIMAL (combined result; facets unknown)` | `0` : `int` | `0` → `DECIMAL` | `UnitNominal >= 0` |
| IC-000273 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.EBWInvestmentInquiry#ORI:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EBWInvestmentInquiry%23ORI.sql#L63) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `isnull (tbp.Bidprice, 0.0)` : `DECIMAL(20,3)` | `tcs.Outstanding` : `MONEY (aggregate; facets unknown)` | `tcs.Outstanding` → `DECIMAL` | `tcs.Outstanding * isnull (tbp.Bidprice, 0.0)` |
| IC-000276 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.EBWInvestmentInquiry#ORI:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EBWInvestmentInquiry%23ORI.sql#L70) | WHERE / P11 / COMPARISON | `NAV` : `DECIMAL(20,3)` | `0` : `int` | `0` → `DECIMAL` | `NAV >= 0` |
| IC-000278 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.EBWInvestmentInquiry#ORI:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EBWInvestmentInquiry%23ORI.sql#L70) | WHERE / P11 / COMPARISON | `UnitNominal` : `DECIMAL (combined result; facets unknown)` | `0` : `int` | `0` → `DECIMAL` | `UnitNominal >= 0` |

<a id="kolom-bidytm"></a>
## Bidytm

**Tipe dasar:** `SMALLMONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003412 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L80) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_tm.Bidytm` : `SMALLMONEY` | `@Bidytm` : `money` | `@Bidytm` → `SMALLMONEY` | `retail.dbo.trsecbondsprice_tm.Bidytm <- @Bidytm` |
| IC-003420 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_tm.Bidytm` : `SMALLMONEY` | `@Bidytm` : `money` | `@Bidytm` → `SMALLMONEY` | `retail.dbo.trsecbondsprice_tm.Bidytm <- @Bidytm` |

<a id="kolom-lastupdatedate"></a>
## LastUpdateDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TM.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000268 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.EBWInvestmentInquiry#ORI:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EBWInvestmentInquiry%23ORI.sql#L48) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `isnull(tbp.LastUpdateDate, tbp.InsertedDate)` : `DATETIME` | `'1900-12-31'` : `varchar` | `'1900-12-31'` → `DATETIME` | `isnull(isnull(tbp.LastUpdateDate, tbp.InsertedDate),'1900-12-31')` |
| IC-000274 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.EBWInvestmentInquiry#ORI:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.EBWInvestmentInquiry%23ORI.sql#L64) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `isnull(tbp.LastUpdateDate, tbp.InsertedDate)` : `DATETIME` | `'1900-12-31'` : `varchar` | `'1900-12-31'` → `DATETIME` | `isnull(isnull(tbp.LastUpdateDate, tbp.InsertedDate),'1900-12-31')` |

<a id="kolom-maturitydate"></a>
## MaturityDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TM.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003413 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L80) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_tm.MaturityDate` : `DATETIME` | `@MaturityDate` : `varchar(8)` | `@MaturityDate` → `DATETIME` | `retail.dbo.trsecbondsprice_tm.MaturityDate <- @MaturityDate` |
| IC-003421 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_tm.MaturityDate` : `DATETIME` | `@MaturityDate` : `varchar(8)` | `@MaturityDate` → `DATETIME` | `retail.dbo.trsecbondsprice_tm.MaturityDate <- @MaturityDate` |

<a id="kolom-offerytm"></a>
## Offerytm

**Tipe dasar:** `SMALLMONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003414 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L80) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_tm.Offerytm` : `SMALLMONEY` | `@Offerytm` : `money` | `@Offerytm` → `SMALLMONEY` | `retail.dbo.trsecbondsprice_tm.Offerytm <- @Offerytm` |
| IC-003422 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_tm.Offerytm` : `SMALLMONEY` | `@Offerytm` : `money` | `@Offerytm` → `SMALLMONEY` | `retail.dbo.trsecbondsprice_tm.Offerytm <- @Offerytm` |

<a id="kolom-performance1d"></a>
## Performance1D

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TM.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003606 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCmprSecurityNo:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNo.sql#L94) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.Performance1D` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.Performance1D,0)` |
| IC-003623 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCmprSecurityNoChannel:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNoChannel.sql#L86) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.Performance1D` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.Performance1D,0)` |

<a id="kolom-performance1m"></a>
## Performance1M

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TM.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003608 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCmprSecurityNo:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNo.sql#L96) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.Performance1M` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.Performance1M,0)` |
| IC-003625 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCmprSecurityNoChannel:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNoChannel.sql#L88) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.Performance1M` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.Performance1M,0)` |

<a id="kolom-performance1w"></a>
## Performance1W

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TM.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003607 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCmprSecurityNo:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNo.sql#L95) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.Performance1W` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.Performance1W,0)` |
| IC-003624 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCmprSecurityNoChannel:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNoChannel.sql#L87) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.Performance1W` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.Performance1W,0)` |

<a id="kolom-performance1y"></a>
## Performance1Y

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TM.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003609 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCmprSecurityNo:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNo.sql#L97) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.Performance1Y` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.Performance1Y,0)` |
| IC-003626 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSCmprSecurityNoChannel:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNoChannel.sql#L89) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.Performance1Y` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.Performance1Y,0)` |

<a id="kolom-pricestatus"></a>
## PriceStatus

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECBondsPrice_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003415 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L80) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_tm.PriceStatus` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.trsecbondsprice_tm.PriceStatus <- 0` |
| IC-003423 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSBondsPriceInsertEChannel:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSBondsPriceInsertEChannel.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsecbondsprice_tm.PriceStatus` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `retail.dbo.trsecbondsprice_tm.PriceStatus <- 1` |
