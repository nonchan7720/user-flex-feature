package grpc

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	inf_grpc "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/grpc"
	"github.com/samber/do"
	"google.golang.org/grpc"
)

func init() {
	do.ProvideNamed(container.Injector, "user-flex-feature-grpc", ProvideGrpcClientConn)
}

func ProvideGrpcClientConn(i *do.Injector) (*grpc.ClientConn, error) {
	ctx := do.MustInvoke[context.Context](i)
	cfg := do.MustInvoke[*config.Config](i)
	return newGrpcConnection(ctx, &cfg.Grpc)
}

func newGrpcConnection(ctx context.Context, cfg *config.Grpc) (*grpc.ClientConn, error) {
	endpoint := cfg.Endpoint()
	conn, err := inf_grpc.NewGrpcConnection(ctx, endpoint, cfg.GrpcCredentials(), cfg.Auth)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err != nil {
			if cerr := conn.Close(); cerr != nil {
				slog.InfoContext(ctx, fmt.Sprintf("Failed to close conn to %s: %v", endpoint, cerr))
			}
			return
		}
		go func() {
			<-ctx.Done()
			if cerr := conn.Close(); cerr != nil {
				slog.InfoContext(ctx, fmt.Sprintf("Failed to close conn to %s: %v", endpoint, cerr))
			}
		}()
	}()
	return conn, nil
}
