package interceptor

import (
	"context"
	"fmt"
	"log/slog"

	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/logging"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func recoveryHandlerContext(ctx context.Context, p interface{}) (err error) {
	slog.With(logging.WithStack(p.(error))).ErrorContext(ctx, fmt.Sprintf("%+v\n", p.(error)))
	return status.Errorf(codes.Internal, "Unexpected error")
}

func RecoveryInterceptor() grpc.UnaryServerInterceptor {
	return grpc_recovery.UnaryServerInterceptor(
		grpc_recovery.WithRecoveryHandlerContext(recoveryHandlerContext),
	)
}
