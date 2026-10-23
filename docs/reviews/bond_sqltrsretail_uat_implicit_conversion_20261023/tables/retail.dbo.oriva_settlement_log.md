# Retail UAT.dbo.oriva_settlement_log — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 17 lokasi source, 3 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-order-id"></a>
## order_id

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.oriva_settlement_log.sql#L4). **Existing leading index:** ix_oriva_settlement_log_order_id.

**Frekuensi source:** 16 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED, TYPE_ANCHOR.

**Pola dan arah review:**

- **P09 — ISNULL dan NULLIF mempunyai aturan berbeda:** ISNULL mengonversi replacement menuju tipe argumen pertama (pengecualian literal NULL). NULLIF membandingkan kedua argumen menurut aturan comparison, tetapi mengembalikan tipe argumen pertama. Validasi sentinel dan kesetaraan NULL behavior.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001032 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OrivaTransactionReport:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L65) | CASE / P11 / COMPARISON | `ISNULL(lg.order_id, 0)` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0) = 0` |
| IC-001033 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OrivaTransactionReport:65](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L65) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.order_id` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0)` |
| IC-001034 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OrivaTransactionReport:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L66) | CASE / P11 / COMPARISON | `ISNULL(lg.order_id, 0)` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0) != 0` |
| IC-001035 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OrivaTransactionReport:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L66) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.order_id` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0)` |
| IC-001036 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OrivaTransactionReport:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L67) | CASE / P11 / COMPARISON | `ISNULL(lg.order_id, 0)` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0) != 0` |
| IC-001037 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OrivaTransactionReport:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L67) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.order_id` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0)` |
| IC-001042 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OrivaTransactionReport:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L97) | WHERE / P11 / COMPARISON | `ISNULL(lg.order_id, 0)` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0) = 0` |
| IC-001043 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OrivaTransactionReport:97](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L97) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.order_id` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0)` |
| IC-001044 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OrivaTransactionReport:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L98) | WHERE / P11 / COMPARISON | `ISNULL(lg.order_id, 0)` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0) != 0` |
| IC-001045 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OrivaTransactionReport:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L98) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.order_id` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0)` |
| IC-001046 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.OrivaTransactionReport:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L99) | WHERE / P11 / COMPARISON | `ISNULL(lg.order_id, 0)` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0) != 0` |
| IC-001047 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.OrivaTransactionReport:99](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.OrivaTransactionReport.sql#L99) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.order_id` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0)` |
| IC-001783 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RPAJurnalSettlementOriva:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql#L96) | WHERE / P11 / COMPARISON | `ISNULL(lg.order_id, 0)` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0) = 0` |
| IC-001784 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.RPAJurnalSettlementOriva:96](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql#L96) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.order_id` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0)` |
| IC-001785 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.RPAJurnalSettlementOriva:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql#L98) | WHERE / P11 / COMPARISON | `ISNULL(lg.order_id, 0)` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0) != 0` |
| IC-001786 | PARTNER_OR_DESTINATION / TYPE_ANCHOR | [dbo.RPAJurnalSettlementOriva:98](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RPAJurnalSettlementOriva.sql#L98) | EXPRESSION / P09 / ISNULL_REPLACEMENT | `lg.order_id` : `BIGINT` | `0` : `int` | `0` → `BIGINT` | `ISNULL(lg.order_id, 0)` |

<a id="kolom-status"></a>
## status

**Tipe dasar:** `VARCHAR(1)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.oriva_settlement_log.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012559 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSUpdatePostingResult:304](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdatePostingResult.sql#L304) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[status]` : `VARCHAR(1)` | `tpj.[Status]` : `CHAR(1)` | `tpj.[Status]` → `VARCHAR(1)` | `[status] = tpj.[Status]` |
