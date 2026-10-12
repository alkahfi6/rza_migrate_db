package migreport

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestProcedureScannerQA31Golden(t *testing.T) {
	root := filepath.Join("..", "..", "fixtures", "qa_rules", "QA-31-desired-state-documentation")
	var expected struct {
		ProcedureScan struct {
			Name               string   `json:"name"`
			Accesses           []string `json:"accesses"`
			Calls              []string `json:"calls"`
			MappingExpressions []string `json:"mapping_expressions"`
			HasClientResult    bool     `json:"has_client_result"`
		} `json:"procedure_scan"`
	}
	raw, err := os.ReadFile(filepath.Join(root, "expected", "output.golden.json"))
	if err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(raw, &expected); err != nil {
		t.Fatal(err)
	}
	docs, err := BuildDesiredStateDocs(filepath.Join(root, "input"))
	if err != nil {
		t.Fatal(err)
	}
	for _, p := range docs.ProcedureProfiles {
		if p.Name != expected.ProcedureScan.Name {
			continue
		}
		var accesses, calls, expressions []string
		for _, a := range p.Accesses {
			accesses = append(accesses, a.Operation+"|"+a.Object+"|"+a.Resolution)
		}
		for _, c := range p.Calls {
			calls = append(calls, c.Object)
		}
		for _, m := range p.Mappings {
			expressions = append(expressions, m.Expression)
		}
		if !reflect.DeepEqual(accesses, expected.ProcedureScan.Accesses) || !reflect.DeepEqual(calls, expected.ProcedureScan.Calls) || !reflect.DeepEqual(expressions, expected.ProcedureScan.MappingExpressions) || p.HasDirectResult != expected.ProcedureScan.HasClientResult {
			t.Fatalf("golden mismatch: accesses=%v calls=%v mappings=%v result=%v", accesses, calls, expressions, p.HasDirectResult)
		}
		return
	}
	t.Fatal("golden procedure missing")
}

