package feature

import (
	"context"
	"fmt"
	"sync"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	domain_feature "github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/feature/retriever"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/storage/cache"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/storage/serialize"
	"github.com/nonchan7720/user-flex-feature/pkg/services/feature/internal"
	"github.com/samber/do"
	"github.com/thomaspoignant/go-feature-flag/model"
	ff_retriever "github.com/thomaspoignant/go-feature-flag/retriever"
)

func init() {
	do.Provide(container.Injector, Provide)
}

type serviceImpl struct {
	cfg        *config.Config
	ff         *feature.Client
	retrievers []ff_retriever.Retriever
	cache      cache.Cache[model.RawVarResult]
}

func Provide(i *do.Injector) (Service, error) {
	cfg := do.MustInvoke[*config.Config](i)
	client := do.MustInvoke[*feature.Client](i)
	retrievers := do.MustInvoke[[]ff_retriever.Retriever](i)
	cacheBackend := do.MustInvoke[cache.Backend](i)
	cache := cache.NewCache(cacheBackend, &serialize.JsonSerializer[model.RawVarResult]{})
	return newService(cfg, retrievers, client, cache), nil
}

func newService(cfg *config.Config, retrievers []ff_retriever.Retriever, ff *feature.Client, cache cache.Cache[model.RawVarResult]) *serviceImpl {
	return &serviceImpl{
		cfg:        cfg,
		ff:         ff,
		retrievers: retrievers,
		cache:      cache,
	}
}

func (impl *serviceImpl) AppendOrUpdateRule(ctx context.Context, key string, rule *domain_feature.Rule) error {
	if err := validator.Validate(&key, validation.Required); err != nil {
		return err
	}
	if err := rule.Validate(); err != nil {
		return err
	}
	updaters := retriever.FindUpdateRetriever(impl.retrievers...)
	wg := sync.WaitGroup{}
	errCh := make(chan error)
	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	for _, updater := range updaters {
		if updater.CanUpdate(ctx) {
			wg.Add(1)
			go func(ctx context.Context, updater retriever.UpdateRetriever, key string, rule *domain_feature.Rule) {
				defer wg.Done()
				if err := updater.AppendOrUpdateRule(ctx, key, rule); err != nil {
					errCh <- err
				}
			}(ctx, updater, key, rule)
		}
	}

	errs := make([]error, 0, len(updaters))
flagLoop:
	for {
		select {
		case <-done:
			close(errCh)
			break flagLoop
		case err, ok := <-errCh:
			if ok {
				if err != nil {
					errs = append(errs, err)
				}
			}
		}
	}
	for _, err := range errs {
		if err != nil {
			return err
		}
	}
	return impl.ff.Reset()
}

func (impl *serviceImpl) evaluateFlag(key string, evalCtx domain_feature.Context) (*model.RawVarResult, error) {
	defaultValue := "thisisadefaultvaluethatItest1233%%"
	val, _ := impl.ff.RawVariation(key, evalCtx, defaultValue)
	if val.Reason == internal.ReasonError {
		msg := fmt.Sprintf("Error while evaluating the flag: %s", key)
		return nil, domain_feature.NewGeneralError(val.ErrorCode, msg, key)
	}
	return &val, nil
}

func (impl *serviceImpl) EvaluateFlag(ctx context.Context, key string, evalCtx domain_feature.Context) (*model.RawVarResult, error) {
	val, _ := impl.evaluateFlag(key, evalCtx)
	if val.Reason == internal.ReasonError {
		msg := fmt.Sprintf("Error while evaluating the flag: %s", key)
		return nil, domain_feature.NewGeneralError(val.ErrorCode, msg, key)
	}
	return val, nil
}

func (impl *serviceImpl) EvaluateFlagsBulk(ctx context.Context, evalCtx domain_feature.Context) map[string]domain_feature.FlagState {
	allFlags := impl.ff.AllFlagsState(evalCtx)
	flags := allFlags.GetFlags()
	results := make(map[string]domain_feature.FlagState)
	for key, val := range flags {
		value := val.Value
		if val.Reason == internal.ReasonError {
			value = nil
		}
		results[key] = domain_feature.FlagState{
			Value:         value,
			Timestamp:     val.Timestamp,
			VariationType: val.VariationType,
			TrackEvents:   val.TrackEvents,
			Failed:        val.Failed,
			ErrorCode:     val.ErrorCode,
			Reason:        val.Reason,
			Metadata:      val.Metadata,
		}
	}
	return results
}
