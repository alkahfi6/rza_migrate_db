# Retail UAT.dbo.TRSReportEmailStatusMap_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 1 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-category"></a>
## Category

**Tipe dasar:** `NVARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSReportEmailStatusMap_TR.sql#L2). **Existing leading index:** PK_TRSReportEmailStatusMap_TR.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006420 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L185) | JOIN_ON / P04 / COMPARISON | `SMap.Category` : `NVARCHAR(50)` | `'Status'` : `varchar` | `'Status'` → `NVARCHAR` | `SMap.Category = 'Status'` |
| IC-006422 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L188) | JOIN_ON / P04 / COMPARISON | `MMap.Category` : `NVARCHAR(50)` | `'Message'` : `varchar` | `'Message'` → `NVARCHAR` | `MMap.Category = 'Message'` |

<a id="kolom-description"></a>
## Description

**Tipe dasar:** `NVARCHAR(255)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSReportEmailStatusMap_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006417 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L166) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `MMap.Description` : `NVARCHAR(255)` | `#tmpDataEmail.ErrorMessage` : `varchar(40)` | `MMap.Description` → `varchar(40)` | `#tmpDataEmail.ErrorMessage <- MMap.Description` |
| IC-006418 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L166) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `SMap.Description` : `NVARCHAR(255)` | `#tmpDataEmail.Status` : `varchar(30)` | `SMap.Description` → `varchar(30)` | `#tmpDataEmail.Status <- SMap.Description` |

<a id="kolom-mapid"></a>
## MapId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSReportEmailStatusMap_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006419 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:184](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L184) | JOIN_ON / P11 / COMPARISON | `SMap.MapId` : `INT` | `TL.StatusId` : `SMALLINT` | `TL.StatusId` → `INT` | `TL.StatusId = SMap.MapId` |
| IC-006421 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L187) | JOIN_ON / P11 / COMPARISON | `MMap.MapId` : `INT` | `TL.MessageId` : `SMALLINT` | `TL.MessageId` → `INT` | `TL.MessageId = MMap.MapId` |
