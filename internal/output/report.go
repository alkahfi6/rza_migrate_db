package output

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Item struct {
	Rule                       string   `json:"rule"`
	Severity                   int      `json:"severity"`
	File                       string   `json:"file"`
	Message                    string   `json:"message"`
	Object                     string   `json:"object,omitempty"`
	Location                   string   `json:"location,omitempty"`
	Suggestion                 string   `json:"suggestion,omitempty"`
	Reason                     string   `json:"reason,omitempty"`
	Action                     string   `json:"action,omitempty"`
	PatchFolderOriginal        string   `json:"patch_folder_original,omitempty"`
	PatchNameCanonical         string   `json:"patch_name_canonical,omitempty"`
	GeneratedMigrationVersions []string `json:"generated_migration_versions,omitempty"`
}

type bucket struct {
	Count int    `json:"count"`
	Items []Item `json:"items"`
}

type naming struct {
	Pattern     string `json:"pattern"`
	LineEndings string `json:"line_endings"`
}

type newMigrations struct {
	Naming naming   `json:"naming"`
	Files  []string `json:"files"`
}

type Report struct {
	Fail          bucket                 `json:"fail"`
	Pass          bucket                 `json:"pass"`
	NewMigrations newMigrations          `json:"new_migrations"`
	Notes         map[string]interface{} `json:"notes,omitempty"`
}

func NewReport(pattern, lineEndings string) Report {
	return Report{
		Fail: bucket{
			Count: 0,
			Items: []Item{},
		},
		Pass: bucket{
			Count: 0,
			Items: []Item{},
		},
		NewMigrations: newMigrations{
			Naming: naming{
				Pattern:     pattern,
				LineEndings: lineEndings,
			},
			Files: []string{},
		},
		Notes: map[string]interface{}{},
	}
}

func (r *Report) AddFail(item Item) {
	r.Fail.Items = append(r.Fail.Items, item)
	r.Fail.Count = len(r.Fail.Items)
}

func (r *Report) AddPass(item Item) {
	r.Pass.Items = append(r.Pass.Items, item)
	r.Pass.Count = len(r.Pass.Items)
}

func (r *Report) SetNewMigrationFiles(files []string) {
	r.NewMigrations.Files = files
}

func (r *Report) Write(path string) error {
	return r.WriteWithOptions(path, true)
}

func (r *Report) WriteWithOptions(path string, includeNotes bool) error {
	dir := filepath.Dir(path)
	if err := os.MkdirAll(dir, 0o755); err != nil {
		return err
	}

	payload := Report{
		Fail:          r.Fail,
		Pass:          r.Pass,
		NewMigrations: r.NewMigrations,
	}
	if includeNotes {
		payload.Notes = r.Notes
	}

	data, err := json.MarshalIndent(payload, "", "  ")
	if err != nil {
		return err
	}
	data = append(data, '\n')
	return os.WriteFile(path, data, 0o644)
}
