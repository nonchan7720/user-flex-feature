package retriever

import (
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/retriever"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
)

type fileRetriever struct {
	*fileretriever.Retriever
}

func (f *fileRetriever) CanUpdate() bool {
	return false
}

func newFileRetriever(cfg *retriever.File) *fileRetriever {
	return &fileRetriever{
		Retriever: &fileretriever.Retriever{
			Path: cfg.Path,
		},
	}
}
