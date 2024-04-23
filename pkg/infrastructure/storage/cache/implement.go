package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/storage/redis"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/storage/serialize"
)

func NewBackend(cfg *config.Cache) Backend {
	switch cfg.Type {
	case config.CacheInMemory:
		return NewInMemory(cfg.InMemory.CleanUp)
	case config.CacheRedis:
		cmd, err := redis.NewRedis(context.Background(), cfg.Redis)
		if err != nil {
			panic(err)
		}
		return NewRedisCache(cmd)
	default:
		panic(fmt.Sprintf("Un support cache type: %s", cfg.Type))
	}
}

func NewCache[V any](backend Backend, serializer serialize.Serialize[V]) Cache[V] {
	return &impl[V]{
		backend:    backend,
		serializer: serializer,
	}
}

type impl[V any] struct {
	backend    Backend
	serializer serialize.Serialize[V]
}

func (c *impl[V]) Get(ctx context.Context, key string) (*V, error) {
	buf, err := c.backend.Get(ctx, key)
	if err != nil {
		return nil, err
	}
	if v, err := c.serializer.Decode(ctx, buf); err == nil {
		return &v, nil
	} else {
		return nil, err
	}
}

func (c *impl[V]) Set(ctx context.Context, key string, value *V) error {
	return c.SetTTL(ctx, key, value, 0)
}

func (c *impl[V]) SetTTL(ctx context.Context, key string, value *V, exp time.Duration) error {
	val, err := c.serializer.Encode(ctx, *value)
	if err != nil {
		return err
	}
	return c.backend.SetTTL(ctx, key, val, exp)
}

func (c *impl[V]) Del(ctx context.Context, key string) error {
	return c.backend.Del(ctx, key)
}
