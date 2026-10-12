package migreport

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"
)

func TestProcedureProfileLinksEscapeFragmentCharacters(t *testing.T) {
	docs := DesiredStateDocs{ProcedureProfiles: []ProcedureProfile{{Name: "dbo.Report#ORI"}}}
	for _, rendered := range []string{renderProcedureIndex(docs), renderProcedureCatalog(docs)} {
		if !strings.Contains(rendered, "profiles/dbo.Report%23ORI.md") || strings.Contains(rendered, "profiles/dbo.Report#ORI.md") {
			t.Fatalf("profile link must URL-escape the fragment marker: %s", rendered)
		}
	}
}

func TestModuleLineageDedupesCaseAndDoesNotTreatReturnAssignmentAsDynamic(t *testing.T) {
	modules := []moduleSource{{
		Name: "dbo.Root", Kind: "procedure",
		SQL: "SELECT * FROM dbo.SourceTable; SELECT * FROM dbo.sourcetable; EXEC @rc = dbo.Child;",
	}}
	_, dependencies, dynamic := analyzeModuleRelationships(modules, nil, nil)
	if len(dependencies) != 1 {
		t.Fatalf("case-only references must collapse to one lineage dependency: %+v", dependencies)
	}
	if len(dynamic) != 0 {
		t.Fatalf("static EXEC return assignment must not be labeled dynamic SQL: %+v", dynamic)
	}

	modules[0].SQL = "EXEC @runtime_sql;"
	_, _, dynamic = analyzeModuleRelationships(modules, nil, nil)
	if len(dynamic) != 1 || dynamic[0] != "dbo.Root" {
		t.Fatalf("actual variable execution must remain dynamic: %+v", dynamic)
	}
}

func TestRelationalERDShowsEveryPairOnceInBoundedRelationalMaps(t *testing.T) {
	docs := DesiredStateDocs{}
	for i := 0; i < 80; i++ {
		docs.Relationships = append(docs.Relationships, ObservedRelationship{
			LeftObject:   fmt.Sprintf("dbo.Table%03d", i),
			LeftColumns:  []string{"ParentId"},
			RightObject:  fmt.Sprintf("dbo.Table%03d", i+1),
			RightColumns: []string{"Id"},
			Evidence:     "OBSERVED_JOIN",
			Occurrences:  1,
		})
	}
	docs.Relationships = append(docs.Relationships, ObservedRelationship{
		LeftObject: "dbo.Table040", LeftColumns: []string{"ParentId"},
		RightObject: "dbo.Table040", RightColumns: []string{"Id"},
		Evidence: "OBSERVED_JOIN", Occurrences: 1,
	})

	markdown := renderRelationalERD(docs)
	if !strings.Contains(markdown, "Relational Map 1 of") || strings.Contains(markdown, relationshipDiagramDirectory+"/") {
		t.Fatalf("large connected group must stay in one navigable document: %s", markdown)
	}
	if strings.Contains(markdown, "### Hub:") || !strings.Contains(markdown, "Shared object navigation") || !strings.Contains(markdown, "### Object index") {
		t.Fatalf("relational maps must not invent a Hub and must preserve indexed navigation across map boundaries: %s", markdown)
	}
	if strings.Contains(markdown, "|\n```mermaid") {
		t.Fatal("Mermaid fence must be separated from a preceding Markdown table by a blank line")
	}
	diagramPairs := 0
	for _, block := range strings.Split(markdown, "```mermaid\n")[1:] {
		block = strings.SplitN(block, "```", 2)[0]
		nodes := strings.Count(block, "[\"`")
		if nodes > relationalMapMaxObjects {
			t.Fatalf("Mermaid block contains %d objects; maximum is %d", nodes, relationalMapMaxObjects)
		}
		edges := strings.Count(block, " ---|") + strings.Count(block, " -.-|")
		diagramPairs += edges
		if edges > relationalMapMaxPairs {
			t.Fatalf("Mermaid block contains %d lines; maximum is %d", edges, relationalMapMaxPairs)
		}
		if len("```mermaid\n"+block+"```") > relationalMapMaxTextCharacters {
			t.Fatalf("Mermaid block exceeds text budget: %d", len(block))
		}
		if !mermaidSetIsConnected(block) {
			t.Fatalf("Mermaid Connection Set must be connected:\n%s", block)
		}
	}
	if diagramPairs != 80 {
		t.Fatalf("every non-self object pair must appear exactly once: got %d, want 80", diagramPairs)
	}
	if strings.Contains(markdown, "peran A") || strings.Contains(markdown, "peran B") || !strings.Contains(markdown, "### Self-join review") || !strings.Contains(markdown, "ParentId = Id") {
		t.Fatal("self-join must be removed from overview diagrams and listed in the review section")
	}
	catalog := renderRelationshipCatalog(docs)
	if strings.Count(catalog, "| `OBSERVED_JOIN` |") != 81 || !strings.Contains(catalog, "### Self-join Review") {
		t.Fatal("relationship catalog must retain every exact mapping and separate self-joins")
	}
	if strings.Contains(catalog, "### Hub:") || !strings.Contains(catalog, "### Relational Map") {
		t.Fatal("catalog navigation must follow Group/Relational Map without artificial Hub ownership")
	}
}

