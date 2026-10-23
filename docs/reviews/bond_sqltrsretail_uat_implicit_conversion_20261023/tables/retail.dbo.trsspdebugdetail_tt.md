# Retail UAT.dbo.TRSSPDebugDetail_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-paramvalue"></a>
## ParamValue

**Tipe dasar:** `NVARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSSPDebugDetail_TT.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004300 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:337](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L337) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsspdebugdetail_tt.ParamValue` : `NVARCHAR(MAX)` | `@pbCommitImmediately` : `bit` | `@pbCommitImmediately` → `NVARCHAR(MAX)` | `retail.dbo.trsspdebugdetail_tt.ParamValue <- @pbCommitImmediately` |
| IC-012937 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.sp_sbn_master_nasabah_datagrid:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.sp_sbn_master_nasabah_datagrid.sql#L22) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsspdebugdetail_tt.ParamValue` : `NVARCHAR(MAX)` | `@pcAPICode` : `VARCHAR(50)` | `@pcAPICode` → `NVARCHAR(MAX)` | `retail.dbo.trsspdebugdetail_tt.ParamValue <- @pcAPICode` |
