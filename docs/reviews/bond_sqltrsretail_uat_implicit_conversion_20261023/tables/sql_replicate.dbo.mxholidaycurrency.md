# SQL_REPLICATE.dbo.MXHolidayCurrency — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 4 procedure, 1 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-currency"></a>
## Currency

**Tipe dasar:** `NVARCHAR(40)`. [Deklarasi kolom](../../../../all_object/bond/bond-sqlreplicate/db_object/tables/dbo.MXHolidayCurrency.sql#L2). **Existing leading index:** IX_MXHolidayCurrency_Currency.

**Frekuensi source:** 15 lokasi / 4 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P04 — Nilai/parameter dikonversi; kolom menjadi pasangan:** Ini tidak otomatis menghilangkan seek pada kolom pasangan. Tetap validasi domain, overflow, format, estimasi, dan plan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-000109 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.CheckMXWorkingDate:23](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.CheckMXWorkingDate.sql#L23) | WHERE / P04 / COMPARISON | `Currency` : `NVARCHAR(40)` | `@pcCcy` : `varchar(3)` | `@pcCcy` → `NVARCHAR` | `Currency = @pcCcy` |
| IC-001771 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.REPOCheckMaturityDate:39](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.REPOCheckMaturityDate.sql#L39) | WHERE / P04 / COMPARISON | `Currency` : `NVARCHAR(40)` | `@CCY` : `VARCHAR(3)` | `@CCY` → `NVARCHAR` | `Currency = @CCY` |
| IC-003590 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSCheckRecordingDate2:48](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSCheckRecordingDate2.sql#L48) | WHERE / P04 / COMPARISON | `Currency` : `NVARCHAR(40)` | `@CCY` : `varchar(3)` | `@CCY` → `NVARCHAR` | `Currency = @CCY` |
| IC-004495 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:566](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L566) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ( 'IDR','USD')` |
| IC-004497 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:574](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L574) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ( 'IDR','EUR')` |
| IC-004499 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:582](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L582) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ( 'IDR','AUD')` |
| IC-004501 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:590](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L590) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ( 'IDR','CNH','HKD')` |
| IC-004507 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:619](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L619) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ('IDR', 'USD')` |
| IC-004509 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:627](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L627) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ('IDR', 'USD')` |
| IC-004511 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:635](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L635) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ('IDR', 'EUR')` |
| IC-004513 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:643](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L643) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ('IDR', 'EUR')` |
| IC-004515 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:651](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L651) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ('IDR', 'AUD')` |
| IC-004517 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:659](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L659) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ('IDR', 'AUD')` |
| IC-004519 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:667](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L667) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ('IDR', 'CNH','HKD')` |
| IC-004521 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSMaintenanceParamHargaModal:675](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSMaintenanceParamHargaModal.sql#L675) | JOIN_ON / P04 / IN_VALUE_TYPE | `hc.Currency` : `NVARCHAR(40)` | `'IDR'` : `varchar` | `'IDR'` → `NVARCHAR` | `hc.Currency in ('IDR', 'CNH','HKD')` |
