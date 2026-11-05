package main

import (
	"errors"
	"os"
	"strings"

	"rza_migrate_db/internal/getobject"
)

// The executable never embeds database credentials. Supply a YAML/JSON file with
// -config or set MSSQL_DSN (RZA_GET_OBJECT_DSN remains a compatibility alias).
var compiledGetObjectConfig = getobject.DefaultAppConfig(".")

func loadCompiledGetObjectConfig() getobject.AppConfig {
	cfg := compiledGetObjectConfig
	if dsn := strings.TrimSpace(os.Getenv("MSSQL_DSN")); dsn != "" {
		cfg.ConnectionString = dsn
	} else if dsn := strings.TrimSpace(os.Getenv("RZA_GET_OBJECT_DSN")); dsn != "" {
		cfg.ConnectionString = dsn
	}
	cfg.Normalize()
	return cfg
}

// Explicit files override the local YAML, which overrides legacy DSN environment
// variables. Offline preview does not require or auto-load database credentials.
func loadGetObjectConfig(path string, preview bool) (getobject.AppConfig, error) {
	if strings.TrimSpace(path) != "" {
		return getobject.LoadAppConfig(path)
	}
	if !preview {
		if _, err := os.Stat("config.yaml"); err == nil {
			return getobject.LoadAppConfig("config.yaml")
		} else if !errors.Is(err, os.ErrNotExist) {
			return getobject.AppConfig{}, err
		}
	}
	return loadCompiledGetObjectConfig(), nil
}
