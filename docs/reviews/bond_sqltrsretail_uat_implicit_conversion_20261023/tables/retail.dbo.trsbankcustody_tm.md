# Retail UAT.dbo.TRSBankCustody_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 7 lokasi source, 4 procedure, 2 kolom. **A:** 3 lokasi pada 3 procedure. **B:** 4 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-idsubreg"></a>
## IdSubReg

**Tipe dasar:** `VARCHAR(10)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBankCustody_TM.sql#L32). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 3 SP. **A:** 3 lokasi / 3 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000960 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.ORIOnlineReportDailyTransaction:366](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql#L366) | JOIN_ON / P01 / COMPARISON | `cusd.IdSubReg` : `VARCHAR(10)` | `sb.IdSubregistry` : `INT` | `cusd.IdSubReg` → `INT` | `sb.IdSubregistry = cusd.IdSubReg` |
| IC-008372 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSReportReportSettlementInstruction:621](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportReportSettlementInstruction.sql#L621) | JOIN_ON / P01 / COMPARISON | `tbctm.IdSubReg` : `VARCHAR(10)` | `sb.IdSubregistry` : `INT` | `tbctm.IdSubReg` → `INT` | `sb.IdSubregistry = tbctm.IdSubReg` |
| IC-010344 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSRptOrderORIDetail:838](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptOrderORIDetail.sql#L838) | JOIN_ON / P01 / COMPARISON | `cusd.IdSubReg` : `VARCHAR(10)` | `sb.IdSubregistry` : `INT` | `cusd.IdSubReg` → `INT` | `sb.IdSubregistry = cusd.IdSubReg` |

<a id="kolom-participant"></a>
## Participant

**Tipe dasar:** `VARCHAR(500)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBankCustody_TM.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006054 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementBIS4:322](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L322) | JOIN_ON / P04 / COMPARISON | `LTRIM(RTRIM(LEFT(mp.Participant, case when CHARINDEX('-', mp.Participant)= 0 then 0 else ( CHARINDEX('-', mp.Participant) - 1 ) END )))` : `VARCHAR (expression; length unknown)` | `bpcs.ParticipantCode` : `NVARCHAR (combined result; facets unknown)` | `LTRIM(RTRIM(LEFT(mp.Participant, case when CHARINDEX('-', mp.Participant)= 0 then 0 else ( CHARINDEX('-', mp.Participant) - 1 ) END )))` → `NVARCHAR` | `bpcs.ParticipantCode = LTRIM(RTRIM(LEFT(mp.Participant, case when CHARINDEX('-', mp.Participant)= 0 then 0 else ( CHARINDEX('-', mp.Participant) - 1 ) END )))` |
| IC-006083 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementBIS4:488](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L488) | JOIN_ON / P04 / COMPARISON | `LTRIM(RTRIM(LEFT(mp.Participant, CASE WHEN CHARINDEX('-', mp.Participant)= 0 THEN 0 ELSE ( CHARINDEX('-', mp.Participant) - 1 ) END )))` : `VARCHAR (expression; length unknown)` | `bpcs.ParticipantCode` : `NVARCHAR (combined result; facets unknown)` | `LTRIM(RTRIM(LEFT(mp.Participant, CASE WHEN CHARINDEX('-', mp.Participant)= 0 THEN 0 ELSE ( CHARINDEX('-', mp.Participant) - 1 ) END )))` → `NVARCHAR` | `bpcs.ParticipantCode = LTRIM(RTRIM(LEFT(mp.Participant, CASE WHEN CHARINDEX('-', mp.Participant)= 0 THEN 0 ELSE ( CHARINDEX('-', mp.Participant) - 1 ) END )))` |
| IC-006127 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementBIS4:1013](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1013) | JOIN_ON / P04 / COMPARISON | `LTRIM(RTRIM(LEFT(mp.Participant, case when CHARINDEX('-', mp.Participant)= 0 then 0 else ( CHARINDEX('-', mp.Participant) - 1 ) END )))` : `VARCHAR (expression; length unknown)` | `bpcs.ParticipantCode` : `NVARCHAR (combined result; facets unknown)` | `LTRIM(RTRIM(LEFT(mp.Participant, case when CHARINDEX('-', mp.Participant)= 0 then 0 else ( CHARINDEX('-', mp.Participant) - 1 ) END )))` → `NVARCHAR` | `bpcs.ParticipantCode = LTRIM(RTRIM(LEFT(mp.Participant, case when CHARINDEX('-', mp.Participant)= 0 then 0 else ( CHARINDEX('-', mp.Participant) - 1 ) END )))` |
| IC-006152 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSPopulateSettlementBIS4:1133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4.sql#L1133) | JOIN_ON / P04 / COMPARISON | `LTRIM(RTRIM(LEFT(mp.Participant, CASE WHEN CHARINDEX('-', mp.Participant)= 0 THEN 0 ELSE ( CHARINDEX('-', mp.Participant) - 1 ) END )))` : `VARCHAR (expression; length unknown)` | `bpcs.ParticipantCode` : `NVARCHAR (combined result; facets unknown)` | `LTRIM(RTRIM(LEFT(mp.Participant, CASE WHEN CHARINDEX('-', mp.Participant)= 0 THEN 0 ELSE ( CHARINDEX('-', mp.Participant) - 1 ) END )))` → `NVARCHAR` | `bpcs.ParticipantCode = LTRIM(RTRIM(LEFT(mp.Participant, CASE WHEN CHARINDEX('-', mp.Participant)= 0 THEN 0 ELSE ( CHARINDEX('-', mp.Participant) - 1 ) END )))` |
