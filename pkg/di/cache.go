package di

import (
	"time"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/storage/cache"
	"github.com/samber/do"
)

func init() {
	do.Provide(container.Injector, cacheBackendDefault)
}

func cacheBackendDefault(i *do.Injector) (cache.Backend, error) {
	return cache.NewInMemory(1 * time.Minute), nil
}
