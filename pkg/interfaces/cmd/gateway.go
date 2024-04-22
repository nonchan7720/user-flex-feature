package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/logging"
	"github.com/pkg/errors"
	"github.com/samber/do"
)

func gatewayCommand() {}

type GatewayShutdown func(ctx context.Context)

func runGatewayServer(cfg *config.Gateway) GatewayShutdown {
	do.Override(container.Injector, func(_ *do.Injector) (*config.Gateway, error) { return cfg, nil })
	srv := do.MustInvokeNamed[*http.Server](container.Injector, "gateway")
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		err := errors.Errorf("failed to listen: %v", err)
		panic(err)
	}
	go func() {
		if err := srv.Serve(lis); err != nil && err != http.ErrServerClosed {
			slog.With(logging.WithStack(err)).Error("Gateway Shutdown.")
		}
	}()
	return func(ctx context.Context) {
		if err := srv.Shutdown(ctx); err != nil {
			slog.With(logging.WithStack(err)).Error("Gateway forced to shutdown")
		}
		slog.Info("Gateway exiting")
	}
}
