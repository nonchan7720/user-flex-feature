package di

import (
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/samber/do"
)

func init() {
	do.Provide(container.Injector, featureNewUpdater)
}

func featureNewUpdater(i *do.Injector) (feature.Updater, error) {
	return feature.NewUpdater(), nil
}
