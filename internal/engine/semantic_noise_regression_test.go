package engine

import "testing"

func TestSemanticNoiseRegressionMatrix(t *testing.T) {
	t.Parallel()

	type tc struct {
		name     string
		category string
		a        string
		b        string
	}

	cases := []tc{
		{
			name:     "procedure comment only",
			category: "procedures",
			a:        "CREATE OR ALTER PROCEDURE dbo.ProcA AS BEGIN\nSELECT 1;\nEND\n",
			b:        "/* banner */\nCREATE OR ALTER PROCEDURE dbo.ProcA AS BEGIN\n-- inline\nSELECT 1;\nEND\n",
		},
		{
			name:     "procedure keyword case only",
			category: "procedures",
			a:        "create or alter procedure dbo.ProcA as begin select 1; end",
			b:        "CREATE OR ALTER PROCEDURE dbo.ProcA AS BEGIN SELECT 1; END",
		},
		{
			name:     "function whitespace churn",
			category: "functions",
			a:        "CREATE OR ALTER FUNCTION dbo.FuncA() RETURNS INT AS BEGIN RETURN 1; END",
			b:        "CREATE OR ALTER FUNCTION dbo.FuncA()\nRETURNS INT\nAS\nBEGIN\n    RETURN 1;\nEND\n",
		},
		{
			name:     "view line ending churn",
			category: "views",
			a:        "CREATE OR ALTER VIEW dbo.ViewA AS SELECT 1 AS id;\n",
			b:        "CREATE OR ALTER VIEW dbo.ViewA AS\r\nSELECT 1 AS id;\r\n",
		},
		{
			name:     "procedure semicolon style only",
			category: "procedures",
			a:        "CREATE OR ALTER PROCEDURE dbo.ProcA AS BEGIN SELECT 1 END",
			b:        "CREATE OR ALTER PROCEDURE dbo.ProcA AS BEGIN SELECT 1; END;",
		},
		{
			name:     "view bracket style only",
			category: "views",
			a:        "CREATE OR ALTER VIEW dbo.ViewA AS SELECT id FROM dbo.TableA;",
			b:        "CREATE OR ALTER VIEW [dbo].[ViewA] AS SELECT [id] FROM [dbo].[TableA];",
		},
		{
			name:     "table leading vs trailing comma style",
			category: "tables",
			a:        "CREATE TABLE dbo.TableA (\n  id INT NOT NULL,\n  status VARCHAR(10) NULL\n);",
			b:        "CREATE TABLE dbo.TableA (\n  id INT NOT NULL\n  ,status VARCHAR(10) NULL\n);",
		},
		{
			name:     "table terminal semicolon optional",
			category: "tables",
			a:        "CREATE TABLE dbo.TableA (\n  id INT NOT NULL,\n  status VARCHAR(10) NULL\n);",
			b:        "CREATE TABLE dbo.TableA (\n  id INT NOT NULL,\n  status VARCHAR(10) NULL\n)",
		},
		{
			name:     "table trailing spaces and blank lines",
			category: "tables",
			a:        "CREATE TABLE dbo.TableA (\n  id INT NOT NULL,\n  status VARCHAR(10) NULL\n);",
			b:        "CREATE TABLE dbo.TableA (\n  id INT NOT NULL,   \n\n  status VARCHAR(10) NULL   \n);\n\n",
		},
		{
			name:     "table default style inline vs table-level",
			category: "tables",
			a: "CREATE TABLE dbo.TableA (\n" +
				"  id INT NOT NULL,\n" +
				"  status VARCHAR(10) NULL CONSTRAINT [DF_TableA__status] DEFAULT ('NEW')\n" +
				");",
			b: "CREATE TABLE dbo.TableA (\n" +
				"  id INT NOT NULL,\n" +
				"  status VARCHAR(10) NULL,\n" +
				"  CONSTRAINT [DF_TableA__status] DEFAULT ('NEW') FOR [status]\n" +
				");",
		},
		{
			name:     "indexes order only",
			category: "indexes",
			a:        "CREATE INDEX IX_TableA_B ON dbo.TableA (ColB);\nCREATE INDEX IX_TableA_A ON dbo.TableA (ColA);\n",
			b:        "CREATE INDEX IX_TableA_A ON dbo.TableA (ColA);\nCREATE INDEX IX_TableA_B ON dbo.TableA (ColB);\n",
		},
		{
			name:     "indexes terminal semicolon optional",
			category: "indexes",
			a:        "CREATE INDEX IX_TableA_A ON dbo.TableA (ColA);\n",
			b:        "CREATE INDEX IX_TableA_A ON dbo.TableA (ColA)\n",
		},
		{
			name:     "indexes spacing churn",
			category: "indexes",
			a:        "CREATE INDEX IX_TableA_A ON dbo.TableA (ColA);",
			b:        "CREATE    INDEX   IX_TableA_A   ON    dbo.TableA (ColA) ;",
		},
	}

	for _, tc := range cases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			h1 := hashSemanticByCategory(tc.category, tc.a)
			h2 := hashSemanticByCategory(tc.category, tc.b)
			if h1 != h2 {
				t.Fatalf("semantic hash mismatch (%s):\nA=%s\nB=%s\nh1=%s\nh2=%s", tc.category, tc.a, tc.b, h1, h2)
			}
		})
	}
}
