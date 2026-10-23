# SQL_REPLICATE.dbo.DDPAR2 — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 86 lokasi source, 15 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-clsfee"></a>
## CLSFEE

**Tipe dasar:** `DECIMAL(11,2)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDPAR2.sql#L48). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004115 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGetBlokirAmount:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetBlokirAmount.sql#L18) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `CLSFEE` : `DECIMAL(11,2)` | `@mClosingFee` : `money` | `CLSFEE` → `money` | `@mClosingFee = CLSFEE` |

<a id="kolom-dp2cur"></a>
## DP2CUR

**Tipe dasar:** `VARCHAR(4)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDPAR2.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000499 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountInfo:100](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L100) | JOIN_ON / P12 / COMPARISON | `dp.DP2CUR` : `VARCHAR(4)` | `ta.CurrencyCode` : `CHAR(3)` | `ta.CurrencyCode` → `VARCHAR` | `ta.CurrencyCode = dp.DP2CUR` |

<a id="kolom-dp2flg3"></a>
## DP2FLG3

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDPAR2.sql#L51). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 65 lokasi / 14 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000503 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetAccountInfo:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L115) | WHERE / P12 / COMPARISON | `dp.DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `dp.DP2FLG3` → `VARCHAR` | `dp.DP2FLG3='Y'` |
| IC-000509 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetAccountInfo:179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L179) | WHERE / P12 / COMPARISON | `dp.DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `dp.DP2FLG3` → `VARCHAR` | `dp.DP2FLG3='Y'` |
| IC-000512 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetAccountSyariah:34](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountSyariah.sql#L34) | WHERE / P12 / COMPARISON | `DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `DP2FLG3` → `VARCHAR` | `DP2FLG3='Y'` |
| IC-000514 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetAccountSyariah:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountSyariah.sql#L46) | WHERE / P12 / COMPARISON | `DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `DP2FLG3` → `VARCHAR` | `DP2FLG3='Y'` |
| IC-000559 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabah:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L124) | WHERE / P12 / COMPARISON | `dp.DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `dp.DP2FLG3` → `VARCHAR` | `dp.DP2FLG3='Y'` |
| IC-000562 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabah:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L140) | WHERE / P12 / COMPARISON | `dp.DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `dp.DP2FLG3` → `VARCHAR` | `dp.DP2FLG3='Y'` |
| IC-000590 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabah:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L301) | WHERE / P12 / COMPARISON | `par.DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `par.DP2FLG3` → `VARCHAR` | `par.DP2FLG3 <> 'Y'` |
| IC-000633 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:272](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L272) | WHERE / P12 / COMPARISON | `ddp.DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `ddp.DP2FLG3` → `VARCHAR` | `ddp.DP2FLG3='Y'` |
| IC-000634 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:279](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L279) | WHERE / P12 / COMPARISON | `ddp.DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `ddp.DP2FLG3` → `VARCHAR` | `ddp.DP2FLG3 <> 'Y'` |
| IC-004321 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:859](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L859) | WHERE / P12 / COMPARISON | `par.DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `par.DP2FLG3` → `VARCHAR` | `par.DP2FLG3 = 'Y'` |
| IC-009831 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:925](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L925) | EXPRESSION / P07 / CASE_BRANCH | `DP2FLG3` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `DP2FLG3` → `VARCHAR` | `case when (isnull(DP2FLG3,'') = '') or (isnull(DP2FLG3,'') != 'Y') then 'N' else DP2FLG3 end` |
| IC-009832 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrx:926](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L926) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `''` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') = ''` |
| IC-009833 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrx:926](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L926) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `'Y'` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') != 'Y'` |
| IC-009838 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:995](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L995) | EXPRESSION / P07 / CASE_BRANCH | `DP2FLG3` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `DP2FLG3` → `VARCHAR` | `case when (isnull(DP2FLG3,'') = '') or (isnull(DP2FLG3,'') != 'Y') then 'N' else DP2FLG3 end` |
| IC-009839 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrx:996](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L996) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `''` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') = ''` |
| IC-009840 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrx:996](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L996) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `'Y'` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') != 'Y'` |
| IC-009846 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:1066](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1066) | EXPRESSION / P07 / CASE_BRANCH | `DP2FLG3` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `DP2FLG3` → `VARCHAR` | `case when (isnull(DP2FLG3,'') = '') or (isnull(DP2FLG3,'') != 'Y') then 'N' else DP2FLG3 end` |
| IC-009847 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrx:1067](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1067) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `''` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') = ''` |
| IC-009848 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrx:1067](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1067) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `'Y'` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') != 'Y'` |
| IC-009854 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:1129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1129) | EXPRESSION / P07 / CASE_BRANCH | `DP2FLG3` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `DP2FLG3` → `VARCHAR` | `case when (isnull(DP2FLG3,'') = '') or (isnull(DP2FLG3,'') != 'Y') then 'N' else DP2FLG3 end` |
| IC-009855 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrx:1130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1130) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `''` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') = ''` |
| IC-009856 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrx:1130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1130) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `'Y'` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') != 'Y'` |
| IC-009861 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrx:1140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1140) | WHERE / P12 / COMPARISON | `DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `DP2FLG3` → `VARCHAR` | `DP2FLG3 = 'Y'` |
| IC-009922 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrxTaxAmnesty:861](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L861) | EXPRESSION / P07 / CASE_BRANCH | `DP2FLG3` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `DP2FLG3` → `VARCHAR` | `case when (isnull(DP2FLG3,'') = '') or (isnull(DP2FLG3,'') != 'Y') then 'N' else DP2FLG3 end` |
| IC-009923 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrxTaxAmnesty:862](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L862) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `''` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') = ''` |
| IC-009924 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrxTaxAmnesty:862](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L862) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `'Y'` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') != 'Y'` |
| IC-009930 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrxTaxAmnesty:951](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L951) | EXPRESSION / P07 / CASE_BRANCH | `DP2FLG3` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `DP2FLG3` → `VARCHAR` | `case when (isnull(DP2FLG3,'') = '') or (isnull(DP2FLG3,'') != 'Y') then 'N' else DP2FLG3 end` |
| IC-009931 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrxTaxAmnesty:952](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L952) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `''` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') = ''` |
| IC-009932 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrxTaxAmnesty:952](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L952) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `'Y'` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') != 'Y'` |
| IC-009939 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDetailTrxTaxAmnesty:1042](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L1042) | EXPRESSION / P07 / CASE_BRANCH | `DP2FLG3` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `DP2FLG3` → `VARCHAR` | `case when (isnull(DP2FLG3,'') = '') or (isnull(DP2FLG3,'') != 'Y') then 'N' else DP2FLG3 end` |
| IC-009940 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrxTaxAmnesty:1043](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L1043) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `''` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') = ''` |
| IC-009941 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptDetailTrxTaxAmnesty:1043](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L1043) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `'Y'` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') != 'Y'` |
| IC-010427 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptOrderORIDetail:1527](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql#L1527) | WHERE / P12 / COMPARISON | `DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `DP2FLG3` → `VARCHAR` | `DP2FLG3 = 'Y'` |
| IC-010446 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptOrderORIDetail:1664](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql#L1664) | WHERE / P12 / COMPARISON | `DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `DP2FLG3` → `VARCHAR` | `DP2FLG3 = 'Y'` |
| IC-011632 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptTransactionDetailWMTA:893](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L893) | EXPRESSION / P07 / CASE_BRANCH | `DP2FLG3` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `DP2FLG3` → `VARCHAR` | `case when (isnull(DP2FLG3,'') = '') or (isnull(DP2FLG3,'') != 'Y') then 'N' else DP2FLG3 end` |
| IC-011633 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptTransactionDetailWMTA:894](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L894) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `''` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') = ''` |
| IC-011634 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptTransactionDetailWMTA:894](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L894) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `'Y'` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') != 'Y'` |
| IC-011639 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptTransactionDetailWMTA:984](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L984) | EXPRESSION / P07 / CASE_BRANCH | `DP2FLG3` : `CHAR(1)` | `result type` : `VARCHAR (combined result; facets unknown)` | `DP2FLG3` → `VARCHAR` | `case when (isnull(DP2FLG3,'') = '') or (isnull(DP2FLG3,'') != 'Y') then 'N' else DP2FLG3 end` |
| IC-011640 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptTransactionDetailWMTA:985](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L985) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `''` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') = ''` |
| IC-011641 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSRptTransactionDetailWMTA:985](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L985) | CASE / P12 / COMPARISON | `isnull(DP2FLG3,'')` : `CHAR(1)` | `'Y'` : `varchar` | `isnull(DP2FLG3,'')` → `VARCHAR` | `isnull(DP2FLG3,'') != 'Y'` |
| IC-013157 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:609](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L609) | WHERE / P12 / COMPARISON | `DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `DP2FLG3` → `VARCHAR` | `DP2FLG3='Y'` |
| IC-013162 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:639](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L639) | WHERE / P12 / COMPARISON | `DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `DP2FLG3` → `VARCHAR` | `DP2FLG3='Y'` |
| IC-013614 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateDataRekeningMaterai:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L80) | WHERE / P12 / COMPARISON | `dp.DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `dp.DP2FLG3` → `VARCHAR` | `dp.DP2FLG3 = 'Y'` |
| IC-013635 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_PopulateDataRekeningObligasi:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L85) | WHERE / P12 / COMPARISON | `dp.DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `dp.DP2FLG3` → `VARCHAR` | `dp.DP2FLG3 = 'Y'` |
| IC-013807 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:247](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L247) | WHERE / P12 / COMPARISON | `DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `DP2FLG3` → `VARCHAR` | `DP2FLG3 = 'Y'` |
| IC-013853 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:507](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L507) | WHERE / P12 / COMPARISON | `DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `DP2FLG3` → `VARCHAR` | `DP2FLG3='Y'` |
| IC-013857 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateTreasuryCustomer_TM:537](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L537) | WHERE / P12 / COMPARISON | `DP2FLG3` : `CHAR(1)` | `'Y'` : `varchar` | `DP2FLG3` → `VARCHAR` | `DP2FLG3='Y'` |
| IC-009834 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrx:926](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L926) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009835 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrx:926](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L926) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009841 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrx:996](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L996) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009842 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrx:996](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L996) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009849 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrx:1067](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1067) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009850 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrx:1067](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1067) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009857 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrx:1130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1130) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009858 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrx:1130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1130) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009925 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrxTaxAmnesty:862](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L862) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009926 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrxTaxAmnesty:862](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L862) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009933 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrxTaxAmnesty:952](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L952) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009934 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrxTaxAmnesty:952](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L952) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009942 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrxTaxAmnesty:1043](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L1043) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-009943 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptDetailTrxTaxAmnesty:1043](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L1043) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-011635 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptTransactionDetailWMTA:894](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L894) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-011636 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptTransactionDetailWMTA:894](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L894) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-011642 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptTransactionDetailWMTA:985](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L985) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |
| IC-011643 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptTransactionDetailWMTA:985](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L985) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `DP2FLG3` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `isnull(DP2FLG3,'')` |

