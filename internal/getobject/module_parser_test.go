package getobject

import (
	"strings"
	"testing"
)

func TestEnsureCreateOrAlterVariants(t *testing.T) {
	t.Parallel()

	cases := []struct {
		name         string
		kind         string
		input        string
		wantPrefix   string
		wantContains string
		expectError  bool
	}{
		{
			name:       "procedure alias proc",
			kind:       "PROCEDURE",
			input:      "CREATE PROC dbo.ProcA AS\nSELECT 1;\n",
			wantPrefix: "CREATE OR ALTER PROCEDURE",
		},
		{
			name:       "procedure with leading comments and sets",
			kind:       "PROCEDURE",
			input:      "-- header\n/* another */\nSET ANSI_NULLS ON\nSET QUOTED_IDENTIFIER ON\nCREATE  PROCEDURE dbo.ProcB AS\nSELECT 1;\n",
			wantPrefix: "CREATE OR ALTER PROCEDURE",
		},
		{
			name:       "function with bom and go",
			kind:       "FUNCTION",
			input:      "\uFEFF\nGO\nCREATE     FUNCTION dbo.FuncA()\nRETURNS INT\nAS\nBEGIN\n  RETURN 1;\nEND\n",
			wantPrefix: "CREATE OR ALTER FUNCTION",
		},
		{
			name:       "view alter normalized",
			kind:       "VIEW",
			input:      "ALTER VIEW dbo.ViewA AS SELECT 1 AS value;\n",
			wantPrefix: "CREATE OR ALTER VIEW",
		},
		{
			name:         "view create or alter without schema",
			kind:         "VIEW",
			input:        "CREATE OR ALTER VIEW ViewNoSchema AS SELECT 1 AS value;\n",
			wantPrefix:   "CREATE OR ALTER VIEW",
			wantContains: "CREATE OR ALTER VIEW dbo.ViewNoSchema AS",
		},
		{
			name:         "view bracket object normalized",
			kind:         "VIEW",
			input:        "CREATE VIEW [dbo].[ViewBracket] AS SELECT 1 AS value;\n",
			wantPrefix:   "CREATE OR ALTER VIEW",
			wantContains: "CREATE OR ALTER VIEW dbo.ViewBracket AS",
		},
		{
			name:         "view mixed bracket object normalized",
			kind:         "VIEW",
			input:        "CREATE VIEW dbo.[ViewMixedBracket] AS SELECT 1 AS value;\n",
			wantPrefix:   "CREATE OR ALTER VIEW",
			wantContains: "CREATE OR ALTER VIEW dbo.ViewMixedBracket AS",
		},
		{
			name:         "procedure without schema normalized",
			kind:         "PROCEDURE",
			input:        "CREATE PROC ProcNoSchema AS\nSELECT 1;\n",
			wantPrefix:   "CREATE OR ALTER PROCEDURE",
			wantContains: "CREATE OR ALTER PROCEDURE dbo.ProcNoSchema AS",
		},
		{
			name:         "function mixed bracket normalized",
			kind:         "FUNCTION",
			input:        "CREATE FUNCTION dbo.[FuncMixed] ()\nRETURNS INT\nAS\nBEGIN\n  RETURN 1;\nEND\n",
			wantPrefix:   "CREATE OR ALTER FUNCTION",
			wantContains: "CREATE OR ALTER FUNCTION dbo.FuncMixed ()",
		},
		{
			name:        "invalid non ddl text",
			kind:        "FUNCTION",
			input:       "definition\nCREATE FUNCTION dbo.FuncB() RETURNS INT AS BEGIN RETURN 1; END",
			expectError: true,
		},
		{
			name:        "invalid non dbo schema",
			kind:        "PROCEDURE",
			input:       "CREATE PROCEDURE sales.ProcInvalid AS SELECT 1;\n",
			expectError: true,
		},
		{
			name:        "invalid select text",
			kind:        "VIEW",
			input:       "SELECT * FROM sys.sql_modules;",
			expectError: true,
		},
		{
			name:        "unsupported kind",
			kind:        "TRIGGER",
			input:       "CREATE TRIGGER dbo.T ON dbo.TableA AFTER INSERT AS SELECT 1;",
			expectError: true,
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			out, err := ensureCreateOrAlter(tc.input, tc.kind)
			if tc.expectError {
				if err == nil {
					t.Fatalf("expected error, got output:\n%s", out)
				}
				return
			}
			if err != nil {
				t.Fatalf("ensureCreateOrAlter failed: %v", err)
			}
			if !strings.HasPrefix(out, tc.wantPrefix+" ") && out != tc.wantPrefix {
				t.Fatalf("unexpected prefix, want %q got:\n%s", tc.wantPrefix, out)
			}
			if tc.wantContains != "" && !strings.Contains(out, tc.wantContains) {
				t.Fatalf("expected output to contain %q, got:\n%s", tc.wantContains, out)
			}
		})
	}
}
