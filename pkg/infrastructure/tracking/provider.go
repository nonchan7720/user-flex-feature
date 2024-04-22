package tracking

import (
	"context"
	"log/slog"
	"time"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/logging"
	"github.com/nonchan7720/user-flex-feature/pkg/version"
	"github.com/samber/do"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

type ServiceRoot string

type TraceProvider struct {
	trace.TracerProvider

	cleanUp func()
}

func (tp *TraceProvider) Shutdown() error {
	tp.cleanUp()
	return nil
}

var (
	_ do.Shutdownable = &TraceProvider{}
)

func init() {
	do.Provide(container.Injector, Provide)
}

func Provide(i *do.Injector) (*TraceProvider, error) {
	var opts []sdktrace.TracerProviderOption
	ctx := do.MustInvoke[context.Context](i)
	tracking := do.MustInvoke[config.Tracking](i)
	serviceRoot := do.MustInvoke[ServiceRoot](i)

	tp, cleanup, err := NewTracerProvider(
		ctx,
		tracking.Enabled,
		tracking.AgentAddr,
		tracking.ServiceName,
		tracking.Environment,
		string(serviceRoot),
		opts...,
	)
	if err != nil {
		return nil, err
	}
	return &TraceProvider{
		TracerProvider: tp,
		cleanUp:        cleanup,
	}, nil
}

func NewTracerProvider(
	ctx context.Context,
	enabled bool,
	otelAgentAddr string,
	serviceName string,
	environment string,
	_serviceRoot string,
	opts ...sdktrace.TracerProviderOption,
) (trace.TracerProvider, func(), error) {
	var (
		tp      trace.TracerProvider
		cleanup func()
	)
	serviceRoot = _serviceRoot
	if !enabled {
		tp = noop.NewTracerProvider()
		cleanup = func() {}
	} else {
		traceClient := otlptracegrpc.NewClient(
			otlptracegrpc.WithInsecure(),
			otlptracegrpc.WithEndpoint(otelAgentAddr),
		)
		timeoutCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()
		exporter, err := otlptrace.New(timeoutCtx, traceClient)
		if err != nil {
			return nil, nil, err
		}
		r := NewResource(serviceName, version.Version, environment)

		sdkTP := sdktrace.NewTracerProvider(
			append([]sdktrace.TracerProviderOption{
				sdktrace.WithBatcher(exporter),
				sdktrace.WithResource(r),
			}, opts...)...,
		)
		pp := NewPropagator()
		otel.SetTextMapPropagator(pp)
		cleanup = func() {
			f := func(fn func(ctx context.Context) error) {
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				if err := fn(ctx); err != nil {
					slog.With(logging.WithStack(err)).Error(err.Error())
				}
				cancel()
			}
			f(sdkTP.ForceFlush)
			f(sdkTP.Shutdown)
			f(exporter.Shutdown)
		}
		tp = sdkTP
	}
	otel.SetTracerProvider(tp)
	return tp, cleanup, nil
}

func NewResource(serviceName string, version string, environment string) *resource.Resource {
	return resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String(serviceName),
		semconv.ServiceVersionKey.String(version),
		semconv.DeploymentEnvironmentKey.String(environment),
		attribute.String("environment", environment),
		attribute.String("env", environment),
	)
}

func NewPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
