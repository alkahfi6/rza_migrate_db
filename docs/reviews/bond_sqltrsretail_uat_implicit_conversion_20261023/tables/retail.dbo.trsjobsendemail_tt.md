# Retail UAT.dbo.TRSJobSendEmail_TT — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 15 lokasi source, 8 procedure, 2 kolom. **A:** 0 lokasi pada 0 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-eventid"></a>
## EventId

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSJobSendEmail_TT.sql#L4). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 8 lokasi / 8 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002978 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH1:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH1.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` : `VARCHAR(MAX)` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-002989 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH3:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH3.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` : `VARCHAR(MAX)` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-003000 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH7:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH7.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` : `VARCHAR(MAX)` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-003349 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApproveSID:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveSID.sql#L172) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` : `VARCHAR(MAX)` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-006706 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovedUploadSRESID:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSRESID.sql#L169) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` : `VARCHAR(MAX)` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-007992 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSRepoInsertEmailProcess:133](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRepoInsertEmailProcess.sql#L133) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `[Value]` : `VARCHAR(MAX)` | `[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- [Value]` |
| IC-012215 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSentEmailHasilRevenueBondGagal:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql#L159) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` : `VARCHAR(MAX)` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |
| IC-012225 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSentEmailTransferAsset:241](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L241) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.EventId` : `INT` | `ev.[Value]` : `VARCHAR(MAX)` | `ev.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.EventId <- ev.[Value]` |

<a id="kolom-maxattempt"></a>
## MaxAttempt

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.TRSJobSendEmail_TT.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 7 lokasi / 7 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002979 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH1:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH1.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` : `VARCHAR(MAX)` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-002990 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH3:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH3.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` : `VARCHAR(MAX)` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-003001 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.SRBIProcessEmailNotifMaturityDateH7:129](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.SRBIProcessEmailNotifMaturityDateH7.sql#L129) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` : `VARCHAR(MAX)` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-003350 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSApproveSID:172](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSApproveSID.sql#L172) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` : `VARCHAR(MAX)` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-006707 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessApprovedUploadSRESID:169](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessApprovedUploadSRESID.sql#L169) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` : `VARCHAR(MAX)` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-012216 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSentEmailHasilRevenueBondGagal:159](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailHasilRevenueBondGagal.sql#L159) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` : `VARCHAR(MAX)` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
| IC-012226 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSSentEmailTransferAsset:241](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSentEmailTransferAsset.sql#L241) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.trsjobsendemail_tt.MaxAttempt` : `INT` | `mx.[Value]` : `VARCHAR(MAX)` | `mx.[Value]` → `INT` | `retail.dbo.trsjobsendemail_tt.MaxAttempt <- mx.[Value]` |
