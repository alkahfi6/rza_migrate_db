# Retail UAT.dbo.PTOSTransaksi_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 197 lokasi source, 14 procedure, 21 kolom. **A:** 16 lokasi pada 3 procedure. **B:** 4 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 2 SP. **A:** 8 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001439 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L31) | JOIN_ON / P02 / COMPARISON | `potr.DealId` : `VARCHAR(20)` | `tmp.[Deal ID]` : `NVARCHAR(255)` | `potr.DealId` → `NVARCHAR` | `tmp.[Deal ID] = potr.DealId` |
| IC-001496 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L102) | JOIN_ON / P02 / COMPARISON | `potr.DealId` : `VARCHAR(20)` | `tmp.DealId` : `NVARCHAR(255)` | `potr.DealId` → `NVARCHAR` | `potr.DealId = tmp.DealId` |
| IC-001535 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L143) | JOIN_ON / P02 / COMPARISON | `pttr.DealId` : `VARCHAR(20)` | `tmp.DealId` : `NVARCHAR(255)` | `pttr.DealId` → `NVARCHAR` | `pttr.DealId = tmp.DealId` |
| IC-001559 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L167) | JOIN_ON / P02 / COMPARISON | `tr.DealId` : `VARCHAR(20)` | `tmp.[Deal ID]` : `NVARCHAR(255)` | `tr.DealId` → `NVARCHAR` | `tr.DealId = tmp.[Deal ID]` |
| IC-001566 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L28) | JOIN_ON / P02 / COMPARISON | `potr.DealId` : `VARCHAR(20)` | `tmp.[Deal ID]` : `NVARCHAR(255)` | `potr.DealId` → `NVARCHAR` | `tmp.[Deal ID] = potr.DealId` |
| IC-001624 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:101](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L101) | JOIN_ON / P02 / COMPARISON | `potr.DealId` : `VARCHAR(20)` | `tmp.DealId` : `NVARCHAR(255)` | `potr.DealId` → `NVARCHAR` | `potr.DealId = tmp.DealId` |
| IC-001663 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L143) | JOIN_ON / P02 / COMPARISON | `pttr.DealId` : `VARCHAR(20)` | `tmp.DealId` : `NVARCHAR(255)` | `pttr.DealId` → `NVARCHAR` | `pttr.DealId = tmp.DealId` |
| IC-001687 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L167) | JOIN_ON / P02 / COMPARISON | `tr.DealId` : `VARCHAR(20)` | `tmp.[Deal ID]` : `NVARCHAR(255)` | `tr.DealId` → `NVARCHAR` | `tr.DealId = tmp.[Deal ID]` |
| IC-001500 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.DealId` : `VARCHAR(20)` | `DealId` : `NVARCHAR(255)` | `DealId` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_tr.DealId <- DealId` |
| IC-001628 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.DealId` : `VARCHAR(20)` | `DealId` : `NVARCHAR(255)` | `DealId` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_tr.DealId <- DealId` |

<a id="kolom-trxtype"></a>
## TrxType

