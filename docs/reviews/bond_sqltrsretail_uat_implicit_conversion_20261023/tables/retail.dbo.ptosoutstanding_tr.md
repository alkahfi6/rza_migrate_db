# Retail UAT.dbo.PTOSOutstanding_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 156 lokasi source, 16 procedure, 14 kolom. **A:** 21 lokasi pada 2 procedure. **B:** 10 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-kodeproduk"></a>
## KodeProduk

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 2 SP. **A:** 10 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001236 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L23) | JOIN_ON / P02 / COMPARISON | `potr.KodeProduk` : `VARCHAR(20)` | `tmp.[Kode Produk]` : `NVARCHAR(255)` | `potr.KodeProduk` → `NVARCHAR` | `tmp.[Kode Produk] = potr.KodeProduk` |
| IC-001282 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L99) | JOIN_ON / P02 / COMPARISON | `potr.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR(255)` | `potr.KodeProduk` → `NVARCHAR` | `potr.KodeProduk = tmp.KodeProduk` |
| IC-001304 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L131) | JOIN_ON / P02 / COMPARISON | `potr.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR(255)` | `potr.KodeProduk` → `NVARCHAR` | `potr.KodeProduk = tmp.KodeProduk` |
| IC-001322 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L157) | JOIN_ON / P02 / COMPARISON | `tr.KodeProduk` : `VARCHAR(20)` | `tmp.[Kode Produk]` : `NVARCHAR(255)` | `tr.KodeProduk` → `NVARCHAR` | `tr.KodeProduk = tmp.[Kode Produk]` |
| IC-001325 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L168) | JOIN_ON / P02 / COMPARISON | `tr.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR(255)` | `tr.KodeProduk` → `NVARCHAR` | `tr.KodeProduk = tmp.KodeProduk` |
| IC-001330 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L25) | JOIN_ON / P02 / COMPARISON | `potr.KodeProduk` : `VARCHAR(20)` | `tmp.[Kode Produk]` : `NVARCHAR(255)` | `potr.KodeProduk` → `NVARCHAR` | `tmp.[Kode Produk] = potr.KodeProduk` |
| IC-001387 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L136) | JOIN_ON / P02 / COMPARISON | `potr.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR(255)` | `potr.KodeProduk` → `NVARCHAR` | `potr.KodeProduk = tmp.KodeProduk` |
| IC-001409 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L171) | JOIN_ON / P02 / COMPARISON | `potr.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR(255)` | `potr.KodeProduk` → `NVARCHAR` | `potr.KodeProduk = tmp.KodeProduk` |
| IC-001428 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L197) | JOIN_ON / P02 / COMPARISON | `tr.KodeProduk` : `VARCHAR(20)` | `tmp.[Kode Produk]` : `NVARCHAR(255)` | `tr.KodeProduk` → `NVARCHAR` | `tr.KodeProduk = tmp.[Kode Produk]` |
| IC-001431 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:207](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L207) | JOIN_ON / P02 / COMPARISON | `tr.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR(255)` | `tr.KodeProduk` → `NVARCHAR` | `tr.KodeProduk = tmp.KodeProduk` |
| IC-001289 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptosoutstanding_tr.KodeProduk <- KodeProduk` |
| IC-001394 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptosoutstanding_tr.KodeProduk <- KodeProduk` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L3). **Existing leading index:** ix_ptosoutstanding_tr_sid&#124;ix_ptosoutstanding_tr_sidnordn.

