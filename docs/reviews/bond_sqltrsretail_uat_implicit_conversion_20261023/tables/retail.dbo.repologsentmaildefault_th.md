# Retail UAT.dbo.REPOLogSentMailDefault_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-haveemail"></a>
## HaveEmail

**Tipe dasar:** `[bit]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOLogSentMailDefault_TH.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004952 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPDFDefaultREPO:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPDFDefaultREPO.sql#L56) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repologsentmaildefault_th.HaveEmail` : `[bit]` | `case when isnull(Email,'') != '' then 1 else 0 end` : `INT (combined result; facets unknown)` | `case when isnull(Email,'') != '' then 1 else 0 end` → `[bit]` | `retail.dbo.repologsentmaildefault_th.HaveEmail <- case when isnull(Email,'') != '' then 1 else 0 end` |

<a id="kolom-maturitydate"></a>
## MaturityDate

**Tipe dasar:** `[datetime]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.REPOLogSentMailDefault_TH.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004953 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSPDFDefaultREPO:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPDFDefaultREPO.sql#L56) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.repologsentmaildefault_th.MaturityDate` : `[datetime]` | `MaturityDate` : `VARCHAR(12)` | `MaturityDate` → `[datetime]` | `retail.dbo.repologsentmaildefault_th.MaturityDate <- MaturityDate` |
