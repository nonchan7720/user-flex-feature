package api

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/logging"
	sloggin "github.com/samber/slog-gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func newGin(cfg config.ConfigLoader) *gin.Engine {
	if cfg.IsStaging() || cfg.IsProduction() {
		gin.SetMode(gin.ReleaseMode)
	}
	l := newGinJsonWriter()
	gin.DefaultWriter = l
	r := gin.New()
	if cfg.IsStaging() || cfg.IsProduction() {
		r.TrustedPlatform = gin.PlatformCloudflare
	}
	r.Use(otelgin.Middleware(cfg.OTEL().ServiceName))
	r.Use(sloggin.NewWithConfig(slog.Default(), sloggin.Config{
		WithUserAgent: true,
		WithRequestID: true,
	}))
	r.Use(Recovery())
	r.Use(gzip.Gzip(gzip.DefaultCompression))
	r.MaxMultipartMemory = 50 * 1024 * 1024 // 50MB
	return r
}

func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()
		defer func() {
			if err := recover(); err != nil {
				e := err.(error)
				slog.With(logging.WithStack(e)).ErrorContext(ctx, e.Error())
				mp := map[string]string{
					"error": e.Error(),
				}
				c.AbortWithStatusJSON(http.StatusInternalServerError, &mp)
			}
		}()
		c.Next()
	}
}

type ginJsonWriter struct{}

var (
	_ io.Writer = &ginJsonWriter{}
)

func newGinJsonWriter() *ginJsonWriter {
	return &ginJsonWriter{}
}

func (w *ginJsonWriter) Write(p []byte) (n int, err error) {
	logging.Debug(string(p))
	return len(p), nil
}
