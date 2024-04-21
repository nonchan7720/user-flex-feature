package retriever

import (
	"fmt"

	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/retriever"
	ff_retriever "github.com/thomaspoignant/go-feature-flag/retriever"
)

func New(cfg *config.Config) ([]ff_retriever.Retriever, error) {
	var retrievers []ff_retriever.Retriever
	for _, r := range cfg.Retrievers {
		switch r.Type {
		case retriever.FileType:
			retrievers = append(retrievers, newFileRetriever(r.File))
		case retriever.InMemoryType:
			retrievers = append(retrievers, newInMemory(r.InMemory))
		default:
			return nil, fmt.Errorf("Un supported type: %s", r.Type)
		}
	}
	return retrievers, nil
}
