# REST API Hardening Policy (`rza_migrate_rest`)

Dokumen ini menetapkan baseline hardening untuk deployment REST migration API.

## 1. Prinsip

- Endpoint migration adalah high-impact.
- Default harus secure-by-default.
- Operasi destruktif (`down`, `run-file`) wajib explicit opt-in.
- Gunakan lock DB untuk mencegah race antar proses/host.

## 2. Baseline Wajib

`config.yaml` minimal:
- `security.enable_api_key=true`
- `security.api_key` bukan placeholder
- `migration.db_lock.enabled=true`
- `migration.allow_down=false`
- `migration.allow_run_specific_file=false`
- timeout operasi best-effort (hardcoded di service)

## 3. Profil Environment

### 3.1 Dev (lokal)

Tujuan: fleksibel debugging.

Rekomendasi:
- `enable_api_key=true` (tetap aktif)
- `allow_down=true` (opsional, jika memang perlu)
- `allow_run_specific_file=true` (opsional)
- `db_lock.enabled=true`

Catatan:
- Hindari expose ke network publik.

### 3.2 Staging

Tujuan: simulasi produksi.

Rekomendasi:
- `enable_api_key=true`
- `allow_down=false` (default)
- `allow_run_specific_file=false` (default)
- `db_lock.enabled=true`

Opsional terbatas:
- boleh aktifkan `allow_down=true` hanya saat rehearsal rollback, lalu nonaktifkan lagi.

### 3.3 Production

Tujuan: stabil dan minim blast radius.

Wajib:
- `enable_api_key=true`
- `allow_down=false`
- `allow_run_specific_file=false`
- `db_lock.enabled=true`

Tambahan disarankan:
- jalankan service di private network/VPN.
- pasang reverse proxy TLS + IP allowlist.
- audit log request (siapa memanggil endpoint dan kapan).

## 4. Timeout Behavior

Mode timeout operasi (best-effort, hardcoded):
- timeout dipakai untuk lock/preflight context.
- tidak menjamin query migrate in-flight langsung berhenti (batasan runtime/driver).

Implikasi:
- siapkan SOP operasional jika migrate lama/tergantung lock.
- observability tetap diperlukan (log + monitor).

## 5. Change Control

- Perubahan `allow_down` atau `allow_run_specific_file` di production harus melalui approval.
- Setelah incident/maintenance selesai, kembalikan ke baseline secure.
