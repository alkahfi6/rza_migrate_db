# SQL_REPLICATE.dbo.PBCustomerFlagging_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-userid"></a>
## UserId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.PBCustomerFlagging_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003999 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSFailedTransactionPrimaryPreBook:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSFailedTransactionPrimaryPreBook.sql#L69) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pb.UserId` : `INT` | `tmp.NIKMarketingFlagging` : `varchar(100)` | `pb.UserId` → `varchar(100)` | `tmp.NIKMarketingFlagging = pb.UserId` |
| IC-010597 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:325](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L325) | JOIN_ON / P01 / COMPARISON | `pbcf.UserId` : `INT` | `um.EMPLID` : `VARCHAR(11)` | `um.EMPLID` → `INT` | `pbcf.UserId = um.EMPLID` |
| IC-010600 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo:381](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOutstandingTaxAmnestyPerSecAccNo.sql#L381) | JOIN_ON / P01 / COMPARISON | `pbcf.UserId` : `INT` | `um.EMPLID` : `VARCHAR(11)` | `um.EMPLID` → `INT` | `pbcf.UserId = um.EMPLID` |
