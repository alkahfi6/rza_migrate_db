# Retail UAT.dbo.CorporateBondLogSentMailKonfirmasi_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-email"></a>
## Email

**Tipe dasar:** `VARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.CorporateBondLogSentMailKonfirmasi_TH.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012228 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSentMailNotifikasiCorporateBond:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentMailNotifikasiCorporateBond.sql#L69) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `th.Email` : `VARCHAR(MAX)` | `#tmpSentMailCorpBond.ToAddress` : `nvarchar(max)` | `th.Email` → `nvarchar(max)` | `#tmpSentMailCorpBond.ToAddress <- th.Email` |
| IC-012227 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSentMailNotifikasiCorporateBond:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentMailNotifikasiCorporateBond.sql#L54) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.corporatebondlogsentmailkonfirmasi_th.Email` : `VARCHAR(MAX)` | `@cDefaultToAddress` : `nvarchar(max)` | `@cDefaultToAddress` → `VARCHAR(MAX)` | `retail.dbo.corporatebondlogsentmailkonfirmasi_th.Email <- @cDefaultToAddress` |
