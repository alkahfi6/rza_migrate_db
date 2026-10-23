# Retail UAT.dbo.TRSParamMaterai_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 2 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-endlimitnominaltransaksi"></a>
## EndLimitNominalTransaksi

**Tipe dasar:** `DECIMAL(25,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamMaterai_TR.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 3 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003537 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateMateraiFee:21](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L21) | CASE / P11 / COMPARISON | `EndLimitNominalTransaksi` : `DECIMAL(25,2)` | `@pnTotalProceed` : `MONEY` | `@pnTotalProceed` → `DECIMAL` | `@pnTotalProceed <= EndLimitNominalTransaksi` |
| IC-003541 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateMateraiFee:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L28) | CASE / P05 / BINARY_Divide | `EndLimitNominalTransaksi` : `DECIMAL(25,2)` | `@nRate_TAX` : `MONEY` | `@nRate_TAX` → `DECIMAL` | `EndLimitNominalTransaksi / @nRate_TAX` |
| IC-003542 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateMateraiFee:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L28) | CASE / P11 / COMPARISON | `EndLimitNominalTransaksi / @nRate_TAX` : `DECIMAL (combined result; facets unknown)` | `@pnTotalProceed` : `MONEY` | `@pnTotalProceed` → `DECIMAL` | `@pnTotalProceed <= EndLimitNominalTransaksi / @nRate_TAX` |

<a id="kolom-startlimitnominaltransaksi"></a>
## StartLimitNominalTransaksi

**Tipe dasar:** `DECIMAL(25,2)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSParamMaterai_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003536 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateMateraiFee:20](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L20) | CASE / P11 / COMPARISON | `StartLimitNominalTransaksi` : `DECIMAL(25,2)` | `@pnTotalProceed` : `MONEY` | `@pnTotalProceed` → `DECIMAL` | `@pnTotalProceed > StartLimitNominalTransaksi` |
| IC-003538 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateMateraiFee:22](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L22) | CASE / P11 / COMPARISON | `StartLimitNominalTransaksi` : `DECIMAL(25,2)` | `@pnTotalProceed` : `MONEY` | `@pnTotalProceed` → `DECIMAL` | `@pnTotalProceed > StartLimitNominalTransaksi` |
| IC-003539 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateMateraiFee:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L27) | CASE / P05 / BINARY_Divide | `StartLimitNominalTransaksi` : `DECIMAL(25,2)` | `@nRate_TAX` : `MONEY` | `@nRate_TAX` → `DECIMAL` | `StartLimitNominalTransaksi / @nRate_TAX` |
| IC-003540 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateMateraiFee:27](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L27) | CASE / P11 / COMPARISON | `StartLimitNominalTransaksi / @nRate_TAX` : `DECIMAL (combined result; facets unknown)` | `@pnTotalProceed` : `MONEY` | `@pnTotalProceed` → `DECIMAL` | `@pnTotalProceed > StartLimitNominalTransaksi / @nRate_TAX` |
| IC-003543 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateMateraiFee:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L29) | CASE / P05 / BINARY_Divide | `StartLimitNominalTransaksi` : `DECIMAL(25,2)` | `@nRate_TAX` : `MONEY` | `@nRate_TAX` → `DECIMAL` | `StartLimitNominalTransaksi / @nRate_TAX` |
| IC-003544 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCalculateMateraiFee:29](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCalculateMateraiFee.sql#L29) | CASE / P11 / COMPARISON | `StartLimitNominalTransaksi / @nRate_TAX` : `DECIMAL (combined result; facets unknown)` | `@pnTotalProceed` : `MONEY` | `@pnTotalProceed` → `DECIMAL` | `@pnTotalProceed > StartLimitNominalTransaksi / @nRate_TAX` |
| IC-006531 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTrxBuyBackPerdana:240](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTrxBuyBackPerdana.sql#L240) | CASE / P11 / COMPARISON | `(select top 1 StartLimitNominalTransaksi from TRSParamMaterai_TR where ProductCode = 'BOND')` : `DECIMAL(25,2)` | `TotalProceed` : `MONEY` | `TotalProceed` → `DECIMAL` | `TotalProceed >= (select top 1 StartLimitNominalTransaksi from TRSParamMaterai_TR where ProductCode = 'BOND')` |
