package retriever

import (
	"fmt"

	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/retriever"
	"github.com/samber/do"
	ff_retriever "github.com/thomaspoignant/go-feature-flag/retriever"
)

func init() {
	do.Provide(container.Injector, Provide)
	do.Provide(container.Injector, ProvideUpdateRetrievers)
}

func Provide(i *do.Injector) ([]ff_retriever.Retriever, error) {
	cfg := do.MustInvoke[*config.Config](i)
	return New(cfg)
}

func ProvideUpdateRetrievers(i *do.Injector) ([]feature.UpdateRetriever, error) {
	retrievers := do.MustInvoke[[]ff_retriever.Retriever](i)
	return FindUpdateRetriever(retrievers...), nil
}

func newRetriever(r retriever.Retriever) (ff_retriever.Retriever, error) {
	switch r.Type {
	case retriever.FileType:
		return newFileRetriever(r.File), nil
	case retriever.InMemoryType:
		return newInMemory(r.InMemory), nil
	default:
		return nil, fmt.Errorf("Un supported type: %s", r.Type)
	}
}

func New(cfg *config.Config) ([]ff_retriever.Retriever, error) {
	var retrievers []ff_retriever.Retriever
	for _, r := range cfg.Retrievers {
		if v, err := newRetriever(r); err != nil {
			return nil, err
		} else {
			retrievers = append(retrievers, v)
		}
	}
	return retrievers, nil
}
