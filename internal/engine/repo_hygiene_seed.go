package engine

import (
	"errors"
	"os"

	"rza_migrate_db/internal/output"
)

var optionalRepoFileTemplates = []struct {
	path string
	body string
}{
	{
		path: "constraints_naming.md",
		body: "# Constraint Naming Guideline\n\nConstraint names must be explicit and stable so migrations can drop or recreate them deterministically. Prefixes such as PK_, DF_, UQ_, CK_, and FK_ are recommendations, not mandatory validation rules.\n",
	},
	{
		path: ".gitignore",
		body: "out/**\n*.output.json\n",
	},
	{
		path: ".gitattributes",
		body: "* text=auto eol=lf\n*.sql text eol=lf\n*.md text eol=lf\n*.go text eol=lf\n",
	},
}

// seedOptionalRepoFilesForBootstrap creates only missing helper files. Existing
// user files are never appended to or overwritten, even when incomplete.
func (e Engine) seedOptionalRepoFilesForBootstrap(report *output.Report) error {
	for _, template := range optionalRepoFileTemplates {
		path := e.cfg.RepoPath(template.path)
		file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o644)
		if errors.Is(err, os.ErrExist) {
			continue
		}
		if err != nil {
			return err
		}
		if _, err := file.WriteString(template.body); err != nil {
			_ = file.Close()
			_ = os.Remove(path)
			return err
		}
		if err := file.Sync(); err != nil {
			_ = file.Close()
			_ = os.Remove(path)
			return err
		}
		if err := file.Close(); err != nil {
			_ = os.Remove(path)
			return err
		}
		report.AddPass(output.Item{
			Rule: "OPTIONAL_REPO_FILE_SEEDED", Severity: 0, File: template.path,
			Message: "bootstrap created optional repository helper file",
			Reason:  "CREATED", Action: "ok",
		})
	}
	return nil
}
