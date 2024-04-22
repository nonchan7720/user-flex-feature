package feature

import (
	"context"
	"sync"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/feature/retriever"
	"github.com/nonchan7720/user-flex-feature/pkg/utils/collection"
	"github.com/samber/do"
	ff_retriever "github.com/thomaspoignant/go-feature-flag/retriever"
)

func init() {
	do.Provide(container.Injector, Provide)
}

type serviceImpl struct {
	retrievers []ff_retriever.Retriever
}

func Provide(i *do.Injector) (Service, error) {
	retrievers := do.MustInvoke[[]ff_retriever.Retriever](i)
	return newService(retrievers), nil
}

func newService(retrievers []ff_retriever.Retriever) *serviceImpl {
	return &serviceImpl{
		retrievers: retrievers,
	}
}

func (impl *serviceImpl) AppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error {
	if err := validator.Validate(&key, validation.Required); err != nil {
		return err
	}
	wg := sync.WaitGroup{}
	updaters := retriever.FindUpdateRetriever(impl.retrievers...)
	variationsCh := make(chan []string)
	go func() {
		wg.Wait()
		close(variationsCh)
	}()
	for _, updater := range updaters {
		wg.Add(1)
		go func(ctx context.Context, updater retriever.UpdateRetriever, key string) {
			defer wg.Done()
			variationsCh <- updater.GetVariations(ctx, key)
		}(ctx, updater, key)
	}
	var variations []string
	for variation := range variationsCh {
		variations = append(variations, variation...)
	}
	variations = collection.Uniq(variations)
	if err := rule.Validate(variations); err != nil {
		return nil
	}

	wg = sync.WaitGroup{}
	errCh := make(chan error)
	go func() {
		wg.Wait()
		close(errCh)
	}()
	for _, updater := range updaters {
		if updater.CanUpdate(ctx) {
			wg.Add(1)
			go func(ctx context.Context, updater retriever.UpdateRetriever, key string, rule *feature.Rule) {
				defer wg.Done()
				if err := updater.AppendOrUpdateRule(ctx, key, rule); err != nil {
					errCh <- err
				}
			}(ctx, updater, key, rule)
		}
	}
	for err := range errCh {
		return err
	}
	return nil
}
