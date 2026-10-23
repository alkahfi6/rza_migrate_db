# Retail UAT.dbo.PTOSTransaksi_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 53 lokasi source, 4 procedure, 21 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001537 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptostransaksi_th.AccruedInterest <- AccruedInterest` |
| IC-001665 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptostransaksi_th.AccruedInterest <- AccruedInterest` |

<a id="kolom-actionflag"></a>
## ActionFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L32). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001494 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L93) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- 'O'` |
| IC-001538 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `case when [Status] = 1 then 'N' else 'I' end` : `VARCHAR (combined result; facets unknown)` | `case when [Status] = 1 then 'N' else 'I' end` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- case when [Status] = 1 then 'N' else 'I' end` |
| IC-001622 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L92) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- 'O'` |
| IC-001666 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `case when [Status] = 1 then 'N' else 'I' end` : `VARCHAR (combined result; facets unknown)` | `case when [Status] = 1 then 'N' else 'I' end` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- case when [Status] = 1 then 'N' else 'I' end` |
| IC-007172 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSTrxAuthen:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrxAuthen.sql#L125) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- 'D'` |
| IC-007174 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSTrxAuthen:191](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrxAuthen.sql#L191) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- 'O'` |
| IC-007176 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSTrxAuthen:260](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrxAuthen.sql#L260) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- 'N'` |
| IC-007178 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSTrxAuthen:333](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrxAuthen.sql#L333) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- 'I'` |
| IC-007234 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:674](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L674) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- 'I'` |
| IC-007236 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:740](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L740) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `'O'` : `varchar` | `'O'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- 'O'` |
| IC-007238 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:809](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L809) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ActionFlag` : `CHAR(1)` | `'N'` : `varchar` | `'N'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ActionFlag <- 'N'` |

<a id="kolom-approveflag"></a>
## ApproveFlag

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L33). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001495 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:93](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L93) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ApproveFlag <- 'A'` |
| IC-001539 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ApproveFlag <- 'A'` |
| IC-001623 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L92) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ApproveFlag <- 'A'` |
| IC-001667 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.ApproveFlag` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `retail.dbo.ptostransaksi_th.ApproveFlag <- 'A'` |

<a id="kolom-ccy"></a>
## Ccy

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001540 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptostransaksi_th.Ccy <- Ccy` |
| IC-001668 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptostransaksi_th.Ccy <- Ccy` |

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001541 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.DealId` : `VARCHAR(20)` | `DealId` : `NVARCHAR(255)` | `DealId` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_th.DealId <- DealId` |
| IC-001669 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.DealId` : `VARCHAR(20)` | `DealId` : `NVARCHAR(255)` | `DealId` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_th.DealId <- DealId` |

<a id="kolom-dealprice"></a>
## DealPrice

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001542 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.DealPrice` : `DECIMAL(20,5)` | `DealPrice` : `FLOAT` | `DealPrice` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_th.DealPrice <- DealPrice` |
| IC-001670 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.DealPrice` : `DECIMAL(20,5)` | `DealPrice` : `FLOAT` | `DealPrice` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_th.DealPrice <- DealPrice` |

<a id="kolom-gainloss"></a>
## GainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001543 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.GainLoss` : `MONEY` | `GainLoss` : `FLOAT` | `GainLoss` → `MONEY` | `retail.dbo.ptostransaksi_th.GainLoss <- GainLoss` |
| IC-001671 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.GainLoss` : `MONEY` | `GainLoss` : `FLOAT` | `GainLoss` → `MONEY` | `retail.dbo.ptostransaksi_th.GainLoss <- GainLoss` |

<a id="kolom-hargaperolehan"></a>
## HargaPerolehan

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001544 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.HargaPerolehan` : `DECIMAL(20,5)` | `HargaPerolehan` : `FLOAT` | `HargaPerolehan` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_th.HargaPerolehan <- HargaPerolehan` |
| IC-001672 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.HargaPerolehan` : `DECIMAL(20,5)` | `HargaPerolehan` : `FLOAT` | `HargaPerolehan` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_th.HargaPerolehan <- HargaPerolehan` |

<a id="kolom-kodeproduk"></a>
## KodeProduk

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001545 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_th.KodeProduk <- KodeProduk` |
| IC-001673 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR (expression; length unknown)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_th.KodeProduk <- KodeProduk` |

<a id="kolom-npwp"></a>
## NPWP

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001546 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.NPWP` : `VARCHAR(40)` | `NPWP` : `NVARCHAR(255)` | `NPWP` → `VARCHAR(40)` | `retail.dbo.ptostransaksi_th.NPWP <- NPWP` |
| IC-001674 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.NPWP` : `VARCHAR(40)` | `NPWP` : `NVARCHAR(255)` | `NPWP` → `VARCHAR(40)` | `retail.dbo.ptostransaksi_th.NPWP <- NPWP` |

<a id="kolom-namanasabah"></a>
## NamaNasabah

**Tipe dasar:** `VARCHAR(95)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001547 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.NamaNasabah` : `VARCHAR(95)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptostransaksi_th.NamaNasabah <- NamaNasabah` |
| IC-001675 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.NamaNasabah` : `VARCHAR(95)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptostransaksi_th.NamaNasabah <- NamaNasabah` |

