# Retail UAT.dbo.TRSCustAccNumber_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 17 lokasi source, 8 procedure, 5 kolom. **A:** 6 lokasi pada 3 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-norekinvestor"></a>
## NoRekInvestor

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustAccNumber_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 4 SP. **A:** 6 lokasi / 3 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000601 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L76) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `b.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = b.ACCTNO` |
| IC-000602 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L78) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000621 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L205) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `b.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = b.ACCTNO` |
| IC-000622 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:207](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L207) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-004441 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSListEarlyRedemption:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSListEarlyRedemption.sql#L109) | JOIN_ON / P01 / COMPARISON | `tcantr.NoRekInvestor` : `VARCHAR(20)` | `DDMAST.ACCTNO` : `DECIMAL(19,0)` | `tcantr.NoRekInvestor` → `DECIMAL` | `tcantr.NoRekInvestor = DDMAST.ACCTNO` |
| IC-013564 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.trs_ListTreasuryCustomer_TM_Original:398](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_ListTreasuryCustomer_TM_Original.sql#L398) | JOIN_ON / P01 / COMPARISON | `tcantr.NoRekInvestor` : `VARCHAR(20)` | `DDMAST.ACCTNO` : `DECIMAL(19,0)` | `tcantr.NoRekInvestor` → `DECIMAL` | `tcantr.NoRekInvestor = DDMAST.ACCTNO` |
| IC-008628 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:718](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L718) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tcan.NoRekInvestor` : `VARCHAR(20)` | `fm.NoRekInvestor` : `BIGINT` | `fm.NoRekInvestor` → `VARCHAR(20)` | `tcan.NoRekInvestor = fm.NoRekInvestor` |
| IC-008632 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:733](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L733) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumber_tr.NoRekInvestor` : `VARCHAR(20)` | `fm.NoRekInvestor` : `BIGINT` | `fm.NoRekInvestor` → `VARCHAR(20)` | `retail.dbo.trscustaccnumber_tr.NoRekInvestor <- fm.NoRekInvestor` |

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustAccNumber_TR.sql#L3). **Existing leading index:** TRSCustAccNumber_TR_NN1&#124;PK_TRSCustAccNumber_TR.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002866 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L127) | WHERE / P11 / COMPARISON | `tr.CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `tr.CIFId = @nCIFId` |
| IC-002874 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:416](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L416) | WHERE / P11 / COMPARISON | `tr.CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `tr.CIFId = @nCIFId` |
| IC-002883 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessRekDana:604](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L604) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumber_tr.CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `retail.dbo.trscustaccnumber_tr.CIFId <- @nCIFId` |

<a id="kolom-isjoin"></a>
## IsJoin

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustAccNumber_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000730 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMSubmitTransactionSekunder:265](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMSubmitTransactionSekunder.sql#L265) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumber_tr.IsJoin` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumber_tr.IsJoin <- 'N'` |
| IC-008631 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:733](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L733) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumber_tr.IsJoin` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumber_tr.IsJoin <- 'N'` |
| IC-012665 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateTaxAmnestyExpired:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L274) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumber_tr.IsJoin` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.trscustaccnumber_tr.IsJoin <- 'N'` |
| IC-013176 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_InsertTreasuryCustomer_TM:854](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTreasuryCustomer_TM.sql#L854) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumber_tr.IsJoin` : `CHAR(1)` | `STATUS_JOIN` : `varchar(1 )` | `STATUS_JOIN` → `CHAR(1)` | `retail.dbo.trscustaccnumber_tr.IsJoin <- STATUS_JOIN` |

<a id="kolom-productcode"></a>
## ProductCode

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustAccNumber_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012666 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateTaxAmnestyExpired:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L274) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumber_tr.ProductCode` : `VARCHAR(10)` | `isnull(d.SCCODE, '')` : `CHAR(10)` | `isnull(d.SCCODE, '')` → `VARCHAR(10)` | `retail.dbo.trscustaccnumber_tr.ProductCode <- isnull(d.SCCODE, '')` |

<a id="kolom-relasi"></a>
## Relasi

**Tipe dasar:** `VARCHAR(2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustAccNumber_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012667 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdateTaxAmnestyExpired:274](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateTaxAmnestyExpired.sql#L274) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscustaccnumber_tr.Relasi` : `VARCHAR(2)` | `isnull(c.CFRELA, '')` : `CHAR(2)` | `isnull(c.CFRELA, '')` → `VARCHAR(2)` | `retail.dbo.trscustaccnumber_tr.Relasi <- isnull(c.CFRELA, '')` |
