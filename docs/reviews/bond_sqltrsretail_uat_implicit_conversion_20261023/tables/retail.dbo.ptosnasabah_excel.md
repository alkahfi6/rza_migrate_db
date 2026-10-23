# Retail UAT.dbo.PTOSNasabah_Excel — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 43 lokasi source, 2 procedure, 16 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 15 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-alamat-email"></a>
## Alamat Email

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001173 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Alamat Email]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.Email` : `VARCHAR(50)` | `[Alamat Email]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Email <- [Alamat Email]` |

<a id="kolom-alamat-ktp"></a>
## Alamat KTP

**Tipe dasar:** `NVARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001171 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Alamat KTP]` : `NVARCHAR(MAX)` | `retail.dbo.ptoscustomer_error.AlamatKTP` : `VARCHAR(1000)` | `[Alamat KTP]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKTP <- [Alamat KTP]` |

<a id="kolom-alamat-korespondensi"></a>
## Alamat Korespondensi

**Tipe dasar:** `NVARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001172 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Alamat Korespondensi]` : `NVARCHAR(MAX)` | `retail.dbo.ptoscustomer_error.AlamatKorespondensi` : `VARCHAR(1000)` | `[Alamat Korespondensi]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_error.AlamatKorespondensi <- [Alamat Korespondensi]` |

<a id="kolom-jenis-kelamin"></a>
## Jenis Kelamin

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001174 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Jenis Kelamin]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.JenisKelamin` : `VARCHAR(20)` | `[Jenis Kelamin]` → `VARCHAR(20)` | `retail.dbo.ptoscustomer_error.JenisKelamin <- [Jenis Kelamin]` |
| IC-001124 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L68) | WHERE / P04 / COMPARISON | `UPPER([Jenis Kelamin])` : `NVARCHAR (expression; length unknown)` | `'PRIA'` : `varchar` | `'PRIA'` → `NVARCHAR` | `UPPER([Jenis Kelamin]) != 'PRIA'` |
| IC-001125 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L68) | WHERE / P04 / COMPARISON | `UPPER([Jenis Kelamin])` : `NVARCHAR (expression; length unknown)` | `'WANITA'` : `varchar` | `'WANITA'` → `NVARCHAR` | `UPPER([Jenis Kelamin]) != 'WANITA'` |

<a id="kolom-nama-nasabah"></a>
## Nama Nasabah

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001176 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Nama Nasabah]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.NamaNasabah` : `VARCHAR(95)` | `[Nama Nasabah]` → `VARCHAR(95)` | `retail.dbo.ptoscustomer_error.NamaNasabah <- [Nama Nasabah]` |
| IC-012258 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L22) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '')` : `NVARCHAR (combined result; facets unknown)` | `'</td> <td class="table-data">'` : `varchar` | `'</td> <td class="table-data">'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">' + isnull([Nomor Rekening RDN], '') + '</td> <td class="table-data">' + isnull([Nama Nasabah], '') + '</td> <td class="table-data">'` |
| IC-012261 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:25](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L25) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nama Nasabah]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nama Nasabah], '')` |

<a id="kolom-nama-perusahaan-efek"></a>
## Nama Perusahaan Efek

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001177 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Nama Perusahaan Efek]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek` : `VARCHAR(50)` | `[Nama Perusahaan Efek]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.NamaPerusahaanEfek <- [Nama Perusahaan Efek]` |

<a id="kolom-nomor-fax"></a>
## Nomor Fax

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001178 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Nomor Fax]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.NoFax` : `VARCHAR(40)` | `[Nomor Fax]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoFax <- [Nomor Fax]` |

<a id="kolom-nomor-hp"></a>
## Nomor HP

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001179 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Nomor HP]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.NoHP` : `VARCHAR(40)` | `[Nomor HP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoHP <- [Nomor HP]` |

