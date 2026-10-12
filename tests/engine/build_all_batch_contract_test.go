package engine_test

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestBuildAllBatchContract(t *testing.T) {
	t.Parallel()
	repo := repoRoot(t)
	raw, err := os.ReadFile(filepath.Join(repo, "build_all.bat"))
	if err != nil {
		t.Fatalf("read build_all.bat: %v", err)
	}
	batch := string(raw)

	required := []string{
		"GOOS=windows",
		"GOOS=linux",
		"GOARCH=amd64",
		"CGO_ENABLED=0",
		"-buildvcs=false",
		"-trimpath",
		".\\cmd\\rza_migrate_db",
		".\\cmd\\rza_migrate_check",
		"tools\\linux-amd64\\rza_migrate_db",
		"tools\\linux-amd64\\rza_migrate_check",
		"license_gen",
		"-license-owner",
		"-license-months",
		"SourceSHA256=%RZA_SOURCE_SHA256%",
		"release_artifacts.ps1\" -Mode Prepare",
		"release_artifacts.ps1\" -Mode Finalize",
		"release_artifacts.ps1\" -Mode Publish",
	}
	for _, value := range required {
		if !strings.Contains(batch, value) {
			t.Errorf("build_all.bat is missing %q", value)
		}
	}

	for _, forbidden := range []string{"mssql_credential.nfm", "password", "SQL_REPLICATE", "SQL_REKSA2", "Get-FileHash", "copy /b /y"} {
		if strings.Contains(batch, forbidden) {
			t.Errorf("build_all.bat must not depend on credential/catalog value %q", forbidden)
		}
	}

	buildDB := strings.Index(batch, "-o \"%RZA_STAGE_ROOT%\\tools\\rza_migrate_db.exe\"")
	publishDB := strings.Index(batch, "release_artifacts.ps1\" -Mode Publish")
	if buildDB < 0 || publishDB < 0 || buildDB >= publishDB {
		t.Fatalf("Windows binary must be built in staging before publication")
	}
	licenseGen := strings.Index(batch, "license_gen")
	publishLicense := publishDB
	if licenseGen < 0 || publishLicense < 0 || licenseGen >= publishLicense {
		t.Fatalf("license must be generated in staging before publication")
	}

	helperRaw, err := os.ReadFile(filepath.Join(repo, "scripts", "release_artifacts.ps1"))
	if err != nil {
		t.Fatalf("read release helper: %v", err)
	}
	helper := string(helperRaw)
	for _, required := range []string{
		`$ErrorActionPreference = "Stop"`,
		"Security.Cryptography.SHA256",
		"^[0-9a-f]{64}$",
		"staged checksum mismatch",
		".rza-backup-",
		"unsafe release path",
		"scripts\\verify_release_source.sh",
	} {
		if !strings.Contains(helper, required) {
			t.Errorf("release helper is missing %q", required)
		}
	}
}
