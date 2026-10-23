# SQL_REPLICATE.dbo.BIS4_SUB_ACCOUNTS — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 30 lokasi source, 9 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 3 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-code"></a>
## code

**Tipe dasar:** `NVARCHAR(35)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.BIS4_SUB_ACCOUNTS.sql#L4). **Existing leading index:** BIS4_SUB_ACCOUNTS_NN1.

**Frekuensi source:** 10 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.
- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006020 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateSettlementBIS4:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L21) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `SubAccCode` : `NVARCHAR(35)` | `@pnCashAccountBKON` : `varchar(50)` | `SubAccCode` → `varchar(50)` | `@pnCashAccountBKON = SubAccCode` |
| IC-007402 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessSettlementBIS4:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBIS4.sql#L97) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sub.SubAccCode` : `NVARCHAR(35)` | `tmp.SubAccountBuyer` : `varchar(500)` | `sub.SubAccCode` → `varchar(500)` | `tmp.SubAccountBuyer = sub.SubAccCode` |
| IC-007405 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessSettlementBIS4:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBIS4.sql#L106) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `sub.SubAccCode` : `NVARCHAR(35)` | `tmp.SubAccountSeller` : `varchar(500)` | `sub.SubAccCode` → `varchar(500)` | `tmp.SubAccountSeller = sub.SubAccCode` |
| IC-005337 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateCashAccountsParam:2](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateCashAccountsParam.sql#L2) | SET_OUTPUT / P07 / UNION_OUTPUT | `right output #1` : `NVARCHAR(35)` | `left output #1` : `varchar` | `left output #1` → `NVARCHAR` | `select '' as 'Value', '' as 'Display' union all select SubAccCode as 'Value' ,SubAccCode + ' - ' + SubAccName as 'Display' from dbo.BIS4SubAccounts_v where SubAccTypeId = 2 order by Value` |
| IC-005338 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateCashAccountsParam:6](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateCashAccountsParam.sql#L6) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `SubAccCode` : `NVARCHAR(35)` | `' - '` : `varchar` | `' - '` → `NVARCHAR` | `SubAccCode + ' - '` |
| IC-005968 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePaymentAgentBIS4:8](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePaymentAgentBIS4.sql#L8) | WHERE / P04 / COMPARISON | `subAcc.SubAccCode` : `NVARCHAR(35)` | `@pnCashAccount` : `varchar(50)` | `@pnCashAccount` → `NVARCHAR` | `subAcc.SubAccCode = @pnCashAccount` |
| IC-005970 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePaymentAgentBIS4:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePaymentAgentBIS4.sql#L17) | WHERE / P04 / COMPARISON | `subAcc.SubAccCode` : `NVARCHAR(35)` | `@pnCashAccount` : `varchar(50)` | `@pnCashAccount` → `NVARCHAR` | `subAcc.SubAccCode = @pnCashAccount` |
| IC-012114 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchCashAccount:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchCashAccount.sql#L17) | WHERE / P04 / COMPARISON | `SubAccCode` : `NVARCHAR(35)` | `@cCol1` : `varchar(40)` | `@cCol1` → `NVARCHAR` | `SubAccCode = @cCol1` |
| IC-012116 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchCashAccount:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchCashAccount.sql#L24) | WHERE / P06 / LIKE | `SubAccCode` : `NVARCHAR(35)` | `@cCol1 + '%'` : `VARCHAR (combined result; facets unknown)` | `@cCol1 + '%'` → `NVARCHAR` | `SubAccCode like @cCol1 + '%'` |
| IC-012122 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchCashAccount:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchCashAccount.sql#L37) | WHERE / P06 / LIKE | `SubAccCode` : `NVARCHAR(35)` | `@cCol1 + '%'` : `VARCHAR (combined result; facets unknown)` | `@cCol1 + '%'` → `NVARCHAR` | `SubAccCode like @cCol1 + '%'` |

<a id="kolom-name"></a>
## name

**Tipe dasar:** `NVARCHAR(120)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.BIS4_SUB_ACCOUNTS.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006021 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L23) | WHERE / P04 / COMPARISON | `SubAccName` : `NVARCHAR(120)` | `'KSEI IDR Cash Account'` : `varchar` | `'KSEI IDR Cash Account'` → `NVARCHAR` | `SubAccName = 'KSEI IDR Cash Account'` |
| IC-012118 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchCashAccount:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchCashAccount.sql#L25) | WHERE / P06 / LIKE | `upper(SubAccName)` : `NVARCHAR (expression; length unknown)` | `@cCol2 + '%'` : `VARCHAR (combined result; facets unknown)` | `@cCol2 + '%'` → `NVARCHAR` | `upper(SubAccName) like @cCol2 + '%'` |
| IC-012120 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchCashAccount:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchCashAccount.sql#L31) | WHERE / P06 / LIKE | `upper(SubAccName)` : `NVARCHAR (expression; length unknown)` | `@cCol2 + '%'` : `VARCHAR (combined result; facets unknown)` | `@cCol2 + '%'` → `NVARCHAR` | `upper(SubAccName) like @cCol2 + '%'` |

<a id="kolom-sub-acc-type-id"></a>
## sub_acc_type_id

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.BIS4_SUB_ACCOUNTS.sql#L6). **Existing leading index:** BIS4_SUB_ACCOUNTS_NN2.

**Frekuensi source:** 17 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005211 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateAccSubAccBIS4:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateAccSubAccBIS4.sql#L13) | WHERE / P11 / COMPARISON | `subAcc.SubAccTypeId` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `subAcc.SubAccTypeId = 1` |
| IC-005339 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateCashAccountsParam:8](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateCashAccountsParam.sql#L8) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `SubAccTypeId = 2` |
| IC-005967 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateParticipantParamBIS4:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateParticipantParamBIS4.sql#L9) | JOIN_ON / P11 / COMPARISON | `subAcc.SubAccTypeId` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `subAcc.SubAccTypeId = 1` |
| IC-005969 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePaymentAgentBIS4:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePaymentAgentBIS4.sql#L9) | WHERE / P11 / COMPARISON | `subAcc.SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `subAcc.SubAccTypeId = 2` |
| IC-005971 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePaymentAgentBIS4:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePaymentAgentBIS4.sql#L18) | WHERE / P11 / COMPARISON | `subAcc.SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `subAcc.SubAccTypeId = 2` |
| IC-006022 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L24) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `SubAccTypeId = 2` |
| IC-006029 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:158](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L158) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `SubAccTypeId = 1` |
| IC-006030 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L162) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `SubAccTypeId = 2` |
| IC-006178 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4Detail:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4Detail.sql#L15) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `SubAccTypeId = 2` |
| IC-006184 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementBIS4RepoDetail:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4RepoDetail.sql#L15) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `SubAccTypeId = 2` |
| IC-007404 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessSettlementBIS4:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBIS4.sql#L103) | JOIN_ON / P11 / COMPARISON | `sub.SubAccTypeId` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `sub.SubAccTypeId = 1` |
| IC-007407 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSProcessSettlementBIS4:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementBIS4.sql#L112) | JOIN_ON / P11 / COMPARISON | `sub.SubAccTypeId` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `sub.SubAccTypeId = 1` |
| IC-012115 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchCashAccount:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchCashAccount.sql#L17) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `SubAccTypeId = 2` |
| IC-012117 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchCashAccount:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchCashAccount.sql#L25) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `SubAccTypeId = 2` |
| IC-012119 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchCashAccount:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchCashAccount.sql#L31) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `SubAccTypeId = 2` |
| IC-012121 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchCashAccount:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchCashAccount.sql#L37) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `SubAccTypeId = 2` |
| IC-012123 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSearchCashAccount:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchCashAccount.sql#L43) | WHERE / P11 / COMPARISON | `SubAccTypeId` : `BIGINT` | `2` : `int` | `2` → `BIGINT` | `SubAccTypeId = 2` |
