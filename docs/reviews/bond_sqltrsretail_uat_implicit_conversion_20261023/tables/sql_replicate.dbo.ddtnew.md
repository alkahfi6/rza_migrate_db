# SQL_REPLICATE.dbo.DDTNEW — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 39 lokasi source, 14 procedure, 7 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 17 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-acctno"></a>
## ACCTNO

**Tipe dasar:** `DECIMAL(19,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDTNEW.sql#L4). **Existing leading index:** idx_acctno_actype.

**Frekuensi source:** 16 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000551 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.GetRekeningMasterNasabah:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L61) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ACCTNO` : `DECIMAL(19,0)` | `#tmpRek.AccountNo` : `VARCHAR(50)` | `ACCTNO` → `VARCHAR(50)` | `#tmpRek.AccountNo <- ACCTNO` |
| IC-000116 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.CheckRekeningIsMcy:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L28) | WHERE / P04 / COMPARISON | `a.ACCTNO` : `DECIMAL(19,0)` | `@pcRek` : `varchar(50)` | `@pcRek` → `DECIMAL` | `a.ACCTNO = @pcRek` |
| IC-000597 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabahTRS:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L51) | JOIN_ON / P01 / COMPARISON | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000602 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabahTRS:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L78) | JOIN_ON / P01 / COMPARISON | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000607 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabahTRS:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L111) | JOIN_ON / P01 / COMPARISON | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000612 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabahTRS:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L145) | JOIN_ON / P01 / COMPARISON | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000617 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabahTRS:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L180) | JOIN_ON / P01 / COMPARISON | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000622 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabahTRS:207](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L207) | JOIN_ON / P01 / COMPARISON | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000627 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabahTRS:234](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L234) | JOIN_ON / P01 / COMPARISON | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000632 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetRekeningMasterNasabahTRS:261](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L261) | JOIN_ON / P01 / COMPARISON | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` : `VARCHAR(20)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000801 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OMUpdateRegisterNasabah:314](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMUpdateRegisterNasabah.sql#L314) | JOIN_ON / P01 / COMPARISON | `v.ACCTNO` : `DECIMAL(19,0)` | `tmp.NoRekIDR` : `varchar(50)` | `tmp.NoRekIDR` → `DECIMAL` | `tmp.NoRekIDR = v.ACCTNO` |
| IC-001770 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ProTRSReportCOD067:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ProTRSReportCOD067.sql#L143) | JOIN_ON / P11 / COMPARISON | `b.ACCTNO` : `DECIMAL(19,0)` | `a.[NoRekTA]` : `[bigint]` | `a.[NoRekTA]` → `DECIMAL` | `b.ACCTNO=a.[NoRekTA]` |
| IC-002175 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSGetSecurityTransaction_TT:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGetSecurityTransaction_TT.sql#L169) | JOIN_ON / P11 / COMPARISON | `dd.ACCTNO` : `DECIMAL(19,0)` | `CONVERT(NUMERIC, pb.NoRekInvestor)` : `NUMERIC` | `CONVERT(NUMERIC, pb.NoRekInvestor)` → `DECIMAL` | `CONVERT(NUMERIC, pb.NoRekInvestor) = dd.ACCTNO` |
| IC-004155 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGetSecurityTransaction_TT:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetSecurityTransaction_TT.sql#L167) | JOIN_ON / P11 / COMPARISON | `dd.ACCTNO` : `DECIMAL(19,0)` | `CONVERT(NUMERIC, pb.NoRekInvestor)` : `NUMERIC` | `CONVERT(NUMERIC, pb.NoRekInvestor)` → `DECIMAL` | `CONVERT(NUMERIC, pb.NoRekInvestor) = dd.ACCTNO` |
| IC-013611 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningMaterai:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L68) | JOIN_ON / P11 / COMPARISON | `dm.ACCTNO` : `DECIMAL(19,0)` | `t.NoRekInvestor` : `BIGINT` | `t.NoRekInvestor` → `DECIMAL` | `t.NoRekInvestor = dm.ACCTNO` |
| IC-013632 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_PopulateDataRekeningObligasi:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L73) | JOIN_ON / P11 / COMPARISON | `dm.ACCTNO` : `DECIMAL(19,0)` | `t.NoRekInvestor` : `bigint` | `t.NoRekInvestor` → `DECIMAL` | `t.NoRekInvestor = dm.ACCTNO` |

