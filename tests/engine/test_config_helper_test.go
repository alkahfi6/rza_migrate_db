package engine_test

import "rza_migrate_db/internal/config"

func testConfig(workRoot string) config.AppConfig {
	cfg := config.Default(workRoot)
	// Fixture expectations in tests/engine are written for non-strict precheck mode.
	cfg.EnableStrictMode = false
	return cfg
}
