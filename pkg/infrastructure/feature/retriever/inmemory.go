package retriever

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/nonchan7720/user-flex-feature/pkg/domain/errors"
	"github.com/nonchan7720/user-flex-feature/pkg/domain/feature"
	config "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/inmemory"
	"github.com/thomaspoignant/go-feature-flag/retriever"
)

type inMemory struct {
	mp  sync.Map
	cfg *config.InMemory

	status retriever.Status
}

var (
	_ retriever.InitializableRetriever = (*inMemory)(nil)
	_ UpdateRetriever                  = (*inMemory)(nil)
)

func newInMemory(cfg *config.InMemory) *inMemory {
	return &inMemory{
		mp:     sync.Map{},
		cfg:    cfg,
		status: retriever.RetrieverNotReady,
	}
}

func (r *inMemory) Retrieve(ctx context.Context) ([]byte, error) {
	m := map[string]interface{}{}
	r.mp.Range(func(key, value any) bool {
		m[fmt.Sprint(key)] = value
		return true
	})
	return json.Marshal(&m)
}

func (r *inMemory) Init(ctx context.Context, logger *log.Logger) error {
	if r.cfg == nil || r.cfg.FilePath == "" {
		r.status = retriever.RetrieverReady
		return nil
	}
	buf, err := os.ReadFile(r.cfg.FilePath)
	if err != nil {
		r.status = retriever.RetrieverError
		return err
	}
	flags, err := ConvertToFlagStruct(buf)
	if err != nil {
		r.status = retriever.RetrieverError
		return err
	}
	for key, flag := range flags {
		r.mp.Store(key, flag)
	}
	r.status = retriever.RetrieverReady
	return nil
}

func (r *inMemory) Shutdown(ctx context.Context) error {
	r.status = retriever.RetrieverNotReady
	r.mp = sync.Map{}
	return nil
}

func (r *inMemory) Status() retriever.Status {
	return r.status
}

func (r *inMemory) CanUpdate(_ context.Context) bool {
	return r.status == retriever.RetrieverReady
}

func (r *inMemory) AppendOrUpdateRule(ctx context.Context, key string, rule *feature.Rule) error {
	value, ok := r.mp.Load(key)
	if !ok {
		return errors.ErrNotfound
	}
	nowFlag, ok := value.(*feature.Flag)
	if !ok {
		return errors.ErrNotfound
	}
	nowFlag.AppendOrUpdateRule(rule)
	r.mp.Store(key, nowFlag)
	return nil
}

func (r *inMemory) GetVariations(ctx context.Context, key string) []string {
	value, ok := r.mp.Load(key)
	if !ok {
		return nil
	}
	nowFlag, ok := value.(*feature.Flag)
	if !ok {
		return nil
	}
	if nowFlag.Variations != nil {
		results := make([]string, len(*nowFlag.Variations))
		for key := range *nowFlag.Variations {
			results = append(results, key)
		}
		return results
	}
	return nil
}
