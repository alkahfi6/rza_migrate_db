# Retail UAT.dbo.REPOLogSentMailKonfirmasi_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 2 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-email"></a>
## Email

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOLogSentMailKonfirmasi_TH.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012229 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSentMailNotifikasiREPO:32](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentMailNotifikasiREPO.sql#L32) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `th.Email` : `VARCHAR(40)` | `#tmpSentMail.ToAddress` : `nvarchar(max)` | `th.Email` → `nvarchar(max)` | `#tmpSentMail.ToAddress <- th.Email` |

<a id="kolom-haveemail"></a>
## HaveEmail

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOLogSentMailKonfirmasi_TH.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004955 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPDFKonfirmasiREPO:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPDFKonfirmasiREPO.sql#L54) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repologsentmailkonfirmasi_th.HaveEmail` : `BIT` | `case when isnull(Email,'') != '' then 1 else 0 end` : `INT (combined result; facets unknown)` | `case when isnull(Email,'') != '' then 1 else 0 end` → `BIT` | `retail.dbo.repologsentmailkonfirmasi_th.HaveEmail <- case when isnull(Email,'') != '' then 1 else 0 end` |

<a id="kolom-processtime"></a>
## ProcessTime

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOLogSentMailKonfirmasi_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012230 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSentMailNotifikasiREPO:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentMailNotifikasiREPO.sql#L46) | WHERE / P04 / COMPARISON | `th.ProcessTime` : `DATETIME` | `@cTodayDate` : `varchar(8)` | `@cTodayDate` → `DATETIME` | `th.ProcessTime >= @cTodayDate` |
| IC-012231 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSentMailNotifikasiREPO:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentMailNotifikasiREPO.sql#L70) | WHERE / P04 / COMPARISON | `th.ProcessTime` : `DATETIME` | `@cTodayDate` : `varchar(8)` | `@cTodayDate` → `DATETIME` | `th.ProcessTime >= @cTodayDate` |
| IC-012232 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSentMailNotifikasiREPO:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentMailNotifikasiREPO.sql#L89) | WHERE / P04 / COMPARISON | `th.ProcessTime` : `DATETIME` | `@cTodayDate` : `varchar(8)` | `@cTodayDate` → `DATETIME` | `th.ProcessTime >= @cTodayDate` |
