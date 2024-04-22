package cmd

import (
	"context"
	"fmt"
	"log/slog"
	"net"
	"os/signal"
	"syscall"
	"time"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/tracking"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc"
	"github.com/pkg/errors"
	"github.com/samber/do"
	"github.com/spf13/cobra"
)

func serverCommand() *cobra.Command {
	var serverArgs commonArgs
	cmd := &cobra.Command{
		Use:           "server",
		SilenceErrors: true,
		SilenceUsage:  true,
		Run: func(cmd *cobra.Command, args []string) {
			serverExecute(cmd.Context(), &serverArgs)
		},
	}
	flag := cmd.Flags()
	flag.StringVarP(&serverArgs.configFilePath, "config", "c", "config.yaml", "configuration file path")
	return cmd
}

func serverExecute(ctx context.Context, args *commonArgs) {
	if err := args.Validate(); err != nil {
		panic(err)
	}

	ctx, stop := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM)
	defer stop()
	cfg := loadConfig[config.Config](args.configFilePath)
	do.Override(container.Injector, func(i *do.Injector) (context.Context, error) { return ctx, nil })
	do.Override(container.Injector, func(_ *do.Injector) (*config.Config, error) { return cfg, nil })
	do.Override(container.Injector, func(_ *do.Injector) (config.Tracking, error) { return cfg.Tracking, nil })
	do.Override(container.Injector, func(i *do.Injector) (tracking.ServiceRoot, error) {
		return tracking.ServiceRoot("user-flex-feature"), nil
	})
	srv := do.MustInvoke[*grpc.UserFlexFeatureServer](container.Injector)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Grpc.Port))
	if err != nil {
		err := errors.Errorf("failed to listen: %v", err)
		panic(err)
	}
	go func() {
		slog.Info(fmt.Sprintf("Start grpc server: :%d", cfg.Grpc.Port))
		srv.Serve(lis)
	}()
	var gatewayShutdown GatewayShutdown
	if cfg.Gateway != nil {
		gatewayShutdown = runGatewayServer(cfg.Gateway)
	}

	<-ctx.Done()
	stop()
	slog.Info("shutting down gracefully, press Ctrl+C again to force")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	srv.GracefulStop()
	if gatewayShutdown != nil {
		gatewayShutdown(ctx)
	}
	_ = container.Injector.Shutdown()
	slog.Info("Server exiting")
}
