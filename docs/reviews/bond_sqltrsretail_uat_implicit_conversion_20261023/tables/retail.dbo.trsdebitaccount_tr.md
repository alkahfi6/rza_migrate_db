# Retail UAT.dbo.TRSDebitAccount_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 1 procedure, 3 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accountdebit"></a>
## AccountDebit

**Tipe dasar:** `VARCHAR(12)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSDebitAccount_TR.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003583 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSCheckDebitAccount:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckDebitAccount.sql#L18) | WHERE / P01 / COMPARISON | `AccountDebit` : `VARCHAR(12)` | `@pnAccountDebit` : `int` | `AccountDebit` → `INT` | `AccountDebit = @pnAccountDebit` |
| IC-003584 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCheckDebitAccount:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckDebitAccount.sql#L28) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsdebitaccount_tr.AccountDebit` : `VARCHAR(12)` | `@pnAccountDebit` : `int` | `@pnAccountDebit` → `VARCHAR(12)` | `retail.dbo.trsdebitaccount_tr.AccountDebit <- @pnAccountDebit` |

<a id="kolom-dull"></a>
## Dull

**Tipe dasar:** `TINYINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSDebitAccount_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003581 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSCheckDebitAccount:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckDebitAccount.sql#L15) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Dull + 1` : `INT (combined result; facets unknown)` | `Dull` : `TINYINT` | `Dull + 1` → `TINYINT` | `Dull = Dull + 1` |
| IC-003582 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSCheckDebitAccount:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckDebitAccount.sql#L15) | DML_CONDITION / P05 / BINARY_Add | `Dull` : `TINYINT` | `1` : `int` | `Dull` → `INT` | `Dull + 1` |
| IC-003581 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCheckDebitAccount:15](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckDebitAccount.sql#L15) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `Dull` : `TINYINT` | `Dull + 1` : `INT (combined result; facets unknown)` | `Dull + 1` → `TINYINT` | `Dull = Dull + 1` |
| IC-003585 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCheckDebitAccount:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckDebitAccount.sql#L28) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsdebitaccount_tr.Dull` : `TINYINT` | `1` : `int` | `1` → `TINYINT` | `retail.dbo.trsdebitaccount_tr.Dull <- 1` |

<a id="kolom-statustrx"></a>
## StatusTrx

**Tipe dasar:** `TINYINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSDebitAccount_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003586 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCheckDebitAccount:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckDebitAccount.sql#L28) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsdebitaccount_tr.StatusTrx` : `TINYINT` | `1` : `int` | `1` → `TINYINT` | `retail.dbo.trsdebitaccount_tr.StatusTrx <- 1` |
| IC-003587 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSCheckDebitAccount:40](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckDebitAccount.sql#L40) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `StatusTrx` : `TINYINT` | `2` : `int` | `2` → `TINYINT` | `StatusTrx=2` |