**Frekuensi source:** 12 lokasi / 2 SP. **A:** 10 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001235 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L21) | JOIN_ON / P02 / COMPARISON | `potr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `potr.[SID]` → `NVARCHAR` | `tmp.[SID] = potr.[SID]` |
| IC-001281 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L97) | JOIN_ON / P02 / COMPARISON | `potr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `potr.[SID]` → `NVARCHAR` | `potr.[SID] = tmp.[SID]` |
| IC-001303 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L129) | JOIN_ON / P02 / COMPARISON | `potr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `potr.[SID]` → `NVARCHAR` | `potr.[SID] = tmp.[SID]` |
| IC-001321 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L155) | JOIN_ON / P02 / COMPARISON | `tr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tr.[SID]` → `NVARCHAR` | `tr.[SID] = tmp.[SID]` |
| IC-001324 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L166) | JOIN_ON / P02 / COMPARISON | `tr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tr.[SID]` → `NVARCHAR` | `tr.[SID] = tmp.[SID]` |
| IC-001329 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L23) | JOIN_ON / P02 / COMPARISON | `potr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `potr.[SID]` → `NVARCHAR` | `tmp.[SID] = potr.[SID]` |
| IC-001386 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L134) | JOIN_ON / P02 / COMPARISON | `potr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `potr.[SID]` → `NVARCHAR` | `potr.[SID] = tmp.[SID]` |
| IC-001408 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L169) | JOIN_ON / P02 / COMPARISON | `potr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `potr.[SID]` → `NVARCHAR` | `potr.[SID] = tmp.[SID]` |
| IC-001426 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:195](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L195) | JOIN_ON / P02 / COMPARISON | `tr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tr.[SID]` → `NVARCHAR` | `tr.[SID] = tmp.[SID]` |
| IC-001430 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L205) | JOIN_ON / P02 / COMPARISON | `tr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tr.[SID]` → `NVARCHAR` | `tr.[SID] = tmp.[SID]` |
| IC-001294 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptosoutstanding_tr.SID <- [SID]` |
| IC-001399 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptosoutstanding_tr.SID <- [SID]` |

<a id="kolom-norekinvestorrdn"></a>
## NoRekInvestorRDN

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001427 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L196) | JOIN_ON / P02 / COMPARISON | `tr.NoRekInvestorRDN` : `VARCHAR(20)` | `tmp.[Nomor Rekening RDN]` : `NVARCHAR(255)` | `tr.NoRekInvestorRDN` → `NVARCHAR` | `tr.NoRekInvestorRDN = tmp.[Nomor Rekening RDN]` |

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001285 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptosoutstanding_tr.AccruedInterest <- AccruedInterest` |
| IC-001301 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L121) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.AccruedInterest` : `MONEY` | `tmp.AccruedInterest` : `FLOAT` | `tmp.AccruedInterest` → `MONEY` | `potr.AccruedInterest = tmp.AccruedInterest` |
| IC-001390 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptosoutstanding_tr.AccruedInterest <- AccruedInterest` |
| IC-001406 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L159) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.AccruedInterest` : `MONEY` | `tmp.AccruedInterest` : `FLOAT` | `tmp.AccruedInterest` → `MONEY` | `potr.AccruedInterest = tmp.AccruedInterest` |
| IC-002372 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:142](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L142) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(potr.AccruedInterest, 0)` |
| IC-002379 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:273](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L273) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(potr.AccruedInterest, 0)` |
| IC-004961 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPTOSPopulateOutstanding:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateOutstanding.sql#L17) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(AccruedInterest,0)` |
| IC-005748 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSOuts:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L33) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(pttr.AccruedInterest,0)` |
| IC-005753 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSOuts:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L70) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `isnull(pttr.AccruedInterest,0)` |
| IC-005831 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L141) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(potr.AccruedInterest, 0)` |
| IC-005838 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:272](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L272) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(potr.AccruedInterest, 0)` |
| IC-010611 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPTOSOutstanding:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSOutstanding.sql#L60) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.AccruedInterest,0)` |

<a id="kolom-approveddate"></a>
## ApprovedDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007130 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOutsAuthen:330](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOutsAuthen.sql#L330) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ApprovedDate` : `DATETIME` | `tmp.ApprovedBy` : `int` | `tmp.ApprovedBy` → `DATETIME` | `ApprovedDate = tmp.ApprovedBy` |

