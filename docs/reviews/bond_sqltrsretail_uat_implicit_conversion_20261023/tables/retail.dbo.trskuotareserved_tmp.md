# Retail UAT.dbo.TRSKuotaReserved_TMP — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-amount"></a>
## Amount

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSKuotaReserved_TMP.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003564 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCalculateTransactionSekunder:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateTransactionSekunder.sql#L61) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sum(-1 * st.Amount)` : `MONEY (aggregate; facets unknown)` | `#tmp_total.TotalAmount` : `decimal(20,4)` | `sum(-1 * st.Amount)` → `decimal(20,4)` | `#tmp_total.TotalAmount <- sum(-1 * st.Amount)` |
| IC-003565 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateTransactionSekunder:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateTransactionSekunder.sql#L63) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `st.Amount` : `MONEY` | `-1` : `int` | `-1` → `MONEY` | `-1 * st.Amount` |
