package redis

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"log/slog"
	"os"
	"time"

	config_redis "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/redis"
	config "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/tls"
	"github.com/redis/go-redis/extra/redisotel/v9"
	"github.com/redis/go-redis/v9"
)

func NewRedis(ctx context.Context, cfg *config_redis.Redis) (Cmdable, error) {
	var (
		c   Cmdable
		err error
	)
	c, err = buildStandalone(cfg)
	if err != nil {
		return nil, err
	}
	_ = redisotel.InstrumentTracing(c)
	_ = redisotel.InstrumentMetrics(c)
	if err := c.Ping(ctx).Err(); err != nil {
		return nil, err
	}
	return c, nil
}

func buildStandalone(cfg *config_redis.Redis) (*redis.Client, error) {
	opt, err := redis.ParseURL(cfg.ConnectionURL)
	if err != nil {
		return nil, err
	}
	if cfg.Password != "" {
		opt.Password = cfg.Password
	}
	if cfg.Username != "" {
		opt.Username = cfg.Username
	}
	if cfg.TLS != nil {
		if err := setupTLSConfig(cfg.TLS, opt); err != nil {
			return nil, err
		}
	}

	opt.ConnMaxIdleTime = time.Duration(cfg.IdleTimeout) * time.Second

	return redis.NewClient(opt), nil
}

func setupTLSConfig(opts *config.TLS, opt *redis.Options) error {
	if opts.InsecureSkipVerify {
		if opt.TLSConfig == nil {
			/* #nosec */
			opt.TLSConfig = &tls.Config{}
		}

		opt.TLSConfig.InsecureSkipVerify = true
	}

	if opts.CaFile != "" {
		rootCAs, err := x509.SystemCertPool()
		if err != nil {
			slog.Error("failed to load system cert pool for redis connection, falling back to empty cert pool")
		}
		if rootCAs == nil {
			rootCAs = x509.NewCertPool()
		}
		certs, err := os.ReadFile(opts.CaFile)
		if err != nil {
			return fmt.Errorf("failed to load %q, %v", opts.CaFile, err)
		}

		// Append our cert to the system pool
		if ok := rootCAs.AppendCertsFromPEM(certs); !ok {
			slog.Error("no certs appended, using system certs only")
		}

		if opt.TLSConfig == nil {
			opt.TLSConfig = &tls.Config{}
		}

		opt.TLSConfig.RootCAs = rootCAs
	}
	return nil
}

type Cmdable interface {
	redis.UniversalClient
	Close() error
}