func TestProcedurePrecisionKeepsEveryStatementAndMapping(t *testing.T) {
	catalog := newProcedureCatalog(map[string]*DocumentedTable{"dbo.T": {Name: "dbo.T"}}, nil, nil, map[string]struct{}{"dbo.Child": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE PROCEDURE dbo.P AS
UPDATE dbo.T SET Id=1 WHERE Id=0; UPDATE dbo.T SET Id=2 WHERE Id=1;
INSERT dbo.T(Id) VALUES(3); INSERT dbo.T(Id) VALUES(4);
EXEC
dbo.Child
@id=1; EXEC dbo.Child @id=2;`}, catalog)
	if len(p.Accesses) != 4 || len(p.Mappings) != 4 || len(p.Calls) != 2 {
		t.Fatalf("lost occurrences: %+v", p)
	}
	for i, want := range []string{"1", "2", "3", "4"} {
		if p.Mappings[i].Expression != want {
			t.Fatalf("mapping crosses statement: %+v", p.Mappings)
		}
	}
	if p.Calls[0].Arguments != "@id=1" || p.Calls[1].Arguments != "@id=2" || p.Calls[0].Line != 4 {
		t.Fatalf("static call scope/position lost: %+v", p.Calls)
	}
}

func TestProcedureSameLineSequenceUsesSourceOffset(t *testing.T) {
	p := analyzeProcedure(moduleSource{Kind: "procedure", SQL: `EXEC dbo.Child; UPDATE dbo.T SET Id=1 WHERE Id=0; EXEC dbo.Child;`}, newProcedureCatalog(map[string]*DocumentedTable{"dbo.T": {Name: "dbo.T"}}, nil, nil, map[string]struct{}{"dbo.Child": {}}))
	events := procedureProcessEvents(&p)
	if len(events) != 3 || events[0].event != "EXEC" || events[1].event != "UPDATE" || events[2].event != "EXEC" {
		t.Fatalf("same-line source order lost: %+v", events)
	}
	text := renderObservedProcessSequence(&p)
	if strings.Index(text, "`EXEC`") > strings.Index(text, "`UPDATE`") {
		t.Fatalf("audit sequence order wrong: %s", text)
	}
}

func TestProcedureCallAndMappingStopBeforeBlockEnd(t *testing.T) {
	p := analyzeProcedure(moduleSource{Kind: "procedure", SQL: `CREATE PROCEDURE dbo.P AS BEGIN
UPDATE dbo.T SET Id=CASE WHEN Id=0 THEN 1 ELSE 2 END WHERE Id=0
EXEC dbo.Child @id=1
END;`}, newProcedureCatalog(map[string]*DocumentedTable{"dbo.T": {Name: "dbo.T"}}, nil, nil, map[string]struct{}{"dbo.Child": {}}))
	if len(p.Calls) != 1 || p.Calls[0].Arguments != "@id=1" || len(p.Mappings) != 1 || p.Mappings[0].Expression != "CASE WHEN Id=0 THEN 1 ELSE 2 END" {
		t.Fatalf("block/case boundaries wrong: calls=%+v mappings=%+v", p.Calls, p.Mappings)
	}
}

func TestProcedureAliasIsScopedAndAmbiguityIsNotDiscarded(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.Real": {Name: "dbo.Real"}, "dbo.Other": {Name: "dbo.Other"}, "dbo.t": {Name: "dbo.t"}}
	catalog := newProcedureCatalog(tables, nil, nil, nil)
	p := analyzeProcedure(moduleSource{Kind: "procedure", SQL: `UPDATE t SET Id=1 FROM dbo.Real t WHERE EXISTS(SELECT 1 FROM dbo.Other t);
UPDATE x SET Id=2 FROM dbo.Real x JOIN dbo.Other x ON 1=1;`}, catalog)
	if !hasAccess(p.Accesses, "UPDATE", "dbo.Real", ResolutionStatic) || !hasAccess(p.Accesses, "UPDATE", "x", ResolutionAmbiguous) {
		t.Fatalf("incorrect alias resolution: %+v", p.Accesses)
	}
	if len(p.Mappings) != 1 || p.Mappings[0].Expression != "1" {
		t.Fatalf("FROM leaked into assignment: %+v", p.Mappings)
	}
}

func TestProcedureQuotedObjectIdentityIsPreserved(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.Order Items": {Name: "dbo.Order Items"}, "dbo.a]b": {Name: "dbo.a]b"}, "dbo.A.B": {Name: "dbo.A.B"}}
	tables["dbo.INTO"] = &DocumentedTable{Name: "dbo.INTO"}
	tables["dbo.OPENJSON"] = &DocumentedTable{Name: "dbo.OPENJSON"}
	p := analyzeProcedure(moduleSource{Kind: "procedure", SQL: `SELECT * FROM [dbo].[Order Items]; SELECT * FROM [dbo].[a]]b]; SELECT * FROM dbo.[A.B]; SELECT * FROM [INTO]; SELECT * FROM [OPENJSON];`}, newProcedureCatalog(tables, nil, nil, nil))
	for name := range tables {
		if !hasAccess(p.Accesses, "READ", name, ResolutionStatic) {
			t.Fatalf("identity %s corrupted: %+v", name, p.Accesses)
		}
	}
}

func TestProcedureCTECaseDependsOnUnknownCollation(t *testing.T) {
	p := analyzeProcedure(moduleSource{Kind: "procedure", SQL: `WITH Recent AS(SELECT 1 Id) SELECT * FROM recent;`}, newProcedureCatalog(map[string]*DocumentedTable{"dbo.recent": {Name: "dbo.recent"}}, nil, nil, nil))
	if len(p.Accesses) != 1 || p.Accesses[0].Resolution != ResolutionAmbiguous {
		t.Fatalf("unknown collation guessed physical target: %+v", p.Accesses)
	}
}

func TestProcedureDoesNotClaimExactMappingForMultipleRowsOrQueryBranches(t *testing.T) {
	catalog := newProcedureCatalog(map[string]*DocumentedTable{"dbo.T": {Name: "dbo.T"}}, nil, nil, nil)
	for _, sql := range []string{`INSERT dbo.T(Id) VALUES(1),(2);`, `INSERT dbo.T(Id) SELECT TOP(1) Id FROM dbo.T;`, `INSERT dbo.T(Id) SELECT Id FROM dbo.T UNION ALL SELECT Id FROM dbo.T;`} {
		p := analyzeProcedure(moduleSource{Kind: "procedure", SQL: sql}, catalog)
		if len(p.Mappings) != 0 {
			t.Fatalf("mapping incomplete/guessed for %s: %+v", sql, p.Mappings)
		}
	}
}

func TestProcedureProfileUsesCatalogAndKeepsDynamicEvidenceConservative(t *testing.T) {
	tables := map[string]*DocumentedTable{
		"dbo.Target": {Name: "dbo.Target"},
		"dbo.Log":    {Name: "dbo.Log"},
	}
	views := map[string]struct{}{"dbo.vTarget": {}}
	functions := map[string]struct{}{"dbo.fnKnown": {}}
	procedures := map[string]struct{}{"dbo.Root": {}, "dbo.Child": {}, "dbo.DynamicChild": {}}
	catalog := newProcedureCatalog(tables, views, functions, procedures)

	root := analyzeProcedure(moduleSource{Name: "dbo.Root", Kind: "procedure", SQL: `
CREATE OR ALTER PROCEDURE dbo.Root @id int, @status varchar(10) = 'NEW' AS
BEGIN TRAN;
UPDATE t SET t.Status = @status FROM dbo.Target t WHERE t.Id = @id;
INSERT INTO #work(Id) VALUES (@id);
SELECT dbo.fnKnown(@id);
SELECT GETDATE();
EXEC dbo.Child @id;
EXEC dbo.DynamicChild 'Target';
EXEC SQL_OTHER.dbo.ExternalProc @id;
EXEC(N'UPDATE dbo.Target SET Status = 1');
COMMIT;
RETURN 0;`}, catalog)

	if len(root.Parameters) != 2 || root.Parameters[1].Default != "'NEW'" {
		t.Fatalf("parameter contract mismatch: %+v", root.Parameters)
	}
	if !root.HasTransaction || !root.HasDirectResult {
		t.Fatalf("execution signals missing: %+v", root)
	}
	if !hasAccess(root.Accesses, "UPDATE", "dbo.Target", ResolutionStatic) {
		t.Fatalf("UPDATE alias was not resolved through db_object: %+v", root.Accesses)
	}
	if !hasAccess(root.Accesses, "INSERT", "#work", ResolutionNotPersisted) {
		t.Fatalf("temporary target must stay non-persisted: %+v", root.Accesses)
	}
	if len(root.Functions) != 1 || root.Functions[0].Object != "dbo.fnKnown" {
		t.Fatalf("built-ins must not become user-defined functions: %+v", root.Functions)
	}
	if !hasCall(root.Calls, "dbo.Child", ResolutionStatic) || !hasCall(root.Calls, "SQL_OTHER.dbo.ExternalProc", ResolutionExternal) {
		t.Fatalf("procedure resolution mismatch: %+v", root.Calls)
	}
	if !hasDynamic(root.DynamicSQL, ResolutionDynamic, "dbo.Target") {
		t.Fatalf("literal dynamic SQL was not re-parsed and catalog-validated: %+v", root.DynamicSQL)
	}
}

func TestProcedureDynamicObjectCanBeAttributedOnlyFromLiteralCaller(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.Target": {Name: "dbo.Target"}}
	views, functions := map[string]struct{}{}, map[string]struct{}{}
	procedures := map[string]struct{}{"dbo.Root": {}, "dbo.DynamicChild": {}}
	catalog := newProcedureCatalog(tables, views, functions, procedures)
	modules := []moduleSource{
		{Name: "dbo.Root", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.Root AS EXEC dbo.DynamicChild 'Target';`},
		{Name: "dbo.DynamicChild", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.DynamicChild @pcTableName varchar(100) AS
DECLARE @sql nvarchar(max);
SET @sql = 'UPDATE dbo.' + @pcTableName + ' SET Flag=1';
EXEC(@sql);`},
	}
	profiles := make([]ProcedureProfile, 0, len(modules))
	for _, module := range modules {
		profiles = append(profiles, analyzeProcedure(module, catalog))
	}
	applyCallerResolution(profiles, catalog)
	var child ProcedureProfile
	for _, profile := range profiles {
		if profile.Name == "dbo.DynamicChild" {
			child = profile
		}
	}
	if !hasDynamic(child.DynamicSQL, ResolutionFromCaller, "dbo.Target") {
		t.Fatalf("literal caller evidence did not resolve dynamic object candidate: child=%+v profiles=%+v", child.DynamicSQL, profiles)
	}
	if !strings.Contains(child.DynamicSQL[0].RuntimeSource, "dbo.Root") {
		t.Fatalf("caller evidence missing: %+v", child.DynamicSQL)
	}
}

func TestProcedureProfileDoesNotCallAssignmentOrExistsAResultSet(t *testing.T) {
	catalog := newProcedureCatalog(map[string]*DocumentedTable{"dbo.T": {Name: "dbo.T"}, "dbo.L": {Name: "dbo.L"}}, nil, nil, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P AS
EXEC @result = sp_xml_preparedocument @handle OUTPUT, @xml;
SELECT @x = Value FROM dbo.T;
IF EXISTS (
  SELECT TOP 1 1 FROM dbo.T
)
BEGIN
  INSERT dbo.L(Id)
  SELECT Id FROM dbo.T;
END
RETURN 0;`}, catalog)
	if p.HasDirectResult {
		t.Fatalf("assignment, EXISTS, and INSERT SELECT must not be reported as direct result sets")
	}
	if len(p.DynamicSQL) != 0 {
		t.Fatalf("EXEC return-value assignment to a static system procedure is not dynamic SQL: %+v", p.DynamicSQL)
	}
}

