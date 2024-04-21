package retriever

import (
	"context"

	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/thomaspoignant/go-feature-flag/retriever"
)

type UpdateRetriever interface {
	CanUpdate(ctx context.Context) bool
	AppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error
	GetVariations(ctx context.Context, key string) []string
}

func FindUpdateRetriever(retrievers ...retriever.Retriever) []UpdateRetriever {
	results := make([]UpdateRetriever, 0, len(retrievers))
	for _, retriever := range retrievers {
		if v, ok := retriever.(UpdateRetriever); ok {
			results = append(results, v)
		}
	}
	return results
}
