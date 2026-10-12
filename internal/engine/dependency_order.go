package engine

import (
	"fmt"
	"sort"
	"strings"

	"rza_migrate_db/internal/diff"
	"rza_migrate_db/internal/sqlparse"
)

type dependencyOrderResult struct {
	Ops                []plannedMigration
	Edges              []string
	ReferencesScanned  int
	ResolvedReferences int
}

// orderPlannedMigrations keeps the desired/history model intact: it only uses
// desired source to discover dependencies and only reorders migrations planned
// for the current generation batch. Existing migration history is never moved.
func orderPlannedMigrations(ops []plannedMigration, desired map[string]desiredObject, order []string) (dependencyOrderResult, error) {
	if len(ops) < 2 {
		return dependencyOrderResult{Ops: ops}, nil
	}

	planned := make(map[string]int, len(ops))
	for index, op := range ops {
		planned[dependencyNodeKey(op.Category, op.Object)] = index
	}

	incoming := make([]map[int]struct{}, len(ops))
	outgoing := make([]map[int]struct{}, len(ops))
	for i := range ops {
		incoming[i] = make(map[int]struct{})
		outgoing[i] = make(map[int]struct{})
	}

	result := dependencyOrderResult{}
	// Desired-state removals have no source SQL to inspect. A table's explicit
	// index history must be removed before the table itself so the generated
	// up sequence remains deterministic; down naturally runs in reverse.
	for tableIndex, tableOp := range ops {
		if tableOp.Category != "tables" {
			continue
		}
		if _, stillDesired := desired[objectKey(tableOp.Category, tableOp.Object)]; stillDesired {
			continue
		}
		indexIndex, exists := planned[dependencyNodeKey("indexes", tableOp.Object)]
		if !exists || indexIndex == tableIndex {
			continue
		}
		outgoing[indexIndex][tableIndex] = struct{}{}
		incoming[tableIndex][indexIndex] = struct{}{}
		result.ResolvedReferences++
		result.Edges = append(result.Edges, ops[indexIndex].Object+" indexes -> "+tableOp.Object+" table")
	}
	for dependentIndex, op := range ops {
		if !isDependencyAwareCategory(op.Category) {
			continue
		}
		source, exists := desired[objectKey(op.Category, op.Object)]
		if !exists {
			// Removal ordering is the reverse of creation ordering: remove the
			// dependent object before the object it references. The down body
			// contains the last known definition and is therefore the correct
			// historical source for dependency discovery.
			for _, ref := range sqlparse.ExtractObjectReferences(op.DownSemanticSQL) {
				result.ReferencesScanned++
				for _, category := range dependencyReferenceCategories(ref.Kind) {
					dependencyIndex, plannedDependency := planned[dependencyNodeKey(category, ref.Name)]
					if !plannedDependency || dependencyIndex == dependentIndex {
						continue
					}
					if _, edgeExists := outgoing[dependentIndex][dependencyIndex]; edgeExists {
						continue
					}
					outgoing[dependentIndex][dependencyIndex] = struct{}{}
					incoming[dependencyIndex][dependentIndex] = struct{}{}
					result.ResolvedReferences++
					result.Edges = append(result.Edges, op.Object+" removal -> "+ops[dependencyIndex].Object+" removal")
				}
			}
			continue
		}
		for _, ref := range sqlparse.ExtractObjectReferences(source.SemanticSQL) {
			result.ReferencesScanned++
			for _, category := range dependencyReferenceCategories(ref.Kind) {
				dependencyIndex, exists := planned[dependencyNodeKey(category, ref.Name)]
				if !exists || dependencyIndex == dependentIndex {
					continue
				}
				if _, exists := outgoing[dependencyIndex][dependentIndex]; exists {
					continue
				}
				outgoing[dependencyIndex][dependentIndex] = struct{}{}
				incoming[dependentIndex][dependencyIndex] = struct{}{}
				result.ResolvedReferences++
				result.Edges = append(result.Edges, ops[dependencyIndex].Object+" -> "+op.Object)
			}
		}
	}

	orderIndex := diff.BuildOrderIndex(order)
	less := func(left, right int) bool {
		return diff.LessPlanned(orderIndex, len(order), ops[left].Category, ops[left].Object, ops[left].VK, ops[right].Category, ops[right].Object, ops[right].VK)
	}
	ready := make([]int, 0, len(ops))
	for index := range ops {
		if len(incoming[index]) == 0 {
			ready = append(ready, index)
		}
	}
	sort.Slice(ready, func(i, j int) bool { return less(ready[i], ready[j]) })

	ordered := make([]plannedMigration, 0, len(ops))
	for len(ready) > 0 {
		current := ready[0]
		ready = ready[1:]
		ordered = append(ordered, ops[current])

		next := make([]int, 0, len(outgoing[current]))
		for target := range outgoing[current] {
			next = append(next, target)
		}
		sort.Slice(next, func(i, j int) bool { return less(next[i], next[j]) })
		for _, target := range next {
			delete(incoming[target], current)
			if len(incoming[target]) == 0 {
				ready = append(ready, target)
			}
		}
		sort.Slice(ready, func(i, j int) bool { return less(ready[i], ready[j]) })
	}

	if len(ordered) != len(ops) {
		blocked := make([]string, 0)
		for index := range ops {
			if len(incoming[index]) > 0 {
				blocked = append(blocked, ops[index].Category+":"+ops[index].Object)
			}
		}
		sort.Strings(blocked)
		return dependencyOrderResult{}, fmt.Errorf("dependency cycle detected among planned objects: %s", strings.Join(blocked, ", "))
	}

	sort.Strings(result.Edges)
	result.Ops = ordered
	return result, nil
}

func isDependencyAwareCategory(category string) bool {
	switch category {
	case "types", "tables", "indexes", "views", "functions", "procedures":
		return true
	default:
		return false
	}
}

func dependencyReferenceCategories(kind sqlparse.ObjectReferenceKind) []string {
	switch kind {
	case sqlparse.ObjectReferenceType:
		return []string{"types"}
	case sqlparse.ObjectReferenceFunction:
		return []string{"functions"}
	case sqlparse.ObjectReferenceProcedure:
		return []string{"procedures"}
	default:
		return []string{"tables", "views"}
	}
}

func dependencyNodeKey(category, object string) string {
	return strings.TrimSpace(category) + "|" + strings.TrimSpace(object)
}
