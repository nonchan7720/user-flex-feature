package controller

import (
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/samber/do"
	ffclient "github.com/thomaspoignant/go-feature-flag"
)

func init() {
	do.Provide(container.Injector, ProvideAPI)
}

type API interface {
	ofrepAPI
}

type api struct {
	ff  *ffclient.GoFeatureFlag
	cfg *config.Config
}

func ProvideAPI(i *do.Injector) (API, error) {
	ff := do.MustInvoke[*ffclient.GoFeatureFlag](i)
	cfg := do.MustInvoke[*config.Config](i)
	return newAPI(ff, cfg), nil
}

func newAPI(ff *ffclient.GoFeatureFlag, cfg *config.Config) *api {
	return &api{
		ff:  ff,
		cfg: cfg,
	}
}
