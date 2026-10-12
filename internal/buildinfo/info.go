package buildinfo

import "strings"

// These values are intentionally linker-overridable. Development builds keep
// explicit, non-empty values so reports can still distinguish an unversioned
// local binary from a release artifact.
var (
	Version      = "dev"
	Commit       = "unknown"
	BuildTime    = "unknown"
	SourceSHA256 = "unknown"
)

type Info struct {
	Version      string `json:"version"`
	Commit       string `json:"commit"`
	BuildTime    string `json:"build_time"`
	SourceSHA256 string `json:"source_sha256"`
}

func Current() Info {
	return Info{
		Version:      normalized(Version, "dev"),
		Commit:       normalized(Commit, "unknown"),
		BuildTime:    normalized(BuildTime, "unknown"),
		SourceSHA256: normalized(SourceSHA256, "unknown"),
	}
}

func normalized(value, fallback string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return fallback
	}
	return value
}
