@echo off
setlocal EnableExtensions DisableDelayedExpansion

rem Compatibility wrapper. A checker-only rebuild can leave the MR tool out of
rem sync with the engine and release metadata, so all artifacts are now built
rem and published as one release by build_all.bat.
rem Usage: scripts\build_migrate_check_linux.bat [license_owner] [license_months] [output_directory]

set "RZA_LICENSE_OWNER=%~1"
if not defined RZA_LICENSE_OWNER set "RZA_LICENSE_OWNER=template_repo"
set "RZA_LICENSE_MONTHS=%~2"
if not defined RZA_LICENSE_MONTHS set "RZA_LICENSE_MONTHS=3"

call "%~dp0..\build_all.bat" "%RZA_LICENSE_OWNER%" "%RZA_LICENSE_MONTHS%" "%~3"
exit /b %errorlevel%
