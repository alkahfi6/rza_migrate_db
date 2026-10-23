# Retail UAT.dbo.trs_getaccount_session — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 1 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-insert-date"></a>
## INSERT_DATE

**Tipe dasar:** `SMALLDATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.trs_getaccount_session.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000477 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GetAccountInfo:19](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L19) | WRITE_ASSIGNMENT / P08 / INSERT_VALUES | `retail.dbo.trs_getaccount_session.INSERT_DATE` : `SMALLDATETIME` | `GETDATE()` : `datetime` | `GETDATE()` → `SMALLDATETIME` | `retail.dbo.trs_getaccount_session.INSERT_DATE <- GETDATE()` |
