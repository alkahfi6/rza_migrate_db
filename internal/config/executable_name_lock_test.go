package config

import "testing"

func TestValidateExecutableNamePasses(t *testing.T) {
	t.Parallel()

	lock := DefaultMigrateExecutableNameLock()
	for _, current := range []string{
		"rza_migrate_db",
		"rza_migrate_db.exe",
		`C:\tools\rza_migrate_db.exe`,
	} {
		if err := ValidateExecutableName(current, lock); err != nil {
			t.Fatalf("expected executable name %q to pass lock: %v", current, err)
		}
	}
}

func TestValidateExecutableNameFailsOnRename(t *testing.T) {
	t.Parallel()

	lock := DefaultGetObjectExecutableNameLock()
	if err := ValidateExecutableName("renamed_tool.exe", lock); err == nil {
		t.Fatalf("expected renamed executable to fail lock")
	}
}

func TestValidateExecutableNamePassesForRestBinary(t *testing.T) {
	t.Parallel()

	lock := DefaultMigrateRestExecutableNameLock()
	for _, current := range []string{
		"rza_migrate_rest",
		"rza_migrate_rest.exe",
		`C:\tools\rza_migrate_rest.exe`,
	} {
		if err := ValidateExecutableName(current, lock); err != nil {
			t.Fatalf("expected executable name %q to pass lock: %v", current, err)
		}
	}
}

func TestValidateExecutableNamePassesForMigrateCheckBinary(t *testing.T) {
	t.Parallel()
	lock := DefaultMigrateCheckExecutableNameLock()
	for _, current := range []string{"rza_migrate_check", "rza_migrate_check.exe", `C:\tools\rza_migrate_check.exe`} {
		if err := ValidateExecutableName(current, lock); err != nil {
			t.Fatalf("expected executable name %q to pass lock: %v", current, err)
		}
	}
}

func TestValidateExecutableNamePassesForMigrateReportBinary(t *testing.T) {
	t.Parallel()
	lock := DefaultMigrateReportExecutableNameLock()
	for _, current := range []string{"rza_migrate_report", "rza_migrate_report.exe", `C:\tools\rza_migrate_report.exe`} {
		if err := ValidateExecutableName(current, lock); err != nil {
			t.Fatalf("expected executable name %q to pass lock: %v", current, err)
		}
	}
}
