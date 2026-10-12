package cmd

import "testing"

func TestRunRESTFailsWhenExecutableRenamed(t *testing.T) {
	t.Parallel()

	exit := RunREST([]string{
		"renamed_rest_tool.exe",
	})
	if exit != 1 {
		t.Fatalf("expected exit code 1 for renamed executable, got %d", exit)
	}
}