<a id="kolom-nomor-ktp"></a>
## Nomor KTP

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001180 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Nomor KTP]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.NoIdentitas` : `VARCHAR(40)` | `[Nomor KTP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoIdentitas <- [Nomor KTP]` |
| IC-001091 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L43) | WHERE / P04 / COMPARISON | `isnull([Nomor KTP], '')` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull([Nomor KTP], '') = ''` |
| IC-001093 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.PTOSProsesDataJobNasabah:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L43) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nomor KTP]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nomor KTP], '')` |

<a id="kolom-nomor-npwp"></a>
## Nomor NPWP

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001175 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Nomor NPWP]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.NPWP` : `VARCHAR(40)` | `[Nomor NPWP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NPWP <- [Nomor NPWP]` |

<a id="kolom-nomor-rekening-rdn"></a>
## Nomor Rekening RDN

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001181 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Nomor Rekening RDN]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.NoRekInvestorRDN` : `VARCHAR(20)` | `[Nomor Rekening RDN]` → `VARCHAR(20)` | `retail.dbo.ptoscustomer_error.NoRekInvestorRDN <- [Nomor Rekening RDN]` |
| IC-001063 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L23) | WHERE / P04 / COMPARISON | `isnull([Nomor Rekening RDN], '')` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull([Nomor Rekening RDN], '') = ''` |
| IC-001065 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.PTOSProsesDataJobNasabah:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L23) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nomor Rekening RDN]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nomor Rekening RDN], '')` |
| IC-012260 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L24) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Nomor Rekening RDN]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([Nomor Rekening RDN], '')` |

<a id="kolom-nomor-telepon"></a>
## Nomor Telepon

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001182 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Nomor Telepon]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.NoTlp` : `VARCHAR(40)` | `[Nomor Telepon]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_error.NoTlp <- [Nomor Telepon]` |

<a id="kolom-pekerjaan"></a>
## Pekerjaan

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001183 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Pekerjaan]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.Pekerjaan` : `VARCHAR(50)` | `[Pekerjaan]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_error.Pekerjaan <- [Pekerjaan]` |

<a id="kolom-result"></a>
## Result

