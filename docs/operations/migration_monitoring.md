# Migration Monitoring Runbook

`dbo.mig_operation_journal` adalah report tingkat request/release. `dbo.mig_event`
adalah detail keputusan setiap file (`run`, `skip`, `ok`, atau `fail`). Semua
timestamp mandatory menggunakan UTC.

## Ringkasan operasi terbaru

```sql
SELECT TOP (50)
  operation_id,
  request_id,
  operation,
  status,
  requested_steps,
  executed_steps,
  before_version,
  after_version,
  result_direction,
  manifest_sha256,
  error_message,
  created_at,
  finished_at
FROM dbo.mig_operation_journal
ORDER BY created_at DESC;
```

Interpretasi penting:

- `result_direction = up|down`: file benar-benar dijalankan.
- `result_direction = ack`: version diakui secara manual; file tidak dieksekusi.
- `status = no_change`: tidak ada perubahan version.
- `manifest_sha256`: checksum report yang mengikat urutan, nama, arah, dan
  SHA-256 setiap file dalam operasi.

## Daftar file dalam satu operasi

```sql
DECLARE @operation_id UNIQUEIDENTIFIER = '00000000-0000-0000-0000-000000000000';

SELECT
  j.operation_id,
  f.[version],
  f.file_name,
  f.direction,
  f.sha256
FROM dbo.mig_operation_journal AS j
CROSS APPLY OPENJSON(j.applied_files)
WITH (
  [version] BIGINT '$.version',
  file_name NVARCHAR(260) '$.file_name',
  direction NVARCHAR(8) '$.direction',
  sha256 CHAR(64) '$.sha256'
) AS f
WHERE j.operation_id = @operation_id
ORDER BY
  CASE WHEN f.direction = N'down' THEN -f.[version] ELSE f.[version] END;
```

## Detail event dan precheck

Event file dapat dikorelasikan dengan report melalui `version + file_name`.
Cara ini juga bekerja pada masa transisi ketika procedure logger lama belum
memiliki kolom `operation_id`.

```sql
DECLARE @operation_id UNIQUEIDENTIFIER = '00000000-0000-0000-0000-000000000000';

WITH applied AS (
  SELECT f.[version], f.file_name
  FROM dbo.mig_operation_journal AS j
  CROSS APPLY OPENJSON(j.applied_files)
  WITH (
    [version] BIGINT '$.version',
    file_name NVARCHAR(260) '$.file_name'
  ) AS f
  WHERE j.operation_id = @operation_id
)
SELECT
  e.created_at,
  e.version,
  e.file_name,
  e.direction,
  e.decision,
  e.severity,
  e.action,
  e.target,
  e.reason,
  e.message,
  e.duration_ms,
  e.error_number,
  e.error_line,
  e.error_procedure
FROM dbo.mig_event AS e
JOIN applied AS a
  ON a.[version] = e.[version]
 AND a.file_name COLLATE Latin1_General_100_BIN2 =
     e.file_name COLLATE Latin1_General_100_BIN2
ORDER BY e.event_id;
```

## Alert minimum production

Alert harus aktif jika salah satu kondisi berikut ditemukan:

```sql
SELECT operation_id, request_id, status, error_message, updated_at
FROM dbo.mig_operation_journal
WHERE status = N'failed'
   OR (status = N'running' AND updated_at < DATEADD(MINUTE, -15, SYSUTCDATETIME()));

SELECT event_id, version, file_name, target, message, detail, created_at
FROM dbo.mig_event
WHERE decision = N'fail' OR severity >= 3;
```

Sebelum setiap operasi, REST runner membandingkan file yang pernah tercatat
dengan SHA-256 di journal. File history yang berubah, hilang, atau mempunyai
bukti hash yang konflik akan memblokir migration.

Sebelum release dinyatakan selesai, simpan `operation_id`,
`manifest_sha256`, `before_version`, `after_version`, dan response API sebagai
bukti deployment. Jangan mengubah migration yang sudah pernah dijalankan;
perubahan body akan ditolak oleh verifikasi checksum.
