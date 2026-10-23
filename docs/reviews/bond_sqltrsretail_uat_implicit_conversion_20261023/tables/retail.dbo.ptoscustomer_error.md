# Retail UAT.dbo.PTOSCustomer_Error — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 67 lokasi source, 1 procedure, 15 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-alamatktp"></a>
## AlamatKTP

**Tipe dasar:** `VARCHAR(1000)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001067 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.AlamatKTP` : `VARCHAR(1000)` | `[Alamat KTP]` : `NVARCHAR(MAX)` | `[Alamat KTP]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKTP <- [Alamat KTP]` |
| IC-001095 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.AlamatKTP` : `VARCHAR(1000)` | `[Alamat KTP]` : `NVARCHAR(MAX)` | `[Alamat KTP]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKTP <- [Alamat KTP]` |
| IC-001127 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.AlamatKTP` : `VARCHAR(1000)` | `[Alamat KTP]` : `NVARCHAR(MAX)` | `[Alamat KTP]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKTP <- [Alamat KTP]` |
| IC-001154 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.AlamatKTP` : `VARCHAR(1000)` | `[Alamat KTP]` : `NVARCHAR(MAX)` | `[Alamat KTP]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKTP <- [Alamat KTP]` |
| IC-001171 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.AlamatKTP` : `VARCHAR(1000)` | `[Alamat KTP]` : `NVARCHAR(MAX)` | `[Alamat KTP]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKTP <- [Alamat KTP]` |

<a id="kolom-alamatkorespondensi"></a>
## AlamatKorespondensi

**Tipe dasar:** `VARCHAR(1000)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001068 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.AlamatKorespondensi` : `VARCHAR(1000)` | `[Alamat Korespondensi]` : `NVARCHAR(MAX)` | `[Alamat Korespondensi]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKorespondensi <- [Alamat Korespondensi]` |
| IC-001096 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.AlamatKorespondensi` : `VARCHAR(1000)` | `[Alamat Korespondensi]` : `NVARCHAR(MAX)` | `[Alamat Korespondensi]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKorespondensi <- [Alamat Korespondensi]` |
| IC-001128 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.AlamatKorespondensi` : `VARCHAR(1000)` | `[Alamat Korespondensi]` : `NVARCHAR(MAX)` | `[Alamat Korespondensi]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKorespondensi <- [Alamat Korespondensi]` |
| IC-001155 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.AlamatKorespondensi` : `VARCHAR(1000)` | `[Alamat Korespondensi]` : `NVARCHAR(MAX)` | `[Alamat Korespondensi]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKorespondensi <- [Alamat Korespondensi]` |
| IC-001172 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.AlamatKorespondensi` : `VARCHAR(1000)` | `[Alamat Korespondensi]` : `NVARCHAR(MAX)` | `[Alamat Korespondensi]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKorespondensi <- [Alamat Korespondensi]` |

<a id="kolom-email"></a>
## Email

**Tipe dasar:** `VARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001069 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.Email` : `VARCHAR(50)` | `[Alamat Email]` : `NVARCHAR(255)` | `[Alamat Email]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Email <- [Alamat Email]` |
| IC-001097 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.Email` : `VARCHAR(50)` | `[Alamat Email]` : `NVARCHAR(255)` | `[Alamat Email]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Email <- [Alamat Email]` |
| IC-001129 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.Email` : `VARCHAR(50)` | `[Alamat Email]` : `NVARCHAR(255)` | `[Alamat Email]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Email <- [Alamat Email]` |
| IC-001156 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.Email` : `VARCHAR(50)` | `[Alamat Email]` : `NVARCHAR(255)` | `[Alamat Email]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Email <- [Alamat Email]` |
| IC-001173 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.Email` : `VARCHAR(50)` | `[Alamat Email]` : `NVARCHAR(255)` | `[Alamat Email]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Email <- [Alamat Email]` |

