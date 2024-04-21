package interceptor

import (
	"context"

	grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func AuthUnaryServerInterceptor(cfg *config.Auth) grpc.UnaryServerInterceptor {
	return grpc_auth.UnaryServerInterceptor(func(ctx context.Context) (context.Context, error) {
		if cfg == nil {
			return ctx, nil
		}
		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "There is no METADATA.")
		}
		if !cfg.Valid(md) {
			return nil, status.Error(codes.Unauthenticated, "Authentication failed.")
		}
		return ctx, nil
	})
}

func AuthUnaryClientInterceptor(cfg *config.Auth) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if cfg != nil {
			if token := cfg.Token(); token != "" {
				ctx = metadata.AppendToOutgoingContext(ctx, "authorization", token)
			}
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
