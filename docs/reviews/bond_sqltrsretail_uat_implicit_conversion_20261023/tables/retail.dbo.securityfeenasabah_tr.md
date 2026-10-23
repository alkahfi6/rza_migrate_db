# Retail UAT.dbo.SecurityFeeNasabah_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 39 lokasi source, 5 procedure, 10 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 3 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-type-1-1-fee"></a>
## Type_1_1_Fee

**Tipe dasar:** `DECIMAL(10,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007271 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:388](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L388) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabah_tr.Type_1_1_Fee` : `DECIMAL(10,3)` | `Type_1_1_Fee` : `MONEY` | `Type_1_1_Fee` → `DECIMAL(10,3)` | `retail.dbo.securityfeenasabah_tr.Type_1_1_Fee <- Type_1_1_Fee` |
| IC-007278 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:564](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L564) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabah_tr.Type_1_1_Fee` : `DECIMAL(10,3)` | `Type_1_1_Fee` : `MONEY` | `Type_1_1_Fee` → `DECIMAL(10,3)` | `retail.dbo.securityfeenasabah_tr.Type_1_1_Fee <- Type_1_1_Fee` |

<a id="kolom-type-1-1-nominaltransaksifrom"></a>
## Type_1_1_NominalTransaksiFrom

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008969 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateTransactionFee:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateTransactionFee.sql#L36) | WHERE / P11 / COMPARISON | `Type_1_1_NominalTransaksiFrom` : `MONEY` | `@mFaceValue` : `decimal (20, 5)` | `Type_1_1_NominalTransaksiFrom` → `DECIMAL` | `Type_1_1_NominalTransaksiFrom <= @mFaceValue` |

<a id="kolom-type-1-1-nominaltransaksito"></a>
## Type_1_1_NominalTransaksiTo

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008970 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSRetailCalculateTransactionFee:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateTransactionFee.sql#L37) | WHERE / P11 / COMPARISON | `Type_1_1_NominalTransaksiTo` : `MONEY` | `@mFaceValue` : `decimal (20, 5)` | `Type_1_1_NominalTransaksiTo` → `DECIMAL` | `@mFaceValue <= Type_1_1_NominalTransaksiTo` |

<a id="kolom-type-1-1-ppn"></a>
## Type_1_1_PPN

**Tipe dasar:** `DECIMAL(6,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008721 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateCoupon:403](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L403) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `Type_1_1_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + Type_1_1_PPN` |
| IC-008968 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateTransactionFee:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateTransactionFee.sql#L33) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `Type_1_1_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + Type_1_1_PPN` |

<a id="kolom-type-1-2-ppn"></a>
## Type_1_2_PPN

**Tipe dasar:** `DECIMAL(6,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TR.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004084 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGenerateSafeKeepingFeeNasabah:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L237) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `(tskf.Outstanding * sfntr.Type_1_2_SafeKeepFee / 1200) * (100 + Type_1_2_PPN) / 100` : `DECIMAL (combined result; facets unknown)` | `SafeKeepingFeeAfterTax` : `MONEY` | `(tskf.Outstanding * sfntr.Type_1_2_SafeKeepFee / 1200) * (100 + Type_1_2_PPN) / 100` → `MONEY` | `SafeKeepingFeeAfterTax = (tskf.Outstanding * sfntr.Type_1_2_SafeKeepFee / 1200) * (100 + Type_1_2_PPN) / 100` |
| IC-008871 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRetailCalculateFee:1017](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1017) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `(temp.NominalJual * (Type_1_2_SafeKeepFee/100)) + ((temp.NominalJual * (Type_1_2_SafeKeepFee/100)) * (Type_1_2_PPN / 100) )` : `DECIMAL (combined result; facets unknown)` | `SafeKeepingFee` : `money` | `(temp.NominalJual * (Type_1_2_SafeKeepFee/100)) + ((temp.NominalJual * (Type_1_2_SafeKeepFee/100)) * (Type_1_2_PPN / 100) )` → `money` | `SafeKeepingFee = (temp.NominalJual * (Type_1_2_SafeKeepFee/100)) + ((temp.NominalJual * (Type_1_2_SafeKeepFee/100)) * (Type_1_2_PPN / 100) )` |
| IC-004085 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeNasabah:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L237) | DML_CONDITION / P05 / BINARY_Add | `Type_1_2_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + Type_1_2_PPN` |
| IC-004086 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeNasabah:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L237) | DML_CONDITION / P05 / BINARY_Divide | `(tskf.Outstanding * sfntr.Type_1_2_SafeKeepFee / 1200) * (100 + Type_1_2_PPN)` : `DECIMAL (combined result; facets unknown)` | `100` : `int` | `100` → `DECIMAL` | `(tskf.Outstanding * sfntr.Type_1_2_SafeKeepFee / 1200) * (100 + Type_1_2_PPN) / 100` |
| IC-008714 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateCoupon:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L166) | CASE / P05 / BINARY_Add | `Type_1_2_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + Type_1_2_PPN` |
| IC-008715 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateCoupon:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L166) | CASE / P05 / BINARY_Divide | `((Type_1_2_SafeKeepFee * xmltl.NominalJual) / 100) * (100 + Type_1_2_PPN)` : `DECIMAL (combined result; facets unknown)` | `100` : `int` | `100` → `DECIMAL` | `((Type_1_2_SafeKeepFee * xmltl.NominalJual) / 100) * (100 + Type_1_2_PPN) / 100` |
| IC-008718 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateCoupon:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L169) | DML_CONDITION / P05 / BINARY_Add | `Type_1_2_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + Type_1_2_PPN` |
| IC-008719 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateCoupon:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L169) | DML_CONDITION / P05 / BINARY_Divide | `((Type_1_2_SafeKeepFee * xmltl.NominalJual) / 100) * (100 + Type_1_2_PPN)` : `DECIMAL (combined result; facets unknown)` | `100` : `int` | `100` → `DECIMAL` | `((Type_1_2_SafeKeepFee * xmltl.NominalJual) / 100) * (100 + Type_1_2_PPN) / 100` |
| IC-008872 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1017](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1017) | DML_CONDITION / P05 / BINARY_Divide | `Type_1_2_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `Type_1_2_PPN / 100` |

