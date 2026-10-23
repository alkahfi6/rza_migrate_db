# Retail UAT.dbo.TRSRepoNotif_TL — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 2 lokasi source, 1 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-guidprocess"></a>
## GuidProcess

**Tipe dasar:** `UNIQUEIDENTIFIER`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSRepoNotif_TL.sql#L3). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P13 — Perbedaan lain yang memerlukan verifikasi:** Periksa apakah konversi tipe tersebut diizinkan pada operator yang dipakai. Tanpa oracle SQL Server, sebagian kasus dapat merupakan type incompatibility/error alih-alih conversion yang sukses.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012644 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateStatusNotifRepo:28](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusNotifRepo.sql#L28) | WHERE / P13 / COMPARISON | `GuidProcess` : `UNIQUEIDENTIFIER` | `@cGuidProcess` : `NVARCHAR(100)` | `@cGuidProcess` → `UNIQUEIDENTIFIER` | `GuidProcess = @cGuidProcess` |
| IC-012645 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSUpdateStatusNotifRepo:55](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSUpdateStatusNotifRepo.sql#L55) | WHERE / P13 / COMPARISON | `GuidProcess` : `UNIQUEIDENTIFIER` | `@cGuidProcess` : `NVARCHAR(100)` | `@cGuidProcess` → `UNIQUEIDENTIFIER` | `GuidProcess = @cGuidProcess` |
