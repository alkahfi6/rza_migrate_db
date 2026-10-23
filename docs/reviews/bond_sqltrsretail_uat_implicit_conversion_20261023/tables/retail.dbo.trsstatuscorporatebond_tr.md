# Retail UAT.dbo.TRSStatusCorporateBond_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 9 lokasi source, 7 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-enable"></a>
## Enable

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSStatusCorporateBond_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000662 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMGetListProductSekunder:247](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMGetListProductSekunder.sql#L247) | WHERE / P11 / COMPARISON | `sts.[Enable]` : `BIT` | `0` : `int` | `sts.[Enable]` → `INT` | `sts.[Enable] = 0` |
| IC-000757 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.OMTrxSecPreBook:483](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OMTrxSecPreBook.sql#L483) | WHERE / P11 / COMPARISON | `[Enable]` : `BIT` | `0` : `int` | `[Enable]` → `INT` | `[Enable] = 0` |
| IC-004326 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSInsertSecurityTransaction_TT_Original:931](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSInsertSecurityTransaction_TT_Original.sql#L931) | WHERE / P11 / COMPARISON | `[Enable]` : `BIT` | `0` : `int` | `[Enable]` → `INT` | `[Enable] = 0` |
| IC-005356 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateDataTransaksiSuratBerharga:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateDataTransaksiSuratBerharga.sql#L91) | WHERE / P11 / COMPARISON | `[Enable]` : `BIT` | `0` : `int` | `[Enable]` → `INT` | `[Enable] = 0` |
| IC-004454 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSMaintenanceParamHargaModal:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L135) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsstatuscorporatebond_tr.Enable` : `BIT` | `1` : `int` | `1` → `BIT` | `retail.dbo.trsstatuscorporatebond_tr.Enable <- 1` |
| IC-006395 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateStatusCorporateBonds:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusCorporateBonds.sql#L61) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `tr.[Enable]` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(tr.[Enable],0)` |
| IC-012360 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitStatusCorporateBonds:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitStatusCorporateBonds.sql#L107) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.[Enable]` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(par.[Enable],0)` |
| IC-012363 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitStatusCorporateBonds:123](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitStatusCorporateBonds.sql#L123) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.[Enable]` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(par.[Enable],0)` |
| IC-012366 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSSubmitStatusCorporateBonds:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitStatusCorporateBonds.sql#L130) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `par.[Enable]` : `BIT` | `0` : `int` | `0` → `BIT` | `ISNULL(par.[Enable],0)` |
