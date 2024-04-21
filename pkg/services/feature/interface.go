package feature

import (
	"context"

	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
)

type Service interface {
	AppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error
}
