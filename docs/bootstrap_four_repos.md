# Bootstrap Empat Repo Service

Gunakan alur yang sama secara terpisah untuk setiap pasangan server dan initial
catalog (`bond-sqlreplicate`, `bond-sqltrsretail`, `reksa-sqlreksa2`, dan
`reksa-sqlreplicate`). Jangan menyalin `migrations` atau manifest lintas catalog.

1. Build binary dengan version/commit/build time melalui
   `scripts/build_dist_windows.ps1`.
2. Di repo service tujuan, jalankan `rza_get_object_db -profile bootstrap`.
3. Review warning dan `unmanaged_objects` dalam
   `out/object_management_manifest.json`. Warning bukan bukti bahwa object aman
   diabaikan; object tersebut tetap dikelola manual sampai extractor mendukungnya.
4. Jalankan `rza_migrate_db bootstrap -repo .`. Default patch formatting adalah
   `off`, dan batch diteruskan sampai `pending_pairs=0`.
5. Untuk review bertahap, gunakan `draft`, review maksimal 999 pair, lalu `seal`;
   ulangi sampai tidak ada drift.
6. Commit pasangan `.up.sql`/`.down.sql`, desired state, dan evidence review.
   Jangan commit credential atau artifact lokal `out/`.
7. Deploy seluruh immutable history ke App-DB. File applied yang hilang atau
   berubah tetap menghentikan operasi dengan HTTP 409 dan kode error spesifik.

Manifest harus tetap tersedia sejak selesai extract sampai generation selesai.
Jika manifest hilang atau rusak, hentikan bootstrap dan ulangi extractor; jangan
mengarang entri unmanaged untuk melewati removal.
