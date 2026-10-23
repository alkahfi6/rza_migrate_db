# Retail UAT.dbo.trs_list_source_reference_lelang — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 1 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-is-internal"></a>
## is_internal

**Tipe dasar:** `[bit]`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.trs_list_source_reference_lelang.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003008 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SecurityGetXMLLelang:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SecurityGetXMLLelang.sql#L77) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trs_list_source_reference_lelang.is_internal` : `[bit]` | `CASE @pcInternal WHEN 'Y' THEN 1 ELSE 0 END` : `INT (combined result; facets unknown)` | `CASE @pcInternal WHEN 'Y' THEN 1 ELSE 0 END` → `[bit]` | `retail.dbo.trs_list_source_reference_lelang.is_internal <- CASE @pcInternal WHEN 'Y' THEN 1 ELSE 0 END` |
