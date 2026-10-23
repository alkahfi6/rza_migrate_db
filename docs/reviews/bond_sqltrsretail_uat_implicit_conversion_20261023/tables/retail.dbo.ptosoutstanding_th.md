# Retail UAT.dbo.PTOSOutstanding_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 39 lokasi source, 4 procedure, 13 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001307 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptosoutstanding_th.AccruedInterest <- AccruedInterest` |
| IC-001412 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptosoutstanding_th.AccruedInterest <- AccruedInterest` |

<a id="kolom-actionflag"></a>
## ActionFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001278 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L90) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- 'O'` |
| IC-001308 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `case when [Status] = 1 then 'N' else 'I' end` : `VARCHAR (combined result; facets unknown)` | `case when [Status] = 1 then 'N' else 'I' end` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- case when [Status] = 1 then 'N' else 'I' end` |
| IC-001383 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- 'O'` |
| IC-001413 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `case when [Status] = 1 then 'N' else 'I' end` : `VARCHAR (combined result; facets unknown)` | `case when [Status] = 1 then 'N' else 'I' end` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- case when [Status] = 1 then 'N' else 'I' end` |
| IC-007119 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOutsAuthen:95](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOutsAuthen.sql#L95) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- 'D'` |
| IC-007121 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOutsAuthen:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOutsAuthen.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- 'O'` |
| IC-007123 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOutsAuthen:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOutsAuthen.sql#L193) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- 'N'` |
| IC-007125 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSOutsAuthen:249](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSOutsAuthen.sql#L249) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- 'I'` |
| IC-007217 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:310](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L310) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- 'I'` |
| IC-007219 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:356](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L356) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- 'O'` |
| IC-007221 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:411](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L411) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ActionFlag <- 'N'` |

<a id="kolom-approveflag"></a>
## ApproveFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001279 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L90) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ApproveFlag <- 'A'` |
| IC-001309 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ApproveFlag <- 'A'` |
| IC-001384 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ApproveFlag <- 'A'` |
| IC-001414 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptosoutstanding_th.ApproveFlag <- 'A'` |

<a id="kolom-ccy"></a>
## Ccy

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001310 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptosoutstanding_th.Ccy <- Ccy` |
| IC-001415 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptosoutstanding_th.Ccy <- Ccy` |

<a id="kolom-hargaperolehannasabah"></a>
## HargaPerolehanNasabah

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001311 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `HargaPerolehanNasabah` : `FLOAT` | `HargaPerolehanNasabah` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_th.HargaPerolehanNasabah <- HargaPerolehanNasabah` |
| IC-001416 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `HargaPerolehanNasabah` : `FLOAT` | `HargaPerolehanNasabah` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_th.HargaPerolehanNasabah <- HargaPerolehanNasabah` |

<a id="kolom-indikasihargabid"></a>
## IndikasiHargaBid

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001312 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.IndikasiHargaBid` : `DECIMAL(20,5)` | `IndikasiHargaBid` : `FLOAT` | `IndikasiHargaBid` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_th.IndikasiHargaBid <- IndikasiHargaBid` |
| IC-001417 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.IndikasiHargaBid` : `DECIMAL(20,5)` | `IndikasiHargaBid` : `FLOAT` | `IndikasiHargaBid` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_th.IndikasiHargaBid <- IndikasiHargaBid` |

<a id="kolom-kodeproduk"></a>
## KodeProduk

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001313 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptosoutstanding_th.KodeProduk <- KodeProduk` |
| IC-001418 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptosoutstanding_th.KodeProduk <- KodeProduk` |

<a id="kolom-namanasabah"></a>
## NamaNasabah

**Tipe dasar:** `VARCHAR(95)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001314 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.NamaNasabah` : `VARCHAR(95)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptosoutstanding_th.NamaNasabah <- NamaNasabah` |
| IC-001419 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.NamaNasabah` : `VARCHAR(95)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptosoutstanding_th.NamaNasabah <- NamaNasabah` |

<a id="kolom-namaproduk"></a>
## NamaProduk

**Tipe dasar:** `VARCHAR(70)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001315 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptosoutstanding_th.NamaProduk <- NamaProduk` |
| IC-001420 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptosoutstanding_th.NamaProduk <- NamaProduk` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001316 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptosoutstanding_th.Nominal <- Nominal` |
| IC-001421 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptosoutstanding_th.Nominal <- Nominal` |

<a id="kolom-ptostype"></a>
## PTOSType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001317 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.PTOSType` : `SMALLINT` | `PTOSType` : `int` | `PTOSType` → `SMALLINT` | `retail.dbo.ptosoutstanding_th.PTOSType <- PTOSType` |
| IC-001422 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.PTOSType` : `SMALLINT` | `PTOSType` : `int` | `PTOSType` → `SMALLINT` | `retail.dbo.ptosoutstanding_th.PTOSType <- PTOSType` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001318 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptosoutstanding_th.SID <- [SID]` |
| IC-001423 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptosoutstanding_th.SID <- [SID]` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_TH.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001280 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L90) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.Status` : `SMALLINT` | `tmp.[Status]` : `INT (combined result; facets unknown)` | `tmp.[Status]` → `SMALLINT` | `retail.dbo.ptosoutstanding_th.Status <- tmp.[Status]` |
| IC-001319 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L137) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.Status` : `SMALLINT` | `[Status]` : `INT (combined result; facets unknown)` | `[Status]` → `SMALLINT` | `retail.dbo.ptosoutstanding_th.Status <- [Status]` |
| IC-001385 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L127) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.Status` : `SMALLINT` | `tmp.[Status]` : `INT (combined result; facets unknown)` | `tmp.[Status]` → `SMALLINT` | `retail.dbo.ptosoutstanding_th.Status <- tmp.[Status]` |
| IC-001424 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:177](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L177) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_th.Status` : `SMALLINT` | `[Status]` : `INT (combined result; facets unknown)` | `[Status]` → `SMALLINT` | `retail.dbo.ptosoutstanding_th.Status <- [Status]` |
