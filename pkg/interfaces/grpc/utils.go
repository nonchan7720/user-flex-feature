package grpc

import (
	"context"
	"fmt"

	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/grpc"
	user_flex_feature "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature"
	user_flex_feature_raft "github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/user-flex-feature-raft"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/raft"
)

func generateServiceName[T any]() string {
	var t T

	// struct
	name := fmt.Sprintf("%T", t)
	if name != "<nil>" {
		return name
	}

	// interface
	return fmt.Sprintf("%T", new(T))
}

func leader[T any](ctx context.Context, raft *raft.Raft) (T, error) {
	var client T
	leaderAddr, _ := raft.LeaderWithID()
	conn, err := grpc.NewGrpcConnection(ctx, string(leaderAddr), nil, nil)
	if err != nil {
		return client, err
	}
	go func() {
		<-ctx.Done()
		_ = conn.Close()
	}()
	name := generateServiceName[T]()
	switch name {
	case "*user_flex_feature_raft.RaftServiceClient":
		client = user_flex_feature_raft.NewRaftServiceClient(conn).(T)
	case "*user_flex_feature.UserFlexFeatureServiceClient":
		client = user_flex_feature.NewUserFlexFeatureServiceClient(conn).(T)
	}
	return client, nil
}
