# Retail UAT.dbo.CIFReplicatePending_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-action"></a>
## Action

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CIFReplicatePending_TT.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004744 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSMxSyncCIFGenerateFile:5](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMxSyncCIFGenerateFile.sql#L5) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `[Action]` : `CHAR(1)` | `right(replicate('0', 19) + convert(varchar,[CIFNo]), 19) + right(replicate(' ', 40) + convert(varchar(40),[Guid]), 40) + right(replicate(' ', 40) + convert(varchar(40),[MxLabel]), 40)` : `VARCHAR (combined result; facets unknown)` | `[Action]` → `VARCHAR` | `right(replicate('0', 19) + convert(varchar,[CIFNo]), 19) + right(replicate(' ', 40) + convert(varchar(40),[Guid]), 40) + right(replicate(' ', 40) + convert(varchar(40),[MxLabel]), 40) + [Action]` |
| IC-004745 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMxSyncCIFOnlineInsert:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMxSyncCIFOnlineInsert.sql#L20) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.cifreplicatepending_tt.Action` : `CHAR(1)` | `'I'` : `varchar` | `'I'` → `CHAR(1)` | `retail.dbo.cifreplicatepending_tt.Action <- 'I'` |

<a id="kolom-cifno"></a>
## CIFNo

**Tipe dasar:** `DECIMAL(19,0)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CIFReplicatePending_TT.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004746 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMxSyncCIFOnlineInsert:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMxSyncCIFOnlineInsert.sql#L20) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.cifreplicatepending_tt.CIFNo` : `DECIMAL(19,0)` | `@pcCIFNo` : `varchar(19)` | `@pcCIFNo` → `DECIMAL(19,0)` | `retail.dbo.cifreplicatepending_tt.CIFNo <- @pcCIFNo` |

<a id="kolom-guid"></a>
## Guid

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CIFReplicatePending_TT.sql#L2). **Existing leading index:** CIFReplicatePending_TT_CN1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004747 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMxSyncCIFOnlineInsert:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMxSyncCIFOnlineInsert.sql#L20) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.cifreplicatepending_tt.Guid` : `VARCHAR(40)` | `@uGuid` : `uniqueidentifier` | `@uGuid` → `VARCHAR(40)` | `retail.dbo.cifreplicatepending_tt.Guid <- @uGuid` |
