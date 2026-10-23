# Retail UAT.dbo.TRSKuponCustomerSaldo_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 4 lokasi source, 1 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-processdate"></a>
## ProcessDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSKuponCustomerSaldo_TT.sql#L2). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002199 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSKuponProcess:255](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L255) | WHERE / P13 / COMPARISON | `ks.ProcessDate` : `DATETIME` | `@dWorkingDate` : `DATE` | `@dWorkingDate` → `DATETIME` | `ks.ProcessDate >= @dWorkingDate` |
| IC-002200 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSKuponProcess:256](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L256) | WHERE / P13 / COMPARISON | `ks.ProcessDate` : `DATETIME` | `DATEADD(DAY, 1, @dWorkingDate)` : `DATE` | `DATEADD(DAY, 1, @dWorkingDate)` → `DATETIME` | `ks.ProcessDate < DATEADD(DAY, 1, @dWorkingDate)` |

<a id="kolom-secid"></a>
## SecId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSKuponCustomerSaldo_TT.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002198 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSKuponProcess:253](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L253) | WHERE / P11 / COMPARISON | `ks.SecId` : `INT` | `@nSecId` : `BIGINT` | `ks.SecId` → `BIGINT` | `ks.SecId = @nSecId` |
| IC-002222 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSKuponProcess:428](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSKuponProcess.sql#L428) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trskuponcustomersaldo_tt.SecId` : `INT` | `ps.sec_id` : `BIGINT` | `ps.sec_id` → `INT` | `retail.dbo.trskuponcustomersaldo_tt.SecId <- ps.sec_id` |
