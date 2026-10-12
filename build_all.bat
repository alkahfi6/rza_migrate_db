@echo off
setlocal EnableExtensions DisableDelayedExpansion

rem Build the complete Windows/Linux release and generate its license.
rem Usage: build_all.bat [license_owner] [license_months] [output_directory]

set "RZA_REPO_ROOT=%~dp0"
if "%RZA_REPO_ROOT:~-1%"=="\" set "RZA_REPO_ROOT=%RZA_REPO_ROOT:~0,-1%"

if /I "%~1"=="/?" goto :usage
if /I "%~1"=="-h" goto :usage
if /I "%~1"=="--help" goto :usage

set "RZA_LICENSE_OWNER=%~1"
if not defined RZA_LICENSE_OWNER (
    set /p "RZA_LICENSE_OWNER=License owner: "
)
if not defined RZA_LICENSE_OWNER (
    echo [error] License owner must not be empty.
    exit /b 2
)

set "RZA_LICENSE_MONTHS=%~2"
if not defined RZA_LICENSE_MONTHS set "RZA_LICENSE_MONTHS=3"

powershell -NoProfile -NonInteractive -Command "$v=$env:RZA_LICENSE_MONTHS; if ($v -notmatch '^[1-9][0-9]*$') { exit 1 }"
if errorlevel 1 (
    echo [error] license_months must be an integer greater than or equal to 1.
    exit /b 2
)

set "RZA_OUTPUT_ROOT=%~3"
if not defined RZA_OUTPUT_ROOT set "RZA_OUTPUT_ROOT=%RZA_REPO_ROOT%"
for %%I in ("%RZA_OUTPUT_ROOT%") do set "RZA_OUTPUT_ROOT=%%~fI"

where go >nul 2>&1
if errorlevel 1 (
    echo [error] Go was not found in PATH.
    exit /b 1
)
where powershell >nul 2>&1
if errorlevel 1 (
    echo [error] PowerShell was not found in PATH.
    exit /b 1
)

set "RZA_STAGE_ROOT=%RZA_REPO_ROOT%\out\.build_all_%RANDOM%_%RANDOM%"
if exist "%RZA_STAGE_ROOT%" (
    echo [error] Staging path already exists: %RZA_STAGE_ROOT%
    exit /b 1
)

mkdir "%RZA_STAGE_ROOT%\tools\linux-amd64" || goto :failed

pushd "%RZA_REPO_ROOT%" || goto :failed

powershell -NoProfile -NonInteractive -ExecutionPolicy Bypass -File "%RZA_REPO_ROOT%\scripts\release_artifacts.ps1" -Mode Prepare -RepoRoot "%RZA_REPO_ROOT%" -StageRoot "%RZA_STAGE_ROOT%"
if errorlevel 1 goto :build_failed
for /f "usebackq tokens=1,* delims==" %%A in ("%RZA_STAGE_ROOT%\build.env") do set "%%A=%%B"
if not defined RZA_BUILD_VERSION goto :build_failed
if not defined RZA_BUILD_TIME goto :build_failed
if not defined RZA_SOURCE_SHA256 goto :build_failed
set "RZA_RELEASE_LDFLAGS=-s -w -X rza_migrate_db/internal/buildinfo.Version=%RZA_BUILD_VERSION% -X rza_migrate_db/internal/buildinfo.BuildTime=%RZA_BUILD_TIME% -X rza_migrate_db/internal/buildinfo.SourceSHA256=%RZA_SOURCE_SHA256%"

echo [build] Windows amd64 rza_migrate_db.exe
set "GOOS=windows"
set "GOARCH=amd64"
set "CGO_ENABLED=0"
go build -buildvcs=false -trimpath -ldflags="%RZA_RELEASE_LDFLAGS%" -o "%RZA_STAGE_ROOT%\tools\rza_migrate_db.exe" .\cmd\rza_migrate_db
if errorlevel 1 goto :build_failed

echo [build] Windows amd64 rza_migrate_check.exe
go build -buildvcs=false -trimpath -ldflags="%RZA_RELEASE_LDFLAGS%" -o "%RZA_STAGE_ROOT%\tools\rza_migrate_check.exe" .\cmd\rza_migrate_check
if errorlevel 1 goto :build_failed

echo [build] Windows amd64 rza_migrate_report.exe
go build -buildvcs=false -trimpath -ldflags="%RZA_RELEASE_LDFLAGS%" -o "%RZA_STAGE_ROOT%\tools\rza_migrate_report.exe" .\cmd\rza_migrate_report
if errorlevel 1 goto :build_failed

echo [build] Linux amd64 rza_migrate_db
set "GOOS=linux"
set "GOARCH=amd64"
set "CGO_ENABLED=0"
go build -buildvcs=false -trimpath -ldflags="%RZA_RELEASE_LDFLAGS%" -o "%RZA_STAGE_ROOT%\tools\linux-amd64\rza_migrate_db" .\cmd\rza_migrate_db
if errorlevel 1 goto :build_failed

