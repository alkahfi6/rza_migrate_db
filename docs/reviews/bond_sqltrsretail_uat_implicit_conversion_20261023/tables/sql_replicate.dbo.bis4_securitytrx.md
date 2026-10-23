# SQL_REPLICATE.dbo.BIS4_SecurityTrx — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 6 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-step-code"></a>
## step_code

**Tipe dasar:** `NVARCHAR(30)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.BIS4_SecurityTrx.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003593 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCheckStatusBIS4:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckStatusBIS4.sql#L22) | WHERE / P04 / IN_VALUE_TYPE | `[BIS4_SecurityTrx_v].step_code` : `NVARCHAR(30)` | `'X'` : `varchar` | `'X'` → `NVARCHAR` | `[BIS4_SecurityTrx_v].step_code in ('X','A','WM','PT/Q')` |
| IC-003596 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCheckStatusBIS4:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckStatusBIS4.sql#L32) | WHERE / P04 / IN_VALUE_TYPE | `biss.step_code` : `NVARCHAR(30)` | `'X'` : `varchar` | `'X'` → `NVARCHAR` | `biss.step_code in ('X','A','WM','PT/Q')` |

<a id="kolom-trn"></a>
## trn

**Tipe dasar:** `NVARCHAR(30)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.BIS4_SecurityTrx.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P06 — LIKE dan representasi karakter:** LIKE pada key numerik memerlukan representasi karakter. Equality numeric yang tervalidasi cocok untuk lookup tepat; prefix/substring membutuhkan requirement pencarian tersendiri, bukan cast kolom sebagai solusi index.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003591 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCheckStatusBIS4:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckStatusBIS4.sql#L18) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `trn` : `NVARCHAR(30)` | `#tmpM_Date.trn` : `varchar(60)` | `trn` → `varchar(60)` | `#tmpM_Date.trn <- trn` |
| IC-003592 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCheckStatusBIS4:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckStatusBIS4.sql#L21) | WHERE / P06 / LIKE | `trn` : `NVARCHAR(30)` | `'%-%'` : `varchar` | `'%-%'` → `NVARCHAR` | `trn like '%-%'` |
| IC-003594 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCheckStatusBIS4:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckStatusBIS4.sql#L29) | JOIN_ON / P02 / COMPARISON | `biss.trn` : `NVARCHAR(30)` | `md.trn` : `varchar(60)` | `md.trn` → `NVARCHAR` | `biss.trn = md.trn` |
| IC-003595 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCheckStatusBIS4:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckStatusBIS4.sql#L31) | WHERE / P06 / LIKE | `biss.trn` : `NVARCHAR(30)` | `'%-%'` : `varchar` | `'%-%'` → `NVARCHAR` | `biss.trn like '%-%'` |
