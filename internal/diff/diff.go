package diff

// State models semantic identity used for drift checks.
type State struct {
	Category       string
	Object         string
	NormalizedHash string
	SourceFile     string
}

// DriftKind classifies desired-vs-history mismatch direction.
type DriftKind string

const (
	DriftMissingInHistory DriftKind = "missing_in_history"
	DriftHashMismatch     DriftKind = "hash_mismatch"
	DriftMissingInDesired DriftKind = "missing_in_desired"
)

// DriftEntry captures one deterministic drift item.
type DriftEntry struct {
	Kind       DriftKind
	Category   string
	Object     string
	SourceFile string
}

// ComputeDrift compares desired state against history state.
func ComputeDrift(desired []State, history []State) []DriftEntry {
	historyByKey := make(map[string]State, len(history))
	for _, h := range history {
		historyByKey[objectKey(h.Category, h.Object)] = h
	}

	desiredKeys := make(map[string]struct{}, len(desired))
	out := make([]DriftEntry, 0)
	for _, d := range desired {
		key := objectKey(d.Category, d.Object)
		desiredKeys[key] = struct{}{}

		h, ok := historyByKey[key]
		if !ok {
			out = append(out, DriftEntry{
				Kind:       DriftMissingInHistory,
				Category:   d.Category,
				Object:     d.Object,
				SourceFile: d.SourceFile,
			})
			continue
		}
		if h.NormalizedHash != d.NormalizedHash {
			out = append(out, DriftEntry{
				Kind:       DriftHashMismatch,
				Category:   d.Category,
				Object:     d.Object,
				SourceFile: d.SourceFile,
			})
		}
	}

	for _, h := range history {
		key := objectKey(h.Category, h.Object)
		if _, ok := desiredKeys[key]; ok {
			continue
		}
		out = append(out, DriftEntry{
			Kind:       DriftMissingInDesired,
			Category:   h.Category,
			Object:     h.Object,
			SourceFile: h.SourceFile,
		})
	}

	return out
}

// BuildOrderIndex builds category ranking map from configured order.
func BuildOrderIndex(order []string) map[string]int {
	out := make(map[string]int, len(order))
	for idx, category := range order {
		out[category] = idx
	}
	return out
}

// LessCategoryObject compares category/object ordering based on orderIndex.
func LessCategoryObject(orderIndex map[string]int, orderSize int, categoryA, objectA, categoryB, objectB string) bool {
	oa, okA := orderIndex[categoryA]
	ob, okB := orderIndex[categoryB]
	if !okA {
		oa = orderSize + 1
	}
	if !okB {
		ob = orderSize + 1
	}
	if oa != ob {
		return oa < ob
	}
	return objectA < objectB
}

// LessPlanned compares category/object/vk ordering based on orderIndex.
func LessPlanned(orderIndex map[string]int, orderSize int, categoryA, objectA string, vkA int, categoryB, objectB string, vkB int) bool {
	if LessCategoryObject(orderIndex, orderSize, categoryA, objectA, categoryB, objectB) {
		return true
	}
	if LessCategoryObject(orderIndex, orderSize, categoryB, objectB, categoryA, objectA) {
		return false
	}
	return vkA < vkB
}

func objectKey(category, object string) string {
	return category + "|" + object
}
