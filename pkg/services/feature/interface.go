package feature

import (
	"context"

	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/thomaspoignant/go-feature-flag/model"
)

type Service interface {
	AppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error
	EvaluateFlag(ctx context.Context, key string, evalCtx feature.Context) (*model.RawVarResult, error)
	EvaluateFlagsBulk(ctx context.Context, evalCtx feature.Context) map[string]feature.FlagState
}
