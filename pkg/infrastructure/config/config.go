package config

import (
	"strings"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/nonchan7720/user-flex-feature/pkg/container"
	"github.com/nonchan7720/user-flex-feature/pkg/infrastructure/config/retriever"
	"github.com/samber/do"
)

func init() {
	do.Provide(container.Injector, func(_ *do.Injector) (Config, error) { return Config{}, nil })
}

type Config struct {
	AppEnv          string                `yaml:"appEnv" default:"stg"`
	Retrievers      []retriever.Retriever `yaml:"retrievers"`
	PollingInterval time.Duration         `yaml:"pollingInterval" default:"1m"`
	Gateway         *Gateway              `yaml:"gateway"`
	Grpc            Grpc                  `yaml:"grpc,alias"`

	Tracking Tracking  `yaml:"tracking"`
	Logging  Logging   `yaml:"logging"`
	CSRF     CSRFToken `yaml:"csrf"`
}

func (c Config) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.Logging),
		validation.Field(&c.Tracking),
		validation.Field(&c.CSRF),
	)
}

func (c Config) IsProduction() bool {
	for _, name := range []string{"prod", "production"} {
		if strings.EqualFold(c.AppEnv, name) {
			return true
		}
	}
	return false
}

func (c Config) IsStaging() bool {
	for _, name := range []string{"stg", "staging"} {
		if strings.EqualFold(c.AppEnv, name) {
			return true
		}
	}
	return false
}

func (c Config) IsLocal() bool {
	for _, name := range []string{"local"} {
		if strings.EqualFold(c.AppEnv, name) {
			return true
		}
	}
	return false
}

func (c Config) OTEL() Tracking {
	return c.Tracking
}

func (c *Config) SetupLog() {
	c.Logging.SetupLog(c.Tracking)
}

var (
	_ do.Shutdownable = &Config{}
)

func (c *Config) Shutdown() error {
	c.Logging.Close()
	return nil
}
