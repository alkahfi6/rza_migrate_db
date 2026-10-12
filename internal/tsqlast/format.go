package tsqlast

import (
	"errors"
	"fmt"
)

var ErrANTLRBackendUnavailable = errors.New("antlr backend unavailable in current build")

func FormatModuleSQL(input string, cfg FormatConfig) (string, FormatReport, error) {
	cfg = normalizeConfig(cfg)

	switch cfg.Backend {
	case BackendANTLR:
		out, rep, err := formatWithANTLR(input, cfg)
		if err == nil {
			return out, rep, nil
		}
		if cfg.FallbackToLite {
			liteOut := formatWithLite(input, cfg)
			return liteOut, FormatReport{
				Backend:      BackendLite,
				FallbackUsed: true,
				Notes:        []string{fmt.Sprintf("antlr backend failed: %v", err)},
			}, nil
		}
		return "", FormatReport{Backend: BackendANTLR}, err
	case BackendLite:
		return formatWithLite(input, cfg), FormatReport{Backend: BackendLite}, nil
	default:
		return "", FormatReport{}, fmt.Errorf("unsupported tsqlast backend: %s", cfg.Backend)
	}
}

func normalizeConfig(cfg FormatConfig) FormatConfig {
	if cfg.IndentSize <= 0 {
		cfg.IndentSize = 2
	}
	if cfg.Backend == "" {
		cfg.Backend = BackendANTLR
	}
	return cfg
}

func formatWithANTLR(input string, cfg FormatConfig) (string, FormatReport, error) {
	return "", FormatReport{Backend: BackendANTLR}, ErrANTLRBackendUnavailable
}
