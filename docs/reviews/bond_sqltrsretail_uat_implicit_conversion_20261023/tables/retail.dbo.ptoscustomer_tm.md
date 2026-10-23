# Retail UAT.dbo.PTOSCustomer_TM — review konversi implisit

[Kembali ke rangkuman](../TABLE_COLUMN_SUMMARY.md) · [Panduan pola](../PATTERN_GUIDE.md)

**Cakupan procedure UAT:** 64 lokasi source, 13 procedure, 18 kolom. **A:** 22 lokasi pada 5 procedure. **B:** 0 lokasi. Jumlah context memasukkan pasangan/destination; tidak semuanya conversion kolom indexed.

Semua lokasi untuk tabel ini dicantumkan di bawah tanpa pemilihan contoh saja. Untuk filter/olah data gunakan column_occurrences.csv; site_id menghubungkan dua operand dari satu kejadian ke conversion_sites.csv. Metadata/index berasal dari snapshot, bukan live catalog.

<a id="kolom-sid"></a>
## SID

**Tipe dasar:** `VARCHAR(100)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L5). **Existing leading index:** ix_ptoscustomer_tm_sidnordn.

**Frekuensi source:** 22 lokasi / 5 SP. **A:** 21 lokasi / 5 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P02 — ANSI bertemu Unicode:** Samakan tipe/kolasi staging dengan key tujuan sesudah validasi panjang dan karakter. Jangan menurunkan Unicode jika karakter akan hilang; perubahan storage memerlukan review tersendiri.
- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001167 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:128](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L128) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001170 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:142](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L142) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001191 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:167](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L167) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001195 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:180](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L180) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001208 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:201](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L201) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001211 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:211](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L211) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001212 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:214](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L214) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001230 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:232](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L232) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001232 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:237](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L237) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001261 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L51) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001263 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:58](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L58) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001273 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstanding:68](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstanding.sql#L68) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001355 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:54](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L54) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001357 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L61) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001367 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobOutstandingBond:71](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobOutstandingBond.sql#L71) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001470 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:61](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L61) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001472 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:67](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L67) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001490 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksi:79](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksi.sql#L79) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001598 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:60](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L60) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR (expression; length unknown)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001600 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:66](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L66) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR(255)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001618 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobTransaksiBond:77](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobTransaksiBond.sql#L77) | JOIN_ON / P02 / COMPARISON | `tm.[SID]` : `VARCHAR(100)` | `tmp.[SID]` : `NVARCHAR (expression; length unknown)` | `tm.[SID]` → `NVARCHAR` | `tmp.[SID] = tm.[SID]` |
| IC-001226 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.SID` : `VARCHAR(100)` | `[SID]` : `NVARCHAR(255)` | `[SID]` → `VARCHAR(100)` | `retail.dbo.ptoscustomer_tm.SID <- [SID]` |

<a id="kolom-norekinvestorrdn"></a>
## NoRekInvestorRDN

**Tipe dasar:** `VARCHAR(20)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L6). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 1 lokasi / 1 SP. **Peran:** CONVERTED_COLUMN.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001233 | A_COLUMN_PREDICATE / CONVERTED_COLUMN | [dbo.PTOSProsesDataJobNasabah:238](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L238) | JOIN_ON / P01 / COMPARISON | `tm.NoRekInvestorRDN` : `VARCHAR(20)` | `CONVERT(bigint,tmp.[Nomor Rekening RDN])` : `bigint` | `tm.NoRekInvestorRDN` → `BIGINT` | `CONVERT(bigint,tmp.[Nomor Rekening RDN]) = tm.NoRekInvestorRDN` |

<a id="kolom-alamatktp"></a>
## AlamatKTP

**Tipe dasar:** `VARCHAR(1000)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L11). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001199 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:187](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L187) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.AlamatKTP` : `VARCHAR(1000)` | `tmp.[Alamat KTP]` : `NVARCHAR(MAX)` | `tmp.[Alamat KTP]` → `VARCHAR(1000)` | `tm.AlamatKTP = tmp.[Alamat KTP]` |
| IC-001213 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.AlamatKTP` : `VARCHAR(1000)` | `[Alamat KTP]` : `NVARCHAR(MAX)` | `[Alamat KTP]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_tm.AlamatKTP <- [Alamat KTP]` |

<a id="kolom-alamatkorespondensi"></a>
## AlamatKorespondensi

