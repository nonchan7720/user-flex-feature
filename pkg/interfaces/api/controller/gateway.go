package controller

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/api/gateway"
	user_flex_feature "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature"
)

func newUserFlexFeatureGatewayServer(ctx context.Context, cfg *config.Config) *runtime.ServeMux {
	gatewayServer := gateway.NewGrpcGateway(
		ctx,
		cfg.Gateway,
		gateway.WithGrpcEndpoint(user_flex_feature.RegisterUserFlexFeatureServiceHandlerFromEndpoint),
	)
	return gatewayServer
}
