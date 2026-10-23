# Retail UAT.dbo.TRSECParamSpread_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 62 lokasi source, 12 procedure, 11 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-bidspread"></a>
## BidSpread

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004473 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMaintenanceParamHargaModal:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L337) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sp.BidSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(sp.BidSpread, 0)` |
| IC-004475 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMaintenanceParamHargaModal:340](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L340) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sp.BidSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(sp.BidSpread, 0)` |
| IC-005883 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamECSpread:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamECSpread.sql#L20) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `BidSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(BidSpread,0)` |
| IC-005945 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSpread:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L69) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.BidSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.BidSpread,0)` |
| IC-012693 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUploadParamSpread:594](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUploadParamSpread.sql#L594) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.BidSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(ps.BidSpread,0)` |

<a id="kolom-bidspreadkaryawan"></a>
## BidSpreadKaryawan

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005885 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamECSpread:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamECSpread.sql#L22) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `BidSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(BidSpreadKaryawan,0)` |
| IC-005954 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSpread:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L85) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.BidSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.BidSpreadKaryawan,0)` |
| IC-005960 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSpread:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L122) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `BidSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(BidSpreadKaryawan,0)` |
| IC-005962 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSpreadExport:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpreadExport.sql#L25) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `BidSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(BidSpreadKaryawan,0)` |
| IC-012695 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUploadParamSpread:596](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUploadParamSpread.sql#L596) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.BidSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(ps.BidSpreadKaryawan,0)` |

