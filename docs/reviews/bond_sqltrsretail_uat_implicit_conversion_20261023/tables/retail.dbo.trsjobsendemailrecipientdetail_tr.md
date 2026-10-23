# Retail UAT.dbo.TRSJobSendEmailRecipientDetail_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 3 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-paramvalue"></a>
## ParamValue

**Tipe dasar:** `VARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSJobSendEmailRecipientDetail_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002980 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH1:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH1.sql#L205) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemailrecipientdetail_tr.ParamValue` : `VARCHAR(MAX)` | `sb.Body` : `NVARCHAR(MAX)` | `sb.Body` → `VARCHAR(MAX)` | `retail.dbo.trsjobsendemailrecipientdetail_tr.ParamValue <- sb.Body` |
| IC-002991 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH3:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH3.sql#L205) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemailrecipientdetail_tr.ParamValue` : `VARCHAR(MAX)` | `sb.Body` : `NVARCHAR(MAX)` | `sb.Body` → `VARCHAR(MAX)` | `retail.dbo.trsjobsendemailrecipientdetail_tr.ParamValue <- sb.Body` |
| IC-003002 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH7:205](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH7.sql#L205) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemailrecipientdetail_tr.ParamValue` : `VARCHAR(MAX)` | `sb.Body` : `NVARCHAR(MAX)` | `sb.Body` → `VARCHAR(MAX)` | `retail.dbo.trsjobsendemailrecipientdetail_tr.ParamValue <- sb.Body` |
