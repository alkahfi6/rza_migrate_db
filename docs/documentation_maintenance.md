# Pemeliharaan dan Praktik Dokumentasi

Dokumentasi software dimulai di [Engine Docs Guide](engine_docs_guide.md). Dokumen ini menjelaskan cara menjaga panduan tetap dapat dipakai, termasuk oleh agent yang baru masuk. Hierarki `rules -> spec -> fsd -> plan -> traceability -> QA` tidak berubah.

## Praktik yang dipakai dan sumbernya

| Praktik | Penerapan di repository | Referensi primer |
| --- | --- | --- |
| Pisahkan kebutuhan belajar, menyelesaikan task, lookup, dan memahami alasan | Tutorial offline, panduan agent/walkthrough, engine reference, architecture overview | [Diátaxis](https://diataxis.fr/) menguraikan empat bentuk dokumentasi tersebut. |
| Diagram bertingkat sesuai kebutuhan pembaca | Context/data flow terpisah dari package catalog dan function call tree | [C4 diagrams](https://c4model.com/diagrams) menjelaskan level zoom; panduan ini mengadaptasi prinsipnya, bukan mengklaim seluruh diagram sebagai notasi C4 formal. |
| Dokumentasi ikut version control/review/validation | Markdown dekat source, source links, runnable example, review evidence dalam commit | [Write the Docs: Docs as Code](https://www.writethedocs.org/guide/docs-as-code/) menjelaskan pemakaian workflow engineering untuk docs. |

Sumber diperiksa pada 2026-10-22. Praktik tersebut membantu navigasi dan review; aturan engine tetap berasal dari repository, bukan dari framework dokumentasi eksternal.

## Peta jenis dokumentasi

| Pertanyaan pembaca | Pemilik jawaban |
| --- | --- |
| Project ini untuk apa dan bagaimana mulai? | [README](../README.md), [tutorial](tutorials/offline_authoring.md) |
| Engine saling berhubungan bagaimana? | [Overview](architecture/overview.md) |
| Command memanggil bagian mana? | [Call tree](architecture/call_tree.md) |
| Paket ini bertugas apa? | [Package catalog](architecture/package_catalog.md) |
| Input, output, mode, side effect tiap aplikasi? | [Engine reference](engine_reference.md) |
| Aturan/requirement/desain lengkap? | [Rules](../rules.md), [spec](../spec.md), [FSD](../fsd.md) |
| Bagaimana mengubah dan menguji engine? | [Agent onboarding](engineering/agent_onboarding.md), QA gate, fixtures, traceability |
| Definisi dan histori database aplikasi? | Generated `docs/database/`; bukan source manual untuk arsitektur software. |
| Apa yang sudah diverifikasi pada perubahan docs ini? | [Review record](qa/documentation_review_20261022.md) |

## Kontrak penulisan

1. Mulai dengan tujuan dokumen, scope, dan link ke parent guide. Gunakan narasi Bahasa Indonesia serta nama teknis source yang persis.
2. Pisahkan requirement normatif, penjelasan implementasi, contoh, rencana, dan bukti run. Jangan mengubah status capability hanya karena dokumentasinya bertambah.
3. Setiap engine dijelaskan dengan input, output, mutation boundary, failure/exit, entrypoint, dan test yang relevan. Read-only source masih dapat menulis diagnostic.
4. Link relatif terhadap lokasi Markdown untuk repository portability. Gunakan nama file/simbol, bukan nomor baris sebagai identitas permanen.
5. Call tree memuat jalur penting, label conditional branch, callback bridge, dan side effect. Hindari dump seluruh helper yang cepat usang; periksa ulang nama fungsi bila source berubah.
6. Contoh harus menyebut shell, working directory, precondition/license, target sandbox, dan expected outcome. Jangan menjalankan extractor/generator/live reset pada repo/DB nyata sebagai contoh onboarding.
7. Dokumentasikan ambiguity, unsupported, dan skipped evidence. `go test` tanpa live DB bukan bukti runtime SQL; static JOIN bukan bukti constraint/uniqueness.
8. Hindari duplikasi normative lists. Ringkas lalu tautkan pemilik kontrak; jika menemukan kontradiksi, telusuri rule/capability/source/test dan selesaikan scope yang relevan secara eksplisit.

## Kapan memperbarui peta ini

| Perubahan source/kontrak | Dokumen navigasi yang ikut diperiksa |
| --- | --- |
| Entrypoint/flag/profile/output | Engine reference, call tree, walkthrough, README, spec |
| Package split atau ownership pindah | Package catalog, call tree, FSD |
| Parser/fixer/model baru | Overview boundary, call tree, agent impact map + minimum docs AGENTS |
| Report layout/inference | Engine reference, report layout, report docs generated sesuai kontraknya |
| Release digest/build/CI | Call tree release, engine reference, production readiness, guide |
| Test/QA gate atau accepted unsupported | Agent onboarding, QA gate/traceability/capability/backlog; regenerate coverage bila facts berubah |

Setiap perubahan behavior tetap mengikuti documentation map lengkap dalam [AGENTS.md](../AGENTS.md). Tabel ini menambah navigasi, bukan menggantikannya.

## Review dan validasi sebelum commit

- Periksa semua link relatif baru/diubah termasuk fragment heading, source path, kapitalisasi filename, serta fenced code yang tertutup.
- Cocokkan call tree dengan caller dan callee aktual, bukan hanya pencarian bahwa simbolnya ada. Periksa branch skip/fail, report writes, dan release wrapper.
- Jalankan tutorial pada direktori latihan baru. Catat outcome dan hash source/history untuk mode yang harus no-write; jalankan generation kedua untuk no-op.
- Jalankan `go test ./...` dan fixture/gate yang relevan. Review `git diff --check` serta generated diffs; jangan commit `out/` atau credential.
- Pada documentation-only change, buktikan tidak ada source input digest, executable, authored SQL, migration history, atau behavior yang berubah. Perubahan file `.md` di luar digest tidak memerlukan build/license baru. Jika source digest berubah, jalankan complete release workflow dan Linux gate setelah perubahan final.
- Catat dokumen updated dan reviewed unchanged beserta alasan, command/exit, validation limitation, dan baseline revision. Review record bukan janji bahwa CI remote/live DB sudah lulus.

Repository ini belum memiliki checker otomatis khusus seluruh link/symbol dokumentasi. Review 2026-10-22 menggunakan pemeriksaan lokal tambahan; jangan menyebutnya sebagai gate CI yang telah terpasang. Untuk evidence aktual baca [review record](qa/documentation_review_20261022.md).
