# Retail UAT.dbo.PTOSOutstanding_Error — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 49 lokasi source, 2 procedure, 9 kolom. **A:** 4 lokasi pada 2 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-kodeproduk"></a>
## KodeProduk

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Error.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001327 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L171) | JOIN_ON / P02 / COMPARISON | `er.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR(255)` | `er.KodeProduk` → `NVARCHAR` | `er.KodeProduk = tmp.KodeProduk` |
| IC-001433 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:210](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L210) | JOIN_ON / P02 / COMPARISON | `er.KodeProduk` : `VARCHAR(20)` | `tmp.KodeProduk` : `NVARCHAR(255)` | `er.KodeProduk` → `NVARCHAR` | `er.KodeProduk = tmp.KodeProduk` |
| IC-001250 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L36) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptosoutstanding_error.KodeProduk <- KodeProduk` |
| IC-001268 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L62) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptosoutstanding_error.KodeProduk <- KodeProduk` |
| IC-001344 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L39) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptosoutstanding_error.KodeProduk <- KodeProduk` |
| IC-001362 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptosoutstanding_error.KodeProduk <- KodeProduk` |
| IC-001374 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L94) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.KodeProduk` : `VARCHAR(20)` | `KodeProduk` : `NVARCHAR(255)` | `KodeProduk` → `VARCHAR(20)` | `retail.dbo.ptosoutstanding_error.KodeProduk <- KodeProduk` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Error.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 2 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001326 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L169) | JOIN_ON / P02 / COMPARISON | `er.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `er.[SID]` → `NVARCHAR` | `er.[SID] = tmp.[SID]` |
| IC-001432 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L208) | JOIN_ON / P02 / COMPARISON | `er.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `er.[SID]` → `NVARCHAR` | `er.[SID] = tmp.[SID]` |
| IC-001254 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L36) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptosoutstanding_error.SID <- [SID]` |
| IC-001272 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L62) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.SID` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tmp.[SID]` → `VARCHAR(100)` | `retail.dbo.ptosoutstanding_error.SID <- tmp.[SID]` |
| IC-001348 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L39) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptosoutstanding_error.SID <- [SID]` |
| IC-001366 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.SID` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tmp.[SID]` → `VARCHAR(100)` | `retail.dbo.ptosoutstanding_error.SID <- tmp.[SID]` |
| IC-001378 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L94) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.SID` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tmp.[SID]` → `VARCHAR(100)` | `retail.dbo.ptosoutstanding_error.SID <- tmp.[SID]` |

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Error.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001246 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L36) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptosoutstanding_error.AccruedInterest <- AccruedInterest` |
| IC-001264 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L62) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptosoutstanding_error.AccruedInterest <- AccruedInterest` |
| IC-001340 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L39) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptosoutstanding_error.AccruedInterest <- AccruedInterest` |
| IC-001358 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptosoutstanding_error.AccruedInterest <- AccruedInterest` |
| IC-001370 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L94) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.AccruedInterest` : `MONEY` | `AccruedInterest` : `FLOAT` | `AccruedInterest` → `MONEY` | `retail.dbo.ptosoutstanding_error.AccruedInterest <- AccruedInterest` |

<a id="kolom-ccy"></a>
## Ccy

**Tipe dasar:** `VARCHAR(3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Error.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001247 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L36) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptosoutstanding_error.Ccy <- Ccy` |
| IC-001265 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L62) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptosoutstanding_error.Ccy <- Ccy` |
| IC-001341 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L39) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptosoutstanding_error.Ccy <- Ccy` |
| IC-001359 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptosoutstanding_error.Ccy <- Ccy` |
| IC-001371 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L94) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.Ccy` : `VARCHAR(3)` | `Ccy` : `NVARCHAR(255)` | `Ccy` → `VARCHAR(3)` | `retail.dbo.ptosoutstanding_error.Ccy <- Ccy` |

<a id="kolom-hargaperolehannasabah"></a>
## HargaPerolehanNasabah

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Error.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001248 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L36) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `HargaPerolehanNasabah` : `FLOAT` | `HargaPerolehanNasabah` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_error.HargaPerolehanNasabah <- HargaPerolehanNasabah` |
| IC-001266 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L62) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `HargaPerolehanNasabah` : `FLOAT` | `HargaPerolehanNasabah` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_error.HargaPerolehanNasabah <- HargaPerolehanNasabah` |
| IC-001342 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L39) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `HargaPerolehanNasabah` : `FLOAT` | `HargaPerolehanNasabah` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_error.HargaPerolehanNasabah <- HargaPerolehanNasabah` |
| IC-001360 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `HargaPerolehanNasabah` : `FLOAT` | `HargaPerolehanNasabah` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_error.HargaPerolehanNasabah <- HargaPerolehanNasabah` |
| IC-001372 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L94) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.HargaPerolehanNasabah` : `DECIMAL(20,5)` | `HargaPerolehanNasabah` : `FLOAT` | `HargaPerolehanNasabah` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_error.HargaPerolehanNasabah <- HargaPerolehanNasabah` |

