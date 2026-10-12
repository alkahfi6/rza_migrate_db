param(
  [string]$RepoRoot = ".",
  [string]$ScenarioRelPath = "fixtures/qa_rules/07_vk_whitespace_only/input",
  [ValidateSet("verify", "gen", "verify_then_gen")]
  [string]$Command = "verify",
  [int]$Iterations = 10,
  [int]$ParallelWorkers = 8
)

$ErrorActionPreference = "Stop"
if ($Iterations -lt 2) {
  throw "Iterations must be >= 2"
}
if ($ParallelWorkers -lt 1) {
  throw "ParallelWorkers must be >= 1"
}

Set-Location $RepoRoot
$scenarioRoot = Join-Path (Get-Location) $ScenarioRelPath
if (-not (Test-Path $scenarioRoot)) {
  throw "Scenario path not found: $ScenarioRelPath"
}

$tmpRoot = Join-Path ([System.IO.Path]::GetTempPath()) ("rza_repeatability_" + [guid]::NewGuid().ToString("N"))
$null = New-Item -ItemType Directory -Path $tmpRoot -Force

try {
  $binPath = Join-Path $tmpRoot "rza_migrate_db.exe"
  & go build -buildvcs=false -o $binPath ./cmd/rza_migrate_db
  if ($LASTEXITCODE -ne 0) {
    throw "go build failed with exit code $LASTEXITCODE"
  }

  function Get-StringSha256([string]$Value) {
    $bytes = [System.Text.Encoding]::UTF8.GetBytes($Value)
    $hash = [System.Security.Cryptography.SHA256]::Create().ComputeHash($bytes)
    return ([System.BitConverter]::ToString($hash)).Replace("-", "").ToLowerInvariant()
  }

  function Get-CanonicalReportHash([string]$OutputJsonPath) {
    $raw = Get-Content -Raw -Path $OutputJsonPath
    $obj = $raw | ConvertFrom-Json
    $canon = [ordered]@{
      fail = $obj.fail
      pass = $obj.pass
      new_migrations = $obj.new_migrations
    }
    $json = $canon | ConvertTo-Json -Depth 100 -Compress
    return Get-StringSha256 $json
  }

  function Get-MigrationsTreeHash([string]$WorkRoot) {
    $migRoot = Join-Path $WorkRoot "migrations"
    if (-not (Test-Path $migRoot)) {
      return Get-StringSha256 ""
    }
    $lines = New-Object System.Collections.Generic.List[string]
    $files = Get-ChildItem -Path $migRoot -File | Sort-Object Name
    foreach ($f in $files) {
      $h = (Get-FileHash -Path $f.FullName -Algorithm SHA256).Hash.ToLowerInvariant()
      $lines.Add("$($f.Name)|$h")
    }
    return Get-StringSha256 ($lines -join "`n")
  }

  function Invoke-OneRun([string]$Mode, [int]$Iteration) {
    $parallelRead = if ($Mode -eq "parallel") { "true" } else { "false" }
    $workRoot = Join-Path $tmpRoot ("work_" + $Mode + "_" + $Iteration.ToString("00"))
    Copy-Item -Path $scenarioRoot -Destination $workRoot -Recurse -Force
    $licenseSource = Join-Path (Get-Location) "tools\license.key"
    if (-not (Test-Path $licenseSource)) {
      $licenseSource = Join-Path (Get-Location) "license.key"
    }
    $licenseTarget = Join-Path $workRoot "tools\license.key"
    if (-not (Test-Path $licenseTarget) -and (Test-Path $licenseSource)) {
      $null = New-Item -ItemType Directory -Path (Split-Path -Parent $licenseTarget) -Force
      Copy-Item -LiteralPath $licenseSource -Destination $licenseTarget -Force
    }

    & $binPath $Command -repo $workRoot -parallel-read $parallelRead -parallel-workers $ParallelWorkers -perf-notes false | Out-Host
    if ($LASTEXITCODE -ne 0) {
      throw "run failed mode=$Mode iter=$Iteration exit=$LASTEXITCODE"
    }

    $outPath = Join-Path $workRoot "out/output.json"
    if (-not (Test-Path $outPath)) {
      throw "output.json missing mode=$Mode iter=$Iteration"
    }

    $reportHash = Get-CanonicalReportHash $outPath
    $migrationsHash = Get-MigrationsTreeHash $workRoot
    return [PSCustomObject]@{
      mode            = $Mode
      iteration       = $Iteration
      report_hash     = $reportHash
      migrations_hash = $migrationsHash
      signature       = "$reportHash|$migrationsHash"
    }
  }

  $results = New-Object System.Collections.Generic.List[object]
  foreach ($mode in @("serial", "parallel")) {
    for ($i = 1; $i -le $Iterations; $i++) {
      $results.Add((Invoke-OneRun -Mode $mode -Iteration $i))
    }
  }

  $serial = $results | Where-Object { $_.mode -eq "serial" }
  $parallel = $results | Where-Object { $_.mode -eq "parallel" }
  $serialUnique = @($serial | Select-Object -ExpandProperty signature -Unique)
  $parallelUnique = @($parallel | Select-Object -ExpandProperty signature -Unique)

  if ($serialUnique.Count -ne 1) {
    throw "serial repeatability failed: found $($serialUnique.Count) unique signatures"
  }
  if ($parallelUnique.Count -ne 1) {
    throw "parallel repeatability failed: found $($parallelUnique.Count) unique signatures"
  }
  if ($serialUnique[0] -ne $parallelUnique[0]) {
    throw "serial vs parallel signature mismatch"
  }

  $summary = [ordered]@{
    scenario_rel_path      = $ScenarioRelPath
    command                = $Command
    iterations             = $Iterations
    parallel_workers       = $ParallelWorkers
    serial_signature       = $serialUnique[0]
    parallel_signature     = $parallelUnique[0]
    serial_iterations      = @($serial)
    parallel_iterations    = @($parallel)
    repeatability_verified = $true
  }

  $outDir = Join-Path (Get-Location) "out"
  $null = New-Item -ItemType Directory -Path $outDir -Force
  $summaryPath = Join-Path $outDir "parallel_repeatability_summary.json"
  ($summary | ConvertTo-Json -Depth 100) | Set-Content -Path $summaryPath -Encoding UTF8

  Write-Host "Parallel repeatability validation passed."
  Write-Host "Summary: $summaryPath"
  exit 0
}
finally {
  if (Test-Path $tmpRoot) {
    Remove-Item -Path $tmpRoot -Recurse -Force
  }
}