**Tipe dasar:** `NVARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001185 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[Result]` : `NVARCHAR(MAX)` | `retail.dbo.ptoscustomer_error.WarningMessage` : `VARCHAR(250)` | `[Result]` → `VARCHAR(250)` | `retail.dbo.ptoscustomer_error.WarningMessage <- [Result]` |
| IC-001062 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L22) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Result` : `NVARCHAR(MAX)` | `'Failed - SID atau Nomor RDN kosong'` : `varchar` | `'Failed - SID atau Nomor RDN kosong'` → `NVARCHAR(MAX)` | `Result = 'Failed - SID atau Nomor RDN kosong'` |
| IC-001090 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L42) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Result` : `NVARCHAR(MAX)` | `'Failed - Nomor KTP atau Tanggal Lahir kosong'` : `varchar` | `'Failed - Nomor KTP atau Tanggal Lahir kosong'` → `NVARCHAR(MAX)` | `Result = 'Failed - Nomor KTP atau Tanggal Lahir kosong'` |
| IC-001123 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L67) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Result` : `NVARCHAR(MAX)` | `'Failed - Jenis Kelamin Tidak Teridentifikasi [WANITA/PRIA]'` : `varchar` | `'Failed - Jenis Kelamin Tidak Teridentifikasi [WANITA/PRIA]'` → `NVARCHAR(MAX)` | `Result = 'Failed - Jenis Kelamin Tidak Teridentifikasi [WANITA/PRIA]'` |
| IC-001153 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:109](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L109) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ex.Result` : `NVARCHAR(MAX)` | `'Failed - CIFNo tidak di temukan'` : `varchar` | `'Failed - CIFNo tidak di temukan'` → `NVARCHAR(MAX)` | `ex.Result = 'Failed - CIFNo tidak di temukan'` |
| IC-001168 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L134) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ex.Result` : `NVARCHAR(MAX)` | `'Failed - Nasabah Sudah Terdaftar [CIFNo dan NoRekInvestorRDN]'` : `varchar` | `'Failed - Nasabah Sudah Terdaftar [CIFNo dan NoRekInvestorRDN]'` → `NVARCHAR(MAX)` | `ex.Result = 'Failed - Nasabah Sudah Terdaftar [CIFNo dan NoRekInvestorRDN]'` |
| IC-001186 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L151) | WHERE / P04 / COMPARISON | `Result` : `NVARCHAR(MAX)` | `'Failed - Nasabah Sudah Terdaftar [CIFNo dan NoRekInvestorRDN]'` : `varchar` | `'Failed - Nasabah Sudah Terdaftar [CIFNo dan NoRekInvestorRDN]'` → `NVARCHAR` | `Result = 'Failed - Nasabah Sudah Terdaftar [CIFNo dan NoRekInvestorRDN]'` |
| IC-001187 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L154) | WHERE / P06 / LIKE | `Result` : `NVARCHAR(MAX)` | `'Failed%'` : `varchar` | `'Failed%'` → `NVARCHAR` | `Result LIKE 'Failed%'` |
| IC-001189 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:160](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L160) | WHERE / P06 / LIKE | `Result` : `NVARCHAR(MAX)` | `'Failed%'` : `varchar` | `'Failed%'` → `NVARCHAR` | `Result LIKE 'Failed%'` |
| IC-001231 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:235](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L235) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tmp.Result` : `NVARCHAR(MAX)` | `'Failed - Insert to Table Master'` : `varchar` | `'Failed - Insert to Table Master'` → `NVARCHAR(MAX)` | `tmp.Result = 'Failed - Insert to Table Master'` |
| IC-012262 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L28) | WHERE / P04 / COMPARISON | `isnull(Result, '')` : `NVARCHAR(MAX)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull(Result, '') != ''` |
| IC-012263 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L28) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `Result` : `NVARCHAR(MAX)` | `''` : `varchar` | `''` → `NVARCHAR(MAX)` | `isnull(Result, '')` |

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001184 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.PTOSProsesDataJobNasabah:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L145) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `[SID]` : `NVARCHAR(255)` | `retail.dbo.ptoscustomer_error.SID` : `VARCHAR(100)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptoscustomer_error.SID <- [SID]` |
| IC-001064 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L23) | WHERE / P04 / COMPARISON | `isnull([SID], '')` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull([SID], '') = ''` |
| IC-001066 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.PTOSProsesDataJobNasabah:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L23) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[SID]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([SID], '')` |
| IC-001126 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L69) | WHERE / P04 / COMPARISON | `UPPER(LEFT([SID],2))` : `NVARCHAR (expression; length unknown)` | `'ID'` : `varchar` | `'ID'` → `NVARCHAR` | `UPPER(LEFT([SID],2)) = 'ID'` |
| IC-001232 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L237) | JOIN_ON / P02 / COMPARISON | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` : `VARCHAR(100)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-012257 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSubmitLogSentMailResult:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L22) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '')` : `NVARCHAR (combined result; facets unknown)` | `'</td> <td class="table-data">'` : `varchar` | `'</td> <td class="table-data">'` → `NVARCHAR` | `@cBodyEmail + '<tr class="table-row"> <td class="table-data">' + isnull([SID], '') +'</td> <td class="table-data">'` |
| IC-012259 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitLogSentMailResult:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L23) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[SID]` : `NVARCHAR(255)` | `''` : `varchar` | `''` → `NVARCHAR(255)` | `isnull([SID], '')` |

<a id="kolom-tanggal-lahir"></a>
## Tanggal Lahir

**Tipe dasar:** `NVARCHAR(15)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSNasabah_Excel.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001092 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.PTOSProsesDataJobNasabah:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L43) | WHERE / P04 / COMPARISON | `isnull([Tanggal Lahir], '')` : `NVARCHAR(15)` | `''` : `varchar` | `''` → `NVARCHAR` | `isnull([Tanggal Lahir], '') = ''` |
| IC-001094 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.PTOSProsesDataJobNasabah:43](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L43) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `[Tanggal Lahir]` : `NVARCHAR(15)` | `''` : `varchar` | `''` → `NVARCHAR(15)` | `isnull([Tanggal Lahir], '')` |