<a id="kolom-actype"></a>
## ACTYPE

**Tipe dasar:** `VARCHAR(1)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDTNEW.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000112 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CheckRekeningIsMcy:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L18) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `a.ACTYPE` : `VARCHAR(1)` | `#tmpRek.ACTYPE` : `CHAR(1)` | `a.ACTYPE` → `CHAR(1)` | `#tmpRek.ACTYPE <- a.ACTYPE` |
| IC-000118 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CheckRekeningIsMcy:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L52) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `a.ACTYPE` : `VARCHAR(1)` | `#tmpListRek.ACTYPE` : `CHAR(1)` | `a.ACTYPE` → `CHAR(1)` | `#tmpListRek.ACTYPE <- a.ACTYPE` |
| IC-002174 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.Rvmp_TRSGetSecurityTransaction_TT:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSGetSecurityTransaction_TT.sql#L166) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `dd.ACTYPE` : `VARCHAR(1)` | `ACTYPE` : `CHAR(1)` | `dd.ACTYPE` → `CHAR(1)` | `ACTYPE = dd.ACTYPE` |
| IC-004154 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGetSecurityTransaction_TT:164](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetSecurityTransaction_TT.sql#L164) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `dd.ACTYPE` : `VARCHAR(1)` | `ACTYPE` : `CHAR(1)` | `dd.ACTYPE` → `CHAR(1)` | `ACTYPE = dd.ACTYPE` |

<a id="kolom-branch"></a>
## BRANCH

**Tipe dasar:** `DECIMAL(5,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDTNEW.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000113 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CheckRekeningIsMcy:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L18) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `a.BRANCH` : `DECIMAL(5,0)` | `#tmpRek.BRANCH` : `INT` | `a.BRANCH` → `INT` | `#tmpRek.BRANCH <- a.BRANCH` |
| IC-000119 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CheckRekeningIsMcy:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L52) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `a.BRANCH` : `DECIMAL(5,0)` | `#tmpListRek.BRANCH` : `INT` | `a.BRANCH` → `INT` | `#tmpListRek.BRANCH <- a.BRANCH` |
| IC-000552 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.GetRekeningMasterNasabah:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L61) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `BRANCH` : `DECIMAL(5,0)` | `#tmpRek.Branch` : `int` | `BRANCH` → `int` | `#tmpRek.Branch <- BRANCH` |
| IC-004245 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSInsertRevenueBondJurnalGL:146](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L146) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `x.[BRANCH]` : `DECIMAL(5,0)` | `z.[KodeCabangRLS]` : `VARCHAR(5)` | `x.[BRANCH]` → `VARCHAR(5)` | `z.[KodeCabangRLS] = x.[BRANCH]` |
| IC-004246 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSInsertRevenueBondJurnalGL:153](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertRevenueBondJurnalGL.sql#L153) | JOIN_ON / P11 / COMPARISON | `x.[BRANCH]` : `DECIMAL(5,0)` | `v.[office_id]` : `INT` | `v.[office_id]` → `DECIMAL` | `x.[BRANCH] = v.[office_id]` |

<a id="kolom-cifno"></a>
## CIFNO

**Tipe dasar:** `DECIMAL(19,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDTNEW.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000513 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountSyariah:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountSyariah.sql#L44) | WHERE / P04 / COMPARISON | `dt.CIFNO` : `DECIMAL(19,0)` | `@pcCIF` : `varchar(19)` | `@pcCIF` → `DECIMAL` | `dt.CIFNO = @pcCIF` |
| IC-000526 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.GetAccountTaxAmnesty:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountTaxAmnesty.sql#L37) | WHERE / P04 / COMPARISON | `dt.CIFNO` : `DECIMAL(19,0)` | `@pcCIF` : `varchar(19)` | `@pcCIF` → `DECIMAL` | `dt.CIFNO = @pcCIF` |
| IC-001769 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.ProTRSReportCOD067:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ProTRSReportCOD067.sql#L143) | JOIN_ON / P11 / COMPARISON | `b.CIFNO` : `DECIMAL(19,0)` | `a.[NoCIF]` : `[bigint]` | `a.[NoCIF]` → `DECIMAL` | `a.[NoCIF] = b.CIFNO` |

<a id="kolom-ddctyp"></a>
## DDCTYP

**Tipe dasar:** `VARCHAR(4)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDTNEW.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000496 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.GetAccountInfo:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L79) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `dt.DDCTYP` : `VARCHAR(4)` | `ta.CurrencyCode` : `CHAR(3)` | `dt.DDCTYP` → `CHAR(3)` | `ta.CurrencyCode = dt.DDCTYP` |
| IC-008638 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSResponseBondFromMX:758](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L758) | JOIN_ON / P12 / COMPARISON | `dt.DDCTYP` : `VARCHAR(4)` | `dd.CURR` : `CHAR(4)` | `dd.CURR` → `VARCHAR` | `dt.DDCTYP = dd.CURR` |
| IC-008649 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSResponseBondFromMX:801](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L801) | JOIN_ON / P12 / COMPARISON | `dt.DDCTYP` : `VARCHAR(4)` | `dd.CURR` : `CHAR(4)` | `dd.CURR` → `VARCHAR` | `dt.DDCTYP = dd.CURR` |

