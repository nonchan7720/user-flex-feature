package feature

import (
	"context"
	"fmt"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/storage/cache"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/storage/serialize"
	"github.com/nonchan7720/user-flex-feature/pkg/services/feature/internal"
	"github.com/samber/do"
	"github.com/thomaspoignant/go-feature-flag/model"
)

func init() {
	do.Provide(container.Injector, Provide)
}

type serviceImpl struct {
	cfg   *config.Config
	ff    feature.Feature
	cache cache.Cache[model.RawVarResult]
}

func Provide(i *do.Injector) (Service, error) {
	cfg := do.MustInvoke[*config.Config](i)
	client := do.MustInvoke[feature.Feature](i)
	cacheBackend := do.MustInvoke[cache.Backend](i)
	cache := cache.NewCache(cacheBackend, &serialize.JsonSerializer[model.RawVarResult]{})
	return newService(cfg, client, cache), nil
}

func newService(cfg *config.Config, ff feature.Feature, cache cache.Cache[model.RawVarResult]) *serviceImpl {
	return &serviceImpl{
		cfg:   cfg,
		ff:    ff,
		cache: cache,
	}
}

func (impl *serviceImpl) AppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error {
	if err := validator.Validate(&key, validation.Required); err != nil {
		return err
	}
	if err := rule.Validate(); err != nil {
		return err
	}

	if err := impl.ff.AppendOrUpdateRule(ctx, key, rule); err != nil {
		return err
	}
	return impl.ff.Reset()
}

func (impl *serviceImpl) evaluateFlag(key string, evalCtx feature.Context) (*model.RawVarResult, error) {
	defaultValue := "thisisadefaultvaluethatItest1233%%"
	val, _ := impl.ff.Provider().RawVariation(key, evalCtx, defaultValue)
	if val.Reason == internal.ReasonError {
		msg := fmt.Sprintf("Error while evaluating the flag: %s", key)
		return nil, feature.NewGeneralError(val.ErrorCode, msg, key)
	}
	return &val, nil
}

func (impl *serviceImpl) EvaluateFlag(ctx context.Context, key string, evalCtx feature.Context) (*model.RawVarResult, error) {
	val, _ := impl.evaluateFlag(key, evalCtx)
	if val.Reason == internal.ReasonError {
		msg := fmt.Sprintf("Error while evaluating the flag: %s", key)
		return nil, feature.NewGeneralError(val.ErrorCode, msg, key)
	}
	return val, nil
}

func (impl *serviceImpl) EvaluateFlagsBulk(ctx context.Context, evalCtx feature.Context) map[string]feature.FlagState {
	allFlags := impl.ff.Provider().AllFlagsState(evalCtx)
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
