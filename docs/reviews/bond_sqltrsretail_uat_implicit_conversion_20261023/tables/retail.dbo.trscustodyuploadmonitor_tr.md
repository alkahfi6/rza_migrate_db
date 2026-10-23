# Retail UAT.dbo.TRSCustodyUploadMonitor_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 3 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-completeflag"></a>
## CompleteFlag

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustodyUploadMonitor_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003672 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCustodyUploadProcessAuth:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyUploadProcessAuth.sql#L139) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `CompleteFlag` : `BIT` | `1` : `int` | `1` → `BIT` | `CompleteFlag = 1` |

<a id="kolom-periodtype"></a>
## PeriodType

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustodyUploadMonitor_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003661 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCustodyMonitorFile:13](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyMonitorFile.sql#L13) | WHERE / P12 / COMPARISON | `a.PeriodType` : `CHAR(1)` | `'D'` : `varchar` | `a.PeriodType` → `VARCHAR` | `a.PeriodType = 'D'` |
| IC-003663 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCustodyMonitorFile:24](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyMonitorFile.sql#L24) | WHERE / P12 / COMPARISON | `a.PeriodType` : `CHAR(1)` | `'M'` : `varchar` | `a.PeriodType` → `VARCHAR` | `a.PeriodType = 'M'` |
| IC-003667 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCustodyUploadJobs:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyUploadJobs.sql#L37) | WHERE / P12 / COMPARISON | `PeriodType` : `CHAR(1)` | `'D'` : `varchar` | `PeriodType` → `VARCHAR` | `PeriodType = 'D'` |
| IC-003668 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCustodyUploadJobs:45](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyUploadJobs.sql#L45) | WHERE / P12 / COMPARISON | `PeriodType` : `CHAR(1)` | `'M'` : `varchar` | `PeriodType` → `VARCHAR` | `PeriodType = 'M'` |

<a id="kolom-uploadnik"></a>
## UploadNik

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustodyUploadMonitor_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003660 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCustodyMonitorFile:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyMonitorFile.sql#L12) | JOIN_ON / P11 / COMPARISON | `a.UploadNik` : `INT` | `c.nik` : `DECIMAL (combined result; facets unknown)` | `a.UploadNik` → `DECIMAL` | `a.UploadNik = c.nik` |
| IC-003662 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCustodyMonitorFile:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCustodyMonitorFile.sql#L23) | JOIN_ON / P11 / COMPARISON | `a.UploadNik` : `INT` | `c.nik` : `DECIMAL (combined result; facets unknown)` | `a.UploadNik` → `DECIMAL` | `a.UploadNik = c.nik` |