<a id="kolom-buysell"></a>
## BuySell

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 21 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000663 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:272](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L272) | WHERE / P12 / COMPARISON | `ps.BuySell` : `CHAR(1)` | `@cBankBuySell` : `varchar(10)` | `ps.BuySell` → `VARCHAR` | `ps.BuySell = @cBankBuySell` |
| IC-000665 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L286) | WHERE / P12 / COMPARISON | `ps.BuySell` : `CHAR(1)` | `@cBankBuySell` : `varchar(10)` | `ps.BuySell` → `VARCHAR` | `ps.BuySell = @cBankBuySell` |
| IC-000666 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:306](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L306) | WHERE / P12 / COMPARISON | `ps.BuySell` : `CHAR(1)` | `@cBankBuySell` : `varchar(10)` | `ps.BuySell` → `VARCHAR` | `ps.BuySell = @cBankBuySell` |
| IC-000667 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:316](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L316) | WHERE / P12 / COMPARISON | `ps.BuySell` : `CHAR(1)` | `@cBankBuySell` : `varchar(10)` | `ps.BuySell` → `VARCHAR` | `ps.BuySell = @cBankBuySell` |
| IC-000668 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:336](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L336) | WHERE / P12 / COMPARISON | `ps.BuySell` : `CHAR(1)` | `@cBankBuySell` : `varchar(10)` | `ps.BuySell` → `VARCHAR` | `ps.BuySell = @cBankBuySell` |
| IC-000669 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:346](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L346) | WHERE / P12 / COMPARISON | `ps.BuySell` : `CHAR(1)` | `@cBankBuySell` : `varchar(10)` | `ps.BuySell` → `VARCHAR` | `ps.BuySell = @cBankBuySell` |
| IC-000683 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:458](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L458) | WHERE / P12 / COMPARISON | `ps.BuySell` : `CHAR(1)` | `@cBankBuySell` : `varchar(10)` | `ps.BuySell` → `VARCHAR` | `ps.BuySell = @cBankBuySell` |
| IC-000686 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:468](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L468) | WHERE / P12 / COMPARISON | `ps.BuySell` : `CHAR(1)` | `@cBankBuySell` : `varchar(10)` | `ps.BuySell` → `VARCHAR` | `ps.BuySell = @cBankBuySell` |
| IC-003604 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCmprSecurityNo:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNo.sql#L70) | WHERE / P12 / COMPARISON | `b.BuySell` : `CHAR(1)` | `'B'` : `varchar` | `b.BuySell` → `VARCHAR` | `b.BuySell = 'B'` |
| IC-003605 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCmprSecurityNo:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNo.sql#L83) | WHERE / P12 / COMPARISON | `b.BuySell` : `CHAR(1)` | `'S'` : `varchar` | `b.BuySell` → `VARCHAR` | `b.BuySell = 'S'` |
| IC-003621 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCmprSecurityNoChannel:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNoChannel.sql#L63) | WHERE / P12 / COMPARISON | `b.BuySell` : `CHAR(1)` | `'B'` : `varchar` | `b.BuySell` → `VARCHAR` | `b.BuySell = 'B'` |
| IC-003622 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCmprSecurityNoChannel:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCmprSecurityNoChannel.sql#L75) | WHERE / P12 / COMPARISON | `b.BuySell` : `CHAR(1)` | `'S'` : `varchar` | `b.BuySell` → `VARCHAR` | `b.BuySell = 'S'` |
| IC-004476 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:348](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L348) | JOIN_ON / P12 / COMPARISON | `sp.BuySell` : `CHAR(1)` | `'B'` : `varchar` | `sp.BuySell` → `VARCHAR` | `sp.BuySell = 'B'` |
| IC-004482 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMaintenanceParamHargaModal:370](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L370) | JOIN_ON / P12 / COMPARISON | `sp.BuySell` : `CHAR(1)` | `'S'` : `varchar` | `sp.BuySell` → `VARCHAR` | `sp.BuySell = 'S'` |
| IC-005942 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParamSpread:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L66) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `BuySell` : `CHAR(1)` | `'B'` : `varchar` | `BuySell` → `VARCHAR` | `WHEN 'B' THEN dbo.YIELD( trs.SettlementDateBankBuy, sec.MaturityDate, sec.CouponRate, sec.JumlahUnit, CASE WHEN hm.HargaModalBeli = 0 THEN ISNULL(tmpB.HargaModalBeli,0) + bp.BidSpread ELSE hm.HargaModalBeli + ISNULL(bp.BidSpread,0) END )` |
| IC-005946 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParamSpread:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L72) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `BuySell` : `CHAR(1)` | `'S'` : `varchar` | `BuySell` → `VARCHAR` | `WHEN 'S' THEN dbo.YIELD( trs.SettlementDateBankSell, sec.MaturityDate, sec.CouponRate, sec.JumlahUnit, CASE WHEN hm.HargaModalJual != 0 THEN hm.HargaModalJual + ISNULL(bp.OfferSpread,0) ELSE ISNULL(tmpJ.HargaModalJual,0) + ISNULL(bp.OfferSpread,0) END )` |
| IC-005951 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParamSpread:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L82) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `BuySell` : `CHAR(1)` | `'B'` : `varchar` | `BuySell` → `VARCHAR` | `WHEN 'B' THEN dbo.YIELD( trs.SettlementDateBankBuy, sec.MaturityDate, sec.CouponRate, sec.JumlahUnit, CASE WHEN hm.HargaModalBeli = 0 THEN ISNULL(tmpB.HargaModalBeli,0) + bp.BidSpreadKaryawan ELSE hm.HargaModalBeli + ISNULL(bp.BidSpreadKaryawan,0) END )` |
| IC-005955 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateParamSpread:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L88) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `BuySell` : `CHAR(1)` | `'S'` : `varchar` | `BuySell` → `VARCHAR` | `WHEN 'S' THEN dbo.YIELD( trs.SettlementDateBankSell, sec.MaturityDate, sec.CouponRate, sec.JumlahUnit, CASE WHEN hm.HargaModalJual != 0 THEN hm.HargaModalJual + ISNULL(bp.OfferSpreadKaryawan,0) ELSE ISNULL(tmpJ.HargaModalJual,0) + ISNULL(bp.OfferSpreadKaryawan,0) END )` |
| IC-007962 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRPTRevenueOBLforWM:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L163) | JOIN_ON / P12 / COMPARISON | `sr.BuySell` : `CHAR(1)` | `CASE obl.KodeJnsTrx WHEN 3 THEN 'B' ELSE 'S' END` : `VARCHAR (combined result; facets unknown)` | `sr.BuySell` → `VARCHAR` | `sr.BuySell = CASE obl.KodeJnsTrx WHEN 3 THEN 'B' ELSE 'S' END` |
| IC-008747 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee:288](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L288) | WHERE / P12 / COMPARISON | `sr.BuySell` : `CHAR(1)` | `CASE @nCalcType WHEN 3 THEN 'B' ELSE 'S' END` : `VARCHAR (combined result; facets unknown)` | `sr.BuySell` → `VARCHAR` | `sr.BuySell = CASE @nCalcType WHEN 3 THEN 'B' ELSE 'S' END` |
| IC-008751 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateFee:310](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L310) | WHERE / P12 / COMPARISON | `sr.BuySell` : `CHAR(1)` | `CASE @nCalcType WHEN 3 THEN 'B' ELSE 'S' END` : `VARCHAR (combined result; facets unknown)` | `sr.BuySell` → `VARCHAR` | `sr.BuySell = CASE @nCalcType WHEN 3 THEN 'B' ELSE 'S' END` |

<a id="kolom-channelcode"></a>
## ChannelCode

