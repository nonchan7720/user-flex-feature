package feature

import (
	"context"
	"fmt"
	"sync"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/feature/retriever"
	"github.com/nonchan7720/user-flex-feature/pkg/services/feature/internal"
	"github.com/samber/do"
	ffclient "github.com/thomaspoignant/go-feature-flag"
	"github.com/thomaspoignant/go-feature-flag/model"
	ff_retriever "github.com/thomaspoignant/go-feature-flag/retriever"
)

func init() {
	do.Provide(container.Injector, Provide)
}

type serviceImpl struct {
	ff         *ffclient.GoFeatureFlag
	retrievers []ff_retriever.Retriever
}

func Provide(i *do.Injector) (Service, error) {
	ff := do.MustInvoke[*ffclient.GoFeatureFlag](i)
	retrievers := do.MustInvoke[[]ff_retriever.Retriever](i)
	return newService(retrievers, ff), nil
}

func newService(retrievers []ff_retriever.Retriever, ff *ffclient.GoFeatureFlag) *serviceImpl {
	return &serviceImpl{
		ff:         ff,
		retrievers: retrievers,
	}
}

func (impl *serviceImpl) AppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error {
	if err := validator.Validate(&key, validation.Required); err != nil {
		return err
	}
	updaters := retriever.FindUpdateRetriever(impl.retrievers...)
	wg := sync.WaitGroup{}
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

func (impl *serviceImpl) EvaluateFlag(ctx context.Context, key string, evalCtx feature.Context) (*model.RawVarResult, error) {
	defaultValue := "thisisadefaultvaluethatItest1233%%"
	val, _ := impl.ff.RawVariation(key, evalCtx, defaultValue)
	if val.Reason == internal.ReasonError {
		msg := fmt.Sprintf("Error while evaluating the flag: %s", key)
		return nil, feature.NewGeneralError(val.ErrorCode, msg, key)
	}
	return &val, nil
}

func (impl *serviceImpl) EvaluateFlagsBulk(ctx context.Context, evalCtx feature.Context) map[string]feature.FlagState {
	allFlags := impl.ff.AllFlagsState(evalCtx)
	flags := allFlags.GetFlags()
	results := make(map[string]feature.FlagState)
	for key, val := range flags {
		value := val.Value
		if val.Reason == internal.ReasonError {
			value = nil
		}
		results[key] = feature.FlagState{
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
