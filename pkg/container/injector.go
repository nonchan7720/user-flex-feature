package container

import (
	"context"
	"fmt"

	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/logging"
	"github.com/samber/do"
)

var Injector *do.Injector = do.NewWithOpts(&do.InjectorOpts{
	Logf: func(format string, args ...any) {
		logging.Info(fmt.Sprintf(format, args...))
	},
})

func init() {
	do.Provide(Injector, func(i *do.Injector) (context.Context, error) { return context.Background(), nil })
}