**Tipe dasar:** `VARCHAR(1000)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L12). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001200 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:188](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L188) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.AlamatKorespondensi` : `VARCHAR(1000)` | `tmp.[Alamat Korespondensi]` : `NVARCHAR(MAX)` | `tmp.[Alamat Korespondensi]` → `VARCHAR(1000)` | `tm.AlamatKorespondensi = tmp.[Alamat Korespondensi]` |
| IC-001214 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.AlamatKorespondensi` : `VARCHAR(1000)` | `[Alamat Korespondensi]` : `NVARCHAR(MAX)` | `[Alamat Korespondensi]` → `VARCHAR(1000)` | `retail.dbo.ptoscustomer_tm.AlamatKorespondensi <- [Alamat Korespondensi]` |

<a id="kolom-approveddate"></a>
## ApprovedDate

**Tipe dasar:** `DATETIME`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L28). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 1 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-007096 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomerAuthen:406](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomerAuthen.sql#L406) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pctm.ApprovedDate` : `DATETIME` | `tmp.ApprovedBy` : `int` | `tmp.ApprovedBy` → `DATETIME` | `pctm.ApprovedDate = tmp.ApprovedBy` |

<a id="kolom-email"></a>
## Email

**Tipe dasar:** `VARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L18). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001206 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:194](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L194) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.Email` : `VARCHAR(50)` | `tmp.[Alamat Email]` : `NVARCHAR(255)` | `tmp.[Alamat Email]` → `VARCHAR(50)` | `tm.Email = tmp.[Alamat Email]` |
| IC-001215 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.Email` : `VARCHAR(50)` | `[Alamat Email]` : `NVARCHAR(255)` | `[Alamat Email]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_tm.Email <- [Alamat Email]` |

<a id="kolom-insertedby"></a>
## InsertedBy

**Tipe dasar:** `INT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L23). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** COUNTERPART_COLUMN_NOT_CONVERTED.

**Pola dan arah review:**

- **P01 — Kolom teks bertemu numerik:** Samakan domain key; parameter mengikuti tipe kolom tujuan. Untuk dua kolom beda storage, validasi dan materialisasikan representasi canonical sekali. Jangan CAST kolom indexed sebagai perbaikan otomatis.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002369 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L51) | JOIN_ON / P01 / COMPARISON | `ptcm.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptcm.InsertedBy = ndme.EMPLID` |
| IC-005828 | PARTNER_OR_DESTINATION / COUNTERPART_COLUMN_NOT_CONVERTED | [dbo.TRSPopulatePTOSUploadAuthen:50](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L50) | JOIN_ON / P01 / COMPARISON | `ptcm.InsertedBy` : `INT` | `ndme.EMPLID` : `CHAR(11)` | `ndme.EMPLID` → `INT` | `ptcm.InsertedBy = ndme.EMPLID` |

<a id="kolom-jeniskelamin"></a>
## JenisKelamin

**Tipe dasar:** `VARCHAR(6)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L13). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001201 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:189](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L189) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.JenisKelamin` : `VARCHAR(6)` | `UPPER(tmp.[Jenis Kelamin])` : `NVARCHAR (expression; length unknown)` | `UPPER(tmp.[Jenis Kelamin])` → `VARCHAR(6)` | `tm.JenisKelamin = UPPER(tmp.[Jenis Kelamin])` |
| IC-001216 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.JenisKelamin` : `VARCHAR(6)` | `UPPER([Jenis Kelamin])` : `NVARCHAR (expression; length unknown)` | `UPPER([Jenis Kelamin])` → `VARCHAR(6)` | `retail.dbo.ptoscustomer_tm.JenisKelamin <- UPPER([Jenis Kelamin])` |

<a id="kolom-npwp"></a>
## NPWP

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L8). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001197 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:184](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L184) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.NPWP` : `VARCHAR(40)` | `tmp.[Nomor NPWP]` : `NVARCHAR(255)` | `tmp.[Nomor NPWP]` → `VARCHAR(40)` | `tm.NPWP = tmp.[Nomor NPWP]` |
| IC-001217 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.NPWP` : `VARCHAR(40)` | `[Nomor NPWP]` : `NVARCHAR(255)` | `[Nomor NPWP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_tm.NPWP <- [Nomor NPWP]` |

<a id="kolom-namanasabah"></a>
## NamaNasabah

**Tipe dasar:** `VARCHAR(95)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L9). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001198 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:185](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L185) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.NamaNasabah` : `VARCHAR(95)` | `tmp.[Nama Nasabah]` : `NVARCHAR(255)` | `tmp.[Nama Nasabah]` → `VARCHAR(95)` | `tm.NamaNasabah = tmp.[Nama Nasabah]` |
| IC-001218 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.NamaNasabah` : `VARCHAR(95)` | `[Nama Nasabah]` : `NVARCHAR(255)` | `[Nama Nasabah]` → `VARCHAR(95)` | `retail.dbo.ptoscustomer_tm.NamaNasabah <- [Nama Nasabah]` |

