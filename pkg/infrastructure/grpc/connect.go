package grpc

import (
	"context"
	"math"
	"time"

	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/grpc/interceptor"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

func NewGrpcConnection(ctx context.Context, endpoint string, creds credentials.TransportCredentials, auth *config.Auth) (*grpc.ClientConn, error) {
	if creds == nil {
		creds = insecure.NewCredentials()
	}
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
			interceptor.AuthUnaryClientInterceptor(auth),
		),
	}
	conn, err := grpc.DialContext(ctx, endpoint, dialOpts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
