# Retail UAT.dbo.TRSRekSyariah_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 11 lokasi source, 3 procedure, 2 kolom. **A:** 4 lokasi pada 1 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-norekinvestor"></a>
## NoRekInvestor

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSRekSyariah_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 6 lokasi / 2 SP. **A:** 4 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000611 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L143) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `b.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = b.ACCTNO` |
| IC-000612 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:145](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L145) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-000631 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:259](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L259) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `b.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = b.ACCTNO` |
| IC-000632 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetRekeningMasterNasabahTRS:261](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetRekeningMasterNasabahTRS.sql#L261) | JOIN_ON / P01 / COMPARISON | `a.NoRekInvestor` : `VARCHAR(20)` | `c.ACCTNO` : `DECIMAL(19,0)` | `a.NoRekInvestor` → `DECIMAL` | `a.NoRekInvestor = c.ACCTNO` |
| IC-008643 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:779](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L779) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `trs.NoRekInvestor` : `VARCHAR(20)` | `fm.NoRekInvestor` : `BIGINT` | `fm.NoRekInvestor` → `VARCHAR(20)` | `trs.NoRekInvestor = fm.NoRekInvestor` |
| IC-008646 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSResponseBondFromMX:786](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L786) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsreksyariah_tr.NoRekInvestor` : `VARCHAR(20)` | `fm.NoRekInvestor` : `BIGINT` | `fm.NoRekInvestor` → `VARCHAR(20)` | `retail.dbo.trsreksyariah_tr.NoRekInvestor <- fm.NoRekInvestor` |

<a id="kolom-cifid"></a>
## CIFId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSRekSyariah_TR.sql#L3). **Existing leading index:** TRSRekSyariah_TR_NN1.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002875 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:424](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L424) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |
| IC-002877 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:445](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L445) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |
| IC-002884 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:632](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L632) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |
| IC-002891 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:888](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L888) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |
| IC-002893 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.SBNSchedProcessRekDana:909](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedProcessRekDana.sql#L909) | WHERE / P11 / COMPARISON | `CIFId` : `BIGINT` | `@nCIFId` : `int` | `@nCIFId` → `BIGINT` | `CIFId = @nCIFId` |
