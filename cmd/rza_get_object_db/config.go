package main

import (
	"os"
	"strings"

	"rza_migrate_db/internal/getobject"
)

// The executable never embeds database credentials. Supply a JSON file with
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