<a id="kolom-namaperusahaanefek"></a>
## NamaPerusahaanEfek

**Tipe dasar:** `VARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L19). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001207 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:199](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L199) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.NamaPerusahaanEfek` : `VARCHAR(50)` | `tmp.[Nama Perusahaan Efek]` : `NVARCHAR(255)` | `tmp.[Nama Perusahaan Efek]` → `VARCHAR(50)` | `tm.NamaPerusahaanEfek = tmp.[Nama Perusahaan Efek]` |
| IC-001219 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.NamaPerusahaanEfek` : `VARCHAR(50)` | `[Nama Perusahaan Efek]` : `NVARCHAR(255)` | `[Nama Perusahaan Efek]` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_tm.NamaPerusahaanEfek <- [Nama Perusahaan Efek]` |

<a id="kolom-nofax"></a>
## NoFax

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L17). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001205 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:193](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L193) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.NoFax` : `VARCHAR(40)` | `tmp.[Nomor Fax]` : `NVARCHAR(255)` | `tmp.[Nomor Fax]` → `VARCHAR(40)` | `tm.NoFax = tmp.[Nomor Fax]` |
| IC-001220 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.NoFax` : `VARCHAR(40)` | `[Nomor Fax]` : `NVARCHAR(255)` | `[Nomor Fax]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_tm.NoFax <- [Nomor Fax]` |

<a id="kolom-nohp"></a>
## NoHP

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L16). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001204 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L192) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.NoHP` : `VARCHAR(40)` | `tmp.[Nomor HP]` : `NVARCHAR (expression; length unknown)` | `tmp.[Nomor HP]` → `VARCHAR(40)` | `tm.NoHP = tmp.[Nomor HP]` |
| IC-001221 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.NoHP` : `VARCHAR(40)` | `[Nomor HP]` : `NVARCHAR (expression; length unknown)` | `[Nomor HP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_tm.NoHP <- [Nomor HP]` |

<a id="kolom-noidentitas"></a>
## NoIdentitas

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L7). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001196 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:183](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L183) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.NoIdentitas` : `VARCHAR(40)` | `tmp.[Nomor KTP]` : `NVARCHAR(255)` | `tmp.[Nomor KTP]` → `VARCHAR(40)` | `tm.NoIdentitas = tmp.[Nomor KTP]` |
| IC-001222 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.NoIdentitas` : `VARCHAR(40)` | `[Nomor KTP]` : `NVARCHAR(255)` | `[Nomor KTP]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_tm.NoIdentitas <- [Nomor KTP]` |

<a id="kolom-notlp"></a>
## NoTlp

**Tipe dasar:** `VARCHAR(40)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L15). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001203 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:191](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L191) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.NoTlp` : `VARCHAR(40)` | `tmp.[Nomor Telepon]` : `NVARCHAR (expression; length unknown)` | `tmp.[Nomor Telepon]` → `VARCHAR(40)` | `tm.NoTlp = tmp.[Nomor Telepon]` |
| IC-001223 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.NoTlp` : `VARCHAR(40)` | `[Nomor Telepon]` : `NVARCHAR (expression; length unknown)` | `[Nomor Telepon]` → `VARCHAR(40)` | `retail.dbo.ptoscustomer_tm.NoTlp <- [Nomor Telepon]` |

<a id="kolom-ptostype"></a>
## PTOSType

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L20). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 2 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001224 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.PTOSType` : `SMALLINT` | `1` : `int` | `1` → `SMALLINT` | `retail.dbo.ptoscustomer_tm.PTOSType <- 1` |
| IC-002320 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L33) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.PTOSType` : `SMALLINT` | `PTOSType` : `INT` | `PTOSType` → `SMALLINT` | `retail.dbo.ptoscustomer_tm.PTOSType <- PTOSType` |

<a id="kolom-pekerjaan"></a>
## Pekerjaan

**Tipe dasar:** `VARCHAR(50)`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L14). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 2 lokasi / 1 SP. **A:** 0 lokasi / 0 SP. **Peran:** DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-001202 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:190](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L190) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `tm.Pekerjaan` : `VARCHAR(50)` | `tmp.[Pekerjaan]` : `NVARCHAR(255)` | `tmp.[Pekerjaan]` → `VARCHAR(50)` | `tm.Pekerjaan = tmp.[Pekerjaan]` |
| IC-001225 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.Pekerjaan` : `VARCHAR(50)` | `Pekerjaan` : `NVARCHAR(255)` | `Pekerjaan` → `VARCHAR(50)` | `retail.dbo.ptoscustomer_tm.Pekerjaan <- Pekerjaan` |

