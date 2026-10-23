# Retail UAT.dbo.TRSBankCustodyPending_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 42 lokasi source, 4 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-iscommitted"></a>
## IsCommitted

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBankCustodyPending_TM.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 11 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006713 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L166) | JOIN_ON / P11 / COMPARISON | `bcptm.IsCommitted` : `BIT` | `0` : `int` | `bcptm.IsCommitted` → `INT` | `bcptm.IsCommitted = 0` |
| IC-006715 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L179) | JOIN_ON / P11 / COMPARISON | `bcptm.IsCommitted` : `BIT` | `0` : `int` | `bcptm.IsCommitted` → `INT` | `bcptm.IsCommitted = 0` |
| IC-006717 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L192) | WHERE / P11 / COMPARISON | `bcptm.IsCommitted` : `BIT` | `0` : `int` | `bcptm.IsCommitted` → `INT` | `bcptm.IsCommitted = 0` |
| IC-006720 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L193) | WHERE / P11 / COMPARISON | `bcptm.IsCommitted` : `BIT` | `0` : `int` | `bcptm.IsCommitted` → `INT` | `bcptm.IsCommitted = 0` |
| IC-006744 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustodyEOD:62](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L62) | WHERE / P11 / COMPARISON | `bcptm.IsCommitted` : `BIT` | `0` : `int` | `bcptm.IsCommitted` → `INT` | `bcptm.IsCommitted = 0` |
| IC-006747 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustodyEOD:70](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L70) | WHERE / P11 / COMPARISON | `bcptm.IsCommitted` : `BIT` | `0` : `int` | `bcptm.IsCommitted` → `INT` | `bcptm.IsCommitted = 0` |
| IC-006750 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustodyEOD:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L166) | WHERE / P11 / COMPARISON | `IsCommitted` : `BIT` | `0` : `int` | `IsCommitted` → `INT` | `IsCommitted = 0` |
| IC-006753 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustodyEOD:171](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L171) | WHERE / P11 / COMPARISON | `IsCommitted` : `BIT` | `0` : `int` | `IsCommitted` → `INT` | `IsCommitted = 0` |
| IC-005288 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSPopulateBankCustody:277](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBankCustody.sql#L277) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `bcptm.IsCommitted` : `BIT` | `1` : `int` | `1` → `BIT` | `ISNULL(bcptm.IsCommitted, 1)` |
| IC-006735 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustody:512](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L512) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsCommitted` : `BIT` | `CASE WHEN temp.OperationCode IN ('U', 'D') THEN 0 ELSE 1 END` : `INT (combined result; facets unknown)` | `CASE WHEN temp.OperationCode IN ('U', 'D') THEN 0 ELSE 1 END` → `BIT` | `IsCommitted = CASE WHEN temp.OperationCode IN ('U', 'D') THEN 0 ELSE 1 END` |
| IC-006751 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustodyEOD:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L169) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `IsCommitted` : `BIT` | `1` : `int` | `1` → `BIT` | `IsCommitted = 1` |

<a id="kolom-lastprocesscode"></a>
## LastProcessCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBankCustodyPending_TM.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 22 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-005283 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBankCustody:53](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBankCustody.sql#L53) | WHERE / P12 / COMPARISON | `LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `LastProcessCode` → `VARCHAR` | `LastProcessCode = 'P'` |
| IC-005284 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBankCustody:80](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBankCustody.sql#L80) | WHERE / P12 / COMPARISON | `LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `LastProcessCode` → `VARCHAR` | `LastProcessCode = 'P'` |
| IC-005285 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBankCustody:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBankCustody.sql#L107) | WHERE / P12 / COMPARISON | `LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `LastProcessCode` → `VARCHAR` | `LastProcessCode = 'P'` |
| IC-005286 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBankCustody:134](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBankCustody.sql#L134) | WHERE / P12 / COMPARISON | `LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `LastProcessCode` → `VARCHAR` | `LastProcessCode = 'P'` |
| IC-005287 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulateBankCustody:161](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateBankCustody.sql#L161) | WHERE / P12 / COMPARISON | `LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `LastProcessCode` → `VARCHAR` | `LastProcessCode = 'P'` |
| IC-006712 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:143](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L143) | JOIN_ON / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode = 'P'` |
| IC-006714 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:166](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L166) | JOIN_ON / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode = 'P'` |
| IC-006716 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:179](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L179) | JOIN_ON / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode = 'P'` |
| IC-006718 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L192) | WHERE / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode = 'P'` |
| IC-006724 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:223](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L223) | JOIN_ON / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode = 'P'` |
| IC-006743 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustodyEOD:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L61) | WHERE / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode = 'A'` |
| IC-006746 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustodyEOD:69](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L69) | WHERE / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode = 'A'` |
| IC-006749 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustodyEOD:165](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L165) | WHERE / P12 / COMPARISON | `LastProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `LastProcessCode` → `VARCHAR` | `LastProcessCode = 'A'` |
| IC-006752 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustodyEOD:170](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L170) | WHERE / P12 / COMPARISON | `LastProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `LastProcessCode` → `VARCHAR` | `LastProcessCode = 'A'` |
| IC-012102 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchBankCustody:18](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchBankCustody.sql#L18) | JOIN_ON / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'F'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode <> 'F'` |
| IC-012104 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchBankCustody:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchBankCustody.sql#L29) | JOIN_ON / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'F'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode <> 'F'` |
| IC-012106 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchBankCustody:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchBankCustody.sql#L39) | JOIN_ON / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'F'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode <> 'F'` |
| IC-012108 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchBankCustody:49](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchBankCustody.sql#L49) | JOIN_ON / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'F'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode <> 'F'` |
| IC-012110 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchBankCustody:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchBankCustody.sql#L60) | JOIN_ON / P12 / COMPARISON | `bcptm.LastProcessCode` : `CHAR(1)` | `'F'` : `varchar` | `bcptm.LastProcessCode` → `VARCHAR` | `bcptm.LastProcessCode <> 'F'` |
| IC-006725 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustody:240](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L240) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsbankcustodypending_tm.LastProcessCode` : `CHAR(1)` | `'P'` : `varchar` | `'P'` → `CHAR(1)` | `retail.dbo.trsbankcustodypending_tm.LastProcessCode <- 'P'` |
| IC-006734 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustody:511](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L511) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastProcessCode` : `CHAR(1)` | `'A'` : `varchar` | `'A'` → `CHAR(1)` | `LastProcessCode = 'A'` |
| IC-006740 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessBankCustody:596](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L596) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `LastProcessCode` : `CHAR(1)` | `'R'` : `varchar` | `'R'` → `CHAR(1)` | `LastProcessCode = 'R'` |