func TestRelationalERDCollapsesColumnMappingsPerObjectPair(t *testing.T) {
	docs := DesiredStateDocs{Relationships: []ObservedRelationship{
		{LeftObject: "dbo.A", LeftColumns: []string{"Id"}, RightObject: "dbo.B", RightColumns: []string{"AId"}, Evidence: "OBSERVED_JOIN", Occurrences: 2},
		{LeftObject: "dbo.A", LeftColumns: []string{"Code"}, RightObject: "dbo.B", RightColumns: []string{"ACode"}, Evidence: "OBSERVED_JOIN", Occurrences: 3},
	}}
	erd := renderRelationalERD(docs)
	if strings.Count(erd, " -.-|") != 1 || !strings.Contains(erd, "Id = AId") || !strings.Contains(erd, "Code = ACode") || strings.Contains(erd, "[JOIN]") || strings.Contains(erd, "usage 5x") {
		t.Fatalf("diagram must show every exact column mapping and omit usage aggregation: %s", erd)
	}
	catalog := renderRelationshipCatalog(docs)
	if strings.Count(catalog, "| `OBSERVED_JOIN` |") != 2 {
		t.Fatalf("catalog must retain both exact column mappings: %s", catalog)
	}
}

func TestRelationalERDShowsObjectTypeAndKeyContextOnNodes(t *testing.T) {
	docs := DesiredStateDocs{
		Tables: []DocumentedTable{
			{
				Name:    "dbo.Parent",
				Columns: []DocumentedColumn{{Name: "Id", Identity: true}, {Name: "Code"}},
				Keys: []DocumentedKey{
					{Columns: []string{"Id"}, Evidence: "PRIMARY_KEY"},
					{Columns: []string{"Code", "Id"}, Evidence: "UNIQUE_INDEX"},
				},
			},
			{
				Name:    "dbo.Child",
				Columns: []DocumentedColumn{{Name: "ParentId"}},
				Indexes: []DocumentedIndex{{Name: "IX_Child_ParentId", Columns: []string{"ParentId"}}},
			},
		},
		Modules: []DocumentedModule{{Name: "dbo.ParentView", Kind: "view"}},
		Relationships: []ObservedRelationship{
			{LeftObject: "dbo.Child", LeftColumns: []string{"ParentId"}, RightObject: "dbo.Parent", RightColumns: []string{"Id"}, Evidence: "OBSERVED_JOIN"},
			{LeftObject: "dbo.Parent", LeftColumns: []string{"Code"}, RightObject: "dbo.ParentView", RightColumns: []string{"Code"}, Evidence: "OBSERVED_JOIN"},
		},
	}

	erd := renderRelationalERD(docs)
	for _, expected := range []string{
		"**dbo.Parent**\nTABLE\nCode — PART OF UNIQUE KEY (Code, Id)\nId — PRIMARY KEY · IDENTITY",
		"**dbo.Child**\nTABLE\nParentId — NON-UNIQUE INDEX",
		"**dbo.ParentView**\nVIEW\nCode — UNIQUENESS UNKNOWN",
		"classDef table", "classDef view", "### Object index",
	} {
		if !strings.Contains(erd, expected) {
			t.Fatalf("ERD missing object/key context %q:\n%s", expected, erd)
		}
	}
}

func TestRelationalERDUsesEvidenceTagsOnlyForMixedPair(t *testing.T) {
	docs := DesiredStateDocs{Relationships: []ObservedRelationship{
		{LeftObject: "dbo.A", LeftColumns: []string{"ParentId"}, RightObject: "dbo.B", RightColumns: []string{"Id"}, Evidence: "DECLARED_FK"},
		{LeftObject: "dbo.A", LeftColumns: []string{"Code"}, RightObject: "dbo.B", RightColumns: []string{"Code"}, Evidence: "OBSERVED_JOIN"},
	}}
	erd := renderRelationalERD(docs)
	if !strings.Contains(erd, "ParentId = Id [FK]") || !strings.Contains(erd, "Code = Code [JOIN]") || strings.Contains(erd, " ---|") {
		t.Fatalf("mixed-evidence pair must remain dashed and label each mapping evidence: %s", erd)
	}
}

func TestRelationalERDHardLimitsDenseHubToTwentyObjects(t *testing.T) {
	docs := DesiredStateDocs{}
	for i := 0; i < 40; i++ {
		docs.Relationships = append(docs.Relationships, ObservedRelationship{
			LeftObject: "dbo.Hub", LeftColumns: []string{"Id"},
			RightObject: fmt.Sprintf("dbo.Leaf%02d", i), RightColumns: []string{"HubId"},
			Evidence: "OBSERVED_JOIN",
		})
	}
	erd := renderRelationalERD(docs)
	totalEdges := 0
	for _, block := range strings.Split(erd, "```mermaid\n")[1:] {
		block = strings.SplitN(block, "```", 2)[0]
		nodes := strings.Count(block, "[\"`")
		edges := strings.Count(block, " -.-|") + strings.Count(block, " ---|")
		totalEdges += edges
		if nodes > relationalMapMaxObjects || edges > relationalMapMaxPairs {
			t.Fatalf("map exceeded hard limits: objects=%d edges=%d", nodes, edges)
		}
	}
	if totalEdges != 40 {
		t.Fatalf("hard-limit partition lost relationships: got %d, want 40", totalEdges)
	}
}