<a id="kolom-namaproduk"></a>
## NamaProduk

**Tipe dasar:** `VARCHAR(70)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001548 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptostransaksi_th.NamaProduk <- NamaProduk` |
| IC-001676 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptostransaksi_th.NamaProduk <- NamaProduk` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001549 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptostransaksi_th.Nominal <- Nominal` |
| IC-001677 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptostransaksi_th.Nominal <- Nominal` |

<a id="kolom-ptostype"></a>
## PTOSType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001550 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.PTOSType` : `SMALLINT` | `PTOSType` : `int` | `PTOSType` → `SMALLINT` | `retail.dbo.ptostransaksi_th.PTOSType <- PTOSType` |
| IC-001678 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.PTOSType` : `SMALLINT` | `PTOSType` : `int` | `PTOSType` → `SMALLINT` | `retail.dbo.ptostransaksi_th.PTOSType <- PTOSType` |

<a id="kolom-proceed"></a>
## Proceed

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001551 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.Proceed` : `MONEY` | `Proceed` : `FLOAT` | `Proceed` → `MONEY` | `retail.dbo.ptostransaksi_th.Proceed <- Proceed` |
| IC-001679 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.Proceed` : `MONEY` | `Proceed` : `FLOAT` | `Proceed` → `MONEY` | `retail.dbo.ptostransaksi_th.Proceed <- Proceed` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001552 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptostransaksi_th.SID <- [SID]` |
| IC-001680 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR (expression; length unknown)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptostransaksi_th.SID <- [SID]` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001553 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.Status` : `SMALLINT` | `[Status]` : `INT (combined result; facets unknown)` | `[Status]` → `SMALLINT` | `retail.dbo.ptostransaksi_th.Status <- [Status]` |
| IC-001681 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.Status` : `SMALLINT` | `[Status]` : `INT (combined result; facets unknown)` | `[Status]` → `SMALLINT` | `retail.dbo.ptostransaksi_th.Status <- [Status]` |

<a id="kolom-taxcapitalgainloss"></a>
## TaxCapitalGainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001554 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.TaxCapitalGainLoss` : `MONEY` | `TaxCapitalGainLoss` : `FLOAT` | `TaxCapitalGainLoss` → `MONEY` | `retail.dbo.ptostransaksi_th.TaxCapitalGainLoss <- TaxCapitalGainLoss` |
| IC-001682 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.TaxCapitalGainLoss` : `MONEY` | `TaxCapitalGainLoss` : `FLOAT` | `TaxCapitalGainLoss` → `MONEY` | `retail.dbo.ptostransaksi_th.TaxCapitalGainLoss <- TaxCapitalGainLoss` |

<a id="kolom-taxonaccrued"></a>
## TaxOnAccrued

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001555 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.TaxOnAccrued` : `MONEY` | `TaxOnAccrued` : `FLOAT` | `TaxOnAccrued` → `MONEY` | `retail.dbo.ptostransaksi_th.TaxOnAccrued <- TaxOnAccrued` |
| IC-001683 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.TaxOnAccrued` : `MONEY` | `TaxOnAccrued` : `FLOAT` | `TaxOnAccrued` → `MONEY` | `retail.dbo.ptostransaksi_th.TaxOnAccrued <- TaxOnAccrued` |

<a id="kolom-totalproceed"></a>
## TotalProceed

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001556 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.TotalProceed` : `MONEY` | `TotalProceed` : `FLOAT` | `TotalProceed` → `MONEY` | `retail.dbo.ptostransaksi_th.TotalProceed <- TotalProceed` |
| IC-001684 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.TotalProceed` : `MONEY` | `TotalProceed` : `FLOAT` | `TotalProceed` → `MONEY` | `retail.dbo.ptostransaksi_th.TotalProceed <- TotalProceed` |

<a id="kolom-trxtype"></a>
## TrxType

**Tipe dasar:** `VARCHAR(8)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001557 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.TrxType` : `VARCHAR(8)` | `TrxType` : `NVARCHAR(255)` | `TrxType` → `VARCHAR(8)` | `retail.dbo.ptostransaksi_th.TrxType <- TrxType` |
| IC-001685 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:147](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L147) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_th.TrxType` : `VARCHAR(8)` | `TrxType` : `NVARCHAR(255)` | `TrxType` → `VARCHAR(8)` | `retail.dbo.ptostransaksi_th.TrxType <- TrxType` |
