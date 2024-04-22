package grpc

import (
	"context"
	"fmt"
	"log/slog"
	"math"
	"time"

	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/grpc/interceptor"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewConnection(ctx context.Context, cfg *config.Grpc) (*grpc.ClientConn, error) {
	endpoint := cfg.Endpoint()
	creds := cfg.GrpcCredentials()
	dialOpts := []grpc.DialOption{
		grpc.WithDefaultCallOptions(
			grpc.MaxCallSendMsgSize(math.MaxInt64),
			grpc.MaxCallRecvMsgSize(math.MaxInt64),
		),
		grpc.WithKeepaliveParams(
			keepalive.ClientParameters{
				Time:                5 * time.Minute,
				Timeout:             10 * time.Second,
				PermitWithoutStream: true,
			},
		),
		grpc.WithStatsHandler(otelgrpc.NewClientHandler()),
		grpc.WithTransportCredentials(creds),
		grpc.WithChainUnaryInterceptor(
			interceptor.AuthUnaryClientInterceptor(cfg.Auth),
		),
	}
	conn, err := grpc.DialContext(ctx, endpoint, dialOpts...)
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
