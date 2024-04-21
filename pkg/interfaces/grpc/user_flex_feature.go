package grpc

import (
	"context"
	"math"
	"time"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/grpc/interceptor"
	user_flex_feature "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature"
	svc_feature "github.com/nonchan7720/user-flex-feature/pkg/services/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/utils"
	"github.com/samber/do"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

func init() {
	do.Provide(container.Injector, ProvideUserFlexFeatureServer)
	do.Provide(container.Injector, newUserFlexFeatureGrpcServer)
}

func newUserFlexFeatureGrpcServer(i *do.Injector) (*grpc.Server, error) {
	cfg := do.MustInvoke[*config.Config](i)
	srv := do.MustInvoke[user_flex_feature.UserFlexFeatureServiceServer](i)
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			interceptor.RecoveryInterceptor(),
			interceptor.AuthUnaryServerInterceptor(cfg.Grpc.Auth),
		),
		grpc.MaxRecvMsgSize(math.MaxInt64),
		grpc.MaxSendMsgSize(math.MaxInt64),
		grpc.KeepaliveParams(
			keepalive.ServerParameters{
				Time:    90 * time.Second,
				Timeout: 60 * time.Second,
			},
		),
		grpc.MaxConcurrentStreams(100),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	}
	s := grpc.NewServer(opts...)
	reflection.Register(s)
	user_flex_feature.RegisterUserFlexFeatureServiceServer(s, srv)
	healthpb.RegisterHealthServer(s, health.NewServer())
	return s, nil
}

func ProvideUserFlexFeatureServer(i *do.Injector) (user_flex_feature.UserFlexFeatureServiceServer, error) {
	svc := do.MustInvoke[svc_feature.Service](i)
	return newUserFlexFeatureServer(svc), nil
}

func newUserFlexFeatureServer(svc svc_feature.Service) *server {
	return &server{
		svc: svc,
	}
}

type server struct {
	svc svc_feature.Service
}

func (s *server) RuleUpdate(ctx context.Context, in *user_flex_feature.RuleUpdateRequest) (*user_flex_feature.RuleUpdateResponse, error) {
	rule := in.GetRule()
	if rule == nil {
		return nil, status.Error(codes.InvalidArgument, "rule is empty.")
	}
	var percentages map[string]float64
	if v := rule.GetPercentageValue(); v != nil {
		percentages = v.GetValue()
	}
	var progressiveRollout *feature.ProgressiveRollout
	if v := rule.GetProgressiveRolloutValue(); v != nil {
		conv := func(v *user_flex_feature.ProgressiveRolloutStep) *feature.ProgressiveRolloutStep {
			if v == nil {
				return nil
			}
			value := &feature.ProgressiveRolloutStep{
				Variation: utils.String(v.GetVariationValue()),
				Date:      utils.ToTime(v.GetDateValue()),
			}
			if v, ok := v.GetPercentage().(*user_flex_feature.ProgressiveRolloutStep_PercentageValue); ok {
				value.Percentage = &v.PercentageValue
			}
			return value
		}
		progressiveRollout = &feature.ProgressiveRollout{
			Initial: conv(v.GetInitial()),
			End:     conv(v.GetEnd()),
		}
	}
	var disable *bool
	if v, ok := rule.GetDisable().(*user_flex_feature.Rule_DisableValue); ok {
		disable = &v.DisableValue
	}
	modelRule := feature.Rule{
		Name:               rule.GetName(),
		Query:              rule.GetQuery(),
		VariationResult:    rule.GetVariationResult(),
		Percentages:        percentages,
		ProgressiveRollout: progressiveRollout,
		Disable:            disable,
	}
	if err := s.svc.AppendOrUpdateRule(ctx, in.GetKey(), &modelRule); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &user_flex_feature.RuleUpdateResponse{
		Result: true,
	}, nil
}
