# Pemeriksaan ulang: temuan yang sebelumnya terlewat

**Ya, ada yang terlewat.** Pemeriksaan ulang pada 23 Oktober 2026 menemukan tambahan konversi melalui view/CTE/UNION, hasil function dan aggregate, serta perbandingan input `CASE` dengan `WHEN`. Ada juga kandidat lama yang dicabut setelah tipe hasil function diperbaiki. Seluruh **1.063 procedure UAT** dibaca kembali dengan metadata **841 tabel** dari kedua snapshot; CFMAST bukan pusat pemeriksaan.

Baca [rangkuman tabel/kolom terbaru](TABLE_COLUMN_SUMMARY.md), [panduan pola](PATTERN_GUIDE.md), dan [cakupan serta metode](REVIEW.md). Dokumen ini membandingkan versi pertama review **20261023**, bukan langsung dengan audit 20261022 yang memakai definisi hitung berbeda.

## Perubahan hasil

| Ukuran statis | Sebelum | Setelah |
|---|---:|---:|
| Lokasi pasangan operand kategori A pada kolom fisik dalam predicate SP | 440 | **456** |
| Tabel dengan lokasi A | 56 | **57** |
| Kolom dengan lokasi A | 78 | **81** |
| Procedure berbeda dengan lokasi A | 146 | **149** |
| Tabel dalam semua conversion context SP, termasuk partner/destination | 407 | **412** |
| Kolom dalam semua context tersebut | 1.548 | **1.580** |
| Seluruh pasangan beda tipe, termasuk family rutin dan write boundary | 13.054 | **14.552** |
| Pasangan comparison yang belum dapat diikat | 1.288 | **591** |

Kategori A berarti aturan tipe memprediksi konversi pada kolom fisik dalam predicate. Ini **kandidat risiko akses index**, belum merupakan observasi scan atau kehilangan seek di server. Selisih 16 lokasi muncul di **7 SP**; empat di antaranya sudah mempunyai lokasi A lain, sehingga jumlah SP bertambah tiga.

Delta seluruh pasangan: **1.517 pasangan lokasi baru**, **82 pasangan berubah tipe/asal**, dan **19 pasangan lama tidak lagi dilaporkan**. Angka ini bukan jumlah query lambat atau temuan kritis. Lokasi diidentifikasi menggunakan database/file/baris/context/operator/operand, karena `site_id` dapat berubah saat pengurutan evidence berubah. Lihat [delta terukur](recheck_delta.json), [pasangan baru/reklasifikasi](recheck_added_or_retyped.csv), dan [pasangan yang dicabut](recheck_removed_pairs.csv).

## Tambahan 16 lokasi A: tabel, kolom, dan SP

Daftar lengkap operand, query, deklarasi asal, dan index tersedia di [CSV tambahan predicate](recheck_new_column_predicates.csv) serta halaman tabel yang ditautkan.

