package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"os/signal"
	"syscall"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/logging"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/tracking"
	"github.com/pkg/errors"
	"github.com/samber/do"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

type severArgs struct {
	configFilePath string
}

func (args *severArgs) Validate() error {
	return validation.ValidateStruct(args,
		validation.Field(&args.configFilePath, validation.Required),
	)
}

func serverCommand() *cobra.Command {
	var operationArgs severArgs
	cmd := &cobra.Command{
		Use:           "server",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			serverExecute(cmd.Context(), &operationArgs)
		},
	}
	flag := cmd.Flags()
	flag.StringVarP(&operationArgs.configFilePath, "config", "c", "config.yaml", "configuration file path")
	return cmd
}

func serverExecute(ctx context.Context, args *severArgs) {
	if err := args.Validate(); err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	cfg := loadConfig[config.Config](args.configFilePath)
	do.Override(container.Injector, func(i *do.Injector) (context.Context, error) { return ctx, nil })
	do.Override(container.Injector, func(_ *do.Injector) (*config.Config, error) { return &cfg, nil })
	do.Override(container.Injector, func(_ *do.Injector) (config.Tracking, error) { return cfg.Tracking, nil })
	do.Override(container.Injector, func(i *do.Injector) (tracking.ServiceRoot, error) {
		return tracking.ServiceRoot("user-flex-feature"), nil
	})
	srv := do.MustInvoke[*grpc.Server](container.Injector)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.Port))
	if err != nil {
		err := errors.Errorf("failed to listen: %v", err)
		panic(err)
	}
	go func() {
		slog.Info(fmt.Sprintf("Start grpc server: :%d", cfg.Grpc.Port))
		if err := srv.Serve(lis); err != nil && err != grpc.ErrServerStopped {
			slog.With(logging.WithStack(err)).Error("Shutdown.")
		}
	}()
	if cfg.Gateway != nil {
		do.Override(container.Injector, func(_ *do.Injector) (*config.Gateway, error) { return cfg.Gateway, nil })
		srv := do.MustInvokeNamed[*http.Server](container.Injector, "gateway")
		lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Gateway.Port))
		if err != nil {
			err := errors.Errorf("failed to listen: %v", err)
			panic(err)
		}
		go func() {
			if err := srv.Serve(lis); err != nil && err != http.ErrServerClosed {
				slog.With(logging.WithStack(err)).Error("Gateway Shutdown.")
			}
		}()
		go func() {
			<-ctx.Done()
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			_ = container.Injector.Shutdown()
			if err := srv.Shutdown(ctx); err != nil {
				slog.With(logging.WithStack(err)).Error("Gateway forced to shutdown")
			}
			slog.Info("Gateway exiting")
		}()
	}

	<-ctx.Done()
	stop()
	slog.Info("shutting down gracefully, press Ctrl+C again to force")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.GracefulStop()
	_ = container.Injector.Shutdown()
	slog.Info("Server exiting")
}