<a id="kolom-ccy"></a>
## Ccy

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001286 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptosoutstanding_tr.Ccy <- Ccy` |
| IC-001298 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L118) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.Ccy` : `VARCHAR(3)` | `tmp.Ccy` : `NVARCHAR(255)` | `tmp.Ccy` → `VARCHAR(3)` | `potr.Ccy = tmp.Ccy` |
| IC-001391 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptosoutstanding_tr.Ccy <- Ccy` |
| IC-001403 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L156) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.Ccy` : `VARCHAR(3)` | `tmp.Ccy` : `NVARCHAR(255)` | `tmp.Ccy` → `VARCHAR(3)` | `potr.Ccy = tmp.Ccy` |

<a id="kolom-hargaperolehannasabah"></a>
## HargaPerolehanNasabah

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001237 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L24) | JOIN_ON / P11 / COMPARISON | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `tmp.[Harga Perolehan Nasabah]` : `FLOAT` | `potr.HargaPerolehanNasabah` → `FLOAT` | `tmp.[Harga Perolehan Nasabah] = potr.HargaPerolehanNasabah` |
| IC-001284 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L101) | JOIN_ON / P11 / COMPARISON | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `tmp.HargaPerolehanNasabah` : `FLOAT` | `potr.HargaPerolehanNasabah` → `FLOAT` | `potr.HargaPerolehanNasabah = tmp.HargaPerolehanNasabah` |
| IC-001306 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L133) | JOIN_ON / P11 / COMPARISON | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `tmp.HargaPerolehanNasabah` : `FLOAT` | `potr.HargaPerolehanNasabah` → `FLOAT` | `potr.HargaPerolehanNasabah = tmp.HargaPerolehanNasabah` |
| IC-001331 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L26) | JOIN_ON / P11 / COMPARISON | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `tmp.[Harga Perolehan Nasabah]` : `FLOAT` | `potr.HargaPerolehanNasabah` → `FLOAT` | `tmp.[Harga Perolehan Nasabah] = potr.HargaPerolehanNasabah` |
| IC-001389 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L138) | JOIN_ON / P11 / COMPARISON | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `tmp.HargaPerolehanNasabah` : `FLOAT` | `potr.HargaPerolehanNasabah` → `FLOAT` | `potr.HargaPerolehanNasabah = tmp.HargaPerolehanNasabah` |
| IC-001411 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:173](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L173) | JOIN_ON / P11 / COMPARISON | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `tmp.HargaPerolehanNasabah` : `FLOAT` | `potr.HargaPerolehanNasabah` → `FLOAT` | `potr.HargaPerolehanNasabah = tmp.HargaPerolehanNasabah` |
| IC-001287 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `HargaPerolehanNasabah` : `FLOAT` | `HargaPerolehanNasabah` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_tr.HargaPerolehanNasabah <- HargaPerolehanNasabah` |
| IC-001392 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `HargaPerolehanNasabah` : `FLOAT` | `HargaPerolehanNasabah` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_tr.HargaPerolehanNasabah <- HargaPerolehanNasabah` |
| IC-002327 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPTOSSubmitFromTemp:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L81) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(potr.HargaPerolehanNasabah, 0)` |
| IC-002334 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPTOSSubmitFromTemp:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L93) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(potr.HargaPerolehanNasabah, 0)` |
| IC-002340 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPTOSSubmitFromTemp:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L111) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(potr.HargaPerolehanNasabah, 0)` |
| IC-002346 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPTOSSubmitFromTemp:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L130) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potr.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `0` : `int` | `0` → `DECIMAL(20,5)` | `ISNULL(potr.HargaPerolehanNasabah, 0)` |

<a id="kolom-indikasihargabid"></a>
## IndikasiHargaBid

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001288 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.IndikasiHargaBid` : `DECIMAL(20,5)` | `IndikasiHargaBid` : `FLOAT` | `IndikasiHargaBid` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_tr.IndikasiHargaBid <- IndikasiHargaBid` |
| IC-001300 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L120) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.IndikasiHargaBid` : `DECIMAL(20,5)` | `tmp.IndikasiHargaBid` : `FLOAT` | `tmp.IndikasiHargaBid` → `DECIMAL(20,5)` | `potr.IndikasiHargaBid = tmp.IndikasiHargaBid` |
| IC-001393 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.IndikasiHargaBid` : `DECIMAL(20,5)` | `IndikasiHargaBid` : `FLOAT` | `IndikasiHargaBid` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_tr.IndikasiHargaBid <- IndikasiHargaBid` |
| IC-001405 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L158) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.IndikasiHargaBid` : `DECIMAL(20,5)` | `tmp.IndikasiHargaBid` : `FLOAT` | `tmp.IndikasiHargaBid` → `DECIMAL(20,5)` | `potr.IndikasiHargaBid = tmp.IndikasiHargaBid` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002373 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L150) | JOIN_ON / P01 / COMPARISON | `potr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `potr.InsertedBy = ndme.EMPLID` |
| IC-002380 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L281) | JOIN_ON / P01 / COMPARISON | `potr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `potr.InsertedBy = ndme.EMPLID` |
| IC-005749 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSOuts:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L40) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-005751 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSOuts:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L51) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-005754 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSOuts:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L77) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-005757 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSOuts:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L89) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-005832 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUploadAuthen:149](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L149) | JOIN_ON / P01 / COMPARISON | `potr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `potr.InsertedBy = ndme.EMPLID` |
| IC-005839 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUploadAuthen:280](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L280) | JOIN_ON / P01 / COMPARISON | `potr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `potr.InsertedBy = ndme.EMPLID` |

<a id="kolom-namanasabah"></a>
## NamaNasabah

**Tipe dasar:** `VARCHAR(95)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001290 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.NamaNasabah` : `VARCHAR(95)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptosoutstanding_tr.NamaNasabah <- NamaNasabah` |
| IC-001296 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L116) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.NamaNasabah` : `VARCHAR(95)` | `tmp.NamaNasabah` : `NVARCHAR(255)` | `tmp.NamaNasabah` → `VARCHAR(95)` | `potr.NamaNasabah = tmp.NamaNasabah` |
| IC-001395 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.NamaNasabah` : `VARCHAR(95)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptosoutstanding_tr.NamaNasabah <- NamaNasabah` |
| IC-001401 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L154) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.NamaNasabah` : `VARCHAR(95)` | `tmp.NamaNasabah` : `NVARCHAR(255)` | `tmp.NamaNasabah` → `VARCHAR(95)` | `potr.NamaNasabah = tmp.NamaNasabah` |

