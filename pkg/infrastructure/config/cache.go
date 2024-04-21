package config

import (
	validation "github.com/go-ozzo/ozzo-validation/v4"
	. "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/inmemory"
	. "github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/redis"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/validator"
)

type CacheType string

const (
	CacheInMemory = CacheType("inmemory")
	CacheRedis    = CacheType("redis")
)

type Cache struct {
	Type     CacheType `yaml:"type" default:"inmemory"`
	Redis    *Redis    `yaml:"redis"`
	InMemory *InMemory `yaml:"inmemory"`
}

func (c Cache) Validate() error {
	return validator.ValidateStruct(&c,
		validation.Field(&c.Type, validation.Required, validation.In(CacheInMemory, CacheRedis)),
		validation.Field(c.Redis, validation.When(c.Type == CacheRedis, validation.Required)),
		validation.Field(c.InMemory, validation.When(c.Type == CacheInMemory, validation.Required)),
	)
}
