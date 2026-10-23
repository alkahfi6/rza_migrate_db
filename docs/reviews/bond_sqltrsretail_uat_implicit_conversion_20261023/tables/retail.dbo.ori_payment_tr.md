# Retail UAT.dbo.ORI_Payment_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-xmlresponsepayinquiry"></a>
## XmlResponsePayInquiry

**Tipe dasar:** `NVARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Payment_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002859 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessORIPayment:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessORIPayment.sql#L46) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_payment_tr.XmlResponsePayInquiry` : `NVARCHAR(MAX)` | `CASE WHEN @pcPaymentType = 'INQ' THEN @pcXmlData ELSE NULL END` : `VARCHAR (combined result; facets unknown)` | `CASE WHEN @pcPaymentType = 'INQ' THEN @pcXmlData ELSE NULL END` → `NVARCHAR(MAX)` | `retail.dbo.ori_payment_tr.XmlResponsePayInquiry <- CASE WHEN @pcPaymentType = 'INQ' THEN @pcXmlData ELSE NULL END` |

<a id="kolom-xmlresponsepayprocess"></a>
## XmlResponsePayProcess

**Tipe dasar:** `NVARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORI_Payment_TR.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002860 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedProcessORIPayment:46](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessORIPayment.sql#L46) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ori_payment_tr.XmlResponsePayProcess` : `NVARCHAR(MAX)` | `CASE WHEN @pcPaymentType = 'PAY' THEN @pcXmlData ELSE NULL END` : `VARCHAR (combined result; facets unknown)` | `CASE WHEN @pcPaymentType = 'PAY' THEN @pcXmlData ELSE NULL END` → `NVARCHAR(MAX)` | `retail.dbo.ori_payment_tr.XmlResponsePayProcess <- CASE WHEN @pcPaymentType = 'PAY' THEN @pcXmlData ELSE NULL END` |