<a id="kolom-minblm"></a>
## MINBLM

**Tipe dasar:** `NUMERIC(17,2)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDPAR2.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004114 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGetBlokirAmount:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetBlokirAmount.sql#L17) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `MINBLM` : `NUMERIC(17,2)` | `@mMinimumBalance` : `money` | `MINBLM` → `money` | `@mMinimumBalance = MINBLM` |

<a id="kolom-sccode"></a>
## SCCODE

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDPAR2.sql#L2). **Existing leading index:** idx_sccode_dp2cur.

**Frekuensi source:** 18 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000510 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountSyariah:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountSyariah.sql#L32) | JOIN_ON / P12 / COMPARISON | `par.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = par.SCCODE` |
| IC-004320 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertSecurityTransaction_TT_Original:856](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L856) | JOIN_ON / P12 / COMPARISON | `par.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = par.SCCODE` |
| IC-009837 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrx:934](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L934) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-009844 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrx:1004](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1004) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-009852 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrx:1075](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1075) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-009860 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrx:1138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrx.sql#L1138) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-009928 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrxTaxAmnesty:870](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L870) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-009936 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrxTaxAmnesty:960](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L960) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-009945 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDetailTrxTaxAmnesty:1051](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDetailTrxTaxAmnesty.sql#L1051) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-010426 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOrderORIDetail:1526](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql#L1526) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-010445 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOrderORIDetail:1663](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql#L1663) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-011638 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptTransactionDetailWMTA:906](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L906) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-011645 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptTransactionDetailWMTA:997](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L997) | JOIN_ON / P12 / COMPARISON | `dp.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = dp.SCCODE` |
| IC-013156 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertTreasuryCustomer_TM:608](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L608) | JOIN_ON / P12 / COMPARISON | `par.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = par.SCCODE` |
| IC-013161 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_InsertTreasuryCustomer_TM:638](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L638) | JOIN_ON / P12 / COMPARISON | `par.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = par.SCCODE` |
| IC-013806 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_UpdateStatusSecurityTransactionAfterMurex:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransactionAfterMurex.sql#L245) | JOIN_ON / P12 / COMPARISON | `par.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = par.SCCODE` |
| IC-013852 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_UpdateTreasuryCustomer_TM:506](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L506) | JOIN_ON / P12 / COMPARISON | `par.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = par.SCCODE` |
| IC-013856 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_UpdateTreasuryCustomer_TM:536](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateTreasuryCustomer_TM.sql#L536) | JOIN_ON / P12 / COMPARISON | `par.SCCODE` : `VARCHAR(10)` | `dd.SCCODE` : `CHAR(10)` | `dd.SCCODE` → `VARCHAR` | `dd.SCCODE = par.SCCODE` |
