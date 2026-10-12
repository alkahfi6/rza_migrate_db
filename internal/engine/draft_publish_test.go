package engine

import (
	"os"
	"path/filepath"
	"testing"
)

func TestPublishDraftFilesPreservesReviewedBytesAndSources(t *testing.T) {
	root := t.TempDir()
	draft := filepath.Join(root, "draft")
	target := filepath.Join(root, "migrations")
	if err := os.MkdirAll(draft, 0o755); err != nil {
		t.Fatal(err)
	}
	source := filepath.Join(draft, "20260101000000000_tables_dbo.A_v1.up.sql")
	want := []byte("SELECT 1;\r\n-- reviewed bytes\r\n")
	if err := os.WriteFile(source, want, 0o600); err != nil {
		t.Fatal(err)
	}
	published, err := publishDraftFiles([]string{source}, target)
	if err != nil || len(published) != 1 {
		t.Fatalf("publish: files=%#v err=%v", published, err)
	}
	got, err := os.ReadFile(published[0])
	if err != nil || string(got) != string(want) {
		t.Fatalf("published bytes changed: %q err=%v", got, err)
	}
	if _, err := os.Stat(source); err != nil {
		t.Fatalf("source must remain until all files are published: %v", err)
	}
}

func TestPublishDraftFilesNeverOverwritesCollision(t *testing.T) {
	root := t.TempDir()
	draft := filepath.Join(root, "draft")
	target := filepath.Join(root, "migrations")
	if err := os.MkdirAll(draft, 0o755); err != nil {
		t.Fatal(err)
	}
	if err := os.MkdirAll(target, 0o755); err != nil {
		t.Fatal(err)
	}
	name := "20260101000000000_tables_dbo.A_v1.up.sql"
	source := filepath.Join(draft, name)
	destination := filepath.Join(target, name)
	if err := os.WriteFile(source, []byte("new"), 0o600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(destination, []byte("existing"), 0o600); err != nil {
		t.Fatal(err)
	}
	if _, err := publishDraftFiles([]string{source}, target); err == nil {
		t.Fatal("expected collision rejection")
	}
	got, _ := os.ReadFile(destination)
	if string(got) != "existing" {
		t.Fatalf("existing migration was overwritten: %q", got)
	}
	if _, err := os.Stat(source); err != nil {
		t.Fatalf("draft was lost on collision: %v", err)
	}
}