func TestProcedureKeyContextDoesNotClaimSingleRowWithoutFullUniqueKey(t *testing.T) {
	docs := DesiredStateDocs{Tables: []DocumentedTable{{
		Name: "dbo.Order", Keys: []DocumentedKey{{Columns: []string{"OrderId", "TrxId"}, Evidence: "UNIQUE_INDEX"}},
		Indexes: []DocumentedIndex{{Name: "IX_Order_TrxId", Columns: []string{"TrxId"}}},
	}}}
	p := ProcedureProfile{Accesses: []ProcedureDataAccess{{Operation: "UPDATE", Object: "dbo.Order", Resolution: ResolutionStatic}}, Mappings: []ProcedureWriteMapping{{Object: "dbo.Order", Filter: "TrxId=@trx"}}}
	context := renderProcedureKeyContext(docs, &p)
	if !strings.Contains(context, "NOT_PROVEN") || !strings.Contains(context, "UNIQUE_INDEX: (OrderId, TrxId)") {
		t.Fatalf("composite unique key coverage must remain conservative: %s", context)
	}
}

func TestProcedureInsertSelectMappingPreservesTargetAndSourceExpressions(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.Source": {Name: "dbo.Source"}, "dbo.Log": {Name: "dbo.Log"}}
	catalog := newProcedureCatalog(tables, nil, nil, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P @id int AS
INSERT dbo.Log(LogDate, Id)
SELECT GETDATE(), SourceId
FROM dbo.Source
WHERE SourceId=@id;`}, catalog)
	if len(p.Mappings) != 2 || p.Mappings[0].Column != "LogDate" || p.Mappings[0].Expression != "GETDATE()" || p.Mappings[1].Column != "Id" || p.Mappings[1].Expression != "SourceId" {
		t.Fatalf("INSERT SELECT mapping mismatch: %+v", p.Mappings)
	}
}

func TestProcedureCatalogDoesNotGuessAmbiguousUnqualifiedObject(t *testing.T) {
	tables := map[string]*DocumentedTable{
		"dbo.Target":   {Name: "dbo.Target"},
		"audit.Target": {Name: "audit.Target"},
	}
	catalog := newProcedureCatalog(tables, nil, nil, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P AS SELECT * FROM Target;`}, catalog)
	if !hasAccess(p.Accesses, "READ", "Target", ResolutionAmbiguous) {
		t.Fatalf("unqualified object with multiple catalog candidates must stay ambiguous: %+v", p.Accesses)
	}
}

