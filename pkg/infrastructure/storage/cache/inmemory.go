package cache

import (
	"context"
	"time"

	"github.com/nonchan7720/user-flex-feature/pkg/domain/errors"
	"github.com/patrickmn/go-cache"
)

type inmemoryCache struct {
	cache *cache.Cache
}

func NewInMemory(cleanUp time.Duration) Backend {
	return &inmemoryCache{
		cache: cache.New(0, cleanUp),
	}
}

func (c *inmemoryCache) Get(ctx context.Context, key string) (string, error) {
	v, ok := c.cache.Get(key)
	if !ok {
		return "", errors.ErrNotfound
	}
	return v.(string), nil
}

func (c *inmemoryCache) Set(ctx context.Context, key string, value string) error {
	return c.SetTTL(ctx, key, value, 0)
}

func (c *inmemoryCache) SetTTL(ctx context.Context, key string, value string, exp time.Duration) error {
	c.cache.Set(key, value, exp)
	return nil
}

func (c *inmemoryCache) Del(ctx context.Context, key string) error {
	c.cache.Delete(key)
	return nil
}
