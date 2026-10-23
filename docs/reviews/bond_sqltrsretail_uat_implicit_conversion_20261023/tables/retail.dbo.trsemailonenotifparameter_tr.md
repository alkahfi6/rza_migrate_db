# Retail UAT.dbo.TRSEmailOneNotifParameter_TR — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 14 lokasi source, 5 procedure, 1 kolom. **A:** 4 lokasi pada 2 procedure. **B:** 10 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-value"></a>
## Value

**Tipe dasar:** `VARCHAR(MAX)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSEmailOneNotifParameter_TR.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 14 lokasi / 5 SP. **A:** 4 lokasi / 2 SP. **Peran:** CONVERTED_COLUMN, SOURCE_TO_DESTINATION.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-012214 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSSentEmailHasilRevenueBondGagal:157](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql#L157) | JOIN_ON / P01 / COMPARISON | `np.Value` : `VARCHAR(MAX)` | `gp.GroupId` : `INT` | `np.Value` → `INT` | `gp.GroupId = np.Value` |
| IC-012218 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSSentEmailTransferAsset:92](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L92) | JOIN_ON / P01 / COMPARISON | `np.Value` : `VARCHAR(MAX)` | `gp.GroupId` : `INT` | `np.Value` → `INT` | `gp.GroupId = np.Value` |
| IC-012219 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSSentEmailTransferAsset:106](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L106) | JOIN_ON / P01 / COMPARISON | `np.Value` : `VARCHAR(MAX)` | `gp.GroupId` : `INT` | `np.Value` → `INT` | `gp.GroupId = np.Value` |
| IC-012220 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.TRSSentEmailTransferAsset:126](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L126) | JOIN_ON / P01 / COMPARISON | `np.Value` : `VARCHAR(MAX)` | `gp.GroupId` : `INT` | `np.Value` → `INT` | `gp.GroupId = np.Value` |
| IC-002978 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SRBIProcessEmailNotifMaturityDateH1:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH1.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ev.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-002979 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SRBIProcessEmailNotifMaturityDateH1:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH1.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mx.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-002989 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SRBIProcessEmailNotifMaturityDateH3:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH3.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ev.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-002990 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SRBIProcessEmailNotifMaturityDateH3:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH3.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mx.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-003000 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SRBIProcessEmailNotifMaturityDateH7:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH7.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ev.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-003001 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.SRBIProcessEmailNotifMaturityDateH7:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH7.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mx.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-012215 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSentEmailHasilRevenueBondGagal:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql#L159) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ev.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-012216 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSentEmailHasilRevenueBondGagal:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql#L159) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mx.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-012225 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSentEmailTransferAsset:241](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L241) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `ev.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-012226 | B_EXPRESSION_OR_BOUNDARY / SOURCE_TO_DESTINATION | [dbo.TRSSentEmailTransferAsset:241](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L241) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `mx.[Value]` : `VARCHAR(MAX)` | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