<a id="kolom-namaproduk"></a>
## NamaProduk

**Tipe dasar:** `VARCHAR(70)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001291 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptosoutstanding_tr.NamaProduk <- NamaProduk` |
| IC-001297 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L117) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.NamaProduk` : `VARCHAR(70)` | `tmp.NamaProduk` : `NVARCHAR(255)` | `tmp.NamaProduk` → `VARCHAR(70)` | `potr.NamaProduk = tmp.NamaProduk` |
| IC-001396 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptosoutstanding_tr.NamaProduk <- NamaProduk` |
| IC-001402 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:155](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L155) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.NamaProduk` : `VARCHAR(70)` | `tmp.NamaProduk` : `NVARCHAR(255)` | `tmp.NamaProduk` → `VARCHAR(70)` | `potr.NamaProduk = tmp.NamaProduk` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001292 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptosoutstanding_tr.Nominal <- Nominal` |
| IC-001299 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L119) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.Nominal` : `MONEY` | `tmp.Nominal` : `FLOAT` | `tmp.Nominal` → `MONEY` | `potr.Nominal = tmp.Nominal` |
| IC-001397 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptosoutstanding_tr.Nominal <- Nominal` |
| IC-001404 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L157) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.Nominal` : `MONEY` | `tmp.Nominal` : `FLOAT` | `tmp.Nominal` → `MONEY` | `potr.Nominal = tmp.Nominal` |
| IC-010610 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPTOSOutstanding:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSOutstanding.sql#L59) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `SUM(tr.Nominal)` : `MONEY (aggregate; facets unknown)` | `0` : `int` | `0` → `MONEY (aggregate; facets unknown)` | `ISNULL(SUM(tr.Nominal), 0)` |

<a id="kolom-ptostype"></a>
## PTOSType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 33 lokasi / 11 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001048 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSJobMasterOutstanding:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSJobMasterOutstanding.sql#L57) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.PTOSType` : `SMALLINT` | `#tmpLogData.PTOSType` : `int` | `tr.PTOSType` → `int` | `#tmpLogData.PTOSType <- tr.PTOSType` |
| IC-001051 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSJobMasterOutstanding:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSJobMasterOutstanding.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.PTOSType` : `SMALLINT` | `#tmpLogData.PTOSType` : `int` | `tr.PTOSType` → `int` | `#tmpLogData.PTOSType <- tr.PTOSType` |
| IC-001703 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSSubmitMasterOutstanding:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L47) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.PTOSType` : `SMALLINT` | `#tmpLogData.PTOSType` : `int` | `tr.PTOSType` → `int` | `#tmpLogData.PTOSType <- tr.PTOSType` |
| IC-001707 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSSubmitMasterOutstanding:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L74) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.PTOSType` : `SMALLINT` | `#tmpLogData.PTOSType` : `int` | `tr.PTOSType` → `int` | `#tmpLogData.PTOSType <- tr.PTOSType` |
| IC-001711 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSSubmitMasterOutstanding:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L87) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.PTOSType` : `SMALLINT` | `#tmpLogData.PTOSType` : `int` | `tr.PTOSType` → `int` | `#tmpLogData.PTOSType <- tr.PTOSType` |
| IC-001238 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:26](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L26) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `2` : `int` | `potr.PTOSType` → `INT` | `potr.PTOSType = 2` |
| IC-001283 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L100) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `potr.PTOSType` → `INT` | `potr.PTOSType = tmp.PTOSType` |
| IC-001305 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L132) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `potr.PTOSType` → `INT` | `potr.PTOSType = tmp.PTOSType` |
| IC-001323 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L158) | WHERE / P11 / COMPARISON | `tr.PTOSType` : `SMALLINT` | `2` : `int` | `tr.PTOSType` → `INT` | `tr.PTOSType = 2` |
| IC-001332 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L28) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `1` : `int` | `potr.PTOSType` → `INT` | `potr.PTOSType = 1` |
| IC-001388 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L137) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `potr.PTOSType` → `INT` | `potr.PTOSType = tmp.PTOSType` |
| IC-001410 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L172) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `potr.PTOSType` → `INT` | `potr.PTOSType = tmp.PTOSType` |
| IC-001429 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L198) | WHERE / P11 / COMPARISON | `tr.PTOSType` : `SMALLINT` | `1` : `int` | `tr.PTOSType` → `INT` | `tr.PTOSType = 1` |
| IC-001705 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSSubmitMasterOutstanding:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L56) | WHERE / P11 / COMPARISON | `tr.PTOSType` : `SMALLINT` | `@pnPTOSType` : `int` | `tr.PTOSType` → `INT` | `tr.PTOSType = @pnPTOSType` |
| IC-001706 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSSubmitMasterOutstanding:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L66) | WHERE / P11 / COMPARISON | `PTOSType` : `SMALLINT` | `@pnPTOSType` : `int` | `PTOSType` → `INT` | `PTOSType = @pnPTOSType` |
| IC-001709 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSSubmitMasterOutstanding:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L82) | WHERE / P11 / COMPARISON | `tr.PTOSType` : `SMALLINT` | `@pnPTOSType` : `int` | `tr.PTOSType` → `INT` | `tr.PTOSType = @pnPTOSType` |
| IC-001713 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSSubmitMasterOutstanding:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L97) | WHERE / P11 / COMPARISON | `tr.PTOSType` : `SMALLINT` | `@pnPTOSType` : `int` | `tr.PTOSType` → `INT` | `tr.PTOSType = @pnPTOSType` |
| IC-002301 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L197) | WHERE / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `potr.PTOSType` → `INT` | `potr.PTOSType = @pnType` |
| IC-002326 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L80) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `potr.PTOSType` → `INT` | `tmp.PTOSType = potr.PTOSType` |
| IC-002333 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L92) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `potr.PTOSType` → `INT` | `potr.PTOSType = tmp.PTOSType` |
| IC-002339 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L110) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `potr.PTOSType` → `INT` | `tmp.PTOSType = potr.PTOSType` |
| IC-002345 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L129) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `potr.PTOSType` → `INT` | `tmp.PTOSType = potr.PTOSType` |
| IC-002382 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:283](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L283) | WHERE / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `potr.PTOSType` → `INT` | `potr.PTOSType = @pnType` |
| IC-004962 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateOutstanding:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateOutstanding.sql#L19) | WHERE / P11 / COMPARISON | `PTOSType` : `SMALLINT` | `1` : `int` | `PTOSType` → `INT` | `PTOSType = 1` |
| IC-004964 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateOutstanding:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateOutstanding.sql#L32) | WHERE / P11 / COMPARISON | `PTOSType` : `SMALLINT` | `2` : `int` | `PTOSType` → `INT` | `PTOSType = 2` |
| IC-005756 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSOuts:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L79) | WHERE / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = @pnType` |
| IC-005759 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSOuts:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L91) | WHERE / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = @pnType` |
| IC-005819 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:503](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L503) | WHERE / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `potr.PTOSType` → `INT` | `potr.PTOSType = @pnType` |
| IC-005822 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:628](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L628) | WHERE / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `potr.PTOSType` → `INT` | `potr.PTOSType = @pnType` |
| IC-005841 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L282) | WHERE / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `potr.PTOSType` → `INT` | `potr.PTOSType = @pnType` |
| IC-001293 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.PTOSType` : `SMALLINT` | `PTOSType` : `int` | `PTOSType` → `SMALLINT` | `retail.dbo.ptosoutstanding_tr.PTOSType <- PTOSType` |
| IC-001398 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.PTOSType` : `SMALLINT` | `PTOSType` : `int` | `PTOSType` → `SMALLINT` | `retail.dbo.ptosoutstanding_tr.PTOSType <- PTOSType` |
| IC-002343 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `tmp.PTOSType` → `SMALLINT` | `retail.dbo.ptosoutstanding_tr.PTOSType <- tmp.PTOSType` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TR.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 44 lokasi / 15 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001049 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSJobMasterOutstanding:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSJobMasterOutstanding.sql#L57) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.[Status]` : `SMALLINT` | `#tmpLogData.Status` : `int` | `tr.[Status]` → `int` | `#tmpLogData.Status <- tr.[Status]` |
| IC-001052 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSJobMasterOutstanding:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSJobMasterOutstanding.sql#L70) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.[Status]` : `SMALLINT` | `#tmpLogData.Status` : `int` | `tr.[Status]` → `int` | `#tmpLogData.Status <- tr.[Status]` |
| IC-001704 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSSubmitMasterOutstanding:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L47) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.[Status]` : `SMALLINT` | `#tmpLogData.Status` : `int` | `tr.[Status]` → `int` | `#tmpLogData.Status <- tr.[Status]` |
| IC-001708 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSSubmitMasterOutstanding:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L74) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.[Status]` : `SMALLINT` | `#tmpLogData.Status` : `int` | `tr.[Status]` → `int` | `#tmpLogData.Status <- tr.[Status]` |
| IC-001712 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSSubmitMasterOutstanding:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSSubmitMasterOutstanding.sql#L87) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tr.[Status]` : `SMALLINT` | `#tmpLogData.Status` : `int` | `tr.[Status]` → `int` | `#tmpLogData.Status <- tr.[Status]` |
| IC-002302 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:197](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L197) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `3` : `int` | `potr.Status` → `INT` | `potr.Status != 3` |
| IC-002329 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.Rvmp_TRSPTOSSubmitFromTemp:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L82) | WHERE / P11 / COMPARISON | `ISNULL(potr.Status, 0)` : `SMALLINT` | `3` : `int` | `ISNULL(potr.Status, 0)` → `INT` | `ISNULL(potr.Status, 0) != 3` |
| IC-002374 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L151) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `0` : `int` | `potr.Status` → `INT` | `potr.Status = 0` |
| IC-002381 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:282](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L282) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `0` : `int` | `potr.Status` → `INT` | `potr.Status = 0` |
| IC-004963 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateOutstanding:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateOutstanding.sql#L21) | WHERE / P11 / COMPARISON | `[Status]` : `SMALLINT` | `3` : `int` | `[Status]` → `INT` | `[Status] = 3` |
| IC-004965 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateOutstanding:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateOutstanding.sql#L34) | WHERE / P11 / COMPARISON | `[Status]` : `SMALLINT` | `3` : `int` | `[Status]` → `INT` | `[Status] = 3` |
| IC-005750 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSOuts:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L41) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `3` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 3` |
| IC-005752 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSOuts:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L52) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `3` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 3` |
| IC-005755 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSOuts:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L78) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `3` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 3` |
| IC-005758 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSOuts:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSOuts.sql#L90) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `3` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 3` |
| IC-005818 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:502](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L502) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `3` : `int` | `potr.Status` → `INT` | `potr.Status != 3` |
| IC-005821 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:627](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L627) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `3` : `int` | `potr.Status` → `INT` | `potr.Status != 3` |
| IC-005833 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:150](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L150) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `0` : `int` | `potr.Status` → `INT` | `potr.Status = 0` |
| IC-005840 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L281) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `0` : `int` | `potr.Status` → `INT` | `potr.Status = 0` |
| IC-007196 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSUpload:537](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L537) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `3` : `int` | `potr.Status` → `INT` | `potr.Status != 3` |
| IC-007200 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSUpload:841](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L841) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `3` : `int` | `potr.Status` → `INT` | `potr.Status != 3` |
| IC-007224 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:480](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L480) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `0` : `int` | `potr.Status` → `INT` | `potr.Status = 0` |
| IC-007226 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:507](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L507) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `1` : `int` | `potr.Status` → `INT` | `potr.Status = 1` |
| IC-007228 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:528](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L528) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `0` : `int` | `potr.Status` → `INT` | `potr.Status = 0` |
| IC-007231 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:547](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L547) | WHERE / P11 / COMPARISON | `potr.Status` : `SMALLINT` | `1` : `int` | `potr.Status` → `INT` | `potr.Status = 1` |
| IC-001295 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L104) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.ptosoutstanding_tr.Status <- 3` |
| IC-001302 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L122) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `potr.[Status] = 3` |
| IC-001400 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.ptosoutstanding_tr.Status <- 3` |
| IC-001407 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:160](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L160) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `potr.[Status] = 3` |
| IC-002331 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPTOSSubmitFromTemp:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L82) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `potr.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `ISNULL(potr.Status, 0)` |
| IC-002332 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:87](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L87) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.[Status]` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `potr.[Status] = 1` |
| IC-002344 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.ptosoutstanding_tr.Status <- 0` |
| IC-007117 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOuts:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOuts.sql#L143) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.Status` : `SMALLINT` | `@pnStatus` : `INT` | `@pnStatus` → `SMALLINT` | `potr.Status = @pnStatus` |
| IC-007129 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOutsAuthen:326](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOutsAuthen.sql#L326) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `Status = 3` |
| IC-007133 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOutsAuthen:350](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOutsAuthen.sql#L350) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.ptosoutstanding_tr.Status <- 3` |
| IC-007136 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOutsAuthen:406](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOutsAuthen.sql#L406) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `[Status] = 3` |
| IC-007139 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOutsAuthen:427](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOutsAuthen.sql#L427) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `[Status] = 3` |
| IC-007198 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:625](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L625) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `[Status] = 1` |
| IC-007199 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:634](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L634) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.ptosoutstanding_tr.Status <- 0` |
| IC-007202 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:926](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L926) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `[Status] = 1` |
| IC-007203 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:935](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L935) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_tr.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.ptosoutstanding_tr.Status <- 0` |
| IC-007223 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:470](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L470) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `potr.Status = 3` |
| IC-007225 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:494](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L494) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `potr.Status = 3` |
| IC-007230 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:540](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L540) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `potr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `potr.Status = 3` |