func TestProcedureCatalogDetectsCaseCollision(t *testing.T) {
	procedures := map[string]struct{}{"dbo.P": {}, "dbo.Child": {}, "dbo.child": {}}
	catalog := newProcedureCatalog(nil, nil, nil, procedures)
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P AS EXEC dbo.Child;`}, catalog)
	if !hasCall(p.Calls, "dbo.Child", ResolutionAmbiguous) {
		t.Fatalf("case-insensitive catalog collision must not be selected automatically: %+v", p.Calls)
	}
}

func TestProcedureParameterDefaultContainingASIsPreserved(t *testing.T) {
	catalog := newProcedureCatalog(nil, nil, nil, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P @label varchar(20) = 'AS text', @id int AS SELECT @id;`}, catalog)
	if len(p.Parameters) != 2 || p.Parameters[0].Default != "'AS text'" || p.Parameters[1].Name != "@id" {
		t.Fatalf("AS inside a string literal must not terminate the parameter list: %+v", p.Parameters)
	}
}

func TestProcedureProfileSeparatesTableValuedFunctionAndCursorFromDataObjects(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.Target": {Name: "dbo.Target"}}
	functions := map[string]struct{}{"dbo.Rows_FN": {}}
	catalog := newProcedureCatalog(tables, nil, functions, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P AS
SELECT * FROM dbo.Rows_FN(1);
FETCH NEXT FROM WorkCursor INTO @id;`}, catalog)
	if len(p.Accesses) != 0 {
		t.Fatalf("TVF and cursor must not be reported as table/view access: %+v", p.Accesses)
	}
	if len(p.Functions) != 1 || p.Functions[0].Object != "dbo.Rows_FN" {
		t.Fatalf("TVF must remain visible as a function dependency: %+v", p.Functions)
	}
}

func TestProcedureScannerHandlesMultilineAccessAndStatementScopedCTE(t *testing.T) {
	tables := map[string]*DocumentedTable{
		"dbo.Source": {Name: "dbo.Source"},
		"dbo.T":      {Name: "dbo.T"},
	}
	catalog := newProcedureCatalog(tables, nil, nil, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P AS
WITH T AS (SELECT Id FROM dbo.Source)
SELECT * FROM T;
SELECT *
FROM
dbo.T;`}, catalog)
	if !hasAccess(p.Accesses, "READ", "T", ResolutionNotPersisted) {
		t.Fatalf("CTE reference must remain query-scoped and non-persisted: %+v", p.Accesses)
	}
	if !hasAccess(p.Accesses, "READ", "dbo.T", ResolutionStatic) {
		t.Fatalf("schema-qualified table after CTE was not resolved across line breaks: %+v", p.Accesses)
	}
	if !hasAccess(p.Accesses, "READ", "dbo.Source", ResolutionStatic) {
		t.Fatalf("physical source inside CTE was not retained: %+v", p.Accesses)
	}
}