**Tipe dasar:** `VARCHAR(8)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 3 SP. **A:** 4 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004970 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L46) | WHERE / P01 / COMPARISON | `TrxType` : `VARCHAR(8)` | `@pnTrxType` : `SMALLINT` | `TrxType` → `SMALLINT` | `TrxType = @pnTrxType` |
| IC-004973 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L75) | WHERE / P01 / COMPARISON | `TrxType` : `VARCHAR(8)` | `@pnTrxType` : `SMALLINT` | `TrxType` → `SMALLINT` | `TrxType = @pnTrxType` |
| IC-004976 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:111](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L111) | WHERE / P01 / COMPARISON | `TrxType` : `VARCHAR(8)` | `@pnTrxType` : `SMALLINT` | `TrxType` → `SMALLINT` | `TrxType = @pnTrxType` |
| IC-004979 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L141) | WHERE / P01 / COMPARISON | `TrxType` : `VARCHAR(8)` | `@pnTrxType` : `SMALLINT` | `TrxType` → `SMALLINT` | `TrxType = @pnTrxType` |
| IC-004968 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L24) | CASE / P01 / COMPARISON | `TrxType` : `VARCHAR(8)` | `3` : `int` | `TrxType` → `INT` | `TrxType = 3` |
| IC-004971 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:57](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L57) | CASE / P01 / COMPARISON | `TrxType` : `VARCHAR(8)` | `3` : `int` | `TrxType` → `INT` | `TrxType = 3` |
| IC-004974 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L89) | CASE / P01 / COMPARISON | `TrxType` : `VARCHAR(8)` | `3` : `int` | `TrxType` → `INT` | `TrxType = 3` |
| IC-004977 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L123) | CASE / P01 / COMPARISON | `TrxType` : `VARCHAR(8)` | `3` : `int` | `TrxType` → `INT` | `TrxType = 3` |
| IC-001516 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.TrxType` : `VARCHAR(8)` | `TrxType` : `NVARCHAR(255)` | `TrxType` → `VARCHAR(8)` | `retail.dbo.ptostransaksi_tr.TrxType <- TrxType` |
| IC-001520 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L122) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.TrxType` : `VARCHAR(8)` | `tmp.TrxType` : `NVARCHAR(255)` | `tmp.TrxType` → `VARCHAR(8)` | `pttr.TrxType = tmp.TrxType` |
| IC-001644 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.TrxType` : `VARCHAR(8)` | `TrxType` : `NVARCHAR(255)` | `TrxType` → `VARCHAR(8)` | `retail.dbo.ptostransaksi_tr.TrxType <- TrxType` |
| IC-001648 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:121](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L121) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.TrxType` : `VARCHAR(8)` | `tmp.TrxType` : `NVARCHAR(255)` | `tmp.TrxType` → `VARCHAR(8)` | `pttr.TrxType = tmp.TrxType` |

<a id="kolom-kodeproduk"></a>
## KodeProduk

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001562 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L180) | JOIN_ON / P02 / COMPARISON | `tr.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR(255)` | `tr.KodeProduk` → `NVARCHAR` | `tr.KodeProduk = tmp.KodeProduk` |
| IC-001690 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L180) | JOIN_ON / P02 / COMPARISON | `tr.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR (expression; length unknown)` | `tr.KodeProduk` → `NVARCHAR` | `tr.KodeProduk = tmp.KodeProduk` |
| IC-001504 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_tr.KodeProduk <- KodeProduk` |
| IC-001521 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L123) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR(255)` | `tmp.KodeProduk` → `VARCHAR(20)` | `pttr.KodeProduk = tmp.KodeProduk` |
| IC-001632 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR (expression; length unknown)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptostransaksi_tr.KodeProduk <- KodeProduk` |
| IC-001649 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:122](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L122) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR (expression; length unknown)` | `tmp.KodeProduk` → `VARCHAR(20)` | `pttr.KodeProduk = tmp.KodeProduk` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L3). **Existing leading index:** ix_ptostransaksi_tr_sid.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001561 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L178) | JOIN_ON / P02 / COMPARISON | `tr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tr.[SID]` → `NVARCHAR` | `tr.[SID] = tmp.[SID]` |
| IC-001689 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L178) | JOIN_ON / P02 / COMPARISON | `tr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR (expression; length unknown)` | `tr.[SID]` → `NVARCHAR` | `tr.[SID] = tmp.[SID]` |
| IC-001511 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptostransaksi_tr.SID <- [SID]` |
| IC-001517 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L116) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tmp.[SID]` → `VARCHAR(100)` | `pttr.[SID] = tmp.[SID]` |
| IC-001639 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR (expression; length unknown)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptostransaksi_tr.SID <- [SID]` |
| IC-001645 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L115) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR (expression; length unknown)` | `tmp.[SID]` → `VARCHAR(100)` | `pttr.[SID] = tmp.[SID]` |

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001498 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptostransaksi_tr.AccruedInterest <- AccruedInterest` |
| IC-001527 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L129) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.AccruedInterest` : `MONEY` | `tmp.AccruedInterest` : `FLOAT` | `tmp.AccruedInterest` → `MONEY` | `pttr.AccruedInterest = tmp.AccruedInterest` |
| IC-001626 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptostransaksi_tr.AccruedInterest <- AccruedInterest` |
| IC-001655 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L128) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.AccruedInterest` : `MONEY` | `tmp.AccruedInterest` : `FLOAT` | `tmp.AccruedInterest` → `MONEY` | `pttr.AccruedInterest = tmp.AccruedInterest` |
| IC-002390 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:460](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L460) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.AccruedInterest, 0)` |
| IC-002404 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:652](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L652) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.AccruedInterest, 0)` |
| IC-005776 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:44](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L44) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.AccruedInterest,0)` |
| IC-005784 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L98) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.AccruedInterest,0)` |
| IC-005793 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L156) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.AccruedInterest,0)` |
| IC-005801 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L208) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.AccruedInterest,0)` |
| IC-005849 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:459](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L459) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.AccruedInterest, 0)` |
| IC-005863 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:651](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L651) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.AccruedInterest, 0)` |
| IC-010615 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPTOSTransaksi:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSTransaksi.sql#L81) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.AccruedInterest,0)` |

