# Diagram Relasi Berdasarkan Implementasi

Dokumen ini menampilkan seluruh hubungan antarobject yang dapat dibuktikan dari source `db_object`. Sumbernya adalah foreign key pada `db_object/tables` serta equality JOIN pada `db_object/views` dan `db_object/procedures`. Nama kolom yang kebetulan sama tidak cukup untuk membuat hubungan.

## Cara membaca document

- `Relationship Group` adalah jaringan object yang masih tersambung melalui minimal satu jalur. Nomornya bukan urutan proses, prioritas, atau hierarchy.
- `Connection Set` adalah potongan kecil dari jaringan yang berisi maksimal delapan relationship lines. Set dibentuk dari relationship yang benar-benar tersambung, bukan dari object yang dipaksa menjadi parent atau Hub.
- Garis solid berarti terdapat `DECLARED_FK`. Garis putus-putus tanpa panah berarti hanya terdapat `OBSERVED_JOIN`. Diagram tidak menyatakan cardinality `1:1` atau `1:N`.
- Setiap label garis menampilkan exact column mapping dalam bentuk `left_column = right_column`, disertai `[FK]` atau `[JOIN]`. Statistik usage tidak ditampilkan karena bukan informasi struktur relasi.
- Detail join type, confidence, occurrences, dan source module tetap tersedia di `relationship_catalog.md` pada Group/Connection Set yang sama.
- `Continued through` menunjukkan shared object yang muncul lagi pada Connection Set lain; cari nama object tersebut di document untuk melanjutkan jalurnya.

Self-join tidak digambar sebagai garis karena dapat terlihat seperti satu record menunjuk dirinya sendiri. Seluruh kasusnya tetap tersedia pada `Self-join Review`.

Tidak ditemukan foreign key atau relasi statis yang dapat dibuktikan.