**Tipe dasar:** `VARCHAR(15)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008748 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:293](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L293) | WHERE / P12 / COMPARISON | `sr.ChannelCode` : `VARCHAR(15)` | `@cInstruction` : `char(15)` | `@cInstruction` → `VARCHAR` | `sr.ChannelCode = @cInstruction` |
| IC-008752 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:315](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L315) | WHERE / P12 / COMPARISON | `sr.ChannelCode` : `VARCHAR(15)` | `@cInstruction` : `char(15)` | `@cInstruction` → `VARCHAR` | `sr.ChannelCode = @cInstruction` |

<a id="kolom-indikasiharga"></a>
## IndikasiHarga

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000664 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L274) | WHERE / P11 / COMPARISON | `IndikasiHarga` : `BIT` | `1` : `int` | `IndikasiHarga` → `INT` | `IndikasiHarga = 1` |
| IC-005887 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamECSpread:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamECSpread.sql#L26) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(IndikasiHarga,0)` |
| IC-012017 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamSpread:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamSpread.sql#L298) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(ps.IndikasiHarga, 0)` |
| IC-012019 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamSpread:324](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamSpread.sql#L324) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(ps.IndikasiHarga, 0)` |
| IC-012020 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamSpread:333](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamSpread.sql#L333) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(ps.IndikasiHarga, 0)` |
| IC-012024 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamSpread:363](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamSpread.sql#L363) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(ps.IndikasiHarga, 0)` |
| IC-012049 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamUploadSpread:255](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamUploadSpread.sql#L255) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(ps.IndikasiHarga, 0)` |
| IC-012053 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSaveParamUploadSpread:307](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamUploadSpread.sql#L307) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(ps.IndikasiHarga, 0)` |
| IC-012691 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUploadParamSpread:566](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUploadParamSpread.sql#L566) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.IndikasiHarga` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(ps.IndikasiHarga, 0)` |

<a id="kolom-nominalfrom"></a>
## NominalFrom

**Tipe dasar:** `DECIMAL(21,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007960 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforWM:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L159) | JOIN_ON / P11 / COMPARISON | `sr.NominalFrom` : `DECIMAL(21,2)` | `obl.FaceValue` : `MONEY` | `obl.FaceValue` → `DECIMAL` | `sr.NominalFrom <= obl.FaceValue` |
| IC-008745 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:286](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L286) | WHERE / P11 / COMPARISON | `sr.NominalFrom` : `DECIMAL(21,2)` | `@nFaceValue` : `money` | `@nFaceValue` → `DECIMAL` | `sr.NominalFrom <= @nFaceValue` |
| IC-008749 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:308](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L308) | WHERE / P11 / COMPARISON | `sr.NominalFrom` : `DECIMAL(21,2)` | `@nFaceValue` : `money` | `@nFaceValue` → `DECIMAL` | `sr.NominalFrom <= @nFaceValue` |

<a id="kolom-nominalto"></a>
## NominalTo

**Tipe dasar:** `DECIMAL(21,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007961 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRPTRevenueOBLforWM:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTRevenueOBLforWM.sql#L161) | JOIN_ON / P11 / COMPARISON | `sr.NominalTo` : `DECIMAL(21,2)` | `obl.FaceValue` : `MONEY` | `obl.FaceValue` → `DECIMAL` | `sr.NominalTo >= obl.FaceValue` |
| IC-008746 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:287](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L287) | WHERE / P11 / COMPARISON | `sr.NominalTo` : `DECIMAL(21,2)` | `@nFaceValue` : `money` | `@nFaceValue` → `DECIMAL` | `sr.NominalTo >= @nFaceValue` |
| IC-008750 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:309](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L309) | WHERE / P11 / COMPARISON | `sr.NominalTo` : `DECIMAL(21,2)` | `@nFaceValue` : `money` | `@nFaceValue` → `DECIMAL` | `sr.NominalTo >= @nFaceValue` |

<a id="kolom-offerspread"></a>
## OfferSpread

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004479 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMaintenanceParamHargaModal:359](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L359) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sp.OfferSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(sp.OfferSpread, 0)` |
| IC-004481 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSMaintenanceParamHargaModal:362](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L362) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sp.OfferSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `isnull(sp.OfferSpread, 0)` |
| IC-005884 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamECSpread:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamECSpread.sql#L21) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `OfferSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(OfferSpread,0)` |
| IC-005948 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSpread:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L74) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.OfferSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.OfferSpread,0)` |
| IC-005949 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSpread:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L75) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.OfferSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.OfferSpread,0)` |
| IC-012694 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUploadParamSpread:595](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUploadParamSpread.sql#L595) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.OfferSpread` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(ps.OfferSpread,0)` |

<a id="kolom-offerspreadkaryawan"></a>
## OfferSpreadKaryawan

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005886 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamECSpread:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamECSpread.sql#L23) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `OfferSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(OfferSpreadKaryawan,0)` |
| IC-005957 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSpread:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L90) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.OfferSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.OfferSpreadKaryawan,0)` |
| IC-005958 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSpread:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L91) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bp.OfferSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(bp.OfferSpreadKaryawan,0)` |
| IC-005961 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSpread:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpread.sql#L123) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `OfferSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(OfferSpreadKaryawan,0)` |
| IC-005963 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateParamSpreadExport:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParamSpreadExport.sql#L26) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `OfferSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(OfferSpreadKaryawan,0)` |
| IC-012696 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUploadParamSpread:597](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUploadParamSpread.sql#L597) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ps.OfferSpreadKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(ps.OfferSpreadKaryawan,0)` |

<a id="kolom-ytm"></a>
## YTM

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000684 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OMGetListProductSekunder:459](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L459) | WHERE / P11 / COMPARISON | `ps.YTM` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `ps.YTM > 0` |

<a id="kolom-ytmkaryawan"></a>
## YTMKaryawan

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSECParamSpread_TR.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000687 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OMGetListProductSekunder:469](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L469) | WHERE / P11 / COMPARISON | `ps.YTMKaryawan` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL` | `ps.YTMKaryawan > 0` |