<a id="kolom-approveddate"></a>
## ApprovedDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L30). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007183 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSTrxAuthen:430](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrxAuthen.sql#L430) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.ApprovedDate` : `DATETIME` | `tmp.ApprovedBy` : `int` | `tmp.ApprovedBy` → `DATETIME` | `pttr.ApprovedDate = tmp.ApprovedBy` |
| IC-007243 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:921](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L921) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.ApprovedDate` : `DATETIME` | `tmp.ApprovedBy` : `int` | `tmp.ApprovedBy` → `DATETIME` | `pttr.ApprovedDate = tmp.ApprovedBy` |

<a id="kolom-ccy"></a>
## Ccy

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001499 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptostransaksi_tr.Ccy <- Ccy` |
| IC-001523 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L125) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.Ccy` : `VARCHAR(3)` | `tmp.Ccy` : `NVARCHAR(255)` | `tmp.Ccy` → `VARCHAR(3)` | `pttr.Ccy = tmp.Ccy` |
| IC-001627 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptostransaksi_tr.Ccy <- Ccy` |
| IC-001651 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L124) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.Ccy` : `VARCHAR(3)` | `tmp.Ccy` : `NVARCHAR(255)` | `tmp.Ccy` → `VARCHAR(3)` | `pttr.Ccy = tmp.Ccy` |

<a id="kolom-dealprice"></a>
## DealPrice

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001501 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.DealPrice` : `DECIMAL(20,5)` | `DealPrice` : `FLOAT` | `DealPrice` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_tr.DealPrice <- DealPrice` |
| IC-001525 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L127) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.DealPrice` : `DECIMAL(20,5)` | `tmp.DealPrice` : `FLOAT` | `tmp.DealPrice` → `DECIMAL(20,5)` | `pttr.DealPrice = tmp.DealPrice` |
| IC-001629 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.DealPrice` : `DECIMAL(20,5)` | `DealPrice` : `FLOAT` | `DealPrice` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_tr.DealPrice <- DealPrice` |
| IC-001653 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L126) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.DealPrice` : `DECIMAL(20,5)` | `tmp.DealPrice` : `FLOAT` | `tmp.DealPrice` → `DECIMAL(20,5)` | `pttr.DealPrice = tmp.DealPrice` |

<a id="kolom-gainloss"></a>
## GainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001502 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.GainLoss` : `MONEY` | `GainLoss` : `FLOAT` | `GainLoss` → `MONEY` | `retail.dbo.ptostransaksi_tr.GainLoss <- GainLoss` |
| IC-001530 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L133) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.GainLoss` : `MONEY` | `tmp.GainLoss` : `FLOAT` | `tmp.GainLoss` → `MONEY` | `pttr.GainLoss = tmp.GainLoss` |
| IC-001630 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.GainLoss` : `MONEY` | `GainLoss` : `FLOAT` | `GainLoss` → `MONEY` | `retail.dbo.ptostransaksi_tr.GainLoss <- GainLoss` |
| IC-001658 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L132) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.GainLoss` : `MONEY` | `tmp.GainLoss` : `FLOAT` | `tmp.GainLoss` → `MONEY` | `pttr.GainLoss = tmp.GainLoss` |
| IC-010617 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPTOSTransaksi:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSTransaksi.sql#L82) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.GainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.GainLoss,0)` |