<a id="kolom-status"></a>
## Status

**Tipe dasar:** `SMALLINT`. [Deklarasi kolom](../../../../bond-sqltrsretail-uat/db_object/tables/dbo.PTOSCustomer_TM.sql#L21). **Existing leading index:** Tidak ada yang tercatat.

**Frekuensi source:** 12 lokasi / 9 SP. **A:** 0 lokasi / 0 SP. **Peran:** CONVERTED_COLUMN, DESTINATION_COLUMN.

**Pola dan arah review:**

- **P08 — Assignment/INSERT/UPDATE mengikuti tipe tujuan:** Konversi menuju tipe destination, bukan selalu highest precedence. Selaraskan staging; uji truncation, Unicode loss, overflow, rounding, tanggal, dan error handling. Bukan bukti access predicate kehilangan seek.
- **P11 — Perbedaan keluarga numerik:** BIT/TINYINT/SMALLINT/INT atau money/decimal dapat berbeda tipe secara rutin. Jangan otomatis menaikkan severity; cek precision/scale dan actual plan bila relevan.

### Daftar procedure dan bagian query

| Site | Tier / peran | Procedure : baris | Context / pola | Operand kolom : tipe | Pasangan : tipe | Operand yang dikonversi → tipe hasil | Query atau mapping assignment |
|---|---|---|---|---|---|---|---|
| IC-002370 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.Rvmp_TRSPopulatePTOSUploadAuthen:52](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPopulatePTOSUploadAuthen.sql#L52) | WHERE / P11 / COMPARISON | `ptcm.Status` : `SMALLINT` | `0` : `int` | `ptcm.Status` → `INT` | `ptcm.Status = 0` |
| IC-005747 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSCustomer_TM:31](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSCustomer_TM.sql#L31) | WHERE / P11 / COMPARISON | `Status` : `SMALLINT` | `3` : `int` | `Status` → `INT` | `Status = 3` |
| IC-005829 | C_FAMILY_OR_ROUTINE / CONVERTED_COLUMN | [dbo.TRSPopulatePTOSUploadAuthen:51](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulatePTOSUploadAuthen.sql#L51) | WHERE / P11 / COMPARISON | `ptcm.Status` : `SMALLINT` | `0` : `int` | `ptcm.Status` → `INT` | `ptcm.Status = 0` |
| IC-001227 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.PTOSProsesDataJobNasabah:216](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.PTOSProsesDataJobNasabah.sql#L216) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.ptoscustomer_tm.Status <- 3` |
| IC-002321 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.Rvmp_TRSPTOSSubmitFromTemp:33](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.Rvmp_TRSPTOSSubmitFromTemp.sql#L33) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.ptoscustomer_tm.Status <- 0` |
| IC-007095 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomerAuthen:402](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomerAuthen.sql#L402) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pctm.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `pctm.Status = 3` |
| IC-007099 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomerAuthen:418](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomerAuthen.sql#L418) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `retail.dbo.ptoscustomer_tm.Status <- 3` |
| IC-007102 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomerAuthen:491](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomerAuthen.sql#L491) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `[Status] = 3` |
| IC-007105 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomerAuthen:504](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomerAuthen.sql#L504) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `[Status]` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `[Status] = 3` |
| IC-007112 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSCustomer_TM:139](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSCustomer_TM.sql#L139) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pctm.Status` : `SMALLINT` | `@pnStatus` : `INT` | `@pnStatus` → `SMALLINT` | `pctm.Status = @pnStatus` |
| IC-007195 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUpload:294](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUpload.sql#L294) | WRITE_ASSIGNMENT / P08 / INSERT_SELECT | `retail.dbo.ptoscustomer_tm.Status` : `SMALLINT` | `0` : `int` | `0` → `SMALLINT` | `retail.dbo.ptoscustomer_tm.Status <- 0` |
| IC-007214 | PARTNER_OR_DESTINATION / DESTINATION_COLUMN | [dbo.TRSProcessPTOSUploadAuthen:192](../../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSProcessPTOSUploadAuthen.sql#L192) | WRITE_ASSIGNMENT / P08 / ASSIGNMENTSETCLAUSE | `pctm.Status` : `SMALLINT` | `3` : `int` | `3` → `SMALLINT` | `pctm.Status = 3` |
