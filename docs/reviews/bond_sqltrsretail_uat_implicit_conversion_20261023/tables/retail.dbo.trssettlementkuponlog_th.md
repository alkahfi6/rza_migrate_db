# Retail UAT.dbo.TRSSettlementKuponLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 36 lokasi source, 8 procedure, 9 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-debitrekeninggnc"></a>
## DebitRekeningGNC

**Tipe dasar:** `VARCHAR(12)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementKuponLog_TH.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009561 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptBatchTransaction:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptBatchTransaction.sql#L63) | JOIN_ON / P01 / COMPARISON | `tsklth.DebitRekeningGNC` : `VARCHAR(12)` | `dvd.ACCTNO` : `DECIMAL(19,0)` | `tsklth.DebitRekeningGNC` → `DECIMAL` | `tsklth.DebitRekeningGNC = dvd.ACCTNO` |

<a id="kolom-coupondate"></a>
## CouponDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementKuponLog_TH.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011531 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSettlementKupon:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementKupon.sql#L69) | WHERE / P04 / COMPARISON | `stl.CouponDate` : `DATETIME` | `@pdPeriodFrom` : `varchar(10)` | `@pdPeriodFrom` → `DATETIME` | `stl.CouponDate >= @pdPeriodFrom` |
| IC-011532 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSettlementKupon:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementKupon.sql#L69) | WHERE / P04 / COMPARISON | `stl.CouponDate` : `DATETIME` | `@pdPeriodTo` : `varchar(10)` | `@pdPeriodTo` → `DATETIME` | `stl.CouponDate <= @pdPeriodTo` |
| IC-011535 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSettlementKupon:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementKupon.sql#L77) | WHERE / P04 / COMPARISON | `stl.CouponDate` : `DATETIME` | `@pdPeriodFrom` : `varchar(10)` | `@pdPeriodFrom` → `DATETIME` | `stl.CouponDate >= @pdPeriodFrom` |
| IC-011536 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptSettlementKupon:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementKupon.sql#L77) | WHERE / P04 / COMPARISON | `stl.CouponDate` : `DATETIME` | `@pdPeriodTo` : `varchar(10)` | `@pdPeriodTo` → `DATETIME` | `stl.CouponDate <= @pdPeriodTo` |
| IC-012591 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:775](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L775) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementkuponlog_th.CouponDate` : `DATETIME` | `CONVERT(VARCHAR(8),GETDATE(),112)` : `VARCHAR(8)` | `CONVERT(VARCHAR(8),GETDATE(),112)` → `DATETIME` | `retail.dbo.trssettlementkuponlog_th.CouponDate <- CONVERT(VARCHAR(8),GETDATE(),112)` |

<a id="kolom-nominalkupon"></a>
## NominalKupon

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementKuponLog_TH.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009558 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptBatchTransaction:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptBatchTransaction.sql#L44) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tsklth.NominalKupon` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tsklth.NominalKupon, 0)` |

<a id="kolom-processdate"></a>
## ProcessDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementKuponLog_TH.sql#L9). **Existing leading index:** TRSSettlementKuponLog_TH_CN1.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006274 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementKupon:374](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementKupon.sql#L374) | WHERE / P13 / COMPARISON | `settlementLog.ProcessDate` : `DATETIME` | `@dCurrentDate` : `DATE` | `@dCurrentDate` → `DATETIME` | `settlementLog.ProcessDate >= @dCurrentDate` |
| IC-006275 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateSettlementKupon:375](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementKupon.sql#L375) | WHERE / P13 / COMPARISON | `settlementLog.ProcessDate` : `DATETIME` | `DATEADD(DAY, 1, @dCurrentDate)` : `DATE` | `DATEADD(DAY, 1, @dCurrentDate)` → `DATETIME` | `settlementLog.ProcessDate < DATEADD(DAY, 1, @dCurrentDate)` |

<a id="kolom-processusernik"></a>
## ProcessUserNik

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementKuponLog_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009698 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptDailyPaymentCouponOrMaturity:525](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L525) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `th.ProcessUserNik` : `INT` | `UserSettlement` : `varchar(50)` | `th.ProcessUserNik` → `varchar(50)` | `UserSettlement = th.ProcessUserNik` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementKuponLog_TH.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009690 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:506](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L506) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009699 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:528](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L528) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009708 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:558](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L558) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009712 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:569](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L569) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009727 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:616](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L616) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009734 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:638](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L638) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009738 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:657](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L657) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-009742 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:668](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L668) | JOIN_ON / P11 / COMPARISON | `th.SecId` : `INT` | `tmp.SecId` : `bigint` | `th.SecId` → `BIGINT` | `th.SecId = tmp.SecId` |
| IC-011533 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettlementKupon:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementKupon.sql#L70) | WHERE / P11 / COMPARISON | `stl.[SecId]` : `INT` | `@nSecId` : `bigint` | `stl.[SecId]` → `BIGINT` | `stl.[SecId] = @nSecId` |
| IC-011537 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptSettlementKupon:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementKupon.sql#L78) | WHERE / P11 / COMPARISON | `stl.[SecId]` : `INT` | `@nSecId` : `bigint` | `stl.[SecId]` → `BIGINT` | `stl.[SecId] = @nSecId` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementKuponLog_TH.sql#L11). **Existing leading index:** TRSSettlementKuponLog_TH_NN2.

