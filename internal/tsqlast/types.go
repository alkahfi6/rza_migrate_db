package tsqlast

type Backend string

const (
	BackendLite  Backend = "lite"
	BackendANTLR Backend = "antlr"
)

type FormatConfig struct {
	StripComments  bool
	Reindent       bool
	IndentSize     int
	Backend        Backend
	FallbackToLite bool
}

type FormatReport struct {
	Backend      Backend
	FallbackUsed bool
	Notes        []string
}