<a id="kolom-type-1-2-safekeepfee"></a>
## Type_1_2_SafeKeepFee

**Tipe dasar:** `DECIMAL(16,10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TR.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004083 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeNasabah:236](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L236) | DML_CONDITION / P05 / BINARY_Multiply | `sfntr.Type_1_2_SafeKeepFee` : `DECIMAL(16,10)` | `tskf.Outstanding` : `MONEY` | `tskf.Outstanding` → `DECIMAL` | `tskf.Outstanding * sfntr.Type_1_2_SafeKeepFee` |
| IC-004087 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeNasabah:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L237) | DML_CONDITION / P05 / BINARY_Multiply | `sfntr.Type_1_2_SafeKeepFee` : `DECIMAL(16,10)` | `tskf.Outstanding` : `MONEY` | `tskf.Outstanding` → `DECIMAL` | `tskf.Outstanding * sfntr.Type_1_2_SafeKeepFee` |
| IC-008716 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateCoupon:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L166) | CASE / P05 / BINARY_Multiply | `Type_1_2_SafeKeepFee` : `DECIMAL(16,10)` | `xmltl.NominalJual` : `MONEY` | `xmltl.NominalJual` → `DECIMAL` | `Type_1_2_SafeKeepFee * xmltl.NominalJual` |
| IC-008720 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateCoupon:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L169) | DML_CONDITION / P05 / BINARY_Multiply | `Type_1_2_SafeKeepFee` : `DECIMAL(16,10)` | `xmltl.NominalJual` : `MONEY` | `xmltl.NominalJual` → `DECIMAL` | `Type_1_2_SafeKeepFee * xmltl.NominalJual` |
| IC-008870 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1010](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1010) | WHERE / P11 / COMPARISON | `Type_1_2_SafeKeepFee` : `DECIMAL(16,10)` | `0` : `int` | `0` → `DECIMAL` | `Type_1_2_SafeKeepFee = 0` |
| IC-008873 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1017](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1017) | DML_CONDITION / P05 / BINARY_Divide | `Type_1_2_SafeKeepFee` : `DECIMAL(16,10)` | `100` : `int` | `100` → `DECIMAL` | `Type_1_2_SafeKeepFee/100` |
| IC-008874 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1017](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1017) | DML_CONDITION / P05 / BINARY_Divide | `Type_1_2_SafeKeepFee` : `DECIMAL(16,10)` | `100` : `int` | `100` → `DECIMAL` | `Type_1_2_SafeKeepFee/100` |
| IC-008875 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1017](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1017) | DML_CONDITION / P05 / BINARY_Multiply | `(Type_1_2_SafeKeepFee/100)` : `DECIMAL (combined result; facets unknown)` | `temp.NominalJual` : `money` | `temp.NominalJual` → `DECIMAL` | `temp.NominalJual * (Type_1_2_SafeKeepFee/100)` |
| IC-008876 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1017](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1017) | DML_CONDITION / P05 / BINARY_Multiply | `(Type_1_2_SafeKeepFee/100)` : `DECIMAL (combined result; facets unknown)` | `temp.NominalJual` : `money` | `temp.NominalJual` → `DECIMAL` | `temp.NominalJual * (Type_1_2_SafeKeepFee/100)` |

<a id="kolom-type-2-1-fee"></a>
## Type_2_1_Fee

**Tipe dasar:** `DECIMAL(10,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TR.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007272 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:388](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L388) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabah_tr.Type_2_1_Fee` : `DECIMAL(10,3)` | `Type_2_1_Fee` : `MONEY` | `Type_2_1_Fee` → `DECIMAL(10,3)` | `retail.dbo.securityfeenasabah_tr.Type_2_1_Fee <- Type_2_1_Fee` |
| IC-007279 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessSecurityFee:564](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessSecurityFee.sql#L564) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.securityfeenasabah_tr.Type_2_1_Fee` : `DECIMAL(10,3)` | `Type_2_1_Fee` : `MONEY` | `Type_2_1_Fee` → `DECIMAL(10,3)` | `retail.dbo.securityfeenasabah_tr.Type_2_1_Fee <- Type_2_1_Fee` |

