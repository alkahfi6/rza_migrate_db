# Retail UAT.dbo.TRSNTIParameter_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 4 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-value"></a>
## Value

**Tipe dasar:** `VARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSNTIParameter_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003349 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSApproveSID:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveSID.sql#L172) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ev.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-003350 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSApproveSID:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveSID.sql#L172) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mx.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-006706 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessApprovedUploadSRESID:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSRESID.sql#L169) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ev.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-006707 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSProcessApprovedUploadSRESID:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSRESID.sql#L169) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mx.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
