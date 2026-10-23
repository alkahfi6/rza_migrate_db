# Retail UAT.dbo.TRSCPT_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 4 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-accrued"></a>
## Accrued

**Tipe dasar:** `DECIMAL(19,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCPT_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-006495 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulateTransaksiBuyBack:148](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateTransaksiBuyBack.sql#L148) | EXPRESSION_OR_CONTROL / P05 / BINARY_Add | `Accrued` : `DECIMAL(19,5)` | `(sum(NominalBuyBack) * (HargaBuyBack/100))` : `MONEY (combined result; facets unknown)` | `(sum(NominalBuyBack) * (HargaBuyBack/100))` → `DECIMAL` | `(sum(NominalBuyBack) * (HargaBuyBack/100)) + Accrued` |

<a id="kolom-ctpid"></a>
## CTPId

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCPT_TM.sql#L2). **Existing leading index:** ix_TRSCPT_TM_2.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-003654 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCreateMT541Content:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCreateMT541Content.sql#L66) | WHERE / P11 / COMPARISON | `CTPId` : `BIGINT` | `@cCTPId` : `int` | `@cCTPId` → `BIGINT` | `@cCTPId=CTPId` |
| IC-011662 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSSaveDataYieldBIS4:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveDataYieldBIS4.sql#L51) | JOIN_ON / P11 / COMPARISON | `cpt.CTPId` : `BIGINT` | `tmp.CTPId` : `INT` | `tmp.CTPId` → `BIGINT` | `cpt.CTPId = tmp.CTPId` |

<a id="kolom-yield"></a>
## Yield

**Tipe dasar:** `DECIMAL(19,5)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSCPT_TM.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012086 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSaveUploadCTP:208](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSaveUploadCTP.sql#L208) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trscpt_tm.Yield` : `DECIMAL(19,5)` | `isnull(Yield,0)` : `float` | `isnull(Yield,0)` → `DECIMAL(19,5)` | `retail.dbo.trscpt_tm.Yield <- isnull(Yield,0)` |
