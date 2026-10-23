# Retail UAT.dbo.TRSParamThreshold_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-equsdthreshold"></a>
## EqUSDThreshold

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamThreshold_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012039 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamThreshold:127](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L127) | WHERE / P11 / COMPARISON | `tm.EqUSDThreshold` : `BIGINT` | `tmp.EqUSDThreshold` : `decimal(25,5)` | `tm.EqUSDThreshold` → `DECIMAL` | `tmp.EqUSDThreshold <> tm.EqUSDThreshold` |
| IC-012041 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSaveParamThreshold:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L140) | WHERE / P11 / COMPARISON | `tm.EqUSDThreshold` : `BIGINT` | `tmp.EqUSDThreshold` : `decimal(25,5)` | `tm.EqUSDThreshold` → `DECIMAL` | `tmp.EqUSDThreshold <> tm.EqUSDThreshold` |
| IC-012036 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamThreshold:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L98) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsparamthreshold_tm.EqUSDThreshold` : `BIGINT` | `tmp.EqUSDThreshold` : `decimal(25,5)` | `tmp.EqUSDThreshold` → `BIGINT` | `retail.dbo.trsparamthreshold_tm.EqUSDThreshold <- tmp.EqUSDThreshold` |
| IC-012040 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveParamThreshold:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveParamThreshold.sql#L133) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.EqUSDThreshold` : `BIGINT` | `tmp.EqUSDThreshold` : `decimal(25,5)` | `tmp.EqUSDThreshold` → `BIGINT` | `tm.EqUSDThreshold = tmp.EqUSDThreshold` |