func TestProcedureSelectIntoIsNotClientResultSet(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.T": {Name: "dbo.T"}}
	catalog := newProcedureCatalog(tables, nil, nil, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P AS
SELECT Id INTO #work FROM dbo.T;`}, catalog)
	if p.HasDirectResult {
		t.Fatal("SELECT INTO populates an object and must not be reported as a client result set")
	}
}

func TestProcedureDeleteAliasResolutionStopsAtNextStatement(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.First": {Name: "dbo.First"}, "dbo.Second": {Name: "dbo.Second"}}
	catalog := newProcedureCatalog(tables, nil, nil, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P AS
DELETE tm
FROM dbo.First tm
WHERE Id = 1
DELETE tm
FROM dbo.Second tm
WHERE Id = 2;`}, catalog)
	if !hasAccess(p.Accesses, "DELETE", "dbo.First", ResolutionStatic) || !hasAccess(p.Accesses, "DELETE", "dbo.Second", ResolutionStatic) {
		t.Fatalf("each DELETE alias must resolve within its own statement: %+v", p.Accesses)
	}
}

func TestProcedureCallArgumentsMapToChildParameters(t *testing.T) {
	docs := DesiredStateDocs{ProcedureProfiles: []ProcedureProfile{
		{Name: "dbo.Root"},
		{Name: "dbo.Child", Parameters: []ProcedureParameter{{Name: "@id"}, {Name: "@status"}}},
	}}
	positional := formatCallArgumentMapping(docs, ProcedureCall{Object: "dbo.Child", Arguments: "@localId, 'OK'", Resolution: ResolutionStatic})
	if positional != "@id ← @localId; @status ← 'OK'" {
		t.Fatalf("positional argument mapping mismatch: %s", positional)
	}
	named := formatCallArgumentMapping(docs, ProcedureCall{Object: "dbo.Child", Arguments: "@status = @localStatus, @id = 7", Resolution: ResolutionStatic})
	if named != "@status ← @localStatus; @id ← 7" {
		t.Fatalf("named argument mapping mismatch: %s", named)
	}
}

func TestProcedureProfileHandlesFormattedInsertCTETempAndSystemObjects(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.Target": {Name: "dbo.Target"}}
	catalog := newProcedureCatalog(tables, nil, nil, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P AS
WITH RowsToDelete AS (SELECT Id FROM dbo.Target)
DELETE FROM RowsToDelete;
INSERT INTO
  dbo.Target(Id)
VALUES (1);
SELECT * FROM dbo.#session;
SELECT * FROM sys.columns;
SELECT * FROM SQL_REPLICATE..RemoteTable;`}, catalog)
	if !hasAccess(p.Accesses, "DELETE", "RowsToDelete", ResolutionNotPersisted) {
		t.Fatalf("CTE must be marked non-persisted: %+v", p.Accesses)
	}
	for _, access := range p.Accesses {
		if access.Operation == "READ" && access.Object == "RowsToDelete" {
			t.Fatalf("DELETE target must not be duplicated as a READ: %+v", p.Accesses)
		}
	}
	if !hasAccess(p.Accesses, "INSERT", "dbo.Target", ResolutionStatic) {
		t.Fatalf("multiline INSERT target was not resolved: %+v", p.Accesses)
	}
	if !hasAccess(p.Accesses, "READ", "dbo.#session", ResolutionNotPersisted) {
		t.Fatalf("schema-qualified temp object must be non-persisted: %+v", p.Accesses)
	}
	if !hasAccess(p.Accesses, "READ", "sys.columns", ResolutionSystem) {
		t.Fatalf("system catalog must be a system object: %+v", p.Accesses)
	}
	if !hasAccess(p.Accesses, "READ", "SQL_REPLICATE..RemoteTable", ResolutionExternal) {
		t.Fatalf("database..object must be external: %+v", p.Accesses)
	}
}