<a id="kolom-indikasihargabid"></a>
## IndikasiHargaBid

**Tipe dasar:** `DECIMAL(20,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Error.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001249 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L36) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.IndikasiHargaBid` : `DECIMAL(20,5)` | `IndikasiHargaBid` : `FLOAT` | `IndikasiHargaBid` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_error.IndikasiHargaBid <- IndikasiHargaBid` |
| IC-001267 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L62) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.IndikasiHargaBid` : `DECIMAL(20,5)` | `IndikasiHargaBid` : `FLOAT` | `IndikasiHargaBid` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_error.IndikasiHargaBid <- IndikasiHargaBid` |
| IC-001343 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L39) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.IndikasiHargaBid` : `DECIMAL(20,5)` | `IndikasiHargaBid` : `FLOAT` | `IndikasiHargaBid` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_error.IndikasiHargaBid <- IndikasiHargaBid` |
| IC-001361 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.IndikasiHargaBid` : `DECIMAL(20,5)` | `IndikasiHargaBid` : `FLOAT` | `IndikasiHargaBid` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_error.IndikasiHargaBid <- IndikasiHargaBid` |
| IC-001373 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L94) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.IndikasiHargaBid` : `DECIMAL(20,5)` | `IndikasiHargaBid` : `FLOAT` | `IndikasiHargaBid` → `DECIMAL(20,5)` | `retail.dbo.ptosoutstanding_error.IndikasiHargaBid <- IndikasiHargaBid` |

<a id="kolom-namanasabah"></a>
## NamaNasabah

**Tipe dasar:** `VARCHAR(95)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Error.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001251 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L36) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.NamaNasabah` : `VARCHAR(95)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptosoutstanding_error.NamaNasabah <- NamaNasabah` |
| IC-001269 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L62) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.NamaNasabah` : `VARCHAR(95)` | `tmp.NamaNasabah` : `NVARCHAR(255)` | `tmp.NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptosoutstanding_error.NamaNasabah <- tmp.NamaNasabah` |
| IC-001345 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L39) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.NamaNasabah` : `VARCHAR(95)` | `NamaNasabah` : `NVARCHAR(255)` | `NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptosoutstanding_error.NamaNasabah <- NamaNasabah` |
| IC-001363 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.NamaNasabah` : `VARCHAR(95)` | `tmp.NamaNasabah` : `NVARCHAR(255)` | `tmp.NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptosoutstanding_error.NamaNasabah <- tmp.NamaNasabah` |
| IC-001375 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L94) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.NamaNasabah` : `VARCHAR(95)` | `tmp.NamaNasabah` : `NVARCHAR(255)` | `tmp.NamaNasabah` → `VARCHAR(95)` | `retail.dbo.ptosoutstanding_error.NamaNasabah <- tmp.NamaNasabah` |

<a id="kolom-namaproduk"></a>
## NamaProduk

**Tipe dasar:** `VARCHAR(70)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Error.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001252 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L36) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptosoutstanding_error.NamaProduk <- NamaProduk` |
| IC-001270 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L62) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptosoutstanding_error.NamaProduk <- NamaProduk` |
| IC-001346 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L39) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptosoutstanding_error.NamaProduk <- NamaProduk` |
| IC-001364 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptosoutstanding_error.NamaProduk <- NamaProduk` |
| IC-001376 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L94) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.NamaProduk` : `VARCHAR(70)` | `NamaProduk` : `NVARCHAR(255)` | `NamaProduk` → `VARCHAR(70)` | `retail.dbo.ptosoutstanding_error.NamaProduk <- NamaProduk` |

<a id="kolom-nominal"></a>
## Nominal

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSOutstanding_Error.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001253 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L36) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptosoutstanding_error.Nominal <- Nominal` |
| IC-001271 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L62) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptosoutstanding_error.Nominal <- Nominal` |
| IC-001347 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L39) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptosoutstanding_error.Nominal <- Nominal` |
| IC-001365 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L65) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptosoutstanding_error.Nominal <- Nominal` |
| IC-001377 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:94](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L94) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosoutstanding_error.Nominal` : `MONEY` | `Nominal` : `FLOAT` | `Nominal` → `MONEY` | `retail.dbo.ptosoutstanding_error.Nominal <- Nominal` |
