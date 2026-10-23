# Retail UAT.dbo.TRSListChannelBonds_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 3 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-channelcode"></a>
## ChannelCode

**Tipe dasar:** `VARCHAR(15)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSListChannelBonds_TR.sql#L3). **Existing leading index:** TRSListChannelBonds_TR_CN1.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P12 — CHAR/VARCHAR dan keluarga teks biasa:** Perbedaan tipe ini bukan bukti scan berbahaya. Panjang, trailing spaces, collation, dan Unicode adalah dimensi berbeda; nilai karakter dan plan menentukan efeknya.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008735 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRetailCalculateFee:228](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRetailCalculateFee.sql#L228) | WHERE / P12 / COMPARISON | `ChannelCode` : `VARCHAR(15)` | `@cInstruction` : `char(15)` | `@cInstruction` → `VARCHAR` | `ChannelCode = @cInstruction` |

<a id="kolom-isupdatenorekinvestoraftertrx"></a>
## isUpdateNoRekInvestorAfterTrx

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSListChannelBonds_TR.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-008620 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSResponseBondFromMX:684](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSResponseBondFromMX.sql#L684) | JOIN_ON / P11 / COMPARISON | `tlcb.isUpdateNoRekInvestorAfterTrx` : `BIT` | `1` : `int` | `tlcb.isUpdateNoRekInvestorAfterTrx` → `INT` | `tlcb.isUpdateNoRekInvestorAfterTrx = 1` |

<a id="kolom-isuseecprice"></a>
## isUseECPrice

**Tipe dasar:** `BIT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSListChannelBonds_TR.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-004134 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSGetParamChannelPrice:12](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSGetParamChannelPrice.sql#L12) | WHERE / P11 / COMPARISON | `isUseECPrice` : `BIT` | `1` : `int` | `isUseECPrice` → `INT` | `isUseECPrice = 1` |
