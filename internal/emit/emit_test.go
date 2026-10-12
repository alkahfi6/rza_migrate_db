package emit

import "testing"

func TestMigrationFileNames(t *testing.T) {
	t.Parallel()

	up, down := MigrationFileNames("20260217073015000", PlannedMigrationRef{
		Category: "tables",
		Object:   "dbo.TableA",
		VK:       2,
	})
	if up != "20260217073015000_tables_dbo.TableA_v2.up.sql" {
		t.Fatalf("unexpected up filename: %s", up)
	}
	if down != "20260217073015000_tables_dbo.TableA_v2.down.sql" {
		t.Fatalf("unexpected down filename: %s", down)
	}

	up, down = MigrationFileNames("20260217073015010", PlannedMigrationRef{
		Category: "patch",
		Object:   "fix_status",
		VK:       1,
	})
	if up != "20260217073015010_patch_fix_status_v1.up.sql" {
		t.Fatalf("unexpected patch up filename: %s", up)
	}
	if down != "20260217073015010_patch_fix_status_v1.down.sql" {
		t.Fatalf("unexpected patch down filename: %s", down)
	}
}

func TestAllocateVersion(t *testing.T) {
	t.Parallel()

	reserved := map[string]struct{}{
		"20260217073015000_tables_dbo.TableA_v1.up.sql":   {},
		"20260217073015000_tables_dbo.TableA_v1.down.sql": {},
	}

	version, nextSeq, err := AllocateVersion("20260217073015", 0, 3, 999, reserved, PlannedMigrationRef{
		Category: "tables",
		Object:   "dbo.TableA",
		VK:       1,
	})
	if err != nil {
		t.Fatalf("allocate version failed: %v", err)
	}
	if version != "20260217073015001" {
		t.Fatalf("unexpected allocated version: %s", version)
	}
	if nextSeq != 2 {
		t.Fatalf("unexpected next sequence: %d", nextSeq)
	}
}

func TestAllocateVersionSkipsVersionOwnedByDifferentObject(t *testing.T) {
	reserved := map[string]struct{}{
		"20260217073015000_tables_dbo.Other_v1.up.sql":   {},
		"20260217073015000_tables_dbo.Other_v1.down.sql": {},
	}
	version, nextSeq, err := AllocateVersion("20260217073015", 0, 3, 999, reserved, PlannedMigrationRef{
		Category: "indexes",
		Object:   "dbo.Target",
		VK:       1,
	})
	if err != nil {
		t.Fatal(err)
	}
	if version != "20260217073015001" || nextSeq != 2 {
		t.Fatalf("global version collision not skipped: version=%s next=%d", version, nextSeq)
	}
}
