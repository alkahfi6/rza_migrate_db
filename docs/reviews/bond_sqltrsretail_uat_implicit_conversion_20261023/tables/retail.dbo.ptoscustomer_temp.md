# Retail UAT.dbo.PTOSCustomer_Temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Temp.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005746 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSCustomerAuthen:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSCustomerAuthen.sql#L33) | JOIN_ON / P01 / COMPARISON | `tmp.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `tmp.InsertedBy = ndme.EMPLID` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_Temp.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007113 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomer_TM:144](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomer_TM.sql#L144) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_temp.Status` : `SMALLINT` | `@pnStatus` : `INT` | `@pnStatus` → `SMALLINT` | `retail.dbo.ptoscustomer_temp.Status <- @pnStatus` |
