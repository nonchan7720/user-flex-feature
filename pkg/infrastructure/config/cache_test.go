package config

import (
	"testing"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/stretchr/testify/require"
)

func TestCacheConfigInValidRedis(t *testing.T) {
	cfg := Cache{
		Type: CacheRedis,
	}
	err := cfg.Validate()
	var validErr validation.Errors
	require.ErrorAs(t, err, &validErr)
	_, ok := validErr["redis"]
	require.True(t, ok)
}

func TestCacheConfigInValidInMemory(t *testing.T) {
	cfg := Cache{
		Type: CacheInMemory,
	}
	err := cfg.Validate()
	var validErr validation.Errors
	require.ErrorAs(t, err, &validErr)
	_, ok := validErr["inmemory"]
	require.True(t, ok)
}
