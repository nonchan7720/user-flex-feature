package gateway

import (
	"context"
	"strings"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Gateway struct {
	*runtime.ServeMux
	conn *grpc.ClientConn
}

func (g *Gateway) Conn() *grpc.ClientConn {
	return g.conn
}

func NewGrpcGateway(ctx context.Context, conn *grpc.ClientConn, opts ...Option) *Gateway {
	opt := &option{
		headerMatchers: []runtime.HeaderMatcherFunc{matcher},
	}
	for _, o := range opts {
		o.apply(opt)
	}
	opt.runtimeOption = append(opt.runtimeOption, runtime.WithIncomingHeaderMatcher(wrapMatchers(opt.headerMatchers)))

	mux := runtime.NewServeMux(
		opt.runtimeOption...,
	)

	for _, endpoint := range opt.endpoints {
		if err := endpoint(ctx, mux, conn); err != nil {
			panic(err)
		}
	}
	return &Gateway{
		ServeMux: mux,
		conn:     conn,
	}
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
