# Retail UAT.dbo.RMMSecurityTransaction_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-processcode"></a>
## ProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.RMMSecurityTransaction_TH.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000645 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMDeleteProcessInputSecurityMaster:141](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMDeleteProcessInputSecurityMaster.sql#L141) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.rmmsecuritytransaction_th.ProcessCode` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.rmmsecuritytransaction_th.ProcessCode <- 'D'` |
| IC-000646 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.OMDeleteProcessInputSecurityMaster:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMDeleteProcessInputSecurityMaster.sql#L196) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.rmmsecuritytransaction_th.ProcessCode` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.rmmsecuritytransaction_th.ProcessCode <- 'D'` |
| IC-012978 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_DeleteSecurityTransaction_TT:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_DeleteSecurityTransaction_TT.sql#L167) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.rmmsecuritytransaction_th.ProcessCode` : `CHAR(1)` | `'D'` : `varchar` | `'D'` → `CHAR(1)` | `retail.dbo.rmmsecuritytransaction_th.ProcessCode <- 'D'` |
