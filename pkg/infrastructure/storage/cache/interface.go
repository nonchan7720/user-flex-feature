package cache

import (
	"context"
	"time"
)

type Cache[V any] interface {
	Get(ctx context.Context, key string) (*V, error)
	Set(ctx context.Context, key string, value *V) error
	SetTTL(ctx context.Context, key string, value *V, exp time.Duration) error
	Del(ctx context.Context, key string) error
}

type Backend interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value string) error
	SetTTL(ctx context.Context, key string, value string, exp time.Duration) error
	Del(ctx context.Context, key string) error
}
