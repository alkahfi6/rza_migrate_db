package engine

import "testing"

func TestVerifyMetasNilReturnsEmptySlice(t *testing.T) {
	t.Parallel()

	metas, err := verifyMetas(nil)
	if err != nil {
		t.Fatalf("verifyMetas(nil) returned error: %v", err)
	}
	if len(metas) != 0 {
		t.Fatalf("verifyMetas(nil) expected empty slice, got len=%d", len(metas))
	}
}

func TestVerifyHistoryMapNilReturnsEmptyMap(t *testing.T) {
	t.Parallel()

	historyMap, err := verifyHistoryMap(nil)
	if err != nil {
		t.Fatalf("verifyHistoryMap(nil) returned error: %v", err)
	}
	if len(historyMap) != 0 {
		t.Fatalf("verifyHistoryMap(nil) expected empty map, got len=%d", len(historyMap))
	}
}
