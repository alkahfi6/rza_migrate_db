param(
  [string]$DistRoot = "dist/engine_windows_amd64",
  [string]$BuildVersion = "dev",
  [string]$BuildCommit = "unknown",
  [string]$BuildTime = (Get-Date).ToUniversalTime().ToString("yyyy-MM-ddTHH:mm:ssZ"),
  [switch]$SkipTests
)

$ErrorActionPreference = "Stop"

if (-not (Test-Path $DistRoot)) {
  throw "Dist root not found: $DistRoot"
}

if (-not $SkipTests) {
  Write-Host "[test] go test ./..." -ForegroundColor Cyan
  go test -buildvcs=false ./...
}

$binDir = Join-Path $DistRoot "bin"
$templateDir = Join-Path $DistRoot "template_repo"

if (-not (Test-Path $binDir)) {
  New-Item -ItemType Directory -Path $binDir | Out-Null
}
if (-not (Test-Path $templateDir)) {
  throw "Template repo not found: $templateDir"
}

$getExe = Join-Path $binDir "rza_get_object_db.exe"
$migExe = Join-Path $binDir "rza_migrate_db.exe"
$checkExe = Join-Path $binDir "rza_migrate_check.exe"
$restExe = Join-Path $binDir "rza_migrate_rest.exe"
$ldflags = "-X rza_migrate_db/internal/buildinfo.Version=$BuildVersion -X rza_migrate_db/internal/buildinfo.Commit=$BuildCommit -X rza_migrate_db/internal/buildinfo.BuildTime=$BuildTime"

Write-Host "[build] rza_get_object_db.exe" -ForegroundColor Cyan
go build -buildvcs=false -ldflags $ldflags -o $getExe ./cmd/rza_get_object_db

Write-Host "[build] rza_migrate_db.exe" -ForegroundColor Cyan
go build -buildvcs=false -ldflags $ldflags -o $migExe ./cmd/rza_migrate_db

Write-Host "[build] rza_migrate_check.exe" -ForegroundColor Cyan
go build -buildvcs=false -ldflags $ldflags -o $checkExe ./cmd/rza_migrate_check

Write-Host "[build] rza_migrate_rest.exe" -ForegroundColor Cyan
go build -buildvcs=false -ldflags $ldflags -o $restExe ./cmd/rza_migrate_rest

Write-Host "[sync] copy binaries to template_repo" -ForegroundColor Cyan
Copy-Item -Force $getExe (Join-Path $templateDir "rza_get_object_db.exe")
Copy-Item -Force $migExe (Join-Path $templateDir "rza_migrate_db.exe")
Copy-Item -Force $checkExe (Join-Path $templateDir "rza_migrate_check.exe")
Copy-Item -Force $restExe (Join-Path $templateDir "rza_migrate_rest.exe")

Write-Host "[sync] copy binaries to dist root" -ForegroundColor Cyan
Copy-Item -Force $getExe (Join-Path $DistRoot "rza_get_object_db.exe")
Copy-Item -Force $migExe (Join-Path $DistRoot "rza_migrate_db.exe")
Copy-Item -Force $checkExe (Join-Path $DistRoot "rza_migrate_check.exe")
Copy-Item -Force $restExe (Join-Path $DistRoot "rza_migrate_rest.exe")

Write-Host "[sync] copy rza_migrate_rest source to template_repo root" -ForegroundColor Cyan

# remove old nested source bundle if still present
$legacyRestSourceDir = Join-Path $templateDir "rza_migrate_rest_src"
if (Test-Path $legacyRestSourceDir) {
  Remove-Item -Recurse $legacyRestSourceDir
}

New-Item -ItemType Directory -Force -Path (Join-Path $templateDir "cmd") | Out-Null
New-Item -ItemType Directory -Force -Path (Join-Path $templateDir "internal/handler") | Out-Null
New-Item -ItemType Directory -Force -Path (Join-Path $templateDir "internal/migrest") | Out-Null
New-Item -ItemType Directory -Force -Path (Join-Path $templateDir "internal/requestmeta") | Out-Null
New-Item -ItemType Directory -Force -Path (Join-Path $templateDir "internal/telemetry") | Out-Null

Get-ChildItem "internal/handler" -File -Filter *.go |
  Where-Object { $_.Name -notlike '*_test.go' } |
  ForEach-Object {
    Copy-Item -Force $_.FullName (Join-Path $templateDir ("internal/handler/" + $_.Name))
  }
Get-ChildItem "internal/handler" -File -Filter *.yaml |
  ForEach-Object {
    Copy-Item -Force $_.FullName (Join-Path $templateDir ("internal/handler/" + $_.Name))
  }
# cleanup legacy spec filename
$legacyOpenAPISpec = Join-Path $templateDir "internal/handler/openapi.yaml"
if (Test-Path $legacyOpenAPISpec) {
  Remove-Item -Force $legacyOpenAPISpec
}
# cleanup legacy folders not needed by rest deployment bundle
$legacyRestConfigDir = Join-Path $templateDir "internal/config"
if (Test-Path $legacyRestConfigDir) {
  Remove-Item -Recurse -Force $legacyRestConfigDir
}
$legacyRestHistoryDir = Join-Path $templateDir "internal/history"
if (Test-Path $legacyRestHistoryDir) {
  Remove-Item -Recurse -Force $legacyRestHistoryDir
}

Get-ChildItem "internal/migrest" -File -Filter *.go |
  Where-Object { $_.Name -notlike '*_test.go' } |
  ForEach-Object {
    Copy-Item -Force $_.FullName (Join-Path $templateDir ("internal/migrest/" + $_.Name))
  }

Get-ChildItem "internal/requestmeta" -File -Filter *.go |
  Where-Object { $_.Name -notlike '*_test.go' } |
  ForEach-Object {
    Copy-Item -Force $_.FullName (Join-Path $templateDir ("internal/requestmeta/" + $_.Name))
  }

Get-ChildItem "internal/telemetry" -File -Filter *.go |
  Where-Object { $_.Name -notlike '*_test.go' } |
  ForEach-Object {
    Copy-Item -Force $_.FullName (Join-Path $templateDir ("internal/telemetry/" + $_.Name))
  }

# ensure wrapper/source templates for root-based REST module exist
if (-not (Test-Path (Join-Path $templateDir "main.go"))) {
  throw "template_repo/main.go not found. Keep source entrypoint in template_repo root."
}
if (-not (Test-Path (Join-Path $templateDir "cmd/rest.go"))) {
  throw "template_repo/cmd/rest.go not found. Keep Cobra command entrypoint in cmd/rest.go."
}
if (-not (Test-Path (Join-Path $templateDir "go.mod"))) {
  throw "template_repo/go.mod not found."
}
if (-not (Test-Path (Join-Path $templateDir "go.sum"))) {
  throw "template_repo/go.sum not found."
}
if (-not (Test-Path (Join-Path $templateDir "Dockerfile"))) {
  throw "template_repo/Dockerfile not found."
}
if (-not (Test-Path (Join-Path $templateDir ".dockerignore"))) {
  throw "template_repo/.dockerignore not found."
}
if (-not (Test-Path (Join-Path $templateDir ".gitlab-ci.rest.example.yml"))) {
  throw "template_repo/.gitlab-ci.rest.example.yml not found."
}

Write-Host "[done] distribution binaries are up to date" -ForegroundColor Green
