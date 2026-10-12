package telemetry

import (
	"context"
	"strings"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

// Config is runtime telemetry configuration for OTel tracing.
type Config struct {
	Enabled     bool
	ServiceName string
	Endpoint    string
	Insecure    bool
	SampleRatio float64
}

// Init initializes global OpenTelemetry tracer provider.
// Returns a shutdown function that must be called during service stop.
func Init(ctx context.Context, cfg Config) (func(context.Context) error, error) {
	if !cfg.Enabled {
		return func(context.Context) error { return nil }, nil
	}

	serviceName := strings.TrimSpace(cfg.ServiceName)
	if serviceName == "" {
		serviceName = "migration_rest_service"
	}

	sampleRatio := cfg.SampleRatio
	if sampleRatio <= 0 {
		sampleRatio = 1.0
	}
	if sampleRatio > 1 {
		sampleRatio = 1.0
	}

	exporterOpts := make([]otlptracehttp.Option, 0, 2)
	endpoint := strings.TrimSpace(cfg.Endpoint)
	if endpoint != "" {
		if strings.HasPrefix(endpoint, "http://") || strings.HasPrefix(endpoint, "https://") {
			exporterOpts = append(exporterOpts, otlptracehttp.WithEndpointURL(endpoint))
		} else {
			exporterOpts = append(exporterOpts, otlptracehttp.WithEndpoint(endpoint))
			if cfg.Insecure {
				exporterOpts = append(exporterOpts, otlptracehttp.WithInsecure())
			}
		}
	}

	exporter, err := otlptracehttp.New(ctx, exporterOpts...)
	if err != nil {
		return nil, err
	}

	res, err := resource.New(
		ctx,
		resource.WithAttributes(
			attribute.String("service.name", serviceName),
		),
	)
	if err != nil {
		return nil, err
	}

	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exporter),
		sdktrace.WithResource(res),
		sdktrace.WithSampler(sdktrace.ParentBased(sdktrace.TraceIDRatioBased(sampleRatio))),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(
		propagation.NewCompositeTextMapPropagator(
			propagation.TraceContext{},
			propagation.Baggage{},
		),
	)

	return tp.Shutdown, nil
}
