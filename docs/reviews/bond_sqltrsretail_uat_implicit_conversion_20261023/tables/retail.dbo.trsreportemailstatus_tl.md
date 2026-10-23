# Retail UAT.dbo.TRSReportEmailStatus_TL — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 1 procedure, 4 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifno"></a>
## CifNo

**Tipe dasar:** `DECIMAL(19)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSReportEmailStatus_TL.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006415 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L166) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `TL.CifNo` : `DECIMAL(19)` | `#tmpDataEmail.CIF` : `varchar(19)` | `TL.CifNo` → `varchar(19)` | `#tmpDataEmail.CIF <- TL.CifNo` |

<a id="kolom-messageid"></a>
## MessageId

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSReportEmailStatus_TL.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006421 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L187) | JOIN_ON / P11 / COMPARISON | `TL.MessageId` : `SMALLINT` | `MMap.MapId` : `INT` | `TL.MessageId` → `INT` | `TL.MessageId = MMap.MapId` |

<a id="kolom-statusid"></a>
## StatusId

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSReportEmailStatus_TL.sql#L14). **Existing leading index:** IX_TRSReportEmailStatus_TL_StatusCallback.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006419 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:184](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L184) | JOIN_ON / P11 / COMPARISON | `TL.StatusId` : `SMALLINT` | `SMap.MapId` : `INT` | `TL.StatusId` → `INT` | `TL.StatusId = SMap.MapId` |

<a id="kolom-transactionid"></a>
## TransactionId

**Tipe dasar:** `VARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSReportEmailStatus_TL.sql#L4). **Existing leading index:** UQ_TRSReportEmailStatus_TL_CompositeKey.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006416 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateStatusEmailKonfirmasiProduct:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateStatusEmailKonfirmasiProduct.sql#L166) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `TL.TransactionId` : `VARCHAR(50)` | `#tmpDataEmail.DealId` : `bigint` | `TL.TransactionId` → `bigint` | `#tmpDataEmail.DealId <- TL.TransactionId` |