<a id="kolom-operationcode"></a>
## OperationCode

**Tipe dasar:** `CHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSBankCustodyPending_TM.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 9 lokasi / 3 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006732 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:500](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L500) | WHERE / P12 / COMPARISON | `bp.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `bp.OperationCode` → `VARCHAR` | `bp.OperationCode = 'U'` |
| IC-006733 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustody:508](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustody.sql#L508) | WHERE / P12 / COMPARISON | `bp.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `bp.OperationCode` → `VARCHAR` | `bp.OperationCode = 'D'` |
| IC-006742 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustodyEOD:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L60) | WHERE / P12 / COMPARISON | `bcptm.OperationCode` : `CHAR(1)` | `'U'` : `varchar` | `bcptm.OperationCode` → `VARCHAR` | `bcptm.OperationCode = 'U'` |
| IC-006745 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSProcessBankCustodyEOD:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessBankCustodyEOD.sql#L68) | WHERE / P12 / COMPARISON | `bcptm.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `bcptm.OperationCode` → `VARCHAR` | `bcptm.OperationCode = 'D'` |
| IC-012101 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchBankCustody:17](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchBankCustody.sql#L17) | JOIN_ON / P12 / COMPARISON | `bcptm.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `bcptm.OperationCode` → `VARCHAR` | `bcptm.OperationCode = 'D'` |
| IC-012103 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchBankCustody:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchBankCustody.sql#L28) | JOIN_ON / P12 / COMPARISON | `bcptm.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `bcptm.OperationCode` → `VARCHAR` | `bcptm.OperationCode = 'D'` |
| IC-012105 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchBankCustody:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchBankCustody.sql#L38) | JOIN_ON / P12 / COMPARISON | `bcptm.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `bcptm.OperationCode` → `VARCHAR` | `bcptm.OperationCode = 'D'` |
| IC-012107 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchBankCustody:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchBankCustody.sql#L48) | JOIN_ON / P12 / COMPARISON | `bcptm.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `bcptm.OperationCode` → `VARCHAR` | `bcptm.OperationCode = 'D'` |
| IC-012109 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSSearchBankCustody:59](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchBankCustody.sql#L59) | JOIN_ON / P12 / COMPARISON | `bcptm.OperationCode` : `CHAR(1)` | `'D'` : `varchar` | `bcptm.OperationCode` → `VARCHAR` | `bcptm.OperationCode = 'D'` |
