# Retail UAT.dbo.PTOSLogHasilBMW_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 12 lokasi source, 6 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-filenameexcel"></a>
## FileNameExcel

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSLogHasilBMW_TH.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN_EXPRESSION.

**Pola dan arah review:**

- **P07 — CASE matching dan tipe hasil CASE/COALESCE/UNION:** Pada simple CASE, samakan input dengan setiap WHEN; ini terpisah dari tipe hasil THEN/ELSE. Samakan branch/output secara sengaja; cek error, rounding, duplikasi UNION, dan output contract.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012306 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSSubmitLogSentMailResult:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L110) | EXPRESSION / P07 / CASE_BRANCH | `'File ' + th.FileNameExcel + ' Berhasil di Upload'` : `VARCHAR (combined result; facets unknown)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `'File ' + th.FileNameExcel + ' Berhasil di Upload'` → `NVARCHAR` | `case when th.[Status] = 'Success' then 'File ' + th.FileNameExcel + ' Berhasil di Upload' when th.[Status] = 'No Data' then 'File ' + th.FileNameExcel + ' tidak ada data dalam excel' else @cResult end` |
| IC-012307 | B_EXPRESSION_OR_BOUNDARY / CONVERTED_COLUMN_EXPRESSION | [dbo.TRSSubmitLogSentMailResult:110](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSubmitLogSentMailResult.sql#L110) | EXPRESSION / P07 / CASE_BRANCH | `'File ' + th.FileNameExcel + ' tidak ada data dalam excel'` : `VARCHAR (combined result; facets unknown)` | `result type` : `NVARCHAR (combined result; facets unknown)` | `'File ' + th.FileNameExcel + ' tidak ada data dalam excel'` → `NVARCHAR` | `case when th.[Status] = 'Success' then 'File ' + th.FileNameExcel + ' Berhasil di Upload' when th.[Status] = 'No Data' then 'File ' + th.FileNameExcel + ' tidak ada data dalam excel' else @cResult end` |

<a id="kolom-valuedate"></a>
## ValueDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSLogHasilBMW_TH.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 10 lokasi / 5 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001188 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:156](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L156) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosloghasilbmw_th.ValueDate` : `DATETIME` | `@cValueDate` : `varchar(8)` | `@cValueDate` → `DATETIME` | `retail.dbo.ptosloghasilbmw_th.ValueDate <- @cValueDate` |
| IC-001234 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:243](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L243) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosloghasilbmw_th.ValueDate` : `DATETIME` | `@cValueDate` : `varchar(8)` | `@cValueDate` → `DATETIME` | `retail.dbo.ptosloghasilbmw_th.ValueDate <- @cValueDate` |
| IC-001275 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:75](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L75) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosloghasilbmw_th.ValueDate` : `DATETIME` | `@cValueDate` : `varchar(8)` | `@cValueDate` → `DATETIME` | `retail.dbo.ptosloghasilbmw_th.ValueDate <- @cValueDate` |
| IC-001328 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstanding:176](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L176) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosloghasilbmw_th.ValueDate` : `DATETIME` | `@cValueDate` : `varchar(8)` | `@cValueDate` → `DATETIME` | `retail.dbo.ptosloghasilbmw_th.ValueDate <- @cValueDate` |
| IC-001380 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:107](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L107) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosloghasilbmw_th.ValueDate` : `DATETIME` | `@cValueDate` : `varchar(8)` | `@cValueDate` → `DATETIME` | `retail.dbo.ptosloghasilbmw_th.ValueDate <- @cValueDate` |
| IC-001434 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:215](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L215) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosloghasilbmw_th.ValueDate` : `DATETIME` | `@cValueDate` : `varchar(8)` | `@cValueDate` → `DATETIME` | `retail.dbo.ptosloghasilbmw_th.ValueDate <- @cValueDate` |
| IC-001492 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:85](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L85) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosloghasilbmw_th.ValueDate` : `DATETIME` | `@cValueDate` : `varchar(8)` | `@cValueDate` → `DATETIME` | `retail.dbo.ptosloghasilbmw_th.ValueDate <- @cValueDate` |
| IC-001564 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksi:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L187) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosloghasilbmw_th.ValueDate` : `DATETIME` | `@cValueDate` : `varchar(8)` | `@cValueDate` → `DATETIME` | `retail.dbo.ptosloghasilbmw_th.ValueDate <- @cValueDate` |
| IC-001620 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:84](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L84) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosloghasilbmw_th.ValueDate` : `DATETIME` | `@cValueDate` : `varchar(8)` | `@cValueDate` → `DATETIME` | `retail.dbo.ptosloghasilbmw_th.ValueDate <- @cValueDate` |
| IC-001692 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:186](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L186) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptosloghasilbmw_th.ValueDate` : `DATETIME` | `@cValueDate` : `varchar(8)` | `@cValueDate` → `DATETIME` | `retail.dbo.ptosloghasilbmw_th.ValueDate <- @cValueDate` |