<a id="kolom-hargaperolehan"></a>
## HargaPerolehan

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001503 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.HargaPerolehan` : `DECIMAL(20,5)` | `HargaPerolehan` : `FLOAT` | `HargaPerolehan` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_tr.HargaPerolehan <- HargaPerolehan` |
| IC-001529 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L132) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.HargaPerolehan` : `DECIMAL(20,5)` | `tmp.HargaPerolehan` : `FLOAT` | `tmp.HargaPerolehan` → `DECIMAL(20,5)` | `pttr.HargaPerolehan = tmp.HargaPerolehan` |
| IC-001631 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.HargaPerolehan` : `DECIMAL(20,5)` | `HargaPerolehan` : `FLOAT` | `HargaPerolehan` → `DECIMAL(20,5)` | `retail.dbo.ptostransaksi_tr.HargaPerolehan <- HargaPerolehan` |
| IC-001657 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:131](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L131) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.HargaPerolehan` : `DECIMAL(20,5)` | `tmp.HargaPerolehan` : `FLOAT` | `tmp.HargaPerolehan` → `DECIMAL(20,5)` | `pttr.HargaPerolehan = tmp.HargaPerolehan` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002394 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:474](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L474) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-002408 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:666](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L666) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-005780 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUTrx:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L61) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-005788 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUTrx:115](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L115) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-005797 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUTrx:173](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L173) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-005805 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUTrx:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L225) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-005853 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUploadAuthen:473](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L473) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |
| IC-005867 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUploadAuthen:665](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L665) | JOIN_ON / P01 / COMPARISON | `pttr.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `pttr.InsertedBy = ndme.EMPLID` |

<a id="kolom-npwp"></a>
## NPWP

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001505 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.NPWP` : `VARCHAR(40)` | `NPWP` : `NVARCHAR(255)` | `NPWP` → `VARCHAR(40)` | `retail.dbo.ptostransaksi_tr.NPWP <- NPWP` |
| IC-001519 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:119](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L119) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.NPWP` : `VARCHAR(40)` | `tmp.NPWP` : `NVARCHAR(255)` | `tmp.NPWP` → `VARCHAR(40)` | `pttr.NPWP = tmp.NPWP` |
| IC-001633 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.NPWP` : `VARCHAR(40)` | `NPWP` : `NVARCHAR(255)` | `NPWP` → `VARCHAR(40)` | `retail.dbo.ptostransaksi_tr.NPWP <- NPWP` |
| IC-001647 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L118) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.NPWP` : `VARCHAR(40)` | `tmp.NPWP` : `NVARCHAR(255)` | `tmp.NPWP` → `VARCHAR(40)` | `pttr.NPWP = tmp.NPWP` |

<a id="kolom-namanasabah"></a>
## NamaNasabah

**Tipe dasar:** `VARCHAR(95)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001506 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.NamaNasabah` : `VARCHAR(95)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptostransaksi_tr.NamaNasabah <- NamaNasabah` |
| IC-001518 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L118) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.NamaNasabah` : `VARCHAR(95)` | `tmp.NamaNasabah` : `NVARCHAR(255)` | `tmp.NamaNasabah` → `VARCHAR(95)` | `pttr.NamaNasabah = tmp.NamaNasabah` |
| IC-001634 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.NamaNasabah` : `VARCHAR(95)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptostransaksi_tr.NamaNasabah <- NamaNasabah` |
| IC-001646 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L117) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.NamaNasabah` : `VARCHAR(95)` | `tmp.NamaNasabah` : `NVARCHAR(255)` | `tmp.NamaNasabah` → `VARCHAR(95)` | `pttr.NamaNasabah = tmp.NamaNasabah` |

<a id="kolom-namaproduk"></a>
## NamaProduk

**Tipe dasar:** `VARCHAR(70)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001507 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptostransaksi_tr.NamaProduk <- NamaProduk` |
| IC-001522 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:124](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L124) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.NamaProduk` : `VARCHAR(70)` | `tmp.NamaProduk` : `NVARCHAR(255)` | `tmp.NamaProduk` → `VARCHAR(70)` | `pttr.NamaProduk = tmp.NamaProduk` |
| IC-001635 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptostransaksi_tr.NamaProduk <- NamaProduk` |
| IC-001650 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L123) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.NamaProduk` : `VARCHAR(70)` | `tmp.NamaProduk` : `NVARCHAR(255)` | `tmp.NamaProduk` → `VARCHAR(70)` | `pttr.NamaProduk = tmp.NamaProduk` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001508 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptostransaksi_tr.Nominal <- Nominal` |
| IC-001524 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L126) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.Nominal` : `MONEY` | `tmp.Nominal` : `FLOAT` | `tmp.Nominal` → `MONEY` | `pttr.Nominal = tmp.Nominal` |
| IC-001636 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptostransaksi_tr.Nominal <- Nominal` |
| IC-001652 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:125](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L125) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.Nominal` : `MONEY` | `tmp.Nominal` : `FLOAT` | `tmp.Nominal` → `MONEY` | `pttr.Nominal = tmp.Nominal` |
| IC-010613 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPTOSTransaksi:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSTransaksi.sql#L80) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.Nominal` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.Nominal,0)` |

