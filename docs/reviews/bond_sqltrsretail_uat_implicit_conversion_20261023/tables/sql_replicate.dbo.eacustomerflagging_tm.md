# SQL_REPLICATE.dbo.EACustomerFlagging_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 3 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-moid"></a>
## MOId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.EACustomerFlagging_TM.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003998 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSFailedTransactionPrimaryPreBook:63](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSFailedTransactionPrimaryPreBook.sql#L63) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ea.MOId` : `INT` | `tmp.NIKMarketingFlagging` : `varchar(100)` | `ea.MOId` → `varchar(100)` | `tmp.NIKMarketingFlagging = ea.MOId` |
| IC-004774 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSOBLExpiredKitasRPT:108](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSOBLExpiredKitasRPT.sql#L108) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ISNULL(EA.MOId, PB.UserId)` : `INT` | `''` : `varchar` | `''` → `INT` | `ISNULL(ISNULL(EA.MOId, PB.UserId), '')` |
| IC-010598 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:327](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L327) | JOIN_ON / P01 / COMPARISON | `eacf.MOId` : `INT` | `unm.EMPLID` : `VARCHAR(11)` | `unm.EMPLID` → `INT` | `eacf.MOId = unm.EMPLID` |
| IC-010601 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:383](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L383) | JOIN_ON / P01 / COMPARISON | `eacf.MOId` : `INT` | `unm.EMPLID` : `VARCHAR(11)` | `unm.EMPLID` → `INT` | `eacf.MOId = unm.EMPLID` |