func TestProcedureAliasResolutionSupportsLongSetAndDerivedTarget(t *testing.T) {
	var filler strings.Builder
	for i := 0; i < 120; i++ {
		filler.WriteString(" + 0\n")
	}
	tables := map[string]*DocumentedTable{"dbo.Target": {Name: "dbo.Target"}}
	catalog := newProcedureCatalog(tables, nil, nil, map[string]struct{}{"dbo.P": {}})
	sql := "CREATE OR ALTER PROCEDURE dbo.P AS\nUPDATE t SET Value = 1" + filler.String() + "FROM dbo.Target t WHERE Id=1;\nUPDATE derived SET Value=2 FROM (SELECT Value FROM dbo.Target) derived;"
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: sql}, catalog)
	if !hasAccess(p.Accesses, "UPDATE", "dbo.Target", ResolutionStatic) {
		t.Fatalf("alias after a long SET clause must resolve: %+v", p.Accesses)
	}
	if !hasAccess(p.Accesses, "UPDATE", "derived", ResolutionNotPersisted) {
		t.Fatalf("derived-table update target must be non-persisted: %+v", p.Accesses)
	}
}

func TestProcedureAliasPreservesMissingCatalogTargetName(t *testing.T) {
	catalog := newProcedureCatalog(nil, nil, nil, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P AS
UPDATE missingAlias SET Flag=1
FROM dbo.MissingTable missingAlias
WHERE Id=1;`}, catalog)
	if !hasAccess(p.Accesses, "UPDATE", "dbo.MissingTable", ResolutionNotFound) {
		t.Fatalf("explicit alias must retain its source target even when db_object validation fails: %+v", p.Accesses)
	}
}

func TestProcedureInsertWithoutTargetColumnsDoesNotInventMapping(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.Target": {Name: "dbo.Target"}, "dbo.Source": {Name: "dbo.Source"}}
	catalog := newProcedureCatalog(tables, nil, nil, map[string]struct{}{"dbo.P": {}})
	p := analyzeProcedure(moduleSource{Name: "dbo.P", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.P AS
INSERT dbo.Target
SELECT COALESCE(Id, 0)
FROM dbo.Source;`}, catalog)
	if !hasAccess(p.Accesses, "INSERT", "dbo.Target", ResolutionStatic) {
		t.Fatalf("INSERT access must remain visible: %+v", p.Accesses)
	}
	if len(p.Mappings) != 0 {
		t.Fatalf("mapping must not be invented when target column list is absent: %+v", p.Mappings)
	}
}

func TestProcedureCatalogCountsUnresolvedDynamicSQL(t *testing.T) {
	docs := DesiredStateDocs{ProcedureProfiles: []ProcedureProfile{{
		Name:       "dbo.P",
		DynamicSQL: []ProcedureDynamicSQL{{Resolution: ResolutionUnresolved}},
	}}}
	report := renderProcedureCatalog(docs)
	if !strings.Contains(report, "| 0 | 0 | 0 | 0 | 0 | 1 | 1 |") {
		t.Fatalf("unresolved dynamic SQL must contribute to unresolved reference count: %s", report)
	}
}