<a id="kolom-sccode"></a>
## SCCODE

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDTNEW.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000114 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CheckRekeningIsMcy:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L18) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `a.SCCODE` : `VARCHAR(10)` | `#tmpRek.SCCODE` : `CHAR(10)` | `a.SCCODE` → `CHAR(10)` | `#tmpRek.SCCODE <- a.SCCODE` |
| IC-000120 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CheckRekeningIsMcy:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L52) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `a.SCCODE` : `VARCHAR(10)` | `#tmpListRek.SCCODE` : `CHAR(10)` | `a.SCCODE` → `CHAR(10)` | `#tmpListRek.SCCODE <- a.SCCODE` |
| IC-000122 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.CheckRekeningIsMcy:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L62) | JOIN_ON / P12 / COMPARISON | `a.SCCODE` : `VARCHAR(10)` | `b.SCCODE` : `CHAR(10)` | `b.SCCODE` → `VARCHAR` | `a.SCCODE = b.SCCODE` |

<a id="kolom-status"></a>
## STATUS

**Tipe dasar:** `DECIMAL(1,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.DDTNEW.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000115 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CheckRekeningIsMcy:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L18) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `a.[STATUS]` : `DECIMAL(1,0)` | `#tmpRek.STATUS` : `INT` | `a.[STATUS]` → `INT` | `#tmpRek.STATUS <- a.[STATUS]` |
| IC-000121 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.CheckRekeningIsMcy:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckRekeningIsMcy.sql#L52) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `a.[STATUS]` : `DECIMAL(1,0)` | `#tmpListRek.STATUS` : `INT` | `a.[STATUS]` → `INT` | `#tmpListRek.STATUS <- a.[STATUS]` |
| IC-000553 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.GetRekeningMasterNasabah:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabah.sql#L61) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `STATUS` : `DECIMAL(1,0)` | `#tmpRek.STATUS` : `int` | `STATUS` → `int` | `#tmpRek.STATUS <- STATUS` |
| IC-013610 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningMaterai:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningMaterai.sql#L65) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `dm.STATUS` : `DECIMAL(1,0)` | `STAT` : `CHAR(1)` | `dm.STATUS` → `CHAR(1)` | `STAT = dm.STATUS` |
| IC-013631 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_PopulateDataRekeningObligasi:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_PopulateDataRekeningObligasi.sql#L70) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `dm.STATUS` : `DECIMAL(1,0)` | `STAT` : `char(1)` | `dm.STATUS` → `char(1)` | `STAT = dm.STATUS` |