<a id="kolom-ptostype"></a>
## PTOSType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 26 lokasi / 8 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001438 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L31) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `2` : `int` | `potr.PTOSType` → `INT` | `potr.PTOSType = 2` |
| IC-001497 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L103) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `potr.PTOSType` → `INT` | `potr.PTOSType = tmp.PTOSType` |
| IC-001536 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L144) | JOIN_ON / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = tmp.PTOSType` |
| IC-001560 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L168) | WHERE / P11 / COMPARISON | `tr.PTOSType` : `SMALLINT` | `2` : `int` | `tr.PTOSType` → `INT` | `tr.PTOSType = 2` |
| IC-001563 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L181) | JOIN_ON / P11 / COMPARISON | `tr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `tr.PTOSType` → `INT` | `tr.PTOSType = tmp.PTOSType` |
| IC-001567 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L29) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `1` : `int` | `potr.PTOSType` → `INT` | `potr.PTOSType = 1` |
| IC-001625 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:102](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L102) | JOIN_ON / P11 / COMPARISON | `potr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `potr.PTOSType` → `INT` | `potr.PTOSType = tmp.PTOSType` |
| IC-001664 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L144) | JOIN_ON / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = tmp.PTOSType` |
| IC-001688 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L168) | WHERE / P11 / COMPARISON | `tr.PTOSType` : `SMALLINT` | `1` : `int` | `tr.PTOSType` → `INT` | `tr.PTOSType = 1` |
| IC-001691 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L181) | JOIN_ON / P11 / COMPARISON | `tr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `tr.PTOSType` → `INT` | `tr.PTOSType = tmp.PTOSType` |
| IC-002311 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L298) | WHERE / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = @pnType` |
| IC-002352 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L162) | JOIN_ON / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `pttr.PTOSType` → `INT` | `tmp.PTOSType = pttr.PTOSType` |
| IC-002357 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L170) | JOIN_ON / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = tmp.PTOSType` |
| IC-002361 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L186) | JOIN_ON / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `pttr.PTOSType` → `INT` | `tmp.PTOSType = pttr.PTOSType` |
| IC-002365 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:202](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L202) | JOIN_ON / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `pttr.PTOSType` → `INT` | `tmp.PTOSType = pttr.PTOSType` |
| IC-002410 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:668](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L668) | WHERE / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = @pnType` |
| IC-005790 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUTrx:117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L117) | WHERE / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = @pnType` |
| IC-005807 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUTrx:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L228) | WHERE / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = @pnType` |
| IC-005825 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:862](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L862) | WHERE / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = @pnType` |
| IC-005827 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:1118](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L1118) | WHERE / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = @pnType` |
| IC-005869 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:667](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L667) | WHERE / P11 / COMPARISON | `pttr.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `pttr.PTOSType` → `INT` | `pttr.PTOSType = @pnType` |
| IC-001509 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.PTOSType` : `SMALLINT` | `PTOSType` : `int` | `PTOSType` → `SMALLINT` | `retail.dbo.ptostransaksi_tr.PTOSType <- PTOSType` |
| IC-001533 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L136) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `tmp.PTOSType` → `SMALLINT` | `pttr.PTOSType = tmp.PTOSType` |
| IC-001637 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.PTOSType` : `SMALLINT` | `PTOSType` : `int` | `PTOSType` → `SMALLINT` | `retail.dbo.ptostransaksi_tr.PTOSType <- PTOSType` |
| IC-001661 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L135) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `int` | `tmp.PTOSType` → `SMALLINT` | `pttr.PTOSType = tmp.PTOSType` |
| IC-002363 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L189) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `tmp.PTOSType` → `SMALLINT` | `retail.dbo.ptostransaksi_tr.PTOSType <- tmp.PTOSType` |

<a id="kolom-proceed"></a>
## Proceed

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001510 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Proceed` : `MONEY` | `Proceed` : `FLOAT` | `Proceed` → `MONEY` | `retail.dbo.ptostransaksi_tr.Proceed <- Proceed` |
| IC-001526 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L128) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.Proceed` : `MONEY` | `tmp.Proceed` : `FLOAT` | `tmp.Proceed` → `MONEY` | `pttr.Proceed = tmp.Proceed` |
| IC-001638 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Proceed` : `MONEY` | `Proceed` : `FLOAT` | `Proceed` → `MONEY` | `retail.dbo.ptostransaksi_tr.Proceed <- Proceed` |
| IC-001654 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L127) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.Proceed` : `MONEY` | `tmp.Proceed` : `FLOAT` | `tmp.Proceed` → `MONEY` | `pttr.Proceed = tmp.Proceed` |
| IC-010614 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPTOSTransaksi:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSTransaksi.sql#L80) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.Proceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.Proceed, 0)` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 36 lokasi / 13 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, CONVERTED_COLUMN_EXPRESSION, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002312 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPTOSProcessUpload:298](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSProcessUpload.sql#L298) | WHERE / P11 / COMPARISON | `pttr.Status` : `SMALLINT` | `3` : `int` | `pttr.Status` → `INT` | `pttr.Status != 3` |
| IC-002353 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN_EXPRESSION | [dbo.Rvmp_TRSPTOSSubmitFromTemp:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L163) | WHERE / P11 / COMPARISON | `ISNULL(pttr.Status, 0)` : `SMALLINT` | `3` : `int` | `ISNULL(pttr.Status, 0)` → `INT` | `ISNULL(pttr.Status, 0) != 3` |
| IC-002395 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:475](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L475) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `0` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 0` |
| IC-002409 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:667](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L667) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `0` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 0` |
| IC-004969 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L41) | WHERE / P11 / COMPARISON | `[Status]` : `SMALLINT` | `3` : `int` | `[Status]` → `INT` | `[Status] = 3` |
| IC-004972 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L70) | WHERE / P11 / COMPARISON | `[Status]` : `SMALLINT` | `3` : `int` | `[Status]` → `INT` | `[Status] = 3` |
| IC-004975 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L106) | WHERE / P11 / COMPARISON | `[Status]` : `SMALLINT` | `3` : `int` | `[Status]` → `INT` | `[Status] = 3` |
| IC-004978 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPTOSPopulateTransaksi:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPTOSPopulateTransaksi.sql#L136) | WHERE / P11 / COMPARISON | `[Status]` : `SMALLINT` | `3` : `int` | `[Status]` → `INT` | `[Status] = 3` |
| IC-005781 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUTrx:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L62) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `3` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 3` |
| IC-005789 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUTrx:116](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L116) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `3` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 3` |
| IC-005798 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUTrx:174](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L174) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `3` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 3` |
| IC-005806 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUTrx:226](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L226) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `3` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 3` |
| IC-005824 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:861](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L861) | WHERE / P11 / COMPARISON | `pttr.Status` : `SMALLINT` | `3` : `int` | `pttr.Status` → `INT` | `pttr.Status != 3` |
| IC-005826 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUpload:1117](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUpload.sql#L1117) | WHERE / P11 / COMPARISON | `pttr.Status` : `SMALLINT` | `3` : `int` | `pttr.Status` → `INT` | `pttr.Status != 3` |
| IC-005854 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:474](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L474) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `0` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 0` |
| IC-005868 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:666](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L666) | WHERE / P11 / COMPARISON | `pttr.[Status]` : `SMALLINT` | `0` : `int` | `pttr.[Status]` → `INT` | `pttr.[Status] = 0` |
| IC-007204 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSUpload:1281](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L1281) | WHERE / P11 / COMPARISON | `pttr.Status` : `SMALLINT` | `3` : `int` | `pttr.Status` → `INT` | `pttr.Status != 3` |
| IC-007208 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessPTOSUpload:1743](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L1743) | WHERE / P11 / COMPARISON | `pttr.Status` : `SMALLINT` | `3` : `int` | `pttr.Status` → `INT` | `pttr.Status != 3` |
| IC-001512 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.ptostransaksi_tr.Status <- 3` |
| IC-001534 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:137](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L137) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `pttr.[Status] = 3` |
| IC-001640 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.ptostransaksi_tr.Status <- 3` |
| IC-001662 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L136) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `pttr.[Status] = 3` |
| IC-002355 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPTOSSubmitFromTemp:163](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L163) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `ISNULL(pttr.Status, 0)` |
| IC-002356 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:168](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L168) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.[Status]` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `pttr.[Status] = 1` |
| IC-002364 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L189) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.ptostransaksi_tr.Status <- 0` |
| IC-007170 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSTrx:175](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrx.sql#L175) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.Status` : `SMALLINT` | `@pnStatus` : `INT` | `@pnStatus` → `SMALLINT` | `pttr.Status = @pnStatus` |
| IC-007182 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSTrxAuthen:426](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrxAuthen.sql#L426) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `pttr.Status = 3` |
| IC-007188 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSTrxAuthen:517](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrxAuthen.sql#L517) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `[Status] = 3` |
| IC-007191 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSTrxAuthen:532](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrxAuthen.sql#L532) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `[Status] = 3` |
| IC-007206 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:1403](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L1403) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `[Status] = 1` |
| IC-007207 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:1409](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L1409) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.ptostransaksi_tr.Status <- 0` |
| IC-007210 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:1853](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L1853) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `[Status] = 1` |
| IC-007211 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:1859](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L1859) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.ptostransaksi_tr.Status <- 0` |
| IC-007240 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:886](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L886) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `pttr.Status = 3` |
| IC-007242 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:917](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L917) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `pttr.Status = 3` |
| IC-007247 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:947](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L947) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `[Status] = 3` |

