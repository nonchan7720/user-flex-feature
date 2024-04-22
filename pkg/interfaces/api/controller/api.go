package controller

import (
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/interfaces/grpc/ofrep"
	"github.com/samber/do"
	"google.golang.org/grpc"
)

func init() {
	do.Provide(container.Injector, ProvideAPI)
}

type API interface {
	ofrepAPI
}

type api struct {
	ff          *feature.Client
	cfg         *config.Config
	ofrepClient ofrep.OFREPServiceClient
}

func ProvideAPI(i *do.Injector) (API, error) {
	ff := do.MustInvoke[*feature.Client](i)
	cfg := do.MustInvoke[*config.Config](i)
	conn := do.MustInvokeNamed[*grpc.ClientConn](i, "user-flex-feature-grpc")
	ofrepClient := ofrep.NewOFREPServiceClient(conn)
	return newAPI(ff, cfg, ofrepClient), nil
}

func newAPI(ff *feature.Client, cfg *config.Config, ofrepClient ofrep.OFREPServiceClient) *api {
	return &api{
		ff:          ff,
		cfg:         cfg,
		ofrepClient: ofrepClient,
	}
}
