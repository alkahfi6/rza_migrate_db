# Retail UAT.dbo.SecurityTransferAset_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 56 lokasi source, 8 procedure, 9 kolom. **A:** 2 lokasi pada 1 procedure. **B:** 22 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAset_TM.sql#L2). **Existing leading index:** SecurityTransferAset_TM_CN1.

**Frekuensi source:** 6 lokasi / 3 SP. **A:** 2 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012441 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSTransferAssetDealIdSrc:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L84) | WHERE / P06 / LIKE | `tm.DealId` : `INT` | `rtrim(ltrim(@cCol1)) + '%'` : `VARCHAR (combined result; facets unknown)` | `tm.DealId` → `VARCHAR (combined result; facets unknown)` | `tm.DealId like rtrim(ltrim(@cCol1)) + '%'` |
| IC-012444 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSTransferAssetDealIdSrc:146](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L146) | WHERE / P06 / LIKE | `tm.DealId` : `INT` | `rtrim(ltrim(@cCol1)) + '%'` : `VARCHAR (combined result; facets unknown)` | `tm.DealId` → `VARCHAR (combined result; facets unknown)` | `tm.DealId like rtrim(ltrim(@cCol1)) + '%'` |
| IC-009948 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L53) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.DealId` : `INT` | `#tempResult.DealIdParent` : `bigint` | `sta.DealId` → `bigint` | `#tempResult.DealIdParent <- sta.DealId` |
| IC-009958 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.DealId` : `INT` | `#tempResult.DealIdParent` : `bigint` | `sta.DealId` → `bigint` | `#tempResult.DealIdParent <- sta.DealId` |
| IC-012221 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSentEmailTransferAsset:175](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L175) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.DealId` : `INT` | `DealNo` : `VARCHAR(40)` | `tm.DealId` → `VARCHAR(40)` | `DealNo = tm.DealId` |
| IC-012439 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSTransferAssetDealIdSrc:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L49) | WHERE / P04 / COMPARISON | `tm.DealId` : `INT` | `@cCol1` : `VARCHAR(20)` | `@cCol1` → `INT` | `tm.DealId = @cCol1` |

<a id="kolom-approvalnik"></a>
## ApprovalNIK

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAset_TM.sql#L26). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009949 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L53) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.ApprovalNIK` : `INT` | `#tempResult.NIKApproval` : `bigint` | `sta.ApprovalNIK` → `bigint` | `#tempResult.NIKApproval <- sta.ApprovalNIK` |
| IC-009959 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.ApprovalNIK` : `INT` | `#tempResult.NIKApproval` : `bigint` | `sta.ApprovalNIK` → `bigint` | `#tempResult.NIKApproval <- sta.ApprovalNIK` |
| IC-013051 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_GetDataTransaksiAktifitas:334](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetDataTransaksiAktifitas.sql#L334) | JOIN_ON / P01 / COMPARISON | `tm.ApprovalNIK` : `INT` | `Supervisor.EMPLID` : `VARCHAR(11)` | `Supervisor.EMPLID` → `INT` | `tm.ApprovalNIK = Supervisor.EMPLID` |

<a id="kolom-confrimnik"></a>
## ConfrimNIK

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAset_TM.sql#L29). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009950 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L53) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.ConfrimNIK` : `INT` | `#tempResult.NIKConfirm` : `bigint` | `sta.ConfrimNIK` → `bigint` | `#tempResult.NIKConfirm <- sta.ConfrimNIK` |
| IC-009960 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.ConfrimNIK` : `INT` | `#tempResult.NIKConfirm` : `bigint` | `sta.ConfrimNIK` → `bigint` | `#tempResult.NIKConfirm <- sta.ConfrimNIK` |

<a id="kolom-creatednik"></a>
## CreatedNIK

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAset_TM.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009951 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L53) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.CreatedNIK` : `INT` | `#tempResult.NIKInputter` : `varchar(20)` | `sta.CreatedNIK` → `varchar(20)` | `#tempResult.NIKInputter <- sta.CreatedNIK` |
| IC-009961 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.CreatedNIK` : `INT` | `#tempResult.NIKInputter` : `varchar(20)` | `sta.CreatedNIK` → `varchar(20)` | `#tempResult.NIKInputter <- sta.CreatedNIK` |
| IC-005257 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalFOPTransferAsset:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalFOPTransferAsset.sql#L75) | JOIN_ON / P11 / COMPARISON | `tm.CreatedNIK` : `INT` | `usr.nik` : `DECIMAL (combined result; facets unknown)` | `tm.CreatedNIK` → `DECIMAL` | `tm.CreatedNIK = usr.nik` |
| IC-005260 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalFOPTransferAsset:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalFOPTransferAsset.sql#L140) | JOIN_ON / P11 / COMPARISON | `tm.CreatedNIK` : `INT` | `usr.nik` : `DECIMAL (combined result; facets unknown)` | `tm.CreatedNIK` → `DECIMAL` | `tm.CreatedNIK = usr.nik` |
| IC-012222 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSentEmailTransferAsset:230](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L230) | JOIN_ON / P11 / COMPARISON | `tm.CreatedNIK` : `INT` | `inp.nik` : `DECIMAL (combined result; facets unknown)` | `tm.CreatedNIK` → `DECIMAL` | `inp.nik = tm.CreatedNIK` |
| IC-013050 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_GetDataTransaksiAktifitas:333](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetDataTransaksiAktifitas.sql#L333) | JOIN_ON / P01 / COMPARISON | `tm.CreatedNIK` : `INT` | `Inputer.EMPLID` : `VARCHAR(11)` | `Inputer.EMPLID` → `INT` | `tm.CreatedNIK = Inputer.EMPLID` |

<a id="kolom-jenistransaksi"></a>
## JenisTransaksi

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAset_TM.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 15 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006852 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L43) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.JenisTransaksi` : `INT` | `retail.dbo.securitytransaction_tt.TrxType` : `SMALLINT` | `tm.JenisTransaksi` → `SMALLINT` | `retail.dbo.securitytransaction_tt.TrxType <- tm.JenisTransaksi` |
| IC-006861 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L85) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.JenisTransaksi` : `INT` | `retail.dbo.securitytransaction_tt.TrxType` : `SMALLINT` | `tm.JenisTransaksi` → `SMALLINT` | `retail.dbo.securitytransaction_tt.TrxType <- tm.JenisTransaksi` |
| IC-006880 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `tm.JenisTransaksi` : `INT` | `retail.dbo.securitytransaction_tt.TrxType` : `SMALLINT` | `tm.JenisTransaksi` → `SMALLINT` | `retail.dbo.securitytransaction_tt.TrxType <- tm.JenisTransaksi` |
| IC-006855 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessFOPTransferAsset:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L74) | JOIN_ON / P11 / COMPARISON | `tm.JenisTransaksi` : `INT` | `rtt.TrxType` : `SMALLINT` | `rtt.TrxType` → `INT` | `rtt.TrxType = tm.JenisTransaksi` |
| IC-006863 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessFOPTransferAsset:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L118) | JOIN_ON / P11 / COMPARISON | `tm.JenisTransaksi` : `INT` | `rtt.TrxType` : `SMALLINT` | `rtt.TrxType` → `INT` | `rtt.TrxType = tm.JenisTransaksi` |
| IC-006884 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessFOPTransferAsset:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L159) | JOIN_ON / P11 / COMPARISON | `tm.JenisTransaksi` : `INT` | `rtt.TrxType` : `SMALLINT` | `rtt.TrxType` → `INT` | `rtt.TrxType = tm.JenisTransaksi` |
| IC-006891 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessFOPTransferAsset:301](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L301) | JOIN_ON / P11 / COMPARISON | `tm.JenisTransaksi` : `INT` | `rtt.TrxType` : `SMALLINT` | `rtt.TrxType` → `INT` | `rtt.TrxType = tm.JenisTransaksi` |
| IC-009955 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptFOPandTransferAsset:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L108) | JOIN_ON / P01 / COMPARISON | `sta.JenisTransaksi` : `INT` | `par2.ParamValue` : `VARCHAR(30)` | `par2.ParamValue` → `INT` | `sta.JenisTransaksi = par2.ParamValue` |
| IC-009965 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptFOPandTransferAsset:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L180) | JOIN_ON / P01 / COMPARISON | `sta.JenisTransaksi` : `INT` | `par2.ParamValue` : `VARCHAR(30)` | `par2.ParamValue` → `INT` | `sta.JenisTransaksi = par2.ParamValue` |
| IC-012438 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSTransferAssetDealIdSrc:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L41) | JOIN_ON / P01 / COMPARISON | `tm.JenisTransaksi` : `INT` | `tr1.ParamValue` : `VARCHAR(30)` | `tr1.ParamValue` → `INT` | `tm.JenisTransaksi = tr1.ParamValue` |
| IC-012440 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSTransferAssetDealIdSrc:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L76) | JOIN_ON / P01 / COMPARISON | `tm.JenisTransaksi` : `INT` | `tr1.ParamValue` : `VARCHAR(30)` | `tr1.ParamValue` → `INT` | `tm.JenisTransaksi = tr1.ParamValue` |
| IC-012442 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSTransferAssetDealIdSrc:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L107) | JOIN_ON / P01 / COMPARISON | `tm.JenisTransaksi` : `INT` | `tr1.ParamValue` : `VARCHAR(30)` | `tr1.ParamValue` → `INT` | `tm.JenisTransaksi = tr1.ParamValue` |
| IC-012443 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSTransferAssetDealIdSrc:138](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L138) | JOIN_ON / P01 / COMPARISON | `tm.JenisTransaksi` : `INT` | `tr1.ParamValue` : `VARCHAR(30)` | `tr1.ParamValue` → `INT` | `tm.JenisTransaksi = tr1.ParamValue` |
| IC-012445 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSTransferAssetDealIdSrc:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSTransferAssetDealIdSrc.sql#L170) | JOIN_ON / P01 / COMPARISON | `tm.JenisTransaksi` : `INT` | `tr1.ParamValue` : `VARCHAR(30)` | `tr1.ParamValue` → `INT` | `tm.JenisTransaksi = tr1.ParamValue` |
| IC-013052 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_GetDataTransaksiAktifitas:336](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetDataTransaksiAktifitas.sql#L336) | JOIN_ON / P01 / COMPARISON | `tm.JenisTransaksi` : `INT` | `par.ParamValue` : `VARCHAR(30)` | `par.ParamValue` → `INT` | `tm.JenisTransaksi = par.ParamValue` |

<a id="kolom-nikseller"></a>
## NIKSeller

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAset_TM.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009952 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L53) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.NIKSeller` : `INT` | `#tempResult.NIKSeller` : `bigint` | `sta.NIKSeller` → `bigint` | `#tempResult.NIKSeller <- sta.NIKSeller` |
| IC-009962 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.NIKSeller` : `INT` | `#tempResult.NIKSeller` : `bigint` | `sta.NIKSeller` → `bigint` | `#tempResult.NIKSeller <- sta.NIKSeller` |
| IC-005258 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalFOPTransferAsset:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalFOPTransferAsset.sql#L77) | JOIN_ON / P11 / COMPARISON | `tm.NIKSeller` : `INT` | `slr.nik` : `DECIMAL (combined result; facets unknown)` | `tm.NIKSeller` → `DECIMAL` | `tm.NIKSeller = slr.nik` |
| IC-005261 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalFOPTransferAsset:142](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalFOPTransferAsset.sql#L142) | JOIN_ON / P11 / COMPARISON | `tm.NIKSeller` : `INT` | `slr.nik` : `DECIMAL (combined result; facets unknown)` | `tm.NIKSeller` → `DECIMAL` | `tm.NIKSeller = slr.nik` |
| IC-012223 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSentEmailTransferAsset:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L232) | JOIN_ON / P11 / COMPARISON | `tm.NIKSeller` : `INT` | `slr.nik` : `DECIMAL (combined result; facets unknown)` | `tm.NIKSeller` → `DECIMAL` | `slr.nik = tm.NIKSeller` |
| IC-013049 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_GetDataTransaksiAktifitas:332](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_GetDataTransaksiAktifitas.sql#L332) | JOIN_ON / P01 / COMPARISON | `tm.NIKSeller` : `INT` | `seller.EMPLID` : `VARCHAR(11)` | `seller.EMPLID` → `INT` | `tm.NIKSeller = seller.EMPLID` |

<a id="kolom-purpose"></a>
## Purpose

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAset_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006893 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessFOPTransferAsset:318](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessFOPTransferAsset.sql#L318) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `Purpose` : `INT` | `retail.dbo.securitytransferaset_th.Purpose` : `VARCHAR(30)` | `Purpose` → `VARCHAR(30)` | `retail.dbo.securitytransferaset_th.Purpose <- Purpose` |
| IC-012372 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSubmitTransferAsset:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitTransferAsset.sql#L274) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `Purpose` : `INT` | `retail.dbo.securitytransferaset_th.Purpose` : `VARCHAR(30)` | `Purpose` → `VARCHAR(30)` | `retail.dbo.securitytransferaset_th.Purpose <- Purpose` |
| IC-012373 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSubmitTransferAsset:333](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitTransferAsset.sql#L333) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `Purpose` : `INT` | `retail.dbo.securitytransferaset_th.Purpose` : `VARCHAR(30)` | `Purpose` → `VARCHAR(30)` | `retail.dbo.securitytransferaset_th.Purpose <- Purpose` |
| IC-012374 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSubmitTransferAsset:372](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitTransferAsset.sql#L372) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `Purpose` : `INT` | `retail.dbo.securitytransferaset_th.Purpose` : `VARCHAR(30)` | `Purpose` → `VARCHAR(30)` | `retail.dbo.securitytransferaset_th.Purpose <- Purpose` |
| IC-012650 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateStatusTransferAsset:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusTransferAsset.sql#L27) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `Purpose` : `INT` | `retail.dbo.securitytransferaset_th.Purpose` : `VARCHAR(30)` | `Purpose` → `VARCHAR(30)` | `retail.dbo.securitytransferaset_th.Purpose <- Purpose` |
| IC-012655 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSUpdateStatusTransferAsset:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusTransferAsset.sql#L112) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `Purpose` : `INT` | `retail.dbo.securitytransferaset_th.Purpose` : `VARCHAR(30)` | `Purpose` → `VARCHAR(30)` | `retail.dbo.securitytransferaset_th.Purpose <- Purpose` |
| IC-005259 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprovalFOPTransferAsset:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalFOPTransferAsset.sql#L79) | JOIN_ON / P01 / COMPARISON | `tm.Purpose` : `INT` | `par1.ParamValue` : `VARCHAR(30)` | `par1.ParamValue` → `INT` | `tm.Purpose = par1.ParamValue` |
| IC-005262 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateApprovalFOPTransferAsset:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalFOPTransferAsset.sql#L144) | JOIN_ON / P01 / COMPARISON | `tm.Purpose` : `INT` | `par1.ParamValue` : `VARCHAR(30)` | `par1.ParamValue` → `INT` | `tm.Purpose = par1.ParamValue` |
| IC-009956 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptFOPandTransferAsset:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L111) | JOIN_ON / P01 / COMPARISON | `sta.Purpose` : `INT` | `par3.ParamValue` : `VARCHAR(30)` | `par3.ParamValue` → `INT` | `sta.Purpose = par3.ParamValue` |
| IC-009966 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptFOPandTransferAsset:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L183) | JOIN_ON / P01 / COMPARISON | `sta.Purpose` : `INT` | `par3.ParamValue` : `VARCHAR(30)` | `par3.ParamValue` → `INT` | `sta.Purpose = par3.ParamValue` |

