package raft

import (
	transport "github.com/Jille/raft-grpc-transport"
	"github.com/hashicorp/raft"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/samber/do"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	do.Provide(container.Injector, ProvideTransport)
}

func ProvideTransport(i *do.Injector) (*transport.Manager, error) {
	cfg := do.MustInvoke[*config.Config](i)
	return newTransport(cfg), nil
}

func newTransport(cfg *config.Config) *transport.Manager {
	return transport.New(raft.ServerAddress(cfg.Grpc.Endpoint()), []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	})
}
