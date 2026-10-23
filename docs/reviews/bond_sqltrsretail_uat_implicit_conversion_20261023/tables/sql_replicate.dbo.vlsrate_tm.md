# SQL_REPLICATE.dbo.VLSRate_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 10 lokasi source, 2 procedure, 3 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 2 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-birateexchange"></a>
## BiRateExchange

**Tipe dasar:** `FLOAT`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.VLSRate_TM.sql#L5). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** SOURCE_TO_DESTINATION, TYPE_ANCHOR.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-011602 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptTransactionDetailWMTA:591](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L591) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ISNULL(sc.BiRateExchange, 1)` : `FLOAT` | `KursAcquisition` : `DECIMAL(13,7)` | `ISNULL(sc.BiRateExchange, 1)` → `DECIMAL(13,7)` | `KursAcquisition = ISNULL(sc.BiRateExchange, 1)` |
| IC-011606 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSRptTransactionDetailWMTA:613](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L613) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `ISNULL(sc.BiRateExchange, 1)` : `FLOAT` | `tmp.Kurs` : `decimal(13,7)` | `ISNULL(sc.BiRateExchange, 1)` → `decimal(13,7)` | `tmp.Kurs = ISNULL(sc.BiRateExchange, 1)` |
| IC-011603 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptTransactionDetailWMTA:591](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L591) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sc.BiRateExchange` : `FLOAT` | `1` : `int` | `1` → `FLOAT` | `ISNULL(sc.BiRateExchange, 1)` |
| IC-011607 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.TRSRptTransactionDetailWMTA:613](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L613) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `sc.BiRateExchange` : `FLOAT` | `1` : `int` | `1` → `FLOAT` | `ISNULL(sc.BiRateExchange, 1)` |

<a id="kolom-jisdorrate"></a>
## JISDORRate

**Tipe dasar:** `FLOAT`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.VLSRate_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P05 — Aritmetika/concatenation dengan tipe campuran:** Pisahkan kebutuhan hitung dan penyusunan string. Numeric mempunyai precedence lebih tinggi daripada teks; operator + dapat menjadi penjumlahan dan bukan penggabungan string.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007800 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:91](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L91) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `jisdor.JISDORRate` : `FLOAT` | `st.TotalProceed` : `MONEY` | `st.TotalProceed` → `FLOAT` | `st.TotalProceed * jisdor.JISDORRate` |
| IC-007819 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:178](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L178) | EXPRESSION_OR_CONTROL / P05 / BINARY_Multiply | `jisdor.JISDORRate` : `FLOAT` | `st.TotalProceed` : `MONEY` | `st.TotalProceed` → `FLOAT` | `st.TotalProceed * jisdor.JISDORRate` |

<a id="kolom-period"></a>
## Period

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.VLSRate_TM.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 4 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P03 — Tanggal teks bertemu date/time:** Tentukan kontrak YYYYMMDD/hari/timestamp. Bentuk nilai pencarian dalam tipe kolom atau gunakan rentang native date; jangan konversi kolom seluruh baris tanpa alasan.
- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007805 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L129) | JOIN_ON / P04 / COMPARISON | `jisdor.Period` : `DATETIME` | `convert(varchar,st.TrxDate,112)` : `varchar` | `convert(varchar,st.TrxDate,112)` → `DATETIME` | `convert(varchar,st.TrxDate,112) = jisdor.Period` |
| IC-007825 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSREPOReportBiayaMaterai:217](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSREPOReportBiayaMaterai.sql#L217) | JOIN_ON / P04 / COMPARISON | `jisdor.Period` : `DATETIME` | `convert(varchar,st.TrxDate,112)` : `varchar` | `convert(varchar,st.TrxDate,112)` → `DATETIME` | `convert(varchar,st.TrxDate,112) = jisdor.Period` |
| IC-011604 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptTransactionDetailWMTA:598](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L598) | JOIN_ON / P04 / COMPARISON | `sc.Period` : `DATETIME` | `CONVERT(VARCHAR, stt.TrxDate, 112)` : `VARCHAR` | `CONVERT(VARCHAR, stt.TrxDate, 112)` → `DATETIME` | `CONVERT(VARCHAR, stt.TrxDate, 112) = sc.Period` |
| IC-011608 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSRptTransactionDetailWMTA:616](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptTransactionDetailWMTA.sql#L616) | JOIN_ON / P03 / COMPARISON | `sc.Period` : `DATETIME` | `tmp.TanggalTransaksi` : `VARCHAR(8)` | `tmp.TanggalTransaksi` → `DATETIME` | `tmp.TanggalTransaksi = sc.Period` |
