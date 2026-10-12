param(
  [string]$RepoRoot = ".",
  [string]$ScenarioRelPath = "fixtures/qa_rules/07_vk_whitespace_only/input",
  [ValidateSet("verify", "gen", "verify_then_gen")]
  [string]$Command = "gen",
  [int]$Iterations = 5,
  [int]$ParallelWorkers = 8,
  [double]$RequireSpeedupPercent = [double]::NaN
)

$ErrorActionPreference = "Stop"
if ($Iterations -lt 1) {
  throw "Iterations must be >= 1"
}
if ($ParallelWorkers -lt 1) {
  throw "ParallelWorkers must be >= 1"
}

Set-Location $RepoRoot
$scenarioRoot = Join-Path (Get-Location) $ScenarioRelPath
if (-not (Test-Path $scenarioRoot)) {
  throw "Scenario path not found: $ScenarioRelPath"
}

$tmpRoot = Join-Path ([System.IO.Path]::GetTempPath()) ("rza_parallel_bench_" + [guid]::NewGuid().ToString("N"))
$null = New-Item -ItemType Directory -Path $tmpRoot -Force

try {
  $binPath = Join-Path $tmpRoot "rza_migrate_db.exe"
  & go build -buildvcs=false -o $binPath ./cmd/rza_migrate_db
  if ($LASTEXITCODE -ne 0) {
    throw "go build failed with exit code $LASTEXITCODE"
  }

  function Get-PerfMap([string]$OutputJsonPath) {
    $raw = Get-Content -Raw -Path $OutputJsonPath
    $obj = $raw | ConvertFrom-Json
    if ($null -eq $obj.notes -or $null -eq $obj.notes.perf_ms) {
      throw "perf_ms missing in output: $OutputJsonPath"
    }
    return $obj.notes.perf_ms
  }

  function Invoke-BenchRun([string]$Mode, [int]$Iteration) {
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

    & $binPath $Command -repo $workRoot -parallel-read $parallelRead -parallel-workers $ParallelWorkers -perf-notes true | Out-Host
    if ($LASTEXITCODE -ne 0) {
      throw "benchmark run failed mode=$Mode iter=$Iteration exit=$LASTEXITCODE"
    }

    $outPath = Join-Path $workRoot "out/output.json"
    if (-not (Test-Path $outPath)) {
      throw "output.json missing mode=$Mode iter=$Iteration"
    }

    $perf = Get-PerfMap $outPath
    $runTotal = [double]$perf.'run.total'

    return [PSCustomObject]@{
      mode       = $Mode
      iteration  = $Iteration
      run_total  = $runTotal
      perf_stage = $perf
    }
  }

  $rows = New-Object System.Collections.Generic.List[object]
  foreach ($mode in @("serial", "parallel")) {
    for ($i = 1; $i -le $Iterations; $i++) {
      $rows.Add((Invoke-BenchRun -Mode $mode -Iteration $i))
    }
  }

  function Build-Stats($modeRows) {
    $vals = @($modeRows | Select-Object -ExpandProperty run_total)
    $avg = ($vals | Measure-Object -Average).Average
    $min = ($vals | Measure-Object -Minimum).Minimum
    $max = ($vals | Measure-Object -Maximum).Maximum
    return [PSCustomObject]@{
      avg_ms = [math]::Round([double]$avg, 3)
      min_ms = [math]::Round([double]$min, 3)
      max_ms = [math]::Round([double]$max, 3)
    }
  }

  $serialRows = @($rows | Where-Object { $_.mode -eq "serial" })
  $parallelRows = @($rows | Where-Object { $_.mode -eq "parallel" })
  $serialStats = Build-Stats $serialRows
  $parallelStats = Build-Stats $parallelRows

  $speedupPercent = 0.0
  if ($serialStats.avg_ms -gt 0) {
    $speedupPercent = (($serialStats.avg_ms - $parallelStats.avg_ms) / $serialStats.avg_ms) * 100.0
  }
  $speedupPercent = [math]::Round($speedupPercent, 3)

  if (-not [double]::IsNaN($RequireSpeedupPercent)) {
    if ($speedupPercent -lt $RequireSpeedupPercent) {
      throw "parallel speedup $speedupPercent% is below required threshold $RequireSpeedupPercent%"
    }
  }

  $rowsSimple = @(
    $rows | ForEach-Object {
      [PSCustomObject]@{
        mode = $_.mode
        iteration = $_.iteration
        run_total = $_.run_total
      }
    }
  )

  $summary = @{
    scenario_rel_path = $ScenarioRelPath
    command = $Command
    iterations = $Iterations
    parallel_workers = $ParallelWorkers
    serial = $serialStats
    parallel = $parallelStats
    speedup_percent = $speedupPercent
    rows = $rowsSimple
  }

  $outDir = Join-Path (Get-Location) "out"
  $null = New-Item -ItemType Directory -Path $outDir -Force
  $summaryPath = Join-Path $outDir "parallel_benchmark_summary.json"
  ($summary | ConvertTo-Json -Depth 100) | Set-Content -Path $summaryPath -Encoding UTF8

  Write-Host "Parallel benchmark completed."
  Write-Host "serial avg(ms): $($serialStats.avg_ms)"
  Write-Host "parallel avg(ms): $($parallelStats.avg_ms)"
  Write-Host "speedup(%): $speedupPercent"
  Write-Host "Summary: $summaryPath"
  exit 0
}
finally {
  if (Test-Path $tmpRoot) {
    Remove-Item -Path $tmpRoot -Recurse -Force
  }
}