func TestProcedureMapsPreserveNestedCallGuardsAndRecursiveDirection(t *testing.T) {
	tables := map[string]*DocumentedTable{"dbo.Target": {Name: "dbo.Target"}}
	procedures := map[string]struct{}{"dbo.Root": {}, "dbo.ChildA": {}, "dbo.ChildB": {}}
	catalog := newProcedureCatalog(tables, nil, nil, procedures)
	root := analyzeProcedure(moduleSource{Name: "dbo.Root", Kind: "procedure", SQL: `CREATE OR ALTER PROCEDURE dbo.Root @status int, @mode varchar(10) AS
BEGIN TRAN
UPDATE dbo.Target SET Flag=1 WHERE Id=1
COMMIT TRAN
IF (@status = 1)
BEGIN
  IF (@mode = 'A')
  BEGIN
    EXEC dbo.ChildA @status
  END
  ELSE BEGIN
    EXEC dbo.ChildB @status
  END
END
RETURN 0;`}, catalog)
	if len(root.Calls) != 2 || !strings.Contains(root.Calls[0].Guard, "@status = 1") || !strings.Contains(root.Calls[0].Guard, "@mode = 'A'") {
		t.Fatalf("nested guard missing from first call: %+v", root.Calls)
	}
	if !strings.Contains(root.Calls[1].Guard, "ELSE of IF (@mode = 'A')") {
		t.Fatalf("ELSE guard missing from second call: %+v", root.Calls)
	}
	localMap := renderProcedureMermaid(&root)
	if !strings.Contains(localMap, "flowchart TD") || !strings.Contains(localMap, "TRANSACTION<br/>") || !strings.Contains(localMap, "PROCEDURE ROUTING") || !strings.Contains(localMap, "condition not met") || strings.Contains(localMap, "root -->|UPDATE|") {
		t.Fatalf("local map must summarize transactions and expose guarded routing instead of a linear/star graph: %s", localMap)
	}
	if !strings.Contains(localMap, "<br/>COMMIT`") || strings.Contains(localMap, "\\n") || strings.Contains(localMap, `-->|"line `) {
		t.Fatalf("local map must show the transaction boundary inside its scope and keep source-line noise out of ordinary edges: %s", localMap)
	}
	if !strings.Contains(localMap, `-->|"`) {
		t.Fatalf("Mermaid edge labels must be quoted so SQL operators remain renderable: %s", localMap)
	}
	objectMap := renderLocalObjectMap(&root)
	if !strings.Contains(objectMap, "dbo.Target") || !strings.Contains(objectMap, "UPDATE") || !strings.Contains(objectMap, "dbo.ChildA") || !strings.Contains(objectMap, "EXEC") {
		t.Fatalf("local object map must include direct data objects and procedure calls: %s", objectMap)
	}
	if strings.Contains(objectMap, "\\n") || !strings.Contains(objectMap, "<br/>PROCEDURE") {
		t.Fatalf("local object map must use Mermaid-compatible line breaks: %s", objectMap)
	}
	docs := DesiredStateDocs{ProcedureProfiles: []ProcedureProfile{root, {Name: "dbo.ChildA"}, {Name: "dbo.ChildB"}}}
	recursiveMap := renderRecursiveProcedureMermaid(docs, "dbo.Root")
	if !strings.Contains(recursiveMap, "p0 -->") || !strings.Contains(recursiveMap, "@mode = 'A'") || !strings.Contains(recursiveMap, "dbo.ChildB") {
		t.Fatalf("recursive procedure map must preserve caller-to-callee edges and guards: %s", recursiveMap)
	}
	if !strings.Contains(recursiveMap, `-->|"`) {
		t.Fatalf("recursive Mermaid labels must be quoted: %s", recursiveMap)
	}
	profileDocument := renderProcedureProfile(docs, &root)
	if !strings.Contains(profileDocument, "### Call 1 — `dbo.ChildA`") || strings.Contains(profileDocument, "| Called procedure |") {
		t.Fatalf("direct calls must use readable per-call blocks instead of a wide table: %s", profileDocument)
	}
	if !strings.Contains(profileDocument, "### Diagram legend") || !strings.Contains(profileDocument, "Cabang dari decision node adalah kemungkinan path") {
		t.Fatalf("profile must explain how to read the process graph without forcing users to infer its semantics: %s", profileDocument)
	}
}