func TestRelationshipCatalogSplitsDenseSinglePairWithoutLosingMappings(t *testing.T) {
	docs := DesiredStateDocs{}
	for i := 0; i < 80; i++ {
		docs.Relationships = append(docs.Relationships, ObservedRelationship{
			LeftObject: "dbo.A", LeftColumns: []string{fmt.Sprintf("Key%02d", i)},
			RightObject: "dbo.B", RightColumns: []string{fmt.Sprintf("AKey%02d", i)},
			Evidence: "OBSERVED_JOIN", Occurrences: 1,
		})
	}
	erd := renderRelationalERD(docs)
	if strings.Count(erd, " -.-|") != 3 || !strings.Contains(erd, "Mapping batch 1/3") || !strings.Contains(erd, "Mapping batch 3/3") || !strings.Contains(erd, "Key00 = AKey00") || !strings.Contains(erd, "Key79 = AKey79") || strings.Contains(erd, "usage 80x") {
		t.Fatalf("dense pair must use bounded mapping batches without losing exact mappings: %s", erd)
	}
	for _, block := range strings.Split(erd, "```mermaid\n")[1:] {
		block = strings.SplitN(block, "```", 2)[0]
		if len(block) > relationalMapMaxTextCharacters {
			t.Fatalf("dense Mermaid block exceeds its text budget: chars=%d", len(block))
		}
	}
	catalog := renderRelationshipCatalog(docs)
	if strings.Count(catalog, "| `OBSERVED_JOIN` |") != 80 {
		t.Fatalf("catalog lost an exact mapping: %s", catalog)
	}
	for _, table := range strings.Split(catalog, "| Evidence |")[1:] {
		rows := strings.Count(strings.SplitN(table, "\n\n", 2)[0], "| `OBSERVED_JOIN` |")
		if rows > 35 {
			t.Fatalf("catalog table contains %d data rows; maximum is 35", rows)
		}
	}
}

func TestDesiredStateWriterRemovesObsoleteSplitDiagramDirectory(t *testing.T) {
	out := t.TempDir()
	stale := filepath.Join(out, relationshipDiagramDirectory)
	if err := os.MkdirAll(stale, 0755); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(stale, "old.md"), []byte("obsolete"), 0644); err != nil {
		t.Fatal(err)
	}
	if err := writeDesiredStateDocs(out, DesiredStateDocs{}); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(stale); !os.IsNotExist(err) {
		t.Fatalf("obsolete generated diagram directory must be removed, stat error: %v", err)
	}
}

func mermaidSetIsConnected(block string) bool {
	nodePattern := regexp.MustCompile(`(?m)^\s+(n\d+)\[`)
	edgePattern := regexp.MustCompile(`(?m)^\s+(n\d+) (?:---|-\.-)\|.*\| (n\d+)$`)
	nodes := nodePattern.FindAllStringSubmatch(block, -1)
	if len(nodes) == 0 {
		return false
	}
	adjacency := map[string][]string{}
	for _, node := range nodes {
		adjacency[node[1]] = nil
	}
	for _, edge := range edgePattern.FindAllStringSubmatch(block, -1) {
		adjacency[edge[1]] = append(adjacency[edge[1]], edge[2])
		adjacency[edge[2]] = append(adjacency[edge[2]], edge[1])
	}
	seen := map[string]bool{nodes[0][1]: true}
	queue := []string{nodes[0][1]}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		for _, next := range adjacency[current] {
			if !seen[next] {
				seen[next] = true
				queue = append(queue, next)
			}
		}
	}
	return len(seen) == len(nodes)
}

func TestDesiredStateMarkdownDoesNotUseRawHTMLBreaks(t *testing.T) {
	docs := DesiredStateDocs{Relationships: []ObservedRelationship{{
		LeftObject: "dbo.A", LeftColumns: []string{"Id"}, RightObject: "dbo.B", RightColumns: []string{"AId"},
		Evidence: "OBSERVED_JOIN", JoinTypes: []string{"INNER"}, Modules: []string{"dbo.p1", "dbo.p2"}, Occurrences: 2, Confidence: "MEDIUM",
	}}}
	for name, markdown := range map[string]string{
		"relationship_catalog.md": renderRelationshipCatalog(docs),
		"relational_erd.md":       renderRelationalERD(docs),
	} {
		if strings.Contains(markdown, "<br>") {
			t.Fatalf("%s contains raw HTML line break", name)
		}
	}
}