<a id="kolom-jeniskelamin"></a>
## JenisKelamin

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001070 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.JenisKelamin` : `VARCHAR(20)` | `[Jenis Kelamin]` : `NVARCHAR(255)` | `[Jenis Kelamin]` → `VARCHAR(20)` | `retail.dbo.ptoscustomer_error.JenisKelamin <- [Jenis Kelamin]` |
| IC-001098 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.JenisKelamin` : `VARCHAR(20)` | `[Jenis Kelamin]` : `NVARCHAR(255)` | `[Jenis Kelamin]` → `VARCHAR(20)` | `retail.dbo.ptoscustomer_error.JenisKelamin <- [Jenis Kelamin]` |
| IC-001130 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.JenisKelamin` : `VARCHAR(20)` | `[Jenis Kelamin]` : `NVARCHAR(255)` | `[Jenis Kelamin]` → `VARCHAR(20)` | `retail.dbo.ptoscustomer_error.JenisKelamin <- [Jenis Kelamin]` |
| IC-001157 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.JenisKelamin` : `VARCHAR(20)` | `[Jenis Kelamin]` : `NVARCHAR(255)` | `[Jenis Kelamin]` → `VARCHAR(20)` | `retail.dbo.ptoscustomer_error.JenisKelamin <- [Jenis Kelamin]` |
| IC-001174 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.JenisKelamin` : `VARCHAR(20)` | `[Jenis Kelamin]` : `NVARCHAR(255)` | `[Jenis Kelamin]` → `VARCHAR(20)` | `retail.dbo.ptoscustomer_error.JenisKelamin <- [Jenis Kelamin]` |

<a id="kolom-npwp"></a>
## NPWP

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001071 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NPWP` : `VARCHAR(40)` | `[Nomor NPWP]` : `NVARCHAR(255)` | `[Nomor NPWP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NPWP <- [Nomor NPWP]` |
| IC-001099 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NPWP` : `VARCHAR(40)` | `[Nomor NPWP]` : `NVARCHAR(255)` | `[Nomor NPWP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NPWP <- [Nomor NPWP]` |
| IC-001131 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NPWP` : `VARCHAR(40)` | `[Nomor NPWP]` : `NVARCHAR(255)` | `[Nomor NPWP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NPWP <- [Nomor NPWP]` |
| IC-001158 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NPWP` : `VARCHAR(40)` | `[Nomor NPWP]` : `NVARCHAR(255)` | `[Nomor NPWP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NPWP <- [Nomor NPWP]` |
| IC-001175 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NPWP` : `VARCHAR(40)` | `[Nomor NPWP]` : `NVARCHAR(255)` | `[Nomor NPWP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NPWP <- [Nomor NPWP]` |

<a id="kolom-namanasabah"></a>
## NamaNasabah

**Tipe dasar:** `VARCHAR(95)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001072 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NamaNasabah` : `VARCHAR(95)` | `[Nama Nasabah]` : `NVARCHAR(255)` | `[Nama Nasabah]` → `VARCHAR(95)` | `retail.dbo.ptoscustomer_error.NamaNasabah <- [Nama Nasabah]` |
| IC-001100 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NamaNasabah` : `VARCHAR(95)` | `[Nama Nasabah]` : `NVARCHAR(255)` | `[Nama Nasabah]` → `VARCHAR(95)` | `retail.dbo.ptoscustomer_error.NamaNasabah <- [Nama Nasabah]` |
| IC-001132 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NamaNasabah` : `VARCHAR(95)` | `[Nama Nasabah]` : `NVARCHAR(255)` | `[Nama Nasabah]` → `VARCHAR(95)` | `retail.dbo.ptoscustomer_error.NamaNasabah <- [Nama Nasabah]` |
| IC-001159 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NamaNasabah` : `VARCHAR(95)` | `[Nama Nasabah]` : `NVARCHAR(255)` | `[Nama Nasabah]` → `VARCHAR(95)` | `retail.dbo.ptoscustomer_error.NamaNasabah <- [Nama Nasabah]` |
| IC-001176 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NamaNasabah` : `VARCHAR(95)` | `[Nama Nasabah]` : `NVARCHAR(255)` | `[Nama Nasabah]` → `VARCHAR(95)` | `retail.dbo.ptoscustomer_error.NamaNasabah <- [Nama Nasabah]` |

<a id="kolom-namaperusahaanefek"></a>
## NamaPerusahaanEfek

**Tipe dasar:** `VARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001073 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek` : `VARCHAR(50)` | `[Nama Perusahaan Efek]` : `NVARCHAR(255)` | `[Nama Perusahaan Efek]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek <- [Nama Perusahaan Efek]` |
| IC-001101 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek` : `VARCHAR(50)` | `[Nama Perusahaan Efek]` : `NVARCHAR(255)` | `[Nama Perusahaan Efek]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek <- [Nama Perusahaan Efek]` |
| IC-001133 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek` : `VARCHAR(50)` | `[Nama Perusahaan Efek]` : `NVARCHAR(255)` | `[Nama Perusahaan Efek]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek <- [Nama Perusahaan Efek]` |
| IC-001160 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek` : `VARCHAR(50)` | `[Nama Perusahaan Efek]` : `NVARCHAR(255)` | `[Nama Perusahaan Efek]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek <- [Nama Perusahaan Efek]` |
| IC-001177 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek` : `VARCHAR(50)` | `[Nama Perusahaan Efek]` : `NVARCHAR(255)` | `[Nama Perusahaan Efek]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek <- [Nama Perusahaan Efek]` |

<a id="kolom-nofax"></a>
## NoFax

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001074 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoFax` : `VARCHAR(40)` | `[Nomor Fax]` : `NVARCHAR(255)` | `[Nomor Fax]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoFax <- [Nomor Fax]` |
| IC-001102 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoFax` : `VARCHAR(40)` | `[Nomor Fax]` : `NVARCHAR(255)` | `[Nomor Fax]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoFax <- [Nomor Fax]` |
| IC-001134 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoFax` : `VARCHAR(40)` | `[Nomor Fax]` : `NVARCHAR(255)` | `[Nomor Fax]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoFax <- [Nomor Fax]` |
| IC-001161 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoFax` : `VARCHAR(40)` | `[Nomor Fax]` : `NVARCHAR(255)` | `[Nomor Fax]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoFax <- [Nomor Fax]` |
| IC-001178 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoFax` : `VARCHAR(40)` | `[Nomor Fax]` : `NVARCHAR(255)` | `[Nomor Fax]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoFax <- [Nomor Fax]` |

<a id="kolom-nohp"></a>
## NoHP

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001075 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoHP` : `VARCHAR(40)` | `[Nomor HP]` : `NVARCHAR (expression; length unknown)` | `[Nomor HP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoHP <- [Nomor HP]` |
| IC-001103 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoHP` : `VARCHAR(40)` | `[Nomor HP]` : `NVARCHAR (expression; length unknown)` | `[Nomor HP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoHP <- [Nomor HP]` |
| IC-001135 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoHP` : `VARCHAR(40)` | `[Nomor HP]` : `NVARCHAR (expression; length unknown)` | `[Nomor HP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoHP <- [Nomor HP]` |
| IC-001162 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoHP` : `VARCHAR(40)` | `[Nomor HP]` : `NVARCHAR (expression; length unknown)` | `[Nomor HP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoHP <- [Nomor HP]` |
| IC-001179 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoHP` : `VARCHAR(40)` | `[Nomor HP]` : `NVARCHAR(255)` | `[Nomor HP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoHP <- [Nomor HP]` |

<a id="kolom-noidentitas"></a>
## NoIdentitas

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001076 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoIdentitas` : `VARCHAR(40)` | `[Nomor KTP]` : `NVARCHAR(255)` | `[Nomor KTP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoIdentitas <- [Nomor KTP]` |
| IC-001104 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoIdentitas` : `VARCHAR(40)` | `[Nomor KTP]` : `NVARCHAR(255)` | `[Nomor KTP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoIdentitas <- [Nomor KTP]` |
| IC-001136 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoIdentitas` : `VARCHAR(40)` | `[Nomor KTP]` : `NVARCHAR(255)` | `[Nomor KTP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoIdentitas <- [Nomor KTP]` |
| IC-001163 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoIdentitas` : `VARCHAR(40)` | `[Nomor KTP]` : `NVARCHAR(255)` | `[Nomor KTP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoIdentitas <- [Nomor KTP]` |
| IC-001180 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoIdentitas` : `VARCHAR(40)` | `[Nomor KTP]` : `NVARCHAR(255)` | `[Nomor KTP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoIdentitas <- [Nomor KTP]` |

<a id="kolom-norekinvestorrdn"></a>
## NoRekInvestorRDN

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001181 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoRekInvestorRDN` : `VARCHAR(20)` | `[Nomor Rekening RDN]` : `NVARCHAR(255)` | `[Nomor Rekening RDN]` → `VARCHAR(20)` | `retail.dbo.ptoscustomer_error.NoRekInvestorRDN <- [Nomor Rekening RDN]` |

<a id="kolom-notlp"></a>
## NoTlp

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001077 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoTlp` : `VARCHAR(40)` | `[Nomor Telepon]` : `NVARCHAR (expression; length unknown)` | `[Nomor Telepon]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoTlp <- [Nomor Telepon]` |
| IC-001105 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoTlp` : `VARCHAR(40)` | `[Nomor Telepon]` : `NVARCHAR (expression; length unknown)` | `[Nomor Telepon]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoTlp <- [Nomor Telepon]` |
| IC-001137 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoTlp` : `VARCHAR(40)` | `[Nomor Telepon]` : `NVARCHAR (expression; length unknown)` | `[Nomor Telepon]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoTlp <- [Nomor Telepon]` |
| IC-001164 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoTlp` : `VARCHAR(40)` | `[Nomor Telepon]` : `NVARCHAR (expression; length unknown)` | `[Nomor Telepon]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoTlp <- [Nomor Telepon]` |
| IC-001182 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.NoTlp` : `VARCHAR(40)` | `[Nomor Telepon]` : `NVARCHAR(255)` | `[Nomor Telepon]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoTlp <- [Nomor Telepon]` |

<a id="kolom-pekerjaan"></a>
## Pekerjaan

**Tipe dasar:** `VARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001078 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.Pekerjaan` : `VARCHAR(50)` | `[Pekerjaan]` : `NVARCHAR(255)` | `[Pekerjaan]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Pekerjaan <- [Pekerjaan]` |
| IC-001106 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.Pekerjaan` : `VARCHAR(50)` | `[Pekerjaan]` : `NVARCHAR(255)` | `[Pekerjaan]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Pekerjaan <- [Pekerjaan]` |
| IC-001138 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.Pekerjaan` : `VARCHAR(50)` | `tmp.[Pekerjaan]` : `NVARCHAR(255)` | `tmp.[Pekerjaan]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Pekerjaan <- tmp.[Pekerjaan]` |
| IC-001165 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.Pekerjaan` : `VARCHAR(50)` | `[Pekerjaan]` : `NVARCHAR(255)` | `[Pekerjaan]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Pekerjaan <- [Pekerjaan]` |
| IC-001183 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.Pekerjaan` : `VARCHAR(50)` | `[Pekerjaan]` : `NVARCHAR(255)` | `[Pekerjaan]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Pekerjaan <- [Pekerjaan]` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001079 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L25) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptoscustomer_error.SID <- [SID]` |
| IC-001107 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L45) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptoscustomer_error.SID <- [SID]` |
| IC-001139 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L71) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.SID` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tmp.[SID]` → `VARCHAR(100)` | `retail.dbo.ptoscustomer_error.SID <- tmp.[SID]` |
| IC-001166 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:114](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L114) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptoscustomer_error.SID <- [SID]` |
| IC-001184 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptoscustomer_error.SID <- [SID]` |

<a id="kolom-warningmessage"></a>
## WarningMessage

**Tipe dasar:** `VARCHAR(250)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Error.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001185 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_error.WarningMessage` : `VARCHAR(250)` | `[Result]` : `NVARCHAR(MAX)` | `[Result]` → `VARCHAR(250)` | `retail.dbo.ptoscustomer_error.WarningMessage <- [Result]` |