func TestCompactGuardForDiagramKeepsBranchDiscriminators(t *testing.T) {
	guard := "IF(@pnStatus = 1) AND IF(ISNULL(@pcXmlResponse, '') <> '') AND IF(@pcWorkflowID = 'ORD_ORDER_ORI')"
	got := compactGuardForDiagram(guard)
	if strings.Count(got, "\\n") != 2 || !strings.Contains(got, "@pnStatus = 1") || !strings.Contains(got, "ORD_ORDER_ORI") {
		t.Fatalf("diagram guard must remain readable while retaining branch discriminators: %q", got)
	}
	elseGuard := "IF (@pcWorkflowID IN ('ORD_PAYMENT_INQ', 'ORD_PAYMENT_PROCESS')) AND isnull(@pcXmlResponse,'') <> '' AND ELSE of IF @cPaymntMode = 'DIRECT'"
	if got := compactGuardForDiagram(elseGuard); !strings.Contains(got, "ELSE of IF @cPaymntMode = 'DIRECT'") {
		t.Fatalf("ELSE discriminator must survive diagram compaction: %q", got)
	}
}

func TestProcedureFlowSplitCountsSeparatedBranchesAndErrorNode(t *testing.T) {
	events := []procedureSequenceEvent{}
	for index := 0; index < 16; index++ {
		events = append(events, procedureSequenceEvent{line: index + 1, event: "READ", target: fmt.Sprintf("dbo.T%d", index)})
	}
	events = append(events,
		procedureSequenceEvent{line: 20, event: "READ", target: "dbo.A", guard: "IF @a = 1"},
		procedureSequenceEvent{line: 21, event: "GOTO", target: "GOTO ERROR"},
		procedureSequenceEvent{line: 22, event: "READ", target: "dbo.B", guard: "IF @b = 1"},
	)
	chunks := splitProcedureFlowEvents(events, 20)
	for _, chunk := range chunks {
		if nodes := estimatedProcedureFlowNodes(chunk); nodes > 20 {
			t.Fatalf("flow split exceeded the render budget: %d nodes in %+v", nodes, chunk)
		}
	}
}

func TestRecursiveProcedureMapSplitsWithoutDroppingCallEdges(t *testing.T) {
	root := ProcedureProfile{Name: "dbo.Root"}
	profiles := []ProcedureProfile{root}
	for i := 0; i < 25; i++ {
		name := fmt.Sprintf("dbo.Child%02d", i)
		root.Calls = append(root.Calls, ProcedureCall{Object: name, Resolution: ResolutionStatic, Line: i + 1})
		profiles = append(profiles, ProcedureProfile{Name: name})
	}
	profiles[0] = root
	report := renderRecursiveProcedureMermaid(DesiredStateDocs{ProcedureProfiles: profiles}, root.Name)
	if strings.Count(report, " -->|") != 25 {
		t.Fatalf("all recursive call edges must remain visible after map splitting: %s", report)
	}
	if !strings.Contains(report, "Recursive call map 2 of 2") {
		t.Fatalf("large recursive graphs must split into renderable maps: %s", report)
	}
}

func TestLocalExecutionFlowDoesNotHideCallsInsideTransaction(t *testing.T) {
	profile := ProcedureProfile{
		Name: "dbo.Root",
		ControlSteps: []ProcedureControlStep{
			{Event: "TRANSACTION START", Evidence: "BEGIN TRAN", Line: 2},
			{Event: "COMMIT", Evidence: "COMMIT TRAN", Line: 5},
		},
		Calls: []ProcedureCall{{Object: "dbo.Child", Resolution: ResolutionStatic, Line: 3}},
	}
	report := renderProcedureMermaid(&profile)
	if !strings.Contains(report, "EXEC<br/>dbo.Child") || strings.Contains(report, "TRANSACTION<br/>lines 2–5") {
		t.Fatalf("transaction summarization must not hide procedure calls: %s", report)
	}
}

func hasAccess(items []ProcedureDataAccess, operation, object, resolution string) bool {
	for _, item := range items {
		if item.Operation == operation && item.Object == object && item.Resolution == resolution {
			return true
		}
	}
	return false
}
func hasCall(items []ProcedureCall, object, resolution string) bool {
	for _, item := range items {
		if item.Object == object && item.Resolution == resolution {
			return true
		}
	}
	return false
}
func hasDynamic(items []ProcedureDynamicSQL, resolution, object string) bool {
	for _, item := range items {
		if item.Resolution == resolution && containsFold(item.ResolvedObjects, object) {
			return true
		}
	}
	return false
}
