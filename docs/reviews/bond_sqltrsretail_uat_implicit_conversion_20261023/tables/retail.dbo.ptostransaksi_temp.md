# Retail UAT.dbo.PTOSTransaksi_Temp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 38 lokasi source, 6 procedure, 7 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accruedinterest"></a>
## AccruedInterest

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Temp.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002397 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:521](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L521) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.AccruedInterest, 0)` |
| IC-002412 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:714](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L714) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.AccruedInterest, 0)` |
| IC-005767 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSTrxAuthen:36](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSTrxAuthen.sql#L36) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.AccruedInterest,0)` |
| IC-005856 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:520](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L520) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.AccruedInterest, 0)` |
| IC-005871 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:713](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L713) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.AccruedInterest` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.AccruedInterest, 0)` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Temp.sql#L25). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002401 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:535](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L535) | JOIN_ON / P01 / COMPARISON | `ptt.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptt.InsertedBy = ndme.EMPLID` |
| IC-002416 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:728](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L728) | JOIN_ON / P01 / COMPARISON | `ptt.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptt.InsertedBy = ndme.EMPLID` |
| IC-005771 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSTrxAuthen:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSTrxAuthen.sql#L53) | JOIN_ON / P01 / COMPARISON | `ptt.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptt.InsertedBy = ndme.EMPLID` |
| IC-005860 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUploadAuthen:534](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L534) | JOIN_ON / P01 / COMPARISON | `ptt.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptt.InsertedBy = ndme.EMPLID` |
| IC-005875 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUploadAuthen:727](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L727) | JOIN_ON / P01 / COMPARISON | `ptt.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptt.InsertedBy = ndme.EMPLID` |

<a id="kolom-ptostype"></a>
## PTOSType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Temp.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002418 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:730](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L730) | WHERE / P11 / COMPARISON | `ptt.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `ptt.PTOSType` → `INT` | `ptt.PTOSType = @pnType` |
| IC-005772 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSTrxAuthen:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSTrxAuthen.sql#L55) | WHERE / P11 / COMPARISON | `ptt.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `ptt.PTOSType` → `INT` | `ptt.PTOSType = @pnType` |
| IC-005877 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:729](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L729) | WHERE / P11 / COMPARISON | `ptt.PTOSType` : `SMALLINT` | `@pnType` : `INT` | `ptt.PTOSType` → `INT` | `ptt.PTOSType = @pnType` |
| IC-002359 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:173](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L173) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_temp.PTOSType` : `SMALLINT` | `tmp.PTOSType` : `INT` | `tmp.PTOSType` → `SMALLINT` | `retail.dbo.ptostransaksi_temp.PTOSType <- tmp.PTOSType` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Temp.sql#L24). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 6 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002402 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:536](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L536) | WHERE / P11 / COMPARISON | `ptt.[Status]` : `SMALLINT` | `1` : `int` | `ptt.[Status]` → `INT` | `ptt.[Status] = 1` |
| IC-002417 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:729](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L729) | WHERE / P11 / COMPARISON | `ptt.[Status]` : `SMALLINT` | `1` : `int` | `ptt.[Status]` → `INT` | `ptt.[Status] = 1` |
| IC-005773 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSTrxAuthen:56](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSTrxAuthen.sql#L56) | WHERE / P11 / IN_VALUE_TYPE | `ptt.[Status]` : `SMALLINT` | `2` : `int` | `ptt.[Status]` → `INT` | `ptt.[Status] IN (2, -1, 4)` |
| IC-005861 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:535](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L535) | WHERE / P11 / COMPARISON | `ptt.[Status]` : `SMALLINT` | `1` : `int` | `ptt.[Status]` → `INT` | `ptt.[Status] = 1` |
| IC-005876 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:728](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L728) | WHERE / P11 / COMPARISON | `ptt.[Status]` : `SMALLINT` | `1` : `int` | `ptt.[Status]` → `INT` | `ptt.[Status] = 1` |
| IC-002360 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:173](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L173) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_temp.Status` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `retail.dbo.ptostransaksi_temp.Status <- 1` |
| IC-007171 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSTrx:181](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSTrx.sql#L181) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_temp.Status` : `SMALLINT` | `@pnStatus` : `INT` | `@pnStatus` → `SMALLINT` | `retail.dbo.ptostransaksi_temp.Status <- @pnStatus` |
| IC-007205 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:1346](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L1346) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_temp.Status` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `retail.dbo.ptostransaksi_temp.Status <- 1` |
| IC-007209 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:1802](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L1802) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptostransaksi_temp.Status` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `retail.dbo.ptostransaksi_temp.Status <- 1` |

<a id="kolom-taxcapitalgainloss"></a>
## TaxCapitalGainLoss

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Temp.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002399 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:526](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L526) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TaxCapitalGainLoss, 0)` |
| IC-002414 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:719](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L719) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TaxCapitalGainLoss, 0)` |
| IC-005769 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSTrxAuthen:41](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSTrxAuthen.sql#L41) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TaxCapitalGainLoss,0)` |
| IC-005858 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:525](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L525) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TaxCapitalGainLoss, 0)` |
| IC-005873 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:718](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L718) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TaxCapitalGainLoss` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TaxCapitalGainLoss, 0)` |

<a id="kolom-taxonaccrued"></a>
## TaxOnAccrued

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Temp.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002398 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:522](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L522) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TaxOnAccrued, 0)` |
| IC-002413 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:715](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L715) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TaxOnAccrued, 0)` |
| IC-005768 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSTrxAuthen:37](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSTrxAuthen.sql#L37) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TaxOnAccrued,0)` |
| IC-005857 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:521](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L521) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TaxOnAccrued, 0)` |
| IC-005872 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:714](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L714) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TaxOnAccrued` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TaxOnAccrued, 0)` |

<a id="kolom-totalproceed"></a>
## TotalProceed

**Tipe dasar:** `MONEY`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSTransaksi_Temp.sql#L22). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002400 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:527](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L527) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TotalProceed, 0)` |
| IC-002415 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:720](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L720) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TotalProceed, 0)` |
| IC-005770 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSTrxAuthen:42](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSTrxAuthen.sql#L42) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TotalProceed,0)` |
| IC-005859 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:526](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L526) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TotalProceed, 0)` |
| IC-005874 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulatePTOSUploadAuthen:719](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L719) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `ptt.TotalProceed` : `MONEY` | `0` : `int` | `0` → `MONEY` | `ISNULL(ptt.TotalProceed, 0)` |
