# Retail UAT.dbo.oriva_security_master — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-sec-id"></a>
## sec_id

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.oriva_security_master.sql#L2). **Existing leading index:** ix_settlement_oriva_log_sec_id.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001038 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OrivaTransactionReport:76](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L76) | JOIN_ON / P11 / COMPARISON | `tr.sec_id` : `BIGINT` | `sec.SecId` : `INT` | `sec.SecId` → `BIGINT` | `sec.SecId = tr.sec_id` |
| IC-001781 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RPAJurnalSettlementOriva:83](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql#L83) | JOIN_ON / P11 / COMPARISON | `tr.sec_id` : `BIGINT` | `sec.SecId` : `INT` | `sec.SecId` → `BIGINT` | `sec.SecId = tr.sec_id` |

<a id="kolom-use-customer-whitelist"></a>
## use_customer_whitelist

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.oriva_security_master.sql#L10). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001027 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNISubmitInvestorPending:195](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitInvestorPending.sql#L195) | WHERE / P11 / COMPARISON | `use_customer_whitelist` : `BIT` | `0` : `int` | `use_customer_whitelist` → `INT` | `use_customer_whitelist = 0` |
| IC-001028 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.ORISBNISubmitInvestorPending:196](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORISBNISubmitInvestorPending.sql#L196) | WHERE / P11 / COMPARISON | `use_customer_whitelist` : `BIT` | `1` : `int` | `use_customer_whitelist` → `INT` | `use_customer_whitelist = 1` |
