package retriever

import (
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/thomaspoignant/go-feature-flag/retriever"
)

func FindUpdateRetriever(retrievers ...retriever.Retriever) []feature.UpdateRetriever {
	results := make([]feature.UpdateRetriever, 0, len(retrievers))
	for _, retriever := range retrievers {
		if v, ok := retriever.(feature.UpdateRetriever); ok {
			results = append(results, v)
		}
	}
	return results
}
