# Retail UAT.dbo.TRSCustodianUploadID_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 2 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 1 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-iddata"></a>
## IdData

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustodianUploadID_TR.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005270 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSPopulateApprovalSRESID:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalSRESID.sql#L51) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `IdData` : `INT` | `#tmpDataParameter.IdData` : `BIGINT` | `IdData` → `BIGINT` | `#tmpDataParameter.IdData <- IdData` |
| IC-006699 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovedUploadSRESID:64](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSRESID.sql#L64) | WHERE / P11 / COMPARISON | `IdData` : `INT` | `@pnIdData` : `BIGINT` | `IdData` → `BIGINT` | `IdData = @pnIdData` |
| IC-006700 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovedUploadSRESID:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSRESID.sql#L69) | WHERE / P11 / COMPARISON | `IdData` : `INT` | `@pnIdData` : `BIGINT` | `IdData` → `BIGINT` | `IdData = @pnIdData` |
| IC-006701 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovedUploadSRESID:73](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSRESID.sql#L73) | WHERE / P11 / COMPARISON | `IdData` : `INT` | `@pnIdData` : `BIGINT` | `IdData` → `BIGINT` | `IdData != @pnIdData` |
| IC-006702 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovedUploadSRESID:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSRESID.sql#L92) | WHERE / P11 / COMPARISON | `IdData` : `INT` | `@pnIdData` : `BIGINT` | `IdData` → `BIGINT` | `IdData = @pnIdData` |
| IC-006710 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessApprovedUploadSRESID:346](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSRESID.sql#L346) | WHERE / P11 / COMPARISON | `IdData` : `INT` | `@pnIdData` : `BIGINT` | `IdData` → `BIGINT` | `IdData = @pnIdData` |

<a id="kolom-bsid"></a>
## bSID

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustodianUploadID_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005271 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalSRESID:78](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalSRESID.sql#L78) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `bSID` : `BIT` | `1` : `int` | `bSID` → `INT` | `WHEN 1 THEN [SID]` |
| IC-005272 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalSRESID:82](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalSRESID.sql#L82) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `bSID` : `BIT` | `1` : `int` | `bSID` → `INT` | `WHEN 1 THEN ''` |

<a id="kolom-bsre"></a>
## bSRE

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCustodianUploadID_TR.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005273 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalSRESID:86](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalSRESID.sql#L86) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `bSRE` : `BIT` | `1` : `int` | `bSRE` → `INT` | `WHEN 1 THEN SRE` |
| IC-005274 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateApprovalSRESID:90](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateApprovalSRESID.sql#L90) | CASE_TEST / P07 / SIMPLE_CASE_MATCH | `bSRE` : `BIT` | `1` : `int` | `bSRE` → `INT` | `WHEN 1 THEN ''` |
