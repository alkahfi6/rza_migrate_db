# Retail UAT.dbo.TRSSettlementMaturityLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 31 lokasi source, 6 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-processusernik"></a>
## ProcessUserNik

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementMaturityLog_TH.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009703 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDailyPaymentCouponOrMaturity:536](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L536) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `th.ProcessUserNik` : `INT` | `UserSettlement` : `varchar(50)` | `th.ProcessUserNik` → `varchar(50)` | `UserSettlement = th.ProcessUserNik` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementMaturityLog_TH.sql#L3). **Existing leading index:** TRSSettlementMaturityLog_TH_CN1.

**Frekuensi source:** 9 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009650 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:418](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L418) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009694 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:517](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L517) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009704 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:539](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L539) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009716 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:580](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L580) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009720 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:591](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L591) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009724 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:606](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L606) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009730 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:626](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L626) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009746 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:679](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L679) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009750 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:690](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L690) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementMaturityLog_TH.sql#L12). **Existing leading index:** TRSSettlementMaturityLog_TH_NN1.

**Frekuensi source:** 20 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006295 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementMaturity:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementMaturity.sql#L60) | CASE / P12 / COMPARISON | `smlth.[Status]` : `CHAR(1)` | `'P'` : `varchar` | `smlth.[Status]` → `VARCHAR` | `smlth.[Status] = 'P'` |
| IC-006296 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementMaturity:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementMaturity.sql#L63) | CASE / P12 / COMPARISON | `smlth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `smlth.[Status]` → `VARCHAR` | `smlth.[Status] = 'S'` |
| IC-006297 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementMaturity:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementMaturity.sql#L66) | CASE / P12 / COMPARISON | `smlth.[Status]` : `CHAR(1)` | `'F'` : `varchar` | `smlth.[Status]` → `VARCHAR` | `smlth.[Status] = 'F'` |
| IC-006304 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateSettlementMaturity:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementMaturity.sql#L132) | JOIN_ON / P12 / COMPARISON | `smlth.[Status]` : `CHAR(1)` | `'F'` : `varchar` | `smlth.[Status]` → `VARCHAR` | `smlth.[Status] <> 'F'` |
| IC-007517 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementMaturity:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementMaturity.sql#L194) | JOIN_ON / P12 / COMPARISON | `smlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `smlth.Status` → `VARCHAR` | `smlth.Status <> 'F'` |
| IC-009582 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptCorporateAction:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptCorporateAction.sql#L167) | WHERE / P12 / COMPARISON | `sl.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `sl.[Status]` → `VARCHAR` | `sl.[Status] = 'S'` |
| IC-009697 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:522](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L522) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'S'` : `varchar` | `Status` → `VARCHAR` | `Status = 'S'` |
| IC-009707 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:544](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L544) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'S'` : `varchar` | `Status` → `VARCHAR` | `Status = 'S'` |
| IC-009719 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:585](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L585) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'F'` : `varchar` | `Status` → `VARCHAR` | `Status = 'F'` |
| IC-009723 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:596](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L596) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'S'` : `varchar` | `Status` → `VARCHAR` | `Status = 'S'` |
| IC-009726 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:609](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L609) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'S'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'S'` |
| IC-009733 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:631](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L631) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'F'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'F'` |
| IC-009748 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:682](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L682) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'S'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'S'` |
| IC-009752 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:693](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L693) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'F'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'F'` |
| IC-011540 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettlementMaturity:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementMaturity.sql#L24) | CASE / P12 / COMPARISON | `pbblth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `pbblth.[Status]` → `VARCHAR` | `pbblth.[Status] = 'S'` |
| IC-011541 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettlementMaturity:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementMaturity.sql#L36) | CASE / P12 / COMPARISON | `pbblth.Status` : `CHAR(1)` | `'P'` : `varchar` | `pbblth.Status` → `VARCHAR` | `pbblth.Status = 'P'` |
| IC-011542 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettlementMaturity:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementMaturity.sql#L37) | CASE / P12 / COMPARISON | `pbblth.Status` : `CHAR(1)` | `'S'` : `varchar` | `pbblth.Status` → `VARCHAR` | `pbblth.Status = 'S'` |
| IC-011543 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettlementMaturity:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementMaturity.sql#L38) | CASE / P12 / COMPARISON | `pbblth.Status` : `CHAR(1)` | `'F'` : `varchar` | `pbblth.Status` → `VARCHAR` | `pbblth.Status = 'F'` |
| IC-012546 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L118) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-012593 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSUpdatePostingResult:784](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L784) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `mt.[Status]` : `CHAR(1)` | `''` : `varchar` | `''` → `CHAR(1)` | `ISNULL(mt.[Status],'')` |

<a id="kolom-uniqueid"></a>
## UniqueId

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementMaturityLog_TH.sql#L11). **Existing leading index:** TRSSettlementMaturityLog_TH_NN2.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009731 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDailyPaymentCouponOrMaturity:630](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L630) | JOIN_ON / P13 / COMPARISON | `th.UniqueId` : `UNIQUEIDENTIFIER` | `o.OBLREM3` : `VARCHAR(40)` | `o.OBLREM3` → `UNIQUEIDENTIFIER` | `th.UniqueId = o.OBLREM3` |
