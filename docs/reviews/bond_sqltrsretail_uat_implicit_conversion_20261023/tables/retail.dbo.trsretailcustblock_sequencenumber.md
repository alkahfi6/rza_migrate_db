# Retail UAT.dbo.TRSRetailCustBlock_SequenceNumber — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-increament"></a>
## Increament

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSRetailCustBlock_SequenceNumber.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013678 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_RetailCustBlock_SequenceNumber:14](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_RetailCustBlock_SequenceNumber.sql#L14) | WRITE_ASSIGNMENT / P08 / INSERT_VALUES | `retail.dbo.trsretailcustblock_sequencenumber.Increament` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `retail.dbo.trsretailcustblock_sequencenumber.Increament <- 1` |
| IC-013679 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_RetailCustBlock_SequenceNumber:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_RetailCustBlock_SequenceNumber.sql#L28) | DML_CONDITION / P05 / BINARY_Add | `Increament` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `Increament + 1` |
| IC-013680 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.trs_RetailCustBlock_SequenceNumber:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_RetailCustBlock_SequenceNumber.sql#L33) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Increament` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `Increament = 1` |
| IC-013681 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.trs_RetailCustBlock_SequenceNumber:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_RetailCustBlock_SequenceNumber.sql#L38) | DML_CONDITION / P05 / BINARY_Add | `Increament` : `BIGINT` | `1` : `int` | `1` → `BIGINT` | `Increament + 1` |
