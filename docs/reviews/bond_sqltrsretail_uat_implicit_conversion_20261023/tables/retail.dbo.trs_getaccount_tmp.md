# Retail UAT.dbo.trs_getaccount_tmp — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 3 lokasi source, 1 procedure, 3 kolom. **A:** 1 lokasi pada 1 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-acctno"></a>
## ACCTNO

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.trs_getaccount_tmp.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000507 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.GetAccountInfo:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L172) | JOIN_ON / P01 / COMPARISON | `tmp.ACCTNO` : `VARCHAR(20)` | `tms.[CFACC#]` : `DECIMAL(19,0)` | `tmp.ACCTNO` → `DECIMAL` | `tmp.ACCTNO = tms.[CFACC#]` |

<a id="kolom-ccy"></a>
## CCY

**Tipe dasar:** `VARCHAR(4)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.trs_getaccount_tmp.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000505 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GetAccountInfo:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L120) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trs_getaccount_tmp.CCY` : `VARCHAR(4)` | `a.CurrencyCode` : `CHAR(3)` | `a.CurrencyCode` → `VARCHAR(4)` | `retail.dbo.trs_getaccount_tmp.CCY <- a.CurrencyCode` |

<a id="kolom-session-ident"></a>
## SESSION_IDENT

**Tipe dasar:** `BIGINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.trs_getaccount_tmp.sql#L2). **Existing leading index:** PK_trs_getaccount_tmp.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000506 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.GetAccountInfo:120](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.GetAccountInfo.sql#L120) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trs_getaccount_tmp.SESSION_IDENT` : `BIGINT` | `CAST(@pnSessionId AS VARCHAR(19))` : `VARCHAR(19)` | `CAST(@pnSessionId AS VARCHAR(19))` → `BIGINT` | `retail.dbo.trs_getaccount_tmp.SESSION_IDENT <- CAST(@pnSessionId AS VARCHAR(19))` |