echo [build] Linux amd64 rza_migrate_check
go build -buildvcs=false -trimpath -ldflags="%RZA_RELEASE_LDFLAGS%" -o "%RZA_STAGE_ROOT%\tools\linux-amd64\rza_migrate_check" .\cmd\rza_migrate_check
if errorlevel 1 goto :build_failed

echo [build] Linux amd64 rza_migrate_report
go build -buildvcs=false -trimpath -ldflags="%RZA_RELEASE_LDFLAGS%" -o "%RZA_STAGE_ROOT%\tools\linux-amd64\rza_migrate_report" .\cmd\rza_migrate_report
if errorlevel 1 goto :build_failed

echo [license] Generating license.key for "%RZA_LICENSE_OWNER%" (%RZA_LICENSE_MONTHS% months)
set "GOOS="
set "GOARCH="
set "CGO_ENABLED="
"%RZA_STAGE_ROOT%\tools\rza_migrate_db.exe" license_gen -repo "%RZA_REPO_ROOT%" -license-owner "%RZA_LICENSE_OWNER%" -license-months "%RZA_LICENSE_MONTHS%" -license-out "%RZA_STAGE_ROOT%\tools\license.key" -pause-on-exit false
if errorlevel 1 goto :build_failed

powershell -NoProfile -NonInteractive -ExecutionPolicy Bypass -File "%RZA_REPO_ROOT%\scripts\release_artifacts.ps1" -Mode Finalize -RepoRoot "%RZA_REPO_ROOT%" -StageRoot "%RZA_STAGE_ROOT%"
if errorlevel 1 goto :build_failed

popd

mkdir "%RZA_OUTPUT_ROOT%" 2>nul
mkdir "%RZA_OUTPUT_ROOT%\tools" 2>nul
mkdir "%RZA_OUTPUT_ROOT%\tools\linux-amd64" 2>nul

powershell -NoProfile -NonInteractive -ExecutionPolicy Bypass -File "%RZA_REPO_ROOT%\scripts\release_artifacts.ps1" -Mode Publish -RepoRoot "%RZA_REPO_ROOT%" -StageRoot "%RZA_STAGE_ROOT%" -OutputRoot "%RZA_OUTPUT_ROOT%"
if errorlevel 1 goto :publish_failed

call :cleanup_stage
echo [ok] Complete release written to: %RZA_OUTPUT_ROOT%
echo [ok] Windows: tools\rza_migrate_db.exe, tools\rza_migrate_check.exe, tools\rza_migrate_report.exe
echo [ok] Linux: tools\linux-amd64\rza_migrate_db, tools\linux-amd64\rza_migrate_check, tools\linux-amd64\rza_migrate_report
echo [ok] License: tools\license.key (%RZA_LICENSE_MONTHS% months)
exit /b 0

:build_failed
popd
:failed
echo [error] Build failed. Existing release files were not published.
call :cleanup_stage
exit /b 1

:publish_failed
echo [error] Publishing failed. Existing release files were rolled back.
echo [error] Valid staged files remain at: %RZA_STAGE_ROOT%
exit /b 1

:cleanup_stage
del /q "%RZA_STAGE_ROOT%\tools\rza_migrate_db.exe" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\rza_migrate_db.exe.sha256" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\rza_migrate_check.exe" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\rza_migrate_check.exe.sha256" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\rza_migrate_report.exe" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\rza_migrate_report.exe.sha256" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\license.key" 2>nul
del /q "%RZA_STAGE_ROOT%\build.env" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\linux-amd64\rza_migrate_db" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\linux-amd64\rza_migrate_db.sha256" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\linux-amd64\rza_migrate_check" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\linux-amd64\rza_migrate_check.sha256" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\linux-amd64\rza_migrate_report" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\linux-amd64\rza_migrate_report.sha256" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\linux-amd64\build-info.txt" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\linux-amd64\release-manifest.json" 2>nul
del /q "%RZA_STAGE_ROOT%\tools\linux-amd64\release.sha256" 2>nul
rmdir "%RZA_STAGE_ROOT%\tools\linux-amd64" 2>nul
rmdir "%RZA_STAGE_ROOT%\tools" 2>nul
rmdir "%RZA_STAGE_ROOT%" 2>nul
exit /b 0

:usage
echo Usage: build_all.bat [license_owner] [license_months] [output_directory]
echo.
echo Builds Windows amd64 and Linux amd64 binaries for rza_migrate_db and
echo rza_migrate_check, creates SHA-256 manifests, and generates license.key.
echo license_months defaults to 3. output_directory defaults to this repository.
exit /b 0
