# Retail UAT.dbo.TRSTransactionFeeLog_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSTransactionFeeLog_TR.sql#L12). **Existing leading index:** TRSTransactionFeeLog_TR_NN1.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007703 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessSettlementTransaksiBankJual:146](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankJual.sql#L146) | JOIN_ON / P12 / IN_VALUE_TYPE | `tfltr.Status` : `CHAR(1)` | `'S'` : `varchar` | `tfltr.Status` → `VARCHAR` | `tfltr.Status IN ('S', 'P')` |
| IC-012545 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSUpdatePostingResult:112](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L112) | WHERE / P12 / COMPARISON | `Status` : `CHAR(1)` | `'P'` : `varchar` | `Status` → `VARCHAR` | `Status = 'P'` |
| IC-007706 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSettlementTransaksiBankJual:200](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSettlementTransaksiBankJual.sql#L200) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trstransactionfeelog_tr.Status` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trstransactionfeelog_tr.Status <- 'P'` |
