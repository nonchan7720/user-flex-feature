package retriever

import (
	"context"
	"os"
	"sync"

	"github.com/goccy/go-yaml"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/retriever"
	"github.com/thomaspoignant/go-feature-flag/retriever/fileretriever"
)

type fileRetriever struct {
	mu sync.Locker
	*fileretriever.Retriever
}

var (
	_ UpdateRetriever = (*fileRetriever)(nil)
)

func (f *fileRetriever) CanUpdate(_ context.Context) bool {
	return true
}

func (f *fileRetriever) muter(fn func(flags feature.Flags) error) error {
	f.mu.Lock()
	defer f.mu.Unlock()
	buf, err := os.ReadFile(f.Path)
	if err != nil {
		return err
	}
	var flags feature.Flags
	if err := yaml.Unmarshal(buf, &flags); err != nil {
		return err
	}
	return fn(flags)
}

func (f *fileRetriever) AppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error {
	return f.muter(func(flags feature.Flags) error {
		v, ok := flags[key]
		if !ok {
			return nil
		}
		v.AppendOrUpdateRule(rule)
		f, err := os.Create(f.Path)
		if err != nil {
			return err
		}
		defer f.Close()
		if err := yaml.NewEncoder(f).Encode(&flags); err != nil {
			return err
		}
		return nil
	})
}

func (f *fileRetriever) Retrieve(_ context.Context) ([]byte, error) {
	f.mu.Lock()
	defer f.mu.Unlock()

	return os.ReadFile(f.Path)
}

func newFileRetriever(cfg *retriever.File) *fileRetriever {
	return &fileRetriever{
		mu: &sync.Mutex{},
		Retriever: &fileretriever.Retriever{
			Path: cfg.Path,
		},
	}
}
