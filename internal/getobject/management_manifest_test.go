package getobject

import "testing"

func TestManagementManifestMarksUnsupportedObjectsUnmanaged(t *testing.T) {
	summary := ExportSummary{
		Categories: []CategoryChange{
			{Dir: "db_object/tables", Unchanged: []string{"dbo.Legacy.sql", "dbo.Managed.sql"}},
			{Dir: "db_object/indexes", Unchanged: []string{"dbo.Legacy.sql", "dbo.Managed.sql"}},
			{Dir: "db_object/views", Unchanged: []string{"dbo.LegacyView.sql", "dbo.ManagedView.sql"}},
		},
		TableWarnings: []ExportWarning{{
			Rule: "TABLE_EXTRACT_SKIPPED_UNSUPPORTED", Object: "dbo.Legacy", Message: "unsupported index options",
		}},
		ObjectWarnings: []ExportWarning{{
			Rule: "OBJECT_EXTRACT_SKIPPED_UNSUPPORTED", Category: "views", Object: "dbo.LegacyView", Message: "unsupported module",
		}},
	}
	manifest := managementManifestFromSummary("dbo", summary)
	unmanaged := manifest.UnmanagedSet()
	for _, key := range []string{"tables|dbo.Legacy", "indexes|dbo.Legacy", "views|dbo.LegacyView"} {
		if _, ok := unmanaged[key]; !ok {
			t.Fatalf("missing unmanaged key %s: %#v", key, unmanaged)
		}
	}
	for _, item := range manifest.Managed {
		if item.Object == "dbo.Legacy" || item.Object == "dbo.LegacyView" {
			t.Fatalf("unsupported object remained managed: %#v", item)
		}
	}
}
