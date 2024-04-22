package gateway

import (
	"context"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

type Endpoint func(ctx context.Context, mux *runtime.ServeMux, conn *grpc.ClientConn) error

type option struct {
	runtimeOption  []runtime.ServeMuxOption
	endpoints      []Endpoint
	headerMatchers []runtime.HeaderMatcherFunc
}

type Option interface {
	apply(opt *option)
}

type optionFn func(opt *option)

func (fn optionFn) apply(opt *option) {
	fn(opt)
}

func WithRuntimeSeverMuxOptions(opts ...runtime.ServeMuxOption) Option {
	return optionFn(func(opt *option) {
		opt.runtimeOption = append(opt.runtimeOption, opts...)
	})
}

func WithGrpcEndpoint(endpoints ...Endpoint) Option {
	return optionFn(func(opt *option) {
		opt.endpoints = append(opt.endpoints, endpoints...)
	})
}

func WithHeaderMatcher(fn runtime.HeaderMatcherFunc) Option {
	return optionFn(func(opt *option) {
		opt.headerMatchers = append(opt.headerMatchers, fn)
	})
}
