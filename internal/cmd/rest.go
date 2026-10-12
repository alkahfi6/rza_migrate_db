package cmd

import (
	"context"
	"errors"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"

	"go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp"
	"go.uber.org/zap"

	"rza_migrate_db/internal/handler"
	"rza_migrate_db/internal/migrest"
	"rza_migrate_db/internal/telemetry"
)

// RunREST is entrypoint for rza_migrate_rest executable.
func RunREST(args []string) int {
	bootstrapCfg := migrest.DefaultConfig()
	logger, err := buildLogger(bootstrapCfg.Logging)
	if err != nil {
		logger = zap.NewNop()
	}
	restoreGlobals := zap.ReplaceGlobals(logger)
	defer func() {
		restoreGlobals()
		_ = logger.Sync()
	}()

	defaultConfigPath := "config.yaml"
	if len(args) > 0 {
		execDir := filepath.Dir(filepath.Clean(args[0]))
		defaultConfigPath = filepath.Join(execDir, "config.yaml")
	}

	fs := flag.NewFlagSet("rza_migrate_rest", flag.ContinueOnError)
	configPath := fs.String("config", defaultConfigPath, "path to runtime YAML config")
	fs.SetOutput(os.Stderr)
	if err := fs.Parse(args[1:]); err != nil {
		logger.Error("parse_cli_flags_failed",
			zap.String("component", "rest"),
			zap.Error(err),
		)
		return 1
	}

	cfgPath := strings.TrimSpace(*configPath)
	if cfgPath == "" {
		logger.Error("config_path_required",
			zap.String("component", "rest"),
		)
		return 1
	}

	cfg, err := migrest.LoadConfig(cfgPath)
	if err != nil {
		logger.Error("load_config_failed",
			zap.String("component", "rest"),
			zap.String("config_path", filepath.ToSlash(cfgPath)),
			zap.Error(err),
		)
		return 1
	}

	if configuredLogger, logErr := buildLogger(cfg.Logging); logErr != nil {
		logger.Warn("build_configured_logger_failed_using_bootstrap",
			zap.String("component", "rest"),
			zap.Error(logErr),
		)
	} else {
		restoreGlobals()
		_ = logger.Sync()
		logger = configuredLogger
		restoreGlobals = zap.ReplaceGlobals(logger)
	}

	shutdownTelemetry, err := telemetry.Init(context.Background(), telemetry.Config{
		Enabled:     cfg.Telemetry.Enabled,
		ServiceName: cfg.ServiceName,
		Endpoint:    cfg.Telemetry.Endpoint,
		Insecure:    cfg.Telemetry.Insecure,
		SampleRatio: cfg.Telemetry.SampleRatio,
	})
	if err != nil {
		logger.Error("initialize_telemetry_failed",
			zap.String("component", "rest"),
			zap.Error(err),
		)
		return 1
	}
	defer func() {
		shutdownCtx, cancel := context.WithTimeout(
			context.Background(),
			time.Duration(cfg.Server.ShutdownTimeoutSeconds)*time.Second,
		)
		defer cancel()
		if shutErr := shutdownTelemetry(shutdownCtx); shutErr != nil {
			logger.Warn("shutdown_telemetry_failed",
				zap.String("component", "rest"),
				zap.Error(shutErr),
			)
		}
	}()

	service, err := migrest.NewService(migrest.NewServiceConfigFromRuntime(cfg))
	if err != nil {
		logger.Error("initialize_migration_service_failed",
			zap.String("component", "rest"),
			zap.Error(err),
		)
		return 1
	}
	defer func() {
		if closeErr := service.Close(); closeErr != nil {
			logger.Warn("close_service_failed",
				zap.String("component", "rest"),
				zap.Error(closeErr),
			)
		}
	}()

	router := handler.NewMigrationHandler(service, cfg.Security, cfg.ServiceName).Router()
	withMiddleware := handler.WithTraceIDHeader(handler.WithAccessLog(router))
	httpHandler := otelhttp.NewHandler(withMiddleware, cfg.ServiceName+".http")
	server := &http.Server{
		Addr:              cfg.Server.Listen,
		Handler:           httpHandler,
		ReadHeaderTimeout: time.Duration(cfg.Server.ReadTimeoutSeconds) * time.Second,
		ReadTimeout:       time.Duration(cfg.Server.ReadTimeoutSeconds) * time.Second,
		WriteTimeout:      time.Duration(cfg.Server.WriteTimeoutSeconds) * time.Second,
		IdleTimeout:       time.Duration(cfg.Server.IdleTimeoutSeconds) * time.Second,
	}

	errCh := make(chan error, 1)
	go func() {
		logger.Info("rest_server_start",
			zap.String("component", "rest"),
			zap.String("listen", cfg.Server.Listen),
			zap.String("config_path", filepath.ToSlash(cfgPath)),
		)
		err := server.ListenAndServe()
		if err != nil && !errors.Is(err, http.ErrServerClosed) {
			errCh <- err
			return
		}
		errCh <- nil
	}()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt, syscall.SIGTERM)
	defer signal.Stop(sigCh)

	select {
	case err := <-errCh:
		if err != nil {
			logger.Error("rest_server_failed",
				zap.String("component", "rest"),
				zap.Error(err),
			)
			return 1
		}
		return 0
	case sig := <-sigCh:
		logger.Info("shutdown_signal_received",
			zap.String("component", "rest"),
			zap.String("signal", sig.String()),
		)
		shutdownCtx, cancel := context.WithTimeout(context.Background(), time.Duration(cfg.Server.ShutdownTimeoutSeconds)*time.Second)
		defer cancel()
		if err := server.Shutdown(shutdownCtx); err != nil {
			logger.Error("graceful_shutdown_failed",
				zap.String("component", "rest"),
				zap.Error(err),
			)
			return 1
		}
		err := <-errCh
		if err != nil {
			logger.Error("rest_server_failed_during_shutdown",
				zap.String("component", "rest"),
				zap.Error(err),
			)
			return 1
		}
		logger.Info("server_stopped_gracefully",
			zap.String("component", "rest"),
		)
		return 0
	}
}
