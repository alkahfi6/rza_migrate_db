# Retail UAT.dbo.ORA_TRN_HDR_DBF — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 1 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-m-nb"></a>
## M_NB

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORA_TRN_HDR_DBF.sql#L3). **Existing leading index:** ORA_TRN_HDR_DBF_NN1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000108 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.CTMCopyDataBrokerRepToBroker:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CTMCopyDataBrokerRepToBroker.sql#L9) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ora_trn_hdr_dbf.M_NB` : `INT` | `rtrim(rep.M_NB)` : `VARCHAR (expression; length unknown)` | `rtrim(rep.M_NB)` → `INT` | `retail.dbo.ora_trn_hdr_dbf.M_NB <- rtrim(rep.M_NB)` |
