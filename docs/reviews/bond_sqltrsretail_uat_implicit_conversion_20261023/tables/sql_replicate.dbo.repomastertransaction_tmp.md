# SQL_REPLICATE.dbo.RepoMasterTransaction_Tmp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 1 procedure, 14 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 12 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-dirtyprice"></a>
## DirtyPrice

**Tipe dasar:** `NUMERIC(24,14)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001868 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.DirtyPrice` : `NUMERIC(24,14)` | `retail.dbo.repomastertransaction_tt.DirtyPrice` : `DECIMAL(15,4)` | `rmt.DirtyPrice` → `DECIMAL(15,4)` | `retail.dbo.repomastertransaction_tt.DirtyPrice <- rmt.DirtyPrice` |

<a id="kolom-facevalue"></a>
## FaceValue

**Tipe dasar:** `NVARCHAR(38)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001869 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.FaceValue` : `NVARCHAR(38)` | `retail.dbo.repomastertransaction_tt.FaceValue` : `DECIMAL(25,2)` | `rmt.FaceValue` → `DECIMAL(25,2)` | `retail.dbo.repomastertransaction_tt.FaceValue <- rmt.FaceValue` |

<a id="kolom-haircutmethod"></a>
## HaircutMethod

**Tipe dasar:** `NVARCHAR(5)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001870 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.HaircutMethod` : `NVARCHAR(5)` | `retail.dbo.repomastertransaction_tt.HaircutMethod` : `VARCHAR(15)` | `rmt.HaircutMethod` → `VARCHAR(15)` | `retail.dbo.repomastertransaction_tt.HaircutMethod <- rmt.HaircutMethod` |

<a id="kolom-isincode"></a>
## ISINCode

**Tipe dasar:** `NVARCHAR(15)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001880 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoRONDataMigrate:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L71) | JOIN_ON / P02 / COMPARISON | `rmt.ISINCode` : `NVARCHAR(15)` | `sm.ISINCode` : `VARCHAR(20)` | `sm.ISINCode` → `NVARCHAR` | `rmt.ISINCode = sm.ISINCode` |

<a id="kolom-maturitydate"></a>
## MaturityDate

**Tipe dasar:** `NVARCHAR(8)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001871 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.MaturityDate` : `NVARCHAR(8)` | `retail.dbo.repomastertransaction_tt.MaturityDate` : `DATETIME` | `rmt.MaturityDate` → `DATETIME` | `retail.dbo.repomastertransaction_tt.MaturityDate <- rmt.MaturityDate` |

<a id="kolom-nodays"></a>
## NoDays

**Tipe dasar:** `NVARCHAR(38)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001872 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.NoDays` : `NVARCHAR(38)` | `retail.dbo.repomastertransaction_tt.NoDays` : `INT` | `rmt.NoDays` → `INT` | `retail.dbo.repomastertransaction_tt.NoDays <- rmt.NoDays` |

<a id="kolom-nomurexexternal"></a>
## NoMurexExternal

**Tipe dasar:** `NUMERIC(10,0)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001873 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.NoMurexExternal` : `NUMERIC(10,0)` | `retail.dbo.repomastertransaction_tt.NoMurexExternal` : `VARCHAR(20)` | `rmt.NoMurexExternal` → `VARCHAR(20)` | `retail.dbo.repomastertransaction_tt.NoMurexExternal <- rmt.NoMurexExternal` |
| IC-001881 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoRONDataMigrate:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L73) | JOIN_ON / P01 / COMPARISON | `rmt.NoMurexExternal` : `NUMERIC(10,0)` | `rmtt.NoMurexExternal` : `VARCHAR(20)` | `rmtt.NoMurexExternal` → `NUMERIC` | `rmt.NoMurexExternal = rmtt.NoMurexExternal` |

<a id="kolom-rateconvention"></a>
## RateConvention

**Tipe dasar:** `NVARCHAR(15)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001874 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.RateConvention` : `NVARCHAR(15)` | `retail.dbo.repomastertransaction_tt.RateConvention` : `VARCHAR(25)` | `rmt.RateConvention` → `VARCHAR(25)` | `retail.dbo.repomastertransaction_tt.RateConvention <- rmt.RateConvention` |

<a id="kolom-repo1stleg"></a>
## Repo1stLeg

**Tipe dasar:** `NVARCHAR(38)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001875 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.Repo1stLeg` : `NVARCHAR(38)` | `retail.dbo.repomastertransaction_tt.Repo1stLeg` : `DECIMAL(25,2)` | `rmt.Repo1stLeg` → `DECIMAL(25,2)` | `retail.dbo.repomastertransaction_tt.Repo1stLeg <- rmt.Repo1stLeg` |

<a id="kolom-repo2ndleg"></a>
## Repo2ndLeg

**Tipe dasar:** `NVARCHAR(38)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001876 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.Repo2ndLeg` : `NVARCHAR(38)` | `retail.dbo.repomastertransaction_tt.Repo2ndLeg` : `DECIMAL(25,2)` | `rmt.Repo2ndLeg` → `DECIMAL(25,2)` | `retail.dbo.repomastertransaction_tt.Repo2ndLeg <- rmt.Repo2ndLeg` |

<a id="kolom-repointerest"></a>
## RepoInterest

**Tipe dasar:** `NVARCHAR(38)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001877 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.RepoInterest` : `NVARCHAR(38)` | `retail.dbo.repomastertransaction_tt.RepoInterest` : `DECIMAL(15,4)` | `rmt.RepoInterest` → `DECIMAL(15,4)` | `retail.dbo.repomastertransaction_tt.RepoInterest <- rmt.RepoInterest` |

<a id="kolom-startdate"></a>
## StartDate

**Tipe dasar:** `NVARCHAR(8)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001878 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.StartDate` : `NVARCHAR(8)` | `retail.dbo.repomastertransaction_tt.StartDate` : `DATETIME` | `rmt.StartDate` → `DATETIME` | `retail.dbo.repomastertransaction_tt.StartDate <- rmt.StartDate` |

<a id="kolom-trxdate"></a>
## TrxDate

**Tipe dasar:** `NVARCHAR(8)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001879 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.RepoRONDataMigrate:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L13) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `rmt.TrxDate` : `NVARCHAR(8)` | `retail.dbo.repomastertransaction_tt.TrxDate` : `DATETIME` | `rmt.TrxDate` → `DATETIME` | `retail.dbo.repomastertransaction_tt.TrxDate <- rmt.TrxDate` |

<a id="kolom-trxstatus"></a>
## TrxStatus

**Tipe dasar:** `NVARCHAR(3)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.RepoMasterTransaction_Tmp.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001882 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RepoRONDataMigrate:74](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoRONDataMigrate.sql#L74) | WHERE / P04 / COMPARISON | `rmt.TrxStatus` : `NVARCHAR(3)` | `'INS'` : `varchar` | `'INS'` → `NVARCHAR` | `rmt.TrxStatus = 'INS'` |