<a id="kolom-taxcapitalgainloss"></a>
## TaxCapitalGainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001513 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.TaxCapitalGainLoss` : `MONEY` | `TaxCapitalGainLoss` : `FLOAT` | `TaxCapitalGainLoss` → `MONEY` | `retail.dbo.ptostransaksi_tr.TaxCapitalGainLoss <- TaxCapitalGainLoss` |
| IC-001531 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L134) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.TaxCapitalGainLoss` : `MONEY` | `tmp.TaxCapitalGainLoss` : `FLOAT` | `tmp.TaxCapitalGainLoss` → `MONEY` | `pttr.TaxCapitalGainLoss = tmp.TaxCapitalGainLoss` |
| IC-001641 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.TaxCapitalGainLoss` : `MONEY` | `TaxCapitalGainLoss` : `FLOAT` | `TaxCapitalGainLoss` → `MONEY` | `retail.dbo.ptostransaksi_tr.TaxCapitalGainLoss <- TaxCapitalGainLoss` |
| IC-001659 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L133) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.TaxCapitalGainLoss` : `MONEY` | `tmp.TaxCapitalGainLoss` : `FLOAT` | `tmp.TaxCapitalGainLoss` → `MONEY` | `pttr.TaxCapitalGainLoss = tmp.TaxCapitalGainLoss` |
| IC-002392 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:465](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L465) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxCapitalGainLoss, 0)` |
| IC-002406 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:657](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L657) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxCapitalGainLoss, 0)` |
| IC-005778 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L49) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxCapitalGainLoss,0)` |
| IC-005786 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:103](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L103) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxCapitalGainLoss,0)` |
| IC-005795 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L161) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxCapitalGainLoss,0)` |
| IC-005803 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L213) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxCapitalGainLoss,0)` |
| IC-005851 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:464](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L464) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxCapitalGainLoss, 0)` |
| IC-005865 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:656](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L656) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxCapitalGainLoss, 0)` |
| IC-010618 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPTOSTransaksi:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSTransaksi.sql#L82) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxCapitalGainLoss,0)` |

