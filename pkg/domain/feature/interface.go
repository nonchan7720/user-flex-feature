package feature

import (
	"context"

	ffclient "github.com/thomaspoignant/go-feature-flag"
)

type Feature interface {
	Provider() *ffclient.GoFeatureFlag
	Shutdown() error
	Reset() error
	AppendOrUpdateRule(ctx context.Context, key string, rule *Rule) error
}

type UpdateRetriever interface {
	CanUpdate(ctx context.Context) bool
	AppendOrUpdateRule(ctx context.Context, key string, rule *Rule) error
}

type Updater interface {
	AppendOrUpdateRule(ctx context.Context, updateRetrievers []UpdateRetriever, key string, rule *Rule) error
}