**Frekuensi source:** 14 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011534 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptSettlementKupon:72](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptSettlementKupon.sql#L72) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `stl.[Status]` : `CHAR(1)` | `@tmp_Staging.StatusCode` : `varchar(10)` | `stl.[Status]` → `varchar(10)` | `@tmp_Staging.StatusCode <- stl.[Status]` |
| IC-006970 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessKuponUploadAuthorization:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessKuponUploadAuthorization.sql#L114) | JOIN_ON / P12 / COMPARISON | `sklth.Status` : `CHAR(1)` | `'F'` : `varchar` | `sklth.Status` → `VARCHAR` | `sklth.Status <> 'F'` |
| IC-007498 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementKupon:221](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementKupon.sql#L221) | JOIN_ON / P12 / COMPARISON | `smlth.Status` : `CHAR(1)` | `'F'` : `varchar` | `smlth.Status` → `VARCHAR` | `smlth.Status <> 'F'` |
| IC-009564 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptBatchTransaction:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptBatchTransaction.sql#L67) | WHERE / P12 / COMPARISON | `tsklth.[Status]` : `CHAR(1)` | `'S'` : `varchar` | `tsklth.[Status]` → `VARCHAR` | `tsklth.[Status] = 'S'` |
| IC-009693 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:511](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L511) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'S'` : `varchar` | `Status` → `VARCHAR` | `Status = 'S'` |
| IC-009702 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:533](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L533) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'S'` : `varchar` | `Status` → `VARCHAR` | `Status = 'S'` |
| IC-009711 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:563](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L563) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'F'` : `varchar` | `Status` → `VARCHAR` | `Status = 'F'` |
| IC-009715 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:574](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L574) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'S'` : `varchar` | `Status` → `VARCHAR` | `Status = 'S'` |
| IC-009729 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:619](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L619) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'S'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'S'` |
| IC-009737 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:643](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L643) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'F'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'F'` |
| IC-009740 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:660](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L660) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'S'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'S'` |
| IC-009744 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptDailyPaymentCouponOrMaturity:671](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L671) | WHERE / P12 / COMPARISON | `th.Status` : `CHAR(1)` | `'F'` : `varchar` | `th.Status` → `VARCHAR` | `th.Status = 'F'` |
| IC-010624 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRptPajakKupon:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPajakKupon.sql#L124) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'S'` : `varchar` | `Status` → `VARCHAR` | `Status = 'S'` |
| IC-012556 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L178) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |

<a id="kolom-taxamnesty"></a>
## TaxAmnesty

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementKuponLog_TH.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012592 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:775](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L775) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trssettlementkuponlog_th.TaxAmnesty` : `VARCHAR(20)` | `0` : `int` | `0` → `VARCHAR(20)` | `retail.dbo.trssettlementkuponlog_th.TaxAmnesty <- 0` |

<a id="kolom-uniqueid"></a>
## UniqueId

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSettlementKuponLog_TH.sql#L10). **Existing leading index:** TRSSettlementKuponLog_TH_NN1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-009735 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptDailyPaymentCouponOrMaturity:642](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptDailyPaymentCouponOrMaturity.sql#L642) | JOIN_ON / P13 / COMPARISON | `th.UniqueId` : `UNIQUEIDENTIFIER` | `o.OBLREM3` : `VARCHAR(40)` | `o.OBLREM3` → `UNIQUEIDENTIFIER` | `th.UniqueId = o.OBLREM3` |
