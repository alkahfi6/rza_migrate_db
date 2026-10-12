package management

import (
	"os"
	"path/filepath"
	"testing"
)

func TestWriteLoadAndUnmanagedSet(t *testing.T) {
	path := filepath.Join(t.TempDir(), "out", "object_management_manifest.json")
	want := Manifest{
		Schema: "dbo",
		Managed: []Object{
			{Category: "views", Object: "dbo.Z"},
			{Category: "tables", Object: "dbo.A"},
		},
		Unmanaged: []Object{{Category: " Tables ", Object: "dbo.Legacy", Reason: "unsupported"}},
	}
	if err := Write(path, want); err != nil {
		t.Fatal(err)
	}
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if got.GeneratedAt == "" || len(got.Managed) != 2 || got.Managed[0].Category != "tables" {
		t.Fatalf("manifest not normalized/sorted: %#v", got)
	}
	if _, ok := got.UnmanagedSet()["tables|dbo.Legacy"]; !ok {
		t.Fatalf("unmanaged set missing normalized key: %#v", got.UnmanagedSet())
	}
}

func TestLoadFallsBackToBackup(t *testing.T) {
	path := filepath.Join(t.TempDir(), "manifest.json")
	if err := os.WriteFile(path+".bak", []byte(`{"schema":"dbo","unmanaged_objects":[{"category":"views","object":"dbo.Legacy"}]}`), 0o600); err != nil {
		t.Fatal(err)
	}
	got, err := Load(path)
	if err != nil {
		t.Fatal(err)
	}
	if _, ok := got.UnmanagedSet()["views|dbo.Legacy"]; !ok {
		t.Fatalf("backup manifest was not recovered: %#v", got)
	}
}
