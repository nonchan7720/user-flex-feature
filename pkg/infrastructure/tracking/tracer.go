package tracking

import (
	"strings"

	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/tracking/trace"
	"go.opentelemetry.io/otel"
	oteltrace "go.opentelemetry.io/otel/trace"
)

var (
	serviceRoot = ""
)

func Tracer(name string, opts ...oteltrace.TracerOption) oteltrace.Tracer {
	t := otel.Tracer(strings.Join([]string{serviceRoot, name}, "/"), opts...)
	return trace.Tracer(t)
}
