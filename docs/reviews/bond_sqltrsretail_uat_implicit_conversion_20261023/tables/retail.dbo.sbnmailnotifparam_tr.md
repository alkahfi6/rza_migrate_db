# Retail UAT.dbo.SBNMailNotifParam_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 12 lokasi source, 2 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-flagsyariah"></a>
## FlagSyariah

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.SBNMailNotifParam_TR.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002913 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L192) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `1` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 1` |
| IC-002914 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:198](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L198) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `0` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 0` |
| IC-002915 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:209](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L209) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `1` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 1` |
| IC-002916 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L216) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `0` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 0` |
| IC-002918 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:347](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L347) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `1` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 1` |
| IC-002919 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:353](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L353) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `0` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 0` |
| IC-002920 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:509](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L509) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `1` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 1` |
| IC-002921 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderCompleted:516](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderCompleted.sql#L516) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `0` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 0` |
| IC-002937 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:213](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L213) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `1` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 1` |
| IC-002938 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:225](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L225) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `0` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 0` |
| IC-002939 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:283](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L283) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `1` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 1` |
| IC-002940 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedPushMailNotifOrderORI:295](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedPushMailNotifOrderORI.sql#L295) | WHERE / P11 / COMPARISON | `[FlagSyariah]` : `BIT` | `0` : `int` | `[FlagSyariah]` → `INT` | `[FlagSyariah] = 0` |