<a id="kolom-taxonaccrued"></a>
## TaxOnAccrued

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001514 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.TaxOnAccrued` : `MONEY` | `TaxOnAccrued` : `FLOAT` | `TaxOnAccrued` → `MONEY` | `retail.dbo.ptostransaksi_tr.TaxOnAccrued <- TaxOnAccrued` |
| IC-001528 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L130) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.TaxOnAccrued` : `MONEY` | `tmp.TaxOnAccrued` : `FLOAT` | `tmp.TaxOnAccrued` → `MONEY` | `pttr.TaxOnAccrued = tmp.TaxOnAccrued` |
| IC-001642 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.TaxOnAccrued` : `MONEY` | `TaxOnAccrued` : `FLOAT` | `TaxOnAccrued` → `MONEY` | `retail.dbo.ptostransaksi_tr.TaxOnAccrued <- TaxOnAccrued` |
| IC-001656 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L129) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.TaxOnAccrued` : `MONEY` | `tmp.TaxOnAccrued` : `FLOAT` | `tmp.TaxOnAccrued` → `MONEY` | `pttr.TaxOnAccrued = tmp.TaxOnAccrued` |
| IC-002391 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:461](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L461) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxOnAccrued, 0)` |
| IC-002405 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:653](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L653) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxOnAccrued, 0)` |
| IC-005777 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L45) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxOnAccrued,0)` |
| IC-005785 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L99) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxOnAccrued,0)` |
| IC-005794 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L157) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxOnAccrued,0)` |
| IC-005802 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:209](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L209) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxOnAccrued,0)` |
| IC-005850 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:460](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L460) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxOnAccrued, 0)` |
| IC-005864 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:652](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L652) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TaxOnAccrued, 0)` |
| IC-010616 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPTOSTransaksi:81](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSTransaksi.sql#L81) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TaxOnAccrued ,0)` |

<a id="kolom-totalproceed"></a>
## TotalProceed

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_TR.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 13 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001515 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L106) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.TotalProceed` : `MONEY` | `TotalProceed` : `FLOAT` | `TotalProceed` → `MONEY` | `retail.dbo.ptostransaksi_tr.TotalProceed <- TotalProceed` |
| IC-001532 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L135) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.TotalProceed` : `MONEY` | `tmp.TotalProceed` : `FLOAT` | `tmp.TotalProceed` → `MONEY` | `pttr.TotalProceed = tmp.TotalProceed` |
| IC-001643 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:105](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L105) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_tr.TotalProceed` : `MONEY` | `TotalProceed` : `FLOAT` | `TotalProceed` → `MONEY` | `retail.dbo.ptostransaksi_tr.TotalProceed <- TotalProceed` |
| IC-001660 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L134) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pttr.TotalProceed` : `MONEY` | `tmp.TotalProceed` : `FLOAT` | `tmp.TotalProceed` → `MONEY` | `pttr.TotalProceed = tmp.TotalProceed` |
| IC-002393 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:466](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L466) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TotalProceed, 0)` |
| IC-002407 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:658](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L658) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TotalProceed, 0)` |
| IC-005779 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L50) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TotalProceed,0)` |
| IC-005787 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:104](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L104) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TotalProceed,0)` |
| IC-005796 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:162](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L162) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TotalProceed,0)` |
| IC-005804 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUTrx:214](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUTrx.sql#L214) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TotalProceed,0)` |
| IC-005852 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:465](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L465) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TotalProceed, 0)` |
| IC-005866 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:657](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L657) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `pttr.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(pttr.TotalProceed, 0)` |
| IC-010619 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptPTOSTransaksi:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptPTOSTransaksi.sql#L83) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(tr.TotalProceed,0)` |
