package cache

import (
	"context"
	"time"

	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/storage/redis"
)

type redisCache struct {
	cmd redis.Cmdable
}

func NewRedisCache(cmd redis.Cmdable) Backend {
	return &redisCache{
		cmd: cmd,
	}
}

func (c *redisCache) Get(ctx context.Context, key string) (string, error) {
	return c.cmd.Get(ctx, key).Result()
}

func (c *redisCache) Set(ctx context.Context, key string, value string) error {
	return c.SetTTL(ctx, key, value, 0)
}

func (c *redisCache) SetTTL(ctx context.Context, key string, value string, exp time.Duration) error {
	return c.cmd.Set(ctx, key, value, exp).Err()
}

func (c *redisCache) Del(ctx context.Context, key string) error {
	return c.cmd.Del(ctx, key).Err()
}
