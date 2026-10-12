package engine

type desiredObject struct {
	Category       string
	Object         string
	SourceFile     string
	SemanticSQL    string
	NormalizedHash string
	Subtype        string
	RenameColumns  map[string]string // lower-case old name -> exact new name; derived from source comments
}

type historyObject struct {
	Category       string
	Object         string
	LastVersion    string
	LastVK         int
	SemanticSQL    string
	NormalizedHash string
	Subtype        string
	Deleted        bool
	ManualDropOnly bool
}

type patchSource struct {
	FolderName    string
	CanonicalName string
	RelDir        string
	ReadmePath    string
	UpPath        string
	DownPath      string
	UpSQL         string
	DownSQL       string
	UpHash        string
	DownHash      string
}

type allocatedMigration struct {
	Op       plannedMigration
	Version  string
	UpName   string
	DownName string
}

type plannedMigration struct {
	Category            string
	Object              string
	VK                  int
	UpSemanticSQL       string
	DownSemanticSQL     string
	HasDestructiveDown  bool
	ManualDropOnly      bool
	RemovedColumns      []string
	SourceFile          string
	PatchFolderOriginal string
	PatchNameCanonical  string
	PatchReadmePath     string
	PatchUpHash         string
	PatchDownHash       string
}

type generatedFile struct {
	Name    string
	Content string
}
