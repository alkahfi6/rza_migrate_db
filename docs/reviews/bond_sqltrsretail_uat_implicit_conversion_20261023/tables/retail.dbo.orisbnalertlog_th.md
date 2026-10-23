# Retail UAT.dbo.ORISBNAlertLog_TH — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 5 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-triggerbit"></a>
## TriggerBit

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.ORISBNAlertLog_TH.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 5 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002784 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedCheckAlert:38](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedCheckAlert.sql#L38) | WHERE / P11 / COMPARISON | `TriggerBit` : `BIT` | `0` : `int` | `TriggerBit` → `INT` | `TriggerBit = 0` |
| IC-002786 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.SBNSchedCheckAlert:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedCheckAlert.sql#L48) | WHERE / P11 / COMPARISON | `TriggerBit` : `BIT` | `0` : `int` | `TriggerBit` → `INT` | `TriggerBit = 0` |
| IC-002779 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedCheckAlert:9](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedCheckAlert.sql#L9) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbnalertlog_th.TriggerBit` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.orisbnalertlog_th.TriggerBit <- 0` |
| IC-002781 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedCheckAlert:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedCheckAlert.sql#L23) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.orisbnalertlog_th.TriggerBit` : `BIT` | `0` : `int` | `0` → `BIT` | `retail.dbo.orisbnalertlog_th.TriggerBit <- 0` |
| IC-002785 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SBNSchedCheckAlert:47](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SBNSchedCheckAlert.sql#L47) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `TriggerBit` : `BIT` | `1` : `int` | `1` → `BIT` | `TriggerBit = 1` |
