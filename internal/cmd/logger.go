package cmd

import (
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"rza_migrate_db/internal/migrest"
)

func buildLogger(cfg migrest.LoggingConfig) (*zap.Logger, error) {
	logCfg := zap.NewProductionConfig()
	logCfg.Encoding = cfg.Format
	logCfg.Level = zap.NewAtomicLevelAt(parseZapLevel(cfg.Level))
	logCfg.DisableCaller = cfg.DisableCaller
	logCfg.DisableStacktrace = true
	logCfg.EncoderConfig.TimeKey = "ts"
	logCfg.EncoderConfig.MessageKey = "msg"
	logCfg.EncoderConfig.LevelKey = "level"
	logCfg.EncoderConfig.CallerKey = "caller"
	logCfg.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	logCfg.EncoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	return logCfg.Build()
}

func parseZapLevel(raw string) zapcore.Level {
	switch strings.ToLower(strings.TrimSpace(raw)) {
	case "debug":
		return zapcore.DebugLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "info":
		fallthrough
	default:
		return zapcore.InfoLevel
	}
}
