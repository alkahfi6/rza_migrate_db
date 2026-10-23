# Retail UAT.dbo.SecurityTransactionDetail_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 8 lokasi source, 4 procedure, 5 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 6 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-cifiddebet"></a>
## CIFIdDebet

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionDetail_TT.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013706 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_StatusSecurityTransaction2nd_TT:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction2nd_TT.sql#L89) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `td.CIFIdDebet` : `BIGINT` | `tt.CIFIdDebet` : `int` | `td.CIFIdDebet` → `int` | `tt.CIFIdDebet = td.CIFIdDebet` |
| IC-013716 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_StatusSecurityTransaction2nd_TT:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction2nd_TT.sql#L130) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `td.CIFIdDebet` : `BIGINT` | `tt.CIFIdDebet` : `int` | `td.CIFIdDebet` → `int` | `tt.CIFIdDebet = td.CIFIdDebet` |

<a id="kolom-cifidkredit"></a>
## CIFIdKredit

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionDetail_TT.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013707 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_StatusSecurityTransaction2nd_TT:89](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction2nd_TT.sql#L89) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `td.CIFIdKredit` : `BIGINT` | `tt.CIFIdKredit` : `int` | `td.CIFIdKredit` → `int` | `tt.CIFIdKredit = td.CIFIdKredit` |
| IC-013717 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_StatusSecurityTransaction2nd_TT:130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_StatusSecurityTransaction2nd_TT.sql#L130) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `td.CIFIdKredit` : `BIGINT` | `tt.CIFIdKredit` : `int` | `td.CIFIdKredit` → `int` | `tt.CIFIdKredit = td.CIFIdKredit` |

<a id="kolom-d-k"></a>
## D/K

**Tipe dasar:** `VARCHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionDetail_TT.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013792 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_UpdateStatusSecurityTransaction2nd_TT:135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransaction2nd_TT.sql#L135) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `[D/K]` : `VARCHAR(1)` | `@cDKStatus` : `char(1)` | `[D/K]` → `char(1)` | `@cDKStatus = [D/K]` |

<a id="kolom-dealid"></a>
## DealId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionDetail_TT.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013237 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.trs_InsertTrx2ndSecurityTransaction_TT:585](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_InsertTrx2ndSecurityTransaction_TT.sql#L585) | VARIABLE_ASSIGNMENT / P08 / SELECTSETVARIABLE | `DealId` : `BIGINT` | `@nCheckDealId` : `varchar(10)` | `DealId` → `varchar(10)` | `@nCheckDealId = DealId` |

<a id="kolom-trxdetailtype"></a>
## TrxDetailType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityTransactionDetail_TT.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-013658 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_RejectStatusSecurityTransaction2nd_TT:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_RejectStatusSecurityTransaction2nd_TT.sql#L156) | WHERE / P11 / COMPARISON | `tt.TrxDetailType` : `SMALLINT` | `1` : `int` | `tt.TrxDetailType` → `INT` | `tt.TrxDetailType = 1` |
| IC-013794 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.trs_UpdateStatusSecurityTransaction2nd_TT:179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.trs_UpdateStatusSecurityTransaction2nd_TT.sql#L179) | WHERE / P11 / COMPARISON | `tt.TrxDetailType` : `SMALLINT` | `1` : `int` | `tt.TrxDetailType` → `INT` | `tt.TrxDetailType = 1` |
