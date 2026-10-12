# Report object precision: bukti snapshot nyata

Fixture ini mempertahankan byte source dari snapshot lokal berikut; file SQL
tidak diperbaiki atau disederhanakan untuk membuat test lulus:

- `all_object/bond/bond-sqltrsretail/db_object`: procedure
  `Rvmp_trs_ListSecurityTransactionForInsertEditBlockTrans`, `TRSParamIssuerSave`
  dan delapan table yang relevan.
- `all_object/bond/bond-sqlreplicate/db_object`: table/index `JHCLDR` dan
  `MXHolidayCurrency`; empat CREATE INDEX memisahkan `ON` pada baris berikutnya.
- `all_object/reksa/reksa-sqlreksa2/db_object`: function `fnReksaGoodFund` dan
  table `ReksaProduct_TM`, `ReksaHolidayTable_TM`, `ReksaWindowPeriod_TM`.

Expected: LEFT OUTER JOIN OrderId tidak berubah menjadi INNER, static function
references menjadi consumer/lineage, empat multiline indexes terbaca,
`TRSParamIssuerSave` mempertahankan evidence `IssuerCode` yang tidak ada pada
snapshot `TRSParamIssuer_TH` sebagai `COLUMN_NOT_FOUND`.
Missing snapshot objects tetap memiliki resolution sendiri; bukan tebakan FK.

`TestRealObjectPrecisionFixtureReadOnlyAndIdempotent` menjalankan report dua kali,
membandingkan kontrak golden serta hash semua source/history. Negative tests
di `object_precision_test.go` mencakup nested query/CTE, collision, keyword
quoted, encoding, filename mismatch, dan target INSERT berqualifier yang tidak
boleh diubah menjadi nama kolom palsu.

Jalankan: `go test ./internal/migreport ./internal/sqlparse`.
Jika empat snapshot lokal tersedia, `TestReportRealSnapshotCorpus` juga mengaudit
seluruh corpus tanpa menulis source; file hasil report uji masuk temporary dir.
