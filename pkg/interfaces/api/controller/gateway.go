package controller

import (
	"context"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/api/gateway"
	user_flex_feature "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature"
	"github.com/samber/do"
	"google.golang.org/grpc"
)

func init() {
	do.ProvideNamed(container.Injector, "user-flex-feature-gateway", ProviderUserFlexFeatureGatewayServer)
}

func ProviderUserFlexFeatureGatewayServer(i *do.Injector) (*gateway.Gateway, error) {
	ctx := do.MustInvoke[context.Context](i)
	conn := do.MustInvokeNamed[*grpc.ClientConn](i, "user-flex-feature-grpc")
	return newUserFlexFeatureGatewayServer(ctx, conn), nil
}

func newUserFlexFeatureGatewayServer(ctx context.Context, conn *grpc.ClientConn) *gateway.Gateway {
	gatewayServer := gateway.NewGrpcGateway(
		ctx,
		conn,
		gateway.WithGrpcEndpoint(user_flex_feature.RegisterUserFlexFeatureServiceHandler),
	)
	return gatewayServer
}
