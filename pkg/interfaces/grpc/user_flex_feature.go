package grpc

import (
	"log/slog"
	"math"
	"net"
	"time"

	"github.com/Jille/raft-grpc-leader-rpc/leaderhealth"
	transport "github.com/Jille/raft-grpc-transport"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/grpc/interceptor"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/logging"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/ofrep"
	user_flex_feature "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature"
	user_flex_feature_raft "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature-raft"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/raft"
	svc_feature "github.com/nonchan7720/user-flex-feature/pkg/services/feature"
	"github.com/samber/do"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	healthpb "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/keepalive"
	"google.golang.org/grpc/reflection"
)

type UserFlexFeatureServer struct {
	*grpc.Server

	svcSrv ServiceServer
}

func (s *UserFlexFeatureServer) Serve(lis net.Listener, isClustering bool) {
	go func() {
		if isClustering {
			if err := s.svcSrv.JoinCluster(); err != nil {
				panic(err)
			}
		}
	}()
	if err := s.Server.Serve(lis); err != nil && err != grpc.ErrServerStopped {
		slog.With(logging.WithStack(err)).Error("Shutdown.")
	}
}

func init() {
	do.Provide(container.Injector, ProvideUserFlexFeatureServer)
	do.Provide(container.Injector, newUserFlexFeatureGrpcServer)
}

func newUserFlexFeatureGrpcServer(i *do.Injector) (*UserFlexFeatureServer, error) {
	cfg := do.MustInvoke[*config.Config](i)
	srv := do.MustInvoke[ServiceServer](i)
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
	ofrep.RegisterOFREPServiceServer(s, srv)
	user_flex_feature_raft.RegisterRaftServiceServer(s, srv)
	hs := health.NewServer()
	healthpb.RegisterHealthServer(s, hs)
	if cfg.IsRaftCluster() {
		raft := do.MustInvoke[*raft.Raft](i)
		tm := do.MustInvoke[*transport.Manager](i)
		tm.Register(s)
		leaderhealth.Report(raft.Raft, hs, []string{"UserFlexFeature"})
	}
	return &UserFlexFeatureServer{
		Server: s,
		svcSrv: srv,
	}, nil
}

func ProvideUserFlexFeatureServer(i *do.Injector) (ServiceServer, error) {
	svc := do.MustInvoke[svc_feature.Service](i)
	ff := do.MustInvoke[feature.Feature](i)
	cfg := do.MustInvoke[*config.Config](i)
	var vRaft *raft.Raft
	if cfg.IsRaftCluster() {
		vRaft = do.MustInvoke[*raft.Raft](i)
	}
	return newUserFlexFeatureServer(svc, ff, vRaft, cfg), nil
}
