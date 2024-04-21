package gateway

import (
	"context"
	"math"
	"strings"
	"time"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/grpc/interceptor"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/keepalive"
)

func NewGrpcGateway(ctx context.Context, cfg *config.Gateway, opts ...Option) *runtime.ServeMux {
	opt := &option{
		headerMatchers: []runtime.HeaderMatcherFunc{matcher},
	}
	for _, o := range opts {
		o.apply(opt)
	}
	opt.runtimeOption = append(opt.runtimeOption, runtime.WithIncomingHeaderMatcher(wrapMatchers(opt.headerMatchers)))

	addr := cfg.Grpc.Endpoint()
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
			interceptor.AuthUnaryClientInterceptor(cfg.Grpc.Auth),
		),
	}
	mux := runtime.NewServeMux(
		opt.runtimeOption...,
	)

	for _, endpoint := range opt.endpoints {
		if err := endpoint(ctx, mux, addr, dialOpts); err != nil {
			panic(err)
		}
	}
	return mux
}

func matcher(key string) (string, bool) {
	if strings.HasPrefix(strings.ToLower(key), "x-") {
		return key, true
	}
	return "", false
}

func wrapMatchers(funcs []runtime.HeaderMatcherFunc) runtime.HeaderMatcherFunc {
	return func(s string) (string, bool) {
		for _, fn := range funcs {
			if _, ok := fn(s); ok {
				return s, ok
			}
		}
		return "", false
	}
}