<a id="kolom-trxstatus"></a>
## TrxStatus

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAset_TM.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009954 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptFOPandTransferAsset:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L105) | JOIN_ON / P01 / COMPARISON | `sta.TrxStatus` : `INT` | `par.ParamValue` : `VARCHAR(30)` | `par.ParamValue` → `INT` | `sta.TrxStatus = par.ParamValue` |
| IC-009964 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptFOPandTransferAsset:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L177) | JOIN_ON / P01 / COMPARISON | `sta.TrxStatus` : `INT` | `par.ParamValue` : `VARCHAR(30)` | `par.ParamValue` → `INT` | `sta.TrxStatus = par.ParamValue` |
| IC-012651 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateStatusTransferAsset:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusTransferAsset.sql#L58) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TrxStatus` : `INT` | `@pnNewStatus` : `VARCHAR(15)` | `@pnNewStatus` → `INT` | `TrxStatus = @pnNewStatus` |
| IC-012652 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateStatusTransferAsset:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusTransferAsset.sql#L73) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TrxStatus` : `INT` | `@pnNewStatus` : `VARCHAR(15)` | `@pnNewStatus` → `INT` | `TrxStatus = @pnNewStatus` |
| IC-012653 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateStatusTransferAsset:88](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusTransferAsset.sql#L88) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TrxStatus` : `INT` | `@pnNewStatus` : `VARCHAR(15)` | `@pnNewStatus` → `INT` | `TrxStatus = @pnNewStatus` |
| IC-012654 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateStatusTransferAsset:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusTransferAsset.sql#L103) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TrxStatus` : `INT` | `@pnNewStatus` : `VARCHAR(15)` | `@pnNewStatus` → `INT` | `TrxStatus = @pnNewStatus` |

<a id="kolom-verifynik"></a>
## VerifyNIK

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransferAset_TM.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009953 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L53) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.VerifyNIK` : `INT` | `#tempResult.NIKVerify` : `bigint` | `sta.VerifyNIK` → `bigint` | `#tempResult.NIKVerify <- sta.VerifyNIK` |
| IC-009963 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptFOPandTransferAsset:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptFOPandTransferAsset.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sta.VerifyNIK` : `INT` | `#tempResult.NIKVerify` : `bigint` | `sta.VerifyNIK` → `bigint` | `#tempResult.NIKVerify <- sta.VerifyNIK` |