<a id="kolom-type-2-1-ppn"></a>
## Type_2_1_PPN

**Tipe dasar:** `DECIMAL(6,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TR.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008708 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateCoupon:140](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L140) | DML_CONDITION / P05 / BINARY_Add | `Type_2_1_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + Type_2_1_PPN` |
| IC-008766 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:394](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L394) | DML_CONDITION / P05 / BINARY_Add | `Type_2_1_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + Type_2_1_PPN` |

<a id="kolom-type-2-2-ppn"></a>
## Type_2_2_PPN

**Tipe dasar:** `DECIMAL(6,3)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TR.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004092 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSGenerateSafeKeepingFeeNasabah:267](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L267) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `sfntr.Type_2_2_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + sfntr.Type_2_2_PPN` |
| IC-008710 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateCoupon:151](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L151) | CASE / P05 / BINARY_Add | `Type_2_2_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + Type_2_2_PPN` |
| IC-008712 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateCoupon:154](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateCoupon.sql#L154) | DML_CONDITION / P05 / BINARY_Add | `Type_2_2_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `100 + Type_2_2_PPN` |
| IC-008889 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1133) | CASE / P05 / BINARY_Divide | `Type_2_2_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `Type_2_2_PPN/100` |
| IC-008892 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1136) | CASE / P05 / BINARY_Divide | `Type_2_2_PPN` : `DECIMAL(6,3)` | `100` : `int` | `100` → `DECIMAL` | `Type_2_2_PPN/100` |

<a id="kolom-type-2-2-safekeepfee"></a>
## Type_2_2_SafeKeepFee

**Tipe dasar:** `DECIMAL(24,10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SecurityFeeNasabah_TR.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004088 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSGenerateSafeKeepingFeeNasabah:245](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGenerateSafeKeepingFeeNasabah.sql#L245) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `sfntr.Type_2_2_SafeKeepFee` : `DECIMAL(24,10)` | `#tSafeKeepingFeePerTransaction.SafeKeepingFee` : `MONEY` | `sfntr.Type_2_2_SafeKeepFee` → `MONEY` | `#tSafeKeepingFeePerTransaction.SafeKeepingFee <- sfntr.Type_2_2_SafeKeepFee` |
| IC-008887 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1130](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1130) | CASE / P11 / COMPARISON | `sfn.Type_2_2_SafeKeepFee` : `DECIMAL(24,10)` | `0` : `int` | `0` → `DECIMAL` | `sfn.Type_2_2_SafeKeepFee = 0` |
| IC-008888 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1132](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1132) | CASE / P05 / BINARY_Multiply | `sfn.Type_2_2_SafeKeepFee` : `DECIMAL(24,10)` | `temp.NominalTersimpan` : `money` | `temp.NominalTersimpan` → `DECIMAL` | `temp.NominalTersimpan * sfn.Type_2_2_SafeKeepFee` |
| IC-008890 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1133) | CASE / P05 / BINARY_Multiply | `sfn.Type_2_2_SafeKeepFee` : `DECIMAL(24,10)` | `temp.NominalTersimpan` : `money` | `temp.NominalTersimpan` → `DECIMAL` | `temp.NominalTersimpan * sfn.Type_2_2_SafeKeepFee` |
| IC-008891 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1135](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1135) | CASE / P05 / BINARY_Multiply | `sfn.Type_2_2_SafeKeepFee` : `DECIMAL(24,10)` | `temp.NominalTersimpan` : `money` | `temp.NominalTersimpan` → `DECIMAL` | `temp.NominalTersimpan * sfn.Type_2_2_SafeKeepFee` |
| IC-008893 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:1136](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L1136) | CASE / P05 / BINARY_Multiply | `sfn.Type_2_2_SafeKeepFee` : `DECIMAL(24,10)` | `temp.NominalTersimpan` : `money` | `temp.NominalTersimpan` → `DECIMAL` | `temp.NominalTersimpan * sfn.Type_2_2_SafeKeepFee` |