| Procedure : baris | Kolom fisik yang dikonversi | Pasangan / tipe tujuan | Lokasi baru | Arah perbaikan |
|---|---|---|---:|---|
| [REPOPopulateHapusTransaksi:55,56,93,94](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.REPOPopulateHapusTransaksi.sql#L55) | [REPOMasterTransaction_TT.NIKSeller / NIKInputer](tables/retail.dbo.repomastertransaction_tt.md), VARCHAR(10) | user_nisp_v.nik, family DECIMAL | 4 | Tetapkan satu kontrak NIK untuk source/view/staging/caller. Validasi ID nonnumeric dan leading zero sebelum memilih storage canonical. |
| [RepoPopulateParamHargaModal:64](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.RepoPopulateParamHargaModal.sql#L64) | [REPOParamTenor_TR.LastUpdateNIK](tables/retail.dbo.repoparamtenor_tr.md#kolom-lastupdatenik), VARCHAR(10) | user_nisp_v.nik, family DECIMAL | 1 | Selaraskan key audit-user dengan kontrak NIK yang sama; jangan CAST kolom per baris sebagai solusi index. |
| [TRSPopulateSettlementBIS4Detail:62,66](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4Detail.sql#L62) | [TRSSettlementBIS4_Temp.ParticipantSeller / ParticipantBuyer](tables/retail.dbo.trssettlementbis4_temp.md), VARCHAR(500) | BIS4Participant_v.ParticipantCode, NVARCHAR | 2 | Sesuaikan tipe dan panjang staging dengan domain participant Unicode sesudah memeriksa data dan interface. |
| [TRSPopulateSettlementBIS4RepoDetail:71,77](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateSettlementBIS4RepoDetail.sql#L71) | TRSSettlementBIS4_Temp.ParticipantSeller / ParticipantBuyer, VARCHAR(500) | BIS4Participant_v.ParticipantCode, NVARCHAR | 2 | Terapkan kontrak participant yang sama untuk jalur repo dan non-repo. |
| [TRSRPTMasterNasabahPerubahanObligasi2:261](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRPTMasterNasabahPerubahanObligasi2.sql#L261) | [PS_NDM_EMPLOYEES.EMPLID](tables/sql_replicate.dbo.ps_ndm_employees.md#kolom-emplid), CHAR(11) | #tmpMAIN.TellerId, INT | 1 | Periksa asal tipe SELECT INTO temp dan kontrak employee key, bukan hanya deklarasi parameter SP. |
| [TRSRptLaporanPajakBankBuy:426](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSRptLaporanPajakBankBuy.sql#L426) | [PPLMAPMasterWorkflow_TM.EMPLID](tables/sql_replicate.dbo.pplmapmasterworkflow_tm.md#kolom-emplid), VARCHAR(11) | td.NIK_CS, INT | 1 | Gunakan domain employee key yang tervalidasi; index EMPLID yang sudah ada tidak menghapus mismatch tipe. |
| [TRSSearchEmployee:16,26,34,42,51](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSSearchEmployee.sql#L16) | PS_NDM_EMPLOYEES.EMPLID, CHAR(11) | user_nisp_v.nik, family DECIMAL | 5 | Selaraskan hasil view dan key pegawai; uji branch pencarian serta pegawai dari setiap source. |

`TRSSettlementBIS4_Temp` pada tabel di atas adalah **tabel fisik bernama demikian**, bukan local temporary table `#...`. Nama tabel tidak menentukan jenis object.

### Mengapa view menyembunyikan mismatch

[user_nisp_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.user_nisp_v.sql) menyatukan NIK dari `SQL_REPLICATE.dbo.TLTEL.TLNUM DECIMAL(7,0)` dan `SQL_REPLICATE.dbo.PPLMAPMasterWorkflow_TM.EMPLID VARCHAR(11)` **melalui [PSEmployee_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.PSEmployee_v.sql)** menggunakan UNION ALL. View PSEmployee_v pada snapshot ini mengambil PPLMAPMasterWorkflow_TM; namanya tidak berarti sumbernya tabel PSEmployee. Family output `nik` menjadi DECIMAL; exact precision/scale hasil gabungan tidak disimpulkan oleh scanner. Join ke NIK VARCHAR atau EMPLID CHAR karena itu memerlukan rekonsiliasi numeric. Binder sebelumnya belum menelusuri bentuk UNION/CTE/TVF/VALUES yang membentuk schema view tersebut.

[BIS4Participant_v](../../../bond-sqltrsretail-uat/db_object/views/dbo.BIS4Participant_v.sql) mengambil `code NVARCHAR(32)` dari `SQL_REPLICATE.dbo.BIS4_PARTICIPANT`, lalu UNION dengan literal participant tambahan. Kolom lokal VARCHAR dibandingkan dengan output Unicode tersebut. Sebuah alias view tidak cukup sebagai bukti tipe: periksa expression setiap branch dan metadata tabel dasar.

Jangan membuat index baru semata-mata untuk menutup mismatch ini. Selaraskan tipe terlebih dahulu, lalu periksa actual plan dan kebutuhan key/include. CAST eksplisit pada kolom indexed hanya membuat konversi terlihat; belum tentu memulihkan seek.

## CASE: perbandingan WHEN sempat tidak diperiksa

Versi awal sudah memeriksa tipe hasil THEN/ELSE, tetapi melewatkan comparison pada simple CASE:

```sql
-- Bentuk aktual ORIOnlineBrowseTrx:98; KemenkeuStatusId VARCHAR(20)
CASE trx.KemenkeuStatusId WHEN 3 THEN ...
-- VARCHAR dibandingkan dengan INT, terpisah dari tipe hasil THEN/ELSE.
```

Ditemukan **390 pasangan beda tipe** pada input/WHEN: 311 family numerik, 60 family teks, dan **19 text-column → numeric**. Kesembilan belas berada di:

- [ORIOnlineBrowseTrx:98–100](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineBrowseTrx.sql#L98): KemenkeuStatusId vs angka 3/4/5.
- [ORIOnlineReportDailyTransaction:177–179](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.ORIOnlineReportDailyTransaction.sql#L177): pola status yang sama.
- [TRSPopulateGagalBayarOri:130–132,217–219,329–331,453–455](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSPopulateGagalBayarOri.sql#L130): empat rangkaian status.
- [TRSReportSKFBankBeli:59](../../../bond-sqltrsretail-uat/db_object/procedures/dbo.TRSReportSKFBankBeli.sql#L59): `TreasuryCustomer_TM.JenisIdentitas CHAR(1)` vs `WHEN 1`.

Jika domain status/jenis identitas adalah teks, pertimbangkan WHEN `'3'`/`'4'`/`'5'` atau `'1'`, setelah memastikan kesetaraan terhadap nilai seperti `'03'`, spasi, NULL, dan kode nonnumeric. Comparison numeric menyamakan sebagian representasi teks yang mungkin berbeda bila menjadi comparison string. Lokasi tersebut berada pada expression/CASE_TEST dan dilaporkan sebagai **B**, bukan ditambahkan ke 456 lokasi A. Aturan input/WHEN dan tipe hasil dijelaskan terpisah dalam [Microsoft CASE](https://learn.microsoft.com/en-us/sql/t-sql/language-elements/case-transact-sql).

## NULLIF, return type, dan koreksi kandidat lama

- **NULLIF** sebelumnya belum diperiksa sebagai comparison. Ada 10 pasangan beda tipe, seluruhnya family numerik pada snapshot ini. `NULLIF(text_column, 0)` membandingkan teks dengan numeric, tetapi return type tetap mengikuti argumen pertama. Ini berbeda dari replacement rule ISNULL. Lihat [Microsoft NULLIF](https://learn.microsoft.com/en-us/sql/t-sql/language-elements/nullif-transact-sql).
- **Return type scalar UDF, SUM/AVG, arithmetic, TVF deklaratif, VALUES, dan UNION** kini dipropagasikan pada bentuk yang didukung. SUM/AVG pada SMALLINT menghasilkan INT, pada SMALLMONEY menghasilkan MONEY, dan pada REAL menghasilkan FLOAT; exact facets tetap tidak diklaim. Lihat [SUM](https://learn.microsoft.com/en-us/sql/t-sql/functions/sum-transact-sql) dan [AVG](https://learn.microsoft.com/en-us/sql/t-sql/functions/avg-transact-sql).
- **LEFT/RIGHT menghasilkan VARCHAR/NVARCHAR**, sehingga tipe input numeric atau CHAR tidak boleh diteruskan mentah sebagai tipe hasil. Beberapa pasangan assignment VARCHAR ← LEFT(CHAR) tidak lagi dilaporkan sebagai mismatch family. Contohnya `TRSDataNasabahRetailInstitusi:386,387` dan `TRSRetailPopulateORIOrder:47`. Lihat [Microsoft LEFT](https://learn.microsoft.com/en-us/sql/t-sql/functions/left-transact-sql).
- Pada `TRSPopulateWarganegaraNasabah`, outer comparison dua hasil RIGHT kini dibaca sebagai teks vs teks. **Konversi pada operator `+` di dalam expression masih perlu ditinjau**: penggabungan string dengan CFCIF# BIGINT dapat mengikuti numeric precedence. Dicabutnya satu pasangan outer tidak membersihkan seluruh expression.

Ke-19 pasangan yang tidak lagi dilaporkan bukan otomatis 19 query yang sepenuhnya aman. Reklasifikasi tipe/asal dan hasil function dapat menghapus mismatch pada satu level, sementara conversion pada level lain tetap ada.

## View/function pendukung juga mempunyai dampak

Contoh tambahan: [PSEmailAddress_REP_v:27,34](../../../bond-sqltrsretail-uat/db_object/views/dbo.PSEmailAddress_REP_v.sql#L27) menggabungkan AdditionalUser dari VALUES ber-ID INT dengan PPLMAPMasterWorkflow_TM.EMPLID VARCHAR(11). Comparison dan UNION output-nya perlu kontrak tipe yang konsisten.

Untuk pasangan baru/reklasifikasi dalam view/function, [ledger jalur pendukung](recheck_supporting_impacts.csv) memuat **641 hubungan site → procedure** yang dapat dijangkau secara statis, melibatkan **102 SP**. Satu body site dapat mencapai banyak SP; satu SP dapat mempunyai banyak jalur. Ini **kemungkinan dampak**, bukan bukti setiap branch dieksekusi atau 641 conversion tambahan pada predicate SP. Angka ini tidak dijumlahkan ke kategori A procedure-local. Ringkasan semua body module ada di [SUPPORTING_MODULES.md](SUPPORTING_MODULES.md).

## Yang masih belum dapat dibuktikan penuh

| Batas | Bukti dan konsekuensi |
|---|---|
| Dynamic SQL | [75 execution site pada 53 SP](remaining_dynamic_sql.csv). Teks final, reaching assignment/branch, dan deklarasi parameter runtime belum diikat penuh. Ledger menyimpan lokasi EXEC, bukan klaim isi query dinamis sudah direview lengkap. |
| Parse failure | `TRSPopulateHapusTransaksiSuratBerharga:190` mempunyai syntax error pada `sttt.InsertedBy,(unInput.fullname, '') as NameInputer`; AST parsial tidak membuktikan sisa procedure bersih. Source tidak diperbaiki dalam audit. |
| Binding statis | 591 pasangan comparison dan 8.664 observasi extension belum dapat diselesaikan. Angka extension meliputi NULL/context dan expression nonkolom, bukan 8.664 query berbahaya. |
| Temp lifecycle | Tidak ditemukan multiple CREATE/SELECT INTO untuk nama temp yang sama dalam satu SP pada AST yang tersedia; [ledger](remaining_temp_lifecycle.csv) hanya berisi header. Ini tidak memodelkan temp lintas callee/dynamic/branch penuh. |
| Konversi lain | Panjang, precision/scale, collation, UDT, internal coercion setiap built-in, implicit EXEC parameter binding, application/driver binding, dan metadata database eksternal belum dicakup penuh. Base type sama tidak berarti kontrak identik. |
| Runtime | Tidak ada actual execution plan, data profiling, execution frequency, logical reads/CPU/duration, atau runtime error yang diukur. Tidak ada klaim pasti index seek hilang atau estimasi peningkatan performa. |

Dengan batas tersebut, kesimpulannya adalah **cakupan statis diperluas dan gap nyata diperbaiki**, bukan sertifikasi bahwa seluruh conversion sudah ditemukan. Prioritas lanjutan yang paling bernilai ialah mengambil generated SQL + parameter dari 53 SP dinamis, melengkapi procedure yang gagal parse, dan menguji actual plan pada kolom A yang paling sering dipakai workload.

## Reproduksi dan validasi

Baseline pertama disimpan lossless dalam `baseline/*.gz`; hash payload asal ada di [recheck_baseline.json](recheck_baseline.json). Baseline validation.json di dalam arsip adalah bukti versi lama, **bukan** validasi hasil terbaru. `baseline/summary.json` memudahkan reproduksi delta tanpa mengekstrak arsip besar.

```powershell
python review_conversions.py --workspace <repo> --retail-ast <retail_ast.jsonl> --replicate-ast <replicate_ast.jsonl>
python recheck_results.py --workspace <repo> --baseline <review-dir>/baseline --retail-ast <retail_ast.jsonl> --replicate-ast <replicate_ast.jsonl>
python render_review.py
python validate_review.py --workspace <repo> --dll <Microsoft.SqlServer.TransactSql.ScriptDom.dll>
```

Lihat [validation.json](validation.json) untuk hasil pemeriksaan **terbaru**, hash input/artefak, dan batas pengujian. Regression offline mencakup simple CASE, NULLIF comparison vs return type, scalar UDF, aggregate, Unicode REPLACE, LEFT/RIGHT, UNION view/CTE, VALUES, TVF, SELECT INTO berurutan, serta penghentian rekursi pada expression yang tipenya tidak diketahui. Pemeriksaan juga merekonsiliasi seluruh lokasi ke dokumen tabel dan menjaga **7.856 file SQL source/migration** tetap identik. Tidak ada perubahan engine produksi, SQL object, index live, atau migration; karena itu build release engine tidak berlaku untuk perubahan laporan ini.
